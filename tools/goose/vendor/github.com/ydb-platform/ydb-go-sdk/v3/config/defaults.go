package config

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"google.golang.org/grpc"
	grpcCredentials "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/ydb-platform/ydb-go-sdk/v3/balancers"
	"github.com/ydb-platform/ydb-go-sdk/v3/credentials"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

var (
	// DefaultKeepaliveInterval contains default duration between grpc keepalive
	DefaultKeepaliveInterval    = 10 * time.Second
	MinKeepaliveInterval        = 10 * time.Second
	DefaultDialTimeout          = 5 * time.Second
	DefaultGRPCMsgSize          = 64 * 1024 * 1024 // 64MB
	DefaultGrpcConnectionPolicy = keepalive.ClientParameters{
		Time:                DefaultKeepaliveInterval,
		Timeout:             MinKeepaliveInterval,
		PermitWithoutStream: true,
	}
)

func defaultGrpcOptions(secure bool, tlsConfig *tls.Config) (opts []grpc.DialOption) {
	opts = append(opts,
		// keep-aliving all connections
		grpc.WithKeepaliveParams(
			DefaultGrpcConnectionPolicy,
		),
		// use round robin balancing policy for fastest dialing
		grpc.WithDefaultServiceConfig(`{
			"loadBalancingPolicy": "round_robin"
		}`),
		// limit size of outgoing and incoming packages
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(DefaultGRPCMsgSize),
			grpc.MaxCallSendMsgSize(DefaultGRPCMsgSize),
		),
	)
	if secure {
		opts = append(opts, grpc.WithTransportCredentials(
			grpcCredentials.NewTLS(tlsConfig),
		))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
	}

	return opts
}

func certPool() *x509.CertPool {
	certPool, err := x509.SystemCertPool()
	if err == nil {
		return certPool
	}

	return x509.NewCertPool()
}

func defaultTLSConfig() *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool(),
	}
}

func defaultConfig() (c *Config) {
	return &Config{
		credentials: credentials.NewAnonymousCredentials(
			credentials.WithSourceInfo(stack.Record(0)),
		),
		balancerConfig: balancers.Default(),
		tlsConfig:      defaultTLSConfig(),
		dialTimeout:    DefaultDialTimeout,
		trace:          &trace.Driver{},
	}
}