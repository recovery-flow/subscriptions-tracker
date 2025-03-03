package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra"
	"github.com/sirupsen/logrus"
)

type Domain interface {
	CreateSubType(ctx context.Context, sub models.SubscriptionType) error
	CreateSubPlan(ctx context.Context, plan models.SubscriptionPlan) error

	UpdateSubType(ctx context.Context, update map[string]any) error
	UpdateSubPlan(ctx context.Context, update map[string]any) error

	ActivateSubType(ctx context.Context, id uuid.UUID) error
	ActivateSubPlan(ctx context.Context, id uuid.UUID) error
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

func (d *domain) CreateSubType(ctx context.Context, subType models.SubscriptionType) error {
	subType.Status = "inactive"
	return d.Infra.Data.SubTypes.Create(ctx, subType)
}

func (d *domain) CreateSubPlan(ctx context.Context, plan models.SubscriptionPlan) error {
	plan.Status = "inactive"
	return d.Infra.Data.SubPlans.Create(ctx, plan)
}

func (d *domain) UpdateSubType(ctx context.Context, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}
	return d.Infra.Data.SubTypes.Update(ctx, update)
}

func (d *domain) UpdateSubPlan(ctx context.Context, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}
	return d.Infra.Data.SubPlans.Update(ctx, update)
}

func (d *domain) ActivateSubType(ctx context.Context, id uuid.UUID) error {
	err := d.Infra.Data.SubTypes.Update(ctx, map[string]any{"id": id.String(), "status": "active"})
	if err != nil {
		return err
	}

	err = d.Infra.Data.SubPlans.DropCache(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache")
	}

	return nil
}

func (d *domain) ActivateSubPlan(ctx context.Context, id uuid.UUID) error {
	curPlan, err := d.Infra.Data.SubPlans.Filter(map[string]any{"id": id.String()}).Get(ctx)
	if err != nil {
		d.log.WithError(err).Error("failed to get subscription plan")
		return err
	}

	subType, err := d.Infra.Data.SubTypes.Filter(map[string]any{"id": curPlan.TypeID.String()}).Get(ctx)
	if err != nil {
		d.log.WithError(err).Error("failed to get subscription type")
		return err
	}

	if subType.Status != "active" {
		return fmt.Errorf("subscription type maust be active to activate a subscription plan")
	}

	return d.Infra.Data.SubPlans.Update(ctx, map[string]any{"id": id, "status": "active"})
}

func (d *domain) DeactivateSubType(ctx context.Context, id uuid.UUID) error {
	plans, err := d.Infra.Data.SubPlans.Filter(map[string]any{"type_id": id.String()}).Select(ctx)
	if err != nil {
		d.log.WithError(err).Error("failed to get subscription plans")
		return err
	}

	err := d.Infra.Data.SubTypes.Update(ctx, map[string]any{"id": id.String(), "status": "inactive"})
	if err != nil {
		return err
	}

	err = d.Infra.Data.SubPlans.DropCache(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache")
	}

	return nil
}
