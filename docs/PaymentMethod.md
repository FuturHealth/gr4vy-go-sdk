# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**AdditionalSchemes** | Pointer to **[]string** | Additional schemes of the card. Only applies to card payment methods. | [optional] 
**ApprovalTarget** | Pointer to **NullableString** | The browser target that an approval URL must be opened in. If &#x60;any&#x60; or &#x60;null&#x60;, then there is no specific requirement. | [optional] 
**ApprovalUrl** | Pointer to **NullableString** | The optional URL that the buyer needs to be redirected to to further authorize their payment. | [optional] 
**Buyer** | Pointer to [**NullableBuyer**](Buyer.md) | The optional buyer for which this payment method has been stored. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used in multiple countries. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this payment method was first created in our system. | [optional] 
**Currency** | Pointer to **NullableString** | The ISO-4217 currency code that this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used for multiple currencies. | [optional] 
**Details** | Pointer to [**PaymentMethodDetailsCard**](PaymentMethodDetailsCard.md) |  | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date for the payment method. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the payment method against your own records. | [optional] 
**HasReplacement** | Pointer to **bool** | Whether this card has a pending replacement that hasn&#39;t been applied yet.  When the Account Updater determines that new card details are available, existing details are not changed immediately, but this field is set to &#x60;true&#x60;. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * &#x60;canceled_payment_method&#x60;     * &#x60;expired_payment_method&#x60;     * &#x60;unavailable_payment_method&#x60;     * &#x60;unknown_payment_method&#x60;  When the replacement is applied, this field is set to &#x60;false&#x60;. For non-card payment methods, the value of this field is always set to &#x60;false&#x60;. | [optional] 
**Label** | Pointer to **NullableString** | A label for the card or the account. For a &#x60;paypal&#x60; payment method this is the user&#39;s email address. For a card it is the last 4 digits of the card. | [optional] 
**LastReplacedAt** | Pointer to **NullableTime** | The date and time when this card was last replaced.  When the Account Updater determines that new card details are available, existing details are not changed immediately. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * &#x60;canceled_payment_method&#x60;     * &#x60;expired_payment_method&#x60;     * &#x60;unavailable_payment_method&#x60;     * &#x60;unknown_payment_method&#x60;  When the replacement is applied, this field is updated. For non-card payment methods, the value of this field is always set to &#x60;null&#x60;. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Method** | Pointer to **string** | The type of this payment method. | [optional] 
**Mode** | Pointer to **string** | The mode to use with this payment method. | [optional] 
**Scheme** | Pointer to **NullableString** | The scheme of the card. Only applies to card payments. | [optional] 
**Status** | Pointer to **string** | The state of the payment method.  - &#x60;processing&#x60; - The payment method is stored but is not ready to be    used yet, as we may be waiting for a notification from a connector    to complete the setup. - &#x60;buyer_approval_required&#x60; - Storing the payment method requires   the buyer to provide approval. Follow the &#x60;approval_url&#x60; for next steps. - &#x60;succeeded&#x60; - The payment method is stored and can be used. - &#x60;failed&#x60; - The payment method could not be stored, or failed verification. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this payment method was last updated in our system. | [optional] 
**Fingerprint** | Pointer to **NullableString** | The unique hash derived from the payment method identifier (e.g. card number). | [optional] 
**LastUsedAt** | Pointer to **NullableTime** | The timestamp when this payment method was last used in a transaction. | [optional] 
**UsageCount** | Pointer to **int32** | The number of times this payment method has been used in transactions. | [optional] 
**CitLastUsedAt** | Pointer to **NullableTime** | The timestamp when this payment method was last used in a transaction for client initiated transactions. | [optional] 
**CitUsageCount** | Pointer to **int32** | The number of times this payment method has been used in transactions for client initiated transactions. | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdditionalSchemes

`func (o *PaymentMethod) GetAdditionalSchemes() []string`

GetAdditionalSchemes returns the AdditionalSchemes field if non-nil, zero value otherwise.

### GetAdditionalSchemesOk

`func (o *PaymentMethod) GetAdditionalSchemesOk() (*[]string, bool)`

GetAdditionalSchemesOk returns a tuple with the AdditionalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSchemes

`func (o *PaymentMethod) SetAdditionalSchemes(v []string)`

SetAdditionalSchemes sets AdditionalSchemes field to given value.

### HasAdditionalSchemes

`func (o *PaymentMethod) HasAdditionalSchemes() bool`

HasAdditionalSchemes returns a boolean if a field has been set.

### SetAdditionalSchemesNil

`func (o *PaymentMethod) SetAdditionalSchemesNil(b bool)`

 SetAdditionalSchemesNil sets the value for AdditionalSchemes to be an explicit nil

### UnsetAdditionalSchemes
`func (o *PaymentMethod) UnsetAdditionalSchemes()`

UnsetAdditionalSchemes ensures that no value is present for AdditionalSchemes, not even an explicit nil
### GetApprovalTarget

`func (o *PaymentMethod) GetApprovalTarget() string`

GetApprovalTarget returns the ApprovalTarget field if non-nil, zero value otherwise.

### GetApprovalTargetOk

`func (o *PaymentMethod) GetApprovalTargetOk() (*string, bool)`

GetApprovalTargetOk returns a tuple with the ApprovalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalTarget

`func (o *PaymentMethod) SetApprovalTarget(v string)`

SetApprovalTarget sets ApprovalTarget field to given value.

### HasApprovalTarget

`func (o *PaymentMethod) HasApprovalTarget() bool`

HasApprovalTarget returns a boolean if a field has been set.

### SetApprovalTargetNil

`func (o *PaymentMethod) SetApprovalTargetNil(b bool)`

 SetApprovalTargetNil sets the value for ApprovalTarget to be an explicit nil

### UnsetApprovalTarget
`func (o *PaymentMethod) UnsetApprovalTarget()`

UnsetApprovalTarget ensures that no value is present for ApprovalTarget, not even an explicit nil
### GetApprovalUrl

`func (o *PaymentMethod) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *PaymentMethod) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *PaymentMethod) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *PaymentMethod) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *PaymentMethod) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *PaymentMethod) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetBuyer

`func (o *PaymentMethod) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *PaymentMethod) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *PaymentMethod) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *PaymentMethod) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *PaymentMethod) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *PaymentMethod) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetCountry

`func (o *PaymentMethod) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethod) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethod) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentMethod) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentMethod) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCreatedAt

`func (o *PaymentMethod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentMethod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentMethod) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentMethod) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentMethod) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentMethod) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentMethod) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentMethod) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDetails

`func (o *PaymentMethod) GetDetails() PaymentMethodDetailsCard`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PaymentMethod) GetDetailsOk() (*PaymentMethodDetailsCard, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PaymentMethod) SetDetails(v PaymentMethodDetailsCard)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PaymentMethod) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetExpirationDate

`func (o *PaymentMethod) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethod) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethod) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethod) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PaymentMethod) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PaymentMethod) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetExternalIdentifier

`func (o *PaymentMethod) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PaymentMethod) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PaymentMethod) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PaymentMethod) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PaymentMethod) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PaymentMethod) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetHasReplacement

`func (o *PaymentMethod) GetHasReplacement() bool`

GetHasReplacement returns the HasReplacement field if non-nil, zero value otherwise.

### GetHasReplacementOk

`func (o *PaymentMethod) GetHasReplacementOk() (*bool, bool)`

GetHasReplacementOk returns a tuple with the HasReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReplacement

`func (o *PaymentMethod) SetHasReplacement(v bool)`

SetHasReplacement sets HasReplacement field to given value.

### HasHasReplacement

`func (o *PaymentMethod) HasHasReplacement() bool`

HasHasReplacement returns a boolean if a field has been set.

### GetLabel

`func (o *PaymentMethod) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentMethod) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentMethod) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentMethod) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PaymentMethod) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PaymentMethod) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetLastReplacedAt

`func (o *PaymentMethod) GetLastReplacedAt() time.Time`

GetLastReplacedAt returns the LastReplacedAt field if non-nil, zero value otherwise.

### GetLastReplacedAtOk

`func (o *PaymentMethod) GetLastReplacedAtOk() (*time.Time, bool)`

GetLastReplacedAtOk returns a tuple with the LastReplacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplacedAt

`func (o *PaymentMethod) SetLastReplacedAt(v time.Time)`

SetLastReplacedAt sets LastReplacedAt field to given value.

### HasLastReplacedAt

`func (o *PaymentMethod) HasLastReplacedAt() bool`

HasLastReplacedAt returns a boolean if a field has been set.

### SetLastReplacedAtNil

`func (o *PaymentMethod) SetLastReplacedAtNil(b bool)`

 SetLastReplacedAtNil sets the value for LastReplacedAt to be an explicit nil

### UnsetLastReplacedAt
`func (o *PaymentMethod) UnsetLastReplacedAt()`

UnsetLastReplacedAt ensures that no value is present for LastReplacedAt, not even an explicit nil
### GetMerchantAccountId

`func (o *PaymentMethod) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PaymentMethod) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PaymentMethod) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *PaymentMethod) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentMethod) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentMethod) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentMethod) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentMethod) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMode

`func (o *PaymentMethod) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaymentMethod) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaymentMethod) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PaymentMethod) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetScheme

`func (o *PaymentMethod) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentMethod) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentMethod) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentMethod) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentMethod) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentMethod) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetStatus

`func (o *PaymentMethod) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethod) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethod) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentMethod) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethod) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentMethod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *PaymentMethod) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *PaymentMethod) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *PaymentMethod) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *PaymentMethod) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *PaymentMethod) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *PaymentMethod) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetLastUsedAt

`func (o *PaymentMethod) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *PaymentMethod) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *PaymentMethod) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *PaymentMethod) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *PaymentMethod) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *PaymentMethod) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetUsageCount

`func (o *PaymentMethod) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *PaymentMethod) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *PaymentMethod) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *PaymentMethod) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetCitLastUsedAt

`func (o *PaymentMethod) GetCitLastUsedAt() time.Time`

GetCitLastUsedAt returns the CitLastUsedAt field if non-nil, zero value otherwise.

### GetCitLastUsedAtOk

`func (o *PaymentMethod) GetCitLastUsedAtOk() (*time.Time, bool)`

GetCitLastUsedAtOk returns a tuple with the CitLastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitLastUsedAt

`func (o *PaymentMethod) SetCitLastUsedAt(v time.Time)`

SetCitLastUsedAt sets CitLastUsedAt field to given value.

### HasCitLastUsedAt

`func (o *PaymentMethod) HasCitLastUsedAt() bool`

HasCitLastUsedAt returns a boolean if a field has been set.

### SetCitLastUsedAtNil

`func (o *PaymentMethod) SetCitLastUsedAtNil(b bool)`

 SetCitLastUsedAtNil sets the value for CitLastUsedAt to be an explicit nil

### UnsetCitLastUsedAt
`func (o *PaymentMethod) UnsetCitLastUsedAt()`

UnsetCitLastUsedAt ensures that no value is present for CitLastUsedAt, not even an explicit nil
### GetCitUsageCount

`func (o *PaymentMethod) GetCitUsageCount() int32`

GetCitUsageCount returns the CitUsageCount field if non-nil, zero value otherwise.

### GetCitUsageCountOk

`func (o *PaymentMethod) GetCitUsageCountOk() (*int32, bool)`

GetCitUsageCountOk returns a tuple with the CitUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitUsageCount

`func (o *PaymentMethod) SetCitUsageCount(v int32)`

SetCitUsageCount sets CitUsageCount field to given value.

### HasCitUsageCount

`func (o *PaymentMethod) HasCitUsageCount() bool`

HasCitUsageCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


