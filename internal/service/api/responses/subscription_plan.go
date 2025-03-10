package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func SubscriptionPlan(subscriptionPlan *models.SubscriptionPlan, subscriptionType *models.SubscriptionType) resources.SubscriptionPlan {
	res := resources.SubscriptionPlan{
		Data: resources.SubscriptionPlanData{
			Id:   subscriptionPlan.ID.String(),
			Type: resources.TypeSubscriptionPlan,
			Attributes: resources.SubscriptionPlanDataAttributes{
				Name:            subscriptionPlan.Name,
				Desc:            subscriptionPlan.Description,
				Price:           subscriptionPlan.Price,
				Currency:        subscriptionPlan.Currency,
				BillingInterval: int32(subscriptionPlan.BillingInterval),
				BillingCycle:    string(subscriptionPlan.BillingCycle),
				UpdatedAt:       subscriptionPlan.UpdatedAt,
				CreatedAt:       subscriptionPlan.CreatedAt,
			},
		},
	}
	if subscriptionType != nil {
		res.Data.Relationships = resources.SubscriptionPlanDataRelationships{
			SubscriptionType: resources.Relationships{
				Data: resources.RelationshipsData{
					Id:   subscriptionType.ID.String(),
					Type: resources.TypeSubscriptionType,
				},
			},
		}
	}
	return res
}
