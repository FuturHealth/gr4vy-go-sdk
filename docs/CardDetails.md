# CardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;card-detail&#x60;. | [optional] 
**Id** | Pointer to **string** | The 8 digit BIN of the card. When looking up card details using a &#x60;payment_method_id&#x60; this value will be &#x60;null&#x60;. | [optional] 
**CardType** | Pointer to **string** | The type of card. | [optional] 
**Scheme** | Pointer to **NullableString** | The scheme/brand of the card. | [optional] 
**SchemeIconUrl** | Pointer to **string** | An icon to display for the card scheme. | [optional] 
**Country** | Pointer to **string** | The 2-letter ISO code of the issuing country of the card. | [optional] 
**RequiredFields** | Pointer to [**RequiredFields**](RequiredFields.md) |  | [optional] 

## Methods

### NewCardDetails

`func NewCardDetails() *CardDetails`

NewCardDetails instantiates a new CardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsWithDefaults

`func NewCardDetailsWithDefaults() *CardDetails`

NewCardDetailsWithDefaults instantiates a new CardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CardDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CardDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CardDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCardType

`func (o *CardDetails) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardDetails) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardDetails) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CardDetails) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetScheme

`func (o *CardDetails) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *CardDetails) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *CardDetails) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *CardDetails) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *CardDetails) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *CardDetails) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetSchemeIconUrl

`func (o *CardDetails) GetSchemeIconUrl() string`

GetSchemeIconUrl returns the SchemeIconUrl field if non-nil, zero value otherwise.

### GetSchemeIconUrlOk

`func (o *CardDetails) GetSchemeIconUrlOk() (*string, bool)`

GetSchemeIconUrlOk returns a tuple with the SchemeIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeIconUrl

`func (o *CardDetails) SetSchemeIconUrl(v string)`

SetSchemeIconUrl sets SchemeIconUrl field to given value.

### HasSchemeIconUrl

`func (o *CardDetails) HasSchemeIconUrl() bool`

HasSchemeIconUrl returns a boolean if a field has been set.

### GetCountry

`func (o *CardDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CardDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CardDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CardDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRequiredFields

`func (o *CardDetails) GetRequiredFields() RequiredFields`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *CardDetails) GetRequiredFieldsOk() (*RequiredFields, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *CardDetails) SetRequiredFields(v RequiredFields)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *CardDetails) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


