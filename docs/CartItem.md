# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the cart item. The value you set for this property may be truncated if the maximum length accepted by a payment service provider is less than 255 characters. | 
**Quantity** | **int32** | The quantity of this item in the cart. This value cannot be negative or zero. | 
**UnitAmount** | **int32** | The amount for an individual item represented as a monetary amount in the smallest currency unit for the given currency, for example &#x60;1299&#x60; USD cents represents &#x60;$12.99&#x60;. The amount sent through to the payment processor as unitary amount will be calculated to include the discount and tax values sent as part of this cart item. | 
**DiscountAmount** | Pointer to **NullableInt32** | The amount discounted for this item represented as a monetary amount in the smallest currency unit for the given currency, for example &#x60;1299&#x60; USD cents represents &#x60;$12.99&#x60;.  Please note that this amount is for the total of the cart item and not for an individual item. For example, if the quantity is 5, this value should be the total discount amount for 5 of the cart item.  You might see unexpected failed transactions if the &#x60;discount_amount&#x60; can not be equally divided by the &#x60;quantity&#x60; value. This is due to the fact that some payment services require this amount to be specified per unit.  In this situation we recommend splitting this item into separate items, each with their own specific discount. | [optional] [default to 0]
**TaxAmount** | Pointer to **NullableInt32** | The tax amount for this item represented as a monetary amount in the smallest currency unit for the given currency, for example &#x60;1299&#x60; USD cents represents &#x60;$12.99&#x60;.  Please not that this amount is for the total of the cart item and not for an individual item. For example, if the quantity is 5, this value should be the total tax amount for 5 of the cart item.  You might see unexpected failed transactions if the &#x60;tax_amount&#x60; can not be equally divided by the &#x60;quantity&#x60; value. This is due to the fact that some payment services require this amount to be specified per unit.  In this situation we recommend splitting this item into separate items, each with their own specific tax amount. | [optional] [default to 0]
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier for the cart item. This can be set to any value and is not sent to the payment service. | [optional] 
**Sku** | Pointer to **NullableString** | The SKU for the item. | [optional] 
**ProductUrl** | Pointer to **NullableString** | The product URL for the item. | [optional] 
**ImageUrl** | Pointer to **NullableString** | The URL for the image of the item. | [optional] 
**Categories** | Pointer to **[]string** | A list of strings containing product categories for the item. Max length per item: 50. | [optional] 
**ProductType** | Pointer to **NullableString** | The product type of the cart item. | [optional] 

## Methods

### NewCartItem

`func NewCartItem(name string, quantity int32, unitAmount int32, ) *CartItem`

NewCartItem instantiates a new CartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemWithDefaults

`func NewCartItemWithDefaults() *CartItem`

NewCartItemWithDefaults instantiates a new CartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CartItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartItem) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *CartItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitAmount

`func (o *CartItem) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *CartItem) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *CartItem) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### GetDiscountAmount

`func (o *CartItem) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *CartItem) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *CartItem) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *CartItem) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### SetDiscountAmountNil

`func (o *CartItem) SetDiscountAmountNil(b bool)`

 SetDiscountAmountNil sets the value for DiscountAmount to be an explicit nil

### UnsetDiscountAmount
`func (o *CartItem) UnsetDiscountAmount()`

UnsetDiscountAmount ensures that no value is present for DiscountAmount, not even an explicit nil
### GetTaxAmount

`func (o *CartItem) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *CartItem) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *CartItem) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *CartItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### SetTaxAmountNil

`func (o *CartItem) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *CartItem) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetExternalIdentifier

`func (o *CartItem) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *CartItem) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *CartItem) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *CartItem) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *CartItem) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *CartItem) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetSku

`func (o *CartItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CartItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CartItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CartItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CartItem) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CartItem) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetProductUrl

`func (o *CartItem) GetProductUrl() string`

GetProductUrl returns the ProductUrl field if non-nil, zero value otherwise.

### GetProductUrlOk

`func (o *CartItem) GetProductUrlOk() (*string, bool)`

GetProductUrlOk returns a tuple with the ProductUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUrl

`func (o *CartItem) SetProductUrl(v string)`

SetProductUrl sets ProductUrl field to given value.

### HasProductUrl

`func (o *CartItem) HasProductUrl() bool`

HasProductUrl returns a boolean if a field has been set.

### SetProductUrlNil

`func (o *CartItem) SetProductUrlNil(b bool)`

 SetProductUrlNil sets the value for ProductUrl to be an explicit nil

### UnsetProductUrl
`func (o *CartItem) UnsetProductUrl()`

UnsetProductUrl ensures that no value is present for ProductUrl, not even an explicit nil
### GetImageUrl

`func (o *CartItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CartItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CartItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CartItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CartItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CartItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetCategories

`func (o *CartItem) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CartItem) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CartItem) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CartItem) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *CartItem) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *CartItem) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetProductType

`func (o *CartItem) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *CartItem) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *CartItem) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *CartItem) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *CartItem) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *CartItem) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


