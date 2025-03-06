package responses

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func Transaction(transaction models.Transaction) resources.Transaction {
	return resources.Transaction{
		Data: resources.TransactionData{
			Id:   transaction.ID.String(),
			Type: resources.TypeTransaction,
			Attributes: resources.TransactionDataAttributes{
				UserId:                transaction.UserID.String(),
				PaymentId:             transaction.PaymentProviderID,
				Amount:                transaction.Amount,
				Currency:              transaction.Currency,
				Status:                string(transaction.Status),
				ProviderTransactionId: transaction.PaymentProviderID,
				TransactionDate:       transaction.TransactionDate,
			},
		},
	}
}
