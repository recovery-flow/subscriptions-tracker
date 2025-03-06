# SubscriptionPlanDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **string** | Subscription Plan ID | 
**Name** | **string** | Subscription Plan Name | 
**Desc** | **string** | Subscription Plan Description | 
**Price** | **float32** | Subscription Plan Price | 
**Currency** | **string** | Subscription Plan Currency | 
**BillingCycle** | **string** | Subscription Plan Billing Interval | 
**BillingInterval** | **int32** | Subscription Plan Billing Interval | 
**Status** | **string** | Subscription Plan Status | 
**UpdatedAt** | **time.Time** | Subscription Plan Updated At | 
**CreatedAt** | **time.Time** | Subscription Plan Created At | 

## Methods

### NewSubscriptionPlanDataAttributes

`func NewSubscriptionPlanDataAttributes(typeId string, name string, desc string, price float32, currency string, billingCycle string, billingInterval int32, status string, updatedAt time.Time, createdAt time.Time, ) *SubscriptionPlanDataAttributes`

NewSubscriptionPlanDataAttributes instantiates a new SubscriptionPlanDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanDataAttributesWithDefaults

`func NewSubscriptionPlanDataAttributesWithDefaults() *SubscriptionPlanDataAttributes`

NewSubscriptionPlanDataAttributesWithDefaults instantiates a new SubscriptionPlanDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *SubscriptionPlanDataAttributes) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *SubscriptionPlanDataAttributes) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *SubscriptionPlanDataAttributes) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.


### GetName

`func (o *SubscriptionPlanDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionPlanDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionPlanDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *SubscriptionPlanDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SubscriptionPlanDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SubscriptionPlanDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetPrice

`func (o *SubscriptionPlanDataAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionPlanDataAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionPlanDataAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *SubscriptionPlanDataAttributes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionPlanDataAttributes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionPlanDataAttributes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBillingCycle

`func (o *SubscriptionPlanDataAttributes) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *SubscriptionPlanDataAttributes) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *SubscriptionPlanDataAttributes) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetBillingInterval

`func (o *SubscriptionPlanDataAttributes) GetBillingInterval() int32`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *SubscriptionPlanDataAttributes) GetBillingIntervalOk() (*int32, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *SubscriptionPlanDataAttributes) SetBillingInterval(v int32)`

SetBillingInterval sets BillingInterval field to given value.


### GetStatus

`func (o *SubscriptionPlanDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionPlanDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionPlanDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *SubscriptionPlanDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionPlanDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionPlanDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *SubscriptionPlanDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionPlanDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionPlanDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


