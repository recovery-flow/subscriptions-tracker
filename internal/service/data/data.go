package data

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/repositories"
)

type Config struct {
	Mongo struct {
		Uri    string
		DbName string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}

type Data struct {
	Subscribers repositories.Subscribers
}

func NewDataBase(cfg Config) (*Data, error) {
	return &Data{
		Subscribers: repositories.NewSubscribers(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, cfg.Mongo.Uri, cfg.Mongo.DbName), //todo
	}, nil
}
