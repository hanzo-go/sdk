# \CloudTaskAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddTask**](CloudTaskAPIAPI.md#CloudApiControllerAddTask) | **Post** /v1/cloud/add-task | Api Controller Add Task
[**CloudApiControllerDeleteTask**](CloudTaskAPIAPI.md#CloudApiControllerDeleteTask) | **Post** /v1/cloud/delete-task | Api Controller Delete Task
[**CloudApiControllerGetGlobalTasks**](CloudTaskAPIAPI.md#CloudApiControllerGetGlobalTasks) | **Get** /v1/cloud/get-global-tasks | Api Controller Get Global Tasks
[**CloudApiControllerGetTask**](CloudTaskAPIAPI.md#CloudApiControllerGetTask) | **Get** /v1/cloud/get-task | Api Controller Get Task
[**CloudApiControllerGetTasks**](CloudTaskAPIAPI.md#CloudApiControllerGetTasks) | **Get** /v1/cloud/get-tasks | Api Controller Get Tasks
[**CloudApiControllerUpdateTask**](CloudTaskAPIAPI.md#CloudApiControllerUpdateTask) | **Post** /v1/cloud/update-task | Api Controller Update Task



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
	resp, r, err := apiClient.CloudTaskAPIAPI.CloudApiControllerAddTask(context.Background()).CloudObjectTask(cloudObjectTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTaskAPIAPI.CloudApiControllerAddTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddTask`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudTaskAPIAPI.CloudApiControllerAddTask`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTaskAPIAPI.CloudApiControllerDeleteTask(context.Background()).CloudObjectTask(cloudObjectTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTaskAPIAPI.CloudApiControllerDeleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteTask`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudTaskAPIAPI.CloudApiControllerDeleteTask`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTaskAPIAPI.CloudApiControllerGetGlobalTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTaskAPIAPI.CloudApiControllerGetGlobalTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalTasks`: []CloudObjectTask
	fmt.Fprintf(os.Stdout, "Response from `CloudTaskAPIAPI.CloudApiControllerGetGlobalTasks`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTaskAPIAPI.CloudApiControllerGetTask(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTaskAPIAPI.CloudApiControllerGetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTask`: CloudObjectTask
	fmt.Fprintf(os.Stdout, "Response from `CloudTaskAPIAPI.CloudApiControllerGetTask`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTaskAPIAPI.CloudApiControllerGetTasks(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTaskAPIAPI.CloudApiControllerGetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetTasks`: []CloudObjectTask
	fmt.Fprintf(os.Stdout, "Response from `CloudTaskAPIAPI.CloudApiControllerGetTasks`: %v\n", resp)
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
	resp, r, err := apiClient.CloudTaskAPIAPI.CloudApiControllerUpdateTask(context.Background()).Id(id).CloudObjectTask(cloudObjectTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudTaskAPIAPI.CloudApiControllerUpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateTask`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudTaskAPIAPI.CloudApiControllerUpdateTask`: %v\n", resp)
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

