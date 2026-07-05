# \ConsoleAnnotationQueuesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateAnnotationQueue**](ConsoleAnnotationQueuesAPI.md#ConsoleCreateAnnotationQueue) | **Post** /v1/console/annotation-queues | Create an annotation queue
[**ConsoleGetAnnotationQueue**](ConsoleAnnotationQueuesAPI.md#ConsoleGetAnnotationQueue) | **Get** /v1/console/annotation-queues/{queueId} | Get an annotation queue
[**ConsoleListAnnotationQueueItems**](ConsoleAnnotationQueuesAPI.md#ConsoleListAnnotationQueueItems) | **Get** /v1/console/annotation-queues/{queueId}/items | Get items for an annotation queue
[**ConsoleListAnnotationQueues**](ConsoleAnnotationQueuesAPI.md#ConsoleListAnnotationQueues) | **Get** /v1/console/annotation-queues | Get all annotation queues



## ConsoleCreateAnnotationQueue

> ConsoleAnnotationQueue ConsoleCreateAnnotationQueue(ctx).ConsoleCreateAnnotationQueueRequest(consoleCreateAnnotationQueueRequest).Execute()

Create an annotation queue

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
	consoleCreateAnnotationQueueRequest := *openapiclient.NewConsoleCreateAnnotationQueueRequest("Name_example") // ConsoleCreateAnnotationQueueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleAnnotationQueuesAPI.ConsoleCreateAnnotationQueue(context.Background()).ConsoleCreateAnnotationQueueRequest(consoleCreateAnnotationQueueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleAnnotationQueuesAPI.ConsoleCreateAnnotationQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateAnnotationQueue`: ConsoleAnnotationQueue
	fmt.Fprintf(os.Stdout, "Response from `ConsoleAnnotationQueuesAPI.ConsoleCreateAnnotationQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateAnnotationQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateAnnotationQueueRequest** | [**ConsoleCreateAnnotationQueueRequest**](ConsoleCreateAnnotationQueueRequest.md) |  | 

### Return type

[**ConsoleAnnotationQueue**](ConsoleAnnotationQueue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetAnnotationQueue

> ConsoleAnnotationQueue ConsoleGetAnnotationQueue(ctx, queueId).Execute()

Get an annotation queue

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
	queueId := "queueId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleAnnotationQueuesAPI.ConsoleGetAnnotationQueue(context.Background(), queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleAnnotationQueuesAPI.ConsoleGetAnnotationQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetAnnotationQueue`: ConsoleAnnotationQueue
	fmt.Fprintf(os.Stdout, "Response from `ConsoleAnnotationQueuesAPI.ConsoleGetAnnotationQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetAnnotationQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleAnnotationQueue**](ConsoleAnnotationQueue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListAnnotationQueueItems

> ConsoleListAnnotationQueueItems200Response ConsoleListAnnotationQueueItems(ctx, queueId).Status(status).Page(page).Limit(limit).Execute()

Get items for an annotation queue

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
	queueId := "queueId_example" // string | 
	status := "status_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleAnnotationQueuesAPI.ConsoleListAnnotationQueueItems(context.Background(), queueId).Status(status).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleAnnotationQueuesAPI.ConsoleListAnnotationQueueItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListAnnotationQueueItems`: ConsoleListAnnotationQueueItems200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleAnnotationQueuesAPI.ConsoleListAnnotationQueueItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListAnnotationQueueItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ConsoleListAnnotationQueueItems200Response**](ConsoleListAnnotationQueueItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListAnnotationQueues

> ConsoleListAnnotationQueues200Response ConsoleListAnnotationQueues(ctx).Page(page).Limit(limit).Execute()

Get all annotation queues

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
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleAnnotationQueuesAPI.ConsoleListAnnotationQueues(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleAnnotationQueuesAPI.ConsoleListAnnotationQueues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListAnnotationQueues`: ConsoleListAnnotationQueues200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleAnnotationQueuesAPI.ConsoleListAnnotationQueues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListAnnotationQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ConsoleListAnnotationQueues200Response**](ConsoleListAnnotationQueues200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

