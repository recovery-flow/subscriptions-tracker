package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func PaymentMethod(paymentMethod models.PaymentMethod) resources.PaymentMethod {
	return resources.PaymentMethod{
		Data: resources.PaymentMethodData{
			Id:   paymentMethod.ID.String(),
			Type: resources.TypePaymentMethod,
			Attributes: resources.PaymentMethodDataAttributes{
				UserId:        paymentMethod.UserID.String(),
				Type:          string(paymentMethod.Type),
				ProviderToken: paymentMethod.ProviderToken,
				IsDefault:     paymentMethod.IsDefault,
				CreatedAt:     paymentMethod.CreatedAt,
			},
		},
	}
}
