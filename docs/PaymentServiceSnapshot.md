# PaymentServiceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **string** | The ID of this payment service. | [optional] 
**DisplayName** | Pointer to **string** | The custom name set for this service. | [optional] 
**Method** | Pointer to **string** | The payment method that this services handles. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The ID of the payment service definition used to create this service.  | [optional] 

## Methods

### NewPaymentServiceSnapshot

`func NewPaymentServiceSnapshot() *PaymentServiceSnapshot`

NewPaymentServiceSnapshot instantiates a new PaymentServiceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceSnapshotWithDefaults

`func NewPaymentServiceSnapshotWithDefaults() *PaymentServiceSnapshot`

NewPaymentServiceSnapshotWithDefaults instantiates a new PaymentServiceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentServiceSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentServiceSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentServiceSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentServiceSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentServiceSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *PaymentServiceSnapshot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceSnapshot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceSnapshot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentServiceSnapshot) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentServiceSnapshot) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentServiceSnapshot) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentServiceSnapshot) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentServiceSnapshot) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentServiceSnapshot) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentServiceSnapshot) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentServiceSnapshot) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentServiceSnapshot) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


