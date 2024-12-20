package conn

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn/badconn"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn/isolation"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type transaction struct {
	tx.Identifier

	conn *Conn
	ctx  context.Context //nolint:containedctx
	tx   table.Transaction
}

var (
	_ driver.Tx             = &transaction{}
	_ driver.ExecerContext  = &transaction{}
	_ driver.QueryerContext = &transaction{}
	_ tx.Identifier         = &transaction{}
)

func beginTx(ctx context.Context, c *Conn, txOptions driver.TxOptions) (currentTx, error) {
	txc, err := isolation.ToYDB(txOptions)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}
	nativeTx, err := c.session.BeginTransaction(ctx, table.TxSettings(txc))
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return &transaction{
		Identifier: tx.ID(nativeTx.ID()),
		conn:       c,
		ctx:        ctx,
		tx:         nativeTx,
	}, nil
}

func (tx *transaction) checkTxState() error {
	if tx.conn.currentTx == tx {
		return nil
	}
	if tx.conn.currentTx == nil {
		return fmt.Errorf("broken conn state: tx=%q not related to conn=%q",
			tx.ID(), tx.conn.ID(),
		)
	}

	return fmt.Errorf("broken conn state: tx=%s not related to conn=%q (conn have current tx=%q)",
		tx.conn.currentTx.ID(), tx.conn.ID(), tx.ID(),
	)
}

func (tx *transaction) Commit() (finalErr error) {
	var (
		ctx    = tx.ctx
		onDone = trace.DatabaseSQLOnTxCommit(tx.conn.parent.Trace(), &ctx,
			stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*transaction).Commit"),
			tx,
		)
	)
	defer func() {
		onDone(finalErr)
	}()
	if err := tx.checkTxState(); err != nil {
		return badconn.Map(xerrors.WithStackTrace(err))
	}
	defer func() {
		tx.conn.currentTx = nil
	}()
	if _, err := tx.tx.CommitTx(tx.ctx); err != nil {
		return badconn.Map(xerrors.WithStackTrace(err))
	}

	return nil
}

func (tx *transaction) Rollback() (finalErr error) {
	var (
		ctx    = tx.ctx
		onDone = trace.DatabaseSQLOnTxRollback(tx.conn.parent.Trace(), &ctx,
			stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*transaction).Rollback"),
			tx,
		)
	)
	defer func() {
		onDone(finalErr)
	}()
	if err := tx.checkTxState(); err != nil {
		return badconn.Map(xerrors.WithStackTrace(err))
	}
	defer func() {
		tx.conn.currentTx = nil
	}()
	err := tx.tx.Rollback(tx.ctx)
	if err != nil {
		return badconn.Map(xerrors.WithStackTrace(err))
	}

	return err
}

func (tx *transaction) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (
	_ driver.Rows, finalErr error,
) {
	onDone := trace.DatabaseSQLOnTxQuery(tx.conn.parent.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*transaction).QueryContext"),
		tx.ctx, tx, query,
	)
	defer func() {
		onDone(finalErr)
	}()
	m := queryModeFromContext(ctx, tx.conn.defaultQueryMode)
	if m != DataQueryMode {
		return nil, badconn.Map(
			xerrors.WithStackTrace(
				xerrors.Retryable(
					fmt.Errorf("wrong query mode: %s", m.String()),
					xerrors.InvalidObject(),
					xerrors.WithName("WRONG_QUERY_MODE"),
				),
			),
		)
	}
	query, parameters, err := tx.conn.normalize(query, args...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}
	res, err := tx.tx.Execute(ctx,
		query, &parameters, tx.conn.dataQueryOptions(ctx)...,
	)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}
	if err = res.Err(); err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return &rows{
		conn:   tx.conn,
		result: res,
	}, nil
}

func (tx *transaction) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (
	_ driver.Result, finalErr error,
) {
	onDone := trace.DatabaseSQLOnTxExec(tx.conn.parent.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*transaction).ExecContext"),
		tx.ctx, tx, query,
	)
	defer func() {
		onDone(finalErr)
	}()
	m := queryModeFromContext(ctx, tx.conn.defaultQueryMode)
	if m != DataQueryMode {
		return nil, badconn.Map(
			xerrors.WithStackTrace(
				xerrors.Retryable(
					fmt.Errorf("wrong query mode: %s", m.String()),
					xerrors.InvalidObject(),
					xerrors.WithName("WRONG_QUERY_MODE"),
				),
			),
		)
	}
	query, parameters, err := tx.conn.normalize(query, args...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}
	_, err = tx.tx.Execute(ctx,
		query, &parameters, tx.conn.dataQueryOptions(ctx)...,
	)
	if err != nil {
		return nil, badconn.Map(xerrors.WithStackTrace(err))
	}

	return resultNoRows{}, nil
}

func (tx *transaction) PrepareContext(ctx context.Context, query string) (_ driver.Stmt, finalErr error) {
	onDone := trace.DatabaseSQLOnTxPrepare(tx.conn.parent.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn.(*transaction).PrepareContext"),
		tx.ctx, tx, query,
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
