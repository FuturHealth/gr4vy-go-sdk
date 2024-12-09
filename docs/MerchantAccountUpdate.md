# MerchantAccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The human-readable name of the merchant account. | [optional] 
**OutboundWebhookUrl** | Pointer to **NullableString** | The optional URL where webhooks will be received. | [optional] 
**OutboundWebhookUsername** | Pointer to **NullableString** | The optional username to use when &#x60;outbound_webhook_url&#x60; is configured and requires basic authentication. | [optional] 
**OutboundWebhookPassword** | Pointer to **NullableString** | The optional password to use when &#x60;outbound_webhook_url&#x60; is configured and requires basic authentication. | [optional] 
**VisaNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Visa after onboarding to use Network Tokens. The requestor ID must be unique across all schemes and merchant accounts. | [optional] 
**VisaNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Visa after onboarding to use Network Tokens. The application ID must be unique across all schemes and merchant accounts. | [optional] 
**AmexNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Amex after onboarding to use Network Tokens. The requestor ID must be unique across all schemes and merchant accounts. | [optional] 
**AmexNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Amex after onboarding to use Network Tokens. The application ID must be unique across all schemes and merchant accounts. | [optional] 
**MastercardNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Mastercard after onboarding to use Network Tokens. The requestor ID must be unique across all schemes and merchant accounts. | [optional] 
**MastercardNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Mastercard after onboarding to use Network Tokens. The application ID must be unique across all schemes and merchant accounts. | [optional] 
**LoonClientKey** | Pointer to **NullableString** | Client key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  * If the field is not set, the Account Updater service configuration is not updated. * If the field is set to &#x60;null&#x60;, the Account Updater service is disabled. * If the field is set to &#x60;null&#x60;, the other &#x60;loon_*&#x60; fields must be set to &#x60;null&#x60; as well. | [optional] 
**LoonSecretKey** | Pointer to **NullableString** | Secret key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  * If the field is not set, the Account Updater service configuration is not updated. * If the field is set to &#x60;null&#x60;, the Account Updater service is disabled. * If the field is set to &#x60;null&#x60;, the other &#x60;loon_*&#x60; fields must be set to &#x60;null&#x60; as well. | [optional] 
**LoonAcceptedSchemes** | Pointer to **[]string** | Card schemes accepted when creating jobs using this set of Loon API keys. Loon is the Account Updater service used by Gr4vy.  * If the field is not set, the Account Updater service configuration is not updated. * If the field is set to &#x60;null&#x60;, the Account Updater service is disabled. * If the field is set to &#x60;null&#x60;, the other &#x60;loon_*&#x60; fields must be set to &#x60;null&#x60; as well. | [optional] 
**OverCaptureAmount** | Pointer to **NullableFloat32** | The maximum monetary amount allowed for over-capture, in the smallest currency unit, for example &#x60;1299&#x60; cents to allow for an over-capture of &#x60;$12.99&#x60;. | [optional] 
**OverCapturePercentage** | Pointer to **NullableFloat32** | The maximum percentage allowed for over-capture, for example &#x60;25&#x60; to allow for an over-capture of 25% of the original transaction amount. | [optional] 

## Methods

### NewMerchantAccountUpdate

`func NewMerchantAccountUpdate() *MerchantAccountUpdate`

NewMerchantAccountUpdate instantiates a new MerchantAccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAccountUpdateWithDefaults

`func NewMerchantAccountUpdateWithDefaults() *MerchantAccountUpdate`

NewMerchantAccountUpdateWithDefaults instantiates a new MerchantAccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MerchantAccountUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MerchantAccountUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MerchantAccountUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MerchantAccountUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOutboundWebhookUrl

`func (o *MerchantAccountUpdate) GetOutboundWebhookUrl() string`

GetOutboundWebhookUrl returns the OutboundWebhookUrl field if non-nil, zero value otherwise.

### GetOutboundWebhookUrlOk

`func (o *MerchantAccountUpdate) GetOutboundWebhookUrlOk() (*string, bool)`

GetOutboundWebhookUrlOk returns a tuple with the OutboundWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUrl

`func (o *MerchantAccountUpdate) SetOutboundWebhookUrl(v string)`

SetOutboundWebhookUrl sets OutboundWebhookUrl field to given value.

### HasOutboundWebhookUrl

`func (o *MerchantAccountUpdate) HasOutboundWebhookUrl() bool`

HasOutboundWebhookUrl returns a boolean if a field has been set.

### SetOutboundWebhookUrlNil

`func (o *MerchantAccountUpdate) SetOutboundWebhookUrlNil(b bool)`

 SetOutboundWebhookUrlNil sets the value for OutboundWebhookUrl to be an explicit nil

### UnsetOutboundWebhookUrl
`func (o *MerchantAccountUpdate) UnsetOutboundWebhookUrl()`

UnsetOutboundWebhookUrl ensures that no value is present for OutboundWebhookUrl, not even an explicit nil
### GetOutboundWebhookUsername

`func (o *MerchantAccountUpdate) GetOutboundWebhookUsername() string`

GetOutboundWebhookUsername returns the OutboundWebhookUsername field if non-nil, zero value otherwise.

### GetOutboundWebhookUsernameOk

`func (o *MerchantAccountUpdate) GetOutboundWebhookUsernameOk() (*string, bool)`

GetOutboundWebhookUsernameOk returns a tuple with the OutboundWebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUsername

`func (o *MerchantAccountUpdate) SetOutboundWebhookUsername(v string)`

SetOutboundWebhookUsername sets OutboundWebhookUsername field to given value.

### HasOutboundWebhookUsername

`func (o *MerchantAccountUpdate) HasOutboundWebhookUsername() bool`

HasOutboundWebhookUsername returns a boolean if a field has been set.

### SetOutboundWebhookUsernameNil

`func (o *MerchantAccountUpdate) SetOutboundWebhookUsernameNil(b bool)`

 SetOutboundWebhookUsernameNil sets the value for OutboundWebhookUsername to be an explicit nil

### UnsetOutboundWebhookUsername
`func (o *MerchantAccountUpdate) UnsetOutboundWebhookUsername()`

UnsetOutboundWebhookUsername ensures that no value is present for OutboundWebhookUsername, not even an explicit nil
### GetOutboundWebhookPassword

`func (o *MerchantAccountUpdate) GetOutboundWebhookPassword() string`

GetOutboundWebhookPassword returns the OutboundWebhookPassword field if non-nil, zero value otherwise.

### GetOutboundWebhookPasswordOk

`func (o *MerchantAccountUpdate) GetOutboundWebhookPasswordOk() (*string, bool)`

GetOutboundWebhookPasswordOk returns a tuple with the OutboundWebhookPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookPassword

`func (o *MerchantAccountUpdate) SetOutboundWebhookPassword(v string)`

SetOutboundWebhookPassword sets OutboundWebhookPassword field to given value.

### HasOutboundWebhookPassword

`func (o *MerchantAccountUpdate) HasOutboundWebhookPassword() bool`

HasOutboundWebhookPassword returns a boolean if a field has been set.

### SetOutboundWebhookPasswordNil

`func (o *MerchantAccountUpdate) SetOutboundWebhookPasswordNil(b bool)`

 SetOutboundWebhookPasswordNil sets the value for OutboundWebhookPassword to be an explicit nil

### UnsetOutboundWebhookPassword
`func (o *MerchantAccountUpdate) UnsetOutboundWebhookPassword()`

UnsetOutboundWebhookPassword ensures that no value is present for OutboundWebhookPassword, not even an explicit nil
### GetVisaNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) GetVisaNetworkTokensRequestorId() string`

GetVisaNetworkTokensRequestorId returns the VisaNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensRequestorIdOk

`func (o *MerchantAccountUpdate) GetVisaNetworkTokensRequestorIdOk() (*string, bool)`

GetVisaNetworkTokensRequestorIdOk returns a tuple with the VisaNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) SetVisaNetworkTokensRequestorId(v string)`

SetVisaNetworkTokensRequestorId sets VisaNetworkTokensRequestorId field to given value.

### HasVisaNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) HasVisaNetworkTokensRequestorId() bool`

HasVisaNetworkTokensRequestorId returns a boolean if a field has been set.

### SetVisaNetworkTokensRequestorIdNil

`func (o *MerchantAccountUpdate) SetVisaNetworkTokensRequestorIdNil(b bool)`

 SetVisaNetworkTokensRequestorIdNil sets the value for VisaNetworkTokensRequestorId to be an explicit nil

### UnsetVisaNetworkTokensRequestorId
`func (o *MerchantAccountUpdate) UnsetVisaNetworkTokensRequestorId()`

UnsetVisaNetworkTokensRequestorId ensures that no value is present for VisaNetworkTokensRequestorId, not even an explicit nil
### GetVisaNetworkTokensAppId

`func (o *MerchantAccountUpdate) GetVisaNetworkTokensAppId() string`

GetVisaNetworkTokensAppId returns the VisaNetworkTokensAppId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensAppIdOk

`func (o *MerchantAccountUpdate) GetVisaNetworkTokensAppIdOk() (*string, bool)`

GetVisaNetworkTokensAppIdOk returns a tuple with the VisaNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensAppId

`func (o *MerchantAccountUpdate) SetVisaNetworkTokensAppId(v string)`

SetVisaNetworkTokensAppId sets VisaNetworkTokensAppId field to given value.

### HasVisaNetworkTokensAppId

`func (o *MerchantAccountUpdate) HasVisaNetworkTokensAppId() bool`

HasVisaNetworkTokensAppId returns a boolean if a field has been set.

### SetVisaNetworkTokensAppIdNil

`func (o *MerchantAccountUpdate) SetVisaNetworkTokensAppIdNil(b bool)`

 SetVisaNetworkTokensAppIdNil sets the value for VisaNetworkTokensAppId to be an explicit nil

### UnsetVisaNetworkTokensAppId
`func (o *MerchantAccountUpdate) UnsetVisaNetworkTokensAppId()`

UnsetVisaNetworkTokensAppId ensures that no value is present for VisaNetworkTokensAppId, not even an explicit nil
### GetAmexNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) GetAmexNetworkTokensRequestorId() string`

GetAmexNetworkTokensRequestorId returns the AmexNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensRequestorIdOk

`func (o *MerchantAccountUpdate) GetAmexNetworkTokensRequestorIdOk() (*string, bool)`

GetAmexNetworkTokensRequestorIdOk returns a tuple with the AmexNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) SetAmexNetworkTokensRequestorId(v string)`

SetAmexNetworkTokensRequestorId sets AmexNetworkTokensRequestorId field to given value.

### HasAmexNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) HasAmexNetworkTokensRequestorId() bool`

HasAmexNetworkTokensRequestorId returns a boolean if a field has been set.

### SetAmexNetworkTokensRequestorIdNil

`func (o *MerchantAccountUpdate) SetAmexNetworkTokensRequestorIdNil(b bool)`

 SetAmexNetworkTokensRequestorIdNil sets the value for AmexNetworkTokensRequestorId to be an explicit nil

### UnsetAmexNetworkTokensRequestorId
`func (o *MerchantAccountUpdate) UnsetAmexNetworkTokensRequestorId()`

UnsetAmexNetworkTokensRequestorId ensures that no value is present for AmexNetworkTokensRequestorId, not even an explicit nil
### GetAmexNetworkTokensAppId

`func (o *MerchantAccountUpdate) GetAmexNetworkTokensAppId() string`

GetAmexNetworkTokensAppId returns the AmexNetworkTokensAppId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensAppIdOk

`func (o *MerchantAccountUpdate) GetAmexNetworkTokensAppIdOk() (*string, bool)`

GetAmexNetworkTokensAppIdOk returns a tuple with the AmexNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensAppId

`func (o *MerchantAccountUpdate) SetAmexNetworkTokensAppId(v string)`

SetAmexNetworkTokensAppId sets AmexNetworkTokensAppId field to given value.

### HasAmexNetworkTokensAppId

`func (o *MerchantAccountUpdate) HasAmexNetworkTokensAppId() bool`

HasAmexNetworkTokensAppId returns a boolean if a field has been set.

### SetAmexNetworkTokensAppIdNil

`func (o *MerchantAccountUpdate) SetAmexNetworkTokensAppIdNil(b bool)`

 SetAmexNetworkTokensAppIdNil sets the value for AmexNetworkTokensAppId to be an explicit nil

### UnsetAmexNetworkTokensAppId
`func (o *MerchantAccountUpdate) UnsetAmexNetworkTokensAppId()`

UnsetAmexNetworkTokensAppId ensures that no value is present for AmexNetworkTokensAppId, not even an explicit nil
### GetMastercardNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) GetMastercardNetworkTokensRequestorId() string`

GetMastercardNetworkTokensRequestorId returns the MastercardNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensRequestorIdOk

`func (o *MerchantAccountUpdate) GetMastercardNetworkTokensRequestorIdOk() (*string, bool)`

GetMastercardNetworkTokensRequestorIdOk returns a tuple with the MastercardNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) SetMastercardNetworkTokensRequestorId(v string)`

SetMastercardNetworkTokensRequestorId sets MastercardNetworkTokensRequestorId field to given value.

### HasMastercardNetworkTokensRequestorId

`func (o *MerchantAccountUpdate) HasMastercardNetworkTokensRequestorId() bool`

HasMastercardNetworkTokensRequestorId returns a boolean if a field has been set.

### SetMastercardNetworkTokensRequestorIdNil

`func (o *MerchantAccountUpdate) SetMastercardNetworkTokensRequestorIdNil(b bool)`

 SetMastercardNetworkTokensRequestorIdNil sets the value for MastercardNetworkTokensRequestorId to be an explicit nil

### UnsetMastercardNetworkTokensRequestorId
`func (o *MerchantAccountUpdate) UnsetMastercardNetworkTokensRequestorId()`

UnsetMastercardNetworkTokensRequestorId ensures that no value is present for MastercardNetworkTokensRequestorId, not even an explicit nil
### GetMastercardNetworkTokensAppId

`func (o *MerchantAccountUpdate) GetMastercardNetworkTokensAppId() string`

GetMastercardNetworkTokensAppId returns the MastercardNetworkTokensAppId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensAppIdOk

`func (o *MerchantAccountUpdate) GetMastercardNetworkTokensAppIdOk() (*string, bool)`

GetMastercardNetworkTokensAppIdOk returns a tuple with the MastercardNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensAppId

`func (o *MerchantAccountUpdate) SetMastercardNetworkTokensAppId(v string)`

SetMastercardNetworkTokensAppId sets MastercardNetworkTokensAppId field to given value.

### HasMastercardNetworkTokensAppId

`func (o *MerchantAccountUpdate) HasMastercardNetworkTokensAppId() bool`

HasMastercardNetworkTokensAppId returns a boolean if a field has been set.

### SetMastercardNetworkTokensAppIdNil

`func (o *MerchantAccountUpdate) SetMastercardNetworkTokensAppIdNil(b bool)`

 SetMastercardNetworkTokensAppIdNil sets the value for MastercardNetworkTokensAppId to be an explicit nil

### UnsetMastercardNetworkTokensAppId
`func (o *MerchantAccountUpdate) UnsetMastercardNetworkTokensAppId()`

UnsetMastercardNetworkTokensAppId ensures that no value is present for MastercardNetworkTokensAppId, not even an explicit nil
### GetLoonClientKey

`func (o *MerchantAccountUpdate) GetLoonClientKey() string`

GetLoonClientKey returns the LoonClientKey field if non-nil, zero value otherwise.

### GetLoonClientKeyOk

`func (o *MerchantAccountUpdate) GetLoonClientKeyOk() (*string, bool)`

GetLoonClientKeyOk returns a tuple with the LoonClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonClientKey

`func (o *MerchantAccountUpdate) SetLoonClientKey(v string)`

SetLoonClientKey sets LoonClientKey field to given value.

### HasLoonClientKey

`func (o *MerchantAccountUpdate) HasLoonClientKey() bool`

HasLoonClientKey returns a boolean if a field has been set.

### SetLoonClientKeyNil

`func (o *MerchantAccountUpdate) SetLoonClientKeyNil(b bool)`

 SetLoonClientKeyNil sets the value for LoonClientKey to be an explicit nil

### UnsetLoonClientKey
`func (o *MerchantAccountUpdate) UnsetLoonClientKey()`

UnsetLoonClientKey ensures that no value is present for LoonClientKey, not even an explicit nil
### GetLoonSecretKey

`func (o *MerchantAccountUpdate) GetLoonSecretKey() string`

GetLoonSecretKey returns the LoonSecretKey field if non-nil, zero value otherwise.

### GetLoonSecretKeyOk

`func (o *MerchantAccountUpdate) GetLoonSecretKeyOk() (*string, bool)`

GetLoonSecretKeyOk returns a tuple with the LoonSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonSecretKey

`func (o *MerchantAccountUpdate) SetLoonSecretKey(v string)`

SetLoonSecretKey sets LoonSecretKey field to given value.

### HasLoonSecretKey

`func (o *MerchantAccountUpdate) HasLoonSecretKey() bool`

HasLoonSecretKey returns a boolean if a field has been set.

### SetLoonSecretKeyNil

`func (o *MerchantAccountUpdate) SetLoonSecretKeyNil(b bool)`

 SetLoonSecretKeyNil sets the value for LoonSecretKey to be an explicit nil

### UnsetLoonSecretKey
`func (o *MerchantAccountUpdate) UnsetLoonSecretKey()`

UnsetLoonSecretKey ensures that no value is present for LoonSecretKey, not even an explicit nil
### GetLoonAcceptedSchemes

`func (o *MerchantAccountUpdate) GetLoonAcceptedSchemes() []string`

GetLoonAcceptedSchemes returns the LoonAcceptedSchemes field if non-nil, zero value otherwise.

### GetLoonAcceptedSchemesOk

`func (o *MerchantAccountUpdate) GetLoonAcceptedSchemesOk() (*[]string, bool)`

GetLoonAcceptedSchemesOk returns a tuple with the LoonAcceptedSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonAcceptedSchemes

`func (o *MerchantAccountUpdate) SetLoonAcceptedSchemes(v []string)`

SetLoonAcceptedSchemes sets LoonAcceptedSchemes field to given value.

### HasLoonAcceptedSchemes

`func (o *MerchantAccountUpdate) HasLoonAcceptedSchemes() bool`

HasLoonAcceptedSchemes returns a boolean if a field has been set.

### SetLoonAcceptedSchemesNil

`func (o *MerchantAccountUpdate) SetLoonAcceptedSchemesNil(b bool)`

 SetLoonAcceptedSchemesNil sets the value for LoonAcceptedSchemes to be an explicit nil

### UnsetLoonAcceptedSchemes
`func (o *MerchantAccountUpdate) UnsetLoonAcceptedSchemes()`

UnsetLoonAcceptedSchemes ensures that no value is present for LoonAcceptedSchemes, not even an explicit nil
### GetOverCaptureAmount

`func (o *MerchantAccountUpdate) GetOverCaptureAmount() float32`

GetOverCaptureAmount returns the OverCaptureAmount field if non-nil, zero value otherwise.

### GetOverCaptureAmountOk

`func (o *MerchantAccountUpdate) GetOverCaptureAmountOk() (*float32, bool)`

GetOverCaptureAmountOk returns a tuple with the OverCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCaptureAmount

`func (o *MerchantAccountUpdate) SetOverCaptureAmount(v float32)`

SetOverCaptureAmount sets OverCaptureAmount field to given value.

### HasOverCaptureAmount

`func (o *MerchantAccountUpdate) HasOverCaptureAmount() bool`

HasOverCaptureAmount returns a boolean if a field has been set.

### SetOverCaptureAmountNil

`func (o *MerchantAccountUpdate) SetOverCaptureAmountNil(b bool)`

 SetOverCaptureAmountNil sets the value for OverCaptureAmount to be an explicit nil

### UnsetOverCaptureAmount
`func (o *MerchantAccountUpdate) UnsetOverCaptureAmount()`

UnsetOverCaptureAmount ensures that no value is present for OverCaptureAmount, not even an explicit nil
### GetOverCapturePercentage

`func (o *MerchantAccountUpdate) GetOverCapturePercentage() float32`

GetOverCapturePercentage returns the OverCapturePercentage field if non-nil, zero value otherwise.

### GetOverCapturePercentageOk

`func (o *MerchantAccountUpdate) GetOverCapturePercentageOk() (*float32, bool)`

GetOverCapturePercentageOk returns a tuple with the OverCapturePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapturePercentage

`func (o *MerchantAccountUpdate) SetOverCapturePercentage(v float32)`

SetOverCapturePercentage sets OverCapturePercentage field to given value.

### HasOverCapturePercentage

`func (o *MerchantAccountUpdate) HasOverCapturePercentage() bool`

HasOverCapturePercentage returns a boolean if a field has been set.

### SetOverCapturePercentageNil

`func (o *MerchantAccountUpdate) SetOverCapturePercentageNil(b bool)`

 SetOverCapturePercentageNil sets the value for OverCapturePercentage to be an explicit nil

### UnsetOverCapturePercentage
`func (o *MerchantAccountUpdate) UnsetOverCapturePercentage()`

UnsetOverCapturePercentage ensures that no value is present for OverCapturePercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


