# \NexusTaskAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddTask**](NexusTaskAPIAPI.md#NexusAddTask) | **Post** /v1/nexus/add-task | add Task
[**NexusDeleteTask**](NexusTaskAPIAPI.md#NexusDeleteTask) | **Post** /v1/nexus/delete-task | delete Task
[**NexusGetGlobalTasks**](NexusTaskAPIAPI.md#NexusGetGlobalTasks) | **Get** /v1/nexus/get-global-tasks | get Global Tasks
[**NexusGetTask**](NexusTaskAPIAPI.md#NexusGetTask) | **Get** /v1/nexus/get-task | get Task
[**NexusGetTasks**](NexusTaskAPIAPI.md#NexusGetTasks) | **Get** /v1/nexus/get-tasks | get Tasks
[**NexusUpdateTask**](NexusTaskAPIAPI.md#NexusUpdateTask) | **Post** /v1/nexus/update-task | update Task



## NexusAddTask

> NexusResponse NexusAddTask(ctx).NexusTask(nexusTask).Execute()

add Task



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
	nexusTask := *openapiclient.NewNexusTask() // NexusTask | The details of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTaskAPIAPI.NexusAddTask(context.Background()).NexusTask(nexusTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTaskAPIAPI.NexusAddTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddTask`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusTaskAPIAPI.NexusAddTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusTask** | [**NexusTask**](NexusTask.md) | The details of the task | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteTask

> NexusResponse NexusDeleteTask(ctx).NexusTask(nexusTask).Execute()

delete Task



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
	nexusTask := *openapiclient.NewNexusTask() // NexusTask | The details of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTaskAPIAPI.NexusDeleteTask(context.Background()).NexusTask(nexusTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTaskAPIAPI.NexusDeleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteTask`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusTaskAPIAPI.NexusDeleteTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusTask** | [**NexusTask**](NexusTask.md) | The details of the task | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalTasks

> []NexusTask NexusGetGlobalTasks(ctx).Execute()

get Global Tasks



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
	resp, r, err := apiClient.NexusTaskAPIAPI.NexusGetGlobalTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTaskAPIAPI.NexusGetGlobalTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalTasks`: []NexusTask
	fmt.Fprintf(os.Stdout, "Response from `NexusTaskAPIAPI.NexusGetGlobalTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalTasksRequest struct via the builder pattern


### Return type

[**[]NexusTask**](NexusTask.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetTask

> NexusTask NexusGetTask(ctx).Id(id).Execute()

get Task



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
	id := "id_example" // string | The id (owner/name) of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTaskAPIAPI.NexusGetTask(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTaskAPIAPI.NexusGetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetTask`: NexusTask
	fmt.Fprintf(os.Stdout, "Response from `NexusTaskAPIAPI.NexusGetTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the task | 

### Return type

[**NexusTask**](NexusTask.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetTasks

> []NexusTask NexusGetTasks(ctx).Owner(owner).Execute()

get Tasks



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
	owner := "owner_example" // string | The owner of the tasks

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTaskAPIAPI.NexusGetTasks(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTaskAPIAPI.NexusGetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetTasks`: []NexusTask
	fmt.Fprintf(os.Stdout, "Response from `NexusTaskAPIAPI.NexusGetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the tasks | 

### Return type

[**[]NexusTask**](NexusTask.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateTask

> NexusResponse NexusUpdateTask(ctx).Id(id).NexusTask(nexusTask).Execute()

update Task



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
	id := "id_example" // string | The id (owner/name) of the task
	nexusTask := *openapiclient.NewNexusTask() // NexusTask | The details of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusTaskAPIAPI.NexusUpdateTask(context.Background()).Id(id).NexusTask(nexusTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusTaskAPIAPI.NexusUpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateTask`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusTaskAPIAPI.NexusUpdateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the task | 
 **nexusTask** | [**NexusTask**](NexusTask.md) | The details of the task | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

