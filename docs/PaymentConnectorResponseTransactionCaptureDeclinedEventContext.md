# PaymentConnectorResponseTransactionCaptureDeclinedEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service used. | [optional] 
**PaymentServiceDisplayName** | Pointer to **string** | The display name of the payment service used. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The payment service definition used. | [optional] 
**PaymentServiceTransactionId** | Pointer to **NullableString** | The external ID of the transaction as set by the payment service. | [optional] 
**Code** | Pointer to **NullableString** | A raw response code returned for the failure. | [optional] 
**RawResponseCode** | Pointer to **NullableString** | This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**RawResponseDescription** | Pointer to **NullableString** | This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**AvsResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - &#x60;no_match&#x60; - neither address or postal code match - &#x60;match&#x60; - both address and postal code match - &#x60;partial_match_address&#x60; - address matches but postal code does not - &#x60;partial_match_postcode&#x60; - postal code matches but address does not - &#x60;unavailable &#x60; - AVS is unavailable for card/country  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**CvvResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - &#x60;no_match&#x60; - the CVV does not match the expected value - &#x60;match&#x60; - the CVV matches the expected value - &#x60;unavailable &#x60; - CVV check unavailable for card our country - &#x60;not_provided &#x60; - CVV not provided  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**PaymentMethodScheme** | Pointer to **NullableString** | The card scheme sent to the connector. | [optional] 

## Methods

### NewPaymentConnectorResponseTransactionCaptureDeclinedEventContext

`func NewPaymentConnectorResponseTransactionCaptureDeclinedEventContext() *PaymentConnectorResponseTransactionCaptureDeclinedEventContext`

NewPaymentConnectorResponseTransactionCaptureDeclinedEventContext instantiates a new PaymentConnectorResponseTransactionCaptureDeclinedEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorResponseTransactionCaptureDeclinedEventContextWithDefaults

`func NewPaymentConnectorResponseTransactionCaptureDeclinedEventContextWithDefaults() *PaymentConnectorResponseTransactionCaptureDeclinedEventContext`

NewPaymentConnectorResponseTransactionCaptureDeclinedEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionCaptureDeclinedEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.

### SetPaymentServiceTransactionIdNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentServiceTransactionIdNil(b bool)`

 SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil

### UnsetPaymentServiceTransactionId
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetPaymentServiceTransactionId()`

UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
### GetCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetRawResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseCode() string`

GetRawResponseCode returns the RawResponseCode field if non-nil, zero value otherwise.

### GetRawResponseCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseCodeOk() (*string, bool)`

GetRawResponseCodeOk returns a tuple with the RawResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseCode(v string)`

SetRawResponseCode sets RawResponseCode field to given value.

### HasRawResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasRawResponseCode() bool`

HasRawResponseCode returns a boolean if a field has been set.

### SetRawResponseCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseCodeNil(b bool)`

 SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil

### UnsetRawResponseCode
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetRawResponseCode()`

UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
### GetRawResponseDescription

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseDescription() string`

GetRawResponseDescription returns the RawResponseDescription field if non-nil, zero value otherwise.

### GetRawResponseDescriptionOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetRawResponseDescriptionOk() (*string, bool)`

GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseDescription

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseDescription(v string)`

SetRawResponseDescription sets RawResponseDescription field to given value.

### HasRawResponseDescription

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasRawResponseDescription() bool`

HasRawResponseDescription returns a boolean if a field has been set.

### SetRawResponseDescriptionNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetRawResponseDescriptionNil(b bool)`

 SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil

### UnsetRawResponseDescription
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetRawResponseDescription()`

UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
### GetAvsResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetAvsResponseCode() string`

GetAvsResponseCode returns the AvsResponseCode field if non-nil, zero value otherwise.

### GetAvsResponseCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetAvsResponseCodeOk() (*string, bool)`

GetAvsResponseCodeOk returns a tuple with the AvsResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetAvsResponseCode(v string)`

SetAvsResponseCode sets AvsResponseCode field to given value.

### HasAvsResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasAvsResponseCode() bool`

HasAvsResponseCode returns a boolean if a field has been set.

### SetAvsResponseCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetAvsResponseCodeNil(b bool)`

 SetAvsResponseCodeNil sets the value for AvsResponseCode to be an explicit nil

### UnsetAvsResponseCode
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetAvsResponseCode()`

UnsetAvsResponseCode ensures that no value is present for AvsResponseCode, not even an explicit nil
### GetCvvResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCvvResponseCode() string`

GetCvvResponseCode returns the CvvResponseCode field if non-nil, zero value otherwise.

### GetCvvResponseCodeOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetCvvResponseCodeOk() (*string, bool)`

GetCvvResponseCodeOk returns a tuple with the CvvResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCvvResponseCode(v string)`

SetCvvResponseCode sets CvvResponseCode field to given value.

### HasCvvResponseCode

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasCvvResponseCode() bool`

HasCvvResponseCode returns a boolean if a field has been set.

### SetCvvResponseCodeNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetCvvResponseCodeNil(b bool)`

 SetCvvResponseCodeNil sets the value for CvvResponseCode to be an explicit nil

### UnsetCvvResponseCode
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetCvvResponseCode()`

UnsetCvvResponseCode ensures that no value is present for CvvResponseCode, not even an explicit nil
### GetPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentMethodScheme() string`

GetPaymentMethodScheme returns the PaymentMethodScheme field if non-nil, zero value otherwise.

### GetPaymentMethodSchemeOk

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) GetPaymentMethodSchemeOk() (*string, bool)`

GetPaymentMethodSchemeOk returns a tuple with the PaymentMethodScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentMethodScheme(v string)`

SetPaymentMethodScheme sets PaymentMethodScheme field to given value.

### HasPaymentMethodScheme

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) HasPaymentMethodScheme() bool`

HasPaymentMethodScheme returns a boolean if a field has been set.

### SetPaymentMethodSchemeNil

`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) SetPaymentMethodSchemeNil(b bool)`

 SetPaymentMethodSchemeNil sets the value for PaymentMethodScheme to be an explicit nil

### UnsetPaymentMethodScheme
`func (o *PaymentConnectorResponseTransactionCaptureDeclinedEventContext) UnsetPaymentMethodScheme()`

UnsetPaymentMethodScheme ensures that no value is present for PaymentMethodScheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


