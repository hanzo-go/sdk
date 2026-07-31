# \ContainerAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddContainer**](ContainerAPIAPI.md#CloudApiControllerAddContainer) | **Post** /v1/cloud/add-container | Api Controller Add Container
[**CloudApiControllerDeleteContainer**](ContainerAPIAPI.md#CloudApiControllerDeleteContainer) | **Post** /v1/cloud/delete-container | Api Controller Delete Container
[**CloudApiControllerGetContainer**](ContainerAPIAPI.md#CloudApiControllerGetContainer) | **Get** /v1/cloud/get-container | Api Controller Get Container
[**CloudApiControllerGetContainers**](ContainerAPIAPI.md#CloudApiControllerGetContainers) | **Get** /v1/cloud/get-containers | Api Controller Get Containers
[**CloudApiControllerUpdateContainer**](ContainerAPIAPI.md#CloudApiControllerUpdateContainer) | **Post** /v1/cloud/update-container | Api Controller Update Container



## CloudApiControllerAddContainer

> CloudControllersResponse CloudApiControllerAddContainer(ctx).CloudObjectContainer(cloudObjectContainer).Execute()

Api Controller Add Container



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
	cloudObjectContainer := *openapiclient.NewCloudObjectContainer() // CloudObjectContainer | The details of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPIAPI.CloudApiControllerAddContainer(context.Background()).CloudObjectContainer(cloudObjectContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPIAPI.CloudApiControllerAddContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddContainer`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPIAPI.CloudApiControllerAddContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectContainer** | [**CloudObjectContainer**](CloudObjectContainer.md) | The details of the container | 

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


## CloudApiControllerDeleteContainer

> CloudControllersResponse CloudApiControllerDeleteContainer(ctx).CloudObjectContainer(cloudObjectContainer).Execute()

Api Controller Delete Container



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
	cloudObjectContainer := *openapiclient.NewCloudObjectContainer() // CloudObjectContainer | The details of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPIAPI.CloudApiControllerDeleteContainer(context.Background()).CloudObjectContainer(cloudObjectContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPIAPI.CloudApiControllerDeleteContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteContainer`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPIAPI.CloudApiControllerDeleteContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectContainer** | [**CloudObjectContainer**](CloudObjectContainer.md) | The details of the container | 

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


## CloudApiControllerGetContainer

> CloudObjectContainer CloudApiControllerGetContainer(ctx).Id(id).Execute()

Api Controller Get Container



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
	id := "id_example" // string | The id ( owner/name ) of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPIAPI.CloudApiControllerGetContainer(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPIAPI.CloudApiControllerGetContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetContainer`: CloudObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPIAPI.CloudApiControllerGetContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the container | 

### Return type

[**CloudObjectContainer**](CloudObjectContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetContainers

> CloudObjectContainer CloudApiControllerGetContainers(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Containers



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
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPIAPI.CloudApiControllerGetContainers(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPIAPI.CloudApiControllerGetContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetContainers`: CloudObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPIAPI.CloudApiControllerGetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectContainer**](CloudObjectContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateContainer

> CloudControllersResponse CloudApiControllerUpdateContainer(ctx).Id(id).CloudObjectContainer(cloudObjectContainer).Execute()

Api Controller Update Container



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
	id := "id_example" // string | The id ( owner/name ) of the container
	cloudObjectContainer := *openapiclient.NewCloudObjectContainer() // CloudObjectContainer | The details of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPIAPI.CloudApiControllerUpdateContainer(context.Background()).Id(id).CloudObjectContainer(cloudObjectContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPIAPI.CloudApiControllerUpdateContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateContainer`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPIAPI.CloudApiControllerUpdateContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the container | 
 **cloudObjectContainer** | [**CloudObjectContainer**](CloudObjectContainer.md) | The details of the container | 

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

