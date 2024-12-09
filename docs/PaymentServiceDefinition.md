# PaymentServiceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the payment service. This is the underlying provider followed by a dash followed by the payment method ID. | [optional] 
**Type** | Pointer to **string** | &#x60;payment-service-definition&#x60;. | [optional] [default to "payment-service-definition"]
**DisplayName** | Pointer to **string** | The display name of this service. | [optional] 
**Method** | Pointer to **string** | The ID of the payment method that this services handles. | [optional] 
**Fields** | Pointer to [**[]AntiFraudServiceDefinitionFieldsInner**](AntiFraudServiceDefinitionFieldsInner.md) | A list of fields that need to be submitted when activating the payment. service. | [optional] 
**ReportingFields** | Pointer to [**[]PaymentServiceDefinitionReportingFieldsInner**](PaymentServiceDefinitionReportingFieldsInner.md) | A list of fields that need to be submitted when enabling the payment service settlement reporting. | [optional] 
**SupportedCurrencies** | Pointer to **[]string** | A list of three-letter ISO currency codes that this service supports. | [optional] 
**SupportedCountries** | Pointer to **[]string** | A list of two-letter ISO country codes that this service supports. | [optional] 
**Mode** | Pointer to **string** | The mode of this payment service. | [optional] 
**SupportedFeatures** | Pointer to [**PaymentServiceDefinitionSupportedFeatures**](PaymentServiceDefinitionSupportedFeatures.md) |  | [optional] 
**IconUrl** | Pointer to **NullableString** | An icon to display for the payment service. | [optional] 
**Configuration** | Pointer to [**PaymentServiceDefinitionConfiguration**](PaymentServiceDefinitionConfiguration.md) |  | [optional] 

## Methods

### NewPaymentServiceDefinition

`func NewPaymentServiceDefinition() *PaymentServiceDefinition`

NewPaymentServiceDefinition instantiates a new PaymentServiceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceDefinitionWithDefaults

`func NewPaymentServiceDefinitionWithDefaults() *PaymentServiceDefinition`

NewPaymentServiceDefinitionWithDefaults instantiates a new PaymentServiceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentServiceDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PaymentServiceDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentServiceDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentServiceDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentServiceDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *PaymentServiceDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentServiceDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentServiceDefinition) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentServiceDefinition) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentServiceDefinition) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentServiceDefinition) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetFields

`func (o *PaymentServiceDefinition) GetFields() []AntiFraudServiceDefinitionFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentServiceDefinition) GetFieldsOk() (*[]AntiFraudServiceDefinitionFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentServiceDefinition) SetFields(v []AntiFraudServiceDefinitionFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PaymentServiceDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetReportingFields

`func (o *PaymentServiceDefinition) GetReportingFields() []PaymentServiceDefinitionReportingFieldsInner`

GetReportingFields returns the ReportingFields field if non-nil, zero value otherwise.

### GetReportingFieldsOk

`func (o *PaymentServiceDefinition) GetReportingFieldsOk() (*[]PaymentServiceDefinitionReportingFieldsInner, bool)`

GetReportingFieldsOk returns a tuple with the ReportingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingFields

`func (o *PaymentServiceDefinition) SetReportingFields(v []PaymentServiceDefinitionReportingFieldsInner)`

SetReportingFields sets ReportingFields field to given value.

### HasReportingFields

`func (o *PaymentServiceDefinition) HasReportingFields() bool`

HasReportingFields returns a boolean if a field has been set.

### GetSupportedCurrencies

`func (o *PaymentServiceDefinition) GetSupportedCurrencies() []string`

GetSupportedCurrencies returns the SupportedCurrencies field if non-nil, zero value otherwise.

### GetSupportedCurrenciesOk

`func (o *PaymentServiceDefinition) GetSupportedCurrenciesOk() (*[]string, bool)`

GetSupportedCurrenciesOk returns a tuple with the SupportedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCurrencies

`func (o *PaymentServiceDefinition) SetSupportedCurrencies(v []string)`

SetSupportedCurrencies sets SupportedCurrencies field to given value.

### HasSupportedCurrencies

`func (o *PaymentServiceDefinition) HasSupportedCurrencies() bool`

HasSupportedCurrencies returns a boolean if a field has been set.

### GetSupportedCountries

`func (o *PaymentServiceDefinition) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *PaymentServiceDefinition) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *PaymentServiceDefinition) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.

### HasSupportedCountries

`func (o *PaymentServiceDefinition) HasSupportedCountries() bool`

HasSupportedCountries returns a boolean if a field has been set.

### GetMode

`func (o *PaymentServiceDefinition) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaymentServiceDefinition) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaymentServiceDefinition) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PaymentServiceDefinition) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PaymentServiceDefinition) GetSupportedFeatures() PaymentServiceDefinitionSupportedFeatures`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PaymentServiceDefinition) GetSupportedFeaturesOk() (*PaymentServiceDefinitionSupportedFeatures, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PaymentServiceDefinition) SetSupportedFeatures(v PaymentServiceDefinitionSupportedFeatures)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PaymentServiceDefinition) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetIconUrl

`func (o *PaymentServiceDefinition) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *PaymentServiceDefinition) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *PaymentServiceDefinition) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *PaymentServiceDefinition) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *PaymentServiceDefinition) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *PaymentServiceDefinition) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetConfiguration

`func (o *PaymentServiceDefinition) GetConfiguration() PaymentServiceDefinitionConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PaymentServiceDefinition) GetConfigurationOk() (*PaymentServiceDefinitionConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PaymentServiceDefinition) SetConfiguration(v PaymentServiceDefinitionConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PaymentServiceDefinition) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


