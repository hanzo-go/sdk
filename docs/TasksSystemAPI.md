# \TasksSystemAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksTasksCluster**](TasksSystemAPI.md#TasksTasksCluster) | **Get** /v1/tasks/cluster | Cluster status (open probe)
[**TasksTasksClusterHealth**](TasksSystemAPI.md#TasksTasksClusterHealth) | **Get** /v1/tasks/cluster/health | Cluster health (open probe)
[**TasksTasksHealth**](TasksSystemAPI.md#TasksTasksHealth) | **Get** /v1/tasks/health | Liveness probe
[**TasksTasksSettings**](TasksSystemAPI.md#TasksTasksSettings) | **Get** /v1/tasks/settings | Capability flags (open bootstrap)



## TasksTasksCluster

> map[string]interface{} TasksTasksCluster(ctx).Execute()

Cluster status (open probe)

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
	resp, r, err := apiClient.TasksSystemAPI.TasksTasksCluster(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksSystemAPI.TasksTasksCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksSystemAPI.TasksTasksCluster`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksClusterRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksClusterHealth

> map[string]interface{} TasksTasksClusterHealth(ctx).Execute()

Cluster health (open probe)

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
	resp, r, err := apiClient.TasksSystemAPI.TasksTasksClusterHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksSystemAPI.TasksTasksClusterHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksClusterHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksSystemAPI.TasksTasksClusterHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksClusterHealthRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksHealth

> V1EvalsHealthGet200Response TasksTasksHealth(ctx).Execute()

Liveness probe

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
	resp, r, err := apiClient.TasksSystemAPI.TasksTasksHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksSystemAPI.TasksTasksHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksHealth`: V1EvalsHealthGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TasksSystemAPI.TasksTasksHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksHealthRequest struct via the builder pattern


### Return type

[**V1EvalsHealthGet200Response**](V1EvalsHealthGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksSettings

> map[string]interface{} TasksTasksSettings(ctx).Execute()

Capability flags (open bootstrap)

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
	resp, r, err := apiClient.TasksSystemAPI.TasksTasksSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksSystemAPI.TasksTasksSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTasksSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksSystemAPI.TasksTasksSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksSettingsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

