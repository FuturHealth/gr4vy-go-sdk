# RolePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to **[]string** | The list of permissions to allow for a role. | [optional] 

## Methods

### NewRolePermissions

`func NewRolePermissions() *RolePermissions`

NewRolePermissions instantiates a new RolePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsWithDefaults

`func NewRolePermissionsWithDefaults() *RolePermissions`

NewRolePermissionsWithDefaults instantiates a new RolePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *RolePermissions) GetAllow() []string`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *RolePermissions) GetAllowOk() (*[]string, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *RolePermissions) SetAllow(v []string)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *RolePermissions) HasAllow() bool`

HasAllow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


