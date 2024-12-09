# PaymentConnectorResponseTransactionVoidFailedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;payment-connector-response-transaction-void-failed&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this event was created in our system. | [optional] 
**Context** | Pointer to [**PaymentConnectorResponseTransactionVoidFailedEventContext**](PaymentConnectorResponseTransactionVoidFailedEventContext.md) |  | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionVoidFailedEvent

`func NewPaymentConnectorResponseTransactionVoidFailedEvent() *PaymentConnectorResponseTransactionVoidFailedEvent`

NewPaymentConnectorResponseTransactionVoidFailedEvent instantiates a new PaymentConnectorResponseTransactionVoidFailedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionVoidFailedEventWithDefaults

`func NewPaymentConnectorResponseTransactionVoidFailedEventWithDefaults() *PaymentConnectorResponseTransactionVoidFailedEvent`

NewPaymentConnectorResponseTransactionVoidFailedEventWithDefaults instantiates a new PaymentConnectorResponseTransactionVoidFailedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetContext() PaymentConnectorResponseTransactionVoidFailedEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetContextOk() (*PaymentConnectorResponseTransactionVoidFailedEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetContext(v PaymentConnectorResponseTransactionVoidFailedEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


