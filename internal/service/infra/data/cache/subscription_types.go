package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

const SubscriptionTypesCollection = "subscription_types"

type Types interface {
	Set(ctx context.Context, sType *models.SubscriptionType) error
	Get(ctx context.Context, ID string) (*models.SubscriptionType, error)

	Update(ctx context.Context, ID string, fields map[string]string) error

	Delete(ctx context.Context, ID string) error

	Drop(ctx context.Context) error
}

type types struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubTypesQueryCache(client *redis.Client, lifeTime time.Duration) Types {
	return &types{
		client:   client,
		lifeTime: lifeTime,
	}
}

func (t *types) Set(ctx context.Context, sub *models.SubscriptionType) error {
	subKey := fmt.Sprintf("%s:id:%s", SubscriptionTypesCollection, sub.ID.String())

	data := map[string]interface{}{
		"name":        sub.Name,
		"description": sub.Description,
		"status":      sub.Status,
		"updated_at":  sub.UpdatedAt.Format(time.RFC3339),
		"created_at":  sub.CreatedAt.Format(time.RFC3339),
	}

	if err := t.client.HSet(ctx, subKey, data).Err(); err != nil {
		return err
	}

	if t.lifeTime > 0 {
		pipe := t.client.Pipeline()
		pipe.Expire(ctx, subKey, t.lifeTime)
		_, err := pipe.Exec(ctx)
		if err != nil && err != redis.Nil {
			return err
		}
	}

	return nil
}

func (t *types) Update(ctx context.Context, ID string, fields map[string]string) error {
	allowed := map[string]bool{
		"name":        true,
		"description": true,
		"status":      true,
		"updated_at":  true,
		"created_at":  true,
	}

	updates := make(map[string]interface{})
	for key, value := range fields {
		if !allowed[key] {
			return fmt.Errorf("field %q cannot be updated", key)
		}
		updates[key] = value
	}

	subKey := fmt.Sprintf("%s:id:%s", SubscriptionTypesCollection, ID)

	if err := t.client.HSet(ctx, subKey, updates).Err(); err != nil {
		return fmt.Errorf("error updating %s in cache: %w", SubscriptionTypesCollection, err)
	}

	if t.lifeTime > 0 {
		if err := t.client.Expire(ctx, subKey, t.lifeTime).Err(); err != nil && err != redis.Nil {
			return fmt.Errorf("error updating %s for key %s: %w", SubscriptionTypesCollection, subKey, err)
		}
	}

	return nil
}

func (t *types) Get(ctx context.Context, ID string) (*models.SubscriptionType, error) {
	subKey := fmt.Sprintf("%s:id:%s", SubscriptionTypesCollection, ID)
	vals, err := t.client.HGetAll(ctx, subKey).Result()
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, redis.Nil
	}
	return parseType(ID, vals)
}

func (t *types) Delete(ctx context.Context, ID string) error {
	subKey := fmt.Sprintf("%s:id:%s", SubscriptionTypesCollection, ID)

	if err := t.client.Del(ctx, subKey).Err(); err != nil {
		return err
	}

	return nil
}

func (t *types) Drop(ctx context.Context) error {
	pattern := fmt.Sprintf("%s:*", SubscriptionTypesCollection)
	keys, err := t.client.Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("error fetching keys with pattern %s: %w", pattern, err)
	}
	if len(keys) == 0 {
		return nil
	}
	if err := t.client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("failed to delete keys with pattern %s: %w", pattern, err)
	}
	return nil
}

func parseType(userID string, vals map[string]string) (*models.SubscriptionType, error) {
	ID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("error parsing ID: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, vals["created_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing created_at: %w", err)
	}

	updatedAt, err := time.Parse(time.RFC3339, vals["updated_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing updated_at: %w", err)
	}

	status, err := models.ParseStatusType(vals["status"])
	if err != nil {
		return nil, fmt.Errorf("error parsing status: %w", err)
	}

	sub := models.SubscriptionType{
		ID:          ID,
		Name:        vals["name"],
		Description: vals["description"],
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	return &sub, nil
}
