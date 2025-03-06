# BillingScheduleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | user ID | 
**Type** | **string** |  | 
**Attributes** | [**BillingScheduleDataAttributes**](BillingScheduleDataAttributes.md) |  | 

## Methods

### NewBillingScheduleData

`func NewBillingScheduleData(id string, type_ string, attributes BillingScheduleDataAttributes, ) *BillingScheduleData`

NewBillingScheduleData instantiates a new BillingScheduleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingScheduleDataWithDefaults

`func NewBillingScheduleDataWithDefaults() *BillingScheduleData`

NewBillingScheduleDataWithDefaults instantiates a new BillingScheduleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingScheduleData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingScheduleData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingScheduleData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BillingScheduleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingScheduleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingScheduleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BillingScheduleData) GetAttributes() BillingScheduleDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingScheduleData) GetAttributesOk() (*BillingScheduleDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingScheduleData) SetAttributes(v BillingScheduleDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


