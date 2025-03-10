package responses

import (
	"net/url"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/resources"
)

func SubscriptionCollection(subscriptions []models.Subscription, baseURL string, queryParams url.Values, totalItems, pageSize, pageNumber int64) resources.SubscriptionsCollection {
	var data []resources.SubscriptionData
	for _, subscription := range subscriptions {
		data = append(data, Subscription(subscription, nil).Data)
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

	return resources.SubscriptionsCollection{
		Data: resources.SubscriptionsCollectionData{
			Type: resources.TypeSubscriptionsCollection,
			Attributes: resources.SubscriptionsCollectionDataAttributes{
				Subscriptions: data,
			},
		},
		Links: links,
	}
}
