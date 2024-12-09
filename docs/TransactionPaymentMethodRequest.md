# TransactionPaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The method to use for this request. | 
**Number** | Pointer to **NullableString** | The 13-19 digit number for this credit card as it can be found on the front of the card. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date of the card, formatted &#x60;MM/YY&#x60;. If a card has been previously stored with us this value is optional. | [optional] 
**SecurityCode** | Pointer to **NullableString** | The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the card against your own records. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The redirect URL to redirect a buyer to after they have authorized their transaction or payment method. This only applies to payment methods that require buyer approval. | [optional] 
**Id** | Pointer to **NullableString** | An identifier for a previously tokenized payment method or checkout-session. This id can represent any type of payment method or checkout-session. | [optional] 
**Currency** | Pointer to **NullableString** | The ISO-4217 currency code to use this payment method for. This is used to select the payment service to use. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country to use this payment method for. This is used to select the payment service to use. | [optional] 
**Token** | Pointer to **map[string]interface{}** | The encrypted (opaque) token that was passed to the &#x60;onpaymentauthorized&#x60; callback by the Apple Pay integration. | [optional] 
**AssuranceDetails** | Pointer to [**NullableGooglePayRequestAssuranceDetails**](GooglePayRequestAssuranceDetails.md) |  | [optional] 
**CardholderName** | Pointer to **NullableString** | Name of the card holder. | [optional] 
**Cryptogram** | Pointer to **NullableString** | The cryptogram of the network token. | [optional] 

## Methods

### NewTransactionPaymentMethodRequest

`func NewTransactionPaymentMethodRequest(method string, ) *TransactionPaymentMethodRequest`

NewTransactionPaymentMethodRequest instantiates a new TransactionPaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPaymentMethodRequestWithDefaults

`func NewTransactionPaymentMethodRequestWithDefaults() *TransactionPaymentMethodRequest`

NewTransactionPaymentMethodRequestWithDefaults instantiates a new TransactionPaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TransactionPaymentMethodRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionPaymentMethodRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionPaymentMethodRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNumber

`func (o *TransactionPaymentMethodRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TransactionPaymentMethodRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TransactionPaymentMethodRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TransactionPaymentMethodRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *TransactionPaymentMethodRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *TransactionPaymentMethodRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetExpirationDate

`func (o *TransactionPaymentMethodRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TransactionPaymentMethodRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TransactionPaymentMethodRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TransactionPaymentMethodRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *TransactionPaymentMethodRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *TransactionPaymentMethodRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetSecurityCode

`func (o *TransactionPaymentMethodRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *TransactionPaymentMethodRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *TransactionPaymentMethodRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *TransactionPaymentMethodRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### SetSecurityCodeNil

`func (o *TransactionPaymentMethodRequest) SetSecurityCodeNil(b bool)`

 SetSecurityCodeNil sets the value for SecurityCode to be an explicit nil

### UnsetSecurityCode
`func (o *TransactionPaymentMethodRequest) UnsetSecurityCode()`

UnsetSecurityCode ensures that no value is present for SecurityCode, not even an explicit nil
### GetExternalIdentifier

`func (o *TransactionPaymentMethodRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionPaymentMethodRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionPaymentMethodRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionPaymentMethodRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionPaymentMethodRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionPaymentMethodRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetRedirectUrl

`func (o *TransactionPaymentMethodRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TransactionPaymentMethodRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TransactionPaymentMethodRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *TransactionPaymentMethodRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *TransactionPaymentMethodRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *TransactionPaymentMethodRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetId

`func (o *TransactionPaymentMethodRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionPaymentMethodRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionPaymentMethodRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionPaymentMethodRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TransactionPaymentMethodRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TransactionPaymentMethodRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCurrency

`func (o *TransactionPaymentMethodRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionPaymentMethodRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionPaymentMethodRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionPaymentMethodRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *TransactionPaymentMethodRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *TransactionPaymentMethodRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCountry

`func (o *TransactionPaymentMethodRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TransactionPaymentMethodRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TransactionPaymentMethodRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TransactionPaymentMethodRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *TransactionPaymentMethodRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *TransactionPaymentMethodRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetToken

`func (o *TransactionPaymentMethodRequest) GetToken() map[string]interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TransactionPaymentMethodRequest) GetTokenOk() (*map[string]interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TransactionPaymentMethodRequest) SetToken(v map[string]interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *TransactionPaymentMethodRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *TransactionPaymentMethodRequest) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *TransactionPaymentMethodRequest) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetAssuranceDetails

`func (o *TransactionPaymentMethodRequest) GetAssuranceDetails() GooglePayRequestAssuranceDetails`

GetAssuranceDetails returns the AssuranceDetails field if non-nil, zero value otherwise.

### GetAssuranceDetailsOk

`func (o *TransactionPaymentMethodRequest) GetAssuranceDetailsOk() (*GooglePayRequestAssuranceDetails, bool)`

GetAssuranceDetailsOk returns a tuple with the AssuranceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceDetails

`func (o *TransactionPaymentMethodRequest) SetAssuranceDetails(v GooglePayRequestAssuranceDetails)`

SetAssuranceDetails sets AssuranceDetails field to given value.

### HasAssuranceDetails

`func (o *TransactionPaymentMethodRequest) HasAssuranceDetails() bool`

HasAssuranceDetails returns a boolean if a field has been set.

### SetAssuranceDetailsNil

`func (o *TransactionPaymentMethodRequest) SetAssuranceDetailsNil(b bool)`

 SetAssuranceDetailsNil sets the value for AssuranceDetails to be an explicit nil

### UnsetAssuranceDetails
`func (o *TransactionPaymentMethodRequest) UnsetAssuranceDetails()`

UnsetAssuranceDetails ensures that no value is present for AssuranceDetails, not even an explicit nil
### GetCardholderName

`func (o *TransactionPaymentMethodRequest) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *TransactionPaymentMethodRequest) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *TransactionPaymentMethodRequest) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *TransactionPaymentMethodRequest) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *TransactionPaymentMethodRequest) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *TransactionPaymentMethodRequest) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
### GetCryptogram

`func (o *TransactionPaymentMethodRequest) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *TransactionPaymentMethodRequest) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *TransactionPaymentMethodRequest) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.

### HasCryptogram

`func (o *TransactionPaymentMethodRequest) HasCryptogram() bool`

HasCryptogram returns a boolean if a field has been set.

### SetCryptogramNil

`func (o *TransactionPaymentMethodRequest) SetCryptogramNil(b bool)`

 SetCryptogramNil sets the value for Cryptogram to be an explicit nil

### UnsetCryptogram
`func (o *TransactionPaymentMethodRequest) UnsetCryptogram()`

UnsetCryptogram ensures that no value is present for Cryptogram, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


