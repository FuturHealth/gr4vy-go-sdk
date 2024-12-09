# Refund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;refund&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the refund. | [optional] 
**TransactionId** | Pointer to **string** | The ID of the transaction associated with this refund. | [optional] 
**PaymentServiceRefundId** | Pointer to **string** | The payment service&#39;s unique ID for the refund. | [optional] 
**Status** | Pointer to **string** | The status of the refund. It may change over time as asynchronous processing events occur.  - &#x60;processing&#x60; - The refund is being processed. - &#x60;succeeded&#x60; - The refund was successful. - &#x60;declined&#x60; - The refund was declined by the underlying PSP. - &#x60;failed&#x60; - The refund could not proceed due to a technical issue. - &#x60;voided&#x60; - The refund was voided and will not proceed. | [optional] 
**Currency** | Pointer to **string** | The currency code for this refund. Will always match that of the associated transaction. | [optional] 
**Amount** | Pointer to **int32** | The amount requested for this refund. | [optional] 
**Reason** | Pointer to **NullableString** | The reason for this refund. Could be a multiline string. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this refund was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this refund was last updated. | [optional] 
**TargetType** | Pointer to **string** | The type of the instrument that was refunded. | [optional] 
**TargetId** | Pointer to **NullableString** | The optional ID of the instrument that was refunded. This may be &#x60;null&#x60; if the instrument was not stored. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the refund against your own records. | [optional] 
**ReconciliationId** | Pointer to **string** | The base62 encoded refund ID. This represents a shorter version of this refund&#39;s &#x60;id&#x60; which is sent to payment services, anti-fraud services, and other connectors. You can use this ID to reconcile a payment service&#39;s refund against our system. | [optional] 
**TransactionExternalIdentifier** | Pointer to **NullableString** | The external identifier of the related transaction. | [optional] 
**TransactionReconciliationId** | Pointer to **string** | The base62 encoded transaction ID. This represents a shorter version of the related transaction&#39;s &#x60;id&#x60; which is sent to payment services, anti-fraud services, and other connectors. You can use this ID to reconcile a payment service&#39;s transaction against our system. | [optional] 

## Methods

### NewRefund

`func NewRefund() *Refund`

NewRefund instantiates a new Refund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundWithDefaults

`func NewRefundWithDefaults() *Refund`

NewRefundWithDefaults instantiates a new Refund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Refund) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Refund) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Refund) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Refund) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Refund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Refund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Refund) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Refund) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTransactionId

`func (o *Refund) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Refund) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Refund) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Refund) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetPaymentServiceRefundId

`func (o *Refund) GetPaymentServiceRefundId() string`

GetPaymentServiceRefundId returns the PaymentServiceRefundId field if non-nil, zero value otherwise.

### GetPaymentServiceRefundIdOk

`func (o *Refund) GetPaymentServiceRefundIdOk() (*string, bool)`

GetPaymentServiceRefundIdOk returns a tuple with the PaymentServiceRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceRefundId

`func (o *Refund) SetPaymentServiceRefundId(v string)`

SetPaymentServiceRefundId sets PaymentServiceRefundId field to given value.

### HasPaymentServiceRefundId

`func (o *Refund) HasPaymentServiceRefundId() bool`

HasPaymentServiceRefundId returns a boolean if a field has been set.

### GetStatus

`func (o *Refund) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Refund) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Refund) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Refund) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *Refund) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Refund) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Refund) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Refund) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *Refund) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Refund) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Refund) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Refund) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetReason

`func (o *Refund) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Refund) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Refund) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Refund) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *Refund) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Refund) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetCreatedAt

`func (o *Refund) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Refund) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Refund) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Refund) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Refund) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Refund) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Refund) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Refund) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTargetType

`func (o *Refund) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Refund) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Refund) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Refund) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetId

`func (o *Refund) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Refund) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Refund) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Refund) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *Refund) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *Refund) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetExternalIdentifier

`func (o *Refund) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *Refund) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *Refund) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *Refund) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *Refund) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *Refund) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetReconciliationId

`func (o *Refund) GetReconciliationId() string`

GetReconciliationId returns the ReconciliationId field if non-nil, zero value otherwise.

### GetReconciliationIdOk

`func (o *Refund) GetReconciliationIdOk() (*string, bool)`

GetReconciliationIdOk returns a tuple with the ReconciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationId

`func (o *Refund) SetReconciliationId(v string)`

SetReconciliationId sets ReconciliationId field to given value.

### HasReconciliationId

`func (o *Refund) HasReconciliationId() bool`

HasReconciliationId returns a boolean if a field has been set.

### GetTransactionExternalIdentifier

`func (o *Refund) GetTransactionExternalIdentifier() string`

GetTransactionExternalIdentifier returns the TransactionExternalIdentifier field if non-nil, zero value otherwise.

### GetTransactionExternalIdentifierOk

`func (o *Refund) GetTransactionExternalIdentifierOk() (*string, bool)`

GetTransactionExternalIdentifierOk returns a tuple with the TransactionExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionExternalIdentifier

`func (o *Refund) SetTransactionExternalIdentifier(v string)`

SetTransactionExternalIdentifier sets TransactionExternalIdentifier field to given value.

### HasTransactionExternalIdentifier

`func (o *Refund) HasTransactionExternalIdentifier() bool`

HasTransactionExternalIdentifier returns a boolean if a field has been set.

### SetTransactionExternalIdentifierNil

`func (o *Refund) SetTransactionExternalIdentifierNil(b bool)`

 SetTransactionExternalIdentifierNil sets the value for TransactionExternalIdentifier to be an explicit nil

### UnsetTransactionExternalIdentifier
`func (o *Refund) UnsetTransactionExternalIdentifier()`

UnsetTransactionExternalIdentifier ensures that no value is present for TransactionExternalIdentifier, not even an explicit nil
### GetTransactionReconciliationId

`func (o *Refund) GetTransactionReconciliationId() string`

GetTransactionReconciliationId returns the TransactionReconciliationId field if non-nil, zero value otherwise.

### GetTransactionReconciliationIdOk

`func (o *Refund) GetTransactionReconciliationIdOk() (*string, bool)`

GetTransactionReconciliationIdOk returns a tuple with the TransactionReconciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReconciliationId

`func (o *Refund) SetTransactionReconciliationId(v string)`

SetTransactionReconciliationId sets TransactionReconciliationId field to given value.

### HasTransactionReconciliationId

`func (o *Refund) HasTransactionReconciliationId() bool`

HasTransactionReconciliationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


