# PaymentMethodDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**Type** | **string** |  | 
**ProviderToken** | **string** | Provider token | 
**IsDefault** | **bool** | Is default | 
**CreatedAt** | **time.Time** | Created at | 

## Methods

### NewPaymentMethodDataAttributes

`func NewPaymentMethodDataAttributes(userId string, type_ string, providerToken string, isDefault bool, createdAt time.Time, ) *PaymentMethodDataAttributes`

NewPaymentMethodDataAttributes instantiates a new PaymentMethodDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodDataAttributesWithDefaults

`func NewPaymentMethodDataAttributesWithDefaults() *PaymentMethodDataAttributes`

NewPaymentMethodDataAttributesWithDefaults instantiates a new PaymentMethodDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PaymentMethodDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentMethodDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentMethodDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetType

`func (o *PaymentMethodDataAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodDataAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodDataAttributes) SetType(v string)`

SetType sets Type field to given value.


### GetProviderToken

`func (o *PaymentMethodDataAttributes) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *PaymentMethodDataAttributes) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *PaymentMethodDataAttributes) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.


### GetIsDefault

`func (o *PaymentMethodDataAttributes) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethodDataAttributes) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethodDataAttributes) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCreatedAt

`func (o *PaymentMethodDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


