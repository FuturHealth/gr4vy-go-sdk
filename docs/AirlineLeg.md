# AirlineLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierCode** | Pointer to **NullableString** | 2 character airline code as set by IATA. | [optional] 
**FlightNumber** | Pointer to **NullableString** | Unique identifier of the flight number. | [optional] 
**DepartureAt** | Pointer to **NullableTime** | The date and time of travel in local time at the departure airport. | [optional] 
**DepartureCountry** | Pointer to **NullableString** | Departure country code in ISO 3166 format. | [optional] 
**DepartureCity** | Pointer to **NullableString** | Departure city name. | [optional] 
**DepartureAirport** | Pointer to **NullableString** | Departure airport code of leg. 3-letter ISO code according to IATA official directory. | [optional] 
**ArrivalAt** | Pointer to **NullableTime** | The date and time of travel in local time at the arrival airport. | [optional] 
**ArrivalCountry** | Pointer to **NullableString** | Arrival country code in ISO 3166 format. | [optional] 
**ArrivalCity** | Pointer to **NullableString** | Arrival city name. | [optional] 
**ArrivalAirport** | Pointer to **NullableString** | Arrival airport code of leg. 3-letter ISO code according to IATA official directory. | [optional] 
**FareBasisCode** | Pointer to **NullableString** | The alphanumeric code for the \&quot;booking class\&quot; of a ticket. | [optional] 
**FlightClass** | Pointer to **NullableString** | Indicates service class (first class, business class, etc.). | [optional] 
**StopOver** | Pointer to **NullableBool** | Indicates whether a stopover is allowed on this ticket. | [optional] 
**RouteType** | Pointer to **string** | The route type of the flight. | [optional] 
**CouponNumber** | Pointer to **NullableString** | Coupon number associated with the leg. | [optional] 
**FareAmount** | Pointer to **NullableInt32** | Amount of the ticket, for current leg of the trip, excluding taxes and fees. | [optional] 
**FeeAmount** | Pointer to **NullableInt32** | Fee amount for current leg of the trip. | [optional] 
**TaxAmount** | Pointer to **NullableInt32** | Amount of the taxes for current leg of the trip. | [optional] 
**DepartureTaxAmount** | Pointer to **NullableInt32** | Departure tax amount charged by a country when a person is leaving the country. | [optional] 

## Methods

### NewAirlineLeg

`func NewAirlineLeg() *AirlineLeg`

NewAirlineLeg instantiates a new AirlineLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlineLegWithDefaults

`func NewAirlineLegWithDefaults() *AirlineLeg`

NewAirlineLegWithDefaults instantiates a new AirlineLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierCode

`func (o *AirlineLeg) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *AirlineLeg) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *AirlineLeg) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *AirlineLeg) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### SetCarrierCodeNil

`func (o *AirlineLeg) SetCarrierCodeNil(b bool)`

 SetCarrierCodeNil sets the value for CarrierCode to be an explicit nil

### UnsetCarrierCode
`func (o *AirlineLeg) UnsetCarrierCode()`

UnsetCarrierCode ensures that no value is present for CarrierCode, not even an explicit nil
### GetFlightNumber

`func (o *AirlineLeg) GetFlightNumber() string`

GetFlightNumber returns the FlightNumber field if non-nil, zero value otherwise.

### GetFlightNumberOk

`func (o *AirlineLeg) GetFlightNumberOk() (*string, bool)`

GetFlightNumberOk returns a tuple with the FlightNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightNumber

`func (o *AirlineLeg) SetFlightNumber(v string)`

SetFlightNumber sets FlightNumber field to given value.

### HasFlightNumber

`func (o *AirlineLeg) HasFlightNumber() bool`

HasFlightNumber returns a boolean if a field has been set.

### SetFlightNumberNil

`func (o *AirlineLeg) SetFlightNumberNil(b bool)`

 SetFlightNumberNil sets the value for FlightNumber to be an explicit nil

### UnsetFlightNumber
`func (o *AirlineLeg) UnsetFlightNumber()`

UnsetFlightNumber ensures that no value is present for FlightNumber, not even an explicit nil
### GetDepartureAt

`func (o *AirlineLeg) GetDepartureAt() time.Time`

GetDepartureAt returns the DepartureAt field if non-nil, zero value otherwise.

### GetDepartureAtOk

`func (o *AirlineLeg) GetDepartureAtOk() (*time.Time, bool)`

GetDepartureAtOk returns a tuple with the DepartureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureAt

`func (o *AirlineLeg) SetDepartureAt(v time.Time)`

SetDepartureAt sets DepartureAt field to given value.

### HasDepartureAt

`func (o *AirlineLeg) HasDepartureAt() bool`

HasDepartureAt returns a boolean if a field has been set.

### SetDepartureAtNil

`func (o *AirlineLeg) SetDepartureAtNil(b bool)`

 SetDepartureAtNil sets the value for DepartureAt to be an explicit nil

### UnsetDepartureAt
`func (o *AirlineLeg) UnsetDepartureAt()`

UnsetDepartureAt ensures that no value is present for DepartureAt, not even an explicit nil
### GetDepartureCountry

`func (o *AirlineLeg) GetDepartureCountry() string`

GetDepartureCountry returns the DepartureCountry field if non-nil, zero value otherwise.

### GetDepartureCountryOk

`func (o *AirlineLeg) GetDepartureCountryOk() (*string, bool)`

GetDepartureCountryOk returns a tuple with the DepartureCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCountry

`func (o *AirlineLeg) SetDepartureCountry(v string)`

SetDepartureCountry sets DepartureCountry field to given value.

### HasDepartureCountry

`func (o *AirlineLeg) HasDepartureCountry() bool`

HasDepartureCountry returns a boolean if a field has been set.

### SetDepartureCountryNil

`func (o *AirlineLeg) SetDepartureCountryNil(b bool)`

 SetDepartureCountryNil sets the value for DepartureCountry to be an explicit nil

### UnsetDepartureCountry
`func (o *AirlineLeg) UnsetDepartureCountry()`

UnsetDepartureCountry ensures that no value is present for DepartureCountry, not even an explicit nil
### GetDepartureCity

`func (o *AirlineLeg) GetDepartureCity() string`

GetDepartureCity returns the DepartureCity field if non-nil, zero value otherwise.

### GetDepartureCityOk

`func (o *AirlineLeg) GetDepartureCityOk() (*string, bool)`

GetDepartureCityOk returns a tuple with the DepartureCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCity

`func (o *AirlineLeg) SetDepartureCity(v string)`

SetDepartureCity sets DepartureCity field to given value.

### HasDepartureCity

`func (o *AirlineLeg) HasDepartureCity() bool`

HasDepartureCity returns a boolean if a field has been set.

### SetDepartureCityNil

`func (o *AirlineLeg) SetDepartureCityNil(b bool)`

 SetDepartureCityNil sets the value for DepartureCity to be an explicit nil

### UnsetDepartureCity
`func (o *AirlineLeg) UnsetDepartureCity()`

UnsetDepartureCity ensures that no value is present for DepartureCity, not even an explicit nil
### GetDepartureAirport

`func (o *AirlineLeg) GetDepartureAirport() string`

GetDepartureAirport returns the DepartureAirport field if non-nil, zero value otherwise.

### GetDepartureAirportOk

`func (o *AirlineLeg) GetDepartureAirportOk() (*string, bool)`

GetDepartureAirportOk returns a tuple with the DepartureAirport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureAirport

`func (o *AirlineLeg) SetDepartureAirport(v string)`

SetDepartureAirport sets DepartureAirport field to given value.

### HasDepartureAirport

`func (o *AirlineLeg) HasDepartureAirport() bool`

HasDepartureAirport returns a boolean if a field has been set.

### SetDepartureAirportNil

`func (o *AirlineLeg) SetDepartureAirportNil(b bool)`

 SetDepartureAirportNil sets the value for DepartureAirport to be an explicit nil

### UnsetDepartureAirport
`func (o *AirlineLeg) UnsetDepartureAirport()`

UnsetDepartureAirport ensures that no value is present for DepartureAirport, not even an explicit nil
### GetArrivalAt

`func (o *AirlineLeg) GetArrivalAt() time.Time`

GetArrivalAt returns the ArrivalAt field if non-nil, zero value otherwise.

### GetArrivalAtOk

`func (o *AirlineLeg) GetArrivalAtOk() (*time.Time, bool)`

GetArrivalAtOk returns a tuple with the ArrivalAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalAt

`func (o *AirlineLeg) SetArrivalAt(v time.Time)`

SetArrivalAt sets ArrivalAt field to given value.

### HasArrivalAt

`func (o *AirlineLeg) HasArrivalAt() bool`

HasArrivalAt returns a boolean if a field has been set.

### SetArrivalAtNil

`func (o *AirlineLeg) SetArrivalAtNil(b bool)`

 SetArrivalAtNil sets the value for ArrivalAt to be an explicit nil

### UnsetArrivalAt
`func (o *AirlineLeg) UnsetArrivalAt()`

UnsetArrivalAt ensures that no value is present for ArrivalAt, not even an explicit nil
### GetArrivalCountry

`func (o *AirlineLeg) GetArrivalCountry() string`

GetArrivalCountry returns the ArrivalCountry field if non-nil, zero value otherwise.

### GetArrivalCountryOk

`func (o *AirlineLeg) GetArrivalCountryOk() (*string, bool)`

GetArrivalCountryOk returns a tuple with the ArrivalCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalCountry

`func (o *AirlineLeg) SetArrivalCountry(v string)`

SetArrivalCountry sets ArrivalCountry field to given value.

### HasArrivalCountry

`func (o *AirlineLeg) HasArrivalCountry() bool`

HasArrivalCountry returns a boolean if a field has been set.

### SetArrivalCountryNil

`func (o *AirlineLeg) SetArrivalCountryNil(b bool)`

 SetArrivalCountryNil sets the value for ArrivalCountry to be an explicit nil

### UnsetArrivalCountry
`func (o *AirlineLeg) UnsetArrivalCountry()`

UnsetArrivalCountry ensures that no value is present for ArrivalCountry, not even an explicit nil
### GetArrivalCity

`func (o *AirlineLeg) GetArrivalCity() string`

GetArrivalCity returns the ArrivalCity field if non-nil, zero value otherwise.

### GetArrivalCityOk

`func (o *AirlineLeg) GetArrivalCityOk() (*string, bool)`

GetArrivalCityOk returns a tuple with the ArrivalCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalCity

`func (o *AirlineLeg) SetArrivalCity(v string)`

SetArrivalCity sets ArrivalCity field to given value.

### HasArrivalCity

`func (o *AirlineLeg) HasArrivalCity() bool`

HasArrivalCity returns a boolean if a field has been set.

### SetArrivalCityNil

`func (o *AirlineLeg) SetArrivalCityNil(b bool)`

 SetArrivalCityNil sets the value for ArrivalCity to be an explicit nil

### UnsetArrivalCity
`func (o *AirlineLeg) UnsetArrivalCity()`

UnsetArrivalCity ensures that no value is present for ArrivalCity, not even an explicit nil
### GetArrivalAirport

`func (o *AirlineLeg) GetArrivalAirport() string`

GetArrivalAirport returns the ArrivalAirport field if non-nil, zero value otherwise.

### GetArrivalAirportOk

`func (o *AirlineLeg) GetArrivalAirportOk() (*string, bool)`

GetArrivalAirportOk returns a tuple with the ArrivalAirport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalAirport

`func (o *AirlineLeg) SetArrivalAirport(v string)`

SetArrivalAirport sets ArrivalAirport field to given value.

### HasArrivalAirport

`func (o *AirlineLeg) HasArrivalAirport() bool`

HasArrivalAirport returns a boolean if a field has been set.

### SetArrivalAirportNil

`func (o *AirlineLeg) SetArrivalAirportNil(b bool)`

 SetArrivalAirportNil sets the value for ArrivalAirport to be an explicit nil

### UnsetArrivalAirport
`func (o *AirlineLeg) UnsetArrivalAirport()`

UnsetArrivalAirport ensures that no value is present for ArrivalAirport, not even an explicit nil
### GetFareBasisCode

`func (o *AirlineLeg) GetFareBasisCode() string`

GetFareBasisCode returns the FareBasisCode field if non-nil, zero value otherwise.

### GetFareBasisCodeOk

`func (o *AirlineLeg) GetFareBasisCodeOk() (*string, bool)`

GetFareBasisCodeOk returns a tuple with the FareBasisCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareBasisCode

`func (o *AirlineLeg) SetFareBasisCode(v string)`

SetFareBasisCode sets FareBasisCode field to given value.

### HasFareBasisCode

`func (o *AirlineLeg) HasFareBasisCode() bool`

HasFareBasisCode returns a boolean if a field has been set.

### SetFareBasisCodeNil

`func (o *AirlineLeg) SetFareBasisCodeNil(b bool)`

 SetFareBasisCodeNil sets the value for FareBasisCode to be an explicit nil

### UnsetFareBasisCode
`func (o *AirlineLeg) UnsetFareBasisCode()`

UnsetFareBasisCode ensures that no value is present for FareBasisCode, not even an explicit nil
### GetFlightClass

`func (o *AirlineLeg) GetFlightClass() string`

GetFlightClass returns the FlightClass field if non-nil, zero value otherwise.

### GetFlightClassOk

`func (o *AirlineLeg) GetFlightClassOk() (*string, bool)`

GetFlightClassOk returns a tuple with the FlightClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightClass

`func (o *AirlineLeg) SetFlightClass(v string)`

SetFlightClass sets FlightClass field to given value.

### HasFlightClass

`func (o *AirlineLeg) HasFlightClass() bool`

HasFlightClass returns a boolean if a field has been set.

### SetFlightClassNil

`func (o *AirlineLeg) SetFlightClassNil(b bool)`

 SetFlightClassNil sets the value for FlightClass to be an explicit nil

### UnsetFlightClass
`func (o *AirlineLeg) UnsetFlightClass()`

UnsetFlightClass ensures that no value is present for FlightClass, not even an explicit nil
### GetStopOver

`func (o *AirlineLeg) GetStopOver() bool`

GetStopOver returns the StopOver field if non-nil, zero value otherwise.

### GetStopOverOk

`func (o *AirlineLeg) GetStopOverOk() (*bool, bool)`

GetStopOverOk returns a tuple with the StopOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOver

`func (o *AirlineLeg) SetStopOver(v bool)`

SetStopOver sets StopOver field to given value.

### HasStopOver

`func (o *AirlineLeg) HasStopOver() bool`

HasStopOver returns a boolean if a field has been set.

### SetStopOverNil

`func (o *AirlineLeg) SetStopOverNil(b bool)`

 SetStopOverNil sets the value for StopOver to be an explicit nil

### UnsetStopOver
`func (o *AirlineLeg) UnsetStopOver()`

UnsetStopOver ensures that no value is present for StopOver, not even an explicit nil
### GetRouteType

`func (o *AirlineLeg) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *AirlineLeg) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *AirlineLeg) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.

### HasRouteType

`func (o *AirlineLeg) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.

### GetCouponNumber

`func (o *AirlineLeg) GetCouponNumber() string`

GetCouponNumber returns the CouponNumber field if non-nil, zero value otherwise.

### GetCouponNumberOk

`func (o *AirlineLeg) GetCouponNumberOk() (*string, bool)`

GetCouponNumberOk returns a tuple with the CouponNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponNumber

`func (o *AirlineLeg) SetCouponNumber(v string)`

SetCouponNumber sets CouponNumber field to given value.

### HasCouponNumber

`func (o *AirlineLeg) HasCouponNumber() bool`

HasCouponNumber returns a boolean if a field has been set.

### SetCouponNumberNil

`func (o *AirlineLeg) SetCouponNumberNil(b bool)`

 SetCouponNumberNil sets the value for CouponNumber to be an explicit nil

### UnsetCouponNumber
`func (o *AirlineLeg) UnsetCouponNumber()`

UnsetCouponNumber ensures that no value is present for CouponNumber, not even an explicit nil
### GetFareAmount

`func (o *AirlineLeg) GetFareAmount() int32`

GetFareAmount returns the FareAmount field if non-nil, zero value otherwise.

### GetFareAmountOk

`func (o *AirlineLeg) GetFareAmountOk() (*int32, bool)`

GetFareAmountOk returns a tuple with the FareAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareAmount

`func (o *AirlineLeg) SetFareAmount(v int32)`

SetFareAmount sets FareAmount field to given value.

### HasFareAmount

`func (o *AirlineLeg) HasFareAmount() bool`

HasFareAmount returns a boolean if a field has been set.

### SetFareAmountNil

`func (o *AirlineLeg) SetFareAmountNil(b bool)`

 SetFareAmountNil sets the value for FareAmount to be an explicit nil

### UnsetFareAmount
`func (o *AirlineLeg) UnsetFareAmount()`

UnsetFareAmount ensures that no value is present for FareAmount, not even an explicit nil
### GetFeeAmount

`func (o *AirlineLeg) GetFeeAmount() int32`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *AirlineLeg) GetFeeAmountOk() (*int32, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *AirlineLeg) SetFeeAmount(v int32)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *AirlineLeg) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### SetFeeAmountNil

`func (o *AirlineLeg) SetFeeAmountNil(b bool)`

 SetFeeAmountNil sets the value for FeeAmount to be an explicit nil

### UnsetFeeAmount
`func (o *AirlineLeg) UnsetFeeAmount()`

UnsetFeeAmount ensures that no value is present for FeeAmount, not even an explicit nil
### GetTaxAmount

`func (o *AirlineLeg) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *AirlineLeg) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *AirlineLeg) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *AirlineLeg) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### SetTaxAmountNil

`func (o *AirlineLeg) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *AirlineLeg) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetDepartureTaxAmount

`func (o *AirlineLeg) GetDepartureTaxAmount() int32`

GetDepartureTaxAmount returns the DepartureTaxAmount field if non-nil, zero value otherwise.

### GetDepartureTaxAmountOk

`func (o *AirlineLeg) GetDepartureTaxAmountOk() (*int32, bool)`

GetDepartureTaxAmountOk returns a tuple with the DepartureTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTaxAmount

`func (o *AirlineLeg) SetDepartureTaxAmount(v int32)`

SetDepartureTaxAmount sets DepartureTaxAmount field to given value.

### HasDepartureTaxAmount

`func (o *AirlineLeg) HasDepartureTaxAmount() bool`

HasDepartureTaxAmount returns a boolean if a field has been set.

### SetDepartureTaxAmountNil

`func (o *AirlineLeg) SetDepartureTaxAmountNil(b bool)`

 SetDepartureTaxAmountNil sets the value for DepartureTaxAmount to be an explicit nil

### UnsetDepartureTaxAmount
`func (o *AirlineLeg) UnsetDepartureTaxAmount()`

UnsetDepartureTaxAmount ensures that no value is present for DepartureTaxAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


