# BillingDetailsRequestAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **NullableString** | The city for the address. | 
**Country** | **NullableString** | The country for the address in ISO 3166 format. | 
**PostalCode** | **NullableString** | The postal code or zip code for the address. | 
**State** | **NullableString** | The state, county, or province for the address. | 
**StateCode** | Pointer to **NullableString** | The code of state, county, or province for the address in ISO 3166-2 format. | [optional] 
**HouseNumberOrName** | Pointer to **NullableString** | The house number or name for the address. Not all payment services use this field but some do. | [optional] 
**Line1** | **NullableString** | The first line of the address. | 
**Line2** | Pointer to **NullableString** | The second line of the address. | [optional] 
**Organization** | Pointer to **NullableString** | The optional name of the company or organisation to add to the address. | [optional] 

## Methods

### NewBillingDetailsRequestAddress

`func NewBillingDetailsRequestAddress(city NullableString, country NullableString, postalCode NullableString, state NullableString, line1 NullableString, ) *BillingDetailsRequestAddress`

NewBillingDetailsRequestAddress instantiates a new BillingDetailsRequestAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDetailsRequestAddressWithDefaults

`func NewBillingDetailsRequestAddressWithDefaults() *BillingDetailsRequestAddress`

NewBillingDetailsRequestAddressWithDefaults instantiates a new BillingDetailsRequestAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *BillingDetailsRequestAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingDetailsRequestAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingDetailsRequestAddress) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *BillingDetailsRequestAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *BillingDetailsRequestAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *BillingDetailsRequestAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingDetailsRequestAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingDetailsRequestAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *BillingDetailsRequestAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *BillingDetailsRequestAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPostalCode

`func (o *BillingDetailsRequestAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingDetailsRequestAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingDetailsRequestAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *BillingDetailsRequestAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *BillingDetailsRequestAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetState

`func (o *BillingDetailsRequestAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingDetailsRequestAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingDetailsRequestAddress) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *BillingDetailsRequestAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BillingDetailsRequestAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStateCode

`func (o *BillingDetailsRequestAddress) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *BillingDetailsRequestAddress) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *BillingDetailsRequestAddress) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *BillingDetailsRequestAddress) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### SetStateCodeNil

`func (o *BillingDetailsRequestAddress) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *BillingDetailsRequestAddress) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetHouseNumberOrName

`func (o *BillingDetailsRequestAddress) GetHouseNumberOrName() string`

GetHouseNumberOrName returns the HouseNumberOrName field if non-nil, zero value otherwise.

### GetHouseNumberOrNameOk

`func (o *BillingDetailsRequestAddress) GetHouseNumberOrNameOk() (*string, bool)`

GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberOrName

`func (o *BillingDetailsRequestAddress) SetHouseNumberOrName(v string)`

SetHouseNumberOrName sets HouseNumberOrName field to given value.

### HasHouseNumberOrName

`func (o *BillingDetailsRequestAddress) HasHouseNumberOrName() bool`

HasHouseNumberOrName returns a boolean if a field has been set.

### SetHouseNumberOrNameNil

`func (o *BillingDetailsRequestAddress) SetHouseNumberOrNameNil(b bool)`

 SetHouseNumberOrNameNil sets the value for HouseNumberOrName to be an explicit nil

### UnsetHouseNumberOrName
`func (o *BillingDetailsRequestAddress) UnsetHouseNumberOrName()`

UnsetHouseNumberOrName ensures that no value is present for HouseNumberOrName, not even an explicit nil
### GetLine1

`func (o *BillingDetailsRequestAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *BillingDetailsRequestAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *BillingDetailsRequestAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### SetLine1Nil

`func (o *BillingDetailsRequestAddress) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *BillingDetailsRequestAddress) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *BillingDetailsRequestAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *BillingDetailsRequestAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *BillingDetailsRequestAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *BillingDetailsRequestAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *BillingDetailsRequestAddress) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *BillingDetailsRequestAddress) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetOrganization

`func (o *BillingDetailsRequestAddress) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BillingDetailsRequestAddress) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BillingDetailsRequestAddress) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BillingDetailsRequestAddress) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *BillingDetailsRequestAddress) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *BillingDetailsRequestAddress) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


