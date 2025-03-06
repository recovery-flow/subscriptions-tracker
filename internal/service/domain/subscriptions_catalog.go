package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type SubscriptionsCatalog interface {
	CreateSubType(ctx context.Context, name, desc string) (*models.SubscriptionType, error)
	CreateSubPlan(
		ctx context.Context,
		name, desc string,
		TypeID uuid.UUID,
		price float64,
		currency string,
		BillingInterval int8,
		BillingIntervalUnit models.BillingCycle,
	) (*models.SubscriptionPlan, error)

	UpdateSubType(ctx context.Context, ID uuid.UUID, update map[string]any) error
	UpdateSubPlan(ctx context.Context, ID uuid.UUID, update map[string]any) error

	ActivateSubType(ctx context.Context, ID uuid.UUID) error
	ActivateSubPlan(ctx context.Context, ID uuid.UUID) error

	DeactivateSubType(ctx context.Context, ID uuid.UUID) error
	DeactivateSubPlan(ctx context.Context, ID uuid.UUID) error

	GetSubTypes(ctx context.Context, ID uuid.UUID) (*models.SubscriptionType, []models.SubscriptionPlan, error)
	GetSubType(ctx context.Context, ID uuid.UUID) (*models.SubscriptionType, error)
	GetSubPlan(ctx context.Context, ID uuid.UUID) (*models.SubscriptionPlan, error)
	GetSubPlanByType(ctx context.Context, typeID uuid.UUID) ([]models.SubscriptionPlan, error)
}

func (d *domain) CreateSubType(ctx context.Context, name, desc string) (*models.SubscriptionType, error) {
	sType := &models.SubscriptionType{
		ID:          uuid.New(),
		Name:        name,
		Description: desc,
		Status:      models.StatusTypeInactive,
	}

	err := d.Infra.Data.SQL.Types.New().Insert(ctx, sType)
	if err != nil {
		return nil, err
	}

	return sType, nil
}

func (d *domain) CreateSubPlan(
	ctx context.Context,
	name, desc string,
	TypeID uuid.UUID,
	price float64,
	currency string,
	BillingInterval int8,
	BillingIntervalUnit models.BillingCycle,
) (*models.SubscriptionPlan, error) {

	plan := &models.SubscriptionPlan{
		ID:              uuid.New(),
		TypeID:          TypeID,
		Price:           price,
		Name:            name,
		Description:     desc,
		BillingInterval: BillingInterval,
		BillingCycle:    BillingIntervalUnit,
		Currency:        currency,
		Status:          models.StatusPlanInactive,
		UpdatedAt:       time.Now().UTC(),
		CreatedAt:       time.Now().UTC(),
	}

	_, err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{
		"id": TypeID.String(),
	}).Get(ctx)
	if err != nil {
		return nil, err
	}

	err = d.Infra.Data.SQL.Plans.Insert(ctx, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (d *domain) UpdateSubType(ctx context.Context, ID uuid.UUID, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}

	err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Update(ctx, update)
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) UpdateSubPlan(ctx context.Context, ID uuid.UUID, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}

	err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Update(ctx, update)
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) ActivateSubType(ctx context.Context, ID uuid.UUID) error {
	err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusTypeActive})
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) ActivateSubPlan(ctx context.Context, ID uuid.UUID) error {
	curPlan, err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Get(ctx)
	if err != nil {
		d.log.WithError(err).Error("failed to filter subscription plans")
		return err
	}

	subType, err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": curPlan.TypeID.String()}).Get(ctx)
	if err != nil {
		d.log.WithError(err).Error("failed to filter subscription types")
		return err
	}

	if subType.Status != "active" {
		return fmt.Errorf("subscription type maust be active to activate a subscription plan")
	}

	err = d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusPlanInactive})
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) DeactivateSubType(ctx context.Context, ID uuid.UUID) error {
	if err := d.Infra.Data.SQL.Types.New().Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusTypeInactive})
		if err != nil {
			return err
		}

		plans, err := d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"type_id": ID.String()}).Select(ctx)
		if err != nil {
			return err
		}

		for _, plan := range plans {
			err = d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"id": plan.ID.String()}).Update(ctx, map[string]any{"status": models.StatusPlanInactive})
			if err != nil {
				return err
			}
			err = d.Infra.Data.SQL.Subscriptions.New().Filter(map[string]any{"plan_id": plan.ID.String()}).Update(ctx, map[string]any{"availability": models.SubPlanUnavailable})
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (d *domain) DeactivateSubPlan(ctx context.Context, ID uuid.UUID) error {
	if err := d.Infra.Data.SQL.Plans.Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusPlanInactive})
		if err != nil {
			return err
		}
		err = d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"plan_id": ID.String()}).Update(ctx, map[string]any{"availability": models.SubPlanUnavailable})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (d *domain) GetSubTypes(ctx context.Context, ID uuid.UUID) (*models.SubscriptionType, []models.SubscriptionPlan, error) {
	subType, err := d.GetSubType(ctx, ID)
	if err != nil {
		return nil, nil, err
	}

	plans, err := d.GetSubPlanByType(ctx, ID)
	if err != nil {
		return nil, nil, err
	}

	return subType, plans, nil
}

func (d *domain) GetSubType(ctx context.Context, ID uuid.UUID) (*models.SubscriptionType, error) {
	sType, err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Get(ctx)
	if err != nil {
		return nil, err
	}

	return sType, nil
}

func (d *domain) GetSubPlan(ctx context.Context, ID uuid.UUID) (*models.SubscriptionPlan, error) {
	plans, err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Select(ctx)
	if err != nil {
		return nil, err
	}

	return &plans[0], nil
}

func (d *domain) GetSubPlanByType(ctx context.Context, typeID uuid.UUID) ([]models.SubscriptionPlan, error) {
	plans, err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"type_id": typeID.String()}).Select(ctx)
	if err != nil {
		return nil, err
	}

	return plans, nil
}
