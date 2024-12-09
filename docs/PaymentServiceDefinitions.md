# PaymentServiceDefinitions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PaymentServiceDefinition**](PaymentServiceDefinition.md) |  | [optional] 
**Limit** | Pointer to **int32** | The limit applied to request. This represents the number of items that are at maximum returned by this request. | [optional] [default to 20]
**NextCursor** | Pointer to **NullableString** | The cursor that represents the next page of results. Use the &#x60;cursor&#x60; query parameter to fetch this page of items. | [optional] 
**PreviousCursor** | Pointer to **NullableString** | The cursor that represents the next page of results. Use the &#x60;cursor&#x60; query parameter to fetch this page of items. | [optional] 

## Methods

### NewPaymentServiceDefinitions

`func NewPaymentServiceDefinitions() *PaymentServiceDefinitions`

NewPaymentServiceDefinitions instantiates a new PaymentServiceDefinitions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceDefinitionsWithDefaults

`func NewPaymentServiceDefinitionsWithDefaults() *PaymentServiceDefinitions`

NewPaymentServiceDefinitionsWithDefaults instantiates a new PaymentServiceDefinitions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PaymentServiceDefinitions) GetItems() []PaymentServiceDefinition`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaymentServiceDefinitions) GetItemsOk() (*[]PaymentServiceDefinition, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaymentServiceDefinitions) SetItems(v []PaymentServiceDefinition)`

SetItems sets Items field to given value.

### HasItems

`func (o *PaymentServiceDefinitions) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *PaymentServiceDefinitions) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaymentServiceDefinitions) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaymentServiceDefinitions) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaymentServiceDefinitions) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *PaymentServiceDefinitions) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaymentServiceDefinitions) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaymentServiceDefinitions) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaymentServiceDefinitions) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *PaymentServiceDefinitions) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PaymentServiceDefinitions) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPreviousCursor

`func (o *PaymentServiceDefinitions) GetPreviousCursor() string`

GetPreviousCursor returns the PreviousCursor field if non-nil, zero value otherwise.

### GetPreviousCursorOk

`func (o *PaymentServiceDefinitions) GetPreviousCursorOk() (*string, bool)`

GetPreviousCursorOk returns a tuple with the PreviousCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCursor

`func (o *PaymentServiceDefinitions) SetPreviousCursor(v string)`

SetPreviousCursor sets PreviousCursor field to given value.

### HasPreviousCursor

`func (o *PaymentServiceDefinitions) HasPreviousCursor() bool`

HasPreviousCursor returns a boolean if a field has been set.

### SetPreviousCursorNil

`func (o *PaymentServiceDefinitions) SetPreviousCursorNil(b bool)`

 SetPreviousCursorNil sets the value for PreviousCursor to be an explicit nil

### UnsetPreviousCursor
`func (o *PaymentServiceDefinitions) UnsetPreviousCursor()`

UnsetPreviousCursor ensures that no value is present for PreviousCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


