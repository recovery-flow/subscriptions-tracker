package repo

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
)

type BillingSchedules interface {
	New() BillingSchedules

	Insert(ctx context.Context, bs models.BillingSchedule) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.BillingSchedule, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.BillingSchedule, error)

	Filter(map[string]any) BillingSchedules

	Page(limit, offset uint64) BillingSchedules
}

type billingSchedule struct {
	sql     *sqldb.billingSchedules
	filters map[string]any
	limit   int64
	skip    int64
}

func NewBillingPlan(sql *sqldb.billingSchedules) BillingSchedules {
	return &billingSchedule{
		sql:     sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,
	}
}

func (b *billingSchedule) New() BillingSchedules {
	return NewBillingPlan(b.sql)
}

func (b *billingSchedule) Insert(ctx context.Context, bs models.BillingSchedule) error {
	return b.sql.New().Insert(ctx, bs)
}

func (b *billingSchedule) Update(ctx context.Context, updates map[string]any) error {
	return b.sql.New().Filter(b.filters).Update(ctx, updates)
}

func (b *billingSchedule) Delete(ctx context.Context) error {
	return b.sql.New().Filter(b.filters).Delete(ctx)
}

func (b *billingSchedule) Select(ctx context.Context) ([]models.BillingSchedule, error) {
	return b.sql.New().Filter(b.filters).Page(uint64(b.limit), uint64(b.skip)).Select(ctx)
}

func (b *billingSchedule) Count(ctx context.Context) (int, error) {
	return b.sql.New().Filter(b.filters).Count(ctx)
}

func (b *billingSchedule) Get(ctx context.Context) (*models.BillingSchedule, error) {
	return b.sql.New().Filter(b.filters).Get(ctx)
}

func (b *billingSchedule) Filter(filters map[string]any) BillingSchedules {
	b.filters = filters
	return b
}

func (b *billingSchedule) Page(limit, offset uint64) BillingSchedules {
	b.limit = int64(limit)
	b.skip = int64(offset)
	return b
}
