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

// checks if the NetworkTokenCryptogramProvisionSucceededContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkTokenCryptogramProvisionSucceededContext{}

// NetworkTokenCryptogramProvisionSucceededContext Additional context for this event.
type NetworkTokenCryptogramProvisionSucceededContext struct {
	// The endpoint for the request.
	Url *string `json:"url,omitempty"`
	// The HTTP body sent to the Network Token provider.
	Request *string `json:"request,omitempty"`
	// The HTTP body received from the Network Token provider.
	Response *string `json:"response,omitempty"`
	// The HTTP response status code from the Network Token provider.
	ResponseStatusCode *float32 `json:"response_status_code,omitempty"`
}

// NewNetworkTokenCryptogramProvisionSucceededContext instantiates a new NetworkTokenCryptogramProvisionSucceededContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkTokenCryptogramProvisionSucceededContext() *NetworkTokenCryptogramProvisionSucceededContext {
	this := NetworkTokenCryptogramProvisionSucceededContext{}
	return &this
}

// NewNetworkTokenCryptogramProvisionSucceededContextWithDefaults instantiates a new NetworkTokenCryptogramProvisionSucceededContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTokenCryptogramProvisionSucceededContextWithDefaults() *NetworkTokenCryptogramProvisionSucceededContext {
	this := NetworkTokenCryptogramProvisionSucceededContext{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworkTokenCryptogramProvisionSucceededContext) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetRequest() string {
	if o == nil || IsNil(o.Request) {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetRequestOk() (*string, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *NetworkTokenCryptogramProvisionSucceededContext) SetRequest(v string) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *NetworkTokenCryptogramProvisionSucceededContext) SetResponse(v string) {
	o.Response = &v
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetResponseStatusCode() float32 {
	if o == nil || IsNil(o.ResponseStatusCode) {
		var ret float32
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) GetResponseStatusCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.ResponseStatusCode) {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *NetworkTokenCryptogramProvisionSucceededContext) HasResponseStatusCode() bool {
	if o != nil && !IsNil(o.ResponseStatusCode) {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given float32 and assigns it to the ResponseStatusCode field.
func (o *NetworkTokenCryptogramProvisionSucceededContext) SetResponseStatusCode(v float32) {
	o.ResponseStatusCode = &v
}

func (o NetworkTokenCryptogramProvisionSucceededContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkTokenCryptogramProvisionSucceededContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.ResponseStatusCode) {
		toSerialize["response_status_code"] = o.ResponseStatusCode
	}
	return toSerialize, nil
}

type NullableNetworkTokenCryptogramProvisionSucceededContext struct {
	value *NetworkTokenCryptogramProvisionSucceededContext
	isSet bool
}

func (v NullableNetworkTokenCryptogramProvisionSucceededContext) Get() *NetworkTokenCryptogramProvisionSucceededContext {
	return v.value
}

func (v *NullableNetworkTokenCryptogramProvisionSucceededContext) Set(val *NetworkTokenCryptogramProvisionSucceededContext) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkTokenCryptogramProvisionSucceededContext) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkTokenCryptogramProvisionSucceededContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkTokenCryptogramProvisionSucceededContext(val *NetworkTokenCryptogramProvisionSucceededContext) *NullableNetworkTokenCryptogramProvisionSucceededContext {
	return &NullableNetworkTokenCryptogramProvisionSucceededContext{value: val, isSet: true}
}

func (v NullableNetworkTokenCryptogramProvisionSucceededContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkTokenCryptogramProvisionSucceededContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


