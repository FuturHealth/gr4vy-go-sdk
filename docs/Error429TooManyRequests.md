# Error429TooManyRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;error&#x60;. | [optional] 
**Code** | Pointer to **string** | &#x60;too_many_requests&#x60;. | [optional] 
**Status** | Pointer to **int32** | &#x60;429&#x60;. | [optional] 
**Message** | Pointer to **string** | Further details on the field that triggered the error. | [optional] 
**Details** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | A list of detail objects that further clarify the reason for the error. Not every error supports more detail. | [optional] 

## Methods

### NewError429TooManyRequests

`func NewError429TooManyRequests() *Error429TooManyRequests`

NewError429TooManyRequests instantiates a new Error429TooManyRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError429TooManyRequestsWithDefaults

`func NewError429TooManyRequestsWithDefaults() *Error429TooManyRequests`

NewError429TooManyRequestsWithDefaults instantiates a new Error429TooManyRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Error429TooManyRequests) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error429TooManyRequests) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error429TooManyRequests) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Error429TooManyRequests) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Error429TooManyRequests) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error429TooManyRequests) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error429TooManyRequests) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error429TooManyRequests) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *Error429TooManyRequests) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error429TooManyRequests) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error429TooManyRequests) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Error429TooManyRequests) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Error429TooManyRequests) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error429TooManyRequests) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error429TooManyRequests) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error429TooManyRequests) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *Error429TooManyRequests) GetDetails() []ErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Error429TooManyRequests) GetDetailsOk() (*[]ErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Error429TooManyRequests) SetDetails(v []ErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Error429TooManyRequests) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


