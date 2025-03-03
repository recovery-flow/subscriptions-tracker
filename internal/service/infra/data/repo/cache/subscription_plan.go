package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

const SubscriptionPlanCollection = "subscription_plans"

type SubPlanQueryCache interface {
	Set(ctx context.Context, key string, plans []models.SubscriptionPlan) error
	Get(ctx context.Context, key string) ([]models.SubscriptionPlan, error)
	Delete(ctx context.Context, key string) error
	Drop(ctx context.Context) error
}

type subPlanQueryCache struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubPlanQueryCache(client *redis.Client, lifeTime time.Duration) SubPlanQueryCache {
	return &subPlanQueryCache{
		client:   client,
		lifeTime: lifeTime,
	}
}

func (c *subPlanQueryCache) Set(ctx context.Context, key string, plans []models.SubscriptionPlan) error {
	data, err := json.Marshal(plans)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription plans: %w", err)
	}
	if err := c.client.Set(ctx, key, data, c.lifeTime).Err(); err != nil {
		return fmt.Errorf("failed to set subscription plans in cache: %w", err)
	}

	if c.lifeTime > 0 {
		pipe := c.client.Pipeline()
		pipe.Expire(ctx, key, c.lifeTime)
		_, err := pipe.Exec(ctx)
		if err != nil && err != redis.Nil {
			return err
		}
	}

	return nil
}

func (c *subPlanQueryCache) Get(ctx context.Context, key string) ([]models.SubscriptionPlan, error) {
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get subscription plans from cache: %w", err)
	}
	var plans []models.SubscriptionPlan
	if err := json.Unmarshal([]byte(data), &plans); err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription plans: %w", err)
	}
	return plans, nil
}

func (c *subPlanQueryCache) Delete(ctx context.Context, key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete cache for key %s: %w", key, err)
	}
	return nil
}

func (c *subPlanQueryCache) Drop(ctx context.Context) error {
	pattern := fmt.Sprintf("%s:*", SubscriptionPlanCollection)
	keys, err := c.client.Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("error fetching keys with pattern %s: %w", pattern, err)
	}
	if len(keys) == 0 {
		return nil
	}
	if err := c.client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("failed to delete keys with pattern %s: %w", pattern, err)
	}
	return nil
}
