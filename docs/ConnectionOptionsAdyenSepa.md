# ConnectionOptionsAdyenSepa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRescue** | Pointer to **bool** | Enabled Adyen&#39;s auto-rescue feature. | [optional] [default to false]
**MaxDaysToRescue** | Pointer to **NullableInt32** | Defines the number of days to auto-retry a payment for if &#x60;autoRescue&#x60; is enabled. | [optional] 
**AutoRescueSepaScenario** | Pointer to **NullableString** | Defines the Adyen auto-rescue test scenario to invoke. | [optional] 
**OwnerName** | Pointer to **NullableString** | Defines the type of chargeback that you want to simulate. | [optional] 

## Methods

### NewConnectionOptionsAdyenSepa

`func NewConnectionOptionsAdyenSepa() *ConnectionOptionsAdyenSepa`

NewConnectionOptionsAdyenSepa instantiates a new ConnectionOptionsAdyenSepa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsAdyenSepaWithDefaults

`func NewConnectionOptionsAdyenSepaWithDefaults() *ConnectionOptionsAdyenSepa`

NewConnectionOptionsAdyenSepaWithDefaults instantiates a new ConnectionOptionsAdyenSepa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRescue

`func (o *ConnectionOptionsAdyenSepa) GetAutoRescue() bool`

GetAutoRescue returns the AutoRescue field if non-nil, zero value otherwise.

### GetAutoRescueOk

`func (o *ConnectionOptionsAdyenSepa) GetAutoRescueOk() (*bool, bool)`

GetAutoRescueOk returns a tuple with the AutoRescue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRescue

`func (o *ConnectionOptionsAdyenSepa) SetAutoRescue(v bool)`

SetAutoRescue sets AutoRescue field to given value.

### HasAutoRescue

`func (o *ConnectionOptionsAdyenSepa) HasAutoRescue() bool`

HasAutoRescue returns a boolean if a field has been set.

### GetMaxDaysToRescue

`func (o *ConnectionOptionsAdyenSepa) GetMaxDaysToRescue() int32`

GetMaxDaysToRescue returns the MaxDaysToRescue field if non-nil, zero value otherwise.

### GetMaxDaysToRescueOk

`func (o *ConnectionOptionsAdyenSepa) GetMaxDaysToRescueOk() (*int32, bool)`

GetMaxDaysToRescueOk returns a tuple with the MaxDaysToRescue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysToRescue

`func (o *ConnectionOptionsAdyenSepa) SetMaxDaysToRescue(v int32)`

SetMaxDaysToRescue sets MaxDaysToRescue field to given value.

### HasMaxDaysToRescue

`func (o *ConnectionOptionsAdyenSepa) HasMaxDaysToRescue() bool`

HasMaxDaysToRescue returns a boolean if a field has been set.

### SetMaxDaysToRescueNil

`func (o *ConnectionOptionsAdyenSepa) SetMaxDaysToRescueNil(b bool)`

 SetMaxDaysToRescueNil sets the value for MaxDaysToRescue to be an explicit nil

### UnsetMaxDaysToRescue
`func (o *ConnectionOptionsAdyenSepa) UnsetMaxDaysToRescue()`

UnsetMaxDaysToRescue ensures that no value is present for MaxDaysToRescue, not even an explicit nil
### GetAutoRescueSepaScenario

`func (o *ConnectionOptionsAdyenSepa) GetAutoRescueSepaScenario() string`

GetAutoRescueSepaScenario returns the AutoRescueSepaScenario field if non-nil, zero value otherwise.

### GetAutoRescueSepaScenarioOk

`func (o *ConnectionOptionsAdyenSepa) GetAutoRescueSepaScenarioOk() (*string, bool)`

GetAutoRescueSepaScenarioOk returns a tuple with the AutoRescueSepaScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRescueSepaScenario

`func (o *ConnectionOptionsAdyenSepa) SetAutoRescueSepaScenario(v string)`

SetAutoRescueSepaScenario sets AutoRescueSepaScenario field to given value.

### HasAutoRescueSepaScenario

`func (o *ConnectionOptionsAdyenSepa) HasAutoRescueSepaScenario() bool`

HasAutoRescueSepaScenario returns a boolean if a field has been set.

### SetAutoRescueSepaScenarioNil

`func (o *ConnectionOptionsAdyenSepa) SetAutoRescueSepaScenarioNil(b bool)`

 SetAutoRescueSepaScenarioNil sets the value for AutoRescueSepaScenario to be an explicit nil

### UnsetAutoRescueSepaScenario
`func (o *ConnectionOptionsAdyenSepa) UnsetAutoRescueSepaScenario()`

UnsetAutoRescueSepaScenario ensures that no value is present for AutoRescueSepaScenario, not even an explicit nil
### GetOwnerName

`func (o *ConnectionOptionsAdyenSepa) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ConnectionOptionsAdyenSepa) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ConnectionOptionsAdyenSepa) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ConnectionOptionsAdyenSepa) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *ConnectionOptionsAdyenSepa) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *ConnectionOptionsAdyenSepa) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


