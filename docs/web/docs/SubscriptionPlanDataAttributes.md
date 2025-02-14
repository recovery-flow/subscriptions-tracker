# SubscriptionPlanDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title | 
**Price** | **float32** | Price | 
**Currency** | **string** | Currency | 
**PayFrequency** | **string** | Pay frequency | 
**Status** | **string** | Status | 
**UpdatedAt** | Pointer to **time.Time** | Updated at | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** | Created at | 

## Methods

### NewSubscriptionPlanDataAttributes

`func NewSubscriptionPlanDataAttributes(title string, price float32, currency string, payFrequency string, status string, createdAt time.Time, ) *SubscriptionPlanDataAttributes`

NewSubscriptionPlanDataAttributes instantiates a new SubscriptionPlanDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanDataAttributesWithDefaults

`func NewSubscriptionPlanDataAttributesWithDefaults() *SubscriptionPlanDataAttributes`

NewSubscriptionPlanDataAttributesWithDefaults instantiates a new SubscriptionPlanDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SubscriptionPlanDataAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SubscriptionPlanDataAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SubscriptionPlanDataAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


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


### GetPayFrequency

`func (o *SubscriptionPlanDataAttributes) GetPayFrequency() string`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *SubscriptionPlanDataAttributes) GetPayFrequencyOk() (*string, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *SubscriptionPlanDataAttributes) SetPayFrequency(v string)`

SetPayFrequency sets PayFrequency field to given value.


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

### HasUpdatedAt

`func (o *SubscriptionPlanDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SubscriptionPlanDataAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubscriptionPlanDataAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubscriptionPlanDataAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubscriptionPlanDataAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

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


