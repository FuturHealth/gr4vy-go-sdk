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

// checks if the AntiFraudDecisionErrorEventContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntiFraudDecisionErrorEventContext{}

// AntiFraudDecisionErrorEventContext Additional context for this event.
type AntiFraudDecisionErrorEventContext struct {
	// The unique ID of the anti-fraud service used.
	AntiFraudServiceId *string `json:"anti_fraud_service_id,omitempty"`
	// The name of the anti-fraud service used.
	AntiFraudServiceName *string `json:"anti_fraud_service_name,omitempty"`
	// The anti-fraud service definition used.
	AntiFraudServiceDefinitionId *string `json:"anti_fraud_service_definition_id,omitempty"`
	// The HTTP response status code from the anti-fraud provider, if we received any.
	StatusCode *float32 `json:"status_code,omitempty"`
	// The reason we could not get the anti-fraud decision.
	Reason *string `json:"reason,omitempty"`
}

// NewAntiFraudDecisionErrorEventContext instantiates a new AntiFraudDecisionErrorEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntiFraudDecisionErrorEventContext() *AntiFraudDecisionErrorEventContext {
	this := AntiFraudDecisionErrorEventContext{}
	return &this
}

// NewAntiFraudDecisionErrorEventContextWithDefaults instantiates a new AntiFraudDecisionErrorEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntiFraudDecisionErrorEventContextWithDefaults() *AntiFraudDecisionErrorEventContext {
	this := AntiFraudDecisionErrorEventContext{}
	return &this
}

// GetAntiFraudServiceId returns the AntiFraudServiceId field value if set, zero value otherwise.
func (o *AntiFraudDecisionErrorEventContext) GetAntiFraudServiceId() string {
	if o == nil || IsNil(o.AntiFraudServiceId) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceId
}

// GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionErrorEventContext) GetAntiFraudServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceId) {
		return nil, false
	}
	return o.AntiFraudServiceId, true
}

// HasAntiFraudServiceId returns a boolean if a field has been set.
func (o *AntiFraudDecisionErrorEventContext) HasAntiFraudServiceId() bool {
	if o != nil && !IsNil(o.AntiFraudServiceId) {
		return true
	}

	return false
}

// SetAntiFraudServiceId gets a reference to the given string and assigns it to the AntiFraudServiceId field.
func (o *AntiFraudDecisionErrorEventContext) SetAntiFraudServiceId(v string) {
	o.AntiFraudServiceId = &v
}

// GetAntiFraudServiceName returns the AntiFraudServiceName field value if set, zero value otherwise.
func (o *AntiFraudDecisionErrorEventContext) GetAntiFraudServiceName() string {
	if o == nil || IsNil(o.AntiFraudServiceName) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceName
}

// GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionErrorEventContext) GetAntiFraudServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceName) {
		return nil, false
	}
	return o.AntiFraudServiceName, true
}

// HasAntiFraudServiceName returns a boolean if a field has been set.
func (o *AntiFraudDecisionErrorEventContext) HasAntiFraudServiceName() bool {
	if o != nil && !IsNil(o.AntiFraudServiceName) {
		return true
	}

	return false
}

// SetAntiFraudServiceName gets a reference to the given string and assigns it to the AntiFraudServiceName field.
func (o *AntiFraudDecisionErrorEventContext) SetAntiFraudServiceName(v string) {
	o.AntiFraudServiceName = &v
}

// GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field value if set, zero value otherwise.
func (o *AntiFraudDecisionErrorEventContext) GetAntiFraudServiceDefinitionId() string {
	if o == nil || IsNil(o.AntiFraudServiceDefinitionId) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceDefinitionId
}

// GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionErrorEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceDefinitionId) {
		return nil, false
	}
	return o.AntiFraudServiceDefinitionId, true
}

// HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.
func (o *AntiFraudDecisionErrorEventContext) HasAntiFraudServiceDefinitionId() bool {
	if o != nil && !IsNil(o.AntiFraudServiceDefinitionId) {
		return true
	}

	return false
}

// SetAntiFraudServiceDefinitionId gets a reference to the given string and assigns it to the AntiFraudServiceDefinitionId field.
func (o *AntiFraudDecisionErrorEventContext) SetAntiFraudServiceDefinitionId(v string) {
	o.AntiFraudServiceDefinitionId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *AntiFraudDecisionErrorEventContext) GetStatusCode() float32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret float32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionErrorEventContext) GetStatusCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *AntiFraudDecisionErrorEventContext) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given float32 and assigns it to the StatusCode field.
func (o *AntiFraudDecisionErrorEventContext) SetStatusCode(v float32) {
	o.StatusCode = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AntiFraudDecisionErrorEventContext) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudDecisionErrorEventContext) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AntiFraudDecisionErrorEventContext) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AntiFraudDecisionErrorEventContext) SetReason(v string) {
	o.Reason = &v
}

func (o AntiFraudDecisionErrorEventContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntiFraudDecisionErrorEventContext) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableAntiFraudDecisionErrorEventContext struct {
	value *AntiFraudDecisionErrorEventContext
	isSet bool
}

func (v NullableAntiFraudDecisionErrorEventContext) Get() *AntiFraudDecisionErrorEventContext {
	return v.value
}

func (v *NullableAntiFraudDecisionErrorEventContext) Set(val *AntiFraudDecisionErrorEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAntiFraudDecisionErrorEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAntiFraudDecisionErrorEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntiFraudDecisionErrorEventContext(val *AntiFraudDecisionErrorEventContext) *NullableAntiFraudDecisionErrorEventContext {
	return &NullableAntiFraudDecisionErrorEventContext{value: val, isSet: true}
}

func (v NullableAntiFraudDecisionErrorEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntiFraudDecisionErrorEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


