package daolib

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	// вернет контест с транзакцией
	BeginTx(context.Context, ...TxOption) (txCtx context.Context, err error)
	RollbackTx(context.Context) error
	CommitTx(context.Context) error
	NewRepo(context.Context, BaseQuerySetter)
}

type daoCtxKey int

const txCtxKey daoCtxKey = iota

type dao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) DAO {
	return &dao{db: db}
}

// вернет контест с транзакцией
func (d *dao) BeginTx(ctx context.Context, opts ...TxOption) (txCtx context.Context, err error) {
	txopts := &sql.TxOptions{}
	for _, o := range opts {
		o(txopts)
	}

	tx, err := d.db.BeginTxx(ctx, txopts)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	return context.WithValue(ctx, txCtxKey, tx), nil
}

func (d *dao) RollbackTx(ctx context.Context) error {
	if !d.hasTx(ctx) {
		return NoTransactionError
	}

	if err := d.getSqlxTx(ctx).Rollback(); err != nil {
		return fmt.Errorf("rollback tx: %w", err)
	}
	return nil
}

func (d *dao) CommitTx(ctx context.Context) error {
	if !d.hasTx(ctx) {
		return NoTransactionError
	}

	if err := d.getSqlxTx(ctx).Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

func (d *dao) NewRepo(ctx context.Context, impl BaseQuerySetter) {
	impl.SetBaseQuery(&BaseQuery{ctx: ctx, Runner: d.getRunner(ctx)})
}

func (d *dao) hasTx(ctx context.Context) bool {
	_, ok := ctx.Value(txCtxKey).(*sqlx.Tx)
	return ok
}

// to check have tx use HasTransaction
func (d *dao) getSqlxTx(ctx context.Context) *sqlx.Tx {
	return ctx.Value(txCtxKey).(*sqlx.Tx)
}

func (d *dao) getRunner(ctx context.Context) Runner {
	if d.hasTx(ctx) {
		return NewTx(d.getSqlxTx(ctx))
	}
	return NewDB(d.db)
}
