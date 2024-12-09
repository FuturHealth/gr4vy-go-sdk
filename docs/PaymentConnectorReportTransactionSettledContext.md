# PaymentConnectorReportTransactionSettledContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementId** | Pointer to **string** | The unique ID of this transaction&#39;s settlement. | [optional] 
**PostedAt** | Pointer to **time.Time** | The date and time when this transaction was settled. | [optional] 
**IngestedAt** | Pointer to **time.Time** | The date and time when information about this transaction&#39;s settlement was ingested into our system. | [optional] 
**Currency** | Pointer to **string** | The currency of this transaction&#39;s settlement in ISO 4217 three-letter code format. | [optional] 
**Amount** | Pointer to **int32** | The net amount settled for this transaction. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat32** | The exchange rate used to convert amounts from the processing currency to the settlement currency. | [optional] 
**Commission** | Pointer to **int32** | The total commission paid on this transaction, expressed in &#x60;currency&#x60;. | [optional] 
**Interchange** | Pointer to **NullableInt32** | The interchange fee for this transaction, expressed in &#x60;currency&#x60;. | [optional] 
**Markup** | Pointer to **NullableInt32** | The markup applied to this transaction by the acquirer, expressed in &#x60;currency&#x60;. | [optional] 
**SchemeFee** | Pointer to **NullableInt32** | The scheme fee for this transaction, expressed in &#x60;currency&#x60;. | [optional] 
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service used. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The payment service definition used. | [optional] 
**PaymentServiceDisplayName** | Pointer to **string** | The display name of the payment service used. | [optional] 
**PaymentServiceReportId** | Pointer to **string** | The ID of the payment service report containing this transaction&#39;s settlement. | [optional] 
**PaymentServiceReportFileIds** | Pointer to **[]string** | The list of payment service report file IDs that make up the payment service report containing this transaction&#39;s settlement. | [optional] 
**PaymentServiceTransactionId** | Pointer to **string** | The external ID of the transaction as set by the payment service. | [optional] 

## Methods

### NewPaymentConnectorReportTransactionSettledContext

`func NewPaymentConnectorReportTransactionSettledContext() *PaymentConnectorReportTransactionSettledContext`

NewPaymentConnectorReportTransactionSettledContext instantiates a new PaymentConnectorReportTransactionSettledContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorReportTransactionSettledContextWithDefaults

`func NewPaymentConnectorReportTransactionSettledContextWithDefaults() *PaymentConnectorReportTransactionSettledContext`

NewPaymentConnectorReportTransactionSettledContextWithDefaults instantiates a new PaymentConnectorReportTransactionSettledContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementId

`func (o *PaymentConnectorReportTransactionSettledContext) GetSettlementId() string`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetSettlementIdOk() (*string, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *PaymentConnectorReportTransactionSettledContext) SetSettlementId(v string)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *PaymentConnectorReportTransactionSettledContext) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetPostedAt

`func (o *PaymentConnectorReportTransactionSettledContext) GetPostedAt() time.Time`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPostedAtOk() (*time.Time, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *PaymentConnectorReportTransactionSettledContext) SetPostedAt(v time.Time)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *PaymentConnectorReportTransactionSettledContext) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetIngestedAt

`func (o *PaymentConnectorReportTransactionSettledContext) GetIngestedAt() time.Time`

GetIngestedAt returns the IngestedAt field if non-nil, zero value otherwise.

### GetIngestedAtOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetIngestedAtOk() (*time.Time, bool)`

GetIngestedAtOk returns a tuple with the IngestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedAt

`func (o *PaymentConnectorReportTransactionSettledContext) SetIngestedAt(v time.Time)`

SetIngestedAt sets IngestedAt field to given value.

### HasIngestedAt

`func (o *PaymentConnectorReportTransactionSettledContext) HasIngestedAt() bool`

HasIngestedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentConnectorReportTransactionSettledContext) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentConnectorReportTransactionSettledContext) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentConnectorReportTransactionSettledContext) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentConnectorReportTransactionSettledContext) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentConnectorReportTransactionSettledContext) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentConnectorReportTransactionSettledContext) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *PaymentConnectorReportTransactionSettledContext) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *PaymentConnectorReportTransactionSettledContext) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *PaymentConnectorReportTransactionSettledContext) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *PaymentConnectorReportTransactionSettledContext) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *PaymentConnectorReportTransactionSettledContext) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCommission

`func (o *PaymentConnectorReportTransactionSettledContext) GetCommission() int32`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetCommissionOk() (*int32, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *PaymentConnectorReportTransactionSettledContext) SetCommission(v int32)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *PaymentConnectorReportTransactionSettledContext) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetInterchange

`func (o *PaymentConnectorReportTransactionSettledContext) GetInterchange() int32`

GetInterchange returns the Interchange field if non-nil, zero value otherwise.

### GetInterchangeOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetInterchangeOk() (*int32, bool)`

GetInterchangeOk returns a tuple with the Interchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterchange

`func (o *PaymentConnectorReportTransactionSettledContext) SetInterchange(v int32)`

SetInterchange sets Interchange field to given value.

### HasInterchange

`func (o *PaymentConnectorReportTransactionSettledContext) HasInterchange() bool`

HasInterchange returns a boolean if a field has been set.

### SetInterchangeNil

`func (o *PaymentConnectorReportTransactionSettledContext) SetInterchangeNil(b bool)`

 SetInterchangeNil sets the value for Interchange to be an explicit nil

### UnsetInterchange
`func (o *PaymentConnectorReportTransactionSettledContext) UnsetInterchange()`

UnsetInterchange ensures that no value is present for Interchange, not even an explicit nil
### GetMarkup

`func (o *PaymentConnectorReportTransactionSettledContext) GetMarkup() int32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetMarkupOk() (*int32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *PaymentConnectorReportTransactionSettledContext) SetMarkup(v int32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *PaymentConnectorReportTransactionSettledContext) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### SetMarkupNil

`func (o *PaymentConnectorReportTransactionSettledContext) SetMarkupNil(b bool)`

 SetMarkupNil sets the value for Markup to be an explicit nil

### UnsetMarkup
`func (o *PaymentConnectorReportTransactionSettledContext) UnsetMarkup()`

UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
### GetSchemeFee

`func (o *PaymentConnectorReportTransactionSettledContext) GetSchemeFee() int32`

GetSchemeFee returns the SchemeFee field if non-nil, zero value otherwise.

### GetSchemeFeeOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetSchemeFeeOk() (*int32, bool)`

GetSchemeFeeOk returns a tuple with the SchemeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeFee

`func (o *PaymentConnectorReportTransactionSettledContext) SetSchemeFee(v int32)`

SetSchemeFee sets SchemeFee field to given value.

### HasSchemeFee

`func (o *PaymentConnectorReportTransactionSettledContext) HasSchemeFee() bool`

HasSchemeFee returns a boolean if a field has been set.

### SetSchemeFeeNil

`func (o *PaymentConnectorReportTransactionSettledContext) SetSchemeFeeNil(b bool)`

 SetSchemeFeeNil sets the value for SchemeFee to be an explicit nil

### UnsetSchemeFee
`func (o *PaymentConnectorReportTransactionSettledContext) UnsetSchemeFee()`

UnsetSchemeFee ensures that no value is present for SchemeFee, not even an explicit nil
### GetPaymentServiceId

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorReportTransactionSettledContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorReportTransactionSettledContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorReportTransactionSettledContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorReportTransactionSettledContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetPaymentServiceDisplayName

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorReportTransactionSettledContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorReportTransactionSettledContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### GetPaymentServiceReportId

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceReportId() string`

GetPaymentServiceReportId returns the PaymentServiceReportId field if non-nil, zero value otherwise.

### GetPaymentServiceReportIdOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceReportIdOk() (*string, bool)`

GetPaymentServiceReportIdOk returns a tuple with the PaymentServiceReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceReportId

`func (o *PaymentConnectorReportTransactionSettledContext) SetPaymentServiceReportId(v string)`

SetPaymentServiceReportId sets PaymentServiceReportId field to given value.

### HasPaymentServiceReportId

`func (o *PaymentConnectorReportTransactionSettledContext) HasPaymentServiceReportId() bool`

HasPaymentServiceReportId returns a boolean if a field has been set.

### GetPaymentServiceReportFileIds

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceReportFileIds() []string`

GetPaymentServiceReportFileIds returns the PaymentServiceReportFileIds field if non-nil, zero value otherwise.

### GetPaymentServiceReportFileIdsOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceReportFileIdsOk() (*[]string, bool)`

GetPaymentServiceReportFileIdsOk returns a tuple with the PaymentServiceReportFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceReportFileIds

`func (o *PaymentConnectorReportTransactionSettledContext) SetPaymentServiceReportFileIds(v []string)`

SetPaymentServiceReportFileIds sets PaymentServiceReportFileIds field to given value.

### HasPaymentServiceReportFileIds

`func (o *PaymentConnectorReportTransactionSettledContext) HasPaymentServiceReportFileIds() bool`

HasPaymentServiceReportFileIds returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *PaymentConnectorReportTransactionSettledContext) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *PaymentConnectorReportTransactionSettledContext) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *PaymentConnectorReportTransactionSettledContext) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


