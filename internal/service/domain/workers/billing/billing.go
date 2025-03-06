package billing

import (
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type BillingTask struct {
	Schedule models.BillingSchedule
	UserID   uuid.UUID
}
