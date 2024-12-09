# AntiFraudDecisionUpdateErrorEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceId** | Pointer to **string** | The unique ID of the anti-fraud service used. | [optional] 
**AntiFraudServiceName** | Pointer to **string** | The name of the anti-fraud service used. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The anti-fraud service definition used. | [optional] 
**AntiFraudServiceCheckId** | Pointer to **string** | The external ID of the decision. | [optional] 
**Decision** | Pointer to **string** | The updated decision sent to the anti-fraud provider. | [optional] 
**Reason** | Pointer to **string** | The reason we could not update the anti-fraud decision. | [optional] 

## Methods

### NewAntiFraudDecisionUpdateErrorEventContext

`func NewAntiFraudDecisionUpdateErrorEventContext() *AntiFraudDecisionUpdateErrorEventContext`

NewAntiFraudDecisionUpdateErrorEventContext instantiates a new AntiFraudDecisionUpdateErrorEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudDecisionUpdateErrorEventContextWithDefaults

`func NewAntiFraudDecisionUpdateErrorEventContextWithDefaults() *AntiFraudDecisionUpdateErrorEventContext`

NewAntiFraudDecisionUpdateErrorEventContextWithDefaults instantiates a new AntiFraudDecisionUpdateErrorEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceId

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceId() string`

GetAntiFraudServiceId returns the AntiFraudServiceId field if non-nil, zero value otherwise.

### GetAntiFraudServiceIdOk

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceIdOk() (*string, bool)`

GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceId

`func (o *AntiFraudDecisionUpdateErrorEventContext) SetAntiFraudServiceId(v string)`

SetAntiFraudServiceId sets AntiFraudServiceId field to given value.

### HasAntiFraudServiceId

`func (o *AntiFraudDecisionUpdateErrorEventContext) HasAntiFraudServiceId() bool`

HasAntiFraudServiceId returns a boolean if a field has been set.

### GetAntiFraudServiceName

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceName() string`

GetAntiFraudServiceName returns the AntiFraudServiceName field if non-nil, zero value otherwise.

### GetAntiFraudServiceNameOk

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceNameOk() (*string, bool)`

GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceName

`func (o *AntiFraudDecisionUpdateErrorEventContext) SetAntiFraudServiceName(v string)`

SetAntiFraudServiceName sets AntiFraudServiceName field to given value.

### HasAntiFraudServiceName

`func (o *AntiFraudDecisionUpdateErrorEventContext) HasAntiFraudServiceName() bool`

HasAntiFraudServiceName returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionUpdateErrorEventContext) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionUpdateErrorEventContext) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.

### GetAntiFraudServiceCheckId

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceCheckId() string`

GetAntiFraudServiceCheckId returns the AntiFraudServiceCheckId field if non-nil, zero value otherwise.

### GetAntiFraudServiceCheckIdOk

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetAntiFraudServiceCheckIdOk() (*string, bool)`

GetAntiFraudServiceCheckIdOk returns a tuple with the AntiFraudServiceCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceCheckId

`func (o *AntiFraudDecisionUpdateErrorEventContext) SetAntiFraudServiceCheckId(v string)`

SetAntiFraudServiceCheckId sets AntiFraudServiceCheckId field to given value.

### HasAntiFraudServiceCheckId

`func (o *AntiFraudDecisionUpdateErrorEventContext) HasAntiFraudServiceCheckId() bool`

HasAntiFraudServiceCheckId returns a boolean if a field has been set.

### GetDecision

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AntiFraudDecisionUpdateErrorEventContext) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AntiFraudDecisionUpdateErrorEventContext) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetReason

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AntiFraudDecisionUpdateErrorEventContext) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AntiFraudDecisionUpdateErrorEventContext) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AntiFraudDecisionUpdateErrorEventContext) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


