package repo

import (
	"context"
	"errors"
	"fmt"

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
	Update(ctx context.Context, update map[string]any) error

	Count(ctx context.Context) (int, error)

	Filter(filters map[string]any) SubTypes

	Page(limit, offset uint64) SubTypes

	DropCache(ctx context.Context) error
}

type subTypes struct {
	redis   cache.SubTypesQueryCache
	sql     sqldb.SubTypes
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewSubTypes(sql sqldb.SubTypes, redis cache.SubTypesQueryCache, log *logrus.Logger) SubTypes {
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
	return NewSubTypes(t.sql, t.redis, t.log)
}

func (t *subTypes) Create(ctx context.Context, subType models.SubscriptionType) error {
	if err := t.sql.New().Insert(ctx, subType); err != nil {
		return err
	}

	if err := t.redis.Drop(ctx); err != nil {
		t.log.WithField("redis", err).Error("error dropping subscription types cache")
	}

	if err := t.redis.Set(ctx, t.cacheKey(), []models.SubscriptionType{subType}); err != nil {
		t.log.WithField("redis", err).Error("error setting subscription plan in cache")
	}

	return nil
}

func (t *subTypes) Get(ctx context.Context) (*models.SubscriptionType, error) {
	resCache, err := t.redis.Get(ctx, t.cacheKey())
	if err != nil || !errors.Is(err, redis.Nil) {
		t.log.WithField("redis", err).Error("error setting subscription plan in cache")
	} else if len(resCache) > 0 {
		return &resCache[0], nil
	}

	res, err := t.sql.New().Filter(t.filters).Get(ctx)
	if err != nil {
		return nil, err
	}

	err = t.redis.Set(ctx, t.cacheKey(), []models.SubscriptionType{*res})
	if err != nil {
		t.log.WithField("redis", err).Error("error setting subscription plan in cache")
	}

	return res, nil
}

func (t *subTypes) Select(ctx context.Context) ([]models.SubscriptionType, error) {
	resCache, err := t.redis.Get(ctx, t.cacheKey())
	if err != nil || !errors.Is(err, redis.Nil) {
		t.log.WithField("redis", err).Error("error setting subscription plan in cache")
	} else if resCache != nil {
		return resCache, nil
	}

	res, err := t.sql.New().Filter(t.filters).Page(uint64(t.limit), uint64(t.skip)).Select(ctx)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	err = t.redis.Set(ctx, t.cacheKey(), res)
	if err != nil {
		t.log.WithField("redis", err).Error("error setting subscription plan in cache")
	}

	return res, nil
}

func (t *subTypes) Update(ctx context.Context, update map[string]any) error {
	if err := t.redis.Drop(ctx); err != nil {
		t.log.WithField("redis", err).Error("error dropping subscription types cache")
	}

	return t.sql.New().Filter(t.filters).Update(ctx, update)
}

func (t *subTypes) Delete(ctx context.Context) error {
	if err := t.redis.Drop(ctx); err != nil {
		t.log.WithField("redis", err).Error("error dropping subscription types cache")
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

func (t *subTypes) Page(limit, offset uint64) SubTypes {
	t.limit = int64(limit)
	t.skip = int64(offset)
	return t
}

func (t *subTypes) DropCache(ctx context.Context) error {
	err := t.redis.Drop(ctx)
	if err != nil {
		t.log.WithField("redis", err).Error("error dropping subscription plans cache")
	}

	return err
}

func (t *subTypes) cacheKey() string {
	key := cache.SubscriptionTypesCollection
	if len(t.filters) > 0 {
		key += fmt.Sprintf(":filters=%v", t.filters)
	}
	key += fmt.Sprintf(":limit=%d:skip=%d", t.limit, t.skip)
	return key
}
