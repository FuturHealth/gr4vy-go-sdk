# GooglePayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;googlepay&#x60;. | 
**Token** | **string** | The encrypted (opaque) token returned by the Google Pay API that represents a payment method. | 
**CardSuffix** | Pointer to **NullableString** | Last 4 digits of the PAN for identification purposes. | [optional] 
**CardScheme** | Pointer to **NullableString** | The scheme/brand of the card. | [optional] 
**CardType** | Pointer to **NullableString** | The type of card. | [optional] 
**AssuranceDetails** | Pointer to [**NullableGooglePayRequestAssuranceDetails**](GooglePayRequestAssuranceDetails.md) |  | [optional] 
**CardholderName** | Pointer to **NullableString** | Name of the card holder. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | We strongly recommend providing a &#x60;redirect_url&#x60; either when 3-D Secure is enabled and &#x60;three_d_secure_data&#x60; is not provided, or when using connections where 3DS is enabled. This value will be appended with both a transaction ID and status (e.g. &#x60;https://example.com/callback?gr4vy_transaction_id&#x3D;123 &amp;gr4vy_transaction_status&#x3D;capture_succeeded&#x60;) after 3-D Secure has completed. For those cases, if the value is not present, the transaction will be marked as failed. | [optional] 

## Methods

### NewGooglePayRequest

`func NewGooglePayRequest(method string, token string, ) *GooglePayRequest`

NewGooglePayRequest instantiates a new GooglePayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayRequestWithDefaults

`func NewGooglePayRequestWithDefaults() *GooglePayRequest`

NewGooglePayRequestWithDefaults instantiates a new GooglePayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *GooglePayRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GooglePayRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GooglePayRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *GooglePayRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GooglePayRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GooglePayRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetCardSuffix

`func (o *GooglePayRequest) GetCardSuffix() string`

GetCardSuffix returns the CardSuffix field if non-nil, zero value otherwise.

### GetCardSuffixOk

`func (o *GooglePayRequest) GetCardSuffixOk() (*string, bool)`

GetCardSuffixOk returns a tuple with the CardSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSuffix

`func (o *GooglePayRequest) SetCardSuffix(v string)`

SetCardSuffix sets CardSuffix field to given value.

### HasCardSuffix

`func (o *GooglePayRequest) HasCardSuffix() bool`

HasCardSuffix returns a boolean if a field has been set.

### SetCardSuffixNil

`func (o *GooglePayRequest) SetCardSuffixNil(b bool)`

 SetCardSuffixNil sets the value for CardSuffix to be an explicit nil

### UnsetCardSuffix
`func (o *GooglePayRequest) UnsetCardSuffix()`

UnsetCardSuffix ensures that no value is present for CardSuffix, not even an explicit nil
### GetCardScheme

`func (o *GooglePayRequest) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *GooglePayRequest) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *GooglePayRequest) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.

### HasCardScheme

`func (o *GooglePayRequest) HasCardScheme() bool`

HasCardScheme returns a boolean if a field has been set.

### SetCardSchemeNil

`func (o *GooglePayRequest) SetCardSchemeNil(b bool)`

 SetCardSchemeNil sets the value for CardScheme to be an explicit nil

### UnsetCardScheme
`func (o *GooglePayRequest) UnsetCardScheme()`

UnsetCardScheme ensures that no value is present for CardScheme, not even an explicit nil
### GetCardType

`func (o *GooglePayRequest) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *GooglePayRequest) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *GooglePayRequest) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *GooglePayRequest) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *GooglePayRequest) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *GooglePayRequest) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetAssuranceDetails

`func (o *GooglePayRequest) GetAssuranceDetails() GooglePayRequestAssuranceDetails`

GetAssuranceDetails returns the AssuranceDetails field if non-nil, zero value otherwise.

### GetAssuranceDetailsOk

`func (o *GooglePayRequest) GetAssuranceDetailsOk() (*GooglePayRequestAssuranceDetails, bool)`

GetAssuranceDetailsOk returns a tuple with the AssuranceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceDetails

`func (o *GooglePayRequest) SetAssuranceDetails(v GooglePayRequestAssuranceDetails)`

SetAssuranceDetails sets AssuranceDetails field to given value.

### HasAssuranceDetails

`func (o *GooglePayRequest) HasAssuranceDetails() bool`

HasAssuranceDetails returns a boolean if a field has been set.

### SetAssuranceDetailsNil

`func (o *GooglePayRequest) SetAssuranceDetailsNil(b bool)`

 SetAssuranceDetailsNil sets the value for AssuranceDetails to be an explicit nil

### UnsetAssuranceDetails
`func (o *GooglePayRequest) UnsetAssuranceDetails()`

UnsetAssuranceDetails ensures that no value is present for AssuranceDetails, not even an explicit nil
### GetCardholderName

`func (o *GooglePayRequest) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *GooglePayRequest) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *GooglePayRequest) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *GooglePayRequest) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *GooglePayRequest) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *GooglePayRequest) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
### GetRedirectUrl

`func (o *GooglePayRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *GooglePayRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *GooglePayRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *GooglePayRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *GooglePayRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *GooglePayRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


