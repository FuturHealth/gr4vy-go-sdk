# TransactionRefundAllRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** | The reason to refund for. This can be used to attach a written reason to the refund request. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the refund against your own records. | [optional] 

## Methods

### NewTransactionRefundAllRequest

`func NewTransactionRefundAllRequest() *TransactionRefundAllRequest`

NewTransactionRefundAllRequest instantiates a new TransactionRefundAllRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRefundAllRequestWithDefaults

`func NewTransactionRefundAllRequestWithDefaults() *TransactionRefundAllRequest`

NewTransactionRefundAllRequestWithDefaults instantiates a new TransactionRefundAllRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *TransactionRefundAllRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransactionRefundAllRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransactionRefundAllRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransactionRefundAllRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *TransactionRefundAllRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TransactionRefundAllRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetExternalIdentifier

`func (o *TransactionRefundAllRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionRefundAllRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionRefundAllRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionRefundAllRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionRefundAllRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionRefundAllRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


