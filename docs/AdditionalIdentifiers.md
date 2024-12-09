# AdditionalIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceAuthorizationId** | Pointer to **NullableString** | The optional ID for the authorization of this transaction. Availability of this ID will vary per connector used. | [optional] 
**PaymentServiceCaptureId** | Pointer to **NullableString** | The optional ID for the capture of this transaction. Availability of this ID will vary per connector used. | [optional] 
**PaymentServiceProcessorId** | Pointer to **NullableString** | The optional ID provided by the processor for this transaction. Availability of this ID will vary per connector used. | [optional] 

## Methods

### NewAdditionalIdentifiers

`func NewAdditionalIdentifiers() *AdditionalIdentifiers`

NewAdditionalIdentifiers instantiates a new AdditionalIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalIdentifiersWithDefaults

`func NewAdditionalIdentifiersWithDefaults() *AdditionalIdentifiers`

NewAdditionalIdentifiersWithDefaults instantiates a new AdditionalIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceAuthorizationId

`func (o *AdditionalIdentifiers) GetPaymentServiceAuthorizationId() string`

GetPaymentServiceAuthorizationId returns the PaymentServiceAuthorizationId field if non-nil, zero value otherwise.

### GetPaymentServiceAuthorizationIdOk

`func (o *AdditionalIdentifiers) GetPaymentServiceAuthorizationIdOk() (*string, bool)`

GetPaymentServiceAuthorizationIdOk returns a tuple with the PaymentServiceAuthorizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceAuthorizationId

`func (o *AdditionalIdentifiers) SetPaymentServiceAuthorizationId(v string)`

SetPaymentServiceAuthorizationId sets PaymentServiceAuthorizationId field to given value.

### HasPaymentServiceAuthorizationId

`func (o *AdditionalIdentifiers) HasPaymentServiceAuthorizationId() bool`

HasPaymentServiceAuthorizationId returns a boolean if a field has been set.

### SetPaymentServiceAuthorizationIdNil

`func (o *AdditionalIdentifiers) SetPaymentServiceAuthorizationIdNil(b bool)`

 SetPaymentServiceAuthorizationIdNil sets the value for PaymentServiceAuthorizationId to be an explicit nil

### UnsetPaymentServiceAuthorizationId
`func (o *AdditionalIdentifiers) UnsetPaymentServiceAuthorizationId()`

UnsetPaymentServiceAuthorizationId ensures that no value is present for PaymentServiceAuthorizationId, not even an explicit nil
### GetPaymentServiceCaptureId

`func (o *AdditionalIdentifiers) GetPaymentServiceCaptureId() string`

GetPaymentServiceCaptureId returns the PaymentServiceCaptureId field if non-nil, zero value otherwise.

### GetPaymentServiceCaptureIdOk

`func (o *AdditionalIdentifiers) GetPaymentServiceCaptureIdOk() (*string, bool)`

GetPaymentServiceCaptureIdOk returns a tuple with the PaymentServiceCaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceCaptureId

`func (o *AdditionalIdentifiers) SetPaymentServiceCaptureId(v string)`

SetPaymentServiceCaptureId sets PaymentServiceCaptureId field to given value.

### HasPaymentServiceCaptureId

`func (o *AdditionalIdentifiers) HasPaymentServiceCaptureId() bool`

HasPaymentServiceCaptureId returns a boolean if a field has been set.

### SetPaymentServiceCaptureIdNil

`func (o *AdditionalIdentifiers) SetPaymentServiceCaptureIdNil(b bool)`

 SetPaymentServiceCaptureIdNil sets the value for PaymentServiceCaptureId to be an explicit nil

### UnsetPaymentServiceCaptureId
`func (o *AdditionalIdentifiers) UnsetPaymentServiceCaptureId()`

UnsetPaymentServiceCaptureId ensures that no value is present for PaymentServiceCaptureId, not even an explicit nil
### GetPaymentServiceProcessorId

`func (o *AdditionalIdentifiers) GetPaymentServiceProcessorId() string`

GetPaymentServiceProcessorId returns the PaymentServiceProcessorId field if non-nil, zero value otherwise.

### GetPaymentServiceProcessorIdOk

`func (o *AdditionalIdentifiers) GetPaymentServiceProcessorIdOk() (*string, bool)`

GetPaymentServiceProcessorIdOk returns a tuple with the PaymentServiceProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceProcessorId

`func (o *AdditionalIdentifiers) SetPaymentServiceProcessorId(v string)`

SetPaymentServiceProcessorId sets PaymentServiceProcessorId field to given value.

### HasPaymentServiceProcessorId

`func (o *AdditionalIdentifiers) HasPaymentServiceProcessorId() bool`

HasPaymentServiceProcessorId returns a boolean if a field has been set.

### SetPaymentServiceProcessorIdNil

`func (o *AdditionalIdentifiers) SetPaymentServiceProcessorIdNil(b bool)`

 SetPaymentServiceProcessorIdNil sets the value for PaymentServiceProcessorId to be an explicit nil

### UnsetPaymentServiceProcessorId
`func (o *AdditionalIdentifiers) UnsetPaymentServiceProcessorId()`

UnsetPaymentServiceProcessorId ensures that no value is present for PaymentServiceProcessorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


