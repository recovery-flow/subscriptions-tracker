# TransactionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User ID | [optional] 
**PlanId** | Pointer to **string** | Plan ID | [optional] 
**SubId** | Pointer to **string** | Subscriber ID | [optional] 
**Amount** | **float32** | Transaction amount | 
**Currency** | **string** | Transaction currency | 
**PaymentMethod** | **string** | Payment method | 
**ProviderTransactionId** | **string** | Provider transaction ID | 
**CreatedAt** | **time.Time** | Transaction creation date | 

## Methods

### NewTransactionDataAttributes

`func NewTransactionDataAttributes(amount float32, currency string, paymentMethod string, providerTransactionId string, createdAt time.Time, ) *TransactionDataAttributes`

NewTransactionDataAttributes instantiates a new TransactionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDataAttributesWithDefaults

`func NewTransactionDataAttributesWithDefaults() *TransactionDataAttributes`

NewTransactionDataAttributesWithDefaults instantiates a new TransactionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TransactionDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TransactionDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TransactionDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TransactionDataAttributes) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPlanId

`func (o *TransactionDataAttributes) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *TransactionDataAttributes) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *TransactionDataAttributes) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *TransactionDataAttributes) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSubId

`func (o *TransactionDataAttributes) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *TransactionDataAttributes) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *TransactionDataAttributes) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *TransactionDataAttributes) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionDataAttributes) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionDataAttributes) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionDataAttributes) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransactionDataAttributes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionDataAttributes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionDataAttributes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *TransactionDataAttributes) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionDataAttributes) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionDataAttributes) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetProviderTransactionId

`func (o *TransactionDataAttributes) GetProviderTransactionId() string`

GetProviderTransactionId returns the ProviderTransactionId field if non-nil, zero value otherwise.

### GetProviderTransactionIdOk

`func (o *TransactionDataAttributes) GetProviderTransactionIdOk() (*string, bool)`

GetProviderTransactionIdOk returns a tuple with the ProviderTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTransactionId

`func (o *TransactionDataAttributes) SetProviderTransactionId(v string)`

SetProviderTransactionId sets ProviderTransactionId field to given value.


### GetCreatedAt

`func (o *TransactionDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


