package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Subscribers interface {
	New() Subscribers

	Insert(ctx context.Context, sub models.Subscriber) (*models.Subscriber, error)

	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Subscriber, error)
	Get(ctx context.Context) (*models.Subscriber, error)

	FilterStrict(filters map[string]any) Subscribers

	UpdateOne(ctx context.Context, fields map[string]any) (*models.Subscriber, error)
	UpdateMany(ctx context.Context, fields map[string]any) (int64, error)

	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) (int64, error)

	SortBy(field string, ascending bool) Subscribers
	Limit(limit int64) Subscribers
	Skip(skip int64) Subscribers
}

type subscribers struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewSubscribers(uri, dbName, collectionName string) (Subscribers, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(collectionName)

	return &subscribers{
		client:     client,
		database:   database,
		collection: coll,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (s *subscribers) New() Subscribers {
	return &subscribers{
		client:     s.client,
		database:   s.database,
		collection: s.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (s *subscribers) Insert(ctx context.Context, sub models.Subscriber) (*models.Subscriber, error) {
	sub.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	sub.ID = primitive.NewObjectID()

	_, err := s.collection.InsertOne(ctx, sub)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (s *subscribers) Count(ctx context.Context) (int64, error) {
	return s.collection.CountDocuments(ctx, s.filters)
}

func (s *subscribers) Select(ctx context.Context) ([]models.Subscriber, error) {
	opts := options.Find().SetSort(s.sort).SetLimit(s.limit).SetSkip(s.skip)
	cursor, err := s.collection.Find(ctx, s.filters, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to select subscribers: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var subs []models.Subscriber
	if err = cursor.All(ctx, &subs); err != nil {
		return nil, fmt.Errorf("failed to decode subscribers: %w", err)
	}
	return subs, nil
}

func (s *subscribers) Get(ctx context.Context) (*models.Subscriber, error) {
	var sub models.Subscriber
	if err := s.collection.FindOne(ctx, s.filters).Decode(&sub); err != nil {
		return nil, fmt.Errorf("failed to get subscribers: %w", err)
	}
	return &sub, nil
}

func (s *subscribers) FilterStrict(filters map[string]any) Subscribers {
	var validFilters = map[string]bool{
		"_id":           true,
		"user_id":       true,
		"plan_id":       true,
		"streak_months": true,
		"status":        true,
		"start_at":      true,
		"end_at":        true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		s.filters[field] = value
	}
	return s
}

func (s *subscribers) UpdateOne(ctx context.Context, fields map[string]any) (*models.Subscriber, error) {
	validFields := map[string]bool{
		"plan_id":  true,
		"status":   true,
		"start_at": true,
		"end_at":   true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if !validFields[key] {
			continue
		}
		if value == nil {
			continue
		}
		updateFields[key] = value
	}

	updateFields["updated_at"] = primitive.NewDateTimeFromTime(time.Now().UTC())

	_, err := s.collection.UpdateOne(ctx, s.filters, bson.M{"$set": updateFields})
	if err != nil {
		return nil, fmt.Errorf("failed to update subscribers: %w", err)
	}

	return s.Get(ctx)
}

func (s *subscribers) UpdateMany(ctx context.Context, fields map[string]any) (int64, error) {
	validFields := map[string]bool{
		"plan_id":  true,
		"status":   true,
		"start_at": true,
		"end_at":   true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if !validFields[key] {
			continue
		}
		if value == nil {
			continue
		}
		updateFields[key] = value
	}

	updateFields["updated_at"] = primitive.NewDateTimeFromTime(time.Now().UTC())

	res, err := s.collection.UpdateMany(ctx, s.filters, bson.M{"$set": updateFields})
	if err != nil {
		return 0, fmt.Errorf("failed to update subscribers: %w", err)
	}

	return res.ModifiedCount, nil
}

func (s *subscribers) DeleteOne(ctx context.Context) error {
	_, err := s.collection.DeleteOne(ctx, s.filters)
	if err != nil {
		return fmt.Errorf("failed to delete subscribers: %w", err)
	}
	return nil
}

func (s *subscribers) DeleteMany(ctx context.Context) (int64, error) {
	res, err := s.collection.DeleteMany(ctx, s.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete subscribers: %w", err)
	}
	return res.DeletedCount, nil
}

func (s *subscribers) SortBy(field string, ascending bool) Subscribers {
	var order int
	if ascending {
		order = 1
	} else {
		order = -1
	}
	s.sort = bson.D{{field, order}}
	return s
}

func (s *subscribers) Limit(limit int64) Subscribers {
	s.limit = limit
	return s
}

func (s *subscribers) Skip(skip int64) Subscribers {
	s.skip = skip
	return s
}
