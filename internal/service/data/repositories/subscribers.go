package repositories

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/cache"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/mongodb"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscribers interface {
	//METHODS...
}

type subscribers struct {
	redis *cache.Subscribers
	mongo *mongodb.Subscribers
}

func NewSubscribers(cfg config.Config) (Subscribers, error) {
	redisRepo := cache.NewSubscribers(redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Redis.Addr,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	}), 15*time.Minute)
	mongo, err := mongodb.NewSubscribers(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, "subscribers")
	if err != nil {
		return nil, err
	}
	return &subscribers{
		redis: &redisRepo,
		mongo: &mongo,
	}, nil
}

func (s *subscribers) Create(r *http.Request, userID uuid.UUID, PlanID primitive.ObjectID) (*models.Subscriber, error) {
	return &models.Subscriber{}, nil
}
