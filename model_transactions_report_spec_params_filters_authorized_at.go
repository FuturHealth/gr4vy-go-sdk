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

// checks if the TransactionsReportSpecParamsFiltersAuthorizedAt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionsReportSpecParamsFiltersAuthorizedAt{}

// TransactionsReportSpecParamsFiltersAuthorizedAt The time window for when the transactions were authorized to filter by.
type TransactionsReportSpecParamsFiltersAuthorizedAt struct {
	// The starting boundary for when the transactions were authorized as either a date-time or date-time placeholder value.
	Start *string `json:"start,omitempty"`
	// The ending boundary for when the transactions were authorized as either a date-time or date-time placeholder value.
	End *string `json:"end,omitempty"`
}

// NewTransactionsReportSpecParamsFiltersAuthorizedAt instantiates a new TransactionsReportSpecParamsFiltersAuthorizedAt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsReportSpecParamsFiltersAuthorizedAt() *TransactionsReportSpecParamsFiltersAuthorizedAt {
	this := TransactionsReportSpecParamsFiltersAuthorizedAt{}
	return &this
}

// NewTransactionsReportSpecParamsFiltersAuthorizedAtWithDefaults instantiates a new TransactionsReportSpecParamsFiltersAuthorizedAt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsReportSpecParamsFiltersAuthorizedAtWithDefaults() *TransactionsReportSpecParamsFiltersAuthorizedAt {
	this := TransactionsReportSpecParamsFiltersAuthorizedAt{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *TransactionsReportSpecParamsFiltersAuthorizedAt) SetEnd(v string) {
	o.End = &v
}

func (o TransactionsReportSpecParamsFiltersAuthorizedAt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionsReportSpecParamsFiltersAuthorizedAt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableTransactionsReportSpecParamsFiltersAuthorizedAt struct {
	value *TransactionsReportSpecParamsFiltersAuthorizedAt
	isSet bool
}

func (v NullableTransactionsReportSpecParamsFiltersAuthorizedAt) Get() *TransactionsReportSpecParamsFiltersAuthorizedAt {
	return v.value
}

func (v *NullableTransactionsReportSpecParamsFiltersAuthorizedAt) Set(val *TransactionsReportSpecParamsFiltersAuthorizedAt) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsReportSpecParamsFiltersAuthorizedAt) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsReportSpecParamsFiltersAuthorizedAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsReportSpecParamsFiltersAuthorizedAt(val *TransactionsReportSpecParamsFiltersAuthorizedAt) *NullableTransactionsReportSpecParamsFiltersAuthorizedAt {
	return &NullableTransactionsReportSpecParamsFiltersAuthorizedAt{value: val, isSet: true}
}

func (v NullableTransactionsReportSpecParamsFiltersAuthorizedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsReportSpecParamsFiltersAuthorizedAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


