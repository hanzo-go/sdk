# \PlanAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldLimits**](PlanAPI.md#WorldWorldLimits) | **Get** /v1/world/limits | Resolved World plan limits (contract echo)



## WorldWorldLimits

> WorldWorldLimits200Response WorldWorldLimits(ctx).Plan(plan).Execute()

Resolved World plan limits (contract echo)



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
	plan := "plan_example" // string | Plan id (world-free | world-pro | world-team | world-enterprise). Defaults to world-free. (optional) (default to "world-free")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.WorldWorldLimits(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.WorldWorldLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldLimits`: WorldWorldLimits200Response
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.WorldWorldLimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **string** | Plan id (world-free | world-pro | world-team | world-enterprise). Defaults to world-free. | [default to &quot;world-free&quot;]

### Return type

[**WorldWorldLimits200Response**](WorldWorldLimits200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

