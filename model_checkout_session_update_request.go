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

// checks if the CheckoutSessionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutSessionUpdateRequest{}

// CheckoutSessionUpdateRequest A request to update a checkout session.
type CheckoutSessionUpdateRequest struct {
	// An array of cart items that represents the line items of a transaction.
	CartItems []CartItem `json:"cart_items,omitempty"`
	// Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The airline addendum data which describes the airline booking associated with this transaction.
	Airline NullableAirline `json:"airline,omitempty"`
}

// NewCheckoutSessionUpdateRequest instantiates a new CheckoutSessionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionUpdateRequest() *CheckoutSessionUpdateRequest {
	this := CheckoutSessionUpdateRequest{}
	return &this
}

// NewCheckoutSessionUpdateRequestWithDefaults instantiates a new CheckoutSessionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionUpdateRequestWithDefaults() *CheckoutSessionUpdateRequest {
	this := CheckoutSessionUpdateRequest{}
	return &this
}

// GetCartItems returns the CartItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionUpdateRequest) GetCartItems() []CartItem {
	if o == nil {
		var ret []CartItem
		return ret
	}
	return o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionUpdateRequest) GetCartItemsOk() ([]CartItem, bool) {
	if o == nil || IsNil(o.CartItems) {
		return nil, false
	}
	return o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *CheckoutSessionUpdateRequest) HasCartItems() bool {
	if o != nil && !IsNil(o.CartItems) {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *CheckoutSessionUpdateRequest) SetCartItems(v []CartItem) {
	o.CartItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionUpdateRequest) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionUpdateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckoutSessionUpdateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CheckoutSessionUpdateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetAirline returns the Airline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionUpdateRequest) GetAirline() Airline {
	if o == nil || IsNil(o.Airline.Get()) {
		var ret Airline
		return ret
	}
	return *o.Airline.Get()
}

// GetAirlineOk returns a tuple with the Airline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionUpdateRequest) GetAirlineOk() (*Airline, bool) {
	if o == nil {
		return nil, false
	}
	return o.Airline.Get(), o.Airline.IsSet()
}

// HasAirline returns a boolean if a field has been set.
func (o *CheckoutSessionUpdateRequest) HasAirline() bool {
	if o != nil && o.Airline.IsSet() {
		return true
	}

	return false
}

// SetAirline gets a reference to the given NullableAirline and assigns it to the Airline field.
func (o *CheckoutSessionUpdateRequest) SetAirline(v Airline) {
	o.Airline.Set(&v)
}
// SetAirlineNil sets the value for Airline to be an explicit nil
func (o *CheckoutSessionUpdateRequest) SetAirlineNil() {
	o.Airline.Set(nil)
}

// UnsetAirline ensures that no value is present for Airline, not even an explicit nil
func (o *CheckoutSessionUpdateRequest) UnsetAirline() {
	o.Airline.Unset()
}

func (o CheckoutSessionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CartItems != nil {
		toSerialize["cart_items"] = o.CartItems
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Airline.IsSet() {
		toSerialize["airline"] = o.Airline.Get()
	}
	return toSerialize, nil
}

type NullableCheckoutSessionUpdateRequest struct {
	value *CheckoutSessionUpdateRequest
	isSet bool
}

func (v NullableCheckoutSessionUpdateRequest) Get() *CheckoutSessionUpdateRequest {
	return v.value
}

func (v *NullableCheckoutSessionUpdateRequest) Set(val *CheckoutSessionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionUpdateRequest(val *CheckoutSessionUpdateRequest) *NullableCheckoutSessionUpdateRequest {
	return &NullableCheckoutSessionUpdateRequest{value: val, isSet: true}
}

func (v NullableCheckoutSessionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


