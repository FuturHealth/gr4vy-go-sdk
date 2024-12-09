# CheckoutSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;checkout-session&#x60;. | 
**Id** | **string** | The ID of the Checkout Session. | 
**RedirectUrl** | Pointer to **NullableString** | The redirect URL to redirect a buyer to after they have authorized their transaction or payment method. This only applies to payment methods that require buyer approval. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the card against your own records. | [optional] 
**BuyerId** | Pointer to **string** | The ID of the buyer to associate this payment method to. If this field is provided then the &#x60;buyer_external_identifier&#x60; field needs to be unset. | [optional] 
**BuyerExternalIdentifier** | Pointer to **string** | The &#x60;external_identifier&#x60; of the buyer to associate this payment method to. If this field is provided then the &#x60;buyer_id&#x60; field needs to be unset. | [optional] 

## Methods

### NewCheckoutSessionRequest

`func NewCheckoutSessionRequest(method string, id string, ) *CheckoutSessionRequest`

NewCheckoutSessionRequest instantiates a new CheckoutSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionRequestWithDefaults

`func NewCheckoutSessionRequestWithDefaults() *CheckoutSessionRequest`

NewCheckoutSessionRequestWithDefaults instantiates a new CheckoutSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CheckoutSessionRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckoutSessionRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckoutSessionRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetId

`func (o *CheckoutSessionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSessionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSessionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRedirectUrl

`func (o *CheckoutSessionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CheckoutSessionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CheckoutSessionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CheckoutSessionRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CheckoutSessionRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CheckoutSessionRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetExternalIdentifier

`func (o *CheckoutSessionRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *CheckoutSessionRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *CheckoutSessionRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *CheckoutSessionRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *CheckoutSessionRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *CheckoutSessionRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetBuyerId

`func (o *CheckoutSessionRequest) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *CheckoutSessionRequest) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *CheckoutSessionRequest) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *CheckoutSessionRequest) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetBuyerExternalIdentifier

`func (o *CheckoutSessionRequest) GetBuyerExternalIdentifier() string`

GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field if non-nil, zero value otherwise.

### GetBuyerExternalIdentifierOk

`func (o *CheckoutSessionRequest) GetBuyerExternalIdentifierOk() (*string, bool)`

GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerExternalIdentifier

`func (o *CheckoutSessionRequest) SetBuyerExternalIdentifier(v string)`

SetBuyerExternalIdentifier sets BuyerExternalIdentifier field to given value.

### HasBuyerExternalIdentifier

`func (o *CheckoutSessionRequest) HasBuyerExternalIdentifier() bool`

HasBuyerExternalIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


