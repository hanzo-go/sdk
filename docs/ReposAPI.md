# \ReposAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GitCreateRepo**](ReposAPI.md#GitCreateRepo) | **Post** /v1/git/repos | Create a bare repo
[**GitDeleteRepo**](ReposAPI.md#GitDeleteRepo) | **Delete** /v1/git/repos/{name} | Delete a repo and purge its storage
[**GitGetRepo**](ReposAPI.md#GitGetRepo) | **Get** /v1/git/repos/{name} | Repo detail (branches + resolved HEAD)
[**GitGitPush**](ReposAPI.md#GitGitPush) | **Post** /v1/git/repos/{name}/push | Client-less push (build a commit from posted files)
[**GitListRepos**](ReposAPI.md#GitListRepos) | **Get** /v1/git/repos | List the tenant&#39;s repos



## GitCreateRepo

> GitRepo GitCreateRepo(ctx).GitCreateRepo(gitCreateRepo).Execute()

Create a bare repo

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
	gitCreateRepo := *openapiclient.NewGitCreateRepo("Name_example") // GitCreateRepo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReposAPI.GitCreateRepo(context.Background()).GitCreateRepo(gitCreateRepo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReposAPI.GitCreateRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GitCreateRepo`: GitRepo
	fmt.Fprintf(os.Stdout, "Response from `ReposAPI.GitCreateRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGitCreateRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitCreateRepo** | [**GitCreateRepo**](GitCreateRepo.md) |  | 

### Return type

[**GitRepo**](GitRepo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitDeleteRepo

> GitDeleteRepo(ctx, name).Execute()

Delete a repo and purge its storage

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
	name := "name_example" // string | Repo name (a trailing \".git\" is stripped)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReposAPI.GitDeleteRepo(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReposAPI.GitDeleteRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Repo name (a trailing \&quot;.git\&quot; is stripped) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitDeleteRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitGetRepo

> GitRepo GitGetRepo(ctx, name).Execute()

Repo detail (branches + resolved HEAD)

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
	name := "name_example" // string | Repo name (a trailing \".git\" is stripped)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReposAPI.GitGetRepo(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReposAPI.GitGetRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GitGetRepo`: GitRepo
	fmt.Fprintf(os.Stdout, "Response from `ReposAPI.GitGetRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Repo name (a trailing \&quot;.git\&quot; is stripped) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitGetRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitRepo**](GitRepo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitGitPush

> GitPushResult GitGitPush(ctx, name).GitPushRequest(gitPushRequest).Execute()

Client-less push (build a commit from posted files)



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
	name := "name_example" // string | Repo name (a trailing \".git\" is stripped)
	gitPushRequest := *openapiclient.NewGitPushRequest([]openapiclient.GitPushFile{*openapiclient.NewGitPushFile("src/index.js", "Content_example")}) // GitPushRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReposAPI.GitGitPush(context.Background(), name).GitPushRequest(gitPushRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReposAPI.GitGitPush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GitGitPush`: GitPushResult
	fmt.Fprintf(os.Stdout, "Response from `ReposAPI.GitGitPush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Repo name (a trailing \&quot;.git\&quot; is stripped) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitGitPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitPushRequest** | [**GitPushRequest**](GitPushRequest.md) |  | 

### Return type

[**GitPushResult**](GitPushResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitListRepos

> GitListRepos200Response GitListRepos(ctx).Execute()

List the tenant's repos

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
	resp, r, err := apiClient.ReposAPI.GitListRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReposAPI.GitListRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GitListRepos`: GitListRepos200Response
	fmt.Fprintf(os.Stdout, "Response from `ReposAPI.GitListRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGitListReposRequest struct via the builder pattern


### Return type

[**GitListRepos200Response**](GitListRepos200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

