package responses

//func SubscriptionTypesCollection(subscriptionTypes []models.SubscriptionType, baseURL string, queryParams url.Values, totalItems, pageSize, pageNumber int64) resources.SubscriptionTypesCollection {
//	var data []resources.SubscriptionTypeData
//	for _, subscriptionType := range subscriptionTypes {
//		data = append(data, SubscriptionType(subscriptionType).Data)
//	}
//
//	links := resources.LinksPagination{
//		Self:     *generatePaginationLink(baseURL, queryParams, pageNumber, pageSize),
//		Previous: generatePaginationLink(baseURL, queryParams, pageNumber-1, pageSize),
//		Next:     generatePaginationLink(baseURL, queryParams, pageNumber+1, pageSize),
//	}
//
//	if pageNumber <= 1 {
//		links.Previous = nil
//	}
//
//	totalPages := (totalItems + pageSize - 1) / pageSize
//	if pageNumber >= totalPages {
//		links.Next = nil
//	}
//
//	return resources.SubscriptionTypesCollection{
//		Data: resources.SubscriptionTypesCollectionData{
//			Type: resources.TypeSubscriptionsTypesCollection,
//			Attributes: resources.SubscriptionTypesCollectionDataAttributes{
//				SubscriptionTypes: data,
//			},
//		},
//		Links: links,
//	}
//}
