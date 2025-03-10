package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/recovery-flow/comtools/jsonkit"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func SubscriptionPlanCreate(r *http.Request) (req resources.SubscriptionPlanCreate, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = jsonkit.NewDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In(resources.TypeSubscriptionPlanCreate)),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	return req, errs.Filter()
}
