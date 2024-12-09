# TokenizedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;id&#x60;. | 
**Id** | **string** | A ID that represents a previously stored payment method. This ID can represent any type of payment method. | 
**RedirectUrl** | Pointer to **string** | This value is mandatory for stored redirect payment methods. For stored cards, we strongly recommend providing a &#x60;redirect_url&#x60; either when 3-D Secure is enabled and &#x60;three_d_secure_data&#x60; is not provided, or when using connections where 3DS is enabled. This value will be appended with both a transaction ID and status (e.g. &#x60;https://example.com/callback?gr4vy_transaction_id&#x3D;123 &amp;gr4vy_transaction_status&#x3D;capture_succeeded&#x60;) after 3-D Secure has completed. For those cases, if the value is not present, the transaction will be marked as failed. | [optional] 
**SecurityCode** | Pointer to **NullableString** | The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD.  The security code can only be set if the stored payment method represents a card. | [optional] 

## Methods

### NewTokenizedRequest

`func NewTokenizedRequest(method string, id string, ) *TokenizedRequest`

NewTokenizedRequest instantiates a new TokenizedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedRequestWithDefaults

`func NewTokenizedRequestWithDefaults() *TokenizedRequest`

NewTokenizedRequestWithDefaults instantiates a new TokenizedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TokenizedRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TokenizedRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TokenizedRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetId

`func (o *TokenizedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenizedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenizedRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRedirectUrl

`func (o *TokenizedRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TokenizedRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TokenizedRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *TokenizedRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetSecurityCode

`func (o *TokenizedRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *TokenizedRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *TokenizedRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *TokenizedRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### SetSecurityCodeNil

`func (o *TokenizedRequest) SetSecurityCodeNil(b bool)`

 SetSecurityCodeNil sets the value for SecurityCode to be an explicit nil

### UnsetSecurityCode
`func (o *TokenizedRequest) UnsetSecurityCode()`

UnsetSecurityCode ensures that no value is present for SecurityCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


