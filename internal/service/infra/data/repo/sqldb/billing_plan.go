package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const billingPlanTable = "billing_plan"

type BillingPlan struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewBillingSchedules(db *sql.DB) *BillingPlan {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	res := BillingPlan{
		db:       db,
		selector: builder.Select("*").From(billingPlanTable),
		inserter: builder.Insert(billingPlanTable),
		updater:  builder.Update(billingPlanTable),
		deleter:  builder.Delete(billingPlanTable),
		counter:  builder.Select("COUNT(*) AS count").From(billingPlanTable),
	}
	return &res
}

func (b *BillingPlan) New() *BillingPlan {
	return NewBillingSchedules(b.db)
}

func (b *BillingPlan) Insert(ctx context.Context, bs models.BillingPlan) error {
	values := map[string]interface{}{
		"id":             bs.ID,
		"user_id":        bs.UserID,
		"scheduled_date": bs.ScheduledDate,
		"status":         bs.Status,
		"created_at":     bs.CreatedAt,
	}

	if bs.AttemptedDate != nil {
		values["attempted_date"] = *bs.AttemptedDate
	}

	query, args, err := b.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for billing_schedules: %w", err)
	}

	if _, err := b.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("inserting billing_schedule: %w", err)
	}
	return nil
}

func (b *BillingPlan) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := b.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for billing_schedules: %w", err)
	}

	if _, err := b.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating billing_schedule: %w", err)
	}
	return nil
}

func (b *BillingPlan) Delete(ctx context.Context) error {
	query, args, err := b.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for billing_schedules: %w", err)
	}

	if _, err := b.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting billing_schedule: %w", err)
	}
	return nil
}

func (b *BillingPlan) Select(ctx context.Context) ([]models.BillingPlan, error) {
	query, args, err := b.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for billing_schedules: %w", err)
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for billing_schedules: %w", err)
	}
	defer rows.Close()

	var results []models.BillingPlan
	for rows.Next() {
		var bs models.BillingPlan
		var attemptedDate *time.Time

		err := rows.Scan(
			&bs.ID,
			&bs.UserID,
			&bs.ScheduledDate,
			&attemptedDate,
			&bs.Status,
			&bs.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning billing_schedule row: %w", err)
		}
		bs.AttemptedDate = attemptedDate
		results = append(results, bs)
	}
	return results, nil
}

func (b *BillingPlan) Count(ctx context.Context) (int, error) {
	query, args, err := b.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for billing_schedules: %w", err)
	}

	var count int
	if err := b.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting billing_schedules: %w", err)
	}
	return count, nil
}

func (b *BillingPlan) Get(ctx context.Context) (*models.BillingPlan, error) {
	query, args, err := b.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for billing_schedules: %w", err)
	}

	var bs models.BillingPlan
	var attemptedDate *time.Time

	err = b.db.QueryRowContext(ctx, query, args...).Scan(
		&bs.ID,
		&bs.UserID,
		&bs.ScheduledDate,
		&attemptedDate,
		&bs.Status,
		&bs.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting billing_schedule: %w", err)
	}
	bs.AttemptedDate = attemptedDate

	return &bs, nil
}

func (b *BillingPlan) Filter(filters map[string]any) *BillingPlan {
	var validFilters = map[string]bool{
		"id":      true,
		"user_id": true,
		"status":  true,
	}

	for key, value := range filters {
		if _, exists := validFilters[key]; !exists {
			continue
		}
		b.selector = b.selector.Where(sq.Eq{key: value})
		b.counter = b.counter.Where(sq.Eq{key: value})
		b.deleter = b.deleter.Where(sq.Eq{key: value})
		b.updater = b.updater.Where(sq.Eq{key: value})
	}
	return b
}

func (b *BillingPlan) Page(limit, offset uint64) *BillingPlan {
	b.selector = b.selector.Limit(limit).Offset(offset)
	b.counter = b.counter.Limit(limit).Offset(offset)
	return b
}
