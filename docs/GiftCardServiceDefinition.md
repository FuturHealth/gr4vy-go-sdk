# GiftCardServiceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the gift card service definition. | [optional] 
**Type** | Pointer to **string** | &#x60;gift-card-service-definition&#x60;. | [optional] [default to "gift-card-service-definition"]
**DisplayName** | Pointer to **string** | The display name of this service. | [optional] 
**Fields** | Pointer to [**[]AntiFraudServiceDefinitionFieldsInner**](AntiFraudServiceDefinitionFieldsInner.md) | A list of fields that need to be submitted when activating the payment. service. | [optional] 
**IconUrl** | Pointer to **string** | An icon to display for the payment service. | [optional] 

## Methods

### NewGiftCardServiceDefinition

`func NewGiftCardServiceDefinition() *GiftCardServiceDefinition`

NewGiftCardServiceDefinition instantiates a new GiftCardServiceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardServiceDefinitionWithDefaults

`func NewGiftCardServiceDefinitionWithDefaults() *GiftCardServiceDefinition`

NewGiftCardServiceDefinitionWithDefaults instantiates a new GiftCardServiceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GiftCardServiceDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardServiceDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardServiceDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardServiceDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GiftCardServiceDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardServiceDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardServiceDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardServiceDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GiftCardServiceDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GiftCardServiceDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GiftCardServiceDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GiftCardServiceDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *GiftCardServiceDefinition) GetFields() []AntiFraudServiceDefinitionFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GiftCardServiceDefinition) GetFieldsOk() (*[]AntiFraudServiceDefinitionFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GiftCardServiceDefinition) SetFields(v []AntiFraudServiceDefinitionFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GiftCardServiceDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIconUrl

`func (o *GiftCardServiceDefinition) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *GiftCardServiceDefinition) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *GiftCardServiceDefinition) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *GiftCardServiceDefinition) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


