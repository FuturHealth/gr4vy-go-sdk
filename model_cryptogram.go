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

// checks if the Cryptogram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cryptogram{}

// Cryptogram A network token cryptogram.
type Cryptogram struct {
	// The cryptogram of the network token.
	Cryptogram *string `json:"cryptogram,omitempty"`
}

// NewCryptogram instantiates a new Cryptogram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptogram() *Cryptogram {
	this := Cryptogram{}
	return &this
}

// NewCryptogramWithDefaults instantiates a new Cryptogram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptogramWithDefaults() *Cryptogram {
	this := Cryptogram{}
	return &this
}

// GetCryptogram returns the Cryptogram field value if set, zero value otherwise.
func (o *Cryptogram) GetCryptogram() string {
	if o == nil || IsNil(o.Cryptogram) {
		var ret string
		return ret
	}
	return *o.Cryptogram
}

// GetCryptogramOk returns a tuple with the Cryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cryptogram) GetCryptogramOk() (*string, bool) {
	if o == nil || IsNil(o.Cryptogram) {
		return nil, false
	}
	return o.Cryptogram, true
}

// HasCryptogram returns a boolean if a field has been set.
func (o *Cryptogram) HasCryptogram() bool {
	if o != nil && !IsNil(o.Cryptogram) {
		return true
	}

	return false
}

// SetCryptogram gets a reference to the given string and assigns it to the Cryptogram field.
func (o *Cryptogram) SetCryptogram(v string) {
	o.Cryptogram = &v
}

func (o Cryptogram) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cryptogram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cryptogram) {
		toSerialize["cryptogram"] = o.Cryptogram
	}
	return toSerialize, nil
}

type NullableCryptogram struct {
	value *Cryptogram
	isSet bool
}

func (v NullableCryptogram) Get() *Cryptogram {
	return v.value
}

func (v *NullableCryptogram) Set(val *Cryptogram) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptogram) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptogram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptogram(val *Cryptogram) *NullableCryptogram {
	return &NullableCryptogram{value: val, isSet: true}
}

func (v NullableCryptogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptogram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


