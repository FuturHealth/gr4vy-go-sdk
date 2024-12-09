# MerchantAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;merchant-account&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID for this merchant account. | [optional] 
**DisplayName** | Pointer to **string** | The display name of this merchant account. | [optional] 
**OutboundWebhookUrl** | Pointer to **NullableString** | The optional URL where webhooks will be received. | [optional] 
**OutboundWebhookUsername** | Pointer to **NullableString** | The optional username to use when &#x60;outbound_webhook_url&#x60; is configured and requires basic authentication. | [optional] 
**OutboundWebhookPassword** | Pointer to **NullableString** | The optional password to use when &#x60;outbound_webhook_url&#x60; is configured and requires basic authentication.  If the field is not &#x60;null&#x60;, the value is masked to avoid exposing sensitive information. | [optional] 
**VisaNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Visa after onboarding to use Network Tokens. | [optional] 
**VisaNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Visa after onboarding to use Network Tokens. | [optional] 
**AmexNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Amex after onboarding to use Network Tokens. | [optional] 
**AmexNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Amex after onboarding to use Network Tokens. | [optional] 
**MastercardNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Mastercard after onboarding to use Network Tokens. | [optional] 
**MastercardNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Mastercard after onboarding to use Network Tokens. | [optional] 
**LoonClientKey** | Pointer to **NullableString** | Client key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy. | [optional] 
**LoonSecretKey** | Pointer to **NullableString** | Secret key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  If the field is not &#x60;null&#x60;, the value is masked to avoid exposing sensitive information. | [optional] 
**LoonAcceptedSchemes** | Pointer to **[]string** | Card schemes accepted when creating jobs using this set of Loon API keys. Loon is the Account Updater service used by Gr4vy. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this merchant account was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this merchant account was updated. | [optional] 
**OverCaptureAmount** | Pointer to **NullableFloat32** | The maximum monetary amount allowed for over-capture, in the smallest currency unit, for example &#x60;1299&#x60; cents to allow for an over-capture of &#x60;$12.99&#x60;. | [optional] 
**OverCapturePercentage** | Pointer to **NullableFloat32** | The maximum percentage allowed for over-capture, for example &#x60;25&#x60; to allow for an over-capture of 25% of the original transaction amount. | [optional] 

## Methods

### NewMerchantAccount

`func NewMerchantAccount() *MerchantAccount`

NewMerchantAccount instantiates a new MerchantAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAccountWithDefaults

`func NewMerchantAccountWithDefaults() *MerchantAccount`

NewMerchantAccountWithDefaults instantiates a new MerchantAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MerchantAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MerchantAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *MerchantAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MerchantAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MerchantAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MerchantAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MerchantAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOutboundWebhookUrl

`func (o *MerchantAccount) GetOutboundWebhookUrl() string`

GetOutboundWebhookUrl returns the OutboundWebhookUrl field if non-nil, zero value otherwise.

### GetOutboundWebhookUrlOk

`func (o *MerchantAccount) GetOutboundWebhookUrlOk() (*string, bool)`

GetOutboundWebhookUrlOk returns a tuple with the OutboundWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUrl

`func (o *MerchantAccount) SetOutboundWebhookUrl(v string)`

SetOutboundWebhookUrl sets OutboundWebhookUrl field to given value.

### HasOutboundWebhookUrl

`func (o *MerchantAccount) HasOutboundWebhookUrl() bool`

HasOutboundWebhookUrl returns a boolean if a field has been set.

### SetOutboundWebhookUrlNil

`func (o *MerchantAccount) SetOutboundWebhookUrlNil(b bool)`

 SetOutboundWebhookUrlNil sets the value for OutboundWebhookUrl to be an explicit nil

### UnsetOutboundWebhookUrl
`func (o *MerchantAccount) UnsetOutboundWebhookUrl()`

UnsetOutboundWebhookUrl ensures that no value is present for OutboundWebhookUrl, not even an explicit nil
### GetOutboundWebhookUsername

`func (o *MerchantAccount) GetOutboundWebhookUsername() string`

GetOutboundWebhookUsername returns the OutboundWebhookUsername field if non-nil, zero value otherwise.

### GetOutboundWebhookUsernameOk

`func (o *MerchantAccount) GetOutboundWebhookUsernameOk() (*string, bool)`

GetOutboundWebhookUsernameOk returns a tuple with the OutboundWebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUsername

`func (o *MerchantAccount) SetOutboundWebhookUsername(v string)`

SetOutboundWebhookUsername sets OutboundWebhookUsername field to given value.

### HasOutboundWebhookUsername

`func (o *MerchantAccount) HasOutboundWebhookUsername() bool`

HasOutboundWebhookUsername returns a boolean if a field has been set.

### SetOutboundWebhookUsernameNil

`func (o *MerchantAccount) SetOutboundWebhookUsernameNil(b bool)`

 SetOutboundWebhookUsernameNil sets the value for OutboundWebhookUsername to be an explicit nil

### UnsetOutboundWebhookUsername
`func (o *MerchantAccount) UnsetOutboundWebhookUsername()`

UnsetOutboundWebhookUsername ensures that no value is present for OutboundWebhookUsername, not even an explicit nil
### GetOutboundWebhookPassword

`func (o *MerchantAccount) GetOutboundWebhookPassword() string`

GetOutboundWebhookPassword returns the OutboundWebhookPassword field if non-nil, zero value otherwise.

### GetOutboundWebhookPasswordOk

`func (o *MerchantAccount) GetOutboundWebhookPasswordOk() (*string, bool)`

GetOutboundWebhookPasswordOk returns a tuple with the OutboundWebhookPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookPassword

`func (o *MerchantAccount) SetOutboundWebhookPassword(v string)`

SetOutboundWebhookPassword sets OutboundWebhookPassword field to given value.

### HasOutboundWebhookPassword

`func (o *MerchantAccount) HasOutboundWebhookPassword() bool`

HasOutboundWebhookPassword returns a boolean if a field has been set.

### SetOutboundWebhookPasswordNil

`func (o *MerchantAccount) SetOutboundWebhookPasswordNil(b bool)`

 SetOutboundWebhookPasswordNil sets the value for OutboundWebhookPassword to be an explicit nil

### UnsetOutboundWebhookPassword
`func (o *MerchantAccount) UnsetOutboundWebhookPassword()`

UnsetOutboundWebhookPassword ensures that no value is present for OutboundWebhookPassword, not even an explicit nil
### GetVisaNetworkTokensRequestorId

`func (o *MerchantAccount) GetVisaNetworkTokensRequestorId() string`

GetVisaNetworkTokensRequestorId returns the VisaNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensRequestorIdOk

`func (o *MerchantAccount) GetVisaNetworkTokensRequestorIdOk() (*string, bool)`

GetVisaNetworkTokensRequestorIdOk returns a tuple with the VisaNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensRequestorId

`func (o *MerchantAccount) SetVisaNetworkTokensRequestorId(v string)`

SetVisaNetworkTokensRequestorId sets VisaNetworkTokensRequestorId field to given value.

### HasVisaNetworkTokensRequestorId

`func (o *MerchantAccount) HasVisaNetworkTokensRequestorId() bool`

HasVisaNetworkTokensRequestorId returns a boolean if a field has been set.

### SetVisaNetworkTokensRequestorIdNil

`func (o *MerchantAccount) SetVisaNetworkTokensRequestorIdNil(b bool)`

 SetVisaNetworkTokensRequestorIdNil sets the value for VisaNetworkTokensRequestorId to be an explicit nil

### UnsetVisaNetworkTokensRequestorId
`func (o *MerchantAccount) UnsetVisaNetworkTokensRequestorId()`

UnsetVisaNetworkTokensRequestorId ensures that no value is present for VisaNetworkTokensRequestorId, not even an explicit nil
### GetVisaNetworkTokensAppId

`func (o *MerchantAccount) GetVisaNetworkTokensAppId() string`

GetVisaNetworkTokensAppId returns the VisaNetworkTokensAppId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensAppIdOk

`func (o *MerchantAccount) GetVisaNetworkTokensAppIdOk() (*string, bool)`

GetVisaNetworkTokensAppIdOk returns a tuple with the VisaNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensAppId

`func (o *MerchantAccount) SetVisaNetworkTokensAppId(v string)`

SetVisaNetworkTokensAppId sets VisaNetworkTokensAppId field to given value.

### HasVisaNetworkTokensAppId

`func (o *MerchantAccount) HasVisaNetworkTokensAppId() bool`

HasVisaNetworkTokensAppId returns a boolean if a field has been set.

### SetVisaNetworkTokensAppIdNil

`func (o *MerchantAccount) SetVisaNetworkTokensAppIdNil(b bool)`

 SetVisaNetworkTokensAppIdNil sets the value for VisaNetworkTokensAppId to be an explicit nil

### UnsetVisaNetworkTokensAppId
`func (o *MerchantAccount) UnsetVisaNetworkTokensAppId()`

UnsetVisaNetworkTokensAppId ensures that no value is present for VisaNetworkTokensAppId, not even an explicit nil
### GetAmexNetworkTokensRequestorId

`func (o *MerchantAccount) GetAmexNetworkTokensRequestorId() string`

GetAmexNetworkTokensRequestorId returns the AmexNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensRequestorIdOk

`func (o *MerchantAccount) GetAmexNetworkTokensRequestorIdOk() (*string, bool)`

GetAmexNetworkTokensRequestorIdOk returns a tuple with the AmexNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensRequestorId

`func (o *MerchantAccount) SetAmexNetworkTokensRequestorId(v string)`

SetAmexNetworkTokensRequestorId sets AmexNetworkTokensRequestorId field to given value.

### HasAmexNetworkTokensRequestorId

`func (o *MerchantAccount) HasAmexNetworkTokensRequestorId() bool`

HasAmexNetworkTokensRequestorId returns a boolean if a field has been set.

### SetAmexNetworkTokensRequestorIdNil

`func (o *MerchantAccount) SetAmexNetworkTokensRequestorIdNil(b bool)`

 SetAmexNetworkTokensRequestorIdNil sets the value for AmexNetworkTokensRequestorId to be an explicit nil

### UnsetAmexNetworkTokensRequestorId
`func (o *MerchantAccount) UnsetAmexNetworkTokensRequestorId()`

UnsetAmexNetworkTokensRequestorId ensures that no value is present for AmexNetworkTokensRequestorId, not even an explicit nil
### GetAmexNetworkTokensAppId

`func (o *MerchantAccount) GetAmexNetworkTokensAppId() string`

GetAmexNetworkTokensAppId returns the AmexNetworkTokensAppId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensAppIdOk

`func (o *MerchantAccount) GetAmexNetworkTokensAppIdOk() (*string, bool)`

GetAmexNetworkTokensAppIdOk returns a tuple with the AmexNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensAppId

`func (o *MerchantAccount) SetAmexNetworkTokensAppId(v string)`

SetAmexNetworkTokensAppId sets AmexNetworkTokensAppId field to given value.

### HasAmexNetworkTokensAppId

`func (o *MerchantAccount) HasAmexNetworkTokensAppId() bool`

HasAmexNetworkTokensAppId returns a boolean if a field has been set.

### SetAmexNetworkTokensAppIdNil

`func (o *MerchantAccount) SetAmexNetworkTokensAppIdNil(b bool)`

 SetAmexNetworkTokensAppIdNil sets the value for AmexNetworkTokensAppId to be an explicit nil

### UnsetAmexNetworkTokensAppId
`func (o *MerchantAccount) UnsetAmexNetworkTokensAppId()`

UnsetAmexNetworkTokensAppId ensures that no value is present for AmexNetworkTokensAppId, not even an explicit nil
### GetMastercardNetworkTokensRequestorId

`func (o *MerchantAccount) GetMastercardNetworkTokensRequestorId() string`

GetMastercardNetworkTokensRequestorId returns the MastercardNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensRequestorIdOk

`func (o *MerchantAccount) GetMastercardNetworkTokensRequestorIdOk() (*string, bool)`

GetMastercardNetworkTokensRequestorIdOk returns a tuple with the MastercardNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensRequestorId

`func (o *MerchantAccount) SetMastercardNetworkTokensRequestorId(v string)`

SetMastercardNetworkTokensRequestorId sets MastercardNetworkTokensRequestorId field to given value.

### HasMastercardNetworkTokensRequestorId

`func (o *MerchantAccount) HasMastercardNetworkTokensRequestorId() bool`

HasMastercardNetworkTokensRequestorId returns a boolean if a field has been set.

### SetMastercardNetworkTokensRequestorIdNil

`func (o *MerchantAccount) SetMastercardNetworkTokensRequestorIdNil(b bool)`

 SetMastercardNetworkTokensRequestorIdNil sets the value for MastercardNetworkTokensRequestorId to be an explicit nil

### UnsetMastercardNetworkTokensRequestorId
`func (o *MerchantAccount) UnsetMastercardNetworkTokensRequestorId()`

UnsetMastercardNetworkTokensRequestorId ensures that no value is present for MastercardNetworkTokensRequestorId, not even an explicit nil
### GetMastercardNetworkTokensAppId

`func (o *MerchantAccount) GetMastercardNetworkTokensAppId() string`

GetMastercardNetworkTokensAppId returns the MastercardNetworkTokensAppId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensAppIdOk

`func (o *MerchantAccount) GetMastercardNetworkTokensAppIdOk() (*string, bool)`

GetMastercardNetworkTokensAppIdOk returns a tuple with the MastercardNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensAppId

`func (o *MerchantAccount) SetMastercardNetworkTokensAppId(v string)`

SetMastercardNetworkTokensAppId sets MastercardNetworkTokensAppId field to given value.

### HasMastercardNetworkTokensAppId

`func (o *MerchantAccount) HasMastercardNetworkTokensAppId() bool`

HasMastercardNetworkTokensAppId returns a boolean if a field has been set.

### SetMastercardNetworkTokensAppIdNil

`func (o *MerchantAccount) SetMastercardNetworkTokensAppIdNil(b bool)`

 SetMastercardNetworkTokensAppIdNil sets the value for MastercardNetworkTokensAppId to be an explicit nil

### UnsetMastercardNetworkTokensAppId
`func (o *MerchantAccount) UnsetMastercardNetworkTokensAppId()`

UnsetMastercardNetworkTokensAppId ensures that no value is present for MastercardNetworkTokensAppId, not even an explicit nil
### GetLoonClientKey

`func (o *MerchantAccount) GetLoonClientKey() string`

GetLoonClientKey returns the LoonClientKey field if non-nil, zero value otherwise.

### GetLoonClientKeyOk

`func (o *MerchantAccount) GetLoonClientKeyOk() (*string, bool)`

GetLoonClientKeyOk returns a tuple with the LoonClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonClientKey

`func (o *MerchantAccount) SetLoonClientKey(v string)`

SetLoonClientKey sets LoonClientKey field to given value.

### HasLoonClientKey

`func (o *MerchantAccount) HasLoonClientKey() bool`

HasLoonClientKey returns a boolean if a field has been set.

### SetLoonClientKeyNil

`func (o *MerchantAccount) SetLoonClientKeyNil(b bool)`

 SetLoonClientKeyNil sets the value for LoonClientKey to be an explicit nil

### UnsetLoonClientKey
`func (o *MerchantAccount) UnsetLoonClientKey()`

UnsetLoonClientKey ensures that no value is present for LoonClientKey, not even an explicit nil
### GetLoonSecretKey

`func (o *MerchantAccount) GetLoonSecretKey() string`

GetLoonSecretKey returns the LoonSecretKey field if non-nil, zero value otherwise.

### GetLoonSecretKeyOk

`func (o *MerchantAccount) GetLoonSecretKeyOk() (*string, bool)`

GetLoonSecretKeyOk returns a tuple with the LoonSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonSecretKey

`func (o *MerchantAccount) SetLoonSecretKey(v string)`

SetLoonSecretKey sets LoonSecretKey field to given value.

### HasLoonSecretKey

`func (o *MerchantAccount) HasLoonSecretKey() bool`

HasLoonSecretKey returns a boolean if a field has been set.

### SetLoonSecretKeyNil

`func (o *MerchantAccount) SetLoonSecretKeyNil(b bool)`

 SetLoonSecretKeyNil sets the value for LoonSecretKey to be an explicit nil

### UnsetLoonSecretKey
`func (o *MerchantAccount) UnsetLoonSecretKey()`

UnsetLoonSecretKey ensures that no value is present for LoonSecretKey, not even an explicit nil
### GetLoonAcceptedSchemes

`func (o *MerchantAccount) GetLoonAcceptedSchemes() []string`

GetLoonAcceptedSchemes returns the LoonAcceptedSchemes field if non-nil, zero value otherwise.

### GetLoonAcceptedSchemesOk

`func (o *MerchantAccount) GetLoonAcceptedSchemesOk() (*[]string, bool)`

GetLoonAcceptedSchemesOk returns a tuple with the LoonAcceptedSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoonAcceptedSchemes

`func (o *MerchantAccount) SetLoonAcceptedSchemes(v []string)`

SetLoonAcceptedSchemes sets LoonAcceptedSchemes field to given value.

### HasLoonAcceptedSchemes

`func (o *MerchantAccount) HasLoonAcceptedSchemes() bool`

HasLoonAcceptedSchemes returns a boolean if a field has been set.

### SetLoonAcceptedSchemesNil

`func (o *MerchantAccount) SetLoonAcceptedSchemesNil(b bool)`

 SetLoonAcceptedSchemesNil sets the value for LoonAcceptedSchemes to be an explicit nil

### UnsetLoonAcceptedSchemes
`func (o *MerchantAccount) UnsetLoonAcceptedSchemes()`

UnsetLoonAcceptedSchemes ensures that no value is present for LoonAcceptedSchemes, not even an explicit nil
### GetCreatedAt

`func (o *MerchantAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOverCaptureAmount

`func (o *MerchantAccount) GetOverCaptureAmount() float32`

GetOverCaptureAmount returns the OverCaptureAmount field if non-nil, zero value otherwise.

### GetOverCaptureAmountOk

`func (o *MerchantAccount) GetOverCaptureAmountOk() (*float32, bool)`

GetOverCaptureAmountOk returns a tuple with the OverCaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCaptureAmount

`func (o *MerchantAccount) SetOverCaptureAmount(v float32)`

SetOverCaptureAmount sets OverCaptureAmount field to given value.

### HasOverCaptureAmount

`func (o *MerchantAccount) HasOverCaptureAmount() bool`

HasOverCaptureAmount returns a boolean if a field has been set.

### SetOverCaptureAmountNil

`func (o *MerchantAccount) SetOverCaptureAmountNil(b bool)`

 SetOverCaptureAmountNil sets the value for OverCaptureAmount to be an explicit nil

### UnsetOverCaptureAmount
`func (o *MerchantAccount) UnsetOverCaptureAmount()`

UnsetOverCaptureAmount ensures that no value is present for OverCaptureAmount, not even an explicit nil
### GetOverCapturePercentage

`func (o *MerchantAccount) GetOverCapturePercentage() float32`

GetOverCapturePercentage returns the OverCapturePercentage field if non-nil, zero value otherwise.

### GetOverCapturePercentageOk

`func (o *MerchantAccount) GetOverCapturePercentageOk() (*float32, bool)`

GetOverCapturePercentageOk returns a tuple with the OverCapturePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapturePercentage

`func (o *MerchantAccount) SetOverCapturePercentage(v float32)`

SetOverCapturePercentage sets OverCapturePercentage field to given value.

### HasOverCapturePercentage

`func (o *MerchantAccount) HasOverCapturePercentage() bool`

HasOverCapturePercentage returns a boolean if a field has been set.

### SetOverCapturePercentageNil

`func (o *MerchantAccount) SetOverCapturePercentageNil(b bool)`

 SetOverCapturePercentageNil sets the value for OverCapturePercentage to be an explicit nil

### UnsetOverCapturePercentage
`func (o *MerchantAccount) UnsetOverCapturePercentage()`

UnsetOverCapturePercentage ensures that no value is present for OverCapturePercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


