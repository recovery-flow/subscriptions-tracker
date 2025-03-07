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

// checks if the SubscriptionCreateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateDataAttributes{}

// SubscriptionCreateDataAttributes struct for SubscriptionCreateDataAttributes
type SubscriptionCreateDataAttributes struct {
	// Plan ID
	PlanId string `json:"plan_id"`
	// Payment Method ID
	PaymentMethodId string `json:"payment_method_id"`
}

type _SubscriptionCreateDataAttributes SubscriptionCreateDataAttributes

// NewSubscriptionCreateDataAttributes instantiates a new SubscriptionCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateDataAttributes(planId string, paymentMethodId string) *SubscriptionCreateDataAttributes {
	this := SubscriptionCreateDataAttributes{}
	this.PlanId = planId
	this.PaymentMethodId = paymentMethodId
	return &this
}

// NewSubscriptionCreateDataAttributesWithDefaults instantiates a new SubscriptionCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateDataAttributesWithDefaults() *SubscriptionCreateDataAttributes {
	this := SubscriptionCreateDataAttributes{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *SubscriptionCreateDataAttributes) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateDataAttributes) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *SubscriptionCreateDataAttributes) SetPlanId(v string) {
	o.PlanId = v
}

// GetPaymentMethodId returns the PaymentMethodId field value
func (o *SubscriptionCreateDataAttributes) GetPaymentMethodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateDataAttributes) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodId, true
}

// SetPaymentMethodId sets field value
func (o *SubscriptionCreateDataAttributes) SetPaymentMethodId(v string) {
	o.PaymentMethodId = v
}

func (o SubscriptionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan_id"] = o.PlanId
	toSerialize["payment_method_id"] = o.PaymentMethodId
	return toSerialize, nil
}

func (o *SubscriptionCreateDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan_id",
		"payment_method_id",
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

	varSubscriptionCreateDataAttributes := _SubscriptionCreateDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionCreateDataAttributes)

	if err != nil {
		return err
	}

	*o = SubscriptionCreateDataAttributes(varSubscriptionCreateDataAttributes)

	return err
}

type NullableSubscriptionCreateDataAttributes struct {
	value *SubscriptionCreateDataAttributes
	isSet bool
}

func (v NullableSubscriptionCreateDataAttributes) Get() *SubscriptionCreateDataAttributes {
	return v.value
}

func (v *NullableSubscriptionCreateDataAttributes) Set(val *SubscriptionCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateDataAttributes(val *SubscriptionCreateDataAttributes) *NullableSubscriptionCreateDataAttributes {
	return &NullableSubscriptionCreateDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


