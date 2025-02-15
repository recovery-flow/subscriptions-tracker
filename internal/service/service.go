package service

import (
	"github.com/recovery-flow/cifra-rabbit"
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Config       *config.Config
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *config.Config, logger *logrus.Logger) (*Service, error) {
	broker, err := cifra_rabbit.NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}

	return &Service{
		Config: cfg,
		Logger: logger,
		Broker: broker,
	}, nil
}
