# \PlatformRedisAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformRedisCreate**](PlatformRedisAPI.md#PlatformRedisCreate) | **Post** /v1/platform/redis/create | Create a Redis service
[**PlatformRedisOne**](PlatformRedisAPI.md#PlatformRedisOne) | **Get** /v1/platform/redis/one | Get Redis details



## PlatformRedisCreate

> PlatformTRPCResult PlatformRedisCreate(ctx).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()

Create a Redis service

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
	platformMariadbCreateRequest := *openapiclient.NewPlatformMariadbCreateRequest() // PlatformMariadbCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformRedisAPI.PlatformRedisCreate(context.Background()).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformRedisAPI.PlatformRedisCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformRedisCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformRedisAPI.PlatformRedisCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformRedisCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformMariadbCreateRequest** | [**PlatformMariadbCreateRequest**](PlatformMariadbCreateRequest.md) |  | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformRedisOne

> PlatformTRPCResult PlatformRedisOne(ctx).Input(input).Execute()

Get Redis details

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
	input := "input_example" // string | URL-encoded JSON input for tRPC queries (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformRedisAPI.PlatformRedisOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformRedisAPI.PlatformRedisOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformRedisOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformRedisAPI.PlatformRedisOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformRedisOneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

