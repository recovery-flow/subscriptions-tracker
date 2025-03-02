package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionTypesTable = "subscription_types"

type SubscriptionTypes interface {
	New() SubscriptionTypes

	Insert(ctx context.Context, sub models.SubscriptionType) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.SubscriptionType, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionType, error)

	Filter(filters map[string]any) SubscriptionTypes

	Page(limit, offset uint64) SubscriptionTypes
}

type subscriptionTypes struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubscriptionTypes(db *sql.DB) SubscriptionTypes {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subscriptionTypes{
		db:       db,
		selector: builder.Select("*").From(subscriptionTypesTable),
		inserter: builder.Insert(subscriptionTypesTable),
		updater:  builder.Update(subscriptionTypesTable),
		deleter:  builder.Delete(subscriptionTypesTable),
		counter:  builder.Select("COUNT(*) AS count").From(subscriptionTypesTable),
	}
}

func (t *subscriptionTypes) New() SubscriptionTypes {
	return NewSubscriptionTypes(t.db)
}

func (t *subscriptionTypes) Insert(ctx context.Context, sub models.SubscriptionType) error {
	query, args, err := t.inserter.SetMap(map[string]interface{}{
		"id":          sub.ID,
		"name":        sub.Name,
		"description": sub.Description,
		"created_at":  sub.CreatedAt,
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

func (t *subscriptionTypes) Update(ctx context.Context, updates map[string]any) error {
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

func (t *subscriptionTypes) Delete(ctx context.Context) error {
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

func (t *subscriptionTypes) Select(ctx context.Context) ([]models.SubscriptionType, error) {
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
		var t models.SubscriptionType
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning subscription_type row: %w", err)
		}
		types = append(types, t)
	}
	return types, nil
}

func (t *subscriptionTypes) Count(ctx context.Context) (int, error) {
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

func (t *subscriptionTypes) Get(ctx context.Context) (*models.SubscriptionType, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building get query for subscription_types: %w", err)
	}

	var subt models.SubscriptionType
	err = t.db.QueryRowContext(ctx, query, args...).Scan(
		&subt.ID,
		&subt.Name,
		&subt.Description,
		&subt.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting subscription_type: %w", err)
	}
	return &subt, nil
}

func (t *subscriptionTypes) Filter(filters map[string]any) SubscriptionTypes {
	var validFilters = map[string]bool{
		"id": true,
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

func (t *subscriptionTypes) Page(limit, offset uint64) SubscriptionTypes {
	t.selector = t.selector.Limit(limit).Offset(offset)
	t.counter = t.counter.Limit(limit).Offset(offset)
	return t
}
