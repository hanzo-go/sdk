# \GatewayCDNAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayGetCDNAnalytics**](GatewayCDNAPI.md#GatewayGetCDNAnalytics) | **Get** /v1/gateway/cdn/analytics | CDN cache analytics
[**GatewayPurgeCDNCache**](GatewayCDNAPI.md#GatewayPurgeCDNCache) | **Post** /v1/gateway/cdn/purge | Purge CDN cache



## GatewayGetCDNAnalytics

> GatewayGetCDNAnalytics200Response GatewayGetCDNAnalytics(ctx).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()

CDN cache analytics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	granularity := "granularity_example" // string |  (optional) (default to "hour")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayCDNAPI.GatewayGetCDNAnalytics(context.Background()).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayCDNAPI.GatewayGetCDNAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetCDNAnalytics`: GatewayGetCDNAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayCDNAPI.GatewayGetCDNAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetCDNAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **granularity** | **string** |  | [default to &quot;hour&quot;]

### Return type

[**GatewayGetCDNAnalytics200Response**](GatewayGetCDNAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayPurgeCDNCache

> GatewayPurgeCDNCache200Response GatewayPurgeCDNCache(ctx).GatewayCDNPurgeRequest(gatewayCDNPurgeRequest).Execute()

Purge CDN cache

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
	gatewayCDNPurgeRequest := *openapiclient.NewGatewayCDNPurgeRequest() // GatewayCDNPurgeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayCDNAPI.GatewayPurgeCDNCache(context.Background()).GatewayCDNPurgeRequest(gatewayCDNPurgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayCDNAPI.GatewayPurgeCDNCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayPurgeCDNCache`: GatewayPurgeCDNCache200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayCDNAPI.GatewayPurgeCDNCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayPurgeCDNCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayCDNPurgeRequest** | [**GatewayCDNPurgeRequest**](GatewayCDNPurgeRequest.md) |  | 

### Return type

[**GatewayPurgeCDNCache200Response**](GatewayPurgeCDNCache200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

