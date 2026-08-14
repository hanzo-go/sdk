# O11yIncidentioConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**AlertSourceToken** | Pointer to **interface{}** |  | [optional] 
**AlertSourceTokenFile** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**MaxAlerts** | Pointer to **int32** | MaxAlerts is the maximum number of alerts to be sent per incident.io message. Alerts exceeding this threshold will be truncated. Setting this to 0 allows an unlimited number of alerts. Note that if the payload exceeds incident.io&#39;s size limits, you will receive a 429 response and alerts will not be ingested. | [optional] 
**Timeout** | Pointer to **int32** | Timeout is the maximum time allowed to invoke incident.io. Setting this to 0 does not impose a timeout. | [optional] 
**Url** | Pointer to **interface{}** |  | [optional] 
**UrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yIncidentioConfig

`func NewO11yIncidentioConfig() *O11yIncidentioConfig`

NewO11yIncidentioConfig instantiates a new O11yIncidentioConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIncidentioConfigWithDefaults

`func NewO11yIncidentioConfigWithDefaults() *O11yIncidentioConfig`

NewO11yIncidentioConfigWithDefaults instantiates a new O11yIncidentioConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yIncidentioConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yIncidentioConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yIncidentioConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yIncidentioConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetAlertSourceToken

`func (o *O11yIncidentioConfig) GetAlertSourceToken() interface{}`

GetAlertSourceToken returns the AlertSourceToken field if non-nil, zero value otherwise.

### GetAlertSourceTokenOk

`func (o *O11yIncidentioConfig) GetAlertSourceTokenOk() (*interface{}, bool)`

GetAlertSourceTokenOk returns a tuple with the AlertSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceToken

`func (o *O11yIncidentioConfig) SetAlertSourceToken(v interface{})`

SetAlertSourceToken sets AlertSourceToken field to given value.

### HasAlertSourceToken

`func (o *O11yIncidentioConfig) HasAlertSourceToken() bool`

HasAlertSourceToken returns a boolean if a field has been set.

### SetAlertSourceTokenNil

`func (o *O11yIncidentioConfig) SetAlertSourceTokenNil(b bool)`

 SetAlertSourceTokenNil sets the value for AlertSourceToken to be an explicit nil

### UnsetAlertSourceToken
`func (o *O11yIncidentioConfig) UnsetAlertSourceToken()`

UnsetAlertSourceToken ensures that no value is present for AlertSourceToken, not even an explicit nil
### GetAlertSourceTokenFile

`func (o *O11yIncidentioConfig) GetAlertSourceTokenFile() string`

GetAlertSourceTokenFile returns the AlertSourceTokenFile field if non-nil, zero value otherwise.

### GetAlertSourceTokenFileOk

`func (o *O11yIncidentioConfig) GetAlertSourceTokenFileOk() (*string, bool)`

GetAlertSourceTokenFileOk returns a tuple with the AlertSourceTokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceTokenFile

`func (o *O11yIncidentioConfig) SetAlertSourceTokenFile(v string)`

SetAlertSourceTokenFile sets AlertSourceTokenFile field to given value.

### HasAlertSourceTokenFile

`func (o *O11yIncidentioConfig) HasAlertSourceTokenFile() bool`

HasAlertSourceTokenFile returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yIncidentioConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yIncidentioConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yIncidentioConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yIncidentioConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMaxAlerts

`func (o *O11yIncidentioConfig) GetMaxAlerts() int32`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *O11yIncidentioConfig) GetMaxAlertsOk() (*int32, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *O11yIncidentioConfig) SetMaxAlerts(v int32)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *O11yIncidentioConfig) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetTimeout

`func (o *O11yIncidentioConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *O11yIncidentioConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *O11yIncidentioConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *O11yIncidentioConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *O11yIncidentioConfig) GetUrl() interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yIncidentioConfig) GetUrlOk() (*interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yIncidentioConfig) SetUrl(v interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yIncidentioConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *O11yIncidentioConfig) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *O11yIncidentioConfig) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUrlFile

`func (o *O11yIncidentioConfig) GetUrlFile() string`

GetUrlFile returns the UrlFile field if non-nil, zero value otherwise.

### GetUrlFileOk

`func (o *O11yIncidentioConfig) GetUrlFileOk() (*string, bool)`

GetUrlFileOk returns a tuple with the UrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFile

`func (o *O11yIncidentioConfig) SetUrlFile(v string)`

SetUrlFile sets UrlFile field to given value.

### HasUrlFile

`func (o *O11yIncidentioConfig) HasUrlFile() bool`

HasUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


