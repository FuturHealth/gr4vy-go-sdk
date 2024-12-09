# RequiredFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **bool** | The first (given) name of the buyer. | [optional] [readonly] 
**LastName** | Pointer to **bool** | The last (family) name of the buyer. | [optional] [readonly] 
**EmailAddress** | Pointer to **bool** | The email address of the buyer. | [optional] [readonly] 
**PhoneNumber** | Pointer to **bool** | The phone number of the buyer. | [optional] [readonly] 
**Address** | Pointer to [**RequiredFieldsAddress**](RequiredFieldsAddress.md) |  | [optional] 
**TaxId** | Pointer to **bool** | The tax id code associated with the billing details. | [optional] [readonly] 

## Methods

### NewRequiredFields

`func NewRequiredFields() *RequiredFields`

NewRequiredFields instantiates a new RequiredFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredFieldsWithDefaults

`func NewRequiredFieldsWithDefaults() *RequiredFields`

NewRequiredFieldsWithDefaults instantiates a new RequiredFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *RequiredFields) GetFirstName() bool`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RequiredFields) GetFirstNameOk() (*bool, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RequiredFields) SetFirstName(v bool)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RequiredFields) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *RequiredFields) GetLastName() bool`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RequiredFields) GetLastNameOk() (*bool, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RequiredFields) SetLastName(v bool)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RequiredFields) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *RequiredFields) GetEmailAddress() bool`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *RequiredFields) GetEmailAddressOk() (*bool, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *RequiredFields) SetEmailAddress(v bool)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *RequiredFields) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *RequiredFields) GetPhoneNumber() bool`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RequiredFields) GetPhoneNumberOk() (*bool, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RequiredFields) SetPhoneNumber(v bool)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RequiredFields) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAddress

`func (o *RequiredFields) GetAddress() RequiredFieldsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RequiredFields) GetAddressOk() (*RequiredFieldsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RequiredFields) SetAddress(v RequiredFieldsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RequiredFields) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTaxId

`func (o *RequiredFields) GetTaxId() bool`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *RequiredFields) GetTaxIdOk() (*bool, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *RequiredFields) SetTaxId(v bool)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *RequiredFields) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


