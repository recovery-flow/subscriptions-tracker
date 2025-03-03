package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionTypesTable = "subscription_types"

type SubTypes interface {
	New() SubTypes

	Insert(ctx context.Context, sub models.SubscriptionType) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.SubscriptionType, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionType, error)

	Filter(filters map[string]any) SubTypes

	Transaction(func() error) error

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

func (t *subTypes) Insert(ctx context.Context, sub models.SubscriptionType) error {
	query, args, err := t.inserter.SetMap(map[string]interface{}{
		"id":          sub.ID,
		"name":        sub.Name,
		"description": sub.Description,
		"status":      sub.Status,
		"updated_at":  time.Now().UTC(),
		"created_at":  time.Now().UTC(),
	}).ToSql()
	if err != nil {
		return fmt.Errorf("error building insert query for subscription_types: %w", err)
	}

	_, err = t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting subscription_type: %w", err)
	}
	return nil
}

func (t *subTypes) Update(ctx context.Context, updates map[string]any) error {
	updates["updated_at"] = time.Now().UTC()
	query, args, err := t.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("error building update query for subscription_types: %w", err)
	}

	_, err = t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating subscription_type: %w", err)
	}
	return nil
}

func (t *subTypes) Delete(ctx context.Context) error {
	query, args, err := t.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("error building delete query for subscription_types: %w", err)
	}

	_, err = t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error deleting subscription_type: %w", err)
	}
	return nil
}

func (t *subTypes) Select(ctx context.Context) ([]models.SubscriptionType, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building select query for subscription_types: %w", err)
	}

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing select query for subscription_types: %w", err)
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
			return nil, fmt.Errorf("error scanning subscription_type row: %w", err)
		}
		types = append(types, st)
	}
	return types, nil
}

func (t *subTypes) Count(ctx context.Context) (int, error) {
	query, args, err := t.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building count query for subscription_types: %w", err)
	}

	var count int
	if err := t.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("error counting subscription_types: %w", err)
	}
	return count, nil
}

func (t *subTypes) Get(ctx context.Context) (*models.SubscriptionType, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building get query for subscription_types: %w", err)
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
		return nil, fmt.Errorf("error getting subscription_type: %w", err)
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

func (t *subTypes) Transaction(f func() error) error {
	ctx := context.Background()

	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	if err := f(); err != nil {
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
