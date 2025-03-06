package responses

import (
	"net/url"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func PaymentMethodsCollection(methods []models.PaymentMethod, baseURL string, queryParams url.Values, totalItems, pageSize, pageNumber int64) resources.PaymentMethodsCollection {
	var data []resources.PaymentMethodData
	for _, method := range methods {
		data = append(data, PaymentMethod(method).Data)
	}

	links := resources.LinksPagination{
		Self:     *generatePaginationLink(baseURL, queryParams, pageNumber, pageSize),
		Previous: generatePaginationLink(baseURL, queryParams, pageNumber-1, pageSize),
		Next:     generatePaginationLink(baseURL, queryParams, pageNumber+1, pageSize),
	}

	if pageNumber <= 1 {
		links.Previous = nil
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	if pageNumber >= totalPages {
		links.Next = nil
	}

	return resources.PaymentMethodsCollection{
		Data: resources.PaymentMethodsCollectionData{
			Type: resources.TypePaymentMethodsCollection,
			Attributes: resources.PaymentMethodsCollectionDataAttributes{
				PaymentMethods: data,
			},
		},
		Links: links,
	}
}
