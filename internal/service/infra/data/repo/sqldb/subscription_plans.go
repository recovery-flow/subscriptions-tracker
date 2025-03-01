package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionPlansTable = "subscription_plans"

type SubscriptionPlan interface {
	New() SubscriptionPlan

	Insert(ctx context.Context, plan models.SubscriptionPlan) error
	Update(ctx context.Context, plan models.SubscriptionPlan) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.SubscriptionPlan, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionPlan, error)

	Transaction(func() error) error

	FilterID(id uuid.UUID) SubscriptionPlan
	FilterTypeID(typeID uuid.UUID) SubscriptionPlan

	Page(limit, offset uint64) SubscriptionPlan
}

type subscriptionPlan struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubscriptionPlan(db *sql.DB) SubscriptionPlan {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subscriptionPlan{
		db:       db,
		selector: builder.Select("*").From(subscriptionPlansTable),
		inserter: builder.Insert(subscriptionPlansTable),
		updater:  builder.Update(subscriptionPlansTable),
		deleter:  builder.Delete(subscriptionPlansTable),
		counter:  builder.Select("COUNT(*) as count").From(subscriptionPlansTable),
	}
}

func (p *subscriptionPlan) New() SubscriptionPlan {
	return NewSubscriptionPlan(p.db)
}

func (p *subscriptionPlan) Insert(ctx context.Context, plan models.SubscriptionPlan) error {
	values := map[string]interface{}{
		"id":                    plan.ID,
		"type_id":               plan.TypeID,
		"price":                 plan.Price,
		"billing_interval":      plan.BillingInterval,
		"billing_interval_unit": plan.BillingIntervalUnit,
		"currency":              plan.Currency,
		"created_at":            plan.CreatedAt,
	}

	query, args, err := p.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for subscription_plans: %w", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("inserting subscription_plan: %w", err)
	}
	return nil
}

func (p *subscriptionPlan) Update(ctx context.Context, plan models.SubscriptionPlan) error {
	updates := map[string]interface{}{
		"type_id":               plan.TypeID,
		"price":                 plan.Price,
		"billing_interval":      plan.BillingInterval,
		"billing_interval_unit": plan.BillingIntervalUnit,
		"currency":              plan.Currency,
	}
	query, args, err := p.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for subscription_plans: %w", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating subscription_plan: %w", err)
	}
	return nil
}

func (p *subscriptionPlan) Delete(ctx context.Context) error {
	query, args, err := p.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for subscription_plans: %w", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting subscription_plan: %w", err)
	}
	return nil
}

func (p *subscriptionPlan) Select(ctx context.Context) ([]models.SubscriptionPlan, error) {
	query, args, err := p.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for subscription_plans: %w", err)
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for subscription_plans: %w", err)
	}
	defer rows.Close()

	var plans []models.SubscriptionPlan
	for rows.Next() {
		var plan models.SubscriptionPlan
		err := rows.Scan(
			&plan.ID,
			&plan.TypeID,
			&plan.Price,
			&plan.BillingInterval,
			&plan.BillingIntervalUnit,
			&plan.Currency,
			&plan.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning subscription_plan row: %w", err)
		}
		plans = append(plans, plan)
	}
	return plans, nil
}

func (p *subscriptionPlan) Count(ctx context.Context) (int, error) {
	query, args, err := p.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for subscription_plans: %w", err)
	}

	var count int
	if err := p.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting subscription_plans: %w", err)
	}
	return count, nil
}

func (p *subscriptionPlan) Get(ctx context.Context) (*models.SubscriptionPlan, error) {
	query, args, err := p.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for subscription_plans: %w", err)
	}

	var plan models.SubscriptionPlan
	err = p.db.QueryRowContext(ctx, query, args...).Scan(
		&plan.ID,
		&plan.TypeID,
		&plan.Price,
		&plan.BillingInterval,
		&plan.BillingIntervalUnit,
		&plan.Currency,
		&plan.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting subscription_plan: %w", err)
	}

	return &plan, nil
}

func (p *subscriptionPlan) FilterID(id uuid.UUID) SubscriptionPlan {
	cond := sq.Eq{"id": id}
	p.selector = p.selector.Where(cond)
	p.updater = p.updater.Where(cond)
	p.deleter = p.deleter.Where(cond)
	p.counter = p.counter.Where(cond)
	return p
}

func (p *subscriptionPlan) FilterTypeID(typeID uuid.UUID) SubscriptionPlan {
	cond := sq.Eq{"type_id": typeID}
	p.selector = p.selector.Where(cond)
	p.updater = p.updater.Where(cond)
	p.deleter = p.deleter.Where(cond)
	p.counter = p.counter.Where(cond)
	return p
}

func (p *subscriptionPlan) Transaction(f func() error) error {
	ctx := context.Background()

	tx, err := p.db.BeginTx(ctx, nil)
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

func (p *subscriptionPlan) Page(limit, offset uint64) SubscriptionPlan {
	p.selector = p.selector.Limit(limit).Offset(offset)
	p.counter = p.counter.Limit(limit).Offset(offset)
	return p
}
