# \AnalyticsCollectAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsBatch**](AnalyticsCollectAPI.md#AnalyticsBatch) | **Post** /v1/analytics/batch | Send a batch of events
[**AnalyticsSend**](AnalyticsCollectAPI.md#AnalyticsSend) | **Post** /v1/analytics/send | Send a single event or identify payload



## AnalyticsBatch

> AnalyticsBatch200Response AnalyticsBatch(ctx).AnalyticsSendPayload(analyticsSendPayload).Execute()

Send a batch of events

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
	analyticsSendPayload := []openapiclient.AnalyticsSendPayload{*openapiclient.NewAnalyticsSendPayload("Type_example", *openapiclient.NewAnalyticsSendPayloadPayload("Website_example"))} // []AnalyticsSendPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsCollectAPI.AnalyticsBatch(context.Background()).AnalyticsSendPayload(analyticsSendPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsCollectAPI.AnalyticsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsBatch`: AnalyticsBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsCollectAPI.AnalyticsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsSendPayload** | [**[]AnalyticsSendPayload**](AnalyticsSendPayload.md) |  | 

### Return type

[**AnalyticsBatch200Response**](AnalyticsBatch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsSend

> AnalyticsSend200Response AnalyticsSend(ctx).AnalyticsSendPayload(analyticsSendPayload).Execute()

Send a single event or identify payload

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
	analyticsSendPayload := *openapiclient.NewAnalyticsSendPayload("Type_example", *openapiclient.NewAnalyticsSendPayloadPayload("Website_example")) // AnalyticsSendPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsCollectAPI.AnalyticsSend(context.Background()).AnalyticsSendPayload(analyticsSendPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsCollectAPI.AnalyticsSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsSend`: AnalyticsSend200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsCollectAPI.AnalyticsSend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsSendPayload** | [**AnalyticsSendPayload**](AnalyticsSendPayload.md) |  | 

### Return type

[**AnalyticsSend200Response**](AnalyticsSend200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

