# \McpAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMcpServersById**](McpAPI.md#DeleteMcpServersById) | **Delete** /v1/mcp/servers/{id} | Deregisters one of the caller org&#39;s external MCP servers, so its tools leave the registry.
[**GetMcpServers**](McpAPI.md#GetMcpServers) | **Get** /v1/mcp/servers | Lists the external MCP servers the caller&#39;s org has registered.
[**PostMcpServers**](McpAPI.md#PostMcpServers) | **Post** /v1/mcp/servers | Gives the caller&#39;s org one more external MCP server, so its tools join the org&#39;s tool plane and the fleet&#39;s MCP door.



## DeleteMcpServersById

> DeleteMcpServersById(ctx, id).Execute()

Deregisters one of the caller org's external MCP servers, so its tools leave the registry.



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
	r, err := apiClient.McpAPI.DeleteMcpServersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpAPI.DeleteMcpServersById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMcpServersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMcpServers

> McpServerList GetMcpServers(ctx).Execute()

Lists the external MCP servers the caller's org has registered.



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
	resp, r, err := apiClient.McpAPI.GetMcpServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpAPI.GetMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMcpServers`: McpServerList
	fmt.Fprintf(os.Stdout, "Response from `McpAPI.GetMcpServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMcpServersRequest struct via the builder pattern


### Return type

[**McpServerList**](McpServerList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMcpServers

> MCPServer PostMcpServers(ctx).CreateServerReq(createServerReq).Execute()

Gives the caller's org one more external MCP server, so its tools join the org's tool plane and the fleet's MCP door.



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
	createServerReq := *openapiclient.NewCreateServerReq() // CreateServerReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpAPI.PostMcpServers(context.Background()).CreateServerReq(createServerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpAPI.PostMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMcpServers`: MCPServer
	fmt.Fprintf(os.Stdout, "Response from `McpAPI.PostMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerReq** | [**CreateServerReq**](CreateServerReq.md) |  | 

### Return type

[**MCPServer**](MCPServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

