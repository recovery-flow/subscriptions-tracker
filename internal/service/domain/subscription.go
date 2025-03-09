package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type SubscriptionsTracker interface {
	GetUserSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error)
	CreateSubscription(ctx context.Context, UserID, PlanID, PaymentMethodID uuid.UUID, frePeriod time.Duration) (*models.Subscription, error)
	DeactivateSubscription(ctx context.Context, UserID uuid.UUID) error
	CanceledSubscription(ctx context.Context, UserID uuid.UUID) error

	CreatePaymentMethod(ctx context.Context, userID uuid.UUID, token string, payType string) (*models.PaymentMethod, error)
	DeletePaymentMethod(ctx context.Context, userID, paymentMethodID uuid.UUID) error
	GetPaymentMethod(ctx context.Context, paymentMethodID uuid.UUID) (*models.PaymentMethod, error)
	GetUserPaymentMethod(ctx context.Context, userID, paymentMethodID uuid.UUID) (*models.PaymentMethod, error)
	GetUserPaymentMethods(ctx context.Context, userID uuid.UUID) ([]models.PaymentMethod, error)
	GetUserDefaultPaymentMethod(ctx context.Context, userID uuid.UUID) (*models.PaymentMethod, error)
	SetPaymentMethodAsDefault(ctx context.Context, userID, paymentMethodID uuid.UUID) error
}

func (d *domain) CreateSubscription(ctx context.Context, UserID, PlanID, PaymentMethodID uuid.UUID, frePeriod time.Duration) (*models.Subscription, error) {
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

		//Create a new subscription

		var endDate time.Time
		uInterval := sPlan.BillingCycle
		interval := sPlan.BillingInterval

		switch uInterval {
		case models.CycleOnce:
			endDate = time.Now().AddDate(100, 0, 0)
		case models.CycleDay:
			endDate = time.Now().AddDate(0, 0, 1*int(interval))
		case models.CycleWeek:
			endDate = time.Now().AddDate(0, 0, 7*int(interval))
		case models.CycleMonth:
			endDate = time.Now().AddDate(0, int(interval), 0)
		case models.CycleYear:
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

		bc := &models.BillingSchedule{
			UserID:        UserID,
			SchedulesDate: time.Now().UTC().Add(frePeriod),
			Status:        models.ScheduleBillingStatusPlanned,
		}
		err = d.Infra.Data.SQL.Schedules.New().Insert(ctx, bc)
		if err != nil {
			return err
		}

		//Create Billing Schedules for the Subscription
		if frePeriod == 0 {
			err = d.Infra.Data.SQL.Schedules.New().Insert(ctx, bc)
			if err != nil {
				return err
			}
			//TODO there should have been a payment :)
			//pay, payErr := code.SomePaymentMethod(pMethod)
			//if payErr != nil {
			//	err = d.Infra.Data.SQL.Transactions.New().Insert(ctx , &models.Transaction{
			//		ID: uuid.New(),
			//		UserID: UserID,
			//		PaymentMethodID: PaymentMethodID,
			//		Amount: sPlan.Price,
			//		Currency: sPlan.Currency,
			//		Status: models.TrnStatusFailed,
			//		PaymentProvider: "TODO",
			//		PaymentProviderID: "TODO",
			//		TransactionDate: time.Now().UTC(),
			//	})
			//	if err != nil {
			//		d.log.WithError(err).Error("failed to insert transaction")
			//	}
			//	return payErr
			//}
			//err = d.Infra.Data.SQL.Transactions.New().Insert(ctx , &models.Transaction{
			//	ID: uuid.New(),
			//	UserID: UserID,
			//	PaymentMethodID: PaymentMethodID,
			//	Amount: sPlan.Price,
			//	Currency: sPlan.Currency,
			//	Status: models.TrnStatusSuccess,
			//	PaymentProvider: "TODO",
			//	PaymentProviderID: "TODO",
			//	TransactionDate: time.Now().UTC(),
			//})
			//if err != nil {
			//	d.log.WithError(err).Error("failed to insert transaction")
			//}
		}

		//TODO to supplement the work with Producer
		//if err := d.Infra.Producer.SubscriptionCreated(evebody.CreateSubscription{
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

		//TODO to supplement the work with Producer

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

		//TODO to supplement the work with Producer

		return nil
	})

	return err
}

func (d *domain) GetUserSubscription(ctx context.Context, userID uuid.UUID) (*models.Subscription, error) {
	res, err := d.Infra.Data.SQL.Subscriptions.Filter(map[string]any{"user_id": userID.String()}).Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *domain) CreatePaymentMethod(ctx context.Context, userID uuid.UUID, token string, payType string) (*models.PaymentMethod, error) {
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

func (d *domain) GetPaymentMethod(ctx context.Context, paymentMethodID uuid.UUID) (*models.PaymentMethod, error) {
	res, err := d.Infra.Data.SQL.PaymentMethods.Filter(map[string]any{
		"id": paymentMethodID.String(),
	}).Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *domain) GetUserPaymentMethod(ctx context.Context, userID, paymentMethodID uuid.UUID) (*models.PaymentMethod, error) {
	method, err := d.GetPaymentMethod(ctx, paymentMethodID)
	if err != nil {
		return nil, err
	}

	if method.UserID != userID {
		return nil, fmt.Errorf("payment method does not belong to user %s", paymentMethodID)
	}

	return method, nil
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

func (d *domain) GetUserDefaultPaymentMethod(ctx context.Context, userID uuid.UUID) (*models.PaymentMethod, error) {
	res, err := d.Infra.Data.SQL.PaymentMethods.Filter(map[string]any{
		"user_id":    userID.String(),
		"is_default": true,
	}).Get(ctx)
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
