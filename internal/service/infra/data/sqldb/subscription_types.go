package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionTypesTable = "subscription_types"

type SubTypes interface {
	New() SubTypes

	Insert(ctx context.Context, sub *models.SubscriptionType) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.SubscriptionType, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionType, error)

	Filter(filters map[string]any) SubTypes

	Transaction(fn func(ctx context.Context) error) error

	Page(limit, offset uint64) SubTypes
}

type subTypes struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubTypes(db *sql.DB) SubTypes {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subTypes{
		db:       db,
		selector: builder.Select("*").From(subscriptionTypesTable),
		inserter: builder.Insert(subscriptionTypesTable),
		updater:  builder.Update(subscriptionTypesTable),
		deleter:  builder.Delete(subscriptionTypesTable),
		counter:  builder.Select("COUNT(*) AS count").From(subscriptionTypesTable),
	}
}

func (t *subTypes) New() SubTypes {
	return NewSubTypes(t.db)
}

func (t *subTypes) Insert(ctx context.Context, sub *models.SubscriptionType) error {
	query, args, err := t.inserter.SetMap(map[string]interface{}{
		"id":          sub.ID,
		"name":        sub.Name,
		"description": sub.Description,
		"status":      sub.Status,
		"updated_at":  sub.UpdatedAt,
		"created_at":  sub.CreatedAt,
	}).ToSql()
	if err != nil {
		return fmt.Errorf("error building insert query for %s: %w", subscriptionTypesTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = t.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("error inserting %s: %w", subscriptionTypesTable, err)
	}
	return nil
}

func (t *subTypes) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := t.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("error building update query for %s: %w", subscriptionTypesTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = t.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("error updating %s: %w", subscriptionTypesTable, err)
	}
	return nil
}

func (t *subTypes) Delete(ctx context.Context) error {
	query, args, err := t.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("error building delete query for %s: %w", subscriptionTypesTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = t.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("error deleting %s: %w", subscriptionTypesTable, err)
	}
	return nil
}

func (t *subTypes) Select(ctx context.Context) ([]models.SubscriptionType, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building select query for %s: %w", subscriptionTypesTable, err)
	}

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing select query for %s: %w", subscriptionTypesTable, err)
	}
	defer rows.Close()

	var types []models.SubscriptionType
	for rows.Next() {
		var st models.SubscriptionType
		if err := rows.Scan(
			&st.ID,
			&st.Name,
			&st.Description,
			&st.Status,
			&st.UpdatedAt,
			&st.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning %s row: %w", subscriptionTypesTable, err)
		}
		types = append(types, st)
	}
	return types, nil
}

func (t *subTypes) Count(ctx context.Context) (int, error) {
	query, args, err := t.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building count query for %s: %w", subscriptionTypesTable, err)
	}

	var count int
	if err := t.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("error counting %s: %w", subscriptionTypesTable, err)
	}
	return count, nil
}

func (t *subTypes) Get(ctx context.Context) (*models.SubscriptionType, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building get query for %s: %w", subscriptionTypesTable, err)
	}

	var st models.SubscriptionType
	err = t.db.QueryRowContext(ctx, query, args...).Scan(
		&st.ID,
		&st.Name,
		&st.Description,
		&st.Status,
		&st.UpdatedAt,
		&st.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting %s: %w", subscriptionTypesTable, err)
	}
	return &st, nil
}

func (t *subTypes) Filter(filters map[string]any) SubTypes {
	var validFilters = map[string]bool{
		"id":     true,
		"status": true,
		"name":   true,
	}
	for key, value := range filters {
		if _, exists := validFilters[key]; !exists {
			continue
		}
		t.selector = t.selector.Where(sq.Eq{key: value})
		t.counter = t.counter.Where(sq.Eq{key: value})
		t.deleter = t.deleter.Where(sq.Eq{key: value})
		t.updater = t.updater.Where(sq.Eq{key: value})
	}
	return t
}

func (t *subTypes) Transaction(fn func(ctx context.Context) error) error {
	ctx := context.Background()

	tx, err := t.db.BeginTx(ctx, nil)
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

func (t *subTypes) Page(limit, offset uint64) SubTypes {
	t.selector = t.selector.Limit(limit).Offset(offset)
	t.counter = t.counter.Limit(limit).Offset(offset)
	return t
}
