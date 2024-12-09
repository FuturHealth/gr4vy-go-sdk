# Airline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassengerNameRecord** | Pointer to **NullableString** | The Passenger Name Record (PNR) in the airline reservation system. | [optional] 
**BookingCode** | Pointer to **NullableString** | The unique identifier of the reservation in the global distribution system. | [optional] 
**TicketNumber** | Pointer to **NullableString** | The airline&#39;s unique ticket number. | [optional] 
**TicketDeliveryMethod** | Pointer to **NullableString** | The delivery method of the ticket. | [optional] [default to "electronic"]
**IssuedAt** | Pointer to **NullableTime** | The date that the ticket was last issued in the airline reservation system. | [optional] 
**IssuedAddress** | Pointer to **NullableString** | The address of the place/agency that issued the ticket. | [optional] 
**TravelAgencyCode** | Pointer to **NullableString** | The IATA travel agency code. | [optional] 
**TravelAgencyName** | Pointer to **NullableString** | The name of the travel agency. | [optional] 
**TravelAgencyInvoiceNumber** | Pointer to **NullableString** | The reference number of the invoice that was issued by the travel agency. | [optional] 
**TravelAgencyPlanName** | Pointer to **NullableString** | The name of the travel agency plan. | [optional] 
**RestrictedTicket** | Pointer to **NullableBool** | Indicates whether the ticket is restricted (refundable). | [optional] 
**IssuingCarrierCode** | Pointer to **NullableString** | For airline aggregators, two-character IATA code of the airline issuing the ticket. | [optional] 
**ReservationSystem** | Pointer to **NullableString** | The name of the reservation system. | [optional] 
**Passengers** | Pointer to [**[]AirlinePassenger**](AirlinePassenger.md) | An array of the travelling passengers. | [optional] 
**Legs** | Pointer to [**[]AirlineLeg**](AirlineLeg.md) | An array of separate trip segments. Each leg contains detailed itinerary information. | [optional] 

## Methods

### NewAirline

`func NewAirline() *Airline`

NewAirline instantiates a new Airline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlineWithDefaults

`func NewAirlineWithDefaults() *Airline`

NewAirlineWithDefaults instantiates a new Airline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassengerNameRecord

`func (o *Airline) GetPassengerNameRecord() string`

GetPassengerNameRecord returns the PassengerNameRecord field if non-nil, zero value otherwise.

### GetPassengerNameRecordOk

`func (o *Airline) GetPassengerNameRecordOk() (*string, bool)`

GetPassengerNameRecordOk returns a tuple with the PassengerNameRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerNameRecord

`func (o *Airline) SetPassengerNameRecord(v string)`

SetPassengerNameRecord sets PassengerNameRecord field to given value.

### HasPassengerNameRecord

`func (o *Airline) HasPassengerNameRecord() bool`

HasPassengerNameRecord returns a boolean if a field has been set.

### SetPassengerNameRecordNil

`func (o *Airline) SetPassengerNameRecordNil(b bool)`

 SetPassengerNameRecordNil sets the value for PassengerNameRecord to be an explicit nil

### UnsetPassengerNameRecord
`func (o *Airline) UnsetPassengerNameRecord()`

UnsetPassengerNameRecord ensures that no value is present for PassengerNameRecord, not even an explicit nil
### GetBookingCode

`func (o *Airline) GetBookingCode() string`

GetBookingCode returns the BookingCode field if non-nil, zero value otherwise.

### GetBookingCodeOk

`func (o *Airline) GetBookingCodeOk() (*string, bool)`

GetBookingCodeOk returns a tuple with the BookingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingCode

`func (o *Airline) SetBookingCode(v string)`

SetBookingCode sets BookingCode field to given value.

### HasBookingCode

`func (o *Airline) HasBookingCode() bool`

HasBookingCode returns a boolean if a field has been set.

### SetBookingCodeNil

`func (o *Airline) SetBookingCodeNil(b bool)`

 SetBookingCodeNil sets the value for BookingCode to be an explicit nil

### UnsetBookingCode
`func (o *Airline) UnsetBookingCode()`

UnsetBookingCode ensures that no value is present for BookingCode, not even an explicit nil
### GetTicketNumber

`func (o *Airline) GetTicketNumber() string`

GetTicketNumber returns the TicketNumber field if non-nil, zero value otherwise.

### GetTicketNumberOk

`func (o *Airline) GetTicketNumberOk() (*string, bool)`

GetTicketNumberOk returns a tuple with the TicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumber

`func (o *Airline) SetTicketNumber(v string)`

SetTicketNumber sets TicketNumber field to given value.

### HasTicketNumber

`func (o *Airline) HasTicketNumber() bool`

HasTicketNumber returns a boolean if a field has been set.

### SetTicketNumberNil

`func (o *Airline) SetTicketNumberNil(b bool)`

 SetTicketNumberNil sets the value for TicketNumber to be an explicit nil

### UnsetTicketNumber
`func (o *Airline) UnsetTicketNumber()`

UnsetTicketNumber ensures that no value is present for TicketNumber, not even an explicit nil
### GetTicketDeliveryMethod

`func (o *Airline) GetTicketDeliveryMethod() string`

GetTicketDeliveryMethod returns the TicketDeliveryMethod field if non-nil, zero value otherwise.

### GetTicketDeliveryMethodOk

`func (o *Airline) GetTicketDeliveryMethodOk() (*string, bool)`

GetTicketDeliveryMethodOk returns a tuple with the TicketDeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketDeliveryMethod

`func (o *Airline) SetTicketDeliveryMethod(v string)`

SetTicketDeliveryMethod sets TicketDeliveryMethod field to given value.

### HasTicketDeliveryMethod

`func (o *Airline) HasTicketDeliveryMethod() bool`

HasTicketDeliveryMethod returns a boolean if a field has been set.

### SetTicketDeliveryMethodNil

`func (o *Airline) SetTicketDeliveryMethodNil(b bool)`

 SetTicketDeliveryMethodNil sets the value for TicketDeliveryMethod to be an explicit nil

### UnsetTicketDeliveryMethod
`func (o *Airline) UnsetTicketDeliveryMethod()`

UnsetTicketDeliveryMethod ensures that no value is present for TicketDeliveryMethod, not even an explicit nil
### GetIssuedAt

`func (o *Airline) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Airline) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Airline) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Airline) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### SetIssuedAtNil

`func (o *Airline) SetIssuedAtNil(b bool)`

 SetIssuedAtNil sets the value for IssuedAt to be an explicit nil

### UnsetIssuedAt
`func (o *Airline) UnsetIssuedAt()`

UnsetIssuedAt ensures that no value is present for IssuedAt, not even an explicit nil
### GetIssuedAddress

`func (o *Airline) GetIssuedAddress() string`

GetIssuedAddress returns the IssuedAddress field if non-nil, zero value otherwise.

### GetIssuedAddressOk

`func (o *Airline) GetIssuedAddressOk() (*string, bool)`

GetIssuedAddressOk returns a tuple with the IssuedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAddress

`func (o *Airline) SetIssuedAddress(v string)`

SetIssuedAddress sets IssuedAddress field to given value.

### HasIssuedAddress

`func (o *Airline) HasIssuedAddress() bool`

HasIssuedAddress returns a boolean if a field has been set.

### SetIssuedAddressNil

`func (o *Airline) SetIssuedAddressNil(b bool)`

 SetIssuedAddressNil sets the value for IssuedAddress to be an explicit nil

### UnsetIssuedAddress
`func (o *Airline) UnsetIssuedAddress()`

UnsetIssuedAddress ensures that no value is present for IssuedAddress, not even an explicit nil
### GetTravelAgencyCode

`func (o *Airline) GetTravelAgencyCode() string`

GetTravelAgencyCode returns the TravelAgencyCode field if non-nil, zero value otherwise.

### GetTravelAgencyCodeOk

`func (o *Airline) GetTravelAgencyCodeOk() (*string, bool)`

GetTravelAgencyCodeOk returns a tuple with the TravelAgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelAgencyCode

`func (o *Airline) SetTravelAgencyCode(v string)`

SetTravelAgencyCode sets TravelAgencyCode field to given value.

### HasTravelAgencyCode

`func (o *Airline) HasTravelAgencyCode() bool`

HasTravelAgencyCode returns a boolean if a field has been set.

### SetTravelAgencyCodeNil

`func (o *Airline) SetTravelAgencyCodeNil(b bool)`

 SetTravelAgencyCodeNil sets the value for TravelAgencyCode to be an explicit nil

### UnsetTravelAgencyCode
`func (o *Airline) UnsetTravelAgencyCode()`

UnsetTravelAgencyCode ensures that no value is present for TravelAgencyCode, not even an explicit nil
### GetTravelAgencyName

`func (o *Airline) GetTravelAgencyName() string`

GetTravelAgencyName returns the TravelAgencyName field if non-nil, zero value otherwise.

### GetTravelAgencyNameOk

`func (o *Airline) GetTravelAgencyNameOk() (*string, bool)`

GetTravelAgencyNameOk returns a tuple with the TravelAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelAgencyName

`func (o *Airline) SetTravelAgencyName(v string)`

SetTravelAgencyName sets TravelAgencyName field to given value.

### HasTravelAgencyName

`func (o *Airline) HasTravelAgencyName() bool`

HasTravelAgencyName returns a boolean if a field has been set.

### SetTravelAgencyNameNil

`func (o *Airline) SetTravelAgencyNameNil(b bool)`

 SetTravelAgencyNameNil sets the value for TravelAgencyName to be an explicit nil

### UnsetTravelAgencyName
`func (o *Airline) UnsetTravelAgencyName()`

UnsetTravelAgencyName ensures that no value is present for TravelAgencyName, not even an explicit nil
### GetTravelAgencyInvoiceNumber

`func (o *Airline) GetTravelAgencyInvoiceNumber() string`

GetTravelAgencyInvoiceNumber returns the TravelAgencyInvoiceNumber field if non-nil, zero value otherwise.

### GetTravelAgencyInvoiceNumberOk

`func (o *Airline) GetTravelAgencyInvoiceNumberOk() (*string, bool)`

GetTravelAgencyInvoiceNumberOk returns a tuple with the TravelAgencyInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelAgencyInvoiceNumber

`func (o *Airline) SetTravelAgencyInvoiceNumber(v string)`

SetTravelAgencyInvoiceNumber sets TravelAgencyInvoiceNumber field to given value.

### HasTravelAgencyInvoiceNumber

`func (o *Airline) HasTravelAgencyInvoiceNumber() bool`

HasTravelAgencyInvoiceNumber returns a boolean if a field has been set.

### SetTravelAgencyInvoiceNumberNil

`func (o *Airline) SetTravelAgencyInvoiceNumberNil(b bool)`

 SetTravelAgencyInvoiceNumberNil sets the value for TravelAgencyInvoiceNumber to be an explicit nil

### UnsetTravelAgencyInvoiceNumber
`func (o *Airline) UnsetTravelAgencyInvoiceNumber()`

UnsetTravelAgencyInvoiceNumber ensures that no value is present for TravelAgencyInvoiceNumber, not even an explicit nil
### GetTravelAgencyPlanName

`func (o *Airline) GetTravelAgencyPlanName() string`

GetTravelAgencyPlanName returns the TravelAgencyPlanName field if non-nil, zero value otherwise.

### GetTravelAgencyPlanNameOk

`func (o *Airline) GetTravelAgencyPlanNameOk() (*string, bool)`

GetTravelAgencyPlanNameOk returns a tuple with the TravelAgencyPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelAgencyPlanName

`func (o *Airline) SetTravelAgencyPlanName(v string)`

SetTravelAgencyPlanName sets TravelAgencyPlanName field to given value.

### HasTravelAgencyPlanName

`func (o *Airline) HasTravelAgencyPlanName() bool`

HasTravelAgencyPlanName returns a boolean if a field has been set.

### SetTravelAgencyPlanNameNil

`func (o *Airline) SetTravelAgencyPlanNameNil(b bool)`

 SetTravelAgencyPlanNameNil sets the value for TravelAgencyPlanName to be an explicit nil

### UnsetTravelAgencyPlanName
`func (o *Airline) UnsetTravelAgencyPlanName()`

UnsetTravelAgencyPlanName ensures that no value is present for TravelAgencyPlanName, not even an explicit nil
### GetRestrictedTicket

`func (o *Airline) GetRestrictedTicket() bool`

GetRestrictedTicket returns the RestrictedTicket field if non-nil, zero value otherwise.

### GetRestrictedTicketOk

`func (o *Airline) GetRestrictedTicketOk() (*bool, bool)`

GetRestrictedTicketOk returns a tuple with the RestrictedTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedTicket

`func (o *Airline) SetRestrictedTicket(v bool)`

SetRestrictedTicket sets RestrictedTicket field to given value.

### HasRestrictedTicket

`func (o *Airline) HasRestrictedTicket() bool`

HasRestrictedTicket returns a boolean if a field has been set.

### SetRestrictedTicketNil

`func (o *Airline) SetRestrictedTicketNil(b bool)`

 SetRestrictedTicketNil sets the value for RestrictedTicket to be an explicit nil

### UnsetRestrictedTicket
`func (o *Airline) UnsetRestrictedTicket()`

UnsetRestrictedTicket ensures that no value is present for RestrictedTicket, not even an explicit nil
### GetIssuingCarrierCode

`func (o *Airline) GetIssuingCarrierCode() string`

GetIssuingCarrierCode returns the IssuingCarrierCode field if non-nil, zero value otherwise.

### GetIssuingCarrierCodeOk

`func (o *Airline) GetIssuingCarrierCodeOk() (*string, bool)`

GetIssuingCarrierCodeOk returns a tuple with the IssuingCarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCarrierCode

`func (o *Airline) SetIssuingCarrierCode(v string)`

SetIssuingCarrierCode sets IssuingCarrierCode field to given value.

### HasIssuingCarrierCode

`func (o *Airline) HasIssuingCarrierCode() bool`

HasIssuingCarrierCode returns a boolean if a field has been set.

### SetIssuingCarrierCodeNil

`func (o *Airline) SetIssuingCarrierCodeNil(b bool)`

 SetIssuingCarrierCodeNil sets the value for IssuingCarrierCode to be an explicit nil

### UnsetIssuingCarrierCode
`func (o *Airline) UnsetIssuingCarrierCode()`

UnsetIssuingCarrierCode ensures that no value is present for IssuingCarrierCode, not even an explicit nil
### GetReservationSystem

`func (o *Airline) GetReservationSystem() string`

GetReservationSystem returns the ReservationSystem field if non-nil, zero value otherwise.

### GetReservationSystemOk

`func (o *Airline) GetReservationSystemOk() (*string, bool)`

GetReservationSystemOk returns a tuple with the ReservationSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationSystem

`func (o *Airline) SetReservationSystem(v string)`

SetReservationSystem sets ReservationSystem field to given value.

### HasReservationSystem

`func (o *Airline) HasReservationSystem() bool`

HasReservationSystem returns a boolean if a field has been set.

### SetReservationSystemNil

`func (o *Airline) SetReservationSystemNil(b bool)`

 SetReservationSystemNil sets the value for ReservationSystem to be an explicit nil

### UnsetReservationSystem
`func (o *Airline) UnsetReservationSystem()`

UnsetReservationSystem ensures that no value is present for ReservationSystem, not even an explicit nil
### GetPassengers

`func (o *Airline) GetPassengers() []AirlinePassenger`

GetPassengers returns the Passengers field if non-nil, zero value otherwise.

### GetPassengersOk

`func (o *Airline) GetPassengersOk() (*[]AirlinePassenger, bool)`

GetPassengersOk returns a tuple with the Passengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengers

`func (o *Airline) SetPassengers(v []AirlinePassenger)`

SetPassengers sets Passengers field to given value.

### HasPassengers

`func (o *Airline) HasPassengers() bool`

HasPassengers returns a boolean if a field has been set.

### SetPassengersNil

`func (o *Airline) SetPassengersNil(b bool)`

 SetPassengersNil sets the value for Passengers to be an explicit nil

### UnsetPassengers
`func (o *Airline) UnsetPassengers()`

UnsetPassengers ensures that no value is present for Passengers, not even an explicit nil
### GetLegs

`func (o *Airline) GetLegs() []AirlineLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *Airline) GetLegsOk() (*[]AirlineLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *Airline) SetLegs(v []AirlineLeg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *Airline) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### SetLegsNil

`func (o *Airline) SetLegsNil(b bool)`

 SetLegsNil sets the value for Legs to be an explicit nil

### UnsetLegs
`func (o *Airline) UnsetLegs()`

UnsetLegs ensures that no value is present for Legs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


