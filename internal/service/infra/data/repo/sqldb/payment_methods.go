package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo"
)

const paymentMethodsTable = "payment_methods"

type paymentMethods struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewPaymentMethods(db *sql.DB) repo.PaymentMethods {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &paymentMethods{
		db:       db,
		selector: builder.Select("*").From(paymentMethodsTable),
		inserter: builder.Insert(paymentMethodsTable),
		updater:  builder.Update(paymentMethodsTable),
		deleter:  builder.Delete(paymentMethodsTable),
		counter:  builder.Select("COUNT(*) AS count").From(paymentMethodsTable),
	}
}

func (p *paymentMethods) New() repo.PaymentMethods {
	return NewPaymentMethods(p.db)
}

func (p *paymentMethods) Insert(ctx context.Context, pm models.PaymentMethod) error {
	values := map[string]interface{}{
		"id":             pm.ID,
		"user_id":        pm.UserID,
		"type":           pm.Type,
		"provider_token": pm.ProviderToken,
		"is_default":     pm.IsDefault,
		"created_at":     pm.CreatedAt,
	}

	query, args, err := p.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for payment_methods: %w", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("inserting payment_method: %w", err)
	}
	return nil
}

func (p *paymentMethods) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := p.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for payment_methods: %w", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating payment_method: %w", err)
	}
	return nil
}

func (p *paymentMethods) Delete(ctx context.Context) error {
	query, args, err := p.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for payment_methods: %w", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting payment_method: %w", err)
	}
	return nil
}

func (p *paymentMethods) Select(ctx context.Context) ([]models.PaymentMethod, error) {
	query, args, err := p.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for payment_methods: %w", err)
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for payment_methods: %w", err)
	}
	defer rows.Close()

	var results []models.PaymentMethod
	for rows.Next() {
		var pm models.PaymentMethod
		if err := rows.Scan(
			&pm.ID,
			&pm.UserID,
			&pm.Type,
			&pm.ProviderToken,
			&pm.IsDefault,
			&pm.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scanning payment_method row: %w", err)
		}
		results = append(results, pm)
	}
	return results, nil
}

func (p *paymentMethods) Count(ctx context.Context) (int, error) {
	query, args, err := p.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for payment_methods: %w", err)
	}

	var count int
	if err := p.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting payment_methods: %w", err)
	}
	return count, nil
}

func (p *paymentMethods) Get(ctx context.Context) (*models.PaymentMethod, error) {
	query, args, err := p.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for payment_methods: %w", err)
	}

	var pm models.PaymentMethod
	err = p.db.QueryRowContext(ctx, query, args...).Scan(
		&pm.ID,
		&pm.UserID,
		&pm.Type,
		&pm.ProviderToken,
		&pm.IsDefault,
		&pm.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting payment_method: %w", err)
	}
	return &pm, nil
}

func (p *paymentMethods) FilterID(id uuid.UUID) repo.PaymentMethods {
	cond := sq.Eq{"id": id}
	p.selector = p.selector.Where(cond)
	p.updater = p.updater.Where(cond)
	p.deleter = p.deleter.Where(cond)
	p.counter = p.counter.Where(cond)
	return p
}

func (p *paymentMethods) FilterUserID(userID uuid.UUID) repo.PaymentMethods {
	cond := sq.Eq{"user_id": userID}
	p.selector = p.selector.Where(cond)
	p.updater = p.updater.Where(cond)
	p.deleter = p.deleter.Where(cond)
	p.counter = p.counter.Where(cond)
	return p
}

func (p *paymentMethods) Page(limit, offset uint64) repo.PaymentMethods {
	p.selector = p.selector.Limit(limit).Offset(offset)
	p.counter = p.counter.Limit(limit).Offset(offset)
	return p
}
