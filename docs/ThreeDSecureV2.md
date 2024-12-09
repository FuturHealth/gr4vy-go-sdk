# ThreeDSecureV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cavv** | Pointer to **NullableString** | The cardholder authentication value or AAV. | [optional] 
**Eci** | Pointer to **NullableString** | The electronic commerce indicator for the 3DS transaction. | [optional] 
**Version** | Pointer to **string** | The version of 3-D Secure that was used. | [optional] 
**AuthenticationResponse** | Pointer to **NullableString** | The transaction status after a the 3DS challenge. This will be null in case of a frictionless 3DS flow. | [optional] 
**DirectoryResponse** | Pointer to **NullableString** | The transaction status received as part of the authentication request. | [optional] 
**DirectoryTransactionId** | Pointer to **NullableString** | The transaction identifier. | [optional] 
**TransactionReason** | Pointer to **NullableString** | The reason identifier for a declined transaction. | [optional] 

## Methods

### NewThreeDSecureV2

`func NewThreeDSecureV2() *ThreeDSecureV2`

NewThreeDSecureV2 instantiates a new ThreeDSecureV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureV2WithDefaults

`func NewThreeDSecureV2WithDefaults() *ThreeDSecureV2`

NewThreeDSecureV2WithDefaults instantiates a new ThreeDSecureV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCavv

`func (o *ThreeDSecureV2) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureV2) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureV2) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ThreeDSecureV2) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### SetCavvNil

`func (o *ThreeDSecureV2) SetCavvNil(b bool)`

 SetCavvNil sets the value for Cavv to be an explicit nil

### UnsetCavv
`func (o *ThreeDSecureV2) UnsetCavv()`

UnsetCavv ensures that no value is present for Cavv, not even an explicit nil
### GetEci

`func (o *ThreeDSecureV2) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureV2) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureV2) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDSecureV2) HasEci() bool`

HasEci returns a boolean if a field has been set.

### SetEciNil

`func (o *ThreeDSecureV2) SetEciNil(b bool)`

 SetEciNil sets the value for Eci to be an explicit nil

### UnsetEci
`func (o *ThreeDSecureV2) UnsetEci()`

UnsetEci ensures that no value is present for Eci, not even an explicit nil
### GetVersion

`func (o *ThreeDSecureV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureV2) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSecureV2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAuthenticationResponse

`func (o *ThreeDSecureV2) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureV2) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureV2) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.

### HasAuthenticationResponse

`func (o *ThreeDSecureV2) HasAuthenticationResponse() bool`

HasAuthenticationResponse returns a boolean if a field has been set.

### SetAuthenticationResponseNil

`func (o *ThreeDSecureV2) SetAuthenticationResponseNil(b bool)`

 SetAuthenticationResponseNil sets the value for AuthenticationResponse to be an explicit nil

### UnsetAuthenticationResponse
`func (o *ThreeDSecureV2) UnsetAuthenticationResponse()`

UnsetAuthenticationResponse ensures that no value is present for AuthenticationResponse, not even an explicit nil
### GetDirectoryResponse

`func (o *ThreeDSecureV2) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureV2) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureV2) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.

### HasDirectoryResponse

`func (o *ThreeDSecureV2) HasDirectoryResponse() bool`

HasDirectoryResponse returns a boolean if a field has been set.

### SetDirectoryResponseNil

`func (o *ThreeDSecureV2) SetDirectoryResponseNil(b bool)`

 SetDirectoryResponseNil sets the value for DirectoryResponse to be an explicit nil

### UnsetDirectoryResponse
`func (o *ThreeDSecureV2) UnsetDirectoryResponse()`

UnsetDirectoryResponse ensures that no value is present for DirectoryResponse, not even an explicit nil
### GetDirectoryTransactionId

`func (o *ThreeDSecureV2) GetDirectoryTransactionId() string`

GetDirectoryTransactionId returns the DirectoryTransactionId field if non-nil, zero value otherwise.

### GetDirectoryTransactionIdOk

`func (o *ThreeDSecureV2) GetDirectoryTransactionIdOk() (*string, bool)`

GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTransactionId

`func (o *ThreeDSecureV2) SetDirectoryTransactionId(v string)`

SetDirectoryTransactionId sets DirectoryTransactionId field to given value.

### HasDirectoryTransactionId

`func (o *ThreeDSecureV2) HasDirectoryTransactionId() bool`

HasDirectoryTransactionId returns a boolean if a field has been set.

### SetDirectoryTransactionIdNil

`func (o *ThreeDSecureV2) SetDirectoryTransactionIdNil(b bool)`

 SetDirectoryTransactionIdNil sets the value for DirectoryTransactionId to be an explicit nil

### UnsetDirectoryTransactionId
`func (o *ThreeDSecureV2) UnsetDirectoryTransactionId()`

UnsetDirectoryTransactionId ensures that no value is present for DirectoryTransactionId, not even an explicit nil
### GetTransactionReason

`func (o *ThreeDSecureV2) GetTransactionReason() string`

GetTransactionReason returns the TransactionReason field if non-nil, zero value otherwise.

### GetTransactionReasonOk

`func (o *ThreeDSecureV2) GetTransactionReasonOk() (*string, bool)`

GetTransactionReasonOk returns a tuple with the TransactionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReason

`func (o *ThreeDSecureV2) SetTransactionReason(v string)`

SetTransactionReason sets TransactionReason field to given value.

### HasTransactionReason

`func (o *ThreeDSecureV2) HasTransactionReason() bool`

HasTransactionReason returns a boolean if a field has been set.

### SetTransactionReasonNil

`func (o *ThreeDSecureV2) SetTransactionReasonNil(b bool)`

 SetTransactionReasonNil sets the value for TransactionReason to be an explicit nil

### UnsetTransactionReason
`func (o *ThreeDSecureV2) UnsetTransactionReason()`

UnsetTransactionReason ensures that no value is present for TransactionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


