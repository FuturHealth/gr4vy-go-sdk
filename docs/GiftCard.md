# GiftCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **string** | The ID of this gift card. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Service** | Pointer to [**GiftCardService**](GiftCardService.md) |  | [optional] 
**Bin** | Pointer to **string** | The first 6 digits of the full gift card number. | [optional] 
**SubBin** | Pointer to **string** | The 3 digits after the &#x60;bin&#x60; of the full gift card number. | [optional] 
**Last4** | Pointer to **string** | The last 4 digits for the gift card. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | The date and time when this gift card expires. This is a full date/time and may be more accurate than the actual expiry date received by the gift card service. | [optional] 
**Buyer** | Pointer to [**NullableBuyer**](Buyer.md) | The optional buyer for which this payment method has been stored. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this gift card was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this gift card was last updated in our system. | [optional] 

## Methods

### NewGiftCard

`func NewGiftCard() *GiftCard`

NewGiftCard instantiates a new GiftCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardWithDefaults

`func NewGiftCardWithDefaults() *GiftCard`

NewGiftCardWithDefaults instantiates a new GiftCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *GiftCard) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *GiftCard) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *GiftCard) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *GiftCard) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetService

`func (o *GiftCard) GetService() GiftCardService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GiftCard) GetServiceOk() (*GiftCardService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GiftCard) SetService(v GiftCardService)`

SetService sets Service field to given value.

### HasService

`func (o *GiftCard) HasService() bool`

HasService returns a boolean if a field has been set.

### GetBin

`func (o *GiftCard) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *GiftCard) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *GiftCard) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *GiftCard) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetSubBin

`func (o *GiftCard) GetSubBin() string`

GetSubBin returns the SubBin field if non-nil, zero value otherwise.

### GetSubBinOk

`func (o *GiftCard) GetSubBinOk() (*string, bool)`

GetSubBinOk returns a tuple with the SubBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBin

`func (o *GiftCard) SetSubBin(v string)`

SetSubBin sets SubBin field to given value.

### HasSubBin

`func (o *GiftCard) HasSubBin() bool`

HasSubBin returns a boolean if a field has been set.

### GetLast4

`func (o *GiftCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *GiftCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *GiftCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *GiftCard) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GiftCard) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GiftCard) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GiftCard) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GiftCard) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *GiftCard) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *GiftCard) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetBuyer

`func (o *GiftCard) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *GiftCard) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *GiftCard) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *GiftCard) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *GiftCard) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *GiftCard) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetCreatedAt

`func (o *GiftCard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GiftCard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GiftCard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GiftCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GiftCard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GiftCard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GiftCard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GiftCard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


