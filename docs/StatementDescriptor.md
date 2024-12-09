# StatementDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Reflects your doing business as (DBA) name.  Other validations:  1. Contains only Latin characters. 2. Contain at least one letter 3. Does not contain any of the special characters &#x60;&lt; &gt; \\ &#39; \&quot; *&#x60; 4. Supports:   1. Lower case: &#x60;a-z&#x60;   2. Upper case: &#x60;A-Z&#x60;   3. Numbers: &#x60;0-9&#x60;   4. Spaces: &#x60; &#x60;   5. Special characters: &#x60;. , _ - ? + /&#x60;. | [optional] 
**Description** | Pointer to **NullableString** | A short description about the purchase.  Other validations: 1. Contains only Latin characters. 2. Contain at least one letter 3. Does not contain any of the special characters &#x60;&lt; &gt; \\ &#39; \&quot; *&#x60; 4. Supports:   1. Lower case: &#x60;a-z&#x60;   2. Upper case: &#x60;A-Z&#x60;   3. Numbers: &#x60;0-9&#x60;   4. Spaces: &#x60; &#x60;   5. Special characters: &#x60;. , _ - ? + /&#x60;. | [optional] 
**City** | Pointer to **NullableString** | The merchant&#39;s city to be displayed in a statement descriptor. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO country code of the merchant to be displayed in a statement descriptor.  | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The value in the phone number field of a customer&#39;s statement which should be formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164). | [optional] 
**Url** | Pointer to **NullableString** | The merchant&#39;s URL to be displayed in a statement descriptor. | [optional] 

## Methods

### NewStatementDescriptor

`func NewStatementDescriptor() *StatementDescriptor`

NewStatementDescriptor instantiates a new StatementDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementDescriptorWithDefaults

`func NewStatementDescriptorWithDefaults() *StatementDescriptor`

NewStatementDescriptorWithDefaults instantiates a new StatementDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StatementDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatementDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatementDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatementDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StatementDescriptor) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StatementDescriptor) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *StatementDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatementDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatementDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatementDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StatementDescriptor) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StatementDescriptor) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCity

`func (o *StatementDescriptor) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *StatementDescriptor) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *StatementDescriptor) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *StatementDescriptor) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *StatementDescriptor) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *StatementDescriptor) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *StatementDescriptor) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StatementDescriptor) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StatementDescriptor) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StatementDescriptor) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *StatementDescriptor) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *StatementDescriptor) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPhoneNumber

`func (o *StatementDescriptor) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *StatementDescriptor) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *StatementDescriptor) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *StatementDescriptor) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *StatementDescriptor) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *StatementDescriptor) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetUrl

`func (o *StatementDescriptor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StatementDescriptor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StatementDescriptor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StatementDescriptor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StatementDescriptor) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StatementDescriptor) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


