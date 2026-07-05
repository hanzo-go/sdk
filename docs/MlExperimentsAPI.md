# \MlExperimentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MlCreateExperiment**](MlExperimentsAPI.md#MlCreateExperiment) | **Post** /v1/ml/experiments | Create an experiment
[**MlGetRunMetrics**](MlExperimentsAPI.md#MlGetRunMetrics) | **Get** /v1/ml/experiments/{experiment_id}/runs/{run_id}/metrics | Get run metrics
[**MlListExperimentRuns**](MlExperimentsAPI.md#MlListExperimentRuns) | **Get** /v1/ml/experiments/{experiment_id}/runs | List experiment runs
[**MlListExperiments**](MlExperimentsAPI.md#MlListExperiments) | **Get** /v1/ml/experiments | List experiments
[**MlLogMetrics**](MlExperimentsAPI.md#MlLogMetrics) | **Post** /v1/ml/experiments/{experiment_id}/runs/{run_id}/metrics | Log metrics
[**MlStartExperimentRun**](MlExperimentsAPI.md#MlStartExperimentRun) | **Post** /v1/ml/experiments/{experiment_id}/runs | Start an experiment run



## MlCreateExperiment

> MlExperiment MlCreateExperiment(ctx).MlCreateExperimentRequest(mlCreateExperimentRequest).Execute()

Create an experiment

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
	mlCreateExperimentRequest := *openapiclient.NewMlCreateExperimentRequest("Name_example") // MlCreateExperimentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlExperimentsAPI.MlCreateExperiment(context.Background()).MlCreateExperimentRequest(mlCreateExperimentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlExperimentsAPI.MlCreateExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlCreateExperiment`: MlExperiment
	fmt.Fprintf(os.Stdout, "Response from `MlExperimentsAPI.MlCreateExperiment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlCreateExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mlCreateExperimentRequest** | [**MlCreateExperimentRequest**](MlCreateExperimentRequest.md) |  | 

### Return type

[**MlExperiment**](MlExperiment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlGetRunMetrics

> MlGetRunMetrics200Response MlGetRunMetrics(ctx, experimentId, runId).Execute()

Get run metrics

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
	experimentId := "experimentId_example" // string | 
	runId := "runId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlExperimentsAPI.MlGetRunMetrics(context.Background(), experimentId, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlExperimentsAPI.MlGetRunMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetRunMetrics`: MlGetRunMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `MlExperimentsAPI.MlGetRunMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experimentId** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlGetRunMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MlGetRunMetrics200Response**](MlGetRunMetrics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlListExperimentRuns

> MlListExperimentRuns200Response MlListExperimentRuns(ctx, experimentId).Execute()

List experiment runs

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
	experimentId := "experimentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlExperimentsAPI.MlListExperimentRuns(context.Background(), experimentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlExperimentsAPI.MlListExperimentRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListExperimentRuns`: MlListExperimentRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `MlExperimentsAPI.MlListExperimentRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlListExperimentRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MlListExperimentRuns200Response**](MlListExperimentRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlListExperiments

> MlListExperiments200Response MlListExperiments(ctx).Execute()

List experiments

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
	resp, r, err := apiClient.MlExperimentsAPI.MlListExperiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlExperimentsAPI.MlListExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListExperiments`: MlListExperiments200Response
	fmt.Fprintf(os.Stdout, "Response from `MlExperimentsAPI.MlListExperiments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMlListExperimentsRequest struct via the builder pattern


### Return type

[**MlListExperiments200Response**](MlListExperiments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlLogMetrics

> MlLogMetrics(ctx, experimentId, runId).RequestBody(requestBody).Execute()

Log metrics



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
	experimentId := "experimentId_example" // string | 
	runId := "runId_example" // string | 
	requestBody := map[string]float32{"key": float32(123)} // map[string]float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MlExperimentsAPI.MlLogMetrics(context.Background(), experimentId, runId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlExperimentsAPI.MlLogMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experimentId** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlLogMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]float32** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlStartExperimentRun

> MlExperimentRun MlStartExperimentRun(ctx, experimentId).MlStartExperimentRunRequest(mlStartExperimentRunRequest).Execute()

Start an experiment run

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
	experimentId := "experimentId_example" // string | 
	mlStartExperimentRunRequest := *openapiclient.NewMlStartExperimentRunRequest() // MlStartExperimentRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlExperimentsAPI.MlStartExperimentRun(context.Background(), experimentId).MlStartExperimentRunRequest(mlStartExperimentRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlExperimentsAPI.MlStartExperimentRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlStartExperimentRun`: MlExperimentRun
	fmt.Fprintf(os.Stdout, "Response from `MlExperimentsAPI.MlStartExperimentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlStartExperimentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mlStartExperimentRunRequest** | [**MlStartExperimentRunRequest**](MlStartExperimentRunRequest.md) |  | 

### Return type

[**MlExperimentRun**](MlExperimentRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

