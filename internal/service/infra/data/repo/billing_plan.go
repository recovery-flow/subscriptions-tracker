package repo

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
)

type BillingPlan interface {
	New() BillingPlan

	Insert(ctx context.Context, bs models.BillingPlan) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.BillingPlan, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.BillingPlan, error)

	Filter(map[string]any) BillingPlan

	Page(limit, offset uint64) BillingPlan
}

type billingPlan struct {
	sql     *sqldb.BillingPlan
	filters map[string]any
	limit   int64
	skip    int64
}

func NewBillingPlan(sql *sqldb.BillingPlan) BillingPlan {
	return &billingPlan{
		sql:     sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,
	}
}

func (b *billingPlan) New() BillingPlan {
	return NewBillingPlan(b.sql)
}

func (b *billingPlan) Insert(ctx context.Context, bs models.BillingPlan) error {
	return b.sql.New().Insert(ctx, bs)
}

func (b *billingPlan) Update(ctx context.Context, updates map[string]any) error {
	return b.sql.New().Filter(b.filters).Update(ctx, updates)
}

func (b *billingPlan) Delete(ctx context.Context) error {
	return b.sql.New().Filter(b.filters).Delete(ctx)
}

func (b *billingPlan) Select(ctx context.Context) ([]models.BillingPlan, error) {
	return b.sql.New().Filter(b.filters).Page(uint64(b.limit), uint64(b.skip)).Select(ctx)
}

func (b *billingPlan) Count(ctx context.Context) (int, error) {
	return b.sql.New().Filter(b.filters).Count(ctx)
}

func (b *billingPlan) Get(ctx context.Context) (*models.BillingPlan, error) {
	return b.sql.New().Filter(b.filters).Get(ctx)
}

func (b *billingPlan) Filter(filters map[string]any) BillingPlan {
	b.filters = filters
	return b
}

func (b *billingPlan) Page(limit, offset uint64) BillingPlan {
	b.limit = int64(limit)
	b.skip = int64(offset)
	return b
}
