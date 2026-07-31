# CloudMcpServerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]CloudMCPServer**](CloudMCPServer.md) | Servers is every external MCP server this org has registered. No secret VALUE is ever included — only whether one is set. | [optional] 

## Methods

### NewCloudMcpServerList

`func NewCloudMcpServerList() *CloudMcpServerList`

NewCloudMcpServerList instantiates a new CloudMcpServerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMcpServerListWithDefaults

`func NewCloudMcpServerListWithDefaults() *CloudMcpServerList`

NewCloudMcpServerListWithDefaults instantiates a new CloudMcpServerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *CloudMcpServerList) GetServers() []CloudMCPServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *CloudMcpServerList) GetServersOk() (*[]CloudMCPServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *CloudMcpServerList) SetServers(v []CloudMCPServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *CloudMcpServerList) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


