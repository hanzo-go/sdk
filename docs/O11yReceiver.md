# O11yReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscordConfigs** | Pointer to [**[]O11yDiscordConfig**](O11yDiscordConfig.md) |  | [optional] 
**EmailConfigs** | Pointer to [**[]O11yEmailConfig**](O11yEmailConfig.md) |  | [optional] 
**IncidentioConfigs** | Pointer to [**[]O11yIncidentioConfig**](O11yIncidentioConfig.md) |  | [optional] 
**JiraConfigs** | Pointer to [**[]O11yJiraConfig**](O11yJiraConfig.md) |  | [optional] 
**MattermostConfigs** | Pointer to [**[]O11yMattermostConfig**](O11yMattermostConfig.md) |  | [optional] 
**MsteamsConfigs** | Pointer to [**[]O11yMSTeamsConfig**](O11yMSTeamsConfig.md) |  | [optional] 
**Msteamsv2Configs** | Pointer to [**[]O11yMSTeamsV2Config**](O11yMSTeamsV2Config.md) |  | [optional] 
**Name** | Pointer to **string** | A unique identifier for this receiver. | [optional] 
**OpsgenieConfigs** | Pointer to [**[]O11yOpsGenieConfig**](O11yOpsGenieConfig.md) |  | [optional] 
**PagerdutyConfigs** | Pointer to [**[]O11yPagerdutyConfig**](O11yPagerdutyConfig.md) |  | [optional] 
**PushoverConfigs** | Pointer to [**[]O11yPushoverConfig**](O11yPushoverConfig.md) |  | [optional] 
**RocketchatConfigs** | Pointer to [**[]O11yRocketchatConfig**](O11yRocketchatConfig.md) |  | [optional] 
**SlackConfigs** | Pointer to [**[]O11ySlackConfig**](O11ySlackConfig.md) |  | [optional] 
**SnsConfigs** | Pointer to [**[]O11ySNSConfig**](O11ySNSConfig.md) |  | [optional] 
**TelegramConfigs** | Pointer to [**[]O11yTelegramConfig**](O11yTelegramConfig.md) |  | [optional] 
**VictoropsConfigs** | Pointer to [**[]O11yVictorOpsConfig**](O11yVictorOpsConfig.md) |  | [optional] 
**WebexConfigs** | Pointer to [**[]O11yWebexConfig**](O11yWebexConfig.md) |  | [optional] 
**WebhookConfigs** | Pointer to [**[]O11yWebhookConfig**](O11yWebhookConfig.md) |  | [optional] 
**WechatConfigs** | Pointer to [**[]O11yWechatConfig**](O11yWechatConfig.md) |  | [optional] 

## Methods

### NewO11yReceiver

`func NewO11yReceiver() *O11yReceiver`

NewO11yReceiver instantiates a new O11yReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yReceiverWithDefaults

`func NewO11yReceiverWithDefaults() *O11yReceiver`

NewO11yReceiverWithDefaults instantiates a new O11yReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscordConfigs

`func (o *O11yReceiver) GetDiscordConfigs() []O11yDiscordConfig`

GetDiscordConfigs returns the DiscordConfigs field if non-nil, zero value otherwise.

### GetDiscordConfigsOk

`func (o *O11yReceiver) GetDiscordConfigsOk() (*[]O11yDiscordConfig, bool)`

GetDiscordConfigsOk returns a tuple with the DiscordConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordConfigs

`func (o *O11yReceiver) SetDiscordConfigs(v []O11yDiscordConfig)`

SetDiscordConfigs sets DiscordConfigs field to given value.

### HasDiscordConfigs

`func (o *O11yReceiver) HasDiscordConfigs() bool`

HasDiscordConfigs returns a boolean if a field has been set.

### GetEmailConfigs

`func (o *O11yReceiver) GetEmailConfigs() []O11yEmailConfig`

GetEmailConfigs returns the EmailConfigs field if non-nil, zero value otherwise.

### GetEmailConfigsOk

`func (o *O11yReceiver) GetEmailConfigsOk() (*[]O11yEmailConfig, bool)`

GetEmailConfigsOk returns a tuple with the EmailConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfigs

`func (o *O11yReceiver) SetEmailConfigs(v []O11yEmailConfig)`

SetEmailConfigs sets EmailConfigs field to given value.

### HasEmailConfigs

`func (o *O11yReceiver) HasEmailConfigs() bool`

HasEmailConfigs returns a boolean if a field has been set.

### GetIncidentioConfigs

`func (o *O11yReceiver) GetIncidentioConfigs() []O11yIncidentioConfig`

GetIncidentioConfigs returns the IncidentioConfigs field if non-nil, zero value otherwise.

### GetIncidentioConfigsOk

`func (o *O11yReceiver) GetIncidentioConfigsOk() (*[]O11yIncidentioConfig, bool)`

GetIncidentioConfigsOk returns a tuple with the IncidentioConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentioConfigs

`func (o *O11yReceiver) SetIncidentioConfigs(v []O11yIncidentioConfig)`

SetIncidentioConfigs sets IncidentioConfigs field to given value.

### HasIncidentioConfigs

`func (o *O11yReceiver) HasIncidentioConfigs() bool`

HasIncidentioConfigs returns a boolean if a field has been set.

### GetJiraConfigs

`func (o *O11yReceiver) GetJiraConfigs() []O11yJiraConfig`

GetJiraConfigs returns the JiraConfigs field if non-nil, zero value otherwise.

### GetJiraConfigsOk

`func (o *O11yReceiver) GetJiraConfigsOk() (*[]O11yJiraConfig, bool)`

GetJiraConfigsOk returns a tuple with the JiraConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraConfigs

`func (o *O11yReceiver) SetJiraConfigs(v []O11yJiraConfig)`

SetJiraConfigs sets JiraConfigs field to given value.

### HasJiraConfigs

`func (o *O11yReceiver) HasJiraConfigs() bool`

HasJiraConfigs returns a boolean if a field has been set.

### GetMattermostConfigs

`func (o *O11yReceiver) GetMattermostConfigs() []O11yMattermostConfig`

GetMattermostConfigs returns the MattermostConfigs field if non-nil, zero value otherwise.

### GetMattermostConfigsOk

`func (o *O11yReceiver) GetMattermostConfigsOk() (*[]O11yMattermostConfig, bool)`

GetMattermostConfigsOk returns a tuple with the MattermostConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMattermostConfigs

`func (o *O11yReceiver) SetMattermostConfigs(v []O11yMattermostConfig)`

SetMattermostConfigs sets MattermostConfigs field to given value.

### HasMattermostConfigs

`func (o *O11yReceiver) HasMattermostConfigs() bool`

HasMattermostConfigs returns a boolean if a field has been set.

### GetMsteamsConfigs

`func (o *O11yReceiver) GetMsteamsConfigs() []O11yMSTeamsConfig`

GetMsteamsConfigs returns the MsteamsConfigs field if non-nil, zero value otherwise.

### GetMsteamsConfigsOk

`func (o *O11yReceiver) GetMsteamsConfigsOk() (*[]O11yMSTeamsConfig, bool)`

GetMsteamsConfigsOk returns a tuple with the MsteamsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsConfigs

`func (o *O11yReceiver) SetMsteamsConfigs(v []O11yMSTeamsConfig)`

SetMsteamsConfigs sets MsteamsConfigs field to given value.

### HasMsteamsConfigs

`func (o *O11yReceiver) HasMsteamsConfigs() bool`

HasMsteamsConfigs returns a boolean if a field has been set.

### GetMsteamsv2Configs

`func (o *O11yReceiver) GetMsteamsv2Configs() []O11yMSTeamsV2Config`

GetMsteamsv2Configs returns the Msteamsv2Configs field if non-nil, zero value otherwise.

### GetMsteamsv2ConfigsOk

`func (o *O11yReceiver) GetMsteamsv2ConfigsOk() (*[]O11yMSTeamsV2Config, bool)`

GetMsteamsv2ConfigsOk returns a tuple with the Msteamsv2Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsv2Configs

`func (o *O11yReceiver) SetMsteamsv2Configs(v []O11yMSTeamsV2Config)`

SetMsteamsv2Configs sets Msteamsv2Configs field to given value.

### HasMsteamsv2Configs

`func (o *O11yReceiver) HasMsteamsv2Configs() bool`

HasMsteamsv2Configs returns a boolean if a field has been set.

### GetName

`func (o *O11yReceiver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yReceiver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yReceiver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yReceiver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpsgenieConfigs

`func (o *O11yReceiver) GetOpsgenieConfigs() []O11yOpsGenieConfig`

GetOpsgenieConfigs returns the OpsgenieConfigs field if non-nil, zero value otherwise.

### GetOpsgenieConfigsOk

`func (o *O11yReceiver) GetOpsgenieConfigsOk() (*[]O11yOpsGenieConfig, bool)`

GetOpsgenieConfigsOk returns a tuple with the OpsgenieConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieConfigs

`func (o *O11yReceiver) SetOpsgenieConfigs(v []O11yOpsGenieConfig)`

SetOpsgenieConfigs sets OpsgenieConfigs field to given value.

### HasOpsgenieConfigs

`func (o *O11yReceiver) HasOpsgenieConfigs() bool`

HasOpsgenieConfigs returns a boolean if a field has been set.

### GetPagerdutyConfigs

`func (o *O11yReceiver) GetPagerdutyConfigs() []O11yPagerdutyConfig`

GetPagerdutyConfigs returns the PagerdutyConfigs field if non-nil, zero value otherwise.

### GetPagerdutyConfigsOk

`func (o *O11yReceiver) GetPagerdutyConfigsOk() (*[]O11yPagerdutyConfig, bool)`

GetPagerdutyConfigsOk returns a tuple with the PagerdutyConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerdutyConfigs

`func (o *O11yReceiver) SetPagerdutyConfigs(v []O11yPagerdutyConfig)`

SetPagerdutyConfigs sets PagerdutyConfigs field to given value.

### HasPagerdutyConfigs

`func (o *O11yReceiver) HasPagerdutyConfigs() bool`

HasPagerdutyConfigs returns a boolean if a field has been set.

### GetPushoverConfigs

`func (o *O11yReceiver) GetPushoverConfigs() []O11yPushoverConfig`

GetPushoverConfigs returns the PushoverConfigs field if non-nil, zero value otherwise.

### GetPushoverConfigsOk

`func (o *O11yReceiver) GetPushoverConfigsOk() (*[]O11yPushoverConfig, bool)`

GetPushoverConfigsOk returns a tuple with the PushoverConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushoverConfigs

`func (o *O11yReceiver) SetPushoverConfigs(v []O11yPushoverConfig)`

SetPushoverConfigs sets PushoverConfigs field to given value.

### HasPushoverConfigs

`func (o *O11yReceiver) HasPushoverConfigs() bool`

HasPushoverConfigs returns a boolean if a field has been set.

### GetRocketchatConfigs

`func (o *O11yReceiver) GetRocketchatConfigs() []O11yRocketchatConfig`

GetRocketchatConfigs returns the RocketchatConfigs field if non-nil, zero value otherwise.

### GetRocketchatConfigsOk

`func (o *O11yReceiver) GetRocketchatConfigsOk() (*[]O11yRocketchatConfig, bool)`

GetRocketchatConfigsOk returns a tuple with the RocketchatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocketchatConfigs

`func (o *O11yReceiver) SetRocketchatConfigs(v []O11yRocketchatConfig)`

SetRocketchatConfigs sets RocketchatConfigs field to given value.

### HasRocketchatConfigs

`func (o *O11yReceiver) HasRocketchatConfigs() bool`

HasRocketchatConfigs returns a boolean if a field has been set.

### GetSlackConfigs

`func (o *O11yReceiver) GetSlackConfigs() []O11ySlackConfig`

GetSlackConfigs returns the SlackConfigs field if non-nil, zero value otherwise.

### GetSlackConfigsOk

`func (o *O11yReceiver) GetSlackConfigsOk() (*[]O11ySlackConfig, bool)`

GetSlackConfigsOk returns a tuple with the SlackConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigs

`func (o *O11yReceiver) SetSlackConfigs(v []O11ySlackConfig)`

SetSlackConfigs sets SlackConfigs field to given value.

### HasSlackConfigs

`func (o *O11yReceiver) HasSlackConfigs() bool`

HasSlackConfigs returns a boolean if a field has been set.

### GetSnsConfigs

`func (o *O11yReceiver) GetSnsConfigs() []O11ySNSConfig`

GetSnsConfigs returns the SnsConfigs field if non-nil, zero value otherwise.

### GetSnsConfigsOk

`func (o *O11yReceiver) GetSnsConfigsOk() (*[]O11ySNSConfig, bool)`

GetSnsConfigsOk returns a tuple with the SnsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnsConfigs

`func (o *O11yReceiver) SetSnsConfigs(v []O11ySNSConfig)`

SetSnsConfigs sets SnsConfigs field to given value.

### HasSnsConfigs

`func (o *O11yReceiver) HasSnsConfigs() bool`

HasSnsConfigs returns a boolean if a field has been set.

### GetTelegramConfigs

`func (o *O11yReceiver) GetTelegramConfigs() []O11yTelegramConfig`

GetTelegramConfigs returns the TelegramConfigs field if non-nil, zero value otherwise.

### GetTelegramConfigsOk

`func (o *O11yReceiver) GetTelegramConfigsOk() (*[]O11yTelegramConfig, bool)`

GetTelegramConfigsOk returns a tuple with the TelegramConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramConfigs

`func (o *O11yReceiver) SetTelegramConfigs(v []O11yTelegramConfig)`

SetTelegramConfigs sets TelegramConfigs field to given value.

### HasTelegramConfigs

`func (o *O11yReceiver) HasTelegramConfigs() bool`

HasTelegramConfigs returns a boolean if a field has been set.

### GetVictoropsConfigs

`func (o *O11yReceiver) GetVictoropsConfigs() []O11yVictorOpsConfig`

GetVictoropsConfigs returns the VictoropsConfigs field if non-nil, zero value otherwise.

### GetVictoropsConfigsOk

`func (o *O11yReceiver) GetVictoropsConfigsOk() (*[]O11yVictorOpsConfig, bool)`

GetVictoropsConfigsOk returns a tuple with the VictoropsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictoropsConfigs

`func (o *O11yReceiver) SetVictoropsConfigs(v []O11yVictorOpsConfig)`

SetVictoropsConfigs sets VictoropsConfigs field to given value.

### HasVictoropsConfigs

`func (o *O11yReceiver) HasVictoropsConfigs() bool`

HasVictoropsConfigs returns a boolean if a field has been set.

### GetWebexConfigs

`func (o *O11yReceiver) GetWebexConfigs() []O11yWebexConfig`

GetWebexConfigs returns the WebexConfigs field if non-nil, zero value otherwise.

### GetWebexConfigsOk

`func (o *O11yReceiver) GetWebexConfigsOk() (*[]O11yWebexConfig, bool)`

GetWebexConfigsOk returns a tuple with the WebexConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebexConfigs

`func (o *O11yReceiver) SetWebexConfigs(v []O11yWebexConfig)`

SetWebexConfigs sets WebexConfigs field to given value.

### HasWebexConfigs

`func (o *O11yReceiver) HasWebexConfigs() bool`

HasWebexConfigs returns a boolean if a field has been set.

### GetWebhookConfigs

`func (o *O11yReceiver) GetWebhookConfigs() []O11yWebhookConfig`

GetWebhookConfigs returns the WebhookConfigs field if non-nil, zero value otherwise.

### GetWebhookConfigsOk

`func (o *O11yReceiver) GetWebhookConfigsOk() (*[]O11yWebhookConfig, bool)`

GetWebhookConfigsOk returns a tuple with the WebhookConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfigs

`func (o *O11yReceiver) SetWebhookConfigs(v []O11yWebhookConfig)`

SetWebhookConfigs sets WebhookConfigs field to given value.

### HasWebhookConfigs

`func (o *O11yReceiver) HasWebhookConfigs() bool`

HasWebhookConfigs returns a boolean if a field has been set.

### GetWechatConfigs

`func (o *O11yReceiver) GetWechatConfigs() []O11yWechatConfig`

GetWechatConfigs returns the WechatConfigs field if non-nil, zero value otherwise.

### GetWechatConfigsOk

`func (o *O11yReceiver) GetWechatConfigsOk() (*[]O11yWechatConfig, bool)`

GetWechatConfigsOk returns a tuple with the WechatConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWechatConfigs

`func (o *O11yReceiver) SetWechatConfigs(v []O11yWechatConfig)`

SetWechatConfigs sets WechatConfigs field to given value.

### HasWechatConfigs

`func (o *O11yReceiver) HasWechatConfigs() bool`

HasWechatConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


