# AntiFraudServiceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiFraudServiceDefinitionId** | **string** | The name of the Anti-Fraud service provider. During update request, this value is used for validation only but the underlying service can not be changed for an existing service. | 
**DisplayName** | **string** | A unique name for this anti-fraud service which is used in the Gr4vy admin panel to give a anti-fraud Service a human readable name. | 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. There can only be one active service at any time. When updating a service to active, the current active service will be deactivated. | [optional] [default to true]
**ReviewsEnabled** | Pointer to **bool** | Defines if this service needs to handle the review status from anti-fraud responses with a proper review workflow. If not, the review status will be treated as any other one. | [optional] [default to false]
**Fields** | [**[]AntiFraudServiceUpdateFieldsInner**](AntiFraudServiceUpdateFieldsInner.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this anti-fraud service e.g. for Sift &#x60;api_key&#x60; must be sent within this field when creating the service.  For updates, only the fields sent here will be updated, existing ones will not be affected if not present. | 

## Methods

### NewAntiFraudServiceCreate

`func NewAntiFraudServiceCreate(antiFraudServiceDefinitionId string, displayName string, fields []AntiFraudServiceUpdateFieldsInner, ) *AntiFraudServiceCreate`

NewAntiFraudServiceCreate instantiates a new AntiFraudServiceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudServiceCreateWithDefaults

`func NewAntiFraudServiceCreateWithDefaults() *AntiFraudServiceCreate`

NewAntiFraudServiceCreateWithDefaults instantiates a new AntiFraudServiceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudServiceCreate) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudServiceCreate) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudServiceCreate) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.


### GetDisplayName

`func (o *AntiFraudServiceCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AntiFraudServiceCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AntiFraudServiceCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetActive

`func (o *AntiFraudServiceCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AntiFraudServiceCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AntiFraudServiceCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AntiFraudServiceCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReviewsEnabled

`func (o *AntiFraudServiceCreate) GetReviewsEnabled() bool`

GetReviewsEnabled returns the ReviewsEnabled field if non-nil, zero value otherwise.

### GetReviewsEnabledOk

`func (o *AntiFraudServiceCreate) GetReviewsEnabledOk() (*bool, bool)`

GetReviewsEnabledOk returns a tuple with the ReviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewsEnabled

`func (o *AntiFraudServiceCreate) SetReviewsEnabled(v bool)`

SetReviewsEnabled sets ReviewsEnabled field to given value.

### HasReviewsEnabled

`func (o *AntiFraudServiceCreate) HasReviewsEnabled() bool`

HasReviewsEnabled returns a boolean if a field has been set.

### GetFields

`func (o *AntiFraudServiceCreate) GetFields() []AntiFraudServiceUpdateFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AntiFraudServiceCreate) GetFieldsOk() (*[]AntiFraudServiceUpdateFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AntiFraudServiceCreate) SetFields(v []AntiFraudServiceUpdateFieldsInner)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


