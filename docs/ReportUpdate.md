# ReportUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the report. | [optional] 
**Description** | Pointer to **NullableString** | The description of the report. | [optional] 
**ScheduleEnabled** | Pointer to **bool** | Indicates whether the report&#39;s scheduling is enabled. This value can only be set to &#x60;true&#x60; if this is a recurring report.  When this value is set to &#x60;true&#x60;, the report will be executed at the &#x60;next_execution_at&#x60; date and time.  When this value is set to &#x60;false&#x60;, future executions of the report are paused until this value is set to &#x60;true&#x60; again.  If scheduling is enabled after being disabled, then the &#x60;next_execution_at&#x60; value is updated if and only if its current value is a past date-time. The &#x60;next_execution_at&#x60; value is then set to the next closest date-time in the future depending on the values of &#x60;schedule&#x60; and &#x60;schedule_timezone&#x60;. | [optional] 

## Methods

### NewReportUpdate

`func NewReportUpdate() *ReportUpdate`

NewReportUpdate instantiates a new ReportUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportUpdateWithDefaults

`func NewReportUpdateWithDefaults() *ReportUpdate`

NewReportUpdateWithDefaults instantiates a new ReportUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReportUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ReportUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReportUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReportUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleEnabled

`func (o *ReportUpdate) GetScheduleEnabled() bool`

GetScheduleEnabled returns the ScheduleEnabled field if non-nil, zero value otherwise.

### GetScheduleEnabledOk

`func (o *ReportUpdate) GetScheduleEnabledOk() (*bool, bool)`

GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnabled

`func (o *ReportUpdate) SetScheduleEnabled(v bool)`

SetScheduleEnabled sets ScheduleEnabled field to given value.

### HasScheduleEnabled

`func (o *ReportUpdate) HasScheduleEnabled() bool`

HasScheduleEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


