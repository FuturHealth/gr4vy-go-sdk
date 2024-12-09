# ConnectionOptionsCybersourceKcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaKeyMerchantId** | Pointer to **NullableString** | An override for the merchant ID configured for the connector, used in combination with meta keys. | [optional] 
**MerchantDefinedInformation** | Pointer to **map[string]string** | This is a key-value object for merchant defined information. Each key needs to be a numeric string identifying the MDI field to set. For example, for field 1 set the key to \&quot;1\&quot;. | [optional] 
**ShipToMethod** | Pointer to **NullableString** | Shipping method for the order. | [optional] 

## Methods

### NewConnectionOptionsCybersourceKcp

`func NewConnectionOptionsCybersourceKcp() *ConnectionOptionsCybersourceKcp`

NewConnectionOptionsCybersourceKcp instantiates a new ConnectionOptionsCybersourceKcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsCybersourceKcpWithDefaults

`func NewConnectionOptionsCybersourceKcpWithDefaults() *ConnectionOptionsCybersourceKcp`

NewConnectionOptionsCybersourceKcpWithDefaults instantiates a new ConnectionOptionsCybersourceKcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceKcp) GetMetaKeyMerchantId() string`

GetMetaKeyMerchantId returns the MetaKeyMerchantId field if non-nil, zero value otherwise.

### GetMetaKeyMerchantIdOk

`func (o *ConnectionOptionsCybersourceKcp) GetMetaKeyMerchantIdOk() (*string, bool)`

GetMetaKeyMerchantIdOk returns a tuple with the MetaKeyMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceKcp) SetMetaKeyMerchantId(v string)`

SetMetaKeyMerchantId sets MetaKeyMerchantId field to given value.

### HasMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceKcp) HasMetaKeyMerchantId() bool`

HasMetaKeyMerchantId returns a boolean if a field has been set.

### SetMetaKeyMerchantIdNil

`func (o *ConnectionOptionsCybersourceKcp) SetMetaKeyMerchantIdNil(b bool)`

 SetMetaKeyMerchantIdNil sets the value for MetaKeyMerchantId to be an explicit nil

### UnsetMetaKeyMerchantId
`func (o *ConnectionOptionsCybersourceKcp) UnsetMetaKeyMerchantId()`

UnsetMetaKeyMerchantId ensures that no value is present for MetaKeyMerchantId, not even an explicit nil
### GetMerchantDefinedInformation

`func (o *ConnectionOptionsCybersourceKcp) GetMerchantDefinedInformation() map[string]string`

GetMerchantDefinedInformation returns the MerchantDefinedInformation field if non-nil, zero value otherwise.

### GetMerchantDefinedInformationOk

`func (o *ConnectionOptionsCybersourceKcp) GetMerchantDefinedInformationOk() (*map[string]string, bool)`

GetMerchantDefinedInformationOk returns a tuple with the MerchantDefinedInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDefinedInformation

`func (o *ConnectionOptionsCybersourceKcp) SetMerchantDefinedInformation(v map[string]string)`

SetMerchantDefinedInformation sets MerchantDefinedInformation field to given value.

### HasMerchantDefinedInformation

`func (o *ConnectionOptionsCybersourceKcp) HasMerchantDefinedInformation() bool`

HasMerchantDefinedInformation returns a boolean if a field has been set.

### GetShipToMethod

`func (o *ConnectionOptionsCybersourceKcp) GetShipToMethod() string`

GetShipToMethod returns the ShipToMethod field if non-nil, zero value otherwise.

### GetShipToMethodOk

`func (o *ConnectionOptionsCybersourceKcp) GetShipToMethodOk() (*string, bool)`

GetShipToMethodOk returns a tuple with the ShipToMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToMethod

`func (o *ConnectionOptionsCybersourceKcp) SetShipToMethod(v string)`

SetShipToMethod sets ShipToMethod field to given value.

### HasShipToMethod

`func (o *ConnectionOptionsCybersourceKcp) HasShipToMethod() bool`

HasShipToMethod returns a boolean if a field has been set.

### SetShipToMethodNil

`func (o *ConnectionOptionsCybersourceKcp) SetShipToMethodNil(b bool)`

 SetShipToMethodNil sets the value for ShipToMethod to be an explicit nil

### UnsetShipToMethod
`func (o *ConnectionOptionsCybersourceKcp) UnsetShipToMethod()`

UnsetShipToMethod ensures that no value is present for ShipToMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


