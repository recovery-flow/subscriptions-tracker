package infra

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/data"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events/producer"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/listener"
	"github.com/sirupsen/logrus"
)

type Infra struct {
	Producer producer.Producer
	Listener listener.Listener

	Data *data.Data
}

func NewInfra(cfg *config.Config, log *logrus.Logger) (*Infra, error) {
	prd := producer.NewProducer(cfg)

	db, err := data.NewData(cfg)
	if err != nil {
		return nil, err
	}

	return &Infra{
		Producer: prd,
		Data:     db,
	}, nil
}
