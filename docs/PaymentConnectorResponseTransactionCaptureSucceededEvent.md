# PaymentConnectorResponseTransactionCaptureSucceededEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;payment-connector-response-transaction-capture-succeeded&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Context** | Pointer to [**PaymentConnectorResponseTransactionCaptureSucceededEventContext**](PaymentConnectorResponseTransactionCaptureSucceededEventContext.md) |  | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionCaptureSucceededEvent

`func NewPaymentConnectorResponseTransactionCaptureSucceededEvent() *PaymentConnectorResponseTransactionCaptureSucceededEvent`

NewPaymentConnectorResponseTransactionCaptureSucceededEvent instantiates a new PaymentConnectorResponseTransactionCaptureSucceededEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionCaptureSucceededEventWithDefaults

`func NewPaymentConnectorResponseTransactionCaptureSucceededEventWithDefaults() *PaymentConnectorResponseTransactionCaptureSucceededEvent`

NewPaymentConnectorResponseTransactionCaptureSucceededEventWithDefaults instantiates a new PaymentConnectorResponseTransactionCaptureSucceededEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetContext() PaymentConnectorResponseTransactionCaptureSucceededEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) GetContextOk() (*PaymentConnectorResponseTransactionCaptureSucceededEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) SetContext(v PaymentConnectorResponseTransactionCaptureSucceededEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


