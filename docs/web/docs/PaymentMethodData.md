# PaymentMethodData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | payment method ID | 
**Type** | **string** |  | 
**Attributes** | [**PaymentMethodDataAttributes**](PaymentMethodDataAttributes.md) |  | 

## Methods

### NewPaymentMethodData

`func NewPaymentMethodData(id string, type_ string, attributes PaymentMethodDataAttributes, ) *PaymentMethodData`

NewPaymentMethodData instantiates a new PaymentMethodData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodDataWithDefaults

`func NewPaymentMethodDataWithDefaults() *PaymentMethodData`

NewPaymentMethodDataWithDefaults instantiates a new PaymentMethodData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethodData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PaymentMethodData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PaymentMethodData) GetAttributes() PaymentMethodDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentMethodData) GetAttributesOk() (*PaymentMethodDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentMethodData) SetAttributes(v PaymentMethodDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


