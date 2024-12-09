# MerchantProfileSchemeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAcquirerBin** | Pointer to **string** | Acquirer bin to use when calling 3DS through this scheme. | [optional] 
**MerchantUrl** | Pointer to **string** | URL to send when calling 3DS through this scheme. | [optional] 
**MerchantAcquirerId** | Pointer to **string** | Merchant ID to use when calling 3DS through this scheme. | [optional] 
**MerchantName** | Pointer to **string** | Merchant name to use when calling 3DS through this scheme. | [optional] 
**MerchantCountryCode** | Pointer to **string** | Merchant country code to use when calling 3DS through this scheme. | [optional] 
**MerchantCategoryCode** | Pointer to **string** | Merchant category code to use when calling 3DS through this scheme. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this profile was created. | [optional] 

## Methods

### NewMerchantProfileSchemeSummary

`func NewMerchantProfileSchemeSummary() *MerchantProfileSchemeSummary`

NewMerchantProfileSchemeSummary instantiates a new MerchantProfileSchemeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProfileSchemeSummaryWithDefaults

`func NewMerchantProfileSchemeSummaryWithDefaults() *MerchantProfileSchemeSummary`

NewMerchantProfileSchemeSummaryWithDefaults instantiates a new MerchantProfileSchemeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAcquirerBin

`func (o *MerchantProfileSchemeSummary) GetMerchantAcquirerBin() string`

GetMerchantAcquirerBin returns the MerchantAcquirerBin field if non-nil, zero value otherwise.

### GetMerchantAcquirerBinOk

`func (o *MerchantProfileSchemeSummary) GetMerchantAcquirerBinOk() (*string, bool)`

GetMerchantAcquirerBinOk returns a tuple with the MerchantAcquirerBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAcquirerBin

`func (o *MerchantProfileSchemeSummary) SetMerchantAcquirerBin(v string)`

SetMerchantAcquirerBin sets MerchantAcquirerBin field to given value.

### HasMerchantAcquirerBin

`func (o *MerchantProfileSchemeSummary) HasMerchantAcquirerBin() bool`

HasMerchantAcquirerBin returns a boolean if a field has been set.

### GetMerchantUrl

`func (o *MerchantProfileSchemeSummary) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *MerchantProfileSchemeSummary) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *MerchantProfileSchemeSummary) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *MerchantProfileSchemeSummary) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### GetMerchantAcquirerId

`func (o *MerchantProfileSchemeSummary) GetMerchantAcquirerId() string`

GetMerchantAcquirerId returns the MerchantAcquirerId field if non-nil, zero value otherwise.

### GetMerchantAcquirerIdOk

`func (o *MerchantProfileSchemeSummary) GetMerchantAcquirerIdOk() (*string, bool)`

GetMerchantAcquirerIdOk returns a tuple with the MerchantAcquirerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAcquirerId

`func (o *MerchantProfileSchemeSummary) SetMerchantAcquirerId(v string)`

SetMerchantAcquirerId sets MerchantAcquirerId field to given value.

### HasMerchantAcquirerId

`func (o *MerchantProfileSchemeSummary) HasMerchantAcquirerId() bool`

HasMerchantAcquirerId returns a boolean if a field has been set.

### GetMerchantName

`func (o *MerchantProfileSchemeSummary) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *MerchantProfileSchemeSummary) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *MerchantProfileSchemeSummary) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *MerchantProfileSchemeSummary) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetMerchantCountryCode

`func (o *MerchantProfileSchemeSummary) GetMerchantCountryCode() string`

GetMerchantCountryCode returns the MerchantCountryCode field if non-nil, zero value otherwise.

### GetMerchantCountryCodeOk

`func (o *MerchantProfileSchemeSummary) GetMerchantCountryCodeOk() (*string, bool)`

GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCountryCode

`func (o *MerchantProfileSchemeSummary) SetMerchantCountryCode(v string)`

SetMerchantCountryCode sets MerchantCountryCode field to given value.

### HasMerchantCountryCode

`func (o *MerchantProfileSchemeSummary) HasMerchantCountryCode() bool`

HasMerchantCountryCode returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *MerchantProfileSchemeSummary) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *MerchantProfileSchemeSummary) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *MerchantProfileSchemeSummary) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *MerchantProfileSchemeSummary) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MerchantProfileSchemeSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantProfileSchemeSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantProfileSchemeSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantProfileSchemeSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


