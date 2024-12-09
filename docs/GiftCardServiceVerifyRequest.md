# GiftCardServiceVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftCardServiceDefinitionId** | **string** | The ID of the gift card service to use. | 
**GiftCardServiceId** | Pointer to **string** | The ID of the gift card service. Required if sending a partial set of credentials in the &#x60;fields&#x60; property. This will merge the provided fields with those already on the service. | [optional] 
**Fields** | [**[]GiftCardServiceVerifyRequestFieldsInner**](GiftCardServiceVerifyRequestFieldsInner.md) | A list of fields where each field is a key-value pair that represents a defined field in the definition of the service. You are not required to send the full list of fields if the credentials for the service are already stored. For example, if your credentials for &#x60;qwikcilver-gift-card&#x60; are stored and you only provide a &#x60;secret_key&#x60; in the request, it will override the stored &#x60;secret_key&#x60; and verify the resulting set of credentials against the service. | 

## Methods

### NewGiftCardServiceVerifyRequest

`func NewGiftCardServiceVerifyRequest(giftCardServiceDefinitionId string, fields []GiftCardServiceVerifyRequestFieldsInner, ) *GiftCardServiceVerifyRequest`

NewGiftCardServiceVerifyRequest instantiates a new GiftCardServiceVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardServiceVerifyRequestWithDefaults

`func NewGiftCardServiceVerifyRequestWithDefaults() *GiftCardServiceVerifyRequest`

NewGiftCardServiceVerifyRequestWithDefaults instantiates a new GiftCardServiceVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftCardServiceDefinitionId

`func (o *GiftCardServiceVerifyRequest) GetGiftCardServiceDefinitionId() string`

GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field if non-nil, zero value otherwise.

### GetGiftCardServiceDefinitionIdOk

`func (o *GiftCardServiceVerifyRequest) GetGiftCardServiceDefinitionIdOk() (*string, bool)`

GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceDefinitionId

`func (o *GiftCardServiceVerifyRequest) SetGiftCardServiceDefinitionId(v string)`

SetGiftCardServiceDefinitionId sets GiftCardServiceDefinitionId field to given value.


### GetGiftCardServiceId

`func (o *GiftCardServiceVerifyRequest) GetGiftCardServiceId() string`

GetGiftCardServiceId returns the GiftCardServiceId field if non-nil, zero value otherwise.

### GetGiftCardServiceIdOk

`func (o *GiftCardServiceVerifyRequest) GetGiftCardServiceIdOk() (*string, bool)`

GetGiftCardServiceIdOk returns a tuple with the GiftCardServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceId

`func (o *GiftCardServiceVerifyRequest) SetGiftCardServiceId(v string)`

SetGiftCardServiceId sets GiftCardServiceId field to given value.

### HasGiftCardServiceId

`func (o *GiftCardServiceVerifyRequest) HasGiftCardServiceId() bool`

HasGiftCardServiceId returns a boolean if a field has been set.

### GetFields

`func (o *GiftCardServiceVerifyRequest) GetFields() []GiftCardServiceVerifyRequestFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GiftCardServiceVerifyRequest) GetFieldsOk() (*[]GiftCardServiceVerifyRequestFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GiftCardServiceVerifyRequest) SetFields(v []GiftCardServiceVerifyRequestFieldsInner)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


