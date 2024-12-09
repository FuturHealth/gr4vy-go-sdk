# PaymentConnectorResponseTransactionVoidSucceededEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service used. | [optional] 
**PaymentServiceDisplayName** | Pointer to **string** | The display name of the payment service used. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The payment service definition used. | [optional] 
**PaymentServiceTransactionId** | Pointer to **NullableString** | The external ID of the transaction as set by the payment service. | [optional] 
**VoidedAt** | Pointer to **time.Time** | The date and time when this transaction was voided. | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionVoidSucceededEventContext

`func NewPaymentConnectorResponseTransactionVoidSucceededEventContext() *PaymentConnectorResponseTransactionVoidSucceededEventContext`

NewPaymentConnectorResponseTransactionVoidSucceededEventContext instantiates a new PaymentConnectorResponseTransactionVoidSucceededEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionVoidSucceededEventContextWithDefaults

`func NewPaymentConnectorResponseTransactionVoidSucceededEventContextWithDefaults() *PaymentConnectorResponseTransactionVoidSucceededEventContext`

NewPaymentConnectorResponseTransactionVoidSucceededEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionVoidSucceededEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.

### SetPaymentServiceTransactionIdNil

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) SetPaymentServiceTransactionIdNil(b bool)`

 SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil

### UnsetPaymentServiceTransactionId
`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) UnsetPaymentServiceTransactionId()`

UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
### GetVoidedAt

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *PaymentConnectorResponseTransactionVoidSucceededEventContext) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


