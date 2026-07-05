# \NexusPodAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddPod**](NexusPodAPIAPI.md#NexusAddPod) | **Post** /v1/nexus/add-pod | add Pod
[**NexusDeletePod**](NexusPodAPIAPI.md#NexusDeletePod) | **Post** /v1/nexus/delete-pod | delete Pod
[**NexusGetPod**](NexusPodAPIAPI.md#NexusGetPod) | **Get** /v1/nexus/get-pod | get Pod
[**NexusGetPods**](NexusPodAPIAPI.md#NexusGetPods) | **Get** /v1/nexus/get-pods | get Pods
[**NexusUpdatePod**](NexusPodAPIAPI.md#NexusUpdatePod) | **Post** /v1/nexus/update-pod | update Pod



## NexusAddPod

> NexusResponse NexusAddPod(ctx).NexusPod(nexusPod).Execute()

add Pod



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
	nexusPod := *openapiclient.NewNexusPod() // NexusPod | The details of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusPodAPIAPI.NexusAddPod(context.Background()).NexusPod(nexusPod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusPodAPIAPI.NexusAddPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddPod`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusPodAPIAPI.NexusAddPod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusPod** | [**NexusPod**](NexusPod.md) | The details of the pod | 

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


## NexusDeletePod

> NexusResponse NexusDeletePod(ctx).NexusPod(nexusPod).Execute()

delete Pod



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
	nexusPod := *openapiclient.NewNexusPod() // NexusPod | The details of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusPodAPIAPI.NexusDeletePod(context.Background()).NexusPod(nexusPod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusPodAPIAPI.NexusDeletePod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeletePod`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusPodAPIAPI.NexusDeletePod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeletePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusPod** | [**NexusPod**](NexusPod.md) | The details of the pod | 

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


## NexusGetPod

> NexusPod NexusGetPod(ctx).Id(id).Execute()

get Pod



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
	id := "id_example" // string | The id (owner/name) of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusPodAPIAPI.NexusGetPod(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusPodAPIAPI.NexusGetPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetPod`: NexusPod
	fmt.Fprintf(os.Stdout, "Response from `NexusPodAPIAPI.NexusGetPod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the pod | 

### Return type

[**NexusPod**](NexusPod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetPods

> NexusPod NexusGetPods(ctx).PageSize(pageSize).P(p).Execute()

get Pods



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
	resp, r, err := apiClient.NexusPodAPIAPI.NexusGetPods(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusPodAPIAPI.NexusGetPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetPods`: NexusPod
	fmt.Fprintf(os.Stdout, "Response from `NexusPodAPIAPI.NexusGetPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusPod**](NexusPod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdatePod

> NexusResponse NexusUpdatePod(ctx).Id(id).NexusPod(nexusPod).Execute()

update Pod



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
	id := "id_example" // string | The id (owner/name) of the pod
	nexusPod := *openapiclient.NewNexusPod() // NexusPod | The details of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusPodAPIAPI.NexusUpdatePod(context.Background()).Id(id).NexusPod(nexusPod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusPodAPIAPI.NexusUpdatePod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdatePod`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusPodAPIAPI.NexusUpdatePod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdatePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the pod | 
 **nexusPod** | [**NexusPod**](NexusPod.md) | The details of the pod | 

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

