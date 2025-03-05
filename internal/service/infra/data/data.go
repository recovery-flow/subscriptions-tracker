package data

import (
	"database/sql"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/sqldb"
	"github.com/redis/go-redis/v9"
)

type Data struct {
	SQL   SQLStorage
	Cache CacheStorage
}

type SQLStorage struct {
	Schedule       sqldb.BillingSchedules
	Transactions   sqldb.Transactions
	PaymentMethods sqldb.PaymentMethods
	Plans          sqldb.SubPlan
	Types          sqldb.SubTypes
	Subscriptions  sqldb.Subscriptions
}

type CacheStorage struct {
	Plans         cache.Plans
	Subscriptions cache.Subscriptions
	Types         cache.Types
}

func NewData(cfg *config.Config) (*Data, error) {
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

	redisPlans := cache.NewPlansCache(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)
	redisSubs := cache.NewSubscriptions(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)
	redisTypes := cache.NewSubTypesQueryCache(redisClient, time.Duration(cfg.Database.Redis.Lifetime)*time.Minute)

	return &Data{
		SQL: SQLStorage{
			Schedule:       sqlSchedule,
			Transactions:   sqlTrans,
			PaymentMethods: sqlPM,
			Plans:          sqlSubPlans,
			Types:          sqlSubTypes,
			Subscriptions:  sqlSub,
		},
		Cache: CacheStorage{
			Plans:         redisPlans,
			Subscriptions: redisSubs,
			Types:         redisTypes,
		},
	}, nil
}
