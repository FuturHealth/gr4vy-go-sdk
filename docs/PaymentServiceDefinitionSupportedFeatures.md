# PaymentServiceDefinitionSupportedFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayedCapture** | Pointer to **bool** | Supports [capturing](#operation/capture-transaction) authorized transactions. | [optional] 
**DigitalWallets** | Pointer to **bool** | Supports passing decrypted digital wallet (e.g. Apple Pay) tokens to the underlying processor. | [optional] 
**NetworkTokensDefault** | Pointer to **bool** | Supports processing network tokens by default. | [optional] 
**NetworkTokensToggle** | Pointer to **bool** | Supports toggling processing of network tokens on or off. | [optional] 
**OpenLoop** | Pointer to **bool** | Supports processing transactions with either raw PAN details or network tokens. | [optional] 
**OpenLoopToggle** | Pointer to **bool** | Supports toggling processing as open-loop on or off. | [optional] 
**PartialRefunds** | Pointer to **bool** | Supports [partially refunding](#operation/refund-transaction) captured transactions. | [optional] 
**PaymentMethodTokenization** | Pointer to **bool** | Supports storing a payment method via tokenization. | [optional] 
**PaymentMethodTokenizationToggle** | Pointer to **bool** | Supports toggling tokenization for a payment method on or off from the dashboard. | [optional] 
**Refunds** | Pointer to **bool** | Supports [refunding](#operation/refund-transaction) captured transactions. | [optional] 
**RequiresWebhookSetup** | Pointer to **bool** | Requires merchant to set up &#x60;webhook_url&#x60; manually with provider. | [optional] 
**ThreeDSecureHosted** | Pointer to **bool** | Supports hosted 3-D Secure with a redirect. | [optional] 
**ThreeDSecurePassThrough** | Pointer to **bool** | Supports passing 3-D Secure data to the underlying processor. | [optional] 
**VerifyCredentials** | Pointer to **bool** | Supports verifying the credentials entered while setting up the underlying processor. This is for internal use only. | [optional] 
**Void** | Pointer to **bool** | Supports [voiding](#operation/void-transaction) authorized transactions. | [optional] 

## Methods

### NewPaymentServiceDefinitionSupportedFeatures

`func NewPaymentServiceDefinitionSupportedFeatures() *PaymentServiceDefinitionSupportedFeatures`

NewPaymentServiceDefinitionSupportedFeatures instantiates a new PaymentServiceDefinitionSupportedFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceDefinitionSupportedFeaturesWithDefaults

`func NewPaymentServiceDefinitionSupportedFeaturesWithDefaults() *PaymentServiceDefinitionSupportedFeatures`

NewPaymentServiceDefinitionSupportedFeaturesWithDefaults instantiates a new PaymentServiceDefinitionSupportedFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayedCapture

`func (o *PaymentServiceDefinitionSupportedFeatures) GetDelayedCapture() bool`

GetDelayedCapture returns the DelayedCapture field if non-nil, zero value otherwise.

### GetDelayedCaptureOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetDelayedCaptureOk() (*bool, bool)`

GetDelayedCaptureOk returns a tuple with the DelayedCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedCapture

`func (o *PaymentServiceDefinitionSupportedFeatures) SetDelayedCapture(v bool)`

SetDelayedCapture sets DelayedCapture field to given value.

### HasDelayedCapture

`func (o *PaymentServiceDefinitionSupportedFeatures) HasDelayedCapture() bool`

HasDelayedCapture returns a boolean if a field has been set.

### GetDigitalWallets

`func (o *PaymentServiceDefinitionSupportedFeatures) GetDigitalWallets() bool`

GetDigitalWallets returns the DigitalWallets field if non-nil, zero value otherwise.

### GetDigitalWalletsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetDigitalWalletsOk() (*bool, bool)`

GetDigitalWalletsOk returns a tuple with the DigitalWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallets

`func (o *PaymentServiceDefinitionSupportedFeatures) SetDigitalWallets(v bool)`

SetDigitalWallets sets DigitalWallets field to given value.

### HasDigitalWallets

`func (o *PaymentServiceDefinitionSupportedFeatures) HasDigitalWallets() bool`

HasDigitalWallets returns a boolean if a field has been set.

### GetNetworkTokensDefault

`func (o *PaymentServiceDefinitionSupportedFeatures) GetNetworkTokensDefault() bool`

GetNetworkTokensDefault returns the NetworkTokensDefault field if non-nil, zero value otherwise.

### GetNetworkTokensDefaultOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetNetworkTokensDefaultOk() (*bool, bool)`

GetNetworkTokensDefaultOk returns a tuple with the NetworkTokensDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokensDefault

`func (o *PaymentServiceDefinitionSupportedFeatures) SetNetworkTokensDefault(v bool)`

SetNetworkTokensDefault sets NetworkTokensDefault field to given value.

### HasNetworkTokensDefault

`func (o *PaymentServiceDefinitionSupportedFeatures) HasNetworkTokensDefault() bool`

HasNetworkTokensDefault returns a boolean if a field has been set.

### GetNetworkTokensToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) GetNetworkTokensToggle() bool`

GetNetworkTokensToggle returns the NetworkTokensToggle field if non-nil, zero value otherwise.

### GetNetworkTokensToggleOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetNetworkTokensToggleOk() (*bool, bool)`

GetNetworkTokensToggleOk returns a tuple with the NetworkTokensToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokensToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) SetNetworkTokensToggle(v bool)`

SetNetworkTokensToggle sets NetworkTokensToggle field to given value.

### HasNetworkTokensToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) HasNetworkTokensToggle() bool`

HasNetworkTokensToggle returns a boolean if a field has been set.

### GetOpenLoop

`func (o *PaymentServiceDefinitionSupportedFeatures) GetOpenLoop() bool`

GetOpenLoop returns the OpenLoop field if non-nil, zero value otherwise.

### GetOpenLoopOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetOpenLoopOk() (*bool, bool)`

GetOpenLoopOk returns a tuple with the OpenLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLoop

`func (o *PaymentServiceDefinitionSupportedFeatures) SetOpenLoop(v bool)`

SetOpenLoop sets OpenLoop field to given value.

### HasOpenLoop

`func (o *PaymentServiceDefinitionSupportedFeatures) HasOpenLoop() bool`

HasOpenLoop returns a boolean if a field has been set.

### GetOpenLoopToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) GetOpenLoopToggle() bool`

GetOpenLoopToggle returns the OpenLoopToggle field if non-nil, zero value otherwise.

### GetOpenLoopToggleOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetOpenLoopToggleOk() (*bool, bool)`

GetOpenLoopToggleOk returns a tuple with the OpenLoopToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLoopToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) SetOpenLoopToggle(v bool)`

SetOpenLoopToggle sets OpenLoopToggle field to given value.

### HasOpenLoopToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) HasOpenLoopToggle() bool`

HasOpenLoopToggle returns a boolean if a field has been set.

### GetPartialRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPartialRefunds() bool`

GetPartialRefunds returns the PartialRefunds field if non-nil, zero value otherwise.

### GetPartialRefundsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPartialRefundsOk() (*bool, bool)`

GetPartialRefundsOk returns a tuple with the PartialRefunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) SetPartialRefunds(v bool)`

SetPartialRefunds sets PartialRefunds field to given value.

### HasPartialRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) HasPartialRefunds() bool`

HasPartialRefunds returns a boolean if a field has been set.

### GetPaymentMethodTokenization

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenization() bool`

GetPaymentMethodTokenization returns the PaymentMethodTokenization field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationOk() (*bool, bool)`

GetPaymentMethodTokenizationOk returns a tuple with the PaymentMethodTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenization

`func (o *PaymentServiceDefinitionSupportedFeatures) SetPaymentMethodTokenization(v bool)`

SetPaymentMethodTokenization sets PaymentMethodTokenization field to given value.

### HasPaymentMethodTokenization

`func (o *PaymentServiceDefinitionSupportedFeatures) HasPaymentMethodTokenization() bool`

HasPaymentMethodTokenization returns a boolean if a field has been set.

### GetPaymentMethodTokenizationToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationToggle() bool`

GetPaymentMethodTokenizationToggle returns the PaymentMethodTokenizationToggle field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationToggleOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationToggleOk() (*bool, bool)`

GetPaymentMethodTokenizationToggleOk returns a tuple with the PaymentMethodTokenizationToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) SetPaymentMethodTokenizationToggle(v bool)`

SetPaymentMethodTokenizationToggle sets PaymentMethodTokenizationToggle field to given value.

### HasPaymentMethodTokenizationToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) HasPaymentMethodTokenizationToggle() bool`

HasPaymentMethodTokenizationToggle returns a boolean if a field has been set.

### GetRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) GetRefunds() bool`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetRefundsOk() (*bool, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) SetRefunds(v bool)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetRequiresWebhookSetup

`func (o *PaymentServiceDefinitionSupportedFeatures) GetRequiresWebhookSetup() bool`

GetRequiresWebhookSetup returns the RequiresWebhookSetup field if non-nil, zero value otherwise.

### GetRequiresWebhookSetupOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetRequiresWebhookSetupOk() (*bool, bool)`

GetRequiresWebhookSetupOk returns a tuple with the RequiresWebhookSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresWebhookSetup

`func (o *PaymentServiceDefinitionSupportedFeatures) SetRequiresWebhookSetup(v bool)`

SetRequiresWebhookSetup sets RequiresWebhookSetup field to given value.

### HasRequiresWebhookSetup

`func (o *PaymentServiceDefinitionSupportedFeatures) HasRequiresWebhookSetup() bool`

HasRequiresWebhookSetup returns a boolean if a field has been set.

### GetThreeDSecureHosted

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecureHosted() bool`

GetThreeDSecureHosted returns the ThreeDSecureHosted field if non-nil, zero value otherwise.

### GetThreeDSecureHostedOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecureHostedOk() (*bool, bool)`

GetThreeDSecureHostedOk returns a tuple with the ThreeDSecureHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureHosted

`func (o *PaymentServiceDefinitionSupportedFeatures) SetThreeDSecureHosted(v bool)`

SetThreeDSecureHosted sets ThreeDSecureHosted field to given value.

### HasThreeDSecureHosted

`func (o *PaymentServiceDefinitionSupportedFeatures) HasThreeDSecureHosted() bool`

HasThreeDSecureHosted returns a boolean if a field has been set.

### GetThreeDSecurePassThrough

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecurePassThrough() bool`

GetThreeDSecurePassThrough returns the ThreeDSecurePassThrough field if non-nil, zero value otherwise.

### GetThreeDSecurePassThroughOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecurePassThroughOk() (*bool, bool)`

GetThreeDSecurePassThroughOk returns a tuple with the ThreeDSecurePassThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecurePassThrough

`func (o *PaymentServiceDefinitionSupportedFeatures) SetThreeDSecurePassThrough(v bool)`

SetThreeDSecurePassThrough sets ThreeDSecurePassThrough field to given value.

### HasThreeDSecurePassThrough

`func (o *PaymentServiceDefinitionSupportedFeatures) HasThreeDSecurePassThrough() bool`

HasThreeDSecurePassThrough returns a boolean if a field has been set.

### GetVerifyCredentials

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVerifyCredentials() bool`

GetVerifyCredentials returns the VerifyCredentials field if non-nil, zero value otherwise.

### GetVerifyCredentialsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVerifyCredentialsOk() (*bool, bool)`

GetVerifyCredentialsOk returns a tuple with the VerifyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentials

`func (o *PaymentServiceDefinitionSupportedFeatures) SetVerifyCredentials(v bool)`

SetVerifyCredentials sets VerifyCredentials field to given value.

### HasVerifyCredentials

`func (o *PaymentServiceDefinitionSupportedFeatures) HasVerifyCredentials() bool`

HasVerifyCredentials returns a boolean if a field has been set.

### GetVoid

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVoid() bool`

GetVoid returns the Void field if non-nil, zero value otherwise.

### GetVoidOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVoidOk() (*bool, bool)`

GetVoidOk returns a tuple with the Void field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoid

`func (o *PaymentServiceDefinitionSupportedFeatures) SetVoid(v bool)`

SetVoid sets Void field to given value.

### HasVoid

`func (o *PaymentServiceDefinitionSupportedFeatures) HasVoid() bool`

HasVoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


