# ReportCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the report. | 
**Description** | Pointer to **NullableString** | The description of the report. | [optional] 
**Schedule** | Pointer to **string** | Specifies the schedule of the report.  If this is a one-off report, set this value to &#x60;once&#x60;.  If this is a recurring report, this value should be set to the frequency by which the report will be executed. For example, a &#x60;monthly&#x60; schedule means that the report will be periodically executed at the start of each month.  Note that a &#x60;weekly&#x60; schedule means that the report will be executed at the start of every Monday. | [optional] [default to "once"]
**ScheduleEnabled** | Pointer to **NullableBool** | Indicates whether the report&#39;s scheduling is enabled. This value can only be set to &#x60;true&#x60; if this is a recurring report.  If this value is set to &#x60;true&#x60;, the report will be executed at the &#x60;next_execution_at&#x60; date and time.  If this is a recurring report and this value is set to &#x60;false&#x60;, executions of the report will not occur until this value is set to &#x60;true&#x60;.  If this value is not provided, &#x60;schedule_enabled&#x60; will automatically be set to &#x60;false&#x60; if &#x60;schedule&#x60; is &#x60;once&#x60; and set to &#x60;true&#x60; otherwise. | [optional] 
**ScheduleTimezone** | Pointer to **string** | The time zone in which the report&#39;s executions will be scheduled. This value is used to compute the report&#39;s &#x60;next_execution_at&#x60; value and is only relevant when this is a recurring report. This time zone is also used to calculate the timestamp range for reports that use date-time placeholders. Date-time placeholders are dynamic timestamps that change with every report execution.  This value must be set to the time zone&#39;s name as presented in the IANA time zone database. For example, to schedule reports in the time zone of New York, set this value to &#x60;America/New_York&#x60;. | [optional] [default to "Etc/UTC"]
**Spec** | [**ReportSpec**](ReportSpec.md) |  | 

## Methods

### NewReportCreate

`func NewReportCreate(name string, spec ReportSpec, ) *ReportCreate`

NewReportCreate instantiates a new ReportCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCreateWithDefaults

`func NewReportCreateWithDefaults() *ReportCreate`

NewReportCreateWithDefaults instantiates a new ReportCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReportCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReportCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReportCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReportCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSchedule

`func (o *ReportCreate) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ReportCreate) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ReportCreate) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ReportCreate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleEnabled

`func (o *ReportCreate) GetScheduleEnabled() bool`

GetScheduleEnabled returns the ScheduleEnabled field if non-nil, zero value otherwise.

### GetScheduleEnabledOk

`func (o *ReportCreate) GetScheduleEnabledOk() (*bool, bool)`

GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnabled

`func (o *ReportCreate) SetScheduleEnabled(v bool)`

SetScheduleEnabled sets ScheduleEnabled field to given value.

### HasScheduleEnabled

`func (o *ReportCreate) HasScheduleEnabled() bool`

HasScheduleEnabled returns a boolean if a field has been set.

### SetScheduleEnabledNil

`func (o *ReportCreate) SetScheduleEnabledNil(b bool)`

 SetScheduleEnabledNil sets the value for ScheduleEnabled to be an explicit nil

### UnsetScheduleEnabled
`func (o *ReportCreate) UnsetScheduleEnabled()`

UnsetScheduleEnabled ensures that no value is present for ScheduleEnabled, not even an explicit nil
### GetScheduleTimezone

`func (o *ReportCreate) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *ReportCreate) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *ReportCreate) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *ReportCreate) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetSpec

`func (o *ReportCreate) GetSpec() ReportSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReportCreate) GetSpecOk() (*ReportSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReportCreate) SetSpec(v ReportSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


