# \FlowFlowRunsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowBulkCancelFlowRuns**](FlowFlowRunsAPI.md#FlowBulkCancelFlowRuns) | **Post** /v1/flow/flow-runs/bulk/cancel | Bulk cancel running flow runs
[**FlowGetFlowRun**](FlowFlowRunsAPI.md#FlowGetFlowRun) | **Get** /v1/flow/flow-runs/{id} | Get a flow run by id
[**FlowListFlowRuns**](FlowFlowRunsAPI.md#FlowListFlowRuns) | **Get** /v1/flow/flow-runs | List flow runs
[**FlowResumeFlowRun**](FlowFlowRunsAPI.md#FlowResumeFlowRun) | **Post** /v1/flow/flow-runs/{id}/requests/{requestId} | Resume a paused flow run with human input
[**FlowRetryFlowRun**](FlowFlowRunsAPI.md#FlowRetryFlowRun) | **Post** /v1/flow/flow-runs/{id}/retry | Retry a failed flow run



## FlowBulkCancelFlowRuns

> map[string]interface{} FlowBulkCancelFlowRuns(ctx).FlowBulkCancelFlowRunsRequest(flowBulkCancelFlowRunsRequest).Execute()

Bulk cancel running flow runs

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
	flowBulkCancelFlowRunsRequest := *openapiclient.NewFlowBulkCancelFlowRunsRequest() // FlowBulkCancelFlowRunsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowRunsAPI.FlowBulkCancelFlowRuns(context.Background()).FlowBulkCancelFlowRunsRequest(flowBulkCancelFlowRunsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowRunsAPI.FlowBulkCancelFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowBulkCancelFlowRuns`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowRunsAPI.FlowBulkCancelFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowBulkCancelFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowBulkCancelFlowRunsRequest** | [**FlowBulkCancelFlowRunsRequest**](FlowBulkCancelFlowRunsRequest.md) |  | 

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


## FlowGetFlowRun

> FlowFlowRun FlowGetFlowRun(ctx, id).Execute()

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
	resp, r, err := apiClient.FlowFlowRunsAPI.FlowGetFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowRunsAPI.FlowGetFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetFlowRun`: FlowFlowRun
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowRunsAPI.FlowGetFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetFlowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowFlowRun**](FlowFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListFlowRuns

> FlowListFlowRuns200Response FlowListFlowRuns(ctx).FlowId(flowId).Status(status).Tags(tags).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

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
	tags := []string{"Inner_example"} // []string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)
	createdAfter := time.Now() // time.Time |  (optional)
	createdBefore := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowRunsAPI.FlowListFlowRuns(context.Background()).FlowId(flowId).Status(status).Tags(tags).Cursor(cursor).Limit(limit).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowRunsAPI.FlowListFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListFlowRuns`: FlowListFlowRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowRunsAPI.FlowListFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** |  | 
 **status** | **[]string** |  | 
 **tags** | **[]string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]
 **createdAfter** | **time.Time** |  | 
 **createdBefore** | **time.Time** |  | 

### Return type

[**FlowListFlowRuns200Response**](FlowListFlowRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowResumeFlowRun

> map[string]interface{} FlowResumeFlowRun(ctx, id, requestId).Body(body).Execute()

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
	resp, r, err := apiClient.FlowFlowRunsAPI.FlowResumeFlowRun(context.Background(), id, requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowRunsAPI.FlowResumeFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowResumeFlowRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowRunsAPI.FlowResumeFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowResumeFlowRunRequest struct via the builder pattern


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


## FlowRetryFlowRun

> FlowFlowRun FlowRetryFlowRun(ctx, id).FlowRetryFlowRunRequest(flowRetryFlowRunRequest).Execute()

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
	flowRetryFlowRunRequest := *openapiclient.NewFlowRetryFlowRunRequest() // FlowRetryFlowRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFlowRunsAPI.FlowRetryFlowRun(context.Background(), id).FlowRetryFlowRunRequest(flowRetryFlowRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFlowRunsAPI.FlowRetryFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowRetryFlowRun`: FlowFlowRun
	fmt.Fprintf(os.Stdout, "Response from `FlowFlowRunsAPI.FlowRetryFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRetryFlowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowRetryFlowRunRequest** | [**FlowRetryFlowRunRequest**](FlowRetryFlowRunRequest.md) |  | 

### Return type

[**FlowFlowRun**](FlowFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

