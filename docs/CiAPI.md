# \CiAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCiFleet**](CiAPI.md#GetCiFleet) | **Get** /v1/ci/fleet | Compares what was written with what is running, one row per service along a single causal line: head, the commit on the branch; built, the image that commit produced; declared, the tag pinned in the universe repository; running, what the cluster serves.
[**GetCiRuns**](CiAPI.md#GetCiRuns) | **Get** /v1/ci/runs | Lists recent builds: the repo, the branch, the commit and how each run ended, newest first.



## GetCiFleet

> Pipelines GetCiFleet(ctx).Execute()

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
	// response from `GetCiFleet`: Pipelines
	fmt.Fprintf(os.Stdout, "Response from `CiAPI.GetCiFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCiFleetRequest struct via the builder pattern


### Return type

[**Pipelines**](Pipelines.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCiRuns

> Executions GetCiRuns(ctx).Execute()

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
	// response from `GetCiRuns`: Executions
	fmt.Fprintf(os.Stdout, "Response from `CiAPI.GetCiRuns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCiRunsRequest struct via the builder pattern


### Return type

[**Executions**](Executions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

