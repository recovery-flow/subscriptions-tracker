# SubscriptionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SubscriptionPlanData**](SubscriptionPlanData.md) |  | 
**Included** | [**[]SubscriptionTypeData**](SubscriptionTypeData.md) |  | 

## Methods

### NewSubscriptionPlan

`func NewSubscriptionPlan(data SubscriptionPlanData, included []SubscriptionTypeData, ) *SubscriptionPlan`

NewSubscriptionPlan instantiates a new SubscriptionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanWithDefaults

`func NewSubscriptionPlanWithDefaults() *SubscriptionPlan`

NewSubscriptionPlanWithDefaults instantiates a new SubscriptionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionPlan) GetData() SubscriptionPlanData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionPlan) GetDataOk() (*SubscriptionPlanData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionPlan) SetData(v SubscriptionPlanData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *SubscriptionPlan) GetIncluded() []SubscriptionTypeData`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SubscriptionPlan) GetIncludedOk() (*[]SubscriptionTypeData, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SubscriptionPlan) SetIncluded(v []SubscriptionTypeData)`

SetIncluded sets Included field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


