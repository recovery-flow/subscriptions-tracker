package data

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/repositories"
)

type Data struct {
	Subscribers  repositories.Subscribers
	Plans        repositories.SubscriptionPlans
	Transactions repositories.Transactions
}

func NewDataBase(cfg config.Config) (*Data, error) {
	subs, err := repositories.NewSubscribers(cfg)
	if err != nil {
		return nil, err
	}
	plans, err := repositories.NewSubscriptionPlans(cfg)
	if err != nil {
		return nil, err
	}
	transactions, err := repositories.NewTransactions(cfg)
	if err != nil {
		return nil, err
	}

	return &Data{
		Subscribers:  subs,
		Plans:        plans,
		Transactions: transactions,
	}, nil
}
