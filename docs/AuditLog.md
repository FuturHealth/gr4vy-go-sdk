# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;audit-log&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID of the audit log entry. | [optional] 
**Timestamp** | Pointer to **time.Time** | The date and time that the action was performed. | [optional] 
**Action** | Pointer to **string** | The action that was performed. | [optional] 
**MerchantAccountId** | Pointer to **NullableString** | The ID of the merchant account this entry was created for. | [optional] 
**User** | Pointer to [**AuditLogUser**](AuditLogUser.md) |  | [optional] 
**Resource** | Pointer to [**AuditLogResource**](AuditLogResource.md) |  | [optional] 

## Methods

### NewAuditLog

`func NewAuditLog() *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuditLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditLog) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AuditLog) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLog) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLog) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditLog) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAction

`func (o *AuditLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *AuditLog) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *AuditLog) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *AuditLog) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *AuditLog) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### SetMerchantAccountIdNil

`func (o *AuditLog) SetMerchantAccountIdNil(b bool)`

 SetMerchantAccountIdNil sets the value for MerchantAccountId to be an explicit nil

### UnsetMerchantAccountId
`func (o *AuditLog) UnsetMerchantAccountId()`

UnsetMerchantAccountId ensures that no value is present for MerchantAccountId, not even an explicit nil
### GetUser

`func (o *AuditLog) GetUser() AuditLogUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditLog) GetUserOk() (*AuditLogUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditLog) SetUser(v AuditLogUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditLog) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetResource

`func (o *AuditLog) GetResource() AuditLogResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AuditLog) GetResourceOk() (*AuditLogResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AuditLog) SetResource(v AuditLogResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AuditLog) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


