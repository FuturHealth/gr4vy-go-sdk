# PaymentOptionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | Gateway used for Google Pay payments. | [optional] 
**GatewayMerchantId** | Pointer to **string** | Gateway merchant identifier used for Google Pay payments. | [optional] 
**MerchantName** | Pointer to **string** | Display name of the merchant as registered with the digital wallet provider. | [optional] 
**SupportedSchemes** | Pointer to **[]string** | Card schemes supported by the digital wallet provider. | [optional] 
**ApprovalUi** | Pointer to [**PaymentOptionApprovalUI**](PaymentOptionApprovalUI.md) |  | [optional] 
**RequiredFields** | Pointer to [**RequiredFields**](RequiredFields.md) |  | [optional] 

## Methods

### NewPaymentOptionContext

`func NewPaymentOptionContext() *PaymentOptionContext`

NewPaymentOptionContext instantiates a new PaymentOptionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionContextWithDefaults

`func NewPaymentOptionContextWithDefaults() *PaymentOptionContext`

NewPaymentOptionContextWithDefaults instantiates a new PaymentOptionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *PaymentOptionContext) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PaymentOptionContext) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PaymentOptionContext) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PaymentOptionContext) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayMerchantId

`func (o *PaymentOptionContext) GetGatewayMerchantId() string`

GetGatewayMerchantId returns the GatewayMerchantId field if non-nil, zero value otherwise.

### GetGatewayMerchantIdOk

`func (o *PaymentOptionContext) GetGatewayMerchantIdOk() (*string, bool)`

GetGatewayMerchantIdOk returns a tuple with the GatewayMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMerchantId

`func (o *PaymentOptionContext) SetGatewayMerchantId(v string)`

SetGatewayMerchantId sets GatewayMerchantId field to given value.

### HasGatewayMerchantId

`func (o *PaymentOptionContext) HasGatewayMerchantId() bool`

HasGatewayMerchantId returns a boolean if a field has been set.

### GetMerchantName

`func (o *PaymentOptionContext) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *PaymentOptionContext) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *PaymentOptionContext) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *PaymentOptionContext) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetSupportedSchemes

`func (o *PaymentOptionContext) GetSupportedSchemes() []string`

GetSupportedSchemes returns the SupportedSchemes field if non-nil, zero value otherwise.

### GetSupportedSchemesOk

`func (o *PaymentOptionContext) GetSupportedSchemesOk() (*[]string, bool)`

GetSupportedSchemesOk returns a tuple with the SupportedSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSchemes

`func (o *PaymentOptionContext) SetSupportedSchemes(v []string)`

SetSupportedSchemes sets SupportedSchemes field to given value.

### HasSupportedSchemes

`func (o *PaymentOptionContext) HasSupportedSchemes() bool`

HasSupportedSchemes returns a boolean if a field has been set.

### GetApprovalUi

`func (o *PaymentOptionContext) GetApprovalUi() PaymentOptionApprovalUI`

GetApprovalUi returns the ApprovalUi field if non-nil, zero value otherwise.

### GetApprovalUiOk

`func (o *PaymentOptionContext) GetApprovalUiOk() (*PaymentOptionApprovalUI, bool)`

GetApprovalUiOk returns a tuple with the ApprovalUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUi

`func (o *PaymentOptionContext) SetApprovalUi(v PaymentOptionApprovalUI)`

SetApprovalUi sets ApprovalUi field to given value.

### HasApprovalUi

`func (o *PaymentOptionContext) HasApprovalUi() bool`

HasApprovalUi returns a boolean if a field has been set.

### GetRequiredFields

`func (o *PaymentOptionContext) GetRequiredFields() RequiredFields`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *PaymentOptionContext) GetRequiredFieldsOk() (*RequiredFields, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *PaymentOptionContext) SetRequiredFields(v RequiredFields)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *PaymentOptionContext) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


