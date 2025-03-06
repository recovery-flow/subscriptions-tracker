/*
User storage service

User storage service for recovery flow

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TransactionDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDataAttributes{}

// TransactionDataAttributes struct for TransactionDataAttributes
type TransactionDataAttributes struct {
	// User ID
	UserId string `json:"user_id"`
	// Plan ID
	PlanId string `json:"plan_id"`
	// Payment ID
	PaymentId string `json:"payment_id"`
	// Transaction amount
	Amount float32 `json:"amount"`
	// Transaction currency
	Currency string `json:"currency"`
	// Provider transaction ID
	ProviderTransactionId string `json:"provider_transaction_id"`
	// Transaction creation date
	TransactionDate time.Time `json:"transaction_date"`
}

type _TransactionDataAttributes TransactionDataAttributes

// NewTransactionDataAttributes instantiates a new TransactionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDataAttributes(userId string, planId string, paymentId string, amount float32, currency string, providerTransactionId string, transactionDate time.Time) *TransactionDataAttributes {
	this := TransactionDataAttributes{}
	this.UserId = userId
	this.PlanId = planId
	this.PaymentId = paymentId
	this.Amount = amount
	this.Currency = currency
	this.ProviderTransactionId = providerTransactionId
	this.TransactionDate = transactionDate
	return &this
}

// NewTransactionDataAttributesWithDefaults instantiates a new TransactionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDataAttributesWithDefaults() *TransactionDataAttributes {
	this := TransactionDataAttributes{}
	return &this
}

// GetUserId returns the UserId field value
func (o *TransactionDataAttributes) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TransactionDataAttributes) SetUserId(v string) {
	o.UserId = v
}

// GetPlanId returns the PlanId field value
func (o *TransactionDataAttributes) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *TransactionDataAttributes) SetPlanId(v string) {
	o.PlanId = v
}

// GetPaymentId returns the PaymentId field value
func (o *TransactionDataAttributes) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *TransactionDataAttributes) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetAmount returns the Amount field value
func (o *TransactionDataAttributes) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionDataAttributes) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransactionDataAttributes) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionDataAttributes) SetCurrency(v string) {
	o.Currency = v
}

// GetProviderTransactionId returns the ProviderTransactionId field value
func (o *TransactionDataAttributes) GetProviderTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderTransactionId
}

// GetProviderTransactionIdOk returns a tuple with the ProviderTransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetProviderTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderTransactionId, true
}

// SetProviderTransactionId sets field value
func (o *TransactionDataAttributes) SetProviderTransactionId(v string) {
	o.ProviderTransactionId = v
}

// GetTransactionDate returns the TransactionDate field value
func (o *TransactionDataAttributes) GetTransactionDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value
// and a boolean to check if the value has been set.
func (o *TransactionDataAttributes) GetTransactionDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionDate, true
}

// SetTransactionDate sets field value
func (o *TransactionDataAttributes) SetTransactionDate(v time.Time) {
	o.TransactionDate = v
}

func (o TransactionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["plan_id"] = o.PlanId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["provider_transaction_id"] = o.ProviderTransactionId
	toSerialize["transaction_date"] = o.TransactionDate
	return toSerialize, nil
}

func (o *TransactionDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"plan_id",
		"payment_id",
		"amount",
		"currency",
		"provider_transaction_id",
		"transaction_date",
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

	varTransactionDataAttributes := _TransactionDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDataAttributes)

	if err != nil {
		return err
	}

	*o = TransactionDataAttributes(varTransactionDataAttributes)

	return err
}

type NullableTransactionDataAttributes struct {
	value *TransactionDataAttributes
	isSet bool
}

func (v NullableTransactionDataAttributes) Get() *TransactionDataAttributes {
	return v.value
}

func (v *NullableTransactionDataAttributes) Set(val *TransactionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDataAttributes(val *TransactionDataAttributes) *NullableTransactionDataAttributes {
	return &NullableTransactionDataAttributes{value: val, isSet: true}
}

func (v NullableTransactionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


