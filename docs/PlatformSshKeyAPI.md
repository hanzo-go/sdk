# \PlatformSshKeyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformSshKeyAll**](PlatformSshKeyAPI.md#PlatformSshKeyAll) | **Get** /v1/platform/sshKey/all | List all SSH keys
[**PlatformSshKeyCreate**](PlatformSshKeyAPI.md#PlatformSshKeyCreate) | **Post** /v1/platform/sshKey/create | Add an SSH key
[**PlatformSshKeyGenerate**](PlatformSshKeyAPI.md#PlatformSshKeyGenerate) | **Post** /v1/platform/sshKey/generate | Generate a new SSH key pair
[**PlatformSshKeyRemove**](PlatformSshKeyAPI.md#PlatformSshKeyRemove) | **Post** /v1/platform/sshKey/remove | Remove an SSH key



## PlatformSshKeyAll

> PlatformTRPCResult PlatformSshKeyAll(ctx).Execute()

List all SSH keys

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
	resp, r, err := apiClient.PlatformSshKeyAPI.PlatformSshKeyAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSshKeyAPI.PlatformSshKeyAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSshKeyAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSshKeyAPI.PlatformSshKeyAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSshKeyAllRequest struct via the builder pattern


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


## PlatformSshKeyCreate

> PlatformTRPCResult PlatformSshKeyCreate(ctx).PlatformSshKeyCreateRequest(platformSshKeyCreateRequest).Execute()

Add an SSH key

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
	platformSshKeyCreateRequest := *openapiclient.NewPlatformSshKeyCreateRequest() // PlatformSshKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformSshKeyAPI.PlatformSshKeyCreate(context.Background()).PlatformSshKeyCreateRequest(platformSshKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSshKeyAPI.PlatformSshKeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSshKeyCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSshKeyAPI.PlatformSshKeyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSshKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformSshKeyCreateRequest** | [**PlatformSshKeyCreateRequest**](PlatformSshKeyCreateRequest.md) |  | 

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


## PlatformSshKeyGenerate

> PlatformTRPCResult PlatformSshKeyGenerate(ctx).Execute()

Generate a new SSH key pair

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
	resp, r, err := apiClient.PlatformSshKeyAPI.PlatformSshKeyGenerate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSshKeyAPI.PlatformSshKeyGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSshKeyGenerate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSshKeyAPI.PlatformSshKeyGenerate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSshKeyGenerateRequest struct via the builder pattern


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


## PlatformSshKeyRemove

> PlatformTRPCResult PlatformSshKeyRemove(ctx).PlatformSshKeyRemoveRequest(platformSshKeyRemoveRequest).Execute()

Remove an SSH key

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
	platformSshKeyRemoveRequest := *openapiclient.NewPlatformSshKeyRemoveRequest() // PlatformSshKeyRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformSshKeyAPI.PlatformSshKeyRemove(context.Background()).PlatformSshKeyRemoveRequest(platformSshKeyRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSshKeyAPI.PlatformSshKeyRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSshKeyRemove`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSshKeyAPI.PlatformSshKeyRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSshKeyRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformSshKeyRemoveRequest** | [**PlatformSshKeyRemoveRequest**](PlatformSshKeyRemoveRequest.md) |  | 

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

