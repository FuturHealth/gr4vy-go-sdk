# PayoutInlineCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The method to use. | 
**Number** | **string** | The 13-19 digit number for this card as it can be found on the front of the card. | 
**ExpirationDate** | **string** | The expiration date of the card, formatted &#x60;MM/YY&#x60;. | 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the card against your own records. | [optional] 
**Scheme** | Pointer to **NullableString** | The card scheme. | [optional] 

## Methods

### NewPayoutInlineCardRequest

`func NewPayoutInlineCardRequest(method string, number string, expirationDate string, ) *PayoutInlineCardRequest`

NewPayoutInlineCardRequest instantiates a new PayoutInlineCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutInlineCardRequestWithDefaults

`func NewPayoutInlineCardRequestWithDefaults() *PayoutInlineCardRequest`

NewPayoutInlineCardRequestWithDefaults instantiates a new PayoutInlineCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *PayoutInlineCardRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PayoutInlineCardRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PayoutInlineCardRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNumber

`func (o *PayoutInlineCardRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PayoutInlineCardRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PayoutInlineCardRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetExpirationDate

`func (o *PayoutInlineCardRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PayoutInlineCardRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PayoutInlineCardRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetExternalIdentifier

`func (o *PayoutInlineCardRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PayoutInlineCardRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PayoutInlineCardRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PayoutInlineCardRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PayoutInlineCardRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PayoutInlineCardRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetScheme

`func (o *PayoutInlineCardRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PayoutInlineCardRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PayoutInlineCardRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PayoutInlineCardRequest) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PayoutInlineCardRequest) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PayoutInlineCardRequest) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


