package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const billingSchedulesTable = "billing_schedules"

type BillingSchedules interface {
	New() BillingSchedules

	Insert(ctx context.Context, bs models.BillingSchedule) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.BillingSchedule, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.BillingSchedule, error)

	FilterID(id uuid.UUID) BillingSchedules
	FilterUserID(userID uuid.UUID) BillingSchedules
	FilterStatus(status string) BillingSchedules

	Page(limit, offset uint64) BillingSchedules
}

type billingSchedules struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewBillingSchedules(db *sql.DB) BillingSchedules {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &billingSchedules{
		db:       db,
		selector: builder.Select("*").From(billingSchedulesTable),
		inserter: builder.Insert(billingSchedulesTable),
		updater:  builder.Update(billingSchedulesTable),
		deleter:  builder.Delete(billingSchedulesTable),
		counter:  builder.Select("COUNT(*) AS count").From(billingSchedulesTable),
	}
}

func (b *billingSchedules) New() BillingSchedules {
	return NewBillingSchedules(b.db)
}

func (b *billingSchedules) Insert(ctx context.Context, bs models.BillingSchedule) error {
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

func (b *billingSchedules) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := b.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for billing_schedules: %w", err)
	}

	if _, err := b.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating billing_schedule: %w", err)
	}
	return nil
}

func (b *billingSchedules) Delete(ctx context.Context) error {
	query, args, err := b.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for billing_schedules: %w", err)
	}

	if _, err := b.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting billing_schedule: %w", err)
	}
	return nil
}

func (b *billingSchedules) Select(ctx context.Context) ([]models.BillingSchedule, error) {
	query, args, err := b.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for billing_schedules: %w", err)
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for billing_schedules: %w", err)
	}
	defer rows.Close()

	var results []models.BillingSchedule
	for rows.Next() {
		var bs models.BillingSchedule
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

func (b *billingSchedules) Count(ctx context.Context) (int, error) {
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

func (b *billingSchedules) Get(ctx context.Context) (*models.BillingSchedule, error) {
	query, args, err := b.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for billing_schedules: %w", err)
	}

	var bs models.BillingSchedule
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

func (b *billingSchedules) FilterID(id uuid.UUID) BillingSchedules {
	cond := sq.Eq{"id": id}
	b.selector = b.selector.Where(cond)
	b.updater = b.updater.Where(cond)
	b.deleter = b.deleter.Where(cond)
	b.counter = b.counter.Where(cond)
	return b
}

func (b *billingSchedules) FilterUserID(userID uuid.UUID) BillingSchedules {
	cond := sq.Eq{"user_id": userID}
	b.selector = b.selector.Where(cond)
	b.updater = b.updater.Where(cond)
	b.deleter = b.deleter.Where(cond)
	b.counter = b.counter.Where(cond)
	return b
}

func (b *billingSchedules) FilterStatus(status string) BillingSchedules {
	cond := sq.Eq{"status": status}
	b.selector = b.selector.Where(cond)
	b.updater = b.updater.Where(cond)
	b.deleter = b.deleter.Where(cond)
	b.counter = b.counter.Where(cond)
	return b
}

func (b *billingSchedules) Page(limit, offset uint64) BillingSchedules {
	b.selector = b.selector.Limit(limit).Offset(offset)
	b.counter = b.counter.Limit(limit).Offset(offset)
	return b
}
