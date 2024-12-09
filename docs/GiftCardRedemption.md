# GiftCardRedemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **NullableString** | The ID of this gift card redemption. This may be &#x60;null&#x60; if the no redemption happened. | [optional] 
**Status** | Pointer to **string** | The status of the gift card redemption for the &#x60;payment_method&#x60;. | [optional] 
**Amount** | Pointer to **int32** | The amount redeemed for this gift card. | [optional] 
**RefundedAmount** | Pointer to **int32** | The amount refunded for this gift card. This can not be larger than &#x60;amount&#x60;. | [optional] 
**GiftCardServiceRedemptionId** | Pointer to **NullableString** | The gift card service&#39;s unique ID for the redemption. | [optional] 
**ErrorCode** | Pointer to **NullableString** | If this gift card redemption resulted in an error, this will contain the internal code for the error. | [optional] 
**RawErrorCode** | Pointer to **NullableString** | If this gift card redemption resulted in an error, this will contain the raw error code received from the gift card provider. | [optional] 
**RawErrorMessage** | Pointer to **NullableString** | If this gift card redemption resulted in an error, this will contain the raw error message received from the gift card provider. | [optional] 
**GiftCard** | Pointer to [**GiftCardSnapshot**](GiftCardSnapshot.md) |  | [optional] 

## Methods

### NewGiftCardRedemption

`func NewGiftCardRedemption() *GiftCardRedemption`

NewGiftCardRedemption instantiates a new GiftCardRedemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardRedemptionWithDefaults

`func NewGiftCardRedemptionWithDefaults() *GiftCardRedemption`

NewGiftCardRedemptionWithDefaults instantiates a new GiftCardRedemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardRedemption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardRedemption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardRedemption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCardRedemption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GiftCardRedemption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCardRedemption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCardRedemption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCardRedemption) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GiftCardRedemption) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GiftCardRedemption) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStatus

`func (o *GiftCardRedemption) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GiftCardRedemption) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GiftCardRedemption) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GiftCardRedemption) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAmount

`func (o *GiftCardRedemption) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GiftCardRedemption) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GiftCardRedemption) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GiftCardRedemption) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRefundedAmount

`func (o *GiftCardRedemption) GetRefundedAmount() int32`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *GiftCardRedemption) GetRefundedAmountOk() (*int32, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *GiftCardRedemption) SetRefundedAmount(v int32)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *GiftCardRedemption) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetGiftCardServiceRedemptionId

`func (o *GiftCardRedemption) GetGiftCardServiceRedemptionId() string`

GetGiftCardServiceRedemptionId returns the GiftCardServiceRedemptionId field if non-nil, zero value otherwise.

### GetGiftCardServiceRedemptionIdOk

`func (o *GiftCardRedemption) GetGiftCardServiceRedemptionIdOk() (*string, bool)`

GetGiftCardServiceRedemptionIdOk returns a tuple with the GiftCardServiceRedemptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardServiceRedemptionId

`func (o *GiftCardRedemption) SetGiftCardServiceRedemptionId(v string)`

SetGiftCardServiceRedemptionId sets GiftCardServiceRedemptionId field to given value.

### HasGiftCardServiceRedemptionId

`func (o *GiftCardRedemption) HasGiftCardServiceRedemptionId() bool`

HasGiftCardServiceRedemptionId returns a boolean if a field has been set.

### SetGiftCardServiceRedemptionIdNil

`func (o *GiftCardRedemption) SetGiftCardServiceRedemptionIdNil(b bool)`

 SetGiftCardServiceRedemptionIdNil sets the value for GiftCardServiceRedemptionId to be an explicit nil

### UnsetGiftCardServiceRedemptionId
`func (o *GiftCardRedemption) UnsetGiftCardServiceRedemptionId()`

UnsetGiftCardServiceRedemptionId ensures that no value is present for GiftCardServiceRedemptionId, not even an explicit nil
### GetErrorCode

`func (o *GiftCardRedemption) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GiftCardRedemption) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GiftCardRedemption) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GiftCardRedemption) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GiftCardRedemption) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GiftCardRedemption) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetRawErrorCode

`func (o *GiftCardRedemption) GetRawErrorCode() string`

GetRawErrorCode returns the RawErrorCode field if non-nil, zero value otherwise.

### GetRawErrorCodeOk

`func (o *GiftCardRedemption) GetRawErrorCodeOk() (*string, bool)`

GetRawErrorCodeOk returns a tuple with the RawErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawErrorCode

`func (o *GiftCardRedemption) SetRawErrorCode(v string)`

SetRawErrorCode sets RawErrorCode field to given value.

### HasRawErrorCode

`func (o *GiftCardRedemption) HasRawErrorCode() bool`

HasRawErrorCode returns a boolean if a field has been set.

### SetRawErrorCodeNil

`func (o *GiftCardRedemption) SetRawErrorCodeNil(b bool)`

 SetRawErrorCodeNil sets the value for RawErrorCode to be an explicit nil

### UnsetRawErrorCode
`func (o *GiftCardRedemption) UnsetRawErrorCode()`

UnsetRawErrorCode ensures that no value is present for RawErrorCode, not even an explicit nil
### GetRawErrorMessage

`func (o *GiftCardRedemption) GetRawErrorMessage() string`

GetRawErrorMessage returns the RawErrorMessage field if non-nil, zero value otherwise.

### GetRawErrorMessageOk

`func (o *GiftCardRedemption) GetRawErrorMessageOk() (*string, bool)`

GetRawErrorMessageOk returns a tuple with the RawErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawErrorMessage

`func (o *GiftCardRedemption) SetRawErrorMessage(v string)`

SetRawErrorMessage sets RawErrorMessage field to given value.

### HasRawErrorMessage

`func (o *GiftCardRedemption) HasRawErrorMessage() bool`

HasRawErrorMessage returns a boolean if a field has been set.

### SetRawErrorMessageNil

`func (o *GiftCardRedemption) SetRawErrorMessageNil(b bool)`

 SetRawErrorMessageNil sets the value for RawErrorMessage to be an explicit nil

### UnsetRawErrorMessage
`func (o *GiftCardRedemption) UnsetRawErrorMessage()`

UnsetRawErrorMessage ensures that no value is present for RawErrorMessage, not even an explicit nil
### GetGiftCard

`func (o *GiftCardRedemption) GetGiftCard() GiftCardSnapshot`

GetGiftCard returns the GiftCard field if non-nil, zero value otherwise.

### GetGiftCardOk

`func (o *GiftCardRedemption) GetGiftCardOk() (*GiftCardSnapshot, bool)`

GetGiftCardOk returns a tuple with the GiftCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCard

`func (o *GiftCardRedemption) SetGiftCard(v GiftCardSnapshot)`

SetGiftCard sets GiftCard field to given value.

### HasGiftCard

`func (o *GiftCardRedemption) HasGiftCard() bool`

HasGiftCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


