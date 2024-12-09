# GiftCardBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the gift card to check a balance for. Required if &#x60;number&#x60; is not set. | [optional] 
**Number** | Pointer to **string** | The 16-19 digit number for this gift card. Required if &#x60;id&#x60; is not set. | [optional] 
**Pin** | Pointer to **string** | The PIN for this gift card. Required if &#x60;number&#x60; is set. | [optional] 

## Methods

### NewGiftCardBalanceRequest

`func NewGiftCardBalanceRequest() *GiftCardBalanceRequest`

NewGiftCardBalanceRequest instantiates a new GiftCardBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardBalanceRequestWithDefaults

`func NewGiftCardBalanceRequestWithDefaults() *GiftCardBalanceRequest`

NewGiftCardBalanceRequestWithDefaults instantiates a new GiftCardBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GiftCardBalanceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardBalanceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardBalanceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardBalanceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *GiftCardBalanceRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GiftCardBalanceRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GiftCardBalanceRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GiftCardBalanceRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPin

`func (o *GiftCardBalanceRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *GiftCardBalanceRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *GiftCardBalanceRequest) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *GiftCardBalanceRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


