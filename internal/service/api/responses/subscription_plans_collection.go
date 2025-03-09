package responses

//func SubscriptionPlansCollection(subscriptionPlans []models.SubscriptionPlan, baseURL string, queryParams url.Values, totalItems, pageSize, pageNumber int64) resources.SubscriptionPlansCollection {
//	var data []resources.SubscriptionPlanData
//	for _, subscriptionPlan := range subscriptionPlans {
//		data = append(data, SubscriptionPlan(subscriptionPlan).Data)
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
//	return resources.SubscriptionPlansCollection{
//		Data: resources.SubscriptionPlansCollectionData{
//			Type: resources.TypeSubscriptionPlansCollection,
//			Attributes: resources.SubscriptionPlansCollectionDataAttributes{
//				SubscriptionPlans: data,
//			},
//		},
//		Links: links,
//	}
//}
