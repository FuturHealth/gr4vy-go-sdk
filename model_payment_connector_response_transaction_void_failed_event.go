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
	"time"
)

// checks if the PaymentConnectorResponseTransactionVoidFailedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentConnectorResponseTransactionVoidFailedEvent{}

// PaymentConnectorResponseTransactionVoidFailedEvent This event logs the exact details parsed details for a failed void as reported by our connector.
type PaymentConnectorResponseTransactionVoidFailedEvent struct {
	// The type of this resource. Is always `transaction-event`.
	Type *string `json:"type,omitempty"`
	// The unique identifier for this event.
	Id *string `json:"id,omitempty"`
	// The name of this resource. Is always `payment-connector-response-transaction-void-failed`.
	Name *string `json:"name,omitempty"`
	// The date and time when this event was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Context *PaymentConnectorResponseTransactionVoidFailedEventContext `json:"context,omitempty"`
}

// NewPaymentConnectorResponseTransactionVoidFailedEvent instantiates a new PaymentConnectorResponseTransactionVoidFailedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConnectorResponseTransactionVoidFailedEvent() *PaymentConnectorResponseTransactionVoidFailedEvent {
	this := PaymentConnectorResponseTransactionVoidFailedEvent{}
	return &this
}

// NewPaymentConnectorResponseTransactionVoidFailedEventWithDefaults instantiates a new PaymentConnectorResponseTransactionVoidFailedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConnectorResponseTransactionVoidFailedEventWithDefaults() *PaymentConnectorResponseTransactionVoidFailedEvent {
	this := PaymentConnectorResponseTransactionVoidFailedEvent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetContext() PaymentConnectorResponseTransactionVoidFailedEventContext {
	if o == nil || IsNil(o.Context) {
		var ret PaymentConnectorResponseTransactionVoidFailedEventContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) GetContextOk() (*PaymentConnectorResponseTransactionVoidFailedEventContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given PaymentConnectorResponseTransactionVoidFailedEventContext and assigns it to the Context field.
func (o *PaymentConnectorResponseTransactionVoidFailedEvent) SetContext(v PaymentConnectorResponseTransactionVoidFailedEventContext) {
	o.Context = &v
}

func (o PaymentConnectorResponseTransactionVoidFailedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentConnectorResponseTransactionVoidFailedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

type NullablePaymentConnectorResponseTransactionVoidFailedEvent struct {
	value *PaymentConnectorResponseTransactionVoidFailedEvent
	isSet bool
}

func (v NullablePaymentConnectorResponseTransactionVoidFailedEvent) Get() *PaymentConnectorResponseTransactionVoidFailedEvent {
	return v.value
}

func (v *NullablePaymentConnectorResponseTransactionVoidFailedEvent) Set(val *PaymentConnectorResponseTransactionVoidFailedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConnectorResponseTransactionVoidFailedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConnectorResponseTransactionVoidFailedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConnectorResponseTransactionVoidFailedEvent(val *PaymentConnectorResponseTransactionVoidFailedEvent) *NullablePaymentConnectorResponseTransactionVoidFailedEvent {
	return &NullablePaymentConnectorResponseTransactionVoidFailedEvent{value: val, isSet: true}
}

func (v NullablePaymentConnectorResponseTransactionVoidFailedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConnectorResponseTransactionVoidFailedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


