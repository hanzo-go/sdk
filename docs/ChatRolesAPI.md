# \ChatRolesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatGetRolesByrolename**](ChatRolesAPI.md#ChatGetRolesByrolename) | **Get** /v1/chat/roles/{roleName} | Get a role by name
[**ChatPutRolesByrolenameAgents**](ChatRolesAPI.md#ChatPutRolesByrolenameAgents) | **Put** /v1/chat/roles/{roleName}/agents | Update agent permissions for a role
[**ChatPutRolesByrolenameMarketplace**](ChatRolesAPI.md#ChatPutRolesByrolenameMarketplace) | **Put** /v1/chat/roles/{roleName}/marketplace | Update marketplace permissions for a role
[**ChatPutRolesByrolenameMcpServers**](ChatRolesAPI.md#ChatPutRolesByrolenameMcpServers) | **Put** /v1/chat/roles/{roleName}/mcp-servers | Update MCP servers permissions for a role
[**ChatPutRolesByrolenameMemories**](ChatRolesAPI.md#ChatPutRolesByrolenameMemories) | **Put** /v1/chat/roles/{roleName}/memories | Update memory permissions for a role
[**ChatPutRolesByrolenamePeoplePicker**](ChatRolesAPI.md#ChatPutRolesByrolenamePeoplePicker) | **Put** /v1/chat/roles/{roleName}/people-picker | Update people picker permissions for a role
[**ChatPutRolesByrolenamePrompts**](ChatRolesAPI.md#ChatPutRolesByrolenamePrompts) | **Put** /v1/chat/roles/{roleName}/prompts | Update prompt permissions for a role
[**ChatPutRolesByrolenameRemoteAgents**](ChatRolesAPI.md#ChatPutRolesByrolenameRemoteAgents) | **Put** /v1/chat/roles/{roleName}/remote-agents | Update remote agents permissions for a role



## ChatGetRolesByrolename

> map[string]interface{} ChatGetRolesByrolename(ctx, roleName).Execute()

Get a role by name

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
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatGetRolesByrolename(context.Background(), roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatGetRolesByrolename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetRolesByrolename`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatGetRolesByrolename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetRolesByrolenameRequest struct via the builder pattern


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


## ChatPutRolesByrolenameAgents

> map[string]interface{} ChatPutRolesByrolenameAgents(ctx, roleName).Body(body).Execute()

Update agent permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenameAgents(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenameAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenameAgents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenameAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenameAgentsRequest struct via the builder pattern


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


## ChatPutRolesByrolenameMarketplace

> map[string]interface{} ChatPutRolesByrolenameMarketplace(ctx, roleName).Body(body).Execute()

Update marketplace permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenameMarketplace(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenameMarketplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenameMarketplace`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenameMarketplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenameMarketplaceRequest struct via the builder pattern


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


## ChatPutRolesByrolenameMcpServers

> map[string]interface{} ChatPutRolesByrolenameMcpServers(ctx, roleName).Body(body).Execute()

Update MCP servers permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenameMcpServers(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenameMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenameMcpServers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenameMcpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenameMcpServersRequest struct via the builder pattern


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


## ChatPutRolesByrolenameMemories

> map[string]interface{} ChatPutRolesByrolenameMemories(ctx, roleName).Body(body).Execute()

Update memory permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenameMemories(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenameMemories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenameMemories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenameMemories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenameMemoriesRequest struct via the builder pattern


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


## ChatPutRolesByrolenamePeoplePicker

> map[string]interface{} ChatPutRolesByrolenamePeoplePicker(ctx, roleName).Body(body).Execute()

Update people picker permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenamePeoplePicker(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenamePeoplePicker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenamePeoplePicker`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenamePeoplePicker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenamePeoplePickerRequest struct via the builder pattern


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


## ChatPutRolesByrolenamePrompts

> map[string]interface{} ChatPutRolesByrolenamePrompts(ctx, roleName).Body(body).Execute()

Update prompt permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenamePrompts(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenamePrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenamePrompts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenamePrompts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenamePromptsRequest struct via the builder pattern


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


## ChatPutRolesByrolenameRemoteAgents

> map[string]interface{} ChatPutRolesByrolenameRemoteAgents(ctx, roleName).Body(body).Execute()

Update remote agents permissions for a role

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
	roleName := "roleName_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatRolesAPI.ChatPutRolesByrolenameRemoteAgents(context.Background(), roleName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatRolesAPI.ChatPutRolesByrolenameRemoteAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutRolesByrolenameRemoteAgents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatRolesAPI.ChatPutRolesByrolenameRemoteAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutRolesByrolenameRemoteAgentsRequest struct via the builder pattern


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

