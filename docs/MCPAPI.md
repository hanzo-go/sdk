# \MCPAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationsMcp**](MCPAPI.md#AutomationsMcp) | **Post** /v1/automations/mcp | JSON-RPC 2.0 tool surface over connector actions
[**CloudDeleteV1McpServersId**](MCPAPI.md#CloudDeleteV1McpServersId) | **Delete** /v1/mcp/servers/{id} | DeleteServer deregisters one of the caller org&#39;s external MCP servers, so its tools leave the registry.
[**CloudGetV1McpServers**](MCPAPI.md#CloudGetV1McpServers) | **Get** /v1/mcp/servers | ListServers lists the external MCP servers the caller&#39;s org has registered.
[**CloudPostV1McpServers**](MCPAPI.md#CloudPostV1McpServers) | **Post** /v1/mcp/servers | CreateServer gives the caller&#39;s org one more external MCP server, so its tools join the org&#39;s tool plane and the fleet&#39;s MCP door.



## AutomationsMcp

> AutomationsMcpResponse AutomationsMcp(ctx).AutomationsMcpRequest(automationsMcpRequest).Execute()

JSON-RPC 2.0 tool surface over connector actions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	automationsMcpRequest := *openapiclient.NewAutomationsMcpRequest("2.0", "Method_example") // AutomationsMcpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPAPI.AutomationsMcp(context.Background()).AutomationsMcpRequest(automationsMcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.AutomationsMcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsMcp`: AutomationsMcpResponse
	fmt.Fprintf(os.Stdout, "Response from `MCPAPI.AutomationsMcp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsMcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationsMcpRequest** | [**AutomationsMcpRequest**](AutomationsMcpRequest.md) |  | 

### Return type

[**AutomationsMcpResponse**](AutomationsMcpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1McpServersId

> CloudDeleteV1McpServersId(ctx, id).Execute()

DeleteServer deregisters one of the caller org's external MCP servers, so its tools leave the registry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the server to deregister, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MCPAPI.CloudDeleteV1McpServersId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.CloudDeleteV1McpServersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the server to deregister, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1McpServersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1McpServers

> CloudMcpServerList CloudGetV1McpServers(ctx).Execute()

ListServers lists the external MCP servers the caller's org has registered.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPAPI.CloudGetV1McpServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.CloudGetV1McpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1McpServers`: CloudMcpServerList
	fmt.Fprintf(os.Stdout, "Response from `MCPAPI.CloudGetV1McpServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1McpServersRequest struct via the builder pattern


### Return type

[**CloudMcpServerList**](CloudMcpServerList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1McpServers

> CloudMCPServer CloudPostV1McpServers(ctx).CloudCreateServerReq(cloudCreateServerReq).Execute()

CreateServer gives the caller's org one more external MCP server, so its tools join the org's tool plane and the fleet's MCP door.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudCreateServerReq := *openapiclient.NewCloudCreateServerReq() // CloudCreateServerReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPAPI.CloudPostV1McpServers(context.Background()).CloudCreateServerReq(cloudCreateServerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.CloudPostV1McpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1McpServers`: CloudMCPServer
	fmt.Fprintf(os.Stdout, "Response from `MCPAPI.CloudPostV1McpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1McpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateServerReq** | [**CloudCreateServerReq**](CloudCreateServerReq.md) |  | 

### Return type

[**CloudMCPServer**](CloudMCPServer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

