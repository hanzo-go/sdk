# \CiAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCiFleet**](CiAPI.md#GetCiFleet) | **Get** /v1/ci/fleet | Compares what was written with what is running, one row per service along a single causal line: head, the commit on the branch; built, the image that commit produced; declared, the tag pinned in the universe repository; running, what the cluster serves.
[**GetCiRuns**](CiAPI.md#GetCiRuns) | **Get** /v1/ci/runs | Lists recent builds: the repo, the branch, the commit and how each run ended, newest first.
[**GetCiRunsByOwnerByRepo**](CiAPI.md#GetCiRunsByOwnerByRepo) | **Get** /v1/ci/runs/{owner}/{repo} | Lists what the forge ran for one ref of one repository: each run&#39;s workflow, status, conclusion and commit, with its jobs and their steps.



## GetCiFleet

> CiPipelines GetCiFleet(ctx).Execute()

Compares what was written with what is running, one row per service along a single causal line: head, the commit on the branch; built, the image that commit produced; declared, the tag pinned in the universe repository; running, what the cluster serves.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiAPI.GetCiFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiAPI.GetCiFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCiFleet`: CiPipelines
	fmt.Fprintf(os.Stdout, "Response from `CiAPI.GetCiFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCiFleetRequest struct via the builder pattern


### Return type

[**CiPipelines**](CiPipelines.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCiRuns

> CiExecutions GetCiRuns(ctx).Execute()

Lists recent builds: the repo, the branch, the commit and how each run ended, newest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiAPI.GetCiRuns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiAPI.GetCiRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCiRuns`: CiExecutions
	fmt.Fprintf(os.Stdout, "Response from `CiAPI.GetCiRuns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCiRunsRequest struct via the builder pattern


### Return type

[**CiExecutions**](CiExecutions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCiRunsByOwnerByRepo

> CiRuns GetCiRunsByOwnerByRepo(ctx, owner, repo).Ref(ref).Execute()

Lists what the forge ran for one ref of one repository: each run's workflow, status, conclusion and commit, with its jobs and their steps.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	owner := "hanzo-inc" // string | Owner is the forge namespace the repository lives under, from the path.
	repo := "cloud" // string | Repo is the repository's name, from the path.
	ref := "v1.2.3" // string | Ref is a branch, a tag or a commit sha. Empty means the default branch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiAPI.GetCiRunsByOwnerByRepo(context.Background(), owner, repo).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiAPI.GetCiRunsByOwnerByRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCiRunsByOwnerByRepo`: CiRuns
	fmt.Fprintf(os.Stdout, "Response from `CiAPI.GetCiRunsByOwnerByRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner is the forge namespace the repository lives under, from the path. | 
**repo** | **string** | Repo is the repository&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCiRunsByOwnerByRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **string** | Ref is a branch, a tag or a commit sha. Empty means the default branch. | 

### Return type

[**CiRuns**](CiRuns.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

