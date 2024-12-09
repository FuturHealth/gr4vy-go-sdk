# TransactionRedirectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The method to use, this can be any of the methods that support redirect requests. | 
**RedirectUrl** | **string** | The redirect URL to redirect a buyer to after they have authorized their transaction. | 
**Currency** | **string** | The ISO-4217 currency code to use this payment method for. This is used to select the payment service to use. | 
**Country** | **string** | The 2-letter ISO code of the country to use this payment method for. This is used to select the payment service to use. | 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the account against your own records. This can only be set if the &#x60;store&#x60; flag is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewTransactionRedirectRequest

`func NewTransactionRedirectRequest(method string, redirectUrl string, currency string, country string, ) *TransactionRedirectRequest`

NewTransactionRedirectRequest instantiates a new TransactionRedirectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRedirectRequestWithDefaults

`func NewTransactionRedirectRequestWithDefaults() *TransactionRedirectRequest`

NewTransactionRedirectRequestWithDefaults instantiates a new TransactionRedirectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TransactionRedirectRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionRedirectRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionRedirectRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRedirectUrl

`func (o *TransactionRedirectRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TransactionRedirectRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TransactionRedirectRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetCurrency

`func (o *TransactionRedirectRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionRedirectRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionRedirectRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCountry

`func (o *TransactionRedirectRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TransactionRedirectRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TransactionRedirectRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetExternalIdentifier

`func (o *TransactionRedirectRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionRedirectRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionRedirectRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionRedirectRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionRedirectRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionRedirectRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


