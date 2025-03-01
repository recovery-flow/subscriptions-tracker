package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

type Subscriptions struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubscriptions(client *redis.Client, lifetime time.Duration) *Subscriptions {
	return &Subscriptions{
		client:   client,
		lifeTime: lifetime,
	}
}

func (s *Subscriptions) Add(ctx context.Context, sub models.Subscription) error {
	subKey := fmt.Sprintf("subscription:user_id:%s", sub.UserID.String())
	planIndexKey := fmt.Sprintf("subscription:plan_id:%s", sub.PlanID.String())
	payMethodIndexKey := fmt.Sprintf("subscription:payment_method_id:%s", sub.PaymentMethodID.String())

	data := map[string]interface{}{
		"plan_id":           sub.PlanID.String(),
		"payment_method_id": sub.PaymentMethodID.String(),
		"status":            sub.Status,
		"start_date":        sub.StartDate.Format(time.RFC3339),
		"end_date":          sub.EndDate.Format(time.RFC3339),
		"created_at":        sub.CreatedAt.Format(time.RFC3339),
		"updated_at":        sub.UpdatedAt.Format(time.RFC3339),
	}

	if err := s.client.HSet(ctx, subKey, data).Err(); err != nil {
		return fmt.Errorf("error adding subscription to Redis: %w", err)
	}

	if err := s.client.SAdd(ctx, planIndexKey, sub.UserID.String()).Err(); err != nil {
		return fmt.Errorf("error adding subscription UserID to plan index: %w", err)
	}
	if err := s.client.SAdd(ctx, payMethodIndexKey, sub.UserID.String()).Err(); err != nil {
		return fmt.Errorf("error adding subscription UserID to payment method index: %w", err)
	}

	if s.lifeTime > 0 {
		pipe := s.client.Pipeline()
		keys := []string{subKey, planIndexKey, payMethodIndexKey}
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

func (s *Subscriptions) GetByUserID(ctx context.Context, userID string) (*models.Subscription, error) {
	subKey := fmt.Sprintf("subscription:user_id:%s", userID)
	vals, err := s.client.HGetAll(ctx, subKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting subscription from Redis: %w", err)
	}
	if len(vals) == 0 {
		return nil, redis.Nil
	}
	return parseSubscription(userID, vals)
}

func (s *Subscriptions) GetByPlanID(ctx context.Context, planID string) ([]models.Subscription, error) {
	planIndexKey := fmt.Sprintf("subscription:plan_id:%s", planID)
	userIDs, err := s.client.SMembers(ctx, planIndexKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting subscription UserIDs by plan ID: %w", err)
	}

	var subs []models.Subscription
	for _, ID := range userIDs {
		vals, err := s.client.HGetAll(ctx, ID).Result()
		if err != nil {
			return nil, fmt.Errorf("error getting subscription: %w", err)
		}
		sub, err := parseSubscription(ID, vals)
		if err != nil {
			return nil, fmt.Errorf("error parsing subscription: %w", err)
		}
		subs = append(subs, *sub)
	}
	return subs, nil
}

func (s *Subscriptions) GetByPaymentMethodID(ctx context.Context, paymentMethodID string) ([]*models.Subscription, error) {
	payMethodIndexKey := fmt.Sprintf("subscription:payment_method_id:%s", paymentMethodID)
	userIDs, err := s.client.SMembers(ctx, payMethodIndexKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting subscription UserIDs by payment method ID: %w", err)
	}

	var subs []*models.Subscription
	for _, userID := range userIDs {
		vals, err := s.client.HGetAll(ctx, userID).Result()
		if err != nil {
			return nil, fmt.Errorf("error getting subscription for user %s: %w", userID, err)
		}
		sub, err := parseSubscription(userID, vals)
		if err != nil {
			return nil, fmt.Errorf("error parsing subscription: %w", err)
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

func (s *Subscriptions) Delete(ctx context.Context, userID string) error {
	subKey := fmt.Sprintf("subscription:user_id:%s", userID)

	sub, err := s.GetByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting subscription to delete: %w", err)
	}

	planIndexKey := fmt.Sprintf("subscription:plan_id:%s", sub.PlanID.String())
	payMethodIndexKey := fmt.Sprintf("subscription:payment_method_id:%s", sub.PaymentMethodID.String())

	if err := s.client.Del(ctx, subKey).Err(); err != nil {
		return fmt.Errorf("error deleting subscription key: %w", err)
	}

	if err := s.client.SRem(ctx, planIndexKey, userID).Err(); err != nil {
		return fmt.Errorf("error removing UserID from plan index: %w", err)
	}

	if err := s.client.SRem(ctx, payMethodIndexKey, userID).Err(); err != nil {
		return fmt.Errorf("error removing UserID from payment method index: %w", err)
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

	status, err := models.ParseSubscriptionStatus(vals["status"])
	if err != nil {
		return nil, fmt.Errorf("error parsing status: %w", err)
	}

	sub := models.Subscription{
		UserID:          uid,
		PlanID:          planID,
		PaymentMethodID: paymentMethodID,
		Status:          status,
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
	return &sub, nil
}
