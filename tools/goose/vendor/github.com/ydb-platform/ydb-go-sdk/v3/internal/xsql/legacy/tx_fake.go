package legacy

import (
	"context"
	"database/sql/driver"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/params"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql/iface"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql/legacy/badconn"
)

type txFake struct {
	conn *Conn
	ctx  context.Context //nolint:containedctx
}

func (t *txFake) Exec(ctx context.Context, sql string, params *params.Params) (driver.Result, error) {
	result, err := t.conn.Exec(ctx, sql, params)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return result, nil
}

func (t *txFake) Query(ctx context.Context, sql string, params *params.Params) (driver.RowsNextResultSet, error) {
	rows, err := t.conn.Query(ctx, sql, params)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return rows, nil
}

func (t *txFake) ID() string {
	return tx.FakeTxID
}

func beginTxFake(ctx context.Context, c *Conn) iface.Tx {
	return &txFake{
		conn: c,
		ctx:  ctx,
	}
}

func (t *txFake) Commit(ctx context.Context) (err error) {
	if !t.conn.isReady() {
		return badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	return nil
}

func (t *txFake) Rollback(ctx context.Context) (err error) {
	if !t.conn.isReady() {
		return badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	return err
}
