# NetworkToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**Id** | Pointer to **string** | The unique ID of the token. | [optional] 
**PaymentMethodId** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**Status** | Pointer to **string** | The state of the network token.  - &#x60;active&#x60; - The network token is active and ready to be used. - &#x60;inactive&#x60; - The network token is being deactivated. - &#x60;suspended&#x60; - The network token is suspended. - &#x60;deleted&#x60; - The network token is deleted. | [optional] 
**Token** | Pointer to **string** | The value of the network token. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date for the network token. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this network token was first created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this network token was last updated in our system. | [optional] 

## Methods

### NewNetworkToken

`func NewNetworkToken() *NetworkToken`

NewNetworkToken instantiates a new NetworkToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTokenWithDefaults

`func NewNetworkTokenWithDefaults() *NetworkToken`

NewNetworkTokenWithDefaults instantiates a new NetworkToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkToken) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *NetworkToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *NetworkToken) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *NetworkToken) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *NetworkToken) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *NetworkToken) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkToken) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkToken) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkToken) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *NetworkToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NetworkToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NetworkToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *NetworkToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpirationDate

`func (o *NetworkToken) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *NetworkToken) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *NetworkToken) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *NetworkToken) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *NetworkToken) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *NetworkToken) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCreatedAt

`func (o *NetworkToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


