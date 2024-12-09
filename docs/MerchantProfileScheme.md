# MerchantProfileScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAcquirerBin** | Pointer to **string** | Acquirer bin to use when calling 3DS through this scheme. | [optional] 
**MerchantUrl** | Pointer to **string** | URL to send when calling 3DS through this scheme. | [optional] 
**MerchantAcquirerId** | Pointer to **string** | Merchant ID to use when calling 3DS through this scheme. | [optional] 
**MerchantName** | Pointer to **string** | Merchant name to use when calling 3DS through this scheme. | [optional] 
**MerchantCountryCode** | Pointer to **string** | Merchant country code to use when calling 3DS through this scheme. | [optional] 
**MerchantCategoryCode** | Pointer to **string** | Merchant category code to use when calling 3DS through this scheme. | [optional] 

## Methods

### NewMerchantProfileScheme

`func NewMerchantProfileScheme() *MerchantProfileScheme`

NewMerchantProfileScheme instantiates a new MerchantProfileScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProfileSchemeWithDefaults

`func NewMerchantProfileSchemeWithDefaults() *MerchantProfileScheme`

NewMerchantProfileSchemeWithDefaults instantiates a new MerchantProfileScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAcquirerBin

`func (o *MerchantProfileScheme) GetMerchantAcquirerBin() string`

GetMerchantAcquirerBin returns the MerchantAcquirerBin field if non-nil, zero value otherwise.

### GetMerchantAcquirerBinOk

`func (o *MerchantProfileScheme) GetMerchantAcquirerBinOk() (*string, bool)`

GetMerchantAcquirerBinOk returns a tuple with the MerchantAcquirerBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAcquirerBin

`func (o *MerchantProfileScheme) SetMerchantAcquirerBin(v string)`

SetMerchantAcquirerBin sets MerchantAcquirerBin field to given value.

### HasMerchantAcquirerBin

`func (o *MerchantProfileScheme) HasMerchantAcquirerBin() bool`

HasMerchantAcquirerBin returns a boolean if a field has been set.

### GetMerchantUrl

`func (o *MerchantProfileScheme) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *MerchantProfileScheme) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *MerchantProfileScheme) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *MerchantProfileScheme) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### GetMerchantAcquirerId

`func (o *MerchantProfileScheme) GetMerchantAcquirerId() string`

GetMerchantAcquirerId returns the MerchantAcquirerId field if non-nil, zero value otherwise.

### GetMerchantAcquirerIdOk

`func (o *MerchantProfileScheme) GetMerchantAcquirerIdOk() (*string, bool)`

GetMerchantAcquirerIdOk returns a tuple with the MerchantAcquirerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAcquirerId

`func (o *MerchantProfileScheme) SetMerchantAcquirerId(v string)`

SetMerchantAcquirerId sets MerchantAcquirerId field to given value.

### HasMerchantAcquirerId

`func (o *MerchantProfileScheme) HasMerchantAcquirerId() bool`

HasMerchantAcquirerId returns a boolean if a field has been set.

### GetMerchantName

`func (o *MerchantProfileScheme) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *MerchantProfileScheme) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *MerchantProfileScheme) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *MerchantProfileScheme) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetMerchantCountryCode

`func (o *MerchantProfileScheme) GetMerchantCountryCode() string`

GetMerchantCountryCode returns the MerchantCountryCode field if non-nil, zero value otherwise.

### GetMerchantCountryCodeOk

`func (o *MerchantProfileScheme) GetMerchantCountryCodeOk() (*string, bool)`

GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCountryCode

`func (o *MerchantProfileScheme) SetMerchantCountryCode(v string)`

SetMerchantCountryCode sets MerchantCountryCode field to given value.

### HasMerchantCountryCode

`func (o *MerchantProfileScheme) HasMerchantCountryCode() bool`

HasMerchantCountryCode returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *MerchantProfileScheme) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *MerchantProfileScheme) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *MerchantProfileScheme) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *MerchantProfileScheme) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


