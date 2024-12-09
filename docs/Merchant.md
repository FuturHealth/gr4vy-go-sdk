# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;merchant&#x60;. | [optional] 
**Name** | Pointer to **string** | The name of the merchant. | [optional] 
**IdentificationNumber** | Pointer to **string** | Unique value which identifies a merchant for processing transactions, also known as a MID. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number for the merchant which should be formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164). | [optional] 
**Url** | Pointer to **string** | Merchant website URL. | [optional] 
**StatementDescriptor** | Pointer to **string** | Value to explain charges or payments on bank statements. | [optional] 
**MerchantCategoryCode** | Pointer to **string** | Merchant classification for the type of goods or services it provides. | [optional] 
**Address** | Pointer to [**NullableAddress**](Address.md) | The address for the merchant. | [optional] 

## Methods

### NewMerchant

`func NewMerchant() *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Merchant) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Merchant) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Merchant) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Merchant) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Merchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentificationNumber

`func (o *Merchant) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *Merchant) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *Merchant) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *Merchant) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Merchant) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Merchant) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Merchant) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Merchant) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetUrl

`func (o *Merchant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Merchant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Merchant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Merchant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *Merchant) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *Merchant) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *Merchant) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *Merchant) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *Merchant) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *Merchant) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *Merchant) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *Merchant) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetAddress

`func (o *Merchant) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Merchant) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Merchant) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Merchant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Merchant) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Merchant) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


