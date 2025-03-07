# PaymentMethodCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**Type** | **string** |  | 
**ProviderToken** | **string** | Provider token | 
**IsDefault** | **bool** | Is default | 

## Methods

### NewPaymentMethodCreateDataAttributes

`func NewPaymentMethodCreateDataAttributes(userId string, type_ string, providerToken string, isDefault bool, ) *PaymentMethodCreateDataAttributes`

NewPaymentMethodCreateDataAttributes instantiates a new PaymentMethodCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreateDataAttributesWithDefaults

`func NewPaymentMethodCreateDataAttributesWithDefaults() *PaymentMethodCreateDataAttributes`

NewPaymentMethodCreateDataAttributesWithDefaults instantiates a new PaymentMethodCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PaymentMethodCreateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentMethodCreateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentMethodCreateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetType

`func (o *PaymentMethodCreateDataAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCreateDataAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCreateDataAttributes) SetType(v string)`

SetType sets Type field to given value.


### GetProviderToken

`func (o *PaymentMethodCreateDataAttributes) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *PaymentMethodCreateDataAttributes) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *PaymentMethodCreateDataAttributes) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.


### GetIsDefault

`func (o *PaymentMethodCreateDataAttributes) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethodCreateDataAttributes) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethodCreateDataAttributes) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


