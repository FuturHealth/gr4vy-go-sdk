# ApiLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;api-log&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID of the API log entry. | [optional] 
**RequestMethod** | Pointer to **string** | The http request method that generated the log entry. | [optional] 
**RequestUrl** | Pointer to **string** | The http request URL which trigged the error log. | [optional] 
**RequestReceivedAt** | Pointer to **time.Time** | The date and time that the request was received. | [optional] 
**ResponseStatusCode** | Pointer to **float32** | The http request status code. | [optional] 
**ResponseBody** | Pointer to [**ApiLogResponseBody**](ApiLogResponseBody.md) |  | [optional] 
**ResponseSentAt** | Pointer to **time.Time** | date-time of when the response was sent. | [optional] 

## Methods

### NewApiLog

`func NewApiLog() *ApiLog`

NewApiLog instantiates a new ApiLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLogWithDefaults

`func NewApiLogWithDefaults() *ApiLog`

NewApiLogWithDefaults instantiates a new ApiLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiLog) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ApiLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestMethod

`func (o *ApiLog) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *ApiLog) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *ApiLog) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *ApiLog) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### GetRequestUrl

`func (o *ApiLog) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *ApiLog) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *ApiLog) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *ApiLog) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetRequestReceivedAt

`func (o *ApiLog) GetRequestReceivedAt() time.Time`

GetRequestReceivedAt returns the RequestReceivedAt field if non-nil, zero value otherwise.

### GetRequestReceivedAtOk

`func (o *ApiLog) GetRequestReceivedAtOk() (*time.Time, bool)`

GetRequestReceivedAtOk returns a tuple with the RequestReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReceivedAt

`func (o *ApiLog) SetRequestReceivedAt(v time.Time)`

SetRequestReceivedAt sets RequestReceivedAt field to given value.

### HasRequestReceivedAt

`func (o *ApiLog) HasRequestReceivedAt() bool`

HasRequestReceivedAt returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *ApiLog) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *ApiLog) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *ApiLog) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *ApiLog) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetResponseBody

`func (o *ApiLog) GetResponseBody() ApiLogResponseBody`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *ApiLog) GetResponseBodyOk() (*ApiLogResponseBody, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *ApiLog) SetResponseBody(v ApiLogResponseBody)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *ApiLog) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseSentAt

`func (o *ApiLog) GetResponseSentAt() time.Time`

GetResponseSentAt returns the ResponseSentAt field if non-nil, zero value otherwise.

### GetResponseSentAtOk

`func (o *ApiLog) GetResponseSentAtOk() (*time.Time, bool)`

GetResponseSentAtOk returns a tuple with the ResponseSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSentAt

`func (o *ApiLog) SetResponseSentAt(v time.Time)`

SetResponseSentAt sets ResponseSentAt field to given value.

### HasResponseSentAt

`func (o *ApiLog) HasResponseSentAt() bool`

HasResponseSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


