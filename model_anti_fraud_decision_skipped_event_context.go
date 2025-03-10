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

// checks if the AntiFraudDecisionSkippedEventContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntiFraudDecisionSkippedEventContext{}

// AntiFraudDecisionSkippedEventContext Additional context for this event.
type AntiFraudDecisionSkippedEventContext struct {
	// The unique ID of the anti-fraud service used.
	AntiFraudServiceId *string `json:"anti_fraud_service_id,omitempty"`
	// The name of the anti-fraud service used.
	AntiFraudServiceName *string `json:"anti_fraud_service_name,omitempty"`
	// The anti-fraud service definition used.
	AntiFraudServiceDefinitionId *string `json:"anti_fraud_service_definition_id,omitempty"`
	// The reason we could not get the anti-fraud decision.
	Reason *string `json:"reason,omitempty"`
}

// NewAntiFraudDecisionSkippedEventContext instantiates a new AntiFraudDecisionSkippedEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntiFraudDecisionSkippedEventContext() *AntiFraudDecisionSkippedEventContext {
	this := AntiFraudDecisionSkippedEventContext{}
	return &this
}

// NewAntiFraudDecisionSkippedEventContextWithDefaults instantiates a new AntiFraudDecisionSkippedEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntiFraudDecisionSkippedEventContextWithDefaults() *AntiFraudDecisionSkippedEventContext {
	this := AntiFraudDecisionSkippedEventContext{}
	return &this
}

// GetAntiFraudServiceId returns the AntiFraudServiceId field value if set, zero value otherwise.
func (o *AntiFraudDecisionSkippedEventContext) GetAntiFraudServiceId() string {
	if o == nil || IsNil(o.AntiFraudServiceId) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceId
}

// GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionSkippedEventContext) GetAntiFraudServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceId) {
		return nil, false
	}
	return o.AntiFraudServiceId, true
}

// HasAntiFraudServiceId returns a boolean if a field has been set.
func (o *AntiFraudDecisionSkippedEventContext) HasAntiFraudServiceId() bool {
	if o != nil && !IsNil(o.AntiFraudServiceId) {
		return true
	}

	return false
}

// SetAntiFraudServiceId gets a reference to the given string and assigns it to the AntiFraudServiceId field.
func (o *AntiFraudDecisionSkippedEventContext) SetAntiFraudServiceId(v string) {
	o.AntiFraudServiceId = &v
}

// GetAntiFraudServiceName returns the AntiFraudServiceName field value if set, zero value otherwise.
func (o *AntiFraudDecisionSkippedEventContext) GetAntiFraudServiceName() string {
	if o == nil || IsNil(o.AntiFraudServiceName) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceName
}

// GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionSkippedEventContext) GetAntiFraudServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceName) {
		return nil, false
	}
	return o.AntiFraudServiceName, true
}

// HasAntiFraudServiceName returns a boolean if a field has been set.
func (o *AntiFraudDecisionSkippedEventContext) HasAntiFraudServiceName() bool {
	if o != nil && !IsNil(o.AntiFraudServiceName) {
		return true
	}

	return false
}

// SetAntiFraudServiceName gets a reference to the given string and assigns it to the AntiFraudServiceName field.
func (o *AntiFraudDecisionSkippedEventContext) SetAntiFraudServiceName(v string) {
	o.AntiFraudServiceName = &v
}

// GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field value if set, zero value otherwise.
func (o *AntiFraudDecisionSkippedEventContext) GetAntiFraudServiceDefinitionId() string {
	if o == nil || IsNil(o.AntiFraudServiceDefinitionId) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceDefinitionId
}

// GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionSkippedEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceDefinitionId) {
		return nil, false
	}
	return o.AntiFraudServiceDefinitionId, true
}

// HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.
func (o *AntiFraudDecisionSkippedEventContext) HasAntiFraudServiceDefinitionId() bool {
	if o != nil && !IsNil(o.AntiFraudServiceDefinitionId) {
		return true
	}

	return false
}

// SetAntiFraudServiceDefinitionId gets a reference to the given string and assigns it to the AntiFraudServiceDefinitionId field.
func (o *AntiFraudDecisionSkippedEventContext) SetAntiFraudServiceDefinitionId(v string) {
	o.AntiFraudServiceDefinitionId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AntiFraudDecisionSkippedEventContext) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionSkippedEventContext) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AntiFraudDecisionSkippedEventContext) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AntiFraudDecisionSkippedEventContext) SetReason(v string) {
	o.Reason = &v
}

func (o AntiFraudDecisionSkippedEventContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntiFraudDecisionSkippedEventContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AntiFraudServiceId) {
		toSerialize["anti_fraud_service_id"] = o.AntiFraudServiceId
	}
	if !IsNil(o.AntiFraudServiceName) {
		toSerialize["anti_fraud_service_name"] = o.AntiFraudServiceName
	}
	if !IsNil(o.AntiFraudServiceDefinitionId) {
		toSerialize["anti_fraud_service_definition_id"] = o.AntiFraudServiceDefinitionId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableAntiFraudDecisionSkippedEventContext struct {
	value *AntiFraudDecisionSkippedEventContext
	isSet bool
}

func (v NullableAntiFraudDecisionSkippedEventContext) Get() *AntiFraudDecisionSkippedEventContext {
	return v.value
}

func (v *NullableAntiFraudDecisionSkippedEventContext) Set(val *AntiFraudDecisionSkippedEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAntiFraudDecisionSkippedEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAntiFraudDecisionSkippedEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntiFraudDecisionSkippedEventContext(val *AntiFraudDecisionSkippedEventContext) *NullableAntiFraudDecisionSkippedEventContext {
	return &NullableAntiFraudDecisionSkippedEventContext{value: val, isSet: true}
}

func (v NullableAntiFraudDecisionSkippedEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntiFraudDecisionSkippedEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


