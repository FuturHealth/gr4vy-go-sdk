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

// checks if the AntiFraudServiceCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntiFraudServiceCreate{}

// AntiFraudServiceCreate A request to create an anti-fraud service.
type AntiFraudServiceCreate struct {
	// The name of the Anti-Fraud service provider. During update request, this value is used for validation only but the underlying service can not be changed for an existing service.
	AntiFraudServiceDefinitionId string `json:"anti_fraud_service_definition_id"`
	// A unique name for this anti-fraud service which is used in the Gr4vy admin panel to give a anti-fraud Service a human readable name.
	DisplayName string `json:"display_name"`
	// Defines if this service is currently active or not. There can only be one active service at any time. When updating a service to active, the current active service will be deactivated.
	Active *bool `json:"active,omitempty"`
	// Defines if this service needs to handle the review status from anti-fraud responses with a proper review workflow. If not, the review status will be treated as any other one.
	ReviewsEnabled *bool `json:"reviews_enabled,omitempty"`
	// A list of fields, each containing a key-value pair for each field defined by the definition for this anti-fraud service e.g. for Sift `api_key` must be sent within this field when creating the service.  For updates, only the fields sent here will be updated, existing ones will not be affected if not present.
	Fields []AntiFraudServiceUpdateFieldsInner `json:"fields"`
}

type _AntiFraudServiceCreate AntiFraudServiceCreate

// NewAntiFraudServiceCreate instantiates a new AntiFraudServiceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntiFraudServiceCreate(antiFraudServiceDefinitionId string, displayName string, fields []AntiFraudServiceUpdateFieldsInner) *AntiFraudServiceCreate {
	this := AntiFraudServiceCreate{}
	this.AntiFraudServiceDefinitionId = antiFraudServiceDefinitionId
	this.DisplayName = displayName
	var active bool = true
	this.Active = &active
	var reviewsEnabled bool = false
	this.ReviewsEnabled = &reviewsEnabled
	this.Fields = fields
	return &this
}

// NewAntiFraudServiceCreateWithDefaults instantiates a new AntiFraudServiceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntiFraudServiceCreateWithDefaults() *AntiFraudServiceCreate {
	this := AntiFraudServiceCreate{}
	var active bool = true
	this.Active = &active
	var reviewsEnabled bool = false
	this.ReviewsEnabled = &reviewsEnabled
	return &this
}

// GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field value
func (o *AntiFraudServiceCreate) GetAntiFraudServiceDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AntiFraudServiceDefinitionId
}

// GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field value
// and a boolean to check if the value has been set.
func (o *AntiFraudServiceCreate) GetAntiFraudServiceDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AntiFraudServiceDefinitionId, true
}

// SetAntiFraudServiceDefinitionId sets field value
func (o *AntiFraudServiceCreate) SetAntiFraudServiceDefinitionId(v string) {
	o.AntiFraudServiceDefinitionId = v
}

// GetDisplayName returns the DisplayName field value
func (o *AntiFraudServiceCreate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AntiFraudServiceCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AntiFraudServiceCreate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AntiFraudServiceCreate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudServiceCreate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AntiFraudServiceCreate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AntiFraudServiceCreate) SetActive(v bool) {
	o.Active = &v
}

// GetReviewsEnabled returns the ReviewsEnabled field value if set, zero value otherwise.
func (o *AntiFraudServiceCreate) GetReviewsEnabled() bool {
	if o == nil || IsNil(o.ReviewsEnabled) {
		var ret bool
		return ret
	}
	return *o.ReviewsEnabled
}

// GetReviewsEnabledOk returns a tuple with the ReviewsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudServiceCreate) GetReviewsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ReviewsEnabled) {
		return nil, false
	}
	return o.ReviewsEnabled, true
}

// HasReviewsEnabled returns a boolean if a field has been set.
func (o *AntiFraudServiceCreate) HasReviewsEnabled() bool {
	if o != nil && !IsNil(o.ReviewsEnabled) {
		return true
	}

	return false
}

// SetReviewsEnabled gets a reference to the given bool and assigns it to the ReviewsEnabled field.
func (o *AntiFraudServiceCreate) SetReviewsEnabled(v bool) {
	o.ReviewsEnabled = &v
}

// GetFields returns the Fields field value
func (o *AntiFraudServiceCreate) GetFields() []AntiFraudServiceUpdateFieldsInner {
	if o == nil {
		var ret []AntiFraudServiceUpdateFieldsInner
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *AntiFraudServiceCreate) GetFieldsOk() ([]AntiFraudServiceUpdateFieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *AntiFraudServiceCreate) SetFields(v []AntiFraudServiceUpdateFieldsInner) {
	o.Fields = v
}

func (o AntiFraudServiceCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntiFraudServiceCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["anti_fraud_service_definition_id"] = o.AntiFraudServiceDefinitionId
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ReviewsEnabled) {
		toSerialize["reviews_enabled"] = o.ReviewsEnabled
	}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *AntiFraudServiceCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"anti_fraud_service_definition_id",
		"display_name",
		"fields",
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

	varAntiFraudServiceCreate := _AntiFraudServiceCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAntiFraudServiceCreate)

	if err != nil {
		return err
	}

	*o = AntiFraudServiceCreate(varAntiFraudServiceCreate)

	return err
}

type NullableAntiFraudServiceCreate struct {
	value *AntiFraudServiceCreate
	isSet bool
}

func (v NullableAntiFraudServiceCreate) Get() *AntiFraudServiceCreate {
	return v.value
}

func (v *NullableAntiFraudServiceCreate) Set(val *AntiFraudServiceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAntiFraudServiceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAntiFraudServiceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntiFraudServiceCreate(val *AntiFraudServiceCreate) *NullableAntiFraudServiceCreate {
	return &NullableAntiFraudServiceCreate{value: val, isSet: true}
}

func (v NullableAntiFraudServiceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntiFraudServiceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


