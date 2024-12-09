# DigitalWalletUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantName** | Pointer to **string** | The name of the merchant. This is used to update the value initially used to register with a digital wallet provider and this name is not displayed to the buyer. | [optional] 
**DomainNames** | Pointer to **[]string** | The list of domain names that a digital wallet can be used on. To use a digital wallet on a website, the domain of the site is required to be in this list. | [optional] 
**MerchantDisplayName** | Pointer to **NullableString** | The consumer facing name of the merchant. | [optional] 
**MerchantCountryCode** | Pointer to **NullableString** | The country code where the merchant is registered. | [optional] 
**MerchantUrl** | Pointer to **string** | The main URL of the merchant. | [optional] 

## Methods

### NewDigitalWalletUpdate

`func NewDigitalWalletUpdate() *DigitalWalletUpdate`

NewDigitalWalletUpdate instantiates a new DigitalWalletUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletUpdateWithDefaults

`func NewDigitalWalletUpdateWithDefaults() *DigitalWalletUpdate`

NewDigitalWalletUpdateWithDefaults instantiates a new DigitalWalletUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantName

`func (o *DigitalWalletUpdate) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *DigitalWalletUpdate) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *DigitalWalletUpdate) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *DigitalWalletUpdate) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetDomainNames

`func (o *DigitalWalletUpdate) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *DigitalWalletUpdate) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *DigitalWalletUpdate) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *DigitalWalletUpdate) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetMerchantDisplayName

`func (o *DigitalWalletUpdate) GetMerchantDisplayName() string`

GetMerchantDisplayName returns the MerchantDisplayName field if non-nil, zero value otherwise.

### GetMerchantDisplayNameOk

`func (o *DigitalWalletUpdate) GetMerchantDisplayNameOk() (*string, bool)`

GetMerchantDisplayNameOk returns a tuple with the MerchantDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDisplayName

`func (o *DigitalWalletUpdate) SetMerchantDisplayName(v string)`

SetMerchantDisplayName sets MerchantDisplayName field to given value.

### HasMerchantDisplayName

`func (o *DigitalWalletUpdate) HasMerchantDisplayName() bool`

HasMerchantDisplayName returns a boolean if a field has been set.

### SetMerchantDisplayNameNil

`func (o *DigitalWalletUpdate) SetMerchantDisplayNameNil(b bool)`

 SetMerchantDisplayNameNil sets the value for MerchantDisplayName to be an explicit nil

### UnsetMerchantDisplayName
`func (o *DigitalWalletUpdate) UnsetMerchantDisplayName()`

UnsetMerchantDisplayName ensures that no value is present for MerchantDisplayName, not even an explicit nil
### GetMerchantCountryCode

`func (o *DigitalWalletUpdate) GetMerchantCountryCode() string`

GetMerchantCountryCode returns the MerchantCountryCode field if non-nil, zero value otherwise.

### GetMerchantCountryCodeOk

`func (o *DigitalWalletUpdate) GetMerchantCountryCodeOk() (*string, bool)`

GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCountryCode

`func (o *DigitalWalletUpdate) SetMerchantCountryCode(v string)`

SetMerchantCountryCode sets MerchantCountryCode field to given value.

### HasMerchantCountryCode

`func (o *DigitalWalletUpdate) HasMerchantCountryCode() bool`

HasMerchantCountryCode returns a boolean if a field has been set.

### SetMerchantCountryCodeNil

`func (o *DigitalWalletUpdate) SetMerchantCountryCodeNil(b bool)`

 SetMerchantCountryCodeNil sets the value for MerchantCountryCode to be an explicit nil

### UnsetMerchantCountryCode
`func (o *DigitalWalletUpdate) UnsetMerchantCountryCode()`

UnsetMerchantCountryCode ensures that no value is present for MerchantCountryCode, not even an explicit nil
### GetMerchantUrl

`func (o *DigitalWalletUpdate) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *DigitalWalletUpdate) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *DigitalWalletUpdate) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *DigitalWalletUpdate) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


