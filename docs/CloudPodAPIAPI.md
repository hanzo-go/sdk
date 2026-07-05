# \CloudPodAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddPod**](CloudPodAPIAPI.md#CloudApiControllerAddPod) | **Post** /v1/cloud/add-pod | Api Controller Add Pod
[**CloudApiControllerDeletePod**](CloudPodAPIAPI.md#CloudApiControllerDeletePod) | **Post** /v1/cloud/delete-pod | Api Controller Delete Pod
[**CloudApiControllerGetPod**](CloudPodAPIAPI.md#CloudApiControllerGetPod) | **Get** /v1/cloud/get-pod | Api Controller Get Pod
[**CloudApiControllerGetPods**](CloudPodAPIAPI.md#CloudApiControllerGetPods) | **Get** /v1/cloud/get-pods | Api Controller Get Pods
[**CloudApiControllerUpdatePod**](CloudPodAPIAPI.md#CloudApiControllerUpdatePod) | **Post** /v1/cloud/update-pod | Api Controller Update Pod



## CloudApiControllerAddPod

> CloudControllersResponse CloudApiControllerAddPod(ctx).CloudObjectPod(cloudObjectPod).Execute()

Api Controller Add Pod



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
	cloudObjectPod := *openapiclient.NewCloudObjectPod() // CloudObjectPod | The details of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPodAPIAPI.CloudApiControllerAddPod(context.Background()).CloudObjectPod(cloudObjectPod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPodAPIAPI.CloudApiControllerAddPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddPod`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudPodAPIAPI.CloudApiControllerAddPod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectPod** | [**CloudObjectPod**](CloudObjectPod.md) | The details of the pod | 

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


## CloudApiControllerDeletePod

> CloudControllersResponse CloudApiControllerDeletePod(ctx).CloudObjectPod(cloudObjectPod).Execute()

Api Controller Delete Pod



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
	cloudObjectPod := *openapiclient.NewCloudObjectPod() // CloudObjectPod | The details of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPodAPIAPI.CloudApiControllerDeletePod(context.Background()).CloudObjectPod(cloudObjectPod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPodAPIAPI.CloudApiControllerDeletePod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeletePod`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudPodAPIAPI.CloudApiControllerDeletePod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeletePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectPod** | [**CloudObjectPod**](CloudObjectPod.md) | The details of the pod | 

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


## CloudApiControllerGetPod

> CloudObjectPod CloudApiControllerGetPod(ctx).Id(id).Execute()

Api Controller Get Pod



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
	id := "id_example" // string | The id ( owner/name ) of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPodAPIAPI.CloudApiControllerGetPod(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPodAPIAPI.CloudApiControllerGetPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetPod`: CloudObjectPod
	fmt.Fprintf(os.Stdout, "Response from `CloudPodAPIAPI.CloudApiControllerGetPod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the pod | 

### Return type

[**CloudObjectPod**](CloudObjectPod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetPods

> CloudObjectPod CloudApiControllerGetPods(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Pods



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
	resp, r, err := apiClient.CloudPodAPIAPI.CloudApiControllerGetPods(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPodAPIAPI.CloudApiControllerGetPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetPods`: CloudObjectPod
	fmt.Fprintf(os.Stdout, "Response from `CloudPodAPIAPI.CloudApiControllerGetPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectPod**](CloudObjectPod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdatePod

> CloudControllersResponse CloudApiControllerUpdatePod(ctx).Id(id).CloudObjectPod(cloudObjectPod).Execute()

Api Controller Update Pod



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
	id := "id_example" // string | The id ( owner/name ) of the pod
	cloudObjectPod := *openapiclient.NewCloudObjectPod() // CloudObjectPod | The details of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPodAPIAPI.CloudApiControllerUpdatePod(context.Background()).Id(id).CloudObjectPod(cloudObjectPod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPodAPIAPI.CloudApiControllerUpdatePod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdatePod`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudPodAPIAPI.CloudApiControllerUpdatePod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdatePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the pod | 
 **cloudObjectPod** | [**CloudObjectPod**](CloudObjectPod.md) | The details of the pod | 

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

