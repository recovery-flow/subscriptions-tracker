package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/requests"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/responses"
	"github.com/recovery-flow/tokens"
)

func SubscriptionPlanCreate(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	req, err := requests.SubscriptionPlanCreate(r)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse request")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse request"))...)
		return
	}

	typeID, err := uuid.Parse(req.Data.Attributes.TypeId)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse type id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse type id"))...)
		return
	}

	name := req.Data.Attributes.Name
	desc := req.Data.Attributes.Desc
	price := req.Data.Attributes.Price
	currency := req.Data.Attributes.Currency
	billingCycle := req.Data.Attributes.BillingCycle
	billingInterval := req.Data.Attributes.BillingInterval

	plans, err := Domain(r).CreatePlan(r.Context(), name, desc, typeID, price, currency, int8(billingInterval), billingCycle)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to create subscription plan")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	sType, _, err := Domain(r).GetType(r.Context(), typeID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription plan %s created, by user %s", plans.ID, accountID)
	httpkit.Render(w, responses.SubscriptionPlan(plans, sType))
}
