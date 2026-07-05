# \EngineGPUsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngineGetGPUAllocations**](EngineGPUsAPI.md#EngineGetGPUAllocations) | **Get** /v1/engine/gpus/allocations | Get GPU allocations
[**EngineListGPUTypes**](EngineGPUsAPI.md#EngineListGPUTypes) | **Get** /v1/engine/gpus | List available GPU types



## EngineGetGPUAllocations

> EngineGetGPUAllocations200Response EngineGetGPUAllocations(ctx).GpuType(gpuType).ClusterId(clusterId).Execute()

Get GPU allocations

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
	gpuType := "gpuType_example" // string |  (optional)
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineGPUsAPI.EngineGetGPUAllocations(context.Background()).GpuType(gpuType).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineGPUsAPI.EngineGetGPUAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetGPUAllocations`: EngineGetGPUAllocations200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineGPUsAPI.EngineGetGPUAllocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetGPUAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpuType** | **string** |  | 
 **clusterId** | **string** |  | 

### Return type

[**EngineGetGPUAllocations200Response**](EngineGetGPUAllocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListGPUTypes

> EngineListGPUTypes200Response EngineListGPUTypes(ctx).Execute()

List available GPU types

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
	resp, r, err := apiClient.EngineGPUsAPI.EngineListGPUTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineGPUsAPI.EngineListGPUTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListGPUTypes`: EngineListGPUTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineGPUsAPI.EngineListGPUTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEngineListGPUTypesRequest struct via the builder pattern


### Return type

[**EngineListGPUTypes200Response**](EngineListGPUTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

