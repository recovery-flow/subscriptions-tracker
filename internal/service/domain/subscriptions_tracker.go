package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events/evebody"
	"github.com/redis/go-redis/v9"
)

type SubscriptionsTracker interface {
	GetSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error)
	CreateSubscription(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
}

func (d *domain) CreateSubscription(ctx context.Context, sub models.Subscription) (*models.Subscription, error) {
	err := d.Infra.Data.SQL.Subscriptions.Transaction(func(ctx context.Context) error {
		if err := d.Infra.Data.SQL.Subscriptions.New().Insert(ctx, sub); err != nil {
			return err
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

	if err := d.Infra.Data.Cache.Subscriptions.Set(ctx, sub); err != nil {
		d.log.WithField("redis", err).Error("failed to set subscription to cache")
	}

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
