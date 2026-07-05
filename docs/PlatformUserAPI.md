# \PlatformUserAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformUserAll**](PlatformUserAPI.md#PlatformUserAll) | **Get** /v1/platform/user/all | List all users (admin only)
[**PlatformUserCreateApiKey**](PlatformUserAPI.md#PlatformUserCreateApiKey) | **Post** /v1/platform/user/createApiKey | Create an API key
[**PlatformUserDeleteApiKey**](PlatformUserAPI.md#PlatformUserDeleteApiKey) | **Post** /v1/platform/user/deleteApiKey | Delete an API key
[**PlatformUserGet**](PlatformUserAPI.md#PlatformUserGet) | **Get** /v1/platform/user/get | Get current authenticated user



## PlatformUserAll

> PlatformTRPCResult PlatformUserAll(ctx).Execute()

List all users (admin only)

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
	resp, r, err := apiClient.PlatformUserAPI.PlatformUserAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformUserAPI.PlatformUserAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformUserAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformUserAPI.PlatformUserAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformUserAllRequest struct via the builder pattern


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


## PlatformUserCreateApiKey

> PlatformTRPCResult PlatformUserCreateApiKey(ctx).PlatformUserCreateApiKeyRequest(platformUserCreateApiKeyRequest).Execute()

Create an API key

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
	platformUserCreateApiKeyRequest := *openapiclient.NewPlatformUserCreateApiKeyRequest() // PlatformUserCreateApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformUserAPI.PlatformUserCreateApiKey(context.Background()).PlatformUserCreateApiKeyRequest(platformUserCreateApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformUserAPI.PlatformUserCreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformUserCreateApiKey`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformUserAPI.PlatformUserCreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformUserCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformUserCreateApiKeyRequest** | [**PlatformUserCreateApiKeyRequest**](PlatformUserCreateApiKeyRequest.md) |  | 

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


## PlatformUserDeleteApiKey

> PlatformTRPCResult PlatformUserDeleteApiKey(ctx).PlatformUserDeleteApiKeyRequest(platformUserDeleteApiKeyRequest).Execute()

Delete an API key

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
	platformUserDeleteApiKeyRequest := *openapiclient.NewPlatformUserDeleteApiKeyRequest() // PlatformUserDeleteApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformUserAPI.PlatformUserDeleteApiKey(context.Background()).PlatformUserDeleteApiKeyRequest(platformUserDeleteApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformUserAPI.PlatformUserDeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformUserDeleteApiKey`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformUserAPI.PlatformUserDeleteApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformUserDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformUserDeleteApiKeyRequest** | [**PlatformUserDeleteApiKeyRequest**](PlatformUserDeleteApiKeyRequest.md) |  | 

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


## PlatformUserGet

> PlatformTRPCResult PlatformUserGet(ctx).Execute()

Get current authenticated user

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
	resp, r, err := apiClient.PlatformUserAPI.PlatformUserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformUserAPI.PlatformUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformUserGet`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformUserAPI.PlatformUserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformUserGetRequest struct via the builder pattern


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

