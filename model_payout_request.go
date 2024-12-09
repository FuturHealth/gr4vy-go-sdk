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

// checks if the PayoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayoutRequest{}

// PayoutRequest Request body for creating new payout.
type PayoutRequest struct {
	// The amount to payout.
	Amount float32 `json:"amount"`
	// The ISO-4217 currency code for the payout.
	Currency string `json:"currency"`
	// The ID of the payment service to use for the payout.
	PaymentServiceId string `json:"payment_service_id"`
	// The type of payout to process.
	Category *string `json:"category,omitempty"`
	// A value that can be used to match the payout against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The ID of the buyer to send the payout.
	BuyerId NullableString `json:"buyer_id,omitempty"`
	// The `external_identifier` of the buyer to send this payout to.
	BuyerExternalIdentifier NullableString `json:"buyer_external_identifier,omitempty"`
	// Inline buyer details for the payout.
	Buyer NullableTransactionBuyerRequest `json:"buyer,omitempty"`
	// Merchant information for the source of the payout.
	Merchant NullableMerchantRequest `json:"merchant,omitempty"`
	PaymentMethod PayoutPaymentMethodRequest `json:"payment_method"`
	// Optional fields for processing payouts on specific payment services.
	ConnectionOptions NullablePayoutConnectionOptionsRequest `json:"connection_options,omitempty"`
}

type _PayoutRequest PayoutRequest

// NewPayoutRequest instantiates a new PayoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutRequest(amount float32, currency string, paymentServiceId string, paymentMethod PayoutPaymentMethodRequest) *PayoutRequest {
	this := PayoutRequest{}
	this.Amount = amount
	this.Currency = currency
	this.PaymentServiceId = paymentServiceId
	var category string = "online_gambling"
	this.Category = &category
	this.PaymentMethod = paymentMethod
	return &this
}

// NewPayoutRequestWithDefaults instantiates a new PayoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutRequestWithDefaults() *PayoutRequest {
	this := PayoutRequest{}
	var category string = "online_gambling"
	this.Category = &category
	return &this
}

// GetAmount returns the Amount field value
func (o *PayoutRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PayoutRequest) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *PayoutRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PayoutRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetPaymentServiceId returns the PaymentServiceId field value
func (o *PayoutRequest) GetPaymentServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentServiceId
}

// GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetPaymentServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentServiceId, true
}

// SetPaymentServiceId sets field value
func (o *PayoutRequest) SetPaymentServiceId(v string) {
	o.PaymentServiceId = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *PayoutRequest) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *PayoutRequest) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *PayoutRequest) SetCategory(v string) {
	o.Category = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutRequest) GetExternalIdentifier() string {
	if o == nil || IsNil(o.ExternalIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *PayoutRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *PayoutRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *PayoutRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *PayoutRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetBuyerId returns the BuyerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutRequest) GetBuyerId() string {
	if o == nil || IsNil(o.BuyerId.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerId.Get()
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutRequest) GetBuyerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerId.Get(), o.BuyerId.IsSet()
}

// HasBuyerId returns a boolean if a field has been set.
func (o *PayoutRequest) HasBuyerId() bool {
	if o != nil && o.BuyerId.IsSet() {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given NullableString and assigns it to the BuyerId field.
func (o *PayoutRequest) SetBuyerId(v string) {
	o.BuyerId.Set(&v)
}
// SetBuyerIdNil sets the value for BuyerId to be an explicit nil
func (o *PayoutRequest) SetBuyerIdNil() {
	o.BuyerId.Set(nil)
}

// UnsetBuyerId ensures that no value is present for BuyerId, not even an explicit nil
func (o *PayoutRequest) UnsetBuyerId() {
	o.BuyerId.Unset()
}

// GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutRequest) GetBuyerExternalIdentifier() string {
	if o == nil || IsNil(o.BuyerExternalIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerExternalIdentifier.Get()
}

// GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutRequest) GetBuyerExternalIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerExternalIdentifier.Get(), o.BuyerExternalIdentifier.IsSet()
}

// HasBuyerExternalIdentifier returns a boolean if a field has been set.
func (o *PayoutRequest) HasBuyerExternalIdentifier() bool {
	if o != nil && o.BuyerExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetBuyerExternalIdentifier gets a reference to the given NullableString and assigns it to the BuyerExternalIdentifier field.
func (o *PayoutRequest) SetBuyerExternalIdentifier(v string) {
	o.BuyerExternalIdentifier.Set(&v)
}
// SetBuyerExternalIdentifierNil sets the value for BuyerExternalIdentifier to be an explicit nil
func (o *PayoutRequest) SetBuyerExternalIdentifierNil() {
	o.BuyerExternalIdentifier.Set(nil)
}

// UnsetBuyerExternalIdentifier ensures that no value is present for BuyerExternalIdentifier, not even an explicit nil
func (o *PayoutRequest) UnsetBuyerExternalIdentifier() {
	o.BuyerExternalIdentifier.Unset()
}

// GetBuyer returns the Buyer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutRequest) GetBuyer() TransactionBuyerRequest {
	if o == nil || IsNil(o.Buyer.Get()) {
		var ret TransactionBuyerRequest
		return ret
	}
	return *o.Buyer.Get()
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutRequest) GetBuyerOk() (*TransactionBuyerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buyer.Get(), o.Buyer.IsSet()
}

// HasBuyer returns a boolean if a field has been set.
func (o *PayoutRequest) HasBuyer() bool {
	if o != nil && o.Buyer.IsSet() {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given NullableTransactionBuyerRequest and assigns it to the Buyer field.
func (o *PayoutRequest) SetBuyer(v TransactionBuyerRequest) {
	o.Buyer.Set(&v)
}
// SetBuyerNil sets the value for Buyer to be an explicit nil
func (o *PayoutRequest) SetBuyerNil() {
	o.Buyer.Set(nil)
}

// UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
func (o *PayoutRequest) UnsetBuyer() {
	o.Buyer.Unset()
}

// GetMerchant returns the Merchant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutRequest) GetMerchant() MerchantRequest {
	if o == nil || IsNil(o.Merchant.Get()) {
		var ret MerchantRequest
		return ret
	}
	return *o.Merchant.Get()
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutRequest) GetMerchantOk() (*MerchantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Merchant.Get(), o.Merchant.IsSet()
}

// HasMerchant returns a boolean if a field has been set.
func (o *PayoutRequest) HasMerchant() bool {
	if o != nil && o.Merchant.IsSet() {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given NullableMerchantRequest and assigns it to the Merchant field.
func (o *PayoutRequest) SetMerchant(v MerchantRequest) {
	o.Merchant.Set(&v)
}
// SetMerchantNil sets the value for Merchant to be an explicit nil
func (o *PayoutRequest) SetMerchantNil() {
	o.Merchant.Set(nil)
}

// UnsetMerchant ensures that no value is present for Merchant, not even an explicit nil
func (o *PayoutRequest) UnsetMerchant() {
	o.Merchant.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *PayoutRequest) GetPaymentMethod() PayoutPaymentMethodRequest {
	if o == nil {
		var ret PayoutPaymentMethodRequest
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *PayoutRequest) GetPaymentMethodOk() (*PayoutPaymentMethodRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *PayoutRequest) SetPaymentMethod(v PayoutPaymentMethodRequest) {
	o.PaymentMethod = v
}

// GetConnectionOptions returns the ConnectionOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayoutRequest) GetConnectionOptions() PayoutConnectionOptionsRequest {
	if o == nil || IsNil(o.ConnectionOptions.Get()) {
		var ret PayoutConnectionOptionsRequest
		return ret
	}
	return *o.ConnectionOptions.Get()
}

// GetConnectionOptionsOk returns a tuple with the ConnectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayoutRequest) GetConnectionOptionsOk() (*PayoutConnectionOptionsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionOptions.Get(), o.ConnectionOptions.IsSet()
}

// HasConnectionOptions returns a boolean if a field has been set.
func (o *PayoutRequest) HasConnectionOptions() bool {
	if o != nil && o.ConnectionOptions.IsSet() {
		return true
	}

	return false
}

// SetConnectionOptions gets a reference to the given NullablePayoutConnectionOptionsRequest and assigns it to the ConnectionOptions field.
func (o *PayoutRequest) SetConnectionOptions(v PayoutConnectionOptionsRequest) {
	o.ConnectionOptions.Set(&v)
}
// SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil
func (o *PayoutRequest) SetConnectionOptionsNil() {
	o.ConnectionOptions.Set(nil)
}

// UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
func (o *PayoutRequest) UnsetConnectionOptions() {
	o.ConnectionOptions.Unset()
}

func (o PayoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["payment_service_id"] = o.PaymentServiceId
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.BuyerId.IsSet() {
		toSerialize["buyer_id"] = o.BuyerId.Get()
	}
	if o.BuyerExternalIdentifier.IsSet() {
		toSerialize["buyer_external_identifier"] = o.BuyerExternalIdentifier.Get()
	}
	if o.Buyer.IsSet() {
		toSerialize["buyer"] = o.Buyer.Get()
	}
	if o.Merchant.IsSet() {
		toSerialize["merchant"] = o.Merchant.Get()
	}
	toSerialize["payment_method"] = o.PaymentMethod
	if o.ConnectionOptions.IsSet() {
		toSerialize["connection_options"] = o.ConnectionOptions.Get()
	}
	return toSerialize, nil
}

func (o *PayoutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"payment_service_id",
		"payment_method",
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

	varPayoutRequest := _PayoutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPayoutRequest)

	if err != nil {
		return err
	}

	*o = PayoutRequest(varPayoutRequest)

	return err
}

type NullablePayoutRequest struct {
	value *PayoutRequest
	isSet bool
}

func (v NullablePayoutRequest) Get() *PayoutRequest {
	return v.value
}

func (v *NullablePayoutRequest) Set(val *PayoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutRequest(val *PayoutRequest) *NullablePayoutRequest {
	return &NullablePayoutRequest{value: val, isSet: true}
}

func (v NullablePayoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


