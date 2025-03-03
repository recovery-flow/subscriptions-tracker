package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

const SubscriptionTypesCollection = "subscription_types"

type SubTypesQueryCache interface {
	Set(ctx context.Context, key string, types []models.SubscriptionType) error
	Get(ctx context.Context, key string) ([]models.SubscriptionType, error)
	Delete(ctx context.Context, key string) error
	Drop(ctx context.Context) error
}

type subTypesQueryCache struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubTypesQueryCache(client *redis.Client, lifeTime time.Duration) SubTypesQueryCache {
	return &subTypesQueryCache{
		client:   client,
		lifeTime: lifeTime,
	}
}

func (c *subTypesQueryCache) Set(ctx context.Context, key string, types []models.SubscriptionType) error {
	data, err := json.Marshal(types)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription types: %w", err)
	}
	if err := c.client.Set(ctx, key, data, c.lifeTime).Err(); err != nil {
		return fmt.Errorf("failed to set subscription types in cache: %w", err)
	}
	return nil
}

func (c *subTypesQueryCache) Get(ctx context.Context, key string) ([]models.SubscriptionType, error) {
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get subscription types from cache: %w", err)
	}
	var types []models.SubscriptionType
	if err := json.Unmarshal([]byte(data), &types); err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription types: %w", err)
	}
	return types, nil
}

func (c *subTypesQueryCache) Delete(ctx context.Context, key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete cache for key %s: %w", key, err)
	}
	return nil
}

func (c *subTypesQueryCache) Drop(ctx context.Context) error {
	pattern := fmt.Sprintf("%s:*", SubscriptionTypesCollection)
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
