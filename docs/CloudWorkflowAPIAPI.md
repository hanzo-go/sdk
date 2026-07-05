# \CloudWorkflowAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddWorkflow**](CloudWorkflowAPIAPI.md#CloudApiControllerAddWorkflow) | **Post** /v1/cloud/add-workflow | Api Controller Add Workflow
[**CloudApiControllerDeleteWorkflow**](CloudWorkflowAPIAPI.md#CloudApiControllerDeleteWorkflow) | **Post** /v1/cloud/delete-workflow | Api Controller Delete Workflow
[**CloudApiControllerGetGlobalWorkflows**](CloudWorkflowAPIAPI.md#CloudApiControllerGetGlobalWorkflows) | **Get** /v1/cloud/get-global-workflows | Api Controller Get Global Workflows
[**CloudApiControllerGetWorkflow**](CloudWorkflowAPIAPI.md#CloudApiControllerGetWorkflow) | **Get** /v1/cloud/get-workflow | Api Controller Get Workflow
[**CloudApiControllerGetWorkflows**](CloudWorkflowAPIAPI.md#CloudApiControllerGetWorkflows) | **Get** /v1/cloud/get-workflows | Api Controller Get Workflows
[**CloudApiControllerUpdateWorkflow**](CloudWorkflowAPIAPI.md#CloudApiControllerUpdateWorkflow) | **Post** /v1/cloud/update-workflow | Api Controller Update Workflow



## CloudApiControllerAddWorkflow

> CloudControllersResponse CloudApiControllerAddWorkflow(ctx).CloudObjectWorkflow(cloudObjectWorkflow).Execute()

Api Controller Add Workflow



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
	cloudObjectWorkflow := *openapiclient.NewCloudObjectWorkflow() // CloudObjectWorkflow | The details of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudWorkflowAPIAPI.CloudApiControllerAddWorkflow(context.Background()).CloudObjectWorkflow(cloudObjectWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkflowAPIAPI.CloudApiControllerAddWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddWorkflow`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkflowAPIAPI.CloudApiControllerAddWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectWorkflow** | [**CloudObjectWorkflow**](CloudObjectWorkflow.md) | The details of the workflow | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteWorkflow

> CloudControllersResponse CloudApiControllerDeleteWorkflow(ctx).CloudObjectWorkflow(cloudObjectWorkflow).Execute()

Api Controller Delete Workflow



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
	cloudObjectWorkflow := *openapiclient.NewCloudObjectWorkflow() // CloudObjectWorkflow | The details of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudWorkflowAPIAPI.CloudApiControllerDeleteWorkflow(context.Background()).CloudObjectWorkflow(cloudObjectWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkflowAPIAPI.CloudApiControllerDeleteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteWorkflow`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkflowAPIAPI.CloudApiControllerDeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectWorkflow** | [**CloudObjectWorkflow**](CloudObjectWorkflow.md) | The details of the workflow | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalWorkflows

> []CloudObjectWorkflow CloudApiControllerGetGlobalWorkflows(ctx).Execute()

Api Controller Get Global Workflows



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
	resp, r, err := apiClient.CloudWorkflowAPIAPI.CloudApiControllerGetGlobalWorkflows(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkflowAPIAPI.CloudApiControllerGetGlobalWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalWorkflows`: []CloudObjectWorkflow
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkflowAPIAPI.CloudApiControllerGetGlobalWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalWorkflowsRequest struct via the builder pattern


### Return type

[**[]CloudObjectWorkflow**](CloudObjectWorkflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetWorkflow

> CloudObjectWorkflow CloudApiControllerGetWorkflow(ctx).Id(id).Execute()

Api Controller Get Workflow



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
	id := "id_example" // string | The id (owner/name) of workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudWorkflowAPIAPI.CloudApiControllerGetWorkflow(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkflowAPIAPI.CloudApiControllerGetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetWorkflow`: CloudObjectWorkflow
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkflowAPIAPI.CloudApiControllerGetWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of workflow | 

### Return type

[**CloudObjectWorkflow**](CloudObjectWorkflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetWorkflows

> []CloudObjectWorkflow CloudApiControllerGetWorkflows(ctx).Owner(owner).Execute()

Api Controller Get Workflows



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
	owner := "owner_example" // string | The owner of workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudWorkflowAPIAPI.CloudApiControllerGetWorkflows(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkflowAPIAPI.CloudApiControllerGetWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetWorkflows`: []CloudObjectWorkflow
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkflowAPIAPI.CloudApiControllerGetWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of workflow | 

### Return type

[**[]CloudObjectWorkflow**](CloudObjectWorkflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateWorkflow

> CloudControllersResponse CloudApiControllerUpdateWorkflow(ctx).Id(id).CloudObjectWorkflow(cloudObjectWorkflow).Execute()

Api Controller Update Workflow



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
	cloudObjectWorkflow := *openapiclient.NewCloudObjectWorkflow() // CloudObjectWorkflow | The details of the workflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudWorkflowAPIAPI.CloudApiControllerUpdateWorkflow(context.Background()).Id(id).CloudObjectWorkflow(cloudObjectWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkflowAPIAPI.CloudApiControllerUpdateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateWorkflow`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkflowAPIAPI.CloudApiControllerUpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the workflow | 
 **cloudObjectWorkflow** | [**CloudObjectWorkflow**](CloudObjectWorkflow.md) | The details of the workflow | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

