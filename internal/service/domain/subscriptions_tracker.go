package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type SubscriptionsTracker interface {
	GetSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error)
	ActivateSubscription(ctx context.Context, UserID, PlanID, PaymentMethodID uuid.UUID) (*models.Subscription, error)
	DeactivateSubscription(ctx context.Context, UserID uuid.UUID) error
	CanceledSubscription(ctx context.Context, UserID uuid.UUID) error

	AddPaymentMethod(ctx context.Context, userID uuid.UUID, token string, payType string) error
	DeletePaymentMethod(ctx context.Context, userID uuid.UUID, paymentMethodID uuid.UUID) error
	GetPaymentMethod(ctx context.Context, ID uuid.UUID) (*models.PaymentMethod, error)
	GetUserPaymentMethods(ctx context.Context, userID uuid.UUID) ([]models.PaymentMethod, error)
	SetPaymentMethodAsDefault(ctx context.Context, userID, paymentMethodID uuid.UUID) error
}

func (d *domain) ActivateSubscription(ctx context.Context, UserID, PlanID, PaymentMethodID uuid.UUID) (*models.Subscription, error) {
	var subscription *models.Subscription
	err := d.Infra.Data.SQL.Subscriptions.Transaction(func(ctx context.Context) error {
		sType, err := d.Infra.Data.SQL.Types.New().Filter(map[string]any{
			"id": PlanID.String(),
		}).Get(ctx)
		if err != nil {
			return err
		}
		if sType == nil {
			return fmt.Errorf("subscription type not found %s", PlanID)
		}

		if sType.Status != models.StatusTypeActive {
			return fmt.Errorf("subscription type is not active %s", PlanID)
		}

		//Get the Subscription Plan

		sPlan, err := d.Infra.Data.SQL.Plans.New().Filter(map[string]any{
			"id": PlanID.String(),
		}).Get(ctx)
		if err != nil {
			return err
		}
		if sPlan == nil {
			return fmt.Errorf("subscription plan not found %s", PlanID)
		}

		//Get the Payment Method

		pMethod, err := d.Infra.Data.SQL.PaymentMethods.New().Filter(map[string]any{
			"id": PaymentMethodID.String(),
		}).Get(ctx)
		if err != nil {
			return err
		}

		if pMethod.UserID != UserID {
			return fmt.Errorf("payment method does not belong to user %s", PaymentMethodID)
		}

		//TODO integrate payment for the subscription

		//Main logic for the subscription activation

		var endDate time.Time
		uInterval := sPlan.BillingIntervalUnit
		interval := sPlan.BillingInterval

		switch uInterval {
		case models.IntervalOnce:
			endDate = time.Now().AddDate(100, 0, 0)
		case models.IntervalDay:
			endDate = time.Now().AddDate(0, 0, 1*int(interval))
		case models.IntervalWeek:
			endDate = time.Now().AddDate(0, 0, 7*int(interval))
		case models.IntervalMonth:
			endDate = time.Now().AddDate(0, int(interval), 0)
		case models.IntervalYear:
			endDate = time.Now().AddDate(int(interval), 0, 0)
		default:
			return fmt.Errorf("invalid billing interval unit %s || %d", uInterval, interval)
		}

		subscription = &models.Subscription{
			UserID:          UserID,
			PlanID:          PlanID,
			PaymentMethodID: PaymentMethodID,
			Status:          models.SubscriptionStatusActive,
			Availability:    models.SubPlanAvailable,
			StartDate:       time.Now().UTC(),
			EndDate:         endDate,
		}

		if err := d.Infra.Data.SQL.Subscriptions.New().Insert(ctx, subscription); err != nil {
			return err
		}

		//TODO to supplement the work with Kafka
		//if err := d.Infra.Kafka.SubscriptionCreated(evebody.CreateSubscription{
		//	UserID:    subscription.UserID.String(),
		//	PlanID:    subscription.PlanID.String(),
		//	TypeID:    typeID.String(),
		//	CreatedAt: subscription.CreatedAt,
		//}); err != nil {
		//	return err
		//}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (d *domain) DeactivateSubscription(ctx context.Context, UserID uuid.UUID) error {
	err := d.Infra.Data.SQL.Subscriptions.Transaction(func(ctx context.Context) error {
		if err := d.Infra.Data.SQL.Subscriptions.New().Filter(map[string]any{
			"user_id": UserID,
		}).Update(ctx, map[string]any{
			"status": models.SubscriptionStatusInactive,
		}); err != nil {
			return err
		}

		//TODO to supplement the work with Kafka

		return nil
	})

	return err
}

func (d *domain) CanceledSubscription(ctx context.Context, UserID uuid.UUID) error {
	err := d.Infra.Data.SQL.Subscriptions.Transaction(func(ctx context.Context) error {
		if err := d.Infra.Data.SQL.Subscriptions.New().Filter(map[string]any{
			"user_id": UserID,
		}).Update(ctx, map[string]any{
			"status": models.SubscriptionStatusCanceled,
		}); err != nil {
			return err
		}

		//TODO to supplement the work with Kafka

		return nil
	})

	return err
}

func (d *domain) GetSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error) {
	res, err := d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"user_id": userID.String()}).Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *domain) AddPaymentMethod(ctx context.Context, userID uuid.UUID, token string, payType string) (*models.PaymentMethod, error) {
	var method *models.PaymentMethod
	pType, err := models.ParsePayType(payType)
	if err != nil {
		return nil, err
	}

	errTrn := d.Infra.Data.SQL.PaymentMethods.New().Transaction(func(ctx context.Context) error {
		method = &models.PaymentMethod{
			ID:            uuid.New(),
			UserID:        userID,
			Type:          pType,
			ProviderToken: token,
			IsDefault:     false,
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
		}

		err := d.Infra.Data.SQL.PaymentMethods.New().Insert(ctx, method)
		if err != nil {
			return err
		}

		return nil
	})

	if errTrn != nil {
		return nil, errTrn
	}

	return method, nil
}

func (d *domain) DeletePaymentMethod(ctx context.Context, userID, paymentMethodID uuid.UUID) error {
	err := d.Infra.Data.SQL.PaymentMethods.New().Filter(map[string]any{
		"id":      paymentMethodID.String(),
		"user_id": userID.String(),
	}).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (d *domain) GetPaymentMethod(ctx context.Context, userID, paymentMethodID uuid.UUID) (*models.PaymentMethod, error) {
	res, err := d.Infra.Data.SQL.PaymentMethods.Filter(map[string]any{
		"id":      paymentMethodID.String(),
		"user_id": userID.String(),
	}).Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *domain) GetUserPaymentMethods(ctx context.Context, userID uuid.UUID) ([]models.PaymentMethod, error) {
	res, err := d.Infra.Data.SQL.PaymentMethods.Filter(map[string]any{
		"user_id": userID.String(),
	}).Select(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *domain) SetPaymentMethodAsDefault(ctx context.Context, userID, paymentMethodID uuid.UUID) error {
	err := d.Infra.Data.SQL.PaymentMethods.Transaction(func(ctx context.Context) error {
		err := d.Infra.Data.SQL.PaymentMethods.New().Filter(map[string]any{
			"user_id": userID.String(),
		}).Update(ctx, map[string]any{
			"is_default": false,
		})
		if err != nil {
			return err
		}

		err = d.Infra.Data.SQL.PaymentMethods.New().Filter(map[string]any{
			"user_id": userID.String(),
			"id":      paymentMethodID.String(),
		}).Update(ctx, map[string]any{
			"is_default": true,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
