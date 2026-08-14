# O11yO11yGlobalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiAssistantUrl** | Pointer to **string** | AIAssistantURL is the AI assistant endpoint, when one is exposed. | [optional] 
**ExternalUrl** | Pointer to **string** | ExternalURL is the deployment&#39;s public URL. | [optional] 
**IdentN** | Pointer to [**O11yO11yIdentN**](O11yO11yIdentN.md) | IdentN says which identity providers are enabled. | [optional] 
**IngestionUrl** | Pointer to **string** | IngestionURL is where telemetry is sent. | [optional] 
**McpUrl** | Pointer to **string** | MCPURL is the MCP endpoint, when one is exposed. | [optional] 

## Methods

### NewO11yO11yGlobalConfig

`func NewO11yO11yGlobalConfig() *O11yO11yGlobalConfig`

NewO11yO11yGlobalConfig instantiates a new O11yO11yGlobalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yGlobalConfigWithDefaults

`func NewO11yO11yGlobalConfigWithDefaults() *O11yO11yGlobalConfig`

NewO11yO11yGlobalConfigWithDefaults instantiates a new O11yO11yGlobalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiAssistantUrl

`func (o *O11yO11yGlobalConfig) GetAiAssistantUrl() string`

GetAiAssistantUrl returns the AiAssistantUrl field if non-nil, zero value otherwise.

### GetAiAssistantUrlOk

`func (o *O11yO11yGlobalConfig) GetAiAssistantUrlOk() (*string, bool)`

GetAiAssistantUrlOk returns a tuple with the AiAssistantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAssistantUrl

`func (o *O11yO11yGlobalConfig) SetAiAssistantUrl(v string)`

SetAiAssistantUrl sets AiAssistantUrl field to given value.

### HasAiAssistantUrl

`func (o *O11yO11yGlobalConfig) HasAiAssistantUrl() bool`

HasAiAssistantUrl returns a boolean if a field has been set.

### GetExternalUrl

`func (o *O11yO11yGlobalConfig) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *O11yO11yGlobalConfig) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *O11yO11yGlobalConfig) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *O11yO11yGlobalConfig) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetIdentN

`func (o *O11yO11yGlobalConfig) GetIdentN() O11yO11yIdentN`

GetIdentN returns the IdentN field if non-nil, zero value otherwise.

### GetIdentNOk

`func (o *O11yO11yGlobalConfig) GetIdentNOk() (*O11yO11yIdentN, bool)`

GetIdentNOk returns a tuple with the IdentN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentN

`func (o *O11yO11yGlobalConfig) SetIdentN(v O11yO11yIdentN)`

SetIdentN sets IdentN field to given value.

### HasIdentN

`func (o *O11yO11yGlobalConfig) HasIdentN() bool`

HasIdentN returns a boolean if a field has been set.

### GetIngestionUrl

`func (o *O11yO11yGlobalConfig) GetIngestionUrl() string`

GetIngestionUrl returns the IngestionUrl field if non-nil, zero value otherwise.

### GetIngestionUrlOk

`func (o *O11yO11yGlobalConfig) GetIngestionUrlOk() (*string, bool)`

GetIngestionUrlOk returns a tuple with the IngestionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionUrl

`func (o *O11yO11yGlobalConfig) SetIngestionUrl(v string)`

SetIngestionUrl sets IngestionUrl field to given value.

### HasIngestionUrl

`func (o *O11yO11yGlobalConfig) HasIngestionUrl() bool`

HasIngestionUrl returns a boolean if a field has been set.

### GetMcpUrl

`func (o *O11yO11yGlobalConfig) GetMcpUrl() string`

GetMcpUrl returns the McpUrl field if non-nil, zero value otherwise.

### GetMcpUrlOk

`func (o *O11yO11yGlobalConfig) GetMcpUrlOk() (*string, bool)`

GetMcpUrlOk returns a tuple with the McpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpUrl

`func (o *O11yO11yGlobalConfig) SetMcpUrl(v string)`

SetMcpUrl sets McpUrl field to given value.

### HasMcpUrl

`func (o *O11yO11yGlobalConfig) HasMcpUrl() bool`

HasMcpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


