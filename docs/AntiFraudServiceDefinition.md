# AntiFraudServiceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the anti fraud service definition. | [optional] 
**Type** | Pointer to **string** | &#x60;anti-fraud-service-definition&#x60;. | [optional] [default to "anti-fraud-service-definition"]
**DisplayName** | Pointer to **string** | The display name of this service. | [optional] 
**Fields** | Pointer to [**[]AntiFraudServiceDefinitionFieldsInner**](AntiFraudServiceDefinitionFieldsInner.md) | A list of fields that need to be submitted when activating the payment. service. | [optional] 
**IconUrl** | Pointer to **string** | An icon to display for the payment service. | [optional] 

## Methods

### NewAntiFraudServiceDefinition

`func NewAntiFraudServiceDefinition() *AntiFraudServiceDefinition`

NewAntiFraudServiceDefinition instantiates a new AntiFraudServiceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudServiceDefinitionWithDefaults

`func NewAntiFraudServiceDefinitionWithDefaults() *AntiFraudServiceDefinition`

NewAntiFraudServiceDefinitionWithDefaults instantiates a new AntiFraudServiceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AntiFraudServiceDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntiFraudServiceDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntiFraudServiceDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AntiFraudServiceDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AntiFraudServiceDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AntiFraudServiceDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AntiFraudServiceDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AntiFraudServiceDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *AntiFraudServiceDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AntiFraudServiceDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AntiFraudServiceDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AntiFraudServiceDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *AntiFraudServiceDefinition) GetFields() []AntiFraudServiceDefinitionFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AntiFraudServiceDefinition) GetFieldsOk() (*[]AntiFraudServiceDefinitionFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AntiFraudServiceDefinition) SetFields(v []AntiFraudServiceDefinitionFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AntiFraudServiceDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIconUrl

`func (o *AntiFraudServiceDefinition) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AntiFraudServiceDefinition) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AntiFraudServiceDefinition) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AntiFraudServiceDefinition) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


