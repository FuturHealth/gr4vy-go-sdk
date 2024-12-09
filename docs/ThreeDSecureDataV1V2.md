# ThreeDSecureDataV1V2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cavv** | **string** | The cardholder authentication value or AAV. | 
**Eci** | **string** | The electronic commerce indicator for the 3DS transaction. | 
**Version** | **string** | The version of 3-D Secure that was used. | 
**DirectoryResponse** | **string** | The transaction status received as part of the authentication request. | 
**Scheme** | Pointer to **NullableString** | The scheme/brand of the card that is used for 3-D Secure. | [optional] 
**AuthenticationResponse** | **NullableString** | The transaction status after a the 3DS challenge. This will be null in case of a frictionless 3DS flow. | 
**CavvAlgorithm** | **string** | The CAVV algorithm used. | 
**Xid** | **string** | The transaction identifier. | 
**DirectoryTransactionId** | **string** | The transaction identifier. | 

## Methods

### NewThreeDSecureDataV1V2

`func NewThreeDSecureDataV1V2(cavv string, eci string, version string, directoryResponse string, authenticationResponse NullableString, cavvAlgorithm string, xid string, directoryTransactionId string, ) *ThreeDSecureDataV1V2`

NewThreeDSecureDataV1V2 instantiates a new ThreeDSecureDataV1V2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureDataV1V2WithDefaults

`func NewThreeDSecureDataV1V2WithDefaults() *ThreeDSecureDataV1V2`

NewThreeDSecureDataV1V2WithDefaults instantiates a new ThreeDSecureDataV1V2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCavv

`func (o *ThreeDSecureDataV1V2) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureDataV1V2) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureDataV1V2) SetCavv(v string)`

SetCavv sets Cavv field to given value.


### GetEci

`func (o *ThreeDSecureDataV1V2) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureDataV1V2) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureDataV1V2) SetEci(v string)`

SetEci sets Eci field to given value.


### GetVersion

`func (o *ThreeDSecureDataV1V2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureDataV1V2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureDataV1V2) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDirectoryResponse

`func (o *ThreeDSecureDataV1V2) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureDataV1V2) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureDataV1V2) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.


### GetScheme

`func (o *ThreeDSecureDataV1V2) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ThreeDSecureDataV1V2) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ThreeDSecureDataV1V2) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ThreeDSecureDataV1V2) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *ThreeDSecureDataV1V2) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *ThreeDSecureDataV1V2) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetAuthenticationResponse

`func (o *ThreeDSecureDataV1V2) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureDataV1V2) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureDataV1V2) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.


### SetAuthenticationResponseNil

`func (o *ThreeDSecureDataV1V2) SetAuthenticationResponseNil(b bool)`

 SetAuthenticationResponseNil sets the value for AuthenticationResponse to be an explicit nil

### UnsetAuthenticationResponse
`func (o *ThreeDSecureDataV1V2) UnsetAuthenticationResponse()`

UnsetAuthenticationResponse ensures that no value is present for AuthenticationResponse, not even an explicit nil
### GetCavvAlgorithm

`func (o *ThreeDSecureDataV1V2) GetCavvAlgorithm() string`

GetCavvAlgorithm returns the CavvAlgorithm field if non-nil, zero value otherwise.

### GetCavvAlgorithmOk

`func (o *ThreeDSecureDataV1V2) GetCavvAlgorithmOk() (*string, bool)`

GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAlgorithm

`func (o *ThreeDSecureDataV1V2) SetCavvAlgorithm(v string)`

SetCavvAlgorithm sets CavvAlgorithm field to given value.


### GetXid

`func (o *ThreeDSecureDataV1V2) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *ThreeDSecureDataV1V2) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *ThreeDSecureDataV1V2) SetXid(v string)`

SetXid sets Xid field to given value.


### GetDirectoryTransactionId

`func (o *ThreeDSecureDataV1V2) GetDirectoryTransactionId() string`

GetDirectoryTransactionId returns the DirectoryTransactionId field if non-nil, zero value otherwise.

### GetDirectoryTransactionIdOk

`func (o *ThreeDSecureDataV1V2) GetDirectoryTransactionIdOk() (*string, bool)`

GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTransactionId

`func (o *ThreeDSecureDataV1V2) SetDirectoryTransactionId(v string)`

SetDirectoryTransactionId sets DirectoryTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


