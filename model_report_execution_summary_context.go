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
	"time"
)

// checks if the ReportExecutionSummaryContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportExecutionSummaryContext{}

// ReportExecutionSummaryContext Contains the context values used to compute the value of date-time placeholders such as `month_start` and `month_end` if present in the report's specification. Date-time placeholders are dynamic timestamps that change with every report execution.
type ReportExecutionSummaryContext struct {
	// The date and time used by the system as a reference point to compute date-time placeholders.
	ReferenceTimestamp *time.Time `json:"reference_timestamp,omitempty"`
	// The time zone used to compute date-time placeholders.
	ReferenceTimezone *string `json:"reference_timezone,omitempty"`
}

// NewReportExecutionSummaryContext instantiates a new ReportExecutionSummaryContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportExecutionSummaryContext() *ReportExecutionSummaryContext {
	this := ReportExecutionSummaryContext{}
	return &this
}

// NewReportExecutionSummaryContextWithDefaults instantiates a new ReportExecutionSummaryContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportExecutionSummaryContextWithDefaults() *ReportExecutionSummaryContext {
	this := ReportExecutionSummaryContext{}
	return &this
}

// GetReferenceTimestamp returns the ReferenceTimestamp field value if set, zero value otherwise.
func (o *ReportExecutionSummaryContext) GetReferenceTimestamp() time.Time {
	if o == nil || IsNil(o.ReferenceTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ReferenceTimestamp
}

// GetReferenceTimestampOk returns a tuple with the ReferenceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecutionSummaryContext) GetReferenceTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReferenceTimestamp) {
		return nil, false
	}
	return o.ReferenceTimestamp, true
}

// HasReferenceTimestamp returns a boolean if a field has been set.
func (o *ReportExecutionSummaryContext) HasReferenceTimestamp() bool {
	if o != nil && !IsNil(o.ReferenceTimestamp) {
		return true
	}

	return false
}

// SetReferenceTimestamp gets a reference to the given time.Time and assigns it to the ReferenceTimestamp field.
func (o *ReportExecutionSummaryContext) SetReferenceTimestamp(v time.Time) {
	o.ReferenceTimestamp = &v
}

// GetReferenceTimezone returns the ReferenceTimezone field value if set, zero value otherwise.
func (o *ReportExecutionSummaryContext) GetReferenceTimezone() string {
	if o == nil || IsNil(o.ReferenceTimezone) {
		var ret string
		return ret
	}
	return *o.ReferenceTimezone
}

// GetReferenceTimezoneOk returns a tuple with the ReferenceTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecutionSummaryContext) GetReferenceTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceTimezone) {
		return nil, false
	}
	return o.ReferenceTimezone, true
}

// HasReferenceTimezone returns a boolean if a field has been set.
func (o *ReportExecutionSummaryContext) HasReferenceTimezone() bool {
	if o != nil && !IsNil(o.ReferenceTimezone) {
		return true
	}

	return false
}

// SetReferenceTimezone gets a reference to the given string and assigns it to the ReferenceTimezone field.
func (o *ReportExecutionSummaryContext) SetReferenceTimezone(v string) {
	o.ReferenceTimezone = &v
}

func (o ReportExecutionSummaryContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportExecutionSummaryContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceTimestamp) {
		toSerialize["reference_timestamp"] = o.ReferenceTimestamp
	}
	if !IsNil(o.ReferenceTimezone) {
		toSerialize["reference_timezone"] = o.ReferenceTimezone
	}
	return toSerialize, nil
}

type NullableReportExecutionSummaryContext struct {
	value *ReportExecutionSummaryContext
	isSet bool
}

func (v NullableReportExecutionSummaryContext) Get() *ReportExecutionSummaryContext {
	return v.value
}

func (v *NullableReportExecutionSummaryContext) Set(val *ReportExecutionSummaryContext) {
	v.value = val
	v.isSet = true
}

func (v NullableReportExecutionSummaryContext) IsSet() bool {
	return v.isSet
}

func (v *NullableReportExecutionSummaryContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportExecutionSummaryContext(val *ReportExecutionSummaryContext) *NullableReportExecutionSummaryContext {
	return &NullableReportExecutionSummaryContext{value: val, isSet: true}
}

func (v NullableReportExecutionSummaryContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportExecutionSummaryContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


