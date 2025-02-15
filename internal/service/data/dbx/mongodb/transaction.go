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

type Transactions interface {
	New() Transactions

	Insert(ctx context.Context, sub models.Transaction) (*models.Transaction, error)

	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Transaction, error)
	Get(ctx context.Context) (*models.Transaction, error)

	FilterStrict(filters map[string]any) Transactions
	FilterDate(filters map[string]any, after bool) Transactions

	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) (int64, error)

	SortBy(field string, ascending bool) Transactions
	Limit(limit int64) Transactions
	Skip(skip int64) Transactions
}

type transactions struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewTransactions(uri, dbName, collectionName string) (Transactions, error) {
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

func (t *transactions) New() Transactions {
	return &transactions{
		client:     t.client,
		database:   t.database,
		collection: t.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (t *transactions) Insert(ctx context.Context, tr models.Transaction) (*models.Transaction, error) {
	tr.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	tr.ID = primitive.NewObjectID()

	_, err := t.collection.InsertOne(ctx, tr)
	if err != nil {
		return nil, err
	}

	return &tr, nil
}

func (t *transactions) Count(ctx context.Context) (int64, error) {
	return t.collection.CountDocuments(ctx, t.filters)
}

func (t *transactions) Select(ctx context.Context) ([]models.Transaction, error) {
	opts := options.Find().SetSort(t.sort).SetLimit(t.limit).SetSkip(t.skip)
	cursor, err := t.collection.Find(ctx, t.filters, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var trs []models.Transaction
	if err = cursor.All(ctx, &trs); err != nil {
		return nil, err
	}
	return trs, nil
}

func (t *transactions) Get(ctx context.Context) (*models.Transaction, error) {
	var tr models.Transaction
	err := t.collection.FindOne(ctx, t.filters).Decode(&tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

func (t *transactions) FilterStrict(filters map[string]any) Transactions {
	var validFilters = map[string]bool{
		"_id":            true,
		"user_id":        true,
		"plan_id":        true,
		"sub_id":         true,
		"amount":         true,
		"currency":       true,
		"status":         true,
		"payment_method": true,
		"prov_tx_id":     true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		t.filters[field] = value
	}
	return t
}

func (t *transactions) FilterDate(filters map[string]any, after bool) Transactions {
	var validFilters = map[string]bool{
		"created_at": true,
	}

	var op string
	if after {
		op = "$gte"
	} else {
		op = "$lte"
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}

		var ti time.Time
		switch val := value.(type) {
		case time.Time:
			ti = val
		case *time.Time:
			ti = *val
		case string:
			parsed, err := time.Parse(time.RFC3339, val)
			if err != nil {
				continue
			}
			ti = parsed
		default:
			continue
		}

		t.filters[field] = bson.M{op: ti}
	}
	return t
}

func (t *transactions) DeleteOne(ctx context.Context) error {
	_, err := t.collection.DeleteOne(ctx, t.filters)
	return err
}

func (t *transactions) DeleteMany(ctx context.Context) (int64, error) {
	res, err := t.collection.DeleteMany(ctx, t.filters)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (t *transactions) SortBy(field string, ascending bool) Transactions {
	order := 1
	if !ascending {
		order = -1
	}
	t.sort = append(t.sort, bson.E{Key: field, Value: order})
	return t
}

func (t *transactions) Limit(limit int64) Transactions {
	t.limit = limit
	return t
}

func (t *transactions) Skip(skip int64) Transactions {
	t.skip = skip
	return t
}
