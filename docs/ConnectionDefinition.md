# ConnectionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the connection. | [optional] 
**Type** | Pointer to **string** | &#x60;connection-definition&#x60;. | [optional] [default to "connection-definition"]
**Name** | Pointer to **string** | The name of this connection. | [optional] 
**Count** | Pointer to **float32** | The number of configured connections. | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **NullableString** | An icon to display for the connection. | [optional] 
**Provider** | Pointer to **NullableString** | The provider for this connection. | [optional] 

## Methods

### NewConnectionDefinition

`func NewConnectionDefinition() *ConnectionDefinition`

NewConnectionDefinition instantiates a new ConnectionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionDefinitionWithDefaults

`func NewConnectionDefinitionWithDefaults() *ConnectionDefinition`

NewConnectionDefinitionWithDefaults instantiates a new ConnectionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConnectionDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConnectionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCount

`func (o *ConnectionDefinition) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConnectionDefinition) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConnectionDefinition) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ConnectionDefinition) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGroup

`func (o *ConnectionDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ConnectionDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ConnectionDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ConnectionDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCategory

`func (o *ConnectionDefinition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConnectionDefinition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConnectionDefinition) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ConnectionDefinition) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIconUrl

`func (o *ConnectionDefinition) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConnectionDefinition) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConnectionDefinition) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConnectionDefinition) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ConnectionDefinition) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ConnectionDefinition) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetProvider

`func (o *ConnectionDefinition) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectionDefinition) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectionDefinition) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ConnectionDefinition) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *ConnectionDefinition) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *ConnectionDefinition) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


