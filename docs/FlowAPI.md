# \FlowAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFlowWorkflowsByWorkflow**](FlowAPI.md#DeleteFlowWorkflowsByWorkflow) | **Delete** /v1/flow/workflows/{workflow} | Deletes one of the caller&#39;s workflows and its runs.
[**GetFlowRuns**](FlowAPI.md#GetFlowRuns) | **Get** /v1/flow/runs | Runs reads one workflow&#39;s recorded runs: every component build with its result, keyed by component.
[**GetFlowStatus**](FlowAPI.md#GetFlowStatus) | **Get** /v1/flow/status | Status reports whether the flow service is reachable and which version it runs.
[**GetFlowWorkflows**](FlowAPI.md#GetFlowWorkflows) | **Get** /v1/flow/workflows | Workflows lists the caller&#39;s workflows, paged.
[**GetFlowWorkflowsByWorkflow**](FlowAPI.md#GetFlowWorkflowsByWorkflow) | **Get** /v1/flow/workflows/{workflow} | Workflow reads one of the caller&#39;s workflows — the full record, graph included.
[**PatchFlowWorkflowsByWorkflow**](FlowAPI.md#PatchFlowWorkflowsByWorkflow) | **Patch** /v1/flow/workflows/{workflow} | Patches one of the caller&#39;s workflows: name, description, graph, or the locked flag — only the stated fields move.
[**PostFlowRuns**](FlowAPI.md#PostFlowRuns) | **Post** /v1/flow/runs | Run executes one of the caller&#39;s workflows synchronously: the graph runs in the flow service and the response carries the run&#39;s session and outputs.
[**PostFlowWorkflows**](FlowAPI.md#PostFlowWorkflows) | **Post** /v1/flow/workflows | Creates a workflow in the caller&#39;s org.



## DeleteFlowWorkflowsByWorkflow

> interface{} DeleteFlowWorkflowsByWorkflow(ctx, workflow).Execute()

Deletes one of the caller's workflows and its runs.



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
	workflow := "workflow_example" // string | Workflow is the workflow's UUID, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.DeleteFlowWorkflowsByWorkflow(context.Background(), workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.DeleteFlowWorkflowsByWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFlowWorkflowsByWorkflow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.DeleteFlowWorkflowsByWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflow** | **string** | Workflow is the workflow&#39;s UUID, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowWorkflowsByWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowRuns

> interface{} GetFlowRuns(ctx).Workflow(workflow).Execute()

Runs reads one workflow's recorded runs: every component build with its result, keyed by component.



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
	workflow := "workflow_example" // string | Workflow is the UUID of the workflow whose run records to read. It rides the query string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.GetFlowRuns(context.Background()).Workflow(workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.GetFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.GetFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflow** | **string** | Workflow is the UUID of the workflow whose run records to read. It rides the query string. | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowStatus

> FlowStatus GetFlowStatus(ctx).Execute()

Status reports whether the flow service is reachable and which version it runs.



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
	resp, r, err := apiClient.FlowAPI.GetFlowStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.GetFlowStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowStatus`: FlowStatus
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.GetFlowStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowStatusRequest struct via the builder pattern


### Return type

[**FlowStatus**](FlowStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowWorkflows

> interface{} GetFlowWorkflows(ctx).Page(page).Size(size).Execute()

Workflows lists the caller's workflows, paged.



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
	page := "page_example" // string | Page is the 1-based page of workflows to return. (optional)
	size := "size_example" // string | Size is how many workflows one page holds (the product caps it at 100). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.GetFlowWorkflows(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.GetFlowWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowWorkflows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.GetFlowWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of workflows to return. | 
 **size** | **string** | Size is how many workflows one page holds (the product caps it at 100). | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowWorkflowsByWorkflow

> interface{} GetFlowWorkflowsByWorkflow(ctx, workflow).Execute()

Workflow reads one of the caller's workflows — the full record, graph included.



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
	workflow := "workflow_example" // string | Workflow is the workflow's UUID, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.GetFlowWorkflowsByWorkflow(context.Background(), workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.GetFlowWorkflowsByWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowWorkflowsByWorkflow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.GetFlowWorkflowsByWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflow** | **string** | Workflow is the workflow&#39;s UUID, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowWorkflowsByWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlowWorkflowsByWorkflow

> interface{} PatchFlowWorkflowsByWorkflow(ctx, workflow).FlowUpdate(flowUpdate).Execute()

Patches one of the caller's workflows: name, description, graph, or the locked flag — only the stated fields move.



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
	workflow := "workflow_example" // string | Workflow is the workflow's UUID, taken from the path.
	flowUpdate := *openapiclient.NewFlowUpdate() // FlowUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.PatchFlowWorkflowsByWorkflow(context.Background(), workflow).FlowUpdate(flowUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.PatchFlowWorkflowsByWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlowWorkflowsByWorkflow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.PatchFlowWorkflowsByWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflow** | **string** | Workflow is the workflow&#39;s UUID, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlowWorkflowsByWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowUpdate** | [**FlowUpdate**](FlowUpdate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlowRuns

> interface{} PostFlowRuns(ctx).FlowRun(flowRun).Execute()

Run executes one of the caller's workflows synchronously: the graph runs in the flow service and the response carries the run's session and outputs.



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
	flowRun := *openapiclient.NewFlowRun() // FlowRun | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.PostFlowRuns(context.Background()).FlowRun(flowRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.PostFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlowRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.PostFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowRun** | [**FlowRun**](FlowRun.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlowWorkflows

> interface{} PostFlowWorkflows(ctx).FlowCreate(flowCreate).Execute()

Creates a workflow in the caller's org.



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
	flowCreate := *openapiclient.NewFlowCreate() // FlowCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.PostFlowWorkflows(context.Background()).FlowCreate(flowCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.PostFlowWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlowWorkflows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.PostFlowWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFlowWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowCreate** | [**FlowCreate**](FlowCreate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

