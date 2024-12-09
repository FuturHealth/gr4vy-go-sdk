# UserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The full name of the user which is used in the Gr4vy admin panel to give a user a human readable name. | [optional] 
**RoleIds** | Pointer to **[]string** | The IDs of the roles to assign to the user. Sending an empty list will remove all roles assigned to the user. | [optional] 
**MerchantAccountIds** | Pointer to **[]string** | A list of merchant account IDs that the user is assigned to. | [optional] 

## Methods

### NewUserUpdate

`func NewUserUpdate() *UserUpdate`

NewUserUpdate instantiates a new UserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateWithDefaults

`func NewUserUpdateWithDefaults() *UserUpdate`

NewUserUpdateWithDefaults instantiates a new UserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleIds

`func (o *UserUpdate) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UserUpdate) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UserUpdate) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UserUpdate) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetMerchantAccountIds

`func (o *UserUpdate) GetMerchantAccountIds() []string`

GetMerchantAccountIds returns the MerchantAccountIds field if non-nil, zero value otherwise.

### GetMerchantAccountIdsOk

`func (o *UserUpdate) GetMerchantAccountIdsOk() (*[]string, bool)`

GetMerchantAccountIdsOk returns a tuple with the MerchantAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountIds

`func (o *UserUpdate) SetMerchantAccountIds(v []string)`

SetMerchantAccountIds sets MerchantAccountIds field to given value.

### HasMerchantAccountIds

`func (o *UserUpdate) HasMerchantAccountIds() bool`

HasMerchantAccountIds returns a boolean if a field has been set.

### SetMerchantAccountIdsNil

`func (o *UserUpdate) SetMerchantAccountIdsNil(b bool)`

 SetMerchantAccountIdsNil sets the value for MerchantAccountIds to be an explicit nil

### UnsetMerchantAccountIds
`func (o *UserUpdate) UnsetMerchantAccountIds()`

UnsetMerchantAccountIds ensures that no value is present for MerchantAccountIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


