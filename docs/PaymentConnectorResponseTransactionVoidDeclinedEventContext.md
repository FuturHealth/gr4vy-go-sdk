# PaymentConnectorResponseTransactionVoidDeclinedEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service used. | [optional] 
**PaymentServiceDisplayName** | Pointer to **string** | The display name of the payment service used. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The payment service definition used. | [optional] 
**PaymentServiceTransactionId** | Pointer to **NullableString** | The external ID of the transaction as set by the payment service. | [optional] 
**Code** | Pointer to **NullableString** | A raw response code returned for the failure. | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionVoidDeclinedEventContext

`func NewPaymentConnectorResponseTransactionVoidDeclinedEventContext() *PaymentConnectorResponseTransactionVoidDeclinedEventContext`

NewPaymentConnectorResponseTransactionVoidDeclinedEventContext instantiates a new PaymentConnectorResponseTransactionVoidDeclinedEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionVoidDeclinedEventContextWithDefaults

`func NewPaymentConnectorResponseTransactionVoidDeclinedEventContextWithDefaults() *PaymentConnectorResponseTransactionVoidDeclinedEventContext`

NewPaymentConnectorResponseTransactionVoidDeclinedEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionVoidDeclinedEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.

### SetPaymentServiceTransactionIdNil

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetPaymentServiceTransactionIdNil(b bool)`

 SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil

### UnsetPaymentServiceTransactionId
`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) UnsetPaymentServiceTransactionId()`

UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
### GetCode

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PaymentConnectorResponseTransactionVoidDeclinedEventContext) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


