# AntiFraudTransactionStatusUpdateEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceId** | Pointer to **string** | The unique ID of the anti-fraud service used. | [optional] 
**AntiFraudServiceName** | Pointer to **string** | The name of the anti-fraud service used. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The anti-fraud service definition used. | [optional] 
**Request** | Pointer to **string** | The HTTP body sent to fetch a decision. | [optional] 
**Response** | Pointer to **string** | The HTTP body received from the anti-fraud provider. | [optional] 
**ResponseStatusCode** | Pointer to **float32** | The HTTP response status code from the anti-fraud provider. | [optional] 

## Methods

### NewAntiFraudTransactionStatusUpdateEventContext

`func NewAntiFraudTransactionStatusUpdateEventContext() *AntiFraudTransactionStatusUpdateEventContext`

NewAntiFraudTransactionStatusUpdateEventContext instantiates a new AntiFraudTransactionStatusUpdateEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudTransactionStatusUpdateEventContextWithDefaults

`func NewAntiFraudTransactionStatusUpdateEventContextWithDefaults() *AntiFraudTransactionStatusUpdateEventContext`

NewAntiFraudTransactionStatusUpdateEventContextWithDefaults instantiates a new AntiFraudTransactionStatusUpdateEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceId

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceId() string`

GetAntiFraudServiceId returns the AntiFraudServiceId field if non-nil, zero value otherwise.

### GetAntiFraudServiceIdOk

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceIdOk() (*string, bool)`

GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceId

`func (o *AntiFraudTransactionStatusUpdateEventContext) SetAntiFraudServiceId(v string)`

SetAntiFraudServiceId sets AntiFraudServiceId field to given value.

### HasAntiFraudServiceId

`func (o *AntiFraudTransactionStatusUpdateEventContext) HasAntiFraudServiceId() bool`

HasAntiFraudServiceId returns a boolean if a field has been set.

### GetAntiFraudServiceName

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceName() string`

GetAntiFraudServiceName returns the AntiFraudServiceName field if non-nil, zero value otherwise.

### GetAntiFraudServiceNameOk

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceNameOk() (*string, bool)`

GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceName

`func (o *AntiFraudTransactionStatusUpdateEventContext) SetAntiFraudServiceName(v string)`

SetAntiFraudServiceName sets AntiFraudServiceName field to given value.

### HasAntiFraudServiceName

`func (o *AntiFraudTransactionStatusUpdateEventContext) HasAntiFraudServiceName() bool`

HasAntiFraudServiceName returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudTransactionStatusUpdateEventContext) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudTransactionStatusUpdateEventContext) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.

### GetRequest

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AntiFraudTransactionStatusUpdateEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AntiFraudTransactionStatusUpdateEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AntiFraudTransactionStatusUpdateEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AntiFraudTransactionStatusUpdateEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *AntiFraudTransactionStatusUpdateEventContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *AntiFraudTransactionStatusUpdateEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


