# APIKeyPairCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A name for this key-pair which is used in the Gr4vy admin panel to give the key-pair a human readable name. | [optional] 
**Algorithm** | Pointer to **string** | The algorithm to use for the API Key Pair. The recommended value is &#x60;ECDSA&#x60;. You should only use the &#x60;RSA&#x60; algorithm in environments that do not support &#x60;ECDSA&#x60;. | [optional] [default to "ECDSA"]
**RoleIds** | Pointer to **[]string** | A list of role IDs that will be assigned to the API Key Pair being created. Only the \&quot;Administrator\&quot; and \&quot;Integration\&quot; roles are supported. | [optional] 
**MerchantAccountId** | Pointer to **NullableString** | The optional ID of the merchant account this API Key Pair should be assigned to. Leave this unset to create an API key that works across all merchant accounts. | [optional] 

## Methods

### NewAPIKeyPairCreate

`func NewAPIKeyPairCreate() *APIKeyPairCreate`

NewAPIKeyPairCreate instantiates a new APIKeyPairCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyPairCreateWithDefaults

`func NewAPIKeyPairCreateWithDefaults() *APIKeyPairCreate`

NewAPIKeyPairCreateWithDefaults instantiates a new APIKeyPairCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *APIKeyPairCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *APIKeyPairCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *APIKeyPairCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *APIKeyPairCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAlgorithm

`func (o *APIKeyPairCreate) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *APIKeyPairCreate) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *APIKeyPairCreate) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *APIKeyPairCreate) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetRoleIds

`func (o *APIKeyPairCreate) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *APIKeyPairCreate) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *APIKeyPairCreate) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *APIKeyPairCreate) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *APIKeyPairCreate) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *APIKeyPairCreate) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *APIKeyPairCreate) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *APIKeyPairCreate) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### SetMerchantAccountIdNil

`func (o *APIKeyPairCreate) SetMerchantAccountIdNil(b bool)`

 SetMerchantAccountIdNil sets the value for MerchantAccountId to be an explicit nil

### UnsetMerchantAccountId
`func (o *APIKeyPairCreate) UnsetMerchantAccountId()`

UnsetMerchantAccountId ensures that no value is present for MerchantAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


