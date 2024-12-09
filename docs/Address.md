# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** | The city for the address. | [optional] 
**Country** | Pointer to **NullableString** | The country for the address in ISO 3166 format. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code or zip code for the address. | [optional] 
**State** | Pointer to **NullableString** | The state, county, or province for the address. | [optional] 
**StateCode** | Pointer to **NullableString** | The code of state, county, or province for the address in ISO 3166-2 format. | [optional] 
**HouseNumberOrName** | Pointer to **NullableString** | The house number or name for the address. Not all payment services use this field but some do. | [optional] 
**Line1** | Pointer to **NullableString** | The first line of the address. | [optional] 
**Line2** | Pointer to **NullableString** | The second line of the address. | [optional] 
**Organization** | Pointer to **NullableString** | The optional name of the company or organisation to add to the address. | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Address) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Address) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Address) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Address) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *Address) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *Address) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Address) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Address) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStateCode

`func (o *Address) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *Address) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *Address) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *Address) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### SetStateCodeNil

`func (o *Address) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *Address) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetHouseNumberOrName

`func (o *Address) GetHouseNumberOrName() string`

GetHouseNumberOrName returns the HouseNumberOrName field if non-nil, zero value otherwise.

### GetHouseNumberOrNameOk

`func (o *Address) GetHouseNumberOrNameOk() (*string, bool)`

GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberOrName

`func (o *Address) SetHouseNumberOrName(v string)`

SetHouseNumberOrName sets HouseNumberOrName field to given value.

### HasHouseNumberOrName

`func (o *Address) HasHouseNumberOrName() bool`

HasHouseNumberOrName returns a boolean if a field has been set.

### SetHouseNumberOrNameNil

`func (o *Address) SetHouseNumberOrNameNil(b bool)`

 SetHouseNumberOrNameNil sets the value for HouseNumberOrName to be an explicit nil

### UnsetHouseNumberOrName
`func (o *Address) UnsetHouseNumberOrName()`

UnsetHouseNumberOrName ensures that no value is present for HouseNumberOrName, not even an explicit nil
### GetLine1

`func (o *Address) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *Address) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *Address) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *Address) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *Address) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *Address) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *Address) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *Address) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *Address) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *Address) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *Address) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *Address) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetOrganization

`func (o *Address) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Address) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Address) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Address) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *Address) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *Address) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


