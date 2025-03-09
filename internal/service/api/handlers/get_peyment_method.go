package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/responses"
	"github.com/recovery-flow/tokens"
)

func GetPaymentMethod(w http.ResponseWriter, r *http.Request) {
	accountID, _, _, _, _, err := tokens.GetAccountData(r.Context())
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get account data")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	ID, err := uuid.Parse(chi.URLParam(r, "payment_id"))
	if err != nil {
		Log(r).WithError(err).Debug("Failed to parse payment ID")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse payment ID"))...)
		return
	}

	payment, err := Domain(r).GetUserPaymentMethod(r.Context(), *accountID, ID)
	if err != nil {
		Log(r).WithError(err).Debug("Failed to get payment method")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.PaymentMethod(*payment))
}
