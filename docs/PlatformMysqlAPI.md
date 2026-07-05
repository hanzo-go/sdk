# \PlatformMysqlAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformMysqlCreate**](PlatformMysqlAPI.md#PlatformMysqlCreate) | **Post** /v1/platform/mysql/create | Create a MySQL service
[**PlatformMysqlOne**](PlatformMysqlAPI.md#PlatformMysqlOne) | **Get** /v1/platform/mysql/one | Get MySQL details



## PlatformMysqlCreate

> PlatformTRPCResult PlatformMysqlCreate(ctx).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()

Create a MySQL service

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
	resp, r, err := apiClient.PlatformMysqlAPI.PlatformMysqlCreate(context.Background()).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformMysqlAPI.PlatformMysqlCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformMysqlCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformMysqlAPI.PlatformMysqlCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformMysqlCreateRequest struct via the builder pattern


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


## PlatformMysqlOne

> PlatformTRPCResult PlatformMysqlOne(ctx).Input(input).Execute()

Get MySQL details

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
	resp, r, err := apiClient.PlatformMysqlAPI.PlatformMysqlOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformMysqlAPI.PlatformMysqlOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformMysqlOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformMysqlAPI.PlatformMysqlOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformMysqlOneRequest struct via the builder pattern


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

