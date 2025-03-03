package domain

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra"
	"github.com/sirupsen/logrus"
)

type Domain interface {
	CreateSubscriptionType(ctx context.Context, sub models.SubscriptionType) error
	CreateSubscriptionPlan(ctx context.Context, plan models.SubscriptionPlan) error

	UpdateSubscriptionType(ctx context.Context, update map[string]any) error
	DeleteSubscriptionType(ctx context.Context, id string) error
}

type domain struct {
	Infra *infra.Infra
	log   *logrus.Logger
}

func NewDomain(infra *infra.Infra, log *logrus.Logger) (Domain, error) {
	return &domain{
		Infra: infra,
		log:   log,
	}, nil
}

func (d *domain) CreateSubscriptionType(ctx context.Context, subType models.SubscriptionType) error {
	return d.Infra.Data.SubTypes.Create(ctx, subType)
}

func (d *domain) CreateSubscriptionPlan(ctx context.Context, plan models.SubscriptionPlan) error {
	return d.Infra.Data.SubPlans.Create(ctx, plan)
}

func (d *domain) UpdateSubscriptionType(ctx context.Context, update map[string]any) error {
	return d.Infra.Data.SubTypes.Update(ctx, update)
}
