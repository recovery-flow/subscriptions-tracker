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
	"github.com/recovery-flow/tokens/identity"
)

func UserSubscriptionActivate(w http.ResponseWriter, r *http.Request) {
	accountID, sessionID, subTypeID, role, server, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	if *role != identity.User {
		Log(r).WithField("role", role).Error("User role is not allowed")
		httpkit.RenderErr(w, problems.Forbidden("Forbidden"))
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

	methodID, err := uuid.Parse(req.Data.Attributes.PaymentMethodId)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse payment method id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse payment method id"))...)
		return
	}

	planID, err := uuid.Parse(req.Data.Attributes.PlanId)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse plan id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse plan id"))...)
		return
	}

	res, err := Domain(r).CreateSubscription(r.Context(), *accountID, planID, methodID, 0)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to create subscription")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	sType, err := Domain(r).GetTypeByPlan(r.Context(), planID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get subscription type")
		httpkit.RenderErr(w, problems.InternalError(err.Error()))
		return
	}

	typeID := sType.ID

	Log(r).Infof("Subscription create, by user %s at session %s", accountID, sessionID)

	httpkit.Render(w, responses.Subscription(*res, &typeID))
}
