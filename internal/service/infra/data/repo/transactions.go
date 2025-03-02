package repo

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
)

type Transactions interface {
	New() Transactions

	Insert(ctx context.Context, trn models.Transaction) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.Transaction, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.Transaction, error)

	Filter(filters map[string]any) Transactions

	Page(limit, offset uint64) Transactions
}

type transactions struct {
	sql     *sqldb.Transactions
	filters map[string]any
	limit   int64
	skip    int64
}

func NewTransactions(sql *sqldb.Transactions) Transactions {
	return &transactions{
		sql:     sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,
	}
}

func (t *transactions) New() Transactions {
	return NewTransactions(t.sql)
}

func (t *transactions) Insert(ctx context.Context, trn models.Transaction) error {
	return t.sql.New().Insert(ctx, trn)
}

func (t *transactions) Update(ctx context.Context, updates map[string]any) error {
	return t.sql.New().Filter(t.filters).Update(ctx, updates)
}

func (t *transactions) Delete(ctx context.Context) error {
	return t.sql.New().Filter(t.filters).Delete(ctx)
}

func (t *transactions) Select(ctx context.Context) ([]models.Transaction, error) {
	return t.sql.New().Filter(t.filters).Select(ctx)
}

func (t *transactions) Count(ctx context.Context) (int, error) {
	return t.sql.New().Filter(t.filters).Count(ctx)
}

func (t *transactions) Get(ctx context.Context) (*models.Transaction, error) {
	return t.sql.New().Filter(t.filters).Get(ctx)
}

func (t *transactions) Filter(filters map[string]any) Transactions {
	t.filters = filters
	return t
}

func (t *transactions) Page(limit, offset uint64) Transactions {
	t.limit = int64(limit)
	t.skip = int64(offset)
	return t
}
