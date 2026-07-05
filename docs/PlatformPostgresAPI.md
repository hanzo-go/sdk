# \PlatformPostgresAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformPostgresCreate**](PlatformPostgresAPI.md#PlatformPostgresCreate) | **Post** /v1/platform/postgres/create | Create a PostgreSQL service
[**PlatformPostgresOne**](PlatformPostgresAPI.md#PlatformPostgresOne) | **Get** /v1/platform/postgres/one | Get PostgreSQL details



## PlatformPostgresCreate

> PlatformTRPCResult PlatformPostgresCreate(ctx).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()

Create a PostgreSQL service

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
	resp, r, err := apiClient.PlatformPostgresAPI.PlatformPostgresCreate(context.Background()).PlatformMariadbCreateRequest(platformMariadbCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformPostgresAPI.PlatformPostgresCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformPostgresCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformPostgresAPI.PlatformPostgresCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostgresCreateRequest struct via the builder pattern


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


## PlatformPostgresOne

> PlatformTRPCResult PlatformPostgresOne(ctx).Input(input).Execute()

Get PostgreSQL details

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
	resp, r, err := apiClient.PlatformPostgresAPI.PlatformPostgresOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformPostgresAPI.PlatformPostgresOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformPostgresOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformPostgresAPI.PlatformPostgresOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostgresOneRequest struct via the builder pattern


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

