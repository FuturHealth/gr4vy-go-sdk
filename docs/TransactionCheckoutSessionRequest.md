# TransactionCheckoutSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;checkout-session&#x60;. | 
**Id** | **string** | The ID of the Checkout Session. | 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the card against your own records. This can only be set if the &#x60;store&#x60; flag is set to &#x60;true&#x60;. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | We strongly recommend providing a &#x60;redirect_url&#x60; either when 3-D Secure is enabled and &#x60;three_d_secure_data&#x60; is not provided, or when using connections where 3DS is enabled. This value will be appended with both a transaction ID and status (e.g. &#x60;https://example.com/callback?gr4vy_transaction_id&#x3D;123 &amp;gr4vy_transaction_status&#x3D;capture_succeeded&#x60;) after 3-D Secure has completed. For those cases, if the value is not present, the transaction will be marked as failed. | [optional] 

## Methods

### NewTransactionCheckoutSessionRequest

`func NewTransactionCheckoutSessionRequest(method string, id string, ) *TransactionCheckoutSessionRequest`

NewTransactionCheckoutSessionRequest instantiates a new TransactionCheckoutSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCheckoutSessionRequestWithDefaults

`func NewTransactionCheckoutSessionRequestWithDefaults() *TransactionCheckoutSessionRequest`

NewTransactionCheckoutSessionRequestWithDefaults instantiates a new TransactionCheckoutSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TransactionCheckoutSessionRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionCheckoutSessionRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionCheckoutSessionRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetId

`func (o *TransactionCheckoutSessionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionCheckoutSessionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionCheckoutSessionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetExternalIdentifier

`func (o *TransactionCheckoutSessionRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionCheckoutSessionRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionCheckoutSessionRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionCheckoutSessionRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionCheckoutSessionRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionCheckoutSessionRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetRedirectUrl

`func (o *TransactionCheckoutSessionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TransactionCheckoutSessionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TransactionCheckoutSessionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *TransactionCheckoutSessionRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *TransactionCheckoutSessionRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *TransactionCheckoutSessionRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


