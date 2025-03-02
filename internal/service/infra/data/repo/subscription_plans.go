package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type SubPlans interface {
	New() SubPlans

	Create(ctx context.Context, plan models.SubscriptionPlan) error
	Get(ctx context.Context) (*models.SubscriptionPlan, error)
	Select(ctx context.Context) ([]models.SubscriptionPlan, error)
	DeleteByID(ctx context.Context) error
	DeleteByTypeID(ctx context.Context) error

	Count(ctx context.Context) (int, error)

	Filter(filters map[string]any) SubPlans

	Page(limit, offset uint64) SubPlans
}

type subPlans struct {
	redis   cache.SubPlans
	sql     sqldb.SubPlan
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewSubPlans(redis *cache.SubPlans, sql *sqldb.SubPlan, log *logrus.Logger) SubPlans {
	return &subPlans{
		redis:   *redis,
		sql:     *sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (p *subPlans) New() SubPlans {
	return NewSubPlans(&p.redis, &p.sql, p.log)
}

func (p *subPlans) Create(ctx context.Context, plan models.SubscriptionPlan) error {
	if err := p.sql.New().Insert(ctx, plan); err != nil {
		return err
	}

	err := p.redis.Add(ctx, plan)
	if err != nil {
		p.log.WithField("error", err).Error("error adding subscription plan to cache")
	}

	return nil
}

func (p *subPlans) Get(ctx context.Context) (*models.SubscriptionPlan, error) {
	res, err := p.redis.GetByID(ctx, p.filters["id"].(string))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			res = nil
		} else {
			p.log.WithField("error", err).Error("error adding subscription plan to cache")
		}
	} else {
		return res, nil
	}

	res, err = p.sql.New().Filter(p.filters).Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *subPlans) Select(ctx context.Context) ([]models.SubscriptionPlan, error) {
	if p.filters["id"] != nil {
		res, err := p.redis.GetByID(ctx, p.filters["id"].(string))
		if err != nil || errors.Is(err, redis.Nil) {
			return nil, err
		}
		return []models.SubscriptionPlan{*res}, nil
	}

	if p.filters["type_id"] != nil {
		return p.redis.GetByTypeID(ctx, p.filters["type_id"].(string))
	}

	res, err := p.sql.New().Filter(p.filters).Page(uint64(p.limit), uint64(p.skip)).Select(ctx)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	return res, nil
}

func (p *subPlans) DeleteByID(ctx context.Context) error {
	err := p.redis.DeleteByID(ctx, p.filters["id"].(string))
	if err != nil {
		p.log.WithField("error", err).Error("error deleting subscription plan from cache")
	}

	return p.sql.New().Filter(p.filters).Delete(ctx)
}

func (p *subPlans) DeleteByTypeID(ctx context.Context) error {
	err := p.redis.DeleteByType(ctx, p.filters["type_id"].(string))
	if err != nil {
		p.log.WithField("error", err).Error("error deleting subscription plan from cache")
	}

	return p.sql.New().Filter(p.filters).Delete(ctx)
}

func (p *subPlans) Count(ctx context.Context) (int, error) {
	return p.sql.New().Filter(p.filters).Count(ctx)
}

func (p *subPlans) Filter(filters map[string]any) SubPlans {
	p.filters = filters
	return p
}

func (p *subPlans) Page(limit, offset uint64) SubPlans {
	p.limit = int64(limit)
	p.skip = int64(offset)
	return p
}
