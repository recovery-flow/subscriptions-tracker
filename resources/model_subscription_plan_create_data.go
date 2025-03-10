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

// checks if the SubscriptionPlanCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlanCreateData{}

// SubscriptionPlanCreateData struct for SubscriptionPlanCreateData
type SubscriptionPlanCreateData struct {
	Type string `json:"type"`
	Attributes SubscriptionPlanCreateDataAttributes `json:"attributes"`
}

type _SubscriptionPlanCreateData SubscriptionPlanCreateData

// NewSubscriptionPlanCreateData instantiates a new SubscriptionPlanCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlanCreateData(type_ string, attributes SubscriptionPlanCreateDataAttributes) *SubscriptionPlanCreateData {
	this := SubscriptionPlanCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSubscriptionPlanCreateDataWithDefaults instantiates a new SubscriptionPlanCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlanCreateDataWithDefaults() *SubscriptionPlanCreateData {
	this := SubscriptionPlanCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPlanCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPlanCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionPlanCreateData) GetAttributes() SubscriptionPlanCreateDataAttributes {
	if o == nil {
		var ret SubscriptionPlanCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanCreateData) GetAttributesOk() (*SubscriptionPlanCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionPlanCreateData) SetAttributes(v SubscriptionPlanCreateDataAttributes) {
	o.Attributes = v
}

func (o SubscriptionPlanCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlanCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *SubscriptionPlanCreateData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varSubscriptionPlanCreateData := _SubscriptionPlanCreateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPlanCreateData)

	if err != nil {
		return err
	}

	*o = SubscriptionPlanCreateData(varSubscriptionPlanCreateData)

	return err
}

type NullableSubscriptionPlanCreateData struct {
	value *SubscriptionPlanCreateData
	isSet bool
}

func (v NullableSubscriptionPlanCreateData) Get() *SubscriptionPlanCreateData {
	return v.value
}

func (v *NullableSubscriptionPlanCreateData) Set(val *SubscriptionPlanCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlanCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlanCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlanCreateData(val *SubscriptionPlanCreateData) *NullableSubscriptionPlanCreateData {
	return &NullableSubscriptionPlanCreateData{value: val, isSet: true}
}

func (v NullableSubscriptionPlanCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlanCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


