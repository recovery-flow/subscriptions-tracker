package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

const SubscriptionCollection = "subscription"

type Subscriptions interface {
	Add(ctx context.Context, sub models.Subscription) error
	Get(ctx context.Context, userID string) (*models.Subscription, error)
	Delete(ctx context.Context, userID string) error
	Drop(ctx context.Context) error
}

type subscriptions struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubscriptions(client *redis.Client, lifetime time.Duration) Subscriptions {
	return &subscriptions{
		client:   client,
		lifeTime: lifetime,
	}
}

func (s *subscriptions) Add(ctx context.Context, sub models.Subscription) error {
	subKey := fmt.Sprintf("%s:user_id:%s", SubscriptionCollection, sub.UserID.String())

	data := map[string]interface{}{
		"plan_id":           sub.PlanID.String(),
		"payment_method_id": sub.PaymentMethodID.String(),
		"status":            sub.State,
		"start_date":        sub.StartDate.Format(time.RFC3339),
		"end_date":          sub.EndDate.Format(time.RFC3339),
		"created_at":        sub.CreatedAt.Format(time.RFC3339),
		"updated_at":        sub.UpdatedAt.Format(time.RFC3339),
	}

	if err := s.client.HSet(ctx, subKey, data).Err(); err != nil {
		return err
	}

	if s.lifeTime > 0 {
		pipe := s.client.Pipeline()
		pipe.Expire(ctx, subKey, s.lifeTime)
		_, err := pipe.Exec(ctx)
		if err != nil && err != redis.Nil {
			return err
		}
	}

	return nil
}

func (s *subscriptions) Get(ctx context.Context, userID string) (*models.Subscription, error) {
	subKey := fmt.Sprintf("%s:user_id:%s", SubscriptionCollection, userID)
	vals, err := s.client.HGetAll(ctx, subKey).Result()
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, redis.Nil
	}
	return parseSubscription(userID, vals)
}

func (s *subscriptions) Delete(ctx context.Context, userID string) error {
	subKey := fmt.Sprintf("%s:user_id:%s", SubscriptionCollection, userID)

	if err := s.client.Del(ctx, subKey).Err(); err != nil {
		return err
	}

	return nil
}

func (s *subscriptions) Drop(ctx context.Context) error {
	pattern := fmt.Sprintf("%s:*", SubscriptionPlanCollection)
	keys, err := s.client.Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("error fetching keys with pattern %s: %w", pattern, err)
	}
	if len(keys) == 0 {
		return nil
	}
	if err := s.client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("failed to delete keys with pattern %s: %w", pattern, err)
	}
	return nil
}

func parseSubscription(userID string, vals map[string]string) (*models.Subscription, error) {
	planID, err := uuid.Parse(vals["plan_id"])
	if err != nil {
		return nil, fmt.Errorf("error parsing plan_id: %w", err)
	}

	paymentMethodID, err := uuid.Parse(vals["payment_method_id"])
	if err != nil {
		return nil, fmt.Errorf("error parsing payment_method_id: %w", err)
	}

	startDate, err := time.Parse(time.RFC3339, vals["start_date"])
	if err != nil {
		return nil, fmt.Errorf("error parsing start_date: %w", err)
	}

	endDate, err := time.Parse(time.RFC3339, vals["end_date"])
	if err != nil {
		return nil, fmt.Errorf("error parsing end_date: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, vals["created_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing created_at: %w", err)
	}

	updatedAt, err := time.Parse(time.RFC3339, vals["updated_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing updated_at: %w", err)
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("error parsing user_id: %w", err)
	}

	status, err := models.ParseSubscriptionState(vals["status"])
	if err != nil {
		return nil, fmt.Errorf("error parsing status: %w", err)
	}

	sub := models.Subscription{
		UserID:          uid,
		PlanID:          planID,
		PaymentMethodID: paymentMethodID,
		State:           status,
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
	return &sub, nil
}
