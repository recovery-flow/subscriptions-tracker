package responses

import (
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func BillingSchedule(schedule models.BillingSchedule) resources.BillingSchedule {
	res := resources.BillingSchedule{
		Data: resources.BillingScheduleData{
			Id:   schedule.UserID.String(),
			Type: resources.TypeSubscriptionType,
			Attributes: resources.BillingScheduleDataAttributes{
				SchedulesDate: schedule.SchedulesDate,
				Status:        string(schedule.Status),
				UpdatedAt:     time.Now(),
				CreatedAt:     schedule.CreatedAt,
			},
		},
	}
	if schedule.AttemptedDate != nil {
		res.Data.Attributes.AttemptedDate = schedule.AttemptedDate
	}
	return res
}
