package table

import (
	"context"

	"github.com/jonboulle/clockwork"
	"github.com/ydb-platform/ydb-go-genproto/Ydb_Table_V1"
	"google.golang.org/grpc"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/allocator"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/pool"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/table/config"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xcontext"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

// sessionBuilder is the interface that holds logic of creating sessions.
type sessionBuilder func(ctx context.Context) (*session, error)

func New(ctx context.Context, cc grpc.ClientConnInterface, config *config.Config) *Client {
	onDone := trace.TableOnInit(config.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table.New"),
	)

	return &Client{
		clock:  config.Clock(),
		config: config,
		cc:     cc,
		build: func(ctx context.Context) (s *session, err error) {
			return newSession(ctx, cc, config)
		},
		pool: pool.New[*session, session](ctx,
			pool.WithLimit[*session, session](config.SizeLimit()),
			pool.WithItemUsageLimit[*session, session](config.SessionUsageLimit()),
			pool.WithIdleTimeToLive[*session, session](config.IdleThreshold()),
			pool.WithCreateItemTimeout[*session, session](config.CreateSessionTimeout()),
			pool.WithCloseItemTimeout[*session, session](config.DeleteTimeout()),
			pool.WithClock[*session, session](config.Clock()),
			pool.WithCreateItemFunc[*session, session](func(ctx context.Context) (*session, error) {
				return newSession(ctx, cc, config)
			}),
			pool.WithTrace[*session, session](&pool.Trace{
				OnNew: func(ctx *context.Context, call stack.Caller) func(limit int) {
					return func(limit int) {
						onDone(limit)
					}
				},
				OnPut: func(ctx *context.Context, call stack.Caller, item any) func(err error) {
					onDone := trace.TableOnPoolPut( //nolint:forcetypeassert
						config.Trace(), ctx, call, item.(*session),
					)

					return func(err error) {
						onDone(err)
					}
				},
				OnGet: func(ctx *context.Context, call stack.Caller) func(item any, attempts int, err error) {
					onDone := trace.TableOnPoolGet(config.Trace(), ctx, call)

					return func(item any, attempts int, err error) {
						onDone(item.(*session), attempts, err) //nolint:forcetypeassert
					}
				},
				OnWith: func(ctx *context.Context, call stack.Caller) func(attempts int, err error) {
					onDone := trace.TableOnPoolWith(config.Trace(), ctx, call)

					return func(attempts int, err error) {
						onDone(attempts, err)
					}
				},
				OnChange: func(stats pool.Stats) {
					trace.TableOnPoolStateChange(config.Trace(),
						stats.Limit, stats.Index, stats.Idle, stats.Wait, stats.CreateInProgress, stats.Index,
					)
				},
			}),
		),
		done: make(chan struct{}),
	}
}

// Client is a set of session instances that may be reused.
// A Client is safe for use by multiple goroutines simultaneously.
type Client struct {
	// read-only fields
	config *config.Config
	build  sessionBuilder
	cc     grpc.ClientConnInterface
	clock  clockwork.Clock
	pool   sessionPool
	done   chan struct{}
}

func (c *Client) CreateSession(ctx context.Context, opts ...table.Option) (_ table.ClosableSession, err error) {
	if c == nil {
		return nil, xerrors.WithStackTrace(errNilClient)
	}
	if c.isClosed() {
		return nil, xerrors.WithStackTrace(errClosedClient)
	}
	createSession := func(ctx context.Context) (*session, error) {
		s, err := c.build(ctx)
		if err != nil {
			return nil, xerrors.WithStackTrace(err)
		}

		return s, nil
	}
	if !c.config.AutoRetry() {
		s, err := createSession(ctx)
		if err != nil {
			return nil, xerrors.WithStackTrace(err)
		}

		return s, nil
	}

	var (
		onDone = trace.TableOnCreateSession(c.config.Trace(), &ctx,
			stack.FunctionID(
				"github.com/ydb-platform/ydb-go-sdk/v3/internal/table.(*Client).CreateSession"),
		)
		attempts = 0
		s        *session
	)
	defer func() {
		if s != nil {
			onDone(s, attempts, err)
		} else {
			onDone(nil, attempts, err)
		}
	}()

	s, err = retry.RetryWithResult(ctx, createSession,
		append(
			[]retry.Option{
				retry.WithIdempotent(true),
				retry.WithTrace(&trace.Retry{
					OnRetry: func(info trace.RetryLoopStartInfo) func(trace.RetryLoopDoneInfo) {
						return func(info trace.RetryLoopDoneInfo) {
							attempts = info.Attempts
						}
					},
				}),
			}, c.retryOptions(opts...).RetryOptions...,
		)...,
	)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return s, nil
}

func (c *Client) isClosed() bool {
	select {
	case <-c.done:
		return true
	default:
		return false
	}
}

// Close deletes all stored sessions inside Client.
// It also stops all underlying timers and goroutines.
// It returns first error occurred during stale sessions' deletion.
// Note that even on error it calls Close() on each session.
func (c *Client) Close(ctx context.Context) (err error) {
	if c == nil {
		return xerrors.WithStackTrace(errNilClient)
	}

	close(c.done)

	onDone := trace.TableOnClose(c.config.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table.(*Client).Close"),
	)
	defer func() {
		onDone(err)
	}()

	return c.pool.Close(ctx)
}

// Do provide the best effort for execute operation
// Do implements internal busy loop until one of the following conditions is met:
// - deadline was canceled or deadlined
// - retry operation returned nil as error
// Warning: if deadline without deadline or cancellation func Retry will be worked infinite
func (c *Client) Do(ctx context.Context, op table.Operation, opts ...table.Option) (finalErr error) {
	if c == nil {
		return xerrors.WithStackTrace(errNilClient)
	}

	if c.isClosed() {
		return xerrors.WithStackTrace(errClosedClient)
	}

	config := c.retryOptions(opts...)

	attempts, onDone := 0, trace.TableOnDo(config.Trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table.(*Client).Do"),
		config.Label, config.Idempotent, xcontext.IsNestedCall(ctx),
	)
	defer func() {
		onDone(attempts, finalErr)
	}()

	err := do(ctx, c.pool, c.config, op, func(err error) {
		attempts++
	}, config.RetryOptions...)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (c *Client) DoTx(ctx context.Context, op table.TxOperation, opts ...table.Option) (finalErr error) {
	if c == nil {
		return xerrors.WithStackTrace(errNilClient)
	}

	if c.isClosed() {
		return xerrors.WithStackTrace(errClosedClient)
	}

	config := c.retryOptions(opts...)

	attempts, onDone := 0, trace.TableOnDoTx(config.Trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table.(*Client).DoTx"),
		config.Label, config.Idempotent, xcontext.IsNestedCall(ctx),
	)
	defer func() {
		onDone(attempts, finalErr)
	}()

	return retryBackoff(ctx, c.pool, func(ctx context.Context, s table.Session) (err error) {
		attempts++

		tx, err := s.BeginTransaction(ctx, config.TxSettings)
		if err != nil {
			return xerrors.WithStackTrace(err)
		}

		defer func() {
			if err != nil && !xerrors.IsOperationError(err) {
				_ = tx.Rollback(ctx)
			}
		}()

		if err = executeTxOperation(ctx, c, op, tx); err != nil {
			return xerrors.WithStackTrace(err)
		}

		_, err = tx.CommitTx(ctx, config.TxCommitOptions...)
		if err != nil {
			return xerrors.WithStackTrace(err)
		}

		return nil
	}, config.RetryOptions...)
}

func (c *Client) BulkUpsert(
	ctx context.Context,
	tableName string,
	data table.BulkUpsertData,
	opts ...table.Option,
) (finalErr error) {
	if c == nil {
		return xerrors.WithStackTrace(errNilClient)
	}

	if c.isClosed() {
		return xerrors.WithStackTrace(errClosedClient)
	}

	a := allocator.New()
	defer a.Free()

	attempts, config := 0, c.retryOptions(opts...)
	config.RetryOptions = append(config.RetryOptions,
		retry.WithIdempotent(true),
		retry.WithTrace(&trace.Retry{
			OnRetry: func(info trace.RetryLoopStartInfo) func(trace.RetryLoopDoneInfo) {
				return func(info trace.RetryLoopDoneInfo) {
					attempts = info.Attempts
				}
			},
		}),
	)

	onDone := trace.TableOnBulkUpsert(config.Trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table.(*Client).BulkUpsert"),
	)
	defer func() {
		onDone(finalErr, attempts)
	}()

	request, err := data.ToYDB(a, tableName)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	client := Ydb_Table_V1.NewTableServiceClient(c.cc)

	err = retry.Retry(ctx,
		func(ctx context.Context) (err error) {
			attempts++
			_, err = client.BulkUpsert(ctx, request)

			return err
		},
		config.RetryOptions...,
	)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func executeTxOperation(ctx context.Context, c *Client, op table.TxOperation, tx table.Transaction) (err error) {
	if panicCallback := c.config.PanicCallback(); panicCallback != nil {
		defer func() {
			if e := recover(); e != nil {
				panicCallback(e)
			}
		}()
	}

	return op(xcontext.MarkRetryCall(ctx), tx)
}
