# PaymentServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceDefinitionId** | **string** | The ID of the payment service to use. | 
**DisplayName** | **string** | A custom name for the payment service. This will be shown in the Admin UI. | 
**Fields** | [**[]PaymentServiceRequestFieldsInner**](PaymentServiceRequestFieldsInner.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card &#x60;secret_key&#x60; is required and so must be sent within this field. | 
**ReportingFields** | Pointer to [**[]PaymentServiceRequestReportingFieldsInner**](PaymentServiceRequestReportingFieldsInner.md) | The &#x60;reporting_fields&#x60; field should contain a list of key-value pairs. Each key-value pair represents a reporting field defined by the payment service. For example, when enabling settlement reporting for &#x60;nuvei-card&#x60;, the &#x60;ssh_username&#x60; field is required and must be included in &#x60;reporting_fields&#x60;. | [optional] 
**AcceptedCountries** | **[]string** | A list of countries that this payment service needs to support in ISO two-letter code format. | 
**AcceptedCurrencies** | **[]string** | A list of currencies that this payment service needs to support in ISO 4217 three-letter code format. | 
**ThreeDSecureEnabled** | Pointer to **bool** | Defines if 3-D Secure is enabled for the service. This feature can only be enabled if the payment service definition supports the &#x60;three_d_secure_hosted&#x60; feature. This does not affect pass through 3-D Secure data. | [optional] [default to false]
**MerchantProfile** | Pointer to [**NullableMerchantProfile**](MerchantProfile.md) | Configuration for each supported card scheme. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**OpenLoop** | Pointer to **NullableBool** | Defines if the service works as an open-loop service. This feature can only be enabled if the PSP is set up to accept previous scheme transaction IDs.  If this value is not provided or is set to &#x60;null&#x60;, it will be set to the value of &#x60;open_loop&#x60; in the payment service definition.  If &#x60;open_loop_toggle&#x60; is &#x60;false&#x60; in the payment service definition, &#x60;open_loop&#x60; should either not be provided or set to &#x60;null&#x60;, or it will fail with a validation error. | [optional] 
**PaymentMethodTokenizationEnabled** | Pointer to **NullableBool** | Defines if tokenization is enabled for the service. This feature can only be enabled if the payment service is NOT set as &#x60;open_loop&#x60; and the PSP is set up to tokenize. | [optional] 
**NetworkTokensEnabled** | Pointer to **NullableBool** | Defines if network tokens are enabled for the service. This feature can only be enabled if the payment service is set as &#x60;open_loop&#x60; and the PSP is set up to accept network tokens.  If this value is not provided or is set to &#x60;null&#x60;, it will be set to the value of &#x60;network_tokens_default&#x60; in the payment service definition.  If &#x60;network_tokens_toggle&#x60; is &#x60;false&#x60; in the payment service definition, &#x60;network_tokens_enabled&#x60; should either not be provided or set to &#x60;null&#x60;, or it will fail with a validation error. | [optional] 
**SettlementReportingEnabled** | Pointer to **bool** | Defines if settlement reporting is enabled for the service. This feature can only be enabled if the payment service definition supports the &#x60;settlement_reporting&#x60; feature. | [optional] [default to false]

## Methods

### NewPaymentServiceRequest

`func NewPaymentServiceRequest(paymentServiceDefinitionId string, displayName string, fields []PaymentServiceRequestFieldsInner, acceptedCountries []string, acceptedCurrencies []string, ) *PaymentServiceRequest`

NewPaymentServiceRequest instantiates a new PaymentServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceRequestWithDefaults

`func NewPaymentServiceRequestWithDefaults() *PaymentServiceRequest`

NewPaymentServiceRequestWithDefaults instantiates a new PaymentServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceDefinitionId

`func (o *PaymentServiceRequest) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentServiceRequest) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentServiceRequest) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.


### GetDisplayName

`func (o *PaymentServiceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFields

`func (o *PaymentServiceRequest) GetFields() []PaymentServiceRequestFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentServiceRequest) GetFieldsOk() (*[]PaymentServiceRequestFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentServiceRequest) SetFields(v []PaymentServiceRequestFieldsInner)`

SetFields sets Fields field to given value.


### GetReportingFields

`func (o *PaymentServiceRequest) GetReportingFields() []PaymentServiceRequestReportingFieldsInner`

GetReportingFields returns the ReportingFields field if non-nil, zero value otherwise.

### GetReportingFieldsOk

`func (o *PaymentServiceRequest) GetReportingFieldsOk() (*[]PaymentServiceRequestReportingFieldsInner, bool)`

GetReportingFieldsOk returns a tuple with the ReportingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingFields

`func (o *PaymentServiceRequest) SetReportingFields(v []PaymentServiceRequestReportingFieldsInner)`

SetReportingFields sets ReportingFields field to given value.

### HasReportingFields

`func (o *PaymentServiceRequest) HasReportingFields() bool`

HasReportingFields returns a boolean if a field has been set.

### GetAcceptedCountries

`func (o *PaymentServiceRequest) GetAcceptedCountries() []string`

GetAcceptedCountries returns the AcceptedCountries field if non-nil, zero value otherwise.

### GetAcceptedCountriesOk

`func (o *PaymentServiceRequest) GetAcceptedCountriesOk() (*[]string, bool)`

GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountries

`func (o *PaymentServiceRequest) SetAcceptedCountries(v []string)`

SetAcceptedCountries sets AcceptedCountries field to given value.


### GetAcceptedCurrencies

`func (o *PaymentServiceRequest) GetAcceptedCurrencies() []string`

GetAcceptedCurrencies returns the AcceptedCurrencies field if non-nil, zero value otherwise.

### GetAcceptedCurrenciesOk

`func (o *PaymentServiceRequest) GetAcceptedCurrenciesOk() (*[]string, bool)`

GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCurrencies

`func (o *PaymentServiceRequest) SetAcceptedCurrencies(v []string)`

SetAcceptedCurrencies sets AcceptedCurrencies field to given value.


### GetThreeDSecureEnabled

`func (o *PaymentServiceRequest) GetThreeDSecureEnabled() bool`

GetThreeDSecureEnabled returns the ThreeDSecureEnabled field if non-nil, zero value otherwise.

### GetThreeDSecureEnabledOk

`func (o *PaymentServiceRequest) GetThreeDSecureEnabledOk() (*bool, bool)`

GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureEnabled

`func (o *PaymentServiceRequest) SetThreeDSecureEnabled(v bool)`

SetThreeDSecureEnabled sets ThreeDSecureEnabled field to given value.

### HasThreeDSecureEnabled

`func (o *PaymentServiceRequest) HasThreeDSecureEnabled() bool`

HasThreeDSecureEnabled returns a boolean if a field has been set.

### GetMerchantProfile

`func (o *PaymentServiceRequest) GetMerchantProfile() MerchantProfile`

GetMerchantProfile returns the MerchantProfile field if non-nil, zero value otherwise.

### GetMerchantProfileOk

`func (o *PaymentServiceRequest) GetMerchantProfileOk() (*MerchantProfile, bool)`

GetMerchantProfileOk returns a tuple with the MerchantProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProfile

`func (o *PaymentServiceRequest) SetMerchantProfile(v MerchantProfile)`

SetMerchantProfile sets MerchantProfile field to given value.

### HasMerchantProfile

`func (o *PaymentServiceRequest) HasMerchantProfile() bool`

HasMerchantProfile returns a boolean if a field has been set.

### SetMerchantProfileNil

`func (o *PaymentServiceRequest) SetMerchantProfileNil(b bool)`

 SetMerchantProfileNil sets the value for MerchantProfile to be an explicit nil

### UnsetMerchantProfile
`func (o *PaymentServiceRequest) UnsetMerchantProfile()`

UnsetMerchantProfile ensures that no value is present for MerchantProfile, not even an explicit nil
### GetActive

`func (o *PaymentServiceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentServiceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentServiceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentServiceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOpenLoop

`func (o *PaymentServiceRequest) GetOpenLoop() bool`

GetOpenLoop returns the OpenLoop field if non-nil, zero value otherwise.

### GetOpenLoopOk

`func (o *PaymentServiceRequest) GetOpenLoopOk() (*bool, bool)`

GetOpenLoopOk returns a tuple with the OpenLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLoop

`func (o *PaymentServiceRequest) SetOpenLoop(v bool)`

SetOpenLoop sets OpenLoop field to given value.

### HasOpenLoop

`func (o *PaymentServiceRequest) HasOpenLoop() bool`

HasOpenLoop returns a boolean if a field has been set.

### SetOpenLoopNil

`func (o *PaymentServiceRequest) SetOpenLoopNil(b bool)`

 SetOpenLoopNil sets the value for OpenLoop to be an explicit nil

### UnsetOpenLoop
`func (o *PaymentServiceRequest) UnsetOpenLoop()`

UnsetOpenLoop ensures that no value is present for OpenLoop, not even an explicit nil
### GetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceRequest) GetPaymentMethodTokenizationEnabled() bool`

GetPaymentMethodTokenizationEnabled returns the PaymentMethodTokenizationEnabled field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationEnabledOk

`func (o *PaymentServiceRequest) GetPaymentMethodTokenizationEnabledOk() (*bool, bool)`

GetPaymentMethodTokenizationEnabledOk returns a tuple with the PaymentMethodTokenizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceRequest) SetPaymentMethodTokenizationEnabled(v bool)`

SetPaymentMethodTokenizationEnabled sets PaymentMethodTokenizationEnabled field to given value.

### HasPaymentMethodTokenizationEnabled

`func (o *PaymentServiceRequest) HasPaymentMethodTokenizationEnabled() bool`

HasPaymentMethodTokenizationEnabled returns a boolean if a field has been set.

### SetPaymentMethodTokenizationEnabledNil

`func (o *PaymentServiceRequest) SetPaymentMethodTokenizationEnabledNil(b bool)`

 SetPaymentMethodTokenizationEnabledNil sets the value for PaymentMethodTokenizationEnabled to be an explicit nil

### UnsetPaymentMethodTokenizationEnabled
`func (o *PaymentServiceRequest) UnsetPaymentMethodTokenizationEnabled()`

UnsetPaymentMethodTokenizationEnabled ensures that no value is present for PaymentMethodTokenizationEnabled, not even an explicit nil
### GetNetworkTokensEnabled

`func (o *PaymentServiceRequest) GetNetworkTokensEnabled() bool`

GetNetworkTokensEnabled returns the NetworkTokensEnabled field if non-nil, zero value otherwise.

### GetNetworkTokensEnabledOk

`func (o *PaymentServiceRequest) GetNetworkTokensEnabledOk() (*bool, bool)`

GetNetworkTokensEnabledOk returns a tuple with the NetworkTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokensEnabled

`func (o *PaymentServiceRequest) SetNetworkTokensEnabled(v bool)`

SetNetworkTokensEnabled sets NetworkTokensEnabled field to given value.

### HasNetworkTokensEnabled

`func (o *PaymentServiceRequest) HasNetworkTokensEnabled() bool`

HasNetworkTokensEnabled returns a boolean if a field has been set.

### SetNetworkTokensEnabledNil

`func (o *PaymentServiceRequest) SetNetworkTokensEnabledNil(b bool)`

 SetNetworkTokensEnabledNil sets the value for NetworkTokensEnabled to be an explicit nil

### UnsetNetworkTokensEnabled
`func (o *PaymentServiceRequest) UnsetNetworkTokensEnabled()`

UnsetNetworkTokensEnabled ensures that no value is present for NetworkTokensEnabled, not even an explicit nil
### GetSettlementReportingEnabled

`func (o *PaymentServiceRequest) GetSettlementReportingEnabled() bool`

GetSettlementReportingEnabled returns the SettlementReportingEnabled field if non-nil, zero value otherwise.

### GetSettlementReportingEnabledOk

`func (o *PaymentServiceRequest) GetSettlementReportingEnabledOk() (*bool, bool)`

GetSettlementReportingEnabledOk returns a tuple with the SettlementReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementReportingEnabled

`func (o *PaymentServiceRequest) SetSettlementReportingEnabled(v bool)`

SetSettlementReportingEnabled sets SettlementReportingEnabled field to given value.

### HasSettlementReportingEnabled

`func (o *PaymentServiceRequest) HasSettlementReportingEnabled() bool`

HasSettlementReportingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


