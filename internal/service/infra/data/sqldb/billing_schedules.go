package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const billingSchedulesTable = "billing_schedule"

type BillingSchedules interface {
	New() BillingSchedules

	Insert(ctx context.Context, bs *models.BillingSchedule) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.BillingSchedule, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.BillingSchedule, error)

	Filter(filters map[string]any) BillingSchedules
	FilterTime(field string, after bool, date time.Time) BillingSchedules

	Transaction(fn func(ctx context.Context) error) error

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

func (b *billingSchedules) Insert(ctx context.Context, bs *models.BillingSchedule) error {
	values := map[string]interface{}{
		"user_id":        bs.UserID,
		"scheduled_date": bs.SchedulesDate,
		"status":         bs.Status,
		//"updated_at":     bs.UpdatedAt,
		"created_at": bs.CreatedAt,
	}

	if bs.AttemptedDate != nil {
		values["attempted_date"] = *bs.AttemptedDate
	}

	query, args, err := b.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for %s: %w", billingSchedulesTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = b.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("inserting %s: %w", billingSchedulesTable, err)
	}
	return nil
}

func (b *billingSchedules) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := b.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for %s: %w", billingSchedulesTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = b.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("updating %s: %w", billingSchedulesTable, err)
	}
	return nil
}

func (b *billingSchedules) Delete(ctx context.Context) error {
	query, args, err := b.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for %s: %w", billingSchedulesTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = b.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("deleting %s: %w", billingSchedulesTable, err)
	}
	return nil
}

func (b *billingSchedules) Select(ctx context.Context) ([]models.BillingSchedule, error) {
	query, args, err := b.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for %s: %w", billingSchedulesTable, err)
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for %s: %w", billingSchedulesTable, err)
	}
	defer rows.Close()

	var results []models.BillingSchedule
	for rows.Next() {
		var bs models.BillingSchedule
		var attemptedDate sql.NullTime

		err := rows.Scan(
			&bs.UserID,
			&bs.SchedulesDate,
			&attemptedDate,
			&bs.Status,
			//&bs.UpdatedAt,
			&bs.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning %s row: %w", billingSchedulesTable, err)
		}

		if attemptedDate.Valid {
			bs.AttemptedDate = &attemptedDate.Time
		} else {
			bs.AttemptedDate = nil
		}

		results = append(results, bs)
	}
	return results, nil
}

func (b *billingSchedules) Count(ctx context.Context) (int, error) {
	query, args, err := b.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for %s: %w", billingSchedulesTable, err)
	}

	var count int
	if err := b.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting %s: %w", billingSchedulesTable, err)
	}
	return count, nil
}

func (b *billingSchedules) Get(ctx context.Context) (*models.BillingSchedule, error) {
	query, args, err := b.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for %s: %w", billingSchedulesTable, err)
	}

	var bs models.BillingSchedule
	var attemptedDate *time.Time

	err = b.db.QueryRowContext(ctx, query, args...).Scan(
		&bs.UserID,
		&bs.SchedulesDate,
		&attemptedDate,
		&bs.Status,
		//&bs.UpdatedAt,
		&bs.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting %s: %w", billingSchedulesTable, err)
	}
	bs.AttemptedDate = attemptedDate

	return &bs, nil
}

func (b *billingSchedules) Transaction(fn func(ctx context.Context) error) error {
	ctx := context.Background()

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	ctxWithTx := context.WithValue(ctx, txKey, tx)

	if err := fn(ctxWithTx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction failed: %v, rollback error: %v", err, rbErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (b *billingSchedules) Filter(filters map[string]any) BillingSchedules {
	var validFilters = map[string]bool{
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

func (b *billingSchedules) FilterTime(field string, after bool, date time.Time) BillingSchedules {
	var validFields = map[string]bool{
		"schedules_date": true,
		"attempted_date": true,
	}
	if _, exists := validFields[field]; !exists {
		return b
	}

	if after {
		b.selector = b.selector.Where(sq.Gt{field: date})
		b.counter = b.counter.Where(sq.Gt{field: date})
		b.deleter = b.deleter.Where(sq.Gt{field: date})
		b.updater = b.updater.Where(sq.Gt{field: date})
	} else {
		b.selector = b.selector.Where(sq.Lt{field: date})
		b.counter = b.counter.Where(sq.Lt{field: date})
		b.deleter = b.deleter.Where(sq.Lt{field: date})
		b.updater = b.updater.Where(sq.Lt{field: date})
	}
	return b
}

func (b *billingSchedules) Page(limit, offset uint64) BillingSchedules {
	b.selector = b.selector.Limit(limit).Offset(offset)
	b.counter = b.counter.Limit(limit).Offset(offset)
	return b
}
