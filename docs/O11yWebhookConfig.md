# O11yWebhookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**MaxAlerts** | Pointer to **int32** | MaxAlerts is the maximum number of alerts to be sent per webhook message. Alerts exceeding this threshold will be truncated. Setting this to 0 allows an unlimited number of alerts. | [optional] 
**Timeout** | Pointer to **int32** | Timeout is the maximum time allowed to invoke the webhook. Setting this to 0 does not impose a timeout. | [optional] 
**Url** | Pointer to **interface{}** |  | [optional] 
**UrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yWebhookConfig

`func NewO11yWebhookConfig() *O11yWebhookConfig`

NewO11yWebhookConfig instantiates a new O11yWebhookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yWebhookConfigWithDefaults

`func NewO11yWebhookConfigWithDefaults() *O11yWebhookConfig`

NewO11yWebhookConfigWithDefaults instantiates a new O11yWebhookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yWebhookConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yWebhookConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yWebhookConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yWebhookConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yWebhookConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yWebhookConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yWebhookConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yWebhookConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMaxAlerts

`func (o *O11yWebhookConfig) GetMaxAlerts() int32`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *O11yWebhookConfig) GetMaxAlertsOk() (*int32, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *O11yWebhookConfig) SetMaxAlerts(v int32)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *O11yWebhookConfig) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetTimeout

`func (o *O11yWebhookConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *O11yWebhookConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *O11yWebhookConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *O11yWebhookConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *O11yWebhookConfig) GetUrl() interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yWebhookConfig) GetUrlOk() (*interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yWebhookConfig) SetUrl(v interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yWebhookConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *O11yWebhookConfig) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *O11yWebhookConfig) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUrlFile

`func (o *O11yWebhookConfig) GetUrlFile() string`

GetUrlFile returns the UrlFile field if non-nil, zero value otherwise.

### GetUrlFileOk

`func (o *O11yWebhookConfig) GetUrlFileOk() (*string, bool)`

GetUrlFileOk returns a tuple with the UrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFile

`func (o *O11yWebhookConfig) SetUrlFile(v string)`

SetUrlFile sets UrlFile field to given value.

### HasUrlFile

`func (o *O11yWebhookConfig) HasUrlFile() bool`

HasUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


