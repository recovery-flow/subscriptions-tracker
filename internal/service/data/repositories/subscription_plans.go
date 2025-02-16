package repositories

import (
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/mongodb"
	"github.com/redis/go-redis/v9"
)

type SubscriptionPlans interface {
}

type subscriptionPlans struct {
	redis *cache.SubscriptionPlans
	mongo *mongodb.SubscriptionPlans
}

func NewSubscriptionPlans(cfg config.Config) (SubscriptionPlans, error) {
	redisRepo := cache.NewSubscriptionPlans(redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	}), 15*time.Minute)
	mongo, err := mongodb.NewSubscriptionPlans(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, "subscription_plans")
	if err != nil {
		return nil, err
	}
	return &subscriptionPlans{
		redis: &redisRepo,
		mongo: &mongo,
	}, nil
}
