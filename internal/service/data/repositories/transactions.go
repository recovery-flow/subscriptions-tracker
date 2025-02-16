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

type Transactions interface {
	New() Transactions

	Insert(r *http.Request, tr models.Transaction) (*models.Transaction, error)
	Select(r *http.Request) ([]models.Transaction, error)
	Get(r *http.Request) (*models.Transaction, error)

	FilterStrict(filters map[string]interface{}) Transactions

	DeleteOne(r *http.Request) error
	DeleteMany(r *http.Request) (int64, error)

	Limit(limit int64) Transactions
	Skip(skip int64) Transactions
}

type transactions struct {
	redis   cache.Transactions
	mongo   mongodb.Transactions
	filters map[string]interface{}
	limit   int64
	skip    int64
}

func NewTransactions(cfg config.Config) (Transactions, error) {
	redisRepo := cache.NewTransactions(redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	}), 15*time.Minute)
	mongoRepo, err := mongodb.NewTransactions(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, "transactions")
	if err != nil {
		return nil, err
	}
	return &transactions{
		redis:   redisRepo,
		mongo:   mongoRepo,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,
	}, nil
}

func (t *transactions) New() Transactions {
	return &transactions{
		redis:   t.redis,
		mongo:   t.mongo,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,
	}
}

func (t *transactions) Insert(r *http.Request, tr models.Transaction) (*models.Transaction, error) {
	created, err := t.mongo.New().Insert(r.Context(), tr)
	if err != nil {
		return nil, err
	}

	err = t.redis.New().Add(r.Context(), *created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (t *transactions) Select(r *http.Request) ([]models.Transaction, error) {
	subs, err := t.mongo.New().FilterStrict(t.filters).Limit(t.limit).Skip(t.skip).Select(r.Context())
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func (t *transactions) Get(r *http.Request) (*models.Transaction, error) {
	sub, err := t.redis.New().Filter(t.filters).Get(r.Context())
	if err != nil {
		return nil, err
	}
	if sub != nil {
		return sub, nil
	}

	sub, err = t.mongo.New().FilterStrict(t.filters).Get(r.Context())
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (t *transactions) FilterStrict(filters map[string]interface{}) Transactions {
	var validFilters = map[string]bool{
		"_id":            true,
		"user_id":        true,
		"amount":         true,
		"currency":       true,
		"status":         true,
		"payment_method": true,
		"prov_tx_id":     true,
		"created_at":     true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		t.filters[field] = value
	}
	return t
}

func (t *transactions) DeleteOne(r *http.Request) error {
	if err := t.redis.New().Filter(t.filters).DeleteOne(r.Context()); err != nil {
		return err
	}
	return t.mongo.New().FilterStrict(t.filters).DeleteOne(r.Context())
}

func (t *transactions) DeleteMany(r *http.Request) (int64, error) {
	if err := t.redis.New().Filter(t.filters).DeleteMany(r.Context()); err != nil {
		return 0, err
	}
	return t.mongo.New().FilterStrict(t.filters).DeleteMany(r.Context())
}

func (t *transactions) Limit(limit int64) Transactions {
	t.limit = limit
	return t
}

func (t *transactions) Skip(skip int64) Transactions {
	t.skip = skip
	return t
}
