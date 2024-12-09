# BrowserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JavaEnabled** | **bool** | Indicates whether the client browser supports Java. | 
**JavascriptEnabled** | **bool** | Indicates whether the client browser supports JavaScript. | 
**Language** | **string** | The preferred language of the buyer, usually the language of the browser UI. | 
**ColorDepth** | **float32** | The color depth of the screen. | 
**ScreenHeight** | **float32** | The height of the screen in pixels. | 
**ScreenWidth** | **float32** | The width of the screen in pixels. | 
**TimeZoneOffset** | **float32** | Time-zone offset in minutes between UTC and buyer location. | 
**UserDevice** | **string** | The platform that is being used to access the website. | 
**UserAgent** | **string** | The user agent string for the current browser. | 
**AcceptHeader** | Pointer to **string** | The &#x60;Accept&#x60; header of the request from the buyer&#39;s browser. | [optional] 

## Methods

### NewBrowserInfo

`func NewBrowserInfo(javaEnabled bool, javascriptEnabled bool, language string, colorDepth float32, screenHeight float32, screenWidth float32, timeZoneOffset float32, userDevice string, userAgent string, ) *BrowserInfo`

NewBrowserInfo instantiates a new BrowserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserInfoWithDefaults

`func NewBrowserInfoWithDefaults() *BrowserInfo`

NewBrowserInfoWithDefaults instantiates a new BrowserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJavaEnabled

`func (o *BrowserInfo) GetJavaEnabled() bool`

GetJavaEnabled returns the JavaEnabled field if non-nil, zero value otherwise.

### GetJavaEnabledOk

`func (o *BrowserInfo) GetJavaEnabledOk() (*bool, bool)`

GetJavaEnabledOk returns a tuple with the JavaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaEnabled

`func (o *BrowserInfo) SetJavaEnabled(v bool)`

SetJavaEnabled sets JavaEnabled field to given value.


### GetJavascriptEnabled

`func (o *BrowserInfo) GetJavascriptEnabled() bool`

GetJavascriptEnabled returns the JavascriptEnabled field if non-nil, zero value otherwise.

### GetJavascriptEnabledOk

`func (o *BrowserInfo) GetJavascriptEnabledOk() (*bool, bool)`

GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptEnabled

`func (o *BrowserInfo) SetJavascriptEnabled(v bool)`

SetJavascriptEnabled sets JavascriptEnabled field to given value.


### GetLanguage

`func (o *BrowserInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BrowserInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BrowserInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetColorDepth

`func (o *BrowserInfo) GetColorDepth() float32`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *BrowserInfo) GetColorDepthOk() (*float32, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *BrowserInfo) SetColorDepth(v float32)`

SetColorDepth sets ColorDepth field to given value.


### GetScreenHeight

`func (o *BrowserInfo) GetScreenHeight() float32`

GetScreenHeight returns the ScreenHeight field if non-nil, zero value otherwise.

### GetScreenHeightOk

`func (o *BrowserInfo) GetScreenHeightOk() (*float32, bool)`

GetScreenHeightOk returns a tuple with the ScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenHeight

`func (o *BrowserInfo) SetScreenHeight(v float32)`

SetScreenHeight sets ScreenHeight field to given value.


### GetScreenWidth

`func (o *BrowserInfo) GetScreenWidth() float32`

GetScreenWidth returns the ScreenWidth field if non-nil, zero value otherwise.

### GetScreenWidthOk

`func (o *BrowserInfo) GetScreenWidthOk() (*float32, bool)`

GetScreenWidthOk returns a tuple with the ScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenWidth

`func (o *BrowserInfo) SetScreenWidth(v float32)`

SetScreenWidth sets ScreenWidth field to given value.


### GetTimeZoneOffset

`func (o *BrowserInfo) GetTimeZoneOffset() float32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *BrowserInfo) GetTimeZoneOffsetOk() (*float32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *BrowserInfo) SetTimeZoneOffset(v float32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.


### GetUserDevice

`func (o *BrowserInfo) GetUserDevice() string`

GetUserDevice returns the UserDevice field if non-nil, zero value otherwise.

### GetUserDeviceOk

`func (o *BrowserInfo) GetUserDeviceOk() (*string, bool)`

GetUserDeviceOk returns a tuple with the UserDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDevice

`func (o *BrowserInfo) SetUserDevice(v string)`

SetUserDevice sets UserDevice field to given value.


### GetUserAgent

`func (o *BrowserInfo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *BrowserInfo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *BrowserInfo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetAcceptHeader

`func (o *BrowserInfo) GetAcceptHeader() string`

GetAcceptHeader returns the AcceptHeader field if non-nil, zero value otherwise.

### GetAcceptHeaderOk

`func (o *BrowserInfo) GetAcceptHeaderOk() (*string, bool)`

GetAcceptHeaderOk returns a tuple with the AcceptHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHeader

`func (o *BrowserInfo) SetAcceptHeader(v string)`

SetAcceptHeader sets AcceptHeader field to given value.

### HasAcceptHeader

`func (o *BrowserInfo) HasAcceptHeader() bool`

HasAcceptHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


