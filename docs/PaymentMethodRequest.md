# PaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The type of the funding source, e.g. &#x60;card&#x60;, &#x60;paypal&#x60;, or &#x60;checkout-session&#x60;. | 
**Id** | Pointer to **string** | The ID of a Checkout Session. | [optional] 
**Number** | Pointer to **string** | The 13-19 digit number for this credit card as it can be found on the front of the card. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date of the card, formatted &#x60;MM/YY&#x60;. If a card has been previously stored with us this value is optional. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the card against your own records. | [optional] 
**BuyerId** | Pointer to **string** | The ID of the buyer to associate this payment method to. If this field is provided then the &#x60;buyer_external_identifier&#x60; field needs to be unset. | [optional] 
**BuyerExternalIdentifier** | Pointer to **string** | The &#x60;external_identifier&#x60; of the buyer to associate this payment method to. If this field is provided then the &#x60;buyer_id&#x60; field needs to be unset. | [optional] 
**RedirectUrl** | Pointer to **string** | The redirect URL to redirect a buyer to after they have authorized their transaction or payment method. This only applies to payment methods that require buyer approval. | [optional] 
**Currency** | Pointer to **string** | The ISO-4217 currency code to store this payment method for. This is used to select the payment service to use.  This only applies to &#x60;redirect&#x60; mode payment methods like &#x60;gocardless&#x60;. | [optional] 
**Country** | Pointer to **string** | The 2-letter ISO code of the country to store this payment method for. This is used to select the payment service to use.  This only applies to &#x60;redirect&#x60; mode payment methods like &#x60;gocardless&#x60;. | [optional] 

## Methods

### NewPaymentMethodRequest

`func NewPaymentMethodRequest(method string, ) *PaymentMethodRequest`

NewPaymentMethodRequest instantiates a new PaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodRequestWithDefaults

`func NewPaymentMethodRequestWithDefaults() *PaymentMethodRequest`

NewPaymentMethodRequestWithDefaults instantiates a new PaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *PaymentMethodRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentMethodRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentMethodRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetId

`func (o *PaymentMethodRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *PaymentMethodRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PaymentMethodRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PaymentMethodRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PaymentMethodRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExpirationDate

`func (o *PaymentMethodRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethodRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethodRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethodRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *PaymentMethodRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PaymentMethodRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PaymentMethodRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PaymentMethodRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PaymentMethodRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PaymentMethodRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetBuyerId

`func (o *PaymentMethodRequest) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *PaymentMethodRequest) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *PaymentMethodRequest) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *PaymentMethodRequest) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetBuyerExternalIdentifier

`func (o *PaymentMethodRequest) GetBuyerExternalIdentifier() string`

GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field if non-nil, zero value otherwise.

### GetBuyerExternalIdentifierOk

`func (o *PaymentMethodRequest) GetBuyerExternalIdentifierOk() (*string, bool)`

GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerExternalIdentifier

`func (o *PaymentMethodRequest) SetBuyerExternalIdentifier(v string)`

SetBuyerExternalIdentifier sets BuyerExternalIdentifier field to given value.

### HasBuyerExternalIdentifier

`func (o *PaymentMethodRequest) HasBuyerExternalIdentifier() bool`

HasBuyerExternalIdentifier returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *PaymentMethodRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentMethodRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentMethodRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *PaymentMethodRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentMethodRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentMethodRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentMethodRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentMethodRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *PaymentMethodRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


