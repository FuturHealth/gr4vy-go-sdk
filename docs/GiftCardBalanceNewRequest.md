# GiftCardBalanceNewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The 16-19 digit number for this gift card. | 
**Pin** | **string** | The PIN for this gift card. | 

## Methods

### NewGiftCardBalanceNewRequest

`func NewGiftCardBalanceNewRequest(number string, pin string, ) *GiftCardBalanceNewRequest`

NewGiftCardBalanceNewRequest instantiates a new GiftCardBalanceNewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardBalanceNewRequestWithDefaults

`func NewGiftCardBalanceNewRequestWithDefaults() *GiftCardBalanceNewRequest`

NewGiftCardBalanceNewRequestWithDefaults instantiates a new GiftCardBalanceNewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GiftCardBalanceNewRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GiftCardBalanceNewRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GiftCardBalanceNewRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetPin

`func (o *GiftCardBalanceNewRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *GiftCardBalanceNewRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *GiftCardBalanceNewRequest) SetPin(v string)`

SetPin sets Pin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


