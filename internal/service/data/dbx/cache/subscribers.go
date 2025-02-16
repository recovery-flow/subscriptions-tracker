package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscribers interface {
	Mew() Subscribers

	Add(ctx context.Context, subscriber models.Subscriber) error
	Get(ctx context.Context) (*models.Subscriber, error)
	Select(ctx context.Context) ([]models.Subscriber, error)
	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) error
	UpdateMany(ctx context.Context, fields map[string]interface{}) (int64, error)

	FilterStrict(filters map[string]any) Subscribers
	Limit(limit int) Subscribers
	Skip(skip int) Subscribers
}

type subscriberCache struct {
	client   *redis.Client
	LifeTime time.Duration
	filters  map[string]string
	skip     int
	limit    int
}

func NewSubscribers(client *redis.Client, lifeTime time.Duration) Subscribers {
	return &subscriberCache{
		client:   client,
		LifeTime: lifeTime,
		filters:  make(map[string]string),
		skip:     0,
		limit:    0,
	}
}

func (s *subscriberCache) Mew() Subscribers {
	return &subscriberCache{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  make(map[string]string),
		skip:     0,
		limit:    0,
	}
}

func (s *subscriberCache) FilterStrict(newFilters map[string]any) Subscribers {
	copied := make(map[string]string)
	for k, v := range s.filters {
		copied[k] = v
	}
	for k, v := range newFilters {
		copied[k] = fmt.Sprintf("%v", v)
	}
	return &subscriberCache{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  copied,
		skip:     s.skip,
		limit:    s.limit,
	}
}

func (s *subscriberCache) Limit(limit int) Subscribers {
	return &subscriberCache{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  s.filters,
		skip:     s.skip,
		limit:    limit,
	}
}

func (s *subscriberCache) Skip(skip int) Subscribers {
	return &subscriberCache{
		client:   s.client,
		LifeTime: s.LifeTime,
		filters:  s.filters,
		skip:     skip,
		limit:    s.limit,
	}
}

func (s *subscriberCache) Add(ctx context.Context, subscriber models.Subscriber) error {
	key := fmt.Sprintf("sub:id:%s", subscriber.ID.Hex())
	userKey := fmt.Sprintf("sub:user_id:%s", subscriber.UserID.String())

	data := map[string]interface{}{
		"id":         subscriber.ID.Hex(),
		"user_id":    subscriber.UserID.String(),
		"plan_id":    subscriber.PlanID.Hex(),
		"status":     subscriber.Status,
		"start_at":   subscriber.StartAt.Time().Format(time.RFC3339),
		"end_at":     subscriber.EndAt.Time().Format(time.RFC3339),
		"created_at": subscriber.CreatedAt.Time().Format(time.RFC3339),
	}
	if subscriber.UpdatedAt != nil {
		data["updated_at"] = subscriber.UpdatedAt.Time().Format(time.RFC3339)
	}

	if err := s.client.HSet(ctx, key, data).Err(); err != nil {
		return fmt.Errorf("error adding subscriber to Redis: %w", err)
	}
	if err := s.client.Set(ctx, userKey, subscriber.ID.Hex(), 0).Err(); err != nil {
		return fmt.Errorf("error creating user index: %w", err)
	}
	if s.LifeTime > 0 {
		_ = s.client.Expire(ctx, key, s.LifeTime).Err()
		_ = s.client.Expire(ctx, userKey, s.LifeTime).Err()
	}
	return nil
}

func (s *subscriberCache) Select(ctx context.Context) ([]models.Subscriber, error) {
	keys, err := s.client.Keys(ctx, "sub:id:*").Result()
	if err != nil {
		return nil, fmt.Errorf("error retrieving keys: %w", err)
	}
	var result []models.Subscriber
	for _, key := range keys {
		data, err := s.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if s.matchesFilters(data) {
			sub, err := mapToSubscriber(data)
			if err == nil {
				result = append(result, *sub)
			}
		}
	}

	if s.skip > 0 && s.skip < len(result) {
		result = result[s.skip:]
	} else if s.skip >= len(result) {
		result = []models.Subscriber{}
	}
	if s.limit > 0 && s.limit < len(result) {
		result = result[:s.limit]
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no matching subscribers found")
	}
	return result, nil
}

func (s *subscriberCache) Get(ctx context.Context) (*models.Subscriber, error) {
	subs, err := s.Select(ctx)
	if err != nil {
		return nil, err
	}
	return &subs[0], nil
}

func (s *subscriberCache) UpdateMany(ctx context.Context, fields map[string]interface{}) (int64, error) {
	keys, err := s.client.Keys(ctx, "sub:id:*").Result()
	if err != nil {
		return 0, fmt.Errorf("error retrieving keys: %w", err)
	}
	var updatedCount int64 = 0
	for _, key := range keys {
		data, err := s.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
		}
		if s.matchesFilters(data) {
			if err := s.client.HSet(ctx, key, fields).Err(); err != nil {
				continue
			}
			updatedCount++
		}
	}
	return updatedCount, nil
}

func (s *subscriberCache) DeleteOne(ctx context.Context) error {
	keys, err := s.client.Keys(ctx, "sub:id:*").Result()
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
				return fmt.Errorf("error deleting subscriber: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no matching subscriber found for deletion")
}

func (s *subscriberCache) DeleteMany(ctx context.Context) error {
	keys, err := s.client.Keys(ctx, "sub:id:*").Result()
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
				delErr = err
			}
		}
	}
	return delErr
}

func (s *subscriberCache) matchesFilters(data map[string]string) bool {
	for fk, fv := range s.filters {
		if data[fk] != fv {
			return false
		}
	}
	return true
}

func mapToSubscriber(data map[string]string) (*models.Subscriber, error) {
	idHex, ok := data["id"]
	if !ok {
		return nil, fmt.Errorf("id not found in data")
	}
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(data["user_id"])
	if err != nil {
		return nil, err
	}
	planID, err := primitive.ObjectIDFromHex(data["plan_id"])
	if err != nil {
		return nil, err
	}
	startAt, err := time.Parse(time.RFC3339, data["start_at"])
	if err != nil {
		return nil, err
	}
	endAt, err := time.Parse(time.RFC3339, data["end_at"])
	if err != nil {
		return nil, err
	}
	createdAt, err := time.Parse(time.RFC3339, data["created_at"])
	if err != nil {
		return nil, err
	}
	var updatedAt *primitive.DateTime
	if val, ok := data["updated_at"]; ok && val != "" {
		t, err := time.Parse(time.RFC3339, val)
		if err == nil {
			tmp := primitive.NewDateTimeFromTime(t)
			updatedAt = &tmp
		}
	}
	return &models.Subscriber{
		ID:        id,
		UserID:    userID,
		PlanID:    planID,
		Status:    models.SubStatus(data["status"]),
		StartAt:   primitive.NewDateTimeFromTime(startAt),
		EndAt:     primitive.NewDateTimeFromTime(endAt),
		CreatedAt: primitive.NewDateTimeFromTime(createdAt),
		UpdatedAt: updatedAt,
	}, nil
}
