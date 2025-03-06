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

// checks if the PaymentMethodsCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsCollection{}

// PaymentMethodsCollection struct for PaymentMethodsCollection
type PaymentMethodsCollection struct {
	Data []PaymentMethodData `json:"data"`
	Links LinksPagination `json:"links"`
}

type _PaymentMethodsCollection PaymentMethodsCollection

// NewPaymentMethodsCollection instantiates a new PaymentMethodsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsCollection(data []PaymentMethodData, links LinksPagination) *PaymentMethodsCollection {
	this := PaymentMethodsCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPaymentMethodsCollectionWithDefaults instantiates a new PaymentMethodsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsCollectionWithDefaults() *PaymentMethodsCollection {
	this := PaymentMethodsCollection{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentMethodsCollection) GetData() []PaymentMethodData {
	if o == nil {
		var ret []PaymentMethodData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodsCollection) GetDataOk() ([]PaymentMethodData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaymentMethodsCollection) SetData(v []PaymentMethodData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *PaymentMethodsCollection) GetLinks() LinksPagination {
	if o == nil {
		var ret LinksPagination
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodsCollection) GetLinksOk() (*LinksPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PaymentMethodsCollection) SetLinks(v LinksPagination) {
	o.Links = v
}

func (o PaymentMethodsCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *PaymentMethodsCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varPaymentMethodsCollection := _PaymentMethodsCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentMethodsCollection)

	if err != nil {
		return err
	}

	*o = PaymentMethodsCollection(varPaymentMethodsCollection)

	return err
}

type NullablePaymentMethodsCollection struct {
	value *PaymentMethodsCollection
	isSet bool
}

func (v NullablePaymentMethodsCollection) Get() *PaymentMethodsCollection {
	return v.value
}

func (v *NullablePaymentMethodsCollection) Set(val *PaymentMethodsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsCollection(val *PaymentMethodsCollection) *NullablePaymentMethodsCollection {
	return &NullablePaymentMethodsCollection{value: val, isSet: true}
}

func (v NullablePaymentMethodsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


