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

// checks if the TransactionRetriesReportSpecParamsFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRetriesReportSpecParamsFilters{}

// TransactionRetriesReportSpecParamsFilters The filters for the report.
type TransactionRetriesReportSpecParamsFilters struct {
	CreatedAt *TransactionRetriesReportSpecParamsFiltersCreatedAt `json:"created_at,omitempty"`
}

// NewTransactionRetriesReportSpecParamsFilters instantiates a new TransactionRetriesReportSpecParamsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRetriesReportSpecParamsFilters() *TransactionRetriesReportSpecParamsFilters {
	this := TransactionRetriesReportSpecParamsFilters{}
	return &this
}

// NewTransactionRetriesReportSpecParamsFiltersWithDefaults instantiates a new TransactionRetriesReportSpecParamsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRetriesReportSpecParamsFiltersWithDefaults() *TransactionRetriesReportSpecParamsFilters {
	this := TransactionRetriesReportSpecParamsFilters{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransactionRetriesReportSpecParamsFilters) GetCreatedAt() TransactionRetriesReportSpecParamsFiltersCreatedAt {
	if o == nil || IsNil(o.CreatedAt) {
		var ret TransactionRetriesReportSpecParamsFiltersCreatedAt
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRetriesReportSpecParamsFilters) GetCreatedAtOk() (*TransactionRetriesReportSpecParamsFiltersCreatedAt, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransactionRetriesReportSpecParamsFilters) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given TransactionRetriesReportSpecParamsFiltersCreatedAt and assigns it to the CreatedAt field.
func (o *TransactionRetriesReportSpecParamsFilters) SetCreatedAt(v TransactionRetriesReportSpecParamsFiltersCreatedAt) {
	o.CreatedAt = &v
}

func (o TransactionRetriesReportSpecParamsFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRetriesReportSpecParamsFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableTransactionRetriesReportSpecParamsFilters struct {
	value *TransactionRetriesReportSpecParamsFilters
	isSet bool
}

func (v NullableTransactionRetriesReportSpecParamsFilters) Get() *TransactionRetriesReportSpecParamsFilters {
	return v.value
}

func (v *NullableTransactionRetriesReportSpecParamsFilters) Set(val *TransactionRetriesReportSpecParamsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRetriesReportSpecParamsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRetriesReportSpecParamsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRetriesReportSpecParamsFilters(val *TransactionRetriesReportSpecParamsFilters) *NullableTransactionRetriesReportSpecParamsFilters {
	return &NullableTransactionRetriesReportSpecParamsFilters{value: val, isSet: true}
}

func (v NullableTransactionRetriesReportSpecParamsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRetriesReportSpecParamsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


