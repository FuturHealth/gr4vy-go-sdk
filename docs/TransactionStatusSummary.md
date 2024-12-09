# TransactionStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this transaction. | [optional] 
**Status** | Pointer to **string** | The status of the transaction. The status may change over time as asynchronous processing events occur. | [optional] 

## Methods

### NewTransactionStatusSummary

`func NewTransactionStatusSummary() *TransactionStatusSummary`

NewTransactionStatusSummary instantiates a new TransactionStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStatusSummaryWithDefaults

`func NewTransactionStatusSummaryWithDefaults() *TransactionStatusSummary`

NewTransactionStatusSummaryWithDefaults instantiates a new TransactionStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionStatusSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionStatusSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionStatusSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionStatusSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TransactionStatusSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionStatusSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionStatusSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionStatusSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionStatusSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionStatusSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionStatusSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionStatusSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


