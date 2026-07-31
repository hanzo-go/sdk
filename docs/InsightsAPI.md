# \InsightsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1InsightsEvents**](InsightsAPI.md#CloudGetV1InsightsEvents) | **Get** /v1/insights/events | InsightsEvents returns the caller org&#39;s most recent product events, newest first.
[**CloudGetV1InsightsHealth**](InsightsAPI.md#CloudGetV1InsightsHealth) | **Get** /v1/insights/health | InsightsHealth reports that the unified insights surface is serving.



## CloudGetV1InsightsEvents

> CloudEventList CloudGetV1InsightsEvents(ctx).Limit(limit).Execute()

InsightsEvents returns the caller org's most recent product events, newest first.



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
	limit := int32(100) // int32 | Limit is how many rows to return, newest first. Default 50, maximum 200; a value at or below zero, or one that is not a number, takes the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.CloudGetV1InsightsEvents(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.CloudGetV1InsightsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1InsightsEvents`: CloudEventList
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.CloudGetV1InsightsEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1InsightsEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit is how many rows to return, newest first. Default 50, maximum 200; a value at or below zero, or one that is not a number, takes the default. | 

### Return type

[**CloudEventList**](CloudEventList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1InsightsHealth

> CloudInsightsStatus CloudGetV1InsightsHealth(ctx).Execute()

InsightsHealth reports that the unified insights surface is serving.



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
	resp, r, err := apiClient.InsightsAPI.CloudGetV1InsightsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.CloudGetV1InsightsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1InsightsHealth`: CloudInsightsStatus
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.CloudGetV1InsightsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1InsightsHealthRequest struct via the builder pattern


### Return type

[**CloudInsightsStatus**](CloudInsightsStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

