# \EntitlementsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntitlements**](EntitlementsAPI.md#GetEntitlements) | **Get** /v1/entitlements | Projection reports which console apps the CALLER&#39;s org may open, and the plan slug that decides it.



## GetEntitlements

> ProjectionView GetEntitlements(ctx).Execute()

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
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlements`: ProjectionView
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsRequest struct via the builder pattern


### Return type

[**ProjectionView**](ProjectionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

