package query

import (
	"context"
	"fmt"

	"github.com/ydb-platform/ydb-go-genproto/Ydb_Query_V1"
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Query"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/allocator"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/session"
	queryTx "github.com/ydb-platform/ydb-go-sdk/v3/internal/query/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	baseTx "github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsync"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

var (
	_ query.Transaction  = (*Transaction)(nil)
	_ baseTx.Transaction = (*Transaction)(nil)
)

type (
	Transaction struct {
		baseTx.LazyID

		s          *Session
		txSettings query.TransactionSettings

		completed bool

		onBeforeCommit xsync.Set[*baseTx.OnTransactionBeforeCommit]
		onCompleted    xsync.Set[*baseTx.OnTransactionCompletedFunc]
	}
)

func begin(
	ctx context.Context,
	client Ydb_Query_V1.QueryServiceClient,
	sessionID string,
	txSettings query.TransactionSettings,
) (txID string, _ error) {
	a := allocator.New()
	defer a.Free()
	response, err := client.BeginTransaction(ctx,
		&Ydb_Query.BeginTransactionRequest{
			SessionId:  sessionID,
			TxSettings: txSettings.ToYDB(a),
		},
	)
	if err != nil {
		return "", xerrors.WithStackTrace(err)
	}

	return response.GetTxMeta().GetId(), nil
}

func (tx *Transaction) UnLazy(ctx context.Context) error {
	if tx.ID() != baseTx.LazyTxID {
		return nil
	}

	txID, err := begin(ctx, tx.s.client, tx.s.ID(), tx.txSettings)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	tx.SetTxID(txID)

	return nil
}

func (tx *Transaction) QueryResultSet(
	ctx context.Context, q string, opts ...options.Execute,
) (rs result.ClosableResultSet, finalErr error) {
	onDone := trace.QueryOnTxQueryResultSet(tx.s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Transaction).QueryResultSet"), tx, q)
	defer func() {
		onDone(finalErr)
	}()

	if tx.completed {
		return nil, xerrors.WithStackTrace(errExecuteOnCompletedTx)
	}

	settings, err := tx.executeSettings(opts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	resultOpts := []resultOption{
		withTrace(tx.s.trace),
		onTxMeta(func(txMeta *Ydb_Query.TransactionMeta) {
			tx.SetTxID(txMeta.GetId())
		}),
	}
	if settings.TxControl().Commit {
		err = tx.waitOnBeforeCommit(ctx)
		if err != nil {
			return nil, err
		}

		// notification about complete transaction must be sended for any error or for successfully read all result if
		// it was execution with commit flag
		resultOpts = append(resultOpts,
			onNextPartErr(func(err error) {
				tx.notifyOnCompleted(xerrors.HideEOF(err))
			}),
		)
	}
	r, err := execute(ctx, tx.s.ID(), tx.s.client, q, settings, resultOpts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	rs, err = readResultSet(ctx, r)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return rs, nil
}

func (tx *Transaction) QueryRow(
	ctx context.Context, q string, opts ...options.Execute,
) (row query.Row, finalErr error) {
	onDone := trace.QueryOnTxQueryRow(tx.s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Transaction).QueryRow"), tx, q)
	defer func() {
		onDone(finalErr)
	}()

	settings := options.ExecuteSettings(
		append(
			[]options.Execute{options.WithTxControl(tx.txControl())},
			opts...,
		)...,
	)

	resultOpts := []resultOption{
		withTrace(tx.s.trace),
		onTxMeta(func(txMeta *Ydb_Query.TransactionMeta) {
			tx.SetTxID(txMeta.GetId())
		}),
	}
	if settings.TxControl().Commit {
		err := tx.waitOnBeforeCommit(ctx)
		if err != nil {
			return nil, err
		}

		// notification about complete transaction must be sended for any error or for successfully read all result if
		// it was execution with commit flag
		resultOpts = append(resultOpts,
			onNextPartErr(func(err error) {
				tx.notifyOnCompleted(xerrors.HideEOF(err))
			}),
		)
	}
	r, err := execute(ctx, tx.s.ID(), tx.s.client, q, settings, resultOpts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	row, err = readRow(ctx, r)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return row, nil
}

func (tx *Transaction) SessionID() string {
	return tx.s.ID()
}

func (tx *Transaction) txControl() *queryTx.Control {
	if tx.ID() != baseTx.LazyTxID {
		return queryTx.NewControl(queryTx.WithTxID(tx.ID()))
	}

	return queryTx.NewControl(
		queryTx.BeginTx(tx.txSettings...),
	)
}

func (tx *Transaction) Exec(ctx context.Context, q string, opts ...options.Execute) (
	finalErr error,
) {
	onDone := trace.QueryOnTxExec(tx.s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Transaction).Exec"), tx.s, tx, q)
	defer func() {
		onDone(finalErr)
	}()

	if tx.completed {
		return xerrors.WithStackTrace(errExecuteOnCompletedTx)
	}

	settings, err := tx.executeSettings(opts...)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	resultOpts := []resultOption{
		withTrace(tx.s.trace),
		onTxMeta(func(txMeta *Ydb_Query.TransactionMeta) {
			tx.SetTxID(txMeta.GetId())
		}),
	}
	if settings.TxControl().Commit {
		err = tx.waitOnBeforeCommit(ctx)
		if err != nil {
			return err
		}

		// notification about complete transaction must be sended for any error or for successfully read all result if
		// it was execution with commit flag
		resultOpts = append(resultOpts,
			onNextPartErr(func(err error) {
				tx.notifyOnCompleted(xerrors.HideEOF(err))
			}),
		)
	}

	r, err := execute(ctx, tx.s.ID(), tx.s.client, q, settings, resultOpts...)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	err = readAll(ctx, r)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (tx *Transaction) executeSettings(opts ...options.Execute) (_ executeSettings, finalErr error) {
	for _, opt := range opts {
		if opt == nil {
			return nil, xerrors.WithStackTrace(errNilOption)
		}
		if _, has := opt.(options.ExecuteNoTx); has {
			return nil, xerrors.WithStackTrace(
				fmt.Errorf("%T: %w", opt, ErrOptionNotForTxExecute),
			)
		}
	}

	return options.ExecuteSettings(append([]options.Execute{
		options.WithTxControl(tx.txControl()),
	}, opts...)...), nil
}

func (tx *Transaction) Query(ctx context.Context, q string, opts ...options.Execute) (
	_ query.Result, finalErr error,
) {
	onDone := trace.QueryOnTxQuery(tx.s.trace, &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/query.(*Transaction).Query"), tx.s, tx, q)
	defer func() {
		onDone(finalErr)
	}()

	if tx.completed {
		return nil, xerrors.WithStackTrace(errExecuteOnCompletedTx)
	}

	settings, err := tx.executeSettings(opts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	resultOpts := []resultOption{
		withTrace(tx.s.trace),
		onTxMeta(func(txMeta *Ydb_Query.TransactionMeta) {
			tx.SetTxID(txMeta.GetId())
		}),
	}
	if settings.TxControl().Commit {
		err = tx.waitOnBeforeCommit(ctx)
		if err != nil {
			return nil, err
		}

		// notification about complete transaction must be sended for any error or for successfully read all result if
		// it was execution with commit flag
		resultOpts = append(resultOpts,
			onNextPartErr(func(err error) {
				tx.notifyOnCompleted(xerrors.HideEOF(err))
			}),
		)
	}
	r, err := execute(ctx, tx.s.ID(), tx.s.client, q, settings, resultOpts...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return r, nil
}

func commitTx(ctx context.Context, client Ydb_Query_V1.QueryServiceClient, sessionID, txID string) error {
	_, err := client.CommitTransaction(ctx, &Ydb_Query.CommitTransactionRequest{
		SessionId: sessionID,
		TxId:      txID,
	})
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (tx *Transaction) CommitTx(ctx context.Context) (finalErr error) {
	if tx.ID() == baseTx.LazyTxID {
		return nil
	}

	if tx.completed {
		return nil
	}

	defer func() {
		tx.notifyOnCompleted(finalErr)
		tx.completed = true
	}()

	err := tx.waitOnBeforeCommit(ctx)
	if err != nil {
		return err
	}

	err = commitTx(ctx, tx.s.client, tx.s.ID(), tx.ID())
	if err != nil {
		if xerrors.IsOperationError(err, Ydb.StatusIds_BAD_SESSION) {
			tx.s.SetStatus(session.StatusClosed)
		}

		return xerrors.WithStackTrace(err)
	}

	return nil
}

func rollback(ctx context.Context, client Ydb_Query_V1.QueryServiceClient, sessionID, txID string) error {
	_, err := client.RollbackTransaction(ctx, &Ydb_Query.RollbackTransactionRequest{
		SessionId: sessionID,
		TxId:      txID,
	})
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (tx *Transaction) Rollback(ctx context.Context) (finalErr error) {
	if tx.ID() == baseTx.LazyTxID {
		// https://github.com/ydb-platform/ydb-go-sdk/issues/1456
		return tx.s.Close(ctx)
	}

	if tx.completed {
		return nil
	}

	tx.completed = true

	tx.notifyOnCompleted(ErrTransactionRollingBack)

	err := rollback(ctx, tx.s.client, tx.s.ID(), tx.ID())
	if err != nil {
		if xerrors.IsOperationError(err, Ydb.StatusIds_BAD_SESSION) {
			tx.s.SetStatus(session.StatusClosed)
		}

		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (tx *Transaction) OnBeforeCommit(f baseTx.OnTransactionBeforeCommit) {
	tx.onBeforeCommit.Add(&f)
}

func (tx *Transaction) OnCompleted(f baseTx.OnTransactionCompletedFunc) {
	tx.onCompleted.Add(&f)
}

func (tx *Transaction) waitOnBeforeCommit(ctx context.Context) (resErr error) {
	tx.onBeforeCommit.Range(func(f *baseTx.OnTransactionBeforeCommit) bool {
		resErr = (*f)(ctx)

		return resErr == nil
	})

	return resErr
}

func (tx *Transaction) notifyOnCompleted(err error) {
	tx.completed = true

	tx.onCompleted.Range(func(f *baseTx.OnTransactionCompletedFunc) bool {
		(*f)(err)

		return tx.onCompleted.Remove(f)
	})
}