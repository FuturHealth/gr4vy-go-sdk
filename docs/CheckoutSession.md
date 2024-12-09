# CheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;checkout-session&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID of the Checkout Session. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The date and time when the Checkout Session will expire. By default this will be set to 1 hour from the date of creation. | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 
**Metadata** | Pointer to **map[string]string** | Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it. | [optional] 
**Airline** | Pointer to [**NullableAirline**](Airline.md) | Contains information about an airline travel, if applicable. | [optional] 
**PaymentMethod** | Pointer to [**NullableCheckoutSessionPaymentMethod**](CheckoutSessionPaymentMethod.md) |  | [optional] 

## Methods

### NewCheckoutSession

`func NewCheckoutSession() *CheckoutSession`

NewCheckoutSession instantiates a new CheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionWithDefaults

`func NewCheckoutSessionWithDefaults() *CheckoutSession`

NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutSession) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutSession) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutSession) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckoutSession) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CheckoutSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckoutSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CheckoutSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutSession) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCartItems

`func (o *CheckoutSession) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CheckoutSession) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *CheckoutSession) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *CheckoutSession) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItemsNil

`func (o *CheckoutSession) SetCartItemsNil(b bool)`

 SetCartItemsNil sets the value for CartItems to be an explicit nil

### UnsetCartItems
`func (o *CheckoutSession) UnsetCartItems()`

UnsetCartItems ensures that no value is present for CartItems, not even an explicit nil
### GetMetadata

`func (o *CheckoutSession) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutSession) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutSession) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutSession) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CheckoutSession) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CheckoutSession) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAirline

`func (o *CheckoutSession) GetAirline() Airline`

GetAirline returns the Airline field if non-nil, zero value otherwise.

### GetAirlineOk

`func (o *CheckoutSession) GetAirlineOk() (*Airline, bool)`

GetAirlineOk returns a tuple with the Airline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirline

`func (o *CheckoutSession) SetAirline(v Airline)`

SetAirline sets Airline field to given value.

### HasAirline

`func (o *CheckoutSession) HasAirline() bool`

HasAirline returns a boolean if a field has been set.

### SetAirlineNil

`func (o *CheckoutSession) SetAirlineNil(b bool)`

 SetAirlineNil sets the value for Airline to be an explicit nil

### UnsetAirline
`func (o *CheckoutSession) UnsetAirline()`

UnsetAirline ensures that no value is present for Airline, not even an explicit nil
### GetPaymentMethod

`func (o *CheckoutSession) GetPaymentMethod() CheckoutSessionPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CheckoutSession) GetPaymentMethodOk() (*CheckoutSessionPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CheckoutSession) SetPaymentMethod(v CheckoutSessionPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CheckoutSession) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *CheckoutSession) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *CheckoutSession) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


