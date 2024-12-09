# GiftCardServiceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftCardServiceDefinitionId** | **string** | The ID of the gift card service to use. | 
**DisplayName** | **string** | A custom name for the service. This will be shown in the Admin UI. | 
**Fields** | [**[]GiftCardServiceCreateRequestFieldsInner**](GiftCardServiceCreateRequestFieldsInner.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this gift card service. | 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]

## Methods

### NewGiftCardServiceCreateRequest

`func NewGiftCardServiceCreateRequest(giftCardServiceDefinitionId string, displayName string, fields []GiftCardServiceCreateRequestFieldsInner, ) *GiftCardServiceCreateRequest`

NewGiftCardServiceCreateRequest instantiates a new GiftCardServiceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardServiceCreateRequestWithDefaults

`func NewGiftCardServiceCreateRequestWithDefaults() *GiftCardServiceCreateRequest`

NewGiftCardServiceCreateRequestWithDefaults instantiates a new GiftCardServiceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftCardServiceDefinitionId

`func (o *GiftCardServiceCreateRequest) GetGiftCardServiceDefinitionId() string`

GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field if non-nil, zero value otherwise.

### GetGiftCardServiceDefinitionIdOk

`func (o *GiftCardServiceCreateRequest) GetGiftCardServiceDefinitionIdOk() (*string, bool)`

GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceDefinitionId

`func (o *GiftCardServiceCreateRequest) SetGiftCardServiceDefinitionId(v string)`

SetGiftCardServiceDefinitionId sets GiftCardServiceDefinitionId field to given value.


### GetDisplayName

`func (o *GiftCardServiceCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GiftCardServiceCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GiftCardServiceCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFields

`func (o *GiftCardServiceCreateRequest) GetFields() []GiftCardServiceCreateRequestFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GiftCardServiceCreateRequest) GetFieldsOk() (*[]GiftCardServiceCreateRequestFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GiftCardServiceCreateRequest) SetFields(v []GiftCardServiceCreateRequestFieldsInner)`

SetFields sets Fields field to given value.


### GetActive

`func (o *GiftCardServiceCreateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GiftCardServiceCreateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GiftCardServiceCreateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GiftCardServiceCreateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


