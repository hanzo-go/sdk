# \MetricsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceGetSaaSMetrics**](MetricsAPI.md#CommerceGetSaaSMetrics) | **Get** /v1/commerce/metrics/saas | SaaS operations snapshot (platform god-view)
[**FunctionsFunctionMetrics**](MetricsAPI.md#FunctionsFunctionMetrics) | **Get** /v1/functions/metrics | Invocation histogram + status breakdown
[**ObserveGetMetrics**](MetricsAPI.md#ObserveGetMetrics) | **Get** /v1/o11y/metrics | Per-org RED metrics and LLM usage for a product



## CommerceGetSaaSMetrics

> CommerceSaaSMetrics CommerceGetSaaSMetrics(ctx).Window(window).Limit(limit).Execute()

SaaS operations snapshot (platform god-view)



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
	window := "window_example" // string | Usage + new/churn window: 7d | 30d | 90d | mtd | all (default 30d) (optional) (default to "30d")
	limit := int32(56) // int32 | Cap for the customers list (1-200, default 20) (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.CommerceGetSaaSMetrics(context.Background()).Window(window).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.CommerceGetSaaSMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetSaaSMetrics`: CommerceSaaSMetrics
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.CommerceGetSaaSMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetSaaSMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **window** | **string** | Usage + new/churn window: 7d | 30d | 90d | mtd | all (default 30d) | [default to &quot;30d&quot;]
 **limit** | **int32** | Cap for the customers list (1-200, default 20) | [default to 20]

### Return type

[**CommerceSaaSMetrics**](CommerceSaaSMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsFunctionMetrics

> FunctionsMetrics FunctionsFunctionMetrics(ctx).Range_(range_).Execute()

Invocation histogram + status breakdown

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
	range_ := "range__example" // string |  (optional) (default to "24H")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.FunctionsFunctionMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.FunctionsFunctionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsFunctionMetrics`: FunctionsMetrics
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.FunctionsFunctionMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsFunctionMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** |  | [default to &quot;24H&quot;]

### Return type

[**FunctionsMetrics**](FunctionsMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObserveGetMetrics

> ObserveMetricsResponse ObserveGetMetrics(ctx).Product(product).Range_(range_).StepSec(stepSec).Execute()

Per-org RED metrics and LLM usage for a product



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
	product := "product_example" // string | Console product slug. Must match `^[a-z0-9][a-z0-9._-]{0,62}$`.
	range_ := int32(56) // int32 | Look-back range in seconds (default 3600, max 604800). (optional) (default to 3600)
	stepSec := int32(56) // int32 | Explicit bucket width in seconds (clamped to [30, 3600]). Omit to auto-derive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ObserveGetMetrics(context.Background()).Product(product).Range_(range_).StepSec(stepSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ObserveGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObserveGetMetrics`: ObserveMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ObserveGetMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObserveGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Console product slug. Must match &#x60;^[a-z0-9][a-z0-9._-]{0,62}$&#x60;. | 
 **range_** | **int32** | Look-back range in seconds (default 3600, max 604800). | [default to 3600]
 **stepSec** | **int32** | Explicit bucket width in seconds (clamped to [30, 3600]). Omit to auto-derive. | 

### Return type

[**ObserveMetricsResponse**](ObserveMetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

