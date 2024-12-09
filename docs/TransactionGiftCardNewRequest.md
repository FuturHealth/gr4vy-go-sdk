# TransactionGiftCardNewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The 16-19 digit number for this gift card. | 
**Pin** | **string** | The PIN for this gift card. | 
**Amount** | **int32** | The monetary amount to charge for this gift card, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;.  All gift card amounts are subtracted from the total transaction amount. The remainder is charged to the provided &#x60;payment_method&#x60;. | 

## Methods

### NewTransactionGiftCardNewRequest

`func NewTransactionGiftCardNewRequest(number string, pin string, amount int32, ) *TransactionGiftCardNewRequest`

NewTransactionGiftCardNewRequest instantiates a new TransactionGiftCardNewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionGiftCardNewRequestWithDefaults

`func NewTransactionGiftCardNewRequestWithDefaults() *TransactionGiftCardNewRequest`

NewTransactionGiftCardNewRequestWithDefaults instantiates a new TransactionGiftCardNewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *TransactionGiftCardNewRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TransactionGiftCardNewRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TransactionGiftCardNewRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetPin

`func (o *TransactionGiftCardNewRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *TransactionGiftCardNewRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *TransactionGiftCardNewRequest) SetPin(v string)`

SetPin sets Pin field to given value.


### GetAmount

`func (o *TransactionGiftCardNewRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionGiftCardNewRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionGiftCardNewRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


