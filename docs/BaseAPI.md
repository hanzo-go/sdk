# \BaseAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBaseHealth**](BaseAPI.md#GetBaseHealth) | **Get** /v1/base/health | Reports that the base subsystem is serving.



## GetBaseHealth

> BaseHealth GetBaseHealth(ctx).Execute()

Reports that the base subsystem is serving.



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
	resp, r, err := apiClient.BaseAPI.GetBaseHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseAPI.GetBaseHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseHealth`: BaseHealth
	fmt.Fprintf(os.Stdout, "Response from `BaseAPI.GetBaseHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseHealthRequest struct via the builder pattern


### Return type

[**BaseHealth**](BaseHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

