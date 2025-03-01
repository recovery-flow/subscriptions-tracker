package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

type SubsTypes struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubsTypes(client *redis.Client, lifetime time.Duration) *SubsTypes {
	return &SubsTypes{
		client:   client,
		lifeTime: lifetime,
	}
}

func (s *SubsTypes) Add(ctx context.Context, subsType models.SubscriptionType) error {
	IDKey := fmt.Sprintf("subscription_type:id:%s", subsType.ID.String())

	data := map[string]interface{}{
		"name":        subsType.Name,
		"description": subsType.Description,
		"created_at":  subsType.CreatedAt.Format(time.RFC3339),
	}

	err := s.client.HSet(ctx, IDKey, data).Err()
	if err != nil {
		return fmt.Errorf("error adding subscription type to Redis: %w", err)
	}

	if s.lifeTime > 0 {
		pipe := s.client.Pipeline()
		keys := []string{IDKey}
		for _, key := range keys {
			pipe.Expire(ctx, key, s.lifeTime)
		}
		_, err := pipe.Exec(ctx)
		if err != nil && err != redis.Nil {
			return fmt.Errorf("error setting expiration for keys: %w", err)
		}
	}

	return nil
}

func (s *SubsTypes) GetByID(ctx context.Context, TypeID string) (*models.SubscriptionType, error) {
	key := fmt.Sprintf("subscription_type:id:%s", TypeID)
	vals, err := s.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting subscription type from Redis: %w", err)
	}

	return parseSubsType(TypeID, vals)
}

func (s *SubsTypes) Delete(ctx context.Context, TypeID string) error {
	key := fmt.Sprintf("subscription_type:id:%s", TypeID)
	err := s.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("error deleting subscription type from Redis: %w", err)
	}

	return nil
}

func parseSubsType(id string, vals map[string]string) (*models.SubscriptionType, error) {
	typeID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("error parsing subscription type ID: %w", err)
	}

	return &models.SubscriptionType{
		ID:          typeID,
		Name:        vals["name"],
		Description: vals["description"],
		CreatedAt:   time.Time{},
	}, nil
}
