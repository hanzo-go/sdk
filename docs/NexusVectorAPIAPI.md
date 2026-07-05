# \NexusVectorAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddVector**](NexusVectorAPIAPI.md#NexusAddVector) | **Post** /v1/nexus/add-vector | add Vector
[**NexusDeleteAllVectors**](NexusVectorAPIAPI.md#NexusDeleteAllVectors) | **Post** /v1/nexus/delete-all-vectors | delete All Vectors
[**NexusDeleteVector**](NexusVectorAPIAPI.md#NexusDeleteVector) | **Post** /v1/nexus/delete-vector | delete Vector
[**NexusGetGlobalVectors**](NexusVectorAPIAPI.md#NexusGetGlobalVectors) | **Get** /v1/nexus/get-global-vectors | get Global Vectors
[**NexusGetVectors**](NexusVectorAPIAPI.md#NexusGetVectors) | **Get** /v1/nexus/get-vectors | get Vectors
[**NexusUpdateVector**](NexusVectorAPIAPI.md#NexusUpdateVector) | **Post** /v1/nexus/update-vector | update Vector



## NexusAddVector

> NexusResponse NexusAddVector(ctx).NexusVector(nexusVector).Execute()

add Vector



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
	nexusVector := *openapiclient.NewNexusVector() // NexusVector | The details of the vector

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVectorAPIAPI.NexusAddVector(context.Background()).NexusVector(nexusVector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVectorAPIAPI.NexusAddVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddVector`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVectorAPIAPI.NexusAddVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusVector** | [**NexusVector**](NexusVector.md) | The details of the vector | 

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


## NexusDeleteAllVectors

> NexusResponse NexusDeleteAllVectors(ctx).Execute()

delete All Vectors



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
	resp, r, err := apiClient.NexusVectorAPIAPI.NexusDeleteAllVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVectorAPIAPI.NexusDeleteAllVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteAllVectors`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVectorAPIAPI.NexusDeleteAllVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteAllVectorsRequest struct via the builder pattern


### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteVector

> NexusResponse NexusDeleteVector(ctx).NexusVector(nexusVector).Execute()

delete Vector



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
	nexusVector := *openapiclient.NewNexusVector() // NexusVector | The details of the vector

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVectorAPIAPI.NexusDeleteVector(context.Background()).NexusVector(nexusVector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVectorAPIAPI.NexusDeleteVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteVector`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVectorAPIAPI.NexusDeleteVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusVector** | [**NexusVector**](NexusVector.md) | The details of the vector | 

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


## NexusGetGlobalVectors

> []NexusVector NexusGetGlobalVectors(ctx).Execute()

get Global Vectors



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
	resp, r, err := apiClient.NexusVectorAPIAPI.NexusGetGlobalVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVectorAPIAPI.NexusGetGlobalVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalVectors`: []NexusVector
	fmt.Fprintf(os.Stdout, "Response from `NexusVectorAPIAPI.NexusGetGlobalVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalVectorsRequest struct via the builder pattern


### Return type

[**[]NexusVector**](NexusVector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetVectors

> []NexusVector NexusGetVectors(ctx).Execute()

get Vectors



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
	resp, r, err := apiClient.NexusVectorAPIAPI.NexusGetVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVectorAPIAPI.NexusGetVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetVectors`: []NexusVector
	fmt.Fprintf(os.Stdout, "Response from `NexusVectorAPIAPI.NexusGetVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetVectorsRequest struct via the builder pattern


### Return type

[**[]NexusVector**](NexusVector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateVector

> NexusResponse NexusUpdateVector(ctx).Id(id).NexusVector(nexusVector).Execute()

update Vector



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
	nexusVector := *openapiclient.NewNexusVector() // NexusVector | The details of the vector

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVectorAPIAPI.NexusUpdateVector(context.Background()).Id(id).NexusVector(nexusVector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVectorAPIAPI.NexusUpdateVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateVector`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVectorAPIAPI.NexusUpdateVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the vector | 
 **nexusVector** | [**NexusVector**](NexusVector.md) | The details of the vector | 

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

