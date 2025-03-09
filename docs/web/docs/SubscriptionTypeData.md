# SubscriptionTypeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Subscription type ID | 
**Type** | **string** |  | 
**Attributes** | [**SubscriptionTypeDataAttributes**](SubscriptionTypeDataAttributes.md) |  | 
**Relationships** | [**SubscriptionTypeDataRelationships**](SubscriptionTypeDataRelationships.md) |  | 

## Methods

### NewSubscriptionTypeData

`func NewSubscriptionTypeData(id string, type_ string, attributes SubscriptionTypeDataAttributes, relationships SubscriptionTypeDataRelationships, ) *SubscriptionTypeData`

NewSubscriptionTypeData instantiates a new SubscriptionTypeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTypeDataWithDefaults

`func NewSubscriptionTypeDataWithDefaults() *SubscriptionTypeData`

NewSubscriptionTypeDataWithDefaults instantiates a new SubscriptionTypeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionTypeData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionTypeData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionTypeData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SubscriptionTypeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionTypeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionTypeData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionTypeData) GetAttributes() SubscriptionTypeDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionTypeData) GetAttributesOk() (*SubscriptionTypeDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionTypeData) SetAttributes(v SubscriptionTypeDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionTypeData) GetRelationships() SubscriptionTypeDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionTypeData) GetRelationshipsOk() (*SubscriptionTypeDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionTypeData) SetRelationships(v SubscriptionTypeDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


