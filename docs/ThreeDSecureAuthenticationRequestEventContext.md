# ThreeDSecureAuthenticationRequestEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL that was called for this request. | [optional] 
**Request** | Pointer to **string** | The request body sent to the &#x60;url&#x60;. | [optional] 
**Response** | Pointer to **string** | The response body received from the &#x60;url&#x60;. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | The response status code received from the &#x60;url&#x60;. | [optional] 
**Cavv** | Pointer to **string** | The 3DS CAVV value parsed from the &#x60;response&#x60;. | [optional] 
**Eci** | Pointer to **string** | The 3DS ECI value parsed from the &#x60;response&#x60;. | [optional] 
**DirectoryResponse** | Pointer to **string** | The &#x60;transStatus&#x60; parsed from the &#x60;response&#x60;. | [optional] 
**DirectoryTransactionId** | Pointer to **string** | The &#x60;dsTransID&#x60; parsed from the &#x60;response&#x60;. | [optional] 
**Version** | Pointer to **string** | The version of 3DS used. | [optional] 

## Methods

### NewThreeDSecureAuthenticationRequestEventContext

`func NewThreeDSecureAuthenticationRequestEventContext() *ThreeDSecureAuthenticationRequestEventContext`

NewThreeDSecureAuthenticationRequestEventContext instantiates a new ThreeDSecureAuthenticationRequestEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureAuthenticationRequestEventContextWithDefaults

`func NewThreeDSecureAuthenticationRequestEventContextWithDefaults() *ThreeDSecureAuthenticationRequestEventContext`

NewThreeDSecureAuthenticationRequestEventContextWithDefaults instantiates a new ThreeDSecureAuthenticationRequestEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetCavv

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetEci

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetDirectoryResponse

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.

### HasDirectoryResponse

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasDirectoryResponse() bool`

HasDirectoryResponse returns a boolean if a field has been set.

### GetDirectoryTransactionId

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetDirectoryTransactionId() string`

GetDirectoryTransactionId returns the DirectoryTransactionId field if non-nil, zero value otherwise.

### GetDirectoryTransactionIdOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetDirectoryTransactionIdOk() (*string, bool)`

GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTransactionId

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetDirectoryTransactionId(v string)`

SetDirectoryTransactionId sets DirectoryTransactionId field to given value.

### HasDirectoryTransactionId

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasDirectoryTransactionId() bool`

HasDirectoryTransactionId returns a boolean if a field has been set.

### GetVersion

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureAuthenticationRequestEventContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureAuthenticationRequestEventContext) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSecureAuthenticationRequestEventContext) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


