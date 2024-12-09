# CheckoutSessionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 
**Metadata** | Pointer to **map[string]string** | Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it. | [optional] 
**Airline** | Pointer to [**NullableAirline**](Airline.md) | The airline addendum data which describes the airline booking associated with this transaction. | [optional] 
**Buyer** | Pointer to [**TransactionBuyerRequest**](TransactionBuyerRequest.md) |  | [optional] 

## Methods

### NewCheckoutSessionCreateRequest

`func NewCheckoutSessionCreateRequest() *CheckoutSessionCreateRequest`

NewCheckoutSessionCreateRequest instantiates a new CheckoutSessionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionCreateRequestWithDefaults

`func NewCheckoutSessionCreateRequestWithDefaults() *CheckoutSessionCreateRequest`

NewCheckoutSessionCreateRequestWithDefaults instantiates a new CheckoutSessionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartItems

`func (o *CheckoutSessionCreateRequest) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CheckoutSessionCreateRequest) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *CheckoutSessionCreateRequest) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *CheckoutSessionCreateRequest) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItemsNil

`func (o *CheckoutSessionCreateRequest) SetCartItemsNil(b bool)`

 SetCartItemsNil sets the value for CartItems to be an explicit nil

### UnsetCartItems
`func (o *CheckoutSessionCreateRequest) UnsetCartItems()`

UnsetCartItems ensures that no value is present for CartItems, not even an explicit nil
### GetMetadata

`func (o *CheckoutSessionCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutSessionCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutSessionCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutSessionCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CheckoutSessionCreateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CheckoutSessionCreateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAirline

`func (o *CheckoutSessionCreateRequest) GetAirline() Airline`

GetAirline returns the Airline field if non-nil, zero value otherwise.

### GetAirlineOk

`func (o *CheckoutSessionCreateRequest) GetAirlineOk() (*Airline, bool)`

GetAirlineOk returns a tuple with the Airline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirline

`func (o *CheckoutSessionCreateRequest) SetAirline(v Airline)`

SetAirline sets Airline field to given value.

### HasAirline

`func (o *CheckoutSessionCreateRequest) HasAirline() bool`

HasAirline returns a boolean if a field has been set.

### SetAirlineNil

`func (o *CheckoutSessionCreateRequest) SetAirlineNil(b bool)`

 SetAirlineNil sets the value for Airline to be an explicit nil

### UnsetAirline
`func (o *CheckoutSessionCreateRequest) UnsetAirline()`

UnsetAirline ensures that no value is present for Airline, not even an explicit nil
### GetBuyer

`func (o *CheckoutSessionCreateRequest) GetBuyer() TransactionBuyerRequest`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *CheckoutSessionCreateRequest) GetBuyerOk() (*TransactionBuyerRequest, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *CheckoutSessionCreateRequest) SetBuyer(v TransactionBuyerRequest)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *CheckoutSessionCreateRequest) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


