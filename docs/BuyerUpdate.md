# BuyerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the buyer against your own records. This value needs to be unique for all buyers. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name. | [optional] 
**BillingDetails** | Pointer to [**NullableBillingDetailsUpdateRequest**](BillingDetailsUpdateRequest.md) | The billing details of the buyer. | [optional] 

## Methods

### NewBuyerUpdate

`func NewBuyerUpdate() *BuyerUpdate`

NewBuyerUpdate instantiates a new BuyerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerUpdateWithDefaults

`func NewBuyerUpdateWithDefaults() *BuyerUpdate`

NewBuyerUpdateWithDefaults instantiates a new BuyerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIdentifier

`func (o *BuyerUpdate) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *BuyerUpdate) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *BuyerUpdate) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *BuyerUpdate) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *BuyerUpdate) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *BuyerUpdate) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetDisplayName

`func (o *BuyerUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BuyerUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BuyerUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BuyerUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BuyerUpdate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BuyerUpdate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetBillingDetails

`func (o *BuyerUpdate) GetBillingDetails() BillingDetailsUpdateRequest`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *BuyerUpdate) GetBillingDetailsOk() (*BillingDetailsUpdateRequest, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *BuyerUpdate) SetBillingDetails(v BillingDetailsUpdateRequest)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *BuyerUpdate) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### SetBillingDetailsNil

`func (o *BuyerUpdate) SetBillingDetailsNil(b bool)`

 SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil

### UnsetBillingDetails
`func (o *BuyerUpdate) UnsetBillingDetails()`

UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


