package config

import (
	"github.com/recovery-flow/cifra-rabbit"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "server"
)

type Service struct {
	Config       *Config
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *Config, logger *logrus.Logger) (*Service, error) {
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
