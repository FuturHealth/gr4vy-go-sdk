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

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address An address for the buyer.
type Address struct {
	// The city for the address.
	City NullableString `json:"city,omitempty"`
	// The country for the address in ISO 3166 format.
	Country NullableString `json:"country,omitempty"`
	// The postal code or zip code for the address.
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The state, county, or province for the address.
	State NullableString `json:"state,omitempty"`
	// The code of state, county, or province for the address in ISO 3166-2 format.
	StateCode NullableString `json:"state_code,omitempty"`
	// The house number or name for the address. Not all payment services use this field but some do.
	HouseNumberOrName NullableString `json:"house_number_or_name,omitempty"`
	// The first line of the address.
	Line1 NullableString `json:"line1,omitempty"`
	// The second line of the address.
	Line2 NullableString `json:"line2,omitempty"`
	// The optional name of the company or organisation to add to the address.
	Organization NullableString `json:"organization,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *Address) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *Address) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *Address) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *Address) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *Address) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *Address) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *Address) UnsetCountry() {
	o.Country.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *Address) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *Address) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *Address) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *Address) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *Address) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *Address) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *Address) UnsetState() {
	o.State.Unset()
}

// GetStateCode returns the StateCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStateCode() string {
	if o == nil || IsNil(o.StateCode.Get()) {
		var ret string
		return ret
	}
	return *o.StateCode.Get()
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateCode.Get(), o.StateCode.IsSet()
}

// HasStateCode returns a boolean if a field has been set.
func (o *Address) HasStateCode() bool {
	if o != nil && o.StateCode.IsSet() {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given NullableString and assigns it to the StateCode field.
func (o *Address) SetStateCode(v string) {
	o.StateCode.Set(&v)
}
// SetStateCodeNil sets the value for StateCode to be an explicit nil
func (o *Address) SetStateCodeNil() {
	o.StateCode.Set(nil)
}

// UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
func (o *Address) UnsetStateCode() {
	o.StateCode.Unset()
}

// GetHouseNumberOrName returns the HouseNumberOrName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetHouseNumberOrName() string {
	if o == nil || IsNil(o.HouseNumberOrName.Get()) {
		var ret string
		return ret
	}
	return *o.HouseNumberOrName.Get()
}

// GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetHouseNumberOrNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HouseNumberOrName.Get(), o.HouseNumberOrName.IsSet()
}

// HasHouseNumberOrName returns a boolean if a field has been set.
func (o *Address) HasHouseNumberOrName() bool {
	if o != nil && o.HouseNumberOrName.IsSet() {
		return true
	}

	return false
}

// SetHouseNumberOrName gets a reference to the given NullableString and assigns it to the HouseNumberOrName field.
func (o *Address) SetHouseNumberOrName(v string) {
	o.HouseNumberOrName.Set(&v)
}
// SetHouseNumberOrNameNil sets the value for HouseNumberOrName to be an explicit nil
func (o *Address) SetHouseNumberOrNameNil() {
	o.HouseNumberOrName.Set(nil)
}

// UnsetHouseNumberOrName ensures that no value is present for HouseNumberOrName, not even an explicit nil
func (o *Address) UnsetHouseNumberOrName() {
	o.HouseNumberOrName.Unset()
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetLine1() string {
	if o == nil || IsNil(o.Line1.Get()) {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *Address) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *Address) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *Address) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *Address) UnsetLine1() {
	o.Line1.Unset()
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetLine2() string {
	if o == nil || IsNil(o.Line2.Get()) {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *Address) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *Address) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *Address) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *Address) UnsetLine2() {
	o.Line2.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetOrganization() string {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *Address) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *Address) SetOrganization(v string) {
	o.Organization.Set(&v)
}
// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *Address) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *Address) UnsetOrganization() {
	o.Organization.Unset()
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.StateCode.IsSet() {
		toSerialize["state_code"] = o.StateCode.Get()
	}
	if o.HouseNumberOrName.IsSet() {
		toSerialize["house_number_or_name"] = o.HouseNumberOrName.Get()
	}
	if o.Line1.IsSet() {
		toSerialize["line1"] = o.Line1.Get()
	}
	if o.Line2.IsSet() {
		toSerialize["line2"] = o.Line2.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


