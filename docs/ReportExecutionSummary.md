# ReportExecutionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;report-execution&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this report execution. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time this report execution was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time this report execution was last updated. | [optional] 
**Status** | Pointer to **string** | The status of this report execution. | [optional] 
**Context** | Pointer to [**ReportExecutionSummaryContext**](ReportExecutionSummaryContext.md) |  | [optional] 

## Methods

### NewReportExecutionSummary

`func NewReportExecutionSummary() *ReportExecutionSummary`

NewReportExecutionSummary instantiates a new ReportExecutionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportExecutionSummaryWithDefaults

`func NewReportExecutionSummaryWithDefaults() *ReportExecutionSummary`

NewReportExecutionSummaryWithDefaults instantiates a new ReportExecutionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportExecutionSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportExecutionSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportExecutionSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportExecutionSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ReportExecutionSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportExecutionSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportExecutionSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportExecutionSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReportExecutionSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportExecutionSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportExecutionSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReportExecutionSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReportExecutionSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReportExecutionSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReportExecutionSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReportExecutionSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ReportExecutionSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportExecutionSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportExecutionSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportExecutionSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContext

`func (o *ReportExecutionSummary) GetContext() ReportExecutionSummaryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ReportExecutionSummary) GetContextOk() (*ReportExecutionSummaryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ReportExecutionSummary) SetContext(v ReportExecutionSummaryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ReportExecutionSummary) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


