# ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryType** | Pointer to **NullableString** | Value to populate the &#x60;deliveryType&#x60; field for this cart item. This overrides the type set at the wider level.  Represents the type of delivery. This can be set to &#x60;PHYSICAL&#x60; for any type of shipped goods, &#x60;DIGITAL&#x60; for non-shipped goods (services, gift cards etc.), or &#x60;HYBRID&#x60; for others. | [optional] 
**DeliveryMethod** | Pointer to **string** | Value to populate the &#x60;deliveryMethod&#x60; field for this cart item. This overrides the method set at the wider level.  Represents the delivery method chosen by customer such as postal service, email, in game transfer, etc. | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails

`func NewConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails() *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails`

NewConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails instantiates a new ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetailsWithDefaults

`func NewConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetailsWithDefaults() *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails`

NewConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetailsWithDefaults instantiates a new ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryType

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### SetDeliveryTypeNil

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) SetDeliveryTypeNil(b bool)`

 SetDeliveryTypeNil sets the value for DeliveryType to be an explicit nil

### UnsetDeliveryType
`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) UnsetDeliveryType()`

UnsetDeliveryType ensures that no value is present for DeliveryType, not even an explicit nil
### GetDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerDeliveryDetails) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


