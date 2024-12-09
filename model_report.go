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

// checks if the Report type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Report{}

// Report A report record.
type Report struct {
	// The date and time this report was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time this report was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The date and time this report will next be executed, provided that `schedule_enabled` is `true`. This value is null if this is a one-off report.
	NextExecutionAt NullableTime `json:"next_execution_at,omitempty"`
	// The description of this report.
	Description NullableString `json:"description,omitempty"`
	// Specifies the schedule of this report.  If this is a one-off report, this value is `once`.  If this is a recurring report, this value is set to the frequency by which the report will be executed. For example, a `monthly` schedule means that this report will be periodically executed at the start of each month.  Note that a `weekly` schedule means that the report is executed at the start of every Monday.
	Schedule *string `json:"schedule,omitempty"`
	// Indicates whether this report's scheduling is enabled. This value can only be set to `true` if this is a recurring report.  When this value is set to `true`, this report will be executed at the `next_execution_at` date and time.  When this value is set to `false`, future executions of this report are paused until this value is set to `true` again.
	ScheduleEnabled *bool `json:"schedule_enabled,omitempty"`
	// The time zone in which the next execution will be scheduled. This value is used to calculate this report's `next_execution_at` value and is only relevant if this is a recurring report. This time zone is also used to calculate the timestamp range for reports that use date-time placeholders. Date-time placeholders are dynamic timestamps that change with every report execution.
	ScheduleTimezone *string `json:"schedule_timezone,omitempty"`
	Spec *ReportSpec `json:"spec,omitempty"`
	// Details of the latest execution of this report.
	LatestExecution NullableReportExecutionSummary `json:"latest_execution,omitempty"`
	// The type of this resource. Is always `report`.
	Type *string `json:"type,omitempty"`
	// The unique identifier for this report.
	Id *string `json:"id,omitempty"`
	// The unique ID for a merchant account.
	MerchantAccountId *string `json:"merchant_account_id,omitempty"`
	// The name of this report.
	Name *string `json:"name,omitempty"`
	// The unique identifier for the creator of this report.
	CreatorId NullableString `json:"creator_id,omitempty"`
	// The name of the creator of this report.
	CreatorDisplayName NullableString `json:"creator_display_name,omitempty"`
	// The type of the creator of this report.
	CreatorType NullableString `json:"creator_type,omitempty"`
}

// NewReport instantiates a new Report object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReport() *Report {
	this := Report{}
	return &this
}

// NewReportWithDefaults instantiates a new Report object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportWithDefaults() *Report {
	this := Report{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Report) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Report) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Report) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Report) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Report) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Report) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNextExecutionAt returns the NextExecutionAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetNextExecutionAt() time.Time {
	if o == nil || IsNil(o.NextExecutionAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.NextExecutionAt.Get()
}

// GetNextExecutionAtOk returns a tuple with the NextExecutionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetNextExecutionAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextExecutionAt.Get(), o.NextExecutionAt.IsSet()
}

// HasNextExecutionAt returns a boolean if a field has been set.
func (o *Report) HasNextExecutionAt() bool {
	if o != nil && o.NextExecutionAt.IsSet() {
		return true
	}

	return false
}

// SetNextExecutionAt gets a reference to the given NullableTime and assigns it to the NextExecutionAt field.
func (o *Report) SetNextExecutionAt(v time.Time) {
	o.NextExecutionAt.Set(&v)
}
// SetNextExecutionAtNil sets the value for NextExecutionAt to be an explicit nil
func (o *Report) SetNextExecutionAtNil() {
	o.NextExecutionAt.Set(nil)
}

// UnsetNextExecutionAt ensures that no value is present for NextExecutionAt, not even an explicit nil
func (o *Report) UnsetNextExecutionAt() {
	o.NextExecutionAt.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Report) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Report) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Report) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Report) UnsetDescription() {
	o.Description.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *Report) GetSchedule() string {
	if o == nil || IsNil(o.Schedule) {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *Report) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *Report) SetSchedule(v string) {
	o.Schedule = &v
}

// GetScheduleEnabled returns the ScheduleEnabled field value if set, zero value otherwise.
func (o *Report) GetScheduleEnabled() bool {
	if o == nil || IsNil(o.ScheduleEnabled) {
		var ret bool
		return ret
	}
	return *o.ScheduleEnabled
}

// GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetScheduleEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScheduleEnabled) {
		return nil, false
	}
	return o.ScheduleEnabled, true
}

// HasScheduleEnabled returns a boolean if a field has been set.
func (o *Report) HasScheduleEnabled() bool {
	if o != nil && !IsNil(o.ScheduleEnabled) {
		return true
	}

	return false
}

// SetScheduleEnabled gets a reference to the given bool and assigns it to the ScheduleEnabled field.
func (o *Report) SetScheduleEnabled(v bool) {
	o.ScheduleEnabled = &v
}

// GetScheduleTimezone returns the ScheduleTimezone field value if set, zero value otherwise.
func (o *Report) GetScheduleTimezone() string {
	if o == nil || IsNil(o.ScheduleTimezone) {
		var ret string
		return ret
	}
	return *o.ScheduleTimezone
}

// GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetScheduleTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleTimezone) {
		return nil, false
	}
	return o.ScheduleTimezone, true
}

// HasScheduleTimezone returns a boolean if a field has been set.
func (o *Report) HasScheduleTimezone() bool {
	if o != nil && !IsNil(o.ScheduleTimezone) {
		return true
	}

	return false
}

// SetScheduleTimezone gets a reference to the given string and assigns it to the ScheduleTimezone field.
func (o *Report) SetScheduleTimezone(v string) {
	o.ScheduleTimezone = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *Report) GetSpec() ReportSpec {
	if o == nil || IsNil(o.Spec) {
		var ret ReportSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetSpecOk() (*ReportSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *Report) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ReportSpec and assigns it to the Spec field.
func (o *Report) SetSpec(v ReportSpec) {
	o.Spec = &v
}

// GetLatestExecution returns the LatestExecution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetLatestExecution() ReportExecutionSummary {
	if o == nil || IsNil(o.LatestExecution.Get()) {
		var ret ReportExecutionSummary
		return ret
	}
	return *o.LatestExecution.Get()
}

// GetLatestExecutionOk returns a tuple with the LatestExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetLatestExecutionOk() (*ReportExecutionSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestExecution.Get(), o.LatestExecution.IsSet()
}

// HasLatestExecution returns a boolean if a field has been set.
func (o *Report) HasLatestExecution() bool {
	if o != nil && o.LatestExecution.IsSet() {
		return true
	}

	return false
}

// SetLatestExecution gets a reference to the given NullableReportExecutionSummary and assigns it to the LatestExecution field.
func (o *Report) SetLatestExecution(v ReportExecutionSummary) {
	o.LatestExecution.Set(&v)
}
// SetLatestExecutionNil sets the value for LatestExecution to be an explicit nil
func (o *Report) SetLatestExecutionNil() {
	o.LatestExecution.Set(nil)
}

// UnsetLatestExecution ensures that no value is present for LatestExecution, not even an explicit nil
func (o *Report) UnsetLatestExecution() {
	o.LatestExecution.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Report) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Report) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Report) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Report) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Report) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Report) SetId(v string) {
	o.Id = &v
}

// GetMerchantAccountId returns the MerchantAccountId field value if set, zero value otherwise.
func (o *Report) GetMerchantAccountId() string {
	if o == nil || IsNil(o.MerchantAccountId) {
		var ret string
		return ret
	}
	return *o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccountId) {
		return nil, false
	}
	return o.MerchantAccountId, true
}

// HasMerchantAccountId returns a boolean if a field has been set.
func (o *Report) HasMerchantAccountId() bool {
	if o != nil && !IsNil(o.MerchantAccountId) {
		return true
	}

	return false
}

// SetMerchantAccountId gets a reference to the given string and assigns it to the MerchantAccountId field.
func (o *Report) SetMerchantAccountId(v string) {
	o.MerchantAccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Report) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Report) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Report) SetName(v string) {
	o.Name = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *Report) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *Report) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *Report) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *Report) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetCreatorDisplayName returns the CreatorDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetCreatorDisplayName() string {
	if o == nil || IsNil(o.CreatorDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorDisplayName.Get()
}

// GetCreatorDisplayNameOk returns a tuple with the CreatorDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetCreatorDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorDisplayName.Get(), o.CreatorDisplayName.IsSet()
}

// HasCreatorDisplayName returns a boolean if a field has been set.
func (o *Report) HasCreatorDisplayName() bool {
	if o != nil && o.CreatorDisplayName.IsSet() {
		return true
	}

	return false
}

// SetCreatorDisplayName gets a reference to the given NullableString and assigns it to the CreatorDisplayName field.
func (o *Report) SetCreatorDisplayName(v string) {
	o.CreatorDisplayName.Set(&v)
}
// SetCreatorDisplayNameNil sets the value for CreatorDisplayName to be an explicit nil
func (o *Report) SetCreatorDisplayNameNil() {
	o.CreatorDisplayName.Set(nil)
}

// UnsetCreatorDisplayName ensures that no value is present for CreatorDisplayName, not even an explicit nil
func (o *Report) UnsetCreatorDisplayName() {
	o.CreatorDisplayName.Unset()
}

// GetCreatorType returns the CreatorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetCreatorType() string {
	if o == nil || IsNil(o.CreatorType.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorType.Get()
}

// GetCreatorTypeOk returns a tuple with the CreatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetCreatorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorType.Get(), o.CreatorType.IsSet()
}

// HasCreatorType returns a boolean if a field has been set.
func (o *Report) HasCreatorType() bool {
	if o != nil && o.CreatorType.IsSet() {
		return true
	}

	return false
}

// SetCreatorType gets a reference to the given NullableString and assigns it to the CreatorType field.
func (o *Report) SetCreatorType(v string) {
	o.CreatorType.Set(&v)
}
// SetCreatorTypeNil sets the value for CreatorType to be an explicit nil
func (o *Report) SetCreatorTypeNil() {
	o.CreatorType.Set(nil)
}

// UnsetCreatorType ensures that no value is present for CreatorType, not even an explicit nil
func (o *Report) UnsetCreatorType() {
	o.CreatorType.Unset()
}

func (o Report) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Report) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.NextExecutionAt.IsSet() {
		toSerialize["next_execution_at"] = o.NextExecutionAt.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.ScheduleEnabled) {
		toSerialize["schedule_enabled"] = o.ScheduleEnabled
	}
	if !IsNil(o.ScheduleTimezone) {
		toSerialize["schedule_timezone"] = o.ScheduleTimezone
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if o.LatestExecution.IsSet() {
		toSerialize["latest_execution"] = o.LatestExecution.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantAccountId) {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.CreatorId.IsSet() {
		toSerialize["creator_id"] = o.CreatorId.Get()
	}
	if o.CreatorDisplayName.IsSet() {
		toSerialize["creator_display_name"] = o.CreatorDisplayName.Get()
	}
	if o.CreatorType.IsSet() {
		toSerialize["creator_type"] = o.CreatorType.Get()
	}
	return toSerialize, nil
}

type NullableReport struct {
	value *Report
	isSet bool
}

func (v NullableReport) Get() *Report {
	return v.value
}

func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

func (v NullableReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


