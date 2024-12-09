# ErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | The location where the error caused an issue. | [optional] 
**Type** | Pointer to **string** | A unique identifier for the type of error that occurred. | [optional] 
**Pointer** | Pointer to **string** | The exact item for which the validation did not succeed. This is a JSON pointer for request bodies, while for query, path, and header parameters it is the name of the parameter. | [optional] 
**Message** | Pointer to **string** | A human readable message for this error detail. | [optional] 

## Methods

### NewErrorDetail

`func NewErrorDetail() *ErrorDetail`

NewErrorDetail instantiates a new ErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailWithDefaults

`func NewErrorDetailWithDefaults() *ErrorDetail`

NewErrorDetailWithDefaults instantiates a new ErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ErrorDetail) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ErrorDetail) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ErrorDetail) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ErrorDetail) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *ErrorDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPointer

`func (o *ErrorDetail) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorDetail) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorDetail) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ErrorDetail) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


