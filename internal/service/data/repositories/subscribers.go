package repositories

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/mongodb"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/data/dbx/redisdb"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscribers interface {
	Create(r *http.Request, userID uuid.UUID, PlanID primitive.ObjectID) (*models.Subscriber, error)
}

type subscribers struct {
	redis *redisdb.Subscribers
	mongo *mongodb.Subscribers
	log   *log.Logger //todo remove and remade
}

func NewSubscribers(redisAddr, redisPassword string, redisDB int, mongoUrl, mongoName string) (Subscribers, error) {
	redisRepo := redisdb.NewSubscribers(redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	}))
	mongo, err := mongodb.NewSubscribers(mongoUrl, mongoName, "subscribers")
	if err != nil {
		return nil, err
	}
	return &subscribers{
		redis: &redisRepo,
		mongo: &mongo,
		log:   log.Default(),
	}, nil
}

func (s *subscribers) Create(r *http.Request, userID uuid.UUID, PlanID primitive.ObjectID) (*models.Subscriber, error) {
	return &models.Subscriber{}, nil
}
