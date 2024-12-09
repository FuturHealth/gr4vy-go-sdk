# MerchantAccountCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID for the merchant account. | [optional] 
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
**LoonClientKey** | Pointer to **NullableString** | Client key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  * If the field is not set or if it&#39;s set to &#x60;null&#x60;, the Account Updater service doesn&#39;t get configured. * If the field is set to &#x60;null&#x60;, the other &#x60;loon_*&#x60; fields must be set to &#x60;null&#x60; as well. | [optional] 
**LoonSecretKey** | Pointer to **NullableString** | Secret key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  * If the field is not set or if it&#39;s set to &#x60;null&#x60;, the Account Updater service doesn&#39;t get configured. * If the field is set to &#x60;null&#x60;, the other &#x60;loon_*&#x60; fields must be set to &#x60;null&#x60; as well. | [optional] 
**LoonAcceptedSchemes** | Pointer to **[]string** | Card schemes accepted when creating jobs using this set of Loon API keys. Loon is the Account Updater service used by Gr4vy.  * If the field is not set or if it&#39;s set to &#x60;null&#x60;, the Account Updater service doesn&#39;t get configured. * If the field is set to &#x60;null&#x60;, the other &#x60;loon_*&#x60; fields must be set to &#x60;null&#x60; as well. | [optional] 
**OverCaptureAmount** | Pointer to **NullableFloat32** | The maximum monetary amount allowed for over-capture, in the smallest currency unit, for example &#x60;1299&#x60; cents to allow for an over-capture of &#x60;$12.99&#x60;. | [optional] 
**OverCapturePercentage** | Pointer to **NullableFloat32** | The maximum percentage allowed for over-capture, for example &#x60;25&#x60; to allow for an over-capture of 25% of the original transaction amount. | [optional] 

## Methods

### NewMerchantAccountCreate

`func NewMerchantAccountCreate() *MerchantAccountCreate`

NewMerchantAccountCreate instantiates a new MerchantAccountCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAccountCreateWithDefaults

`func NewMerchantAccountCreateWithDefaults() *MerchantAccountCreate`

NewMerchantAccountCreateWithDefaults instantiates a new MerchantAccountCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantAccountCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantAccountCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantAccountCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantAccountCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MerchantAccountCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MerchantAccountCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MerchantAccountCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MerchantAccountCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOutboundWebhookUrl

`func (o *MerchantAccountCreate) GetOutboundWebhookUrl() string`

GetOutboundWebhookUrl returns the OutboundWebhookUrl field if non-nil, zero value otherwise.

### GetOutboundWebhookUrlOk

`func (o *MerchantAccountCreate) GetOutboundWebhookUrlOk() (*string, bool)`

GetOutboundWebhookUrlOk returns a tuple with the OutboundWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUrl

`func (o *MerchantAccountCreate) SetOutboundWebhookUrl(v string)`

SetOutboundWebhookUrl sets OutboundWebhookUrl field to given value.

### HasOutboundWebhookUrl

`func (o *MerchantAccountCreate) HasOutboundWebhookUrl() bool`

HasOutboundWebhookUrl returns a boolean if a field has been set.

### SetOutboundWebhookUrlNil

`func (o *MerchantAccountCreate) SetOutboundWebhookUrlNil(b bool)`

 SetOutboundWebhookUrlNil sets the value for OutboundWebhookUrl to be an explicit nil

### UnsetOutboundWebhookUrl
`func (o *MerchantAccountCreate) UnsetOutboundWebhookUrl()`

UnsetOutboundWebhookUrl ensures that no value is present for OutboundWebhookUrl, not even an explicit nil
### GetOutboundWebhookUsername

`func (o *MerchantAccountCreate) GetOutboundWebhookUsername() string`

GetOutboundWebhookUsername returns the OutboundWebhookUsername field if non-nil, zero value otherwise.

### GetOutboundWebhookUsernameOk

`func (o *MerchantAccountCreate) GetOutboundWebhookUsernameOk() (*string, bool)`

GetOutboundWebhookUsernameOk returns a tuple with the OutboundWebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUsername

`func (o *MerchantAccountCreate) SetOutboundWebhookUsername(v string)`

SetOutboundWebhookUsername sets OutboundWebhookUsername field to given value.

### HasOutboundWebhookUsername

`func (o *MerchantAccountCreate) HasOutboundWebhookUsername() bool`

HasOutboundWebhookUsername returns a boolean if a field has been set.

### SetOutboundWebhookUsernameNil

`func (o *MerchantAccountCreate) SetOutboundWebhookUsernameNil(b bool)`

 SetOutboundWebhookUsernameNil sets the value for OutboundWebhookUsername to be an explicit nil

### UnsetOutboundWebhookUsername
`func (o *MerchantAccountCreate) UnsetOutboundWebhookUsername()`

UnsetOutboundWebhookUsername ensures that no value is present for OutboundWebhookUsername, not even an explicit nil
### GetOutboundWebhookPassword

`func (o *MerchantAccountCreate) GetOutboundWebhookPassword() string`

GetOutboundWebhookPassword returns the OutboundWebhookPassword field if non-nil, zero value otherwise.

### GetOutboundWebhookPasswordOk

`func (o *MerchantAccountCreate) GetOutboundWebhookPasswordOk() (*string, bool)`

GetOutboundWebhookPasswordOk returns a tuple with the OutboundWebhookPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookPassword

`func (o *MerchantAccountCreate) SetOutboundWebhookPassword(v string)`

SetOutboundWebhookPassword sets OutboundWebhookPassword field to given value.

### HasOutboundWebhookPassword

`func (o *MerchantAccountCreate) HasOutboundWebhookPassword() bool`

HasOutboundWebhookPassword returns a boolean if a field has been set.

### SetOutboundWebhookPasswordNil

`func (o *MerchantAccountCreate) SetOutboundWebhookPasswordNil(b bool)`

 SetOutboundWebhookPasswordNil sets the value for OutboundWebhookPassword to be an explicit nil

### UnsetOutboundWebhookPassword
`func (o *MerchantAccountCreate) UnsetOutboundWebhookPassword()`

UnsetOutboundWebhookPassword ensures that no value is present for OutboundWebhookPassword, not even an explicit nil
### GetVisaNetworkTokensRequestorId

`func (o *MerchantAccountCreate) GetVisaNetworkTokensRequestorId() string`

GetVisaNetworkTokensRequestorId returns the VisaNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensRequestorIdOk

`func (o *MerchantAccountCreate) GetVisaNetworkTokensRequestorIdOk() (*string, bool)`

GetVisaNetworkTokensRequestorIdOk returns a tuple with the VisaNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensRequestorId

`func (o *MerchantAccountCreate) SetVisaNetworkTokensRequestorId(v string)`

SetVisaNetworkTokensRequestorId sets VisaNetworkTokensRequestorId field to given value.

### HasVisaNetworkTokensRequestorId

`func (o *MerchantAccountCreate) HasVisaNetworkTokensRequestorId() bool`

HasVisaNetworkTokensRequestorId returns a boolean if a field has been set.

### SetVisaNetworkTokensRequestorIdNil

`func (o *MerchantAccountCreate) SetVisaNetworkTokensRequestorIdNil(b bool)`

 SetVisaNetworkTokensRequestorIdNil sets the value for VisaNetworkTokensRequestorId to be an explicit nil

### UnsetVisaNetworkTokensRequestorId
`func (o *MerchantAccountCreate) UnsetVisaNetworkTokensRequestorId()`

UnsetVisaNetworkTokensRequestorId ensures that no value is present for VisaNetworkTokensRequestorId, not even an explicit nil
### GetVisaNetworkTokensAppId

`func (o *MerchantAccountCreate) GetVisaNetworkTokensAppId() string`

GetVisaNetworkTokensAppId returns the VisaNetworkTokensAppId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensAppIdOk

`func (o *MerchantAccountCreate) GetVisaNetworkTokensAppIdOk() (*string, bool)`

GetVisaNetworkTokensAppIdOk returns a tuple with the VisaNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensAppId

`func (o *MerchantAccountCreate) SetVisaNetworkTokensAppId(v string)`

SetVisaNetworkTokensAppId sets VisaNetworkTokensAppId field to given value.

### HasVisaNetworkTokensAppId

`func (o *MerchantAccountCreate) HasVisaNetworkTokensAppId() bool`

HasVisaNetworkTokensAppId returns a boolean if a field has been set.

### SetVisaNetworkTokensAppIdNil

`func (o *MerchantAccountCreate) SetVisaNetworkTokensAppIdNil(b bool)`

 SetVisaNetworkTokensAppIdNil sets the value for VisaNetworkTokensAppId to be an explicit nil

### UnsetVisaNetworkTokensAppId
`func (o *MerchantAccountCreate) UnsetVisaNetworkTokensAppId()`

UnsetVisaNetworkTokensAppId ensures that no value is present for VisaNetworkTokensAppId, not even an explicit nil
### GetAmexNetworkTokensRequestorId

`func (o *MerchantAccountCreate) GetAmexNetworkTokensRequestorId() string`

GetAmexNetworkTokensRequestorId returns the AmexNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensRequestorIdOk

`func (o *MerchantAccountCreate) GetAmexNetworkTokensRequestorIdOk() (*string, bool)`

GetAmexNetworkTokensRequestorIdOk returns a tuple with the AmexNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensRequestorId

`func (o *MerchantAccountCreate) SetAmexNetworkTokensRequestorId(v string)`

SetAmexNetworkTokensRequestorId sets AmexNetworkTokensRequestorId field to given value.

### HasAmexNetworkTokensRequestorId

`func (o *MerchantAccountCreate) HasAmexNetworkTokensRequestorId() bool`

HasAmexNetworkTokensRequestorId returns a boolean if a field has been set.

### SetAmexNetworkTokensRequestorIdNil

`func (o *MerchantAccountCreate) SetAmexNetworkTokensRequestorIdNil(b bool)`

 SetAmexNetworkTokensRequestorIdNil sets the value for AmexNetworkTokensRequestorId to be an explicit nil

### UnsetAmexNetworkTokensRequestorId
`func (o *MerchantAccountCreate) UnsetAmexNetworkTokensRequestorId()`

UnsetAmexNetworkTokensRequestorId ensures that no value is present for AmexNetworkTokensRequestorId, not even an explicit nil
### GetAmexNetworkTokensAppId

`func (o *MerchantAccountCreate) GetAmexNetworkTokensAppId() string`

GetAmexNetworkTokensAppId returns the AmexNetworkTokensAppId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensAppIdOk

`func (o *MerchantAccountCreate) GetAmexNetworkTokensAppIdOk() (*string, bool)`

GetAmexNetworkTokensAppIdOk returns a tuple with the AmexNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensAppId

`func (o *MerchantAccountCreate) SetAmexNetworkTokensAppId(v string)`

SetAmexNetworkTokensAppId sets AmexNetworkTokensAppId field to given value.

### HasAmexNetworkTokensAppId

`func (o *MerchantAccountCreate) HasAmexNetworkTokensAppId() bool`

HasAmexNetworkTokensAppId returns a boolean if a field has been set.

### SetAmexNetworkTokensAppIdNil

`func (o *MerchantAccountCreate) SetAmexNetworkTokensAppIdNil(b bool)`

 SetAmexNetworkTokensAppIdNil sets the value for AmexNetworkTokensAppId to be an explicit nil

### UnsetAmexNetworkTokensAppId
`func (o *MerchantAccountCreate) UnsetAmexNetworkTokensAppId()`

UnsetAmexNetworkTokensAppId ensures that no value is present for AmexNetworkTokensAppId, not even an explicit nil
### GetMastercardNetworkTokensRequestorId

`func (o *MerchantAccountCreate) GetMastercardNetworkTokensRequestorId() string`

GetMastercardNetworkTokensRequestorId returns the MastercardNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensRequestorIdOk

`func (o *MerchantAccountCreate) GetMastercardNetworkTokensRequestorIdOk() (*string, bool)`

GetMastercardNetworkTokensRequestorIdOk returns a tuple with the MastercardNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensRequestorId

`func (o *MerchantAccountCreate) SetMastercardNetworkTokensRequestorId(v string)`

SetMastercardNetworkTokensRequestorId sets MastercardNetworkTokensRequestorId field to given value.

### HasMastercardNetworkTokensRequestorId

`func (o *MerchantAccountCreate) HasMastercardNetworkTokensRequestorId() bool`

HasMastercardNetworkTokensRequestorId returns a boolean if a field has been set.

### SetMastercardNetworkTokensRequestorIdNil

`func (o *MerchantAccountCreate) SetMastercardNetworkTokensRequestorIdNil(b bool)`

 SetMastercardNetworkTokensRequestorIdNil sets the value for MastercardNetworkTokensRequestorId to be an explicit nil

### UnsetMastercardNetworkTokensRequestorId
`func (o *MerchantAccountCreate) UnsetMastercardNetworkTokensRequestorId()`

UnsetMastercardNetworkTokensRequestorId ensures that no value is present for MastercardNetworkTokensRequestorId, not even an explicit nil
### GetMastercardNetworkTokensAppId

`func (o *MerchantAccountCreate) GetMastercardNetworkTokensAppId() string`

GetMastercardNetworkTokensAppId returns the MastercardNetworkTokensAppId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensAppIdOk

`func (o *MerchantAccountCreate) GetMastercardNetworkTokensAppIdOk() (*string, bool)`

GetMastercardNetworkTokensAppIdOk returns a tuple with the MastercardNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensAppId

`func (o *MerchantAccountCreate) SetMastercardNetworkTokensAppId(v string)`

SetMastercardNetworkTokensAppId sets MastercardNetworkTokensAppId field to given value.

### HasMastercardNetworkTokensAppId

`func (o *MerchantAccountCreate) HasMastercardNetworkTokensAppId() bool`

HasMastercardNetworkTokensAppId returns a boolean if a field has been set.

### SetMastercardNetworkTokensAppIdNil

`func (o *MerchantAccountCreate) SetMastercardNetworkTokensAppIdNil(b bool)`

 SetMastercardNetworkTokensAppIdNil sets the value for MastercardNetworkTokensAppId to be an explicit nil

### UnsetMastercardNetworkTokensAppId
`func (o *MerchantAccountCreate) UnsetMastercardNetworkTokensAppId()`

UnsetMastercardNetworkTokensAppId ensures that no value is present for MastercardNetworkTokensAppId, not even an explicit nil
### GetLoonClientKey

`func (o *MerchantAccountCreate) GetLoonClientKey() string`

GetLoonClientKey returns the LoonClientKey field if non-nil, zero value otherwise.

### GetLoonClientKeyOk

`func (o *MerchantAccountCreate) GetLoonClientKeyOk() (*string, bool)`

GetLoonClientKeyOk returns a tuple with the LoonClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonClientKey

`func (o *MerchantAccountCreate) SetLoonClientKey(v string)`

SetLoonClientKey sets LoonClientKey field to given value.

### HasLoonClientKey

`func (o *MerchantAccountCreate) HasLoonClientKey() bool`

HasLoonClientKey returns a boolean if a field has been set.

### SetLoonClientKeyNil

`func (o *MerchantAccountCreate) SetLoonClientKeyNil(b bool)`

 SetLoonClientKeyNil sets the value for LoonClientKey to be an explicit nil

### UnsetLoonClientKey
`func (o *MerchantAccountCreate) UnsetLoonClientKey()`

UnsetLoonClientKey ensures that no value is present for LoonClientKey, not even an explicit nil
### GetLoonSecretKey

`func (o *MerchantAccountCreate) GetLoonSecretKey() string`

GetLoonSecretKey returns the LoonSecretKey field if non-nil, zero value otherwise.

### GetLoonSecretKeyOk

`func (o *MerchantAccountCreate) GetLoonSecretKeyOk() (*string, bool)`

GetLoonSecretKeyOk returns a tuple with the LoonSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonSecretKey

`func (o *MerchantAccountCreate) SetLoonSecretKey(v string)`

SetLoonSecretKey sets LoonSecretKey field to given value.

### HasLoonSecretKey

`func (o *MerchantAccountCreate) HasLoonSecretKey() bool`

HasLoonSecretKey returns a boolean if a field has been set.

### SetLoonSecretKeyNil

`func (o *MerchantAccountCreate) SetLoonSecretKeyNil(b bool)`

 SetLoonSecretKeyNil sets the value for LoonSecretKey to be an explicit nil

### UnsetLoonSecretKey
`func (o *MerchantAccountCreate) UnsetLoonSecretKey()`

UnsetLoonSecretKey ensures that no value is present for LoonSecretKey, not even an explicit nil
### GetLoonAcceptedSchemes

`func (o *MerchantAccountCreate) GetLoonAcceptedSchemes() []string`

GetLoonAcceptedSchemes returns the LoonAcceptedSchemes field if non-nil, zero value otherwise.

### GetLoonAcceptedSchemesOk

`func (o *MerchantAccountCreate) GetLoonAcceptedSchemesOk() (*[]string, bool)`

GetLoonAcceptedSchemesOk returns a tuple with the LoonAcceptedSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonAcceptedSchemes

`func (o *MerchantAccountCreate) SetLoonAcceptedSchemes(v []string)`

SetLoonAcceptedSchemes sets LoonAcceptedSchemes field to given value.

### HasLoonAcceptedSchemes

`func (o *MerchantAccountCreate) HasLoonAcceptedSchemes() bool`

HasLoonAcceptedSchemes returns a boolean if a field has been set.

### SetLoonAcceptedSchemesNil

`func (o *MerchantAccountCreate) SetLoonAcceptedSchemesNil(b bool)`

 SetLoonAcceptedSchemesNil sets the value for LoonAcceptedSchemes to be an explicit nil

### UnsetLoonAcceptedSchemes
`func (o *MerchantAccountCreate) UnsetLoonAcceptedSchemes()`

UnsetLoonAcceptedSchemes ensures that no value is present for LoonAcceptedSchemes, not even an explicit nil
### GetOverCaptureAmount

`func (o *MerchantAccountCreate) GetOverCaptureAmount() float32`

GetOverCaptureAmount returns the OverCaptureAmount field if non-nil, zero value otherwise.

### GetOverCaptureAmountOk

`func (o *MerchantAccountCreate) GetOverCaptureAmountOk() (*float32, bool)`

GetOverCaptureAmountOk returns a tuple with the OverCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCaptureAmount

`func (o *MerchantAccountCreate) SetOverCaptureAmount(v float32)`

SetOverCaptureAmount sets OverCaptureAmount field to given value.

### HasOverCaptureAmount

`func (o *MerchantAccountCreate) HasOverCaptureAmount() bool`

HasOverCaptureAmount returns a boolean if a field has been set.

### SetOverCaptureAmountNil

`func (o *MerchantAccountCreate) SetOverCaptureAmountNil(b bool)`

 SetOverCaptureAmountNil sets the value for OverCaptureAmount to be an explicit nil

### UnsetOverCaptureAmount
`func (o *MerchantAccountCreate) UnsetOverCaptureAmount()`

UnsetOverCaptureAmount ensures that no value is present for OverCaptureAmount, not even an explicit nil
### GetOverCapturePercentage

`func (o *MerchantAccountCreate) GetOverCapturePercentage() float32`

GetOverCapturePercentage returns the OverCapturePercentage field if non-nil, zero value otherwise.

### GetOverCapturePercentageOk

`func (o *MerchantAccountCreate) GetOverCapturePercentageOk() (*float32, bool)`

GetOverCapturePercentageOk returns a tuple with the OverCapturePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapturePercentage

`func (o *MerchantAccountCreate) SetOverCapturePercentage(v float32)`

SetOverCapturePercentage sets OverCapturePercentage field to given value.

### HasOverCapturePercentage

`func (o *MerchantAccountCreate) HasOverCapturePercentage() bool`

HasOverCapturePercentage returns a boolean if a field has been set.

### SetOverCapturePercentageNil

`func (o *MerchantAccountCreate) SetOverCapturePercentageNil(b bool)`

 SetOverCapturePercentageNil sets the value for OverCapturePercentage to be an explicit nil

### UnsetOverCapturePercentage
`func (o *MerchantAccountCreate) UnsetOverCapturePercentage()`

UnsetOverCapturePercentage ensures that no value is present for OverCapturePercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


