package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type SubPlan interface {
	New() SubPlan

	Create(ctx context.Context, plan models.SubscriptionPlan) error
	Get(ctx context.Context) (*models.SubscriptionPlan, error)
	Select(ctx context.Context) ([]models.SubscriptionPlan, error)
	DeleteOne(ctx context.Context) error
	Update(ctx context.Context, update map[string]any) error

	Count(ctx context.Context) (int, error)

	Filter(filters map[string]any) SubPlan

	Page(limit, offset uint64) SubPlan

	DropCache(ctx context.Context) error
}

type subPlan struct {
	redis   cache.SubPlanQueryCache
	sql     sqldb.SubPlan
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewSubPlan(sql sqldb.SubPlan, redis cache.SubPlanQueryCache, log *logrus.Logger) SubPlan {
	return &subPlan{
		redis:   redis,
		sql:     sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (p *subPlan) New() SubPlan {
	return NewSubPlan(p.sql, p.redis, p.log)
}

func (p *subPlan) Create(ctx context.Context, plan models.SubscriptionPlan) error {
	if err := p.sql.New().Insert(ctx, plan); err != nil {
		return err
	}

	if err := p.redis.Drop(ctx); err != nil {
		p.log.WithField("redis", err).Error("error dropping subscription plan cache")
	}

	if err := p.redis.Set(ctx, p.cacheKey(), []models.SubscriptionPlan{plan}); err != nil {
		p.log.WithField("redis", err).Error("error setting subscription plan in cache")
	}

	return nil
}

func (p *subPlan) Get(ctx context.Context) (*models.SubscriptionPlan, error) {
	resCache, err := p.redis.Get(ctx, p.cacheKey())
	if err != nil || !errors.Is(err, redis.Nil) {
		p.log.WithField("redis", err).Error("error setting subscription plan in cache")
	} else if len(resCache) > 0 {
		return &resCache[0], nil
	}

	res, err := p.sql.New().Filter(p.filters).Get(ctx)
	if err != nil {
		return nil, err
	}

	err = p.redis.Set(ctx, p.cacheKey(), []models.SubscriptionPlan{*res})
	if err != nil {
		p.log.WithField("redis", err).Error("error setting subscription plan in cache")
	}

	return res, nil
}

func (p *subPlan) Select(ctx context.Context) ([]models.SubscriptionPlan, error) {
	resCache, err := p.redis.Get(ctx, p.cacheKey())
	if err != nil || !errors.Is(err, redis.Nil) {
		p.log.WithField("redis", err).Error("error setting subscription plan in cache")
	} else if resCache != nil {
		return resCache, nil
	}

	res, err := p.sql.New().Filter(p.filters).Page(uint64(p.limit), uint64(p.skip)).Select(ctx)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	err = p.redis.Set(ctx, p.cacheKey(), res)
	if err != nil {
		p.log.WithField("redis", err).Error("error setting subscription plan in cache")
	}

	return res, nil
}

func (p *subPlan) Update(ctx context.Context, update map[string]any) error {
	if err := p.redis.Drop(ctx); err != nil {
		p.log.WithField("redis", err).Error("error dropping subscription plan cache")
	}

	return p.sql.New().Filter(p.filters).Update(ctx, update)
}

func (p *subPlan) DeleteOne(ctx context.Context) error {
	if err := p.redis.Drop(ctx); err != nil {
		p.log.WithField("redis", err).Error("error dropping subscription plan cache")
	}

	return p.sql.New().Filter(p.filters).Delete(ctx)
}

func (p *subPlan) DropCache(ctx context.Context) error {
	err := p.redis.Drop(ctx)
	if err != nil {
		p.log.WithField("redis", err).Error("error dropping subscription plan cache")
	}

	return err
}

func (p *subPlan) Count(ctx context.Context) (int, error) {
	return p.sql.New().Filter(p.filters).Count(ctx)
}

func (p *subPlan) Filter(filters map[string]any) SubPlan {
	p.filters = filters
	return p
}

func (p *subPlan) Page(limit, offset uint64) SubPlan {
	p.limit = int64(limit)
	p.skip = int64(offset)
	return p
}

func (p *subPlan) cacheKey() string {
	key := cache.SubscriptionPlanCollection
	if len(p.filters) > 0 {
		key += fmt.Sprintf(":filters=%v", p.filters)
	}
	key += fmt.Sprintf(":limit=%d:skip=%d", p.limit, p.skip)
	return key
}
