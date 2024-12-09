# NetworkTokenProvisionSucceeded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;network-token-provision-succeeded&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this event was created in our system. | [optional] 
**Context** | Pointer to [**NetworkTokenCryptogramProvisionSucceededContext**](NetworkTokenCryptogramProvisionSucceededContext.md) |  | [optional] 

## Methods

### NewNetworkTokenProvisionSucceeded

`func NewNetworkTokenProvisionSucceeded() *NetworkTokenProvisionSucceeded`

NewNetworkTokenProvisionSucceeded instantiates a new NetworkTokenProvisionSucceeded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTokenProvisionSucceededWithDefaults

`func NewNetworkTokenProvisionSucceededWithDefaults() *NetworkTokenProvisionSucceeded`

NewNetworkTokenProvisionSucceededWithDefaults instantiates a new NetworkTokenProvisionSucceeded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkTokenProvisionSucceeded) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTokenProvisionSucceeded) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTokenProvisionSucceeded) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkTokenProvisionSucceeded) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *NetworkTokenProvisionSucceeded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTokenProvisionSucceeded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTokenProvisionSucceeded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTokenProvisionSucceeded) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTokenProvisionSucceeded) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTokenProvisionSucceeded) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTokenProvisionSucceeded) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTokenProvisionSucceeded) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkTokenProvisionSucceeded) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTokenProvisionSucceeded) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTokenProvisionSucceeded) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkTokenProvisionSucceeded) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *NetworkTokenProvisionSucceeded) GetContext() NetworkTokenCryptogramProvisionSucceededContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NetworkTokenProvisionSucceeded) GetContextOk() (*NetworkTokenCryptogramProvisionSucceededContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NetworkTokenProvisionSucceeded) SetContext(v NetworkTokenCryptogramProvisionSucceededContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *NetworkTokenProvisionSucceeded) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


