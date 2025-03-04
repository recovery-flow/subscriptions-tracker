package domain

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events/evebody"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Domain interface {
	CreateSubType(ctx context.Context, sub models.SubscriptionType) error
	CreateSubPlan(ctx context.Context, plan models.SubscriptionPlan) error

	UpdateSubType(ctx context.Context, ID uuid.UUID, update map[string]any) error
	UpdateSubPlan(ctx context.Context, ID uuid.UUID, update map[string]any) error

	ActivateSubType(ctx context.Context, ID uuid.UUID) error
	ActivateSubPlan(ctx context.Context, ID uuid.UUID) error

	DeactivateSubType(ctx context.Context, ID uuid.UUID) error
	DeactivateSubPlan(ctx context.Context, ID uuid.UUID) error

	GetSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error)
	CreateSubscription(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
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
	//subType.Status = "inactive"

	err := d.Infra.Data.SQL.Types.Insert(ctx, subType)
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Types.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop subscription_types cache")
	}

	return nil
}

func (d *domain) CreateSubPlan(ctx context.Context, plan models.SubscriptionPlan) error {
	//plan.Status = "inactive"

	err := d.Infra.Data.SQL.Plans.Insert(ctx, plan)
	if err != nil {
		return err
	}

	err = d.Infra.Data.Cache.Plans.Drop(ctx)
	if err != nil {
		d.log.WithField("redis", err).Error("failed to drop subscription_plans cache")
	}

	return nil
}

func (d *domain) UpdateSubType(ctx context.Context, id uuid.UUID, update map[string]any) error {
	if update["status"] != nil {
		return fmt.Errorf("status field is not allowed to update at this method")
	}
	err := d.Infra.Data.SQL.Types.Filter(map[string]any{"id": id.String()}).Update(ctx, update)
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
	return d.Infra.Data.SQL.Types.Transaction(func(ctx context.Context) error {
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

		err = d.Infra.Data.Cache.Types.Drop(ctx)
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
	})
}

func (d *domain) DeactivateSubPlan(ctx context.Context, ID uuid.UUID) error {
	return d.Infra.Data.SQL.Plans.Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.Plans.Filter(map[string]any{"id": ID.String()}).Update(ctx, map[string]any{"status": "inactive"})
		if err != nil {
			return err
		}
		err = d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"plan_id": ID.String()}).Update(ctx, map[string]any{"availability": "deprecated"})
		if err != nil {
			return err
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
	})
}

func (d *domain) CreateSubscription(ctx context.Context, sub models.Subscription) (*models.Subscription, error) {
	err := d.Infra.Data.SQL.Subscriptions.Transaction(func(ctx context.Context) error {
		if err := d.Infra.Data.SQL.Subscriptions.New().Insert(ctx, sub); err != nil {
			return err
		}

		if err := d.Infra.Data.Cache.Subscriptions.Set(ctx, sub); err != nil {
			d.log.WithField("redis", err).Error("failed to set subscription to cache")
		}

		var typeID uuid.UUID

		plans, err := d.Infra.Data.Cache.Plans.Get(ctx, cache.KeyPlans(
			map[string]any{"id": sub.PlanID.String()}, 1, 1,
		))
		if err != nil && !errors.Is(err, redis.Nil) {
			d.log.WithField("redis", err).Error("failed to get subscription plan from cache")
		}

		if plans != nil && len(plans) == 1 {
			typeID = plans[0].TypeID
		} else {
			plans, err = d.Infra.Data.SQL.Plans.New().Filter(map[string]any{
				"id": sub.PlanID.String(),
			}).Select(ctx)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return fmt.Errorf("subscription plan not found")
				}
				return err
			}
			if len(plans) == 0 {
				return fmt.Errorf("subscription plan not found %s %s", sub.PlanID, err)
			}

			typeID = plans[0].TypeID

			if err := d.Infra.Data.Cache.Plans.Set(
				ctx,
				cache.KeyPlans(map[string]any{"id": sub.PlanID.String()}, 1, 1),
				plans,
			); err != nil {
				d.log.WithField("redis", err).Error("failed to set subscription plan to cache")
			}
		}

		if err := d.Infra.Kafka.SubscriptionCreated(evebody.CreateSubscription{
			UserID:    sub.UserID.String(),
			PlanID:    sub.PlanID.String(),
			TypeID:    typeID.String(),
			CreatedAt: sub.CreatedAt,
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (d *domain) GetSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error) {
	res, err := d.Infra.Data.Cache.Subscriptions.Get(ctx, userID.String())
	if err != nil || !errors.Is(err, redis.Nil) {
		d.log.WithError(err).Error("failed to get user subscription from cache")
	} else if res != nil {
		return res, nil
	}

	res, err = d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"user_id": userID.String()}).Get(ctx)
	if err != nil {
		return nil, err
	}

	err = d.Infra.Data.Cache.Subscriptions.Set(ctx, *res)
	if err != nil || !errors.Is(err, redis.Nil) {
		d.log.WithError(err).Error("failed to set user subscription to cache")
	}

	return res, nil
}
