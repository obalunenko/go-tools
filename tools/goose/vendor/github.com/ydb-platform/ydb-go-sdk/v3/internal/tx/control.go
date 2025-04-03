package tx

import (
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Query"
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Table"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/allocator"
)

type (
	Selector interface {
		applyQueryTxSelector(a *allocator.Allocator, txControl *Ydb_Query.TransactionControl)
		applyTableTxSelector(a *allocator.Allocator, txControl *Ydb_Table.TransactionControl)
	}
	ControlOption interface {
		applyTxControlOption(txControl *Control)
	}
	Control struct {
		selector Selector
		commit   bool
	}
)

func (ctrl *Control) Commit() bool {
	return ctrl.commit
}

func (ctrl *Control) ToYdbQueryTransactionControl(a *allocator.Allocator) *Ydb_Query.TransactionControl {
	if ctrl == nil {
		return nil
	}

	txControl := a.QueryTransactionControl()
	ctrl.selector.applyQueryTxSelector(a, txControl)
	txControl.CommitTx = ctrl.commit

	return txControl
}

func (ctrl *Control) ToYdbTableTransactionControl(a *allocator.Allocator) *Ydb_Table.TransactionControl {
	if ctrl == nil {
		return nil
	}

	txControl := a.TableTransactionControl()
	ctrl.selector.applyTableTxSelector(a, txControl)
	txControl.CommitTx = ctrl.commit

	return txControl
}

func (ctrl *Control) Selector() Selector {
	return ctrl.selector
}

var (
	_ ControlOption = beginTxOptions{}
	_ Selector      = beginTxOptions{}
)

type beginTxOptions []SettingsOption

func (opts beginTxOptions) applyTxControlOption(txControl *Control) {
	txControl.selector = opts
}

func (opts beginTxOptions) applyQueryTxSelector(a *allocator.Allocator, txControl *Ydb_Query.TransactionControl) {
	selector := a.QueryTransactionControlBeginTx()
	selector.BeginTx = a.QueryTransactionSettings()
	for _, opt := range opts {
		if opt != nil {
			opt.ApplyQueryTxSettingsOption(a, selector.BeginTx)
		}
	}
	txControl.TxSelector = selector
}

func (opts beginTxOptions) applyTableTxSelector(a *allocator.Allocator, txControl *Ydb_Table.TransactionControl) {
	selector := a.TableTransactionControlBeginTx()
	selector.BeginTx = a.TableTransactionSettings()
	for _, opt := range opts {
		if opt != nil {
			opt.ApplyTableTxSettingsOption(a, selector.BeginTx)
		}
	}
	txControl.TxSelector = selector
}

// BeginTx returns selector transaction control option
func BeginTx(opts ...SettingsOption) beginTxOptions {
	return opts
}

var (
	_ ControlOption = txIDTxControlOption("")
	_ Selector      = txIDTxControlOption("")
)

type txIDTxControlOption string

func (id txIDTxControlOption) applyTxControlOption(txControl *Control) {
	txControl.selector = id
}

func (id txIDTxControlOption) applyQueryTxSelector(a *allocator.Allocator, txControl *Ydb_Query.TransactionControl) {
	selector := a.QueryTransactionControlTxID()
	selector.TxId = string(id)
	txControl.TxSelector = selector
}

func (id txIDTxControlOption) applyTableTxSelector(a *allocator.Allocator, txControl *Ydb_Table.TransactionControl) {
	selector := a.TableTransactionControlTxID()
	selector.TxId = string(id)
	txControl.TxSelector = selector
}

func WithTx(t Identifier) txIDTxControlOption {
	return txIDTxControlOption(t.ID())
}

func WithTxID(txID string) txIDTxControlOption {
	return txIDTxControlOption(txID)
}

type commitTxOption struct{}

func (c commitTxOption) applyTxControlOption(txControl *Control) {
	txControl.commit = true
}

// CommitTx returns commit transaction control option
func CommitTx() ControlOption {
	return commitTxOption{}
}

// NewControl makes transaction control from given options
func NewControl(opts ...ControlOption) *Control {
	txControl := &Control{
		selector: BeginTx(WithSerializableReadWrite()),
		commit:   false,
	}
	for _, opt := range opts {
		if opt != nil {
			opt.applyTxControlOption(txControl)
		}
	}

	return txControl
}

func WithCommit(ctrl *Control) *Control {
	ctrl.commit = true

	return ctrl
}

func NoTx() *Control {
	return nil
}

// DefaultTxControl returns default transaction control with serializable read-write isolation mode and auto-commit
func DefaultTxControl() *Control {
	return NoTx()
}

// SerializableReadWriteTxControl returns transaction control with serializable read-write isolation mode
func SerializableReadWriteTxControl(opts ...ControlOption) *Control {
	return NewControl(
		append([]ControlOption{
			BeginTx(WithSerializableReadWrite()),
		}, opts...)...,
	)
}

// OnlineReadOnlyTxControl returns online read-only transaction control
func OnlineReadOnlyTxControl(opts ...OnlineReadOnlyOption) *Control {
	return NewControl(
		BeginTx(WithOnlineReadOnly(opts...)),
		CommitTx(), // open transactions not supported for OnlineReadOnly
	)
}

// StaleReadOnlyTxControl returns stale read-only transaction control
func StaleReadOnlyTxControl() *Control {
	return NewControl(
		BeginTx(WithStaleReadOnly()),
		CommitTx(), // open transactions not supported for StaleReadOnly
	)
}

// SnapshotReadOnlyTxControl returns snapshot read-only transaction control
func SnapshotReadOnlyTxControl() *Control {
	return NewControl(
		BeginTx(WithSnapshotReadOnly()),
		CommitTx(), // open transactions not supported for StaleReadOnly
	)
}
