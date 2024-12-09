# MerchantProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amex** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for Amex. | [optional] 
**Dankort** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for Dankort. | [optional] 
**Discover** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for Discover. | [optional] 
**Jcb** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for JCB. | [optional] 
**Mastercard** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for Mastercard. | [optional] 
**Unionpay** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for UnionPay. | [optional] 
**Visa** | Pointer to [**NullableMerchantProfileScheme**](MerchantProfileScheme.md) | Merchant profile for Visa. | [optional] 

## Methods

### NewMerchantProfile

`func NewMerchantProfile() *MerchantProfile`

NewMerchantProfile instantiates a new MerchantProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProfileWithDefaults

`func NewMerchantProfileWithDefaults() *MerchantProfile`

NewMerchantProfileWithDefaults instantiates a new MerchantProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmex

`func (o *MerchantProfile) GetAmex() MerchantProfileScheme`

GetAmex returns the Amex field if non-nil, zero value otherwise.

### GetAmexOk

`func (o *MerchantProfile) GetAmexOk() (*MerchantProfileScheme, bool)`

GetAmexOk returns a tuple with the Amex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmex

`func (o *MerchantProfile) SetAmex(v MerchantProfileScheme)`

SetAmex sets Amex field to given value.

### HasAmex

`func (o *MerchantProfile) HasAmex() bool`

HasAmex returns a boolean if a field has been set.

### SetAmexNil

`func (o *MerchantProfile) SetAmexNil(b bool)`

 SetAmexNil sets the value for Amex to be an explicit nil

### UnsetAmex
`func (o *MerchantProfile) UnsetAmex()`

UnsetAmex ensures that no value is present for Amex, not even an explicit nil
### GetDankort

`func (o *MerchantProfile) GetDankort() MerchantProfileScheme`

GetDankort returns the Dankort field if non-nil, zero value otherwise.

### GetDankortOk

`func (o *MerchantProfile) GetDankortOk() (*MerchantProfileScheme, bool)`

GetDankortOk returns a tuple with the Dankort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDankort

`func (o *MerchantProfile) SetDankort(v MerchantProfileScheme)`

SetDankort sets Dankort field to given value.

### HasDankort

`func (o *MerchantProfile) HasDankort() bool`

HasDankort returns a boolean if a field has been set.

### SetDankortNil

`func (o *MerchantProfile) SetDankortNil(b bool)`

 SetDankortNil sets the value for Dankort to be an explicit nil

### UnsetDankort
`func (o *MerchantProfile) UnsetDankort()`

UnsetDankort ensures that no value is present for Dankort, not even an explicit nil
### GetDiscover

`func (o *MerchantProfile) GetDiscover() MerchantProfileScheme`

GetDiscover returns the Discover field if non-nil, zero value otherwise.

### GetDiscoverOk

`func (o *MerchantProfile) GetDiscoverOk() (*MerchantProfileScheme, bool)`

GetDiscoverOk returns a tuple with the Discover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscover

`func (o *MerchantProfile) SetDiscover(v MerchantProfileScheme)`

SetDiscover sets Discover field to given value.

### HasDiscover

`func (o *MerchantProfile) HasDiscover() bool`

HasDiscover returns a boolean if a field has been set.

### SetDiscoverNil

`func (o *MerchantProfile) SetDiscoverNil(b bool)`

 SetDiscoverNil sets the value for Discover to be an explicit nil

### UnsetDiscover
`func (o *MerchantProfile) UnsetDiscover()`

UnsetDiscover ensures that no value is present for Discover, not even an explicit nil
### GetJcb

`func (o *MerchantProfile) GetJcb() MerchantProfileScheme`

GetJcb returns the Jcb field if non-nil, zero value otherwise.

### GetJcbOk

`func (o *MerchantProfile) GetJcbOk() (*MerchantProfileScheme, bool)`

GetJcbOk returns a tuple with the Jcb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJcb

`func (o *MerchantProfile) SetJcb(v MerchantProfileScheme)`

SetJcb sets Jcb field to given value.

### HasJcb

`func (o *MerchantProfile) HasJcb() bool`

HasJcb returns a boolean if a field has been set.

### SetJcbNil

`func (o *MerchantProfile) SetJcbNil(b bool)`

 SetJcbNil sets the value for Jcb to be an explicit nil

### UnsetJcb
`func (o *MerchantProfile) UnsetJcb()`

UnsetJcb ensures that no value is present for Jcb, not even an explicit nil
### GetMastercard

`func (o *MerchantProfile) GetMastercard() MerchantProfileScheme`

GetMastercard returns the Mastercard field if non-nil, zero value otherwise.

### GetMastercardOk

`func (o *MerchantProfile) GetMastercardOk() (*MerchantProfileScheme, bool)`

GetMastercardOk returns a tuple with the Mastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercard

`func (o *MerchantProfile) SetMastercard(v MerchantProfileScheme)`

SetMastercard sets Mastercard field to given value.

### HasMastercard

`func (o *MerchantProfile) HasMastercard() bool`

HasMastercard returns a boolean if a field has been set.

### SetMastercardNil

`func (o *MerchantProfile) SetMastercardNil(b bool)`

 SetMastercardNil sets the value for Mastercard to be an explicit nil

### UnsetMastercard
`func (o *MerchantProfile) UnsetMastercard()`

UnsetMastercard ensures that no value is present for Mastercard, not even an explicit nil
### GetUnionpay

`func (o *MerchantProfile) GetUnionpay() MerchantProfileScheme`

GetUnionpay returns the Unionpay field if non-nil, zero value otherwise.

### GetUnionpayOk

`func (o *MerchantProfile) GetUnionpayOk() (*MerchantProfileScheme, bool)`

GetUnionpayOk returns a tuple with the Unionpay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionpay

`func (o *MerchantProfile) SetUnionpay(v MerchantProfileScheme)`

SetUnionpay sets Unionpay field to given value.

### HasUnionpay

`func (o *MerchantProfile) HasUnionpay() bool`

HasUnionpay returns a boolean if a field has been set.

### SetUnionpayNil

`func (o *MerchantProfile) SetUnionpayNil(b bool)`

 SetUnionpayNil sets the value for Unionpay to be an explicit nil

### UnsetUnionpay
`func (o *MerchantProfile) UnsetUnionpay()`

UnsetUnionpay ensures that no value is present for Unionpay, not even an explicit nil
### GetVisa

`func (o *MerchantProfile) GetVisa() MerchantProfileScheme`

GetVisa returns the Visa field if non-nil, zero value otherwise.

### GetVisaOk

`func (o *MerchantProfile) GetVisaOk() (*MerchantProfileScheme, bool)`

GetVisaOk returns a tuple with the Visa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisa

`func (o *MerchantProfile) SetVisa(v MerchantProfileScheme)`

SetVisa sets Visa field to given value.

### HasVisa

`func (o *MerchantProfile) HasVisa() bool`

HasVisa returns a boolean if a field has been set.

### SetVisaNil

`func (o *MerchantProfile) SetVisaNil(b bool)`

 SetVisaNil sets the value for Visa to be an explicit nil

### UnsetVisa
`func (o *MerchantProfile) UnsetVisa()`

UnsetVisa ensures that no value is present for Visa, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


