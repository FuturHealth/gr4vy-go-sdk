# SetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResetToken** | Pointer to **string** | Unique reset token valid for 7 days. | [optional] 
**Password** | Pointer to **string** | The password the user to log in with. | [optional] 

## Methods

### NewSetPasswordRequest

`func NewSetPasswordRequest() *SetPasswordRequest`

NewSetPasswordRequest instantiates a new SetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPasswordRequestWithDefaults

`func NewSetPasswordRequestWithDefaults() *SetPasswordRequest`

NewSetPasswordRequestWithDefaults instantiates a new SetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResetToken

`func (o *SetPasswordRequest) GetResetToken() string`

GetResetToken returns the ResetToken field if non-nil, zero value otherwise.

### GetResetTokenOk

`func (o *SetPasswordRequest) GetResetTokenOk() (*string, bool)`

GetResetTokenOk returns a tuple with the ResetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetToken

`func (o *SetPasswordRequest) SetResetToken(v string)`

SetResetToken sets ResetToken field to given value.

### HasResetToken

`func (o *SetPasswordRequest) HasResetToken() bool`

HasResetToken returns a boolean if a field has been set.

### GetPassword

`func (o *SetPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SetPasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


