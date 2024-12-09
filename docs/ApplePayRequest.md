# ApplePayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;applepay&#x60;. | 
**Token** | **map[string]interface{}** | The encrypted (opaque) token that was passed to the &#x60;onpaymentauthorized&#x60; callback by the Apple Pay integration. | 
**CardSuffix** | Pointer to **NullableString** | Last 4 digits of the PAN for identification purposes. | [optional] 
**CardScheme** | Pointer to **NullableString** | The scheme/brand of the card. | [optional] 
**CardType** | Pointer to **NullableString** | The type of card. | [optional] 

## Methods

### NewApplePayRequest

`func NewApplePayRequest(method string, token map[string]interface{}, ) *ApplePayRequest`

NewApplePayRequest instantiates a new ApplePayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayRequestWithDefaults

`func NewApplePayRequestWithDefaults() *ApplePayRequest`

NewApplePayRequestWithDefaults instantiates a new ApplePayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ApplePayRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApplePayRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApplePayRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *ApplePayRequest) GetToken() map[string]interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApplePayRequest) GetTokenOk() (*map[string]interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApplePayRequest) SetToken(v map[string]interface{})`

SetToken sets Token field to given value.


### GetCardSuffix

`func (o *ApplePayRequest) GetCardSuffix() string`

GetCardSuffix returns the CardSuffix field if non-nil, zero value otherwise.

### GetCardSuffixOk

`func (o *ApplePayRequest) GetCardSuffixOk() (*string, bool)`

GetCardSuffixOk returns a tuple with the CardSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSuffix

`func (o *ApplePayRequest) SetCardSuffix(v string)`

SetCardSuffix sets CardSuffix field to given value.

### HasCardSuffix

`func (o *ApplePayRequest) HasCardSuffix() bool`

HasCardSuffix returns a boolean if a field has been set.

### SetCardSuffixNil

`func (o *ApplePayRequest) SetCardSuffixNil(b bool)`

 SetCardSuffixNil sets the value for CardSuffix to be an explicit nil

### UnsetCardSuffix
`func (o *ApplePayRequest) UnsetCardSuffix()`

UnsetCardSuffix ensures that no value is present for CardSuffix, not even an explicit nil
### GetCardScheme

`func (o *ApplePayRequest) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *ApplePayRequest) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *ApplePayRequest) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.

### HasCardScheme

`func (o *ApplePayRequest) HasCardScheme() bool`

HasCardScheme returns a boolean if a field has been set.

### SetCardSchemeNil

`func (o *ApplePayRequest) SetCardSchemeNil(b bool)`

 SetCardSchemeNil sets the value for CardScheme to be an explicit nil

### UnsetCardScheme
`func (o *ApplePayRequest) UnsetCardScheme()`

UnsetCardScheme ensures that no value is present for CardScheme, not even an explicit nil
### GetCardType

`func (o *ApplePayRequest) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *ApplePayRequest) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *ApplePayRequest) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *ApplePayRequest) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *ApplePayRequest) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *ApplePayRequest) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


