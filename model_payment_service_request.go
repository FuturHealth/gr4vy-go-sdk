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
	"bytes"
	"fmt"
)

// checks if the PaymentServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServiceRequest{}

// PaymentServiceRequest Request body for activating a payment service.
type PaymentServiceRequest struct {
	// The ID of the payment service to use.
	PaymentServiceDefinitionId string `json:"payment_service_definition_id"`
	// A custom name for the payment service. This will be shown in the Admin UI.
	DisplayName string `json:"display_name"`
	// A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card `secret_key` is required and so must be sent within this field.
	Fields []PaymentServiceRequestFieldsInner `json:"fields"`
	// The `reporting_fields` field should contain a list of key-value pairs. Each key-value pair represents a reporting field defined by the payment service. For example, when enabling settlement reporting for `nuvei-card`, the `ssh_username` field is required and must be included in `reporting_fields`.
	ReportingFields []PaymentServiceRequestReportingFieldsInner `json:"reporting_fields,omitempty"`
	// A list of countries that this payment service needs to support in ISO two-letter code format.
	AcceptedCountries []string `json:"accepted_countries"`
	// A list of currencies that this payment service needs to support in ISO 4217 three-letter code format.
	AcceptedCurrencies []string `json:"accepted_currencies"`
	// Defines if 3-D Secure is enabled for the service. This feature can only be enabled if the payment service definition supports the `three_d_secure_hosted` feature. This does not affect pass through 3-D Secure data.
	ThreeDSecureEnabled *bool `json:"three_d_secure_enabled,omitempty"`
	// Configuration for each supported card scheme.
	MerchantProfile NullableMerchantProfile `json:"merchant_profile,omitempty"`
	// Defines if this service is currently active or not.
	Active *bool `json:"active,omitempty"`
	// Defines if the service works as an open-loop service. This feature can only be enabled if the PSP is set up to accept previous scheme transaction IDs.  If this value is not provided or is set to `null`, it will be set to the value of `open_loop` in the payment service definition.  If `open_loop_toggle` is `false` in the payment service definition, `open_loop` should either not be provided or set to `null`, or it will fail with a validation error.
	OpenLoop NullableBool `json:"open_loop,omitempty"`
	// Defines if tokenization is enabled for the service. This feature can only be enabled if the payment service is NOT set as `open_loop` and the PSP is set up to tokenize.
	PaymentMethodTokenizationEnabled NullableBool `json:"payment_method_tokenization_enabled,omitempty"`
	// Defines if network tokens are enabled for the service. This feature can only be enabled if the payment service is set as `open_loop` and the PSP is set up to accept network tokens.  If this value is not provided or is set to `null`, it will be set to the value of `network_tokens_default` in the payment service definition.  If `network_tokens_toggle` is `false` in the payment service definition, `network_tokens_enabled` should either not be provided or set to `null`, or it will fail with a validation error.
	NetworkTokensEnabled NullableBool `json:"network_tokens_enabled,omitempty"`
	// Defines if settlement reporting is enabled for the service. This feature can only be enabled if the payment service definition supports the `settlement_reporting` feature.
	SettlementReportingEnabled *bool `json:"settlement_reporting_enabled,omitempty"`
}

type _PaymentServiceRequest PaymentServiceRequest

// NewPaymentServiceRequest instantiates a new PaymentServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceRequest(paymentServiceDefinitionId string, displayName string, fields []PaymentServiceRequestFieldsInner, acceptedCountries []string, acceptedCurrencies []string) *PaymentServiceRequest {
	this := PaymentServiceRequest{}
	this.PaymentServiceDefinitionId = paymentServiceDefinitionId
	this.DisplayName = displayName
	this.Fields = fields
	this.AcceptedCountries = acceptedCountries
	this.AcceptedCurrencies = acceptedCurrencies
	var threeDSecureEnabled bool = false
	this.ThreeDSecureEnabled = &threeDSecureEnabled
	var active bool = true
	this.Active = &active
	var settlementReportingEnabled bool = false
	this.SettlementReportingEnabled = &settlementReportingEnabled
	return &this
}

// NewPaymentServiceRequestWithDefaults instantiates a new PaymentServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceRequestWithDefaults() *PaymentServiceRequest {
	this := PaymentServiceRequest{}
	var threeDSecureEnabled bool = false
	this.ThreeDSecureEnabled = &threeDSecureEnabled
	var active bool = true
	this.Active = &active
	var settlementReportingEnabled bool = false
	this.SettlementReportingEnabled = &settlementReportingEnabled
	return &this
}

// GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field value
func (o *PaymentServiceRequest) GetPaymentServiceDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentServiceDefinitionId
}

// GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetPaymentServiceDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentServiceDefinitionId, true
}

// SetPaymentServiceDefinitionId sets field value
func (o *PaymentServiceRequest) SetPaymentServiceDefinitionId(v string) {
	o.PaymentServiceDefinitionId = v
}

// GetDisplayName returns the DisplayName field value
func (o *PaymentServiceRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PaymentServiceRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFields returns the Fields field value
func (o *PaymentServiceRequest) GetFields() []PaymentServiceRequestFieldsInner {
	if o == nil {
		var ret []PaymentServiceRequestFieldsInner
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetFieldsOk() ([]PaymentServiceRequestFieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *PaymentServiceRequest) SetFields(v []PaymentServiceRequestFieldsInner) {
	o.Fields = v
}

// GetReportingFields returns the ReportingFields field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetReportingFields() []PaymentServiceRequestReportingFieldsInner {
	if o == nil || IsNil(o.ReportingFields) {
		var ret []PaymentServiceRequestReportingFieldsInner
		return ret
	}
	return o.ReportingFields
}

// GetReportingFieldsOk returns a tuple with the ReportingFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetReportingFieldsOk() ([]PaymentServiceRequestReportingFieldsInner, bool) {
	if o == nil || IsNil(o.ReportingFields) {
		return nil, false
	}
	return o.ReportingFields, true
}

// HasReportingFields returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasReportingFields() bool {
	if o != nil && !IsNil(o.ReportingFields) {
		return true
	}

	return false
}

// SetReportingFields gets a reference to the given []PaymentServiceRequestReportingFieldsInner and assigns it to the ReportingFields field.
func (o *PaymentServiceRequest) SetReportingFields(v []PaymentServiceRequestReportingFieldsInner) {
	o.ReportingFields = v
}

// GetAcceptedCountries returns the AcceptedCountries field value
func (o *PaymentServiceRequest) GetAcceptedCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcceptedCountries
}

// GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetAcceptedCountriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedCountries, true
}

// SetAcceptedCountries sets field value
func (o *PaymentServiceRequest) SetAcceptedCountries(v []string) {
	o.AcceptedCountries = v
}

// GetAcceptedCurrencies returns the AcceptedCurrencies field value
func (o *PaymentServiceRequest) GetAcceptedCurrencies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcceptedCurrencies
}

// GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetAcceptedCurrenciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedCurrencies, true
}

// SetAcceptedCurrencies sets field value
func (o *PaymentServiceRequest) SetAcceptedCurrencies(v []string) {
	o.AcceptedCurrencies = v
}

// GetThreeDSecureEnabled returns the ThreeDSecureEnabled field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetThreeDSecureEnabled() bool {
	if o == nil || IsNil(o.ThreeDSecureEnabled) {
		var ret bool
		return ret
	}
	return *o.ThreeDSecureEnabled
}

// GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetThreeDSecureEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ThreeDSecureEnabled) {
		return nil, false
	}
	return o.ThreeDSecureEnabled, true
}

// HasThreeDSecureEnabled returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasThreeDSecureEnabled() bool {
	if o != nil && !IsNil(o.ThreeDSecureEnabled) {
		return true
	}

	return false
}

// SetThreeDSecureEnabled gets a reference to the given bool and assigns it to the ThreeDSecureEnabled field.
func (o *PaymentServiceRequest) SetThreeDSecureEnabled(v bool) {
	o.ThreeDSecureEnabled = &v
}

// GetMerchantProfile returns the MerchantProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetMerchantProfile() MerchantProfile {
	if o == nil || IsNil(o.MerchantProfile.Get()) {
		var ret MerchantProfile
		return ret
	}
	return *o.MerchantProfile.Get()
}

// GetMerchantProfileOk returns a tuple with the MerchantProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetMerchantProfileOk() (*MerchantProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantProfile.Get(), o.MerchantProfile.IsSet()
}

// HasMerchantProfile returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasMerchantProfile() bool {
	if o != nil && o.MerchantProfile.IsSet() {
		return true
	}

	return false
}

// SetMerchantProfile gets a reference to the given NullableMerchantProfile and assigns it to the MerchantProfile field.
func (o *PaymentServiceRequest) SetMerchantProfile(v MerchantProfile) {
	o.MerchantProfile.Set(&v)
}
// SetMerchantProfileNil sets the value for MerchantProfile to be an explicit nil
func (o *PaymentServiceRequest) SetMerchantProfileNil() {
	o.MerchantProfile.Set(nil)
}

// UnsetMerchantProfile ensures that no value is present for MerchantProfile, not even an explicit nil
func (o *PaymentServiceRequest) UnsetMerchantProfile() {
	o.MerchantProfile.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PaymentServiceRequest) SetActive(v bool) {
	o.Active = &v
}

// GetOpenLoop returns the OpenLoop field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetOpenLoop() bool {
	if o == nil || IsNil(o.OpenLoop.Get()) {
		var ret bool
		return ret
	}
	return *o.OpenLoop.Get()
}

// GetOpenLoopOk returns a tuple with the OpenLoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetOpenLoopOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenLoop.Get(), o.OpenLoop.IsSet()
}

// HasOpenLoop returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasOpenLoop() bool {
	if o != nil && o.OpenLoop.IsSet() {
		return true
	}

	return false
}

// SetOpenLoop gets a reference to the given NullableBool and assigns it to the OpenLoop field.
func (o *PaymentServiceRequest) SetOpenLoop(v bool) {
	o.OpenLoop.Set(&v)
}
// SetOpenLoopNil sets the value for OpenLoop to be an explicit nil
func (o *PaymentServiceRequest) SetOpenLoopNil() {
	o.OpenLoop.Set(nil)
}

// UnsetOpenLoop ensures that no value is present for OpenLoop, not even an explicit nil
func (o *PaymentServiceRequest) UnsetOpenLoop() {
	o.OpenLoop.Unset()
}

// GetPaymentMethodTokenizationEnabled returns the PaymentMethodTokenizationEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetPaymentMethodTokenizationEnabled() bool {
	if o == nil || IsNil(o.PaymentMethodTokenizationEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.PaymentMethodTokenizationEnabled.Get()
}

// GetPaymentMethodTokenizationEnabledOk returns a tuple with the PaymentMethodTokenizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetPaymentMethodTokenizationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodTokenizationEnabled.Get(), o.PaymentMethodTokenizationEnabled.IsSet()
}

// HasPaymentMethodTokenizationEnabled returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasPaymentMethodTokenizationEnabled() bool {
	if o != nil && o.PaymentMethodTokenizationEnabled.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodTokenizationEnabled gets a reference to the given NullableBool and assigns it to the PaymentMethodTokenizationEnabled field.
func (o *PaymentServiceRequest) SetPaymentMethodTokenizationEnabled(v bool) {
	o.PaymentMethodTokenizationEnabled.Set(&v)
}
// SetPaymentMethodTokenizationEnabledNil sets the value for PaymentMethodTokenizationEnabled to be an explicit nil
func (o *PaymentServiceRequest) SetPaymentMethodTokenizationEnabledNil() {
	o.PaymentMethodTokenizationEnabled.Set(nil)
}

// UnsetPaymentMethodTokenizationEnabled ensures that no value is present for PaymentMethodTokenizationEnabled, not even an explicit nil
func (o *PaymentServiceRequest) UnsetPaymentMethodTokenizationEnabled() {
	o.PaymentMethodTokenizationEnabled.Unset()
}

// GetNetworkTokensEnabled returns the NetworkTokensEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetNetworkTokensEnabled() bool {
	if o == nil || IsNil(o.NetworkTokensEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.NetworkTokensEnabled.Get()
}

// GetNetworkTokensEnabledOk returns a tuple with the NetworkTokensEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetNetworkTokensEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkTokensEnabled.Get(), o.NetworkTokensEnabled.IsSet()
}

// HasNetworkTokensEnabled returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasNetworkTokensEnabled() bool {
	if o != nil && o.NetworkTokensEnabled.IsSet() {
		return true
	}

	return false
}

// SetNetworkTokensEnabled gets a reference to the given NullableBool and assigns it to the NetworkTokensEnabled field.
func (o *PaymentServiceRequest) SetNetworkTokensEnabled(v bool) {
	o.NetworkTokensEnabled.Set(&v)
}
// SetNetworkTokensEnabledNil sets the value for NetworkTokensEnabled to be an explicit nil
func (o *PaymentServiceRequest) SetNetworkTokensEnabledNil() {
	o.NetworkTokensEnabled.Set(nil)
}

// UnsetNetworkTokensEnabled ensures that no value is present for NetworkTokensEnabled, not even an explicit nil
func (o *PaymentServiceRequest) UnsetNetworkTokensEnabled() {
	o.NetworkTokensEnabled.Unset()
}

// GetSettlementReportingEnabled returns the SettlementReportingEnabled field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetSettlementReportingEnabled() bool {
	if o == nil || IsNil(o.SettlementReportingEnabled) {
		var ret bool
		return ret
	}
	return *o.SettlementReportingEnabled
}

// GetSettlementReportingEnabledOk returns a tuple with the SettlementReportingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetSettlementReportingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SettlementReportingEnabled) {
		return nil, false
	}
	return o.SettlementReportingEnabled, true
}

// HasSettlementReportingEnabled returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasSettlementReportingEnabled() bool {
	if o != nil && !IsNil(o.SettlementReportingEnabled) {
		return true
	}

	return false
}

// SetSettlementReportingEnabled gets a reference to the given bool and assigns it to the SettlementReportingEnabled field.
func (o *PaymentServiceRequest) SetSettlementReportingEnabled(v bool) {
	o.SettlementReportingEnabled = &v
}

func (o PaymentServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payment_service_definition_id"] = o.PaymentServiceDefinitionId
	toSerialize["display_name"] = o.DisplayName
	toSerialize["fields"] = o.Fields
	if !IsNil(o.ReportingFields) {
		toSerialize["reporting_fields"] = o.ReportingFields
	}
	toSerialize["accepted_countries"] = o.AcceptedCountries
	toSerialize["accepted_currencies"] = o.AcceptedCurrencies
	if !IsNil(o.ThreeDSecureEnabled) {
		toSerialize["three_d_secure_enabled"] = o.ThreeDSecureEnabled
	}
	if o.MerchantProfile.IsSet() {
		toSerialize["merchant_profile"] = o.MerchantProfile.Get()
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if o.OpenLoop.IsSet() {
		toSerialize["open_loop"] = o.OpenLoop.Get()
	}
	if o.PaymentMethodTokenizationEnabled.IsSet() {
		toSerialize["payment_method_tokenization_enabled"] = o.PaymentMethodTokenizationEnabled.Get()
	}
	if o.NetworkTokensEnabled.IsSet() {
		toSerialize["network_tokens_enabled"] = o.NetworkTokensEnabled.Get()
	}
	if !IsNil(o.SettlementReportingEnabled) {
		toSerialize["settlement_reporting_enabled"] = o.SettlementReportingEnabled
	}
	return toSerialize, nil
}

func (o *PaymentServiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payment_service_definition_id",
		"display_name",
		"fields",
		"accepted_countries",
		"accepted_currencies",
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

	varPaymentServiceRequest := _PaymentServiceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentServiceRequest)

	if err != nil {
		return err
	}

	*o = PaymentServiceRequest(varPaymentServiceRequest)

	return err
}

type NullablePaymentServiceRequest struct {
	value *PaymentServiceRequest
	isSet bool
}

func (v NullablePaymentServiceRequest) Get() *PaymentServiceRequest {
	return v.value
}

func (v *NullablePaymentServiceRequest) Set(val *PaymentServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceRequest(val *PaymentServiceRequest) *NullablePaymentServiceRequest {
	return &NullablePaymentServiceRequest{value: val, isSet: true}
}

func (v NullablePaymentServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


