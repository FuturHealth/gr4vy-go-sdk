# ConnectionOptionsStripeCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorOnRequiresAction** | Pointer to **bool** | Defines if Stripe should automatically fail the payment if it requires two-factor authentication from the user. | [optional] [default to false]
**StripeConnect** | Pointer to [**NullableConnectionOptionsStripeCardStripeConnect**](ConnectionOptionsStripeCardStripeConnect.md) |  | [optional] 

## Methods

### NewConnectionOptionsStripeCard

`func NewConnectionOptionsStripeCard() *ConnectionOptionsStripeCard`

NewConnectionOptionsStripeCard instantiates a new ConnectionOptionsStripeCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsStripeCardWithDefaults

`func NewConnectionOptionsStripeCardWithDefaults() *ConnectionOptionsStripeCard`

NewConnectionOptionsStripeCardWithDefaults instantiates a new ConnectionOptionsStripeCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorOnRequiresAction

`func (o *ConnectionOptionsStripeCard) GetErrorOnRequiresAction() bool`

GetErrorOnRequiresAction returns the ErrorOnRequiresAction field if non-nil, zero value otherwise.

### GetErrorOnRequiresActionOk

`func (o *ConnectionOptionsStripeCard) GetErrorOnRequiresActionOk() (*bool, bool)`

GetErrorOnRequiresActionOk returns a tuple with the ErrorOnRequiresAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOnRequiresAction

`func (o *ConnectionOptionsStripeCard) SetErrorOnRequiresAction(v bool)`

SetErrorOnRequiresAction sets ErrorOnRequiresAction field to given value.

### HasErrorOnRequiresAction

`func (o *ConnectionOptionsStripeCard) HasErrorOnRequiresAction() bool`

HasErrorOnRequiresAction returns a boolean if a field has been set.

### GetStripeConnect

`func (o *ConnectionOptionsStripeCard) GetStripeConnect() ConnectionOptionsStripeCardStripeConnect`

GetStripeConnect returns the StripeConnect field if non-nil, zero value otherwise.

### GetStripeConnectOk

`func (o *ConnectionOptionsStripeCard) GetStripeConnectOk() (*ConnectionOptionsStripeCardStripeConnect, bool)`

GetStripeConnectOk returns a tuple with the StripeConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeConnect

`func (o *ConnectionOptionsStripeCard) SetStripeConnect(v ConnectionOptionsStripeCardStripeConnect)`

SetStripeConnect sets StripeConnect field to given value.

### HasStripeConnect

`func (o *ConnectionOptionsStripeCard) HasStripeConnect() bool`

HasStripeConnect returns a boolean if a field has been set.

### SetStripeConnectNil

`func (o *ConnectionOptionsStripeCard) SetStripeConnectNil(b bool)`

 SetStripeConnectNil sets the value for StripeConnect to be an explicit nil

### UnsetStripeConnect
`func (o *ConnectionOptionsStripeCard) UnsetStripeConnect()`

UnsetStripeConnect ensures that no value is present for StripeConnect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


