package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const transactionsTable = "subscription_transactions"

type Transactions interface {
	New() Transactions

	Insert(ctx context.Context, trn *models.Transaction) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error
	Select(ctx context.Context) ([]models.Transaction, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.Transaction, error)

	Filter(filters map[string]any) Transactions

	Transaction(fn func(ctx context.Context) error) error

	Page(limit, offset uint64) Transactions
}

type transactions struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewTransactions(db *sql.DB) Transactions {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &transactions{
		db:       db,
		selector: builder.Select("*").From(transactionsTable),
		inserter: builder.Insert(transactionsTable),
		updater:  builder.Update(transactionsTable),
		deleter:  builder.Delete(transactionsTable),
		counter:  builder.Select("COUNT(*) AS count").From(transactionsTable),
	}
}

func (t *transactions) New() Transactions {
	return NewTransactions(t.db)
}

func (t *transactions) Insert(ctx context.Context, trn *models.Transaction) error {
	values := map[string]interface{}{
		"id":                trn.ID,
		"user_id":           trn.UserID,
		"payment_method_id": trn.PaymentMethodID,
		"amount":            trn.Amount,
		"currency":          trn.Currency,
		"status":            trn.Status,
		"payment_provider":  trn.PaymentProvider,
		"payment_id":        trn.PaymentID,
		"transaction_date":  trn.TransactionDate,
	}

	query, args, err := t.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for %s: %w", transactionsTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = t.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("inserting %s: %w", transactionsTable, err)
	}
	return nil
}

func (t *transactions) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := t.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for %s: %w", transactionsTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = t.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("updating %s: %w", transactionsTable, err)
	}
	return nil
}

func (t *transactions) Delete(ctx context.Context) error {
	query, args, err := t.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for %s: %w", transactionsTable, err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = t.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return fmt.Errorf("deleting %s: %w", transactionsTable, err)
	}
	return nil
}

func (t *transactions) Select(ctx context.Context) ([]models.Transaction, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for %s: %w", transactionsTable, err)
	}

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for %s: %w", transactionsTable, err)
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
			return nil, fmt.Errorf("scanning %s: %w", transactionsTable, err)
		}

		results = append(results, trn)
	}
	return results, nil
}

func (t *transactions) Count(ctx context.Context) (int, error) {
	query, args, err := t.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for %s: %w", transactionsTable, err)
	}

	var count int
	if err := t.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting %s: %w", transactionsTable, err)
	}
	return count, nil
}

func (t *transactions) Get(ctx context.Context) (*models.Transaction, error) {
	query, args, err := t.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for %s: %w", transactionsTable, err)
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
		return nil, fmt.Errorf("getting %s: %w", transactionsTable, err)
	}

	return &trn, nil
}

func (t *transactions) Transaction(fn func(ctx context.Context) error) error {
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

func (t *transactions) Filter(filters map[string]any) Transactions {
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

func (t *transactions) Page(limit, offset uint64) Transactions {
	t.selector = t.selector.Limit(limit).Offset(offset)
	t.counter = t.counter.Limit(limit).Offset(offset)
	return t
}
