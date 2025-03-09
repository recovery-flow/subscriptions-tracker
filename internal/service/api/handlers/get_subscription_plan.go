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

func GetSubscriptionPlan(w http.ResponseWriter, r *http.Request) {
	planID, err := uuid.Parse(chi.URLParam(r, "plan_id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse plan id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse plan id"))...)
		return
	}

	plan, err := Domain(r).GetPlan(r.Context(), planID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	sType, _, err := Domain(r).GetType(r.Context(), plan.TypeID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	httpkit.Render(w, responses.SubscriptionPlan(plan, sType))
}
