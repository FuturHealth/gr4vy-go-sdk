# Buyer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;buyer&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique Gr4vy ID for this buyer. | [optional] 
**BillingDetails** | Pointer to [**NullableBillingDetails**](BillingDetails.md) | The billing details associated with a buyer. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this buyer was created in our system. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the buyer against your own records. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this buyer was last updated in our system. | [optional] 

## Methods

### NewBuyer

`func NewBuyer() *Buyer`

NewBuyer instantiates a new Buyer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerWithDefaults

`func NewBuyerWithDefaults() *Buyer`

NewBuyerWithDefaults instantiates a new Buyer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Buyer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Buyer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Buyer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Buyer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Buyer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Buyer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Buyer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Buyer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillingDetails

`func (o *Buyer) GetBillingDetails() BillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *Buyer) GetBillingDetailsOk() (*BillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *Buyer) SetBillingDetails(v BillingDetails)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *Buyer) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### SetBillingDetailsNil

`func (o *Buyer) SetBillingDetailsNil(b bool)`

 SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil

### UnsetBillingDetails
`func (o *Buyer) UnsetBillingDetails()`

UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil
### GetCreatedAt

`func (o *Buyer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Buyer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Buyer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Buyer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *Buyer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Buyer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Buyer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Buyer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Buyer) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Buyer) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExternalIdentifier

`func (o *Buyer) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *Buyer) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *Buyer) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *Buyer) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *Buyer) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *Buyer) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetMerchantAccountId

`func (o *Buyer) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *Buyer) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *Buyer) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *Buyer) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Buyer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Buyer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Buyer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Buyer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


