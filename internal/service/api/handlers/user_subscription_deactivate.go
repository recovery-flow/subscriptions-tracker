package handlers

import (
	"net/http"

	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
)

func UserSubscriptionDeactivate(w http.ResponseWriter, r *http.Request) {
	accountID, _, subTypeID, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	err = Domain(r).DeactivateSubscription(r.Context(), *accountID, *subTypeID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to deactivate user subscription")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
