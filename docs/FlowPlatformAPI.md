# \FlowPlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowGetAnalytics**](FlowPlatformAPI.md#FlowGetAnalytics) | **Get** /v1/flow/analytics | Get platform analytics data
[**FlowGetPlatform**](FlowPlatformAPI.md#FlowGetPlatform) | **Get** /v1/flow/platforms/{id} | Get platform settings
[**FlowGetQueueMetrics**](FlowPlatformAPI.md#FlowGetQueueMetrics) | **Get** /v1/flow/queue-metrics | Get worker queue metrics (EE)
[**FlowUpdatePlatform**](FlowPlatformAPI.md#FlowUpdatePlatform) | **Post** /v1/flow/platforms/{id} | Update platform settings



## FlowGetAnalytics

> map[string]interface{} FlowGetAnalytics(ctx).Execute()

Get platform analytics data

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
	resp, r, err := apiClient.FlowPlatformAPI.FlowGetAnalytics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPlatformAPI.FlowGetAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetAnalytics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPlatformAPI.FlowGetAnalytics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetAnalyticsRequest struct via the builder pattern


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


## FlowGetPlatform

> FlowPlatform FlowGetPlatform(ctx, id).Execute()

Get platform settings

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
	resp, r, err := apiClient.FlowPlatformAPI.FlowGetPlatform(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPlatformAPI.FlowGetPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetPlatform`: FlowPlatform
	fmt.Fprintf(os.Stdout, "Response from `FlowPlatformAPI.FlowGetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowPlatform**](FlowPlatform.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetQueueMetrics

> map[string]interface{} FlowGetQueueMetrics(ctx).Execute()

Get worker queue metrics (EE)

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
	resp, r, err := apiClient.FlowPlatformAPI.FlowGetQueueMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPlatformAPI.FlowGetQueueMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetQueueMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPlatformAPI.FlowGetQueueMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetQueueMetricsRequest struct via the builder pattern


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


## FlowUpdatePlatform

> map[string]interface{} FlowUpdatePlatform(ctx, id).FlowUpdatePlatformRequest(flowUpdatePlatformRequest).Execute()

Update platform settings

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
	flowUpdatePlatformRequest := *openapiclient.NewFlowUpdatePlatformRequest() // FlowUpdatePlatformRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowPlatformAPI.FlowUpdatePlatform(context.Background(), id).FlowUpdatePlatformRequest(flowUpdatePlatformRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowPlatformAPI.FlowUpdatePlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpdatePlatform`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowPlatformAPI.FlowUpdatePlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpdatePlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowUpdatePlatformRequest** | [**FlowUpdatePlatformRequest**](FlowUpdatePlatformRequest.md) |  | 

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

