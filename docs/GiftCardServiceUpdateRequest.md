# GiftCardServiceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A custom name for the gift card service. This will be shown in the Admin UI. | [optional] 
**Fields** | Pointer to [**[]GiftCardServiceUpdateRequestFieldsInner**](GiftCardServiceUpdateRequestFieldsInner.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this gift card service. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] 

## Methods

### NewGiftCardServiceUpdateRequest

`func NewGiftCardServiceUpdateRequest() *GiftCardServiceUpdateRequest`

NewGiftCardServiceUpdateRequest instantiates a new GiftCardServiceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardServiceUpdateRequestWithDefaults

`func NewGiftCardServiceUpdateRequestWithDefaults() *GiftCardServiceUpdateRequest`

NewGiftCardServiceUpdateRequestWithDefaults instantiates a new GiftCardServiceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GiftCardServiceUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GiftCardServiceUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GiftCardServiceUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GiftCardServiceUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *GiftCardServiceUpdateRequest) GetFields() []GiftCardServiceUpdateRequestFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GiftCardServiceUpdateRequest) GetFieldsOk() (*[]GiftCardServiceUpdateRequestFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GiftCardServiceUpdateRequest) SetFields(v []GiftCardServiceUpdateRequestFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GiftCardServiceUpdateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetActive

`func (o *GiftCardServiceUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GiftCardServiceUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GiftCardServiceUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GiftCardServiceUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


