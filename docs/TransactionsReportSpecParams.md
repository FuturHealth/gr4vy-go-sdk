# TransactionsReportSpecParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to **[]string** | A list of fields for the report. | [optional] 
**Filters** | Pointer to [**TransactionsReportSpecParamsFilters**](TransactionsReportSpecParamsFilters.md) |  | [optional] 
**Sort** | Pointer to **[]map[string]interface{}** | A list of fields to sort the report. | [optional] 

## Methods

### NewTransactionsReportSpecParams

`func NewTransactionsReportSpecParams() *TransactionsReportSpecParams`

NewTransactionsReportSpecParams instantiates a new TransactionsReportSpecParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsReportSpecParamsWithDefaults

`func NewTransactionsReportSpecParamsWithDefaults() *TransactionsReportSpecParams`

NewTransactionsReportSpecParamsWithDefaults instantiates a new TransactionsReportSpecParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *TransactionsReportSpecParams) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TransactionsReportSpecParams) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TransactionsReportSpecParams) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TransactionsReportSpecParams) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilters

`func (o *TransactionsReportSpecParams) GetFilters() TransactionsReportSpecParamsFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TransactionsReportSpecParams) GetFiltersOk() (*TransactionsReportSpecParamsFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TransactionsReportSpecParams) SetFilters(v TransactionsReportSpecParamsFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TransactionsReportSpecParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *TransactionsReportSpecParams) GetSort() []map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TransactionsReportSpecParams) GetSortOk() (*[]map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TransactionsReportSpecParams) SetSort(v []map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *TransactionsReportSpecParams) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


