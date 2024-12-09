# BuyerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the buyer against your own records. This value needs to be unique for all buyers. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name. | [optional] 
**BillingDetails** | Pointer to [**NullableBillingDetailsRequest**](BillingDetailsRequest.md) | The optional billing details to associate with a buyer. | [optional] 

## Methods

### NewBuyerRequest

`func NewBuyerRequest() *BuyerRequest`

NewBuyerRequest instantiates a new BuyerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerRequestWithDefaults

`func NewBuyerRequestWithDefaults() *BuyerRequest`

NewBuyerRequestWithDefaults instantiates a new BuyerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIdentifier

`func (o *BuyerRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *BuyerRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *BuyerRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *BuyerRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *BuyerRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *BuyerRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetDisplayName

`func (o *BuyerRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BuyerRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BuyerRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BuyerRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *BuyerRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *BuyerRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetBillingDetails

`func (o *BuyerRequest) GetBillingDetails() BillingDetailsRequest`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *BuyerRequest) GetBillingDetailsOk() (*BillingDetailsRequest, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *BuyerRequest) SetBillingDetails(v BillingDetailsRequest)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *BuyerRequest) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### SetBillingDetailsNil

`func (o *BuyerRequest) SetBillingDetailsNil(b bool)`

 SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil

### UnsetBillingDetails
`func (o *BuyerRequest) UnsetBillingDetails()`

UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


