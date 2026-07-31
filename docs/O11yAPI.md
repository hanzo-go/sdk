# \O11yAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1O11yAnnotationQueuesId**](O11yAPI.md#CloudDeleteV1O11yAnnotationQueuesId) | **Delete** /v1/o11y/annotation-queues/{id} | DeleteAnnotationQueue removes one review queue and every item in it.
[**CloudDeleteV1O11yByWildcard1**](O11yAPI.md#CloudDeleteV1O11yByWildcard1) | **Delete** /v1/o11y/{wildcard1} | 
[**CloudGetV1O11yAlertsLast**](O11yAPI.md#CloudGetV1O11yAlertsLast) | **Get** /v1/o11y/alerts/last | 
[**CloudGetV1O11yAnnotationQueues**](O11yAPI.md#CloudGetV1O11yAnnotationQueues) | **Get** /v1/o11y/annotation-queues | ListAnnotationQueues returns a page of the caller org&#39;s human-review queues, newest first, narrowed to the caller&#39;s project.
[**CloudGetV1O11yAnnotationQueuesId**](O11yAPI.md#CloudGetV1O11yAnnotationQueuesId) | **Get** /v1/o11y/annotation-queues/{id} | GetAnnotationQueue returns one review queue with its pending and completed counts and its first page of items.
[**CloudGetV1O11yAnnotationQueuesIdItems**](O11yAPI.md#CloudGetV1O11yAnnotationQueuesIdItems) | **Get** /v1/o11y/annotation-queues/{id}/items | ListAnnotationQueueItems returns a page of one review queue&#39;s items, newest first, optionally filtered to PENDING or COMPLETED.
[**CloudGetV1O11yApiV2Healthz**](O11yAPI.md#CloudGetV1O11yApiV2Healthz) | **Get** /v1/o11y/api/v2/healthz | 
[**CloudGetV1O11yApiV2Livez**](O11yAPI.md#CloudGetV1O11yApiV2Livez) | **Get** /v1/o11y/api/v2/livez | 
[**CloudGetV1O11yApiV2Readyz**](O11yAPI.md#CloudGetV1O11yApiV2Readyz) | **Get** /v1/o11y/api/v2/readyz | 
[**CloudGetV1O11yByWildcard1**](O11yAPI.md#CloudGetV1O11yByWildcard1) | **Get** /v1/o11y/{wildcard1} | 
[**CloudGetV1O11yLogs**](O11yAPI.md#CloudGetV1O11yLogs) | **Get** /v1/o11y/logs | GetO11yLogs returns a page of one product&#39;s logs for the caller&#39;s org.
[**CloudGetV1O11yMetrics**](O11yAPI.md#CloudGetV1O11yMetrics) | **Get** /v1/o11y/metrics | GetO11yMetrics returns one product&#39;s RED series — request rate, errors, p50 and p95 latency — for the caller&#39;s org, plus that org&#39;s LLM usage rollup over the same window.
[**CloudGetV1O11ySessions**](O11yAPI.md#CloudGetV1O11ySessions) | **Get** /v1/o11y/sessions | 
[**CloudGetV1O11yStatus**](O11yAPI.md#CloudGetV1O11yStatus) | **Get** /v1/o11y/status | GetO11yStatus reports whether a product&#39;s service is live: an in-cluster health probe with its measured latency, fused with the per-replica up inventory.
[**CloudGetV1O11yVmQuery**](O11yAPI.md#CloudGetV1O11yVmQuery) | **Get** /v1/o11y/vm/query | 
[**CloudGetV1O11yVmQueryRange**](O11yAPI.md#CloudGetV1O11yVmQueryRange) | **Get** /v1/o11y/vm/query_range | 
[**CloudOptionsV1O11yByWildcard1**](O11yAPI.md#CloudOptionsV1O11yByWildcard1) | **Options** /v1/o11y/{wildcard1} | 
[**CloudPatchV1O11yAnnotationQueuesId**](O11yAPI.md#CloudPatchV1O11yAnnotationQueuesId) | **Patch** /v1/o11y/annotation-queues/{id} | UpdateAnnotationQueue changes a review queue&#39;s name, description or score-config set.
[**CloudPatchV1O11yAnnotationQueuesIdItemsItemId**](O11yAPI.md#CloudPatchV1O11yAnnotationQueuesIdItemsItemId) | **Patch** /v1/o11y/annotation-queues/{id}/items/{itemId} | UpdateAnnotationQueueItem moves one queue item between PENDING and COMPLETED and sets its assignee.
[**CloudPatchV1O11yByWildcard1**](O11yAPI.md#CloudPatchV1O11yByWildcard1) | **Patch** /v1/o11y/{wildcard1} | 
[**CloudPostV1O11yAlertsByReceiver**](O11yAPI.md#CloudPostV1O11yAlertsByReceiver) | **Post** /v1/o11y/alerts/{receiver} | 
[**CloudPostV1O11yAnnotationQueues**](O11yAPI.md#CloudPostV1O11yAnnotationQueues) | **Post** /v1/o11y/annotation-queues | CreateAnnotationQueue creates a human-review queue in the caller&#39;s org and project.
[**CloudPostV1O11yAnnotationQueuesIdItems**](O11yAPI.md#CloudPostV1O11yAnnotationQueuesIdItems) | **Post** /v1/o11y/annotation-queues/{id}/items | AddAnnotationQueueItems enqueues traces, observations or sessions on a review queue.
[**CloudPostV1O11yByWildcard1**](O11yAPI.md#CloudPostV1O11yByWildcard1) | **Post** /v1/o11y/{wildcard1} | 
[**CloudPostV1O11yQuery**](O11yAPI.md#CloudPostV1O11yQuery) | **Post** /v1/o11y/query | 
[**CloudPostV1O11yQueryRange**](O11yAPI.md#CloudPostV1O11yQueryRange) | **Post** /v1/o11y/query_range | 
[**CloudPutV1O11yByWildcard1**](O11yAPI.md#CloudPutV1O11yByWildcard1) | **Put** /v1/o11y/{wildcard1} | 
[**CloudTraceV1O11yByWildcard1**](O11yAPI.md#CloudTraceV1O11yByWildcard1) | **Trace** /v1/o11y/{wildcard1} | 
[**O11yO11yDashboards**](O11yAPI.md#O11yO11yDashboards) | **Get** /v1/o11y/dashboards | List dashboards
[**O11yO11yHealth**](O11yAPI.md#O11yO11yHealth) | **Get** /v1/o11y/health | O11y health
[**O11yO11yIngest**](O11yAPI.md#O11yO11yIngest) | **Post** /v1/o11y/ingestion | Ingest LLM-observability events
[**O11yO11yRules**](O11yAPI.md#O11yO11yRules) | **Get** /v1/o11y/rules | List alert rules
[**O11yO11yServices**](O11yAPI.md#O11yO11yServices) | **Get** /v1/o11y/services | List traced services



## CloudDeleteV1O11yAnnotationQueuesId

> CloudAnnQueueDeleted CloudDeleteV1O11yAnnotationQueuesId(ctx, id).Execute()

DeleteAnnotationQueue removes one review queue and every item in it.



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
	id := "annq_1" // string | ID is the annotation queue to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudDeleteV1O11yAnnotationQueuesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudDeleteV1O11yAnnotationQueuesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1O11yAnnotationQueuesId`: CloudAnnQueueDeleted
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudDeleteV1O11yAnnotationQueuesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1O11yAnnotationQueuesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAnnQueueDeleted**](CloudAnnQueueDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1O11yByWildcard1

> CloudDeleteV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudDeleteV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudDeleteV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1O11yByWildcard1Request struct via the builder pattern


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


## CloudGetV1O11yAlertsLast

> CloudGetV1O11yAlertsLast(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11yAlertsLast(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yAlertsLast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yAlertsLastRequest struct via the builder pattern


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


## CloudGetV1O11yAnnotationQueues

> CloudAnnQueueList CloudGetV1O11yAnnotationQueues(ctx).Page(page).Limit(limit).Execute()

ListAnnotationQueues returns a page of the caller org's human-review queues, newest first, narrowed to the caller's project.



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
	page := int32(1) // int32 | Page is the 1-based page to read. Default 1. (optional)
	limit := int32(20) // int32 | Limit is how many rows to return. Default 20, capped at 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudGetV1O11yAnnotationQueues(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yAnnotationQueues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1O11yAnnotationQueues`: CloudAnnQueueList
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudGetV1O11yAnnotationQueues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yAnnotationQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page is the 1-based page to read. Default 1. | 
 **limit** | **int32** | Limit is how many rows to return. Default 20, capped at 100. | 

### Return type

[**CloudAnnQueueList**](CloudAnnQueueList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1O11yAnnotationQueuesId

> CloudAnnQueueDetailView CloudGetV1O11yAnnotationQueuesId(ctx, id).Execute()

GetAnnotationQueue returns one review queue with its pending and completed counts and its first page of items.



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
	id := "annq_1" // string | ID is the annotation queue to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudGetV1O11yAnnotationQueuesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yAnnotationQueuesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1O11yAnnotationQueuesId`: CloudAnnQueueDetailView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudGetV1O11yAnnotationQueuesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yAnnotationQueuesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAnnQueueDetailView**](CloudAnnQueueDetailView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1O11yAnnotationQueuesIdItems

> CloudAnnItemList CloudGetV1O11yAnnotationQueuesIdItems(ctx, id).Status(status).Page(page).Limit(limit).Execute()

ListAnnotationQueueItems returns a page of one review queue's items, newest first, optionally filtered to PENDING or COMPLETED.



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
	id := "annq_1" // string | ID is the annotation queue whose items to list, from the path.
	status := "PENDING" // string | Status filters to PENDING or COMPLETED items. Absent returns both. (optional)
	page := int32(56) // int32 | Page is the 1-based page to read. Default 1. (optional)
	limit := int32(56) // int32 | Limit is how many rows to return. Default 20, capped at 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudGetV1O11yAnnotationQueuesIdItems(context.Background(), id).Status(status).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yAnnotationQueuesIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1O11yAnnotationQueuesIdItems`: CloudAnnItemList
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudGetV1O11yAnnotationQueuesIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue whose items to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yAnnotationQueuesIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status filters to PENDING or COMPLETED items. Absent returns both. | 
 **page** | **int32** | Page is the 1-based page to read. Default 1. | 
 **limit** | **int32** | Limit is how many rows to return. Default 20, capped at 100. | 

### Return type

[**CloudAnnItemList**](CloudAnnItemList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1O11yApiV2Healthz

> CloudGetV1O11yApiV2Healthz(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11yApiV2Healthz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yApiV2Healthz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yApiV2HealthzRequest struct via the builder pattern


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


## CloudGetV1O11yApiV2Livez

> CloudGetV1O11yApiV2Livez(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11yApiV2Livez(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yApiV2Livez``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yApiV2LivezRequest struct via the builder pattern


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


## CloudGetV1O11yApiV2Readyz

> CloudGetV1O11yApiV2Readyz(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11yApiV2Readyz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yApiV2Readyz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yApiV2ReadyzRequest struct via the builder pattern


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


## CloudGetV1O11yByWildcard1

> CloudGetV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudGetV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yByWildcard1Request struct via the builder pattern


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


## CloudGetV1O11yLogs

> CloudLogsResponse CloudGetV1O11yLogs(ctx).Product(product).SinceNs(sinceNs).Window(window).Limit(limit).Execute()

GetO11yLogs returns a page of one product's logs for the caller's org.



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
	product := "kms" // string | Product is the console product slug whose logs to read, e.g. \"kms\". Required. (optional)
	sinceNs := int32(56) // int32 | SinceNs is the nanosecond cursor from a previous response's nextCursor. Absent (0) reads the last `window` seconds instead. (optional)
	window := int32(56) // int32 | Window is how many seconds back to read when there is no cursor. Default 900, capped at 86400. (optional)
	limit := int32(200) // int32 | Limit caps the returned lines. Default 200, capped at 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudGetV1O11yLogs(context.Background()).Product(product).SinceNs(sinceNs).Window(window).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1O11yLogs`: CloudLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudGetV1O11yLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Product is the console product slug whose logs to read, e.g. \&quot;kms\&quot;. Required. | 
 **sinceNs** | **int32** | SinceNs is the nanosecond cursor from a previous response&#39;s nextCursor. Absent (0) reads the last &#x60;window&#x60; seconds instead. | 
 **window** | **int32** | Window is how many seconds back to read when there is no cursor. Default 900, capped at 86400. | 
 **limit** | **int32** | Limit caps the returned lines. Default 200, capped at 1000. | 

### Return type

[**CloudLogsResponse**](CloudLogsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1O11yMetrics

> CloudMetricsResponse CloudGetV1O11yMetrics(ctx).Product(product).Range_(range_).StepSec(stepSec).Execute()

GetO11yMetrics returns one product's RED series — request rate, errors, p50 and p95 latency — for the caller's org, plus that org's LLM usage rollup over the same window.



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
	product := "kms" // string | Product is the console product slug to read, e.g. \"kms\". Required. (optional)
	range_ := int32(3600) // int32 | Range is the window in seconds. Default 3600, capped at 604800 (7d). (optional)
	stepSec := int32(56) // int32 | StepSec is the bucket width in seconds, clamped to [30, 3600]. Absent picks ~60 buckets across the range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudGetV1O11yMetrics(context.Background()).Product(product).Range_(range_).StepSec(stepSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1O11yMetrics`: CloudMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudGetV1O11yMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Product is the console product slug to read, e.g. \&quot;kms\&quot;. Required. | 
 **range_** | **int32** | Range is the window in seconds. Default 3600, capped at 604800 (7d). | 
 **stepSec** | **int32** | StepSec is the bucket width in seconds, clamped to [30, 3600]. Absent picks ~60 buckets across the range. | 

### Return type

[**CloudMetricsResponse**](CloudMetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1O11ySessions

> CloudGetV1O11ySessions(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11ySessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11ySessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11ySessionsRequest struct via the builder pattern


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


## CloudGetV1O11yStatus

> CloudStatusResult CloudGetV1O11yStatus(ctx).Product(product).Execute()

GetO11yStatus reports whether a product's service is live: an in-cluster health probe with its measured latency, fused with the per-replica up inventory.



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
	product := "kms" // string | Product is the console product slug to probe, e.g. \"kms\". Required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudGetV1O11yStatus(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1O11yStatus`: CloudStatusResult
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudGetV1O11yStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Product is the console product slug to probe, e.g. \&quot;kms\&quot;. Required. | 

### Return type

[**CloudStatusResult**](CloudStatusResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1O11yVmQuery

> CloudGetV1O11yVmQuery(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11yVmQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yVmQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yVmQueryRequest struct via the builder pattern


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


## CloudGetV1O11yVmQueryRange

> CloudGetV1O11yVmQueryRange(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudGetV1O11yVmQueryRange(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudGetV1O11yVmQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1O11yVmQueryRangeRequest struct via the builder pattern


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


## CloudOptionsV1O11yByWildcard1

> CloudOptionsV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudOptionsV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudOptionsV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudOptionsV1O11yByWildcard1Request struct via the builder pattern


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


## CloudPatchV1O11yAnnotationQueuesId

> CloudAnnQueueView CloudPatchV1O11yAnnotationQueuesId(ctx, id).CloudUpdateQueueIn(cloudUpdateQueueIn).Execute()

UpdateAnnotationQueue changes a review queue's name, description or score-config set.



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
	id := "annq_1" // string | ID is the annotation queue to update, from the path.
	cloudUpdateQueueIn := *openapiclient.NewCloudUpdateQueueIn() // CloudUpdateQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudPatchV1O11yAnnotationQueuesId(context.Background(), id).CloudUpdateQueueIn(cloudUpdateQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPatchV1O11yAnnotationQueuesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1O11yAnnotationQueuesId`: CloudAnnQueueView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudPatchV1O11yAnnotationQueuesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1O11yAnnotationQueuesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudUpdateQueueIn** | [**CloudUpdateQueueIn**](CloudUpdateQueueIn.md) |  | 

### Return type

[**CloudAnnQueueView**](CloudAnnQueueView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1O11yAnnotationQueuesIdItemsItemId

> CloudAnnItemView CloudPatchV1O11yAnnotationQueuesIdItemsItemId(ctx, id, itemId).CloudUpdateItemIn(cloudUpdateItemIn).Execute()

UpdateAnnotationQueueItem moves one queue item between PENDING and COMPLETED and sets its assignee.



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
	id := "annq_1" // string | ID is the annotation queue the item belongs to, from the path.
	itemId := "annqi_1" // string | ItemID is the item to update, from the path.
	cloudUpdateItemIn := *openapiclient.NewCloudUpdateItemIn() // CloudUpdateItemIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudPatchV1O11yAnnotationQueuesIdItemsItemId(context.Background(), id, itemId).CloudUpdateItemIn(cloudUpdateItemIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPatchV1O11yAnnotationQueuesIdItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1O11yAnnotationQueuesIdItemsItemId`: CloudAnnItemView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudPatchV1O11yAnnotationQueuesIdItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue the item belongs to, from the path. | 
**itemId** | **string** | ItemID is the item to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1O11yAnnotationQueuesIdItemsItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudUpdateItemIn** | [**CloudUpdateItemIn**](CloudUpdateItemIn.md) |  | 

### Return type

[**CloudAnnItemView**](CloudAnnItemView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1O11yByWildcard1

> CloudPatchV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudPatchV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPatchV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1O11yByWildcard1Request struct via the builder pattern


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


## CloudPostV1O11yAlertsByReceiver

> CloudPostV1O11yAlertsByReceiver(ctx, receiver).Execute()



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
	receiver := "receiver_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudPostV1O11yAlertsByReceiver(context.Background(), receiver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPostV1O11yAlertsByReceiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1O11yAlertsByReceiverRequest struct via the builder pattern


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


## CloudPostV1O11yAnnotationQueues

> CloudAnnQueueView CloudPostV1O11yAnnotationQueues(ctx).CloudCreateQueueReq(cloudCreateQueueReq).Execute()

CreateAnnotationQueue creates a human-review queue in the caller's org and project.



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
	cloudCreateQueueReq := *openapiclient.NewCloudCreateQueueReq() // CloudCreateQueueReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudPostV1O11yAnnotationQueues(context.Background()).CloudCreateQueueReq(cloudCreateQueueReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPostV1O11yAnnotationQueues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1O11yAnnotationQueues`: CloudAnnQueueView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudPostV1O11yAnnotationQueues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1O11yAnnotationQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateQueueReq** | [**CloudCreateQueueReq**](CloudCreateQueueReq.md) |  | 

### Return type

[**CloudAnnQueueView**](CloudAnnQueueView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1O11yAnnotationQueuesIdItems

> CloudAnnItemsCreated CloudPostV1O11yAnnotationQueuesIdItems(ctx, id).CloudAddItemsIn(cloudAddItemsIn).Execute()

AddAnnotationQueueItems enqueues traces, observations or sessions on a review queue.



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
	id := "annq_1" // string | ID is the annotation queue to add to, from the path.
	cloudAddItemsIn := *openapiclient.NewCloudAddItemsIn() // CloudAddItemsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloudPostV1O11yAnnotationQueuesIdItems(context.Background(), id).CloudAddItemsIn(cloudAddItemsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPostV1O11yAnnotationQueuesIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1O11yAnnotationQueuesIdItems`: CloudAnnItemsCreated
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloudPostV1O11yAnnotationQueuesIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to add to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1O11yAnnotationQueuesIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAddItemsIn** | [**CloudAddItemsIn**](CloudAddItemsIn.md) |  | 

### Return type

[**CloudAnnItemsCreated**](CloudAnnItemsCreated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1O11yByWildcard1

> CloudPostV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudPostV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPostV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1O11yByWildcard1Request struct via the builder pattern


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


## CloudPostV1O11yQuery

> CloudPostV1O11yQuery(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudPostV1O11yQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPostV1O11yQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1O11yQueryRequest struct via the builder pattern


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


## CloudPostV1O11yQueryRange

> CloudPostV1O11yQueryRange(ctx).Execute()



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
	r, err := apiClient.O11yAPI.CloudPostV1O11yQueryRange(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPostV1O11yQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1O11yQueryRangeRequest struct via the builder pattern


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


## CloudPutV1O11yByWildcard1

> CloudPutV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudPutV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudPutV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1O11yByWildcard1Request struct via the builder pattern


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


## CloudTraceV1O11yByWildcard1

> CloudTraceV1O11yByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CloudTraceV1O11yByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloudTraceV1O11yByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudTraceV1O11yByWildcard1Request struct via the builder pattern


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


## O11yO11yDashboards

> []O11yDashboardSummary O11yO11yDashboards(ctx).Execute()

List dashboards



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
	resp, r, err := apiClient.O11yAPI.O11yO11yDashboards(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yDashboards`: []O11yDashboardSummary
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yDashboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yDashboardsRequest struct via the builder pattern


### Return type

[**[]O11yDashboardSummary**](O11yDashboardSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yHealth

> O11yHealthResponse O11yO11yHealth(ctx).Execute()

O11y health



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
	resp, r, err := apiClient.O11yAPI.O11yO11yHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yHealth`: O11yHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yHealthRequest struct via the builder pattern


### Return type

[**O11yHealthResponse**](O11yHealthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yIngest

> O11yIngestResult O11yO11yIngest(ctx).O11yIngestBatch(o11yIngestBatch).Execute()

Ingest LLM-observability events



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
	o11yIngestBatch := *openapiclient.NewO11yIngestBatch([]openapiclient.O11yIngestEvent{*openapiclient.NewO11yIngestEvent("Type_example")}) // O11yIngestBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.O11yO11yIngest(context.Background()).O11yIngestBatch(o11yIngestBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yIngest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yIngest`: O11yIngestResult
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yIngest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yIngestBatch** | [**O11yIngestBatch**](O11yIngestBatch.md) |  | 

### Return type

[**O11yIngestResult**](O11yIngestResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yRules

> []O11yAlertRule O11yO11yRules(ctx).Execute()

List alert rules



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
	resp, r, err := apiClient.O11yAPI.O11yO11yRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yRules`: []O11yAlertRule
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yRulesRequest struct via the builder pattern


### Return type

[**[]O11yAlertRule**](O11yAlertRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yServices

> O11yO11yServices200Response O11yO11yServices(ctx).Execute()

List traced services



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
	resp, r, err := apiClient.O11yAPI.O11yO11yServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yServices`: O11yO11yServices200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yServicesRequest struct via the builder pattern


### Return type

[**O11yO11yServices200Response**](O11yO11yServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

