package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionTable = "subscriptions"

type Subscriptions interface {
	New() Subscriptions

	Insert(ctx context.Context, sub models.Subscription) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.Subscription, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.Subscription, error)

	FilterUserID(userID string) Subscriptions
	FilterPlanID(planID string) Subscriptions
	FilterStatus(status string) Subscriptions
	FilterPaymentMethodID(paymentMethodID string) Subscriptions

	//TODO: Add FilterStartDate, FilterEndDate, FilterCreatedAt, FilterUpdatedAt

	Page(limit, offset uint64) Subscriptions
}

type subscriptions struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubscriptions(db *sql.DB) Subscriptions {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subscriptions{
		db:       db,
		selector: builder.Select("*").From(subscriptionTable),
		inserter: builder.Insert(subscriptionTable),
		updater:  builder.Update(subscriptionTable),
		deleter:  builder.Delete(subscriptionTable),
		counter:  builder.Select("COUNT(*) AS count").From(subscriptionTable),
	}
}

func (s *subscriptions) New() Subscriptions {
	return NewSubscriptions(s.db)
}

func (s *subscriptions) Insert(ctx context.Context, sub models.Subscription) error {
	query, args, err := s.inserter.SetMap(map[string]interface{}{
		"user_id":           sub.UserID,
		"plan_id":           sub.PlanID,
		"payment_method_id": sub.PaymentMethodID,
		"status":            sub.Status,
		"start_date":        sub.StartDate,
		"end_date":          sub.EndDate,
		"created_at":        sub.CreatedAt,
		"updated_at":        sub.UpdatedAt,
	}).ToSql()

	fmt.Printf("query: %s, args: %v", query, args)
	time.Sleep(1 * time.Second)
	if err != nil {
		return fmt.Errorf("error building insert query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting subscription: %w", err)
	}

	return nil
}

func (s *subscriptions) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := s.updater.
		SetMap(updates).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building update query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating subscription: %w", err)
	}

	return nil
}

func (s *subscriptions) Delete(ctx context.Context) error {
	query, args, err := s.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("error building delete query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error deleting subscription: %w", err)
	}

	return nil
}

func (s *subscriptions) Count(ctx context.Context) (int, error) {
	query, args, err := s.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building count query: %w", err)
	}

	var count int
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting subscriptions: %w", err)
	}

	return count, nil
}

func (s *subscriptions) Select(ctx context.Context) ([]models.Subscription, error) {
	query, args, err := s.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building select query: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing select query: %w", err)
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		err := rows.Scan(
			&sub.UserID,
			&sub.PlanID,
			&sub.PaymentMethodID,
			&sub.Status,
			&sub.StartDate,
			&sub.EndDate,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning subscription row: %w", err)
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func (s *subscriptions) Get(ctx context.Context) (*models.Subscription, error) {
	query, args, err := s.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building get query: %w", err)
	}

	var sub models.Subscription
	err = s.db.QueryRowContext(ctx, query, args...).Scan(
		&sub.UserID,
		&sub.PlanID,
		&sub.PaymentMethodID,
		&sub.Status,
		&sub.StartDate,
		&sub.EndDate,
		&sub.CreatedAt,
		&sub.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting subscription: %w", err)
	}

	return &sub, nil
}

func (s *subscriptions) FilterUserID(userID string) Subscriptions {
	cond := sq.Eq{"user_id": userID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptions) FilterPlanID(planID string) Subscriptions {
	cond := sq.Eq{"plan_id": planID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptions) FilterPaymentMethodID(paymentMethodID string) Subscriptions {
	cond := sq.Eq{"payment_method_id": paymentMethodID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptions) FilterStatus(status string) Subscriptions {
	cond := sq.Eq{"status": status}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptions) Page(limit, offset uint64) Subscriptions {
	s.selector = s.selector.Limit(limit).Offset(offset)
	s.counter = s.counter.Limit(limit).Offset(offset)
	return s
}
