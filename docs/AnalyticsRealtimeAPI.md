# \AnalyticsRealtimeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetRealtimeData**](AnalyticsRealtimeAPI.md#AnalyticsGetRealtimeData) | **Get** /v1/analytics/realtime/{websiteId} | Get realtime visitor data for the last 30 minutes



## AnalyticsGetRealtimeData

> map[string]interface{} AnalyticsGetRealtimeData(ctx, websiteId).Timezone(timezone).Execute()

Get realtime visitor data for the last 30 minutes

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	timezone := "America/Los_Angeles" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsRealtimeAPI.AnalyticsGetRealtimeData(context.Background(), websiteId).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsRealtimeAPI.AnalyticsGetRealtimeData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetRealtimeData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsRealtimeAPI.AnalyticsGetRealtimeData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetRealtimeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timezone** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

