# \FlowGitReposAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowConnectGitRepo**](FlowGitReposAPI.md#FlowConnectGitRepo) | **Post** /v1/flow/git-repos | Connect a git repo for sync (EE)
[**FlowDisconnectGitRepo**](FlowGitReposAPI.md#FlowDisconnectGitRepo) | **Delete** /v1/flow/git-repos/{id} | Disconnect a git repo (EE)
[**FlowListGitRepos**](FlowGitReposAPI.md#FlowListGitRepos) | **Get** /v1/flow/git-repos | List connected git repos (EE)
[**FlowPullFromGitRepo**](FlowGitReposAPI.md#FlowPullFromGitRepo) | **Post** /v1/flow/git-repos/{id}/pull | Pull flows from git (EE)
[**FlowPushToGitRepo**](FlowGitReposAPI.md#FlowPushToGitRepo) | **Post** /v1/flow/git-repos/{id}/push | Push flows to git (EE)



## FlowConnectGitRepo

> map[string]interface{} FlowConnectGitRepo(ctx).FlowConnectGitRepoRequest(flowConnectGitRepoRequest).Execute()

Connect a git repo for sync (EE)

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
	flowConnectGitRepoRequest := *openapiclient.NewFlowConnectGitRepoRequest("RemoteUrl_example") // FlowConnectGitRepoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowGitReposAPI.FlowConnectGitRepo(context.Background()).FlowConnectGitRepoRequest(flowConnectGitRepoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGitReposAPI.FlowConnectGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowConnectGitRepo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGitReposAPI.FlowConnectGitRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowConnectGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowConnectGitRepoRequest** | [**FlowConnectGitRepoRequest**](FlowConnectGitRepoRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowDisconnectGitRepo

> FlowDisconnectGitRepo(ctx, id).Execute()

Disconnect a git repo (EE)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowGitReposAPI.FlowDisconnectGitRepo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGitReposAPI.FlowDisconnectGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowDisconnectGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListGitRepos

> map[string]interface{} FlowListGitRepos(ctx).Execute()

List connected git repos (EE)

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
	resp, r, err := apiClient.FlowGitReposAPI.FlowListGitRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGitReposAPI.FlowListGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListGitRepos`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGitReposAPI.FlowListGitRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListGitReposRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowPullFromGitRepo

> map[string]interface{} FlowPullFromGitRepo(ctx, id).Execute()

Pull flows from git (EE)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowGitReposAPI.FlowPullFromGitRepo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGitReposAPI.FlowPullFromGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowPullFromGitRepo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGitReposAPI.FlowPullFromGitRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowPullFromGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowPushToGitRepo

> map[string]interface{} FlowPushToGitRepo(ctx, id).Execute()

Push flows to git (EE)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowGitReposAPI.FlowPushToGitRepo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGitReposAPI.FlowPushToGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowPushToGitRepo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGitReposAPI.FlowPushToGitRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowPushToGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

