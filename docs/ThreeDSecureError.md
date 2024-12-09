# ThreeDSecureError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | The error description. | 
**Detail** | **NullableString** | Detail for the error. | 
**Code** | **NullableString** | The error code. | 
**Component** | **NullableString** | Code indicating the 3-D Secure component that identified the error.. | 

## Methods

### NewThreeDSecureError

`func NewThreeDSecureError(description NullableString, detail NullableString, code NullableString, component NullableString, ) *ThreeDSecureError`

NewThreeDSecureError instantiates a new ThreeDSecureError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureErrorWithDefaults

`func NewThreeDSecureErrorWithDefaults() *ThreeDSecureError`

NewThreeDSecureErrorWithDefaults instantiates a new ThreeDSecureError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ThreeDSecureError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreeDSecureError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreeDSecureError) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ThreeDSecureError) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ThreeDSecureError) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDetail

`func (o *ThreeDSecureError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ThreeDSecureError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ThreeDSecureError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### SetDetailNil

`func (o *ThreeDSecureError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ThreeDSecureError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetCode

`func (o *ThreeDSecureError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ThreeDSecureError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ThreeDSecureError) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *ThreeDSecureError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ThreeDSecureError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetComponent

`func (o *ThreeDSecureError) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ThreeDSecureError) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ThreeDSecureError) SetComponent(v string)`

SetComponent sets Component field to given value.


### SetComponentNil

`func (o *ThreeDSecureError) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *ThreeDSecureError) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


