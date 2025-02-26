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

type SubscriptionPlans interface {
	New() SubscriptionPlans

	Insert(ctx context.Context, plan models.SubscriptionPlan) (*models.SubscriptionPlan, error)

	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.SubscriptionPlan, error)
	Get(ctx context.Context) (*models.SubscriptionPlan, error)

	FilterStrict(filters map[string]interface{}) SubscriptionPlans
	FilterDate(filters map[string]interface{}, after bool) SubscriptionPlans

	UpdateOne(ctx context.Context, fields map[string]any) (*models.SubscriptionPlan, error)
	UpdateMany(ctx context.Context, fields map[string]any) (int64, error)

	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) (int64, error)

	SortBy(field string, ascending bool) SubscriptionPlans
	Limit(limit int64) SubscriptionPlans
	Skip(skip int64) SubscriptionPlans
}

type subscriptionPlans struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewSubscriptionPlans(uri, dbName, collectionName string) (SubscriptionPlans, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(collectionName)

	return &subscriptionPlans{
		client:     client,
		database:   database,
		collection: coll,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (s *subscriptionPlans) New() SubscriptionPlans {
	return &subscriptionPlans{
		client:     s.client,
		database:   s.database,
		collection: s.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (s *subscriptionPlans) Insert(ctx context.Context, plan models.SubscriptionPlan) (*models.SubscriptionPlan, error) {
	plan.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	plan.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(ctx, plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func (s *subscriptionPlans) Count(ctx context.Context) (int64, error) {
	return s.collection.CountDocuments(ctx, s.filters)
}

func (s *subscriptionPlans) Select(ctx context.Context) ([]models.SubscriptionPlan, error) {
	cursor, err := s.collection.Find(ctx, s.filters)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var plans []models.SubscriptionPlan
	if err = cursor.All(ctx, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

func (s *subscriptionPlans) Get(ctx context.Context) (*models.SubscriptionPlan, error) {
	var plan models.SubscriptionPlan
	err := s.collection.FindOne(ctx, s.filters).Decode(&plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func (s *subscriptionPlans) FilterStrict(filters map[string]interface{}) SubscriptionPlans {
	var validFilters = map[string]bool{
		"_id":      true,
		"name":     true,
		"price":    true,
		"currency": true,
		"pay_freq": true,
		"status":   true,
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

func (s *subscriptionPlans) FilterDate(filters map[string]interface{}, after bool) SubscriptionPlans {
	var validFilters = map[string]bool{
		"created_at":  true,
		"updated_at":  true,
		"canceled_at": true,
	}

	var op string
	if after {
		op = "$gt"
	} else {
		op = "$lt"
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}

		var t time.Time
		switch val := value.(type) {
		case time.Time:
			t = val
		case *time.Time:
			t = *val
		case string:
			parsed, err := time.Parse(time.RFC3339, val)
			if err != nil {
				continue
			}
			t = parsed
		default:
			continue
		}

		s.filters[field] = bson.M{op: t}
	}

	return s
}

func (s *subscriptionPlans) UpdateOne(ctx context.Context, fields map[string]any) (*models.SubscriptionPlan, error) {
	var ValidFields = map[string]bool{
		"name":     true,
		"title":    true,
		"desc":     true,
		"price":    true,
		"currency": true,
		"pay_freq": true,
	}

	updateFields := bson.M{}
	for field, value := range fields {
		if !ValidFields[field] {
			continue
		}
		if value == nil {
			continue
		}
		updateFields[field] = value
	}

	updateFields["updated_at"] = primitive.NewDateTimeFromTime(time.Now())

	_, err := s.collection.UpdateOne(ctx, s.filters, bson.M{"$set": updateFields})
	if err != nil {
		return nil, err
	}

	return s.Get(ctx)
}

func (s *subscriptionPlans) UpdateMany(ctx context.Context, fields map[string]any) (int64, error) {
	var ValidFields = map[string]bool{
		"name":     true,
		"title":    true,
		"desc":     true,
		"price":    true,
		"currency": true,
		"pay_freq": true,
	}

	updateFields := bson.M{}
	for field, value := range fields {
		if !ValidFields[field] {
			continue
		}
		if value == nil {
			continue
		}
		updateFields[field] = value
	}

	updateFields["updated_at"] = primitive.NewDateTimeFromTime(time.Now())

	res, err := s.collection.UpdateMany(ctx, s.filters, bson.M{"$set": updateFields})
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (s *subscriptionPlans) DeleteOne(ctx context.Context) error {
	_, err := s.collection.DeleteOne(ctx, s.filters)
	if err != nil {
		return err
	}

	return nil
}

func (s *subscriptionPlans) DeleteMany(ctx context.Context) (int64, error) {
	res, err := s.collection.DeleteMany(ctx, s.filters)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (s *subscriptionPlans) SortBy(field string, ascending bool) SubscriptionPlans {
	var order int
	if ascending {
		order = 1
	} else {
		order = -1
	}

	s.sort = bson.D{{field, order}}

	return s
}

func (s *subscriptionPlans) Limit(limit int64) SubscriptionPlans {
	s.limit = limit
	return s
}

func (s *subscriptionPlans) Skip(skip int64) SubscriptionPlans {
	s.skip = skip
	return s
}
