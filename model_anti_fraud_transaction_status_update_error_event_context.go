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

// checks if the AntiFraudTransactionStatusUpdateErrorEventContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntiFraudTransactionStatusUpdateErrorEventContext{}

// AntiFraudTransactionStatusUpdateErrorEventContext Additional context for this event.
type AntiFraudTransactionStatusUpdateErrorEventContext struct {
	// The unique ID of the anti-fraud service used.
	AntiFraudServiceId *string `json:"anti_fraud_service_id,omitempty"`
	// The name of the anti-fraud service used.
	AntiFraudServiceName *string `json:"anti_fraud_service_name,omitempty"`
	// The anti-fraud service definition used.
	AntiFraudServiceDefinitionId *string `json:"anti_fraud_service_definition_id,omitempty"`
	// The reason we could not get the anti-fraud decision.
	Reason *string `json:"reason,omitempty"`
	// The HTTP body sent to fetch a decision.
	Request NullableString `json:"request,omitempty"`
	// The HTTP body received from the anti-fraud provider.
	Response NullableString `json:"response,omitempty"`
	// The HTTP response status code from the anti-fraud provider.
	ResponseStatusCode NullableFloat32 `json:"response_status_code,omitempty"`
}

// NewAntiFraudTransactionStatusUpdateErrorEventContext instantiates a new AntiFraudTransactionStatusUpdateErrorEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntiFraudTransactionStatusUpdateErrorEventContext() *AntiFraudTransactionStatusUpdateErrorEventContext {
	this := AntiFraudTransactionStatusUpdateErrorEventContext{}
	return &this
}

// NewAntiFraudTransactionStatusUpdateErrorEventContextWithDefaults instantiates a new AntiFraudTransactionStatusUpdateErrorEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntiFraudTransactionStatusUpdateErrorEventContextWithDefaults() *AntiFraudTransactionStatusUpdateErrorEventContext {
	this := AntiFraudTransactionStatusUpdateErrorEventContext{}
	return &this
}

// GetAntiFraudServiceId returns the AntiFraudServiceId field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceId() string {
	if o == nil || IsNil(o.AntiFraudServiceId) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceId
}

// GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceId) {
		return nil, false
	}
	return o.AntiFraudServiceId, true
}

// HasAntiFraudServiceId returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasAntiFraudServiceId() bool {
	if o != nil && !IsNil(o.AntiFraudServiceId) {
		return true
	}

	return false
}

// SetAntiFraudServiceId gets a reference to the given string and assigns it to the AntiFraudServiceId field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetAntiFraudServiceId(v string) {
	o.AntiFraudServiceId = &v
}

// GetAntiFraudServiceName returns the AntiFraudServiceName field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceName() string {
	if o == nil || IsNil(o.AntiFraudServiceName) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceName
}

// GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceName) {
		return nil, false
	}
	return o.AntiFraudServiceName, true
}

// HasAntiFraudServiceName returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasAntiFraudServiceName() bool {
	if o != nil && !IsNil(o.AntiFraudServiceName) {
		return true
	}

	return false
}

// SetAntiFraudServiceName gets a reference to the given string and assigns it to the AntiFraudServiceName field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetAntiFraudServiceName(v string) {
	o.AntiFraudServiceName = &v
}

// GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceDefinitionId() string {
	if o == nil || IsNil(o.AntiFraudServiceDefinitionId) {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceDefinitionId
}

// GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFraudServiceDefinitionId) {
		return nil, false
	}
	return o.AntiFraudServiceDefinitionId, true
}

// HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasAntiFraudServiceDefinitionId() bool {
	if o != nil && !IsNil(o.AntiFraudServiceDefinitionId) {
		return true
	}

	return false
}

// SetAntiFraudServiceDefinitionId gets a reference to the given string and assigns it to the AntiFraudServiceDefinitionId field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetAntiFraudServiceDefinitionId(v string) {
	o.AntiFraudServiceDefinitionId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetReason(v string) {
	o.Reason = &v
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetRequest() string {
	if o == nil || IsNil(o.Request.Get()) {
		var ret string
		return ret
	}
	return *o.Request.Get()
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Request.Get(), o.Request.IsSet()
}

// HasRequest returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasRequest() bool {
	if o != nil && o.Request.IsSet() {
		return true
	}

	return false
}

// SetRequest gets a reference to the given NullableString and assigns it to the Request field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetRequest(v string) {
	o.Request.Set(&v)
}
// SetRequestNil sets the value for Request to be an explicit nil
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetRequestNil() {
	o.Request.Set(nil)
}

// UnsetRequest ensures that no value is present for Request, not even an explicit nil
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) UnsetRequest() {
	o.Request.Unset()
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponse() string {
	if o == nil || IsNil(o.Response.Get()) {
		var ret string
		return ret
	}
	return *o.Response.Get()
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Response.Get(), o.Response.IsSet()
}

// HasResponse returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasResponse() bool {
	if o != nil && o.Response.IsSet() {
		return true
	}

	return false
}

// SetResponse gets a reference to the given NullableString and assigns it to the Response field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponse(v string) {
	o.Response.Set(&v)
}
// SetResponseNil sets the value for Response to be an explicit nil
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponseNil() {
	o.Response.Set(nil)
}

// UnsetResponse ensures that no value is present for Response, not even an explicit nil
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) UnsetResponse() {
	o.Response.Unset()
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponseStatusCode() float32 {
	if o == nil || IsNil(o.ResponseStatusCode.Get()) {
		var ret float32
		return ret
	}
	return *o.ResponseStatusCode.Get()
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponseStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseStatusCode.Get(), o.ResponseStatusCode.IsSet()
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasResponseStatusCode() bool {
	if o != nil && o.ResponseStatusCode.IsSet() {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given NullableFloat32 and assigns it to the ResponseStatusCode field.
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponseStatusCode(v float32) {
	o.ResponseStatusCode.Set(&v)
}
// SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponseStatusCodeNil() {
	o.ResponseStatusCode.Set(nil)
}

// UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
func (o *AntiFraudTransactionStatusUpdateErrorEventContext) UnsetResponseStatusCode() {
	o.ResponseStatusCode.Unset()
}

func (o AntiFraudTransactionStatusUpdateErrorEventContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntiFraudTransactionStatusUpdateErrorEventContext) ToMap() (map[string]interface{}, error) {
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
	if o.Request.IsSet() {
		toSerialize["request"] = o.Request.Get()
	}
	if o.Response.IsSet() {
		toSerialize["response"] = o.Response.Get()
	}
	if o.ResponseStatusCode.IsSet() {
		toSerialize["response_status_code"] = o.ResponseStatusCode.Get()
	}
	return toSerialize, nil
}

type NullableAntiFraudTransactionStatusUpdateErrorEventContext struct {
	value *AntiFraudTransactionStatusUpdateErrorEventContext
	isSet bool
}

func (v NullableAntiFraudTransactionStatusUpdateErrorEventContext) Get() *AntiFraudTransactionStatusUpdateErrorEventContext {
	return v.value
}

func (v *NullableAntiFraudTransactionStatusUpdateErrorEventContext) Set(val *AntiFraudTransactionStatusUpdateErrorEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAntiFraudTransactionStatusUpdateErrorEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAntiFraudTransactionStatusUpdateErrorEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntiFraudTransactionStatusUpdateErrorEventContext(val *AntiFraudTransactionStatusUpdateErrorEventContext) *NullableAntiFraudTransactionStatusUpdateErrorEventContext {
	return &NullableAntiFraudTransactionStatusUpdateErrorEventContext{value: val, isSet: true}
}

func (v NullableAntiFraudTransactionStatusUpdateErrorEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntiFraudTransactionStatusUpdateErrorEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


