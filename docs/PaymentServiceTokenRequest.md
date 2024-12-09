# PaymentServiceTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityCode** | Pointer to **NullableString** | The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD.  The security code can only be set if the stored payment method represents a card. | [optional] 
**PaymentServiceId** | **string** | The ID of the payment service. | 
**RedirectUrl** | **string** | The redirect URL to redirect a buyer to after they have authorized their payment method. This only applies to payment methods that require buyer approval. | 

## Methods

### NewPaymentServiceTokenRequest

`func NewPaymentServiceTokenRequest(paymentServiceId string, redirectUrl string, ) *PaymentServiceTokenRequest`

NewPaymentServiceTokenRequest instantiates a new PaymentServiceTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceTokenRequestWithDefaults

`func NewPaymentServiceTokenRequestWithDefaults() *PaymentServiceTokenRequest`

NewPaymentServiceTokenRequestWithDefaults instantiates a new PaymentServiceTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityCode

`func (o *PaymentServiceTokenRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *PaymentServiceTokenRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *PaymentServiceTokenRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *PaymentServiceTokenRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### SetSecurityCodeNil

`func (o *PaymentServiceTokenRequest) SetSecurityCodeNil(b bool)`

 SetSecurityCodeNil sets the value for SecurityCode to be an explicit nil

### UnsetSecurityCode
`func (o *PaymentServiceTokenRequest) UnsetSecurityCode()`

UnsetSecurityCode ensures that no value is present for SecurityCode, not even an explicit nil
### GetPaymentServiceId

`func (o *PaymentServiceTokenRequest) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentServiceTokenRequest) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentServiceTokenRequest) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.


### GetRedirectUrl

`func (o *PaymentServiceTokenRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentServiceTokenRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentServiceTokenRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


