# \FlowAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1FlowWorkflowsWorkflow**](FlowAPI.md#CloudDeleteV1FlowWorkflowsWorkflow) | **Delete** /v1/flow/workflows/{workflow} | WorkflowDelete deletes one of the caller&#39;s workflows and its runs.
[**CloudGetV1FlowRuns**](FlowAPI.md#CloudGetV1FlowRuns) | **Get** /v1/flow/runs | Runs reads one workflow&#39;s recorded runs: every component build with its result, keyed by component.
[**CloudGetV1FlowStatus**](FlowAPI.md#CloudGetV1FlowStatus) | **Get** /v1/flow/status | Status reports whether the flow service is reachable and which version it runs.
[**CloudGetV1FlowWorkflows**](FlowAPI.md#CloudGetV1FlowWorkflows) | **Get** /v1/flow/workflows | Workflows lists the caller&#39;s workflows, paged.
[**CloudGetV1FlowWorkflowsWorkflow**](FlowAPI.md#CloudGetV1FlowWorkflowsWorkflow) | **Get** /v1/flow/workflows/{workflow} | Workflow reads one of the caller&#39;s workflows — the full record, graph included.
[**CloudPatchV1FlowWorkflowsWorkflow**](FlowAPI.md#CloudPatchV1FlowWorkflowsWorkflow) | **Patch** /v1/flow/workflows/{workflow} | WorkflowUpdate patches one of the caller&#39;s workflows: name, description, graph, or the locked flag — only the stated fields move.
[**CloudPostV1FlowRuns**](FlowAPI.md#CloudPostV1FlowRuns) | **Post** /v1/flow/runs | Run executes one of the caller&#39;s workflows synchronously: the graph runs in the flow service and the response carries the run&#39;s session and outputs.
[**CloudPostV1FlowWorkflows**](FlowAPI.md#CloudPostV1FlowWorkflows) | **Post** /v1/flow/workflows | WorkflowCreate creates a workflow in the caller&#39;s org.



## CloudDeleteV1FlowWorkflowsWorkflow

> interface{} CloudDeleteV1FlowWorkflowsWorkflow(ctx, workflow).Execute()

WorkflowDelete deletes one of the caller's workflows and its runs.



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
	resp, r, err := apiClient.FlowAPI.CloudDeleteV1FlowWorkflowsWorkflow(context.Background(), workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudDeleteV1FlowWorkflowsWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1FlowWorkflowsWorkflow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudDeleteV1FlowWorkflowsWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflow** | **string** | Workflow is the workflow&#39;s UUID, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1FlowWorkflowsWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlowRuns

> interface{} CloudGetV1FlowRuns(ctx).Workflow(workflow).Execute()

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
	resp, r, err := apiClient.FlowAPI.CloudGetV1FlowRuns(context.Background()).Workflow(workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudGetV1FlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlowRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudGetV1FlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflow** | **string** | Workflow is the UUID of the workflow whose run records to read. It rides the query string. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlowStatus

> CloudFlowStatus CloudGetV1FlowStatus(ctx).Execute()

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
	resp, r, err := apiClient.FlowAPI.CloudGetV1FlowStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudGetV1FlowStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlowStatus`: CloudFlowStatus
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudGetV1FlowStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlowStatusRequest struct via the builder pattern


### Return type

[**CloudFlowStatus**](CloudFlowStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlowWorkflows

> interface{} CloudGetV1FlowWorkflows(ctx).Page(page).Size(size).Execute()

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
	resp, r, err := apiClient.FlowAPI.CloudGetV1FlowWorkflows(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudGetV1FlowWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlowWorkflows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudGetV1FlowWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlowWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of workflows to return. | 
 **size** | **string** | Size is how many workflows one page holds (the product caps it at 100). | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlowWorkflowsWorkflow

> interface{} CloudGetV1FlowWorkflowsWorkflow(ctx, workflow).Execute()

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
	resp, r, err := apiClient.FlowAPI.CloudGetV1FlowWorkflowsWorkflow(context.Background(), workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudGetV1FlowWorkflowsWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlowWorkflowsWorkflow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudGetV1FlowWorkflowsWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflow** | **string** | Workflow is the workflow&#39;s UUID, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlowWorkflowsWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1FlowWorkflowsWorkflow

> interface{} CloudPatchV1FlowWorkflowsWorkflow(ctx, workflow).CloudFlowUpdate(cloudFlowUpdate).Execute()

WorkflowUpdate patches one of the caller's workflows: name, description, graph, or the locked flag — only the stated fields move.



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
	cloudFlowUpdate := *openapiclient.NewCloudFlowUpdate() // CloudFlowUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.CloudPatchV1FlowWorkflowsWorkflow(context.Background(), workflow).CloudFlowUpdate(cloudFlowUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudPatchV1FlowWorkflowsWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1FlowWorkflowsWorkflow`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudPatchV1FlowWorkflowsWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflow** | **string** | Workflow is the workflow&#39;s UUID, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1FlowWorkflowsWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudFlowUpdate** | [**CloudFlowUpdate**](CloudFlowUpdate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FlowRuns

> interface{} CloudPostV1FlowRuns(ctx).CloudFlowRun(cloudFlowRun).Execute()

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
	cloudFlowRun := *openapiclient.NewCloudFlowRun() // CloudFlowRun | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.CloudPostV1FlowRuns(context.Background()).CloudFlowRun(cloudFlowRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudPostV1FlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FlowRuns`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudPostV1FlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFlowRun** | [**CloudFlowRun**](CloudFlowRun.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FlowWorkflows

> interface{} CloudPostV1FlowWorkflows(ctx).CloudFlowCreate(cloudFlowCreate).Execute()

WorkflowCreate creates a workflow in the caller's org.



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
	cloudFlowCreate := *openapiclient.NewCloudFlowCreate() // CloudFlowCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAPI.CloudPostV1FlowWorkflows(context.Background()).CloudFlowCreate(cloudFlowCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAPI.CloudPostV1FlowWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FlowWorkflows`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAPI.CloudPostV1FlowWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FlowWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFlowCreate** | [**CloudFlowCreate**](CloudFlowCreate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

