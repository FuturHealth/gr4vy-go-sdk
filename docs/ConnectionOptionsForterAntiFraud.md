# ConnectionOptionsForterAntiFraud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryType** | Pointer to **NullableString** | Value to populate the &#x60;deliveryType&#x60; field in &#x60;primaryDeliveryDetails&#x60;.  Represents the type of delivery. This can be set to &#x60;PHYSICAL&#x60; for any type of shipped goods, &#x60;DIGITAL&#x60; for non-shipped goods (services, gift cards etc.), or &#x60;HYBRID&#x60; for others. | [optional] 
**DeliveryMethod** | Pointer to **NullableString** | Value to populate the &#x60;deliveryMethod&#x60; field in &#x60;primaryDeliveryDetails&#x60;.  Represents the delivery method chosen by customer such as postal service, email, in game transfer, etc. | [optional] 
**IsGuestBuyer** | Pointer to **bool** | Defines if this is a guest check-out. This will redact the &#x60;accountId&#x60; and &#x60;created&#x60; fields from the &#x60;accountOwner&#x60; details sent to Forter. | [optional] [default to false]
**CartItems** | Pointer to [**[]ConnectionOptionsForterAntiFraudCartItemsInner**](ConnectionOptionsForterAntiFraudCartItemsInner.md) | A list of Forter cart item objects. These will be merged into the &#x60;cart_items&#x60; passed to the transaction. Every cart item here will be merged with a cart item on the transaction with the same index.  Together, these will augment the &#x60;cartItems&#x60; values sent to the Forter validation API. | [optional] 
**TotalDiscount** | Pointer to [**NullableConnectionOptionsForterAntiFraudTotalDiscount**](ConnectionOptionsForterAntiFraudTotalDiscount.md) |  | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraud

`func NewConnectionOptionsForterAntiFraud() *ConnectionOptionsForterAntiFraud`

NewConnectionOptionsForterAntiFraud instantiates a new ConnectionOptionsForterAntiFraud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudWithDefaults

`func NewConnectionOptionsForterAntiFraudWithDefaults() *ConnectionOptionsForterAntiFraud`

NewConnectionOptionsForterAntiFraudWithDefaults instantiates a new ConnectionOptionsForterAntiFraud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryType

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *ConnectionOptionsForterAntiFraud) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### SetDeliveryTypeNil

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryTypeNil(b bool)`

 SetDeliveryTypeNil sets the value for DeliveryType to be an explicit nil

### UnsetDeliveryType
`func (o *ConnectionOptionsForterAntiFraud) UnsetDeliveryType()`

UnsetDeliveryType ensures that no value is present for DeliveryType, not even an explicit nil
### GetDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraud) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### SetDeliveryMethodNil

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryMethodNil(b bool)`

 SetDeliveryMethodNil sets the value for DeliveryMethod to be an explicit nil

### UnsetDeliveryMethod
`func (o *ConnectionOptionsForterAntiFraud) UnsetDeliveryMethod()`

UnsetDeliveryMethod ensures that no value is present for DeliveryMethod, not even an explicit nil
### GetIsGuestBuyer

`func (o *ConnectionOptionsForterAntiFraud) GetIsGuestBuyer() bool`

GetIsGuestBuyer returns the IsGuestBuyer field if non-nil, zero value otherwise.

### GetIsGuestBuyerOk

`func (o *ConnectionOptionsForterAntiFraud) GetIsGuestBuyerOk() (*bool, bool)`

GetIsGuestBuyerOk returns a tuple with the IsGuestBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuestBuyer

`func (o *ConnectionOptionsForterAntiFraud) SetIsGuestBuyer(v bool)`

SetIsGuestBuyer sets IsGuestBuyer field to given value.

### HasIsGuestBuyer

`func (o *ConnectionOptionsForterAntiFraud) HasIsGuestBuyer() bool`

HasIsGuestBuyer returns a boolean if a field has been set.

### GetCartItems

`func (o *ConnectionOptionsForterAntiFraud) GetCartItems() []ConnectionOptionsForterAntiFraudCartItemsInner`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *ConnectionOptionsForterAntiFraud) GetCartItemsOk() (*[]ConnectionOptionsForterAntiFraudCartItemsInner, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *ConnectionOptionsForterAntiFraud) SetCartItems(v []ConnectionOptionsForterAntiFraudCartItemsInner)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *ConnectionOptionsForterAntiFraud) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *ConnectionOptionsForterAntiFraud) GetTotalDiscount() ConnectionOptionsForterAntiFraudTotalDiscount`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *ConnectionOptionsForterAntiFraud) GetTotalDiscountOk() (*ConnectionOptionsForterAntiFraudTotalDiscount, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *ConnectionOptionsForterAntiFraud) SetTotalDiscount(v ConnectionOptionsForterAntiFraudTotalDiscount)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *ConnectionOptionsForterAntiFraud) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscountNil

`func (o *ConnectionOptionsForterAntiFraud) SetTotalDiscountNil(b bool)`

 SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil

### UnsetTotalDiscount
`func (o *ConnectionOptionsForterAntiFraud) UnsetTotalDiscount()`

UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


