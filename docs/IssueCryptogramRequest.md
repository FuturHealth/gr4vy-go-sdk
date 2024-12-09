# IssueCryptogramRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantInitiated** | **bool** | Defines if the request is merchant initiated or not. | [default to false]

## Methods

### NewIssueCryptogramRequest

`func NewIssueCryptogramRequest(merchantInitiated bool, ) *IssueCryptogramRequest`

NewIssueCryptogramRequest instantiates a new IssueCryptogramRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCryptogramRequestWithDefaults

`func NewIssueCryptogramRequestWithDefaults() *IssueCryptogramRequest`

NewIssueCryptogramRequestWithDefaults instantiates a new IssueCryptogramRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantInitiated

`func (o *IssueCryptogramRequest) GetMerchantInitiated() bool`

GetMerchantInitiated returns the MerchantInitiated field if non-nil, zero value otherwise.

### GetMerchantInitiatedOk

`func (o *IssueCryptogramRequest) GetMerchantInitiatedOk() (*bool, bool)`

GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInitiated

`func (o *IssueCryptogramRequest) SetMerchantInitiated(v bool)`

SetMerchantInitiated sets MerchantInitiated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


