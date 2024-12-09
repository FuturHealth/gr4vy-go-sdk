# ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonalDetails** | [**ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPersonalDetails**](ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPersonalDetails.md) |  | 
**Address** | Pointer to [**NullableConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerAddress**](ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerAddress.md) |  | [optional] 
**Phone** | Pointer to [**[]ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPhoneInner**](ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPhoneInner.md) | List of all phone numbers for the beneficiary. | [optional] 
**Comments** | Pointer to [**NullableConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerComments**](ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerComments.md) |  | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner

`func NewConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner(personalDetails ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPersonalDetails, ) *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner`

NewConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner instantiates a new ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerWithDefaults

`func NewConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerWithDefaults() *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner`

NewConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerWithDefaults instantiates a new ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonalDetails

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetPersonalDetails() ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPersonalDetails`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetPersonalDetailsOk() (*ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPersonalDetails, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) SetPersonalDetails(v ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPersonalDetails)`

SetPersonalDetails sets PersonalDetails field to given value.


### GetAddress

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetAddress() ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetAddressOk() (*ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) SetAddress(v ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPhone

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetPhone() []ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPhoneInner`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetPhoneOk() (*[]ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPhoneInner, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) SetPhone(v []ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerPhoneInner)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetComments

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetComments() ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) GetCommentsOk() (*ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) SetComments(v ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInnerComments)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ConnectionOptionsForterAntiFraudCartItemsInnerBeneficiariesInner) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


