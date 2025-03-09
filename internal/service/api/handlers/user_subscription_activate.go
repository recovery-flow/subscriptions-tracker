package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/requests"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/responses"
	"github.com/recovery-flow/tokens"
)

func UserSubscriptionCreate(w http.ResponseWriter, r *http.Request) {
	accountID, sessionID, subTypeID, _, server, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	req, err := requests.SubscriptionCreate(r)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse request")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse request"))...)
		return
	}

	if server != nil {
		Log(r).WithError(fmt.Errorf("server try to make subscription")).Errorf("Server cant make subscription")
		httpkit.RenderErr(w, problems.Forbidden("Forbidden"))
		return
	}

	if subTypeID != nil {
		Log(r).WithError(fmt.Errorf("user already have subscription")).Errorf("User already have subscription")
		httpkit.RenderErr(w, problems.Forbidden("Forbidden"))
		return
	}

	method, err := uuid.Parse(req.Data.Attributes.PaymentMethodId)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse payment method id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse payment method id"))...)
		return
	}

	plan, err := uuid.Parse(req.Data.Attributes.PlanId)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse plan id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse plan id"))...)
		return
	}

	res, err := Domain(r).CreateSubscription(r.Context(), *accountID, plan, method, 0)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to create subscription")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	Log(r).Infof("Subscription create, by user %s at session %s", accountID, sessionID)

	httpkit.Render(w, responses.Subscription(*res))
}
