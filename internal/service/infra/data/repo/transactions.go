package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type Transactions interface {
	New() Transactions

	Insert(ctx context.Context, trn models.Transaction) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.Transaction, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.Transaction, error)

	FilterID(ID uuid.UUID) Transactions
	FilterUserID(userID uuid.UUID) Transactions
	FilterPaymentMethodID(paymentMethodID uuid.UUID) Transactions
	FilterStatus(status models.TrnStatus) Transactions
	FilterPaymentProvider(provider models.PaymentProvider) Transactions

	Page(limit, offset uint64) Transactions
}
