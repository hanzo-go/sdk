# AiMCPSurface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]AiMCPApp**](AiMCPApp.md) | Apps is one row per subsystem this deployment composes, in manifest order. | [optional] 
**Names** | Pointer to **[]string** | Names are this process&#39;s own tool names, present only when the query asked for them. | [optional] 
**Tools** | Pointer to **int32** | Tools is how many tools THIS PROCESS&#39;s MCP server carries: its own typed-op registry, projected. It is the only number a subsystem can state honestly — what the FLEET&#39;s server carries is a question only the host can ask, and it asks it by asking every subsystem (POST /v1/mcp, tools/list). | [optional] 

## Methods

### NewAiMCPSurface

`func NewAiMCPSurface() *AiMCPSurface`

NewAiMCPSurface instantiates a new AiMCPSurface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMCPSurfaceWithDefaults

`func NewAiMCPSurfaceWithDefaults() *AiMCPSurface`

NewAiMCPSurfaceWithDefaults instantiates a new AiMCPSurface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *AiMCPSurface) GetApps() []AiMCPApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AiMCPSurface) GetAppsOk() (*[]AiMCPApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AiMCPSurface) SetApps(v []AiMCPApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AiMCPSurface) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetNames

`func (o *AiMCPSurface) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *AiMCPSurface) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *AiMCPSurface) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *AiMCPSurface) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetTools

`func (o *AiMCPSurface) GetTools() int32`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiMCPSurface) GetToolsOk() (*int32, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiMCPSurface) SetTools(v int32)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiMCPSurface) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


