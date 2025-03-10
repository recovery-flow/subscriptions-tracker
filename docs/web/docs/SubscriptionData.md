# SubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID | 
**Type** | **string** |  | 
**Attributes** | [**SubscriptionDataAttributes**](SubscriptionDataAttributes.md) |  | 
**Relationships** | [**SubscriptionDataRelationships**](SubscriptionDataRelationships.md) |  | 

## Methods

### NewSubscriptionData

`func NewSubscriptionData(id string, type_ string, attributes SubscriptionDataAttributes, relationships SubscriptionDataRelationships, ) *SubscriptionData`

NewSubscriptionData instantiates a new SubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataWithDefaults

`func NewSubscriptionDataWithDefaults() *SubscriptionData`

NewSubscriptionDataWithDefaults instantiates a new SubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SubscriptionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionData) GetAttributes() SubscriptionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionData) GetAttributesOk() (*SubscriptionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionData) SetAttributes(v SubscriptionDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionData) GetRelationships() SubscriptionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionData) GetRelationshipsOk() (*SubscriptionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionData) SetRelationships(v SubscriptionDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


