# GiftCardRefundFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;gift-card-refund-failed&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this event was created in our system. | [optional] 
**Context** | Pointer to [**GiftCardRefundFailedContext**](GiftCardRefundFailedContext.md) |  | [optional] 

## Methods

### NewGiftCardRefundFailed

`func NewGiftCardRefundFailed() *GiftCardRefundFailed`

NewGiftCardRefundFailed instantiates a new GiftCardRefundFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRefundFailedWithDefaults

`func NewGiftCardRefundFailedWithDefaults() *GiftCardRefundFailed`

NewGiftCardRefundFailedWithDefaults instantiates a new GiftCardRefundFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardRefundFailed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardRefundFailed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardRefundFailed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardRefundFailed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCardRefundFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardRefundFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardRefundFailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardRefundFailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GiftCardRefundFailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GiftCardRefundFailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GiftCardRefundFailed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GiftCardRefundFailed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GiftCardRefundFailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GiftCardRefundFailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GiftCardRefundFailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GiftCardRefundFailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *GiftCardRefundFailed) GetContext() GiftCardRefundFailedContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GiftCardRefundFailed) GetContextOk() (*GiftCardRefundFailedContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GiftCardRefundFailed) SetContext(v GiftCardRefundFailedContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GiftCardRefundFailed) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


