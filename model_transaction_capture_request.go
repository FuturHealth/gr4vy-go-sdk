/*
Gr4vy API

Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.

API version: 1.1.0-beta
Contact: code@gr4vy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransactionCaptureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionCaptureRequest{}

// TransactionCaptureRequest A request to capture a transaction.
type TransactionCaptureRequest struct {
	// The monetary amount to capture an authorization for, in the smallest currency unit for the given currency, for example `1299` cents to create an authorization for `$12.99`.  When omitted blank, this will capture the entire amount.
	Amount *int32 `json:"amount,omitempty"`
	// The airline addendum data which describes the airline booking associated with this transaction. When provided, this will override any airline data provided when authorizing the transaction. Only the data on this request will be passed to the payment service, and none of the original data will be sent or kept.
	Airline NullableAirline `json:"airline,omitempty"`
}

// NewTransactionCaptureRequest instantiates a new TransactionCaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCaptureRequest() *TransactionCaptureRequest {
	this := TransactionCaptureRequest{}
	return &this
}

// NewTransactionCaptureRequestWithDefaults instantiates a new TransactionCaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCaptureRequestWithDefaults() *TransactionCaptureRequest {
	this := TransactionCaptureRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionCaptureRequest) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptureRequest) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionCaptureRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *TransactionCaptureRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetAirline returns the Airline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionCaptureRequest) GetAirline() Airline {
	if o == nil || IsNil(o.Airline.Get()) {
		var ret Airline
		return ret
	}
	return *o.Airline.Get()
}

// GetAirlineOk returns a tuple with the Airline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionCaptureRequest) GetAirlineOk() (*Airline, bool) {
	if o == nil {
		return nil, false
	}
	return o.Airline.Get(), o.Airline.IsSet()
}

// HasAirline returns a boolean if a field has been set.
func (o *TransactionCaptureRequest) HasAirline() bool {
	if o != nil && o.Airline.IsSet() {
		return true
	}

	return false
}

// SetAirline gets a reference to the given NullableAirline and assigns it to the Airline field.
func (o *TransactionCaptureRequest) SetAirline(v Airline) {
	o.Airline.Set(&v)
}
// SetAirlineNil sets the value for Airline to be an explicit nil
func (o *TransactionCaptureRequest) SetAirlineNil() {
	o.Airline.Set(nil)
}

// UnsetAirline ensures that no value is present for Airline, not even an explicit nil
func (o *TransactionCaptureRequest) UnsetAirline() {
	o.Airline.Unset()
}

func (o TransactionCaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionCaptureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if o.Airline.IsSet() {
		toSerialize["airline"] = o.Airline.Get()
	}
	return toSerialize, nil
}

type NullableTransactionCaptureRequest struct {
	value *TransactionCaptureRequest
	isSet bool
}

func (v NullableTransactionCaptureRequest) Get() *TransactionCaptureRequest {
	return v.value
}

func (v *NullableTransactionCaptureRequest) Set(val *TransactionCaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCaptureRequest(val *TransactionCaptureRequest) *NullableTransactionCaptureRequest {
	return &NullableTransactionCaptureRequest{value: val, isSet: true}
}

func (v NullableTransactionCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


