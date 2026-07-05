# \GatewayDNSAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayListDNSRecordsProxy**](GatewayDNSAPI.md#GatewayListDNSRecordsProxy) | **Get** /v1/gateway/dns/zones/{zone}/records | List DNS records (proxy to dns.hanzo.ai)
[**GatewayListDNSZonesProxy**](GatewayDNSAPI.md#GatewayListDNSZonesProxy) | **Get** /v1/gateway/dns/zones | List DNS zones (proxy to dns.hanzo.ai)



## GatewayListDNSRecordsProxy

> map[string]interface{} GatewayListDNSRecordsProxy(ctx, zone).Execute()

List DNS records (proxy to dns.hanzo.ai)

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
	zone := "zone_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayDNSAPI.GatewayListDNSRecordsProxy(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayDNSAPI.GatewayListDNSRecordsProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListDNSRecordsProxy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GatewayDNSAPI.GatewayListDNSRecordsProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListDNSRecordsProxyRequest struct via the builder pattern


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


## GatewayListDNSZonesProxy

> map[string]interface{} GatewayListDNSZonesProxy(ctx).Execute()

List DNS zones (proxy to dns.hanzo.ai)

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
	resp, r, err := apiClient.GatewayDNSAPI.GatewayListDNSZonesProxy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayDNSAPI.GatewayListDNSZonesProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListDNSZonesProxy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GatewayDNSAPI.GatewayListDNSZonesProxy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListDNSZonesProxyRequest struct via the builder pattern


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

