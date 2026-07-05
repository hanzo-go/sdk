# \AutoFlowRunsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoGetFlowRun**](AutoFlowRunsAPI.md#AutoGetFlowRun) | **Get** /v1/auto/flow-runs/{id} | Get a flow run by id
[**AutoListFlowRuns**](AutoFlowRunsAPI.md#AutoListFlowRuns) | **Get** /v1/auto/flow-runs | List flow runs
[**AutoResumeFlowRun**](AutoFlowRunsAPI.md#AutoResumeFlowRun) | **Post** /v1/auto/flow-runs/{id}/requests/{requestId} | Resume a paused flow run with human input
[**AutoRetryFlowRun**](AutoFlowRunsAPI.md#AutoRetryFlowRun) | **Post** /v1/auto/flow-runs/{id}/retry | Retry a failed flow run



## AutoGetFlowRun

> AutoFlowRun AutoGetFlowRun(ctx, id).Execute()

Get a flow run by id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoFlowRunsAPI.AutoGetFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowRunsAPI.AutoGetFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetFlowRun`: AutoFlowRun
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowRunsAPI.AutoGetFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetFlowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoFlowRun**](AutoFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoListFlowRuns

> AutoListFlowRuns200Response AutoListFlowRuns(ctx).FlowId(flowId).Status(status).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

List flow runs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	flowId := "flowId_example" // string |  (optional)
	status := []string{"Status_example"} // []string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)
	createdAfter := time.Now() // time.Time |  (optional)
	createdBefore := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoFlowRunsAPI.AutoListFlowRuns(context.Background()).FlowId(flowId).Status(status).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowRunsAPI.AutoListFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListFlowRuns`: AutoListFlowRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowRunsAPI.AutoListFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** |  | 
 **status** | **[]string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]
 **createdAfter** | **time.Time** |  | 
 **createdBefore** | **time.Time** |  | 

### Return type

[**AutoListFlowRuns200Response**](AutoListFlowRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoResumeFlowRun

> map[string]interface{} AutoResumeFlowRun(ctx, id, requestId).Body(body).Execute()

Resume a paused flow run with human input

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
	id := "id_example" // string | 
	requestId := "requestId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoFlowRunsAPI.AutoResumeFlowRun(context.Background(), id, requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowRunsAPI.AutoResumeFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoResumeFlowRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowRunsAPI.AutoResumeFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoResumeFlowRunRequest struct via the builder pattern


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


## AutoRetryFlowRun

> map[string]interface{} AutoRetryFlowRun(ctx, id).Execute()

Retry a failed flow run

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoFlowRunsAPI.AutoRetryFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFlowRunsAPI.AutoRetryFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoRetryFlowRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoFlowRunsAPI.AutoRetryFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoRetryFlowRunRequest struct via the builder pattern


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

