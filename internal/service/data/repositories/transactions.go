package repositories

import (
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/mongodb"
	"github.com/redis/go-redis/v9"
)

type Transactions interface {
}

type transactions struct {
	redis *cache.Transactions
	mongo *mongodb.Transactions
}

func NewTransactions(cfg config.Config) (Transactions, error) {
	redisRepo := cache.NewTransactions(redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	}), 15*time.Minute)
	mongo, err := mongodb.NewTransactions(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, "transactions")
	if err != nil {
		return nil, err
	}
	return &transactions{
		redis: &redisRepo,
		mongo: &mongo,
	}, nil
}
