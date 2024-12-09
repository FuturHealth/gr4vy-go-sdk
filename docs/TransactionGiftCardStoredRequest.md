# TransactionGiftCardStoredRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the gift card to check a balance for. | 
**Amount** | **int32** | The monetary amount to charge for this gift card, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;.  All gift card amounts are subtracted from the total transaction amount. The remainder is charged to the provided &#x60;payment_method&#x60;. | 

## Methods

### NewTransactionGiftCardStoredRequest

`func NewTransactionGiftCardStoredRequest(id string, amount int32, ) *TransactionGiftCardStoredRequest`

NewTransactionGiftCardStoredRequest instantiates a new TransactionGiftCardStoredRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionGiftCardStoredRequestWithDefaults

`func NewTransactionGiftCardStoredRequestWithDefaults() *TransactionGiftCardStoredRequest`

NewTransactionGiftCardStoredRequestWithDefaults instantiates a new TransactionGiftCardStoredRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionGiftCardStoredRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionGiftCardStoredRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionGiftCardStoredRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *TransactionGiftCardStoredRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionGiftCardStoredRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionGiftCardStoredRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


