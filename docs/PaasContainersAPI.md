# \PaasContainersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasCreateContainer**](PaasContainersAPI.md#PaasCreateContainer) | **Post** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers | Create container
[**PaasDeleteContainer**](PaasContainersAPI.md#PaasDeleteContainer) | **Delete** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers/{containerId} | Delete container
[**PaasDeployContainer**](PaasContainersAPI.md#PaasDeployContainer) | **Post** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers/{containerId}/deploy | Trigger deployment (Nixpacks build + deploy)
[**PaasGetContainer**](PaasContainersAPI.md#PaasGetContainer) | **Get** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers/{containerId} | Get container
[**PaasListContainerPods**](PaasContainersAPI.md#PaasListContainerPods) | **Get** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers/{containerId}/pods | List container pods
[**PaasListContainers**](PaasContainersAPI.md#PaasListContainers) | **Get** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers | List containers
[**PaasUpdateContainer**](PaasContainersAPI.md#PaasUpdateContainer) | **Put** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers/{containerId} | Update container



## PaasCreateContainer

> PaasContainer PaasCreateContainer(ctx, orgId, projectId, envId).PaasCreateContainerRequest(paasCreateContainerRequest).Execute()

Create container

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	paasCreateContainerRequest := *openapiclient.NewPaasCreateContainerRequest("Name_example", "Type_example") // PaasCreateContainerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasCreateContainer(context.Background(), orgId, projectId, envId).PaasCreateContainerRequest(paasCreateContainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasCreateContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasCreateContainer`: PaasContainer
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasCreateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasCreateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **paasCreateContainerRequest** | [**PaasCreateContainerRequest**](PaasCreateContainerRequest.md) |  | 

### Return type

[**PaasContainer**](PaasContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasDeleteContainer

> map[string]interface{} PaasDeleteContainer(ctx, orgId, projectId, envId, containerId).Execute()

Delete container

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	containerId := "containerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasDeleteContainer(context.Background(), orgId, projectId, envId, containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasDeleteContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDeleteContainer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasDeleteContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDeleteContainerRequest struct via the builder pattern


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


## PaasDeployContainer

> map[string]interface{} PaasDeployContainer(ctx, orgId, projectId, envId, containerId).Execute()

Trigger deployment (Nixpacks build + deploy)

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	containerId := "containerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasDeployContainer(context.Background(), orgId, projectId, envId, containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasDeployContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDeployContainer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasDeployContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDeployContainerRequest struct via the builder pattern


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


## PaasGetContainer

> PaasContainer PaasGetContainer(ctx, orgId, projectId, envId, containerId).Execute()

Get container

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	containerId := "containerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasGetContainer(context.Background(), orgId, projectId, envId, containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasGetContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetContainer`: PaasContainer
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasGetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PaasContainer**](PaasContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListContainerPods

> []PaasListContainerPods200ResponseInner PaasListContainerPods(ctx, orgId, projectId, envId, containerId).Execute()

List container pods

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	containerId := "containerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasListContainerPods(context.Background(), orgId, projectId, envId, containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasListContainerPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListContainerPods`: []PaasListContainerPods200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasListContainerPods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListContainerPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]PaasListContainerPods200ResponseInner**](PaasListContainerPods200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListContainers

> []PaasContainer PaasListContainers(ctx, orgId, projectId, envId).Execute()

List containers

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasListContainers(context.Background(), orgId, projectId, envId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasListContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListContainers`: []PaasContainer
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasListContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]PaasContainer**](PaasContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasUpdateContainer

> map[string]interface{} PaasUpdateContainer(ctx, orgId, projectId, envId, containerId).Body(body).Execute()

Update container

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	containerId := "containerId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasContainersAPI.PaasUpdateContainer(context.Background(), orgId, projectId, envId, containerId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasContainersAPI.PaasUpdateContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasUpdateContainer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasContainersAPI.PaasUpdateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasUpdateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

