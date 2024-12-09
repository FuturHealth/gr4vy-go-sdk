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

// checks if the ReportCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportCreate{}

// ReportCreate A request to create a report.
type ReportCreate struct {
	// The name of the report.
	Name string `json:"name"`
	// The description of the report.
	Description NullableString `json:"description,omitempty"`
	// Specifies the schedule of the report.  If this is a one-off report, set this value to `once`.  If this is a recurring report, this value should be set to the frequency by which the report will be executed. For example, a `monthly` schedule means that the report will be periodically executed at the start of each month.  Note that a `weekly` schedule means that the report will be executed at the start of every Monday.
	Schedule *string `json:"schedule,omitempty"`
	// Indicates whether the report's scheduling is enabled. This value can only be set to `true` if this is a recurring report.  If this value is set to `true`, the report will be executed at the `next_execution_at` date and time.  If this is a recurring report and this value is set to `false`, executions of the report will not occur until this value is set to `true`.  If this value is not provided, `schedule_enabled` will automatically be set to `false` if `schedule` is `once` and set to `true` otherwise.
	ScheduleEnabled NullableBool `json:"schedule_enabled,omitempty"`
	// The time zone in which the report's executions will be scheduled. This value is used to compute the report's `next_execution_at` value and is only relevant when this is a recurring report. This time zone is also used to calculate the timestamp range for reports that use date-time placeholders. Date-time placeholders are dynamic timestamps that change with every report execution.  This value must be set to the time zone's name as presented in the IANA time zone database. For example, to schedule reports in the time zone of New York, set this value to `America/New_York`.
	ScheduleTimezone *string `json:"schedule_timezone,omitempty"`
	Spec ReportSpec `json:"spec"`
}

type _ReportCreate ReportCreate

// NewReportCreate instantiates a new ReportCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCreate(name string, spec ReportSpec) *ReportCreate {
	this := ReportCreate{}
	this.Name = name
	var schedule string = "once"
	this.Schedule = &schedule
	var scheduleTimezone string = "Etc/UTC"
	this.ScheduleTimezone = &scheduleTimezone
	this.Spec = spec
	return &this
}

// NewReportCreateWithDefaults instantiates a new ReportCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCreateWithDefaults() *ReportCreate {
	this := ReportCreate{}
	var schedule string = "once"
	this.Schedule = &schedule
	var scheduleTimezone string = "Etc/UTC"
	this.ScheduleTimezone = &scheduleTimezone
	return &this
}

// GetName returns the Name field value
func (o *ReportCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReportCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReportCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreate) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ReportCreate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ReportCreate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ReportCreate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ReportCreate) UnsetDescription() {
	o.Description.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ReportCreate) GetSchedule() string {
	if o == nil || IsNil(o.Schedule) {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCreate) GetScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ReportCreate) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *ReportCreate) SetSchedule(v string) {
	o.Schedule = &v
}

// GetScheduleEnabled returns the ScheduleEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreate) GetScheduleEnabled() bool {
	if o == nil || IsNil(o.ScheduleEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.ScheduleEnabled.Get()
}

// GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreate) GetScheduleEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleEnabled.Get(), o.ScheduleEnabled.IsSet()
}

// HasScheduleEnabled returns a boolean if a field has been set.
func (o *ReportCreate) HasScheduleEnabled() bool {
	if o != nil && o.ScheduleEnabled.IsSet() {
		return true
	}

	return false
}

// SetScheduleEnabled gets a reference to the given NullableBool and assigns it to the ScheduleEnabled field.
func (o *ReportCreate) SetScheduleEnabled(v bool) {
	o.ScheduleEnabled.Set(&v)
}
// SetScheduleEnabledNil sets the value for ScheduleEnabled to be an explicit nil
func (o *ReportCreate) SetScheduleEnabledNil() {
	o.ScheduleEnabled.Set(nil)
}

// UnsetScheduleEnabled ensures that no value is present for ScheduleEnabled, not even an explicit nil
func (o *ReportCreate) UnsetScheduleEnabled() {
	o.ScheduleEnabled.Unset()
}

// GetScheduleTimezone returns the ScheduleTimezone field value if set, zero value otherwise.
func (o *ReportCreate) GetScheduleTimezone() string {
	if o == nil || IsNil(o.ScheduleTimezone) {
		var ret string
		return ret
	}
	return *o.ScheduleTimezone
}

// GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportCreate) GetScheduleTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleTimezone) {
		return nil, false
	}
	return o.ScheduleTimezone, true
}

// HasScheduleTimezone returns a boolean if a field has been set.
func (o *ReportCreate) HasScheduleTimezone() bool {
	if o != nil && !IsNil(o.ScheduleTimezone) {
		return true
	}

	return false
}

// SetScheduleTimezone gets a reference to the given string and assigns it to the ScheduleTimezone field.
func (o *ReportCreate) SetScheduleTimezone(v string) {
	o.ScheduleTimezone = &v
}

// GetSpec returns the Spec field value
func (o *ReportCreate) GetSpec() ReportSpec {
	if o == nil {
		var ret ReportSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ReportCreate) GetSpecOk() (*ReportSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ReportCreate) SetSpec(v ReportSpec) {
	o.Spec = v
}

func (o ReportCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if o.ScheduleEnabled.IsSet() {
		toSerialize["schedule_enabled"] = o.ScheduleEnabled.Get()
	}
	if !IsNil(o.ScheduleTimezone) {
		toSerialize["schedule_timezone"] = o.ScheduleTimezone
	}
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

func (o *ReportCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"spec",
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

	varReportCreate := _ReportCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportCreate)

	if err != nil {
		return err
	}

	*o = ReportCreate(varReportCreate)

	return err
}

type NullableReportCreate struct {
	value *ReportCreate
	isSet bool
}

func (v NullableReportCreate) Get() *ReportCreate {
	return v.value
}

func (v *NullableReportCreate) Set(val *ReportCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCreate(val *ReportCreate) *NullableReportCreate {
	return &NullableReportCreate{value: val, isSet: true}
}

func (v NullableReportCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


