# BillingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;billing-details&#x60;. | [optional] 
**FirstName** | Pointer to **NullableString** | The first name(s) or given name of the buyer. | [optional] 
**LastName** | Pointer to **NullableString** | The last name, or family name, of the buyer. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address of the buyer. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the buyer. This number is formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164). | [optional] 
**Address** | Pointer to [**NullableAddress**](Address.md) | The billing address of the buyer. | [optional] 
**TaxId** | Pointer to [**NullableTaxId**](TaxId.md) | The tax information associated with the billing details. | [optional] 

## Methods

### NewBillingDetails

`func NewBillingDetails() *BillingDetails`

NewBillingDetails instantiates a new BillingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDetailsWithDefaults

`func NewBillingDetailsWithDefaults() *BillingDetails`

NewBillingDetailsWithDefaults instantiates a new BillingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFirstName

`func (o *BillingDetails) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BillingDetails) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BillingDetails) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BillingDetails) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *BillingDetails) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *BillingDetails) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *BillingDetails) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BillingDetails) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BillingDetails) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BillingDetails) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *BillingDetails) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *BillingDetails) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmailAddress

`func (o *BillingDetails) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *BillingDetails) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *BillingDetails) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *BillingDetails) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *BillingDetails) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *BillingDetails) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetPhoneNumber

`func (o *BillingDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BillingDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BillingDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BillingDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *BillingDetails) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *BillingDetails) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetAddress

`func (o *BillingDetails) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingDetails) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingDetails) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BillingDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *BillingDetails) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *BillingDetails) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetTaxId

`func (o *BillingDetails) GetTaxId() TaxId`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BillingDetails) GetTaxIdOk() (*TaxId, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BillingDetails) SetTaxId(v TaxId)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *BillingDetails) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *BillingDetails) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *BillingDetails) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


