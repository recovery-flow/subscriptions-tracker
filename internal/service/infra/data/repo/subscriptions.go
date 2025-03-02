package repo

import (
	"context"
	"errors"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Subscription interface {
	New() Subscription

	Create(ctx context.Context, sub models.Subscription) error
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]models.Subscription, error)
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context) (*models.Subscription, error)

	Filter(filters map[string]interface{}) Subscription

	Page(limit, offset uint64) Subscription
}

type subscription struct {
	redis   cache.Subscriptions
	sql     sqldb.Subscriptions
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewSubscription(redis *cache.Subscriptions, sql *sqldb.Subscriptions, log *logrus.Logger) Subscription {
	return &subscription{
		redis:   *redis,
		sql:     *sql,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (s *subscription) New() Subscription {
	return NewSubscription(&s.redis, &s.sql, s.log)
}

func (s *subscription) Create(ctx context.Context, sub models.Subscription) error {
	if err := s.sql.New().Insert(ctx, sub); err != nil {
		return err
	}

	if err := s.redis.Add(ctx, sub); err != nil {
		s.log.WithField("redis", err).Error("error adding subscription to cache")
	}

	return nil
}

func (s *subscription) Delete(ctx context.Context) error {
	if err := s.sql.New().Delete(ctx); err != nil {
		return err
	}

	if err := s.redis.Delete(ctx, s.filters["user_id"].(string)); err != nil && s.filters["user_id"] != nil {
		s.log.WithField("redis", err).Error("error deleting subscription from cache")
	}

	return nil
}

func (s *subscription) Select(ctx context.Context) ([]models.Subscription, error) {
	if s.filters["user_id"] != nil {
		sub, err := s.redis.Get(ctx, s.filters["user_id"].(string))
		if err != nil || !errors.Is(err, redis.Nil) {
			s.log.WithField("redis", err).Error("error getting subscription from cache")
		} else {
			return []models.Subscription{*sub}, nil
		}
	}

	return s.sql.New().Filter(s.filters).Select(ctx)
}

func (s *subscription) Count(ctx context.Context) (int, error) {
	return s.sql.New().Filter(s.filters).Count(ctx)
}

func (s *subscription) Get(ctx context.Context) (*models.Subscription, error) {
	if s.filters["user_id"] != nil {
		sub, err := s.redis.Get(ctx, s.filters["user_id"].(string))
		if err != nil || !errors.Is(err, redis.Nil) {
			s.log.WithField("redis", err).Error("error getting subscription from cache")
		} else {
			return sub, nil
		}
	}

	return s.sql.New().Filter(s.filters).Get(ctx)
}

func (s *subscription) Filter(filters map[string]interface{}) Subscription {
	s.filters = filters
	return s
}

func (s *subscription) Page(limit, offset uint64) Subscription {
	s.limit = int64(limit)
	s.skip = int64(offset)
	return s
}
