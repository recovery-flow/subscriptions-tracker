package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type SubscriptionsCatalog interface {
	CreateType(ctx context.Context, name, desc string) (*models.SubscriptionType, error)
	CreatePlan(
		ctx context.Context,
		name, desc string,
		TypeID uuid.UUID,
		price float32,
		currency string,
		BillingInterval int8,
		BillingCycle string,
	) (*models.SubscriptionPlan, error)

	UpdateType(ctx context.Context, ID uuid.UUID, update map[string]any) error
	UpdatePlan(ctx context.Context, ID uuid.UUID, update map[string]any) error

	ActivateType(ctx context.Context, ID uuid.UUID) error
	ActivatePlan(ctx context.Context, ID uuid.UUID) error

	DeactivateType(ctx context.Context, ID uuid.UUID) error
	DeactivatePlan(ctx context.Context, ID uuid.UUID) error

	GetType(ctx context.Context, ID uuid.UUID) (*models.SubscriptionType, []models.SubscriptionPlan, error)
	GetPlan(ctx context.Context, ID uuid.UUID) (*models.SubscriptionPlan, error)
	GetAllType(ctx context.Context, statusType *models.StatusType) ([]models.SubscriptionTypeDepends, error)
}

func (d *domain) CreateType(ctx context.Context, name, desc string) (*models.SubscriptionType, error) {
	sType := &models.SubscriptionType{
		ID:          uuid.New(),
		Name:        name,
		Description: desc,
		Status:      models.StatusTypeInactive,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}

	err := d.Infra.Data.SQL.Types.New().Insert(ctx, sType)
	if err != nil {
		return nil, err
	}

	return sType, nil
}

func (d *domain) CreatePlan(
	ctx context.Context,
	name, desc string,
	TypeID uuid.UUID,
	price float32,
	currency string,
	BillingInterval int8,
	BillingCycle string,
) (*models.SubscriptionPlan, error) {
	cycle, err := models.ParseBillingCycle(BillingCycle)
	if err != nil {
		return nil, err
	}

	plan := &models.SubscriptionPlan{
		ID:              uuid.New(),
		TypeID:          TypeID,
		Price:           price,
		Name:            name,
		Description:     desc,
		BillingInterval: BillingInterval,
		BillingCycle:    cycle,
		Currency:        currency,
		Status:          models.StatusPlanInactive,
		UpdatedAt:       time.Now().UTC(),
		CreatedAt:       time.Now().UTC(),
	}

	_, err = d.Infra.Data.SQL.Types.New().Filter(map[string]any{
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

func (d *domain) UpdateType(ctx context.Context, ID uuid.UUID, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}

	err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{"id": ID.String()}).Update(ctx, update)
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) UpdatePlan(ctx context.Context, ID uuid.UUID, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}

	err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{"id": ID.String()}).Update(ctx, update)
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) ActivateType(ctx context.Context, ID uuid.UUID) error {
	err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusTypeActive})
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) ActivatePlan(ctx context.Context, ID uuid.UUID) error {
	curPlan, err := d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"id": ID.String()}).Get(ctx)
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

	err = d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusPlanInactive})
	if err != nil {
		return err
	}

	return nil
}

func (d *domain) DeactivateType(ctx context.Context, ID uuid.UUID) error {
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

func (d *domain) DeactivatePlan(ctx context.Context, ID uuid.UUID) error {
	if err := d.Infra.Data.SQL.Plans.New().Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": models.StatusPlanInactive})
		if err != nil {
			return err
		}
		err = d.Infra.Data.SQL.Subscriptions.New().Filter(map[string]any{"plan_id": ID.String()}).Update(ctx, map[string]any{"availability": models.SubPlanUnavailable})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (d *domain) GetType(ctx context.Context, ID uuid.UUID) (*models.SubscriptionType, []models.SubscriptionPlan, error) {
	sType, err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{"id": ID.String()}).Get(ctx)
	if err != nil {
		return nil, nil, err
	}

	plans, err := d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"type_id": ID.String()}).Select(ctx)
	if err != nil {
		return nil, nil, err
	}

	return sType, plans, nil
}

func (d *domain) GetPlan(ctx context.Context, ID uuid.UUID) (*models.SubscriptionPlan, error) {
	plans, err := d.Infra.Data.SQL.Plans.New().Filter(map[string]any{"id": ID.String()}).Select(ctx)
	if err != nil {
		return nil, err
	}

	return &plans[0], nil
}

func (d *domain) GetAllType(ctx context.Context, statusType *models.StatusType) ([]models.SubscriptionTypeDepends, error) {
	typesQ := d.Infra.Data.SQL.Types.New()
	if statusType != nil {
		typesQ.Filter(map[string]any{"status": *statusType})
	}
	types, err := typesQ.Select(ctx)
	if err != nil {
		return nil, err
	}

	typesDepends := make([]models.SubscriptionTypeDepends, len(types))

	for i, sType := range types {
		plansQ := d.Infra.Data.SQL.Plans.New()
		if statusType != nil {
			if *statusType == models.StatusTypeActive {
				plansQ.Filter(map[string]any{"status": models.StatusPlanActive})
			} else {
				plansQ.Filter(map[string]any{"status": models.StatusPlanInactive})
			}
		}

		plans, err := plansQ.Select(ctx)
		if err != nil {
			return nil, err
		}

		tmp := sType
		typesDepends[i] = models.SubscriptionTypeDepends{
			SType: &tmp,
			Plans: plans,
		}
	}

	return typesDepends, nil
}
