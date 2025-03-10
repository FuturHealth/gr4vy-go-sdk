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

// checks if the RoleAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAssignmentRequest{}

// RoleAssignmentRequest A request to create a role assignment.
type RoleAssignmentRequest struct {
	Role RoleAssignmentRequestRole `json:"role"`
	Assignee RoleAssignmentRequestAssignee `json:"assignee"`
}

type _RoleAssignmentRequest RoleAssignmentRequest

// NewRoleAssignmentRequest instantiates a new RoleAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentRequest(role RoleAssignmentRequestRole, assignee RoleAssignmentRequestAssignee) *RoleAssignmentRequest {
	this := RoleAssignmentRequest{}
	this.Role = role
	this.Assignee = assignee
	return &this
}

// NewRoleAssignmentRequestWithDefaults instantiates a new RoleAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentRequestWithDefaults() *RoleAssignmentRequest {
	this := RoleAssignmentRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *RoleAssignmentRequest) GetRole() RoleAssignmentRequestRole {
	if o == nil {
		var ret RoleAssignmentRequestRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRequest) GetRoleOk() (*RoleAssignmentRequestRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleAssignmentRequest) SetRole(v RoleAssignmentRequestRole) {
	o.Role = v
}

// GetAssignee returns the Assignee field value
func (o *RoleAssignmentRequest) GetAssignee() RoleAssignmentRequestAssignee {
	if o == nil {
		var ret RoleAssignmentRequestAssignee
		return ret
	}

	return o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRequest) GetAssigneeOk() (*RoleAssignmentRequestAssignee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignee, true
}

// SetAssignee sets field value
func (o *RoleAssignmentRequest) SetAssignee(v RoleAssignmentRequestAssignee) {
	o.Assignee = v
}

func (o RoleAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["assignee"] = o.Assignee
	return toSerialize, nil
}

func (o *RoleAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
		"assignee",
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

	varRoleAssignmentRequest := _RoleAssignmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoleAssignmentRequest)

	if err != nil {
		return err
	}

	*o = RoleAssignmentRequest(varRoleAssignmentRequest)

	return err
}

type NullableRoleAssignmentRequest struct {
	value *RoleAssignmentRequest
	isSet bool
}

func (v NullableRoleAssignmentRequest) Get() *RoleAssignmentRequest {
	return v.value
}

func (v *NullableRoleAssignmentRequest) Set(val *RoleAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentRequest(val *RoleAssignmentRequest) *NullableRoleAssignmentRequest {
	return &NullableRoleAssignmentRequest{value: val, isSet: true}
}

func (v NullableRoleAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


