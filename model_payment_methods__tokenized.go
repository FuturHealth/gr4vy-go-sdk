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

// checks if the PaymentMethodsTokenized type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsTokenized{}

// PaymentMethodsTokenized A list of stored payment methods in summarized format.
type PaymentMethodsTokenized struct {
	// A list of stored payment methods in summarized format.
	Items []PaymentMethodTokenized `json:"items,omitempty"`
}

// NewPaymentMethodsTokenized instantiates a new PaymentMethodsTokenized object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsTokenized() *PaymentMethodsTokenized {
	this := PaymentMethodsTokenized{}
	return &this
}

// NewPaymentMethodsTokenizedWithDefaults instantiates a new PaymentMethodsTokenized object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsTokenizedWithDefaults() *PaymentMethodsTokenized {
	this := PaymentMethodsTokenized{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PaymentMethodsTokenized) GetItems() []PaymentMethodTokenized {
	if o == nil || IsNil(o.Items) {
		var ret []PaymentMethodTokenized
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsTokenized) GetItemsOk() ([]PaymentMethodTokenized, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PaymentMethodsTokenized) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentMethodTokenized and assigns it to the Items field.
func (o *PaymentMethodsTokenized) SetItems(v []PaymentMethodTokenized) {
	o.Items = v
}

func (o PaymentMethodsTokenized) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsTokenized) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePaymentMethodsTokenized struct {
	value *PaymentMethodsTokenized
	isSet bool
}

func (v NullablePaymentMethodsTokenized) Get() *PaymentMethodsTokenized {
	return v.value
}

func (v *NullablePaymentMethodsTokenized) Set(val *PaymentMethodsTokenized) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsTokenized) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsTokenized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsTokenized(val *PaymentMethodsTokenized) *NullablePaymentMethodsTokenized {
	return &NullablePaymentMethodsTokenized{value: val, isSet: true}
}

func (v NullablePaymentMethodsTokenized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsTokenized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


