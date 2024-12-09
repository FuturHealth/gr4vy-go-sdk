# ThreeDSecurePreparationRequestEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL that was called for this request. | [optional] 
**Request** | Pointer to **string** | The request body sent to the &#x60;url&#x60;. | [optional] 
**Response** | Pointer to **string** | The response body received from the &#x60;url&#x60;. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | The response status code received from the &#x60;url&#x60;. | [optional] 
**IsEnrolled** | Pointer to **bool** | If the card is enrolled for 3DS. | [optional] 
**Version** | Pointer to **string** | The version of 3DS extracted from the &#x60;response. | [optional] 

## Methods

### NewThreeDSecurePreparationRequestEventContext

`func NewThreeDSecurePreparationRequestEventContext() *ThreeDSecurePreparationRequestEventContext`

NewThreeDSecurePreparationRequestEventContext instantiates a new ThreeDSecurePreparationRequestEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecurePreparationRequestEventContextWithDefaults

`func NewThreeDSecurePreparationRequestEventContextWithDefaults() *ThreeDSecurePreparationRequestEventContext`

NewThreeDSecurePreparationRequestEventContextWithDefaults instantiates a new ThreeDSecurePreparationRequestEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ThreeDSecurePreparationRequestEventContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThreeDSecurePreparationRequestEventContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThreeDSecurePreparationRequestEventContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThreeDSecurePreparationRequestEventContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *ThreeDSecurePreparationRequestEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ThreeDSecurePreparationRequestEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ThreeDSecurePreparationRequestEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ThreeDSecurePreparationRequestEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *ThreeDSecurePreparationRequestEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ThreeDSecurePreparationRequestEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ThreeDSecurePreparationRequestEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ThreeDSecurePreparationRequestEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *ThreeDSecurePreparationRequestEventContext) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *ThreeDSecurePreparationRequestEventContext) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *ThreeDSecurePreparationRequestEventContext) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *ThreeDSecurePreparationRequestEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetIsEnrolled

`func (o *ThreeDSecurePreparationRequestEventContext) GetIsEnrolled() bool`

GetIsEnrolled returns the IsEnrolled field if non-nil, zero value otherwise.

### GetIsEnrolledOk

`func (o *ThreeDSecurePreparationRequestEventContext) GetIsEnrolledOk() (*bool, bool)`

GetIsEnrolledOk returns a tuple with the IsEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnrolled

`func (o *ThreeDSecurePreparationRequestEventContext) SetIsEnrolled(v bool)`

SetIsEnrolled sets IsEnrolled field to given value.

### HasIsEnrolled

`func (o *ThreeDSecurePreparationRequestEventContext) HasIsEnrolled() bool`

HasIsEnrolled returns a boolean if a field has been set.

### GetVersion

`func (o *ThreeDSecurePreparationRequestEventContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecurePreparationRequestEventContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecurePreparationRequestEventContext) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSecurePreparationRequestEventContext) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


