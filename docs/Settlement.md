# Settlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;settlement&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this settlement. | [optional] 
**MerchantAccountId** | Pointer to **string** | The ID of the merchant account to which this settlement belongs to. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this settlement was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this settlement was last updated. | [optional] 
**PostedAt** | Pointer to **time.Time** | The date and time when this settlement occurred. | [optional] 
**IngestedAt** | Pointer to **time.Time** | The date and time when information about this settlement was ingested into our system. | [optional] 
**Currency** | Pointer to **string** | The currency code of this settlement in ISO 4217 three-letter code format. | [optional] 
**Amount** | Pointer to **int32** | The net amount of this settlement. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat32** | The exchange rate used to convert amounts from the processing currency to the settlement currency. | [optional] 
**Commission** | Pointer to **int32** | The total commission of this settlement, expressed in &#x60;currency&#x60;. | [optional] 
**Interchange** | Pointer to **NullableInt32** | The interchange fee of this settlement, expressed in &#x60;currency&#x60;. | [optional] 
**Markup** | Pointer to **NullableInt32** | The markup of this settlement by the acquirer, expressed in &#x60;currency&#x60;. | [optional] 
**SchemeFee** | Pointer to **NullableInt32** | The scheme fee of this settlement, expressed in &#x60;currency&#x60;. | [optional] 
**PaymentServiceReportId** | Pointer to **string** | The ID of the payment service report containing this settlement. | [optional] 
**PaymentServiceReportFileIds** | Pointer to **[]string** | The list of payment service report file IDs that make up the payment service report containing this settlement. | [optional] 
**TransactionId** | Pointer to **string** | The ID of the transaction associated with this settlement. | [optional] 

## Methods

### NewSettlement

`func NewSettlement() *Settlement`

NewSettlement instantiates a new Settlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementWithDefaults

`func NewSettlementWithDefaults() *Settlement`

NewSettlementWithDefaults instantiates a new Settlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Settlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Settlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Settlement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Settlement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Settlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Settlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Settlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Settlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *Settlement) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *Settlement) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *Settlement) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *Settlement) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Settlement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Settlement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Settlement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Settlement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Settlement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Settlement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Settlement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Settlement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPostedAt

`func (o *Settlement) GetPostedAt() time.Time`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *Settlement) GetPostedAtOk() (*time.Time, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *Settlement) SetPostedAt(v time.Time)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *Settlement) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetIngestedAt

`func (o *Settlement) GetIngestedAt() time.Time`

GetIngestedAt returns the IngestedAt field if non-nil, zero value otherwise.

### GetIngestedAtOk

`func (o *Settlement) GetIngestedAtOk() (*time.Time, bool)`

GetIngestedAtOk returns a tuple with the IngestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedAt

`func (o *Settlement) SetIngestedAt(v time.Time)`

SetIngestedAt sets IngestedAt field to given value.

### HasIngestedAt

`func (o *Settlement) HasIngestedAt() bool`

HasIngestedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Settlement) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Settlement) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Settlement) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Settlement) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *Settlement) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Settlement) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Settlement) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Settlement) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *Settlement) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Settlement) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Settlement) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Settlement) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *Settlement) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *Settlement) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCommission

`func (o *Settlement) GetCommission() int32`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *Settlement) GetCommissionOk() (*int32, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *Settlement) SetCommission(v int32)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *Settlement) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetInterchange

`func (o *Settlement) GetInterchange() int32`

GetInterchange returns the Interchange field if non-nil, zero value otherwise.

### GetInterchangeOk

`func (o *Settlement) GetInterchangeOk() (*int32, bool)`

GetInterchangeOk returns a tuple with the Interchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterchange

`func (o *Settlement) SetInterchange(v int32)`

SetInterchange sets Interchange field to given value.

### HasInterchange

`func (o *Settlement) HasInterchange() bool`

HasInterchange returns a boolean if a field has been set.

### SetInterchangeNil

`func (o *Settlement) SetInterchangeNil(b bool)`

 SetInterchangeNil sets the value for Interchange to be an explicit nil

### UnsetInterchange
`func (o *Settlement) UnsetInterchange()`

UnsetInterchange ensures that no value is present for Interchange, not even an explicit nil
### GetMarkup

`func (o *Settlement) GetMarkup() int32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *Settlement) GetMarkupOk() (*int32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *Settlement) SetMarkup(v int32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *Settlement) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### SetMarkupNil

`func (o *Settlement) SetMarkupNil(b bool)`

 SetMarkupNil sets the value for Markup to be an explicit nil

### UnsetMarkup
`func (o *Settlement) UnsetMarkup()`

UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
### GetSchemeFee

`func (o *Settlement) GetSchemeFee() int32`

GetSchemeFee returns the SchemeFee field if non-nil, zero value otherwise.

### GetSchemeFeeOk

`func (o *Settlement) GetSchemeFeeOk() (*int32, bool)`

GetSchemeFeeOk returns a tuple with the SchemeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeFee

`func (o *Settlement) SetSchemeFee(v int32)`

SetSchemeFee sets SchemeFee field to given value.

### HasSchemeFee

`func (o *Settlement) HasSchemeFee() bool`

HasSchemeFee returns a boolean if a field has been set.

### SetSchemeFeeNil

`func (o *Settlement) SetSchemeFeeNil(b bool)`

 SetSchemeFeeNil sets the value for SchemeFee to be an explicit nil

### UnsetSchemeFee
`func (o *Settlement) UnsetSchemeFee()`

UnsetSchemeFee ensures that no value is present for SchemeFee, not even an explicit nil
### GetPaymentServiceReportId

`func (o *Settlement) GetPaymentServiceReportId() string`

GetPaymentServiceReportId returns the PaymentServiceReportId field if non-nil, zero value otherwise.

### GetPaymentServiceReportIdOk

`func (o *Settlement) GetPaymentServiceReportIdOk() (*string, bool)`

GetPaymentServiceReportIdOk returns a tuple with the PaymentServiceReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceReportId

`func (o *Settlement) SetPaymentServiceReportId(v string)`

SetPaymentServiceReportId sets PaymentServiceReportId field to given value.

### HasPaymentServiceReportId

`func (o *Settlement) HasPaymentServiceReportId() bool`

HasPaymentServiceReportId returns a boolean if a field has been set.

### GetPaymentServiceReportFileIds

`func (o *Settlement) GetPaymentServiceReportFileIds() []string`

GetPaymentServiceReportFileIds returns the PaymentServiceReportFileIds field if non-nil, zero value otherwise.

### GetPaymentServiceReportFileIdsOk

`func (o *Settlement) GetPaymentServiceReportFileIdsOk() (*[]string, bool)`

GetPaymentServiceReportFileIdsOk returns a tuple with the PaymentServiceReportFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceReportFileIds

`func (o *Settlement) SetPaymentServiceReportFileIds(v []string)`

SetPaymentServiceReportFileIds sets PaymentServiceReportFileIds field to given value.

### HasPaymentServiceReportFileIds

`func (o *Settlement) HasPaymentServiceReportFileIds() bool`

HasPaymentServiceReportFileIds returns a boolean if a field has been set.

### GetTransactionId

`func (o *Settlement) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Settlement) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Settlement) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Settlement) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


