# \RunAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudPostV1Run**](RunAPI.md#CloudPostV1Run) | **Post** /v1/run | 



## CloudPostV1Run

> CloudRunView CloudPostV1Run(ctx).CloudRunReq(cloudRunReq).Execute()



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
	cloudRunReq := *openapiclient.NewCloudRunReq() // CloudRunReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunAPI.CloudPostV1Run(context.Background()).CloudRunReq(cloudRunReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunAPI.CloudPostV1Run``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Run`: CloudRunView
	fmt.Fprintf(os.Stdout, "Response from `RunAPI.CloudPostV1Run`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1RunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRunReq** | [**CloudRunReq**](CloudRunReq.md) |  | 

### Return type

[**CloudRunView**](CloudRunView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

