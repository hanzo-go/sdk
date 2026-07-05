# O11yNotificationChannelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Webhook URL or Slack incoming webhook | [optional] 
**Channel** | Pointer to **string** | Slack channel name | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**RoutingKey** | Pointer to **string** | PagerDuty routing key | [optional] 
**SeverityMap** | Pointer to **map[string]string** | Map alert severity to PagerDuty severity | [optional] 

## Methods

### NewO11yNotificationChannelConfig

`func NewO11yNotificationChannelConfig() *O11yNotificationChannelConfig`

NewO11yNotificationChannelConfig instantiates a new O11yNotificationChannelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNotificationChannelConfigWithDefaults

`func NewO11yNotificationChannelConfigWithDefaults() *O11yNotificationChannelConfig`

NewO11yNotificationChannelConfigWithDefaults instantiates a new O11yNotificationChannelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *O11yNotificationChannelConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yNotificationChannelConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yNotificationChannelConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yNotificationChannelConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetChannel

`func (o *O11yNotificationChannelConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *O11yNotificationChannelConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *O11yNotificationChannelConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *O11yNotificationChannelConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEmail

`func (o *O11yNotificationChannelConfig) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yNotificationChannelConfig) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yNotificationChannelConfig) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yNotificationChannelConfig) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRoutingKey

`func (o *O11yNotificationChannelConfig) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *O11yNotificationChannelConfig) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *O11yNotificationChannelConfig) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *O11yNotificationChannelConfig) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetSeverityMap

`func (o *O11yNotificationChannelConfig) GetSeverityMap() map[string]string`

GetSeverityMap returns the SeverityMap field if non-nil, zero value otherwise.

### GetSeverityMapOk

`func (o *O11yNotificationChannelConfig) GetSeverityMapOk() (*map[string]string, bool)`

GetSeverityMapOk returns a tuple with the SeverityMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMap

`func (o *O11yNotificationChannelConfig) SetSeverityMap(v map[string]string)`

SetSeverityMap sets SeverityMap field to given value.

### HasSeverityMap

`func (o *O11yNotificationChannelConfig) HasSeverityMap() bool`

HasSeverityMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


