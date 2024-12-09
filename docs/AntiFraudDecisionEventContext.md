# AntiFraudDecisionEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceId** | Pointer to **string** | The unique ID of the anti-fraud service used. | [optional] 
**AntiFraudServiceName** | Pointer to **string** | The name of the anti-fraud service used. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The anti-fraud service definition used. | [optional] 
**AntiFraudServiceCheckId** | Pointer to **string** | The external ID of the decision. | [optional] 
**Request** | Pointer to **string** | The HTTP body sent to fetch a decision. | [optional] 
**Response** | Pointer to **string** | The HTTP body received from the anti-fraud provider. | [optional] 
**ResponseStatusCode** | Pointer to **float32** | The HTTP response status code from the anti-fraud provider. | [optional] 
**Decision** | Pointer to **string** | The parsed decision response from the anti-fraud provider response. | [optional] 

## Methods

### NewAntiFraudDecisionEventContext

`func NewAntiFraudDecisionEventContext() *AntiFraudDecisionEventContext`

NewAntiFraudDecisionEventContext instantiates a new AntiFraudDecisionEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudDecisionEventContextWithDefaults

`func NewAntiFraudDecisionEventContextWithDefaults() *AntiFraudDecisionEventContext`

NewAntiFraudDecisionEventContextWithDefaults instantiates a new AntiFraudDecisionEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceId

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceId() string`

GetAntiFraudServiceId returns the AntiFraudServiceId field if non-nil, zero value otherwise.

### GetAntiFraudServiceIdOk

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceIdOk() (*string, bool)`

GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceId

`func (o *AntiFraudDecisionEventContext) SetAntiFraudServiceId(v string)`

SetAntiFraudServiceId sets AntiFraudServiceId field to given value.

### HasAntiFraudServiceId

`func (o *AntiFraudDecisionEventContext) HasAntiFraudServiceId() bool`

HasAntiFraudServiceId returns a boolean if a field has been set.

### GetAntiFraudServiceName

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceName() string`

GetAntiFraudServiceName returns the AntiFraudServiceName field if non-nil, zero value otherwise.

### GetAntiFraudServiceNameOk

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceNameOk() (*string, bool)`

GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceName

`func (o *AntiFraudDecisionEventContext) SetAntiFraudServiceName(v string)`

SetAntiFraudServiceName sets AntiFraudServiceName field to given value.

### HasAntiFraudServiceName

`func (o *AntiFraudDecisionEventContext) HasAntiFraudServiceName() bool`

HasAntiFraudServiceName returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionEventContext) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionEventContext) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.

### GetAntiFraudServiceCheckId

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceCheckId() string`

GetAntiFraudServiceCheckId returns the AntiFraudServiceCheckId field if non-nil, zero value otherwise.

### GetAntiFraudServiceCheckIdOk

`func (o *AntiFraudDecisionEventContext) GetAntiFraudServiceCheckIdOk() (*string, bool)`

GetAntiFraudServiceCheckIdOk returns a tuple with the AntiFraudServiceCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceCheckId

`func (o *AntiFraudDecisionEventContext) SetAntiFraudServiceCheckId(v string)`

SetAntiFraudServiceCheckId sets AntiFraudServiceCheckId field to given value.

### HasAntiFraudServiceCheckId

`func (o *AntiFraudDecisionEventContext) HasAntiFraudServiceCheckId() bool`

HasAntiFraudServiceCheckId returns a boolean if a field has been set.

### GetRequest

`func (o *AntiFraudDecisionEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AntiFraudDecisionEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AntiFraudDecisionEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AntiFraudDecisionEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *AntiFraudDecisionEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AntiFraudDecisionEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AntiFraudDecisionEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AntiFraudDecisionEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *AntiFraudDecisionEventContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *AntiFraudDecisionEventContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *AntiFraudDecisionEventContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *AntiFraudDecisionEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetDecision

`func (o *AntiFraudDecisionEventContext) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AntiFraudDecisionEventContext) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AntiFraudDecisionEventContext) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AntiFraudDecisionEventContext) HasDecision() bool`

HasDecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


