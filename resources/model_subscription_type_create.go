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

// checks if the SubscriptionTypeCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionTypeCreate{}

// SubscriptionTypeCreate struct for SubscriptionTypeCreate
type SubscriptionTypeCreate struct {
	Data SubscriptionTypeCreateData `json:"data"`
}

type _SubscriptionTypeCreate SubscriptionTypeCreate

// NewSubscriptionTypeCreate instantiates a new SubscriptionTypeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionTypeCreate(data SubscriptionTypeCreateData) *SubscriptionTypeCreate {
	this := SubscriptionTypeCreate{}
	this.Data = data
	return &this
}

// NewSubscriptionTypeCreateWithDefaults instantiates a new SubscriptionTypeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionTypeCreateWithDefaults() *SubscriptionTypeCreate {
	this := SubscriptionTypeCreate{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionTypeCreate) GetData() SubscriptionTypeCreateData {
	if o == nil {
		var ret SubscriptionTypeCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionTypeCreate) GetDataOk() (*SubscriptionTypeCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionTypeCreate) SetData(v SubscriptionTypeCreateData) {
	o.Data = v
}

func (o SubscriptionTypeCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionTypeCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionTypeCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varSubscriptionTypeCreate := _SubscriptionTypeCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionTypeCreate)

	if err != nil {
		return err
	}

	*o = SubscriptionTypeCreate(varSubscriptionTypeCreate)

	return err
}

type NullableSubscriptionTypeCreate struct {
	value *SubscriptionTypeCreate
	isSet bool
}

func (v NullableSubscriptionTypeCreate) Get() *SubscriptionTypeCreate {
	return v.value
}

func (v *NullableSubscriptionTypeCreate) Set(val *SubscriptionTypeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionTypeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionTypeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionTypeCreate(val *SubscriptionTypeCreate) *NullableSubscriptionTypeCreate {
	return &NullableSubscriptionTypeCreate{value: val, isSet: true}
}

func (v NullableSubscriptionTypeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionTypeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


