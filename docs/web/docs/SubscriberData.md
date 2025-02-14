# SubscriberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Subscriber ID | 
**Type** | **string** |  | 
**Attributes** | [**SubscriberDataAttributes**](SubscriberDataAttributes.md) |  | 

## Methods

### NewSubscriberData

`func NewSubscriberData(id string, type_ string, attributes SubscriberDataAttributes, ) *SubscriberData`

NewSubscriberData instantiates a new SubscriberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberDataWithDefaults

`func NewSubscriberDataWithDefaults() *SubscriberData`

NewSubscriberDataWithDefaults instantiates a new SubscriberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriberData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriberData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriberData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SubscriberData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriberData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriberData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriberData) GetAttributes() SubscriberDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriberData) GetAttributesOk() (*SubscriberDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriberData) SetAttributes(v SubscriberDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


