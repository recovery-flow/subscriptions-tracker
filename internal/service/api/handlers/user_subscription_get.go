package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/responses"
)

func UserSubscriptionGet(w http.ResponseWriter, r *http.Request) {
	subID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse user id"))...)
		return
	}

	res, err := Domain(r).GetUserSubscription(r.Context(), subID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get user subscription")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	sType, err := Domain(r).GetTypeByPlan(r.Context(), res.PlanID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	typeID := sType.ID

	httpkit.Render(w, responses.Subscription(*res, &typeID))
}
