# PaymentConnectorResponseTransactionCaptureSucceededEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service used. | [optional] 
**PaymentServiceDisplayName** | Pointer to **string** | The display name of the payment service used. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The payment service definition used. | [optional] 
**PaymentServiceTransactionId** | Pointer to **NullableString** | The external ID of the transaction as set by the payment service. | [optional] 
**Status** | Pointer to **string** | The new status code for the transaction. This is always set to &#x60;capture_succeeded&#x60;. | [optional] 
**InstrumentType** | Pointer to **NullableString** | The type of instrument used for this transaction. | [optional] 
**RetryRule** | Pointer to **NullableString** | This will always be &#x60;null&#x60; because the transaction succeeded. | [optional] 
**RawResponseCode** | Pointer to **NullableString** | This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**RawResponseDescription** | Pointer to **NullableString** | This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**AvsResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - &#x60;no_match&#x60; - neither address or postal code match - &#x60;match&#x60; - both address and postal code match - &#x60;partial_match_address&#x60; - address matches but postal code does not - &#x60;partial_match_postcode&#x60; - postal code matches but address does not - &#x60;unavailable &#x60; - AVS is unavailable for card/country  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**CvvResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - &#x60;no_match&#x60; - the CVV does not match the expected value - &#x60;match&#x60; - the CVV matches the expected value - &#x60;unavailable &#x60; - CVV check unavailable for card our country - &#x60;not_provided &#x60; - CVV not provided  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**PaymentMethodScheme** | Pointer to **NullableString** | The card scheme sent to the connector. | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionCaptureSucceededEventContext

`func NewPaymentConnectorResponseTransactionCaptureSucceededEventContext() *PaymentConnectorResponseTransactionCaptureSucceededEventContext`

NewPaymentConnectorResponseTransactionCaptureSucceededEventContext instantiates a new PaymentConnectorResponseTransactionCaptureSucceededEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionCaptureSucceededEventContextWithDefaults

`func NewPaymentConnectorResponseTransactionCaptureSucceededEventContextWithDefaults() *PaymentConnectorResponseTransactionCaptureSucceededEventContext`

NewPaymentConnectorResponseTransactionCaptureSucceededEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionCaptureSucceededEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.

### SetPaymentServiceTransactionIdNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceTransactionIdNil(b bool)`

 SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil

### UnsetPaymentServiceTransactionId
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetPaymentServiceTransactionId()`

UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
### GetStatus

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInstrumentType

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### SetInstrumentTypeNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetInstrumentTypeNil(b bool)`

 SetInstrumentTypeNil sets the value for InstrumentType to be an explicit nil

### UnsetInstrumentType
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetInstrumentType()`

UnsetInstrumentType ensures that no value is present for InstrumentType, not even an explicit nil
### GetRetryRule

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRetryRule() string`

GetRetryRule returns the RetryRule field if non-nil, zero value otherwise.

### GetRetryRuleOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRetryRuleOk() (*string, bool)`

GetRetryRuleOk returns a tuple with the RetryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryRule

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRetryRule(v string)`

SetRetryRule sets RetryRule field to given value.

### HasRetryRule

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasRetryRule() bool`

HasRetryRule returns a boolean if a field has been set.

### SetRetryRuleNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRetryRuleNil(b bool)`

 SetRetryRuleNil sets the value for RetryRule to be an explicit nil

### UnsetRetryRule
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetRetryRule()`

UnsetRetryRule ensures that no value is present for RetryRule, not even an explicit nil
### GetRawResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseCode() string`

GetRawResponseCode returns the RawResponseCode field if non-nil, zero value otherwise.

### GetRawResponseCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseCodeOk() (*string, bool)`

GetRawResponseCodeOk returns a tuple with the RawResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseCode(v string)`

SetRawResponseCode sets RawResponseCode field to given value.

### HasRawResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasRawResponseCode() bool`

HasRawResponseCode returns a boolean if a field has been set.

### SetRawResponseCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseCodeNil(b bool)`

 SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil

### UnsetRawResponseCode
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetRawResponseCode()`

UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
### GetRawResponseDescription

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseDescription() string`

GetRawResponseDescription returns the RawResponseDescription field if non-nil, zero value otherwise.

### GetRawResponseDescriptionOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseDescriptionOk() (*string, bool)`

GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseDescription

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseDescription(v string)`

SetRawResponseDescription sets RawResponseDescription field to given value.

### HasRawResponseDescription

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasRawResponseDescription() bool`

HasRawResponseDescription returns a boolean if a field has been set.

### SetRawResponseDescriptionNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseDescriptionNil(b bool)`

 SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil

### UnsetRawResponseDescription
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetRawResponseDescription()`

UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
### GetAvsResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetAvsResponseCode() string`

GetAvsResponseCode returns the AvsResponseCode field if non-nil, zero value otherwise.

### GetAvsResponseCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetAvsResponseCodeOk() (*string, bool)`

GetAvsResponseCodeOk returns a tuple with the AvsResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetAvsResponseCode(v string)`

SetAvsResponseCode sets AvsResponseCode field to given value.

### HasAvsResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasAvsResponseCode() bool`

HasAvsResponseCode returns a boolean if a field has been set.

### SetAvsResponseCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetAvsResponseCodeNil(b bool)`

 SetAvsResponseCodeNil sets the value for AvsResponseCode to be an explicit nil

### UnsetAvsResponseCode
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetAvsResponseCode()`

UnsetAvsResponseCode ensures that no value is present for AvsResponseCode, not even an explicit nil
### GetCvvResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetCvvResponseCode() string`

GetCvvResponseCode returns the CvvResponseCode field if non-nil, zero value otherwise.

### GetCvvResponseCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetCvvResponseCodeOk() (*string, bool)`

GetCvvResponseCodeOk returns a tuple with the CvvResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetCvvResponseCode(v string)`

SetCvvResponseCode sets CvvResponseCode field to given value.

### HasCvvResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasCvvResponseCode() bool`

HasCvvResponseCode returns a boolean if a field has been set.

### SetCvvResponseCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetCvvResponseCodeNil(b bool)`

 SetCvvResponseCodeNil sets the value for CvvResponseCode to be an explicit nil

### UnsetCvvResponseCode
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetCvvResponseCode()`

UnsetCvvResponseCode ensures that no value is present for CvvResponseCode, not even an explicit nil
### GetPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentMethodScheme() string`

GetPaymentMethodScheme returns the PaymentMethodScheme field if non-nil, zero value otherwise.

### GetPaymentMethodSchemeOk

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentMethodSchemeOk() (*string, bool)`

GetPaymentMethodSchemeOk returns a tuple with the PaymentMethodScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentMethodScheme(v string)`

SetPaymentMethodScheme sets PaymentMethodScheme field to given value.

### HasPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentMethodScheme() bool`

HasPaymentMethodScheme returns a boolean if a field has been set.

### SetPaymentMethodSchemeNil

`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentMethodSchemeNil(b bool)`

 SetPaymentMethodSchemeNil sets the value for PaymentMethodScheme to be an explicit nil

### UnsetPaymentMethodScheme
`func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetPaymentMethodScheme()`

UnsetPaymentMethodScheme ensures that no value is present for PaymentMethodScheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


