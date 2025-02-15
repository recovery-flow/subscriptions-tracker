package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

type Subscribers interface {
	Add(ctx context.Context, account models.Subscriber) error
}

type subscriber struct {
	client *redis.Client
	ttl    time.Duration
}

func NewSubscribers(client *redis.Client) Subscribers {
	return &subscriber{
		client: client,
	}
}

func (s *subscriber) Add(ctx context.Context, subscriber models.Subscriber) error {
	subKey := fmt.Sprintf("sub:id:%s", subscriber.ID)
	userKey := fmt.Sprintf("sub:user_id:%s", subscriber.UserID)

	//todo some pole can be nil => check
	data := map[string]interface{}{
		"user_id":    subscriber.UserID.String(),
		"plan_id":    subscriber.PlanID.String(),
		"status":     subscriber.Status,
		"start_at":   subscriber.StartAt.Time(),
		"end_at":     subscriber.EndAt.Time(),
		"updated_at": subscriber.UpdatedAt.Time(),
		"created_at": subscriber.CreatedAt.Time(),
	}

	err := s.client.HSet(ctx, subKey, data).Err()
	if err != nil {
		return fmt.Errorf("error adding subscriber to Redis: %w", err)
	}

	err = s.client.Set(ctx, userKey, subscriber.ID.String(), 0).Err()
	if err != nil {
		return fmt.Errorf("error creating user index: %w", err)
	}

	if s.ttl > 0 {
		_ = s.client.Expire(ctx, subKey, s.ttl).Err()
		_ = s.client.Expire(ctx, userKey, s.ttl).Err()
	}

	return nil
}
