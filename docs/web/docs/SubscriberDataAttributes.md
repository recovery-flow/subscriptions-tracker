# SubscriberDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID UUID | 
**PlanId** | **string** | Plan ID object ID | 
**StreakMonths** | **int32** | Streak months | 
**Status** | **string** | Status | 
**StartAt** | **time.Time** | Start at | 
**EndAt** | **time.Time** | End at | 
**UpdatedAt** | Pointer to **time.Time** | Updated at | [optional] 
**CreatedAt** | **time.Time** | Created at | 

## Methods

### NewSubscriberDataAttributes

`func NewSubscriberDataAttributes(userId string, planId string, streakMonths int32, status string, startAt time.Time, endAt time.Time, createdAt time.Time, ) *SubscriberDataAttributes`

NewSubscriberDataAttributes instantiates a new SubscriberDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberDataAttributesWithDefaults

`func NewSubscriberDataAttributesWithDefaults() *SubscriberDataAttributes`

NewSubscriberDataAttributesWithDefaults instantiates a new SubscriberDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SubscriberDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubscriberDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubscriberDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlanId

`func (o *SubscriberDataAttributes) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriberDataAttributes) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriberDataAttributes) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetStreakMonths

`func (o *SubscriberDataAttributes) GetStreakMonths() int32`

GetStreakMonths returns the StreakMonths field if non-nil, zero value otherwise.

### GetStreakMonthsOk

`func (o *SubscriberDataAttributes) GetStreakMonthsOk() (*int32, bool)`

GetStreakMonthsOk returns a tuple with the StreakMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreakMonths

`func (o *SubscriberDataAttributes) SetStreakMonths(v int32)`

SetStreakMonths sets StreakMonths field to given value.


### GetStatus

`func (o *SubscriberDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriberDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriberDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartAt

`func (o *SubscriberDataAttributes) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *SubscriberDataAttributes) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *SubscriberDataAttributes) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *SubscriberDataAttributes) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *SubscriberDataAttributes) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *SubscriberDataAttributes) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.


### GetUpdatedAt

`func (o *SubscriberDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriberDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriberDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriberDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriberDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriberDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriberDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


