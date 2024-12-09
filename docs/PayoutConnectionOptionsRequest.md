# PayoutConnectionOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutCard** | Pointer to [**NullablePayoutConnectionOptionsRequestCheckoutCard**](PayoutConnectionOptionsRequestCheckoutCard.md) |  | [optional] 

## Methods

### NewPayoutConnectionOptionsRequest

`func NewPayoutConnectionOptionsRequest() *PayoutConnectionOptionsRequest`

NewPayoutConnectionOptionsRequest instantiates a new PayoutConnectionOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutConnectionOptionsRequestWithDefaults

`func NewPayoutConnectionOptionsRequestWithDefaults() *PayoutConnectionOptionsRequest`

NewPayoutConnectionOptionsRequestWithDefaults instantiates a new PayoutConnectionOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutCard

`func (o *PayoutConnectionOptionsRequest) GetCheckoutCard() PayoutConnectionOptionsRequestCheckoutCard`

GetCheckoutCard returns the CheckoutCard field if non-nil, zero value otherwise.

### GetCheckoutCardOk

`func (o *PayoutConnectionOptionsRequest) GetCheckoutCardOk() (*PayoutConnectionOptionsRequestCheckoutCard, bool)`

GetCheckoutCardOk returns a tuple with the CheckoutCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutCard

`func (o *PayoutConnectionOptionsRequest) SetCheckoutCard(v PayoutConnectionOptionsRequestCheckoutCard)`

SetCheckoutCard sets CheckoutCard field to given value.

### HasCheckoutCard

`func (o *PayoutConnectionOptionsRequest) HasCheckoutCard() bool`

HasCheckoutCard returns a boolean if a field has been set.

### SetCheckoutCardNil

`func (o *PayoutConnectionOptionsRequest) SetCheckoutCardNil(b bool)`

 SetCheckoutCardNil sets the value for CheckoutCard to be an explicit nil

### UnsetCheckoutCard
`func (o *PayoutConnectionOptionsRequest) UnsetCheckoutCard()`

UnsetCheckoutCard ensures that no value is present for CheckoutCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


