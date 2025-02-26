package service

import (
	"github.com/recovery-flow/rerabbit"
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Config       *config.Config
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Rabbit       *rerabbit.RabbitBroker
	DB           *data.Data
}

func NewService(cfg *config.Config, logger *logrus.Logger) (*Service, error) {
	rabbit, err := rerabbit.NewBroker(cfg.Rabbit.URL)
	if err != nil {
		return nil, err
	}

	database, err := data.NewDataBase(cfg)
	if err != nil {
		return nil, err
	}

	tm := tokens.NewTokenManager(cfg.Database.Redis.Addr, cfg.Database.Redis.Password, cfg.Database.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)

	return &Service{
		Config:       cfg,
		Logger:       logger,
		Rabbit:       &rabbit,
		DB:           database,
		TokenManager: &tm,
	}, nil
}
