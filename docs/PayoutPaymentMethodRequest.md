# PayoutPaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The method to use. | 
**Id** | Pointer to **NullableString** | The ID of the stored payment method. | [optional] 
**Number** | Pointer to **NullableString** | The 13-19 digit number for this card as it can be found on the front of the card. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date of the card, formatted &#x60;MM/YY&#x60;. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the card against your own records. | [optional] 
**Scheme** | Pointer to **NullableString** | The card scheme. | [optional] 

## Methods

### NewPayoutPaymentMethodRequest

`func NewPayoutPaymentMethodRequest(method string, ) *PayoutPaymentMethodRequest`

NewPayoutPaymentMethodRequest instantiates a new PayoutPaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutPaymentMethodRequestWithDefaults

`func NewPayoutPaymentMethodRequestWithDefaults() *PayoutPaymentMethodRequest`

NewPayoutPaymentMethodRequestWithDefaults instantiates a new PayoutPaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *PayoutPaymentMethodRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PayoutPaymentMethodRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PayoutPaymentMethodRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetId

`func (o *PayoutPaymentMethodRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutPaymentMethodRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutPaymentMethodRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayoutPaymentMethodRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PayoutPaymentMethodRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PayoutPaymentMethodRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNumber

`func (o *PayoutPaymentMethodRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PayoutPaymentMethodRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PayoutPaymentMethodRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PayoutPaymentMethodRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *PayoutPaymentMethodRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *PayoutPaymentMethodRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetExpirationDate

`func (o *PayoutPaymentMethodRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PayoutPaymentMethodRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PayoutPaymentMethodRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PayoutPaymentMethodRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PayoutPaymentMethodRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PayoutPaymentMethodRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetExternalIdentifier

`func (o *PayoutPaymentMethodRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PayoutPaymentMethodRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PayoutPaymentMethodRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PayoutPaymentMethodRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PayoutPaymentMethodRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PayoutPaymentMethodRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetScheme

`func (o *PayoutPaymentMethodRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PayoutPaymentMethodRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PayoutPaymentMethodRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PayoutPaymentMethodRequest) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PayoutPaymentMethodRequest) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PayoutPaymentMethodRequest) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


