package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
)

type PayMethods interface {
	Add(ctx context.Context, payMethod models.PaymentMethod) error
	GetByID(ctx context.Context, methodID string) (*models.PaymentMethod, error)
	GetByUserID(ctx context.Context, userID string) ([]models.PaymentMethod, error)
	Delete(ctx context.Context, methodID string) error
	DeleteByUserID(ctx context.Context, userID string) error
}

type payMethods struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewPayMethods(client *redis.Client, lifetime time.Duration) PayMethods {
	return &payMethods{
		client:   client,
		lifeTime: lifetime,
	}
}

func (m *payMethods) Add(ctx context.Context, payMethod models.PaymentMethod) error {
	payMethodKey := fmt.Sprintf("payment_method:id:%s", payMethod.ID.String())
	userIDKey := fmt.Sprintf("payment_method:user_id:%s", payMethod.UserID.String())

	data := map[string]interface{}{
		"user_id":        payMethod.UserID.String(),
		"type":           payMethod.Type,
		"provider_token": payMethod.ProviderToken,
		"is_default":     payMethod.IsDefault,
		"created_at":     payMethod.CreatedAt.Format(time.RFC3339),
	}

	if err := m.client.HSet(ctx, payMethodKey, data).Err(); err != nil {
		return fmt.Errorf("error adding payment method to Redis: %w", err)
	}

	if err := m.client.SAdd(ctx, userIDKey, payMethod.ID.String()).Err(); err != nil {
		return fmt.Errorf("error adding payment method ID to user set: %w", err)
	}

	if m.lifeTime > 0 {
		pipe := m.client.Pipeline()
		keys := []string{payMethodKey, userIDKey}
		for _, key := range keys {
			pipe.Expire(ctx, key, m.lifeTime)
		}
		_, err := pipe.Exec(ctx)
		if err != nil && err != redis.Nil {
			return fmt.Errorf("error setting expiration for keys: %w", err)
		}
	}

	return nil
}

func (m *payMethods) GetByID(ctx context.Context, methodID string) (*models.PaymentMethod, error) {
	IDKey := fmt.Sprintf("payment_method:id:%s", methodID)
	vals, err := m.client.HGetAll(ctx, IDKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting payment method from Redis: %w", err)
	}

	if len(vals) == 0 {
		return nil, fmt.Errorf("pyment_method not found, id=%s", methodID)
	}

	return parsePayMethod(IDKey, vals)
}

func (m *payMethods) GetByUserID(ctx context.Context, userID string) ([]models.PaymentMethod, error) {
	userIDKey := fmt.Sprintf("payment_method:user_id:%s", userID)

	IDs, err := m.client.SMembers(ctx, userIDKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting payment method IDs from Redis: %w", err)
	}

	if len(IDs) == 0 {
		return nil, nil
	}

	var payMethods []models.PaymentMethod
	for _, ID := range IDs {
		vals, err := m.client.HGetAll(ctx, ID).Result()
		if err != nil {
			return nil, fmt.Errorf("error getting payment method: %w", err)
		}
		method, err := parsePayMethod(ID, vals)
		if err != nil {
			return nil, fmt.Errorf("error parsing payment method: %w", err)
		}
		payMethods = append(payMethods, *method)
	}

	return payMethods, nil
}

func (m *payMethods) Delete(ctx context.Context, methodID string) error {
	payMethodKey := fmt.Sprintf("payment_method:id:%s", methodID)

	userID, err := m.client.HGet(ctx, payMethodKey, "user_id").Result()
	if err != nil {
		return fmt.Errorf("error getting user_id from Redis: %w", err)
	}

	err = m.client.Del(ctx, payMethodKey).Err()
	if err != nil {
		return fmt.Errorf("error deleting payment method from Redis: %w", err)
	}

	userIDKey := fmt.Sprintf("payment_method:user_id:%s", userID)
	err = m.client.SRem(ctx, userIDKey, methodID).Err()
	if err != nil {
		return fmt.Errorf("error removing payment method methodID from user set: %w", err)
	}

	return nil
}

func (m *payMethods) DeleteByUserID(ctx context.Context, userID string) error {
	userIDKey := fmt.Sprintf("payment_method:user_id:%s", userID)

	IDs, err := m.client.SMembers(ctx, userIDKey).Result()
	if err != nil {
		return fmt.Errorf("error getting payment method IDs from Redis: %w", err)
	}

	if len(IDs) == 0 {
		return nil
	}

	pipe := m.client.Pipeline()
	for _, ID := range IDs {
		payMethodKey := fmt.Sprintf("payment_method:id:%s", ID)
		pipe.Del(ctx, payMethodKey)
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("error deleting payment methods: %w", err)
	}

	err = m.client.Del(ctx, userIDKey).Err()
	if err != nil {
		return fmt.Errorf("error deleting user set: %w", err)
	}

	return nil
}

func parsePayMethod(payMethodID string, vals map[string]string) (*models.PaymentMethod, error) {
	createdAt, err := time.Parse(time.RFC3339, vals["created_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing created_at: %w", err)
	}

	ID, err := uuid.Parse(payMethodID)
	if err != nil {
		return nil, fmt.Errorf("error parsing AccountID: %w", err)
	}

	userID, err := uuid.Parse(vals["user_id"])
	if err != nil {
		return nil, fmt.Errorf("error parsing userID: %w", err)
	}

	typeMethod, err := models.ParsePayType(vals["type"])
	if err != nil {
		return nil, fmt.Errorf("error parsing type: %w", err)
	}

	account := &models.PaymentMethod{
		ID:            ID,
		UserID:        userID,
		Type:          typeMethod,
		ProviderToken: vals["provider_token"],
		IsDefault:     vals["is_default"] == "true",
		CreatedAt:     createdAt,
	}

	return account, nil
}
