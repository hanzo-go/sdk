# \EntitlementsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Entitlements**](EntitlementsAPI.md#CloudGetV1Entitlements) | **Get** /v1/entitlements | Projection reports which console apps the CALLER&#39;s org may open, and the plan slug that decides it.



## CloudGetV1Entitlements

> CloudProjectionView CloudGetV1Entitlements(ctx).Execute()

Projection reports which console apps the CALLER's org may open, and the plan slug that decides it.



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
	resp, r, err := apiClient.EntitlementsAPI.CloudGetV1Entitlements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CloudGetV1Entitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Entitlements`: CloudProjectionView
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CloudGetV1Entitlements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EntitlementsRequest struct via the builder pattern


### Return type

[**CloudProjectionView**](CloudProjectionView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

