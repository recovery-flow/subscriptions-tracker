package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func SubscriptionType(sType *models.SubscriptionType, plans []models.SubscriptionPlan) resources.SubscriptionType {
	res := resources.SubscriptionType{
		Data: resources.SubscriptionTypeData{
			Id:   sType.ID.String(),
			Type: resources.TypeSubscriptionType,
			Attributes: resources.SubscriptionTypeDataAttributes{
				Name:      sType.Name,
				Desc:      sType.Description,
				Status:    string(sType.Status),
				UpdatedAt: sType.UpdatedAt,
				CreatedAt: sType.CreatedAt,
			},
		},
	}

	if len(plans) > 0 {
		relationships := make([]resources.Relationships, 0)
		for _, plan := range plans {
			relationships = append(relationships, resources.Relationships{
				Data: resources.RelationshipsData{
					Id:   plan.ID.String(),
					Type: resources.TypeSubscriptionPlan,
				},
			})
		}
		res.Data.Relationships = resources.SubscriptionTypeDataRelationships{
			SubscriptionPlanRelation: resources.SubscriptionTypeDataRelationshipsSubscriptionPlanRelation{
				Data: relationships,
			},
		}
	}

	return res
}
