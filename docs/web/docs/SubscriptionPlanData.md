# SubscriptionPlanData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Subscription Plan ID | 
**Type** | **string** |  | 
**Attributes** | [**SubscriptionPlanDataAttributes**](SubscriptionPlanDataAttributes.md) |  | 
**Relationships** | [**SubscriptionPlanDataRelationships**](SubscriptionPlanDataRelationships.md) |  | 

## Methods

### NewSubscriptionPlanData

`func NewSubscriptionPlanData(id string, type_ string, attributes SubscriptionPlanDataAttributes, relationships SubscriptionPlanDataRelationships, ) *SubscriptionPlanData`

NewSubscriptionPlanData instantiates a new SubscriptionPlanData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanDataWithDefaults

`func NewSubscriptionPlanDataWithDefaults() *SubscriptionPlanData`

NewSubscriptionPlanDataWithDefaults instantiates a new SubscriptionPlanData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPlanData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPlanData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPlanData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SubscriptionPlanData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionPlanData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionPlanData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionPlanData) GetAttributes() SubscriptionPlanDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionPlanData) GetAttributesOk() (*SubscriptionPlanDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionPlanData) SetAttributes(v SubscriptionPlanDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionPlanData) GetRelationships() SubscriptionPlanDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionPlanData) GetRelationshipsOk() (*SubscriptionPlanDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionPlanData) SetRelationships(v SubscriptionPlanDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


