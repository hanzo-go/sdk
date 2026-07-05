# \PlatformMongoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformMongoCreate**](PlatformMongoAPI.md#PlatformMongoCreate) | **Post** /v1/platform/mongo/create | Create a MongoDB service
[**PlatformMongoOne**](PlatformMongoAPI.md#PlatformMongoOne) | **Get** /v1/platform/mongo/one | Get MongoDB details



## PlatformMongoCreate

> PlatformTRPCResult PlatformMongoCreate(ctx).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()

Create a MongoDB service

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
	resp, r, err := apiClient.PlatformMongoAPI.PlatformMongoCreate(context.Background()).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformMongoAPI.PlatformMongoCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformMongoCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformMongoAPI.PlatformMongoCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformMongoCreateRequest struct via the builder pattern


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


## PlatformMongoOne

> PlatformTRPCResult PlatformMongoOne(ctx).Input(input).Execute()

Get MongoDB details

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
	resp, r, err := apiClient.PlatformMongoAPI.PlatformMongoOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformMongoAPI.PlatformMongoOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformMongoOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformMongoAPI.PlatformMongoOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformMongoOneRequest struct via the builder pattern


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

