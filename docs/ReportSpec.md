# ReportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | The model (dataset) that the data used for the report is retrieved from. | 
**Params** | **map[string]interface{}** | Parameters used to configure the report. Acceptable values for this property depend on the value specified for &#x60;model&#x60;. | 

## Methods

### NewReportSpec

`func NewReportSpec(model string, params map[string]interface{}, ) *ReportSpec`

NewReportSpec instantiates a new ReportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSpecWithDefaults

`func NewReportSpecWithDefaults() *ReportSpec`

NewReportSpecWithDefaults instantiates a new ReportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ReportSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ReportSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ReportSpec) SetModel(v string)`

SetModel sets Model field to given value.


### GetParams

`func (o *ReportSpec) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ReportSpec) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ReportSpec) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


