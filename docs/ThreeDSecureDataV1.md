# ThreeDSecureDataV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cavv** | **string** | The cardholder authentication value or AAV. | 
**Eci** | **string** | The electronic commerce indicator for the 3DS transaction. | 
**Version** | **string** | The version of 3-D Secure that was used. | 
**DirectoryResponse** | **string** | For 3-D Secure version 1, the enrolment response. For 3-D Secure version , the transaction status from the &#x60;ARes&#x60;. | 
**Scheme** | Pointer to **NullableString** | The scheme/brand of the card that is used for 3-D Secure. | [optional] 
**AuthenticationResponse** | **string** | The response for the 3DS authentication call. | 
**CavvAlgorithm** | **string** | The CAVV algorithm used. | 
**Xid** | **string** | The transaction identifier. | 

## Methods

### NewThreeDSecureDataV1

`func NewThreeDSecureDataV1(cavv string, eci string, version string, directoryResponse string, authenticationResponse string, cavvAlgorithm string, xid string, ) *ThreeDSecureDataV1`

NewThreeDSecureDataV1 instantiates a new ThreeDSecureDataV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureDataV1WithDefaults

`func NewThreeDSecureDataV1WithDefaults() *ThreeDSecureDataV1`

NewThreeDSecureDataV1WithDefaults instantiates a new ThreeDSecureDataV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCavv

`func (o *ThreeDSecureDataV1) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureDataV1) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureDataV1) SetCavv(v string)`

SetCavv sets Cavv field to given value.


### GetEci

`func (o *ThreeDSecureDataV1) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureDataV1) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureDataV1) SetEci(v string)`

SetEci sets Eci field to given value.


### GetVersion

`func (o *ThreeDSecureDataV1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureDataV1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureDataV1) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDirectoryResponse

`func (o *ThreeDSecureDataV1) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureDataV1) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureDataV1) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.


### GetScheme

`func (o *ThreeDSecureDataV1) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ThreeDSecureDataV1) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ThreeDSecureDataV1) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ThreeDSecureDataV1) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *ThreeDSecureDataV1) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *ThreeDSecureDataV1) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetAuthenticationResponse

`func (o *ThreeDSecureDataV1) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureDataV1) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureDataV1) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.


### GetCavvAlgorithm

`func (o *ThreeDSecureDataV1) GetCavvAlgorithm() string`

GetCavvAlgorithm returns the CavvAlgorithm field if non-nil, zero value otherwise.

### GetCavvAlgorithmOk

`func (o *ThreeDSecureDataV1) GetCavvAlgorithmOk() (*string, bool)`

GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAlgorithm

`func (o *ThreeDSecureDataV1) SetCavvAlgorithm(v string)`

SetCavvAlgorithm sets CavvAlgorithm field to given value.


### GetXid

`func (o *ThreeDSecureDataV1) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *ThreeDSecureDataV1) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *ThreeDSecureDataV1) SetXid(v string)`

SetXid sets Xid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


