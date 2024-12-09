# PaymentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-option&#x60;. | [optional] 
**Method** | Pointer to **string** | The type of payment method that is available. | [optional] 
**IconUrl** | Pointer to **NullableString** | An icon to display for the payment option. | [optional] 
**Mode** | Pointer to **string** | The mode to use with this payment option. | [optional] 
**Label** | Pointer to **string** | A label that describes this payment option. This label is returned in the language defined by the &#x60;locale&#x60; query parameter. The label can be used to display a list of payment options to the buyer in their language. | [optional] 
**CanStorePaymentMethod** | Pointer to **bool** | A flag to indicate if storing the payment method is supported. | [optional] 
**CanDelayCapture** | Pointer to **bool** | A flag to indicate if delayed capture is supported. | [optional] 
**Context** | Pointer to [**PaymentOptionContext**](PaymentOptionContext.md) |  | [optional] 

## Methods

### NewPaymentOption

`func NewPaymentOption() *PaymentOption`

NewPaymentOption instantiates a new PaymentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionWithDefaults

`func NewPaymentOptionWithDefaults() *PaymentOption`

NewPaymentOptionWithDefaults instantiates a new PaymentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentOption) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentOption) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentOption) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentOption) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetIconUrl

`func (o *PaymentOption) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *PaymentOption) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *PaymentOption) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *PaymentOption) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *PaymentOption) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *PaymentOption) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetMode

`func (o *PaymentOption) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaymentOption) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaymentOption) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PaymentOption) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetLabel

`func (o *PaymentOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentOption) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentOption) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCanStorePaymentMethod

`func (o *PaymentOption) GetCanStorePaymentMethod() bool`

GetCanStorePaymentMethod returns the CanStorePaymentMethod field if non-nil, zero value otherwise.

### GetCanStorePaymentMethodOk

`func (o *PaymentOption) GetCanStorePaymentMethodOk() (*bool, bool)`

GetCanStorePaymentMethodOk returns a tuple with the CanStorePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanStorePaymentMethod

`func (o *PaymentOption) SetCanStorePaymentMethod(v bool)`

SetCanStorePaymentMethod sets CanStorePaymentMethod field to given value.

### HasCanStorePaymentMethod

`func (o *PaymentOption) HasCanStorePaymentMethod() bool`

HasCanStorePaymentMethod returns a boolean if a field has been set.

### GetCanDelayCapture

`func (o *PaymentOption) GetCanDelayCapture() bool`

GetCanDelayCapture returns the CanDelayCapture field if non-nil, zero value otherwise.

### GetCanDelayCaptureOk

`func (o *PaymentOption) GetCanDelayCaptureOk() (*bool, bool)`

GetCanDelayCaptureOk returns a tuple with the CanDelayCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelayCapture

`func (o *PaymentOption) SetCanDelayCapture(v bool)`

SetCanDelayCapture sets CanDelayCapture field to given value.

### HasCanDelayCapture

`func (o *PaymentOption) HasCanDelayCapture() bool`

HasCanDelayCapture returns a boolean if a field has been set.

### GetContext

`func (o *PaymentOption) GetContext() PaymentOptionContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PaymentOption) GetContextOk() (*PaymentOptionContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PaymentOption) SetContext(v PaymentOptionContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PaymentOption) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


