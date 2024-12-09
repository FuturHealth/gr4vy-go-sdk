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

// checks if the ConnectionOptionsAdyenCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionOptionsAdyenCard{}

// ConnectionOptionsAdyenCard Additional options to be passed through to Adyen when processing card transactions.
type ConnectionOptionsAdyenCard struct {
	// A key-value object representing additional data to be passed to Adyen.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// Enabled Adyen's auto-rescue feature.
	AutoRescue *bool `json:"autoRescue,omitempty"`
	// Defines the number of days to auto-retry a payment for if `autoRescue` is enabled.
	MaxDaysToRescue NullableInt32 `json:"maxDaysToRescue,omitempty"`
	// Defines the Adyen auto-rescue test scenario to invoke.
	AutoRescueScenario NullableString `json:"autoRescueScenario,omitempty"`
}

// NewConnectionOptionsAdyenCard instantiates a new ConnectionOptionsAdyenCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsAdyenCard() *ConnectionOptionsAdyenCard {
	this := ConnectionOptionsAdyenCard{}
	var autoRescue bool = false
	this.AutoRescue = &autoRescue
	return &this
}

// NewConnectionOptionsAdyenCardWithDefaults instantiates a new ConnectionOptionsAdyenCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsAdyenCardWithDefaults() *ConnectionOptionsAdyenCard {
	this := ConnectionOptionsAdyenCard{}
	var autoRescue bool = false
	this.AutoRescue = &autoRescue
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ConnectionOptionsAdyenCard) GetAdditionalData() map[string]string {
	if o == nil || IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsAdyenCard) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenCard) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ConnectionOptionsAdyenCard) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAutoRescue returns the AutoRescue field value if set, zero value otherwise.
func (o *ConnectionOptionsAdyenCard) GetAutoRescue() bool {
	if o == nil || IsNil(o.AutoRescue) {
		var ret bool
		return ret
	}
	return *o.AutoRescue
}

// GetAutoRescueOk returns a tuple with the AutoRescue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsAdyenCard) GetAutoRescueOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRescue) {
		return nil, false
	}
	return o.AutoRescue, true
}

// HasAutoRescue returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenCard) HasAutoRescue() bool {
	if o != nil && !IsNil(o.AutoRescue) {
		return true
	}

	return false
}

// SetAutoRescue gets a reference to the given bool and assigns it to the AutoRescue field.
func (o *ConnectionOptionsAdyenCard) SetAutoRescue(v bool) {
	o.AutoRescue = &v
}

// GetMaxDaysToRescue returns the MaxDaysToRescue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsAdyenCard) GetMaxDaysToRescue() int32 {
	if o == nil || IsNil(o.MaxDaysToRescue.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxDaysToRescue.Get()
}

// GetMaxDaysToRescueOk returns a tuple with the MaxDaysToRescue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsAdyenCard) GetMaxDaysToRescueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDaysToRescue.Get(), o.MaxDaysToRescue.IsSet()
}

// HasMaxDaysToRescue returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenCard) HasMaxDaysToRescue() bool {
	if o != nil && o.MaxDaysToRescue.IsSet() {
		return true
	}

	return false
}

// SetMaxDaysToRescue gets a reference to the given NullableInt32 and assigns it to the MaxDaysToRescue field.
func (o *ConnectionOptionsAdyenCard) SetMaxDaysToRescue(v int32) {
	o.MaxDaysToRescue.Set(&v)
}
// SetMaxDaysToRescueNil sets the value for MaxDaysToRescue to be an explicit nil
func (o *ConnectionOptionsAdyenCard) SetMaxDaysToRescueNil() {
	o.MaxDaysToRescue.Set(nil)
}

// UnsetMaxDaysToRescue ensures that no value is present for MaxDaysToRescue, not even an explicit nil
func (o *ConnectionOptionsAdyenCard) UnsetMaxDaysToRescue() {
	o.MaxDaysToRescue.Unset()
}

// GetAutoRescueScenario returns the AutoRescueScenario field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsAdyenCard) GetAutoRescueScenario() string {
	if o == nil || IsNil(o.AutoRescueScenario.Get()) {
		var ret string
		return ret
	}
	return *o.AutoRescueScenario.Get()
}

// GetAutoRescueScenarioOk returns a tuple with the AutoRescueScenario field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsAdyenCard) GetAutoRescueScenarioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRescueScenario.Get(), o.AutoRescueScenario.IsSet()
}

// HasAutoRescueScenario returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenCard) HasAutoRescueScenario() bool {
	if o != nil && o.AutoRescueScenario.IsSet() {
		return true
	}

	return false
}

// SetAutoRescueScenario gets a reference to the given NullableString and assigns it to the AutoRescueScenario field.
func (o *ConnectionOptionsAdyenCard) SetAutoRescueScenario(v string) {
	o.AutoRescueScenario.Set(&v)
}
// SetAutoRescueScenarioNil sets the value for AutoRescueScenario to be an explicit nil
func (o *ConnectionOptionsAdyenCard) SetAutoRescueScenarioNil() {
	o.AutoRescueScenario.Set(nil)
}

// UnsetAutoRescueScenario ensures that no value is present for AutoRescueScenario, not even an explicit nil
func (o *ConnectionOptionsAdyenCard) UnsetAutoRescueScenario() {
	o.AutoRescueScenario.Unset()
}

func (o ConnectionOptionsAdyenCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionOptionsAdyenCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !IsNil(o.AutoRescue) {
		toSerialize["autoRescue"] = o.AutoRescue
	}
	if o.MaxDaysToRescue.IsSet() {
		toSerialize["maxDaysToRescue"] = o.MaxDaysToRescue.Get()
	}
	if o.AutoRescueScenario.IsSet() {
		toSerialize["autoRescueScenario"] = o.AutoRescueScenario.Get()
	}
	return toSerialize, nil
}

type NullableConnectionOptionsAdyenCard struct {
	value *ConnectionOptionsAdyenCard
	isSet bool
}

func (v NullableConnectionOptionsAdyenCard) Get() *ConnectionOptionsAdyenCard {
	return v.value
}

func (v *NullableConnectionOptionsAdyenCard) Set(val *ConnectionOptionsAdyenCard) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsAdyenCard) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsAdyenCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsAdyenCard(val *ConnectionOptionsAdyenCard) *NullableConnectionOptionsAdyenCard {
	return &NullableConnectionOptionsAdyenCard{value: val, isSet: true}
}

func (v NullableConnectionOptionsAdyenCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsAdyenCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


