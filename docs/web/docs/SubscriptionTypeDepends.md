# SubscriptionTypeDepends

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriptionTypeData**](SubscriptionTypeData.md) |  | 
**Included** | [**[]SubscriptionPlanData**](SubscriptionPlanData.md) |  | 

## Methods

### NewSubscriptionTypeDepends

`func NewSubscriptionTypeDepends(data []SubscriptionTypeData, included []SubscriptionPlanData, ) *SubscriptionTypeDepends`

NewSubscriptionTypeDepends instantiates a new SubscriptionTypeDepends object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTypeDependsWithDefaults

`func NewSubscriptionTypeDependsWithDefaults() *SubscriptionTypeDepends`

NewSubscriptionTypeDependsWithDefaults instantiates a new SubscriptionTypeDepends object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionTypeDepends) GetData() []SubscriptionTypeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionTypeDepends) GetDataOk() (*[]SubscriptionTypeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionTypeDepends) SetData(v []SubscriptionTypeData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *SubscriptionTypeDepends) GetIncluded() []SubscriptionPlanData`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SubscriptionTypeDepends) GetIncludedOk() (*[]SubscriptionPlanData, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SubscriptionTypeDepends) SetIncluded(v []SubscriptionPlanData)`

SetIncluded sets Included field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


