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

func SubscriptionTypeDeactivate(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse type id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse type id"))...)
		return
	}

	err = Domain(r).DeactivateType(r.Context(), ID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to update subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription type %s activate, by user %s", ID, accountID)

	httpkit.Render(w, http.StatusAccepted)
}
