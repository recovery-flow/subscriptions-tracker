package billing

import (
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type Task struct {
	Schedule models.BillingSchedule
	UserID   uuid.UUID
}
