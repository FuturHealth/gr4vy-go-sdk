# ThreeDSecurePreparationRequestEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;three-d-secure-preparation-request&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Context** | Pointer to [**ThreeDSecurePreparationRequestEventContext**](ThreeDSecurePreparationRequestEventContext.md) |  | [optional] 

## Methods

### NewThreeDSecurePreparationRequestEvent

`func NewThreeDSecurePreparationRequestEvent() *ThreeDSecurePreparationRequestEvent`

NewThreeDSecurePreparationRequestEvent instantiates a new ThreeDSecurePreparationRequestEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecurePreparationRequestEventWithDefaults

`func NewThreeDSecurePreparationRequestEventWithDefaults() *ThreeDSecurePreparationRequestEvent`

NewThreeDSecurePreparationRequestEventWithDefaults instantiates a new ThreeDSecurePreparationRequestEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThreeDSecurePreparationRequestEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreeDSecurePreparationRequestEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreeDSecurePreparationRequestEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreeDSecurePreparationRequestEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ThreeDSecurePreparationRequestEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreeDSecurePreparationRequestEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreeDSecurePreparationRequestEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreeDSecurePreparationRequestEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ThreeDSecurePreparationRequestEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreeDSecurePreparationRequestEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreeDSecurePreparationRequestEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreeDSecurePreparationRequestEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ThreeDSecurePreparationRequestEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ThreeDSecurePreparationRequestEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ThreeDSecurePreparationRequestEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ThreeDSecurePreparationRequestEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *ThreeDSecurePreparationRequestEvent) GetContext() ThreeDSecurePreparationRequestEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ThreeDSecurePreparationRequestEvent) GetContextOk() (*ThreeDSecurePreparationRequestEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ThreeDSecurePreparationRequestEvent) SetContext(v ThreeDSecurePreparationRequestEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ThreeDSecurePreparationRequestEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


