package conn

import (
	"context"
	"database/sql/driver"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn/badconn"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type txFake struct {
	tx.Identifier

	beginCtx context.Context //nolint:containedctx
	conn     *Conn
	ctx      context.Context //nolint:containedctx
}

func (tx *txFake) PrepareContext(ctx context.Context, query string) (_ driver.Stmt, finalErr error) {
	onDone := trace.DatabaseSQLOnTxPrepare(tx.conn.parent.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*txFake).PrepareContext"),
		tx.beginCtx, tx, query,
	)
	defer func() {
		onDone(finalErr)
	}()
	if !tx.conn.isReady() {
		return nil, badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	return &stmt{
		conn:      tx.conn,
		processor: tx,
		ctx:       ctx,
		query:     query,
	}, nil
}

var (
	_ driver.Tx             = &txFake{}
	_ driver.ExecerContext  = &txFake{}
	_ driver.QueryerContext = &txFake{}
	_ tx.Identifier         = &txFake{}
)

func beginTxFake(ctx context.Context, c *Conn) currentTx {
	return &txFake{
		Identifier: tx.ID("FAKE"),
		conn:       c,
		ctx:        ctx,
	}
}

func (tx *txFake) Commit() (err error) {
	var (
		ctx    = tx.ctx
		onDone = trace.DatabaseSQLOnTxCommit(tx.conn.parent.Trace(), &ctx,
			stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*txFake).Commit"),
			tx,
		)
	)
	defer func() {
		onDone(err)
	}()
	defer func() {
		tx.conn.currentTx = nil
	}()
	if !tx.conn.isReady() {
		return badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	return nil
}

func (tx *txFake) Rollback() (err error) {
	var (
		ctx    = tx.ctx
		onDone = trace.DatabaseSQLOnTxRollback(tx.conn.parent.Trace(), &ctx,
			stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*txFake).Rollback"),
			tx,
		)
	)
	defer func() {
		onDone(err)
	}()
	defer func() {
		tx.conn.currentTx = nil
	}()
	if !tx.conn.isReady() {
		return badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	return err
}

func (tx *txFake) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (
	rows driver.Rows, err error,
) {
	onDone := trace.DatabaseSQLOnTxQuery(tx.conn.parent.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*txFake).QueryContext"),
		tx.ctx, tx, query,
	)
	defer func() {
		onDone(err)
	}()
	rows, err = tx.conn.QueryContext(ctx, query, args)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return rows, nil
}

func (tx *txFake) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (
	result driver.Result, err error,
) {
	onDone := trace.DatabaseSQLOnTxExec(tx.conn.parent.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*txFake).ExecContext"),
		tx.ctx, tx, query,
	)
	defer func() {
		onDone(err)
	}()
	result, err = tx.conn.ExecContext(ctx, query, args)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return result, nil
}
