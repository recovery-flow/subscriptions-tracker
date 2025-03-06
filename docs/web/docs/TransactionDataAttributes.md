# TransactionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**PaymentId** | **string** | Payment ID | 
**Amount** | **float32** | Transaction amount | 
**Currency** | **string** | Transaction currency | 
**Status** | **string** | Transaction status | 
**PaymentProvider** | **string** | Payment provider | 
**ProviderTransactionId** | **string** | Provider transaction ID | 
**TransactionDate** | **time.Time** | Transaction creation date | 

## Methods

### NewTransactionDataAttributes

`func NewTransactionDataAttributes(userId string, paymentId string, amount float32, currency string, status string, paymentProvider string, providerTransactionId string, transactionDate time.Time, ) *TransactionDataAttributes`

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


### GetPaymentId

`func (o *TransactionDataAttributes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *TransactionDataAttributes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *TransactionDataAttributes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


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


### GetStatus

`func (o *TransactionDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentProvider

`func (o *TransactionDataAttributes) GetPaymentProvider() string`

GetPaymentProvider returns the PaymentProvider field if non-nil, zero value otherwise.

### GetPaymentProviderOk

`func (o *TransactionDataAttributes) GetPaymentProviderOk() (*string, bool)`

GetPaymentProviderOk returns a tuple with the PaymentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProvider

`func (o *TransactionDataAttributes) SetPaymentProvider(v string)`

SetPaymentProvider sets PaymentProvider field to given value.


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


### GetTransactionDate

`func (o *TransactionDataAttributes) GetTransactionDate() time.Time`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *TransactionDataAttributes) GetTransactionDateOk() (*time.Time, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *TransactionDataAttributes) SetTransactionDate(v time.Time)`

SetTransactionDate sets TransactionDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


