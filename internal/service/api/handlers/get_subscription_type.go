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

func GetSubscriptionType(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "type_id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse plan id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse plan id"))...)
		return
	}

	sType, plans, err := Domain(r).GetType(r.Context(), ID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription plan")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	httpkit.Render(w, responses.SubscriptionType(sType, plans))

}
