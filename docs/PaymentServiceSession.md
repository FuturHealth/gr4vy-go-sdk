# PaymentServiceSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Status** | Pointer to **string** | The status of the response.  - &#x60;succeeded&#x60; - The session was successfully generated. - &#x60;failed&#x60; - The session could not be generated. | [optional] 
**Code** | Pointer to **NullableString** | A generic error code that may be returned when the session could not be generated. | [optional] 
**StatusCode** | Pointer to **NullableFloat32** | The HTTP status code received from the payment service. | [optional] 
**ResponseBody** | Pointer to **map[string]interface{}** | The parsed JSON received from the payment service. | [optional] 

## Methods

### NewPaymentServiceSession

`func NewPaymentServiceSession() *PaymentServiceSession`

NewPaymentServiceSession instantiates a new PaymentServiceSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceSessionWithDefaults

`func NewPaymentServiceSessionWithDefaults() *PaymentServiceSession`

NewPaymentServiceSessionWithDefaults instantiates a new PaymentServiceSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentServiceSession) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentServiceSession) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentServiceSession) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentServiceSession) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentServiceSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentServiceSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentServiceSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentServiceSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *PaymentServiceSession) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentServiceSession) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentServiceSession) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentServiceSession) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PaymentServiceSession) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PaymentServiceSession) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetStatusCode

`func (o *PaymentServiceSession) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *PaymentServiceSession) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *PaymentServiceSession) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *PaymentServiceSession) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### SetStatusCodeNil

`func (o *PaymentServiceSession) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *PaymentServiceSession) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
### GetResponseBody

`func (o *PaymentServiceSession) GetResponseBody() map[string]interface{}`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *PaymentServiceSession) GetResponseBodyOk() (*map[string]interface{}, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *PaymentServiceSession) SetResponseBody(v map[string]interface{})`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *PaymentServiceSession) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *PaymentServiceSession) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *PaymentServiceSession) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


