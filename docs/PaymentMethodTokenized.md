# PaymentMethodTokenized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Method** | Pointer to **string** | The type of this payment method. | [optional] 
**Label** | Pointer to **string** | A label for the payment method. For a &#x60;card&#x60; payment method this is the last 4 digits on the card. For others it would be the email address. | [optional] 
**Scheme** | Pointer to **NullableString** | The type of the card, if the payment method is a card. | [optional] 
**AdditionalSchemes** | Pointer to **[]string** | Additional schemes of the card. Only applies to card payment methods. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date for the payment method. | [optional] 
**ApprovalTarget** | Pointer to **NullableString** | The browser target that an approval URL must be opened in. If &#x60;any&#x60; or &#x60;null&#x60;, then there is no specific requirement. | [optional] 
**ApprovalUrl** | Pointer to **NullableString** | The optional URL that the buyer needs to be redirected to to further authorize their payment. | [optional] 
**Currency** | Pointer to **NullableString** | The ISO-4217 currency code that this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used for multiple currencies. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used in multiple countries. | [optional] 
**LastReplacedAt** | Pointer to **NullableTime** | The date and time when this card was last replaced.  When the Account Updater determines that new card details are available, existing details are not changed immediately. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * &#x60;canceled_payment_method&#x60;     * &#x60;expired_payment_method&#x60;     * &#x60;unavailable_payment_method&#x60;     * &#x60;unknown_payment_method&#x60;  When the replacement is applied, this field is updated. For non-card payment methods, the value of this field is always set to &#x60;null&#x60;. | [optional] 
**HasReplacement** | Pointer to **bool** | Whether this card has a pending replacement that hasn&#39;t been applied yet.  When the Account Updater determines that new card details are available, existing details are not changed immediately, but this field is set to &#x60;true&#x60;. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * &#x60;canceled_payment_method&#x60;     * &#x60;expired_payment_method&#x60;     * &#x60;unavailable_payment_method&#x60;     * &#x60;unknown_payment_method&#x60;  When the replacement is applied, this field is set to &#x60;false&#x60;. For non-card payment methods, the value of this field is always set to &#x60;false&#x60;. | [optional] 
**LastUsedAt** | Pointer to **NullableTime** | The timestamp when this payment method was last used in a transaction. | [optional] 
**UsageCount** | Pointer to **int32** | The number of times this payment method has been used in transactions. | [optional] 
**CitLastUsedAt** | Pointer to **NullableTime** | The timestamp when this payment method was last used in a transaction for client initiated transactions. | [optional] 
**CitUsageCount** | Pointer to **int32** | The number of times this payment method has been used in transactions for client initiated transactions. | [optional] 

## Methods

### NewPaymentMethodTokenized

`func NewPaymentMethodTokenized() *PaymentMethodTokenized`

NewPaymentMethodTokenized instantiates a new PaymentMethodTokenized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodTokenizedWithDefaults

`func NewPaymentMethodTokenizedWithDefaults() *PaymentMethodTokenized`

NewPaymentMethodTokenizedWithDefaults instantiates a new PaymentMethodTokenized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodTokenized) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodTokenized) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodTokenized) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodTokenized) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodTokenized) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodTokenized) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodTokenized) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodTokenized) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *PaymentMethodTokenized) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PaymentMethodTokenized) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PaymentMethodTokenized) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *PaymentMethodTokenized) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentMethodTokenized) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentMethodTokenized) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentMethodTokenized) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentMethodTokenized) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetLabel

`func (o *PaymentMethodTokenized) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentMethodTokenized) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentMethodTokenized) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentMethodTokenized) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScheme

`func (o *PaymentMethodTokenized) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentMethodTokenized) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentMethodTokenized) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentMethodTokenized) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentMethodTokenized) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentMethodTokenized) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetAdditionalSchemes

`func (o *PaymentMethodTokenized) GetAdditionalSchemes() []string`

GetAdditionalSchemes returns the AdditionalSchemes field if non-nil, zero value otherwise.

### GetAdditionalSchemesOk

`func (o *PaymentMethodTokenized) GetAdditionalSchemesOk() (*[]string, bool)`

GetAdditionalSchemesOk returns a tuple with the AdditionalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSchemes

`func (o *PaymentMethodTokenized) SetAdditionalSchemes(v []string)`

SetAdditionalSchemes sets AdditionalSchemes field to given value.

### HasAdditionalSchemes

`func (o *PaymentMethodTokenized) HasAdditionalSchemes() bool`

HasAdditionalSchemes returns a boolean if a field has been set.

### SetAdditionalSchemesNil

`func (o *PaymentMethodTokenized) SetAdditionalSchemesNil(b bool)`

 SetAdditionalSchemesNil sets the value for AdditionalSchemes to be an explicit nil

### UnsetAdditionalSchemes
`func (o *PaymentMethodTokenized) UnsetAdditionalSchemes()`

UnsetAdditionalSchemes ensures that no value is present for AdditionalSchemes, not even an explicit nil
### GetExpirationDate

`func (o *PaymentMethodTokenized) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethodTokenized) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethodTokenized) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethodTokenized) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PaymentMethodTokenized) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PaymentMethodTokenized) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetApprovalTarget

`func (o *PaymentMethodTokenized) GetApprovalTarget() string`

GetApprovalTarget returns the ApprovalTarget field if non-nil, zero value otherwise.

### GetApprovalTargetOk

`func (o *PaymentMethodTokenized) GetApprovalTargetOk() (*string, bool)`

GetApprovalTargetOk returns a tuple with the ApprovalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalTarget

`func (o *PaymentMethodTokenized) SetApprovalTarget(v string)`

SetApprovalTarget sets ApprovalTarget field to given value.

### HasApprovalTarget

`func (o *PaymentMethodTokenized) HasApprovalTarget() bool`

HasApprovalTarget returns a boolean if a field has been set.

### SetApprovalTargetNil

`func (o *PaymentMethodTokenized) SetApprovalTargetNil(b bool)`

 SetApprovalTargetNil sets the value for ApprovalTarget to be an explicit nil

### UnsetApprovalTarget
`func (o *PaymentMethodTokenized) UnsetApprovalTarget()`

UnsetApprovalTarget ensures that no value is present for ApprovalTarget, not even an explicit nil
### GetApprovalUrl

`func (o *PaymentMethodTokenized) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *PaymentMethodTokenized) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *PaymentMethodTokenized) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *PaymentMethodTokenized) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *PaymentMethodTokenized) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *PaymentMethodTokenized) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetCurrency

`func (o *PaymentMethodTokenized) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentMethodTokenized) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentMethodTokenized) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentMethodTokenized) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentMethodTokenized) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentMethodTokenized) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCountry

`func (o *PaymentMethodTokenized) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodTokenized) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodTokenized) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodTokenized) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentMethodTokenized) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentMethodTokenized) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLastReplacedAt

`func (o *PaymentMethodTokenized) GetLastReplacedAt() time.Time`

GetLastReplacedAt returns the LastReplacedAt field if non-nil, zero value otherwise.

### GetLastReplacedAtOk

`func (o *PaymentMethodTokenized) GetLastReplacedAtOk() (*time.Time, bool)`

GetLastReplacedAtOk returns a tuple with the LastReplacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplacedAt

`func (o *PaymentMethodTokenized) SetLastReplacedAt(v time.Time)`

SetLastReplacedAt sets LastReplacedAt field to given value.

### HasLastReplacedAt

`func (o *PaymentMethodTokenized) HasLastReplacedAt() bool`

HasLastReplacedAt returns a boolean if a field has been set.

### SetLastReplacedAtNil

`func (o *PaymentMethodTokenized) SetLastReplacedAtNil(b bool)`

 SetLastReplacedAtNil sets the value for LastReplacedAt to be an explicit nil

### UnsetLastReplacedAt
`func (o *PaymentMethodTokenized) UnsetLastReplacedAt()`

UnsetLastReplacedAt ensures that no value is present for LastReplacedAt, not even an explicit nil
### GetHasReplacement

`func (o *PaymentMethodTokenized) GetHasReplacement() bool`

GetHasReplacement returns the HasReplacement field if non-nil, zero value otherwise.

### GetHasReplacementOk

`func (o *PaymentMethodTokenized) GetHasReplacementOk() (*bool, bool)`

GetHasReplacementOk returns a tuple with the HasReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReplacement

`func (o *PaymentMethodTokenized) SetHasReplacement(v bool)`

SetHasReplacement sets HasReplacement field to given value.

### HasHasReplacement

`func (o *PaymentMethodTokenized) HasHasReplacement() bool`

HasHasReplacement returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *PaymentMethodTokenized) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *PaymentMethodTokenized) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *PaymentMethodTokenized) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *PaymentMethodTokenized) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *PaymentMethodTokenized) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *PaymentMethodTokenized) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetUsageCount

`func (o *PaymentMethodTokenized) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *PaymentMethodTokenized) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *PaymentMethodTokenized) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *PaymentMethodTokenized) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetCitLastUsedAt

`func (o *PaymentMethodTokenized) GetCitLastUsedAt() time.Time`

GetCitLastUsedAt returns the CitLastUsedAt field if non-nil, zero value otherwise.

### GetCitLastUsedAtOk

`func (o *PaymentMethodTokenized) GetCitLastUsedAtOk() (*time.Time, bool)`

GetCitLastUsedAtOk returns a tuple with the CitLastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitLastUsedAt

`func (o *PaymentMethodTokenized) SetCitLastUsedAt(v time.Time)`

SetCitLastUsedAt sets CitLastUsedAt field to given value.

### HasCitLastUsedAt

`func (o *PaymentMethodTokenized) HasCitLastUsedAt() bool`

HasCitLastUsedAt returns a boolean if a field has been set.

### SetCitLastUsedAtNil

`func (o *PaymentMethodTokenized) SetCitLastUsedAtNil(b bool)`

 SetCitLastUsedAtNil sets the value for CitLastUsedAt to be an explicit nil

### UnsetCitLastUsedAt
`func (o *PaymentMethodTokenized) UnsetCitLastUsedAt()`

UnsetCitLastUsedAt ensures that no value is present for CitLastUsedAt, not even an explicit nil
### GetCitUsageCount

`func (o *PaymentMethodTokenized) GetCitUsageCount() int32`

GetCitUsageCount returns the CitUsageCount field if non-nil, zero value otherwise.

### GetCitUsageCountOk

`func (o *PaymentMethodTokenized) GetCitUsageCountOk() (*int32, bool)`

GetCitUsageCountOk returns a tuple with the CitUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitUsageCount

`func (o *PaymentMethodTokenized) SetCitUsageCount(v int32)`

SetCitUsageCount sets CitUsageCount field to given value.

### HasCitUsageCount

`func (o *PaymentMethodTokenized) HasCitUsageCount() bool`

HasCitUsageCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


