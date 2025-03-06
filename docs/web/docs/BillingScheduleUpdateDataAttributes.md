# BillingScheduleUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**SchedulesDate** | **time.Time** | Billing Schedule Date | 
**AttemptedDate** | Pointer to **time.Time** | Billing Schedule Attempted Date | [optional] 
**Status** | **string** | Billing Schedule Status | 
**UpdatedAt** | **time.Time** | Billing Schedule Updated At | 
**CreatedAt** | **time.Time** | Billing Schedule Created At | 

## Methods

### NewBillingScheduleUpdateDataAttributes

`func NewBillingScheduleUpdateDataAttributes(userId string, schedulesDate time.Time, status string, updatedAt time.Time, createdAt time.Time, ) *BillingScheduleUpdateDataAttributes`

NewBillingScheduleUpdateDataAttributes instantiates a new BillingScheduleUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingScheduleUpdateDataAttributesWithDefaults

`func NewBillingScheduleUpdateDataAttributesWithDefaults() *BillingScheduleUpdateDataAttributes`

NewBillingScheduleUpdateDataAttributesWithDefaults instantiates a new BillingScheduleUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BillingScheduleUpdateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingScheduleUpdateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingScheduleUpdateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSchedulesDate

`func (o *BillingScheduleUpdateDataAttributes) GetSchedulesDate() time.Time`

GetSchedulesDate returns the SchedulesDate field if non-nil, zero value otherwise.

### GetSchedulesDateOk

`func (o *BillingScheduleUpdateDataAttributes) GetSchedulesDateOk() (*time.Time, bool)`

GetSchedulesDateOk returns a tuple with the SchedulesDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulesDate

`func (o *BillingScheduleUpdateDataAttributes) SetSchedulesDate(v time.Time)`

SetSchedulesDate sets SchedulesDate field to given value.


### GetAttemptedDate

`func (o *BillingScheduleUpdateDataAttributes) GetAttemptedDate() time.Time`

GetAttemptedDate returns the AttemptedDate field if non-nil, zero value otherwise.

### GetAttemptedDateOk

`func (o *BillingScheduleUpdateDataAttributes) GetAttemptedDateOk() (*time.Time, bool)`

GetAttemptedDateOk returns a tuple with the AttemptedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedDate

`func (o *BillingScheduleUpdateDataAttributes) SetAttemptedDate(v time.Time)`

SetAttemptedDate sets AttemptedDate field to given value.

### HasAttemptedDate

`func (o *BillingScheduleUpdateDataAttributes) HasAttemptedDate() bool`

HasAttemptedDate returns a boolean if a field has been set.

### GetStatus

`func (o *BillingScheduleUpdateDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingScheduleUpdateDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingScheduleUpdateDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *BillingScheduleUpdateDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingScheduleUpdateDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingScheduleUpdateDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *BillingScheduleUpdateDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingScheduleUpdateDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingScheduleUpdateDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


