# AntiFraudDecisionUpdateErrorEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;anti-fraud-decision-update-error&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction event was created in our system. | [optional] 
**Context** | Pointer to [**AntiFraudDecisionUpdateErrorEventContext**](AntiFraudDecisionUpdateErrorEventContext.md) |  | [optional] 

## Methods

### NewAntiFraudDecisionUpdateErrorEvent

`func NewAntiFraudDecisionUpdateErrorEvent() *AntiFraudDecisionUpdateErrorEvent`

NewAntiFraudDecisionUpdateErrorEvent instantiates a new AntiFraudDecisionUpdateErrorEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudDecisionUpdateErrorEventWithDefaults

`func NewAntiFraudDecisionUpdateErrorEventWithDefaults() *AntiFraudDecisionUpdateErrorEvent`

NewAntiFraudDecisionUpdateErrorEventWithDefaults instantiates a new AntiFraudDecisionUpdateErrorEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AntiFraudDecisionUpdateErrorEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AntiFraudDecisionUpdateErrorEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AntiFraudDecisionUpdateErrorEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AntiFraudDecisionUpdateErrorEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AntiFraudDecisionUpdateErrorEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntiFraudDecisionUpdateErrorEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntiFraudDecisionUpdateErrorEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AntiFraudDecisionUpdateErrorEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AntiFraudDecisionUpdateErrorEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AntiFraudDecisionUpdateErrorEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AntiFraudDecisionUpdateErrorEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AntiFraudDecisionUpdateErrorEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AntiFraudDecisionUpdateErrorEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AntiFraudDecisionUpdateErrorEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AntiFraudDecisionUpdateErrorEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AntiFraudDecisionUpdateErrorEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *AntiFraudDecisionUpdateErrorEvent) GetContext() AntiFraudDecisionUpdateErrorEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AntiFraudDecisionUpdateErrorEvent) GetContextOk() (*AntiFraudDecisionUpdateErrorEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AntiFraudDecisionUpdateErrorEvent) SetContext(v AntiFraudDecisionUpdateErrorEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AntiFraudDecisionUpdateErrorEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


