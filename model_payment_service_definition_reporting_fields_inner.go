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

// checks if the PaymentServiceDefinitionReportingFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServiceDefinitionReportingFieldsInner{}

// PaymentServiceDefinitionReportingFieldsInner A single field that needs to be submitted for a payment service when the settlement reporting is being enabled for it.
type PaymentServiceDefinitionReportingFieldsInner struct {
	// The key of a field that needs to be submitted.
	Key *string `json:"key,omitempty"`
	// The name to display for a field in the dashboard.
	DisplayName *string `json:"display_name,omitempty"`
	// Defines if this field is required when the reporting is being enabled.
	Required *bool `json:"required,omitempty"`
	// Defines the type of input that needs to be rendered for this field.
	Format *string `json:"format,omitempty"`
	// Defines if this field is secret. When `true` the field is not returned when querying the payment service.
	Secret *bool `json:"secret,omitempty"`
}

// NewPaymentServiceDefinitionReportingFieldsInner instantiates a new PaymentServiceDefinitionReportingFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceDefinitionReportingFieldsInner() *PaymentServiceDefinitionReportingFieldsInner {
	this := PaymentServiceDefinitionReportingFieldsInner{}
	return &this
}

// NewPaymentServiceDefinitionReportingFieldsInnerWithDefaults instantiates a new PaymentServiceDefinitionReportingFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceDefinitionReportingFieldsInnerWithDefaults() *PaymentServiceDefinitionReportingFieldsInner {
	this := PaymentServiceDefinitionReportingFieldsInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PaymentServiceDefinitionReportingFieldsInner) SetKey(v string) {
	o.Key = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PaymentServiceDefinitionReportingFieldsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PaymentServiceDefinitionReportingFieldsInner) SetRequired(v bool) {
	o.Required = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PaymentServiceDefinitionReportingFieldsInner) SetFormat(v string) {
	o.Format = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetSecret() bool {
	if o == nil || IsNil(o.Secret) {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) GetSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionReportingFieldsInner) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *PaymentServiceDefinitionReportingFieldsInner) SetSecret(v bool) {
	o.Secret = &v
}

func (o PaymentServiceDefinitionReportingFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServiceDefinitionReportingFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullablePaymentServiceDefinitionReportingFieldsInner struct {
	value *PaymentServiceDefinitionReportingFieldsInner
	isSet bool
}

func (v NullablePaymentServiceDefinitionReportingFieldsInner) Get() *PaymentServiceDefinitionReportingFieldsInner {
	return v.value
}

func (v *NullablePaymentServiceDefinitionReportingFieldsInner) Set(val *PaymentServiceDefinitionReportingFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceDefinitionReportingFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceDefinitionReportingFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceDefinitionReportingFieldsInner(val *PaymentServiceDefinitionReportingFieldsInner) *NullablePaymentServiceDefinitionReportingFieldsInner {
	return &NullablePaymentServiceDefinitionReportingFieldsInner{value: val, isSet: true}
}

func (v NullablePaymentServiceDefinitionReportingFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceDefinitionReportingFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


