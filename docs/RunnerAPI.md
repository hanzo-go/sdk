# \RunnerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRunnerReleases**](RunnerAPI.md#GetRunnerReleases) | **Get** /v1/runner/releases | Lists the self-publish releases this process has run.
[**GetRunnerReleasesById**](RunnerAPI.md#GetRunnerReleasesById) | **Get** /v1/runner/releases/{id} | Returns one self-publish release by the id its 202 returned.
[**PostRunner**](RunnerAPI.md#PostRunner) | **Post** /v1/runner | Triggers a native build — an image, or the binaries a repo declares.



## GetRunnerReleases

> SelfReleaseList GetRunnerReleases(ctx).Execute()

Lists the self-publish releases this process has run.



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
	resp, r, err := apiClient.RunnerAPI.GetRunnerReleases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerReleases`: SelfReleaseList
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerReleases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerReleasesRequest struct via the builder pattern


### Return type

[**SelfReleaseList**](SelfReleaseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerReleasesById

> ReleaseState GetRunnerReleasesById(ctx, id).Execute()

Returns one self-publish release by the id its 202 returned.



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
	id := "id_example" // string | ID is the build id the release trigger answered with, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerReleasesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerReleasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerReleasesById`: ReleaseState
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerReleasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the build id the release trigger answered with, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerReleasesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseState**](ReleaseState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunner

> RunnerBuildResp PostRunner(ctx).RunnerBuildReq(runnerBuildReq).Execute()

Triggers a native build — an image, or the binaries a repo declares.



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
	runnerBuildReq := *openapiclient.NewRunnerBuildReq() // RunnerBuildReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunner(context.Background()).RunnerBuildReq(runnerBuildReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunner`: RunnerBuildResp
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerBuildReq** | [**RunnerBuildReq**](RunnerBuildReq.md) |  | 

### Return type

[**RunnerBuildResp**](RunnerBuildResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

