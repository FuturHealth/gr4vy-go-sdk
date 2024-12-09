# TransactionsReportSpecParamsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **[]string** | A list of statuses to filter by. | [optional] 
**CreatedAt** | Pointer to [**TransactionRetriesReportSpecParamsFiltersCreatedAt**](TransactionRetriesReportSpecParamsFiltersCreatedAt.md) |  | [optional] 
**UpdatedAt** | Pointer to [**TransactionsReportSpecParamsFiltersUpdatedAt**](TransactionsReportSpecParamsFiltersUpdatedAt.md) |  | [optional] 
**AuthorizedAt** | Pointer to [**TransactionsReportSpecParamsFiltersAuthorizedAt**](TransactionsReportSpecParamsFiltersAuthorizedAt.md) |  | [optional] 
**CapturedAt** | Pointer to [**TransactionsReportSpecParamsFiltersCapturedAt**](TransactionsReportSpecParamsFiltersCapturedAt.md) |  | [optional] 
**VoidedAt** | Pointer to [**TransactionsReportSpecParamsFiltersVoidedAt**](TransactionsReportSpecParamsFiltersVoidedAt.md) |  | [optional] 
**Currency** | Pointer to **[]string** | A list of ISO-4217 currency codes to filter by. | [optional] 
**Method** | Pointer to **[]string** | A list of payment methods to filter by. | [optional] 
**Scheme** | Pointer to **[]string** | A list of card schemes to filter by. | [optional] 
**Metadata** | Pointer to **[]map[string]interface{}** | A list of metadata as key-value pairs to filter by. | [optional] 
**ThreeDSecureStatus** | Pointer to **[]string** | A list of 3DS challenge statuses to filter by. | [optional] 
**ThreeDSecureEci** | Pointer to **[]string** | A list of electric commerce indicators to filter by. | [optional] 
**ThreeDSecureAuthResp** | Pointer to **[]string** | A list of 3DS responses to filter by. | [optional] 
**PaymentSource** | Pointer to **[]string** | A list of payment sources to filter by. | [optional] 
**MerchantInitiated** | Pointer to **bool** | A flag indicating transactions initiated by the merchant to filter by. | [optional] 
**IsSubsequentPayment** | Pointer to **bool** | A flag indicating transactions with subsequent payments to filter by. | [optional] 

## Methods

### NewTransactionsReportSpecParamsFilters

`func NewTransactionsReportSpecParamsFilters() *TransactionsReportSpecParamsFilters`

NewTransactionsReportSpecParamsFilters instantiates a new TransactionsReportSpecParamsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsReportSpecParamsFiltersWithDefaults

`func NewTransactionsReportSpecParamsFiltersWithDefaults() *TransactionsReportSpecParamsFilters`

NewTransactionsReportSpecParamsFiltersWithDefaults instantiates a new TransactionsReportSpecParamsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TransactionsReportSpecParamsFilters) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionsReportSpecParamsFilters) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionsReportSpecParamsFilters) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionsReportSpecParamsFilters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransactionsReportSpecParamsFilters) GetCreatedAt() TransactionRetriesReportSpecParamsFiltersCreatedAt`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionsReportSpecParamsFilters) GetCreatedAtOk() (*TransactionRetriesReportSpecParamsFiltersCreatedAt, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionsReportSpecParamsFilters) SetCreatedAt(v TransactionRetriesReportSpecParamsFiltersCreatedAt)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionsReportSpecParamsFilters) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TransactionsReportSpecParamsFilters) GetUpdatedAt() TransactionsReportSpecParamsFiltersUpdatedAt`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionsReportSpecParamsFilters) GetUpdatedAtOk() (*TransactionsReportSpecParamsFiltersUpdatedAt, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionsReportSpecParamsFilters) SetUpdatedAt(v TransactionsReportSpecParamsFiltersUpdatedAt)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionsReportSpecParamsFilters) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAuthorizedAt

`func (o *TransactionsReportSpecParamsFilters) GetAuthorizedAt() TransactionsReportSpecParamsFiltersAuthorizedAt`

GetAuthorizedAt returns the AuthorizedAt field if non-nil, zero value otherwise.

### GetAuthorizedAtOk

`func (o *TransactionsReportSpecParamsFilters) GetAuthorizedAtOk() (*TransactionsReportSpecParamsFiltersAuthorizedAt, bool)`

GetAuthorizedAtOk returns a tuple with the AuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAt

`func (o *TransactionsReportSpecParamsFilters) SetAuthorizedAt(v TransactionsReportSpecParamsFiltersAuthorizedAt)`

SetAuthorizedAt sets AuthorizedAt field to given value.

### HasAuthorizedAt

`func (o *TransactionsReportSpecParamsFilters) HasAuthorizedAt() bool`

HasAuthorizedAt returns a boolean if a field has been set.

### GetCapturedAt

`func (o *TransactionsReportSpecParamsFilters) GetCapturedAt() TransactionsReportSpecParamsFiltersCapturedAt`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *TransactionsReportSpecParamsFilters) GetCapturedAtOk() (*TransactionsReportSpecParamsFiltersCapturedAt, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *TransactionsReportSpecParamsFilters) SetCapturedAt(v TransactionsReportSpecParamsFiltersCapturedAt)`

SetCapturedAt sets CapturedAt field to given value.

### HasCapturedAt

`func (o *TransactionsReportSpecParamsFilters) HasCapturedAt() bool`

HasCapturedAt returns a boolean if a field has been set.

### GetVoidedAt

`func (o *TransactionsReportSpecParamsFilters) GetVoidedAt() TransactionsReportSpecParamsFiltersVoidedAt`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *TransactionsReportSpecParamsFilters) GetVoidedAtOk() (*TransactionsReportSpecParamsFiltersVoidedAt, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *TransactionsReportSpecParamsFilters) SetVoidedAt(v TransactionsReportSpecParamsFiltersVoidedAt)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *TransactionsReportSpecParamsFilters) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionsReportSpecParamsFilters) GetCurrency() []string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionsReportSpecParamsFilters) GetCurrencyOk() (*[]string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionsReportSpecParamsFilters) SetCurrency(v []string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionsReportSpecParamsFilters) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMethod

`func (o *TransactionsReportSpecParamsFilters) GetMethod() []string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionsReportSpecParamsFilters) GetMethodOk() (*[]string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionsReportSpecParamsFilters) SetMethod(v []string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TransactionsReportSpecParamsFilters) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetScheme

`func (o *TransactionsReportSpecParamsFilters) GetScheme() []string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *TransactionsReportSpecParamsFilters) GetSchemeOk() (*[]string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *TransactionsReportSpecParamsFilters) SetScheme(v []string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *TransactionsReportSpecParamsFilters) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetMetadata

`func (o *TransactionsReportSpecParamsFilters) GetMetadata() []map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionsReportSpecParamsFilters) GetMetadataOk() (*[]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionsReportSpecParamsFilters) SetMetadata(v []map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionsReportSpecParamsFilters) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetThreeDSecureStatus

`func (o *TransactionsReportSpecParamsFilters) GetThreeDSecureStatus() []string`

GetThreeDSecureStatus returns the ThreeDSecureStatus field if non-nil, zero value otherwise.

### GetThreeDSecureStatusOk

`func (o *TransactionsReportSpecParamsFilters) GetThreeDSecureStatusOk() (*[]string, bool)`

GetThreeDSecureStatusOk returns a tuple with the ThreeDSecureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureStatus

`func (o *TransactionsReportSpecParamsFilters) SetThreeDSecureStatus(v []string)`

SetThreeDSecureStatus sets ThreeDSecureStatus field to given value.

### HasThreeDSecureStatus

`func (o *TransactionsReportSpecParamsFilters) HasThreeDSecureStatus() bool`

HasThreeDSecureStatus returns a boolean if a field has been set.

### GetThreeDSecureEci

`func (o *TransactionsReportSpecParamsFilters) GetThreeDSecureEci() []string`

GetThreeDSecureEci returns the ThreeDSecureEci field if non-nil, zero value otherwise.

### GetThreeDSecureEciOk

`func (o *TransactionsReportSpecParamsFilters) GetThreeDSecureEciOk() (*[]string, bool)`

GetThreeDSecureEciOk returns a tuple with the ThreeDSecureEci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureEci

`func (o *TransactionsReportSpecParamsFilters) SetThreeDSecureEci(v []string)`

SetThreeDSecureEci sets ThreeDSecureEci field to given value.

### HasThreeDSecureEci

`func (o *TransactionsReportSpecParamsFilters) HasThreeDSecureEci() bool`

HasThreeDSecureEci returns a boolean if a field has been set.

### GetThreeDSecureAuthResp

`func (o *TransactionsReportSpecParamsFilters) GetThreeDSecureAuthResp() []string`

GetThreeDSecureAuthResp returns the ThreeDSecureAuthResp field if non-nil, zero value otherwise.

### GetThreeDSecureAuthRespOk

`func (o *TransactionsReportSpecParamsFilters) GetThreeDSecureAuthRespOk() (*[]string, bool)`

GetThreeDSecureAuthRespOk returns a tuple with the ThreeDSecureAuthResp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureAuthResp

`func (o *TransactionsReportSpecParamsFilters) SetThreeDSecureAuthResp(v []string)`

SetThreeDSecureAuthResp sets ThreeDSecureAuthResp field to given value.

### HasThreeDSecureAuthResp

`func (o *TransactionsReportSpecParamsFilters) HasThreeDSecureAuthResp() bool`

HasThreeDSecureAuthResp returns a boolean if a field has been set.

### GetPaymentSource

`func (o *TransactionsReportSpecParamsFilters) GetPaymentSource() []string`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *TransactionsReportSpecParamsFilters) GetPaymentSourceOk() (*[]string, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *TransactionsReportSpecParamsFilters) SetPaymentSource(v []string)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *TransactionsReportSpecParamsFilters) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetMerchantInitiated

`func (o *TransactionsReportSpecParamsFilters) GetMerchantInitiated() bool`

GetMerchantInitiated returns the MerchantInitiated field if non-nil, zero value otherwise.

### GetMerchantInitiatedOk

`func (o *TransactionsReportSpecParamsFilters) GetMerchantInitiatedOk() (*bool, bool)`

GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInitiated

`func (o *TransactionsReportSpecParamsFilters) SetMerchantInitiated(v bool)`

SetMerchantInitiated sets MerchantInitiated field to given value.

### HasMerchantInitiated

`func (o *TransactionsReportSpecParamsFilters) HasMerchantInitiated() bool`

HasMerchantInitiated returns a boolean if a field has been set.

### GetIsSubsequentPayment

`func (o *TransactionsReportSpecParamsFilters) GetIsSubsequentPayment() bool`

GetIsSubsequentPayment returns the IsSubsequentPayment field if non-nil, zero value otherwise.

### GetIsSubsequentPaymentOk

`func (o *TransactionsReportSpecParamsFilters) GetIsSubsequentPaymentOk() (*bool, bool)`

GetIsSubsequentPaymentOk returns a tuple with the IsSubsequentPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubsequentPayment

`func (o *TransactionsReportSpecParamsFilters) SetIsSubsequentPayment(v bool)`

SetIsSubsequentPayment sets IsSubsequentPayment field to given value.

### HasIsSubsequentPayment

`func (o *TransactionsReportSpecParamsFilters) HasIsSubsequentPayment() bool`

HasIsSubsequentPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


