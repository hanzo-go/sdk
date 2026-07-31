# \WorldAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1WorldLimits**](WorldAPI.md#CloudGetV1WorldLimits) | **Get** /v1/world/limits | Echoes a World plan&#39;s rate limits, alert quota and model-API grant, read straight from the live @hanzo/plans catalog, so agents and dashboards configure themselves against the catalog instead of hardcoding tier numbers.
[**CloudGetV1WorldNews**](WorldAPI.md#CloudGetV1WorldNews) | **Get** /v1/world/news | Returns the caller&#39;s merged world-news feed: every source their project&#39;s pipeline names — GDELT once per keyword, plus each allowlisted RSS or Atom feed — fetched concurrently, narrowed by the pipeline&#39;s keyword/region/source filters, deduplicated by link and sorted freshest first, capped at 50 items.
[**CloudGetV1WorldPipeline**](WorldAPI.md#CloudGetV1WorldPipeline) | **Get** /v1/world/pipeline | Returns the caller project&#39;s news pipeline: which feeds it reads and how the merged result is filtered.
[**CloudGetV1WorldStream**](WorldAPI.md#CloudGetV1WorldStream) | **Get** /v1/world/stream | Live news refreshes for the caller&#39;s org and project, as Server-Sent Events.
[**CloudPutV1WorldPipeline**](WorldAPI.md#CloudPutV1WorldPipeline) | **Put** /v1/world/pipeline | Replaces the caller project&#39;s news pipeline and returns what was stored.



## CloudGetV1WorldLimits

> CloudLimitsView CloudGetV1WorldLimits(ctx).Plan(plan).Execute()

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
	resp, r, err := apiClient.WorldAPI.CloudGetV1WorldLimits(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.CloudGetV1WorldLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WorldLimits`: CloudLimitsView
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.CloudGetV1WorldLimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WorldLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **string** | Plan is a World plan id from the live @hanzo/plans catalog, e.g. world-pro. Empty means world-free, and so does an id the catalog does not know — this never fails on an unknown plan. | 

### Return type

[**CloudLimitsView**](CloudLimitsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WorldNews

> CloudNewsResponse CloudGetV1WorldNews(ctx).Execute()

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
	resp, r, err := apiClient.WorldAPI.CloudGetV1WorldNews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.CloudGetV1WorldNews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WorldNews`: CloudNewsResponse
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.CloudGetV1WorldNews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WorldNewsRequest struct via the builder pattern


### Return type

[**CloudNewsResponse**](CloudNewsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WorldPipeline

> CloudPipelineView CloudGetV1WorldPipeline(ctx).Execute()

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
	resp, r, err := apiClient.WorldAPI.CloudGetV1WorldPipeline(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.CloudGetV1WorldPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WorldPipeline`: CloudPipelineView
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.CloudGetV1WorldPipeline`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WorldPipelineRequest struct via the builder pattern


### Return type

[**CloudPipelineView**](CloudPipelineView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WorldStream

> CloudGetV1WorldStream(ctx).Execute()

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
	r, err := apiClient.WorldAPI.CloudGetV1WorldStream(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.CloudGetV1WorldStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WorldStreamRequest struct via the builder pattern


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


## CloudPutV1WorldPipeline

> CloudPipelineView CloudPutV1WorldPipeline(ctx).CloudPipelineReq(cloudPipelineReq).Execute()

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
	cloudPipelineReq := *openapiclient.NewCloudPipelineReq() // CloudPipelineReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldAPI.CloudPutV1WorldPipeline(context.Background()).CloudPipelineReq(cloudPipelineReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldAPI.CloudPutV1WorldPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1WorldPipeline`: CloudPipelineView
	fmt.Fprintf(os.Stdout, "Response from `WorldAPI.CloudPutV1WorldPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1WorldPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPipelineReq** | [**CloudPipelineReq**](CloudPipelineReq.md) |  | 

### Return type

[**CloudPipelineView**](CloudPipelineView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

