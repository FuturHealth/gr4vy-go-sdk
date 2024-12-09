# RequiredFieldsAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **bool** | The city for the billing address. | [optional] [readonly] 
**Country** | Pointer to **bool** | The country for the billing address. | [optional] [readonly] 
**PostalCode** | Pointer to **bool** | The postal code or zip code for the billing address. | [optional] [readonly] 
**State** | Pointer to **bool** | The state, county, or province for the billing address. | [optional] [readonly] 
**HouseNumberOrName** | Pointer to **bool** | The house number or name for the billing address. Not all payment services use this field but some do. | [optional] [readonly] 
**Line1** | Pointer to **bool** | The first line of the billing address. | [optional] [readonly] 

## Methods

### NewRequiredFieldsAddress

`func NewRequiredFieldsAddress() *RequiredFieldsAddress`

NewRequiredFieldsAddress instantiates a new RequiredFieldsAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredFieldsAddressWithDefaults

`func NewRequiredFieldsAddressWithDefaults() *RequiredFieldsAddress`

NewRequiredFieldsAddressWithDefaults instantiates a new RequiredFieldsAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *RequiredFieldsAddress) GetCity() bool`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RequiredFieldsAddress) GetCityOk() (*bool, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RequiredFieldsAddress) SetCity(v bool)`

SetCity sets City field to given value.

### HasCity

`func (o *RequiredFieldsAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *RequiredFieldsAddress) GetCountry() bool`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RequiredFieldsAddress) GetCountryOk() (*bool, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RequiredFieldsAddress) SetCountry(v bool)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RequiredFieldsAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *RequiredFieldsAddress) GetPostalCode() bool`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *RequiredFieldsAddress) GetPostalCodeOk() (*bool, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *RequiredFieldsAddress) SetPostalCode(v bool)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *RequiredFieldsAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *RequiredFieldsAddress) GetState() bool`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RequiredFieldsAddress) GetStateOk() (*bool, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RequiredFieldsAddress) SetState(v bool)`

SetState sets State field to given value.

### HasState

`func (o *RequiredFieldsAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHouseNumberOrName

`func (o *RequiredFieldsAddress) GetHouseNumberOrName() bool`

GetHouseNumberOrName returns the HouseNumberOrName field if non-nil, zero value otherwise.

### GetHouseNumberOrNameOk

`func (o *RequiredFieldsAddress) GetHouseNumberOrNameOk() (*bool, bool)`

GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberOrName

`func (o *RequiredFieldsAddress) SetHouseNumberOrName(v bool)`

SetHouseNumberOrName sets HouseNumberOrName field to given value.

### HasHouseNumberOrName

`func (o *RequiredFieldsAddress) HasHouseNumberOrName() bool`

HasHouseNumberOrName returns a boolean if a field has been set.

### GetLine1

`func (o *RequiredFieldsAddress) GetLine1() bool`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *RequiredFieldsAddress) GetLine1Ok() (*bool, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *RequiredFieldsAddress) SetLine1(v bool)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *RequiredFieldsAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


