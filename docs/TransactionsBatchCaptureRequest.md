# TransactionsBatchCaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The (partial) amount to capture.  When left blank, this will capture the entire amount. | 
**Currency** | **string** | A supported ISO-4217 currency code.  | 
**ExternalIdentifier** | Pointer to **string** | An external identifier that can be used to match the transaction against your own records. | [optional] 
**TransactionId** | **string** | The ID of the transaction to capture. | 

## Methods

### NewTransactionsBatchCaptureRequest

`func NewTransactionsBatchCaptureRequest(amount int32, currency string, transactionId string, ) *TransactionsBatchCaptureRequest`

NewTransactionsBatchCaptureRequest instantiates a new TransactionsBatchCaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsBatchCaptureRequestWithDefaults

`func NewTransactionsBatchCaptureRequestWithDefaults() *TransactionsBatchCaptureRequest`

NewTransactionsBatchCaptureRequestWithDefaults instantiates a new TransactionsBatchCaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionsBatchCaptureRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionsBatchCaptureRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionsBatchCaptureRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransactionsBatchCaptureRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionsBatchCaptureRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionsBatchCaptureRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExternalIdentifier

`func (o *TransactionsBatchCaptureRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionsBatchCaptureRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionsBatchCaptureRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionsBatchCaptureRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransactionsBatchCaptureRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionsBatchCaptureRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionsBatchCaptureRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


