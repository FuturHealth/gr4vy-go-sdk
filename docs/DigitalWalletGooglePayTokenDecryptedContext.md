# DigitalWalletGooglePayTokenDecryptedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version information about the payment token. | [optional] 
**Type** | Pointer to **NullableString** | The type of payment instrument. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration of the decrypted data. | [optional] 
**HasCryptogram** | Pointer to **bool** | Online payment cryptogram, as defined by 3D Secure. | [optional] 
**Eci** | Pointer to **NullableString** | ECI indicator, as defined by 3D Secure. | [optional] 
**MessageExpiration** | Pointer to **string** | Date and time at which the message expires as UTC milliseconds since epoch. | [optional] 
**MessageId** | Pointer to **string** | A unique ID that identifies the message in case it needs to be revoked or located at a later time. | [optional] 
**PaymentMethod** | Pointer to **string** | The type of the payment credential. | [optional] 

## Methods

### NewDigitalWalletGooglePayTokenDecryptedContext

`func NewDigitalWalletGooglePayTokenDecryptedContext() *DigitalWalletGooglePayTokenDecryptedContext`

NewDigitalWalletGooglePayTokenDecryptedContext instantiates a new DigitalWalletGooglePayTokenDecryptedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletGooglePayTokenDecryptedContextWithDefaults

`func NewDigitalWalletGooglePayTokenDecryptedContextWithDefaults() *DigitalWalletGooglePayTokenDecryptedContext`

NewDigitalWalletGooglePayTokenDecryptedContextWithDefaults instantiates a new DigitalWalletGooglePayTokenDecryptedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DigitalWalletGooglePayTokenDecryptedContext) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExpirationDate

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHasCryptogram

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetHasCryptogram() bool`

GetHasCryptogram returns the HasCryptogram field if non-nil, zero value otherwise.

### GetHasCryptogramOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetHasCryptogramOk() (*bool, bool)`

GetHasCryptogramOk returns a tuple with the HasCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCryptogram

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetHasCryptogram(v bool)`

SetHasCryptogram sets HasCryptogram field to given value.

### HasHasCryptogram

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasHasCryptogram() bool`

HasHasCryptogram returns a boolean if a field has been set.

### GetEci

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasEci() bool`

HasEci returns a boolean if a field has been set.

### SetEciNil

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetEciNil(b bool)`

 SetEciNil sets the value for Eci to be an explicit nil

### UnsetEci
`func (o *DigitalWalletGooglePayTokenDecryptedContext) UnsetEci()`

UnsetEci ensures that no value is present for Eci, not even an explicit nil
### GetMessageExpiration

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetMessageExpiration() string`

GetMessageExpiration returns the MessageExpiration field if non-nil, zero value otherwise.

### GetMessageExpirationOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetMessageExpirationOk() (*string, bool)`

GetMessageExpirationOk returns a tuple with the MessageExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageExpiration

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetMessageExpiration(v string)`

SetMessageExpiration sets MessageExpiration field to given value.

### HasMessageExpiration

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasMessageExpiration() bool`

HasMessageExpiration returns a boolean if a field has been set.

### GetMessageId

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *DigitalWalletGooglePayTokenDecryptedContext) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *DigitalWalletGooglePayTokenDecryptedContext) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *DigitalWalletGooglePayTokenDecryptedContext) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


