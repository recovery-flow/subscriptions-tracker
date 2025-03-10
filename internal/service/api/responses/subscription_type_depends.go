package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func SubscriptionTypeDepends(subscriptions []models.SubscriptionTypeDepends) resources.SubscriptionTypeDepends {
	types := make([]resources.SubscriptionTypeData, 0)
	plans := make([]resources.SubscriptionPlanData, 0)

	for _, subscription := range subscriptions {
		typeRelations := make([]resources.Relationships, 0)

		for _, plan := range subscription.Plans {
			plans = append(plans, resources.SubscriptionPlanData{
				Id:   plan.ID.String(),
				Type: resources.TypeSubscriptionPlan,
				Attributes: resources.SubscriptionPlanDataAttributes{
					Name:            plan.Name,
					Desc:            plan.Description,
					Price:           plan.Price,
					Currency:        plan.Currency,
					BillingInterval: int32(plan.BillingInterval),
					BillingCycle:    string(plan.BillingCycle),
					Status:          string(plan.Status),
					UpdatedAt:       plan.UpdatedAt,
					CreatedAt:       plan.CreatedAt,
				},
				Relationships: resources.SubscriptionPlanDataRelationships{
					SubscriptionType: resources.Relationships{
						Data: resources.RelationshipsData{
							Id:   subscription.SType.ID.String(),
							Type: resources.TypeSubscriptionType,
						},
					},
				},
			})
			if plan.TypeID == subscription.SType.ID {
				typeRelations = append(typeRelations, resources.Relationships{
					Data: resources.RelationshipsData{
						Id:   plan.ID.String(),
						Type: resources.TypeSubscriptionPlan,
					},
				})
			}
		}

		types = append(types, resources.SubscriptionTypeData{
			Id:   subscription.SType.ID.String(),
			Type: resources.TypeSubscriptionType,
			Attributes: resources.SubscriptionTypeDataAttributes{
				Name:      subscription.SType.Name,
				Desc:      subscription.SType.Description,
				Status:    string(subscription.SType.Status),
				UpdatedAt: subscription.SType.UpdatedAt,
				CreatedAt: subscription.SType.CreatedAt,
			},
			Relationships: resources.SubscriptionTypeDataRelationships{
				SubscriptionPlanRelation: resources.SubscriptionTypeDataRelationshipsSubscriptionPlanRelation{
					Data: typeRelations,
				},
			},
		})
	}

	return resources.SubscriptionTypeDepends{
		Data:     types,
		Included: plans,
	}
}
