# \MlPipelinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MlCreatePipeline**](MlPipelinesAPI.md#MlCreatePipeline) | **Post** /v1/ml/pipelines | Create a pipeline
[**MlGetPipeline**](MlPipelinesAPI.md#MlGetPipeline) | **Get** /v1/ml/pipelines/{pipeline_id} | Get pipeline details
[**MlListPipelineRuns**](MlPipelinesAPI.md#MlListPipelineRuns) | **Get** /v1/ml/pipelines/{pipeline_id}/runs | List pipeline runs
[**MlListPipelines**](MlPipelinesAPI.md#MlListPipelines) | **Get** /v1/ml/pipelines | List pipelines
[**MlStartPipelineRun**](MlPipelinesAPI.md#MlStartPipelineRun) | **Post** /v1/ml/pipelines/{pipeline_id}/runs | Start a pipeline run



## MlCreatePipeline

> MlPipeline MlCreatePipeline(ctx).MlCreatePipelineRequest(mlCreatePipelineRequest).Execute()

Create a pipeline



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
	mlCreatePipelineRequest := *openapiclient.NewMlCreatePipelineRequest("Name_example") // MlCreatePipelineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlPipelinesAPI.MlCreatePipeline(context.Background()).MlCreatePipelineRequest(mlCreatePipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlPipelinesAPI.MlCreatePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlCreatePipeline`: MlPipeline
	fmt.Fprintf(os.Stdout, "Response from `MlPipelinesAPI.MlCreatePipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlCreatePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mlCreatePipelineRequest** | [**MlCreatePipelineRequest**](MlCreatePipelineRequest.md) |  | 

### Return type

[**MlPipeline**](MlPipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlGetPipeline

> MlPipeline MlGetPipeline(ctx, pipelineId).Execute()

Get pipeline details

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
	pipelineId := "pipelineId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlPipelinesAPI.MlGetPipeline(context.Background(), pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlPipelinesAPI.MlGetPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetPipeline`: MlPipeline
	fmt.Fprintf(os.Stdout, "Response from `MlPipelinesAPI.MlGetPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlGetPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MlPipeline**](MlPipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlListPipelineRuns

> MlListPipelineRuns200Response MlListPipelineRuns(ctx, pipelineId).Execute()

List pipeline runs

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
	pipelineId := "pipelineId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlPipelinesAPI.MlListPipelineRuns(context.Background(), pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlPipelinesAPI.MlListPipelineRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListPipelineRuns`: MlListPipelineRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `MlPipelinesAPI.MlListPipelineRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlListPipelineRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MlListPipelineRuns200Response**](MlListPipelineRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlListPipelines

> MlListPipelines200Response MlListPipelines(ctx).Status(status).Limit(limit).Execute()

List pipelines

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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlPipelinesAPI.MlListPipelines(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlPipelinesAPI.MlListPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListPipelines`: MlListPipelines200Response
	fmt.Fprintf(os.Stdout, "Response from `MlPipelinesAPI.MlListPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlListPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**MlListPipelines200Response**](MlListPipelines200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlStartPipelineRun

> MlPipelineRun MlStartPipelineRun(ctx, pipelineId).MlStartPipelineRunRequest(mlStartPipelineRunRequest).Execute()

Start a pipeline run

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
	pipelineId := "pipelineId_example" // string | 
	mlStartPipelineRunRequest := *openapiclient.NewMlStartPipelineRunRequest() // MlStartPipelineRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlPipelinesAPI.MlStartPipelineRun(context.Background(), pipelineId).MlStartPipelineRunRequest(mlStartPipelineRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlPipelinesAPI.MlStartPipelineRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlStartPipelineRun`: MlPipelineRun
	fmt.Fprintf(os.Stdout, "Response from `MlPipelinesAPI.MlStartPipelineRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlStartPipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mlStartPipelineRunRequest** | [**MlStartPipelineRunRequest**](MlStartPipelineRunRequest.md) |  | 

### Return type

[**MlPipelineRun**](MlPipelineRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

