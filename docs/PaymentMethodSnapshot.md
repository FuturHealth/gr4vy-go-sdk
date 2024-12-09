# PaymentMethodSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method&#x60;. | [optional] 
**Id** | Pointer to **NullableString** | The unique ID of the payment method. | [optional] 
**ApprovalTarget** | Pointer to **NullableString** | The browser target that an approval URL must be opened in. If &#x60;any&#x60; or &#x60;null&#x60;, then there is no specific requirement. | [optional] 
**ApprovalUrl** | Pointer to **NullableString** | The optional URL that the buyer needs to be redirected to to further authorize their payment. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used in multiple countries. | [optional] 
**Currency** | Pointer to **NullableString** | The ISO-4217 currency code that this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used for multiple currencies. | [optional] 
**Details** | Pointer to [**NullablePaymentMethodDetailsCard**](PaymentMethodDetailsCard.md) |  | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date for this payment method. This is mostly used by cards where the card might have an expiration date. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the payment method against your own records. | [optional] 
**Label** | Pointer to **NullableString** | A label for the payment method. This can be the last 4 digits for a card, or the email address for an alternative payment method. | [optional] 
**LastReplacedAt** | Pointer to **NullableTime** | The date and time when this card was last replaced.  When the Account Updater determines that new card details are available, existing details are not changed immediately. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * &#x60;canceled_payment_method&#x60;     * &#x60;expired_payment_method&#x60;     * &#x60;unavailable_payment_method&#x60;     * &#x60;unknown_payment_method&#x60;  When the replacement is applied, this field is updated. For non-card payment methods, the value of this field is always set to &#x60;null&#x60;. | [optional] 
**Method** | Pointer to **string** | The type of this payment method. | [optional] 
**PaymentAccountReference** | Pointer to **NullableString** | The payment account reference (PAR) returned by the card scheme. This is a unique reference to the underlying account that has been used to fund this payment method. This value will be unique if the same underlying account was used, regardless of the actual payment method used. For example, a network token or an Apple Pay device token will return the same PAR when possible.  The uniqueness of this value will depend on the card scheme, please refer to their documentation for further details. The availability of the PAR in our API depends on the availability of its value in the API of the payment service used for the transaction. | [optional] 
**Scheme** | Pointer to **NullableString** | An additional label used to differentiate different sub-types of a payment method. Most notably this can include the type of card used in a transaction. This field is &#x60;null&#x60; for the non-card payment methods. This represents the card scheme sent to the connector and it could be different from the actual card scheme that is being used by the PSP to process the transaction in the following situations: 1. &#x60;use_additional_scheme&#x60; transformation is used with the &#x60;PAN&#x60; instrument but we already have a PSP token for the card. 2. &#x60;use_additional_scheme&#x60; transformation is used but PSP has fallen back to the main card scheme internally. | [optional] 
**Fingerprint** | Pointer to **NullableString** | The unique hash derived from the payment method identifier (e.g. card number). | [optional] 

## Methods

### NewPaymentMethodSnapshot

`func NewPaymentMethodSnapshot() *PaymentMethodSnapshot`

NewPaymentMethodSnapshot instantiates a new PaymentMethodSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodSnapshotWithDefaults

`func NewPaymentMethodSnapshotWithDefaults() *PaymentMethodSnapshot`

NewPaymentMethodSnapshotWithDefaults instantiates a new PaymentMethodSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentMethodSnapshot) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentMethodSnapshot) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetApprovalTarget

`func (o *PaymentMethodSnapshot) GetApprovalTarget() string`

GetApprovalTarget returns the ApprovalTarget field if non-nil, zero value otherwise.

### GetApprovalTargetOk

`func (o *PaymentMethodSnapshot) GetApprovalTargetOk() (*string, bool)`

GetApprovalTargetOk returns a tuple with the ApprovalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalTarget

`func (o *PaymentMethodSnapshot) SetApprovalTarget(v string)`

SetApprovalTarget sets ApprovalTarget field to given value.

### HasApprovalTarget

`func (o *PaymentMethodSnapshot) HasApprovalTarget() bool`

HasApprovalTarget returns a boolean if a field has been set.

### SetApprovalTargetNil

`func (o *PaymentMethodSnapshot) SetApprovalTargetNil(b bool)`

 SetApprovalTargetNil sets the value for ApprovalTarget to be an explicit nil

### UnsetApprovalTarget
`func (o *PaymentMethodSnapshot) UnsetApprovalTarget()`

UnsetApprovalTarget ensures that no value is present for ApprovalTarget, not even an explicit nil
### GetApprovalUrl

`func (o *PaymentMethodSnapshot) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *PaymentMethodSnapshot) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *PaymentMethodSnapshot) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *PaymentMethodSnapshot) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *PaymentMethodSnapshot) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *PaymentMethodSnapshot) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetCountry

`func (o *PaymentMethodSnapshot) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodSnapshot) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodSnapshot) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodSnapshot) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentMethodSnapshot) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentMethodSnapshot) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCurrency

`func (o *PaymentMethodSnapshot) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentMethodSnapshot) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentMethodSnapshot) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentMethodSnapshot) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentMethodSnapshot) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentMethodSnapshot) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDetails

`func (o *PaymentMethodSnapshot) GetDetails() PaymentMethodDetailsCard`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PaymentMethodSnapshot) GetDetailsOk() (*PaymentMethodDetailsCard, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PaymentMethodSnapshot) SetDetails(v PaymentMethodDetailsCard)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PaymentMethodSnapshot) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PaymentMethodSnapshot) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PaymentMethodSnapshot) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetExpirationDate

`func (o *PaymentMethodSnapshot) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethodSnapshot) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethodSnapshot) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethodSnapshot) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PaymentMethodSnapshot) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PaymentMethodSnapshot) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetExternalIdentifier

`func (o *PaymentMethodSnapshot) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PaymentMethodSnapshot) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PaymentMethodSnapshot) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PaymentMethodSnapshot) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PaymentMethodSnapshot) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PaymentMethodSnapshot) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetLabel

`func (o *PaymentMethodSnapshot) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentMethodSnapshot) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentMethodSnapshot) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentMethodSnapshot) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PaymentMethodSnapshot) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PaymentMethodSnapshot) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetLastReplacedAt

`func (o *PaymentMethodSnapshot) GetLastReplacedAt() time.Time`

GetLastReplacedAt returns the LastReplacedAt field if non-nil, zero value otherwise.

### GetLastReplacedAtOk

`func (o *PaymentMethodSnapshot) GetLastReplacedAtOk() (*time.Time, bool)`

GetLastReplacedAtOk returns a tuple with the LastReplacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplacedAt

`func (o *PaymentMethodSnapshot) SetLastReplacedAt(v time.Time)`

SetLastReplacedAt sets LastReplacedAt field to given value.

### HasLastReplacedAt

`func (o *PaymentMethodSnapshot) HasLastReplacedAt() bool`

HasLastReplacedAt returns a boolean if a field has been set.

### SetLastReplacedAtNil

`func (o *PaymentMethodSnapshot) SetLastReplacedAtNil(b bool)`

 SetLastReplacedAtNil sets the value for LastReplacedAt to be an explicit nil

### UnsetLastReplacedAt
`func (o *PaymentMethodSnapshot) UnsetLastReplacedAt()`

UnsetLastReplacedAt ensures that no value is present for LastReplacedAt, not even an explicit nil
### GetMethod

`func (o *PaymentMethodSnapshot) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentMethodSnapshot) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentMethodSnapshot) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentMethodSnapshot) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPaymentAccountReference

`func (o *PaymentMethodSnapshot) GetPaymentAccountReference() string`

GetPaymentAccountReference returns the PaymentAccountReference field if non-nil, zero value otherwise.

### GetPaymentAccountReferenceOk

`func (o *PaymentMethodSnapshot) GetPaymentAccountReferenceOk() (*string, bool)`

GetPaymentAccountReferenceOk returns a tuple with the PaymentAccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountReference

`func (o *PaymentMethodSnapshot) SetPaymentAccountReference(v string)`

SetPaymentAccountReference sets PaymentAccountReference field to given value.

### HasPaymentAccountReference

`func (o *PaymentMethodSnapshot) HasPaymentAccountReference() bool`

HasPaymentAccountReference returns a boolean if a field has been set.

### SetPaymentAccountReferenceNil

`func (o *PaymentMethodSnapshot) SetPaymentAccountReferenceNil(b bool)`

 SetPaymentAccountReferenceNil sets the value for PaymentAccountReference to be an explicit nil

### UnsetPaymentAccountReference
`func (o *PaymentMethodSnapshot) UnsetPaymentAccountReference()`

UnsetPaymentAccountReference ensures that no value is present for PaymentAccountReference, not even an explicit nil
### GetScheme

`func (o *PaymentMethodSnapshot) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentMethodSnapshot) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentMethodSnapshot) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentMethodSnapshot) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentMethodSnapshot) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentMethodSnapshot) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetFingerprint

`func (o *PaymentMethodSnapshot) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *PaymentMethodSnapshot) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *PaymentMethodSnapshot) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *PaymentMethodSnapshot) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *PaymentMethodSnapshot) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *PaymentMethodSnapshot) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


