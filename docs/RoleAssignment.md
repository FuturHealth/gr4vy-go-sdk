# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Always &#x60;role-assignment&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID for this role assignment. | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**Assignee** | Pointer to [**RoleAssignmentAssignee**](RoleAssignmentAssignee.md) |  | [optional] 

## Methods

### NewRoleAssignment

`func NewRoleAssignment() *RoleAssignment`

NewRoleAssignment instantiates a new RoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentWithDefaults

`func NewRoleAssignmentWithDefaults() *RoleAssignment`

NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoleAssignment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleAssignment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleAssignment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleAssignment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *RoleAssignment) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleAssignment) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleAssignment) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAssignee

`func (o *RoleAssignment) GetAssignee() RoleAssignmentAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *RoleAssignment) GetAssigneeOk() (*RoleAssignmentAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *RoleAssignment) SetAssignee(v RoleAssignmentAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *RoleAssignment) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


