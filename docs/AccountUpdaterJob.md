# AccountUpdaterJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;account-updater-job&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this Account Updater job. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this Account Updater job was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this Account Updater job was last updated. | [optional] 
**Inquiries** | Pointer to [**[]AccountUpdaterInquirySummary**](AccountUpdaterInquirySummary.md) | A list of inquiries associated with this Account Updater job. | [optional] 

## Methods

### NewAccountUpdaterJob

`func NewAccountUpdaterJob() *AccountUpdaterJob`

NewAccountUpdaterJob instantiates a new AccountUpdaterJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdaterJobWithDefaults

`func NewAccountUpdaterJobWithDefaults() *AccountUpdaterJob`

NewAccountUpdaterJobWithDefaults instantiates a new AccountUpdaterJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountUpdaterJob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountUpdaterJob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountUpdaterJob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountUpdaterJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AccountUpdaterJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountUpdaterJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountUpdaterJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountUpdaterJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *AccountUpdaterJob) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *AccountUpdaterJob) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *AccountUpdaterJob) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *AccountUpdaterJob) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountUpdaterJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountUpdaterJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountUpdaterJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountUpdaterJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountUpdaterJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountUpdaterJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountUpdaterJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountUpdaterJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInquiries

`func (o *AccountUpdaterJob) GetInquiries() []AccountUpdaterInquirySummary`

GetInquiries returns the Inquiries field if non-nil, zero value otherwise.

### GetInquiriesOk

`func (o *AccountUpdaterJob) GetInquiriesOk() (*[]AccountUpdaterInquirySummary, bool)`

GetInquiriesOk returns a tuple with the Inquiries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiries

`func (o *AccountUpdaterJob) SetInquiries(v []AccountUpdaterInquirySummary)`

SetInquiries sets Inquiries field to given value.

### HasInquiries

`func (o *AccountUpdaterJob) HasInquiries() bool`

HasInquiries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


