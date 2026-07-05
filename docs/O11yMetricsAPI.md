# \O11yMetricsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**O11yPromLabelValues**](O11yMetricsAPI.md#O11yPromLabelValues) | **Get** /v1/o11y/label/{name}/values | Get metric label values
[**O11yPromLabels**](O11yMetricsAPI.md#O11yPromLabels) | **Get** /v1/o11y/labels | List metric label names
[**O11yPromQuery**](O11yMetricsAPI.md#O11yPromQuery) | **Get** /v1/o11y/query | Instant PromQL query
[**O11yPromQueryRange**](O11yMetricsAPI.md#O11yPromQueryRange) | **Get** /v1/o11y/query_range | Range PromQL query
[**O11yPromSeries**](O11yMetricsAPI.md#O11yPromSeries) | **Get** /v1/o11y/series | Find series by label matchers



## O11yPromLabelValues

> O11yPromLabelValues200Response O11yPromLabelValues(ctx, name).Start(start).End(end).Match(match).Execute()

Get metric label values

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
	name := "name_example" // string | 
	start := "start_example" // string |  (optional)
	end := "end_example" // string |  (optional)
	match := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yMetricsAPI.O11yPromLabelValues(context.Background(), name).Start(start).End(end).Match(match).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yMetricsAPI.O11yPromLabelValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yPromLabelValues`: O11yPromLabelValues200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yMetricsAPI.O11yPromLabelValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yPromLabelValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** |  | 
 **end** | **string** |  | 
 **match** | **[]string** |  | 

### Return type

[**O11yPromLabelValues200Response**](O11yPromLabelValues200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yPromLabels

> O11yPromLabelValues200Response O11yPromLabels(ctx).Start(start).End(end).Match(match).Execute()

List metric label names

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
	start := "start_example" // string |  (optional)
	end := "end_example" // string |  (optional)
	match := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yMetricsAPI.O11yPromLabels(context.Background()).Start(start).End(end).Match(match).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yMetricsAPI.O11yPromLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yPromLabels`: O11yPromLabelValues200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yMetricsAPI.O11yPromLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yPromLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **match** | **[]string** |  | 

### Return type

[**O11yPromLabelValues200Response**](O11yPromLabelValues200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yPromQuery

> O11yMetricQueryResult O11yPromQuery(ctx).Query(query).Time(time).Timeout(timeout).Execute()

Instant PromQL query

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
	query := "query_example" // string | PromQL expression
	time := "time_example" // string | Evaluation timestamp (optional)
	timeout := "timeout_example" // string | Evaluation timeout (e.g. 30s) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yMetricsAPI.O11yPromQuery(context.Background()).Query(query).Time(time).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yMetricsAPI.O11yPromQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yPromQuery`: O11yMetricQueryResult
	fmt.Fprintf(os.Stdout, "Response from `O11yMetricsAPI.O11yPromQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yPromQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | PromQL expression | 
 **time** | **string** | Evaluation timestamp | 
 **timeout** | **string** | Evaluation timeout (e.g. 30s) | 

### Return type

[**O11yMetricQueryResult**](O11yMetricQueryResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yPromQueryRange

> O11yMetricQueryResult O11yPromQueryRange(ctx).Query(query).Start(start).End(end).Step(step).Timeout(timeout).Execute()

Range PromQL query

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
	query := "query_example" // string | 
	start := "start_example" // string | 
	end := "end_example" // string | 
	step := "step_example" // string | Query resolution step (e.g. 15s, 1m)
	timeout := "timeout_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yMetricsAPI.O11yPromQueryRange(context.Background()).Query(query).Start(start).End(end).Step(step).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yMetricsAPI.O11yPromQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yPromQueryRange`: O11yMetricQueryResult
	fmt.Fprintf(os.Stdout, "Response from `O11yMetricsAPI.O11yPromQueryRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yPromQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 
 **step** | **string** | Query resolution step (e.g. 15s, 1m) | 
 **timeout** | **string** |  | 

### Return type

[**O11yMetricQueryResult**](O11yMetricQueryResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yPromSeries

> O11ySeriesResult O11yPromSeries(ctx).Match(match).Start(start).End(end).Execute()

Find series by label matchers

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
	match := []string{"Inner_example"} // []string | Series selector (at least one required)
	start := "start_example" // string |  (optional)
	end := "end_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yMetricsAPI.O11yPromSeries(context.Background()).Match(match).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yMetricsAPI.O11yPromSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yPromSeries`: O11ySeriesResult
	fmt.Fprintf(os.Stdout, "Response from `O11yMetricsAPI.O11yPromSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yPromSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **match** | **[]string** | Series selector (at least one required) | 
 **start** | **string** |  | 
 **end** | **string** |  | 

### Return type

[**O11ySeriesResult**](O11ySeriesResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

