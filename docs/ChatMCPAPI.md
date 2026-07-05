# \ChatMCPAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteMcpServersByservername**](ChatMCPAPI.md#ChatDeleteMcpServersByservername) | **Delete** /v1/chat/mcp/servers/{serverName} | Delete an MCP server
[**ChatGetMcpByservernameAuthValues**](ChatMCPAPI.md#ChatGetMcpByservernameAuthValues) | **Get** /v1/chat/mcp/{serverName}/auth-values | Check which auth values exist for an MCP server
[**ChatGetMcpByservernameOauthCallback**](ChatMCPAPI.md#ChatGetMcpByservernameOauthCallback) | **Get** /v1/chat/mcp/{serverName}/oauth/callback | MCP OAuth callback
[**ChatGetMcpByservernameOauthInitiate**](ChatMCPAPI.md#ChatGetMcpByservernameOauthInitiate) | **Get** /v1/chat/mcp/{serverName}/oauth/initiate | Initiate MCP OAuth flow
[**ChatGetMcpConnectionStatus**](ChatMCPAPI.md#ChatGetMcpConnectionStatus) | **Get** /v1/chat/mcp/connection/status | Get connection status for all MCP servers
[**ChatGetMcpConnectionStatusByservername**](ChatMCPAPI.md#ChatGetMcpConnectionStatusByservername) | **Get** /v1/chat/mcp/connection/status/{serverName} | Get connection status for a specific MCP server
[**ChatGetMcpOauthStatusByflowid**](ChatMCPAPI.md#ChatGetMcpOauthStatusByflowid) | **Get** /v1/chat/mcp/oauth/status/{flowId} | Check OAuth flow status
[**ChatGetMcpOauthTokensByflowid**](ChatMCPAPI.md#ChatGetMcpOauthTokensByflowid) | **Get** /v1/chat/mcp/oauth/tokens/{flowId} | Get OAuth tokens for a completed flow
[**ChatGetMcpServers**](ChatMCPAPI.md#ChatGetMcpServers) | **Get** /v1/chat/mcp/servers | List user-managed MCP servers
[**ChatGetMcpServersByservername**](ChatMCPAPI.md#ChatGetMcpServersByservername) | **Get** /v1/chat/mcp/servers/{serverName} | Get an MCP server by name
[**ChatGetMcpTools**](ChatMCPAPI.md#ChatGetMcpTools) | **Get** /v1/chat/mcp/tools | Get all available MCP tools
[**ChatPatchMcpServersByservername**](ChatMCPAPI.md#ChatPatchMcpServersByservername) | **Patch** /v1/chat/mcp/servers/{serverName} | Update an MCP server
[**ChatPostMcpByservernameOauthBind**](ChatMCPAPI.md#ChatPostMcpByservernameOauthBind) | **Post** /v1/chat/mcp/{serverName}/oauth/bind | Set CSRF binding cookie for MCP OAuth
[**ChatPostMcpByservernameReinitialize**](ChatMCPAPI.md#ChatPostMcpByservernameReinitialize) | **Post** /v1/chat/mcp/{serverName}/reinitialize | Reinitialize an MCP server
[**ChatPostMcpOauthCancelByservername**](ChatMCPAPI.md#ChatPostMcpOauthCancelByservername) | **Post** /v1/chat/mcp/oauth/cancel/{serverName} | Cancel an OAuth flow
[**ChatPostMcpServers**](ChatMCPAPI.md#ChatPostMcpServers) | **Post** /v1/chat/mcp/servers | Create a user-managed MCP server



## ChatDeleteMcpServersByservername

> map[string]interface{} ChatDeleteMcpServersByservername(ctx, serverName).Execute()

Delete an MCP server

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatDeleteMcpServersByservername(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatDeleteMcpServersByservername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteMcpServersByservername`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatDeleteMcpServersByservername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteMcpServersByservernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpByservernameAuthValues

> ChatGetMcpByservernameAuthValues200Response ChatGetMcpByservernameAuthValues(ctx, serverName).Execute()

Check which auth values exist for an MCP server

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpByservernameAuthValues(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpByservernameAuthValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpByservernameAuthValues`: ChatGetMcpByservernameAuthValues200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpByservernameAuthValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpByservernameAuthValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatGetMcpByservernameAuthValues200Response**](ChatGetMcpByservernameAuthValues200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpByservernameOauthCallback

> ChatGetMcpByservernameOauthCallback(ctx, serverName).Code(code).State(state).Execute()

MCP OAuth callback

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
	serverName := "serverName_example" // string | 
	code := "code_example" // string |  (optional)
	state := "state_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChatMCPAPI.ChatGetMcpByservernameOauthCallback(context.Background(), serverName).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpByservernameOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpByservernameOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** |  | 
 **state** | **string** |  | 

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


## ChatGetMcpByservernameOauthInitiate

> ChatGetMcpByservernameOauthInitiate(ctx, serverName).UserId(userId).FlowId(flowId).Execute()

Initiate MCP OAuth flow

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
	serverName := "serverName_example" // string | 
	userId := "userId_example" // string | 
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChatMCPAPI.ChatGetMcpByservernameOauthInitiate(context.Background(), serverName).UserId(userId).FlowId(flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpByservernameOauthInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpByservernameOauthInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 
 **flowId** | **string** |  | 

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


## ChatGetMcpConnectionStatus

> ChatGetMcpConnectionStatus200Response ChatGetMcpConnectionStatus(ctx).Execute()

Get connection status for all MCP servers

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
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpConnectionStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpConnectionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpConnectionStatus`: ChatGetMcpConnectionStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpConnectionStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpConnectionStatusRequest struct via the builder pattern


### Return type

[**ChatGetMcpConnectionStatus200Response**](ChatGetMcpConnectionStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpConnectionStatusByservername

> map[string]interface{} ChatGetMcpConnectionStatusByservername(ctx, serverName).Execute()

Get connection status for a specific MCP server

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpConnectionStatusByservername(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpConnectionStatusByservername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpConnectionStatusByservername`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpConnectionStatusByservername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpConnectionStatusByservernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpOauthStatusByflowid

> ChatGetMcpOauthStatusByflowid200Response ChatGetMcpOauthStatusByflowid(ctx, flowId).Execute()

Check OAuth flow status

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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpOauthStatusByflowid(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpOauthStatusByflowid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpOauthStatusByflowid`: ChatGetMcpOauthStatusByflowid200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpOauthStatusByflowid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpOauthStatusByflowidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatGetMcpOauthStatusByflowid200Response**](ChatGetMcpOauthStatusByflowid200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpOauthTokensByflowid

> map[string]interface{} ChatGetMcpOauthTokensByflowid(ctx, flowId).Execute()

Get OAuth tokens for a completed flow

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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpOauthTokensByflowid(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpOauthTokensByflowid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpOauthTokensByflowid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpOauthTokensByflowid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpOauthTokensByflowidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpServers

> map[string]interface{} ChatGetMcpServers(ctx).Limit(limit).After(after).Search(search).Execute()

List user-managed MCP servers

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
	limit := int32(56) // int32 |  (optional)
	after := "after_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpServers(context.Background()).Limit(limit).After(after).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpServers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **after** | **string** |  | 
 **search** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpServersByservername

> map[string]interface{} ChatGetMcpServersByservername(ctx, serverName).Execute()

Get an MCP server by name

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpServersByservername(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpServersByservername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpServersByservername`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpServersByservername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpServersByservernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMcpTools

> map[string]interface{} ChatGetMcpTools(ctx).Execute()

Get all available MCP tools

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
	resp, r, err := apiClient.ChatMCPAPI.ChatGetMcpTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatGetMcpTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMcpTools`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatGetMcpTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMcpToolsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPatchMcpServersByservername

> map[string]interface{} ChatPatchMcpServersByservername(ctx, serverName).Body(body).Execute()

Update an MCP server

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
	serverName := "serverName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatPatchMcpServersByservername(context.Background(), serverName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatPatchMcpServersByservername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchMcpServersByservername`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatPatchMcpServersByservername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchMcpServersByservernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMcpByservernameOauthBind

> map[string]interface{} ChatPostMcpByservernameOauthBind(ctx, serverName).Execute()

Set CSRF binding cookie for MCP OAuth

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatPostMcpByservernameOauthBind(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatPostMcpByservernameOauthBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMcpByservernameOauthBind`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatPostMcpByservernameOauthBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMcpByservernameOauthBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMcpByservernameReinitialize

> ChatPostMcpByservernameReinitialize200Response ChatPostMcpByservernameReinitialize(ctx, serverName).Execute()

Reinitialize an MCP server

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatPostMcpByservernameReinitialize(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatPostMcpByservernameReinitialize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMcpByservernameReinitialize`: ChatPostMcpByservernameReinitialize200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatPostMcpByservernameReinitialize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMcpByservernameReinitializeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatPostMcpByservernameReinitialize200Response**](ChatPostMcpByservernameReinitialize200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMcpOauthCancelByservername

> map[string]interface{} ChatPostMcpOauthCancelByservername(ctx, serverName).Execute()

Cancel an OAuth flow

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
	serverName := "serverName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatPostMcpOauthCancelByservername(context.Background(), serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatPostMcpOauthCancelByservername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMcpOauthCancelByservername`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatPostMcpOauthCancelByservername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMcpOauthCancelByservernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMcpServers

> map[string]interface{} ChatPostMcpServers(ctx).Body(body).Execute()

Create a user-managed MCP server

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMCPAPI.ChatPostMcpServers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMCPAPI.ChatPostMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMcpServers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMCPAPI.ChatPostMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

