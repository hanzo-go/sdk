# \VectorAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddVector**](VectorAPIAPI.md#CloudApiControllerAddVector) | **Post** /v1/cloud/add-vector | Api Controller Add Vector
[**CloudApiControllerDeleteAllVectors**](VectorAPIAPI.md#CloudApiControllerDeleteAllVectors) | **Post** /v1/cloud/delete-all-vectors | Api Controller Delete All Vectors
[**CloudApiControllerDeleteVector**](VectorAPIAPI.md#CloudApiControllerDeleteVector) | **Post** /v1/cloud/delete-vector | Api Controller Delete Vector
[**CloudApiControllerGetGlobalVectors**](VectorAPIAPI.md#CloudApiControllerGetGlobalVectors) | **Get** /v1/cloud/get-global-vectors | Api Controller Get Global Vectors
[**CloudApiControllerGetVectors**](VectorAPIAPI.md#CloudApiControllerGetVectors) | **Get** /v1/cloud/get-vectors | Api Controller Get Vectors
[**CloudApiControllerUpdateVector**](VectorAPIAPI.md#CloudApiControllerUpdateVector) | **Post** /v1/cloud/update-vector | Api Controller Update Vector



## CloudApiControllerAddVector

> CloudControllersResponse CloudApiControllerAddVector(ctx).CloudObjectVector(cloudObjectVector).Execute()

Api Controller Add Vector



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
	cloudObjectVector := *openapiclient.NewCloudObjectVector() // CloudObjectVector | The details of the vector

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorAPIAPI.CloudApiControllerAddVector(context.Background()).CloudObjectVector(cloudObjectVector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPIAPI.CloudApiControllerAddVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddVector`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VectorAPIAPI.CloudApiControllerAddVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectVector** | [**CloudObjectVector**](CloudObjectVector.md) | The details of the vector | 

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


## CloudApiControllerDeleteAllVectors

> CloudControllersResponse CloudApiControllerDeleteAllVectors(ctx).Execute()

Api Controller Delete All Vectors



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
	resp, r, err := apiClient.VectorAPIAPI.CloudApiControllerDeleteAllVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPIAPI.CloudApiControllerDeleteAllVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteAllVectors`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VectorAPIAPI.CloudApiControllerDeleteAllVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteAllVectorsRequest struct via the builder pattern


### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteVector

> CloudControllersResponse CloudApiControllerDeleteVector(ctx).CloudObjectVector(cloudObjectVector).Execute()

Api Controller Delete Vector



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
	cloudObjectVector := *openapiclient.NewCloudObjectVector() // CloudObjectVector | The details of the vector

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorAPIAPI.CloudApiControllerDeleteVector(context.Background()).CloudObjectVector(cloudObjectVector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPIAPI.CloudApiControllerDeleteVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteVector`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VectorAPIAPI.CloudApiControllerDeleteVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectVector** | [**CloudObjectVector**](CloudObjectVector.md) | The details of the vector | 

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


## CloudApiControllerGetGlobalVectors

> []CloudObjectVector CloudApiControllerGetGlobalVectors(ctx).Execute()

Api Controller Get Global Vectors



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
	resp, r, err := apiClient.VectorAPIAPI.CloudApiControllerGetGlobalVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPIAPI.CloudApiControllerGetGlobalVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalVectors`: []CloudObjectVector
	fmt.Fprintf(os.Stdout, "Response from `VectorAPIAPI.CloudApiControllerGetGlobalVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalVectorsRequest struct via the builder pattern


### Return type

[**[]CloudObjectVector**](CloudObjectVector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetVectors

> []CloudObjectVector CloudApiControllerGetVectors(ctx).Execute()

Api Controller Get Vectors



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
	resp, r, err := apiClient.VectorAPIAPI.CloudApiControllerGetVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPIAPI.CloudApiControllerGetVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetVectors`: []CloudObjectVector
	fmt.Fprintf(os.Stdout, "Response from `VectorAPIAPI.CloudApiControllerGetVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetVectorsRequest struct via the builder pattern


### Return type

[**[]CloudObjectVector**](CloudObjectVector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateVector

> CloudControllersResponse CloudApiControllerUpdateVector(ctx).Id(id).CloudObjectVector(cloudObjectVector).Execute()

Api Controller Update Vector



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
	id := "id_example" // string | The id (owner/name) of the vector
	cloudObjectVector := *openapiclient.NewCloudObjectVector() // CloudObjectVector | The details of the vector

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorAPIAPI.CloudApiControllerUpdateVector(context.Background()).Id(id).CloudObjectVector(cloudObjectVector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPIAPI.CloudApiControllerUpdateVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateVector`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VectorAPIAPI.CloudApiControllerUpdateVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the vector | 
 **cloudObjectVector** | [**CloudObjectVector**](CloudObjectVector.md) | The details of the vector | 

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

