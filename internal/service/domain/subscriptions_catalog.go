package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/cache"
	"github.com/redis/go-redis/v9"
)

type SubscriptionsCatalog interface {
	CreateSubType(ctx context.Context, sub models.SubscriptionType) error
	CreateSubPlan(ctx context.Context, plan models.SubscriptionPlan) error

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

func (d *domain) CreateSubType(ctx context.Context, subType models.SubscriptionType) error {
	subType.Status = "inactive"

	err := d.Infra.Data.SQL.Types.Insert(ctx, subType)
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Types.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop subscription_types cache")
	}

	err = d.Infra.Data.Cache.Types.Set(ctx, cache.KeyTypes(map[string]any{"id": subType.ID.String()}, 1, 1), []models.SubscriptionType{subType})

	return nil
}

func (d *domain) CreateSubPlan(ctx context.Context, plan models.SubscriptionPlan) error {
	plan.Status = "inactive"

	err := d.Infra.Data.SQL.Plans.Insert(ctx, plan)
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Plans.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop subscription_plans cache")
	}

	err = d.Infra.Data.Cache.Plans.Set(ctx, cache.KeyPlans(map[string]any{"id": plan.ID.String()}, 1, 1), []models.SubscriptionPlan{plan})

	return nil
}

func (d *domain) UpdateSubType(ctx context.Context, ID uuid.UUID, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}
	err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Update(ctx, update)
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Types.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop subscription_types cache")
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

	err = d.Infra.Data.Cache.Plans.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop subscription_plans cache")
	}

	return nil
}

func (d *domain) ActivateSubType(ctx context.Context, ID uuid.UUID) error {
	err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": "active"})
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Types.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache")
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

	err = d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": "active"})
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Plans.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache subscription_plans")
	}
	return nil
}

func (d *domain) DeactivateSubType(ctx context.Context, ID uuid.UUID) error {
	if err := d.Infra.Data.SQL.Types.Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": "inactive"})
		if err != nil {
			return err
		}

		plans, err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"type_id": ID.String()}).Select(ctx)
		if err != nil {
			return err
		}

		for _, plan := range plans {
			err = d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": plan.ID.String()}).Update(ctx, map[string]any{"status": "inactive"})
			if err != nil {
				return err
			}
			err = d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"plan_id": plan.ID.String()}).Update(ctx, map[string]any{"availability": "deprecated"})
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	err := d.Infra.Data.Cache.Types.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache subscription_types")
	}

	err = d.Infra.Data.Cache.Plans.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache subscription_plans")
	}

	err = d.Infra.Data.Cache.Subscriptions.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache subscriptions")
	}

	return nil
}

func (d *domain) DeactivateSubPlan(ctx context.Context, ID uuid.UUID) error {
	if err := d.Infra.Data.SQL.Plans.Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": "inactive"})
		if err != nil {
			return err
		}
		err = d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"plan_id": ID.String()}).Update(ctx, map[string]any{"availability": "deprecated"})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	err := d.Infra.Data.Cache.Plans.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache subscription_plans")
	}

	err = d.Infra.Data.Cache.Subscriptions.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop cache subscriptions")
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
	types, err := d.Infra.Data.Cache.Types.Get(ctx, cache.KeyTypes(map[string]any{"id": ID.String()}, 1, 1))
	if err != nil || !errors.Is(err, redis.Nil) {
		d.log.WithError(err).Error("failed to get subscription type from cache")
	}
	if types == nil || len(types) == 0 {
		types, err = d.Infra.Data.SQL.Types.Filter(map[string]any{"id": ID.String()}).Select(ctx)
		if err != nil {
			return nil, err
		}
		if err := d.Infra.Data.Cache.Types.Set(ctx, cache.KeyTypes(map[string]any{"id": ID.String()}, 1, 1), types); err != nil {
			d.log.WithError(err).Error("failed to set subscription type to cache")
		}
	}

	if len(types) == 0 {
		return nil, fmt.Errorf("subscription type not found")
	}

	return &types[0], nil
}

func (d *domain) GetSubPlan(ctx context.Context, ID uuid.UUID) (*models.SubscriptionPlan, error) {
	plans, err := d.Infra.Data.Cache.Plans.Get(ctx, cache.KeyPlans(map[string]any{"id": ID.String()}, 1, 10))
	if err != nil || !errors.Is(err, redis.Nil) {
		d.log.WithError(err).Error("failed to get subscription plans from cache")
	}
	if plans == nil || len(plans) == 0 {
		plans, err = d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Select(ctx)
		if err != nil {
			return nil, err
		}
		if err := d.Infra.Data.Cache.Plans.Set(ctx, cache.KeyPlans(map[string]any{"id": ID.String()}, 1, 10), plans); err != nil {
			d.log.WithError(err).Error("failed to set subscription plans to cache")
		}
	}

	if len(plans) == 0 {
		return nil, fmt.Errorf("subscription type not found")
	}

	return &plans[0], nil
}

func (d *domain) GetSubPlanByType(ctx context.Context, typeID uuid.UUID) ([]models.SubscriptionPlan, error) {
	plans, err := d.Infra.Data.Cache.Plans.Get(ctx, cache.KeyPlans(map[string]any{"type_id": typeID.String()}, 1, 10))
	if err != nil || !errors.Is(err, redis.Nil) {
		d.log.WithError(err).Error("failed to get subscription plans from cache")
	}
	if plans == nil || len(plans) == 0 {
		plans, err = d.Infra.Data.SQL.Plans.Filter(map[string]any{"type_id": typeID.String()}).Select(ctx)
		if err != nil {
			return nil, err
		}
		if err := d.Infra.Data.Cache.Plans.Set(ctx, cache.KeyPlans(map[string]any{"type_id": typeID.String()}, 1, 10), plans); err != nil {
			d.log.WithError(err).Error("failed to set subscription plans to cache")
		}
	}

	if len(plans) == 0 {
		return nil, fmt.Errorf("subscription type not found")
	}

	return plans, nil
}
