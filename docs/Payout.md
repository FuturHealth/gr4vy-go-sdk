# Payout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;payout&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID of a payout. | [optional] 
**MerchantAccountId** | Pointer to **NullableString** | The optional ID of the merchant account this payout should be assigned to. | [optional] 
**Amount** | Pointer to **int32** | The monetary amount for this payout, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;. | [optional] 
**Currency** | Pointer to **string** | A supported ISO-4217 currency code. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this payout was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this payout was created. | [optional] 
**PaymentService** | Pointer to [**PaymentServiceSnapshot**](PaymentServiceSnapshot.md) | The payment service used for this payout. | [optional] 
**PaymentServicePayoutId** | Pointer to **NullableString** | The ID of the payout in the underlying payment service. | [optional] 
**Category** | Pointer to **string** | The type of payout to process. | [optional] 
**Status** | Pointer to **string** | The status of the payout. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the payout against your own records. This value needs to be unique for all buyers. | [optional] 
**Merchant** | Pointer to [**NullableMerchant**](Merchant.md) | The merchant details associated to this payout. | [optional] 
**Buyer** | Pointer to [**BuyerSnapshot**](BuyerSnapshot.md) | The buyer used for this transaction. | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethodSnapshot**](PaymentMethodSnapshot.md) | The payment method used for this payout. | [optional] 

## Methods

### NewPayout

`func NewPayout() *Payout`

NewPayout instantiates a new Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutWithDefaults

`func NewPayoutWithDefaults() *Payout`

NewPayoutWithDefaults instantiates a new Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Payout) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payout) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payout) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Payout) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Payout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payout) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *Payout) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *Payout) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *Payout) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *Payout) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### SetMerchantAccountIdNil

`func (o *Payout) SetMerchantAccountIdNil(b bool)`

 SetMerchantAccountIdNil sets the value for MerchantAccountId to be an explicit nil

### UnsetMerchantAccountId
`func (o *Payout) UnsetMerchantAccountId()`

UnsetMerchantAccountId ensures that no value is present for MerchantAccountId, not even an explicit nil
### GetAmount

`func (o *Payout) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payout) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payout) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Payout) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Payout) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payout) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payout) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Payout) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Payout) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Payout) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Payout) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Payout) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Payout) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Payout) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Payout) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Payout) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPaymentService

`func (o *Payout) GetPaymentService() PaymentServiceSnapshot`

GetPaymentService returns the PaymentService field if non-nil, zero value otherwise.

### GetPaymentServiceOk

`func (o *Payout) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool)`

GetPaymentServiceOk returns a tuple with the PaymentService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentService

`func (o *Payout) SetPaymentService(v PaymentServiceSnapshot)`

SetPaymentService sets PaymentService field to given value.

### HasPaymentService

`func (o *Payout) HasPaymentService() bool`

HasPaymentService returns a boolean if a field has been set.

### GetPaymentServicePayoutId

`func (o *Payout) GetPaymentServicePayoutId() string`

GetPaymentServicePayoutId returns the PaymentServicePayoutId field if non-nil, zero value otherwise.

### GetPaymentServicePayoutIdOk

`func (o *Payout) GetPaymentServicePayoutIdOk() (*string, bool)`

GetPaymentServicePayoutIdOk returns a tuple with the PaymentServicePayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServicePayoutId

`func (o *Payout) SetPaymentServicePayoutId(v string)`

SetPaymentServicePayoutId sets PaymentServicePayoutId field to given value.

### HasPaymentServicePayoutId

`func (o *Payout) HasPaymentServicePayoutId() bool`

HasPaymentServicePayoutId returns a boolean if a field has been set.

### SetPaymentServicePayoutIdNil

`func (o *Payout) SetPaymentServicePayoutIdNil(b bool)`

 SetPaymentServicePayoutIdNil sets the value for PaymentServicePayoutId to be an explicit nil

### UnsetPaymentServicePayoutId
`func (o *Payout) UnsetPaymentServicePayoutId()`

UnsetPaymentServicePayoutId ensures that no value is present for PaymentServicePayoutId, not even an explicit nil
### GetCategory

`func (o *Payout) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Payout) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Payout) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Payout) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *Payout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payout) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *Payout) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *Payout) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *Payout) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *Payout) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *Payout) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *Payout) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetMerchant

`func (o *Payout) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *Payout) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *Payout) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *Payout) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### SetMerchantNil

`func (o *Payout) SetMerchantNil(b bool)`

 SetMerchantNil sets the value for Merchant to be an explicit nil

### UnsetMerchant
`func (o *Payout) UnsetMerchant()`

UnsetMerchant ensures that no value is present for Merchant, not even an explicit nil
### GetBuyer

`func (o *Payout) GetBuyer() BuyerSnapshot`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Payout) GetBuyerOk() (*BuyerSnapshot, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Payout) SetBuyer(v BuyerSnapshot)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Payout) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *Payout) GetPaymentMethod() PaymentMethodSnapshot`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Payout) GetPaymentMethodOk() (*PaymentMethodSnapshot, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Payout) SetPaymentMethod(v PaymentMethodSnapshot)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Payout) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


