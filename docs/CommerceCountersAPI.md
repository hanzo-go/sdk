# \CommerceCountersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceGetDailyDashboard**](CommerceCountersAPI.md#CommerceGetDailyDashboard) | **Post** /v1/commerce/counter/dashboard/daily | Get daily dashboard metrics
[**CommerceGetProductCounters**](CommerceCountersAPI.md#CommerceGetProductCounters) | **Get** /v1/commerce/counter/product/{productid} | Get product counters
[**CommerceGetToplineMetrics**](CommerceCountersAPI.md#CommerceGetToplineMetrics) | **Get** /v1/commerce/counter/topline | Get topline metrics
[**CommerceSearchCounters**](CommerceCountersAPI.md#CommerceSearchCounters) | **Post** /v1/commerce/counter | Search counters



## CommerceGetDailyDashboard

> map[string]interface{} CommerceGetDailyDashboard(ctx).CommerceGetDailyDashboardRequest(commerceGetDailyDashboardRequest).Execute()

Get daily dashboard metrics

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
	commerceGetDailyDashboardRequest := *openapiclient.NewCommerceGetDailyDashboardRequest() // CommerceGetDailyDashboardRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCountersAPI.CommerceGetDailyDashboard(context.Background()).CommerceGetDailyDashboardRequest(commerceGetDailyDashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCountersAPI.CommerceGetDailyDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetDailyDashboard`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceCountersAPI.CommerceGetDailyDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetDailyDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceGetDailyDashboardRequest** | [**CommerceGetDailyDashboardRequest**](CommerceGetDailyDashboardRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetProductCounters

> map[string]interface{} CommerceGetProductCounters(ctx, productid).Execute()

Get product counters

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
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCountersAPI.CommerceGetProductCounters(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCountersAPI.CommerceGetProductCounters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetProductCounters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceCountersAPI.CommerceGetProductCounters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetProductCountersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CommerceGetToplineMetrics

> map[string]interface{} CommerceGetToplineMetrics(ctx).Execute()

Get topline metrics

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
	resp, r, err := apiClient.CommerceCountersAPI.CommerceGetToplineMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCountersAPI.CommerceGetToplineMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetToplineMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceCountersAPI.CommerceGetToplineMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetToplineMetricsRequest struct via the builder pattern


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


## CommerceSearchCounters

> map[string]interface{} CommerceSearchCounters(ctx).CommerceSearchCountersRequest(commerceSearchCountersRequest).Execute()

Search counters

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
	commerceSearchCountersRequest := *openapiclient.NewCommerceSearchCountersRequest() // CommerceSearchCountersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCountersAPI.CommerceSearchCounters(context.Background()).CommerceSearchCountersRequest(commerceSearchCountersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCountersAPI.CommerceSearchCounters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchCounters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceCountersAPI.CommerceSearchCounters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchCountersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceSearchCountersRequest** | [**CommerceSearchCountersRequest**](CommerceSearchCountersRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

