# MerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the merchant. | 
**IdentificationNumber** | **string** | Unique value which identifies a merchant for processing transactions, also known as a MID. | 
**PhoneNumber** | **string** | The phone number for the merchant which should be formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164). | 
**Url** | **string** | Merchant website URL. | 
**StatementDescriptor** | **string** | Value to explain charges or payments on bank statements. | 
**MerchantCategoryCode** | **string** | Merchant classification for the type of goods or services it provides. | 
**Address** | Pointer to [**NullableAddress**](Address.md) | The address for the merchant. | [optional] 

## Methods

### NewMerchantRequest

`func NewMerchantRequest(name string, identificationNumber string, phoneNumber string, url string, statementDescriptor string, merchantCategoryCode string, ) *MerchantRequest`

NewMerchantRequest instantiates a new MerchantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantRequestWithDefaults

`func NewMerchantRequestWithDefaults() *MerchantRequest`

NewMerchantRequestWithDefaults instantiates a new MerchantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIdentificationNumber

`func (o *MerchantRequest) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *MerchantRequest) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *MerchantRequest) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.


### GetPhoneNumber

`func (o *MerchantRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MerchantRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MerchantRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetUrl

`func (o *MerchantRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MerchantRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MerchantRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatementDescriptor

`func (o *MerchantRequest) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *MerchantRequest) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *MerchantRequest) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.


### GetMerchantCategoryCode

`func (o *MerchantRequest) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *MerchantRequest) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *MerchantRequest) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.


### GetAddress

`func (o *MerchantRequest) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MerchantRequest) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MerchantRequest) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MerchantRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MerchantRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MerchantRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


