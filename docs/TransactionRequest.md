# TransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The monetary amount for this transaction, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;.  If the &#x60;intent&#x60; is set to &#x60;capture&#x60;, an amount greater than zero must be supplied.  All gift card amounts are subtracted from this amount before the remainder is charged to the provided &#x60;payment_method&#x60;. | 
**Currency** | **string** | A supported ISO-4217 currency code.  For redirect requests, this value must match the one specified for &#x60;currency&#x60; in &#x60;payment_method&#x60;.  | 
**PaymentMethod** | Pointer to [**TransactionPaymentMethodRequest**](TransactionPaymentMethodRequest.md) |  | [optional] 
**AntiFraudFingerprint** | Pointer to **NullableString** | This field represents the fingerprint data to be passed to the active anti-fraud service. | [optional] 
**AsyncCapture** | Pointer to **bool** | Whether to capture the transaction asynchronously.  - When &#x60;async_capture&#x60; is &#x60;false&#x60; (default), the transaction is captured   in the same request. - When &#x60;async_capture&#x60; is &#x60;true&#x60;, the transaction is automatically   captured at a later time.  Redirect transactions are not affected by this flag.  This flag can only be set to &#x60;true&#x60; when &#x60;intent&#x60; is set to &#x60;capture&#x60;. | [optional] [default to false]
**BrowserInfo** | Pointer to [**NullableBrowserInfo**](BrowserInfo.md) | Information about the browser used by the buyer. | [optional] 
**BuyerExternalIdentifier** | Pointer to **string** | The &#x60;external_identifier&#x60; of the buyer to associate this payment method to. If this field is provided then the &#x60;buyer_id&#x60; field needs to be unset.  If a stored payment method or gift card is provided, then the buyer for that payment method needs to match the buyer for this field. | [optional] 
**BuyerId** | Pointer to **string** | The ID of the buyer to associate this payment method to. If this field is provided then the &#x60;buyer_external_identifier&#x60; field needs to be unset.  If a stored payment method or gift card is provided, then the buyer for that payment method needs to match the buyer for this field. | [optional] 
**Buyer** | Pointer to [**TransactionBuyerRequest**](TransactionBuyerRequest.md) |  | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 
**ConnectionOptions** | Pointer to [**NullableConnectionOptions**](ConnectionOptions.md) | Allows for passing optional configuration per connection to take advantage of connection specific features. When provided, the data is only passed to the target connection type to prevent sharing configuration across connections.  Please note that each of the keys this object are in kebab-case, for example &#x60;cybersource-anti-fraud&#x60; as they represent the ID of the connector. All the other keys will be snake case, for example &#x60;merchant_defined_data&#x60; or camel case to match an external API that the connector uses. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country where the transaction is processed. This is also used to filter the payment services that can process the transaction.  If this value is provided for redirect requests and it&#39;s not &#x60;null&#x60;, it must match the one specified for &#x60;country&#x60; in &#x60;payment_method&#x60;. Otherwise, the value specified for &#x60;country&#x60; in &#x60;payment_method&#x60; will be assumed implicitly.  | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the transaction against your own records. | [optional] 
**GiftCards** | Pointer to [**[]TransactionGiftCardRequest**](TransactionGiftCardRequest.md) | The optional gift card(s) to use for this transaction. At least one gift card is required if no other &#x60;payment_method&#x60; has been added. By default, only a maximum limit of 10 gift cards may be used in a single transaction. Please contact our team to change this limit. | [optional] 
**Intent** | Pointer to **string** | Defines the intent of this API call. This determines the desired initial state of the transaction.  * &#x60;authorize&#x60; - (Default) Optionally approves and then authorizes a transaction but does not capture the funds. * &#x60;capture&#x60; - Optionally approves and then authorizes and captures the funds of the transaction. | [optional] [default to "authorize"]
**IsSubsequentPayment** | Pointer to **bool** | Indicates whether the transaction represents a subsequent payment coming from a setup recurring payment. Please note there are some restrictions on how this flag may be used.  The flag can only be &#x60;false&#x60; (or not set) when the transaction meets one of the following criteria:  * It is not &#x60;merchant_initiated&#x60;. * &#x60;payment_source&#x60; is set to &#x60;card_on_file&#x60;.  The flag can only be set to &#x60;true&#x60; when the transaction meets one of the following criteria:  * It is not &#x60;merchant_initiated&#x60;. * &#x60;payment_source&#x60; is set to &#x60;recurring&#x60; or &#x60;installment&#x60; and &#x60;merchant_initiated&#x60; is set to &#x60;true&#x60;. * &#x60;payment_source&#x60; is set to &#x60;card_on_file&#x60;. | [optional] [default to false]
**MerchantInitiated** | Pointer to **bool** | Indicates whether the transaction was initiated by the merchant (true) or customer (false). | [optional] [default to false]
**Metadata** | Pointer to **map[string]string** | Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it. | [optional] 
**PaymentSource** | Pointer to **string** | The source of the transaction. Defaults to &#x60;ecommerce&#x60;. | [optional] 
**PreviousSchemeTransactionId** | Pointer to **NullableString** | A scheme&#39;s transaction identifier to use in connecting a merchant initiated transaction to a previous customer initiated transaction.  If not provided, and a qualifying customer initiated transaction has been previously made, then Gr4vy will populate this value with the identifier returned for that transaction.  e.g. the Visa Transaction Identifier, or Mastercard Trace ID. | [optional] 
**ShippingDetailsId** | Pointer to **NullableString** | The unique identifier of a set of shipping details stored for the buyer.  If provided, the created transaction will include a copy of the details at the point of transaction creation; i.e. it will not be affected by later changes to the detail in the database. | [optional] 
**StatementDescriptor** | Pointer to [**NullableStatementDescriptor**](StatementDescriptor.md) |  | [optional] 
**Store** | Pointer to **bool** | Whether or not to also try and store the payment method with us so that it can be used again for future use. This is only supported for payment methods that support this feature. There are also a few restrictions on how the flag may be set:  * The flag has to be set to &#x60;true&#x60; when the &#x60;payment_source&#x60; is set to &#x60;recurring&#x60; or &#x60;installment&#x60;, and &#x60;merchant_initiated&#x60; is set to &#x60;false&#x60;.  * The flag has to be set to &#x60;false&#x60; (or not set) when using a previously vaulted payment method. | [optional] [default to false]
**ThreeDSecureData** | Pointer to [**ThreeDSecureDataV1V2**](ThreeDSecureDataV1V2.md) |  | [optional] 
**PaymentServiceId** | Pointer to **NullableString** | The unique identifier of an existing payment service. When provided, the created transaction will be processed by the given payment service and any routing rules will be skipped. | [optional] 
**Airline** | Pointer to [**NullableAirline**](Airline.md) | The airline addendum data which describes the airline booking associated with this transaction. | [optional] 

## Methods

### NewTransactionRequest

`func NewTransactionRequest(amount int32, currency string, ) *TransactionRequest`

NewTransactionRequest instantiates a new TransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestWithDefaults

`func NewTransactionRequestWithDefaults() *TransactionRequest`

NewTransactionRequestWithDefaults instantiates a new TransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransactionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *TransactionRequest) GetPaymentMethod() TransactionPaymentMethodRequest`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionRequest) GetPaymentMethodOk() (*TransactionPaymentMethodRequest, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionRequest) SetPaymentMethod(v TransactionPaymentMethodRequest)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *TransactionRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetAntiFraudFingerprint

`func (o *TransactionRequest) GetAntiFraudFingerprint() string`

GetAntiFraudFingerprint returns the AntiFraudFingerprint field if non-nil, zero value otherwise.

### GetAntiFraudFingerprintOk

`func (o *TransactionRequest) GetAntiFraudFingerprintOk() (*string, bool)`

GetAntiFraudFingerprintOk returns a tuple with the AntiFraudFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudFingerprint

`func (o *TransactionRequest) SetAntiFraudFingerprint(v string)`

SetAntiFraudFingerprint sets AntiFraudFingerprint field to given value.

### HasAntiFraudFingerprint

`func (o *TransactionRequest) HasAntiFraudFingerprint() bool`

HasAntiFraudFingerprint returns a boolean if a field has been set.

### SetAntiFraudFingerprintNil

`func (o *TransactionRequest) SetAntiFraudFingerprintNil(b bool)`

 SetAntiFraudFingerprintNil sets the value for AntiFraudFingerprint to be an explicit nil

### UnsetAntiFraudFingerprint
`func (o *TransactionRequest) UnsetAntiFraudFingerprint()`

UnsetAntiFraudFingerprint ensures that no value is present for AntiFraudFingerprint, not even an explicit nil
### GetAsyncCapture

`func (o *TransactionRequest) GetAsyncCapture() bool`

GetAsyncCapture returns the AsyncCapture field if non-nil, zero value otherwise.

### GetAsyncCaptureOk

`func (o *TransactionRequest) GetAsyncCaptureOk() (*bool, bool)`

GetAsyncCaptureOk returns a tuple with the AsyncCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncCapture

`func (o *TransactionRequest) SetAsyncCapture(v bool)`

SetAsyncCapture sets AsyncCapture field to given value.

### HasAsyncCapture

`func (o *TransactionRequest) HasAsyncCapture() bool`

HasAsyncCapture returns a boolean if a field has been set.

### GetBrowserInfo

`func (o *TransactionRequest) GetBrowserInfo() BrowserInfo`

GetBrowserInfo returns the BrowserInfo field if non-nil, zero value otherwise.

### GetBrowserInfoOk

`func (o *TransactionRequest) GetBrowserInfoOk() (*BrowserInfo, bool)`

GetBrowserInfoOk returns a tuple with the BrowserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserInfo

`func (o *TransactionRequest) SetBrowserInfo(v BrowserInfo)`

SetBrowserInfo sets BrowserInfo field to given value.

### HasBrowserInfo

`func (o *TransactionRequest) HasBrowserInfo() bool`

HasBrowserInfo returns a boolean if a field has been set.

### SetBrowserInfoNil

`func (o *TransactionRequest) SetBrowserInfoNil(b bool)`

 SetBrowserInfoNil sets the value for BrowserInfo to be an explicit nil

### UnsetBrowserInfo
`func (o *TransactionRequest) UnsetBrowserInfo()`

UnsetBrowserInfo ensures that no value is present for BrowserInfo, not even an explicit nil
### GetBuyerExternalIdentifier

`func (o *TransactionRequest) GetBuyerExternalIdentifier() string`

GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field if non-nil, zero value otherwise.

### GetBuyerExternalIdentifierOk

`func (o *TransactionRequest) GetBuyerExternalIdentifierOk() (*string, bool)`

GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerExternalIdentifier

`func (o *TransactionRequest) SetBuyerExternalIdentifier(v string)`

SetBuyerExternalIdentifier sets BuyerExternalIdentifier field to given value.

### HasBuyerExternalIdentifier

`func (o *TransactionRequest) HasBuyerExternalIdentifier() bool`

HasBuyerExternalIdentifier returns a boolean if a field has been set.

### GetBuyerId

`func (o *TransactionRequest) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *TransactionRequest) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *TransactionRequest) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *TransactionRequest) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetBuyer

`func (o *TransactionRequest) GetBuyer() TransactionBuyerRequest`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *TransactionRequest) GetBuyerOk() (*TransactionBuyerRequest, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *TransactionRequest) SetBuyer(v TransactionBuyerRequest)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *TransactionRequest) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCartItems

`func (o *TransactionRequest) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *TransactionRequest) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *TransactionRequest) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *TransactionRequest) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### GetConnectionOptions

`func (o *TransactionRequest) GetConnectionOptions() ConnectionOptions`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *TransactionRequest) GetConnectionOptionsOk() (*ConnectionOptions, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *TransactionRequest) SetConnectionOptions(v ConnectionOptions)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *TransactionRequest) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### SetConnectionOptionsNil

`func (o *TransactionRequest) SetConnectionOptionsNil(b bool)`

 SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil

### UnsetConnectionOptions
`func (o *TransactionRequest) UnsetConnectionOptions()`

UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
### GetCountry

`func (o *TransactionRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TransactionRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TransactionRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TransactionRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *TransactionRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *TransactionRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetExternalIdentifier

`func (o *TransactionRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetGiftCards

`func (o *TransactionRequest) GetGiftCards() []TransactionGiftCardRequest`

GetGiftCards returns the GiftCards field if non-nil, zero value otherwise.

### GetGiftCardsOk

`func (o *TransactionRequest) GetGiftCardsOk() (*[]TransactionGiftCardRequest, bool)`

GetGiftCardsOk returns a tuple with the GiftCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCards

`func (o *TransactionRequest) SetGiftCards(v []TransactionGiftCardRequest)`

SetGiftCards sets GiftCards field to given value.

### HasGiftCards

`func (o *TransactionRequest) HasGiftCards() bool`

HasGiftCards returns a boolean if a field has been set.

### SetGiftCardsNil

`func (o *TransactionRequest) SetGiftCardsNil(b bool)`

 SetGiftCardsNil sets the value for GiftCards to be an explicit nil

### UnsetGiftCards
`func (o *TransactionRequest) UnsetGiftCards()`

UnsetGiftCards ensures that no value is present for GiftCards, not even an explicit nil
### GetIntent

`func (o *TransactionRequest) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *TransactionRequest) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *TransactionRequest) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *TransactionRequest) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetIsSubsequentPayment

`func (o *TransactionRequest) GetIsSubsequentPayment() bool`

GetIsSubsequentPayment returns the IsSubsequentPayment field if non-nil, zero value otherwise.

### GetIsSubsequentPaymentOk

`func (o *TransactionRequest) GetIsSubsequentPaymentOk() (*bool, bool)`

GetIsSubsequentPaymentOk returns a tuple with the IsSubsequentPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubsequentPayment

`func (o *TransactionRequest) SetIsSubsequentPayment(v bool)`

SetIsSubsequentPayment sets IsSubsequentPayment field to given value.

### HasIsSubsequentPayment

`func (o *TransactionRequest) HasIsSubsequentPayment() bool`

HasIsSubsequentPayment returns a boolean if a field has been set.

### GetMerchantInitiated

`func (o *TransactionRequest) GetMerchantInitiated() bool`

GetMerchantInitiated returns the MerchantInitiated field if non-nil, zero value otherwise.

### GetMerchantInitiatedOk

`func (o *TransactionRequest) GetMerchantInitiatedOk() (*bool, bool)`

GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInitiated

`func (o *TransactionRequest) SetMerchantInitiated(v bool)`

SetMerchantInitiated sets MerchantInitiated field to given value.

### HasMerchantInitiated

`func (o *TransactionRequest) HasMerchantInitiated() bool`

HasMerchantInitiated returns a boolean if a field has been set.

### GetMetadata

`func (o *TransactionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentSource

`func (o *TransactionRequest) GetPaymentSource() string`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *TransactionRequest) GetPaymentSourceOk() (*string, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *TransactionRequest) SetPaymentSource(v string)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *TransactionRequest) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetPreviousSchemeTransactionId

`func (o *TransactionRequest) GetPreviousSchemeTransactionId() string`

GetPreviousSchemeTransactionId returns the PreviousSchemeTransactionId field if non-nil, zero value otherwise.

### GetPreviousSchemeTransactionIdOk

`func (o *TransactionRequest) GetPreviousSchemeTransactionIdOk() (*string, bool)`

GetPreviousSchemeTransactionIdOk returns a tuple with the PreviousSchemeTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSchemeTransactionId

`func (o *TransactionRequest) SetPreviousSchemeTransactionId(v string)`

SetPreviousSchemeTransactionId sets PreviousSchemeTransactionId field to given value.

### HasPreviousSchemeTransactionId

`func (o *TransactionRequest) HasPreviousSchemeTransactionId() bool`

HasPreviousSchemeTransactionId returns a boolean if a field has been set.

### SetPreviousSchemeTransactionIdNil

`func (o *TransactionRequest) SetPreviousSchemeTransactionIdNil(b bool)`

 SetPreviousSchemeTransactionIdNil sets the value for PreviousSchemeTransactionId to be an explicit nil

### UnsetPreviousSchemeTransactionId
`func (o *TransactionRequest) UnsetPreviousSchemeTransactionId()`

UnsetPreviousSchemeTransactionId ensures that no value is present for PreviousSchemeTransactionId, not even an explicit nil
### GetShippingDetailsId

`func (o *TransactionRequest) GetShippingDetailsId() string`

GetShippingDetailsId returns the ShippingDetailsId field if non-nil, zero value otherwise.

### GetShippingDetailsIdOk

`func (o *TransactionRequest) GetShippingDetailsIdOk() (*string, bool)`

GetShippingDetailsIdOk returns a tuple with the ShippingDetailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDetailsId

`func (o *TransactionRequest) SetShippingDetailsId(v string)`

SetShippingDetailsId sets ShippingDetailsId field to given value.

### HasShippingDetailsId

`func (o *TransactionRequest) HasShippingDetailsId() bool`

HasShippingDetailsId returns a boolean if a field has been set.

### SetShippingDetailsIdNil

`func (o *TransactionRequest) SetShippingDetailsIdNil(b bool)`

 SetShippingDetailsIdNil sets the value for ShippingDetailsId to be an explicit nil

### UnsetShippingDetailsId
`func (o *TransactionRequest) UnsetShippingDetailsId()`

UnsetShippingDetailsId ensures that no value is present for ShippingDetailsId, not even an explicit nil
### GetStatementDescriptor

`func (o *TransactionRequest) GetStatementDescriptor() StatementDescriptor`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *TransactionRequest) GetStatementDescriptorOk() (*StatementDescriptor, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *TransactionRequest) SetStatementDescriptor(v StatementDescriptor)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *TransactionRequest) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### SetStatementDescriptorNil

`func (o *TransactionRequest) SetStatementDescriptorNil(b bool)`

 SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil

### UnsetStatementDescriptor
`func (o *TransactionRequest) UnsetStatementDescriptor()`

UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
### GetStore

`func (o *TransactionRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *TransactionRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *TransactionRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *TransactionRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetThreeDSecureData

`func (o *TransactionRequest) GetThreeDSecureData() ThreeDSecureDataV1V2`

GetThreeDSecureData returns the ThreeDSecureData field if non-nil, zero value otherwise.

### GetThreeDSecureDataOk

`func (o *TransactionRequest) GetThreeDSecureDataOk() (*ThreeDSecureDataV1V2, bool)`

GetThreeDSecureDataOk returns a tuple with the ThreeDSecureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureData

`func (o *TransactionRequest) SetThreeDSecureData(v ThreeDSecureDataV1V2)`

SetThreeDSecureData sets ThreeDSecureData field to given value.

### HasThreeDSecureData

`func (o *TransactionRequest) HasThreeDSecureData() bool`

HasThreeDSecureData returns a boolean if a field has been set.

### GetPaymentServiceId

`func (o *TransactionRequest) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *TransactionRequest) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *TransactionRequest) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *TransactionRequest) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### SetPaymentServiceIdNil

`func (o *TransactionRequest) SetPaymentServiceIdNil(b bool)`

 SetPaymentServiceIdNil sets the value for PaymentServiceId to be an explicit nil

### UnsetPaymentServiceId
`func (o *TransactionRequest) UnsetPaymentServiceId()`

UnsetPaymentServiceId ensures that no value is present for PaymentServiceId, not even an explicit nil
### GetAirline

`func (o *TransactionRequest) GetAirline() Airline`

GetAirline returns the Airline field if non-nil, zero value otherwise.

### GetAirlineOk

`func (o *TransactionRequest) GetAirlineOk() (*Airline, bool)`

GetAirlineOk returns a tuple with the Airline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirline

`func (o *TransactionRequest) SetAirline(v Airline)`

SetAirline sets Airline field to given value.

### HasAirline

`func (o *TransactionRequest) HasAirline() bool`

HasAirline returns a boolean if a field has been set.

### SetAirlineNil

`func (o *TransactionRequest) SetAirlineNil(b bool)`

 SetAirlineNil sets the value for Airline to be an explicit nil

### UnsetAirline
`func (o *TransactionRequest) UnsetAirline()`

UnsetAirline ensures that no value is present for Airline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


