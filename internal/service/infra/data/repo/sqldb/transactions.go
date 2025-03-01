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

const transactionsTable = "subscription_transactions"

type transactions struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewTransactions(db *sql.DB) repo.Transactions {
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

func (s *transactions) New() repo.Transactions {
	return NewTransactions(s.db)
}

func (s *transactions) Insert(ctx context.Context, trn models.Transaction) error {
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

	query, args, err := s.inserter.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building insert query for subscription_transactions: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("inserting subscription_transaction: %w", err)
	}
	return nil
}

func (s *transactions) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := s.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for subscription_transactions: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating subscription_transaction: %w", err)
	}
	return nil
}

func (s *transactions) Delete(ctx context.Context) error {
	query, args, err := s.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for subscription_transactions: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting subscription_transaction: %w", err)
	}
	return nil
}

func (s *transactions) Select(ctx context.Context) ([]models.Transaction, error) {
	query, args, err := s.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for subscription_transactions: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
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

func (s *transactions) Count(ctx context.Context) (int, error) {
	query, args, err := s.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for subscription_transactions: %w", err)
	}

	var count int
	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting subscription_transactions: %w", err)
	}
	return count, nil
}

func (s *transactions) Get(ctx context.Context) (*models.Transaction, error) {
	query, args, err := s.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for subscription_transactions: %w", err)
	}

	var trn models.Transaction

	err = s.db.QueryRowContext(ctx, query, args...).Scan(
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

func (s *transactions) FilterID(ID uuid.UUID) repo.Transactions {
	cond := sq.Eq{"id": ID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *transactions) FilterUserID(userID uuid.UUID) repo.Transactions {
	cond := sq.Eq{"user_id": userID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *transactions) FilterPaymentMethodID(paymentMethodID uuid.UUID) repo.Transactions {
	cond := sq.Eq{"payment_method_id": paymentMethodID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *transactions) FilterStatus(status models.TrnStatus) repo.Transactions {
	cond := sq.Eq{"status": status}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *transactions) FilterPaymentProvider(provider models.PaymentProvider) repo.Transactions {
	cond := sq.Eq{"payment_provider": provider}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *transactions) Page(limit, offset uint64) repo.Transactions {
	s.selector = s.selector.Limit(limit).Offset(offset)
	s.counter = s.counter.Limit(limit).Offset(offset)
	return s
}
