# UserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Always &#x60;role&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID for this role. | [optional] 
**Name** | Pointer to **string** | The unique name for this role. | [optional] 
**Description** | Pointer to **string** | The description for this role. | [optional] 
**Permissions** | Pointer to [**RolePermissions**](RolePermissions.md) |  | [optional] 

## Methods

### NewUserRole

`func NewUserRole() *UserRole`

NewUserRole instantiates a new UserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleWithDefaults

`func NewUserRoleWithDefaults() *UserRole`

NewUserRoleWithDefaults instantiates a new UserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *UserRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UserRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *UserRole) GetPermissions() RolePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserRole) GetPermissionsOk() (*RolePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserRole) SetPermissions(v RolePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


