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

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	req, err := requests.PaymentMethodCreate(r)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse request")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse request"))...)
		return
	}

	def := req.Data.Attributes.IsDefault
	token := req.Data.Attributes.ProviderToken
	mType := req.Data.Attributes.Type

	payment, err := Domain(r).CreatePaymentMethod(r.Context(), *accountID, token, mType)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to create payment method")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if def {
		err = Domain(r).SetPaymentMethodAsDefault(r.Context(), *accountID, payment.ID)
		if err != nil {
			Log(r).WithError(err).Debug("Failed to set payment method as default")
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	httpkit.Render(w, responses.PaymentMethod(*payment))
}
