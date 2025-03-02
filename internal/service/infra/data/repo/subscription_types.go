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

type SubscriptionTypes interface {
	New() SubscriptionTypes

	Create(ctx context.Context, sub models.SubscriptionType) error
	Get(ctx context.Context) (*models.SubscriptionType, error)
	Select(ctx context.Context) ([]models.SubscriptionType, error)
	Delete(ctx context.Context) error

	Count(ctx context.Context) (int, error)

	Filter(filters map[string]any) SubscriptionTypes
}

type SubTypes struct {
	redis   cache.SubTypes
	sql     sqldb.SubscriptionTypes
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewSubTypes(redis *cache.SubTypes, sql *sqldb.SubscriptionTypes, log *logrus.Logger) SubscriptionTypes {
	return &SubTypes{
		redis:   *redis,
		sql:     *sql,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (t *SubTypes) New() SubscriptionTypes {
	return NewSubTypes(&t.redis, &t.sql, t.log)
}

func (t *SubTypes) Create(ctx context.Context, sub models.SubscriptionType) error {
	if err := t.sql.New().Insert(ctx, sub); err != nil {
		return err
	}

	err := t.redis.Add(ctx, sub)
	if err != nil {
		t.log.WithField("redis", err).Error("error adding subscription type to cache")
	}
	return nil
}

func (t *SubTypes) Get(ctx context.Context) (*models.SubscriptionType, error) {
	if t.filters["id"] != nil {
		return t.redis.Get(ctx, t.filters["id"].(string))
	}

	return t.sql.New().Filter(t.filters).Get(ctx)
}

func (t *SubTypes) Select(ctx context.Context) ([]models.SubscriptionType, error) {
	if t.filters["id"] != nil {
		res, err := t.redis.Get(ctx, t.filters["id"].(string))
		if err != nil || !errors.Is(err, redis.Nil) {
			t.log.WithField("redis", err).Error("error adding subscription type to cache")
		} else {
			return []models.SubscriptionType{*res}, nil
		}
	}

	return t.sql.New().Filter(t.filters).Page(uint64(t.limit), uint64(t.skip)).Select(ctx)
}

func (t *SubTypes) Delete(ctx context.Context) error {
	if t.filters["id"] != nil {
		return t.redis.Delete(ctx, t.filters["id"].(string))
	}

	return t.sql.New().Filter(t.filters).Delete(ctx)
}

func (t *SubTypes) Count(ctx context.Context) (int, error) {
	return t.sql.New().Filter(t.filters).Count(ctx)
}

func (t *SubTypes) Filter(filters map[string]any) SubscriptionTypes {
	t.filters = filters
	return t
}
