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

func SubscriptionTypeUpdate(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	req, err := requests.SubscriptionTypeUpdate(r)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse request")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse request"))...)
		return
	}

	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse type id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse type id"))...)
	}

	err = Domain(r).UpdateSubType(r.Context(), ID, map[string]any{
		"name": req.Data.Attributes.Name,
	})
	if err != nil {
		Log(r).WithError(err).Debug("Failed to update subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription type %s updated, by user %s", ID, accountID)
	httpkit.Render(w, http.StatusAccepted)
}
