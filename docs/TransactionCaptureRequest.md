# TransactionCaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The monetary amount to capture an authorization for, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;.  When omitted blank, this will capture the entire amount. | [optional] 
**Airline** | Pointer to [**NullableAirline**](Airline.md) | The airline addendum data which describes the airline booking associated with this transaction. When provided, this will override any airline data provided when authorizing the transaction. Only the data on this request will be passed to the payment service, and none of the original data will be sent or kept. | [optional] 

## Methods

### NewTransactionCaptureRequest

`func NewTransactionCaptureRequest() *TransactionCaptureRequest`

NewTransactionCaptureRequest instantiates a new TransactionCaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCaptureRequestWithDefaults

`func NewTransactionCaptureRequestWithDefaults() *TransactionCaptureRequest`

NewTransactionCaptureRequestWithDefaults instantiates a new TransactionCaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionCaptureRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionCaptureRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionCaptureRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionCaptureRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAirline

`func (o *TransactionCaptureRequest) GetAirline() Airline`

GetAirline returns the Airline field if non-nil, zero value otherwise.

### GetAirlineOk

`func (o *TransactionCaptureRequest) GetAirlineOk() (*Airline, bool)`

GetAirlineOk returns a tuple with the Airline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirline

`func (o *TransactionCaptureRequest) SetAirline(v Airline)`

SetAirline sets Airline field to given value.

### HasAirline

`func (o *TransactionCaptureRequest) HasAirline() bool`

HasAirline returns a boolean if a field has been set.

### SetAirlineNil

`func (o *TransactionCaptureRequest) SetAirlineNil(b bool)`

 SetAirlineNil sets the value for Airline to be an explicit nil

### UnsetAirline
`func (o *TransactionCaptureRequest) UnsetAirline()`

UnsetAirline ensures that no value is present for Airline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


