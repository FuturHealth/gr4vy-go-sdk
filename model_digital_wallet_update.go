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

// checks if the DigitalWalletUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletUpdate{}

// DigitalWalletUpdate Request body to update a registered digital wallet's details.
type DigitalWalletUpdate struct {
	// The name of the merchant. This is used to update the value initially used to register with a digital wallet provider and this name is not displayed to the buyer.
	MerchantName *string `json:"merchant_name,omitempty"`
	// The list of domain names that a digital wallet can be used on. To use a digital wallet on a website, the domain of the site is required to be in this list.
	DomainNames []string `json:"domain_names,omitempty"`
	// The consumer facing name of the merchant.
	MerchantDisplayName NullableString `json:"merchant_display_name,omitempty"`
	// The country code where the merchant is registered.
	MerchantCountryCode NullableString `json:"merchant_country_code,omitempty"`
	// The main URL of the merchant.
	MerchantUrl *string `json:"merchant_url,omitempty"`
}

// NewDigitalWalletUpdate instantiates a new DigitalWalletUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletUpdate() *DigitalWalletUpdate {
	this := DigitalWalletUpdate{}
	return &this
}

// NewDigitalWalletUpdateWithDefaults instantiates a new DigitalWalletUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletUpdateWithDefaults() *DigitalWalletUpdate {
	this := DigitalWalletUpdate{}
	return &this
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *DigitalWalletUpdate) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletUpdate) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *DigitalWalletUpdate) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *DigitalWalletUpdate) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *DigitalWalletUpdate) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletUpdate) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *DigitalWalletUpdate) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *DigitalWalletUpdate) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetMerchantDisplayName returns the MerchantDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletUpdate) GetMerchantDisplayName() string {
	if o == nil || IsNil(o.MerchantDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantDisplayName.Get()
}

// GetMerchantDisplayNameOk returns a tuple with the MerchantDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletUpdate) GetMerchantDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantDisplayName.Get(), o.MerchantDisplayName.IsSet()
}

// HasMerchantDisplayName returns a boolean if a field has been set.
func (o *DigitalWalletUpdate) HasMerchantDisplayName() bool {
	if o != nil && o.MerchantDisplayName.IsSet() {
		return true
	}

	return false
}

// SetMerchantDisplayName gets a reference to the given NullableString and assigns it to the MerchantDisplayName field.
func (o *DigitalWalletUpdate) SetMerchantDisplayName(v string) {
	o.MerchantDisplayName.Set(&v)
}
// SetMerchantDisplayNameNil sets the value for MerchantDisplayName to be an explicit nil
func (o *DigitalWalletUpdate) SetMerchantDisplayNameNil() {
	o.MerchantDisplayName.Set(nil)
}

// UnsetMerchantDisplayName ensures that no value is present for MerchantDisplayName, not even an explicit nil
func (o *DigitalWalletUpdate) UnsetMerchantDisplayName() {
	o.MerchantDisplayName.Unset()
}

// GetMerchantCountryCode returns the MerchantCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletUpdate) GetMerchantCountryCode() string {
	if o == nil || IsNil(o.MerchantCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantCountryCode.Get()
}

// GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletUpdate) GetMerchantCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantCountryCode.Get(), o.MerchantCountryCode.IsSet()
}

// HasMerchantCountryCode returns a boolean if a field has been set.
func (o *DigitalWalletUpdate) HasMerchantCountryCode() bool {
	if o != nil && o.MerchantCountryCode.IsSet() {
		return true
	}

	return false
}

// SetMerchantCountryCode gets a reference to the given NullableString and assigns it to the MerchantCountryCode field.
func (o *DigitalWalletUpdate) SetMerchantCountryCode(v string) {
	o.MerchantCountryCode.Set(&v)
}
// SetMerchantCountryCodeNil sets the value for MerchantCountryCode to be an explicit nil
func (o *DigitalWalletUpdate) SetMerchantCountryCodeNil() {
	o.MerchantCountryCode.Set(nil)
}

// UnsetMerchantCountryCode ensures that no value is present for MerchantCountryCode, not even an explicit nil
func (o *DigitalWalletUpdate) UnsetMerchantCountryCode() {
	o.MerchantCountryCode.Unset()
}

// GetMerchantUrl returns the MerchantUrl field value if set, zero value otherwise.
func (o *DigitalWalletUpdate) GetMerchantUrl() string {
	if o == nil || IsNil(o.MerchantUrl) {
		var ret string
		return ret
	}
	return *o.MerchantUrl
}

// GetMerchantUrlOk returns a tuple with the MerchantUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletUpdate) GetMerchantUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantUrl) {
		return nil, false
	}
	return o.MerchantUrl, true
}

// HasMerchantUrl returns a boolean if a field has been set.
func (o *DigitalWalletUpdate) HasMerchantUrl() bool {
	if o != nil && !IsNil(o.MerchantUrl) {
		return true
	}

	return false
}

// SetMerchantUrl gets a reference to the given string and assigns it to the MerchantUrl field.
func (o *DigitalWalletUpdate) SetMerchantUrl(v string) {
	o.MerchantUrl = &v
}

func (o DigitalWalletUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !IsNil(o.DomainNames) {
		toSerialize["domain_names"] = o.DomainNames
	}
	if o.MerchantDisplayName.IsSet() {
		toSerialize["merchant_display_name"] = o.MerchantDisplayName.Get()
	}
	if o.MerchantCountryCode.IsSet() {
		toSerialize["merchant_country_code"] = o.MerchantCountryCode.Get()
	}
	if !IsNil(o.MerchantUrl) {
		toSerialize["merchant_url"] = o.MerchantUrl
	}
	return toSerialize, nil
}

type NullableDigitalWalletUpdate struct {
	value *DigitalWalletUpdate
	isSet bool
}

func (v NullableDigitalWalletUpdate) Get() *DigitalWalletUpdate {
	return v.value
}

func (v *NullableDigitalWalletUpdate) Set(val *DigitalWalletUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletUpdate(val *DigitalWalletUpdate) *NullableDigitalWalletUpdate {
	return &NullableDigitalWalletUpdate{value: val, isSet: true}
}

func (v NullableDigitalWalletUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


