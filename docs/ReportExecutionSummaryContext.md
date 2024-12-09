# ReportExecutionSummaryContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceTimestamp** | Pointer to **time.Time** | The date and time used by the system as a reference point to compute date-time placeholders. | [optional] 
**ReferenceTimezone** | Pointer to **string** | The time zone used to compute date-time placeholders. | [optional] 

## Methods

### NewReportExecutionSummaryContext

`func NewReportExecutionSummaryContext() *ReportExecutionSummaryContext`

NewReportExecutionSummaryContext instantiates a new ReportExecutionSummaryContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportExecutionSummaryContextWithDefaults

`func NewReportExecutionSummaryContextWithDefaults() *ReportExecutionSummaryContext`

NewReportExecutionSummaryContextWithDefaults instantiates a new ReportExecutionSummaryContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceTimestamp

`func (o *ReportExecutionSummaryContext) GetReferenceTimestamp() time.Time`

GetReferenceTimestamp returns the ReferenceTimestamp field if non-nil, zero value otherwise.

### GetReferenceTimestampOk

`func (o *ReportExecutionSummaryContext) GetReferenceTimestampOk() (*time.Time, bool)`

GetReferenceTimestampOk returns a tuple with the ReferenceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTimestamp

`func (o *ReportExecutionSummaryContext) SetReferenceTimestamp(v time.Time)`

SetReferenceTimestamp sets ReferenceTimestamp field to given value.

### HasReferenceTimestamp

`func (o *ReportExecutionSummaryContext) HasReferenceTimestamp() bool`

HasReferenceTimestamp returns a boolean if a field has been set.

### GetReferenceTimezone

`func (o *ReportExecutionSummaryContext) GetReferenceTimezone() string`

GetReferenceTimezone returns the ReferenceTimezone field if non-nil, zero value otherwise.

### GetReferenceTimezoneOk

`func (o *ReportExecutionSummaryContext) GetReferenceTimezoneOk() (*string, bool)`

GetReferenceTimezoneOk returns a tuple with the ReferenceTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTimezone

`func (o *ReportExecutionSummaryContext) SetReferenceTimezone(v string)`

SetReferenceTimezone sets ReferenceTimezone field to given value.

### HasReferenceTimezone

`func (o *ReportExecutionSummaryContext) HasReferenceTimezone() bool`

HasReferenceTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


