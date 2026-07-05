# \PlatformDockerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformDockerGetConfig**](PlatformDockerAPI.md#PlatformDockerGetConfig) | **Get** /v1/platform/docker/getConfig | Get Docker inspect for a container
[**PlatformDockerGetContainers**](PlatformDockerAPI.md#PlatformDockerGetContainers) | **Get** /v1/platform/docker/getContainers | List all Docker containers
[**PlatformDockerRestartContainer**](PlatformDockerAPI.md#PlatformDockerRestartContainer) | **Post** /v1/platform/docker/restartContainer | Restart a container by ID



## PlatformDockerGetConfig

> PlatformTRPCResult PlatformDockerGetConfig(ctx).Input(input).Execute()

Get Docker inspect for a container

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
	resp, r, err := apiClient.PlatformDockerAPI.PlatformDockerGetConfig(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDockerAPI.PlatformDockerGetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDockerGetConfig`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDockerAPI.PlatformDockerGetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDockerGetConfigRequest struct via the builder pattern


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


## PlatformDockerGetContainers

> PlatformTRPCResult PlatformDockerGetContainers(ctx).Input(input).Execute()

List all Docker containers

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
	resp, r, err := apiClient.PlatformDockerAPI.PlatformDockerGetContainers(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDockerAPI.PlatformDockerGetContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDockerGetContainers`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDockerAPI.PlatformDockerGetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDockerGetContainersRequest struct via the builder pattern


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


## PlatformDockerRestartContainer

> PlatformTRPCResult PlatformDockerRestartContainer(ctx).PlatformDockerRestartContainerRequest(platformDockerRestartContainerRequest).Execute()

Restart a container by ID

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
	platformDockerRestartContainerRequest := *openapiclient.NewPlatformDockerRestartContainerRequest() // PlatformDockerRestartContainerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDockerAPI.PlatformDockerRestartContainer(context.Background()).PlatformDockerRestartContainerRequest(platformDockerRestartContainerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDockerAPI.PlatformDockerRestartContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDockerRestartContainer`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDockerAPI.PlatformDockerRestartContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDockerRestartContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDockerRestartContainerRequest** | [**PlatformDockerRestartContainerRequest**](PlatformDockerRestartContainerRequest.md) |  | 

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

