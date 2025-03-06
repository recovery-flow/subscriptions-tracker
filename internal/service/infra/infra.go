package infra

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/sirupsen/logrus"
)

type Infra struct {
	Kafka events.Kafka

	Data *data.Data
}

func NewInfra(cfg *config.Config, log *logrus.Logger) (*Infra, error) {
	eve := events.NewBroker(cfg)

	db, err := data.NewData(cfg)
	if err != nil {
		return nil, err
	}

	return &Infra{
		Kafka: eve,
		Data:  db,
	}, nil
}
