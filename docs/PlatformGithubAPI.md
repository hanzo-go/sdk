# \PlatformGithubAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformGithubGetGithubBranches**](PlatformGithubAPI.md#PlatformGithubGetGithubBranches) | **Get** /v1/platform/github/getGithubBranches | List branches for a repository
[**PlatformGithubGetGithubRepositories**](PlatformGithubAPI.md#PlatformGithubGetGithubRepositories) | **Get** /v1/platform/github/getGithubRepositories | List repos from a GitHub installation
[**PlatformGithubGithubProviders**](PlatformGithubAPI.md#PlatformGithubGithubProviders) | **Get** /v1/platform/github/githubProviders | List GitHub installations



## PlatformGithubGetGithubBranches

> PlatformTRPCResult PlatformGithubGetGithubBranches(ctx).Input(input).Execute()

List branches for a repository

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
	resp, r, err := apiClient.PlatformGithubAPI.PlatformGithubGetGithubBranches(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformGithubAPI.PlatformGithubGetGithubBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGithubGetGithubBranches`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformGithubAPI.PlatformGithubGetGithubBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGithubGetGithubBranchesRequest struct via the builder pattern


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


## PlatformGithubGetGithubRepositories

> PlatformTRPCResult PlatformGithubGetGithubRepositories(ctx).Input(input).Execute()

List repos from a GitHub installation

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
	resp, r, err := apiClient.PlatformGithubAPI.PlatformGithubGetGithubRepositories(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformGithubAPI.PlatformGithubGetGithubRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGithubGetGithubRepositories`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformGithubAPI.PlatformGithubGetGithubRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGithubGetGithubRepositoriesRequest struct via the builder pattern


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


## PlatformGithubGithubProviders

> PlatformTRPCResult PlatformGithubGithubProviders(ctx).Execute()

List GitHub installations

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
	resp, r, err := apiClient.PlatformGithubAPI.PlatformGithubGithubProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformGithubAPI.PlatformGithubGithubProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGithubGithubProviders`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformGithubAPI.PlatformGithubGithubProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGithubGithubProvidersRequest struct via the builder pattern


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

