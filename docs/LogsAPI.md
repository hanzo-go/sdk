# \LogsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeGetFunctionLogs**](LogsAPI.md#EdgeGetFunctionLogs) | **Get** /v1/edge/functions/{slug}/logs | Get function logs
[**ObserveGetLogs**](LogsAPI.md#ObserveGetLogs) | **Get** /v1/o11y/logs | Live org-scoped log stream for a product
[**SearchGetStderrLogs**](LogsAPI.md#SearchGetStderrLogs) | **Get** /v1/search/logs/stderr | Get stderr log configuration
[**SearchResetStderrLogs**](LogsAPI.md#SearchResetStderrLogs) | **Delete** /v1/search/logs/stderr | Reset stderr log level to default
[**SearchStreamLogs**](LogsAPI.md#SearchStreamLogs) | **Post** /v1/search/logs/stream | Stream logs
[**SearchUpdateStderrLogs**](LogsAPI.md#SearchUpdateStderrLogs) | **Put** /v1/search/logs/stderr | Update stderr log level



## EdgeGetFunctionLogs

> []EdgeLogEntry EdgeGetFunctionLogs(ctx, slug).Since(since).Until(until).Level(level).Limit(limit).Execute()

Get function logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	slug := "slug_example" // string | 
	since := time.Now() // time.Time | Return logs after this timestamp (optional)
	until := time.Now() // time.Time |  (optional)
	level := "level_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.EdgeGetFunctionLogs(context.Background(), slug).Since(since).Until(until).Level(level).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.EdgeGetFunctionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeGetFunctionLogs`: []EdgeLogEntry
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.EdgeGetFunctionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeGetFunctionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **time.Time** | Return logs after this timestamp | 
 **until** | **time.Time** |  | 
 **level** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]EdgeLogEntry**](EdgeLogEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObserveGetLogs

> ObserveLogsResponse ObserveGetLogs(ctx).Product(product).SinceNs(sinceNs).Window(window).Limit(limit).Execute()

Live org-scoped log stream for a product



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
	sinceNs := int64(789) // int64 | Nanosecond cursor from the prior response's `nextCursor`; 0/absent starts a fresh window. (optional)
	window := int32(56) // int32 | Initial look-back window in seconds when no cursor is supplied (default 900, max 86400). (optional) (default to 900)
	limit := int32(56) // int32 | Max lines returned (default 200, max 1000). (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.ObserveGetLogs(context.Background()).Product(product).SinceNs(sinceNs).Window(window).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ObserveGetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObserveGetLogs`: ObserveLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.ObserveGetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObserveGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Console product slug. Must match &#x60;^[a-z0-9][a-z0-9._-]{0,62}$&#x60;. | 
 **sinceNs** | **int64** | Nanosecond cursor from the prior response&#39;s &#x60;nextCursor&#x60;; 0/absent starts a fresh window. | 
 **window** | **int32** | Initial look-back window in seconds when no cursor is supplied (default 900, max 86400). | [default to 900]
 **limit** | **int32** | Max lines returned (default 200, max 1000). | [default to 200]

### Return type

[**ObserveLogsResponse**](ObserveLogsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGetStderrLogs

> SearchGetStderrLogs200Response SearchGetStderrLogs(ctx).Execute()

Get stderr log configuration

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
	resp, r, err := apiClient.LogsAPI.SearchGetStderrLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.SearchGetStderrLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetStderrLogs`: SearchGetStderrLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.SearchGetStderrLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetStderrLogsRequest struct via the builder pattern


### Return type

[**SearchGetStderrLogs200Response**](SearchGetStderrLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchResetStderrLogs

> SearchGetStderrLogs200Response SearchResetStderrLogs(ctx).Execute()

Reset stderr log level to default

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
	resp, r, err := apiClient.LogsAPI.SearchResetStderrLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.SearchResetStderrLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchResetStderrLogs`: SearchGetStderrLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.SearchResetStderrLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchResetStderrLogsRequest struct via the builder pattern


### Return type

[**SearchGetStderrLogs200Response**](SearchGetStderrLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStreamLogs

> string SearchStreamLogs(ctx).SearchStreamLogsRequest(searchStreamLogsRequest).Execute()

Stream logs

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
	searchStreamLogsRequest := *openapiclient.NewSearchStreamLogsRequest("Target_example", "Mode_example") // SearchStreamLogsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.SearchStreamLogs(context.Background()).SearchStreamLogsRequest(searchStreamLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.SearchStreamLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchStreamLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.SearchStreamLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStreamLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchStreamLogsRequest** | [**SearchStreamLogsRequest**](SearchStreamLogsRequest.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateStderrLogs

> SearchGetStderrLogs200Response SearchUpdateStderrLogs(ctx).SearchUpdateStderrLogsRequest(searchUpdateStderrLogsRequest).Execute()

Update stderr log level

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
	searchUpdateStderrLogsRequest := *openapiclient.NewSearchUpdateStderrLogsRequest("Target_example") // SearchUpdateStderrLogsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.SearchUpdateStderrLogs(context.Background()).SearchUpdateStderrLogsRequest(searchUpdateStderrLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.SearchUpdateStderrLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateStderrLogs`: SearchGetStderrLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.SearchUpdateStderrLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateStderrLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchUpdateStderrLogsRequest** | [**SearchUpdateStderrLogsRequest**](SearchUpdateStderrLogsRequest.md) |  | 

### Return type

[**SearchGetStderrLogs200Response**](SearchGetStderrLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

