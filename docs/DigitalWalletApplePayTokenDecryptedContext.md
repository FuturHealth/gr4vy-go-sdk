# DigitalWalletApplePayTokenDecryptedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version information about the payment token. | [optional] 
**Type** | Pointer to **NullableString** | The type of payment instrument. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration of the decrypted data. | [optional] 
**HasCryptogram** | Pointer to **bool** | Online payment cryptogram, as defined by 3D Secure. | [optional] 
**Eci** | Pointer to **NullableString** | ECI indicator, as defined by 3D Secure. | [optional] 
**ApplicationData** | Pointer to **NullableString** | Hash of the application data property of the original request. | [optional] 
**TransactionIdentifier** | Pointer to **string** | The unique identifier from Apple Pay. | [optional] 
**CardholderName** | Pointer to **NullableString** | The cardholder name. | [optional] 
**CurrencyCode** | Pointer to **string** | ISO 4217 numeric currency code for the transaction. | [optional] 
**TransactionAmount** | Pointer to **int32** | The amount for the transaction. | [optional] 
**DeviceManufacturerIdentifier** | Pointer to **string** | Hex-encoded device manufacturer identifier which initiated the transaction. | [optional] 
**PaymentDataType** | Pointer to **string** | Either \&quot;3DSecure\&quot; or \&quot;EMV\&quot;. | [optional] 
**MerchantTokenIdentifier** | Pointer to **NullableString** | For a merchant token request, the provisioned merchant token identifier from the payment network. | [optional] 
**CardExpirationDate** | Pointer to **NullableString** | Expiration date of card. | [optional] 
**CardSuffix** | Pointer to **NullableString** | Last four digits of card PAN. | [optional] 

## Methods

### NewDigitalWalletApplePayTokenDecryptedContext

`func NewDigitalWalletApplePayTokenDecryptedContext() *DigitalWalletApplePayTokenDecryptedContext`

NewDigitalWalletApplePayTokenDecryptedContext instantiates a new DigitalWalletApplePayTokenDecryptedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletApplePayTokenDecryptedContextWithDefaults

`func NewDigitalWalletApplePayTokenDecryptedContextWithDefaults() *DigitalWalletApplePayTokenDecryptedContext`

NewDigitalWalletApplePayTokenDecryptedContextWithDefaults instantiates a new DigitalWalletApplePayTokenDecryptedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExpirationDate

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHasCryptogram

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetHasCryptogram() bool`

GetHasCryptogram returns the HasCryptogram field if non-nil, zero value otherwise.

### GetHasCryptogramOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetHasCryptogramOk() (*bool, bool)`

GetHasCryptogramOk returns a tuple with the HasCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCryptogram

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetHasCryptogram(v bool)`

SetHasCryptogram sets HasCryptogram field to given value.

### HasHasCryptogram

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasHasCryptogram() bool`

HasHasCryptogram returns a boolean if a field has been set.

### GetEci

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasEci() bool`

HasEci returns a boolean if a field has been set.

### SetEciNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetEciNil(b bool)`

 SetEciNil sets the value for Eci to be an explicit nil

### UnsetEci
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetEci()`

UnsetEci ensures that no value is present for Eci, not even an explicit nil
### GetApplicationData

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetApplicationData() string`

GetApplicationData returns the ApplicationData field if non-nil, zero value otherwise.

### GetApplicationDataOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetApplicationDataOk() (*string, bool)`

GetApplicationDataOk returns a tuple with the ApplicationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationData

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetApplicationData(v string)`

SetApplicationData sets ApplicationData field to given value.

### HasApplicationData

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasApplicationData() bool`

HasApplicationData returns a boolean if a field has been set.

### SetApplicationDataNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetApplicationDataNil(b bool)`

 SetApplicationDataNil sets the value for ApplicationData to be an explicit nil

### UnsetApplicationData
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetApplicationData()`

UnsetApplicationData ensures that no value is present for ApplicationData, not even an explicit nil
### GetTransactionIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetTransactionIdentifier() string`

GetTransactionIdentifier returns the TransactionIdentifier field if non-nil, zero value otherwise.

### GetTransactionIdentifierOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetTransactionIdentifierOk() (*string, bool)`

GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetTransactionIdentifier(v string)`

SetTransactionIdentifier sets TransactionIdentifier field to given value.

### HasTransactionIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasTransactionIdentifier() bool`

HasTransactionIdentifier returns a boolean if a field has been set.

### GetCardholderName

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
### GetCurrencyCode

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetTransactionAmount() int32`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetTransactionAmountOk() (*int32, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetTransactionAmount(v int32)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetDeviceManufacturerIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetDeviceManufacturerIdentifier() string`

GetDeviceManufacturerIdentifier returns the DeviceManufacturerIdentifier field if non-nil, zero value otherwise.

### GetDeviceManufacturerIdentifierOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetDeviceManufacturerIdentifierOk() (*string, bool)`

GetDeviceManufacturerIdentifierOk returns a tuple with the DeviceManufacturerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManufacturerIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetDeviceManufacturerIdentifier(v string)`

SetDeviceManufacturerIdentifier sets DeviceManufacturerIdentifier field to given value.

### HasDeviceManufacturerIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasDeviceManufacturerIdentifier() bool`

HasDeviceManufacturerIdentifier returns a boolean if a field has been set.

### GetPaymentDataType

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetPaymentDataType() string`

GetPaymentDataType returns the PaymentDataType field if non-nil, zero value otherwise.

### GetPaymentDataTypeOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetPaymentDataTypeOk() (*string, bool)`

GetPaymentDataTypeOk returns a tuple with the PaymentDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDataType

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetPaymentDataType(v string)`

SetPaymentDataType sets PaymentDataType field to given value.

### HasPaymentDataType

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasPaymentDataType() bool`

HasPaymentDataType returns a boolean if a field has been set.

### GetMerchantTokenIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetMerchantTokenIdentifier() string`

GetMerchantTokenIdentifier returns the MerchantTokenIdentifier field if non-nil, zero value otherwise.

### GetMerchantTokenIdentifierOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetMerchantTokenIdentifierOk() (*string, bool)`

GetMerchantTokenIdentifierOk returns a tuple with the MerchantTokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTokenIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetMerchantTokenIdentifier(v string)`

SetMerchantTokenIdentifier sets MerchantTokenIdentifier field to given value.

### HasMerchantTokenIdentifier

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasMerchantTokenIdentifier() bool`

HasMerchantTokenIdentifier returns a boolean if a field has been set.

### SetMerchantTokenIdentifierNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetMerchantTokenIdentifierNil(b bool)`

 SetMerchantTokenIdentifierNil sets the value for MerchantTokenIdentifier to be an explicit nil

### UnsetMerchantTokenIdentifier
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetMerchantTokenIdentifier()`

UnsetMerchantTokenIdentifier ensures that no value is present for MerchantTokenIdentifier, not even an explicit nil
### GetCardExpirationDate

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCardExpirationDate() string`

GetCardExpirationDate returns the CardExpirationDate field if non-nil, zero value otherwise.

### GetCardExpirationDateOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCardExpirationDateOk() (*string, bool)`

GetCardExpirationDateOk returns a tuple with the CardExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationDate

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCardExpirationDate(v string)`

SetCardExpirationDate sets CardExpirationDate field to given value.

### HasCardExpirationDate

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasCardExpirationDate() bool`

HasCardExpirationDate returns a boolean if a field has been set.

### SetCardExpirationDateNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCardExpirationDateNil(b bool)`

 SetCardExpirationDateNil sets the value for CardExpirationDate to be an explicit nil

### UnsetCardExpirationDate
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetCardExpirationDate()`

UnsetCardExpirationDate ensures that no value is present for CardExpirationDate, not even an explicit nil
### GetCardSuffix

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCardSuffix() string`

GetCardSuffix returns the CardSuffix field if non-nil, zero value otherwise.

### GetCardSuffixOk

`func (o *DigitalWalletApplePayTokenDecryptedContext) GetCardSuffixOk() (*string, bool)`

GetCardSuffixOk returns a tuple with the CardSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSuffix

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCardSuffix(v string)`

SetCardSuffix sets CardSuffix field to given value.

### HasCardSuffix

`func (o *DigitalWalletApplePayTokenDecryptedContext) HasCardSuffix() bool`

HasCardSuffix returns a boolean if a field has been set.

### SetCardSuffixNil

`func (o *DigitalWalletApplePayTokenDecryptedContext) SetCardSuffixNil(b bool)`

 SetCardSuffixNil sets the value for CardSuffix to be an explicit nil

### UnsetCardSuffix
`func (o *DigitalWalletApplePayTokenDecryptedContext) UnsetCardSuffix()`

UnsetCardSuffix ensures that no value is present for CardSuffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


