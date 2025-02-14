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

// checks if the SubscribersCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscribersCollection{}

// SubscribersCollection struct for SubscribersCollection
type SubscribersCollection struct {
	Data []SubscriberData `json:"data"`
	Links LinksPagination `json:"links"`
}

type _SubscribersCollection SubscribersCollection

// NewSubscribersCollection instantiates a new SubscribersCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribersCollection(data []SubscriberData, links LinksPagination) *SubscribersCollection {
	this := SubscribersCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscribersCollectionWithDefaults instantiates a new SubscribersCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribersCollectionWithDefaults() *SubscribersCollection {
	this := SubscribersCollection{}
	return &this
}

// GetData returns the Data field value
func (o *SubscribersCollection) GetData() []SubscriberData {
	if o == nil {
		var ret []SubscriberData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscribersCollection) GetDataOk() ([]SubscriberData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscribersCollection) SetData(v []SubscriberData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *SubscribersCollection) GetLinks() LinksPagination {
	if o == nil {
		var ret LinksPagination
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscribersCollection) GetLinksOk() (*LinksPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscribersCollection) SetLinks(v LinksPagination) {
	o.Links = v
}

func (o SubscribersCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribersCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SubscribersCollection) UnmarshalJSON(data []byte) (err error) {
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

	varSubscribersCollection := _SubscribersCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscribersCollection)

	if err != nil {
		return err
	}

	*o = SubscribersCollection(varSubscribersCollection)

	return err
}

type NullableSubscribersCollection struct {
	value *SubscribersCollection
	isSet bool
}

func (v NullableSubscribersCollection) Get() *SubscribersCollection {
	return v.value
}

func (v *NullableSubscribersCollection) Set(val *SubscribersCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribersCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribersCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribersCollection(val *SubscribersCollection) *NullableSubscribersCollection {
	return &NullableSubscribersCollection{value: val, isSet: true}
}

func (v NullableSubscribersCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribersCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


