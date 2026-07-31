# \ExperimentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Experiments**](ExperimentsAPI.md#CloudGetV1Experiments) | **Get** /v1/experiments | 
[**CloudGetV1ExperimentsById**](ExperimentsAPI.md#CloudGetV1ExperimentsById) | **Get** /v1/experiments/{id} | 
[**CloudGetV1ExperimentsByIdAssign**](ExperimentsAPI.md#CloudGetV1ExperimentsByIdAssign) | **Get** /v1/experiments/{id}/assign | 
[**CloudGetV1ExperimentsHealth**](ExperimentsAPI.md#CloudGetV1ExperimentsHealth) | **Get** /v1/experiments/health | 
[**CloudPostV1Experiments**](ExperimentsAPI.md#CloudPostV1Experiments) | **Post** /v1/experiments | 
[**CloudPostV1ExperimentsByIdAnalyze**](ExperimentsAPI.md#CloudPostV1ExperimentsByIdAnalyze) | **Post** /v1/experiments/{id}/analyze | 
[**CloudPostV1ExperimentsByIdDecide**](ExperimentsAPI.md#CloudPostV1ExperimentsByIdDecide) | **Post** /v1/experiments/{id}/decide | 
[**MlCreateExperiment**](ExperimentsAPI.md#MlCreateExperiment) | **Post** /v1/ml/experiments | Create an experiment
[**MlGetRunMetrics**](ExperimentsAPI.md#MlGetRunMetrics) | **Get** /v1/ml/experiments/{experiment_id}/runs/{run_id}/metrics | Get run metrics
[**MlListExperimentRuns**](ExperimentsAPI.md#MlListExperimentRuns) | **Get** /v1/ml/experiments/{experiment_id}/runs | List experiment runs
[**MlListExperiments**](ExperimentsAPI.md#MlListExperiments) | **Get** /v1/ml/experiments | List experiments
[**MlLogMetrics**](ExperimentsAPI.md#MlLogMetrics) | **Post** /v1/ml/experiments/{experiment_id}/runs/{run_id}/metrics | Log metrics
[**MlStartExperimentRun**](ExperimentsAPI.md#MlStartExperimentRun) | **Post** /v1/ml/experiments/{experiment_id}/runs | Start an experiment run



## CloudGetV1Experiments

> CloudGetV1Experiments(ctx).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudGetV1Experiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudGetV1Experiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ExperimentsRequest struct via the builder pattern


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


## CloudGetV1ExperimentsById

> CloudGetV1ExperimentsById(ctx, id).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudGetV1ExperimentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudGetV1ExperimentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ExperimentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1ExperimentsByIdAssign

> CloudGetV1ExperimentsByIdAssign(ctx, id).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudGetV1ExperimentsByIdAssign(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudGetV1ExperimentsByIdAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ExperimentsByIdAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1ExperimentsHealth

> CloudGetV1ExperimentsHealth(ctx).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudGetV1ExperimentsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudGetV1ExperimentsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ExperimentsHealthRequest struct via the builder pattern


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


## CloudPostV1Experiments

> CloudPostV1Experiments(ctx).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudPostV1Experiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudPostV1Experiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ExperimentsRequest struct via the builder pattern


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


## CloudPostV1ExperimentsByIdAnalyze

> CloudPostV1ExperimentsByIdAnalyze(ctx, id).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudPostV1ExperimentsByIdAnalyze(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudPostV1ExperimentsByIdAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ExperimentsByIdAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudPostV1ExperimentsByIdDecide

> CloudPostV1ExperimentsByIdDecide(ctx, id).Execute()



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
	r, err := apiClient.ExperimentsAPI.CloudPostV1ExperimentsByIdDecide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.CloudPostV1ExperimentsByIdDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ExperimentsByIdDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
	resp, r, err := apiClient.ExperimentsAPI.MlCreateExperiment(context.Background()).MlCreateExperimentRequest(mlCreateExperimentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.MlCreateExperiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlCreateExperiment`: MlExperiment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.MlCreateExperiment`: %v\n", resp)
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
	resp, r, err := apiClient.ExperimentsAPI.MlGetRunMetrics(context.Background(), experimentId, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.MlGetRunMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetRunMetrics`: MlGetRunMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.MlGetRunMetrics`: %v\n", resp)
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
	resp, r, err := apiClient.ExperimentsAPI.MlListExperimentRuns(context.Background(), experimentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.MlListExperimentRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListExperimentRuns`: MlListExperimentRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.MlListExperimentRuns`: %v\n", resp)
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
	resp, r, err := apiClient.ExperimentsAPI.MlListExperiments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.MlListExperiments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListExperiments`: MlListExperiments200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.MlListExperiments`: %v\n", resp)
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
	r, err := apiClient.ExperimentsAPI.MlLogMetrics(context.Background(), experimentId, runId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.MlLogMetrics``: %v\n", err)
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
	resp, r, err := apiClient.ExperimentsAPI.MlStartExperimentRun(context.Background(), experimentId).MlStartExperimentRunRequest(mlStartExperimentRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.MlStartExperimentRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlStartExperimentRun`: MlExperimentRun
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.MlStartExperimentRun`: %v\n", resp)
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

