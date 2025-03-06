/*
Gr4vy API

Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.

API version: 1.1.0-beta
Contact: code@gr4vy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CartItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartItem{}

// CartItem A cart item that represents a single cart line item for a transaction. Note that some optional properties are required for certain payment service providers. If no value is set for these properties, we will use their default value.  If the total due to be paid for the item is required by the payment service provider, generally referred to as the \"total amount\", the formula below will usually be used to calculate this amount:  `(unit_amount * quantity) - discount_amount + tax_amount`  It's highly recommended that the total amount to pay for all items should match the transaction's amount to reduce the risk of the transaction being declined by the payment service provider.
type CartItem struct {
	// The name of the cart item. The value you set for this property may be truncated if the maximum length accepted by a payment service provider is less than 255 characters.
	Name string `json:"name"`
	// The quantity of this item in the cart. This value cannot be negative or zero.
	Quantity int32 `json:"quantity"`
	// The amount for an individual item represented as a monetary amount in the smallest currency unit for the given currency, for example `1299` USD cents represents `$12.99`. The amount sent through to the payment processor as unitary amount will be calculated to include the discount and tax values sent as part of this cart item.
	UnitAmount int32 `json:"unit_amount"`
	// The amount discounted for this item represented as a monetary amount in the smallest currency unit for the given currency, for example `1299` USD cents represents `$12.99`.  Please note that this amount is for the total of the cart item and not for an individual item. For example, if the quantity is 5, this value should be the total discount amount for 5 of the cart item.  You might see unexpected failed transactions if the `discount_amount` can not be equally divided by the `quantity` value. This is due to the fact that some payment services require this amount to be specified per unit.  In this situation we recommend splitting this item into separate items, each with their own specific discount.
	DiscountAmount NullableInt32 `json:"discount_amount,omitempty"`
	// The tax amount for this item represented as a monetary amount in the smallest currency unit for the given currency, for example `1299` USD cents represents `$12.99`.  Please not that this amount is for the total of the cart item and not for an individual item. For example, if the quantity is 5, this value should be the total tax amount for 5 of the cart item.  You might see unexpected failed transactions if the `tax_amount` can not be equally divided by the `quantity` value. This is due to the fact that some payment services require this amount to be specified per unit.  In this situation we recommend splitting this item into separate items, each with their own specific tax amount.
	TaxAmount NullableInt32 `json:"tax_amount,omitempty"`
	// An external identifier for the cart item. This can be set to any value and is not sent to the payment service.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The SKU for the item.
	Sku NullableString `json:"sku,omitempty"`
	// The product URL for the item.
	ProductUrl NullableString `json:"product_url,omitempty"`
	// The URL for the image of the item.
	ImageUrl NullableString `json:"image_url,omitempty"`
	// A list of strings containing product categories for the item. Max length per item: 50.
	Categories []string `json:"categories,omitempty"`
	// The product type of the cart item.
	ProductType NullableString `json:"product_type,omitempty"`
	// Sellect country
	SellerCountry NullableString `json:"seller_country,omitempty"`
}

type _CartItem CartItem

// NewCartItem instantiates a new CartItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartItem(name string, quantity int32, unitAmount int32) *CartItem {
	this := CartItem{}
	this.Name = name
	this.Quantity = quantity
	this.UnitAmount = unitAmount
	var discountAmount int32 = 0
	this.DiscountAmount = *NewNullableInt32(&discountAmount)
	var taxAmount int32 = 0
	this.TaxAmount = *NewNullableInt32(&taxAmount)
	return &this
}

// NewCartItemWithDefaults instantiates a new CartItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartItemWithDefaults() *CartItem {
	this := CartItem{}
	var discountAmount int32 = 0
	this.DiscountAmount = *NewNullableInt32(&discountAmount)
	var taxAmount int32 = 0
	this.TaxAmount = *NewNullableInt32(&taxAmount)
	return &this
}

// GetName returns the Name field value
func (o *CartItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CartItem) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *CartItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CartItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUnitAmount returns the UnitAmount field value
func (o *CartItem) GetUnitAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value
// and a boolean to check if the value has been set.
func (o *CartItem) GetUnitAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmount, true
}

// SetUnitAmount sets field value
func (o *CartItem) SetUnitAmount(v int32) {
	o.UnitAmount = v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetDiscountAmount() int32 {
	if o == nil || IsNil(o.DiscountAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.DiscountAmount.Get()
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetDiscountAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountAmount.Get(), o.DiscountAmount.IsSet()
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *CartItem) HasDiscountAmount() bool {
	if o != nil && o.DiscountAmount.IsSet() {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given NullableInt32 and assigns it to the DiscountAmount field.
func (o *CartItem) SetDiscountAmount(v int32) {
	o.DiscountAmount.Set(&v)
}
// SetDiscountAmountNil sets the value for DiscountAmount to be an explicit nil
func (o *CartItem) SetDiscountAmountNil() {
	o.DiscountAmount.Set(nil)
}

// UnsetDiscountAmount ensures that no value is present for DiscountAmount, not even an explicit nil
func (o *CartItem) UnsetDiscountAmount() {
	o.DiscountAmount.Unset()
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetTaxAmount() int32 {
	if o == nil || IsNil(o.TaxAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.TaxAmount.Get()
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetTaxAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxAmount.Get(), o.TaxAmount.IsSet()
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *CartItem) HasTaxAmount() bool {
	if o != nil && o.TaxAmount.IsSet() {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given NullableInt32 and assigns it to the TaxAmount field.
func (o *CartItem) SetTaxAmount(v int32) {
	o.TaxAmount.Set(&v)
}
// SetTaxAmountNil sets the value for TaxAmount to be an explicit nil
func (o *CartItem) SetTaxAmountNil() {
	o.TaxAmount.Set(nil)
}

// UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
func (o *CartItem) UnsetTaxAmount() {
	o.TaxAmount.Unset()
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetExternalIdentifier() string {
	if o == nil || IsNil(o.ExternalIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetExternalIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *CartItem) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *CartItem) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *CartItem) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *CartItem) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetSku() string {
	if o == nil || IsNil(o.Sku.Get()) {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *CartItem) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *CartItem) SetSku(v string) {
	o.Sku.Set(&v)
}
// SetSkuNil sets the value for Sku to be an explicit nil
func (o *CartItem) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *CartItem) UnsetSku() {
	o.Sku.Unset()
}

// GetProductUrl returns the ProductUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetProductUrl() string {
	if o == nil || IsNil(o.ProductUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProductUrl.Get()
}

// GetProductUrlOk returns a tuple with the ProductUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetProductUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductUrl.Get(), o.ProductUrl.IsSet()
}

// HasProductUrl returns a boolean if a field has been set.
func (o *CartItem) HasProductUrl() bool {
	if o != nil && o.ProductUrl.IsSet() {
		return true
	}

	return false
}

// SetProductUrl gets a reference to the given NullableString and assigns it to the ProductUrl field.
func (o *CartItem) SetProductUrl(v string) {
	o.ProductUrl.Set(&v)
}
// SetProductUrlNil sets the value for ProductUrl to be an explicit nil
func (o *CartItem) SetProductUrlNil() {
	o.ProductUrl.Set(nil)
}

// UnsetProductUrl ensures that no value is present for ProductUrl, not even an explicit nil
func (o *CartItem) UnsetProductUrl() {
	o.ProductUrl.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CartItem) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *CartItem) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *CartItem) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *CartItem) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetCategories returns the Categories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CartItem) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *CartItem) SetCategories(v []string) {
	o.Categories = v
}

// GetProductType returns the ProductType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartItem) GetProductType() string {
	if o == nil || IsNil(o.ProductType.Get()) {
		var ret string
		return ret
	}
	return *o.ProductType.Get()
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartItem) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductType.Get(), o.ProductType.IsSet()
}

// HasProductType returns a boolean if a field has been set.
func (o *CartItem) HasProductType() bool {
	if o != nil && o.ProductType.IsSet() {
		return true
	}

	return false
}

// SetProductType gets a reference to the given NullableString and assigns it to the ProductType field.
func (o *CartItem) SetProductType(v string) {
	o.ProductType.Set(&v)
}
// SetProductTypeNil sets the value for ProductType to be an explicit nil
func (o *CartItem) SetProductTypeNil() {
	o.ProductType.Set(nil)
}

// UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
func (o *CartItem) UnsetProductType() {
	o.ProductType.Unset()
}

func (o CartItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["quantity"] = o.Quantity
	toSerialize["unit_amount"] = o.UnitAmount
	if o.DiscountAmount.IsSet() {
		toSerialize["discount_amount"] = o.DiscountAmount.Get()
	}
	if o.TaxAmount.IsSet() {
		toSerialize["tax_amount"] = o.TaxAmount.Get()
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	if o.ProductUrl.IsSet() {
		toSerialize["product_url"] = o.ProductUrl.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.ProductType.IsSet() {
		toSerialize["product_type"] = o.ProductType.Get()
	}
	return toSerialize, nil
}

func (o *CartItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"quantity",
		"unit_amount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCartItem := _CartItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCartItem)

	if err != nil {
		return err
	}

	*o = CartItem(varCartItem)

	return err
}

type NullableCartItem struct {
	value *CartItem
	isSet bool
}

func (v NullableCartItem) Get() *CartItem {
	return v.value
}

func (v *NullableCartItem) Set(val *CartItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCartItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCartItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartItem(val *CartItem) *NullableCartItem {
	return &NullableCartItem{value: val, isSet: true}
}

func (v NullableCartItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


