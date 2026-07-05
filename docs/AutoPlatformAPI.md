# \AutoPlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoGetPlatform**](AutoPlatformAPI.md#AutoGetPlatform) | **Get** /v1/auto/platforms/{id} | Get platform settings
[**AutoUpdatePlatform**](AutoPlatformAPI.md#AutoUpdatePlatform) | **Post** /v1/auto/platforms/{id} | Update platform settings



## AutoGetPlatform

> map[string]interface{} AutoGetPlatform(ctx, id).Execute()

Get platform settings

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPlatformAPI.AutoGetPlatform(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPlatformAPI.AutoGetPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetPlatform`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoPlatformAPI.AutoGetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetPlatformRequest struct via the builder pattern


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


## AutoUpdatePlatform

> map[string]interface{} AutoUpdatePlatform(ctx, id).AutoUpdatePlatformRequest(autoUpdatePlatformRequest).Execute()

Update platform settings

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
	id := "id_example" // string | 
	autoUpdatePlatformRequest := *openapiclient.NewAutoUpdatePlatformRequest() // AutoUpdatePlatformRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPlatformAPI.AutoUpdatePlatform(context.Background(), id).AutoUpdatePlatformRequest(autoUpdatePlatformRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPlatformAPI.AutoUpdatePlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUpdatePlatform`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoPlatformAPI.AutoUpdatePlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoUpdatePlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUpdatePlatformRequest** | [**AutoUpdatePlatformRequest**](AutoUpdatePlatformRequest.md) |  | 

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

