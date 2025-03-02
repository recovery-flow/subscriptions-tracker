package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const transactionsTable = "subscription_transactions"

type Transactions struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewTransactions(db *sql.DB) *Transactions {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	res := Transactions{
		db:       db,
		selector: builder.Select("*").From(transactionsTable),
		inserter: builder.Insert(transactionsTable),
		updater:  builder.Update(transactionsTable),
		deleter:  builder.Delete(transactionsTable),
		counter:  builder.Select("COUNT(*) AS count").From(transactionsTable),
	}
	return &res
}

func (t *Transactions) New() *Transactions {
	return NewTransactions(t.db)
}

func (t *Transactions) Insert(ctx context.Context, trn models.Transaction) error {
	values := map[string]interface{}{
		"id":                trn.ID,
		"user_id":           trn.UserID,
		"amount":            trn.Amount,
		"currency":          trn.Currency,
		"status":            trn.Status,
		"payment_provider":  trn.PaymentProvider,
		"transaction_date":  trn.TransactionDate,
		"payment_id":        trn.PaymentID,
		"payment_method_id": trn.PaymentMethodID,
	}

	query, args, err := t.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for subscription_transactions: %w", err)
	}

	if _, err := t.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("inserting subscription_transaction: %w", err)
	}
	return nil
}

func (t *Transactions) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := t.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for subscription_transactions: %w", err)
	}

	if _, err := t.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating subscription_transaction: %w", err)
	}
	return nil
}

func (t *Transactions) Delete(ctx context.Context) error {
	query, args, err := t.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for subscription_transactions: %w", err)
	}

	if _, err := t.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting subscription_transaction: %w", err)
	}
	return nil
}

func (t *Transactions) Select(ctx context.Context) ([]models.Transaction, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for subscription_transactions: %w", err)
	}

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for subscription_transactions: %w", err)
	}
	defer rows.Close()

	var results []models.Transaction
	for rows.Next() {
		var trn models.Transaction

		err := rows.Scan(
			&trn.ID,
			&trn.UserID,
			&trn.PaymentMethodID,
			&trn.Amount,
			&trn.Currency,
			&trn.Status,
			&trn.PaymentProvider,
			&trn.PaymentID,
			&trn.TransactionDate,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning subscription_transaction: %w", err)
		}

		results = append(results, trn)
	}
	return results, nil
}

func (t *Transactions) Count(ctx context.Context) (int, error) {
	query, args, err := t.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for subscription_transactions: %w", err)
	}

	var count int
	if err := t.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting subscription_transactions: %w", err)
	}
	return count, nil
}

func (t *Transactions) Get(ctx context.Context) (*models.Transaction, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for subscription_transactions: %w", err)
	}

	var trn models.Transaction

	err = t.db.QueryRowContext(ctx, query, args...).Scan(
		&trn.ID,
		&trn.UserID,
		&trn.PaymentMethodID,
		&trn.Amount,
		&trn.Currency,
		&trn.Status,
		&trn.PaymentProvider,
		&trn.PaymentID,
		&trn.TransactionDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting subscription_transaction: %w", err)
	}

	return &trn, nil
}

func (t *Transactions) Filter(filters map[string]any) *Transactions {
	var validFilters = map[string]bool{
		"id":                true,
		"user_id":           true,
		"status":            true,
		"payment_provider":  true,
		"payment_method_id": true,
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

func (t *Transactions) Page(limit, offset uint64) *Transactions {
	t.selector = t.selector.Limit(limit).Offset(offset)
	t.counter = t.counter.Limit(limit).Offset(offset)
	return t
}
