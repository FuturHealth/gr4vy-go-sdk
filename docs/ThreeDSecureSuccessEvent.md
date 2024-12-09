# ThreeDSecureSuccessEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;three-d-secure-success&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Context** | Pointer to [**ThreeDSecureSuccessEventContext**](ThreeDSecureSuccessEventContext.md) |  | [optional] 

## Methods

### NewThreeDSecureSuccessEvent

`func NewThreeDSecureSuccessEvent() *ThreeDSecureSuccessEvent`

NewThreeDSecureSuccessEvent instantiates a new ThreeDSecureSuccessEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureSuccessEventWithDefaults

`func NewThreeDSecureSuccessEventWithDefaults() *ThreeDSecureSuccessEvent`

NewThreeDSecureSuccessEventWithDefaults instantiates a new ThreeDSecureSuccessEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThreeDSecureSuccessEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreeDSecureSuccessEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreeDSecureSuccessEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreeDSecureSuccessEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ThreeDSecureSuccessEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreeDSecureSuccessEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreeDSecureSuccessEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreeDSecureSuccessEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ThreeDSecureSuccessEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreeDSecureSuccessEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreeDSecureSuccessEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreeDSecureSuccessEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ThreeDSecureSuccessEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ThreeDSecureSuccessEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ThreeDSecureSuccessEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ThreeDSecureSuccessEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *ThreeDSecureSuccessEvent) GetContext() ThreeDSecureSuccessEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ThreeDSecureSuccessEvent) GetContextOk() (*ThreeDSecureSuccessEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ThreeDSecureSuccessEvent) SetContext(v ThreeDSecureSuccessEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ThreeDSecureSuccessEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


