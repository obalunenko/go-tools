package legacy

import (
	"context"
	"database/sql/driver"
	"fmt"
	"io"
	"slices"
	"sync/atomic"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/params"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xcontext"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql/iface"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql/legacy/badconn"
	"github.com/ydb-platform/ydb-go-sdk/v3/scripting"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
)

type (
	Parent interface {
		Table() table.Client
		Scripting() scripting.Client
	}
	Conn struct {
		ctx context.Context //nolint:containedctx

		parent  Parent
		session table.ClosableSession // Immutable and r/o usage.

		fakeTxModes []QueryMode

		closed           atomic.Bool
		defaultQueryMode QueryMode

		defaultTxControl *table.TransactionControl
		dataOpts         []options.ExecuteDataQueryOption

		scanOpts []options.ExecuteScanQueryOption

		idleThreshold time.Duration
		onClose       []func()
	}
)

func (c *Conn) Exec(ctx context.Context, sql string, params *params.Params) (result driver.Result, err error) {
	if !c.isReady() {
		return nil, badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	m := queryModeFromContext(ctx, c.defaultQueryMode)

	switch m {
	case DataQueryMode:
		return c.executeDataQuery(ctx, sql, params)
	case SchemeQueryMode:
		return c.executeSchemeQuery(ctx, sql)
	case ScriptingQueryMode:
		return c.executeScriptingQuery(ctx, sql, params)
	default:
		return nil, fmt.Errorf("unsupported query mode '%s' for execute query", m)
	}
}

func (c *Conn) Query(ctx context.Context, sql string, params *params.Params) (
	result driver.RowsNextResultSet, finalErr error,
) {
	if !c.isReady() {
		return nil, badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	switch queryMode := queryModeFromContext(ctx, c.defaultQueryMode); queryMode {
	case DataQueryMode:
		return c.execDataQuery(ctx, sql, params)
	case ScanQueryMode:
		return c.execScanQuery(ctx, sql, params)
	case ScriptingQueryMode:
		return c.execScriptingQuery(ctx, sql, params)
	default:
		return nil, fmt.Errorf("unsupported query mode '%s' on iface query", queryMode)
	}
}

func (c *Conn) Explain(ctx context.Context, sql string, _ *params.Params) (ast string, plan string, err error) {
	exp, err := c.session.Explain(ctx, sql)
	if err != nil {
		return "", "", badconn.Map(xerrors.WithStackTrace(err))
	}

	return exp.AST, exp.Plan, nil
}

func (c *Conn) CheckNamedValue(*driver.NamedValue) error {
	// on this stage allows all values
	return nil
}

func (c *Conn) IsValid() bool {
	return c.isReady()
}

type resultNoRows struct{}

func (resultNoRows) LastInsertId() (int64, error) { return 0, ErrUnsupported }
func (resultNoRows) RowsAffected() (int64, error) { return 0, ErrUnsupported }

func New(ctx context.Context, parent Parent, s table.ClosableSession, opts ...Option) *Conn {
	cc := &Conn{
		ctx:              ctx,
		parent:           parent,
		session:          s,
		defaultQueryMode: DataQueryMode,
		defaultTxControl: table.DefaultTxControl(),
	}

	for _, opt := range opts {
		if opt != nil {
			opt(cc)
		}
	}

	return cc
}

func (c *Conn) isReady() bool {
	return c.session.Status() == table.SessionReady
}

func (c *Conn) executeDataQuery(ctx context.Context, sql string, params *params.Params) (driver.Result, error) {
	_, res, err := c.session.Execute(ctx,
		txControl(ctx, c.defaultTxControl),
		sql, params, c.dataQueryOptions(ctx)...,
	)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	defer res.Close()

	if err := res.NextResultSetErr(ctx); err != nil && !xerrors.Is(err, nil, io.EOF) {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	if err := res.Err(); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return resultNoRows{}, nil
}

func (c *Conn) executeSchemeQuery(ctx context.Context, sql string) (driver.Result, error) {
	if err := c.session.ExecuteSchemeQuery(ctx, sql); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return resultNoRows{}, nil
}

func (c *Conn) executeScriptingQuery(ctx context.Context, sql string, params *params.Params) (
	driver.Result, error,
) {
	res, err := c.parent.Scripting().StreamExecute(ctx, sql, params)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	defer res.Close()

	if err := res.NextResultSetErr(ctx); err != nil && !xerrors.Is(err, nil, io.EOF) {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	if err := res.Err(); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return resultNoRows{}, nil
}

func (c *Conn) execDataQuery(ctx context.Context, sql string, params *params.Params) (
	driver.RowsNextResultSet, error,
) {
	_, res, err := c.session.Execute(ctx,
		txControl(ctx, c.defaultTxControl),
		sql, params, c.dataQueryOptions(ctx)...,
	)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	if err = res.Err(); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return &rows{
		conn:   c,
		result: res,
	}, nil
}

func (c *Conn) execScanQuery(ctx context.Context, sql string, params *params.Params) (
	driver.RowsNextResultSet, error,
) {
	res, err := c.session.StreamExecuteScanQuery(ctx,
		sql, params, c.scanOpts...,
	)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	if err = res.Err(); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return &rows{
		conn:   c,
		result: res,
	}, nil
}

func (c *Conn) execScriptingQuery(ctx context.Context, sql string, params *params.Params) (
	driver.RowsNextResultSet, error,
) {
	res, err := c.parent.Scripting().StreamExecute(ctx, sql, params)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	if err = res.Err(); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return &rows{
		conn:   c,
		result: res,
	}, nil
}

func (c *Conn) Ping(ctx context.Context) (finalErr error) {
	if !c.isReady() {
		return badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}
	if err := c.session.KeepAlive(ctx); err != nil {
		return badconn.Map(xerrors.WithStackTrace(err))
	}

	return nil
}

func (c *Conn) Close() (finalErr error) {
	if !c.closed.CompareAndSwap(false, true) {
		return badconn.Map(xerrors.WithStackTrace(errConnClosedEarly))
	}

	defer func() {
		for _, onClose := range c.onClose {
			onClose()
		}
	}()

	err := c.session.Close(xcontext.ValueOnly(c.ctx))
	if err != nil {
		return badconn.Map(xerrors.WithStackTrace(err))
	}

	return nil
}

func (c *Conn) ID() string {
	return c.session.ID()
}

func (c *Conn) beginTx(ctx context.Context, txOptions driver.TxOptions) (tx iface.Tx, finalErr error) {
	m := queryModeFromContext(ctx, c.defaultQueryMode)

	if slices.Contains(c.fakeTxModes, m) {
		return beginTxFake(ctx, c), nil
	}

	tx, err := beginTx(ctx, c, txOptions)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return tx, nil
}

func (c *Conn) BeginTx(ctx context.Context, txOptions driver.TxOptions) (iface.Tx, error) {
	tx, err := c.beginTx(ctx, txOptions)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return tx, nil
}
