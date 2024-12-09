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

// checks if the PaymentServiceSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServiceSnapshot{}

// PaymentServiceSnapshot An active, configured payment service.
type PaymentServiceSnapshot struct {
	// The type of this resource.
	Type *string `json:"type,omitempty"`
	// The ID of this payment service.
	Id *string `json:"id,omitempty"`
	// The custom name set for this service.
	DisplayName *string `json:"display_name,omitempty"`
	// The payment method that this services handles.
	Method *string `json:"method,omitempty"`
	// The ID of the payment service definition used to create this service. 
	PaymentServiceDefinitionId *string `json:"payment_service_definition_id,omitempty"`
}

// NewPaymentServiceSnapshot instantiates a new PaymentServiceSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceSnapshot() *PaymentServiceSnapshot {
	this := PaymentServiceSnapshot{}
	return &this
}

// NewPaymentServiceSnapshotWithDefaults instantiates a new PaymentServiceSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceSnapshotWithDefaults() *PaymentServiceSnapshot {
	this := PaymentServiceSnapshot{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentServiceSnapshot) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSnapshot) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentServiceSnapshot) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentServiceSnapshot) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentServiceSnapshot) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSnapshot) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentServiceSnapshot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentServiceSnapshot) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PaymentServiceSnapshot) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSnapshot) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PaymentServiceSnapshot) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PaymentServiceSnapshot) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentServiceSnapshot) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSnapshot) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentServiceSnapshot) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentServiceSnapshot) SetMethod(v string) {
	o.Method = &v
}

// GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field value if set, zero value otherwise.
func (o *PaymentServiceSnapshot) GetPaymentServiceDefinitionId() string {
	if o == nil || IsNil(o.PaymentServiceDefinitionId) {
		var ret string
		return ret
	}
	return *o.PaymentServiceDefinitionId
}

// GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSnapshot) GetPaymentServiceDefinitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceDefinitionId) {
		return nil, false
	}
	return o.PaymentServiceDefinitionId, true
}

// HasPaymentServiceDefinitionId returns a boolean if a field has been set.
func (o *PaymentServiceSnapshot) HasPaymentServiceDefinitionId() bool {
	if o != nil && !IsNil(o.PaymentServiceDefinitionId) {
		return true
	}

	return false
}

// SetPaymentServiceDefinitionId gets a reference to the given string and assigns it to the PaymentServiceDefinitionId field.
func (o *PaymentServiceSnapshot) SetPaymentServiceDefinitionId(v string) {
	o.PaymentServiceDefinitionId = &v
}

func (o PaymentServiceSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServiceSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.PaymentServiceDefinitionId) {
		toSerialize["payment_service_definition_id"] = o.PaymentServiceDefinitionId
	}
	return toSerialize, nil
}

type NullablePaymentServiceSnapshot struct {
	value *PaymentServiceSnapshot
	isSet bool
}

func (v NullablePaymentServiceSnapshot) Get() *PaymentServiceSnapshot {
	return v.value
}

func (v *NullablePaymentServiceSnapshot) Set(val *PaymentServiceSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceSnapshot(val *PaymentServiceSnapshot) *NullablePaymentServiceSnapshot {
	return &NullablePaymentServiceSnapshot{value: val, isSet: true}
}

func (v NullablePaymentServiceSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


