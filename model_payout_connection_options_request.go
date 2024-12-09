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

// checks if the PayoutConnectionOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayoutConnectionOptionsRequest{}

// PayoutConnectionOptionsRequest struct for PayoutConnectionOptionsRequest
type PayoutConnectionOptionsRequest struct {
	CheckoutCard NullablePayoutConnectionOptionsRequestCheckoutCard `json:"checkout-card,omitempty"`
}

// NewPayoutConnectionOptionsRequest instantiates a new PayoutConnectionOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutConnectionOptionsRequest() *PayoutConnectionOptionsRequest {
	this := PayoutConnectionOptionsRequest{}
	return &this
}

// NewPayoutConnectionOptionsRequestWithDefaults instantiates a new PayoutConnectionOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutConnectionOptionsRequestWithDefaults() *PayoutConnectionOptionsRequest {
	this := PayoutConnectionOptionsRequest{}
	return &this
}

// GetCheckoutCard returns the CheckoutCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutConnectionOptionsRequest) GetCheckoutCard() PayoutConnectionOptionsRequestCheckoutCard {
	if o == nil || IsNil(o.CheckoutCard.Get()) {
		var ret PayoutConnectionOptionsRequestCheckoutCard
		return ret
	}
	return *o.CheckoutCard.Get()
}

// GetCheckoutCardOk returns a tuple with the CheckoutCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutConnectionOptionsRequest) GetCheckoutCardOk() (*PayoutConnectionOptionsRequestCheckoutCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutCard.Get(), o.CheckoutCard.IsSet()
}

// HasCheckoutCard returns a boolean if a field has been set.
func (o *PayoutConnectionOptionsRequest) HasCheckoutCard() bool {
	if o != nil && o.CheckoutCard.IsSet() {
		return true
	}

	return false
}

// SetCheckoutCard gets a reference to the given NullablePayoutConnectionOptionsRequestCheckoutCard and assigns it to the CheckoutCard field.
func (o *PayoutConnectionOptionsRequest) SetCheckoutCard(v PayoutConnectionOptionsRequestCheckoutCard) {
	o.CheckoutCard.Set(&v)
}
// SetCheckoutCardNil sets the value for CheckoutCard to be an explicit nil
func (o *PayoutConnectionOptionsRequest) SetCheckoutCardNil() {
	o.CheckoutCard.Set(nil)
}

// UnsetCheckoutCard ensures that no value is present for CheckoutCard, not even an explicit nil
func (o *PayoutConnectionOptionsRequest) UnsetCheckoutCard() {
	o.CheckoutCard.Unset()
}

func (o PayoutConnectionOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutConnectionOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckoutCard.IsSet() {
		toSerialize["checkout-card"] = o.CheckoutCard.Get()
	}
	return toSerialize, nil
}

type NullablePayoutConnectionOptionsRequest struct {
	value *PayoutConnectionOptionsRequest
	isSet bool
}

func (v NullablePayoutConnectionOptionsRequest) Get() *PayoutConnectionOptionsRequest {
	return v.value
}

func (v *NullablePayoutConnectionOptionsRequest) Set(val *PayoutConnectionOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutConnectionOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutConnectionOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutConnectionOptionsRequest(val *PayoutConnectionOptionsRequest) *NullablePayoutConnectionOptionsRequest {
	return &NullablePayoutConnectionOptionsRequest{value: val, isSet: true}
}

func (v NullablePayoutConnectionOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutConnectionOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


