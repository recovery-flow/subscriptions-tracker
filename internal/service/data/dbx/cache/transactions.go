package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transactions interface {
	Mew() Transactions

	Add(ctx context.Context, tx models.Transaction) error
	Get(ctx context.Context) (*models.Transaction, error)
	Select(ctx context.Context) ([]models.Transaction, error)
	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) error

	Filter(ctx context.Context, filters map[string]any) Transactions
	Limit(limit int) Transactions
	Skip(skip int) Transactions
}

type transactions struct {
	client   *redis.Client
	LifeTime time.Duration
	filters  map[string]string
	skip     int
	limit    int
}

func NewTransactions(client *redis.Client, lifeTime time.Duration) Transactions {
	return &transactions{
		client:   client,
		LifeTime: lifeTime,
		filters:  make(map[string]string),
		skip:     0,
		limit:    0,
	}
}

func (t *transactions) Mew() Transactions {
	return &transactions{
		client:   t.client,
		LifeTime: t.LifeTime,
		filters:  make(map[string]string),
		skip:     0,
		limit:    0,
	}
}

func (t *transactions) Filter(ctx context.Context, newFilters map[string]any) Transactions {
	copied := make(map[string]string)
	for k, v := range t.filters {
		copied[k] = v
	}
	for k, v := range newFilters {
		copied[k] = fmt.Sprintf("%v", v)
	}
	return &transactions{
		client:   t.client,
		LifeTime: t.LifeTime,
		filters:  copied,
		skip:     t.skip,
		limit:    t.limit,
	}
}

func (t *transactions) Limit(limit int) Transactions {
	return &transactions{
		client:   t.client,
		LifeTime: t.LifeTime,
		filters:  t.filters,
		skip:     t.skip,
		limit:    limit,
	}
}

func (t *transactions) Skip(skip int) Transactions {
	return &transactions{
		client:   t.client,
		LifeTime: t.LifeTime,
		filters:  t.filters,
		skip:     skip,
		limit:    t.limit,
	}
}

func (t *transactions) Add(ctx context.Context, tx models.Transaction) error {
	key := fmt.Sprintf("tx:id:%s", tx.ID.Hex())
	data := map[string]interface{}{
		"id":             tx.ID.Hex(),
		"amount":         tx.Amount,
		"currency":       tx.Currency,
		"status":         string(tx.Status),
		"payment_method": tx.PaymentMethod,
		"prov_tx_id":     tx.ProvTxID,
		"created_at":     tx.CreatedAt.Time().Format(time.RFC3339),
	}
	if tx.UserID != nil {
		data["user_id"] = tx.UserID.Hex()
	}
	if tx.PlanID != nil {
		data["plan_id"] = tx.PlanID.Hex()
	}
	if tx.SubID != nil {
		data["sub_id"] = tx.SubID.Hex()
	}
	if err := t.client.HSet(ctx, key, data).Err(); err != nil {
		return fmt.Errorf("error adding transaction to Redis: %w", err)
	}
	if t.LifeTime > 0 {
		_ = t.client.Expire(ctx, key, t.LifeTime).Err()
	}
	return nil
}

func (t *transactions) Select(ctx context.Context) ([]models.Transaction, error) {
	keys, err := t.client.Keys(ctx, "tx:id:*").Result()
	if err != nil {
		return nil, fmt.Errorf("error retrieving keys: %w", err)
	}
	var result []models.Transaction
	for _, key := range keys {
		data, err := t.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if t.matchesFilters(data) {
			tx, err := mapToTransaction(data)
			if err == nil {
				result = append(result, *tx)
			}
		}
	}
	// Применяем Skip
	if t.skip > 0 && t.skip < len(result) {
		result = result[t.skip:]
	} else if t.skip >= len(result) {
		result = []models.Transaction{}
	}
	// Применяем Limit
	if t.limit > 0 && t.limit < len(result) {
		result = result[:t.limit]
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no matching transactions found")
	}
	return result, nil
}

func (t *transactions) Get(ctx context.Context) (*models.Transaction, error) {
	list, err := t.Select(ctx)
	if err != nil {
		return nil, err
	}
	return &list[0], nil
}

func (t *transactions) DeleteOne(ctx context.Context) error {
	keys, err := t.client.Keys(ctx, "tx:id:*").Result()
	if err != nil {
		return fmt.Errorf("error retrieving keys: %w", err)
	}
	for _, key := range keys {
		data, err := t.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if t.matchesFilters(data) {
			if err := t.client.Del(ctx, key).Err(); err != nil {
				return fmt.Errorf("error deleting transaction: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no matching transaction found for deletion")
}

func (t *transactions) DeleteMany(ctx context.Context) error {
	keys, err := t.client.Keys(ctx, "tx:id:*").Result()
	if err != nil {
		return fmt.Errorf("error retrieving keys: %w", err)
	}
	var delErr error
	for _, key := range keys {
		data, err := t.client.HGetAll(ctx, key).Result()
		if err != nil || len(data) == 0 {
			continue
		}
		if t.matchesFilters(data) {
			if err := t.client.Del(ctx, key).Err(); err != nil {
				delErr = err
			}
		}
	}
	return delErr
}

func (t *transactions) matchesFilters(data map[string]string) bool {
	for fk, fv := range t.filters {
		if data[fk] != fv {
			return false
		}
	}
	return true
}

func mapToTransaction(data map[string]string) (*models.Transaction, error) {
	idHex, ok := data["id"]
	if !ok {
		return nil, fmt.Errorf("id not found in data")
	}
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, err
	}
	amountStr, ok := data["amount"]
	if !ok {
		return nil, fmt.Errorf("amount not found")
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return nil, err
	}
	createdAt, err := time.Parse(time.RFC3339, data["created_at"])
	if err != nil {
		return nil, err
	}
	var userID *primitive.ObjectID
	if uid, ok := data["user_id"]; ok && uid != "" {
		tmp, err := primitive.ObjectIDFromHex(uid)
		if err == nil {
			userID = &tmp
		}
	}
	var planID *primitive.ObjectID
	if pid, ok := data["plan_id"]; ok && pid != "" {
		tmp, err := primitive.ObjectIDFromHex(pid)
		if err == nil {
			planID = &tmp
		}
	}
	var subID *primitive.ObjectID
	if sid, ok := data["sub_id"]; ok && sid != "" {
		tmp, err := primitive.ObjectIDFromHex(sid)
		if err == nil {
			subID = &tmp
		}
	}
	return &models.Transaction{
		ID:            id,
		UserID:        userID,
		PlanID:        planID,
		SubID:         subID,
		Amount:        amount,
		Currency:      data["currency"],
		Status:        models.TrStatus(data["status"]),
		PaymentMethod: data["payment_method"],
		ProvTxID:      data["prov_tx_id"],
		CreatedAt:     primitive.NewDateTimeFromTime(createdAt),
	}, nil
}
