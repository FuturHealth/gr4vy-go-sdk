# GiftCardSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **NullableString** | The ID of this gift card. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Bin** | Pointer to **string** | The first 6 digits of the full gift card number. | [optional] 
**SubBin** | Pointer to **string** | The 3 digits after the &#x60;bin&#x60; of the full gift card number. | [optional] 
**Last4** | Pointer to **string** | The last 4 digits for the gift card. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | The date and time when this gift card expires. This is a full date/time and may be more accurate than the actual expiry date received by the gift card service. | [optional] 
**Balance** | Pointer to **NullableInt32** | The amount remaining on the balance for this gift card according to the gift card service. This may be &#x60;null&#x60; if the balance could not be fetched. | [optional] 
**Currency** | Pointer to **NullableString** | The ISO-4217 currency code that this gift card has a balance for. | [optional] 
**BalanceErrorCode** | Pointer to **NullableString** | If the last balance update failed, this will contain the internal code for this error. | [optional] 
**BalanceRawErrorCode** | Pointer to **NullableString** | If the last balance update failed, this will contain the the raw error code received from the gift card provider. | [optional] 
**BalanceRawErrorMessage** | Pointer to **NullableString** | If the last balance update failed, this will contain the the raw error message received from the gift card provider. | [optional] 

## Methods

### NewGiftCardSummary

`func NewGiftCardSummary() *GiftCardSummary`

NewGiftCardSummary instantiates a new GiftCardSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardSummaryWithDefaults

`func NewGiftCardSummaryWithDefaults() *GiftCardSummary`

NewGiftCardSummaryWithDefaults instantiates a new GiftCardSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCardSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GiftCardSummary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GiftCardSummary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMerchantAccountId

`func (o *GiftCardSummary) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *GiftCardSummary) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *GiftCardSummary) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *GiftCardSummary) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetBin

`func (o *GiftCardSummary) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *GiftCardSummary) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *GiftCardSummary) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *GiftCardSummary) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetSubBin

`func (o *GiftCardSummary) GetSubBin() string`

GetSubBin returns the SubBin field if non-nil, zero value otherwise.

### GetSubBinOk

`func (o *GiftCardSummary) GetSubBinOk() (*string, bool)`

GetSubBinOk returns a tuple with the SubBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBin

`func (o *GiftCardSummary) SetSubBin(v string)`

SetSubBin sets SubBin field to given value.

### HasSubBin

`func (o *GiftCardSummary) HasSubBin() bool`

HasSubBin returns a boolean if a field has been set.

### GetLast4

`func (o *GiftCardSummary) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *GiftCardSummary) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *GiftCardSummary) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *GiftCardSummary) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GiftCardSummary) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GiftCardSummary) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GiftCardSummary) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GiftCardSummary) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *GiftCardSummary) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *GiftCardSummary) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetBalance

`func (o *GiftCardSummary) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GiftCardSummary) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GiftCardSummary) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GiftCardSummary) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *GiftCardSummary) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *GiftCardSummary) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetCurrency

`func (o *GiftCardSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GiftCardSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GiftCardSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GiftCardSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *GiftCardSummary) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *GiftCardSummary) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetBalanceErrorCode

`func (o *GiftCardSummary) GetBalanceErrorCode() string`

GetBalanceErrorCode returns the BalanceErrorCode field if non-nil, zero value otherwise.

### GetBalanceErrorCodeOk

`func (o *GiftCardSummary) GetBalanceErrorCodeOk() (*string, bool)`

GetBalanceErrorCodeOk returns a tuple with the BalanceErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceErrorCode

`func (o *GiftCardSummary) SetBalanceErrorCode(v string)`

SetBalanceErrorCode sets BalanceErrorCode field to given value.

### HasBalanceErrorCode

`func (o *GiftCardSummary) HasBalanceErrorCode() bool`

HasBalanceErrorCode returns a boolean if a field has been set.

### SetBalanceErrorCodeNil

`func (o *GiftCardSummary) SetBalanceErrorCodeNil(b bool)`

 SetBalanceErrorCodeNil sets the value for BalanceErrorCode to be an explicit nil

### UnsetBalanceErrorCode
`func (o *GiftCardSummary) UnsetBalanceErrorCode()`

UnsetBalanceErrorCode ensures that no value is present for BalanceErrorCode, not even an explicit nil
### GetBalanceRawErrorCode

`func (o *GiftCardSummary) GetBalanceRawErrorCode() string`

GetBalanceRawErrorCode returns the BalanceRawErrorCode field if non-nil, zero value otherwise.

### GetBalanceRawErrorCodeOk

`func (o *GiftCardSummary) GetBalanceRawErrorCodeOk() (*string, bool)`

GetBalanceRawErrorCodeOk returns a tuple with the BalanceRawErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceRawErrorCode

`func (o *GiftCardSummary) SetBalanceRawErrorCode(v string)`

SetBalanceRawErrorCode sets BalanceRawErrorCode field to given value.

### HasBalanceRawErrorCode

`func (o *GiftCardSummary) HasBalanceRawErrorCode() bool`

HasBalanceRawErrorCode returns a boolean if a field has been set.

### SetBalanceRawErrorCodeNil

`func (o *GiftCardSummary) SetBalanceRawErrorCodeNil(b bool)`

 SetBalanceRawErrorCodeNil sets the value for BalanceRawErrorCode to be an explicit nil

### UnsetBalanceRawErrorCode
`func (o *GiftCardSummary) UnsetBalanceRawErrorCode()`

UnsetBalanceRawErrorCode ensures that no value is present for BalanceRawErrorCode, not even an explicit nil
### GetBalanceRawErrorMessage

`func (o *GiftCardSummary) GetBalanceRawErrorMessage() string`

GetBalanceRawErrorMessage returns the BalanceRawErrorMessage field if non-nil, zero value otherwise.

### GetBalanceRawErrorMessageOk

`func (o *GiftCardSummary) GetBalanceRawErrorMessageOk() (*string, bool)`

GetBalanceRawErrorMessageOk returns a tuple with the BalanceRawErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceRawErrorMessage

`func (o *GiftCardSummary) SetBalanceRawErrorMessage(v string)`

SetBalanceRawErrorMessage sets BalanceRawErrorMessage field to given value.

### HasBalanceRawErrorMessage

`func (o *GiftCardSummary) HasBalanceRawErrorMessage() bool`

HasBalanceRawErrorMessage returns a boolean if a field has been set.

### SetBalanceRawErrorMessageNil

`func (o *GiftCardSummary) SetBalanceRawErrorMessageNil(b bool)`

 SetBalanceRawErrorMessageNil sets the value for BalanceRawErrorMessage to be an explicit nil

### UnsetBalanceRawErrorMessage
`func (o *GiftCardSummary) UnsetBalanceRawErrorMessage()`

UnsetBalanceRawErrorMessage ensures that no value is present for BalanceRawErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


