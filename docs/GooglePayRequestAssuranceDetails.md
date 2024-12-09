# GooglePayRequestAssuranceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountVerified** | Pointer to **NullableBool** | Indicates that card holder possession validation has been performed. | [optional] 
**CardHolderAuthenticated** | Pointer to **NullableBool** | Indicates that identification and verifications was performed. | [optional] 

## Methods

### NewGooglePayRequestAssuranceDetails

`func NewGooglePayRequestAssuranceDetails() *GooglePayRequestAssuranceDetails`

NewGooglePayRequestAssuranceDetails instantiates a new GooglePayRequestAssuranceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayRequestAssuranceDetailsWithDefaults

`func NewGooglePayRequestAssuranceDetailsWithDefaults() *GooglePayRequestAssuranceDetails`

NewGooglePayRequestAssuranceDetailsWithDefaults instantiates a new GooglePayRequestAssuranceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountVerified

`func (o *GooglePayRequestAssuranceDetails) GetAccountVerified() bool`

GetAccountVerified returns the AccountVerified field if non-nil, zero value otherwise.

### GetAccountVerifiedOk

`func (o *GooglePayRequestAssuranceDetails) GetAccountVerifiedOk() (*bool, bool)`

GetAccountVerifiedOk returns a tuple with the AccountVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVerified

`func (o *GooglePayRequestAssuranceDetails) SetAccountVerified(v bool)`

SetAccountVerified sets AccountVerified field to given value.

### HasAccountVerified

`func (o *GooglePayRequestAssuranceDetails) HasAccountVerified() bool`

HasAccountVerified returns a boolean if a field has been set.

### SetAccountVerifiedNil

`func (o *GooglePayRequestAssuranceDetails) SetAccountVerifiedNil(b bool)`

 SetAccountVerifiedNil sets the value for AccountVerified to be an explicit nil

### UnsetAccountVerified
`func (o *GooglePayRequestAssuranceDetails) UnsetAccountVerified()`

UnsetAccountVerified ensures that no value is present for AccountVerified, not even an explicit nil
### GetCardHolderAuthenticated

`func (o *GooglePayRequestAssuranceDetails) GetCardHolderAuthenticated() bool`

GetCardHolderAuthenticated returns the CardHolderAuthenticated field if non-nil, zero value otherwise.

### GetCardHolderAuthenticatedOk

`func (o *GooglePayRequestAssuranceDetails) GetCardHolderAuthenticatedOk() (*bool, bool)`

GetCardHolderAuthenticatedOk returns a tuple with the CardHolderAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderAuthenticated

`func (o *GooglePayRequestAssuranceDetails) SetCardHolderAuthenticated(v bool)`

SetCardHolderAuthenticated sets CardHolderAuthenticated field to given value.

### HasCardHolderAuthenticated

`func (o *GooglePayRequestAssuranceDetails) HasCardHolderAuthenticated() bool`

HasCardHolderAuthenticated returns a boolean if a field has been set.

### SetCardHolderAuthenticatedNil

`func (o *GooglePayRequestAssuranceDetails) SetCardHolderAuthenticatedNil(b bool)`

 SetCardHolderAuthenticatedNil sets the value for CardHolderAuthenticated to be an explicit nil

### UnsetCardHolderAuthenticated
`func (o *GooglePayRequestAssuranceDetails) UnsetCardHolderAuthenticated()`

UnsetCardHolderAuthenticated ensures that no value is present for CardHolderAuthenticated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


