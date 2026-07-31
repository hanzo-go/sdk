# \PipelinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Pipelines**](PipelinesAPI.md#CloudGetV1Pipelines) | **Get** /v1/pipelines | 
[**MlCreatePipeline**](PipelinesAPI.md#MlCreatePipeline) | **Post** /v1/ml/pipelines | Create a pipeline
[**MlGetPipeline**](PipelinesAPI.md#MlGetPipeline) | **Get** /v1/ml/pipelines/{pipeline_id} | Get pipeline details
[**MlListPipelineRuns**](PipelinesAPI.md#MlListPipelineRuns) | **Get** /v1/ml/pipelines/{pipeline_id}/runs | List pipeline runs
[**MlListPipelines**](PipelinesAPI.md#MlListPipelines) | **Get** /v1/ml/pipelines | List pipelines
[**MlStartPipelineRun**](PipelinesAPI.md#MlStartPipelineRun) | **Post** /v1/ml/pipelines/{pipeline_id}/runs | Start a pipeline run



## CloudGetV1Pipelines

> CloudGetV1Pipelines(ctx).Execute()



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
	r, err := apiClient.PipelinesAPI.CloudGetV1Pipelines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.CloudGetV1Pipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PipelinesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.PipelinesAPI.MlCreatePipeline(context.Background()).MlCreatePipelineRequest(mlCreatePipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.MlCreatePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlCreatePipeline`: MlPipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.MlCreatePipeline`: %v\n", resp)
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
	resp, r, err := apiClient.PipelinesAPI.MlGetPipeline(context.Background(), pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.MlGetPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetPipeline`: MlPipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.MlGetPipeline`: %v\n", resp)
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
	resp, r, err := apiClient.PipelinesAPI.MlListPipelineRuns(context.Background(), pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.MlListPipelineRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListPipelineRuns`: MlListPipelineRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.MlListPipelineRuns`: %v\n", resp)
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
	resp, r, err := apiClient.PipelinesAPI.MlListPipelines(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.MlListPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListPipelines`: MlListPipelines200Response
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.MlListPipelines`: %v\n", resp)
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
	resp, r, err := apiClient.PipelinesAPI.MlStartPipelineRun(context.Background(), pipelineId).MlStartPipelineRunRequest(mlStartPipelineRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.MlStartPipelineRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlStartPipelineRun`: MlPipelineRun
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.MlStartPipelineRun`: %v\n", resp)
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

