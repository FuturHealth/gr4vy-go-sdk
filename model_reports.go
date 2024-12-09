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

// checks if the Reports type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reports{}

// Reports A list of reports.
type Reports struct {
	// A list of reports.
	Items []Report `json:"items,omitempty"`
	// The limit applied to request. This represents the number of items that are at maximum returned by this request.
	Limit *int32 `json:"limit,omitempty"`
	// The cursor that represents the next page of results. Use the `cursor` query parameter to fetch this page of items.
	NextCursor NullableString `json:"next_cursor,omitempty"`
	// The cursor that represents the next page of results. Use the `cursor` query parameter to fetch this page of items.
	PreviousCursor NullableString `json:"previous_cursor,omitempty"`
}

// NewReports instantiates a new Reports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReports() *Reports {
	this := Reports{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewReportsWithDefaults instantiates a new Reports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsWithDefaults() *Reports {
	this := Reports{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Reports) GetItems() []Report {
	if o == nil || IsNil(o.Items) {
		var ret []Report
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reports) GetItemsOk() ([]Report, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Reports) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Report and assigns it to the Items field.
func (o *Reports) SetItems(v []Report) {
	o.Items = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Reports) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reports) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Reports) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *Reports) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Reports) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor.Get()) {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Reports) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *Reports) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *Reports) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *Reports) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *Reports) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Reports) GetPreviousCursor() string {
	if o == nil || IsNil(o.PreviousCursor.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousCursor.Get()
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Reports) GetPreviousCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousCursor.Get(), o.PreviousCursor.IsSet()
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *Reports) HasPreviousCursor() bool {
	if o != nil && o.PreviousCursor.IsSet() {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given NullableString and assigns it to the PreviousCursor field.
func (o *Reports) SetPreviousCursor(v string) {
	o.PreviousCursor.Set(&v)
}
// SetPreviousCursorNil sets the value for PreviousCursor to be an explicit nil
func (o *Reports) SetPreviousCursorNil() {
	o.PreviousCursor.Set(nil)
}

// UnsetPreviousCursor ensures that no value is present for PreviousCursor, not even an explicit nil
func (o *Reports) UnsetPreviousCursor() {
	o.PreviousCursor.Unset()
}

func (o Reports) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reports) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if o.PreviousCursor.IsSet() {
		toSerialize["previous_cursor"] = o.PreviousCursor.Get()
	}
	return toSerialize, nil
}

type NullableReports struct {
	value *Reports
	isSet bool
}

func (v NullableReports) Get() *Reports {
	return v.value
}

func (v *NullableReports) Set(val *Reports) {
	v.value = val
	v.isSet = true
}

func (v NullableReports) IsSet() bool {
	return v.isSet
}

func (v *NullableReports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReports(val *Reports) *NullableReports {
	return &NullableReports{value: val, isSet: true}
}

func (v NullableReports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


