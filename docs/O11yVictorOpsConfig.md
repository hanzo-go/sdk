# O11yVictorOpsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**ApiKey** | Pointer to **interface{}** |  | [optional] 
**ApiKeyFile** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 
**EntityDisplayName** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**MonitoringTool** | Pointer to **string** |  | [optional] 
**RoutingKey** | Pointer to **string** |  | [optional] 
**StateMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yVictorOpsConfig

`func NewO11yVictorOpsConfig() *O11yVictorOpsConfig`

NewO11yVictorOpsConfig instantiates a new O11yVictorOpsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yVictorOpsConfigWithDefaults

`func NewO11yVictorOpsConfigWithDefaults() *O11yVictorOpsConfig`

NewO11yVictorOpsConfigWithDefaults instantiates a new O11yVictorOpsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yVictorOpsConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yVictorOpsConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yVictorOpsConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yVictorOpsConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetApiKey

`func (o *O11yVictorOpsConfig) GetApiKey() interface{}`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *O11yVictorOpsConfig) GetApiKeyOk() (*interface{}, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *O11yVictorOpsConfig) SetApiKey(v interface{})`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *O11yVictorOpsConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *O11yVictorOpsConfig) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *O11yVictorOpsConfig) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetApiKeyFile

`func (o *O11yVictorOpsConfig) GetApiKeyFile() string`

GetApiKeyFile returns the ApiKeyFile field if non-nil, zero value otherwise.

### GetApiKeyFileOk

`func (o *O11yVictorOpsConfig) GetApiKeyFileOk() (*string, bool)`

GetApiKeyFileOk returns a tuple with the ApiKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyFile

`func (o *O11yVictorOpsConfig) SetApiKeyFile(v string)`

SetApiKeyFile sets ApiKeyFile field to given value.

### HasApiKeyFile

`func (o *O11yVictorOpsConfig) HasApiKeyFile() bool`

HasApiKeyFile returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yVictorOpsConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yVictorOpsConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yVictorOpsConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yVictorOpsConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yVictorOpsConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yVictorOpsConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetCustomFields

`func (o *O11yVictorOpsConfig) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *O11yVictorOpsConfig) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *O11yVictorOpsConfig) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *O11yVictorOpsConfig) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityDisplayName

`func (o *O11yVictorOpsConfig) GetEntityDisplayName() string`

GetEntityDisplayName returns the EntityDisplayName field if non-nil, zero value otherwise.

### GetEntityDisplayNameOk

`func (o *O11yVictorOpsConfig) GetEntityDisplayNameOk() (*string, bool)`

GetEntityDisplayNameOk returns a tuple with the EntityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDisplayName

`func (o *O11yVictorOpsConfig) SetEntityDisplayName(v string)`

SetEntityDisplayName sets EntityDisplayName field to given value.

### HasEntityDisplayName

`func (o *O11yVictorOpsConfig) HasEntityDisplayName() bool`

HasEntityDisplayName returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yVictorOpsConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yVictorOpsConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yVictorOpsConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yVictorOpsConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessageType

`func (o *O11yVictorOpsConfig) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *O11yVictorOpsConfig) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *O11yVictorOpsConfig) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *O11yVictorOpsConfig) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetMonitoringTool

`func (o *O11yVictorOpsConfig) GetMonitoringTool() string`

GetMonitoringTool returns the MonitoringTool field if non-nil, zero value otherwise.

### GetMonitoringToolOk

`func (o *O11yVictorOpsConfig) GetMonitoringToolOk() (*string, bool)`

GetMonitoringToolOk returns a tuple with the MonitoringTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTool

`func (o *O11yVictorOpsConfig) SetMonitoringTool(v string)`

SetMonitoringTool sets MonitoringTool field to given value.

### HasMonitoringTool

`func (o *O11yVictorOpsConfig) HasMonitoringTool() bool`

HasMonitoringTool returns a boolean if a field has been set.

### GetRoutingKey

`func (o *O11yVictorOpsConfig) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *O11yVictorOpsConfig) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *O11yVictorOpsConfig) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *O11yVictorOpsConfig) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetStateMessage

`func (o *O11yVictorOpsConfig) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *O11yVictorOpsConfig) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *O11yVictorOpsConfig) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *O11yVictorOpsConfig) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


