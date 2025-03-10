package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionPlansTable = "subscription_plans"

type SubPlan interface {
	New() SubPlan

	Insert(ctx context.Context, plan *models.SubscriptionPlan) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.SubscriptionPlan, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionPlan, error)

	Filter(filters map[string]any) SubPlan

	Transaction(fn func(ctx context.Context) error) error

	Page(limit, offset uint64) SubPlan
}

type subPlan struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubPlan(db *sql.DB) SubPlan {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subPlan{
		db:       db,
		selector: builder.Select("*").From(subscriptionPlansTable),
		inserter: builder.Insert(subscriptionPlansTable),
		updater:  builder.Update(subscriptionPlansTable),
		deleter:  builder.Delete(subscriptionPlansTable),
		counter:  builder.Select("COUNT(*) as count").From(subscriptionPlansTable),
	}
}

func (p *subPlan) New() SubPlan {
	return NewSubPlan(p.db)
}

func (p *subPlan) Insert(ctx context.Context, plan *models.SubscriptionPlan) error {
	values := map[string]interface{}{
		"id":               plan.ID,
		"type_id":          plan.TypeID,
		"name":             plan.Name,
		"description":      plan.Description,
		"price":            plan.Price,
		"currency":         plan.Currency,
		"billing_interval": plan.BillingInterval,
		"billing_cycle":    plan.BillingCycle,
		"status":           plan.Status,
		"updated_at":       plan.UpdatedAt,
		"created_at":       plan.CreatedAt,
	}

	query, args, err := p.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for %s: %w", subscriptionPlansTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = p.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("inserting %s: %w", subscriptionPlansTable, err)
	}
	return nil
}

func (p *subPlan) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := p.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for %s: %w", subscriptionPlansTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = p.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("updating %s: %w", subscriptionPlansTable, err)
	}
	return nil
}

func (p *subPlan) Delete(ctx context.Context) error {
	query, args, err := p.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for %s: %w", subscriptionPlansTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = p.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("deleting %s: %w", subscriptionPlansTable, err)
	}
	return nil
}

func (p *subPlan) Select(ctx context.Context) ([]models.SubscriptionPlan, error) {
	query, args, err := p.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for %s: %w", subscriptionPlansTable, err)
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for %s: %w", subscriptionPlansTable, err)
	}
	defer rows.Close()

	var plans []models.SubscriptionPlan
	for rows.Next() {
		var plan models.SubscriptionPlan
		err := rows.Scan(
			&plan.ID,
			&plan.TypeID,
			&plan.Name,
			&plan.Description,
			&plan.Price,
			&plan.Currency,
			&plan.BillingInterval,
			&plan.BillingCycle,
			&plan.Status,
			&plan.UpdatedAt,
			&plan.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning %s row: %w", subscriptionPlansTable, err)
		}
		plans = append(plans, plan)
	}
	return plans, nil
}

func (p *subPlan) Count(ctx context.Context) (int, error) {
	query, args, err := p.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for %s: %w", subscriptionPlansTable, err)
	}

	var count int
	if err := p.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting %s: %w", subscriptionPlansTable, err)
	}
	return count, nil
}

func (p *subPlan) Get(ctx context.Context) (*models.SubscriptionPlan, error) {
	query, args, err := p.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for %s: %w", subscriptionPlansTable, err)
	}

	var plan models.SubscriptionPlan
	err = p.db.QueryRowContext(ctx, query, args...).Scan(
		&plan.ID,
		&plan.TypeID,
		&plan.Name,
		&plan.Description,
		&plan.Price,
		&plan.Currency,
		&plan.BillingInterval,
		&plan.BillingCycle,
		&plan.Status,
		&plan.UpdatedAt,
		&plan.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting %s: %w", subscriptionPlansTable, err)
	}

	return &plan, nil
}

func (p *subPlan) Filter(filters map[string]any) SubPlan {
	var validFilters = map[string]bool{
		"id":      true,
		"type_id": true,
		"status":  true,
		"name":    true,
	}
	for key, value := range filters {
		if _, exists := validFilters[key]; !exists {
			continue
		}
		p.selector = p.selector.Where(sq.Eq{key: value})
		p.counter = p.counter.Where(sq.Eq{key: value})
		p.deleter = p.deleter.Where(sq.Eq{key: value})
		p.updater = p.updater.Where(sq.Eq{key: value})
	}
	return p
}

func (p *subPlan) Transaction(fn func(ctx context.Context) error) error {
	ctx := context.Background()

	tx, err := p.db.BeginTx(ctx, nil)
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

func (p *subPlan) Page(limit, offset uint64) SubPlan {
	p.selector = p.selector.Limit(limit).Offset(offset)
	p.counter = p.counter.Limit(limit).Offset(offset)
	return p
}
