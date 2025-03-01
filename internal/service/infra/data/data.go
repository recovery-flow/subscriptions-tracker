package data

import "github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo"

type Data struct {
	BillingPlan    repo.BillingPlan
	Transactions   repo.Transactions
	PaymentMethods repo.PaymentMethods
}
