package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type PaymentMethods interface {
	New() PaymentMethods

	Insert(ctx context.Context, pm models.PaymentMethod) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.PaymentMethod, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.PaymentMethod, error)

	FilterID(id uuid.UUID) PaymentMethods
	FilterUserID(userID uuid.UUID) PaymentMethods

	Page(limit, offset uint64) PaymentMethods
}
