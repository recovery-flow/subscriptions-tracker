# TransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Transaction ID | 
**Type** | **string** |  | 
**Attributes** | [**TransactionDataAttributes**](TransactionDataAttributes.md) |  | 

## Methods

### NewTransactionData

`func NewTransactionData(id string, type_ string, attributes TransactionDataAttributes, ) *TransactionData`

NewTransactionData instantiates a new TransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDataWithDefaults

`func NewTransactionDataWithDefaults() *TransactionData`

NewTransactionDataWithDefaults instantiates a new TransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *TransactionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TransactionData) GetAttributes() TransactionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TransactionData) GetAttributesOk() (*TransactionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TransactionData) SetAttributes(v TransactionDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


