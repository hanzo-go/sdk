# \NexusWorkflowAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddWorkflow**](NexusWorkflowAPIAPI.md#NexusAddWorkflow) | **Post** /v1/nexus/add-workflow | add Workflow
[**NexusDeleteWorkflow**](NexusWorkflowAPIAPI.md#NexusDeleteWorkflow) | **Post** /v1/nexus/delete-workflow | delete Workflow
[**NexusGetGlobalWorkflows**](NexusWorkflowAPIAPI.md#NexusGetGlobalWorkflows) | **Get** /v1/nexus/get-global-workflows | get Global Workflows
[**NexusGetWorkflow**](NexusWorkflowAPIAPI.md#NexusGetWorkflow) | **Get** /v1/nexus/get-workflow | get Workflow
[**NexusGetWorkflows**](NexusWorkflowAPIAPI.md#NexusGetWorkflows) | **Get** /v1/nexus/get-workflows | get Workflows
[**NexusUpdateWorkflow**](NexusWorkflowAPIAPI.md#NexusUpdateWorkflow) | **Post** /v1/nexus/update-workflow | update Workflow



## NexusAddWorkflow

> NexusResponse NexusAddWorkflow(ctx).NexusWorkflow(nexusWorkflow).Execute()

add Workflow



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
	nexusWorkflow := *openapiclient.NewNexusWorkflow() // NexusWorkflow | The details of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusWorkflowAPIAPI.NexusAddWorkflow(context.Background()).NexusWorkflow(nexusWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusWorkflowAPIAPI.NexusAddWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddWorkflow`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusWorkflowAPIAPI.NexusAddWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusWorkflow** | [**NexusWorkflow**](NexusWorkflow.md) | The details of the workflow | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteWorkflow

> NexusResponse NexusDeleteWorkflow(ctx).NexusWorkflow(nexusWorkflow).Execute()

delete Workflow



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
	nexusWorkflow := *openapiclient.NewNexusWorkflow() // NexusWorkflow | The details of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusWorkflowAPIAPI.NexusDeleteWorkflow(context.Background()).NexusWorkflow(nexusWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusWorkflowAPIAPI.NexusDeleteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteWorkflow`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusWorkflowAPIAPI.NexusDeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusWorkflow** | [**NexusWorkflow**](NexusWorkflow.md) | The details of the workflow | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalWorkflows

> []NexusWorkflow NexusGetGlobalWorkflows(ctx).Execute()

get Global Workflows



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
	resp, r, err := apiClient.NexusWorkflowAPIAPI.NexusGetGlobalWorkflows(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusWorkflowAPIAPI.NexusGetGlobalWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalWorkflows`: []NexusWorkflow
	fmt.Fprintf(os.Stdout, "Response from `NexusWorkflowAPIAPI.NexusGetGlobalWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalWorkflowsRequest struct via the builder pattern


### Return type

[**[]NexusWorkflow**](NexusWorkflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetWorkflow

> NexusWorkflow NexusGetWorkflow(ctx).Id(id).Execute()

get Workflow



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
	id := "id_example" // string | The id (owner/name) of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusWorkflowAPIAPI.NexusGetWorkflow(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusWorkflowAPIAPI.NexusGetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetWorkflow`: NexusWorkflow
	fmt.Fprintf(os.Stdout, "Response from `NexusWorkflowAPIAPI.NexusGetWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the workflow | 

### Return type

[**NexusWorkflow**](NexusWorkflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetWorkflows

> []NexusWorkflow NexusGetWorkflows(ctx).Owner(owner).Execute()

get Workflows



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
	owner := "owner_example" // string | The owner of the workflows

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusWorkflowAPIAPI.NexusGetWorkflows(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusWorkflowAPIAPI.NexusGetWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetWorkflows`: []NexusWorkflow
	fmt.Fprintf(os.Stdout, "Response from `NexusWorkflowAPIAPI.NexusGetWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the workflows | 

### Return type

[**[]NexusWorkflow**](NexusWorkflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateWorkflow

> NexusResponse NexusUpdateWorkflow(ctx).Id(id).NexusWorkflow(nexusWorkflow).Execute()

update Workflow



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
	id := "id_example" // string | The id (owner/name) of the workflow
	nexusWorkflow := *openapiclient.NewNexusWorkflow() // NexusWorkflow | The details of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusWorkflowAPIAPI.NexusUpdateWorkflow(context.Background()).Id(id).NexusWorkflow(nexusWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusWorkflowAPIAPI.NexusUpdateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateWorkflow`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusWorkflowAPIAPI.NexusUpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the workflow | 
 **nexusWorkflow** | [**NexusWorkflow**](NexusWorkflow.md) | The details of the workflow | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

