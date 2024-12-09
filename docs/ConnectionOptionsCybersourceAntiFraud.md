# ConnectionOptionsCybersourceAntiFraud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaKeyMerchantId** | Pointer to **NullableString** | An override for the merchant ID configured for the connector, used in combination with meta keys. | [optional] 
**MerchantDefinedData** | Pointer to **map[string]string** | This is a key-value object for merchant defined data. Each key needs to be a numeric string identifying the MDD field to set. For example, for field 1 set the key to \&quot;1\&quot;.  Please avoid fields \&quot;31\&quot;, \&quot;48\&quot;, \&quot;50\&quot;-\&quot;56\&quot; and \&quot;63\&quot;-\&quot;76\&quot; as these are auto-populated based on the transaction profile. | [optional] 
**ShippingMethod** | Pointer to **NullableString** | Shipping method for the order. | [optional] 

## Methods

### NewConnectionOptionsCybersourceAntiFraud

`func NewConnectionOptionsCybersourceAntiFraud() *ConnectionOptionsCybersourceAntiFraud`

NewConnectionOptionsCybersourceAntiFraud instantiates a new ConnectionOptionsCybersourceAntiFraud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsCybersourceAntiFraudWithDefaults

`func NewConnectionOptionsCybersourceAntiFraudWithDefaults() *ConnectionOptionsCybersourceAntiFraud`

NewConnectionOptionsCybersourceAntiFraudWithDefaults instantiates a new ConnectionOptionsCybersourceAntiFraud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMetaKeyMerchantId() string`

GetMetaKeyMerchantId returns the MetaKeyMerchantId field if non-nil, zero value otherwise.

### GetMetaKeyMerchantIdOk

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMetaKeyMerchantIdOk() (*string, bool)`

GetMetaKeyMerchantIdOk returns a tuple with the MetaKeyMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceAntiFraud) SetMetaKeyMerchantId(v string)`

SetMetaKeyMerchantId sets MetaKeyMerchantId field to given value.

### HasMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceAntiFraud) HasMetaKeyMerchantId() bool`

HasMetaKeyMerchantId returns a boolean if a field has been set.

### SetMetaKeyMerchantIdNil

`func (o *ConnectionOptionsCybersourceAntiFraud) SetMetaKeyMerchantIdNil(b bool)`

 SetMetaKeyMerchantIdNil sets the value for MetaKeyMerchantId to be an explicit nil

### UnsetMetaKeyMerchantId
`func (o *ConnectionOptionsCybersourceAntiFraud) UnsetMetaKeyMerchantId()`

UnsetMetaKeyMerchantId ensures that no value is present for MetaKeyMerchantId, not even an explicit nil
### GetMerchantDefinedData

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMerchantDefinedData() map[string]string`

GetMerchantDefinedData returns the MerchantDefinedData field if non-nil, zero value otherwise.

### GetMerchantDefinedDataOk

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMerchantDefinedDataOk() (*map[string]string, bool)`

GetMerchantDefinedDataOk returns a tuple with the MerchantDefinedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDefinedData

`func (o *ConnectionOptionsCybersourceAntiFraud) SetMerchantDefinedData(v map[string]string)`

SetMerchantDefinedData sets MerchantDefinedData field to given value.

### HasMerchantDefinedData

`func (o *ConnectionOptionsCybersourceAntiFraud) HasMerchantDefinedData() bool`

HasMerchantDefinedData returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ConnectionOptionsCybersourceAntiFraud) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ConnectionOptionsCybersourceAntiFraud) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ConnectionOptionsCybersourceAntiFraud) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ConnectionOptionsCybersourceAntiFraud) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### SetShippingMethodNil

`func (o *ConnectionOptionsCybersourceAntiFraud) SetShippingMethodNil(b bool)`

 SetShippingMethodNil sets the value for ShippingMethod to be an explicit nil

### UnsetShippingMethod
`func (o *ConnectionOptionsCybersourceAntiFraud) UnsetShippingMethod()`

UnsetShippingMethod ensures that no value is present for ShippingMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


