# McpServerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]MCPServer**](MCPServer.md) | Servers is every external MCP server this org has registered. No secret VALUE is ever included — only whether one is set. | [optional] 

## Methods

### NewMcpServerList

`func NewMcpServerList() *McpServerList`

NewMcpServerList instantiates a new McpServerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerListWithDefaults

`func NewMcpServerListWithDefaults() *McpServerList`

NewMcpServerListWithDefaults instantiates a new McpServerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *McpServerList) GetServers() []MCPServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *McpServerList) GetServersOk() (*[]MCPServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *McpServerList) SetServers(v []MCPServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *McpServerList) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


