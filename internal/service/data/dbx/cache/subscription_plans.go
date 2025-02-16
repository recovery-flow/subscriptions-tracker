package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubscriptionPlans описывает кэш-операции для планов подписки с поддержкой фильтрации и пагинации.
type SubscriptionPlans interface {
	Mew() SubscriptionPlans

	Add(ctx context.Context, plan models.SubscriptionPlan) error
	Get(ctx context.Context) (*models.SubscriptionPlan, error)
	Select(ctx context.Context) ([]models.SubscriptionPlan, error)
	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) error

	Filter(ctx context.Context, filters map[string]any) SubscriptionPlans
	Limit(limit int) SubscriptionPlans
	Skip(skip int) SubscriptionPlans
}

type subscriptionPlans struct {
	client   *redis.Client
	LifeTime time.Duration
	filters  map[string]string
	skip     int
	limit    int
}

func NewSubscriptionPlans(client *redis.Client, lifeTime time.Duration) SubscriptionPlans {
	return &subscriptionPlans{
		client:   client,
		LifeTime: lifeTime,
		filters:  make(map[string]string),
		skip:     0,
		limit:    0,
	}
}

func (s *subscriptionPlans) Mew() SubscriptionPlans {
	return &subscriptionPlans{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  make(map[string]string),
		skip:     0,
		limit:    0,
	}
}

func (s *subscriptionPlans) Filter(ctx context.Context, newFilters map[string]any) SubscriptionPlans {
	copied := make(map[string]string)
	for k, v := range s.filters {
		copied[k] = v
	}
	for k, v := range newFilters {
		copied[k] = fmt.Sprintf("%v", v)
	}
	return &subscriptionPlans{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  copied,
		skip:     s.skip,
		limit:    s.limit,
	}
}

func (s *subscriptionPlans) Limit(limit int) SubscriptionPlans {
	return &subscriptionPlans{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  s.filters,
		skip:     s.skip,
		limit:    limit,
	}
}

func (s *subscriptionPlans) Skip(skip int) SubscriptionPlans {
	return &subscriptionPlans{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  s.filters,
		skip:     skip,
		limit:    s.limit,
	}
}

func (s *subscriptionPlans) Add(ctx context.Context, plan models.SubscriptionPlan) error {
	key := fmt.Sprintf("subplan:id:%s", plan.ID.Hex())
	data := map[string]interface{}{
		"id":         plan.ID.Hex(),
		"name":       plan.Name,
		"title":      plan.Title,
		"desc":       plan.Desc,
		"price":      plan.Price,
		"currency":   plan.Currency,
		"pay_freq":   string(plan.PayFreq),
		"status":     string(plan.Status),
		"created_at": plan.CreatedAt.Time().Format(time.RFC3339),
	}
	if plan.CanceledAt != nil {
		data["canceled_at"] = plan.CanceledAt.Time().Format(time.RFC3339)
	}
	if plan.UpdatedAt != nil {
		data["updated_at"] = plan.UpdatedAt.Time().Format(time.RFC3339)
	}
	if err := s.client.HSet(ctx, key, data).Err(); err != nil {
		return fmt.Errorf("error adding subscription plan to Redis: %w", err)
	}
	if s.LifeTime > 0 {
		_ = s.client.Expire(ctx, key, s.LifeTime).Err()
	}
	return nil
}

func (s *subscriptionPlans) Select(ctx context.Context) ([]models.SubscriptionPlan, error) {
	keys, err := s.client.Keys(ctx, "subplan:id:*").Result()
	if err != nil {
		return nil, fmt.Errorf("error retrieving keys: %w", err)
	}
	var result []models.SubscriptionPlan
	for _, key := range keys {
		data, err := s.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if s.matchesFilters(data) {
			plan, err := mapToSubscriptionPlan(data)
			if err == nil {
				result = append(result, *plan)
			}
		}
	}

	if s.skip > 0 && s.skip < len(result) {
		result = result[s.skip:]
	} else if s.skip >= len(result) {
		result = []models.SubscriptionPlan{}
	}
	if s.limit > 0 && s.limit < len(result) {
		result = result[:s.limit]
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no matching subscription plans found")
	}
	return result, nil
}

func (s *subscriptionPlans) Get(ctx context.Context) (*models.SubscriptionPlan, error) {
	res, err := s.Select(ctx)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (s *subscriptionPlans) DeleteOne(ctx context.Context) error {
	keys, err := s.client.Keys(ctx, "subplan:id:*").Result()
	if err != nil {
		return fmt.Errorf("error retrieving keys: %w", err)
	}
	for _, key := range keys {
		data, err := s.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if s.matchesFilters(data) {
			if err := s.client.Del(ctx, key).Err(); err != nil {
				return fmt.Errorf("error deleting subscription plan: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no matching subscription plan found for deletion")
}

func (s *subscriptionPlans) DeleteMany(ctx context.Context) error {
	keys, err := s.client.Keys(ctx, "subplan:id:*").Result()
	if err != nil {
		return fmt.Errorf("error retrieving keys: %w", err)
	}
	var delErr error
	for _, key := range keys {
		data, err := s.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if s.matchesFilters(data) {
			if err := s.client.Del(ctx, key).Err(); err != nil {
				delErr = err // сохраняем ошибку, но продолжаем удаление
			}
		}
	}
	return delErr
}

func (s *subscriptionPlans) matchesFilters(data map[string]string) bool {
	for fk, fv := range s.filters {
		if data[fk] != fv {
			return false
		}
	}
	return true
}

func mapToSubscriptionPlan(data map[string]string) (*models.SubscriptionPlan, error) {
	idHex, ok := data["id"]
	if !ok {
		return nil, fmt.Errorf("id not found in data")
	}
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, err
	}
	priceStr, ok := data["price"]
	if !ok {
		return nil, fmt.Errorf("price not found")
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return nil, err
	}
	createdAt, err := time.Parse(time.RFC3339, data["created_at"])
	if err != nil {
		return nil, err
	}
	var canceledAt *primitive.DateTime
	if val, ok := data["canceled_at"]; ok && val != "" {
		t, err := time.Parse(time.RFC3339, val)
		if err == nil {
			tmp := primitive.NewDateTimeFromTime(t)
			canceledAt = &tmp
		}
	}
	var updatedAt *primitive.DateTime
	if val, ok := data["updated_at"]; ok && val != "" {
		t, err := time.Parse(time.RFC3339, val)
		if err == nil {
			tmp := primitive.NewDateTimeFromTime(t)
			updatedAt = &tmp
		}
	}
	return &models.SubscriptionPlan{
		ID:         id,
		Name:       data["name"],
		Title:      data["title"],
		Desc:       data["desc"],
		Price:      price,
		Currency:   data["currency"],
		PayFreq:    models.PayFreq(data["pay_freq"]),
		Status:     models.PlanStatus(data["status"]),
		CanceledAt: canceledAt,
		UpdatedAt:  updatedAt,
		CreatedAt:  primitive.NewDateTimeFromTime(createdAt),
	}, nil
}
