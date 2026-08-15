# \EnvironmentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironments**](EnvironmentsAPI.md#GetEnvironments) | **Get** /v1/environments | Returns your deploy targets, and what is running on each.



## GetEnvironments

> EnvironmentBoard GetEnvironments(ctx).Execute()

Returns your deploy targets, and what is running on each.



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
	resp, r, err := apiClient.EnvironmentsAPI.GetEnvironments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironments`: EnvironmentBoard
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetEnvironments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentsRequest struct via the builder pattern


### Return type

[**EnvironmentBoard**](EnvironmentBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

