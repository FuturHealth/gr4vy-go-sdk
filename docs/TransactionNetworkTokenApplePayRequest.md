# TransactionNetworkTokenApplePayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;network-token&#x60;. | 
**Token** | **string** | The value of the decrypted Apple Pay token. | 
**ExpirationDate** | **string** | The expiration date of the network token, formatted &#x60;MM/YY&#x60;. | 
**Cryptogram** | Pointer to **NullableString** | The cryptogram of the network token. | [optional] 
**Eci** | Pointer to **string** | The electronic commerce indicator for 3D-Secure. | [optional] 
**CardSource** | **string** |  | 
**CardSuffix** | Pointer to **NullableString** | Last four digits of card PAN. | [optional] 
**CardScheme** | Pointer to **NullableString** | The scheme/brand of the card. | [optional] 
**CardholderName** | Pointer to **NullableString** | The cardholder name. | [optional] 

## Methods

### NewTransactionNetworkTokenApplePayRequest

`func NewTransactionNetworkTokenApplePayRequest(method string, token string, expirationDate string, cardSource string, ) *TransactionNetworkTokenApplePayRequest`

NewTransactionNetworkTokenApplePayRequest instantiates a new TransactionNetworkTokenApplePayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionNetworkTokenApplePayRequestWithDefaults

`func NewTransactionNetworkTokenApplePayRequestWithDefaults() *TransactionNetworkTokenApplePayRequest`

NewTransactionNetworkTokenApplePayRequestWithDefaults instantiates a new TransactionNetworkTokenApplePayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TransactionNetworkTokenApplePayRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionNetworkTokenApplePayRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionNetworkTokenApplePayRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *TransactionNetworkTokenApplePayRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TransactionNetworkTokenApplePayRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TransactionNetworkTokenApplePayRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpirationDate

`func (o *TransactionNetworkTokenApplePayRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TransactionNetworkTokenApplePayRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TransactionNetworkTokenApplePayRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetCryptogram

`func (o *TransactionNetworkTokenApplePayRequest) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *TransactionNetworkTokenApplePayRequest) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *TransactionNetworkTokenApplePayRequest) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.

### HasCryptogram

`func (o *TransactionNetworkTokenApplePayRequest) HasCryptogram() bool`

HasCryptogram returns a boolean if a field has been set.

### SetCryptogramNil

`func (o *TransactionNetworkTokenApplePayRequest) SetCryptogramNil(b bool)`

 SetCryptogramNil sets the value for Cryptogram to be an explicit nil

### UnsetCryptogram
`func (o *TransactionNetworkTokenApplePayRequest) UnsetCryptogram()`

UnsetCryptogram ensures that no value is present for Cryptogram, not even an explicit nil
### GetEci

`func (o *TransactionNetworkTokenApplePayRequest) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *TransactionNetworkTokenApplePayRequest) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *TransactionNetworkTokenApplePayRequest) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *TransactionNetworkTokenApplePayRequest) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetCardSource

`func (o *TransactionNetworkTokenApplePayRequest) GetCardSource() string`

GetCardSource returns the CardSource field if non-nil, zero value otherwise.

### GetCardSourceOk

`func (o *TransactionNetworkTokenApplePayRequest) GetCardSourceOk() (*string, bool)`

GetCardSourceOk returns a tuple with the CardSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSource

`func (o *TransactionNetworkTokenApplePayRequest) SetCardSource(v string)`

SetCardSource sets CardSource field to given value.


### GetCardSuffix

`func (o *TransactionNetworkTokenApplePayRequest) GetCardSuffix() string`

GetCardSuffix returns the CardSuffix field if non-nil, zero value otherwise.

### GetCardSuffixOk

`func (o *TransactionNetworkTokenApplePayRequest) GetCardSuffixOk() (*string, bool)`

GetCardSuffixOk returns a tuple with the CardSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSuffix

`func (o *TransactionNetworkTokenApplePayRequest) SetCardSuffix(v string)`

SetCardSuffix sets CardSuffix field to given value.

### HasCardSuffix

`func (o *TransactionNetworkTokenApplePayRequest) HasCardSuffix() bool`

HasCardSuffix returns a boolean if a field has been set.

### SetCardSuffixNil

`func (o *TransactionNetworkTokenApplePayRequest) SetCardSuffixNil(b bool)`

 SetCardSuffixNil sets the value for CardSuffix to be an explicit nil

### UnsetCardSuffix
`func (o *TransactionNetworkTokenApplePayRequest) UnsetCardSuffix()`

UnsetCardSuffix ensures that no value is present for CardSuffix, not even an explicit nil
### GetCardScheme

`func (o *TransactionNetworkTokenApplePayRequest) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *TransactionNetworkTokenApplePayRequest) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *TransactionNetworkTokenApplePayRequest) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.

### HasCardScheme

`func (o *TransactionNetworkTokenApplePayRequest) HasCardScheme() bool`

HasCardScheme returns a boolean if a field has been set.

### SetCardSchemeNil

`func (o *TransactionNetworkTokenApplePayRequest) SetCardSchemeNil(b bool)`

 SetCardSchemeNil sets the value for CardScheme to be an explicit nil

### UnsetCardScheme
`func (o *TransactionNetworkTokenApplePayRequest) UnsetCardScheme()`

UnsetCardScheme ensures that no value is present for CardScheme, not even an explicit nil
### GetCardholderName

`func (o *TransactionNetworkTokenApplePayRequest) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *TransactionNetworkTokenApplePayRequest) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *TransactionNetworkTokenApplePayRequest) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *TransactionNetworkTokenApplePayRequest) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *TransactionNetworkTokenApplePayRequest) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *TransactionNetworkTokenApplePayRequest) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


