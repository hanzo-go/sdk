# \FlowSolutionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowApplySolution**](FlowSolutionsAPI.md#FlowApplySolution) | **Post** /v1/flow/solutions/{id}/apply | Apply a solution to the project (EE)
[**FlowCreateSolution**](FlowSolutionsAPI.md#FlowCreateSolution) | **Post** /v1/flow/solutions | Create a solution from flows (EE)
[**FlowListSolutions**](FlowSolutionsAPI.md#FlowListSolutions) | **Get** /v1/flow/solutions | List packaged solutions (EE)



## FlowApplySolution

> map[string]interface{} FlowApplySolution(ctx, id).Execute()

Apply a solution to the project (EE)

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
	resp, r, err := apiClient.FlowSolutionsAPI.FlowApplySolution(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSolutionsAPI.FlowApplySolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowApplySolution`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSolutionsAPI.FlowApplySolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowApplySolutionRequest struct via the builder pattern


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


## FlowCreateSolution

> map[string]interface{} FlowCreateSolution(ctx).FlowCreateSolutionRequest(flowCreateSolutionRequest).Execute()

Create a solution from flows (EE)

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
	flowCreateSolutionRequest := *openapiclient.NewFlowCreateSolutionRequest([]string{"FlowIds_example"}) // FlowCreateSolutionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowSolutionsAPI.FlowCreateSolution(context.Background()).FlowCreateSolutionRequest(flowCreateSolutionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSolutionsAPI.FlowCreateSolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateSolution`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSolutionsAPI.FlowCreateSolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowCreateSolutionRequest** | [**FlowCreateSolutionRequest**](FlowCreateSolutionRequest.md) |  | 

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


## FlowListSolutions

> map[string]interface{} FlowListSolutions(ctx).Execute()

List packaged solutions (EE)

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
	resp, r, err := apiClient.FlowSolutionsAPI.FlowListSolutions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSolutionsAPI.FlowListSolutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListSolutions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSolutionsAPI.FlowListSolutions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListSolutionsRequest struct via the builder pattern


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

