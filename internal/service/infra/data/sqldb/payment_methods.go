package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const paymentMethodsTable = "payment_methods"

type PaymentMethods interface {
	New() PaymentMethods

	Insert(ctx context.Context, pm *models.PaymentMethod) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.PaymentMethod, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.PaymentMethod, error)

	Filter(filters map[string]any) PaymentMethods

	Update(ctx context.Context, updates map[string]any) error

	Transaction(fn func(ctx context.Context) error) error

	Page(limit, offset uint64) PaymentMethods
}

type paymentMethods struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewPaymentMethods(db *sql.DB) PaymentMethods {
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

func (m *paymentMethods) New() PaymentMethods {
	return NewPaymentMethods(m.db)
}

func (m *paymentMethods) Insert(ctx context.Context, pm *models.PaymentMethod) error {
	values := map[string]interface{}{
		"id":             pm.ID,
		"user_id":        pm.UserID,
		"type":           pm.Type,
		"provider_token": pm.ProviderToken,
		"is_default":     pm.IsDefault,
		"created_at":     time.Now().UTC(),
	}

	query, args, err := m.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for %s: %w", paymentMethodsTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = m.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("inserting %s: %w", paymentMethodsTable, err)
	}
	return nil
}

func (m *paymentMethods) Update(ctx context.Context, updates map[string]any) error {
	updates["updated_at"] = time.Now().UTC()
	query, args, err := m.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for %s: %w", paymentMethodsTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = m.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("updating %s: %w", paymentMethodsTable, err)
	}
	return nil
}

func (m *paymentMethods) Delete(ctx context.Context) error {
	query, args, err := m.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for %s: %w", paymentMethodsTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = m.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("deleting %s: %w", paymentMethodsTable, err)
	}
	return nil
}

func (m *paymentMethods) Select(ctx context.Context) ([]models.PaymentMethod, error) {
	query, args, err := m.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for %s: %w", paymentMethodsTable, err)
	}

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for %s: %w", paymentMethodsTable, err)
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
			return nil, fmt.Errorf("scanning %s row: %w", paymentMethodsTable, err)
		}
		results = append(results, pm)
	}
	return results, nil
}

func (m *paymentMethods) Count(ctx context.Context) (int, error) {
	query, args, err := m.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for %s: %w", paymentMethodsTable, err)
	}

	var count int
	if err := m.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting %s: %w", paymentMethodsTable, err)
	}
	return count, nil
}

func (m *paymentMethods) Get(ctx context.Context) (*models.PaymentMethod, error) {
	query, args, err := m.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for %s: %w", paymentMethodsTable, err)
	}

	var pm models.PaymentMethod
	err = m.db.QueryRowContext(ctx, query, args...).Scan(
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
		return nil, fmt.Errorf("getting %s: %w", paymentMethodsTable, err)
	}
	return &pm, nil
}

func (m *paymentMethods) Transaction(fn func(ctx context.Context) error) error {
	ctx := context.Background()

	tx, err := m.db.BeginTx(ctx, nil)
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

func (m *paymentMethods) Filter(filters map[string]any) PaymentMethods {
	var validFilters = map[string]bool{
		"id":      true,
		"user_id": true,
	}
	for key, value := range filters {
		if _, exists := validFilters[key]; !exists {
			continue
		}
		m.selector = m.selector.Where(sq.Eq{key: value})
		m.counter = m.counter.Where(sq.Eq{key: value})
		m.deleter = m.deleter.Where(sq.Eq{key: value})
		m.updater = m.updater.Where(sq.Eq{key: value})
	}
	return m
}

func (m *paymentMethods) Page(limit, offset uint64) PaymentMethods {
	m.selector = m.selector.Limit(limit).Offset(offset)
	m.counter = m.counter.Limit(limit).Offset(offset)
	return m
}
