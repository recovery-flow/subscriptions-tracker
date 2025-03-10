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

// checks if the SubscriptionTypesCollectionDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionTypesCollectionDataAttributes{}

// SubscriptionTypesCollectionDataAttributes struct for SubscriptionTypesCollectionDataAttributes
type SubscriptionTypesCollectionDataAttributes struct {
	SubscriptionTypes []SubscriptionTypeData `json:"subscription_types"`
}

type _SubscriptionTypesCollectionDataAttributes SubscriptionTypesCollectionDataAttributes

// NewSubscriptionTypesCollectionDataAttributes instantiates a new SubscriptionTypesCollectionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionTypesCollectionDataAttributes(subscriptionTypes []SubscriptionTypeData) *SubscriptionTypesCollectionDataAttributes {
	this := SubscriptionTypesCollectionDataAttributes{}
	this.SubscriptionTypes = subscriptionTypes
	return &this
}

// NewSubscriptionTypesCollectionDataAttributesWithDefaults instantiates a new SubscriptionTypesCollectionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionTypesCollectionDataAttributesWithDefaults() *SubscriptionTypesCollectionDataAttributes {
	this := SubscriptionTypesCollectionDataAttributes{}
	return &this
}

// GetSubscriptionTypes returns the SubscriptionTypes field value
func (o *SubscriptionTypesCollectionDataAttributes) GetSubscriptionTypes() []SubscriptionTypeData {
	if o == nil {
		var ret []SubscriptionTypeData
		return ret
	}

	return o.SubscriptionTypes
}

// GetSubscriptionTypesOk returns a tuple with the SubscriptionTypes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionTypesCollectionDataAttributes) GetSubscriptionTypesOk() ([]SubscriptionTypeData, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionTypes, true
}

// SetSubscriptionTypes sets field value
func (o *SubscriptionTypesCollectionDataAttributes) SetSubscriptionTypes(v []SubscriptionTypeData) {
	o.SubscriptionTypes = v
}

func (o SubscriptionTypesCollectionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionTypesCollectionDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription_types"] = o.SubscriptionTypes
	return toSerialize, nil
}

func (o *SubscriptionTypesCollectionDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscription_types",
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

	varSubscriptionTypesCollectionDataAttributes := _SubscriptionTypesCollectionDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionTypesCollectionDataAttributes)

	if err != nil {
		return err
	}

	*o = SubscriptionTypesCollectionDataAttributes(varSubscriptionTypesCollectionDataAttributes)

	return err
}

type NullableSubscriptionTypesCollectionDataAttributes struct {
	value *SubscriptionTypesCollectionDataAttributes
	isSet bool
}

func (v NullableSubscriptionTypesCollectionDataAttributes) Get() *SubscriptionTypesCollectionDataAttributes {
	return v.value
}

func (v *NullableSubscriptionTypesCollectionDataAttributes) Set(val *SubscriptionTypesCollectionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionTypesCollectionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionTypesCollectionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionTypesCollectionDataAttributes(val *SubscriptionTypesCollectionDataAttributes) *NullableSubscriptionTypesCollectionDataAttributes {
	return &NullableSubscriptionTypesCollectionDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionTypesCollectionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionTypesCollectionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


