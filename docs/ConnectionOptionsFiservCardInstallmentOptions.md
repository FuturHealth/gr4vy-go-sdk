# ConnectionOptionsFiservCardInstallmentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfInstallments** | Pointer to **NullableFloat32** | Number of installments for a sale transaction if the customer pays the total amount in multiple transactions. | [optional] 
**InstallmentsInterest** | Pointer to **NullableBool** | Indicates whether the installment interest amount has been applied. | [optional] 
**InstallmentDelayMonths** | Pointer to **NullableFloat32** | The number of months the first installment payment will be delayed. | [optional] 
**MerchantAdviceCodeSupported** | Pointer to **NullableBool** | Indicates if the merchant supports merchant advice code (MAC) in order to receive table 55 code for a declined recurring transaction. | [optional] 

## Methods

### NewConnectionOptionsFiservCardInstallmentOptions

`func NewConnectionOptionsFiservCardInstallmentOptions() *ConnectionOptionsFiservCardInstallmentOptions`

NewConnectionOptionsFiservCardInstallmentOptions instantiates a new ConnectionOptionsFiservCardInstallmentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsFiservCardInstallmentOptionsWithDefaults

`func NewConnectionOptionsFiservCardInstallmentOptionsWithDefaults() *ConnectionOptionsFiservCardInstallmentOptions`

NewConnectionOptionsFiservCardInstallmentOptionsWithDefaults instantiates a new ConnectionOptionsFiservCardInstallmentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfInstallments

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetNumberOfInstallments() float32`

GetNumberOfInstallments returns the NumberOfInstallments field if non-nil, zero value otherwise.

### GetNumberOfInstallmentsOk

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetNumberOfInstallmentsOk() (*float32, bool)`

GetNumberOfInstallmentsOk returns a tuple with the NumberOfInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfInstallments

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetNumberOfInstallments(v float32)`

SetNumberOfInstallments sets NumberOfInstallments field to given value.

### HasNumberOfInstallments

`func (o *ConnectionOptionsFiservCardInstallmentOptions) HasNumberOfInstallments() bool`

HasNumberOfInstallments returns a boolean if a field has been set.

### SetNumberOfInstallmentsNil

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetNumberOfInstallmentsNil(b bool)`

 SetNumberOfInstallmentsNil sets the value for NumberOfInstallments to be an explicit nil

### UnsetNumberOfInstallments
`func (o *ConnectionOptionsFiservCardInstallmentOptions) UnsetNumberOfInstallments()`

UnsetNumberOfInstallments ensures that no value is present for NumberOfInstallments, not even an explicit nil
### GetInstallmentsInterest

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetInstallmentsInterest() bool`

GetInstallmentsInterest returns the InstallmentsInterest field if non-nil, zero value otherwise.

### GetInstallmentsInterestOk

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetInstallmentsInterestOk() (*bool, bool)`

GetInstallmentsInterestOk returns a tuple with the InstallmentsInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentsInterest

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetInstallmentsInterest(v bool)`

SetInstallmentsInterest sets InstallmentsInterest field to given value.

### HasInstallmentsInterest

`func (o *ConnectionOptionsFiservCardInstallmentOptions) HasInstallmentsInterest() bool`

HasInstallmentsInterest returns a boolean if a field has been set.

### SetInstallmentsInterestNil

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetInstallmentsInterestNil(b bool)`

 SetInstallmentsInterestNil sets the value for InstallmentsInterest to be an explicit nil

### UnsetInstallmentsInterest
`func (o *ConnectionOptionsFiservCardInstallmentOptions) UnsetInstallmentsInterest()`

UnsetInstallmentsInterest ensures that no value is present for InstallmentsInterest, not even an explicit nil
### GetInstallmentDelayMonths

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetInstallmentDelayMonths() float32`

GetInstallmentDelayMonths returns the InstallmentDelayMonths field if non-nil, zero value otherwise.

### GetInstallmentDelayMonthsOk

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetInstallmentDelayMonthsOk() (*float32, bool)`

GetInstallmentDelayMonthsOk returns a tuple with the InstallmentDelayMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentDelayMonths

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetInstallmentDelayMonths(v float32)`

SetInstallmentDelayMonths sets InstallmentDelayMonths field to given value.

### HasInstallmentDelayMonths

`func (o *ConnectionOptionsFiservCardInstallmentOptions) HasInstallmentDelayMonths() bool`

HasInstallmentDelayMonths returns a boolean if a field has been set.

### SetInstallmentDelayMonthsNil

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetInstallmentDelayMonthsNil(b bool)`

 SetInstallmentDelayMonthsNil sets the value for InstallmentDelayMonths to be an explicit nil

### UnsetInstallmentDelayMonths
`func (o *ConnectionOptionsFiservCardInstallmentOptions) UnsetInstallmentDelayMonths()`

UnsetInstallmentDelayMonths ensures that no value is present for InstallmentDelayMonths, not even an explicit nil
### GetMerchantAdviceCodeSupported

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetMerchantAdviceCodeSupported() bool`

GetMerchantAdviceCodeSupported returns the MerchantAdviceCodeSupported field if non-nil, zero value otherwise.

### GetMerchantAdviceCodeSupportedOk

`func (o *ConnectionOptionsFiservCardInstallmentOptions) GetMerchantAdviceCodeSupportedOk() (*bool, bool)`

GetMerchantAdviceCodeSupportedOk returns a tuple with the MerchantAdviceCodeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAdviceCodeSupported

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetMerchantAdviceCodeSupported(v bool)`

SetMerchantAdviceCodeSupported sets MerchantAdviceCodeSupported field to given value.

### HasMerchantAdviceCodeSupported

`func (o *ConnectionOptionsFiservCardInstallmentOptions) HasMerchantAdviceCodeSupported() bool`

HasMerchantAdviceCodeSupported returns a boolean if a field has been set.

### SetMerchantAdviceCodeSupportedNil

`func (o *ConnectionOptionsFiservCardInstallmentOptions) SetMerchantAdviceCodeSupportedNil(b bool)`

 SetMerchantAdviceCodeSupportedNil sets the value for MerchantAdviceCodeSupported to be an explicit nil

### UnsetMerchantAdviceCodeSupported
`func (o *ConnectionOptionsFiservCardInstallmentOptions) UnsetMerchantAdviceCodeSupported()`

UnsetMerchantAdviceCodeSupported ensures that no value is present for MerchantAdviceCodeSupported, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


