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

// checks if the PaymentConnectorResponseTransactionCaptureDeclinedEventContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentConnectorResponseTransactionCaptureDeclinedEventContext{}

// PaymentConnectorResponseTransactionCaptureDeclinedEventContext Additional context for this event.
type PaymentConnectorResponseTransactionCaptureDeclinedEventContext struct {
	// The unique ID of the payment service used.
	PaymentServiceId *string `json:"payment_service_id,omitempty"`
	// The display name of the payment service used.
	PaymentServiceDisplayName *string `json:"payment_service_display_name,omitempty"`
	// The payment service definition used.
	PaymentServiceDefinitionId *string `json:"payment_service_definition_id,omitempty"`
	// The external ID of the transaction as set by the payment service.
	PaymentServiceTransactionId NullableString `json:"payment_service_transaction_id,omitempty"`
	// A raw response code returned for the failure.
	Code NullableString `json:"code,omitempty"`
	// This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseCode NullableString `json:"raw_response_code,omitempty"`
	// This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseDescription NullableString `json:"raw_response_description,omitempty"`
	// The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - `no_match` - neither address or postal code match - `match` - both address and postal code match - `partial_match_address` - address matches but postal code does not - `partial_match_postcode` - postal code matches but address does not - `unavailable ` - AVS is unavailable for card/country  The value of this field can be `null` if the payment service did not provide a response.
	AvsResponseCode NullableString `json:"avs_response_code,omitempty"`
	// The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - `no_match` - the CVV does not match the expected value - `match` - the CVV matches the expected value - `unavailable ` - CVV check unavailable for card our country - `not_provided ` - CVV not provided  The value of this field can be `null` if the payment service did not provide a response.
	CvvResponseCode NullableString `json:"cvv_response_code,omitempty"`
	// The card scheme sent to the connector.
	PaymentMethodScheme NullableString `json:"payment_method_scheme,omitempty"`
}

// NewPaymentConnectorResponseTransactionCaptureDeclinedEventContext instantiates a new PaymentConnectorResponseTransactionCaptureDeclinedEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConnectorResponseTransactionCaptureDeclinedEventContext() *PaymentConnectorResponseTransactionCaptureDeclinedEventContext {
	this := PaymentConnectorResponseTransactionCaptureDeclinedEventContext{}
	return &this
}

// NewPaymentConnectorResponseTransactionCaptureDeclinedEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionCaptureDeclinedEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConnectorResponseTransactionCaptureDeclinedEventContextWithDefaults() *PaymentConnectorResponseTransactionCaptureDeclinedEventContext {
	this := PaymentConnectorResponseTransactionCaptureDeclinedEventContext{}
	return &this
}

// GetPaymentServiceId returns the PaymentServiceId field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceId() string {
	if o == nil || IsNil(o.PaymentServiceId) {
		var ret string
		return ret
	}
	return *o.PaymentServiceId
}

// GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceId) {
		return nil, false
	}
	return o.PaymentServiceId, true
}

// HasPaymentServiceId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceId() bool {
	if o != nil && !IsNil(o.PaymentServiceId) {
		return true
	}

	return false
}

// SetPaymentServiceId gets a reference to the given string and assigns it to the PaymentServiceId field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceId(v string) {
	o.PaymentServiceId = &v
}

// GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDisplayName() string {
	if o == nil || IsNil(o.PaymentServiceDisplayName) {
		var ret string
		return ret
	}
	return *o.PaymentServiceDisplayName
}

// GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceDisplayName) {
		return nil, false
	}
	return o.PaymentServiceDisplayName, true
}

// HasPaymentServiceDisplayName returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceDisplayName() bool {
	if o != nil && !IsNil(o.PaymentServiceDisplayName) {
		return true
	}

	return false
}

// SetPaymentServiceDisplayName gets a reference to the given string and assigns it to the PaymentServiceDisplayName field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceDisplayName(v string) {
	o.PaymentServiceDisplayName = &v
}

// GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDefinitionId() string {
	if o == nil || IsNil(o.PaymentServiceDefinitionId) {
		var ret string
		return ret
	}
	return *o.PaymentServiceDefinitionId
}

// GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceDefinitionId) {
		return nil, false
	}
	return o.PaymentServiceDefinitionId, true
}

// HasPaymentServiceDefinitionId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceDefinitionId() bool {
	if o != nil && !IsNil(o.PaymentServiceDefinitionId) {
		return true
	}

	return false
}

// SetPaymentServiceDefinitionId gets a reference to the given string and assigns it to the PaymentServiceDefinitionId field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceDefinitionId(v string) {
	o.PaymentServiceDefinitionId = &v
}

// GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceTransactionId() string {
	if o == nil || IsNil(o.PaymentServiceTransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentServiceTransactionId.Get()
}

// GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentServiceTransactionId.Get(), o.PaymentServiceTransactionId.IsSet()
}

// HasPaymentServiceTransactionId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceTransactionId() bool {
	if o != nil && o.PaymentServiceTransactionId.IsSet() {
		return true
	}

	return false
}

// SetPaymentServiceTransactionId gets a reference to the given NullableString and assigns it to the PaymentServiceTransactionId field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceTransactionId(v string) {
	o.PaymentServiceTransactionId.Set(&v)
}
// SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceTransactionIdNil() {
	o.PaymentServiceTransactionId.Set(nil)
}

// UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetPaymentServiceTransactionId() {
	o.PaymentServiceTransactionId.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetCode() {
	o.Code.Unset()
}

// GetRawResponseCode returns the RawResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseCode() string {
	if o == nil || IsNil(o.RawResponseCode.Get()) {
		var ret string
		return ret
	}
	return *o.RawResponseCode.Get()
}

// GetRawResponseCodeOk returns a tuple with the RawResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawResponseCode.Get(), o.RawResponseCode.IsSet()
}

// HasRawResponseCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasRawResponseCode() bool {
	if o != nil && o.RawResponseCode.IsSet() {
		return true
	}

	return false
}

// SetRawResponseCode gets a reference to the given NullableString and assigns it to the RawResponseCode field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseCode(v string) {
	o.RawResponseCode.Set(&v)
}
// SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseCodeNil() {
	o.RawResponseCode.Set(nil)
}

// UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetRawResponseCode() {
	o.RawResponseCode.Unset()
}

// GetRawResponseDescription returns the RawResponseDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseDescription() string {
	if o == nil || IsNil(o.RawResponseDescription.Get()) {
		var ret string
		return ret
	}
	return *o.RawResponseDescription.Get()
}

// GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawResponseDescription.Get(), o.RawResponseDescription.IsSet()
}

// HasRawResponseDescription returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasRawResponseDescription() bool {
	if o != nil && o.RawResponseDescription.IsSet() {
		return true
	}

	return false
}

// SetRawResponseDescription gets a reference to the given NullableString and assigns it to the RawResponseDescription field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseDescription(v string) {
	o.RawResponseDescription.Set(&v)
}
// SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseDescriptionNil() {
	o.RawResponseDescription.Set(nil)
}

// UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetRawResponseDescription() {
	o.RawResponseDescription.Unset()
}

// GetAvsResponseCode returns the AvsResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetAvsResponseCode() string {
	if o == nil || IsNil(o.AvsResponseCode.Get()) {
		var ret string
		return ret
	}
	return *o.AvsResponseCode.Get()
}

// GetAvsResponseCodeOk returns a tuple with the AvsResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetAvsResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvsResponseCode.Get(), o.AvsResponseCode.IsSet()
}

// HasAvsResponseCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasAvsResponseCode() bool {
	if o != nil && o.AvsResponseCode.IsSet() {
		return true
	}

	return false
}

// SetAvsResponseCode gets a reference to the given NullableString and assigns it to the AvsResponseCode field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetAvsResponseCode(v string) {
	o.AvsResponseCode.Set(&v)
}
// SetAvsResponseCodeNil sets the value for AvsResponseCode to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetAvsResponseCodeNil() {
	o.AvsResponseCode.Set(nil)
}

// UnsetAvsResponseCode ensures that no value is present for AvsResponseCode, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetAvsResponseCode() {
	o.AvsResponseCode.Unset()
}

// GetCvvResponseCode returns the CvvResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCvvResponseCode() string {
	if o == nil || IsNil(o.CvvResponseCode.Get()) {
		var ret string
		return ret
	}
	return *o.CvvResponseCode.Get()
}

// GetCvvResponseCodeOk returns a tuple with the CvvResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCvvResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CvvResponseCode.Get(), o.CvvResponseCode.IsSet()
}

// HasCvvResponseCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasCvvResponseCode() bool {
	if o != nil && o.CvvResponseCode.IsSet() {
		return true
	}

	return false
}

// SetCvvResponseCode gets a reference to the given NullableString and assigns it to the CvvResponseCode field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCvvResponseCode(v string) {
	o.CvvResponseCode.Set(&v)
}
// SetCvvResponseCodeNil sets the value for CvvResponseCode to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCvvResponseCodeNil() {
	o.CvvResponseCode.Set(nil)
}

// UnsetCvvResponseCode ensures that no value is present for CvvResponseCode, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetCvvResponseCode() {
	o.CvvResponseCode.Unset()
}

// GetPaymentMethodScheme returns the PaymentMethodScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentMethodScheme() string {
	if o == nil || IsNil(o.PaymentMethodScheme.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodScheme.Get()
}

// GetPaymentMethodSchemeOk returns a tuple with the PaymentMethodScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentMethodSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodScheme.Get(), o.PaymentMethodScheme.IsSet()
}

// HasPaymentMethodScheme returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentMethodScheme() bool {
	if o != nil && o.PaymentMethodScheme.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodScheme gets a reference to the given NullableString and assigns it to the PaymentMethodScheme field.
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentMethodScheme(v string) {
	o.PaymentMethodScheme.Set(&v)
}
// SetPaymentMethodSchemeNil sets the value for PaymentMethodScheme to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentMethodSchemeNil() {
	o.PaymentMethodScheme.Set(nil)
}

// UnsetPaymentMethodScheme ensures that no value is present for PaymentMethodScheme, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetPaymentMethodScheme() {
	o.PaymentMethodScheme.Unset()
}

func (o PaymentConnectorResponseTransactionCaptureDeclinedEventContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentConnectorResponseTransactionCaptureDeclinedEventContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentServiceId) {
		toSerialize["payment_service_id"] = o.PaymentServiceId
	}
	if !IsNil(o.PaymentServiceDisplayName) {
		toSerialize["payment_service_display_name"] = o.PaymentServiceDisplayName
	}
	if !IsNil(o.PaymentServiceDefinitionId) {
		toSerialize["payment_service_definition_id"] = o.PaymentServiceDefinitionId
	}
	if o.PaymentServiceTransactionId.IsSet() {
		toSerialize["payment_service_transaction_id"] = o.PaymentServiceTransactionId.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.RawResponseCode.IsSet() {
		toSerialize["raw_response_code"] = o.RawResponseCode.Get()
	}
	if o.RawResponseDescription.IsSet() {
		toSerialize["raw_response_description"] = o.RawResponseDescription.Get()
	}
	if o.AvsResponseCode.IsSet() {
		toSerialize["avs_response_code"] = o.AvsResponseCode.Get()
	}
	if o.CvvResponseCode.IsSet() {
		toSerialize["cvv_response_code"] = o.CvvResponseCode.Get()
	}
	if o.PaymentMethodScheme.IsSet() {
		toSerialize["payment_method_scheme"] = o.PaymentMethodScheme.Get()
	}
	return toSerialize, nil
}

type NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext struct {
	value *PaymentConnectorResponseTransactionCaptureDeclinedEventContext
	isSet bool
}

func (v NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext) Get() *PaymentConnectorResponseTransactionCaptureDeclinedEventContext {
	return v.value
}

func (v *NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext) Set(val *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext(val *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) *NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext {
	return &NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext{value: val, isSet: true}
}

func (v NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


