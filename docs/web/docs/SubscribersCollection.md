# SubscribersCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriberData**](SubscriberData.md) |  | 
**Links** | [**LinksPagination**](LinksPagination.md) |  | 

## Methods

### NewSubscribersCollection

`func NewSubscribersCollection(data []SubscriberData, links LinksPagination, ) *SubscribersCollection`

NewSubscribersCollection instantiates a new SubscribersCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribersCollectionWithDefaults

`func NewSubscribersCollectionWithDefaults() *SubscribersCollection`

NewSubscribersCollectionWithDefaults instantiates a new SubscribersCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscribersCollection) GetData() []SubscriberData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscribersCollection) GetDataOk() (*[]SubscriberData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscribersCollection) SetData(v []SubscriberData)`

SetData sets Data field to given value.


### GetLinks

`func (o *SubscribersCollection) GetLinks() LinksPagination`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SubscribersCollection) GetLinksOk() (*LinksPagination, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SubscribersCollection) SetLinks(v LinksPagination)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


