# \ShareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetSharedWebsite**](ShareAPI.md#AnalyticsGetSharedWebsite) | **Get** /v1/analytics/share/{shareId} | Get a shared website by share ID (no auth required)



## AnalyticsGetSharedWebsite

> AnalyticsGetSharedWebsite200Response AnalyticsGetSharedWebsite(ctx, shareId).Execute()

Get a shared website by share ID (no auth required)

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
	shareId := "shareId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.AnalyticsGetSharedWebsite(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.AnalyticsGetSharedWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetSharedWebsite`: AnalyticsGetSharedWebsite200Response
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.AnalyticsGetSharedWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetSharedWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsGetSharedWebsite200Response**](AnalyticsGetSharedWebsite200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

