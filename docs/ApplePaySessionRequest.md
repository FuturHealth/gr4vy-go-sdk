# ApplePaySessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationUrl** | **string** | Validation URL obtained from the event passed to a &#x60;onvalidatemerchant&#x60; callback. | 
**DomainName** | **string** | Fully qualified domain name of the merchant. | 

## Methods

### NewApplePaySessionRequest

`func NewApplePaySessionRequest(validationUrl string, domainName string, ) *ApplePaySessionRequest`

NewApplePaySessionRequest instantiates a new ApplePaySessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePaySessionRequestWithDefaults

`func NewApplePaySessionRequestWithDefaults() *ApplePaySessionRequest`

NewApplePaySessionRequestWithDefaults instantiates a new ApplePaySessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationUrl

`func (o *ApplePaySessionRequest) GetValidationUrl() string`

GetValidationUrl returns the ValidationUrl field if non-nil, zero value otherwise.

### GetValidationUrlOk

`func (o *ApplePaySessionRequest) GetValidationUrlOk() (*string, bool)`

GetValidationUrlOk returns a tuple with the ValidationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationUrl

`func (o *ApplePaySessionRequest) SetValidationUrl(v string)`

SetValidationUrl sets ValidationUrl field to given value.


### GetDomainName

`func (o *ApplePaySessionRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ApplePaySessionRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ApplePaySessionRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


