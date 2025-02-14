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
	MongoDB      *nosql.Repo
	Cloud        *cloud.Repo
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	MongoDb, err := nosql.NewRepositoryNoSql(cfg.Mongo.URI, cfg.Mongo.DbName)
	if err != nil {
		return nil, err
	}
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)
	Storage, err := cloud.NewRepositoryCloud(cfg.Cloud.CloudName, cfg.Cloud.APIKey, cfg.Cloud.APISecret)
	if err != nil {
		return nil, err
	}

	broker, err := cifra_rabbit.NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}

	return &Service{
		Config:       cfg,
		MongoDB:      MongoDb,
		Cloud:        Storage,
		Logger:       logger,
		TokenManager: &TokenManager,
		Broker:       broker,
	}, nil
}
