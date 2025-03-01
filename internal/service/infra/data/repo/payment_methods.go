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

type PaymentMethods interface {
	New() PaymentMethods

	Create(ctx context.Context, pm models.PaymentMethod) error
	Get(ctx context.Context) (*models.PaymentMethod, error)
	Select(ctx context.Context) ([]models.PaymentMethod, error)
	DeleteByID(ctx context.Context) error
	DeleteByUserID(ctx context.Context) error

	Count(ctx context.Context) (int, error)

	SelectForUser(ctx context.Context) ([]models.PaymentMethod, error)

	Filter(filters map[string]any) PaymentMethods

	Page(limit, offset uint64) PaymentMethods
}

type paymentMethods struct {
	redis   cache.PayMethods
	sql     sqldb.PaymentMethods
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewPaymentMethods(redis cache.PayMethods, sql sqldb.PaymentMethods, log *logrus.Logger) PaymentMethods {
	return &paymentMethods{
		redis:   redis,
		sql:     sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (m *paymentMethods) New() PaymentMethods {
	return NewPaymentMethods(m.redis, m.sql, m.log)
}

func (m *paymentMethods) Create(ctx context.Context, pm models.PaymentMethod) error {
	if err := m.sql.New().Insert(ctx, pm); err != nil {
		return err
	}

	return nil
}

func (m *paymentMethods) Get(ctx context.Context) (*models.PaymentMethod, error) {
	res, err := m.redis.GetByID(ctx, m.filters["id"].(string))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			res = nil
		} else {
			m.log.WithField("redis", err).Error("error getting payment method from cache")
		}
	} else {
		return res, nil
	}

	res, err = m.sql.New().Filter(m.filters).Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *paymentMethods) Select(ctx context.Context) ([]models.PaymentMethod, error) {
	res, err := m.sql.New().Filter(m.filters).Select(ctx)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	return res, nil
}

func (m *paymentMethods) DeleteByID(ctx context.Context) error {
	if err := m.sql.New().Filter(m.filters).Delete(ctx); err != nil {
		return err
	}

	if err := m.redis.Delete(ctx, m.filters["id"].(string)); err != nil {
		m.log.WithField("redis", err).Error("error deleting payment method from cache")
	}

	return nil
}

func (m *paymentMethods) DeleteByUserID(ctx context.Context) error {
	if err := m.sql.New().Filter(m.filters).Delete(ctx); err != nil {
		return err
	}

	if err := m.redis.DeleteByUserID(ctx, m.filters["user_id"].(string)); err != nil {
		m.log.WithField("redis", err).Error("error deleting payment methods from cache")
	}

	return nil
}

func (m *paymentMethods) Count(ctx context.Context) (int, error) {
	return m.sql.New().Filter(m.filters).Count(ctx)
}

func (m *paymentMethods) SelectForUser(ctx context.Context) ([]models.PaymentMethod, error) {
	res, err := m.redis.GetByUserID(ctx, m.filters["user_id"].(string))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			res = nil
		} else {
			m.log.WithField("redis", err).Error("error getting payment methods from cache")
		}
	} else {
		return res, nil
	}

	res, err = m.sql.New().Filter(m.filters).Select(ctx)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	go func() {
		for _, ses := range res {
			err = m.redis.Add(ctx, ses)
			if err != nil {
				m.log.WithError(err).Error("error adding session to Redis")
			}
		}
	}()

	return res, nil
}

func (m *paymentMethods) Filter(filters map[string]any) PaymentMethods {
	m.filters = filters
	return m
}

func (m *paymentMethods) Page(limit, offset uint64) PaymentMethods {
	m.limit = int64(limit)
	m.skip = int64(offset)
	return m
}
