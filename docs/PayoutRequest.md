# PayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | The amount to payout. | 
**Currency** | **string** | The ISO-4217 currency code for the payout. | 
**PaymentServiceId** | **string** | The ID of the payment service to use for the payout. | 
**Category** | Pointer to **string** | The type of payout to process. | [optional] [default to "online_gambling"]
**ExternalIdentifier** | Pointer to **NullableString** | A value that can be used to match the payout against your own records. | [optional] 
**BuyerId** | Pointer to **NullableString** | The ID of the buyer to send the payout. | [optional] 
**BuyerExternalIdentifier** | Pointer to **NullableString** | The &#x60;external_identifier&#x60; of the buyer to send this payout to. | [optional] 
**Buyer** | Pointer to [**NullableTransactionBuyerRequest**](TransactionBuyerRequest.md) | Inline buyer details for the payout. | [optional] 
**Merchant** | Pointer to [**NullableMerchantRequest**](MerchantRequest.md) | Merchant information for the source of the payout. | [optional] 
**PaymentMethod** | [**PayoutPaymentMethodRequest**](PayoutPaymentMethodRequest.md) |  | 
**ConnectionOptions** | Pointer to [**NullablePayoutConnectionOptionsRequest**](PayoutConnectionOptionsRequest.md) | Optional fields for processing payouts on specific payment services. | [optional] 

## Methods

### NewPayoutRequest

`func NewPayoutRequest(amount float32, currency string, paymentServiceId string, paymentMethod PayoutPaymentMethodRequest, ) *PayoutRequest`

NewPayoutRequest instantiates a new PayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutRequestWithDefaults

`func NewPayoutRequestWithDefaults() *PayoutRequest`

NewPayoutRequestWithDefaults instantiates a new PayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PayoutRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PayoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentServiceId

`func (o *PayoutRequest) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PayoutRequest) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PayoutRequest) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.


### GetCategory

`func (o *PayoutRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PayoutRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PayoutRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PayoutRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *PayoutRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PayoutRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PayoutRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PayoutRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PayoutRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PayoutRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetBuyerId

`func (o *PayoutRequest) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *PayoutRequest) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *PayoutRequest) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *PayoutRequest) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### SetBuyerIdNil

`func (o *PayoutRequest) SetBuyerIdNil(b bool)`

 SetBuyerIdNil sets the value for BuyerId to be an explicit nil

### UnsetBuyerId
`func (o *PayoutRequest) UnsetBuyerId()`

UnsetBuyerId ensures that no value is present for BuyerId, not even an explicit nil
### GetBuyerExternalIdentifier

`func (o *PayoutRequest) GetBuyerExternalIdentifier() string`

GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field if non-nil, zero value otherwise.

### GetBuyerExternalIdentifierOk

`func (o *PayoutRequest) GetBuyerExternalIdentifierOk() (*string, bool)`

GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerExternalIdentifier

`func (o *PayoutRequest) SetBuyerExternalIdentifier(v string)`

SetBuyerExternalIdentifier sets BuyerExternalIdentifier field to given value.

### HasBuyerExternalIdentifier

`func (o *PayoutRequest) HasBuyerExternalIdentifier() bool`

HasBuyerExternalIdentifier returns a boolean if a field has been set.

### SetBuyerExternalIdentifierNil

`func (o *PayoutRequest) SetBuyerExternalIdentifierNil(b bool)`

 SetBuyerExternalIdentifierNil sets the value for BuyerExternalIdentifier to be an explicit nil

### UnsetBuyerExternalIdentifier
`func (o *PayoutRequest) UnsetBuyerExternalIdentifier()`

UnsetBuyerExternalIdentifier ensures that no value is present for BuyerExternalIdentifier, not even an explicit nil
### GetBuyer

`func (o *PayoutRequest) GetBuyer() TransactionBuyerRequest`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *PayoutRequest) GetBuyerOk() (*TransactionBuyerRequest, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *PayoutRequest) SetBuyer(v TransactionBuyerRequest)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *PayoutRequest) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *PayoutRequest) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *PayoutRequest) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetMerchant

`func (o *PayoutRequest) GetMerchant() MerchantRequest`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *PayoutRequest) GetMerchantOk() (*MerchantRequest, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *PayoutRequest) SetMerchant(v MerchantRequest)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *PayoutRequest) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### SetMerchantNil

`func (o *PayoutRequest) SetMerchantNil(b bool)`

 SetMerchantNil sets the value for Merchant to be an explicit nil

### UnsetMerchant
`func (o *PayoutRequest) UnsetMerchant()`

UnsetMerchant ensures that no value is present for Merchant, not even an explicit nil
### GetPaymentMethod

`func (o *PayoutRequest) GetPaymentMethod() PayoutPaymentMethodRequest`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PayoutRequest) GetPaymentMethodOk() (*PayoutPaymentMethodRequest, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PayoutRequest) SetPaymentMethod(v PayoutPaymentMethodRequest)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetConnectionOptions

`func (o *PayoutRequest) GetConnectionOptions() PayoutConnectionOptionsRequest`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *PayoutRequest) GetConnectionOptionsOk() (*PayoutConnectionOptionsRequest, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *PayoutRequest) SetConnectionOptions(v PayoutConnectionOptionsRequest)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *PayoutRequest) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### SetConnectionOptionsNil

`func (o *PayoutRequest) SetConnectionOptionsNil(b bool)`

 SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil

### UnsetConnectionOptions
`func (o *PayoutRequest) UnsetConnectionOptions()`

UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


