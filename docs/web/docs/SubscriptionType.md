# SubscriptionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SubscriptionTypeData**](SubscriptionTypeData.md) |  | 
**Included** | [**[]SubscriptionPlanData**](SubscriptionPlanData.md) |  | 

## Methods

### NewSubscriptionType

`func NewSubscriptionType(data SubscriptionTypeData, included []SubscriptionPlanData, ) *SubscriptionType`

NewSubscriptionType instantiates a new SubscriptionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTypeWithDefaults

`func NewSubscriptionTypeWithDefaults() *SubscriptionType`

NewSubscriptionTypeWithDefaults instantiates a new SubscriptionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionType) GetData() SubscriptionTypeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionType) GetDataOk() (*SubscriptionTypeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionType) SetData(v SubscriptionTypeData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *SubscriptionType) GetIncluded() []SubscriptionPlanData`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SubscriptionType) GetIncludedOk() (*[]SubscriptionPlanData, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SubscriptionType) SetIncluded(v []SubscriptionPlanData)`

SetIncluded sets Included field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


