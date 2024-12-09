# NetworkTokenProvisionFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;network-token-provision-failed&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this event was created in our system. | [optional] 
**Context** | Pointer to [**NetworkTokenProvisionFailedContext**](NetworkTokenProvisionFailedContext.md) |  | [optional] 

## Methods

### NewNetworkTokenProvisionFailed

`func NewNetworkTokenProvisionFailed() *NetworkTokenProvisionFailed`

NewNetworkTokenProvisionFailed instantiates a new NetworkTokenProvisionFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTokenProvisionFailedWithDefaults

`func NewNetworkTokenProvisionFailedWithDefaults() *NetworkTokenProvisionFailed`

NewNetworkTokenProvisionFailedWithDefaults instantiates a new NetworkTokenProvisionFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkTokenProvisionFailed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTokenProvisionFailed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTokenProvisionFailed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkTokenProvisionFailed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *NetworkTokenProvisionFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTokenProvisionFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTokenProvisionFailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTokenProvisionFailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTokenProvisionFailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTokenProvisionFailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTokenProvisionFailed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTokenProvisionFailed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkTokenProvisionFailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTokenProvisionFailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTokenProvisionFailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkTokenProvisionFailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *NetworkTokenProvisionFailed) GetContext() NetworkTokenProvisionFailedContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NetworkTokenProvisionFailed) GetContextOk() (*NetworkTokenProvisionFailedContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NetworkTokenProvisionFailed) SetContext(v NetworkTokenProvisionFailedContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *NetworkTokenProvisionFailed) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


