# AntiFraudServiceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceDefinitionId** | **string** | The name of the Anti-Fraud service provider. During update request, this value is used for validation only but the underlying service can not be changed for an existing service. | 
**DisplayName** | Pointer to **string** | A unique name for this anti-fraud service which is used in the Gr4vy admin panel to give a anti-fraud Service a human readable name. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. There can only be one active service at any time. When updating a service to active, the current active service will be deactivated. | [optional] [default to true]
**ReviewsEnabled** | Pointer to **bool** | Defines if this service needs to handle the review status from anti-fraud responses with a proper review workflow. If not, the review status will be treated as any other one. | [optional] [default to false]
**Fields** | Pointer to [**[]AntiFraudServiceUpdateFieldsInner**](AntiFraudServiceUpdateFieldsInner.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this anti-fraud service e.g. for Sift &#x60;api_key&#x60; must be sent within this field when creating the service.  For updates, only the fields sent here will be updated, existing ones will not be affected if not present. | [optional] 

## Methods

### NewAntiFraudServiceUpdate

`func NewAntiFraudServiceUpdate(antiFraudServiceDefinitionId string, ) *AntiFraudServiceUpdate`

NewAntiFraudServiceUpdate instantiates a new AntiFraudServiceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudServiceUpdateWithDefaults

`func NewAntiFraudServiceUpdateWithDefaults() *AntiFraudServiceUpdate`

NewAntiFraudServiceUpdateWithDefaults instantiates a new AntiFraudServiceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudServiceUpdate) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudServiceUpdate) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudServiceUpdate) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.


### GetDisplayName

`func (o *AntiFraudServiceUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AntiFraudServiceUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AntiFraudServiceUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AntiFraudServiceUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetActive

`func (o *AntiFraudServiceUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AntiFraudServiceUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AntiFraudServiceUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AntiFraudServiceUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReviewsEnabled

`func (o *AntiFraudServiceUpdate) GetReviewsEnabled() bool`

GetReviewsEnabled returns the ReviewsEnabled field if non-nil, zero value otherwise.

### GetReviewsEnabledOk

`func (o *AntiFraudServiceUpdate) GetReviewsEnabledOk() (*bool, bool)`

GetReviewsEnabledOk returns a tuple with the ReviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewsEnabled

`func (o *AntiFraudServiceUpdate) SetReviewsEnabled(v bool)`

SetReviewsEnabled sets ReviewsEnabled field to given value.

### HasReviewsEnabled

`func (o *AntiFraudServiceUpdate) HasReviewsEnabled() bool`

HasReviewsEnabled returns a boolean if a field has been set.

### GetFields

`func (o *AntiFraudServiceUpdate) GetFields() []AntiFraudServiceUpdateFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AntiFraudServiceUpdate) GetFieldsOk() (*[]AntiFraudServiceUpdateFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AntiFraudServiceUpdate) SetFields(v []AntiFraudServiceUpdateFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AntiFraudServiceUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


