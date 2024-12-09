# UserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The full name of the user which is used in the Gr4vy admin panel to give an user a human readable name. | [optional] 
**EmailAddress** | Pointer to **string** | The email address for this user. | [optional] 
**RoleIds** | Pointer to **[]string** | A list of role ids that will be assigned to the user being created. The creator must have &#x60;roles.write&#x60; or the role that is being assigned. | [optional] 
**MerchantAccountIds** | Pointer to **[]string** | A list of merchant account IDs that the user being created will be assigned to. | [optional] 

## Methods

### NewUserRequest

`func NewUserRequest() *UserRequest`

NewUserRequest instantiates a new UserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestWithDefaults

`func NewUserRequestWithDefaults() *UserRequest`

NewUserRequestWithDefaults instantiates a new UserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *UserRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UserRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetRoleIds

`func (o *UserRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UserRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UserRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UserRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetMerchantAccountIds

`func (o *UserRequest) GetMerchantAccountIds() []string`

GetMerchantAccountIds returns the MerchantAccountIds field if non-nil, zero value otherwise.

### GetMerchantAccountIdsOk

`func (o *UserRequest) GetMerchantAccountIdsOk() (*[]string, bool)`

GetMerchantAccountIdsOk returns a tuple with the MerchantAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountIds

`func (o *UserRequest) SetMerchantAccountIds(v []string)`

SetMerchantAccountIds sets MerchantAccountIds field to given value.

### HasMerchantAccountIds

`func (o *UserRequest) HasMerchantAccountIds() bool`

HasMerchantAccountIds returns a boolean if a field has been set.

### SetMerchantAccountIdsNil

`func (o *UserRequest) SetMerchantAccountIdsNil(b bool)`

 SetMerchantAccountIdsNil sets the value for MerchantAccountIds to be an explicit nil

### UnsetMerchantAccountIds
`func (o *UserRequest) UnsetMerchantAccountIds()`

UnsetMerchantAccountIds ensures that no value is present for MerchantAccountIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


