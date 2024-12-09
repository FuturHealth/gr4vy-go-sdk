# GiftCardRedemptionFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;gift-card-redemption-failed&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this event was created in our system. | [optional] 
**Context** | Pointer to [**GiftCardRedemptionFailedContext**](GiftCardRedemptionFailedContext.md) |  | [optional] 

## Methods

### NewGiftCardRedemptionFailed

`func NewGiftCardRedemptionFailed() *GiftCardRedemptionFailed`

NewGiftCardRedemptionFailed instantiates a new GiftCardRedemptionFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRedemptionFailedWithDefaults

`func NewGiftCardRedemptionFailedWithDefaults() *GiftCardRedemptionFailed`

NewGiftCardRedemptionFailedWithDefaults instantiates a new GiftCardRedemptionFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardRedemptionFailed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardRedemptionFailed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardRedemptionFailed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardRedemptionFailed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCardRedemptionFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardRedemptionFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardRedemptionFailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardRedemptionFailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GiftCardRedemptionFailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GiftCardRedemptionFailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GiftCardRedemptionFailed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GiftCardRedemptionFailed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GiftCardRedemptionFailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GiftCardRedemptionFailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GiftCardRedemptionFailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GiftCardRedemptionFailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *GiftCardRedemptionFailed) GetContext() GiftCardRedemptionFailedContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GiftCardRedemptionFailed) GetContextOk() (*GiftCardRedemptionFailedContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GiftCardRedemptionFailed) SetContext(v GiftCardRedemptionFailedContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GiftCardRedemptionFailed) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


