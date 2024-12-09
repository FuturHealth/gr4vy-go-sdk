# GiftCardStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The 16-19 digit number for this gift card. | 
**Pin** | **string** | The PIN for this gift card. | 
**BuyerId** | Pointer to **string** | The ID of the buyer to associate this gift card to. If this field is provided then the &#x60;buyer_external_identifier&#x60; field needs to be unset. | [optional] 
**BuyerExternalIdentifier** | Pointer to **string** | The &#x60;external_identifier&#x60; of the buyer to associate this gift card to. If this field is provided then the &#x60;buyer_id&#x60; field needs to be unset. | [optional] 

## Methods

### NewGiftCardStoreRequest

`func NewGiftCardStoreRequest(number string, pin string, ) *GiftCardStoreRequest`

NewGiftCardStoreRequest instantiates a new GiftCardStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardStoreRequestWithDefaults

`func NewGiftCardStoreRequestWithDefaults() *GiftCardStoreRequest`

NewGiftCardStoreRequestWithDefaults instantiates a new GiftCardStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GiftCardStoreRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GiftCardStoreRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GiftCardStoreRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetPin

`func (o *GiftCardStoreRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *GiftCardStoreRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *GiftCardStoreRequest) SetPin(v string)`

SetPin sets Pin field to given value.


### GetBuyerId

`func (o *GiftCardStoreRequest) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *GiftCardStoreRequest) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *GiftCardStoreRequest) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *GiftCardStoreRequest) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetBuyerExternalIdentifier

`func (o *GiftCardStoreRequest) GetBuyerExternalIdentifier() string`

GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field if non-nil, zero value otherwise.

### GetBuyerExternalIdentifierOk

`func (o *GiftCardStoreRequest) GetBuyerExternalIdentifierOk() (*string, bool)`

GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerExternalIdentifier

`func (o *GiftCardStoreRequest) SetBuyerExternalIdentifier(v string)`

SetBuyerExternalIdentifier sets BuyerExternalIdentifier field to given value.

### HasBuyerExternalIdentifier

`func (o *GiftCardStoreRequest) HasBuyerExternalIdentifier() bool`

HasBuyerExternalIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


