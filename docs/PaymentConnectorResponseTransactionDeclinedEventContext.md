# PaymentConnectorResponseTransactionDeclinedEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service used. | [optional] 
**PaymentServiceDisplayName** | Pointer to **string** | The display name of the payment service used. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The payment service definition used. | [optional] 
**PaymentServiceTransactionId** | Pointer to **NullableString** | The external ID of the transaction as set by the payment service. | [optional] 
**Status** | Pointer to **string** | The new status code for the transaction. This is always set to &#x60;authorization_declined&#x60;. | [optional] 
**Code** | Pointer to **NullableString** | A raw response code returned for the decline. | [optional] 
**InstrumentType** | Pointer to **string** | The type of instrument used for this transaction. | [optional] 
**RetryRule** | Pointer to **NullableString** | Defines why the transaction might be retried. A retry is not guaranteed because the maximum number of retries might already have been attempted.  * &#x60;failure&#x60; - the transaction will be retried because of a failure calling   the payment service. * &#x60;retriable_decline&#x60; - the transaction will be retried because a decline code   was received that can be retried. * &#x60;payment_method_replacement&#x60; - the transaction will be retried because a   decline code was received that triggered a payment method replacement. | [optional] 
**RawResponseCode** | Pointer to **NullableString** | This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**RawResponseDescription** | Pointer to **NullableString** | This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**AvsResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - &#x60;no_match&#x60; - neither address or postal code match - &#x60;match&#x60; - both address and postal code match - &#x60;partial_match_address&#x60; - address matches but postal code does not - &#x60;partial_match_postcode&#x60; - postal code matches but address does not - &#x60;unavailable &#x60; - AVS is unavailable for card/country  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**CvvResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - &#x60;no_match&#x60; - the CVV does not match the expected value - &#x60;match&#x60; - the CVV matches the expected value - &#x60;unavailable &#x60; - CVV check unavailable for card our country - &#x60;not_provided &#x60; - CVV not provided  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**PaymentMethodScheme** | Pointer to **NullableString** | The card scheme sent to the connector. | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionDeclinedEventContext

`func NewPaymentConnectorResponseTransactionDeclinedEventContext() *PaymentConnectorResponseTransactionDeclinedEventContext`

NewPaymentConnectorResponseTransactionDeclinedEventContext instantiates a new PaymentConnectorResponseTransactionDeclinedEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionDeclinedEventContextWithDefaults

`func NewPaymentConnectorResponseTransactionDeclinedEventContextWithDefaults() *PaymentConnectorResponseTransactionDeclinedEventContext`

NewPaymentConnectorResponseTransactionDeclinedEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionDeclinedEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.

### SetPaymentServiceTransactionIdNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentServiceTransactionIdNil(b bool)`

 SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil

### UnsetPaymentServiceTransactionId
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetPaymentServiceTransactionId()`

UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
### GetStatus

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetInstrumentType

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetRetryRule

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetRetryRule() string`

GetRetryRule returns the RetryRule field if non-nil, zero value otherwise.

### GetRetryRuleOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetRetryRuleOk() (*string, bool)`

GetRetryRuleOk returns a tuple with the RetryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryRule

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetRetryRule(v string)`

SetRetryRule sets RetryRule field to given value.

### HasRetryRule

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasRetryRule() bool`

HasRetryRule returns a boolean if a field has been set.

### SetRetryRuleNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetRetryRuleNil(b bool)`

 SetRetryRuleNil sets the value for RetryRule to be an explicit nil

### UnsetRetryRule
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetRetryRule()`

UnsetRetryRule ensures that no value is present for RetryRule, not even an explicit nil
### GetRawResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetRawResponseCode() string`

GetRawResponseCode returns the RawResponseCode field if non-nil, zero value otherwise.

### GetRawResponseCodeOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetRawResponseCodeOk() (*string, bool)`

GetRawResponseCodeOk returns a tuple with the RawResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetRawResponseCode(v string)`

SetRawResponseCode sets RawResponseCode field to given value.

### HasRawResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasRawResponseCode() bool`

HasRawResponseCode returns a boolean if a field has been set.

### SetRawResponseCodeNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetRawResponseCodeNil(b bool)`

 SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil

### UnsetRawResponseCode
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetRawResponseCode()`

UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
### GetRawResponseDescription

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetRawResponseDescription() string`

GetRawResponseDescription returns the RawResponseDescription field if non-nil, zero value otherwise.

### GetRawResponseDescriptionOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetRawResponseDescriptionOk() (*string, bool)`

GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseDescription

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetRawResponseDescription(v string)`

SetRawResponseDescription sets RawResponseDescription field to given value.

### HasRawResponseDescription

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasRawResponseDescription() bool`

HasRawResponseDescription returns a boolean if a field has been set.

### SetRawResponseDescriptionNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetRawResponseDescriptionNil(b bool)`

 SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil

### UnsetRawResponseDescription
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetRawResponseDescription()`

UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
### GetAvsResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetAvsResponseCode() string`

GetAvsResponseCode returns the AvsResponseCode field if non-nil, zero value otherwise.

### GetAvsResponseCodeOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetAvsResponseCodeOk() (*string, bool)`

GetAvsResponseCodeOk returns a tuple with the AvsResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetAvsResponseCode(v string)`

SetAvsResponseCode sets AvsResponseCode field to given value.

### HasAvsResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasAvsResponseCode() bool`

HasAvsResponseCode returns a boolean if a field has been set.

### SetAvsResponseCodeNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetAvsResponseCodeNil(b bool)`

 SetAvsResponseCodeNil sets the value for AvsResponseCode to be an explicit nil

### UnsetAvsResponseCode
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetAvsResponseCode()`

UnsetAvsResponseCode ensures that no value is present for AvsResponseCode, not even an explicit nil
### GetCvvResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetCvvResponseCode() string`

GetCvvResponseCode returns the CvvResponseCode field if non-nil, zero value otherwise.

### GetCvvResponseCodeOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetCvvResponseCodeOk() (*string, bool)`

GetCvvResponseCodeOk returns a tuple with the CvvResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetCvvResponseCode(v string)`

SetCvvResponseCode sets CvvResponseCode field to given value.

### HasCvvResponseCode

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasCvvResponseCode() bool`

HasCvvResponseCode returns a boolean if a field has been set.

### SetCvvResponseCodeNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetCvvResponseCodeNil(b bool)`

 SetCvvResponseCodeNil sets the value for CvvResponseCode to be an explicit nil

### UnsetCvvResponseCode
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetCvvResponseCode()`

UnsetCvvResponseCode ensures that no value is present for CvvResponseCode, not even an explicit nil
### GetPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentMethodScheme() string`

GetPaymentMethodScheme returns the PaymentMethodScheme field if non-nil, zero value otherwise.

### GetPaymentMethodSchemeOk

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) GetPaymentMethodSchemeOk() (*string, bool)`

GetPaymentMethodSchemeOk returns a tuple with the PaymentMethodScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentMethodScheme(v string)`

SetPaymentMethodScheme sets PaymentMethodScheme field to given value.

### HasPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) HasPaymentMethodScheme() bool`

HasPaymentMethodScheme returns a boolean if a field has been set.

### SetPaymentMethodSchemeNil

`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) SetPaymentMethodSchemeNil(b bool)`

 SetPaymentMethodSchemeNil sets the value for PaymentMethodScheme to be an explicit nil

### UnsetPaymentMethodScheme
`func (o *PaymentConnectorResponseTransactionDeclinedEventContext) UnsetPaymentMethodScheme()`

UnsetPaymentMethodScheme ensures that no value is present for PaymentMethodScheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


