# Error403Forbidden

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;error&#x60;. | [optional] 
**Code** | Pointer to **string** | &#x60;forbidden&#x60;. | [optional] 
**Status** | Pointer to **int32** | &#x60;403&#x60;. | [optional] 
**Message** | Pointer to **string** | Invalid credentials. | [optional] 
**Details** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | A list of detail objects that further clarify the reason for the error. Not every error supports more detail. | [optional] 

## Methods

### NewError403Forbidden

`func NewError403Forbidden() *Error403Forbidden`

NewError403Forbidden instantiates a new Error403Forbidden object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError403ForbiddenWithDefaults

`func NewError403ForbiddenWithDefaults() *Error403Forbidden`

NewError403ForbiddenWithDefaults instantiates a new Error403Forbidden object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Error403Forbidden) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error403Forbidden) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error403Forbidden) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Error403Forbidden) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Error403Forbidden) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error403Forbidden) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error403Forbidden) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error403Forbidden) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *Error403Forbidden) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error403Forbidden) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error403Forbidden) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Error403Forbidden) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Error403Forbidden) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error403Forbidden) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error403Forbidden) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error403Forbidden) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *Error403Forbidden) GetDetails() []ErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Error403Forbidden) GetDetailsOk() (*[]ErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Error403Forbidden) SetDetails(v []ErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Error403Forbidden) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


