# GiftCardRedemptionFailedContext

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
**Reason** | Pointer to **string** | The reason we could not redeem the gift cards. | [optional] 

## Methods

### NewGiftCardRedemptionFailedContext

`func NewGiftCardRedemptionFailedContext() *GiftCardRedemptionFailedContext`

NewGiftCardRedemptionFailedContext instantiates a new GiftCardRedemptionFailedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRedemptionFailedContextWithDefaults

`func NewGiftCardRedemptionFailedContextWithDefaults() *GiftCardRedemptionFailedContext`

NewGiftCardRedemptionFailedContextWithDefaults instantiates a new GiftCardRedemptionFailedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftCardServiceId

`func (o *GiftCardRedemptionFailedContext) GetGiftCardServiceId() string`

GetGiftCardServiceId returns the GiftCardServiceId field if non-nil, zero value otherwise.

### GetGiftCardServiceIdOk

`func (o *GiftCardRedemptionFailedContext) GetGiftCardServiceIdOk() (*string, bool)`

GetGiftCardServiceIdOk returns a tuple with the GiftCardServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceId

`func (o *GiftCardRedemptionFailedContext) SetGiftCardServiceId(v string)`

SetGiftCardServiceId sets GiftCardServiceId field to given value.

### HasGiftCardServiceId

`func (o *GiftCardRedemptionFailedContext) HasGiftCardServiceId() bool`

HasGiftCardServiceId returns a boolean if a field has been set.

### GetGiftCardServiceName

`func (o *GiftCardRedemptionFailedContext) GetGiftCardServiceName() string`

GetGiftCardServiceName returns the GiftCardServiceName field if non-nil, zero value otherwise.

### GetGiftCardServiceNameOk

`func (o *GiftCardRedemptionFailedContext) GetGiftCardServiceNameOk() (*string, bool)`

GetGiftCardServiceNameOk returns a tuple with the GiftCardServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceName

`func (o *GiftCardRedemptionFailedContext) SetGiftCardServiceName(v string)`

SetGiftCardServiceName sets GiftCardServiceName field to given value.

### HasGiftCardServiceName

`func (o *GiftCardRedemptionFailedContext) HasGiftCardServiceName() bool`

HasGiftCardServiceName returns a boolean if a field has been set.

### GetGiftCardServiceDefinitionId

`func (o *GiftCardRedemptionFailedContext) GetGiftCardServiceDefinitionId() string`

GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field if non-nil, zero value otherwise.

### GetGiftCardServiceDefinitionIdOk

`func (o *GiftCardRedemptionFailedContext) GetGiftCardServiceDefinitionIdOk() (*string, bool)`

GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceDefinitionId

`func (o *GiftCardRedemptionFailedContext) SetGiftCardServiceDefinitionId(v string)`

SetGiftCardServiceDefinitionId sets GiftCardServiceDefinitionId field to given value.

### HasGiftCardServiceDefinitionId

`func (o *GiftCardRedemptionFailedContext) HasGiftCardServiceDefinitionId() bool`

HasGiftCardServiceDefinitionId returns a boolean if a field has been set.

### GetUrl

`func (o *GiftCardRedemptionFailedContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GiftCardRedemptionFailedContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GiftCardRedemptionFailedContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GiftCardRedemptionFailedContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GiftCardRedemptionFailedContext) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GiftCardRedemptionFailedContext) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRequest

`func (o *GiftCardRedemptionFailedContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GiftCardRedemptionFailedContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GiftCardRedemptionFailedContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GiftCardRedemptionFailedContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *GiftCardRedemptionFailedContext) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *GiftCardRedemptionFailedContext) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetResponse

`func (o *GiftCardRedemptionFailedContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GiftCardRedemptionFailedContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GiftCardRedemptionFailedContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GiftCardRedemptionFailedContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *GiftCardRedemptionFailedContext) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *GiftCardRedemptionFailedContext) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseStatusCode

`func (o *GiftCardRedemptionFailedContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *GiftCardRedemptionFailedContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *GiftCardRedemptionFailedContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *GiftCardRedemptionFailedContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *GiftCardRedemptionFailedContext) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *GiftCardRedemptionFailedContext) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetReason

`func (o *GiftCardRedemptionFailedContext) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GiftCardRedemptionFailedContext) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GiftCardRedemptionFailedContext) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GiftCardRedemptionFailedContext) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


