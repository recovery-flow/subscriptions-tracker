package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func SubscriptionType(types models.SubscriptionType, plans []models.SubscriptionPlan) resources.SubscriptionType {
	res := resources.SubscriptionType{
		Data: resources.SubscriptionTypeData{
			Id:   types.ID.String(),
			Type: resources.TypeSubscriptionType,
			Attributes: resources.SubscriptionTypeDataAttributes{
				Name:      types.Name,
				Desc:      types.Description,
				Status:    string(types.Status),
				UpdatedAt: types.UpdatedAt,
				CreatedAt: types.CreatedAt,
			},
		},
	}

	if len(plans) > 0 {
		var data []resources.SubscriptionPlanData
		for _, plan := range plans {
			data = append(data, SubscriptionPlan(plan, &types).Data)
		}
		res.Included = data
	}

	return res
}
