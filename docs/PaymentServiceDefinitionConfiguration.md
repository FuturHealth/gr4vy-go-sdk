# PaymentServiceDefinitionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalUiHeight** | Pointer to **string** | Height of the approval interface in either pixels or view height (vh). | [optional] 
**ApprovalUiWidth** | Pointer to **string** | Width of the approval interface in either pixels or view width (vw). | [optional] 
**ApprovalUiTarget** | Pointer to **NullableString** | The browser target that an approval URL must be opened in. If &#x60;any&#x60; or &#x60;null&#x60;, then there is no specific requirement. | [optional] 

## Methods

### NewPaymentServiceDefinitionConfiguration

`func NewPaymentServiceDefinitionConfiguration() *PaymentServiceDefinitionConfiguration`

NewPaymentServiceDefinitionConfiguration instantiates a new PaymentServiceDefinitionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceDefinitionConfigurationWithDefaults

`func NewPaymentServiceDefinitionConfigurationWithDefaults() *PaymentServiceDefinitionConfiguration`

NewPaymentServiceDefinitionConfigurationWithDefaults instantiates a new PaymentServiceDefinitionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalUiHeight

`func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiHeight() string`

GetApprovalUiHeight returns the ApprovalUiHeight field if non-nil, zero value otherwise.

### GetApprovalUiHeightOk

`func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiHeightOk() (*string, bool)`

GetApprovalUiHeightOk returns a tuple with the ApprovalUiHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUiHeight

`func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiHeight(v string)`

SetApprovalUiHeight sets ApprovalUiHeight field to given value.

### HasApprovalUiHeight

`func (o *PaymentServiceDefinitionConfiguration) HasApprovalUiHeight() bool`

HasApprovalUiHeight returns a boolean if a field has been set.

### GetApprovalUiWidth

`func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiWidth() string`

GetApprovalUiWidth returns the ApprovalUiWidth field if non-nil, zero value otherwise.

### GetApprovalUiWidthOk

`func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiWidthOk() (*string, bool)`

GetApprovalUiWidthOk returns a tuple with the ApprovalUiWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUiWidth

`func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiWidth(v string)`

SetApprovalUiWidth sets ApprovalUiWidth field to given value.

### HasApprovalUiWidth

`func (o *PaymentServiceDefinitionConfiguration) HasApprovalUiWidth() bool`

HasApprovalUiWidth returns a boolean if a field has been set.

### GetApprovalUiTarget

`func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiTarget() string`

GetApprovalUiTarget returns the ApprovalUiTarget field if non-nil, zero value otherwise.

### GetApprovalUiTargetOk

`func (o *PaymentServiceDefinitionConfiguration) GetApprovalUiTargetOk() (*string, bool)`

GetApprovalUiTargetOk returns a tuple with the ApprovalUiTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUiTarget

`func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiTarget(v string)`

SetApprovalUiTarget sets ApprovalUiTarget field to given value.

### HasApprovalUiTarget

`func (o *PaymentServiceDefinitionConfiguration) HasApprovalUiTarget() bool`

HasApprovalUiTarget returns a boolean if a field has been set.

### SetApprovalUiTargetNil

`func (o *PaymentServiceDefinitionConfiguration) SetApprovalUiTargetNil(b bool)`

 SetApprovalUiTargetNil sets the value for ApprovalUiTarget to be an explicit nil

### UnsetApprovalUiTarget
`func (o *PaymentServiceDefinitionConfiguration) UnsetApprovalUiTarget()`

UnsetApprovalUiTarget ensures that no value is present for ApprovalUiTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


