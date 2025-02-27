package infra

import (
	"github.com/recovery-flow/rerabbit"
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/sirupsen/logrus"
)

type Infra struct {
	Rabbit rerabbit.RabbitBroker
}

func NewInfra(cfg *config.Config, log *logrus.Logger) (*Infra, error) {
	eve, err := rerabbit.NewBroker(cfg.Rabbit.URL)
	if err != nil {
		return nil, err
	}

	return &Infra{
		Rabbit: eve,
	}, nil
}
