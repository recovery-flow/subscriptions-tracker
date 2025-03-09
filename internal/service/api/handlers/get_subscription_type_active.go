package handlers

import (
	"net/http"

	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/responses"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
)

func GetSubscriptionTypeActive(w http.ResponseWriter, r *http.Request) {
	status := models.StatusTypeActive

	res, err := Domain(r).GetAllType(r.Context(), &status)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription types")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	httpkit.Render(w, responses.SubscriptionTypeDepends(res))
}
