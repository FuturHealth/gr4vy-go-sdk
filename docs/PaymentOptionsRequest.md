# PaymentOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableInt32** | The monetary amount to create an authorization for, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;.  If the &#x60;intent&#x60; is set to &#x60;capture&#x60;, an amount greater than zero must be supplied. | [optional] 
**Locale** | Pointer to **NullableString** | An ISO 639-1 Language Code and optional ISO 3166 Country Code. This locale determines the language for the labels returned for every payment option. | [optional] [default to "en"]
**Currency** | Pointer to **NullableString** | A supported ISO-4217 currency code.  For redirect requests, this value must match the one specified for &#x60;currency&#x60; in &#x60;payment_method&#x60;.  | [optional] 
**Country** | Pointer to **NullableString** | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Used by the Flow engine to filter available options based on various client-defined parameters. If present, this must be a string representing a valid JSON dictionary. | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 

## Methods

### NewPaymentOptionsRequest

`func NewPaymentOptionsRequest() *PaymentOptionsRequest`

NewPaymentOptionsRequest instantiates a new PaymentOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionsRequestWithDefaults

`func NewPaymentOptionsRequestWithDefaults() *PaymentOptionsRequest`

NewPaymentOptionsRequestWithDefaults instantiates a new PaymentOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentOptionsRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentOptionsRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentOptionsRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentOptionsRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PaymentOptionsRequest) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PaymentOptionsRequest) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetLocale

`func (o *PaymentOptionsRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PaymentOptionsRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PaymentOptionsRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PaymentOptionsRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *PaymentOptionsRequest) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *PaymentOptionsRequest) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetCurrency

`func (o *PaymentOptionsRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentOptionsRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentOptionsRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentOptionsRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentOptionsRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentOptionsRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCountry

`func (o *PaymentOptionsRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentOptionsRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentOptionsRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentOptionsRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentOptionsRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentOptionsRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetMetadata

`func (o *PaymentOptionsRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentOptionsRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentOptionsRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentOptionsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentOptionsRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentOptionsRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCartItems

`func (o *PaymentOptionsRequest) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *PaymentOptionsRequest) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *PaymentOptionsRequest) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *PaymentOptionsRequest) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItemsNil

`func (o *PaymentOptionsRequest) SetCartItemsNil(b bool)`

 SetCartItemsNil sets the value for CartItems to be an explicit nil

### UnsetCartItems
`func (o *PaymentOptionsRequest) UnsetCartItems()`

UnsetCartItems ensures that no value is present for CartItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


