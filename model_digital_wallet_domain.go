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

// checks if the DigitalWalletDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletDomain{}

// DigitalWalletDomain Domain name for a digital wallet.
type DigitalWalletDomain struct {
	// The domain name that a digital wallet can be used on. To use a digital wallet on a website, the domain of the site is required to be in this list.
	DomainName string `json:"domain_name"`
}

type _DigitalWalletDomain DigitalWalletDomain

// NewDigitalWalletDomain instantiates a new DigitalWalletDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletDomain(domainName string) *DigitalWalletDomain {
	this := DigitalWalletDomain{}
	this.DomainName = domainName
	return &this
}

// NewDigitalWalletDomainWithDefaults instantiates a new DigitalWalletDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletDomainWithDefaults() *DigitalWalletDomain {
	this := DigitalWalletDomain{}
	return &this
}

// GetDomainName returns the DomainName field value
func (o *DigitalWalletDomain) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletDomain) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *DigitalWalletDomain) SetDomainName(v string) {
	o.DomainName = v
}

func (o DigitalWalletDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain_name"] = o.DomainName
	return toSerialize, nil
}

func (o *DigitalWalletDomain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain_name",
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

	varDigitalWalletDomain := _DigitalWalletDomain{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDigitalWalletDomain)

	if err != nil {
		return err
	}

	*o = DigitalWalletDomain(varDigitalWalletDomain)

	return err
}

type NullableDigitalWalletDomain struct {
	value *DigitalWalletDomain
	isSet bool
}

func (v NullableDigitalWalletDomain) Get() *DigitalWalletDomain {
	return v.value
}

func (v *NullableDigitalWalletDomain) Set(val *DigitalWalletDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletDomain(val *DigitalWalletDomain) *NullableDigitalWalletDomain {
	return &NullableDigitalWalletDomain{value: val, isSet: true}
}

func (v NullableDigitalWalletDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


