package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type BillingPlan interface {
	New() BillingPlan

	Insert(ctx context.Context, bs models.BillingPlan) error
	Update(ctx context.Context, updates map[string]any) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.BillingPlan, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.BillingPlan, error)

	FilterID(id uuid.UUID) BillingPlan
	FilterUserID(userID uuid.UUID) BillingPlan
	FilterStatus(status string) BillingPlan

	Page(limit, offset uint64) BillingPlan
}
