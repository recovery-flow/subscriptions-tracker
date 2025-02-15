package service

import (
	"github.com/recovery-flow/rerabbit"
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Config *config.Config
	Logger *logrus.Logger
	Rabbit rerabbit.RabbitBroker
	DB     *data.Data
}

func NewService(cfg *config.Config, logger *logrus.Logger) (*Service, error) {
	rabbit, err := rerabbit.NewBroker(cfg.Rabbit.URL)
	if err != nil {
		return nil, err
	}

	database, err := data.NewDataBase(data.Config{
		Mongo: struct {
			Uri    string
			DbName string
		}{
			Uri:    cfg.Database.Mongo.URI,
			DbName: cfg.Database.Mongo.Name,
		},
		Redis: struct {
			Addr     string
			Password string
			DB       int
		}{
			Addr:     cfg.Database.Redis.Addr,
			Password: cfg.Database.Redis.Password,
			DB:       cfg.Database.Redis.DB,
		},
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		Config: cfg,
		Logger: logger,
		Rabbit: rabbit,
		DB:     database,
	}, nil
}
