# TransactionNetworkTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;network-token&#x60;. | 
**Token** | **string** | The value of the network token. | 
**ExpirationDate** | **string** | The expiration date of the network token, formatted &#x60;MM/YY&#x60;. | 
**Cryptogram** | Pointer to **NullableString** | The cryptogram of the network token. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | We strongly recommend providing a &#x60;redirect_url&#x60; either when 3-D Secure is enabled and &#x60;three_d_secure_data&#x60; is not provided, or when using connections where 3DS is enabled. This value will be appended with both a transaction ID and status (e.g. &#x60;https://example.com/callback?gr4vy_transaction_id&#x3D;123 &amp;gr4vy_transaction_status&#x3D;capture_succeeded&#x60;) after 3-D Secure has completed. For those cases, if the value is not present, the transaction will be marked as failed. | [optional] 

## Methods

### NewTransactionNetworkTokenRequest

`func NewTransactionNetworkTokenRequest(method string, token string, expirationDate string, ) *TransactionNetworkTokenRequest`

NewTransactionNetworkTokenRequest instantiates a new TransactionNetworkTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionNetworkTokenRequestWithDefaults

`func NewTransactionNetworkTokenRequestWithDefaults() *TransactionNetworkTokenRequest`

NewTransactionNetworkTokenRequestWithDefaults instantiates a new TransactionNetworkTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TransactionNetworkTokenRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionNetworkTokenRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionNetworkTokenRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *TransactionNetworkTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TransactionNetworkTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TransactionNetworkTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpirationDate

`func (o *TransactionNetworkTokenRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TransactionNetworkTokenRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TransactionNetworkTokenRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetCryptogram

`func (o *TransactionNetworkTokenRequest) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *TransactionNetworkTokenRequest) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *TransactionNetworkTokenRequest) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.

### HasCryptogram

`func (o *TransactionNetworkTokenRequest) HasCryptogram() bool`

HasCryptogram returns a boolean if a field has been set.

### SetCryptogramNil

`func (o *TransactionNetworkTokenRequest) SetCryptogramNil(b bool)`

 SetCryptogramNil sets the value for Cryptogram to be an explicit nil

### UnsetCryptogram
`func (o *TransactionNetworkTokenRequest) UnsetCryptogram()`

UnsetCryptogram ensures that no value is present for Cryptogram, not even an explicit nil
### GetRedirectUrl

`func (o *TransactionNetworkTokenRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TransactionNetworkTokenRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TransactionNetworkTokenRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *TransactionNetworkTokenRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *TransactionNetworkTokenRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *TransactionNetworkTokenRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


