package data

import (
	"database/sql"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Data struct {
	BillingPlan    repo.BillingPlan
	Transactions   repo.Transactions
	PaymentMethods repo.PaymentMethods
	SubPlans       repo.SubPlans
	SubTypes       repo.SubTypes
	Subscription   repo.Subscription
}

func NewData(cfg *config.Config, log *logrus.Logger) (*Data, error) {
	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	})

	sqlBP := sqldb.NewBillingSchedules(db)
	sqlTrans := sqldb.NewTransactions(db)
	sqlPM := sqldb.NewPaymentMethods(db)
	sqlSubPlans := sqldb.NewSubPlan(db)
	sqlSubTypes := sqldb.NewSubTypes(db)
	sqlSub := sqldb.NewSubscriptions(db)

	redisPM := cache.NewPayMethods(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)
	redisPlans := cache.NewSubPlans(redisClient)
	redisSubs := cache.NewSubscriptions(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)
	redisTypes := cache.NewSubTypes(redisClient)

	return &Data{
		BillingPlan:    repo.NewBillingPlan(sqlBP),
		Transactions:   repo.NewTransactions(sqlTrans),
		PaymentMethods: repo.NewPaymentMethods(sqlPM, redisPM, log),
		SubPlans:       repo.NewSubPlans(sqlSubPlans, redisPlans, log),
		SubTypes:       repo.NewSubTypes(sqlSubTypes, redisTypes, log),
		Subscription:   repo.NewSubscription(sqlSub, redisSubs, log),
	}, nil
}
