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

// checks if the BillingSchedulesCollectionDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingSchedulesCollectionDataAttributes{}

// BillingSchedulesCollectionDataAttributes struct for BillingSchedulesCollectionDataAttributes
type BillingSchedulesCollectionDataAttributes struct {
	BillingSchedule []BillingScheduleData `json:"billing_schedule"`
}

type _BillingSchedulesCollectionDataAttributes BillingSchedulesCollectionDataAttributes

// NewBillingSchedulesCollectionDataAttributes instantiates a new BillingSchedulesCollectionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingSchedulesCollectionDataAttributes(billingSchedule []BillingScheduleData) *BillingSchedulesCollectionDataAttributes {
	this := BillingSchedulesCollectionDataAttributes{}
	this.BillingSchedule = billingSchedule
	return &this
}

// NewBillingSchedulesCollectionDataAttributesWithDefaults instantiates a new BillingSchedulesCollectionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingSchedulesCollectionDataAttributesWithDefaults() *BillingSchedulesCollectionDataAttributes {
	this := BillingSchedulesCollectionDataAttributes{}
	return &this
}

// GetBillingSchedule returns the BillingSchedule field value
func (o *BillingSchedulesCollectionDataAttributes) GetBillingSchedule() []BillingScheduleData {
	if o == nil {
		var ret []BillingScheduleData
		return ret
	}

	return o.BillingSchedule
}

// GetBillingScheduleOk returns a tuple with the BillingSchedule field value
// and a boolean to check if the value has been set.
func (o *BillingSchedulesCollectionDataAttributes) GetBillingScheduleOk() ([]BillingScheduleData, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingSchedule, true
}

// SetBillingSchedule sets field value
func (o *BillingSchedulesCollectionDataAttributes) SetBillingSchedule(v []BillingScheduleData) {
	o.BillingSchedule = v
}

func (o BillingSchedulesCollectionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingSchedulesCollectionDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billing_schedule"] = o.BillingSchedule
	return toSerialize, nil
}

func (o *BillingSchedulesCollectionDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billing_schedule",
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

	varBillingSchedulesCollectionDataAttributes := _BillingSchedulesCollectionDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingSchedulesCollectionDataAttributes)

	if err != nil {
		return err
	}

	*o = BillingSchedulesCollectionDataAttributes(varBillingSchedulesCollectionDataAttributes)

	return err
}

type NullableBillingSchedulesCollectionDataAttributes struct {
	value *BillingSchedulesCollectionDataAttributes
	isSet bool
}

func (v NullableBillingSchedulesCollectionDataAttributes) Get() *BillingSchedulesCollectionDataAttributes {
	return v.value
}

func (v *NullableBillingSchedulesCollectionDataAttributes) Set(val *BillingSchedulesCollectionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingSchedulesCollectionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingSchedulesCollectionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingSchedulesCollectionDataAttributes(val *BillingSchedulesCollectionDataAttributes) *NullableBillingSchedulesCollectionDataAttributes {
	return &NullableBillingSchedulesCollectionDataAttributes{value: val, isSet: true}
}

func (v NullableBillingSchedulesCollectionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingSchedulesCollectionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


