package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type SubscriptionSchedule interface {
	GetUserSchedule(ctx context.Context, userID uuid.UUID) (*models.BillingSchedule, error)

	SelectSchedules(ctx context.Context, after bool, date time.Time, status string) ([]models.BillingSchedule, error)

	MadeTransaction(ctx context.Context, userID uuid.UUID) (*models.Transaction, error)
	SubscriptionPaymentPayPal(ctx context.Context, sub *models.Subscription, plan *models.SubscriptionPlan, method *models.PaymentMethod) (*models.Transaction, error)
}

func (d *domain) GetUserSchedule(ctx context.Context, userID uuid.UUID) (*models.BillingSchedule, error) {
	return d.Infra.Data.SQL.Schedules.New().Filter(map[string]interface{}{
		"user_id": userID.String(),
	}).Get(ctx)
}

func (d *domain) SelectSchedules(ctx context.Context, after bool, date time.Time, status string) ([]models.BillingSchedule, error) {
	scheduleStatus, err := models.ParseBillingStatus(status)
	if err != nil {
		return nil, err
	}

	return d.Infra.Data.SQL.Schedules.New().Filter(map[string]any{
		"status": scheduleStatus,
	}).FilterTime("schedules_date", after, date).Select(ctx)
}

func (d *domain) MadeTransaction(ctx context.Context, userID uuid.UUID) (*models.Transaction, error) {
	sub, err := d.GetUserSubscription(ctx, userID)
	if err != nil {
		return nil, err
	}

	if sub.Availability != models.SubPlanAvailable {
		return nil, fmt.Errorf("subscription is not available")
	}

	plan, err := d.GetPlan(ctx, sub.PlanID)
	if err != nil {
		return nil, err
	}

	if plan.Status != models.StatusPlanActive {
		return nil, fmt.Errorf("plan is not active")
	}

	method, err := d.GetUserDefaultPaymentMethod(ctx, userID)
	if err != nil {
		return nil, err
	}

	return d.SubscriptionPaymentPayPal(ctx, sub, plan, method)
}

func (d *domain) SubscriptionPaymentPayPal(ctx context.Context, sub *models.Subscription, plan *models.SubscriptionPlan, method *models.PaymentMethod) (*models.Transaction, error) {
	trn := &models.Transaction{
		ID:              uuid.New(),
		UserID:          sub.UserID,
		PaymentMethodID: method.ID,
		Amount:          plan.Price,
		Currency:        plan.Currency,
		Status:          models.TrnStatusSuccess,
		PaymentProvider: models.PaymentProviderPaypal,
		TransactionDate: time.Now().UTC(),
	}

	err := d.Infra.Data.SQL.Schedules.New().Filter(
		map[string]interface{}{
			"user_id": sub.UserID.String(),
		}).Update(ctx,
		map[string]interface{}{
			"status": models.ScheduleBillingStatusProcessing,
		})
	if err != nil {
		return nil, err
	}

	res := d.Infra.Data.SQL.Schedules.New().Transaction(func(ctx context.Context) error {
		resTrn := true //TODO: implement payment provider

		if resTrn != true {
			trn.Status = models.TrnStatusFailed
			err = d.Infra.Data.SQL.Transactions.New().Insert(ctx, trn)
			if err != nil {
				return err
			}

			err = d.Infra.Data.SQL.Schedules.New().Filter(map[string]any{
				"user_id": sub.UserID.String(),
			}).Update(ctx, map[string]any{
				"status":         models.ScheduleBillingStatusFailed,
				"attempted_date": time.Now().UTC(),
			})
			if err != nil {
				return err
			}

			return nil
		}

		err = d.Infra.Data.SQL.Transactions.New().Insert(ctx, trn)
		if err != nil {
			return err
		}

		err = d.Infra.Data.SQL.Schedules.New().Filter(map[string]any{
			"user_id": sub.UserID.String(),
		}).Update(ctx, map[string]any{
			"status":         models.StatusTypeActive,
			"schedules_date": time.Now().UTC().AddDate(0, 1, 0),
		})
		if err != nil {
			return err
		}

		return nil
	})
	if res != nil {
		return nil, res
	}

	return trn, nil
}
