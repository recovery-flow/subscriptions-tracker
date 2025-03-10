/*
User storage service

User storage service for recovery flow

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the SubscriptionPlanUpdateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlanUpdateDataAttributes{}

// SubscriptionPlanUpdateDataAttributes struct for SubscriptionPlanUpdateDataAttributes
type SubscriptionPlanUpdateDataAttributes struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Description
	Desc *string `json:"desc,omitempty"`
	// Price
	Price *float32 `json:"price,omitempty"`
	// Currency
	Currency *string `json:"currency,omitempty"`
}

// NewSubscriptionPlanUpdateDataAttributes instantiates a new SubscriptionPlanUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlanUpdateDataAttributes() *SubscriptionPlanUpdateDataAttributes {
	this := SubscriptionPlanUpdateDataAttributes{}
	return &this
}

// NewSubscriptionPlanUpdateDataAttributesWithDefaults instantiates a new SubscriptionPlanUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlanUpdateDataAttributesWithDefaults() *SubscriptionPlanUpdateDataAttributes {
	this := SubscriptionPlanUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriptionPlanUpdateDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionPlanUpdateDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriptionPlanUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *SubscriptionPlanUpdateDataAttributes) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanUpdateDataAttributes) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *SubscriptionPlanUpdateDataAttributes) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *SubscriptionPlanUpdateDataAttributes) SetDesc(v string) {
	o.Desc = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SubscriptionPlanUpdateDataAttributes) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanUpdateDataAttributes) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SubscriptionPlanUpdateDataAttributes) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *SubscriptionPlanUpdateDataAttributes) SetPrice(v float32) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SubscriptionPlanUpdateDataAttributes) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanUpdateDataAttributes) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SubscriptionPlanUpdateDataAttributes) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SubscriptionPlanUpdateDataAttributes) SetCurrency(v string) {
	o.Currency = &v
}

func (o SubscriptionPlanUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlanUpdateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableSubscriptionPlanUpdateDataAttributes struct {
	value *SubscriptionPlanUpdateDataAttributes
	isSet bool
}

func (v NullableSubscriptionPlanUpdateDataAttributes) Get() *SubscriptionPlanUpdateDataAttributes {
	return v.value
}

func (v *NullableSubscriptionPlanUpdateDataAttributes) Set(val *SubscriptionPlanUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlanUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlanUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlanUpdateDataAttributes(val *SubscriptionPlanUpdateDataAttributes) *NullableSubscriptionPlanUpdateDataAttributes {
	return &NullableSubscriptionPlanUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionPlanUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlanUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


