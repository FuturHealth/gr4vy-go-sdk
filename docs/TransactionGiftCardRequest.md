# TransactionGiftCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the gift card to check a balance for. Required if &#x60;number&#x60; is not set. | [optional] 
**Number** | Pointer to **string** | The 16-19 digit number for this gift card. Required if &#x60;id&#x60; is not set. | [optional] 
**Pin** | Pointer to **string** | The PIN for this gift card. Required if &#x60;number&#x60; is set. | [optional] 
**Amount** | **int32** | The monetary amount to charge for this gift card, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;.  All gift card amounts are subtracted from the total transaction amount. The remainder is charged to the provided &#x60;payment_method&#x60;. | 

## Methods

### NewTransactionGiftCardRequest

`func NewTransactionGiftCardRequest(amount int32, ) *TransactionGiftCardRequest`

NewTransactionGiftCardRequest instantiates a new TransactionGiftCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionGiftCardRequestWithDefaults

`func NewTransactionGiftCardRequestWithDefaults() *TransactionGiftCardRequest`

NewTransactionGiftCardRequestWithDefaults instantiates a new TransactionGiftCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionGiftCardRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionGiftCardRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionGiftCardRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionGiftCardRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *TransactionGiftCardRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TransactionGiftCardRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TransactionGiftCardRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TransactionGiftCardRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPin

`func (o *TransactionGiftCardRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *TransactionGiftCardRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *TransactionGiftCardRequest) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *TransactionGiftCardRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionGiftCardRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionGiftCardRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionGiftCardRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


