# \BaseAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1BaseHealth**](BaseAPI.md#CloudGetV1BaseHealth) | **Get** /v1/base/health | BaseHealth reports that the base subsystem is serving.



## CloudGetV1BaseHealth

> CloudBaseHealth CloudGetV1BaseHealth(ctx).Execute()

BaseHealth reports that the base subsystem is serving.



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
	resp, r, err := apiClient.BaseAPI.CloudGetV1BaseHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.CloudGetV1BaseHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1BaseHealth`: CloudBaseHealth
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.CloudGetV1BaseHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BaseHealthRequest struct via the builder pattern


### Return type

[**CloudBaseHealth**](CloudBaseHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

