# AntiFraudTransactionStatusUpdateErrorEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceId** | Pointer to **string** | The unique ID of the anti-fraud service used. | [optional] 
**AntiFraudServiceName** | Pointer to **string** | The name of the anti-fraud service used. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The anti-fraud service definition used. | [optional] 
**Reason** | Pointer to **string** | The reason we could not get the anti-fraud decision. | [optional] 
**Request** | Pointer to **NullableString** | The HTTP body sent to fetch a decision. | [optional] 
**Response** | Pointer to **NullableString** | The HTTP body received from the anti-fraud provider. | [optional] 
**ResponseStatusCode** | Pointer to **NullableFloat32** | The HTTP response status code from the anti-fraud provider. | [optional] 

## Methods

### NewAntiFraudTransactionStatusUpdateErrorEventContext

`func NewAntiFraudTransactionStatusUpdateErrorEventContext() *AntiFraudTransactionStatusUpdateErrorEventContext`

NewAntiFraudTransactionStatusUpdateErrorEventContext instantiates a new AntiFraudTransactionStatusUpdateErrorEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudTransactionStatusUpdateErrorEventContextWithDefaults

`func NewAntiFraudTransactionStatusUpdateErrorEventContextWithDefaults() *AntiFraudTransactionStatusUpdateErrorEventContext`

NewAntiFraudTransactionStatusUpdateErrorEventContextWithDefaults instantiates a new AntiFraudTransactionStatusUpdateErrorEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceId

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceId() string`

GetAntiFraudServiceId returns the AntiFraudServiceId field if non-nil, zero value otherwise.

### GetAntiFraudServiceIdOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceIdOk() (*string, bool)`

GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceId

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetAntiFraudServiceId(v string)`

SetAntiFraudServiceId sets AntiFraudServiceId field to given value.

### HasAntiFraudServiceId

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasAntiFraudServiceId() bool`

HasAntiFraudServiceId returns a boolean if a field has been set.

### GetAntiFraudServiceName

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceName() string`

GetAntiFraudServiceName returns the AntiFraudServiceName field if non-nil, zero value otherwise.

### GetAntiFraudServiceNameOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceNameOk() (*string, bool)`

GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceName

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetAntiFraudServiceName(v string)`

SetAntiFraudServiceName sets AntiFraudServiceName field to given value.

### HasAntiFraudServiceName

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasAntiFraudServiceName() bool`

HasAntiFraudServiceName returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.

### GetReason

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequest

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetResponse

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseStatusCode

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *AntiFraudTransactionStatusUpdateErrorEventContext) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


