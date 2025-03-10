package responses

import (
	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func Subscription(subscription models.Subscription, typeID *uuid.UUID) resources.Subscription {
	res := resources.Subscription{
		Data: resources.SubscriptionData{
			Id:   subscription.UserID.String(),
			Type: resources.TypeSubscription,
			Attributes: resources.SubscriptionDataAttributes{
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

	if typeID != nil {
		res.Data.Relationships = resources.SubscriptionDataRelationships{
			Plan: resources.Relationships{
				Data: resources.RelationshipsData{
					Id:   subscription.PlanID.String(),
					Type: resources.TypeSubscriptionPlan,
				},
			},
			Type: resources.Relationships{
				Data: resources.RelationshipsData{
					Id:   typeID.String(),
					Type: resources.TypeSubscriptionType,
				},
			},
			PaymentMethod: resources.Relationships{
				Data: resources.RelationshipsData{
					Id:   subscription.PaymentMethodID.String(),
					Type: resources.TypePaymentMethod,
				},
			},
		}
	}

	return res
}
