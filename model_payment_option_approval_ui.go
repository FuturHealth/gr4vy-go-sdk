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

// checks if the PaymentOptionApprovalUI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentOptionApprovalUI{}

// PaymentOptionApprovalUI Configuration for the approval interface that should be shown to the buyer.
type PaymentOptionApprovalUI struct {
	// Height of the approval interface in either pixels or view height (vh).
	Height *string `json:"height,omitempty"`
	// Width of the approval interface in either pixels or view width (vw).
	Width *string `json:"width,omitempty"`
}

// NewPaymentOptionApprovalUI instantiates a new PaymentOptionApprovalUI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOptionApprovalUI() *PaymentOptionApprovalUI {
	this := PaymentOptionApprovalUI{}
	return &this
}

// NewPaymentOptionApprovalUIWithDefaults instantiates a new PaymentOptionApprovalUI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOptionApprovalUIWithDefaults() *PaymentOptionApprovalUI {
	this := PaymentOptionApprovalUI{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *PaymentOptionApprovalUI) GetHeight() string {
	if o == nil || IsNil(o.Height) {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOptionApprovalUI) GetHeightOk() (*string, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *PaymentOptionApprovalUI) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *PaymentOptionApprovalUI) SetHeight(v string) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *PaymentOptionApprovalUI) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOptionApprovalUI) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *PaymentOptionApprovalUI) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *PaymentOptionApprovalUI) SetWidth(v string) {
	o.Width = &v
}

func (o PaymentOptionApprovalUI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentOptionApprovalUI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullablePaymentOptionApprovalUI struct {
	value *PaymentOptionApprovalUI
	isSet bool
}

func (v NullablePaymentOptionApprovalUI) Get() *PaymentOptionApprovalUI {
	return v.value
}

func (v *NullablePaymentOptionApprovalUI) Set(val *PaymentOptionApprovalUI) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOptionApprovalUI) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOptionApprovalUI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOptionApprovalUI(val *PaymentOptionApprovalUI) *NullablePaymentOptionApprovalUI {
	return &NullablePaymentOptionApprovalUI{value: val, isSet: true}
}

func (v NullablePaymentOptionApprovalUI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOptionApprovalUI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


