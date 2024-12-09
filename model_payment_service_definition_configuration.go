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

// checks if the PaymentServiceDefinitionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServiceDefinitionConfiguration{}

// PaymentServiceDefinitionConfiguration Configuration items for the payment service.
type PaymentServiceDefinitionConfiguration struct {
	// Height of the approval interface in either pixels or view height (vh).
	ApprovalUiHeight *string `json:"approval_ui_height,omitempty"`
	// Width of the approval interface in either pixels or view width (vw).
	ApprovalUiWidth *string `json:"approval_ui_width,omitempty"`
	// The browser target that an approval URL must be opened in. If `any` or `null`, then there is no specific requirement.
	ApprovalUiTarget NullableString `json:"approval_ui_target,omitempty"`
}

// NewPaymentServiceDefinitionConfiguration instantiates a new PaymentServiceDefinitionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceDefinitionConfiguration() *PaymentServiceDefinitionConfiguration {
	this := PaymentServiceDefinitionConfiguration{}
	return &this
}

// NewPaymentServiceDefinitionConfigurationWithDefaults instantiates a new PaymentServiceDefinitionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceDefinitionConfigurationWithDefaults() *PaymentServiceDefinitionConfiguration {
	this := PaymentServiceDefinitionConfiguration{}
	return &this
}

// GetApprovalUiHeight returns the ApprovalUiHeight field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiHeight() string {
	if o == nil || IsNil(o.ApprovalUiHeight) {
		var ret string
		return ret
	}
	return *o.ApprovalUiHeight
}

// GetApprovalUiHeightOk returns a tuple with the ApprovalUiHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiHeightOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalUiHeight) {
		return nil, false
	}
	return o.ApprovalUiHeight, true
}

// HasApprovalUiHeight returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionConfiguration) HasApprovalUiHeight() bool {
	if o != nil && !IsNil(o.ApprovalUiHeight) {
		return true
	}

	return false
}

// SetApprovalUiHeight gets a reference to the given string and assigns it to the ApprovalUiHeight field.
func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiHeight(v string) {
	o.ApprovalUiHeight = &v
}

// GetApprovalUiWidth returns the ApprovalUiWidth field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiWidth() string {
	if o == nil || IsNil(o.ApprovalUiWidth) {
		var ret string
		return ret
	}
	return *o.ApprovalUiWidth
}

// GetApprovalUiWidthOk returns a tuple with the ApprovalUiWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiWidthOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalUiWidth) {
		return nil, false
	}
	return o.ApprovalUiWidth, true
}

// HasApprovalUiWidth returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionConfiguration) HasApprovalUiWidth() bool {
	if o != nil && !IsNil(o.ApprovalUiWidth) {
		return true
	}

	return false
}

// SetApprovalUiWidth gets a reference to the given string and assigns it to the ApprovalUiWidth field.
func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiWidth(v string) {
	o.ApprovalUiWidth = &v
}

// GetApprovalUiTarget returns the ApprovalUiTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiTarget() string {
	if o == nil || IsNil(o.ApprovalUiTarget.Get()) {
		var ret string
		return ret
	}
	return *o.ApprovalUiTarget.Get()
}

// GetApprovalUiTargetOk returns a tuple with the ApprovalUiTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovalUiTarget.Get(), o.ApprovalUiTarget.IsSet()
}

// HasApprovalUiTarget returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionConfiguration) HasApprovalUiTarget() bool {
	if o != nil && o.ApprovalUiTarget.IsSet() {
		return true
	}

	return false
}

// SetApprovalUiTarget gets a reference to the given NullableString and assigns it to the ApprovalUiTarget field.
func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiTarget(v string) {
	o.ApprovalUiTarget.Set(&v)
}
// SetApprovalUiTargetNil sets the value for ApprovalUiTarget to be an explicit nil
func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiTargetNil() {
	o.ApprovalUiTarget.Set(nil)
}

// UnsetApprovalUiTarget ensures that no value is present for ApprovalUiTarget, not even an explicit nil
func (o *PaymentServiceDefinitionConfiguration) UnsetApprovalUiTarget() {
	o.ApprovalUiTarget.Unset()
}

func (o PaymentServiceDefinitionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServiceDefinitionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApprovalUiHeight) {
		toSerialize["approval_ui_height"] = o.ApprovalUiHeight
	}
	if !IsNil(o.ApprovalUiWidth) {
		toSerialize["approval_ui_width"] = o.ApprovalUiWidth
	}
	if o.ApprovalUiTarget.IsSet() {
		toSerialize["approval_ui_target"] = o.ApprovalUiTarget.Get()
	}
	return toSerialize, nil
}

type NullablePaymentServiceDefinitionConfiguration struct {
	value *PaymentServiceDefinitionConfiguration
	isSet bool
}

func (v NullablePaymentServiceDefinitionConfiguration) Get() *PaymentServiceDefinitionConfiguration {
	return v.value
}

func (v *NullablePaymentServiceDefinitionConfiguration) Set(val *PaymentServiceDefinitionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceDefinitionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceDefinitionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceDefinitionConfiguration(val *PaymentServiceDefinitionConfiguration) *NullablePaymentServiceDefinitionConfiguration {
	return &NullablePaymentServiceDefinitionConfiguration{value: val, isSet: true}
}

func (v NullablePaymentServiceDefinitionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceDefinitionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


