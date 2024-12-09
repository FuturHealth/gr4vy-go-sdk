# PaymentServiceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A custom name for the payment service. This will be shown in the Admin UI. | [optional] 
**Fields** | Pointer to [**[]PaymentServiceUpdateFieldsInner**](PaymentServiceUpdateFieldsInner.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card &#x60;secret_key&#x60; is required and so must be sent within this field. | [optional] 
**ReportingFields** | Pointer to [**[]PaymentServiceUpdateReportingFieldsInner**](PaymentServiceUpdateReportingFieldsInner.md) | The &#x60;reporting_fields&#x60; field should contain a list of key-value pairs. Each key-value pair represents a reporting field defined by the payment service. For example, when enabling settlement reporting for &#x60;nuvei-card&#x60;, the &#x60;ssh_username&#x60; field is required and must be included in &#x60;reporting_fields&#x60;. | [optional] 
**AcceptedCountries** | Pointer to **[]string** | A list of countries that this payment service needs to support in ISO two-letter code format. | [optional] 
**AcceptedCurrencies** | Pointer to **[]string** | A list of currencies that this payment service needs to support in ISO 4217 three-letter code format. | [optional] 
**ThreeDSecureEnabled** | Pointer to **bool** | Defines if 3-D Secure is enabled for the service. This feature can only be enabled if the payment service definition supports the &#x60;three_d_secure_hosted&#x60; feature. This does not affect pass through 3-D Secure data. | [optional] [default to false]
**MerchantProfile** | Pointer to [**NullableMerchantProfile**](MerchantProfile.md) | Configuration for each supported card scheme. When updating a Payment Service, a key not being present will indicate no updates to be done on that scheme, whereas an object being sent as Null for a key will empty the configuration for that scheme. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**OpenLoop** | Pointer to **NullableBool** | Defines if the service works as an open-loop service. This feature can only be enabled if the PSP is set up to accept previous scheme transaction IDs.  If this value is set to &#x60;null&#x60;, it will be set to the value of &#x60;open_loop&#x60; in the payment service definition.  If &#x60;open_loop_toggle&#x60; is &#x60;false&#x60; in the payment service definition, &#x60;open_loop&#x60; should either not be provided or set to &#x60;null&#x60;, or it will fail with a validation error. | [optional] 
**PaymentMethodTokenizationEnabled** | Pointer to **NullableBool** | Defines if tokenization is enabled for the service. This feature can only be enabled if the payment service is NOT set as &#x60;open_loop&#x60; and the PSP is set up to tokenize. | [optional] 
**NetworkTokensEnabled** | Pointer to **NullableBool** | Defines if network tokens are enabled for the service. This feature can only be enabled if the payment service is set as &#x60;open_loop&#x60; and the PSP is set up to accept network tokens.  If this value is set to &#x60;null&#x60;, it will be set to the value of &#x60;network_tokens_default&#x60; in the payment service definition.  If &#x60;network_tokens_toggle&#x60; is &#x60;false&#x60; in the payment service definition, &#x60;network_tokens_enabled&#x60; should either not be provided or set to &#x60;null&#x60;, or it will fail with a validation error. | [optional] 
**SettlementReportingEnabled** | Pointer to **bool** | Defines if settlement reporting is enabled for the service. This feature can only be enabled if the payment service definition supports the &#x60;settlement_reporting&#x60; feature. | [optional] [default to false]

## Methods

### NewPaymentServiceUpdate

`func NewPaymentServiceUpdate() *PaymentServiceUpdate`

NewPaymentServiceUpdate instantiates a new PaymentServiceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceUpdateWithDefaults

`func NewPaymentServiceUpdateWithDefaults() *PaymentServiceUpdate`

NewPaymentServiceUpdateWithDefaults instantiates a new PaymentServiceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PaymentServiceUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentServiceUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *PaymentServiceUpdate) GetFields() []PaymentServiceUpdateFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentServiceUpdate) GetFieldsOk() (*[]PaymentServiceUpdateFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentServiceUpdate) SetFields(v []PaymentServiceUpdateFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PaymentServiceUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetReportingFields

`func (o *PaymentServiceUpdate) GetReportingFields() []PaymentServiceUpdateReportingFieldsInner`

GetReportingFields returns the ReportingFields field if non-nil, zero value otherwise.

### GetReportingFieldsOk

`func (o *PaymentServiceUpdate) GetReportingFieldsOk() (*[]PaymentServiceUpdateReportingFieldsInner, bool)`

GetReportingFieldsOk returns a tuple with the ReportingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingFields

`func (o *PaymentServiceUpdate) SetReportingFields(v []PaymentServiceUpdateReportingFieldsInner)`

SetReportingFields sets ReportingFields field to given value.

### HasReportingFields

`func (o *PaymentServiceUpdate) HasReportingFields() bool`

HasReportingFields returns a boolean if a field has been set.

### GetAcceptedCountries

`func (o *PaymentServiceUpdate) GetAcceptedCountries() []string`

GetAcceptedCountries returns the AcceptedCountries field if non-nil, zero value otherwise.

### GetAcceptedCountriesOk

`func (o *PaymentServiceUpdate) GetAcceptedCountriesOk() (*[]string, bool)`

GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountries

`func (o *PaymentServiceUpdate) SetAcceptedCountries(v []string)`

SetAcceptedCountries sets AcceptedCountries field to given value.

### HasAcceptedCountries

`func (o *PaymentServiceUpdate) HasAcceptedCountries() bool`

HasAcceptedCountries returns a boolean if a field has been set.

### GetAcceptedCurrencies

`func (o *PaymentServiceUpdate) GetAcceptedCurrencies() []string`

GetAcceptedCurrencies returns the AcceptedCurrencies field if non-nil, zero value otherwise.

### GetAcceptedCurrenciesOk

`func (o *PaymentServiceUpdate) GetAcceptedCurrenciesOk() (*[]string, bool)`

GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCurrencies

`func (o *PaymentServiceUpdate) SetAcceptedCurrencies(v []string)`

SetAcceptedCurrencies sets AcceptedCurrencies field to given value.

### HasAcceptedCurrencies

`func (o *PaymentServiceUpdate) HasAcceptedCurrencies() bool`

HasAcceptedCurrencies returns a boolean if a field has been set.

### GetThreeDSecureEnabled

`func (o *PaymentServiceUpdate) GetThreeDSecureEnabled() bool`

GetThreeDSecureEnabled returns the ThreeDSecureEnabled field if non-nil, zero value otherwise.

### GetThreeDSecureEnabledOk

`func (o *PaymentServiceUpdate) GetThreeDSecureEnabledOk() (*bool, bool)`

GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureEnabled

`func (o *PaymentServiceUpdate) SetThreeDSecureEnabled(v bool)`

SetThreeDSecureEnabled sets ThreeDSecureEnabled field to given value.

### HasThreeDSecureEnabled

`func (o *PaymentServiceUpdate) HasThreeDSecureEnabled() bool`

HasThreeDSecureEnabled returns a boolean if a field has been set.

### GetMerchantProfile

`func (o *PaymentServiceUpdate) GetMerchantProfile() MerchantProfile`

GetMerchantProfile returns the MerchantProfile field if non-nil, zero value otherwise.

### GetMerchantProfileOk

`func (o *PaymentServiceUpdate) GetMerchantProfileOk() (*MerchantProfile, bool)`

GetMerchantProfileOk returns a tuple with the MerchantProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProfile

`func (o *PaymentServiceUpdate) SetMerchantProfile(v MerchantProfile)`

SetMerchantProfile sets MerchantProfile field to given value.

### HasMerchantProfile

`func (o *PaymentServiceUpdate) HasMerchantProfile() bool`

HasMerchantProfile returns a boolean if a field has been set.

### SetMerchantProfileNil

`func (o *PaymentServiceUpdate) SetMerchantProfileNil(b bool)`

 SetMerchantProfileNil sets the value for MerchantProfile to be an explicit nil

### UnsetMerchantProfile
`func (o *PaymentServiceUpdate) UnsetMerchantProfile()`

UnsetMerchantProfile ensures that no value is present for MerchantProfile, not even an explicit nil
### GetActive

`func (o *PaymentServiceUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentServiceUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentServiceUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentServiceUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOpenLoop

`func (o *PaymentServiceUpdate) GetOpenLoop() bool`

GetOpenLoop returns the OpenLoop field if non-nil, zero value otherwise.

### GetOpenLoopOk

`func (o *PaymentServiceUpdate) GetOpenLoopOk() (*bool, bool)`

GetOpenLoopOk returns a tuple with the OpenLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLoop

`func (o *PaymentServiceUpdate) SetOpenLoop(v bool)`

SetOpenLoop sets OpenLoop field to given value.

### HasOpenLoop

`func (o *PaymentServiceUpdate) HasOpenLoop() bool`

HasOpenLoop returns a boolean if a field has been set.

### SetOpenLoopNil

`func (o *PaymentServiceUpdate) SetOpenLoopNil(b bool)`

 SetOpenLoopNil sets the value for OpenLoop to be an explicit nil

### UnsetOpenLoop
`func (o *PaymentServiceUpdate) UnsetOpenLoop()`

UnsetOpenLoop ensures that no value is present for OpenLoop, not even an explicit nil
### GetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceUpdate) GetPaymentMethodTokenizationEnabled() bool`

GetPaymentMethodTokenizationEnabled returns the PaymentMethodTokenizationEnabled field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationEnabledOk

`func (o *PaymentServiceUpdate) GetPaymentMethodTokenizationEnabledOk() (*bool, bool)`

GetPaymentMethodTokenizationEnabledOk returns a tuple with the PaymentMethodTokenizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceUpdate) SetPaymentMethodTokenizationEnabled(v bool)`

SetPaymentMethodTokenizationEnabled sets PaymentMethodTokenizationEnabled field to given value.

### HasPaymentMethodTokenizationEnabled

`func (o *PaymentServiceUpdate) HasPaymentMethodTokenizationEnabled() bool`

HasPaymentMethodTokenizationEnabled returns a boolean if a field has been set.

### SetPaymentMethodTokenizationEnabledNil

`func (o *PaymentServiceUpdate) SetPaymentMethodTokenizationEnabledNil(b bool)`

 SetPaymentMethodTokenizationEnabledNil sets the value for PaymentMethodTokenizationEnabled to be an explicit nil

### UnsetPaymentMethodTokenizationEnabled
`func (o *PaymentServiceUpdate) UnsetPaymentMethodTokenizationEnabled()`

UnsetPaymentMethodTokenizationEnabled ensures that no value is present for PaymentMethodTokenizationEnabled, not even an explicit nil
### GetNetworkTokensEnabled

`func (o *PaymentServiceUpdate) GetNetworkTokensEnabled() bool`

GetNetworkTokensEnabled returns the NetworkTokensEnabled field if non-nil, zero value otherwise.

### GetNetworkTokensEnabledOk

`func (o *PaymentServiceUpdate) GetNetworkTokensEnabledOk() (*bool, bool)`

GetNetworkTokensEnabledOk returns a tuple with the NetworkTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokensEnabled

`func (o *PaymentServiceUpdate) SetNetworkTokensEnabled(v bool)`

SetNetworkTokensEnabled sets NetworkTokensEnabled field to given value.

### HasNetworkTokensEnabled

`func (o *PaymentServiceUpdate) HasNetworkTokensEnabled() bool`

HasNetworkTokensEnabled returns a boolean if a field has been set.

### SetNetworkTokensEnabledNil

`func (o *PaymentServiceUpdate) SetNetworkTokensEnabledNil(b bool)`

 SetNetworkTokensEnabledNil sets the value for NetworkTokensEnabled to be an explicit nil

### UnsetNetworkTokensEnabled
`func (o *PaymentServiceUpdate) UnsetNetworkTokensEnabled()`

UnsetNetworkTokensEnabled ensures that no value is present for NetworkTokensEnabled, not even an explicit nil
### GetSettlementReportingEnabled

`func (o *PaymentServiceUpdate) GetSettlementReportingEnabled() bool`

GetSettlementReportingEnabled returns the SettlementReportingEnabled field if non-nil, zero value otherwise.

### GetSettlementReportingEnabledOk

`func (o *PaymentServiceUpdate) GetSettlementReportingEnabledOk() (*bool, bool)`

GetSettlementReportingEnabledOk returns a tuple with the SettlementReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementReportingEnabled

`func (o *PaymentServiceUpdate) SetSettlementReportingEnabled(v bool)`

SetSettlementReportingEnabled sets SettlementReportingEnabled field to given value.

### HasSettlementReportingEnabled

`func (o *PaymentServiceUpdate) HasSettlementReportingEnabled() bool`

HasSettlementReportingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


