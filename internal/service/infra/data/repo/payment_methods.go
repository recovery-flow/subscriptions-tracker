package repo

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data/repo/sqldb"
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
	sql     sqldb.PaymentMethods
	filters map[string]any
	limit   int64
	skip    int64

	log *logrus.Logger
}

func NewPaymentMethods(sql sqldb.PaymentMethods, log *logrus.Logger) PaymentMethods {
	return &paymentMethods{
		sql:     sql,
		filters: make(map[string]any),
		limit:   0,
		skip:    0,

		log: log,
	}
}

func (m *paymentMethods) New() PaymentMethods {
	return NewPaymentMethods(m.sql, m.log)
}

func (m *paymentMethods) Create(ctx context.Context, pm models.PaymentMethod) error {
	if err := m.sql.New().Insert(ctx, pm); err != nil {
		return err
	}

	return nil
}

func (m *paymentMethods) Get(ctx context.Context) (*models.PaymentMethod, error) {
	return m.sql.New().Filter(m.filters).Get(ctx)
}

func (m *paymentMethods) Select(ctx context.Context) ([]models.PaymentMethod, error) {
	return m.sql.New().Filter(m.filters).Select(ctx)
}

func (m *paymentMethods) DeleteByID(ctx context.Context) error {
	return m.sql.New().Filter(m.filters).Delete(ctx)

}

func (m *paymentMethods) DeleteByUserID(ctx context.Context) error {
	return m.sql.New().Filter(m.filters).Delete(ctx)
}

func (m *paymentMethods) Count(ctx context.Context) (int, error) {
	return m.sql.New().Filter(m.filters).Count(ctx)
}

func (m *paymentMethods) SelectForUser(ctx context.Context) ([]models.PaymentMethod, error) {
	return m.sql.New().Filter(m.filters).Select(ctx)
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
