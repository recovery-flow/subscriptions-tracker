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

type SubPlans struct {
	client *redis.Client
}

func NewSubPlans(client *redis.Client) *SubPlans {
	return &SubPlans{
		client: client,
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

func (p *SubPlans) DeleteByID(ctx context.Context, planID string) error {
	IDKey := fmt.Sprintf("subscription_plan:id:%s", planID)

	subPlan, err := p.GetByID(ctx, planID)
	if err != nil {
		return err
	}
	if subPlan == nil {
		return redis.Nil
	}

	TypeKey := fmt.Sprintf("subscription_plan:type_id:%s", subPlan.TypeID.String())

	exists, err := p.client.Exists(ctx, IDKey).Result()
	if err != nil {
		return fmt.Errorf("error checking existence in Redis: %w", err)
	}
	if exists == 0 {
		return redis.Nil
	}

	if err := p.client.Del(ctx, IDKey).Err(); err != nil {
		return fmt.Errorf("error deleting subscription plan by ID: %w", err)
	}

	if err := p.client.SRem(ctx, TypeKey, planID).Err(); err != nil {
		return fmt.Errorf("error removing planID from type set: %w", err)
	}

	return nil
}

func (p *SubPlans) DeleteByType(ctx context.Context, typeID string) error {
	TypeKey := fmt.Sprintf("subscription_plan:type_id:%s", typeID)

	planIDs, err := p.client.SMembers(ctx, TypeKey).Result()
	if err != nil {
		return fmt.Errorf("error getting plan IDs from type set: %w", err)
	}

	if len(planIDs) == 0 {
		return nil
	}

	for _, planID := range planIDs {
		IDKey := fmt.Sprintf("subscription_plan:id:%s", planID)
		if err := p.client.Del(ctx, IDKey).Err(); err != nil {
			return fmt.Errorf("error deleting subscription plan with id %s: %w", planID, err)
		}
	}

	if err := p.client.Del(ctx, TypeKey).Err(); err != nil {
		return fmt.Errorf("error deleting type set for type %s: %w", typeID, err)
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
