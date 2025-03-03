package data

import (
	"database/sql"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/sqldb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Data struct {
	SQL   SQLStorage
	Cache CacheStorage
}

type SQLStorage struct {
	Schedule      sqldb.BillingSchedules
	Transactions  sqldb.Transactions
	Methods       sqldb.PaymentMethods
	Plans         sqldb.SubPlan
	Types         sqldb.SubTypes
	Subscriptions sqldb.Subscriptions
}

type CacheStorage struct {
	Plans         cache.SubPlanQueryCache
	Subscriptions cache.Subscriptions
	Types         cache.SubTypesQueryCache
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

	sqlSchedule := sqldb.NewBillingSchedules(db)
	sqlTrans := sqldb.NewTransactions(db)
	sqlPM := sqldb.NewPaymentMethods(db)
	sqlSubPlans := sqldb.NewSubPlan(db)
	sqlSubTypes := sqldb.NewSubTypes(db)
	sqlSub := sqldb.NewSubscriptions(db)

	redisPlans := cache.NewSubPlanQueryCache(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)
	redisSubs := cache.NewSubscriptions(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)
	redisTypes := cache.NewSubTypesQueryCache(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)

	return &Data{
		SQL: SQLStorage{
			Schedule:      sqlSchedule,
			Transactions:  sqlTrans,
			Methods:       sqlPM,
			Plans:         sqlSubPlans,
			Types:         sqlSubTypes,
			Subscriptions: sqlSub,
		},
		Cache: CacheStorage{
			Plans:         redisPlans,
			Subscriptions: redisSubs,
			Types:         redisTypes,
		},
	}, nil
}
