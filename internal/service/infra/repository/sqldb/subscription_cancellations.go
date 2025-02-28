package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionCancellationsTable = "subscription_cancellations"

// TODO DELETE NOT USED
type SubscriptionCancellations interface {
	New() SubscriptionCancellations

	Insert(ctx context.Context, c models.SubscriptionCancellation) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.SubscriptionCancellation, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionCancellation, error)

	FilterID(id uuid.UUID) SubscriptionCancellations
	FilterUserID(userID uuid.UUID) SubscriptionCancellations

	Page(limit, offset uint64) SubscriptionCancellations
}

type subscriptionCancellations struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubscriptionCancellations(db *sql.DB) SubscriptionCancellations {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subscriptionCancellations{
		db:       db,
		selector: builder.Select("*").From(subscriptionCancellationsTable),
		inserter: builder.Insert(subscriptionCancellationsTable),
		updater:  builder.Update(subscriptionCancellationsTable),
		deleter:  builder.Delete(subscriptionCancellationsTable),
		counter:  builder.Select("COUNT(*) AS count").From(subscriptionCancellationsTable),
	}
}

func (c *subscriptionCancellations) New() SubscriptionCancellations {
	return NewSubscriptionCancellations(c.db)
}

func (c *subscriptionCancellations) Insert(ctx context.Context, sc models.SubscriptionCancellation) error {
	values := map[string]interface{}{
		"id":                sc.ID,
		"user_id":           sc.UserID,
		"cancellation_date": sc.CancellationDate,
	}

	if sc.Reason != nil {
		values["reason"] = *sc.Reason
	}

	query, args, err := c.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for subscription_cancellations: %w", err)
	}

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("inserting subscription_cancellation: %w", err)
	}
	return nil
}

func (c *subscriptionCancellations) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := c.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for subscription_cancellations: %w", err)
	}

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating subscription_cancellation: %w", err)
	}
	return nil
}

func (c *subscriptionCancellations) Delete(ctx context.Context) error {
	query, args, err := c.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for subscription_cancellations: %w", err)
	}

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting subscription_cancellation: %w", err)
	}
	return nil
}

func (c *subscriptionCancellations) Select(ctx context.Context) ([]models.SubscriptionCancellation, error) {
	query, args, err := c.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for subscription_cancellations: %w", err)
	}

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for subscription_cancellations: %w", err)
	}
	defer rows.Close()

	var results []models.SubscriptionCancellation
	for rows.Next() {
		var sc models.SubscriptionCancellation
		var reason *string

		err := rows.Scan(
			&sc.ID,
			&sc.UserID,
			&sc.CancellationDate,
			&reason,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning subscription_cancellation row: %w", err)
		}
		sc.Reason = reason
		results = append(results, sc)
	}
	return results, nil
}

func (c *subscriptionCancellations) Count(ctx context.Context) (int, error) {
	query, args, err := c.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for subscription_cancellations: %w", err)
	}

	var count int
	if err := c.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting subscription_cancellations: %w", err)
	}
	return count, nil
}

func (c *subscriptionCancellations) Get(ctx context.Context) (*models.SubscriptionCancellation, error) {
	query, args, err := c.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for subscription_cancellations: %w", err)
	}

	var sc models.SubscriptionCancellation
	var reason *string

	err = c.db.QueryRowContext(ctx, query, args...).Scan(
		&sc.ID,
		&sc.UserID,
		&sc.CancellationDate,
		&reason,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting subscription_cancellation: %w", err)
	}
	sc.Reason = reason

	return &sc, nil
}

func (c *subscriptionCancellations) FilterID(id uuid.UUID) SubscriptionCancellations {
	cond := sq.Eq{"id": id}
	c.selector = c.selector.Where(cond)
	c.updater = c.updater.Where(cond)
	c.deleter = c.deleter.Where(cond)
	c.counter = c.counter.Where(cond)
	return c
}

func (c *subscriptionCancellations) FilterUserID(userID uuid.UUID) SubscriptionCancellations {
	cond := sq.Eq{"user_id": userID}
	c.selector = c.selector.Where(cond)
	c.updater = c.updater.Where(cond)
	c.deleter = c.deleter.Where(cond)
	c.counter = c.counter.Where(cond)
	return c
}

func (c *subscriptionCancellations) Page(limit, offset uint64) SubscriptionCancellations {
	c.selector = c.selector.Limit(limit).Offset(offset)
	c.counter = c.counter.Limit(limit).Offset(offset)
	return c
}
