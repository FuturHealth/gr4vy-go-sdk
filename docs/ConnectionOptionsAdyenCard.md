# ConnectionOptionsAdyenCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | A key-value object representing additional data to be passed to Adyen. | [optional] 
**AutoRescue** | Pointer to **bool** | Enabled Adyen&#39;s auto-rescue feature. | [optional] [default to false]
**MaxDaysToRescue** | Pointer to **NullableInt32** | Defines the number of days to auto-retry a payment for if &#x60;autoRescue&#x60; is enabled. | [optional] 
**AutoRescueScenario** | Pointer to **NullableString** | Defines the Adyen auto-rescue test scenario to invoke. | [optional] 

## Methods

### NewConnectionOptionsAdyenCard

`func NewConnectionOptionsAdyenCard() *ConnectionOptionsAdyenCard`

NewConnectionOptionsAdyenCard instantiates a new ConnectionOptionsAdyenCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsAdyenCardWithDefaults

`func NewConnectionOptionsAdyenCardWithDefaults() *ConnectionOptionsAdyenCard`

NewConnectionOptionsAdyenCardWithDefaults instantiates a new ConnectionOptionsAdyenCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ConnectionOptionsAdyenCard) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ConnectionOptionsAdyenCard) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ConnectionOptionsAdyenCard) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ConnectionOptionsAdyenCard) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAutoRescue

`func (o *ConnectionOptionsAdyenCard) GetAutoRescue() bool`

GetAutoRescue returns the AutoRescue field if non-nil, zero value otherwise.

### GetAutoRescueOk

`func (o *ConnectionOptionsAdyenCard) GetAutoRescueOk() (*bool, bool)`

GetAutoRescueOk returns a tuple with the AutoRescue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRescue

`func (o *ConnectionOptionsAdyenCard) SetAutoRescue(v bool)`

SetAutoRescue sets AutoRescue field to given value.

### HasAutoRescue

`func (o *ConnectionOptionsAdyenCard) HasAutoRescue() bool`

HasAutoRescue returns a boolean if a field has been set.

### GetMaxDaysToRescue

`func (o *ConnectionOptionsAdyenCard) GetMaxDaysToRescue() int32`

GetMaxDaysToRescue returns the MaxDaysToRescue field if non-nil, zero value otherwise.

### GetMaxDaysToRescueOk

`func (o *ConnectionOptionsAdyenCard) GetMaxDaysToRescueOk() (*int32, bool)`

GetMaxDaysToRescueOk returns a tuple with the MaxDaysToRescue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysToRescue

`func (o *ConnectionOptionsAdyenCard) SetMaxDaysToRescue(v int32)`

SetMaxDaysToRescue sets MaxDaysToRescue field to given value.

### HasMaxDaysToRescue

`func (o *ConnectionOptionsAdyenCard) HasMaxDaysToRescue() bool`

HasMaxDaysToRescue returns a boolean if a field has been set.

### SetMaxDaysToRescueNil

`func (o *ConnectionOptionsAdyenCard) SetMaxDaysToRescueNil(b bool)`

 SetMaxDaysToRescueNil sets the value for MaxDaysToRescue to be an explicit nil

### UnsetMaxDaysToRescue
`func (o *ConnectionOptionsAdyenCard) UnsetMaxDaysToRescue()`

UnsetMaxDaysToRescue ensures that no value is present for MaxDaysToRescue, not even an explicit nil
### GetAutoRescueScenario

`func (o *ConnectionOptionsAdyenCard) GetAutoRescueScenario() string`

GetAutoRescueScenario returns the AutoRescueScenario field if non-nil, zero value otherwise.

### GetAutoRescueScenarioOk

`func (o *ConnectionOptionsAdyenCard) GetAutoRescueScenarioOk() (*string, bool)`

GetAutoRescueScenarioOk returns a tuple with the AutoRescueScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRescueScenario

`func (o *ConnectionOptionsAdyenCard) SetAutoRescueScenario(v string)`

SetAutoRescueScenario sets AutoRescueScenario field to given value.

### HasAutoRescueScenario

`func (o *ConnectionOptionsAdyenCard) HasAutoRescueScenario() bool`

HasAutoRescueScenario returns a boolean if a field has been set.

### SetAutoRescueScenarioNil

`func (o *ConnectionOptionsAdyenCard) SetAutoRescueScenarioNil(b bool)`

 SetAutoRescueScenarioNil sets the value for AutoRescueScenario to be an explicit nil

### UnsetAutoRescueScenario
`func (o *ConnectionOptionsAdyenCard) UnsetAutoRescueScenario()`

UnsetAutoRescueScenario ensures that no value is present for AutoRescueScenario, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


