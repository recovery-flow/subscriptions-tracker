package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

const subscriptionTransactionsTable = "subscription_transactions"

type SubscriptionTransactions interface {
	New() SubscriptionTransactions

	Insert(ctx context.Context, trn models.SubscriptionTransaction) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.SubscriptionTransaction, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.SubscriptionTransaction, error)

	FilterID(ID uuid.UUID) SubscriptionTransactions
	FilterUserID(userID uuid.UUID) SubscriptionTransactions
	FilterPaymentMethodID(paymentMethodID uuid.UUID) SubscriptionTransactions
	FilterStatus(status models.TrnStatus) SubscriptionTransactions
	FilterPaymentProvider(provider models.PaymentProvider) SubscriptionTransactions

	Page(limit, offset uint64) SubscriptionTransactions
}

type subscriptionTransactions struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewSubscriptionTransactions(db *sql.DB) SubscriptionTransactions {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &subscriptionTransactions{
		db:       db,
		selector: builder.Select("*").From(subscriptionTransactionsTable),
		inserter: builder.Insert(subscriptionTransactionsTable),
		updater:  builder.Update(subscriptionTransactionsTable),
		deleter:  builder.Delete(subscriptionTransactionsTable),
		counter:  builder.Select("COUNT(*) AS count").From(subscriptionTransactionsTable),
	}
}

func (s *subscriptionTransactions) New() SubscriptionTransactions {
	return NewSubscriptionTransactions(s.db)
}

func (s *subscriptionTransactions) Insert(ctx context.Context, trn models.SubscriptionTransaction) error {
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

func (s *subscriptionTransactions) Update(ctx context.Context, updates map[string]any) error {
	query, args, err := s.updater.SetMap(updates).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for subscription_transactions: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("updating subscription_transaction: %w", err)
	}
	return nil
}

func (s *subscriptionTransactions) Delete(ctx context.Context) error {
	query, args, err := s.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for subscription_transactions: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("deleting subscription_transaction: %w", err)
	}
	return nil
}

func (s *subscriptionTransactions) Select(ctx context.Context) ([]models.SubscriptionTransaction, error) {
	query, args, err := s.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for subscription_transactions: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing select query for subscription_transactions: %w", err)
	}
	defer rows.Close()

	var results []models.SubscriptionTransaction
	for rows.Next() {
		var trn models.SubscriptionTransaction

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

func (s *subscriptionTransactions) Count(ctx context.Context) (int, error) {
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

func (s *subscriptionTransactions) Get(ctx context.Context) (*models.SubscriptionTransaction, error) {
	query, args, err := s.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building get query for subscription_transactions: %w", err)
	}

	var trn models.SubscriptionTransaction

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

func (s *subscriptionTransactions) FilterID(ID uuid.UUID) SubscriptionTransactions {
	cond := sq.Eq{"id": ID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptionTransactions) FilterUserID(userID uuid.UUID) SubscriptionTransactions {
	cond := sq.Eq{"user_id": userID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptionTransactions) FilterPaymentMethodID(paymentMethodID uuid.UUID) SubscriptionTransactions {
	cond := sq.Eq{"payment_method_id": paymentMethodID}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.deleter = s.deleter.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptionTransactions) FilterStatus(status models.TrnStatus) SubscriptionTransactions {
	cond := sq.Eq{"status": status}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptionTransactions) FilterPaymentProvider(provider models.PaymentProvider) SubscriptionTransactions {
	cond := sq.Eq{"payment_provider": provider}
	s.selector = s.selector.Where(cond)
	s.updater = s.updater.Where(cond)
	s.counter = s.counter.Where(cond)
	return s
}

func (s *subscriptionTransactions) Page(limit, offset uint64) SubscriptionTransactions {
	s.selector = s.selector.Limit(limit).Offset(offset)
	s.counter = s.counter.Limit(limit).Offset(offset)
	return s
}
