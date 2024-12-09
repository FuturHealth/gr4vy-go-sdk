# ReportSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;report&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this report. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Name** | Pointer to **string** | The name of this report. | [optional] 
**CreatorId** | Pointer to **NullableString** | The unique identifier for the creator of this report. | [optional] 
**CreatorDisplayName** | Pointer to **NullableString** | The name of the creator of this report. | [optional] 
**CreatorType** | Pointer to **NullableString** | The type of the creator of this report. | [optional] 

## Methods

### NewReportSummary

`func NewReportSummary() *ReportSummary`

NewReportSummary instantiates a new ReportSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSummaryWithDefaults

`func NewReportSummaryWithDefaults() *ReportSummary`

NewReportSummaryWithDefaults instantiates a new ReportSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ReportSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *ReportSummary) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *ReportSummary) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *ReportSummary) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *ReportSummary) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetName

`func (o *ReportSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatorId

`func (o *ReportSummary) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ReportSummary) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ReportSummary) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ReportSummary) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *ReportSummary) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *ReportSummary) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetCreatorDisplayName

`func (o *ReportSummary) GetCreatorDisplayName() string`

GetCreatorDisplayName returns the CreatorDisplayName field if non-nil, zero value otherwise.

### GetCreatorDisplayNameOk

`func (o *ReportSummary) GetCreatorDisplayNameOk() (*string, bool)`

GetCreatorDisplayNameOk returns a tuple with the CreatorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorDisplayName

`func (o *ReportSummary) SetCreatorDisplayName(v string)`

SetCreatorDisplayName sets CreatorDisplayName field to given value.

### HasCreatorDisplayName

`func (o *ReportSummary) HasCreatorDisplayName() bool`

HasCreatorDisplayName returns a boolean if a field has been set.

### SetCreatorDisplayNameNil

`func (o *ReportSummary) SetCreatorDisplayNameNil(b bool)`

 SetCreatorDisplayNameNil sets the value for CreatorDisplayName to be an explicit nil

### UnsetCreatorDisplayName
`func (o *ReportSummary) UnsetCreatorDisplayName()`

UnsetCreatorDisplayName ensures that no value is present for CreatorDisplayName, not even an explicit nil
### GetCreatorType

`func (o *ReportSummary) GetCreatorType() string`

GetCreatorType returns the CreatorType field if non-nil, zero value otherwise.

### GetCreatorTypeOk

`func (o *ReportSummary) GetCreatorTypeOk() (*string, bool)`

GetCreatorTypeOk returns a tuple with the CreatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorType

`func (o *ReportSummary) SetCreatorType(v string)`

SetCreatorType sets CreatorType field to given value.

### HasCreatorType

`func (o *ReportSummary) HasCreatorType() bool`

HasCreatorType returns a boolean if a field has been set.

### SetCreatorTypeNil

`func (o *ReportSummary) SetCreatorTypeNil(b bool)`

 SetCreatorTypeNil sets the value for CreatorType to be an explicit nil

### UnsetCreatorType
`func (o *ReportSummary) UnsetCreatorType()`

UnsetCreatorType ensures that no value is present for CreatorType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


