package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
)

func SubscriptionPlanActivate(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse plan id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse plan id"))...)
		return
	}

	err = Domain(r).ActivatePlan(r.Context(), ID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to update subscription plan")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription plan %s activate, by user %s", ID, accountID)

	httpkit.Render(w, http.StatusAccepted)
}
