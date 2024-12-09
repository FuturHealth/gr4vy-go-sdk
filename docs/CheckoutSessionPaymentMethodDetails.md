# CheckoutSessionPaymentMethodDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **NullableString** | First six digits of PAN. | [optional] 
**CardType** | Pointer to **NullableString** |  | [optional] 
**CardCountry** | Pointer to **NullableString** | ISO 3166 two letter country code. | [optional] 

## Methods

### NewCheckoutSessionPaymentMethodDetails

`func NewCheckoutSessionPaymentMethodDetails() *CheckoutSessionPaymentMethodDetails`

NewCheckoutSessionPaymentMethodDetails instantiates a new CheckoutSessionPaymentMethodDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionPaymentMethodDetailsWithDefaults

`func NewCheckoutSessionPaymentMethodDetailsWithDefaults() *CheckoutSessionPaymentMethodDetails`

NewCheckoutSessionPaymentMethodDetailsWithDefaults instantiates a new CheckoutSessionPaymentMethodDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *CheckoutSessionPaymentMethodDetails) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *CheckoutSessionPaymentMethodDetails) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *CheckoutSessionPaymentMethodDetails) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *CheckoutSessionPaymentMethodDetails) HasBin() bool`

HasBin returns a boolean if a field has been set.

### SetBinNil

`func (o *CheckoutSessionPaymentMethodDetails) SetBinNil(b bool)`

 SetBinNil sets the value for Bin to be an explicit nil

### UnsetBin
`func (o *CheckoutSessionPaymentMethodDetails) UnsetBin()`

UnsetBin ensures that no value is present for Bin, not even an explicit nil
### GetCardType

`func (o *CheckoutSessionPaymentMethodDetails) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CheckoutSessionPaymentMethodDetails) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CheckoutSessionPaymentMethodDetails) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CheckoutSessionPaymentMethodDetails) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *CheckoutSessionPaymentMethodDetails) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *CheckoutSessionPaymentMethodDetails) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetCardCountry

`func (o *CheckoutSessionPaymentMethodDetails) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *CheckoutSessionPaymentMethodDetails) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *CheckoutSessionPaymentMethodDetails) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *CheckoutSessionPaymentMethodDetails) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### SetCardCountryNil

`func (o *CheckoutSessionPaymentMethodDetails) SetCardCountryNil(b bool)`

 SetCardCountryNil sets the value for CardCountry to be an explicit nil

### UnsetCardCountry
`func (o *CheckoutSessionPaymentMethodDetails) UnsetCardCountry()`

UnsetCardCountry ensures that no value is present for CardCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


