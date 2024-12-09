# ReportExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | Pointer to [**ReportSummary**](ReportSummary.md) | The report associated with this report execution. | [optional] 
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;report-execution&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this report execution. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time this report execution was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time this report execution was last updated. | [optional] 
**Status** | Pointer to **string** | The status of this report execution. | [optional] 
**Context** | Pointer to [**ReportExecutionSummaryContext**](ReportExecutionSummaryContext.md) |  | [optional] 

## Methods

### NewReportExecution

`func NewReportExecution() *ReportExecution`

NewReportExecution instantiates a new ReportExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportExecutionWithDefaults

`func NewReportExecutionWithDefaults() *ReportExecution`

NewReportExecutionWithDefaults instantiates a new ReportExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *ReportExecution) GetReport() ReportSummary`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ReportExecution) GetReportOk() (*ReportSummary, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ReportExecution) SetReport(v ReportSummary)`

SetReport sets Report field to given value.

### HasReport

`func (o *ReportExecution) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetType

`func (o *ReportExecution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportExecution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportExecution) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportExecution) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ReportExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReportExecution) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportExecution) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportExecution) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReportExecution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReportExecution) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReportExecution) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReportExecution) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReportExecution) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ReportExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContext

`func (o *ReportExecution) GetContext() ReportExecutionSummaryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ReportExecution) GetContextOk() (*ReportExecutionSummaryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ReportExecution) SetContext(v ReportExecutionSummaryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ReportExecution) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


