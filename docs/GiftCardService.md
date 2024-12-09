# GiftCardService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **string** | The ID of this gift card service. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**GiftCardServiceDefinitionId** | Pointer to **string** | The ID of the gift card service definition used to create this service.  | [optional] 
**DisplayName** | Pointer to **string** | The custom name set for this service. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] 
**Fields** | Pointer to [**[]GiftCardServiceFieldsInner**](GiftCardServiceFieldsInner.md) | A list of fields, each containing a key-value pair for each field configured for this gift card service. Fields marked as &#x60;secret&#x60; are not returned. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this service was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this service was last updated. | [optional] 

## Methods

### NewGiftCardService

`func NewGiftCardService() *GiftCardService`

NewGiftCardService instantiates a new GiftCardService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardServiceWithDefaults

`func NewGiftCardServiceWithDefaults() *GiftCardService`

NewGiftCardServiceWithDefaults instantiates a new GiftCardService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCardService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *GiftCardService) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *GiftCardService) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *GiftCardService) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *GiftCardService) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetGiftCardServiceDefinitionId

`func (o *GiftCardService) GetGiftCardServiceDefinitionId() string`

GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field if non-nil, zero value otherwise.

### GetGiftCardServiceDefinitionIdOk

`func (o *GiftCardService) GetGiftCardServiceDefinitionIdOk() (*string, bool)`

GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceDefinitionId

`func (o *GiftCardService) SetGiftCardServiceDefinitionId(v string)`

SetGiftCardServiceDefinitionId sets GiftCardServiceDefinitionId field to given value.

### HasGiftCardServiceDefinitionId

`func (o *GiftCardService) HasGiftCardServiceDefinitionId() bool`

HasGiftCardServiceDefinitionId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GiftCardService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GiftCardService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GiftCardService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GiftCardService) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetActive

`func (o *GiftCardService) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GiftCardService) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GiftCardService) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GiftCardService) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFields

`func (o *GiftCardService) GetFields() []GiftCardServiceFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GiftCardService) GetFieldsOk() (*[]GiftCardServiceFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GiftCardService) SetFields(v []GiftCardServiceFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GiftCardService) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GiftCardService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GiftCardService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GiftCardService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GiftCardService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GiftCardService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GiftCardService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GiftCardService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GiftCardService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


