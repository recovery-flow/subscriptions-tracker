package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

const SubscriptionPlanCollection = "subscription_plans"

type Plans interface {
	Set(ctx context.Context, plan models.SubscriptionPlan) error
	Get(ctx context.Context, ID string) (*models.SubscriptionPlan, error)
	GetByType(ctx context.Context, typeID string) ([]models.SubscriptionPlan, error)

	Update(ctx context.Context, ID string, fields map[string]string) error

	Delete(ctx context.Context, ID string) error
	DeleteByType(ctx context.Context, typeID string) error

	Drop(ctx context.Context) error
}

type plans struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewPlansCache(client *redis.Client, lifeTime time.Duration) Plans {
	return &plans{
		client:   client,
		lifeTime: lifeTime,
	}
}

func (p *plans) Set(ctx context.Context, plan models.SubscriptionPlan) error {
	planKey := fmt.Sprintf("%s:id:%s", SubscriptionPlanCollection, plan.ID.String())

	data := map[string]interface{}{
		"type_id":               plan.TypeID.String(),
		"price":                 plan.Price,
		"name":                  plan.Name,
		"description":           plan.Description,
		"billing_interval":      plan.BillingInterval,
		"billing_interval_unit": plan.BillingCycle,
		"currency":              plan.Currency,
		"status":                plan.Status,
		"created_at":            plan.CreatedAt.Format(time.RFC3339),
		"updated_at":            plan.UpdatedAt.Format(time.RFC3339),
	}

	if err := p.client.HSet(ctx, planKey, data).Err(); err != nil {
		return err
	}

	typeKey := fmt.Sprintf("%s:type:%s", SubscriptionPlanCollection, plan.TypeID.String())
	return p.client.SAdd(ctx, typeKey, plan.ID.String()).Err()
}

func (p *plans) Get(ctx context.Context, ID string) (*models.SubscriptionPlan, error) {
	planKey := fmt.Sprintf("%s:id:%s", SubscriptionPlanCollection, ID)

	vals, err := p.client.HGetAll(ctx, planKey).Result()
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, redis.Nil
	}

	return parsePlan(ID, vals)
}

func (p *plans) GetByType(ctx context.Context, typeID string) ([]models.SubscriptionPlan, error) {
	typeKey := fmt.Sprintf("%s:type:%s", SubscriptionPlanCollection, typeID)
	planIDs, err := p.client.SMembers(ctx, typeKey).Result()
	if err != nil {
		return nil, err
	}

	var plans []models.SubscriptionPlan
	for _, planID := range planIDs {
		plan, err := p.Get(ctx, planID)
		if err == redis.Nil {
			continue
		} else if err != nil {
			return nil, err
		}
		plans = append(plans, *plan)
	}

	return plans, nil
}

func (p *plans) Update(ctx context.Context, ID string, fields map[string]string) error {
	allowed := map[string]bool{
		"price":                 true,
		"name":                  true,
		"description":           true,
		"billing_interval":      true,
		"billing_interval_unit": true,
		"currency":              true,
		"status":                true,
		"updated_at":            true,
	}

	updates := make(map[string]interface{})
	for key, value := range fields {
		if !allowed[key] {
			return fmt.Errorf("field %q cannot be updated", key)
		}
		updates[key] = value
	}

	planKey := fmt.Sprintf("%s:id:%s", SubscriptionPlanCollection, ID)

	if err := p.client.HSet(ctx, planKey, updates).Err(); err != nil {
		return fmt.Errorf("error updating %s in cache: %w", SubscriptionPlanCollection, err)
	}

	if p.lifeTime > 0 {
		if err := p.client.Expire(ctx, planKey, p.lifeTime).Err(); err != nil && err != redis.Nil {
			return err
		}
	}

	return nil
}

func (p *plans) Delete(ctx context.Context, ID string) error {
	planKey := fmt.Sprintf("%s:id:%s", SubscriptionPlanCollection, ID)
	return p.client.Del(ctx, planKey).Err()
}

func (p *plans) DeleteByType(ctx context.Context, typeID string) error {
	plans, err := p.GetByType(ctx, typeID)
	if err != nil {
		return err
	}

	for _, plan := range plans {
		if err := p.Delete(ctx, plan.ID.String()); err != nil {
			return err
		}
	}
	return nil
}

func (p *plans) Drop(ctx context.Context) error {
	pattern := fmt.Sprintf("%s:*", SubscriptionPlanCollection)
	keys, err := p.client.Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("error fetching keys with pattern %s: %w", pattern, err)
	}
	if len(keys) == 0 {
		return nil
	}
	return p.client.Del(ctx, keys...).Err()
}

func parsePlan(ID string, vals map[string]string) (*models.SubscriptionPlan, error) {
	id, err := uuid.Parse(ID)
	if err != nil {
		return nil, err
	}

	typeID, err := uuid.Parse(vals["type_id"])
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(vals["price"], 64)
	if err != nil {
		return nil, err
	}

	billingInterval, err := strconv.ParseInt(vals["billing_interval"], 10, 8)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, vals["created_at"])
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, vals["updated_at"])
	if err != nil {
		return nil, err
	}

	status, err := models.ParseStatusPlan(vals["status"])
	if err != nil {
		return nil, err
	}

	plan := models.SubscriptionPlan{
		ID:              id,
		TypeID:          typeID,
		Price:           float32(price),
		Name:            vals["name"],
		Description:     vals["description"],
		BillingInterval: int8(billingInterval),
		BillingCycle:    models.BillingCycle(vals["billing_interval_unit"]),
		Currency:        vals["currency"],
		Status:          status,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}

	return &plan, nil
}
