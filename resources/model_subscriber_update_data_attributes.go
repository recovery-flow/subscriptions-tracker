/*
User storage service

User storage service for recovery flow

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
)

// checks if the SubscriberUpdateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberUpdateDataAttributes{}

// SubscriberUpdateDataAttributes struct for SubscriberUpdateDataAttributes
type SubscriberUpdateDataAttributes struct {
	// Plan ID object ID
	PlanId *string `json:"plan_id,omitempty"`
	// Streak months
	StreakMonths *int32 `json:"streak_months,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	// Start at
	StartAt *time.Time `json:"start_at,omitempty"`
	// End at
	EndAt *time.Time `json:"end_at,omitempty"`
}

// NewSubscriberUpdateDataAttributes instantiates a new SubscriberUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberUpdateDataAttributes() *SubscriberUpdateDataAttributes {
	this := SubscriberUpdateDataAttributes{}
	return &this
}

// NewSubscriberUpdateDataAttributesWithDefaults instantiates a new SubscriberUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberUpdateDataAttributesWithDefaults() *SubscriberUpdateDataAttributes {
	this := SubscriberUpdateDataAttributes{}
	return &this
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriberUpdateDataAttributes) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberUpdateDataAttributes) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriberUpdateDataAttributes) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriberUpdateDataAttributes) SetPlanId(v string) {
	o.PlanId = &v
}

// GetStreakMonths returns the StreakMonths field value if set, zero value otherwise.
func (o *SubscriberUpdateDataAttributes) GetStreakMonths() int32 {
	if o == nil || IsNil(o.StreakMonths) {
		var ret int32
		return ret
	}
	return *o.StreakMonths
}

// GetStreakMonthsOk returns a tuple with the StreakMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberUpdateDataAttributes) GetStreakMonthsOk() (*int32, bool) {
	if o == nil || IsNil(o.StreakMonths) {
		return nil, false
	}
	return o.StreakMonths, true
}

// HasStreakMonths returns a boolean if a field has been set.
func (o *SubscriberUpdateDataAttributes) HasStreakMonths() bool {
	if o != nil && !IsNil(o.StreakMonths) {
		return true
	}

	return false
}

// SetStreakMonths gets a reference to the given int32 and assigns it to the StreakMonths field.
func (o *SubscriberUpdateDataAttributes) SetStreakMonths(v int32) {
	o.StreakMonths = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriberUpdateDataAttributes) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberUpdateDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriberUpdateDataAttributes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubscriberUpdateDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *SubscriberUpdateDataAttributes) GetStartAt() time.Time {
	if o == nil || IsNil(o.StartAt) {
		var ret time.Time
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberUpdateDataAttributes) GetStartAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartAt) {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *SubscriberUpdateDataAttributes) HasStartAt() bool {
	if o != nil && !IsNil(o.StartAt) {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given time.Time and assigns it to the StartAt field.
func (o *SubscriberUpdateDataAttributes) SetStartAt(v time.Time) {
	o.StartAt = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *SubscriberUpdateDataAttributes) GetEndAt() time.Time {
	if o == nil || IsNil(o.EndAt) {
		var ret time.Time
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberUpdateDataAttributes) GetEndAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndAt) {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *SubscriberUpdateDataAttributes) HasEndAt() bool {
	if o != nil && !IsNil(o.EndAt) {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given time.Time and assigns it to the EndAt field.
func (o *SubscriberUpdateDataAttributes) SetEndAt(v time.Time) {
	o.EndAt = &v
}

func (o SubscriberUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberUpdateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.StreakMonths) {
		toSerialize["streak_months"] = o.StreakMonths
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StartAt) {
		toSerialize["start_at"] = o.StartAt
	}
	if !IsNil(o.EndAt) {
		toSerialize["end_at"] = o.EndAt
	}
	return toSerialize, nil
}

type NullableSubscriberUpdateDataAttributes struct {
	value *SubscriberUpdateDataAttributes
	isSet bool
}

func (v NullableSubscriberUpdateDataAttributes) Get() *SubscriberUpdateDataAttributes {
	return v.value
}

func (v *NullableSubscriberUpdateDataAttributes) Set(val *SubscriberUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberUpdateDataAttributes(val *SubscriberUpdateDataAttributes) *NullableSubscriberUpdateDataAttributes {
	return &NullableSubscriberUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriberUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


