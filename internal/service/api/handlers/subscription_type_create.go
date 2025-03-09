package handlers

import (
	"fmt"
	"net/http"

	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/requests"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/responses"
	"github.com/recovery-flow/tokens"
)

func SubscriptionTypeCreate(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	req, err := requests.SubscriptionTypeCreate(r)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse request")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse request"))...)
		return
	}

	name := req.Data.Attributes.Name
	desc := req.Data.Attributes.Description

	sType, err := Domain(r).CreateType(r.Context(), name, desc)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to create subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription type %s created, by user %s", sType.ID, accountID)
	httpkit.Render(w, responses.SubscriptionType(sType, nil))
}
