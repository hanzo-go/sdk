# \SearchTasksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCancelTasks**](SearchTasksAPI.md#SearchCancelTasks) | **Post** /v1/search/tasks/cancel | Cancel enqueued or processing tasks
[**SearchDeleteTasks**](SearchTasksAPI.md#SearchDeleteTasks) | **Delete** /v1/search/tasks | Delete completed tasks
[**SearchGetTask**](SearchTasksAPI.md#SearchGetTask) | **Get** /v1/search/tasks/{taskUid} | Get task details
[**SearchListTasks**](SearchTasksAPI.md#SearchListTasks) | **Get** /v1/search/tasks | List all tasks



## SearchCancelTasks

> SearchSummarizedTaskView SearchCancelTasks(ctx).Uids(uids).Statuses(statuses).Types(types).IndexUids(indexUids).Execute()

Cancel enqueued or processing tasks

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
	uids := "uids_example" // string |  (optional)
	statuses := "statuses_example" // string |  (optional)
	types := "types_example" // string |  (optional)
	indexUids := "indexUids_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchTasksAPI.SearchCancelTasks(context.Background()).Uids(uids).Statuses(statuses).Types(types).IndexUids(indexUids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchTasksAPI.SearchCancelTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCancelTasks`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchTasksAPI.SearchCancelTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCancelTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uids** | **string** |  | 
 **statuses** | **string** |  | 
 **types** | **string** |  | 
 **indexUids** | **string** |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeleteTasks

> SearchSummarizedTaskView SearchDeleteTasks(ctx).Uids(uids).Statuses(statuses).Types(types).IndexUids(indexUids).Execute()

Delete completed tasks

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
	uids := "uids_example" // string |  (optional)
	statuses := "statuses_example" // string |  (optional)
	types := "types_example" // string |  (optional)
	indexUids := "indexUids_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchTasksAPI.SearchDeleteTasks(context.Background()).Uids(uids).Statuses(statuses).Types(types).IndexUids(indexUids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchTasksAPI.SearchDeleteTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteTasks`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchTasksAPI.SearchDeleteTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uids** | **string** |  | 
 **statuses** | **string** |  | 
 **types** | **string** |  | 
 **indexUids** | **string** |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGetTask

> SearchTaskView SearchGetTask(ctx, taskUid).Execute()

Get task details

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
	taskUid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchTasksAPI.SearchGetTask(context.Background(), taskUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchTasksAPI.SearchGetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetTask`: SearchTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchTasksAPI.SearchGetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskUid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchTaskView**](SearchTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListTasks

> SearchListTasks200Response SearchListTasks(ctx).Limit(limit).From(from).Uids(uids).Statuses(statuses).Types(types).IndexUids(indexUids).Execute()

List all tasks

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
	limit := int32(56) // int32 |  (optional) (default to 20)
	from := int32(56) // int32 | Task UID to start from (optional)
	uids := "uids_example" // string | Comma-separated task UIDs (optional)
	statuses := "statuses_example" // string | Comma-separated statuses (enqueued, processing, succeeded, failed, canceled) (optional)
	types := "types_example" // string | Comma-separated task types (optional)
	indexUids := "indexUids_example" // string | Comma-separated index UIDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchTasksAPI.SearchListTasks(context.Background()).Limit(limit).From(from).Uids(uids).Statuses(statuses).Types(types).IndexUids(indexUids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchTasksAPI.SearchListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListTasks`: SearchListTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchTasksAPI.SearchListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **from** | **int32** | Task UID to start from | 
 **uids** | **string** | Comma-separated task UIDs | 
 **statuses** | **string** | Comma-separated statuses (enqueued, processing, succeeded, failed, canceled) | 
 **types** | **string** | Comma-separated task types | 
 **indexUids** | **string** | Comma-separated index UIDs | 

### Return type

[**SearchListTasks200Response**](SearchListTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

