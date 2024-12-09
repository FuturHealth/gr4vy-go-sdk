# Error401Unauthorized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;error&#x60;. | [optional] 
**Code** | Pointer to **string** | &#x60;unauthorized&#x60;. | [optional] 
**Status** | Pointer to **int32** | &#x60;401&#x60;. | [optional] 
**Message** | Pointer to **string** | No valid API authentication found. | [optional] 
**Details** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | A list of detail objects that further clarify the reason for the error. Not every error supports more detail. | [optional] 

## Methods

### NewError401Unauthorized

`func NewError401Unauthorized() *Error401Unauthorized`

NewError401Unauthorized instantiates a new Error401Unauthorized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError401UnauthorizedWithDefaults

`func NewError401UnauthorizedWithDefaults() *Error401Unauthorized`

NewError401UnauthorizedWithDefaults instantiates a new Error401Unauthorized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Error401Unauthorized) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error401Unauthorized) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error401Unauthorized) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Error401Unauthorized) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Error401Unauthorized) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error401Unauthorized) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error401Unauthorized) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error401Unauthorized) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *Error401Unauthorized) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error401Unauthorized) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error401Unauthorized) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Error401Unauthorized) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Error401Unauthorized) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error401Unauthorized) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error401Unauthorized) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error401Unauthorized) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *Error401Unauthorized) GetDetails() []ErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Error401Unauthorized) GetDetailsOk() (*[]ErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Error401Unauthorized) SetDetails(v []ErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Error401Unauthorized) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


