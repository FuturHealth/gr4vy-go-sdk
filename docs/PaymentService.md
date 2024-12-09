# PaymentService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **string** | The ID of this payment service. | [optional] 
**AcceptedCountries** | Pointer to **[]string** | A list of countries for which this service is enabled, in ISO two-letter code format. | [optional] 
**AcceptedCurrencies** | Pointer to **[]string** | A list of currencies for which this service is enabled, in ISO 4217 three-letter code format. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**CreatedAt** | Pointer to **time.Time** | The date and time when this service was created. | [optional] 
**DisplayName** | Pointer to **string** | The custom name set for this service. | [optional] 
**Fields** | Pointer to [**[]GiftCardServiceFieldsInner**](GiftCardServiceFieldsInner.md) | A list of fields, each containing a key-value pair for each field configured for this payment service. Fields marked as &#x60;secret&#x60; (see Payment Service Definition) are not returned. | [optional] 
**ReportingFields** | Pointer to [**[]PaymentServiceReportingFieldsInner**](PaymentServiceReportingFieldsInner.md) | A list of fields, each containing a key-value pair for each field configured for reporting for this payment service. Fields marked as &#x60;secret&#x60; (see Payment Service Definition) are not returned. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**MerchantProfile** | Pointer to [**NullableMerchantProfileSummary**](MerchantProfileSummary.md) | An object containing a key for each supported card scheme (Amex, Discover, Mastercard and Visa), and for each key an object with the merchant profile for this service and the corresponding scheme. | [optional] 
**Method** | Pointer to **string** | The payment method that this service handles. | [optional] 
**NetworkTokensEnabled** | Pointer to **bool** | Defines if network tokens are enabled for the service. This feature can only be enabled if the payment service is set as &#x60;open_loop&#x60; and the PSP is set up to accept network tokens. | [optional] 
**OpenLoop** | Pointer to **bool** | Defines if the service works as an open-loop service. This feature can only be enabled if the PSP is set up to accept previous scheme transaction IDs. | [optional] 
**PaymentMethodTokenizationEnabled** | Pointer to **bool** | Defines if tokenization is enabled for the service. This feature can only be enabled if the payment service is NOT set as &#x60;open_loop&#x60; and the PSP is set up to tokenize. | [optional] [default to false]
**PaymentServiceDefinitionId** | Pointer to **string** | The ID of the payment service definition used to create this service.  | [optional] 
**SettlementReportingEnabled** | Pointer to **bool** | Defines if settlement reporting is enabled for the service. This feature can only be enabled if the payment service definition supports the &#x60;settlement_reporting&#x60; feature. | [optional] 
**Status** | Pointer to **string** | The current status of this service. This will start off as pending, move to created, and might eventually move to an error status if and when the credentials are no longer valid.  | [optional] 
**ThreeDSecureEnabled** | Pointer to **bool** | Defines if 3-D Secure is enabled for the service. This feature can only be enabled if the payment service definition supports the &#x60;three_d_secure_hosted&#x60; feature. This does not affect pass through 3-D Secure data. | [optional] [default to false]
**UpdatedAt** | Pointer to **time.Time** | The date and time when this service was last updated. | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL that needs to be configured with this payment service as the receiving endpoint for webhooks from the service to Gr4vy. Currently, Gr4vy does not yet automatically register webhooks on setup, and therefore webhooks need to be registered manually by the merchant. | [optional] 

## Methods

### NewPaymentService

`func NewPaymentService() *PaymentService`

NewPaymentService instantiates a new PaymentService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceWithDefaults

`func NewPaymentServiceWithDefaults() *PaymentService`

NewPaymentServiceWithDefaults instantiates a new PaymentService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAcceptedCountries

`func (o *PaymentService) GetAcceptedCountries() []string`

GetAcceptedCountries returns the AcceptedCountries field if non-nil, zero value otherwise.

### GetAcceptedCountriesOk

`func (o *PaymentService) GetAcceptedCountriesOk() (*[]string, bool)`

GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountries

`func (o *PaymentService) SetAcceptedCountries(v []string)`

SetAcceptedCountries sets AcceptedCountries field to given value.

### HasAcceptedCountries

`func (o *PaymentService) HasAcceptedCountries() bool`

HasAcceptedCountries returns a boolean if a field has been set.

### GetAcceptedCurrencies

`func (o *PaymentService) GetAcceptedCurrencies() []string`

GetAcceptedCurrencies returns the AcceptedCurrencies field if non-nil, zero value otherwise.

### GetAcceptedCurrenciesOk

`func (o *PaymentService) GetAcceptedCurrenciesOk() (*[]string, bool)`

GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCurrencies

`func (o *PaymentService) SetAcceptedCurrencies(v []string)`

SetAcceptedCurrencies sets AcceptedCurrencies field to given value.

### HasAcceptedCurrencies

`func (o *PaymentService) HasAcceptedCurrencies() bool`

HasAcceptedCurrencies returns a boolean if a field has been set.

### GetActive

`func (o *PaymentService) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentService) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentService) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentService) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *PaymentService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentService) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *PaymentService) GetFields() []GiftCardServiceFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentService) GetFieldsOk() (*[]GiftCardServiceFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentService) SetFields(v []GiftCardServiceFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PaymentService) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetReportingFields

`func (o *PaymentService) GetReportingFields() []PaymentServiceReportingFieldsInner`

GetReportingFields returns the ReportingFields field if non-nil, zero value otherwise.

### GetReportingFieldsOk

`func (o *PaymentService) GetReportingFieldsOk() (*[]PaymentServiceReportingFieldsInner, bool)`

GetReportingFieldsOk returns a tuple with the ReportingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingFields

`func (o *PaymentService) SetReportingFields(v []PaymentServiceReportingFieldsInner)`

SetReportingFields sets ReportingFields field to given value.

### HasReportingFields

`func (o *PaymentService) HasReportingFields() bool`

HasReportingFields returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *PaymentService) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PaymentService) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PaymentService) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *PaymentService) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetMerchantProfile

`func (o *PaymentService) GetMerchantProfile() MerchantProfileSummary`

GetMerchantProfile returns the MerchantProfile field if non-nil, zero value otherwise.

### GetMerchantProfileOk

`func (o *PaymentService) GetMerchantProfileOk() (*MerchantProfileSummary, bool)`

GetMerchantProfileOk returns a tuple with the MerchantProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProfile

`func (o *PaymentService) SetMerchantProfile(v MerchantProfileSummary)`

SetMerchantProfile sets MerchantProfile field to given value.

### HasMerchantProfile

`func (o *PaymentService) HasMerchantProfile() bool`

HasMerchantProfile returns a boolean if a field has been set.

### SetMerchantProfileNil

`func (o *PaymentService) SetMerchantProfileNil(b bool)`

 SetMerchantProfileNil sets the value for MerchantProfile to be an explicit nil

### UnsetMerchantProfile
`func (o *PaymentService) UnsetMerchantProfile()`

UnsetMerchantProfile ensures that no value is present for MerchantProfile, not even an explicit nil
### GetMethod

`func (o *PaymentService) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentService) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentService) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentService) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNetworkTokensEnabled

`func (o *PaymentService) GetNetworkTokensEnabled() bool`

GetNetworkTokensEnabled returns the NetworkTokensEnabled field if non-nil, zero value otherwise.

### GetNetworkTokensEnabledOk

`func (o *PaymentService) GetNetworkTokensEnabledOk() (*bool, bool)`

GetNetworkTokensEnabledOk returns a tuple with the NetworkTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokensEnabled

`func (o *PaymentService) SetNetworkTokensEnabled(v bool)`

SetNetworkTokensEnabled sets NetworkTokensEnabled field to given value.

### HasNetworkTokensEnabled

`func (o *PaymentService) HasNetworkTokensEnabled() bool`

HasNetworkTokensEnabled returns a boolean if a field has been set.

### GetOpenLoop

`func (o *PaymentService) GetOpenLoop() bool`

GetOpenLoop returns the OpenLoop field if non-nil, zero value otherwise.

### GetOpenLoopOk

`func (o *PaymentService) GetOpenLoopOk() (*bool, bool)`

GetOpenLoopOk returns a tuple with the OpenLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLoop

`func (o *PaymentService) SetOpenLoop(v bool)`

SetOpenLoop sets OpenLoop field to given value.

### HasOpenLoop

`func (o *PaymentService) HasOpenLoop() bool`

HasOpenLoop returns a boolean if a field has been set.

### GetPaymentMethodTokenizationEnabled

`func (o *PaymentService) GetPaymentMethodTokenizationEnabled() bool`

GetPaymentMethodTokenizationEnabled returns the PaymentMethodTokenizationEnabled field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationEnabledOk

`func (o *PaymentService) GetPaymentMethodTokenizationEnabledOk() (*bool, bool)`

GetPaymentMethodTokenizationEnabledOk returns a tuple with the PaymentMethodTokenizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationEnabled

`func (o *PaymentService) SetPaymentMethodTokenizationEnabled(v bool)`

SetPaymentMethodTokenizationEnabled sets PaymentMethodTokenizationEnabled field to given value.

### HasPaymentMethodTokenizationEnabled

`func (o *PaymentService) HasPaymentMethodTokenizationEnabled() bool`

HasPaymentMethodTokenizationEnabled returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentService) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentService) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentService) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentService) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetSettlementReportingEnabled

`func (o *PaymentService) GetSettlementReportingEnabled() bool`

GetSettlementReportingEnabled returns the SettlementReportingEnabled field if non-nil, zero value otherwise.

### GetSettlementReportingEnabledOk

`func (o *PaymentService) GetSettlementReportingEnabledOk() (*bool, bool)`

GetSettlementReportingEnabledOk returns a tuple with the SettlementReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementReportingEnabled

`func (o *PaymentService) SetSettlementReportingEnabled(v bool)`

SetSettlementReportingEnabled sets SettlementReportingEnabled field to given value.

### HasSettlementReportingEnabled

`func (o *PaymentService) HasSettlementReportingEnabled() bool`

HasSettlementReportingEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThreeDSecureEnabled

`func (o *PaymentService) GetThreeDSecureEnabled() bool`

GetThreeDSecureEnabled returns the ThreeDSecureEnabled field if non-nil, zero value otherwise.

### GetThreeDSecureEnabledOk

`func (o *PaymentService) GetThreeDSecureEnabledOk() (*bool, bool)`

GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureEnabled

`func (o *PaymentService) SetThreeDSecureEnabled(v bool)`

SetThreeDSecureEnabled sets ThreeDSecureEnabled field to given value.

### HasThreeDSecureEnabled

`func (o *PaymentService) HasThreeDSecureEnabled() bool`

HasThreeDSecureEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PaymentService) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PaymentService) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PaymentService) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PaymentService) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *PaymentService) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *PaymentService) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


