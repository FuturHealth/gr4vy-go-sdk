# DigitalWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The name of the digital wallet provider. | 
**MerchantName** | **string** | The name of the merchant. This is used to register the merchant with a digital wallet provider and this name is not displayed to the buyer. | 
**MerchantUrl** | Pointer to **NullableString** | The main URL of the merchant. | [optional] 
**MerchantDisplayName** | Pointer to **NullableString** | The consumer facing name of the merchant. | [optional] 
**MerchantCountryCode** | Pointer to **NullableString** | The country code where the merchant is registered. | [optional] 
**DomainNames** | **[]string** | The list of domain names that a digital wallet can be used on. To use a digital wallet on a website, the domain of the site is required to be in this list. | 
**AcceptTermsAndConditions** | **bool** | The explicit acceptance of the digital wallet provider&#39;s terms and conditions by the merchant. Needs to be &#x60;true&#x60; to register a new digital wallet. | 

## Methods

### NewDigitalWalletRequest

`func NewDigitalWalletRequest(provider string, merchantName string, domainNames []string, acceptTermsAndConditions bool, ) *DigitalWalletRequest`

NewDigitalWalletRequest instantiates a new DigitalWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletRequestWithDefaults

`func NewDigitalWalletRequestWithDefaults() *DigitalWalletRequest`

NewDigitalWalletRequestWithDefaults instantiates a new DigitalWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *DigitalWalletRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DigitalWalletRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DigitalWalletRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMerchantName

`func (o *DigitalWalletRequest) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *DigitalWalletRequest) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *DigitalWalletRequest) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.


### GetMerchantUrl

`func (o *DigitalWalletRequest) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *DigitalWalletRequest) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *DigitalWalletRequest) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *DigitalWalletRequest) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### SetMerchantUrlNil

`func (o *DigitalWalletRequest) SetMerchantUrlNil(b bool)`

 SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil

### UnsetMerchantUrl
`func (o *DigitalWalletRequest) UnsetMerchantUrl()`

UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
### GetMerchantDisplayName

`func (o *DigitalWalletRequest) GetMerchantDisplayName() string`

GetMerchantDisplayName returns the MerchantDisplayName field if non-nil, zero value otherwise.

### GetMerchantDisplayNameOk

`func (o *DigitalWalletRequest) GetMerchantDisplayNameOk() (*string, bool)`

GetMerchantDisplayNameOk returns a tuple with the MerchantDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDisplayName

`func (o *DigitalWalletRequest) SetMerchantDisplayName(v string)`

SetMerchantDisplayName sets MerchantDisplayName field to given value.

### HasMerchantDisplayName

`func (o *DigitalWalletRequest) HasMerchantDisplayName() bool`

HasMerchantDisplayName returns a boolean if a field has been set.

### SetMerchantDisplayNameNil

`func (o *DigitalWalletRequest) SetMerchantDisplayNameNil(b bool)`

 SetMerchantDisplayNameNil sets the value for MerchantDisplayName to be an explicit nil

### UnsetMerchantDisplayName
`func (o *DigitalWalletRequest) UnsetMerchantDisplayName()`

UnsetMerchantDisplayName ensures that no value is present for MerchantDisplayName, not even an explicit nil
### GetMerchantCountryCode

`func (o *DigitalWalletRequest) GetMerchantCountryCode() string`

GetMerchantCountryCode returns the MerchantCountryCode field if non-nil, zero value otherwise.

### GetMerchantCountryCodeOk

`func (o *DigitalWalletRequest) GetMerchantCountryCodeOk() (*string, bool)`

GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCountryCode

`func (o *DigitalWalletRequest) SetMerchantCountryCode(v string)`

SetMerchantCountryCode sets MerchantCountryCode field to given value.

### HasMerchantCountryCode

`func (o *DigitalWalletRequest) HasMerchantCountryCode() bool`

HasMerchantCountryCode returns a boolean if a field has been set.

### SetMerchantCountryCodeNil

`func (o *DigitalWalletRequest) SetMerchantCountryCodeNil(b bool)`

 SetMerchantCountryCodeNil sets the value for MerchantCountryCode to be an explicit nil

### UnsetMerchantCountryCode
`func (o *DigitalWalletRequest) UnsetMerchantCountryCode()`

UnsetMerchantCountryCode ensures that no value is present for MerchantCountryCode, not even an explicit nil
### GetDomainNames

`func (o *DigitalWalletRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *DigitalWalletRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *DigitalWalletRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetAcceptTermsAndConditions

`func (o *DigitalWalletRequest) GetAcceptTermsAndConditions() bool`

GetAcceptTermsAndConditions returns the AcceptTermsAndConditions field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsOk

`func (o *DigitalWalletRequest) GetAcceptTermsAndConditionsOk() (*bool, bool)`

GetAcceptTermsAndConditionsOk returns a tuple with the AcceptTermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditions

`func (o *DigitalWalletRequest) SetAcceptTermsAndConditions(v bool)`

SetAcceptTermsAndConditions sets AcceptTermsAndConditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


