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

// checks if the PaymentMethodCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCreateData{}

// PaymentMethodCreateData struct for PaymentMethodCreateData
type PaymentMethodCreateData struct {
	Type string `json:"type"`
	Attributes PaymentMethodCreateDataAttributes `json:"attributes"`
}

type _PaymentMethodCreateData PaymentMethodCreateData

// NewPaymentMethodCreateData instantiates a new PaymentMethodCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCreateData(type_ string, attributes PaymentMethodCreateDataAttributes) *PaymentMethodCreateData {
	this := PaymentMethodCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaymentMethodCreateDataWithDefaults instantiates a new PaymentMethodCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCreateDataWithDefaults() *PaymentMethodCreateData {
	this := PaymentMethodCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaymentMethodCreateData) GetAttributes() PaymentMethodCreateDataAttributes {
	if o == nil {
		var ret PaymentMethodCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateData) GetAttributesOk() (*PaymentMethodCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaymentMethodCreateData) SetAttributes(v PaymentMethodCreateDataAttributes) {
	o.Attributes = v
}

func (o PaymentMethodCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *PaymentMethodCreateData) UnmarshalJSON(data []byte) (err error) {
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

	varPaymentMethodCreateData := _PaymentMethodCreateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentMethodCreateData)

	if err != nil {
		return err
	}

	*o = PaymentMethodCreateData(varPaymentMethodCreateData)

	return err
}

type NullablePaymentMethodCreateData struct {
	value *PaymentMethodCreateData
	isSet bool
}

func (v NullablePaymentMethodCreateData) Get() *PaymentMethodCreateData {
	return v.value
}

func (v *NullablePaymentMethodCreateData) Set(val *PaymentMethodCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCreateData(val *PaymentMethodCreateData) *NullablePaymentMethodCreateData {
	return &NullablePaymentMethodCreateData{value: val, isSet: true}
}

func (v NullablePaymentMethodCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


