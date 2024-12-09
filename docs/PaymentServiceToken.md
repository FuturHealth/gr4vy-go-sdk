# PaymentServiceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **string** | The unique ID of the token. | [optional] 
**PaymentMethodId** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**PaymentServiceId** | Pointer to **string** | The unique ID of the payment service. | [optional] 
**Status** | Pointer to **string** | The state of the token.  - &#x60;processing&#x60; - The payment method is still being stored. - &#x60;buyer_approval_required&#x60; - Storing the payment method requires   the buyer to provide approval. Follow the &#x60;approval_url&#x60; for next steps. - &#x60;succeeded&#x60; - The payment method is approved and stored with all   relevant payment services. - &#x60;failed&#x60; - Storing the payment method did not succeed. | [optional] 
**ApprovalUrl** | Pointer to **NullableString** | The optional URL that the buyer needs to be redirected to to further authorize their payment. | [optional] 
**Token** | Pointer to **NullableString** | The token value. Will be present if succeeded. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this token was first created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this token was last updated in our system. | [optional] 

## Methods

### NewPaymentServiceToken

`func NewPaymentServiceToken() *PaymentServiceToken`

NewPaymentServiceToken instantiates a new PaymentServiceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceTokenWithDefaults

`func NewPaymentServiceTokenWithDefaults() *PaymentServiceToken`

NewPaymentServiceTokenWithDefaults instantiates a new PaymentServiceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentServiceToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentServiceToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentServiceToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentServiceToken) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentServiceToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *PaymentServiceToken) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentServiceToken) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentServiceToken) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PaymentServiceToken) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentServiceId

`func (o *PaymentServiceToken) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentServiceToken) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentServiceToken) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentServiceToken) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentServiceToken) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentServiceToken) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentServiceToken) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentServiceToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovalUrl

`func (o *PaymentServiceToken) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *PaymentServiceToken) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *PaymentServiceToken) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *PaymentServiceToken) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *PaymentServiceToken) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *PaymentServiceToken) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetToken

`func (o *PaymentServiceToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentServiceToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentServiceToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentServiceToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *PaymentServiceToken) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *PaymentServiceToken) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetCreatedAt

`func (o *PaymentServiceToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentServiceToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentServiceToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentServiceToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentServiceToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentServiceToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentServiceToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentServiceToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


