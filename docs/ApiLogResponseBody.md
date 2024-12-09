# ApiLogResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The error code. | [optional] 
**Message** | Pointer to **string** | The error message. | [optional] 
**Status** | Pointer to **float32** | The HTTP error code. | [optional] 
**Type** | Pointer to **string** | Type of the log entry. | [optional] 
**Details** | Pointer to [**ApiLogResponseBodyDetails**](ApiLogResponseBodyDetails.md) |  | [optional] 

## Methods

### NewApiLogResponseBody

`func NewApiLogResponseBody() *ApiLogResponseBody`

NewApiLogResponseBody instantiates a new ApiLogResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLogResponseBodyWithDefaults

`func NewApiLogResponseBodyWithDefaults() *ApiLogResponseBody`

NewApiLogResponseBodyWithDefaults instantiates a new ApiLogResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiLogResponseBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiLogResponseBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiLogResponseBody) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiLogResponseBody) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ApiLogResponseBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiLogResponseBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiLogResponseBody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiLogResponseBody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ApiLogResponseBody) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiLogResponseBody) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiLogResponseBody) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiLogResponseBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ApiLogResponseBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiLogResponseBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiLogResponseBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiLogResponseBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetails

`func (o *ApiLogResponseBody) GetDetails() ApiLogResponseBodyDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApiLogResponseBody) GetDetailsOk() (*ApiLogResponseBodyDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApiLogResponseBody) SetDetails(v ApiLogResponseBodyDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApiLogResponseBody) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


