# \PipelinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPipelines**](PipelinesAPI.md#GetPipelines) | **Get** /v1/pipelines | Returns one build-and-deploy pipeline per app, with its latest run.



## GetPipelines

> PipelineBoard GetPipelines(ctx).Execute()

Returns one build-and-deploy pipeline per app, with its latest run.



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
	resp, r, err := apiClient.PipelinesAPI.GetPipelines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelines`: PipelineBoard
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetPipelines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesRequest struct via the builder pattern


### Return type

[**PipelineBoard**](PipelineBoard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

