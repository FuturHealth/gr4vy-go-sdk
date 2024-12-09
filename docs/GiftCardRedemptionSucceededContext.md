# GiftCardRedemptionSucceededContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftCardServiceId** | Pointer to **string** | The unique ID of the Gift Card service used. | [optional] 
**GiftCardServiceName** | Pointer to **string** | The name of the Gift Card service used. | [optional] 
**GiftCardServiceDefinitionId** | Pointer to **string** | The Gift Card service definition used. | [optional] 
**Url** | Pointer to **NullableString** | The endpoint for the request. | [optional] 
**Request** | Pointer to **string** | The HTTP body sent to the Gift Card provider. | [optional] 
**Response** | Pointer to **string** | The HTTP body received from the Gift Card provider. | [optional] 
**ResponseStatusCode** | Pointer to **float32** | The HTTP response status code from the Gift Card provider. | [optional] 

## Methods

### NewGiftCardRedemptionSucceededContext

`func NewGiftCardRedemptionSucceededContext() *GiftCardRedemptionSucceededContext`

NewGiftCardRedemptionSucceededContext instantiates a new GiftCardRedemptionSucceededContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRedemptionSucceededContextWithDefaults

`func NewGiftCardRedemptionSucceededContextWithDefaults() *GiftCardRedemptionSucceededContext`

NewGiftCardRedemptionSucceededContextWithDefaults instantiates a new GiftCardRedemptionSucceededContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftCardServiceId

`func (o *GiftCardRedemptionSucceededContext) GetGiftCardServiceId() string`

GetGiftCardServiceId returns the GiftCardServiceId field if non-nil, zero value otherwise.

### GetGiftCardServiceIdOk

`func (o *GiftCardRedemptionSucceededContext) GetGiftCardServiceIdOk() (*string, bool)`

GetGiftCardServiceIdOk returns a tuple with the GiftCardServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceId

`func (o *GiftCardRedemptionSucceededContext) SetGiftCardServiceId(v string)`

SetGiftCardServiceId sets GiftCardServiceId field to given value.

### HasGiftCardServiceId

`func (o *GiftCardRedemptionSucceededContext) HasGiftCardServiceId() bool`

HasGiftCardServiceId returns a boolean if a field has been set.

### GetGiftCardServiceName

`func (o *GiftCardRedemptionSucceededContext) GetGiftCardServiceName() string`

GetGiftCardServiceName returns the GiftCardServiceName field if non-nil, zero value otherwise.

### GetGiftCardServiceNameOk

`func (o *GiftCardRedemptionSucceededContext) GetGiftCardServiceNameOk() (*string, bool)`

GetGiftCardServiceNameOk returns a tuple with the GiftCardServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceName

`func (o *GiftCardRedemptionSucceededContext) SetGiftCardServiceName(v string)`

SetGiftCardServiceName sets GiftCardServiceName field to given value.

### HasGiftCardServiceName

`func (o *GiftCardRedemptionSucceededContext) HasGiftCardServiceName() bool`

HasGiftCardServiceName returns a boolean if a field has been set.

### GetGiftCardServiceDefinitionId

`func (o *GiftCardRedemptionSucceededContext) GetGiftCardServiceDefinitionId() string`

GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field if non-nil, zero value otherwise.

### GetGiftCardServiceDefinitionIdOk

`func (o *GiftCardRedemptionSucceededContext) GetGiftCardServiceDefinitionIdOk() (*string, bool)`

GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceDefinitionId

`func (o *GiftCardRedemptionSucceededContext) SetGiftCardServiceDefinitionId(v string)`

SetGiftCardServiceDefinitionId sets GiftCardServiceDefinitionId field to given value.

### HasGiftCardServiceDefinitionId

`func (o *GiftCardRedemptionSucceededContext) HasGiftCardServiceDefinitionId() bool`

HasGiftCardServiceDefinitionId returns a boolean if a field has been set.

### GetUrl

`func (o *GiftCardRedemptionSucceededContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GiftCardRedemptionSucceededContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GiftCardRedemptionSucceededContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GiftCardRedemptionSucceededContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GiftCardRedemptionSucceededContext) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GiftCardRedemptionSucceededContext) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRequest

`func (o *GiftCardRedemptionSucceededContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GiftCardRedemptionSucceededContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GiftCardRedemptionSucceededContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GiftCardRedemptionSucceededContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *GiftCardRedemptionSucceededContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GiftCardRedemptionSucceededContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GiftCardRedemptionSucceededContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GiftCardRedemptionSucceededContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *GiftCardRedemptionSucceededContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *GiftCardRedemptionSucceededContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *GiftCardRedemptionSucceededContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *GiftCardRedemptionSucceededContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


