# \PlatformProjectAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformProjectAll**](PlatformProjectAPI.md#PlatformProjectAll) | **Get** /v1/platform/project/all | List all projects in the active organization
[**PlatformProjectCreate**](PlatformProjectAPI.md#PlatformProjectCreate) | **Post** /v1/platform/project/create | Create a new project
[**PlatformProjectDuplicate**](PlatformProjectAPI.md#PlatformProjectDuplicate) | **Post** /v1/platform/project/duplicate | Duplicate a project or environment with selected services
[**PlatformProjectOne**](PlatformProjectAPI.md#PlatformProjectOne) | **Get** /v1/platform/project/one | Get a single project by ID
[**PlatformProjectRemove**](PlatformProjectAPI.md#PlatformProjectRemove) | **Post** /v1/platform/project/remove | Delete a project
[**PlatformProjectUpdate**](PlatformProjectAPI.md#PlatformProjectUpdate) | **Post** /v1/platform/project/update | Update a project



## PlatformProjectAll

> PlatformTRPCResult PlatformProjectAll(ctx).Execute()

List all projects in the active organization

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
	resp, r, err := apiClient.PlatformProjectAPI.PlatformProjectAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformProjectAPI.PlatformProjectAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformProjectAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformProjectAPI.PlatformProjectAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformProjectAllRequest struct via the builder pattern


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


## PlatformProjectCreate

> PlatformTRPCResult PlatformProjectCreate(ctx).PlatformProjectCreateRequest(platformProjectCreateRequest).Execute()

Create a new project

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
	platformProjectCreateRequest := *openapiclient.NewPlatformProjectCreateRequest() // PlatformProjectCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformProjectAPI.PlatformProjectCreate(context.Background()).PlatformProjectCreateRequest(platformProjectCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformProjectAPI.PlatformProjectCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformProjectCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformProjectAPI.PlatformProjectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformProjectCreateRequest** | [**PlatformProjectCreateRequest**](PlatformProjectCreateRequest.md) |  | 

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


## PlatformProjectDuplicate

> PlatformTRPCResult PlatformProjectDuplicate(ctx).PlatformProjectDuplicateRequest(platformProjectDuplicateRequest).Execute()

Duplicate a project or environment with selected services

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
	platformProjectDuplicateRequest := *openapiclient.NewPlatformProjectDuplicateRequest() // PlatformProjectDuplicateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformProjectAPI.PlatformProjectDuplicate(context.Background()).PlatformProjectDuplicateRequest(platformProjectDuplicateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformProjectAPI.PlatformProjectDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformProjectDuplicate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformProjectAPI.PlatformProjectDuplicate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformProjectDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformProjectDuplicateRequest** | [**PlatformProjectDuplicateRequest**](PlatformProjectDuplicateRequest.md) |  | 

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


## PlatformProjectOne

> PlatformTRPCResult PlatformProjectOne(ctx).Input(input).Execute()

Get a single project by ID

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
	resp, r, err := apiClient.PlatformProjectAPI.PlatformProjectOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformProjectAPI.PlatformProjectOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformProjectOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformProjectAPI.PlatformProjectOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformProjectOneRequest struct via the builder pattern


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


## PlatformProjectRemove

> PlatformTRPCResult PlatformProjectRemove(ctx).PlatformProjectRemoveRequest(platformProjectRemoveRequest).Execute()

Delete a project

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
	platformProjectRemoveRequest := *openapiclient.NewPlatformProjectRemoveRequest() // PlatformProjectRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformProjectAPI.PlatformProjectRemove(context.Background()).PlatformProjectRemoveRequest(platformProjectRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformProjectAPI.PlatformProjectRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformProjectRemove`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformProjectAPI.PlatformProjectRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformProjectRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformProjectRemoveRequest** | [**PlatformProjectRemoveRequest**](PlatformProjectRemoveRequest.md) |  | 

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


## PlatformProjectUpdate

> PlatformTRPCResult PlatformProjectUpdate(ctx).PlatformProjectUpdateRequest(platformProjectUpdateRequest).Execute()

Update a project

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
	platformProjectUpdateRequest := *openapiclient.NewPlatformProjectUpdateRequest() // PlatformProjectUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformProjectAPI.PlatformProjectUpdate(context.Background()).PlatformProjectUpdateRequest(platformProjectUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformProjectAPI.PlatformProjectUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformProjectUpdate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformProjectAPI.PlatformProjectUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformProjectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformProjectUpdateRequest** | [**PlatformProjectUpdateRequest**](PlatformProjectUpdateRequest.md) |  | 

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

