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

// checks if the CheckoutSessionPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutSessionPaymentMethod{}

// CheckoutSessionPaymentMethod Details about the payment method for card type only.
type CheckoutSessionPaymentMethod struct {
	Type *string `json:"type,omitempty"`
	// Unique ID for the payment method.
	Id NullableString `json:"id,omitempty"`
	// Payment method type.
	Method *string `json:"method,omitempty"`
	// The scheme/brand of the card.
	Scheme NullableString `json:"scheme,omitempty"`
	// Last four digits of PAN.
	Label NullableString `json:"label,omitempty"`
	Details NullableCheckoutSessionPaymentMethodDetails `json:"details,omitempty"`
	// The unique hash derived from the payment method identifier (e.g. card number).
	Fingerprint NullableString `json:"fingerprint,omitempty"`
}

// NewCheckoutSessionPaymentMethod instantiates a new CheckoutSessionPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionPaymentMethod() *CheckoutSessionPaymentMethod {
	this := CheckoutSessionPaymentMethod{}
	return &this
}

// NewCheckoutSessionPaymentMethodWithDefaults instantiates a new CheckoutSessionPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionPaymentMethodWithDefaults() *CheckoutSessionPaymentMethod {
	this := CheckoutSessionPaymentMethod{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CheckoutSessionPaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionPaymentMethod) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionPaymentMethod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CheckoutSessionPaymentMethod) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CheckoutSessionPaymentMethod) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CheckoutSessionPaymentMethod) UnsetId() {
	o.Id.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethod) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethod) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CheckoutSessionPaymentMethod) SetMethod(v string) {
	o.Method = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionPaymentMethod) GetScheme() string {
	if o == nil || IsNil(o.Scheme.Get()) {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionPaymentMethod) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *CheckoutSessionPaymentMethod) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *CheckoutSessionPaymentMethod) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *CheckoutSessionPaymentMethod) UnsetScheme() {
	o.Scheme.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionPaymentMethod) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionPaymentMethod) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *CheckoutSessionPaymentMethod) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *CheckoutSessionPaymentMethod) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *CheckoutSessionPaymentMethod) UnsetLabel() {
	o.Label.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionPaymentMethod) GetDetails() CheckoutSessionPaymentMethodDetails {
	if o == nil || IsNil(o.Details.Get()) {
		var ret CheckoutSessionPaymentMethodDetails
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionPaymentMethod) GetDetailsOk() (*CheckoutSessionPaymentMethodDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableCheckoutSessionPaymentMethodDetails and assigns it to the Details field.
func (o *CheckoutSessionPaymentMethod) SetDetails(v CheckoutSessionPaymentMethodDetails) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *CheckoutSessionPaymentMethod) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *CheckoutSessionPaymentMethod) UnsetDetails() {
	o.Details.Unset()
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionPaymentMethod) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionPaymentMethod) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethod) HasFingerprint() bool {
	if o != nil && o.Fingerprint.IsSet() {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given NullableString and assigns it to the Fingerprint field.
func (o *CheckoutSessionPaymentMethod) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}
// SetFingerprintNil sets the value for Fingerprint to be an explicit nil
func (o *CheckoutSessionPaymentMethod) SetFingerprintNil() {
	o.Fingerprint.Set(nil)
}

// UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
func (o *CheckoutSessionPaymentMethod) UnsetFingerprint() {
	o.Fingerprint.Unset()
}

func (o CheckoutSessionPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.Fingerprint.IsSet() {
		toSerialize["fingerprint"] = o.Fingerprint.Get()
	}
	return toSerialize, nil
}

type NullableCheckoutSessionPaymentMethod struct {
	value *CheckoutSessionPaymentMethod
	isSet bool
}

func (v NullableCheckoutSessionPaymentMethod) Get() *CheckoutSessionPaymentMethod {
	return v.value
}

func (v *NullableCheckoutSessionPaymentMethod) Set(val *CheckoutSessionPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionPaymentMethod(val *CheckoutSessionPaymentMethod) *NullableCheckoutSessionPaymentMethod {
	return &NullableCheckoutSessionPaymentMethod{value: val, isSet: true}
}

func (v NullableCheckoutSessionPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


