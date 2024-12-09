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
)

// checks if the BillingDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingDetailsRequest{}

// BillingDetailsRequest Billing details to use associated to a buyer.
type BillingDetailsRequest struct {
	// The first name(s) or given name for the buyer.
	FirstName NullableString `json:"first_name,omitempty"`
	// The last name, or family name, of the buyer.
	LastName NullableString `json:"last_name,omitempty"`
	// The email address for the buyer.
	EmailAddress NullableString `json:"email_address,omitempty"`
	// The phone number for the buyer which should be formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164).
	PhoneNumber NullableString `json:"phone_number,omitempty" validate:"regexp=^\\\\+[1-9]\\\\d{1,14}$"`
	Address NullableBillingDetailsRequestAddress `json:"address,omitempty"`
	// The tax ID information associated with the billing details.
	TaxId NullableTaxId `json:"tax_id,omitempty"`
}

// NewBillingDetailsRequest instantiates a new BillingDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingDetailsRequest() *BillingDetailsRequest {
	this := BillingDetailsRequest{}
	return &this
}

// NewBillingDetailsRequestWithDefaults instantiates a new BillingDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingDetailsRequestWithDefaults() *BillingDetailsRequest {
	this := BillingDetailsRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetailsRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *BillingDetailsRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *BillingDetailsRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *BillingDetailsRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *BillingDetailsRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetailsRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *BillingDetailsRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *BillingDetailsRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *BillingDetailsRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *BillingDetailsRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetailsRequest) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress.Get()) {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *BillingDetailsRequest) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *BillingDetailsRequest) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *BillingDetailsRequest) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *BillingDetailsRequest) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetailsRequest) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *BillingDetailsRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *BillingDetailsRequest) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *BillingDetailsRequest) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *BillingDetailsRequest) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetailsRequest) GetAddress() BillingDetailsRequestAddress {
	if o == nil || IsNil(o.Address.Get()) {
		var ret BillingDetailsRequestAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsRequest) GetAddressOk() (*BillingDetailsRequestAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *BillingDetailsRequest) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableBillingDetailsRequestAddress and assigns it to the Address field.
func (o *BillingDetailsRequest) SetAddress(v BillingDetailsRequestAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *BillingDetailsRequest) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *BillingDetailsRequest) UnsetAddress() {
	o.Address.Unset()
}

// GetTaxId returns the TaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetailsRequest) GetTaxId() TaxId {
	if o == nil || IsNil(o.TaxId.Get()) {
		var ret TaxId
		return ret
	}
	return *o.TaxId.Get()
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsRequest) GetTaxIdOk() (*TaxId, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxId.Get(), o.TaxId.IsSet()
}

// HasTaxId returns a boolean if a field has been set.
func (o *BillingDetailsRequest) HasTaxId() bool {
	if o != nil && o.TaxId.IsSet() {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given NullableTaxId and assigns it to the TaxId field.
func (o *BillingDetailsRequest) SetTaxId(v TaxId) {
	o.TaxId.Set(&v)
}
// SetTaxIdNil sets the value for TaxId to be an explicit nil
func (o *BillingDetailsRequest) SetTaxIdNil() {
	o.TaxId.Set(nil)
}

// UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
func (o *BillingDetailsRequest) UnsetTaxId() {
	o.TaxId.Unset()
}

func (o BillingDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.TaxId.IsSet() {
		toSerialize["tax_id"] = o.TaxId.Get()
	}
	return toSerialize, nil
}

type NullableBillingDetailsRequest struct {
	value *BillingDetailsRequest
	isSet bool
}

func (v NullableBillingDetailsRequest) Get() *BillingDetailsRequest {
	return v.value
}

func (v *NullableBillingDetailsRequest) Set(val *BillingDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingDetailsRequest(val *BillingDetailsRequest) *NullableBillingDetailsRequest {
	return &NullableBillingDetailsRequest{value: val, isSet: true}
}

func (v NullableBillingDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


