package query

import (
	"context"

	"github.com/ydb-platform/ydb-go-genproto/Ydb_Query_V1"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/session"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	baseTx "github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

var _ query.Session = (*Session)(nil)

type (
	Session struct {
		session.Core

		client Ydb_Query_V1.QueryServiceClient
		trace  *trace.Query
		laztTx bool
	}
)

func (s *Session) QueryResultSet(
	ctx context.Context, q string, opts ...options.Execute,
) (rs result.ClosableResultSet, finalErr error) {
	onDone := trace.QueryOnSessionQueryResultSet(s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Session).QueryResultSet"), s, q)
	defer func() {
		onDone(finalErr)
	}()

	r, err := execute(ctx, s.ID(), s.client, q, options.ExecuteSettings(opts...), withTrace(s.trace))
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	rs, err = readResultSet(ctx, r)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return rs, nil
}

func (s *Session) queryRow(
	ctx context.Context, q string, settings executeSettings, resultOpts ...resultOption,
) (row query.Row, finalErr error) {
	r, err := execute(ctx, s.ID(), s.client, q, settings, resultOpts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	row, err = readRow(ctx, r)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return row, nil
}

func (s *Session) QueryRow(ctx context.Context, q string, opts ...options.Execute) (_ query.Row, finalErr error) {
	onDone := trace.QueryOnSessionQueryRow(s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Session).QueryRow"), s, q)
	defer func() {
		onDone(finalErr)
	}()

	row, err := s.queryRow(ctx, q, options.ExecuteSettings(opts...), withTrace(s.trace))
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return row, nil
}

func createSession(
	ctx context.Context, client Ydb_Query_V1.QueryServiceClient, opts ...session.Option,
) (*Session, error) {
	core, err := session.Open(ctx, client, opts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return &Session{
		Core:   core,
		trace:  core.Trace,
		client: core.Client,
	}, nil
}

func (s *Session) Begin(
	ctx context.Context,
	txSettings query.TransactionSettings,
) (
	tx query.Transaction, finalErr error,
) {
	onDone := trace.QueryOnSessionBegin(s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Session).Begin"), s)
	defer func() {
		if finalErr != nil {
			onDone(finalErr, nil)
		} else {
			onDone(nil, tx)
		}
	}()

	if s.laztTx {
		return &Transaction{
			s:          s,
			txSettings: txSettings,
		}, nil
	}

	txID, err := begin(ctx, s.client, s.ID(), txSettings)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return &Transaction{
		LazyID: baseTx.ID(txID),
		s:      s,
	}, nil
}

func (s *Session) Exec(
	ctx context.Context, q string, opts ...options.Execute,
) (finalErr error) {
	onDone := trace.QueryOnSessionExec(s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Session).Exec"), s, q)
	defer func() {
		onDone(finalErr)
	}()

	r, err := execute(ctx, s.ID(), s.client, q, options.ExecuteSettings(opts...), withTrace(s.trace))
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	err = readAll(ctx, r)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (s *Session) Query(
	ctx context.Context, q string, opts ...options.Execute,
) (_ query.Result, finalErr error) {
	onDone := trace.QueryOnSessionQuery(s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Session).Query"), s, q)
	defer func() {
		onDone(finalErr)
	}()

	r, err := execute(ctx, s.ID(), s.client, q, options.ExecuteSettings(opts...), withTrace(s.trace))
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return r, nil
}