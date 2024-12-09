# GiftCardSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **NullableString** | The ID of this gift card. This may be &#x60;null&#x60; if the gift card is not stored. | [optional] 
**Bin** | Pointer to **string** | The first 6 digits of the full gift card number. | [optional] 
**SubBin** | Pointer to **string** | The 3 digits after the &#x60;bin&#x60; of the full gift card number. | [optional] 
**Last4** | Pointer to **string** | The last 4 digits for the gift card. | [optional] 

## Methods

### NewGiftCardSnapshot

`func NewGiftCardSnapshot() *GiftCardSnapshot`

NewGiftCardSnapshot instantiates a new GiftCardSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardSnapshotWithDefaults

`func NewGiftCardSnapshotWithDefaults() *GiftCardSnapshot`

NewGiftCardSnapshotWithDefaults instantiates a new GiftCardSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCardSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GiftCardSnapshot) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GiftCardSnapshot) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetBin

`func (o *GiftCardSnapshot) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *GiftCardSnapshot) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *GiftCardSnapshot) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *GiftCardSnapshot) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetSubBin

`func (o *GiftCardSnapshot) GetSubBin() string`

GetSubBin returns the SubBin field if non-nil, zero value otherwise.

### GetSubBinOk

`func (o *GiftCardSnapshot) GetSubBinOk() (*string, bool)`

GetSubBinOk returns a tuple with the SubBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBin

`func (o *GiftCardSnapshot) SetSubBin(v string)`

SetSubBin sets SubBin field to given value.

### HasSubBin

`func (o *GiftCardSnapshot) HasSubBin() bool`

HasSubBin returns a boolean if a field has been set.

### GetLast4

`func (o *GiftCardSnapshot) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *GiftCardSnapshot) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *GiftCardSnapshot) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *GiftCardSnapshot) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


