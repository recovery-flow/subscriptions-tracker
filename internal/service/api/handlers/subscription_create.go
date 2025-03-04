package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

func SubscriptionCreate(w http.ResponseWriter, r *http.Request) {
	//accountID, _, sub, _, _, err := tokens.GetAccountData(r.Context())
	//if err != nil {
	//	Log(r).WithError(err).Debug("Failed to get account data")
	//	httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
	//	return
	//}

	//if sub != nil {
	//	Log(r).Error("Account already has a subscription")
	//	httpkit.RenderErr(w, problems.Forbidden("Account already has a subscription"))
	//}

	//TypeID, err := uuid.Parse(chi.URLParam(r, "type_id"))
	//if err != nil {
	//	Log(r).Errorf("Failed to parse sub_id: %v", err)
	//	httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
	//		"sub_id": validation.NewError("sub_id", "Invalid sub_id"),
	//	})...)
	//	return
	//}

	Log(r).Debug("Creating subscription")

	typID := uuid.New()

	typ := models.SubscriptionType{
		ID:          typID,
		Name:        "test",
		Description: "test",
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := Domain(r).CreateSubType(r.Context(), typ)
	if err != nil {
		Log(r).WithError(err).Error("Failed to create subscription type")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	pId := uuid.New()

	plan := models.SubscriptionPlan{
		ID:                  pId,
		TypeID:              typID,
		Price:               100,
		Name:                "test",
		Description:         "test",
		BillingInterval:     1,
		BillingIntervalUnit: "month",
		Currency:            "USD",
		Status:              "active",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	err = Domain(r).CreateSubPlan(r.Context(), plan)
	if err != nil {
		Log(r).WithError(err).Error("Failed to create subscription plan")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	res := models.Subscription{
		UserID:          uuid.New(),
		PlanID:          pId,
		PaymentMethodID: uuid.New(),
		Status:          "active",
		Availability:    "available",
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 1, 0),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	subscription, err := Domain(r).CreateSubscription(r.Context(), res)
	if err != nil {
		Log(r).WithError(err).Error("Failed to create subscription")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, subscription)
}
