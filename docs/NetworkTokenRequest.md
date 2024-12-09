# NetworkTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityCode** | Pointer to **NullableString** | The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD.  The security code can only be set if the stored payment method represents a card. | [optional] 
**MerchantInitiated** | **bool** | Defines if the request is merchant initiated or not. | 
**IsSubsequentPayment** | **bool** | Defines if the request is a subsequent of another request or not. | 

## Methods

### NewNetworkTokenRequest

`func NewNetworkTokenRequest(merchantInitiated bool, isSubsequentPayment bool, ) *NetworkTokenRequest`

NewNetworkTokenRequest instantiates a new NetworkTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTokenRequestWithDefaults

`func NewNetworkTokenRequestWithDefaults() *NetworkTokenRequest`

NewNetworkTokenRequestWithDefaults instantiates a new NetworkTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityCode

`func (o *NetworkTokenRequest) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *NetworkTokenRequest) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *NetworkTokenRequest) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *NetworkTokenRequest) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### SetSecurityCodeNil

`func (o *NetworkTokenRequest) SetSecurityCodeNil(b bool)`

 SetSecurityCodeNil sets the value for SecurityCode to be an explicit nil

### UnsetSecurityCode
`func (o *NetworkTokenRequest) UnsetSecurityCode()`

UnsetSecurityCode ensures that no value is present for SecurityCode, not even an explicit nil
### GetMerchantInitiated

`func (o *NetworkTokenRequest) GetMerchantInitiated() bool`

GetMerchantInitiated returns the MerchantInitiated field if non-nil, zero value otherwise.

### GetMerchantInitiatedOk

`func (o *NetworkTokenRequest) GetMerchantInitiatedOk() (*bool, bool)`

GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInitiated

`func (o *NetworkTokenRequest) SetMerchantInitiated(v bool)`

SetMerchantInitiated sets MerchantInitiated field to given value.


### GetIsSubsequentPayment

`func (o *NetworkTokenRequest) GetIsSubsequentPayment() bool`

GetIsSubsequentPayment returns the IsSubsequentPayment field if non-nil, zero value otherwise.

### GetIsSubsequentPaymentOk

`func (o *NetworkTokenRequest) GetIsSubsequentPaymentOk() (*bool, bool)`

GetIsSubsequentPaymentOk returns a tuple with the IsSubsequentPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubsequentPayment

`func (o *NetworkTokenRequest) SetIsSubsequentPayment(v bool)`

SetIsSubsequentPayment sets IsSubsequentPayment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


