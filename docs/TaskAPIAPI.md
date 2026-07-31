# \TaskAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddTask**](TaskAPIAPI.md#CloudApiControllerAddTask) | **Post** /v1/cloud/add-task | Api Controller Add Task
[**CloudApiControllerDeleteTask**](TaskAPIAPI.md#CloudApiControllerDeleteTask) | **Post** /v1/cloud/delete-task | Api Controller Delete Task
[**CloudApiControllerGetGlobalTasks**](TaskAPIAPI.md#CloudApiControllerGetGlobalTasks) | **Get** /v1/cloud/get-global-tasks | Api Controller Get Global Tasks
[**CloudApiControllerGetTask**](TaskAPIAPI.md#CloudApiControllerGetTask) | **Get** /v1/cloud/get-task | Api Controller Get Task
[**CloudApiControllerGetTasks**](TaskAPIAPI.md#CloudApiControllerGetTasks) | **Get** /v1/cloud/get-tasks | Api Controller Get Tasks
[**CloudApiControllerUpdateTask**](TaskAPIAPI.md#CloudApiControllerUpdateTask) | **Post** /v1/cloud/update-task | Api Controller Update Task



## CloudApiControllerAddTask

> CloudControllersResponse CloudApiControllerAddTask(ctx).CloudObjectTask(cloudObjectTask).Execute()

Api Controller Add Task



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
	cloudObjectTask := *openapiclient.NewCloudObjectTask() // CloudObjectTask | The details of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPIAPI.CloudApiControllerAddTask(context.Background()).CloudObjectTask(cloudObjectTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPIAPI.CloudApiControllerAddTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddTask`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TaskAPIAPI.CloudApiControllerAddTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectTask** | [**CloudObjectTask**](CloudObjectTask.md) | The details of the task | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteTask

> CloudControllersResponse CloudApiControllerDeleteTask(ctx).CloudObjectTask(cloudObjectTask).Execute()

Api Controller Delete Task



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
	cloudObjectTask := *openapiclient.NewCloudObjectTask() // CloudObjectTask | The details of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPIAPI.CloudApiControllerDeleteTask(context.Background()).CloudObjectTask(cloudObjectTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPIAPI.CloudApiControllerDeleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteTask`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TaskAPIAPI.CloudApiControllerDeleteTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectTask** | [**CloudObjectTask**](CloudObjectTask.md) | The details of the task | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalTasks

> []CloudObjectTask CloudApiControllerGetGlobalTasks(ctx).Execute()

Api Controller Get Global Tasks



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
	resp, r, err := apiClient.TaskAPIAPI.CloudApiControllerGetGlobalTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPIAPI.CloudApiControllerGetGlobalTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalTasks`: []CloudObjectTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPIAPI.CloudApiControllerGetGlobalTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalTasksRequest struct via the builder pattern


### Return type

[**[]CloudObjectTask**](CloudObjectTask.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetTask

> CloudObjectTask CloudApiControllerGetTask(ctx).Id(id).Execute()

Api Controller Get Task



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
	id := "id_example" // string | The id (owner/name) of task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPIAPI.CloudApiControllerGetTask(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPIAPI.CloudApiControllerGetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTask`: CloudObjectTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPIAPI.CloudApiControllerGetTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of task | 

### Return type

[**CloudObjectTask**](CloudObjectTask.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetTasks

> []CloudObjectTask CloudApiControllerGetTasks(ctx).Owner(owner).Execute()

Api Controller Get Tasks



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
	owner := "owner_example" // string | The owner of task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPIAPI.CloudApiControllerGetTasks(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPIAPI.CloudApiControllerGetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTasks`: []CloudObjectTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPIAPI.CloudApiControllerGetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of task | 

### Return type

[**[]CloudObjectTask**](CloudObjectTask.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateTask

> CloudControllersResponse CloudApiControllerUpdateTask(ctx).Id(id).CloudObjectTask(cloudObjectTask).Execute()

Api Controller Update Task



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
	cloudObjectTask := *openapiclient.NewCloudObjectTask() // CloudObjectTask | The details of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPIAPI.CloudApiControllerUpdateTask(context.Background()).Id(id).CloudObjectTask(cloudObjectTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPIAPI.CloudApiControllerUpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateTask`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `TaskAPIAPI.CloudApiControllerUpdateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the task | 
 **cloudObjectTask** | [**CloudObjectTask**](CloudObjectTask.md) | The details of the task | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

