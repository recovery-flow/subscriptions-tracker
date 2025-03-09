package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
)

func UserSubscriptionDeactivate(w http.ResponseWriter, r *http.Request) {
	subID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse user id"))...)
		return
	}

	err = Domain(r).DeactivateSubscription(r.Context(), subID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to deactivate user subscription")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
