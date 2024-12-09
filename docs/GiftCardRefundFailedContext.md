# GiftCardRefundFailedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftCardServiceId** | Pointer to **string** | The unique ID of the Gift Card service used. | [optional] 
**GiftCardServiceName** | Pointer to **string** | The name of the Gift Card service used. | [optional] 
**GiftCardServiceDefinitionId** | Pointer to **string** | The Gift Card service definition used. | [optional] 
**Url** | Pointer to **NullableString** | The endpoint for the request, if performed. | [optional] 
**Request** | Pointer to **NullableString** | The HTTP body sent to the Gift Card provider, if performed. | [optional] 
**Response** | Pointer to **NullableString** | The HTTP body received from the Gift Card provider, if any. | [optional] 
**ResponseStatusCode** | Pointer to **NullableFloat32** | The HTTP response status code from the Gift Card provider, if any. | [optional] 
**Reason** | Pointer to **string** | The reason we could not refund the gift cards. | [optional] 

## Methods

### NewGiftCardRefundFailedContext

`func NewGiftCardRefundFailedContext() *GiftCardRefundFailedContext`

NewGiftCardRefundFailedContext instantiates a new GiftCardRefundFailedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRefundFailedContextWithDefaults

`func NewGiftCardRefundFailedContextWithDefaults() *GiftCardRefundFailedContext`

NewGiftCardRefundFailedContextWithDefaults instantiates a new GiftCardRefundFailedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftCardServiceId

`func (o *GiftCardRefundFailedContext) GetGiftCardServiceId() string`

GetGiftCardServiceId returns the GiftCardServiceId field if non-nil, zero value otherwise.

### GetGiftCardServiceIdOk

`func (o *GiftCardRefundFailedContext) GetGiftCardServiceIdOk() (*string, bool)`

GetGiftCardServiceIdOk returns a tuple with the GiftCardServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceId

`func (o *GiftCardRefundFailedContext) SetGiftCardServiceId(v string)`

SetGiftCardServiceId sets GiftCardServiceId field to given value.

### HasGiftCardServiceId

`func (o *GiftCardRefundFailedContext) HasGiftCardServiceId() bool`

HasGiftCardServiceId returns a boolean if a field has been set.

### GetGiftCardServiceName

`func (o *GiftCardRefundFailedContext) GetGiftCardServiceName() string`

GetGiftCardServiceName returns the GiftCardServiceName field if non-nil, zero value otherwise.

### GetGiftCardServiceNameOk

`func (o *GiftCardRefundFailedContext) GetGiftCardServiceNameOk() (*string, bool)`

GetGiftCardServiceNameOk returns a tuple with the GiftCardServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceName

`func (o *GiftCardRefundFailedContext) SetGiftCardServiceName(v string)`

SetGiftCardServiceName sets GiftCardServiceName field to given value.

### HasGiftCardServiceName

`func (o *GiftCardRefundFailedContext) HasGiftCardServiceName() bool`

HasGiftCardServiceName returns a boolean if a field has been set.

### GetGiftCardServiceDefinitionId

`func (o *GiftCardRefundFailedContext) GetGiftCardServiceDefinitionId() string`

GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field if non-nil, zero value otherwise.

### GetGiftCardServiceDefinitionIdOk

`func (o *GiftCardRefundFailedContext) GetGiftCardServiceDefinitionIdOk() (*string, bool)`

GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceDefinitionId

`func (o *GiftCardRefundFailedContext) SetGiftCardServiceDefinitionId(v string)`

SetGiftCardServiceDefinitionId sets GiftCardServiceDefinitionId field to given value.

### HasGiftCardServiceDefinitionId

`func (o *GiftCardRefundFailedContext) HasGiftCardServiceDefinitionId() bool`

HasGiftCardServiceDefinitionId returns a boolean if a field has been set.

### GetUrl

`func (o *GiftCardRefundFailedContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GiftCardRefundFailedContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GiftCardRefundFailedContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GiftCardRefundFailedContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GiftCardRefundFailedContext) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GiftCardRefundFailedContext) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRequest

`func (o *GiftCardRefundFailedContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GiftCardRefundFailedContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GiftCardRefundFailedContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GiftCardRefundFailedContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *GiftCardRefundFailedContext) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *GiftCardRefundFailedContext) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetResponse

`func (o *GiftCardRefundFailedContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GiftCardRefundFailedContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GiftCardRefundFailedContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GiftCardRefundFailedContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *GiftCardRefundFailedContext) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *GiftCardRefundFailedContext) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseStatusCode

`func (o *GiftCardRefundFailedContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *GiftCardRefundFailedContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *GiftCardRefundFailedContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *GiftCardRefundFailedContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *GiftCardRefundFailedContext) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *GiftCardRefundFailedContext) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetReason

`func (o *GiftCardRefundFailedContext) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GiftCardRefundFailedContext) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GiftCardRefundFailedContext) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GiftCardRefundFailedContext) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


