# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The date and time this report was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time this report was last updated. | [optional] 
**NextExecutionAt** | Pointer to **NullableTime** | The date and time this report will next be executed, provided that &#x60;schedule_enabled&#x60; is &#x60;true&#x60;. This value is null if this is a one-off report. | [optional] 
**Description** | Pointer to **NullableString** | The description of this report. | [optional] 
**Schedule** | Pointer to **string** | Specifies the schedule of this report.  If this is a one-off report, this value is &#x60;once&#x60;.  If this is a recurring report, this value is set to the frequency by which the report will be executed. For example, a &#x60;monthly&#x60; schedule means that this report will be periodically executed at the start of each month.  Note that a &#x60;weekly&#x60; schedule means that the report is executed at the start of every Monday. | [optional] 
**ScheduleEnabled** | Pointer to **bool** | Indicates whether this report&#39;s scheduling is enabled. This value can only be set to &#x60;true&#x60; if this is a recurring report.  When this value is set to &#x60;true&#x60;, this report will be executed at the &#x60;next_execution_at&#x60; date and time.  When this value is set to &#x60;false&#x60;, future executions of this report are paused until this value is set to &#x60;true&#x60; again. | [optional] 
**ScheduleTimezone** | Pointer to **string** | The time zone in which the next execution will be scheduled. This value is used to calculate this report&#39;s &#x60;next_execution_at&#x60; value and is only relevant if this is a recurring report. This time zone is also used to calculate the timestamp range for reports that use date-time placeholders. Date-time placeholders are dynamic timestamps that change with every report execution. | [optional] 
**Spec** | Pointer to [**ReportSpec**](ReportSpec.md) |  | [optional] 
**LatestExecution** | Pointer to [**NullableReportExecutionSummary**](ReportExecutionSummary.md) | Details of the latest execution of this report. | [optional] 
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;report&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this report. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Name** | Pointer to **string** | The name of this report. | [optional] 
**CreatorId** | Pointer to **NullableString** | The unique identifier for the creator of this report. | [optional] 
**CreatorDisplayName** | Pointer to **NullableString** | The name of the creator of this report. | [optional] 
**CreatorType** | Pointer to **NullableString** | The type of the creator of this report. | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Report) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Report) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Report) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Report) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Report) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Report) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Report) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Report) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNextExecutionAt

`func (o *Report) GetNextExecutionAt() time.Time`

GetNextExecutionAt returns the NextExecutionAt field if non-nil, zero value otherwise.

### GetNextExecutionAtOk

`func (o *Report) GetNextExecutionAtOk() (*time.Time, bool)`

GetNextExecutionAtOk returns a tuple with the NextExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionAt

`func (o *Report) SetNextExecutionAt(v time.Time)`

SetNextExecutionAt sets NextExecutionAt field to given value.

### HasNextExecutionAt

`func (o *Report) HasNextExecutionAt() bool`

HasNextExecutionAt returns a boolean if a field has been set.

### SetNextExecutionAtNil

`func (o *Report) SetNextExecutionAtNil(b bool)`

 SetNextExecutionAtNil sets the value for NextExecutionAt to be an explicit nil

### UnsetNextExecutionAt
`func (o *Report) UnsetNextExecutionAt()`

UnsetNextExecutionAt ensures that no value is present for NextExecutionAt, not even an explicit nil
### GetDescription

`func (o *Report) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Report) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Report) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Report) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Report) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Report) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSchedule

`func (o *Report) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Report) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Report) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Report) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleEnabled

`func (o *Report) GetScheduleEnabled() bool`

GetScheduleEnabled returns the ScheduleEnabled field if non-nil, zero value otherwise.

### GetScheduleEnabledOk

`func (o *Report) GetScheduleEnabledOk() (*bool, bool)`

GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnabled

`func (o *Report) SetScheduleEnabled(v bool)`

SetScheduleEnabled sets ScheduleEnabled field to given value.

### HasScheduleEnabled

`func (o *Report) HasScheduleEnabled() bool`

HasScheduleEnabled returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *Report) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *Report) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *Report) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *Report) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetSpec

`func (o *Report) GetSpec() ReportSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Report) GetSpecOk() (*ReportSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Report) SetSpec(v ReportSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Report) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetLatestExecution

`func (o *Report) GetLatestExecution() ReportExecutionSummary`

GetLatestExecution returns the LatestExecution field if non-nil, zero value otherwise.

### GetLatestExecutionOk

`func (o *Report) GetLatestExecutionOk() (*ReportExecutionSummary, bool)`

GetLatestExecutionOk returns a tuple with the LatestExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecution

`func (o *Report) SetLatestExecution(v ReportExecutionSummary)`

SetLatestExecution sets LatestExecution field to given value.

### HasLatestExecution

`func (o *Report) HasLatestExecution() bool`

HasLatestExecution returns a boolean if a field has been set.

### SetLatestExecutionNil

`func (o *Report) SetLatestExecutionNil(b bool)`

 SetLatestExecutionNil sets the value for LatestExecution to be an explicit nil

### UnsetLatestExecution
`func (o *Report) UnsetLatestExecution()`

UnsetLatestExecution ensures that no value is present for LatestExecution, not even an explicit nil
### GetType

`func (o *Report) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Report) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Report) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Report) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Report) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Report) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *Report) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *Report) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *Report) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *Report) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetName

`func (o *Report) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Report) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Report) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Report) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatorId

`func (o *Report) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Report) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Report) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Report) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *Report) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Report) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetCreatorDisplayName

`func (o *Report) GetCreatorDisplayName() string`

GetCreatorDisplayName returns the CreatorDisplayName field if non-nil, zero value otherwise.

### GetCreatorDisplayNameOk

`func (o *Report) GetCreatorDisplayNameOk() (*string, bool)`

GetCreatorDisplayNameOk returns a tuple with the CreatorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorDisplayName

`func (o *Report) SetCreatorDisplayName(v string)`

SetCreatorDisplayName sets CreatorDisplayName field to given value.

### HasCreatorDisplayName

`func (o *Report) HasCreatorDisplayName() bool`

HasCreatorDisplayName returns a boolean if a field has been set.

### SetCreatorDisplayNameNil

`func (o *Report) SetCreatorDisplayNameNil(b bool)`

 SetCreatorDisplayNameNil sets the value for CreatorDisplayName to be an explicit nil

### UnsetCreatorDisplayName
`func (o *Report) UnsetCreatorDisplayName()`

UnsetCreatorDisplayName ensures that no value is present for CreatorDisplayName, not even an explicit nil
### GetCreatorType

`func (o *Report) GetCreatorType() string`

GetCreatorType returns the CreatorType field if non-nil, zero value otherwise.

### GetCreatorTypeOk

`func (o *Report) GetCreatorTypeOk() (*string, bool)`

GetCreatorTypeOk returns a tuple with the CreatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorType

`func (o *Report) SetCreatorType(v string)`

SetCreatorType sets CreatorType field to given value.

### HasCreatorType

`func (o *Report) HasCreatorType() bool`

HasCreatorType returns a boolean if a field has been set.

### SetCreatorTypeNil

`func (o *Report) SetCreatorTypeNil(b bool)`

 SetCreatorTypeNil sets the value for CreatorType to be an explicit nil

### UnsetCreatorType
`func (o *Report) UnsetCreatorType()`

UnsetCreatorType ensures that no value is present for CreatorType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


