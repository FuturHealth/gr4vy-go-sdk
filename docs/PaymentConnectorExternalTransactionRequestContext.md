# PaymentConnectorExternalTransactionRequestContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorRequestId** | Pointer to **string** | The connector request id. | [optional] 
**Request** | Pointer to **NullableString** | The request. | [optional] 
**Response** | Pointer to **NullableString** | The response. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **NullableString** | The payment service definition id. | [optional] 
**Success** | Pointer to **NullableBool** | Whether the transaction was successful. | [optional] 
**PaymentServiceId** | Pointer to **NullableString** | The payment service id. | [optional] 
**ResponseStatusCode** | Pointer to **NullableString** | The response status code. | [optional] 
**PaymentServiceDisplayName** | Pointer to **NullableString** | The payment service display name. | [optional] 
**Method** | Pointer to **NullableString** | The HTTP method. | [optional] 
**Url** | Pointer to **NullableString** | The endpoint for the request. | [optional] 
**RefundId** | Pointer to **NullableString** | The ID of the refund, in case this links to a refund. | [optional] 
**RefundXid** | Pointer to **NullableString** | The external ID of the refund. | [optional] 

## Methods

### NewPaymentConnectorExternalTransactionRequestContext

`func NewPaymentConnectorExternalTransactionRequestContext() *PaymentConnectorExternalTransactionRequestContext`

NewPaymentConnectorExternalTransactionRequestContext instantiates a new PaymentConnectorExternalTransactionRequestContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConnectorExternalTransactionRequestContextWithDefaults

`func NewPaymentConnectorExternalTransactionRequestContextWithDefaults() *PaymentConnectorExternalTransactionRequestContext`

NewPaymentConnectorExternalTransactionRequestContextWithDefaults instantiates a new PaymentConnectorExternalTransactionRequestContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorRequestId

`func (o *PaymentConnectorExternalTransactionRequestContext) GetConnectorRequestId() string`

GetConnectorRequestId returns the ConnectorRequestId field if non-nil, zero value otherwise.

### GetConnectorRequestIdOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetConnectorRequestIdOk() (*string, bool)`

GetConnectorRequestIdOk returns a tuple with the ConnectorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorRequestId

`func (o *PaymentConnectorExternalTransactionRequestContext) SetConnectorRequestId(v string)`

SetConnectorRequestId sets ConnectorRequestId field to given value.

### HasConnectorRequestId

`func (o *PaymentConnectorExternalTransactionRequestContext) HasConnectorRequestId() bool`

HasConnectorRequestId returns a boolean if a field has been set.

### GetRequest

`func (o *PaymentConnectorExternalTransactionRequestContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PaymentConnectorExternalTransactionRequestContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *PaymentConnectorExternalTransactionRequestContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetResponse

`func (o *PaymentConnectorExternalTransactionRequestContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PaymentConnectorExternalTransactionRequestContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PaymentConnectorExternalTransactionRequestContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetPaymentServiceDefinitionId

`func (o *PaymentConnectorExternalTransactionRequestContext) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentConnectorExternalTransactionRequestContext) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentConnectorExternalTransactionRequestContext) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### SetPaymentServiceDefinitionIdNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetPaymentServiceDefinitionIdNil(b bool)`

 SetPaymentServiceDefinitionIdNil sets the value for PaymentServiceDefinitionId to be an explicit nil

### UnsetPaymentServiceDefinitionId
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetPaymentServiceDefinitionId()`

UnsetPaymentServiceDefinitionId ensures that no value is present for PaymentServiceDefinitionId, not even an explicit nil
### GetSuccess

`func (o *PaymentConnectorExternalTransactionRequestContext) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PaymentConnectorExternalTransactionRequestContext) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PaymentConnectorExternalTransactionRequestContext) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### SetSuccessNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetSuccessNil(b bool)`

 SetSuccessNil sets the value for Success to be an explicit nil

### UnsetSuccess
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetSuccess()`

UnsetSuccess ensures that no value is present for Success, not even an explicit nil
### GetPaymentServiceId

`func (o *PaymentConnectorExternalTransactionRequestContext) GetPaymentServiceId() string`

GetPaymentServiceId returns the PaymentServiceId field if non-nil, zero value otherwise.

### GetPaymentServiceIdOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetPaymentServiceIdOk() (*string, bool)`

GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceId

`func (o *PaymentConnectorExternalTransactionRequestContext) SetPaymentServiceId(v string)`

SetPaymentServiceId sets PaymentServiceId field to given value.

### HasPaymentServiceId

`func (o *PaymentConnectorExternalTransactionRequestContext) HasPaymentServiceId() bool`

HasPaymentServiceId returns a boolean if a field has been set.

### SetPaymentServiceIdNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetPaymentServiceIdNil(b bool)`

 SetPaymentServiceIdNil sets the value for PaymentServiceId to be an explicit nil

### UnsetPaymentServiceId
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetPaymentServiceId()`

UnsetPaymentServiceId ensures that no value is present for PaymentServiceId, not even an explicit nil
### GetResponseStatusCode

`func (o *PaymentConnectorExternalTransactionRequestContext) GetResponseStatusCode() string`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetResponseStatusCodeOk() (*string, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *PaymentConnectorExternalTransactionRequestContext) SetResponseStatusCode(v string)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *PaymentConnectorExternalTransactionRequestContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetPaymentServiceDisplayName

`func (o *PaymentConnectorExternalTransactionRequestContext) GetPaymentServiceDisplayName() string`

GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field if non-nil, zero value otherwise.

### GetPaymentServiceDisplayNameOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetPaymentServiceDisplayNameOk() (*string, bool)`

GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDisplayName

`func (o *PaymentConnectorExternalTransactionRequestContext) SetPaymentServiceDisplayName(v string)`

SetPaymentServiceDisplayName sets PaymentServiceDisplayName field to given value.

### HasPaymentServiceDisplayName

`func (o *PaymentConnectorExternalTransactionRequestContext) HasPaymentServiceDisplayName() bool`

HasPaymentServiceDisplayName returns a boolean if a field has been set.

### SetPaymentServiceDisplayNameNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetPaymentServiceDisplayNameNil(b bool)`

 SetPaymentServiceDisplayNameNil sets the value for PaymentServiceDisplayName to be an explicit nil

### UnsetPaymentServiceDisplayName
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetPaymentServiceDisplayName()`

UnsetPaymentServiceDisplayName ensures that no value is present for PaymentServiceDisplayName, not even an explicit nil
### GetMethod

`func (o *PaymentConnectorExternalTransactionRequestContext) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentConnectorExternalTransactionRequestContext) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentConnectorExternalTransactionRequestContext) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetUrl

`func (o *PaymentConnectorExternalTransactionRequestContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentConnectorExternalTransactionRequestContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PaymentConnectorExternalTransactionRequestContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRefundId

`func (o *PaymentConnectorExternalTransactionRequestContext) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *PaymentConnectorExternalTransactionRequestContext) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *PaymentConnectorExternalTransactionRequestContext) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### SetRefundIdNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetRefundIdNil(b bool)`

 SetRefundIdNil sets the value for RefundId to be an explicit nil

### UnsetRefundId
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetRefundId()`

UnsetRefundId ensures that no value is present for RefundId, not even an explicit nil
### GetRefundXid

`func (o *PaymentConnectorExternalTransactionRequestContext) GetRefundXid() string`

GetRefundXid returns the RefundXid field if non-nil, zero value otherwise.

### GetRefundXidOk

`func (o *PaymentConnectorExternalTransactionRequestContext) GetRefundXidOk() (*string, bool)`

GetRefundXidOk returns a tuple with the RefundXid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundXid

`func (o *PaymentConnectorExternalTransactionRequestContext) SetRefundXid(v string)`

SetRefundXid sets RefundXid field to given value.

### HasRefundXid

`func (o *PaymentConnectorExternalTransactionRequestContext) HasRefundXid() bool`

HasRefundXid returns a boolean if a field has been set.

### SetRefundXidNil

`func (o *PaymentConnectorExternalTransactionRequestContext) SetRefundXidNil(b bool)`

 SetRefundXidNil sets the value for RefundXid to be an explicit nil

### UnsetRefundXid
`func (o *PaymentConnectorExternalTransactionRequestContext) UnsetRefundXid()`

UnsetRefundXid ensures that no value is present for RefundXid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


