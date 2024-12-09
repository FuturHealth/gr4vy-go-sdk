# AirlinePassenger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Title of the passenger. | [optional] 
**FirstName** | Pointer to **NullableString** | The first name(s) or given name of the passenger. | [optional] 
**LastName** | Pointer to **NullableString** | The last name, or family name, of the passenger. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address of the passenger. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number of the passenger. This number is formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164). | [optional] 
**PassportNumber** | Pointer to **NullableString** | The passenger&#39;s unique passport number. | [optional] 
**TicketNumber** | Pointer to **NullableString** | The ticket number for a flight. | [optional] 
**FrequentFlyerNumber** | Pointer to **NullableString** | The passenger&#39;s frequent flyer number. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The passenger&#39;s date of birth. | [optional] 
**AgeGroup** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAirlinePassenger

`func NewAirlinePassenger() *AirlinePassenger`

NewAirlinePassenger instantiates a new AirlinePassenger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlinePassengerWithDefaults

`func NewAirlinePassengerWithDefaults() *AirlinePassenger`

NewAirlinePassengerWithDefaults instantiates a new AirlinePassenger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AirlinePassenger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AirlinePassenger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AirlinePassenger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AirlinePassenger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AirlinePassenger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AirlinePassenger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFirstName

`func (o *AirlinePassenger) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AirlinePassenger) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AirlinePassenger) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AirlinePassenger) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *AirlinePassenger) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *AirlinePassenger) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *AirlinePassenger) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AirlinePassenger) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AirlinePassenger) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AirlinePassenger) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *AirlinePassenger) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *AirlinePassenger) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmailAddress

`func (o *AirlinePassenger) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AirlinePassenger) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AirlinePassenger) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *AirlinePassenger) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *AirlinePassenger) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *AirlinePassenger) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetPhoneNumber

`func (o *AirlinePassenger) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AirlinePassenger) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AirlinePassenger) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AirlinePassenger) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *AirlinePassenger) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *AirlinePassenger) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPassportNumber

`func (o *AirlinePassenger) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *AirlinePassenger) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *AirlinePassenger) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *AirlinePassenger) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### SetPassportNumberNil

`func (o *AirlinePassenger) SetPassportNumberNil(b bool)`

 SetPassportNumberNil sets the value for PassportNumber to be an explicit nil

### UnsetPassportNumber
`func (o *AirlinePassenger) UnsetPassportNumber()`

UnsetPassportNumber ensures that no value is present for PassportNumber, not even an explicit nil
### GetTicketNumber

`func (o *AirlinePassenger) GetTicketNumber() string`

GetTicketNumber returns the TicketNumber field if non-nil, zero value otherwise.

### GetTicketNumberOk

`func (o *AirlinePassenger) GetTicketNumberOk() (*string, bool)`

GetTicketNumberOk returns a tuple with the TicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumber

`func (o *AirlinePassenger) SetTicketNumber(v string)`

SetTicketNumber sets TicketNumber field to given value.

### HasTicketNumber

`func (o *AirlinePassenger) HasTicketNumber() bool`

HasTicketNumber returns a boolean if a field has been set.

### SetTicketNumberNil

`func (o *AirlinePassenger) SetTicketNumberNil(b bool)`

 SetTicketNumberNil sets the value for TicketNumber to be an explicit nil

### UnsetTicketNumber
`func (o *AirlinePassenger) UnsetTicketNumber()`

UnsetTicketNumber ensures that no value is present for TicketNumber, not even an explicit nil
### GetFrequentFlyerNumber

`func (o *AirlinePassenger) GetFrequentFlyerNumber() string`

GetFrequentFlyerNumber returns the FrequentFlyerNumber field if non-nil, zero value otherwise.

### GetFrequentFlyerNumberOk

`func (o *AirlinePassenger) GetFrequentFlyerNumberOk() (*string, bool)`

GetFrequentFlyerNumberOk returns a tuple with the FrequentFlyerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentFlyerNumber

`func (o *AirlinePassenger) SetFrequentFlyerNumber(v string)`

SetFrequentFlyerNumber sets FrequentFlyerNumber field to given value.

### HasFrequentFlyerNumber

`func (o *AirlinePassenger) HasFrequentFlyerNumber() bool`

HasFrequentFlyerNumber returns a boolean if a field has been set.

### SetFrequentFlyerNumberNil

`func (o *AirlinePassenger) SetFrequentFlyerNumberNil(b bool)`

 SetFrequentFlyerNumberNil sets the value for FrequentFlyerNumber to be an explicit nil

### UnsetFrequentFlyerNumber
`func (o *AirlinePassenger) UnsetFrequentFlyerNumber()`

UnsetFrequentFlyerNumber ensures that no value is present for FrequentFlyerNumber, not even an explicit nil
### GetDateOfBirth

`func (o *AirlinePassenger) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *AirlinePassenger) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *AirlinePassenger) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *AirlinePassenger) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *AirlinePassenger) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *AirlinePassenger) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetAgeGroup

`func (o *AirlinePassenger) GetAgeGroup() string`

GetAgeGroup returns the AgeGroup field if non-nil, zero value otherwise.

### GetAgeGroupOk

`func (o *AirlinePassenger) GetAgeGroupOk() (*string, bool)`

GetAgeGroupOk returns a tuple with the AgeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGroup

`func (o *AirlinePassenger) SetAgeGroup(v string)`

SetAgeGroup sets AgeGroup field to given value.

### HasAgeGroup

`func (o *AirlinePassenger) HasAgeGroup() bool`

HasAgeGroup returns a boolean if a field has been set.

### SetAgeGroupNil

`func (o *AirlinePassenger) SetAgeGroupNil(b bool)`

 SetAgeGroupNil sets the value for AgeGroup to be an explicit nil

### UnsetAgeGroup
`func (o *AirlinePassenger) UnsetAgeGroup()`

UnsetAgeGroup ensures that no value is present for AgeGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


