package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func Subscription(subscription models.Subscription) resources.Subscription {
	return resources.Subscription{
		Data: resources.SubscriptionData{
			Id:   subscription.UserID.String(),
			Type: resources.TypeSubscription,
			Attributes: resources.SubscriptionDataAttributes{
				PlanId:          subscription.PlanID.String(),
				PaymentMethodId: subscription.PaymentMethodID.String(),
				Status:          string(subscription.Status),
				Availability:    string(subscription.Availability),
				StartDate:       subscription.StartDate,
				EndDate:         &subscription.EndDate,
				UpdatedAt:       subscription.UpdatedAt,
				CreatedAt:       subscription.CreatedAt,
			},
		},
	}
}
