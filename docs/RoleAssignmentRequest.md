# RoleAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**RoleAssignmentRequestRole**](RoleAssignmentRequestRole.md) |  | 
**Assignee** | [**RoleAssignmentRequestAssignee**](RoleAssignmentRequestAssignee.md) |  | 

## Methods

### NewRoleAssignmentRequest

`func NewRoleAssignmentRequest(role RoleAssignmentRequestRole, assignee RoleAssignmentRequestAssignee, ) *RoleAssignmentRequest`

NewRoleAssignmentRequest instantiates a new RoleAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentRequestWithDefaults

`func NewRoleAssignmentRequestWithDefaults() *RoleAssignmentRequest`

NewRoleAssignmentRequestWithDefaults instantiates a new RoleAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *RoleAssignmentRequest) GetRole() RoleAssignmentRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleAssignmentRequest) GetRoleOk() (*RoleAssignmentRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleAssignmentRequest) SetRole(v RoleAssignmentRequestRole)`

SetRole sets Role field to given value.


### GetAssignee

`func (o *RoleAssignmentRequest) GetAssignee() RoleAssignmentRequestAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *RoleAssignmentRequest) GetAssigneeOk() (*RoleAssignmentRequestAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *RoleAssignmentRequest) SetAssignee(v RoleAssignmentRequestAssignee)`

SetAssignee sets Assignee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


