# DigitalWalletClickToPayTokenDecryptedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | Correlation ID for transaction. | [optional] 
**MerchantTransactionId** | Pointer to **string** | Merchant Checkout Transaction Identifier which links the client-side JavaScript calls and server-side API calls for a specific transaction. | [optional] 
**Type** | Pointer to **string** | The type of payment instrument. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration of the card/token. | [optional] 
**HasCryptogram** | Pointer to **bool** | Online payment cryptogram, as defined by 3-D Secure. | [optional] 
**CardholderName** | Pointer to **NullableString** | The cardholder name. | [optional] 
**CardBin** | Pointer to **NullableString** | First six digits of underlying card. | [optional] 
**CardLastFour** | Pointer to **NullableString** | Last four digits of underlying card. | [optional] 
**CardExpirationDate** | Pointer to **NullableString** | Expiration date of underlying card. | [optional] 
**CardType** | Pointer to **NullableString** | Card type. | [optional] 
**BillingLine1** | Pointer to **NullableString** | Address line 1. | [optional] 
**BillingLine2** | Pointer to **NullableString** | Address line 2. | [optional] 
**BillingCity** | Pointer to **NullableString** | Address city. | [optional] 
**BillingState** | Pointer to **NullableString** | Address state. | [optional] 
**BillingZip** | Pointer to **NullableString** | Address zip/postal code. | [optional] 
**BillingCountryCode** | Pointer to **NullableString** | ISO 3166-1 alpha 2 address country code. | [optional] 

## Methods

### NewDigitalWalletClickToPayTokenDecryptedContext

`func NewDigitalWalletClickToPayTokenDecryptedContext() *DigitalWalletClickToPayTokenDecryptedContext`

NewDigitalWalletClickToPayTokenDecryptedContext instantiates a new DigitalWalletClickToPayTokenDecryptedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletClickToPayTokenDecryptedContextWithDefaults

`func NewDigitalWalletClickToPayTokenDecryptedContextWithDefaults() *DigitalWalletClickToPayTokenDecryptedContext`

NewDigitalWalletClickToPayTokenDecryptedContextWithDefaults instantiates a new DigitalWalletClickToPayTokenDecryptedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetMerchantTransactionId

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetMerchantTransactionId() string`

GetMerchantTransactionId returns the MerchantTransactionId field if non-nil, zero value otherwise.

### GetMerchantTransactionIdOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetMerchantTransactionIdOk() (*string, bool)`

GetMerchantTransactionIdOk returns a tuple with the MerchantTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransactionId

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetMerchantTransactionId(v string)`

SetMerchantTransactionId sets MerchantTransactionId field to given value.

### HasMerchantTransactionId

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasMerchantTransactionId() bool`

HasMerchantTransactionId returns a boolean if a field has been set.

### GetType

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpirationDate

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHasCryptogram

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetHasCryptogram() bool`

GetHasCryptogram returns the HasCryptogram field if non-nil, zero value otherwise.

### GetHasCryptogramOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetHasCryptogramOk() (*bool, bool)`

GetHasCryptogramOk returns a tuple with the HasCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCryptogram

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetHasCryptogram(v bool)`

SetHasCryptogram sets HasCryptogram field to given value.

### HasHasCryptogram

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasHasCryptogram() bool`

HasHasCryptogram returns a boolean if a field has been set.

### GetCardholderName

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
### GetCardBin

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardBin() string`

GetCardBin returns the CardBin field if non-nil, zero value otherwise.

### GetCardBinOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardBinOk() (*string, bool)`

GetCardBinOk returns a tuple with the CardBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBin

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardBin(v string)`

SetCardBin sets CardBin field to given value.

### HasCardBin

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardBin() bool`

HasCardBin returns a boolean if a field has been set.

### SetCardBinNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardBinNil(b bool)`

 SetCardBinNil sets the value for CardBin to be an explicit nil

### UnsetCardBin
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardBin()`

UnsetCardBin ensures that no value is present for CardBin, not even an explicit nil
### GetCardLastFour

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardLastFour() string`

GetCardLastFour returns the CardLastFour field if non-nil, zero value otherwise.

### GetCardLastFourOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardLastFourOk() (*string, bool)`

GetCardLastFourOk returns a tuple with the CardLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFour

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardLastFour(v string)`

SetCardLastFour sets CardLastFour field to given value.

### HasCardLastFour

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardLastFour() bool`

HasCardLastFour returns a boolean if a field has been set.

### SetCardLastFourNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardLastFourNil(b bool)`

 SetCardLastFourNil sets the value for CardLastFour to be an explicit nil

### UnsetCardLastFour
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardLastFour()`

UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
### GetCardExpirationDate

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardExpirationDate() string`

GetCardExpirationDate returns the CardExpirationDate field if non-nil, zero value otherwise.

### GetCardExpirationDateOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardExpirationDateOk() (*string, bool)`

GetCardExpirationDateOk returns a tuple with the CardExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationDate

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardExpirationDate(v string)`

SetCardExpirationDate sets CardExpirationDate field to given value.

### HasCardExpirationDate

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardExpirationDate() bool`

HasCardExpirationDate returns a boolean if a field has been set.

### SetCardExpirationDateNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardExpirationDateNil(b bool)`

 SetCardExpirationDateNil sets the value for CardExpirationDate to be an explicit nil

### UnsetCardExpirationDate
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardExpirationDate()`

UnsetCardExpirationDate ensures that no value is present for CardExpirationDate, not even an explicit nil
### GetCardType

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetBillingLine1

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine1() string`

GetBillingLine1 returns the BillingLine1 field if non-nil, zero value otherwise.

### GetBillingLine1Ok

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine1Ok() (*string, bool)`

GetBillingLine1Ok returns a tuple with the BillingLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLine1

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine1(v string)`

SetBillingLine1 sets BillingLine1 field to given value.

### HasBillingLine1

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingLine1() bool`

HasBillingLine1 returns a boolean if a field has been set.

### SetBillingLine1Nil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine1Nil(b bool)`

 SetBillingLine1Nil sets the value for BillingLine1 to be an explicit nil

### UnsetBillingLine1
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingLine1()`

UnsetBillingLine1 ensures that no value is present for BillingLine1, not even an explicit nil
### GetBillingLine2

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine2() string`

GetBillingLine2 returns the BillingLine2 field if non-nil, zero value otherwise.

### GetBillingLine2Ok

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine2Ok() (*string, bool)`

GetBillingLine2Ok returns a tuple with the BillingLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLine2

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine2(v string)`

SetBillingLine2 sets BillingLine2 field to given value.

### HasBillingLine2

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingLine2() bool`

HasBillingLine2 returns a boolean if a field has been set.

### SetBillingLine2Nil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine2Nil(b bool)`

 SetBillingLine2Nil sets the value for BillingLine2 to be an explicit nil

### UnsetBillingLine2
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingLine2()`

UnsetBillingLine2 ensures that no value is present for BillingLine2, not even an explicit nil
### GetBillingCity

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCity() string`

GetBillingCity returns the BillingCity field if non-nil, zero value otherwise.

### GetBillingCityOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCityOk() (*string, bool)`

GetBillingCityOk returns a tuple with the BillingCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCity

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCity(v string)`

SetBillingCity sets BillingCity field to given value.

### HasBillingCity

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingCity() bool`

HasBillingCity returns a boolean if a field has been set.

### SetBillingCityNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCityNil(b bool)`

 SetBillingCityNil sets the value for BillingCity to be an explicit nil

### UnsetBillingCity
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingCity()`

UnsetBillingCity ensures that no value is present for BillingCity, not even an explicit nil
### GetBillingState

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingState() string`

GetBillingState returns the BillingState field if non-nil, zero value otherwise.

### GetBillingStateOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingStateOk() (*string, bool)`

GetBillingStateOk returns a tuple with the BillingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingState

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingState(v string)`

SetBillingState sets BillingState field to given value.

### HasBillingState

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingState() bool`

HasBillingState returns a boolean if a field has been set.

### SetBillingStateNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingStateNil(b bool)`

 SetBillingStateNil sets the value for BillingState to be an explicit nil

### UnsetBillingState
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingState()`

UnsetBillingState ensures that no value is present for BillingState, not even an explicit nil
### GetBillingZip

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingZip() string`

GetBillingZip returns the BillingZip field if non-nil, zero value otherwise.

### GetBillingZipOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingZipOk() (*string, bool)`

GetBillingZipOk returns a tuple with the BillingZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingZip

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingZip(v string)`

SetBillingZip sets BillingZip field to given value.

### HasBillingZip

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingZip() bool`

HasBillingZip returns a boolean if a field has been set.

### SetBillingZipNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingZipNil(b bool)`

 SetBillingZipNil sets the value for BillingZip to be an explicit nil

### UnsetBillingZip
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingZip()`

UnsetBillingZip ensures that no value is present for BillingZip, not even an explicit nil
### GetBillingCountryCode

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCountryCode() string`

GetBillingCountryCode returns the BillingCountryCode field if non-nil, zero value otherwise.

### GetBillingCountryCodeOk

`func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCountryCodeOk() (*string, bool)`

GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountryCode

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCountryCode(v string)`

SetBillingCountryCode sets BillingCountryCode field to given value.

### HasBillingCountryCode

`func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingCountryCode() bool`

HasBillingCountryCode returns a boolean if a field has been set.

### SetBillingCountryCodeNil

`func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCountryCodeNil(b bool)`

 SetBillingCountryCodeNil sets the value for BillingCountryCode to be an explicit nil

### UnsetBillingCountryCode
`func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingCountryCode()`

UnsetBillingCountryCode ensures that no value is present for BillingCountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


