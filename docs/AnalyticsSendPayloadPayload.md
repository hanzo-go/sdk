# AnalyticsSendPayloadPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Website** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** |  | [optional] 
**Screen** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** | Unix timestamp in seconds | [optional] 
**Id** | Pointer to **string** | Distinct user identifier | [optional] 

## Methods

### NewAnalyticsSendPayloadPayload

`func NewAnalyticsSendPayloadPayload(website string, ) *AnalyticsSendPayloadPayload`

NewAnalyticsSendPayloadPayload instantiates a new AnalyticsSendPayloadPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsSendPayloadPayloadWithDefaults

`func NewAnalyticsSendPayloadPayloadWithDefaults() *AnalyticsSendPayloadPayload`

NewAnalyticsSendPayloadPayloadWithDefaults instantiates a new AnalyticsSendPayloadPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsite

`func (o *AnalyticsSendPayloadPayload) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AnalyticsSendPayloadPayload) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AnalyticsSendPayloadPayload) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetData

`func (o *AnalyticsSendPayloadPayload) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsSendPayloadPayload) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsSendPayloadPayload) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AnalyticsSendPayloadPayload) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHostname

`func (o *AnalyticsSendPayloadPayload) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AnalyticsSendPayloadPayload) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AnalyticsSendPayloadPayload) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AnalyticsSendPayloadPayload) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLanguage

`func (o *AnalyticsSendPayloadPayload) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AnalyticsSendPayloadPayload) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AnalyticsSendPayloadPayload) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AnalyticsSendPayloadPayload) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetReferrer

`func (o *AnalyticsSendPayloadPayload) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *AnalyticsSendPayloadPayload) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *AnalyticsSendPayloadPayload) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *AnalyticsSendPayloadPayload) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetScreen

`func (o *AnalyticsSendPayloadPayload) GetScreen() string`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *AnalyticsSendPayloadPayload) GetScreenOk() (*string, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *AnalyticsSendPayloadPayload) SetScreen(v string)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *AnalyticsSendPayloadPayload) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetTitle

`func (o *AnalyticsSendPayloadPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnalyticsSendPayloadPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnalyticsSendPayloadPayload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AnalyticsSendPayloadPayload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AnalyticsSendPayloadPayload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AnalyticsSendPayloadPayload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AnalyticsSendPayloadPayload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AnalyticsSendPayloadPayload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *AnalyticsSendPayloadPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyticsSendPayloadPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyticsSendPayloadPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnalyticsSendPayloadPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTag

`func (o *AnalyticsSendPayloadPayload) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AnalyticsSendPayloadPayload) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AnalyticsSendPayloadPayload) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AnalyticsSendPayloadPayload) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetIp

`func (o *AnalyticsSendPayloadPayload) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AnalyticsSendPayloadPayload) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AnalyticsSendPayloadPayload) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AnalyticsSendPayloadPayload) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *AnalyticsSendPayloadPayload) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AnalyticsSendPayloadPayload) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AnalyticsSendPayloadPayload) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AnalyticsSendPayloadPayload) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetTimestamp

`func (o *AnalyticsSendPayloadPayload) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AnalyticsSendPayloadPayload) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AnalyticsSendPayloadPayload) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AnalyticsSendPayloadPayload) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetId

`func (o *AnalyticsSendPayloadPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsSendPayloadPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsSendPayloadPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsSendPayloadPayload) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


