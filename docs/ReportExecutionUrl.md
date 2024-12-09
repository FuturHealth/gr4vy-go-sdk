# ReportExecutionUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL to download the report execution. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The date and time when the download URL expires. | [optional] 

## Methods

### NewReportExecutionUrl

`func NewReportExecutionUrl() *ReportExecutionUrl`

NewReportExecutionUrl instantiates a new ReportExecutionUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportExecutionUrlWithDefaults

`func NewReportExecutionUrlWithDefaults() *ReportExecutionUrl`

NewReportExecutionUrlWithDefaults instantiates a new ReportExecutionUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ReportExecutionUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReportExecutionUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReportExecutionUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ReportExecutionUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ReportExecutionUrl) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ReportExecutionUrl) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ReportExecutionUrl) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ReportExecutionUrl) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


