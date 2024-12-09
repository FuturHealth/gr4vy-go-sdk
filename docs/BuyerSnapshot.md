# BuyerSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;buyer&#x60;. | [optional] 
**Id** | Pointer to **NullableString** | The unique Gr4vy ID for this buyer. | [optional] 
**BillingDetails** | Pointer to [**NullableBillingDetails**](BillingDetails.md) | The billing details associated with the buyer, which include the address and tax ID. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the buyer against your own records. | [optional] 

## Methods

### NewBuyerSnapshot

`func NewBuyerSnapshot() *BuyerSnapshot`

NewBuyerSnapshot instantiates a new BuyerSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerSnapshotWithDefaults

`func NewBuyerSnapshotWithDefaults() *BuyerSnapshot`

NewBuyerSnapshotWithDefaults instantiates a new BuyerSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BuyerSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BuyerSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BuyerSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BuyerSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *BuyerSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuyerSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuyerSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BuyerSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BuyerSnapshot) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BuyerSnapshot) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetBillingDetails

`func (o *BuyerSnapshot) GetBillingDetails() BillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *BuyerSnapshot) GetBillingDetailsOk() (*BillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *BuyerSnapshot) SetBillingDetails(v BillingDetails)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *BuyerSnapshot) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### SetBillingDetailsNil

`func (o *BuyerSnapshot) SetBillingDetailsNil(b bool)`

 SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil

### UnsetBillingDetails
`func (o *BuyerSnapshot) UnsetBillingDetails()`

UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil
### GetDisplayName

`func (o *BuyerSnapshot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BuyerSnapshot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BuyerSnapshot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BuyerSnapshot) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BuyerSnapshot) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BuyerSnapshot) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExternalIdentifier

`func (o *BuyerSnapshot) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *BuyerSnapshot) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *BuyerSnapshot) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *BuyerSnapshot) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *BuyerSnapshot) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *BuyerSnapshot) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


