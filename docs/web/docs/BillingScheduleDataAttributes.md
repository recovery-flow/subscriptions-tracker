# BillingScheduleDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchedulesDate** | **time.Time** | Billing Schedule Date | 
**AttemptedDate** | Pointer to **time.Time** | Billing Schedule Attempted Date | [optional] 
**Status** | **string** | Billing Schedule Status | 
**UpdatedAt** | **time.Time** | Billing Schedule Updated At | 
**CreatedAt** | **time.Time** | Billing Schedule Created At | 

## Methods

### NewBillingScheduleDataAttributes

`func NewBillingScheduleDataAttributes(schedulesDate time.Time, status string, updatedAt time.Time, createdAt time.Time, ) *BillingScheduleDataAttributes`

NewBillingScheduleDataAttributes instantiates a new BillingScheduleDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingScheduleDataAttributesWithDefaults

`func NewBillingScheduleDataAttributesWithDefaults() *BillingScheduleDataAttributes`

NewBillingScheduleDataAttributesWithDefaults instantiates a new BillingScheduleDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedulesDate

`func (o *BillingScheduleDataAttributes) GetSchedulesDate() time.Time`

GetSchedulesDate returns the SchedulesDate field if non-nil, zero value otherwise.

### GetSchedulesDateOk

`func (o *BillingScheduleDataAttributes) GetSchedulesDateOk() (*time.Time, bool)`

GetSchedulesDateOk returns a tuple with the SchedulesDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulesDate

`func (o *BillingScheduleDataAttributes) SetSchedulesDate(v time.Time)`

SetSchedulesDate sets SchedulesDate field to given value.


### GetAttemptedDate

`func (o *BillingScheduleDataAttributes) GetAttemptedDate() time.Time`

GetAttemptedDate returns the AttemptedDate field if non-nil, zero value otherwise.

### GetAttemptedDateOk

`func (o *BillingScheduleDataAttributes) GetAttemptedDateOk() (*time.Time, bool)`

GetAttemptedDateOk returns a tuple with the AttemptedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedDate

`func (o *BillingScheduleDataAttributes) SetAttemptedDate(v time.Time)`

SetAttemptedDate sets AttemptedDate field to given value.

### HasAttemptedDate

`func (o *BillingScheduleDataAttributes) HasAttemptedDate() bool`

HasAttemptedDate returns a boolean if a field has been set.

### GetStatus

`func (o *BillingScheduleDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingScheduleDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingScheduleDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *BillingScheduleDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingScheduleDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingScheduleDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *BillingScheduleDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingScheduleDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingScheduleDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


