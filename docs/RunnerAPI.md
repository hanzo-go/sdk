# \RunnerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRunner**](RunnerAPI.md#PostRunner) | **Post** /v1/runner | Triggers a native build — an image, or the binaries a repo declares.



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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

