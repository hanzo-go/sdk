# \AnalyticsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsGetZoneAnalytics**](AnalyticsAPI.md#DnsGetZoneAnalytics) | **Get** /v1/dns/zones/{zone}/analytics | Get query analytics



## DnsGetZoneAnalytics

> DnsQueryAnalytics DnsGetZoneAnalytics(ctx, zone).From(from).To(to).Granularity(granularity).Execute()

Get query analytics

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
	zone := "zone_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	granularity := "granularity_example" // string |  (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.DnsGetZoneAnalytics(context.Background(), zone).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.DnsGetZoneAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsGetZoneAnalytics`: DnsQueryAnalytics
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.DnsGetZoneAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsGetZoneAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **granularity** | **string** |  | [default to &quot;day&quot;]

### Return type

[**DnsQueryAnalytics**](DnsQueryAnalytics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

