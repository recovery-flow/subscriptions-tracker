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

type Subscribers interface {
	New() Subscribers

	Insert(r *http.Request, sub models.Subscriber) (*models.Subscriber, error)
	Select(r *http.Request) ([]models.Subscriber, error)
	Get(r *http.Request) (*models.Subscriber, error)

	FilterStrict(filters map[string]interface{}) Subscribers

	UpdateOne(r *http.Request, fields map[string]interface{}) (*models.Subscriber, error)
	UpdateMany(r *http.Request, fields map[string]interface{}) (int64, error)

	DeleteOne(r *http.Request) error
	DeleteMany(r *http.Request) (int64, error)

	Limit(limit int64) Subscribers
	Skip(skip int64) Subscribers
}

type subscribers struct {
	redis   cache.Subscribers
	mongo   mongodb.Subscribers
	filters map[string]interface{}
	limit   int64
	skip    int64
}

func NewSubscribers(cfg *config.Config) (Subscribers, error) {
	redisRepo := cache.NewSubscribers(
		redis.NewClient(&redis.Options{
			Addr:     cfg.Database.Redis.Addr,
			Password: cfg.Database.Redis.Password,
			DB:       cfg.Database.Redis.DB,
		}),
		15*time.Minute,
	)
	mongoRepo, err := mongodb.NewSubscribers(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, "subscribers")
	if err != nil {
		return nil, err
	}
	return &subscribers{
		redis:   redisRepo,
		mongo:   mongoRepo,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,
	}, nil
}

func (s *subscribers) New() Subscribers {
	return &subscribers{
		redis:   s.redis,
		mongo:   s.mongo,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,
	}
}

func (s *subscribers) Insert(r *http.Request, sub models.Subscriber) (*models.Subscriber, error) {
	created, err := s.mongo.New().Insert(r.Context(), sub)
	if err != nil {
		return nil, err
	}

	err = s.redis.New().Add(r.Context(), *created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *subscribers) Select(r *http.Request) ([]models.Subscriber, error) {
	subs, err := s.mongo.New().FilterStrict(s.filters).Limit(s.limit).Skip(s.skip).Select(r.Context())
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func (s *subscribers) Get(r *http.Request) (*models.Subscriber, error) {
	sub, err := s.redis.New().FilterStrict(s.filters).Get(r.Context())
	if err != nil {
		return nil, err
	}
	if sub != nil {
		return sub, nil
	}

	sub, err = s.mongo.FilterStrict(s.filters).Get(r.Context())
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *subscribers) FilterStrict(filters map[string]interface{}) Subscribers {
	var validFilters = map[string]bool{
		"_id":      true,
		"user_id":  true,
		"plan_id":  true,
		"status":   true,
		"start_at": true,
		"end_at":   true,
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

func (s *subscribers) UpdateOne(r *http.Request, fields map[string]interface{}) (*models.Subscriber, error) {
	updated, err := s.mongo.New().FilterStrict(s.filters).UpdateOne(r.Context(), fields)
	if err != nil {
		return nil, err
	}

	err = s.redis.New().Add(r.Context(), *updated)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
func (s *subscribers) UpdateMany(r *http.Request, fields map[string]interface{}) (int64, error) {
	count, err := s.mongo.New().FilterStrict(s.filters).UpdateMany(r.Context(), fields)
	if err != nil {
		return 0, err
	}

	_, err = s.redis.New().FilterStrict(s.filters).UpdateMany(r.Context(), fields)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *subscribers) DeleteOne(r *http.Request) error {
	err := s.mongo.New().FilterStrict(s.filters).DeleteOne(r.Context())
	if err != nil {
		return err
	}

	err = s.redis.New().FilterStrict(s.filters).DeleteOne(r.Context())
	if err != nil {
		return err
	}
	return nil
}

func (s *subscribers) DeleteMany(r *http.Request) (int64, error) {
	count, err := s.mongo.New().FilterStrict(s.filters).DeleteMany(r.Context())
	if err != nil {
		return 0, err
	}

	err = s.redis.New().FilterStrict(s.filters).DeleteMany(r.Context())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *subscribers) Limit(limit int64) Subscribers {
	s.limit = limit
	return s
}

func (s *subscribers) Skip(skip int64) Subscribers {
	s.skip = skip
	return s
}
