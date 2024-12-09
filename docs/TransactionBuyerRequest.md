# TransactionBuyerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the buyer against your own records. This value needs to be unique for all buyers. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name. | [optional] 
**BillingDetails** | Pointer to [**NullableBillingDetailsRequest**](BillingDetailsRequest.md) | The optional billing details for the a buyer. | [optional] 
**ShippingDetails** | Pointer to [**NullableBillingDetailsRequest**](BillingDetailsRequest.md) | The optional shipping details for the buyer. | [optional] 

## Methods

### NewTransactionBuyerRequest

`func NewTransactionBuyerRequest() *TransactionBuyerRequest`

NewTransactionBuyerRequest instantiates a new TransactionBuyerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBuyerRequestWithDefaults

`func NewTransactionBuyerRequestWithDefaults() *TransactionBuyerRequest`

NewTransactionBuyerRequestWithDefaults instantiates a new TransactionBuyerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIdentifier

`func (o *TransactionBuyerRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionBuyerRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionBuyerRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionBuyerRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionBuyerRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionBuyerRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetDisplayName

`func (o *TransactionBuyerRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TransactionBuyerRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TransactionBuyerRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TransactionBuyerRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TransactionBuyerRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TransactionBuyerRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetBillingDetails

`func (o *TransactionBuyerRequest) GetBillingDetails() BillingDetailsRequest`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *TransactionBuyerRequest) GetBillingDetailsOk() (*BillingDetailsRequest, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *TransactionBuyerRequest) SetBillingDetails(v BillingDetailsRequest)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *TransactionBuyerRequest) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### SetBillingDetailsNil

`func (o *TransactionBuyerRequest) SetBillingDetailsNil(b bool)`

 SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil

### UnsetBillingDetails
`func (o *TransactionBuyerRequest) UnsetBillingDetails()`

UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil
### GetShippingDetails

`func (o *TransactionBuyerRequest) GetShippingDetails() BillingDetailsRequest`

GetShippingDetails returns the ShippingDetails field if non-nil, zero value otherwise.

### GetShippingDetailsOk

`func (o *TransactionBuyerRequest) GetShippingDetailsOk() (*BillingDetailsRequest, bool)`

GetShippingDetailsOk returns a tuple with the ShippingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDetails

`func (o *TransactionBuyerRequest) SetShippingDetails(v BillingDetailsRequest)`

SetShippingDetails sets ShippingDetails field to given value.

### HasShippingDetails

`func (o *TransactionBuyerRequest) HasShippingDetails() bool`

HasShippingDetails returns a boolean if a field has been set.

### SetShippingDetailsNil

`func (o *TransactionBuyerRequest) SetShippingDetailsNil(b bool)`

 SetShippingDetailsNil sets the value for ShippingDetails to be an explicit nil

### UnsetShippingDetails
`func (o *TransactionBuyerRequest) UnsetShippingDetails()`

UnsetShippingDetails ensures that no value is present for ShippingDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


