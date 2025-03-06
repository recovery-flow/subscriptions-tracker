package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

type SubscriptionSchedule interface {
	GetUserSchedule(ctx context.Context, userID uuid.UUID) (*models.BillingSchedule, error)

	SelectSchedule(ctx context.Context, after bool, date time.Time) ([]models.BillingSchedule, error)
	MadeTransaction(ctx context.Context, userID uuid.UUID) error
}

func (d *domain) GetUserSchedule(ctx context.Context, userID uuid.UUID) (*models.BillingSchedule, error) {
	return d.Infra.Data.SQL.Schedules.New().Filter(map[string]interface{}{
		"user_id": userID.String(),
	}).Get(ctx)
}

func (d *domain) SelectSchedule(ctx context.Context, after bool, date time.Time) ([]models.BillingSchedule, error) {
	return d.Infra.Data.SQL.Schedules.New().FilterTime("scheduled_date", after, date).Select(ctx)
}

func (d *domain) MadeTransaction(ctx context.Context, userID uuid.UUID) error {
	sub, err := d.GetUserSubscription(ctx, userID)
	if err != nil {
		return err
	}

	plan, err := d.GetSubPlan(ctx, sub.PlanID)
	if err != nil {
		return err
	}

	method, err := d.GetUserDefaultPaymentMethod(ctx, userID)
	if err != nil {
		return err
	}

	trn := &models.Transaction{
		ID:              uuid.New(),
		UserID:          userID,
		PaymentMethodID: method.ID,
		Amount:          plan.Price,
		Currency:        plan.Currency,
		PaymentProvider: models.PaymentProviderStripe,
		TransactionDate: time.Now().UTC(),
	}

	err = d.Infra.Data.SQL.Schedules.New().Filter(
		map[string]interface{}{
			"user_id": userID.String(),
		}).Update(ctx,
		map[string]interface{}{
			"status": models.BillingStatusProcessing,
		})
	if err != nil {
		return err
	}

	res := d.Infra.Data.SQL.Schedules.New().Transaction(func(ctx context.Context) error {
		//TODO make transaction
		resTrn := true
		statusTrn := models.TrnStatusFailed
		scheduleStatus := models.BillingStatusFailed

		if resTrn == true {
			statusTrn = models.TrnStatusSuccess
			scheduleStatus = models.BillingStatusPlanned
		}

		trn.Status = statusTrn
		err = d.Infra.Data.SQL.Transactions.New().Insert(ctx, trn)
		if err != nil {
			return err
		}

		err = d.Infra.Data.SQL.Schedules.New().Filter(map[string]any{
			"user_id": userID.String(),
		}).Update(ctx, map[string]any{
			"status": scheduleStatus,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if res != nil {
		return res
	}
	return nil
}
