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

type SubTypes interface {
	New() SubTypes

	Create(ctx context.Context, sub models.SubscriptionType) error
	Get(ctx context.Context) (*models.SubscriptionType, error)
	Select(ctx context.Context) ([]models.SubscriptionType, error)
	Delete(ctx context.Context) error

	Count(ctx context.Context) (int, error)

	Filter(filters map[string]any) SubTypes
}

type subTypes struct {
	redis   cache.SubTypes
	sql     sqldb.SubscriptionTypes
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewSubTypes(sql sqldb.SubscriptionTypes, redis cache.SubTypes, log *logrus.Logger) SubTypes {
	return &subTypes{
		redis:   redis,
		sql:     sql,
		filters: make(map[string]interface{}),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (t *subTypes) New() SubTypes {
	return NewSubTypes(&t.redis, &t.sql, t.log)
}

func (t *subTypes) Create(ctx context.Context, sub models.SubscriptionType) error {
	if err := t.sql.New().Insert(ctx, sub); err != nil {
		return err
	}

	err := t.redis.Add(ctx, sub)
	if err != nil {
		t.log.WithField("redis", err).Error("error adding subscription type to cache")
	}
	return nil
}

func (t *subTypes) Get(ctx context.Context) (*models.SubscriptionType, error) {
	if t.filters["id"] != nil {
		return t.redis.Get(ctx, t.filters["id"].(string))
	}

	return t.sql.New().Filter(t.filters).Get(ctx)
}

func (t *subTypes) Select(ctx context.Context) ([]models.SubscriptionType, error) {
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

func (t *subTypes) Delete(ctx context.Context) error {
	if t.filters["id"] != nil {
		return t.redis.Delete(ctx, t.filters["id"].(string))
	}

	return t.sql.New().Filter(t.filters).Delete(ctx)
}

func (t *subTypes) Count(ctx context.Context) (int, error) {
	return t.sql.New().Filter(t.filters).Count(ctx)
}

func (t *subTypes) Filter(filters map[string]any) SubTypes {
	t.filters = filters
	return t
}
