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

// checks if the BINLookupRequestContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BINLookupRequestContext{}

// BINLookupRequestContext BIN lookup request context.
type BINLookupRequestContext struct {
	// The response body received from the `url`.
	Response NullableString `json:"response,omitempty"`
	// The response status code received from the `url`.
	ResponseStatusCode *int32 `json:"response_status_code,omitempty"`
	// Whether the BIN lookup was successful or not.
	Success *bool `json:"success,omitempty"`
	// The value used to lookup BIN details.
	Bin NullableString `json:"bin,omitempty"`
	// The instrument type used to lookup BIN details.
	Instrument NullableString `json:"instrument,omitempty"`
	// The type of card, i.e. credit or debit, from the lookup response.
	Type NullableString `json:"type,omitempty"`
	// The card scheme result from the lookup response.
	Scheme NullableString `json:"scheme,omitempty"`
	// The card additional schemes from the lookup response.
	AdditionalSchemes []string `json:"additional_schemes,omitempty"`
	// The card country code result from the lookup response.
	CountryCode NullableString `json:"country_code,omitempty"`
	// Whether the issuing bank supports network tokenization for this card.
	SupportsNetworkTokens NullableBool `json:"supports_network_tokens,omitempty"`
}

// NewBINLookupRequestContext instantiates a new BINLookupRequestContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBINLookupRequestContext() *BINLookupRequestContext {
	this := BINLookupRequestContext{}
	return &this
}

// NewBINLookupRequestContextWithDefaults instantiates a new BINLookupRequestContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBINLookupRequestContextWithDefaults() *BINLookupRequestContext {
	this := BINLookupRequestContext{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetResponse() string {
	if o == nil || IsNil(o.Response.Get()) {
		var ret string
		return ret
	}
	return *o.Response.Get()
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Response.Get(), o.Response.IsSet()
}

// HasResponse returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasResponse() bool {
	if o != nil && o.Response.IsSet() {
		return true
	}

	return false
}

// SetResponse gets a reference to the given NullableString and assigns it to the Response field.
func (o *BINLookupRequestContext) SetResponse(v string) {
	o.Response.Set(&v)
}
// SetResponseNil sets the value for Response to be an explicit nil
func (o *BINLookupRequestContext) SetResponseNil() {
	o.Response.Set(nil)
}

// UnsetResponse ensures that no value is present for Response, not even an explicit nil
func (o *BINLookupRequestContext) UnsetResponse() {
	o.Response.Unset()
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *BINLookupRequestContext) GetResponseStatusCode() int32 {
	if o == nil || IsNil(o.ResponseStatusCode) {
		var ret int32
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BINLookupRequestContext) GetResponseStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponseStatusCode) {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasResponseStatusCode() bool {
	if o != nil && !IsNil(o.ResponseStatusCode) {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given int32 and assigns it to the ResponseStatusCode field.
func (o *BINLookupRequestContext) SetResponseStatusCode(v int32) {
	o.ResponseStatusCode = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BINLookupRequestContext) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BINLookupRequestContext) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *BINLookupRequestContext) SetSuccess(v bool) {
	o.Success = &v
}

// GetBin returns the Bin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetBin() string {
	if o == nil || IsNil(o.Bin.Get()) {
		var ret string
		return ret
	}
	return *o.Bin.Get()
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetBinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bin.Get(), o.Bin.IsSet()
}

// HasBin returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasBin() bool {
	if o != nil && o.Bin.IsSet() {
		return true
	}

	return false
}

// SetBin gets a reference to the given NullableString and assigns it to the Bin field.
func (o *BINLookupRequestContext) SetBin(v string) {
	o.Bin.Set(&v)
}
// SetBinNil sets the value for Bin to be an explicit nil
func (o *BINLookupRequestContext) SetBinNil() {
	o.Bin.Set(nil)
}

// UnsetBin ensures that no value is present for Bin, not even an explicit nil
func (o *BINLookupRequestContext) UnsetBin() {
	o.Bin.Unset()
}

// GetInstrument returns the Instrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetInstrument() string {
	if o == nil || IsNil(o.Instrument.Get()) {
		var ret string
		return ret
	}
	return *o.Instrument.Get()
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetInstrumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instrument.Get(), o.Instrument.IsSet()
}

// HasInstrument returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasInstrument() bool {
	if o != nil && o.Instrument.IsSet() {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given NullableString and assigns it to the Instrument field.
func (o *BINLookupRequestContext) SetInstrument(v string) {
	o.Instrument.Set(&v)
}
// SetInstrumentNil sets the value for Instrument to be an explicit nil
func (o *BINLookupRequestContext) SetInstrumentNil() {
	o.Instrument.Set(nil)
}

// UnsetInstrument ensures that no value is present for Instrument, not even an explicit nil
func (o *BINLookupRequestContext) UnsetInstrument() {
	o.Instrument.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *BINLookupRequestContext) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *BINLookupRequestContext) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *BINLookupRequestContext) UnsetType() {
	o.Type.Unset()
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetScheme() string {
	if o == nil || IsNil(o.Scheme.Get()) {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *BINLookupRequestContext) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *BINLookupRequestContext) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *BINLookupRequestContext) UnsetScheme() {
	o.Scheme.Unset()
}

// GetAdditionalSchemes returns the AdditionalSchemes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetAdditionalSchemes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AdditionalSchemes
}

// GetAdditionalSchemesOk returns a tuple with the AdditionalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetAdditionalSchemesOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalSchemes) {
		return nil, false
	}
	return o.AdditionalSchemes, true
}

// HasAdditionalSchemes returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasAdditionalSchemes() bool {
	if o != nil && !IsNil(o.AdditionalSchemes) {
		return true
	}

	return false
}

// SetAdditionalSchemes gets a reference to the given []string and assigns it to the AdditionalSchemes field.
func (o *BINLookupRequestContext) SetAdditionalSchemes(v []string) {
	o.AdditionalSchemes = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *BINLookupRequestContext) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *BINLookupRequestContext) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *BINLookupRequestContext) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetSupportsNetworkTokens returns the SupportsNetworkTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BINLookupRequestContext) GetSupportsNetworkTokens() bool {
	if o == nil || IsNil(o.SupportsNetworkTokens.Get()) {
		var ret bool
		return ret
	}
	return *o.SupportsNetworkTokens.Get()
}

// GetSupportsNetworkTokensOk returns a tuple with the SupportsNetworkTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BINLookupRequestContext) GetSupportsNetworkTokensOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportsNetworkTokens.Get(), o.SupportsNetworkTokens.IsSet()
}

// HasSupportsNetworkTokens returns a boolean if a field has been set.
func (o *BINLookupRequestContext) HasSupportsNetworkTokens() bool {
	if o != nil && o.SupportsNetworkTokens.IsSet() {
		return true
	}

	return false
}

// SetSupportsNetworkTokens gets a reference to the given NullableBool and assigns it to the SupportsNetworkTokens field.
func (o *BINLookupRequestContext) SetSupportsNetworkTokens(v bool) {
	o.SupportsNetworkTokens.Set(&v)
}
// SetSupportsNetworkTokensNil sets the value for SupportsNetworkTokens to be an explicit nil
func (o *BINLookupRequestContext) SetSupportsNetworkTokensNil() {
	o.SupportsNetworkTokens.Set(nil)
}

// UnsetSupportsNetworkTokens ensures that no value is present for SupportsNetworkTokens, not even an explicit nil
func (o *BINLookupRequestContext) UnsetSupportsNetworkTokens() {
	o.SupportsNetworkTokens.Unset()
}

func (o BINLookupRequestContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BINLookupRequestContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Response.IsSet() {
		toSerialize["response"] = o.Response.Get()
	}
	if !IsNil(o.ResponseStatusCode) {
		toSerialize["response_status_code"] = o.ResponseStatusCode
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if o.Bin.IsSet() {
		toSerialize["bin"] = o.Bin.Get()
	}
	if o.Instrument.IsSet() {
		toSerialize["instrument"] = o.Instrument.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.AdditionalSchemes != nil {
		toSerialize["additional_schemes"] = o.AdditionalSchemes
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.SupportsNetworkTokens.IsSet() {
		toSerialize["supports_network_tokens"] = o.SupportsNetworkTokens.Get()
	}
	return toSerialize, nil
}

type NullableBINLookupRequestContext struct {
	value *BINLookupRequestContext
	isSet bool
}

func (v NullableBINLookupRequestContext) Get() *BINLookupRequestContext {
	return v.value
}

func (v *NullableBINLookupRequestContext) Set(val *BINLookupRequestContext) {
	v.value = val
	v.isSet = true
}

func (v NullableBINLookupRequestContext) IsSet() bool {
	return v.isSet
}

func (v *NullableBINLookupRequestContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBINLookupRequestContext(val *BINLookupRequestContext) *NullableBINLookupRequestContext {
	return &NullableBINLookupRequestContext{value: val, isSet: true}
}

func (v NullableBINLookupRequestContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBINLookupRequestContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


