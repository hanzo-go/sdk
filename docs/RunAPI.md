# \RunAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRun**](RunAPI.md#PostRun) | **Post** /v1/run | Runs a container image and gives back a URL.



## PostRun

> RunView PostRun(ctx).RunReq(runReq).Execute()

Runs a container image and gives back a URL.



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
	runReq := *openapiclient.NewRunReq() // RunReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunAPI.PostRun(context.Background()).RunReq(runReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunAPI.PostRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRun`: RunView
	fmt.Fprintf(os.Stdout, "Response from `RunAPI.PostRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runReq** | [**RunReq**](RunReq.md) |  | 

### Return type

[**RunView**](RunView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

