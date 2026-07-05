# \AutoGitReposAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoConnectGitRepo**](AutoGitReposAPI.md#AutoConnectGitRepo) | **Post** /v1/auto/git-repos | Connect a git repo for sync (EE)
[**AutoListGitRepos**](AutoGitReposAPI.md#AutoListGitRepos) | **Get** /v1/auto/git-repos | List connected git repos (EE)



## AutoConnectGitRepo

> map[string]interface{} AutoConnectGitRepo(ctx).AutoConnectGitRepoRequest(autoConnectGitRepoRequest).Execute()

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
	autoConnectGitRepoRequest := *openapiclient.NewAutoConnectGitRepoRequest("RemoteUrl_example") // AutoConnectGitRepoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoGitReposAPI.AutoConnectGitRepo(context.Background()).AutoConnectGitRepoRequest(autoConnectGitRepoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoGitReposAPI.AutoConnectGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoConnectGitRepo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoGitReposAPI.AutoConnectGitRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoConnectGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoConnectGitRepoRequest** | [**AutoConnectGitRepoRequest**](AutoConnectGitRepoRequest.md) |  | 

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


## AutoListGitRepos

> map[string]interface{} AutoListGitRepos(ctx).Execute()

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
	resp, r, err := apiClient.AutoGitReposAPI.AutoListGitRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoGitReposAPI.AutoListGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListGitRepos`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoGitReposAPI.AutoListGitRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListGitReposRequest struct via the builder pattern


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

