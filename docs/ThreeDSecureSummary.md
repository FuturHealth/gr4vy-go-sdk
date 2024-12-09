# ThreeDSecureSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of 3DS used for this transaction. | [optional] 
**Status** | Pointer to **string** | The status of the 3DS challenge for this transaction. | [optional] 
**Method** | Pointer to **string** | The method used for 3DS authentication for this transaction. | [optional] 
**ErrorData** | Pointer to [**NullableThreeDSecureError**](ThreeDSecureError.md) | The error data received from our 3DS server. This will not be populated if the customer failed the authentication with a status code of &#x60;N&#x60;, &#x60;R&#x60;, or &#x60;U&#x60;.  To see full details about the 3DS calls in those situations please use our transaction events API. | [optional] 
**ResponseData** | Pointer to [**ThreeDSecureDataV1V2**](ThreeDSecureDataV1V2.md) |  | [optional] 

## Methods

### NewThreeDSecureSummary

`func NewThreeDSecureSummary() *ThreeDSecureSummary`

NewThreeDSecureSummary instantiates a new ThreeDSecureSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureSummaryWithDefaults

`func NewThreeDSecureSummaryWithDefaults() *ThreeDSecureSummary`

NewThreeDSecureSummaryWithDefaults instantiates a new ThreeDSecureSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ThreeDSecureSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureSummary) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSecureSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ThreeDSecureSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ThreeDSecureSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ThreeDSecureSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ThreeDSecureSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMethod

`func (o *ThreeDSecureSummary) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ThreeDSecureSummary) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ThreeDSecureSummary) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ThreeDSecureSummary) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetErrorData

`func (o *ThreeDSecureSummary) GetErrorData() ThreeDSecureError`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ThreeDSecureSummary) GetErrorDataOk() (*ThreeDSecureError, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ThreeDSecureSummary) SetErrorData(v ThreeDSecureError)`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ThreeDSecureSummary) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### SetErrorDataNil

`func (o *ThreeDSecureSummary) SetErrorDataNil(b bool)`

 SetErrorDataNil sets the value for ErrorData to be an explicit nil

### UnsetErrorData
`func (o *ThreeDSecureSummary) UnsetErrorData()`

UnsetErrorData ensures that no value is present for ErrorData, not even an explicit nil
### GetResponseData

`func (o *ThreeDSecureSummary) GetResponseData() ThreeDSecureDataV1V2`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *ThreeDSecureSummary) GetResponseDataOk() (*ThreeDSecureDataV1V2, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *ThreeDSecureSummary) SetResponseData(v ThreeDSecureDataV1V2)`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *ThreeDSecureSummary) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


