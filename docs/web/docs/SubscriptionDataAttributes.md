# SubscriptionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | Plan ID | 
**PaymentMethodId** | **string** | Payment Method ID | 
**Status** | **string** | State | 
**Availability** | **string** | Availability | 
**StartAt** | Pointer to **time.Time** | Start at | [optional] 
**UpdatedAt** | **time.Time** | Updated at | 
**CreatedAt** | **time.Time** | Created at | 

## Methods

### NewSubscriptionDataAttributes

`func NewSubscriptionDataAttributes(planId string, paymentMethodId string, status string, availability string, updatedAt time.Time, createdAt time.Time, ) *SubscriptionDataAttributes`

NewSubscriptionDataAttributes instantiates a new SubscriptionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataAttributesWithDefaults

`func NewSubscriptionDataAttributesWithDefaults() *SubscriptionDataAttributes`

NewSubscriptionDataAttributesWithDefaults instantiates a new SubscriptionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *SubscriptionDataAttributes) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionDataAttributes) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionDataAttributes) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPaymentMethodId

`func (o *SubscriptionDataAttributes) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *SubscriptionDataAttributes) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *SubscriptionDataAttributes) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetStatus

`func (o *SubscriptionDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAvailability

`func (o *SubscriptionDataAttributes) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *SubscriptionDataAttributes) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *SubscriptionDataAttributes) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetStartAt

`func (o *SubscriptionDataAttributes) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *SubscriptionDataAttributes) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *SubscriptionDataAttributes) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *SubscriptionDataAttributes) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *SubscriptionDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


