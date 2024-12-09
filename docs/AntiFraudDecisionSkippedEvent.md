# AntiFraudDecisionSkippedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;anti-fraud-decision-skipped&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Context** | Pointer to [**AntiFraudDecisionSkippedEventContext**](AntiFraudDecisionSkippedEventContext.md) |  | [optional] 

## Methods

### NewAntiFraudDecisionSkippedEvent

`func NewAntiFraudDecisionSkippedEvent() *AntiFraudDecisionSkippedEvent`

NewAntiFraudDecisionSkippedEvent instantiates a new AntiFraudDecisionSkippedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudDecisionSkippedEventWithDefaults

`func NewAntiFraudDecisionSkippedEventWithDefaults() *AntiFraudDecisionSkippedEvent`

NewAntiFraudDecisionSkippedEventWithDefaults instantiates a new AntiFraudDecisionSkippedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AntiFraudDecisionSkippedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AntiFraudDecisionSkippedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AntiFraudDecisionSkippedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AntiFraudDecisionSkippedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AntiFraudDecisionSkippedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntiFraudDecisionSkippedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntiFraudDecisionSkippedEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AntiFraudDecisionSkippedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AntiFraudDecisionSkippedEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AntiFraudDecisionSkippedEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AntiFraudDecisionSkippedEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AntiFraudDecisionSkippedEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AntiFraudDecisionSkippedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AntiFraudDecisionSkippedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AntiFraudDecisionSkippedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AntiFraudDecisionSkippedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *AntiFraudDecisionSkippedEvent) GetContext() AntiFraudDecisionSkippedEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AntiFraudDecisionSkippedEvent) GetContextOk() (*AntiFraudDecisionSkippedEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AntiFraudDecisionSkippedEvent) SetContext(v AntiFraudDecisionSkippedEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AntiFraudDecisionSkippedEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


