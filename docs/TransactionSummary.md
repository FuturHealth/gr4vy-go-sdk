# TransactionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this transaction. | [optional] 
**Amount** | Pointer to **int32** | The authorized amount for this transaction. This can be more than the actual captured amount and part of this amount may be refunded. | [optional] 
**AuthorizedAmount** | Pointer to **int32** | The amount for this transaction that has been authorized for the &#x60;payment_method&#x60;. This can be less than the &#x60;amount&#x60; if gift cards were used. | [optional] 
**Buyer** | Pointer to [**NullableBuyerSnapshot**](BuyerSnapshot.md) | The buyer used for this transaction. | [optional] 
**CapturedAmount** | Pointer to **int32** | The captured amount for this transaction. This can be the full value of the &#x60;authorized_amount&#x60; or less. | [optional] 
**CheckoutSessionId** | Pointer to **NullableString** | The identifier for the checkout session this transaction is associated with. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Currency** | Pointer to **string** | The currency code for this transaction. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the transaction against your own records. | [optional] 
**GiftCardRedemptions** | Pointer to [**[]GiftCardRedemption**](GiftCardRedemption.md) | The gift cards redeemed for this transaction. | [optional] 
**InstrumentType** | Pointer to **NullableString** | The name of the instrument used to process the transaction.  | [optional] 
**Intent** | Pointer to **string** | The original &#x60;intent&#x60; used when the transaction was [created](#operation/authorize-new-transaction). | [optional] 
**MerchantAccountId** | Pointer to **string** | The ID of the merchant account to which this transaction belongs to. | [optional] 
**Method** | Pointer to **NullableString** |  | [optional] 
**PaymentMethod** | Pointer to [**NullablePaymentMethodSnapshot**](PaymentMethodSnapshot.md) | The payment method used for this transaction. | [optional] 
**PaymentService** | Pointer to [**NullablePaymentServiceSnapshot**](PaymentServiceSnapshot.md) | The payment service used for this transaction. | [optional] 
**PendingReview** | Pointer to **bool** | Whether a manual review is pending. | [optional] 
**RawResponseCode** | Pointer to **NullableString** | This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**RawResponseDescription** | Pointer to **NullableString** | This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**ReconciliationId** | Pointer to **string** | The base62 encoded transaction ID. This represents a shorter version of this transaction&#39;s &#x60;id&#x60; which is sent to payment services, anti-fraud services, and other connectors. You can use this ID to reconcile a payment service&#39;s transaction against our system.  This ID is sent instead of the transaction ID because not all services support 36 digit identifiers. | [optional] 
**RefundedAmount** | Pointer to **int32** | The refunded amount for this transaction. This can be the full value of the &#x60;captured_amount&#x60; or less. | [optional] 
**Status** | Pointer to **string** | The status of the transaction. The status may change over time as asynchronous processing events occur. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defines when the transaction was last updated. | [optional] 
**SettledCurrency** | Pointer to **NullableString** | The currency of this transaction&#39;s settlement in ISO 4217 three-letter code format. | [optional] 
**SettledAmount** | Pointer to **int32** | The net amount settled for this transaction. | [optional] 
**Settled** | Pointer to **bool** | Indicates whether this transaction has been settled. | [optional] 

## Methods

### NewTransactionSummary

`func NewTransactionSummary() *TransactionSummary`

NewTransactionSummary instantiates a new TransactionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSummaryWithDefaults

`func NewTransactionSummaryWithDefaults() *TransactionSummary`

NewTransactionSummaryWithDefaults instantiates a new TransactionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TransactionSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionSummary) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSummary) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSummary) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionSummary) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthorizedAmount

`func (o *TransactionSummary) GetAuthorizedAmount() int32`

GetAuthorizedAmount returns the AuthorizedAmount field if non-nil, zero value otherwise.

### GetAuthorizedAmountOk

`func (o *TransactionSummary) GetAuthorizedAmountOk() (*int32, bool)`

GetAuthorizedAmountOk returns a tuple with the AuthorizedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAmount

`func (o *TransactionSummary) SetAuthorizedAmount(v int32)`

SetAuthorizedAmount sets AuthorizedAmount field to given value.

### HasAuthorizedAmount

`func (o *TransactionSummary) HasAuthorizedAmount() bool`

HasAuthorizedAmount returns a boolean if a field has been set.

### GetBuyer

`func (o *TransactionSummary) GetBuyer() BuyerSnapshot`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *TransactionSummary) GetBuyerOk() (*BuyerSnapshot, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *TransactionSummary) SetBuyer(v BuyerSnapshot)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *TransactionSummary) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *TransactionSummary) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *TransactionSummary) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetCapturedAmount

`func (o *TransactionSummary) GetCapturedAmount() int32`

GetCapturedAmount returns the CapturedAmount field if non-nil, zero value otherwise.

### GetCapturedAmountOk

`func (o *TransactionSummary) GetCapturedAmountOk() (*int32, bool)`

GetCapturedAmountOk returns a tuple with the CapturedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAmount

`func (o *TransactionSummary) SetCapturedAmount(v int32)`

SetCapturedAmount sets CapturedAmount field to given value.

### HasCapturedAmount

`func (o *TransactionSummary) HasCapturedAmount() bool`

HasCapturedAmount returns a boolean if a field has been set.

### GetCheckoutSessionId

`func (o *TransactionSummary) GetCheckoutSessionId() string`

GetCheckoutSessionId returns the CheckoutSessionId field if non-nil, zero value otherwise.

### GetCheckoutSessionIdOk

`func (o *TransactionSummary) GetCheckoutSessionIdOk() (*string, bool)`

GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSessionId

`func (o *TransactionSummary) SetCheckoutSessionId(v string)`

SetCheckoutSessionId sets CheckoutSessionId field to given value.

### HasCheckoutSessionId

`func (o *TransactionSummary) HasCheckoutSessionId() bool`

HasCheckoutSessionId returns a boolean if a field has been set.

### SetCheckoutSessionIdNil

`func (o *TransactionSummary) SetCheckoutSessionIdNil(b bool)`

 SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil

### UnsetCheckoutSessionId
`func (o *TransactionSummary) UnsetCheckoutSessionId()`

UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
### GetCountry

`func (o *TransactionSummary) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TransactionSummary) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TransactionSummary) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TransactionSummary) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *TransactionSummary) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *TransactionSummary) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCreatedAt

`func (o *TransactionSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *TransactionSummary) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionSummary) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionSummary) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionSummary) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionSummary) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionSummary) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetGiftCardRedemptions

`func (o *TransactionSummary) GetGiftCardRedemptions() []GiftCardRedemption`

GetGiftCardRedemptions returns the GiftCardRedemptions field if non-nil, zero value otherwise.

### GetGiftCardRedemptionsOk

`func (o *TransactionSummary) GetGiftCardRedemptionsOk() (*[]GiftCardRedemption, bool)`

GetGiftCardRedemptionsOk returns a tuple with the GiftCardRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardRedemptions

`func (o *TransactionSummary) SetGiftCardRedemptions(v []GiftCardRedemption)`

SetGiftCardRedemptions sets GiftCardRedemptions field to given value.

### HasGiftCardRedemptions

`func (o *TransactionSummary) HasGiftCardRedemptions() bool`

HasGiftCardRedemptions returns a boolean if a field has been set.

### GetInstrumentType

`func (o *TransactionSummary) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *TransactionSummary) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *TransactionSummary) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *TransactionSummary) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### SetInstrumentTypeNil

`func (o *TransactionSummary) SetInstrumentTypeNil(b bool)`

 SetInstrumentTypeNil sets the value for InstrumentType to be an explicit nil

### UnsetInstrumentType
`func (o *TransactionSummary) UnsetInstrumentType()`

UnsetInstrumentType ensures that no value is present for InstrumentType, not even an explicit nil
### GetIntent

`func (o *TransactionSummary) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *TransactionSummary) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *TransactionSummary) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *TransactionSummary) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *TransactionSummary) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *TransactionSummary) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *TransactionSummary) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *TransactionSummary) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetMethod

`func (o *TransactionSummary) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionSummary) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionSummary) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TransactionSummary) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *TransactionSummary) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *TransactionSummary) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetPaymentMethod

`func (o *TransactionSummary) GetPaymentMethod() PaymentMethodSnapshot`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionSummary) GetPaymentMethodOk() (*PaymentMethodSnapshot, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionSummary) SetPaymentMethod(v PaymentMethodSnapshot)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *TransactionSummary) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *TransactionSummary) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *TransactionSummary) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetPaymentService

`func (o *TransactionSummary) GetPaymentService() PaymentServiceSnapshot`

GetPaymentService returns the PaymentService field if non-nil, zero value otherwise.

### GetPaymentServiceOk

`func (o *TransactionSummary) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool)`

GetPaymentServiceOk returns a tuple with the PaymentService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentService

`func (o *TransactionSummary) SetPaymentService(v PaymentServiceSnapshot)`

SetPaymentService sets PaymentService field to given value.

### HasPaymentService

`func (o *TransactionSummary) HasPaymentService() bool`

HasPaymentService returns a boolean if a field has been set.

### SetPaymentServiceNil

`func (o *TransactionSummary) SetPaymentServiceNil(b bool)`

 SetPaymentServiceNil sets the value for PaymentService to be an explicit nil

### UnsetPaymentService
`func (o *TransactionSummary) UnsetPaymentService()`

UnsetPaymentService ensures that no value is present for PaymentService, not even an explicit nil
### GetPendingReview

`func (o *TransactionSummary) GetPendingReview() bool`

GetPendingReview returns the PendingReview field if non-nil, zero value otherwise.

### GetPendingReviewOk

`func (o *TransactionSummary) GetPendingReviewOk() (*bool, bool)`

GetPendingReviewOk returns a tuple with the PendingReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReview

`func (o *TransactionSummary) SetPendingReview(v bool)`

SetPendingReview sets PendingReview field to given value.

### HasPendingReview

`func (o *TransactionSummary) HasPendingReview() bool`

HasPendingReview returns a boolean if a field has been set.

### GetRawResponseCode

`func (o *TransactionSummary) GetRawResponseCode() string`

GetRawResponseCode returns the RawResponseCode field if non-nil, zero value otherwise.

### GetRawResponseCodeOk

`func (o *TransactionSummary) GetRawResponseCodeOk() (*string, bool)`

GetRawResponseCodeOk returns a tuple with the RawResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseCode

`func (o *TransactionSummary) SetRawResponseCode(v string)`

SetRawResponseCode sets RawResponseCode field to given value.

### HasRawResponseCode

`func (o *TransactionSummary) HasRawResponseCode() bool`

HasRawResponseCode returns a boolean if a field has been set.

### SetRawResponseCodeNil

`func (o *TransactionSummary) SetRawResponseCodeNil(b bool)`

 SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil

### UnsetRawResponseCode
`func (o *TransactionSummary) UnsetRawResponseCode()`

UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
### GetRawResponseDescription

`func (o *TransactionSummary) GetRawResponseDescription() string`

GetRawResponseDescription returns the RawResponseDescription field if non-nil, zero value otherwise.

### GetRawResponseDescriptionOk

`func (o *TransactionSummary) GetRawResponseDescriptionOk() (*string, bool)`

GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseDescription

`func (o *TransactionSummary) SetRawResponseDescription(v string)`

SetRawResponseDescription sets RawResponseDescription field to given value.

### HasRawResponseDescription

`func (o *TransactionSummary) HasRawResponseDescription() bool`

HasRawResponseDescription returns a boolean if a field has been set.

### SetRawResponseDescriptionNil

`func (o *TransactionSummary) SetRawResponseDescriptionNil(b bool)`

 SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil

### UnsetRawResponseDescription
`func (o *TransactionSummary) UnsetRawResponseDescription()`

UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
### GetReconciliationId

`func (o *TransactionSummary) GetReconciliationId() string`

GetReconciliationId returns the ReconciliationId field if non-nil, zero value otherwise.

### GetReconciliationIdOk

`func (o *TransactionSummary) GetReconciliationIdOk() (*string, bool)`

GetReconciliationIdOk returns a tuple with the ReconciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationId

`func (o *TransactionSummary) SetReconciliationId(v string)`

SetReconciliationId sets ReconciliationId field to given value.

### HasReconciliationId

`func (o *TransactionSummary) HasReconciliationId() bool`

HasReconciliationId returns a boolean if a field has been set.

### GetRefundedAmount

`func (o *TransactionSummary) GetRefundedAmount() int32`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *TransactionSummary) GetRefundedAmountOk() (*int32, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *TransactionSummary) SetRefundedAmount(v int32)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *TransactionSummary) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TransactionSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSettledCurrency

`func (o *TransactionSummary) GetSettledCurrency() string`

GetSettledCurrency returns the SettledCurrency field if non-nil, zero value otherwise.

### GetSettledCurrencyOk

`func (o *TransactionSummary) GetSettledCurrencyOk() (*string, bool)`

GetSettledCurrencyOk returns a tuple with the SettledCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledCurrency

`func (o *TransactionSummary) SetSettledCurrency(v string)`

SetSettledCurrency sets SettledCurrency field to given value.

### HasSettledCurrency

`func (o *TransactionSummary) HasSettledCurrency() bool`

HasSettledCurrency returns a boolean if a field has been set.

### SetSettledCurrencyNil

`func (o *TransactionSummary) SetSettledCurrencyNil(b bool)`

 SetSettledCurrencyNil sets the value for SettledCurrency to be an explicit nil

### UnsetSettledCurrency
`func (o *TransactionSummary) UnsetSettledCurrency()`

UnsetSettledCurrency ensures that no value is present for SettledCurrency, not even an explicit nil
### GetSettledAmount

`func (o *TransactionSummary) GetSettledAmount() int32`

GetSettledAmount returns the SettledAmount field if non-nil, zero value otherwise.

### GetSettledAmountOk

`func (o *TransactionSummary) GetSettledAmountOk() (*int32, bool)`

GetSettledAmountOk returns a tuple with the SettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAmount

`func (o *TransactionSummary) SetSettledAmount(v int32)`

SetSettledAmount sets SettledAmount field to given value.

### HasSettledAmount

`func (o *TransactionSummary) HasSettledAmount() bool`

HasSettledAmount returns a boolean if a field has been set.

### GetSettled

`func (o *TransactionSummary) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *TransactionSummary) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *TransactionSummary) SetSettled(v bool)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *TransactionSummary) HasSettled() bool`

HasSettled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


