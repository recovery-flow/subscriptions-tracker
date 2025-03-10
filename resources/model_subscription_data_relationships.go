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

// checks if the SubscriptionDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDataRelationships{}

// SubscriptionDataRelationships struct for SubscriptionDataRelationships
type SubscriptionDataRelationships struct {
	Plan Relationships `json:"plan"`
	Type Relationships `json:"type"`
	PaymentMethod Relationships `json:"payment_method"`
}

type _SubscriptionDataRelationships SubscriptionDataRelationships

// NewSubscriptionDataRelationships instantiates a new SubscriptionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDataRelationships(plan Relationships, type_ Relationships, paymentMethod Relationships) *SubscriptionDataRelationships {
	this := SubscriptionDataRelationships{}
	this.Plan = plan
	this.Type = type_
	this.PaymentMethod = paymentMethod
	return &this
}

// NewSubscriptionDataRelationshipsWithDefaults instantiates a new SubscriptionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDataRelationshipsWithDefaults() *SubscriptionDataRelationships {
	this := SubscriptionDataRelationships{}
	return &this
}

// GetPlan returns the Plan field value
func (o *SubscriptionDataRelationships) GetPlan() Relationships {
	if o == nil {
		var ret Relationships
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataRelationships) GetPlanOk() (*Relationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *SubscriptionDataRelationships) SetPlan(v Relationships) {
	o.Plan = v
}

// GetType returns the Type field value
func (o *SubscriptionDataRelationships) GetType() Relationships {
	if o == nil {
		var ret Relationships
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataRelationships) GetTypeOk() (*Relationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionDataRelationships) SetType(v Relationships) {
	o.Type = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *SubscriptionDataRelationships) GetPaymentMethod() Relationships {
	if o == nil {
		var ret Relationships
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataRelationships) GetPaymentMethodOk() (*Relationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *SubscriptionDataRelationships) SetPaymentMethod(v Relationships) {
	o.PaymentMethod = v
}

func (o SubscriptionDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan"] = o.Plan
	toSerialize["type"] = o.Type
	toSerialize["payment_method"] = o.PaymentMethod
	return toSerialize, nil
}

func (o *SubscriptionDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan",
		"type",
		"payment_method",
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

	varSubscriptionDataRelationships := _SubscriptionDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionDataRelationships)

	if err != nil {
		return err
	}

	*o = SubscriptionDataRelationships(varSubscriptionDataRelationships)

	return err
}

type NullableSubscriptionDataRelationships struct {
	value *SubscriptionDataRelationships
	isSet bool
}

func (v NullableSubscriptionDataRelationships) Get() *SubscriptionDataRelationships {
	return v.value
}

func (v *NullableSubscriptionDataRelationships) Set(val *SubscriptionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDataRelationships(val *SubscriptionDataRelationships) *NullableSubscriptionDataRelationships {
	return &NullableSubscriptionDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


