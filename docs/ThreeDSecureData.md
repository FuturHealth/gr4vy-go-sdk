# ThreeDSecureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cavv** | **string** | The cardholder authentication value or AAV. | 
**Eci** | **string** | The electronic commerce indicator for the 3DS transaction. | 
**Version** | **string** | The version of 3-D Secure that was used. | 
**DirectoryResponse** | **string** | For 3-D Secure version 1, the enrolment response. For 3-D Secure version , the transaction status from the &#x60;ARes&#x60;. | 
**Scheme** | Pointer to **NullableString** | The scheme/brand of the card that is used for 3-D Secure. | [optional] 

## Methods

### NewThreeDSecureData

`func NewThreeDSecureData(cavv string, eci string, version string, directoryResponse string, ) *ThreeDSecureData`

NewThreeDSecureData instantiates a new ThreeDSecureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureDataWithDefaults

`func NewThreeDSecureDataWithDefaults() *ThreeDSecureData`

NewThreeDSecureDataWithDefaults instantiates a new ThreeDSecureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCavv

`func (o *ThreeDSecureData) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureData) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureData) SetCavv(v string)`

SetCavv sets Cavv field to given value.


### GetEci

`func (o *ThreeDSecureData) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureData) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureData) SetEci(v string)`

SetEci sets Eci field to given value.


### GetVersion

`func (o *ThreeDSecureData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDirectoryResponse

`func (o *ThreeDSecureData) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureData) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureData) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.


### GetScheme

`func (o *ThreeDSecureData) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ThreeDSecureData) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ThreeDSecureData) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ThreeDSecureData) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *ThreeDSecureData) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *ThreeDSecureData) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


