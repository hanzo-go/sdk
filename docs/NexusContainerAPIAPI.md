# \NexusContainerAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddContainer**](NexusContainerAPIAPI.md#NexusAddContainer) | **Post** /v1/nexus/add-container | add Container
[**NexusDeleteContainer**](NexusContainerAPIAPI.md#NexusDeleteContainer) | **Post** /v1/nexus/delete-container | delete Container
[**NexusGetContainer**](NexusContainerAPIAPI.md#NexusGetContainer) | **Get** /v1/nexus/get-container | get Container
[**NexusGetContainers**](NexusContainerAPIAPI.md#NexusGetContainers) | **Get** /v1/nexus/get-containers | get Containers
[**NexusUpdateContainer**](NexusContainerAPIAPI.md#NexusUpdateContainer) | **Post** /v1/nexus/update-container | update Container



## NexusAddContainer

> NexusResponse NexusAddContainer(ctx).NexusContainer(nexusContainer).Execute()

add Container



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
	nexusContainer := *openapiclient.NewNexusContainer() // NexusContainer | The details of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusContainerAPIAPI.NexusAddContainer(context.Background()).NexusContainer(nexusContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusContainerAPIAPI.NexusAddContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddContainer`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusContainerAPIAPI.NexusAddContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusContainer** | [**NexusContainer**](NexusContainer.md) | The details of the container | 

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


## NexusDeleteContainer

> NexusResponse NexusDeleteContainer(ctx).NexusContainer(nexusContainer).Execute()

delete Container



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
	nexusContainer := *openapiclient.NewNexusContainer() // NexusContainer | The details of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusContainerAPIAPI.NexusDeleteContainer(context.Background()).NexusContainer(nexusContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusContainerAPIAPI.NexusDeleteContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteContainer`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusContainerAPIAPI.NexusDeleteContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusContainer** | [**NexusContainer**](NexusContainer.md) | The details of the container | 

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


## NexusGetContainer

> NexusContainer NexusGetContainer(ctx).Id(id).Execute()

get Container



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
	id := "id_example" // string | The id (owner/name) of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusContainerAPIAPI.NexusGetContainer(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusContainerAPIAPI.NexusGetContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetContainer`: NexusContainer
	fmt.Fprintf(os.Stdout, "Response from `NexusContainerAPIAPI.NexusGetContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the container | 

### Return type

[**NexusContainer**](NexusContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetContainers

> NexusContainer NexusGetContainers(ctx).PageSize(pageSize).P(p).Execute()

get Containers



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
	p := "p_example" // string | The page number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusContainerAPIAPI.NexusGetContainers(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusContainerAPIAPI.NexusGetContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetContainers`: NexusContainer
	fmt.Fprintf(os.Stdout, "Response from `NexusContainerAPIAPI.NexusGetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusContainer**](NexusContainer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateContainer

> NexusResponse NexusUpdateContainer(ctx).Id(id).NexusContainer(nexusContainer).Execute()

update Container



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
	id := "id_example" // string | The id (owner/name) of the container
	nexusContainer := *openapiclient.NewNexusContainer() // NexusContainer | The details of the container

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusContainerAPIAPI.NexusUpdateContainer(context.Background()).Id(id).NexusContainer(nexusContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusContainerAPIAPI.NexusUpdateContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateContainer`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusContainerAPIAPI.NexusUpdateContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the container | 
 **nexusContainer** | [**NexusContainer**](NexusContainer.md) | The details of the container | 

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

