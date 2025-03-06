# SubPlanData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Subscription Plan ID | 
**Type** | **string** |  | 
**Attributes** | [**SubPlanDataAttributes**](SubPlanDataAttributes.md) |  | 

## Methods

### NewSubPlanData

`func NewSubPlanData(id string, type_ string, attributes SubPlanDataAttributes, ) *SubPlanData`

NewSubPlanData instantiates a new SubPlanData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubPlanDataWithDefaults

`func NewSubPlanDataWithDefaults() *SubPlanData`

NewSubPlanDataWithDefaults instantiates a new SubPlanData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubPlanData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubPlanData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubPlanData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SubPlanData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubPlanData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubPlanData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubPlanData) GetAttributes() SubPlanDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubPlanData) GetAttributesOk() (*SubPlanDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubPlanData) SetAttributes(v SubPlanDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


