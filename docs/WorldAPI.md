# \WorldAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorld**](WorldAPI.md#GetWorld) | **Get** /v1/world | Answers GET /v1/world — the product&#39;s front door, naming every wire this surface answers on.
[**GetWorldLimits**](WorldAPI.md#GetWorldLimits) | **Get** /v1/world/limits | Echoes a World plan&#39;s rate limits, alert quota and model-API grant, read straight from the live @hanzo/plans catalog, so agents and dashboards configure themselves against the catalog instead of hardcoding tier numbers.
[**GetWorldNews**](WorldAPI.md#GetWorldNews) | **Get** /v1/world/news | Returns the caller&#39;s merged world-news feed: every source their project&#39;s pipeline names — GDELT once per keyword, plus each allowlisted RSS or Atom feed — fetched concurrently, narrowed by the pipeline&#39;s keyword/region/source filters, deduplicated by link and sorted freshest first, capped at 50 items.
[**GetWorldPipeline**](WorldAPI.md#GetWorldPipeline) | **Get** /v1/world/pipeline | Returns the caller project&#39;s news pipeline: which feeds it reads and how the merged result is filtered.
[**GetWorldStream**](WorldAPI.md#GetWorldStream) | **Get** /v1/world/stream | Live news refreshes for the caller&#39;s org and project, as Server-Sent Events.
[**PutWorldPipeline**](WorldAPI.md#PutWorldPipeline) | **Put** /v1/world/pipeline | Replaces the caller project&#39;s news pipeline and returns what was stored.



## GetWorld

> WorldIndex GetWorld(ctx).Execute()

Answers GET /v1/world — the product's front door, naming every wire this surface answers on.



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
	resp, r, err := apiClient.WorldAPI.GetWorld(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.GetWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorld`: WorldIndex
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.GetWorld`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldRequest struct via the builder pattern


### Return type

[**WorldIndex**](WorldIndex.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorldLimits

> LimitsView GetWorldLimits(ctx).Plan(plan).Execute()

Echoes a World plan's rate limits, alert quota and model-API grant, read straight from the live @hanzo/plans catalog, so agents and dashboards configure themselves against the catalog instead of hardcoding tier numbers.



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
	plan := "plan_example" // string | Plan is a World plan id from the live @hanzo/plans catalog, e.g. world-pro. Empty means world-free, and so does an id the catalog does not know — this never fails on an unknown plan. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldAPI.GetWorldLimits(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.GetWorldLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorldLimits`: LimitsView
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.GetWorldLimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **string** | Plan is a World plan id from the live @hanzo/plans catalog, e.g. world-pro. Empty means world-free, and so does an id the catalog does not know — this never fails on an unknown plan. | 

### Return type

[**LimitsView**](LimitsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorldNews

> NewsResponse GetWorldNews(ctx).Execute()

Returns the caller's merged world-news feed: every source their project's pipeline names — GDELT once per keyword, plus each allowlisted RSS or Atom feed — fetched concurrently, narrowed by the pipeline's keyword/region/source filters, deduplicated by link and sorted freshest first, capped at 50 items.



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
	resp, r, err := apiClient.WorldAPI.GetWorldNews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.GetWorldNews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorldNews`: NewsResponse
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.GetWorldNews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldNewsRequest struct via the builder pattern


### Return type

[**NewsResponse**](NewsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorldPipeline

> PipelineView GetWorldPipeline(ctx).Execute()

Returns the caller project's news pipeline: which feeds it reads and how the merged result is filtered.



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
	resp, r, err := apiClient.WorldAPI.GetWorldPipeline(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.GetWorldPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorldPipeline`: PipelineView
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.GetWorldPipeline`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPipelineRequest struct via the builder pattern


### Return type

[**PipelineView**](PipelineView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorldStream

> GetWorldStream(ctx).Execute()

Live news refreshes for the caller's org and project, as Server-Sent Events.



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
	r, err := apiClient.WorldAPI.GetWorldStream(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.GetWorldStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldStreamRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWorldPipeline

> PipelineView PutWorldPipeline(ctx).PipelineReq(pipelineReq).Execute()

Replaces the caller project's news pipeline and returns what was stored.



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
	pipelineReq := *openapiclient.NewPipelineReq() // PipelineReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldAPI.PutWorldPipeline(context.Background()).PipelineReq(pipelineReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.PutWorldPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWorldPipeline`: PipelineView
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.PutWorldPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutWorldPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineReq** | [**PipelineReq**](PipelineReq.md) |  | 

### Return type

[**PipelineView**](PipelineView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

