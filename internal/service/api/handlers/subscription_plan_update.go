package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/requests"
	"github.com/recovery-flow/tokens"
)

func SubscriptionPlanUpdate(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	req, err := requests.SubscriptionPlanUpdate(r)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse request")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse request"))...)
		return
	}

	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse plan id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse plan id"))...)
		return
	}

	err = Domain(r).UpdateSubPlan(r.Context(), ID, map[string]any{
		"name":     req.Data.Attributes.Name,
		"desc":     req.Data.Attributes.Desc,
		"price":    req.Data.Attributes.Price,
		"currency": req.Data.Attributes.Currency,
	})
	if err != nil {
		Log(r).WithError(err).Debug("Failed to update subscription plan")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription plan %s updated, by user %s", ID, accountID)

	httpkit.Render(w, http.StatusAccepted)
}
