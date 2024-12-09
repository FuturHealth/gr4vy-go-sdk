# AntiFraudWebhookEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceId** | Pointer to **string** | The unique ID of the anti-fraud service used. | [optional] 
**AntiFraudServiceName** | Pointer to **string** | The name of the anti-fraud service used. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The anti-fraud service definition used. | [optional] 
**AntiFraudServiceCheckId** | Pointer to **string** | The external ID of the decision that&#39;s being updated. | [optional] 
**Content** | Pointer to **string** | The raw payload sent as a webhook. | [optional] 
**ContentType** | Pointer to **string** | The content type of the payload sent as a webhook. | [optional] 
**Decision** | Pointer to **string** | The parsed decision response from the anti-fraud provider webhook. | [optional] 
**Comment** | Pointer to **string** | Any comment that may have come with the webhook event. | [optional] 

## Methods

### NewAntiFraudWebhookEventContext

`func NewAntiFraudWebhookEventContext() *AntiFraudWebhookEventContext`

NewAntiFraudWebhookEventContext instantiates a new AntiFraudWebhookEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudWebhookEventContextWithDefaults

`func NewAntiFraudWebhookEventContextWithDefaults() *AntiFraudWebhookEventContext`

NewAntiFraudWebhookEventContextWithDefaults instantiates a new AntiFraudWebhookEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceId

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceId() string`

GetAntiFraudServiceId returns the AntiFraudServiceId field if non-nil, zero value otherwise.

### GetAntiFraudServiceIdOk

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceIdOk() (*string, bool)`

GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceId

`func (o *AntiFraudWebhookEventContext) SetAntiFraudServiceId(v string)`

SetAntiFraudServiceId sets AntiFraudServiceId field to given value.

### HasAntiFraudServiceId

`func (o *AntiFraudWebhookEventContext) HasAntiFraudServiceId() bool`

HasAntiFraudServiceId returns a boolean if a field has been set.

### GetAntiFraudServiceName

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceName() string`

GetAntiFraudServiceName returns the AntiFraudServiceName field if non-nil, zero value otherwise.

### GetAntiFraudServiceNameOk

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceNameOk() (*string, bool)`

GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceName

`func (o *AntiFraudWebhookEventContext) SetAntiFraudServiceName(v string)`

SetAntiFraudServiceName sets AntiFraudServiceName field to given value.

### HasAntiFraudServiceName

`func (o *AntiFraudWebhookEventContext) HasAntiFraudServiceName() bool`

HasAntiFraudServiceName returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudWebhookEventContext) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudWebhookEventContext) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.

### GetAntiFraudServiceCheckId

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceCheckId() string`

GetAntiFraudServiceCheckId returns the AntiFraudServiceCheckId field if non-nil, zero value otherwise.

### GetAntiFraudServiceCheckIdOk

`func (o *AntiFraudWebhookEventContext) GetAntiFraudServiceCheckIdOk() (*string, bool)`

GetAntiFraudServiceCheckIdOk returns a tuple with the AntiFraudServiceCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceCheckId

`func (o *AntiFraudWebhookEventContext) SetAntiFraudServiceCheckId(v string)`

SetAntiFraudServiceCheckId sets AntiFraudServiceCheckId field to given value.

### HasAntiFraudServiceCheckId

`func (o *AntiFraudWebhookEventContext) HasAntiFraudServiceCheckId() bool`

HasAntiFraudServiceCheckId returns a boolean if a field has been set.

### GetContent

`func (o *AntiFraudWebhookEventContext) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AntiFraudWebhookEventContext) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AntiFraudWebhookEventContext) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AntiFraudWebhookEventContext) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *AntiFraudWebhookEventContext) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AntiFraudWebhookEventContext) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AntiFraudWebhookEventContext) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AntiFraudWebhookEventContext) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDecision

`func (o *AntiFraudWebhookEventContext) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AntiFraudWebhookEventContext) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AntiFraudWebhookEventContext) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AntiFraudWebhookEventContext) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetComment

`func (o *AntiFraudWebhookEventContext) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AntiFraudWebhookEventContext) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AntiFraudWebhookEventContext) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AntiFraudWebhookEventContext) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


