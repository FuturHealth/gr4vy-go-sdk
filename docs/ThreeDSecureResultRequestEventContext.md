# ThreeDSecureResultRequestEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL that was called for this request. | [optional] 
**Request** | Pointer to **string** | The request body sent to the &#x60;url&#x60;. | [optional] 
**Response** | Pointer to **string** | The response body received from the &#x60;url&#x60;. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | The response status code received from the &#x60;url&#x60;. | [optional] 
**Cavv** | Pointer to **string** | The 3DS CAVV value parsed from the &#x60;response&#x60;. | [optional] 
**Eci** | Pointer to **string** | The 3DS ECI value parsed from the &#x60;response&#x60;. | [optional] 
**AuthenticationResponse** | Pointer to **string** | The &#x60;transStatus&#x60; parsed from the post-authorization &#x60;response&#x60;. | [optional] 
**DirectoryResponse** | Pointer to **string** | The &#x60;transStatus&#x60; parsed from the authorization &#x60;response&#x60;. | [optional] 
**DirectoryTransactionId** | Pointer to **string** | The &#x60;dsTransID&#x60; parsed from the &#x60;response&#x60;. | [optional] 
**Version** | Pointer to **string** | The version of 3DS used. | [optional] 

## Methods

### NewThreeDSecureResultRequestEventContext

`func NewThreeDSecureResultRequestEventContext() *ThreeDSecureResultRequestEventContext`

NewThreeDSecureResultRequestEventContext instantiates a new ThreeDSecureResultRequestEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureResultRequestEventContextWithDefaults

`func NewThreeDSecureResultRequestEventContextWithDefaults() *ThreeDSecureResultRequestEventContext`

NewThreeDSecureResultRequestEventContextWithDefaults instantiates a new ThreeDSecureResultRequestEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ThreeDSecureResultRequestEventContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThreeDSecureResultRequestEventContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThreeDSecureResultRequestEventContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThreeDSecureResultRequestEventContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *ThreeDSecureResultRequestEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ThreeDSecureResultRequestEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ThreeDSecureResultRequestEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ThreeDSecureResultRequestEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *ThreeDSecureResultRequestEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ThreeDSecureResultRequestEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ThreeDSecureResultRequestEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ThreeDSecureResultRequestEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *ThreeDSecureResultRequestEventContext) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *ThreeDSecureResultRequestEventContext) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *ThreeDSecureResultRequestEventContext) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *ThreeDSecureResultRequestEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetCavv

`func (o *ThreeDSecureResultRequestEventContext) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureResultRequestEventContext) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureResultRequestEventContext) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ThreeDSecureResultRequestEventContext) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetEci

`func (o *ThreeDSecureResultRequestEventContext) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureResultRequestEventContext) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureResultRequestEventContext) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDSecureResultRequestEventContext) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetAuthenticationResponse

`func (o *ThreeDSecureResultRequestEventContext) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureResultRequestEventContext) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureResultRequestEventContext) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.

### HasAuthenticationResponse

`func (o *ThreeDSecureResultRequestEventContext) HasAuthenticationResponse() bool`

HasAuthenticationResponse returns a boolean if a field has been set.

### GetDirectoryResponse

`func (o *ThreeDSecureResultRequestEventContext) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureResultRequestEventContext) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureResultRequestEventContext) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.

### HasDirectoryResponse

`func (o *ThreeDSecureResultRequestEventContext) HasDirectoryResponse() bool`

HasDirectoryResponse returns a boolean if a field has been set.

### GetDirectoryTransactionId

`func (o *ThreeDSecureResultRequestEventContext) GetDirectoryTransactionId() string`

GetDirectoryTransactionId returns the DirectoryTransactionId field if non-nil, zero value otherwise.

### GetDirectoryTransactionIdOk

`func (o *ThreeDSecureResultRequestEventContext) GetDirectoryTransactionIdOk() (*string, bool)`

GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTransactionId

`func (o *ThreeDSecureResultRequestEventContext) SetDirectoryTransactionId(v string)`

SetDirectoryTransactionId sets DirectoryTransactionId field to given value.

### HasDirectoryTransactionId

`func (o *ThreeDSecureResultRequestEventContext) HasDirectoryTransactionId() bool`

HasDirectoryTransactionId returns a boolean if a field has been set.

### GetVersion

`func (o *ThreeDSecureResultRequestEventContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureResultRequestEventContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureResultRequestEventContext) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSecureResultRequestEventContext) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


