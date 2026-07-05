# \PlatformMariadbAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformMariadbCreate**](PlatformMariadbAPI.md#PlatformMariadbCreate) | **Post** /v1/platform/mariadb/create | Create a MariaDB service
[**PlatformMariadbOne**](PlatformMariadbAPI.md#PlatformMariadbOne) | **Get** /v1/platform/mariadb/one | Get MariaDB details



## PlatformMariadbCreate

> PlatformTRPCResult PlatformMariadbCreate(ctx).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()

Create a MariaDB service

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
	resp, r, err := apiClient.PlatformMariadbAPI.PlatformMariadbCreate(context.Background()).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformMariadbAPI.PlatformMariadbCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformMariadbCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformMariadbAPI.PlatformMariadbCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformMariadbCreateRequest struct via the builder pattern


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


## PlatformMariadbOne

> PlatformTRPCResult PlatformMariadbOne(ctx).Input(input).Execute()

Get MariaDB details

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
	resp, r, err := apiClient.PlatformMariadbAPI.PlatformMariadbOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformMariadbAPI.PlatformMariadbOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformMariadbOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformMariadbAPI.PlatformMariadbOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformMariadbOneRequest struct via the builder pattern


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

