# \RunsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationsGetRun**](RunsAPI.md#AutomationsGetRun) | **Get** /v1/automations/runs/{id} | Get a run (non-terminal status is refreshed from the engine)
[**AutomationsListRuns**](RunsAPI.md#AutomationsListRuns) | **Get** /v1/automations/runs | List runs
[**AutomationsResumeRun**](RunsAPI.md#AutomationsResumeRun) | **Post** /v1/automations/runs/{id}/resume | Resume a paused run
[**AutomationsRunFlow**](RunsAPI.md#AutomationsRunFlow) | **Post** /v1/automations/flows/{id}/run | Start a durable run of a flow&#39;s runnable version
[**EvalsGetV1EvalsHealth**](RunsAPI.md#EvalsGetV1EvalsHealth) | **Get** /v1/evals/health | Health check
[**EvalsPostV1EvalsRuns**](RunsAPI.md#EvalsPostV1EvalsRuns) | **Post** /v1/evals/runs | Run a dataset against a model with an LLM-as-a-Judge



## AutomationsGetRun

> AutomationsFlowRun AutomationsGetRun(ctx, id).Execute()

Get a run (non-terminal status is refreshed from the engine)

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
	resp, r, err := apiClient.RunsAPI.AutomationsGetRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.AutomationsGetRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsGetRun`: AutomationsFlowRun
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.AutomationsGetRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsGetRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationsFlowRun**](AutomationsFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsListRuns

> AutomationsListRuns200Response AutomationsListRuns(ctx).FlowId(flowId).Limit(limit).Execute()

List runs

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
	flowId := "flowId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.AutomationsListRuns(context.Background()).FlowId(flowId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.AutomationsListRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsListRuns`: AutomationsListRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.AutomationsListRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsListRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** |  | 
 **limit** | **int32** |  | [default to 200]

### Return type

[**AutomationsListRuns200Response**](AutomationsListRuns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsResumeRun

> AutomationsResumeRun200Response AutomationsResumeRun(ctx, id).Body(body).Execute()

Resume a paused run



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
	body := interface{}(987) // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.AutomationsResumeRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.AutomationsResumeRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsResumeRun`: AutomationsResumeRun200Response
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.AutomationsResumeRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsResumeRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**AutomationsResumeRun200Response**](AutomationsResumeRun200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationsRunFlow

> AutomationsFlowRun AutomationsRunFlow(ctx, id).Execute()

Start a durable run of a flow's runnable version

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
	resp, r, err := apiClient.RunsAPI.AutomationsRunFlow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.AutomationsRunFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsRunFlow`: AutomationsFlowRun
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.AutomationsRunFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsRunFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationsFlowRun**](AutomationsFlowRun.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvalsGetV1EvalsHealth

> EvalsGetV1EvalsHealth200Response EvalsGetV1EvalsHealth(ctx).Execute()

Health check

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
	resp, r, err := apiClient.RunsAPI.EvalsGetV1EvalsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.EvalsGetV1EvalsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvalsGetV1EvalsHealth`: EvalsGetV1EvalsHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.EvalsGetV1EvalsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEvalsGetV1EvalsHealthRequest struct via the builder pattern


### Return type

[**EvalsGetV1EvalsHealth200Response**](EvalsGetV1EvalsHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvalsPostV1EvalsRuns

> EvalsRunSummary EvalsPostV1EvalsRuns(ctx).EvalsRunRequest(evalsRunRequest).Execute()

Run a dataset against a model with an LLM-as-a-Judge



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
	evalsRunRequest := *openapiclient.NewEvalsRunRequest("Dataset_example", "Model_example") // EvalsRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.EvalsPostV1EvalsRuns(context.Background()).EvalsRunRequest(evalsRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.EvalsPostV1EvalsRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvalsPostV1EvalsRuns`: EvalsRunSummary
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.EvalsPostV1EvalsRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvalsPostV1EvalsRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalsRunRequest** | [**EvalsRunRequest**](EvalsRunRequest.md) |  | 

### Return type

[**EvalsRunSummary**](EvalsRunSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

