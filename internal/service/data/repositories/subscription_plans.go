package repositories

import (
	"net/http"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/mongodb"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

type SubscriptionPlans interface {
	New() SubscriptionPlans

	Insert(r *http.Request, plan models.SubscriptionPlan) (*models.SubscriptionPlan, error)
	Select(r *http.Request) ([]models.SubscriptionPlan, error)
	Get(r *http.Request) (*models.SubscriptionPlan, error)

	FilterStrict(filters map[string]interface{}) SubscriptionPlans

	UpdateOne(r *http.Request, fields map[string]interface{}) (*models.SubscriptionPlan, error)
	UpdateMany(r *http.Request, fields map[string]interface{}) (int64, error)

	DeleteOne(r *http.Request) error
	DeleteMany(r *http.Request) (int64, error)

	Limit(limit int64) SubscriptionPlans
	Skip(skip int64) SubscriptionPlans
}

type subscriptionPlans struct {
	redis   cache.SubscriptionPlans
	mongo   mongodb.SubscriptionPlans
	filters map[string]interface{}
	limit   int64
	skip    int64
}

func NewSubscriptionPlans(cfg config.Config) (SubscriptionPlans, error) {
	redisRepo := cache.NewSubscriptionPlans(redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	}), 15*time.Minute)
	mongoRepo, err := mongodb.NewSubscriptionPlans(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, "subscription_plans")
	if err != nil {
		return nil, err
	}
	return &subscriptionPlans{
		redis:   redisRepo,
		mongo:   mongoRepo,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,
	}, nil
}

func (s *subscriptionPlans) New() SubscriptionPlans {
	return &subscriptionPlans{
		redis:   s.redis,
		mongo:   s.mongo,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,
	}
}

func (s *subscriptionPlans) Insert(r *http.Request, plan models.SubscriptionPlan) (*models.SubscriptionPlan, error) {
	created, err := s.mongo.Insert(r.Context(), plan)
	if err != nil {
		return nil, err
	}

	err = s.redis.Add(r.Context(), plan)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *subscriptionPlans) Select(r *http.Request) ([]models.SubscriptionPlan, error) {
	plans, err := s.mongo.New().FilterStrict(s.filters).Limit(s.limit).Skip(s.skip).Select(r.Context())
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (s *subscriptionPlans) Get(r *http.Request) (*models.SubscriptionPlan, error) {
	plan, err := s.redis.New().Filter(s.filters).Get(r.Context())
	if err != nil {
		return nil, err
	}
	if plan != nil {
		return plan, nil
	}

	plan, err = s.mongo.New().FilterStrict(s.filters).Get(r.Context())
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (s *subscriptionPlans) FilterStrict(filters map[string]interface{}) SubscriptionPlans {
	var validFilters = map[string]bool{
		"_id":        true,
		"name":       true,
		"title":      true,
		"price":      true,
		"currency":   true,
		"pay_freq":   true,
		"status":     true,
		"cancel_at":  true,
		"updated_at": true,
		"created_at": true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		s.filters[field] = value
	}
	return s
}

func (s *subscriptionPlans) UpdateOne(r *http.Request, fields map[string]interface{}) (*models.SubscriptionPlan, error) {
	updated, err := s.mongo.New().FilterStrict(s.filters).UpdateOne(r.Context(), fields)
	if err != nil {
		return nil, err
	}

	err = s.redis.Add(r.Context(), *updated)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *subscriptionPlans) UpdateMany(r *http.Request, fields map[string]interface{}) (int64, error) {
	count, err := s.mongo.New().FilterStrict(s.filters).UpdateMany(r.Context(), fields)
	if err != nil {
		return 0, err
	}

	_, err = s.redis.New().Filter(s.filters).UpdateMany(r.Context(), fields)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *subscriptionPlans) DeleteOne(r *http.Request) error {
	err := s.mongo.New().FilterStrict(s.filters).DeleteOne(r.Context())
	if err != nil {
		return err
	}

	err = s.redis.New().Filter(s.filters).DeleteOne(r.Context())
	if err != nil {
		return err
	}
	return nil
}

func (s *subscriptionPlans) DeleteMany(r *http.Request) (int64, error) {
	count, err := s.mongo.New().FilterStrict(s.filters).DeleteMany(r.Context())
	if err != nil {
		return 0, err
	}

	err = s.redis.New().Filter(s.filters).DeleteMany(r.Context())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *subscriptionPlans) Limit(limit int64) SubscriptionPlans {
	s.limit = limit
	return s
}

func (s *subscriptionPlans) Skip(skip int64) SubscriptionPlans {
	s.skip = skip
	return s
}
