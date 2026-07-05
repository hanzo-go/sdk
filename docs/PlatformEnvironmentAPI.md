# \PlatformEnvironmentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformEnvironmentByProjectId**](PlatformEnvironmentAPI.md#PlatformEnvironmentByProjectId) | **Get** /v1/platform/environment/byProjectId | List all environments for a project
[**PlatformEnvironmentCreate**](PlatformEnvironmentAPI.md#PlatformEnvironmentCreate) | **Post** /v1/platform/environment/create | Create a new environment within a project
[**PlatformEnvironmentDuplicate**](PlatformEnvironmentAPI.md#PlatformEnvironmentDuplicate) | **Post** /v1/platform/environment/duplicate | Duplicate an environment with all services
[**PlatformEnvironmentOne**](PlatformEnvironmentAPI.md#PlatformEnvironmentOne) | **Get** /v1/platform/environment/one | Get a single environment by ID
[**PlatformEnvironmentRemove**](PlatformEnvironmentAPI.md#PlatformEnvironmentRemove) | **Post** /v1/platform/environment/remove | Delete an environment
[**PlatformEnvironmentUpdate**](PlatformEnvironmentAPI.md#PlatformEnvironmentUpdate) | **Post** /v1/platform/environment/update | Update an environment



## PlatformEnvironmentByProjectId

> PlatformTRPCResult PlatformEnvironmentByProjectId(ctx).Input(input).Execute()

List all environments for a project

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
	resp, r, err := apiClient.PlatformEnvironmentAPI.PlatformEnvironmentByProjectId(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformEnvironmentAPI.PlatformEnvironmentByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformEnvironmentByProjectId`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformEnvironmentAPI.PlatformEnvironmentByProjectId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformEnvironmentByProjectIdRequest struct via the builder pattern


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


## PlatformEnvironmentCreate

> PlatformTRPCResult PlatformEnvironmentCreate(ctx).PlatformEnvironmentCreateRequest(platformEnvironmentCreateRequest).Execute()

Create a new environment within a project

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
	platformEnvironmentCreateRequest := *openapiclient.NewPlatformEnvironmentCreateRequest() // PlatformEnvironmentCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformEnvironmentAPI.PlatformEnvironmentCreate(context.Background()).PlatformEnvironmentCreateRequest(platformEnvironmentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformEnvironmentAPI.PlatformEnvironmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformEnvironmentCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformEnvironmentAPI.PlatformEnvironmentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformEnvironmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformEnvironmentCreateRequest** | [**PlatformEnvironmentCreateRequest**](PlatformEnvironmentCreateRequest.md) |  | 

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


## PlatformEnvironmentDuplicate

> PlatformTRPCResult PlatformEnvironmentDuplicate(ctx).PlatformEnvironmentDuplicateRequest(platformEnvironmentDuplicateRequest).Execute()

Duplicate an environment with all services

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
	platformEnvironmentDuplicateRequest := *openapiclient.NewPlatformEnvironmentDuplicateRequest() // PlatformEnvironmentDuplicateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformEnvironmentAPI.PlatformEnvironmentDuplicate(context.Background()).PlatformEnvironmentDuplicateRequest(platformEnvironmentDuplicateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformEnvironmentAPI.PlatformEnvironmentDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformEnvironmentDuplicate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformEnvironmentAPI.PlatformEnvironmentDuplicate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformEnvironmentDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformEnvironmentDuplicateRequest** | [**PlatformEnvironmentDuplicateRequest**](PlatformEnvironmentDuplicateRequest.md) |  | 

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


## PlatformEnvironmentOne

> PlatformTRPCResult PlatformEnvironmentOne(ctx).Input(input).Execute()

Get a single environment by ID

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
	resp, r, err := apiClient.PlatformEnvironmentAPI.PlatformEnvironmentOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformEnvironmentAPI.PlatformEnvironmentOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformEnvironmentOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformEnvironmentAPI.PlatformEnvironmentOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformEnvironmentOneRequest struct via the builder pattern


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


## PlatformEnvironmentRemove

> PlatformTRPCResult PlatformEnvironmentRemove(ctx).PlatformEnvironmentRemoveRequest(platformEnvironmentRemoveRequest).Execute()

Delete an environment

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
	platformEnvironmentRemoveRequest := *openapiclient.NewPlatformEnvironmentRemoveRequest() // PlatformEnvironmentRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformEnvironmentAPI.PlatformEnvironmentRemove(context.Background()).PlatformEnvironmentRemoveRequest(platformEnvironmentRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformEnvironmentAPI.PlatformEnvironmentRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformEnvironmentRemove`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformEnvironmentAPI.PlatformEnvironmentRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformEnvironmentRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformEnvironmentRemoveRequest** | [**PlatformEnvironmentRemoveRequest**](PlatformEnvironmentRemoveRequest.md) |  | 

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


## PlatformEnvironmentUpdate

> PlatformTRPCResult PlatformEnvironmentUpdate(ctx).PlatformEnvironmentUpdateRequest(platformEnvironmentUpdateRequest).Execute()

Update an environment

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
	platformEnvironmentUpdateRequest := *openapiclient.NewPlatformEnvironmentUpdateRequest() // PlatformEnvironmentUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformEnvironmentAPI.PlatformEnvironmentUpdate(context.Background()).PlatformEnvironmentUpdateRequest(platformEnvironmentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformEnvironmentAPI.PlatformEnvironmentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformEnvironmentUpdate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformEnvironmentAPI.PlatformEnvironmentUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformEnvironmentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformEnvironmentUpdateRequest** | [**PlatformEnvironmentUpdateRequest**](PlatformEnvironmentUpdateRequest.md) |  | 

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

