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

// checks if the RoleAssignmentRequestRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAssignmentRequestRole{}

// RoleAssignmentRequestRole The role to associate with the role assignment.
type RoleAssignmentRequestRole struct {
	// The ID of the role to associate with the role assignment.
	Id string `json:"id"`
}

type _RoleAssignmentRequestRole RoleAssignmentRequestRole

// NewRoleAssignmentRequestRole instantiates a new RoleAssignmentRequestRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentRequestRole(id string) *RoleAssignmentRequestRole {
	this := RoleAssignmentRequestRole{}
	this.Id = id
	return &this
}

// NewRoleAssignmentRequestRoleWithDefaults instantiates a new RoleAssignmentRequestRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentRequestRoleWithDefaults() *RoleAssignmentRequestRole {
	this := RoleAssignmentRequestRole{}
	return &this
}

// GetId returns the Id field value
func (o *RoleAssignmentRequestRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentRequestRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleAssignmentRequestRole) SetId(v string) {
	o.Id = v
}

func (o RoleAssignmentRequestRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAssignmentRequestRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *RoleAssignmentRequestRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varRoleAssignmentRequestRole := _RoleAssignmentRequestRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoleAssignmentRequestRole)

	if err != nil {
		return err
	}

	*o = RoleAssignmentRequestRole(varRoleAssignmentRequestRole)

	return err
}

type NullableRoleAssignmentRequestRole struct {
	value *RoleAssignmentRequestRole
	isSet bool
}

func (v NullableRoleAssignmentRequestRole) Get() *RoleAssignmentRequestRole {
	return v.value
}

func (v *NullableRoleAssignmentRequestRole) Set(val *RoleAssignmentRequestRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentRequestRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentRequestRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentRequestRole(val *RoleAssignmentRequestRole) *NullableRoleAssignmentRequestRole {
	return &NullableRoleAssignmentRequestRole{value: val, isSet: true}
}

func (v NullableRoleAssignmentRequestRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentRequestRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


