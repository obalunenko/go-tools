package balancer

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/ydb-platform/ydb-go-genproto/Ydb_Discovery_V1"
	"google.golang.org/grpc"

	"github.com/ydb-platform/ydb-go-sdk/v3/config"
	balancerConfig "github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer/config"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/conn"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/credentials"
	internalDiscovery "github.com/ydb-platform/ydb-go-sdk/v3/internal/discovery"
	discoveryConfig "github.com/ydb-platform/ydb-go-sdk/v3/internal/discovery/config"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/endpoint"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/meta"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/repeater"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xcontext"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xresolver"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xslices"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

var ErrNoEndpoints = xerrors.Wrap(fmt.Errorf("no endpoints"))

type Balancer struct {
	driverConfig      *config.Config
	balancerConfig    balancerConfig.Config
	discoveryConfig   *discoveryConfig.Config
	pool              *conn.Pool
	discoveryRepeater repeater.Repeater

	discover        func(ctx context.Context) (endpoints []endpoint.Endpoint, location string, err error)
	localDCDetector func(ctx context.Context, endpoints []endpoint.Endpoint) (string, error)

	connectionsState atomic.Pointer[connectionsState]
}

func (b *Balancer) clusterDiscovery(ctx context.Context) (err error) {
	return retry.Retry(
		repeater.WithEvent(ctx, repeater.EventInit),
		func(childCtx context.Context) (err error) {
			if err = b.clusterDiscoveryAttempt(childCtx); err != nil {
				if credentials.IsAccessError(err) {
					return credentials.AccessError("cluster discovery failed", err,
						credentials.WithEndpoint(b.driverConfig.Endpoint()),
						credentials.WithDatabase(b.driverConfig.Database()),
						credentials.WithCredentials(b.driverConfig.Credentials()),
					)
				}
				// if got err but parent context is not done - mark error as retryable
				if ctx.Err() == nil && xerrors.IsTimeoutError(err) {
					return xerrors.WithStackTrace(xerrors.Retryable(err))
				}

				return xerrors.WithStackTrace(err)
			}

			return nil
		},
		retry.WithIdempotent(true),
		retry.WithTrace(b.driverConfig.TraceRetry()),
		retry.WithBudget(b.driverConfig.RetryBudget()),
	)
}

func (b *Balancer) clusterDiscoveryAttempt(ctx context.Context) (err error) {
	var (
		address = "ydb:///" + b.driverConfig.Endpoint()
		onDone  = trace.DriverOnBalancerClusterDiscoveryAttempt(
			b.driverConfig.Trace(), &ctx,
			stack.FunctionID(
				"github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer.(*Balancer).clusterDiscoveryAttempt",
			),
			address,
			b.driverConfig.Database(),
		)
	)
	defer func() {
		onDone(err)
	}()

	if dialTimeout := b.driverConfig.DialTimeout(); dialTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = xcontext.WithTimeout(ctx, dialTimeout)
		defer cancel()
	}

	endpoints, location, err := b.discover(ctx)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	if b.balancerConfig.DetectNearestDC {
		location, err := b.localDCDetector(ctx, endpoints)
		if err != nil {
			return xerrors.WithStackTrace(err)
		}

		b.applyDiscoveredEndpoints(ctx, endpoints, location)
	} else {
		b.applyDiscoveredEndpoints(ctx, endpoints, location)
	}

	return nil
}

func (b *Balancer) applyDiscoveredEndpoints(ctx context.Context, newest []endpoint.Endpoint, localDC string) {
	var (
		onDone = trace.DriverOnBalancerUpdate(
			b.driverConfig.Trace(), &ctx,
			stack.FunctionID(
				"github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer.(*Balancer).applyDiscoveredEndpoints"),
			b.balancerConfig.DetectNearestDC,
			b.driverConfig.Database(),
		)
		previous = b.connections().All()
	)
	defer func() {
		_, added, dropped := xslices.Diff(previous, newest, func(lhs, rhs endpoint.Endpoint) int {
			return strings.Compare(lhs.Address(), rhs.Address())
		})
		onDone(
			xslices.Transform(newest, func(t endpoint.Endpoint) trace.EndpointInfo { return t }),
			xslices.Transform(added, func(t endpoint.Endpoint) trace.EndpointInfo { return t }),
			xslices.Transform(dropped, func(t endpoint.Endpoint) trace.EndpointInfo { return t }),
			localDC,
		)
	}()

	connections := endpointsToConnections(b.pool, newest)
	for _, c := range connections {
		b.pool.Allow(ctx, c)
		c.Endpoint().Touch()
	}

	info := balancerConfig.Info{SelfLocation: localDC}
	state := newConnectionsState(connections, b.balancerConfig.Filter, info, b.balancerConfig.AllowFallback)

	endpointsInfo := make([]endpoint.Info, len(newest))
	for i, e := range newest {
		endpointsInfo[i] = e
	}

	b.connectionsState.Store(state)
}

func (b *Balancer) Close(ctx context.Context) (err error) {
	onDone := trace.DriverOnBalancerClose(
		b.driverConfig.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer.(*Balancer).Close"),
	)
	defer func() {
		onDone(err)
	}()

	if b.discoveryRepeater != nil {
		b.discoveryRepeater.Stop()
	}

	return nil
}

func makeDiscoveryFunc(
	driverConfig *config.Config, discoveryConfig *discoveryConfig.Config,
) func(ctx context.Context) (endpoints []endpoint.Endpoint, location string, err error) {
	return func(ctx context.Context) (endpoints []endpoint.Endpoint, location string, err error) {
		ctx, traceID, err := meta.TraceID(ctx)
		if err != nil {
			return endpoints, location, xerrors.WithStackTrace(err)
		}

		ctx, err = driverConfig.Meta().Context(ctx)
		if err != nil {
			return endpoints, location, xerrors.WithStackTrace(
				fmt.Errorf("failed to enrich context with meta, traceID %q: %w", traceID, err),
			)
		}

		cc, err := grpc.DialContext(ctx,
			"ydb:///"+driverConfig.Endpoint(),
			append(
				driverConfig.GrpcDialOptions(),
				grpc.WithResolvers(
					xresolver.New("ydb", driverConfig.Trace()),
				),
				grpc.WithBlock(),
				grpc.WithDefaultServiceConfig(`{
					"loadBalancingPolicy": "pick_first"
				}`),
			)...,
		)
		if err != nil {
			return endpoints, location, xerrors.WithStackTrace(
				fmt.Errorf("failed to dial %q, traceID %q: %w", driverConfig.Endpoint(), traceID, err),
			)
		}
		defer func() {
			_ = cc.Close()
		}()

		endpoints, location, err = internalDiscovery.Discover(ctx,
			Ydb_Discovery_V1.NewDiscoveryServiceClient(cc), discoveryConfig,
		)
		if err != nil {
			return endpoints, location, xerrors.WithStackTrace(
				fmt.Errorf("failed to discover database %q, address %q, traceID %q: %w",
					driverConfig.Database(), driverConfig.Endpoint(), traceID, err,
				),
			)
		}

		return endpoints, location, nil
	}
}

func New(ctx context.Context, driverConfig *config.Config, pool *conn.Pool, opts ...discoveryConfig.Option) (
	b *Balancer, finalErr error,
) {
	onDone := trace.DriverOnBalancerInit(driverConfig.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer.New"),
		driverConfig.Balancer().String(),
	)
	defer func() {
		onDone(finalErr)
	}()

	b = &Balancer{
		driverConfig: driverConfig,
		pool:         pool,
		discoveryConfig: discoveryConfig.New(append(opts,
			discoveryConfig.With(driverConfig.Common),
			discoveryConfig.WithEndpoint(driverConfig.Endpoint()),
			discoveryConfig.WithDatabase(driverConfig.Database()),
			discoveryConfig.WithSecure(driverConfig.Secure()),
			discoveryConfig.WithMeta(driverConfig.Meta()),
		)...),
		localDCDetector: detectLocalDC,
	}

	b.discover = makeDiscoveryFunc(b.driverConfig, b.discoveryConfig)

	if config := driverConfig.Balancer(); config == nil {
		b.balancerConfig = balancerConfig.Config{}
	} else {
		b.balancerConfig = *config
	}

	if b.balancerConfig.SingleConn {
		b.applyDiscoveredEndpoints(ctx, []endpoint.Endpoint{
			endpoint.New(driverConfig.Endpoint()),
		}, "")
	} else {
		// initialization of balancer state
		if err := b.clusterDiscovery(ctx); err != nil {
			return nil, xerrors.WithStackTrace(err)
		}
		// run background discovering
		if d := b.discoveryConfig.Interval(); d > 0 {
			b.discoveryRepeater = repeater.New(xcontext.ValueOnly(ctx),
				d, b.clusterDiscoveryAttempt,
				repeater.WithName("discovery"),
				repeater.WithTrace(b.driverConfig.Trace()),
			)
		}
	}

	return b, nil
}

func (b *Balancer) Invoke(
	ctx context.Context,
	method string,
	args interface{},
	reply interface{},
	opts ...grpc.CallOption,
) error {
	return b.wrapCall(ctx, func(ctx context.Context, cc conn.Conn) error {
		return cc.Invoke(ctx, method, args, reply, opts...)
	})
}

func (b *Balancer) NewStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	method string,
	opts ...grpc.CallOption,
) (_ grpc.ClientStream, err error) {
	var client grpc.ClientStream
	err = b.wrapCall(ctx, func(ctx context.Context, cc conn.Conn) error {
		client, err = cc.NewStream(ctx, desc, method, opts...)

		return err
	})
	if err == nil {
		return client, nil
	}

	return nil, err
}

func (b *Balancer) wrapCall(ctx context.Context, f func(ctx context.Context, cc conn.Conn) error) (err error) {
	cc, err := b.getConn(ctx)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	defer func() {
		if err == nil {
			if cc.GetState() == conn.Banned {
				b.pool.Allow(ctx, cc)
			}
		} else if conn.IsBadConn(err, b.driverConfig.ExcludeGRPCCodesForPessimization()...) {
			b.pool.Ban(ctx, cc, err)
		}
	}()

	if err = f(ctx, cc); err != nil {
		if conn.UseWrapping(ctx) {
			if credentials.IsAccessError(err) {
				err = credentials.AccessError("no access", err,
					credentials.WithAddress(cc.Endpoint().String()),
					credentials.WithNodeID(cc.Endpoint().NodeID()),
					credentials.WithCredentials(b.driverConfig.Credentials()),
				)
			}

			return xerrors.WithStackTrace(err)
		}

		return err
	}

	return nil
}

func (b *Balancer) connections() *connectionsState {
	return b.connectionsState.Load()
}

func (b *Balancer) getConn(ctx context.Context) (c conn.Conn, err error) {
	onDone := trace.DriverOnBalancerChooseEndpoint(
		b.driverConfig.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer.(*Balancer).getConn"),
	)
	defer func() {
		if err == nil {
			onDone(c.Endpoint(), nil)
		} else {
			onDone(nil, err)
		}
	}()

	if err = ctx.Err(); err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	var (
		state       = b.connections()
		failedCount int
	)

	defer func() {
		if failedCount*2 > state.PreferredCount() && b.discoveryRepeater != nil {
			b.discoveryRepeater.Force()
		}
	}()

	c, failedCount = state.GetConnection(ctx)
	if c == nil {
		return nil, xerrors.WithStackTrace(
			fmt.Errorf("%w: cannot get connection from Balancer after %d attempts", ErrNoEndpoints, failedCount),
		)
	}

	return c, nil
}

func endpointsToConnections(p *conn.Pool, endpoints []endpoint.Endpoint) []conn.Conn {
	conns := make([]conn.Conn, 0, len(endpoints))
	for _, e := range endpoints {
		conns = append(conns, p.Get(e))
	}

	return conns
}
