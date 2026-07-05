# \EngineRayAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngineCreateRayCluster**](EngineRayAPI.md#EngineCreateRayCluster) | **Post** /v1/engine/ray/clusters | Create Ray cluster
[**EngineDeleteRayCluster**](EngineRayAPI.md#EngineDeleteRayCluster) | **Delete** /v1/engine/ray/clusters/{name} | Delete Ray cluster
[**EngineGetRayCluster**](EngineRayAPI.md#EngineGetRayCluster) | **Get** /v1/engine/ray/clusters/{name} | Get Ray cluster
[**EngineGetRayClusterStatus**](EngineRayAPI.md#EngineGetRayClusterStatus) | **Get** /v1/engine/ray/clusters/{name}/status | Get Ray cluster status
[**EngineGetRayDashboard**](EngineRayAPI.md#EngineGetRayDashboard) | **Get** /v1/engine/ray/clusters/{name}/dashboard | Get Ray dashboard URL
[**EngineListRayClusters**](EngineRayAPI.md#EngineListRayClusters) | **Get** /v1/engine/ray/clusters | List Ray clusters
[**EngineScaleRayCluster**](EngineRayAPI.md#EngineScaleRayCluster) | **Put** /v1/engine/ray/clusters/{name}/scale | Scale Ray workers



## EngineCreateRayCluster

> EngineRayCluster EngineCreateRayCluster(ctx).EngineRayClusterCreate(engineRayClusterCreate).Execute()

Create Ray cluster

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
	engineRayClusterCreate := *openapiclient.NewEngineRayClusterCreate("Name_example") // EngineRayClusterCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineCreateRayCluster(context.Background()).EngineRayClusterCreate(engineRayClusterCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineCreateRayCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineCreateRayCluster`: EngineRayCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineCreateRayCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineCreateRayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineRayClusterCreate** | [**EngineRayClusterCreate**](EngineRayClusterCreate.md) |  | 

### Return type

[**EngineRayCluster**](EngineRayCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineDeleteRayCluster

> map[string]interface{} EngineDeleteRayCluster(ctx, name).Execute()

Delete Ray cluster

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineDeleteRayCluster(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineDeleteRayCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineDeleteRayCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineDeleteRayCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineDeleteRayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineGetRayCluster

> EngineRayCluster EngineGetRayCluster(ctx, name).Execute()

Get Ray cluster

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineGetRayCluster(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineGetRayCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetRayCluster`: EngineRayCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineGetRayCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetRayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineRayCluster**](EngineRayCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineGetRayClusterStatus

> EngineRayCluster EngineGetRayClusterStatus(ctx, name).Execute()

Get Ray cluster status

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineGetRayClusterStatus(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineGetRayClusterStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetRayClusterStatus`: EngineRayCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineGetRayClusterStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetRayClusterStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineRayCluster**](EngineRayCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineGetRayDashboard

> EngineGetRayDashboard200Response EngineGetRayDashboard(ctx, name).Execute()

Get Ray dashboard URL

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineGetRayDashboard(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineGetRayDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetRayDashboard`: EngineGetRayDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineGetRayDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetRayDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineGetRayDashboard200Response**](EngineGetRayDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListRayClusters

> EngineListRayClusters200Response EngineListRayClusters(ctx).Namespace(namespace).Status(status).Execute()

List Ray clusters

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
	namespace := "namespace_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineListRayClusters(context.Background()).Namespace(namespace).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineListRayClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListRayClusters`: EngineListRayClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineListRayClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineListRayClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**EngineListRayClusters200Response**](EngineListRayClusters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineScaleRayCluster

> EngineRayCluster EngineScaleRayCluster(ctx, name).EngineRayClusterScale(engineRayClusterScale).Execute()

Scale Ray workers

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
	name := "name_example" // string | 
	engineRayClusterScale := *openapiclient.NewEngineRayClusterScale([]openapiclient.EngineRayClusterScaleWorkersInner{*openapiclient.NewEngineRayClusterScaleWorkersInner("GroupName_example", int32(123))}) // EngineRayClusterScale | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineRayAPI.EngineScaleRayCluster(context.Background(), name).EngineRayClusterScale(engineRayClusterScale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineRayAPI.EngineScaleRayCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineScaleRayCluster`: EngineRayCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineRayAPI.EngineScaleRayCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineScaleRayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engineRayClusterScale** | [**EngineRayClusterScale**](EngineRayClusterScale.md) |  | 

### Return type

[**EngineRayCluster**](EngineRayCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

