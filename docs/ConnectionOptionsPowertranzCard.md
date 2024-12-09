# ConnectionOptionsPowertranzCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipThreeDSecure** | Pointer to **bool** | Allows skipping of 3DS. Without this, every transaction will be sent via Powertranz&#39;s 3DS server, even though 3DS might not be triggered. | [optional] [default to false]

## Methods

### NewConnectionOptionsPowertranzCard

`func NewConnectionOptionsPowertranzCard() *ConnectionOptionsPowertranzCard`

NewConnectionOptionsPowertranzCard instantiates a new ConnectionOptionsPowertranzCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsPowertranzCardWithDefaults

`func NewConnectionOptionsPowertranzCardWithDefaults() *ConnectionOptionsPowertranzCard`

NewConnectionOptionsPowertranzCardWithDefaults instantiates a new ConnectionOptionsPowertranzCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipThreeDSecure

`func (o *ConnectionOptionsPowertranzCard) GetSkipThreeDSecure() bool`

GetSkipThreeDSecure returns the SkipThreeDSecure field if non-nil, zero value otherwise.

### GetSkipThreeDSecureOk

`func (o *ConnectionOptionsPowertranzCard) GetSkipThreeDSecureOk() (*bool, bool)`

GetSkipThreeDSecureOk returns a tuple with the SkipThreeDSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipThreeDSecure

`func (o *ConnectionOptionsPowertranzCard) SetSkipThreeDSecure(v bool)`

SetSkipThreeDSecure sets SkipThreeDSecure field to given value.

### HasSkipThreeDSecure

`func (o *ConnectionOptionsPowertranzCard) HasSkipThreeDSecure() bool`

HasSkipThreeDSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


