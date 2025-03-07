/*
User storage service

User storage service for recovery flow

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SubscriptionPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlan{}

// SubscriptionPlan struct for SubscriptionPlan
type SubscriptionPlan struct {
	Data SubscriptionPlanData `json:"data"`
	Included []SubscriptionTypeData `json:"included"`
}

type _SubscriptionPlan SubscriptionPlan

// NewSubscriptionPlan instantiates a new SubscriptionPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlan(data SubscriptionPlanData, included []SubscriptionTypeData) *SubscriptionPlan {
	this := SubscriptionPlan{}
	this.Data = data
	this.Included = included
	return &this
}

// NewSubscriptionPlanWithDefaults instantiates a new SubscriptionPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlanWithDefaults() *SubscriptionPlan {
	this := SubscriptionPlan{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPlan) GetData() SubscriptionPlanData {
	if o == nil {
		var ret SubscriptionPlanData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlan) GetDataOk() (*SubscriptionPlanData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionPlan) SetData(v SubscriptionPlanData) {
	o.Data = v
}

// GetIncluded returns the Included field value
func (o *SubscriptionPlan) GetIncluded() []SubscriptionTypeData {
	if o == nil {
		var ret []SubscriptionTypeData
		return ret
	}

	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlan) GetIncludedOk() ([]SubscriptionTypeData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Included, true
}

// SetIncluded sets field value
func (o *SubscriptionPlan) SetIncluded(v []SubscriptionTypeData) {
	o.Included = v
}

func (o SubscriptionPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["included"] = o.Included
	return toSerialize, nil
}

func (o *SubscriptionPlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"included",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubscriptionPlan := _SubscriptionPlan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPlan)

	if err != nil {
		return err
	}

	*o = SubscriptionPlan(varSubscriptionPlan)

	return err
}

type NullableSubscriptionPlan struct {
	value *SubscriptionPlan
	isSet bool
}

func (v NullableSubscriptionPlan) Get() *SubscriptionPlan {
	return v.value
}

func (v *NullableSubscriptionPlan) Set(val *SubscriptionPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlan(val *SubscriptionPlan) *NullableSubscriptionPlan {
	return &NullableSubscriptionPlan{value: val, isSet: true}
}

func (v NullableSubscriptionPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


