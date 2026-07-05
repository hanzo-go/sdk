# \EnginePipelinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngineCreatePipeline**](EnginePipelinesAPI.md#EngineCreatePipeline) | **Post** /v1/engine/pipelines | Create ML pipeline
[**EngineCreatePipelineRun**](EnginePipelinesAPI.md#EngineCreatePipelineRun) | **Post** /v1/engine/pipelines/{id}/runs | Create pipeline run
[**EngineDeletePipeline**](EnginePipelinesAPI.md#EngineDeletePipeline) | **Delete** /v1/engine/pipelines/{id} | Delete pipeline
[**EngineGetPipeline**](EnginePipelinesAPI.md#EngineGetPipeline) | **Get** /v1/engine/pipelines/{id} | Get pipeline
[**EngineGetPipelineRun**](EnginePipelinesAPI.md#EngineGetPipelineRun) | **Get** /v1/engine/pipelines/{id}/runs/{run_id} | Get pipeline run
[**EngineListPipelineRuns**](EnginePipelinesAPI.md#EngineListPipelineRuns) | **Get** /v1/engine/pipelines/{id}/runs | List pipeline runs
[**EngineListPipelines**](EnginePipelinesAPI.md#EngineListPipelines) | **Get** /v1/engine/pipelines | List ML pipelines



## EngineCreatePipeline

> EnginePipeline EngineCreatePipeline(ctx).EnginePipelineCreate(enginePipelineCreate).Execute()

Create ML pipeline

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
	enginePipelineCreate := *openapiclient.NewEnginePipelineCreate("Name_example", map[string]interface{}{"key": interface{}(123)}) // EnginePipelineCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineCreatePipeline(context.Background()).EnginePipelineCreate(enginePipelineCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineCreatePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineCreatePipeline`: EnginePipeline
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineCreatePipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineCreatePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enginePipelineCreate** | [**EnginePipelineCreate**](EnginePipelineCreate.md) |  | 

### Return type

[**EnginePipeline**](EnginePipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineCreatePipelineRun

> EnginePipelineRun EngineCreatePipelineRun(ctx, id).EnginePipelineRunCreate(enginePipelineRunCreate).Execute()

Create pipeline run

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	enginePipelineRunCreate := *openapiclient.NewEnginePipelineRunCreate() // EnginePipelineRunCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineCreatePipelineRun(context.Background(), id).EnginePipelineRunCreate(enginePipelineRunCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineCreatePipelineRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineCreatePipelineRun`: EnginePipelineRun
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineCreatePipelineRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineCreatePipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enginePipelineRunCreate** | [**EnginePipelineRunCreate**](EnginePipelineRunCreate.md) |  | 

### Return type

[**EnginePipelineRun**](EnginePipelineRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineDeletePipeline

> map[string]interface{} EngineDeletePipeline(ctx, id).Execute()

Delete pipeline

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineDeletePipeline(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineDeletePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineDeletePipeline`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineDeletePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineDeletePipelineRequest struct via the builder pattern


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


## EngineGetPipeline

> EnginePipeline EngineGetPipeline(ctx, id).Execute()

Get pipeline

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineGetPipeline(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineGetPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetPipeline`: EnginePipeline
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineGetPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnginePipeline**](EnginePipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineGetPipelineRun

> EnginePipelineRun EngineGetPipelineRun(ctx, id, runId).Execute()

Get pipeline run

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineGetPipelineRun(context.Background(), id, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineGetPipelineRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetPipelineRun`: EnginePipelineRun
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineGetPipelineRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetPipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnginePipelineRun**](EnginePipelineRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListPipelineRuns

> EngineListPipelineRuns200Response EngineListPipelineRuns(ctx, id).Status(status).Page(page).PageSize(pageSize).Execute()

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	status := "status_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineListPipelineRuns(context.Background(), id).Status(status).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineListPipelineRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListPipelineRuns`: EngineListPipelineRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineListPipelineRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineListPipelineRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**EngineListPipelineRuns200Response**](EngineListPipelineRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListPipelines

> EngineListPipelines200Response EngineListPipelines(ctx).Page(page).PageSize(pageSize).Execute()

List ML pipelines

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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnginePipelinesAPI.EngineListPipelines(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnginePipelinesAPI.EngineListPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListPipelines`: EngineListPipelines200Response
	fmt.Fprintf(os.Stdout, "Response from `EnginePipelinesAPI.EngineListPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineListPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**EngineListPipelines200Response**](EngineListPipelines200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

