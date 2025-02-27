package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

type SubPlans struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewSubPlans(client *redis.Client, lifetime time.Duration) *SubPlans {
	return &SubPlans{
		client:   client,
		lifeTime: lifetime,
	}
}

func (p *SubPlans) Add(ctx context.Context, subPlan models.SubscriptionPlan) error {
	subPlanKey := fmt.Sprintf("subscription_plan:id:%s", subPlan.ID.String())
	subTypeKey := fmt.Sprintf("subscription_plan:type_id:%s", subPlan.TypeID.String())

	data := map[string]interface{}{
		"type_id":               subPlan.TypeID.String(),
		"price":                 subPlan.Price,
		"billing_interval":      subPlan.BillingInterval,
		"billing_interval_unit": subPlan.BillingIntervalUnit,
		"currency":              subPlan.Currency,
		"created_at":            subPlan.CreatedAt.Format(time.RFC3339),
	}

	if err := p.client.HSet(ctx, subPlanKey, data).Err(); err != nil {
		return fmt.Errorf("error adding subscription plan to Redis: %w", err)
	}

	if err := p.client.SAdd(ctx, subTypeKey, subPlan.ID.String()).Err(); err != nil {
		return fmt.Errorf("error adding subscription plan ID to set: %w", err)
	}

	if p.lifeTime > 0 {
		pipe := p.client.Pipeline()
		keys := []string{subPlanKey, subTypeKey}
		for _, key := range keys {
			pipe.Expire(ctx, key, p.lifeTime)
		}
		_, err := pipe.Exec(ctx)
		if err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("error setting expiration for keys: %w", err)
		}
	}
	return nil
}

func (p *SubPlans) GetByID(ctx context.Context, planID string) (*models.SubscriptionPlan, error) {
	IDKey := fmt.Sprintf("subscription_plan:id:%s", planID)
	vals, err := p.client.HGetAll(ctx, IDKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting subscription plan from Redis: %w", err)
	}

	return parseSubPlan(planID, vals)
}

func (p *SubPlans) GetByTypeID(ctx context.Context, TypeID string) ([]models.SubscriptionPlan, error) {
	TypeKey := fmt.Sprintf("subscription_plan:type_id:%s", TypeID)
	IDs, err := p.client.SMembers(ctx, TypeKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting subscription plan from Redis: %w", err)
	}

	if len(IDs) == 0 {
		return nil, nil
	}

	var subPlans []models.SubscriptionPlan
	for _, ID := range IDs {
		vals, err := p.client.HGetAll(ctx, ID).Result()
		if err != nil {
			return nil, fmt.Errorf("error getting subscription plan: %w", err)
		}
		subPlan, err := parseSubPlan(ID, vals)
		if err != nil {
			return nil, fmt.Errorf("error parsing subscription plan: %w", err)
		}
		subPlans = append(subPlans, *subPlan)
	}

	return subPlans, nil
}

func (p *SubPlans) Delete(ctx context.Context, planID string) error {
	IDKey := fmt.Sprintf("subscription_plan:id:%s", planID)

	subPlan, err := p.GetByID(ctx, planID)
	if err != nil {
		return err
	}

	TypeKey := fmt.Sprintf("subscription_plan:type_id:%s", subPlan.TypeID)

	exists, err := p.client.Exists(ctx, IDKey).Result()
	if err != nil {
		return fmt.Errorf("error checking account existence in Redis: %w", err)
	}

	if exists == 0 {
		return redis.Nil
	}

	err = p.client.Del(ctx, IDKey).Err()
	if err != nil {
		return fmt.Errorf("error deleting subscription plan: %w", err)
	}

	err = p.client.SRem(ctx, TypeKey, planID).Err()
	if err != nil {
		return fmt.Errorf("error deleting subscription plan planID from type set: %w", err)
	}

	return nil
}

func parseSubPlan(ID string, vals map[string]string) (*models.SubscriptionPlan, error) {
	createdAt, err := time.Parse(time.RFC3339, vals["created_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing created_at: %w", err)
	}

	planID, err := uuid.Parse(ID)
	if err != nil {
		return nil, fmt.Errorf("error parsing ID: %w", err)
	}

	typeID, err := uuid.Parse(vals["type_id"])
	if err != nil {
		return nil, fmt.Errorf("error parsing type_id: %w", err)
	}

	price, err := strconv.ParseFloat(vals["price"], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing price: %w", err)
	}

	billingInterval64, err := strconv.ParseInt(vals["billing_interval"], 10, 8) // Основание 10, битность 8
	if err != nil {
		return nil, fmt.Errorf("error parsing billing_interval: %w", err)
	}

	billingIntervalUnit, err := models.ParseBillingIntervalUnit(vals["billing_interval_unit"])
	if err != nil {
		return nil, fmt.Errorf("error parsing billing_interval_unit, %s", err)
	}

	subPlan := models.SubscriptionPlan{
		ID:                  planID,
		TypeID:              typeID,
		Price:               price,
		BillingInterval:     int8(billingInterval64),
		BillingIntervalUnit: billingIntervalUnit,
		Currency:            vals["currency"],
		CreatedAt:           createdAt,
	}

	return &subPlan, nil
}
