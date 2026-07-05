# \TasksTasksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksTasksEvents**](TasksTasksAPI.md#TasksTasksEvents) | **Get** /v1/tasks/events | Realtime event stream (SSE, identity-gated)
[**TasksTasksGet**](TasksTasksAPI.md#TasksTasksGet) | **Get** /v1/tasks/{resource} | Engine JSON API (namespaces, workflows, activities, …), identity-gated
[**TasksTasksMcp**](TasksTasksAPI.md#TasksTasksMcp) | **Post** /v1/tasks/mcp | Tasks MCP tool surface (JSON-RPC, identity-gated)
[**TasksTasksPost**](TasksTasksAPI.md#TasksTasksPost) | **Post** /v1/tasks/{resource} | Engine JSON API write (identity-gated)



## TasksTasksEvents

> string TasksTasksEvents(ctx).Execute()

Realtime event stream (SSE, identity-gated)

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
	resp, r, err := apiClient.TasksTasksAPI.TasksTasksEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksTasksAPI.TasksTasksEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksEvents`: string
	fmt.Fprintf(os.Stdout, "Response from `TasksTasksAPI.TasksTasksEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksEventsRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksGet

> map[string]interface{} TasksTasksGet(ctx, resource).Execute()

Engine JSON API (namespaces, workflows, activities, …), identity-gated

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
	resource := "resource_example" // string | 'Engine resource path (e.g. namespaces, nexus)'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksTasksAPI.TasksTasksGet(context.Background(), resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksTasksAPI.TasksTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksTasksAPI.TasksTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | &#39;Engine resource path (e.g. namespaces, nexus)&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksGetRequest struct via the builder pattern


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


## TasksTasksMcp

> map[string]interface{} TasksTasksMcp(ctx).RequestBody(requestBody).Execute()

Tasks MCP tool surface (JSON-RPC, identity-gated)

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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksTasksAPI.TasksTasksMcp(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksTasksAPI.TasksTasksMcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksMcp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksTasksAPI.TasksTasksMcp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksMcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

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


## TasksTasksPost

> map[string]interface{} TasksTasksPost(ctx, resource).RequestBody(requestBody).Execute()

Engine JSON API write (identity-gated)

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
	resource := "resource_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksTasksAPI.TasksTasksPost(context.Background(), resource).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksTasksAPI.TasksTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksTasksAPI.TasksTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

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

