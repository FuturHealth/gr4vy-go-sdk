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

// checks if the IssueCryptogramRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueCryptogramRequest{}

// IssueCryptogramRequest Request body for issue a cryptogram for a network token.
type IssueCryptogramRequest struct {
	// Defines if the request is merchant initiated or not.
	MerchantInitiated bool `json:"merchant_initiated"`
}

type _IssueCryptogramRequest IssueCryptogramRequest

// NewIssueCryptogramRequest instantiates a new IssueCryptogramRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueCryptogramRequest(merchantInitiated bool) *IssueCryptogramRequest {
	this := IssueCryptogramRequest{}
	this.MerchantInitiated = merchantInitiated
	return &this
}

// NewIssueCryptogramRequestWithDefaults instantiates a new IssueCryptogramRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueCryptogramRequestWithDefaults() *IssueCryptogramRequest {
	this := IssueCryptogramRequest{}
	var merchantInitiated bool = false
	this.MerchantInitiated = merchantInitiated
	return &this
}

// GetMerchantInitiated returns the MerchantInitiated field value
func (o *IssueCryptogramRequest) GetMerchantInitiated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MerchantInitiated
}

// GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field value
// and a boolean to check if the value has been set.
func (o *IssueCryptogramRequest) GetMerchantInitiatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantInitiated, true
}

// SetMerchantInitiated sets field value
func (o *IssueCryptogramRequest) SetMerchantInitiated(v bool) {
	o.MerchantInitiated = v
}

func (o IssueCryptogramRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueCryptogramRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchant_initiated"] = o.MerchantInitiated
	return toSerialize, nil
}

func (o *IssueCryptogramRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchant_initiated",
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

	varIssueCryptogramRequest := _IssueCryptogramRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIssueCryptogramRequest)

	if err != nil {
		return err
	}

	*o = IssueCryptogramRequest(varIssueCryptogramRequest)

	return err
}

type NullableIssueCryptogramRequest struct {
	value *IssueCryptogramRequest
	isSet bool
}

func (v NullableIssueCryptogramRequest) Get() *IssueCryptogramRequest {
	return v.value
}

func (v *NullableIssueCryptogramRequest) Set(val *IssueCryptogramRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueCryptogramRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueCryptogramRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueCryptogramRequest(val *IssueCryptogramRequest) *NullableIssueCryptogramRequest {
	return &NullableIssueCryptogramRequest{value: val, isSet: true}
}

func (v NullableIssueCryptogramRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueCryptogramRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


