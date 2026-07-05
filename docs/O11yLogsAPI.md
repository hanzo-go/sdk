# \O11yLogsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**O11yLokiLabelValues**](O11yLogsAPI.md#O11yLokiLabelValues) | **Get** /v1/o11y/logs/labels/{name}/values | Get label values
[**O11yLokiLabels**](O11yLogsAPI.md#O11yLokiLabels) | **Get** /v1/o11y/logs/labels | List log labels
[**O11yLokiPush**](O11yLogsAPI.md#O11yLokiPush) | **Post** /v1/o11y/logs/push | Push log entries
[**O11yLokiQuery**](O11yLogsAPI.md#O11yLokiQuery) | **Get** /v1/o11y/logs/query | Instant LogQL query
[**O11yLokiQueryRange**](O11yLogsAPI.md#O11yLokiQueryRange) | **Get** /v1/o11y/logs/query_range | Range LogQL query
[**O11yLokiTail**](O11yLogsAPI.md#O11yLokiTail) | **Get** /v1/o11y/logs/tail | Live tail logs (WebSocket)



## O11yLokiLabelValues

> O11yPromLabelValues200Response O11yLokiLabelValues(ctx, name).Start(start).End(end).Query(query).Execute()

Get label values

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
	query := "query_example" // string | LogQL stream selector to filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yLogsAPI.O11yLokiLabelValues(context.Background(), name).Start(start).End(end).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yLogsAPI.O11yLokiLabelValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yLokiLabelValues`: O11yPromLabelValues200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yLogsAPI.O11yLokiLabelValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yLokiLabelValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** |  | 
 **end** | **string** |  | 
 **query** | **string** | LogQL stream selector to filter | 

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


## O11yLokiLabels

> O11yPromLabelValues200Response O11yLokiLabels(ctx).Start(start).End(end).Execute()

List log labels

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yLogsAPI.O11yLokiLabels(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yLogsAPI.O11yLokiLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yLokiLabels`: O11yPromLabelValues200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yLogsAPI.O11yLokiLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yLokiLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 

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


## O11yLokiPush

> O11yLokiPush(ctx).O11yLogPushRequest(o11yLogPushRequest).Execute()

Push log entries

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
	o11yLogPushRequest := *openapiclient.NewO11yLogPushRequest([]openapiclient.O11yLogPushRequestStreamsInner{*openapiclient.NewO11yLogPushRequestStreamsInner(map[string]string{"key": "Inner_example"}, [][]string{[]string{"Values_example"}})}) // O11yLogPushRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yLogsAPI.O11yLokiPush(context.Background()).O11yLogPushRequest(o11yLogPushRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yLogsAPI.O11yLokiPush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yLokiPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yLogPushRequest** | [**O11yLogPushRequest**](O11yLogPushRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yLokiQuery

> O11yLogQueryResult O11yLokiQuery(ctx).Query(query).Time(time).Limit(limit).Direction(direction).Execute()

Instant LogQL query



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
	query := "query_example" // string | LogQL query expression
	time := "time_example" // string | Evaluation timestamp (Unix epoch or RFC3339, default now) (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	direction := "direction_example" // string |  (optional) (default to "backward")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yLogsAPI.O11yLokiQuery(context.Background()).Query(query).Time(time).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yLogsAPI.O11yLokiQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yLokiQuery`: O11yLogQueryResult
	fmt.Fprintf(os.Stdout, "Response from `O11yLogsAPI.O11yLokiQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yLokiQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | LogQL query expression | 
 **time** | **string** | Evaluation timestamp (Unix epoch or RFC3339, default now) | 
 **limit** | **int32** |  | [default to 100]
 **direction** | **string** |  | [default to &quot;backward&quot;]

### Return type

[**O11yLogQueryResult**](O11yLogQueryResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yLokiQueryRange

> O11yLogQueryResult O11yLokiQueryRange(ctx).Query(query).Start(start).End(end).Step(step).Limit(limit).Direction(direction).Execute()

Range LogQL query



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
	start := "start_example" // string | Start timestamp (Unix epoch or RFC3339) (optional)
	end := "end_example" // string | End timestamp (Unix epoch or RFC3339) (optional)
	step := "step_example" // string | Query step (e.g. 15s, 1m, 5m) (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	direction := "direction_example" // string |  (optional) (default to "backward")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yLogsAPI.O11yLokiQueryRange(context.Background()).Query(query).Start(start).End(end).Step(step).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yLogsAPI.O11yLokiQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yLokiQueryRange`: O11yLogQueryResult
	fmt.Fprintf(os.Stdout, "Response from `O11yLogsAPI.O11yLokiQueryRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yLokiQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **start** | **string** | Start timestamp (Unix epoch or RFC3339) | 
 **end** | **string** | End timestamp (Unix epoch or RFC3339) | 
 **step** | **string** | Query step (e.g. 15s, 1m, 5m) | 
 **limit** | **int32** |  | [default to 100]
 **direction** | **string** |  | [default to &quot;backward&quot;]

### Return type

[**O11yLogQueryResult**](O11yLogQueryResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yLokiTail

> O11yLokiTail(ctx).Query(query).DelayFor(delayFor).Limit(limit).Execute()

Live tail logs (WebSocket)



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
	query := "query_example" // string | LogQL query expression
	delayFor := int32(56) // int32 | Seconds to delay entries for deduplication (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yLogsAPI.O11yLokiTail(context.Background()).Query(query).DelayFor(delayFor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yLogsAPI.O11yLokiTail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yLokiTailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | LogQL query expression | 
 **delayFor** | **int32** | Seconds to delay entries for deduplication | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

