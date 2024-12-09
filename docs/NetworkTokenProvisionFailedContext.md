# NetworkTokenProvisionFailedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The endpoint for the request, if performed. | [optional] 
**Request** | Pointer to **NullableString** | The HTTP body sent to the Network Token provider, if performed. | [optional] 
**Response** | Pointer to **NullableString** | The HTTP body received from the Network Token provider, if any. | [optional] 
**ResponseStatusCode** | Pointer to **NullableFloat32** | The HTTP response status code from the Network Token provider, if any. | [optional] 
**Reason** | Pointer to **string** | The reason we could not provision the network token. | [optional] 

## Methods

### NewNetworkTokenProvisionFailedContext

`func NewNetworkTokenProvisionFailedContext() *NetworkTokenProvisionFailedContext`

NewNetworkTokenProvisionFailedContext instantiates a new NetworkTokenProvisionFailedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTokenProvisionFailedContextWithDefaults

`func NewNetworkTokenProvisionFailedContextWithDefaults() *NetworkTokenProvisionFailedContext`

NewNetworkTokenProvisionFailedContextWithDefaults instantiates a new NetworkTokenProvisionFailedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NetworkTokenProvisionFailedContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NetworkTokenProvisionFailedContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NetworkTokenProvisionFailedContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NetworkTokenProvisionFailedContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NetworkTokenProvisionFailedContext) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NetworkTokenProvisionFailedContext) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRequest

`func (o *NetworkTokenProvisionFailedContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *NetworkTokenProvisionFailedContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *NetworkTokenProvisionFailedContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *NetworkTokenProvisionFailedContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *NetworkTokenProvisionFailedContext) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *NetworkTokenProvisionFailedContext) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetResponse

`func (o *NetworkTokenProvisionFailedContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *NetworkTokenProvisionFailedContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *NetworkTokenProvisionFailedContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *NetworkTokenProvisionFailedContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *NetworkTokenProvisionFailedContext) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *NetworkTokenProvisionFailedContext) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseStatusCode

`func (o *NetworkTokenProvisionFailedContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *NetworkTokenProvisionFailedContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *NetworkTokenProvisionFailedContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *NetworkTokenProvisionFailedContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *NetworkTokenProvisionFailedContext) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *NetworkTokenProvisionFailedContext) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetReason

`func (o *NetworkTokenProvisionFailedContext) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NetworkTokenProvisionFailedContext) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NetworkTokenProvisionFailedContext) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NetworkTokenProvisionFailedContext) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


