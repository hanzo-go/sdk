# \EventAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostEvent**](EventAPI.md#PostEvent) | **Post** /v1/event | Capture product events into your org&#39;s warehouse
[**PostEventByProjectEnvelope**](EventAPI.md#PostEventByProjectEnvelope) | **Post** /v1/event/{project}/envelope | Sentry SDK envelope ingest — errors and traces from an unmodified Sentry client
[**PostEventByProjectStore**](EventAPI.md#PostEventByProjectStore) | **Post** /v1/event/{project}/store | Sentry SDK store ingest — the legacy single-event wire



## PostEvent

> CaptureResult PostEvent(ctx).PostEventRequest(postEventRequest).Execute()

Capture product events into your org's warehouse



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
	postEventRequest := openapiclient.post_event_request{CaptureBatch: openapiclient.NewCaptureBatch()} // PostEventRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.PostEvent(context.Background()).PostEventRequest(postEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.PostEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEvent`: CaptureResult
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.PostEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postEventRequest** | [**PostEventRequest**](PostEventRequest.md) |  | 

### Return type

[**CaptureResult**](CaptureResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventByProjectEnvelope

> PostEventByProjectEnvelope(ctx, project).Body(body).Execute()

Sentry SDK envelope ingest — errors and traces from an unmodified Sentry client



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
	project := "project_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventAPI.PostEventByProjectEnvelope(context.Background(), project).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.PostEventByProjectEnvelope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEventByProjectEnvelopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventByProjectStore

> PostEventByProjectStore(ctx, project).Body(body).Execute()

Sentry SDK store ingest — the legacy single-event wire



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
	project := "project_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventAPI.PostEventByProjectStore(context.Background(), project).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.PostEventByProjectStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEventByProjectStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

