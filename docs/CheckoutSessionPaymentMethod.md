# CheckoutSessionPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **NullableString** | Unique ID for the payment method. | [optional] 
**Method** | Pointer to **string** | Payment method type. | [optional] 
**Scheme** | Pointer to **NullableString** | The scheme/brand of the card. | [optional] 
**Label** | Pointer to **NullableString** | Last four digits of PAN. | [optional] 
**Details** | Pointer to [**NullableCheckoutSessionPaymentMethodDetails**](CheckoutSessionPaymentMethodDetails.md) |  | [optional] 
**Fingerprint** | Pointer to **NullableString** | The unique hash derived from the payment method identifier (e.g. card number). | [optional] 

## Methods

### NewCheckoutSessionPaymentMethod

`func NewCheckoutSessionPaymentMethod() *CheckoutSessionPaymentMethod`

NewCheckoutSessionPaymentMethod instantiates a new CheckoutSessionPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionPaymentMethodWithDefaults

`func NewCheckoutSessionPaymentMethodWithDefaults() *CheckoutSessionPaymentMethod`

NewCheckoutSessionPaymentMethodWithDefaults instantiates a new CheckoutSessionPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutSessionPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutSessionPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutSessionPaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckoutSessionPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CheckoutSessionPaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSessionPaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSessionPaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckoutSessionPaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CheckoutSessionPaymentMethod) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CheckoutSessionPaymentMethod) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMethod

`func (o *CheckoutSessionPaymentMethod) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckoutSessionPaymentMethod) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckoutSessionPaymentMethod) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CheckoutSessionPaymentMethod) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetScheme

`func (o *CheckoutSessionPaymentMethod) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *CheckoutSessionPaymentMethod) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *CheckoutSessionPaymentMethod) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *CheckoutSessionPaymentMethod) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *CheckoutSessionPaymentMethod) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *CheckoutSessionPaymentMethod) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetLabel

`func (o *CheckoutSessionPaymentMethod) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CheckoutSessionPaymentMethod) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CheckoutSessionPaymentMethod) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CheckoutSessionPaymentMethod) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CheckoutSessionPaymentMethod) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CheckoutSessionPaymentMethod) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDetails

`func (o *CheckoutSessionPaymentMethod) GetDetails() CheckoutSessionPaymentMethodDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CheckoutSessionPaymentMethod) GetDetailsOk() (*CheckoutSessionPaymentMethodDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CheckoutSessionPaymentMethod) SetDetails(v CheckoutSessionPaymentMethodDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CheckoutSessionPaymentMethod) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CheckoutSessionPaymentMethod) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CheckoutSessionPaymentMethod) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetFingerprint

`func (o *CheckoutSessionPaymentMethod) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CheckoutSessionPaymentMethod) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CheckoutSessionPaymentMethod) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CheckoutSessionPaymentMethod) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *CheckoutSessionPaymentMethod) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *CheckoutSessionPaymentMethod) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


