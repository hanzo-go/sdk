# \DnsZonesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsCreateZone**](DnsZonesAPI.md#DnsCreateZone) | **Post** /v1/dns/zones | Create zone
[**DnsDeleteZone**](DnsZonesAPI.md#DnsDeleteZone) | **Delete** /v1/dns/zones/{zone} | Delete zone
[**DnsExportZone**](DnsZonesAPI.md#DnsExportZone) | **Get** /v1/dns/zones/{zone}/export | Export zone file
[**DnsGetZone**](DnsZonesAPI.md#DnsGetZone) | **Get** /v1/dns/zones/{zone} | Get zone
[**DnsImportZone**](DnsZonesAPI.md#DnsImportZone) | **Post** /v1/dns/zones/{zone}/import | Import zone file
[**DnsListZones**](DnsZonesAPI.md#DnsListZones) | **Get** /v1/dns/zones | List zones
[**DnsUpdateZone**](DnsZonesAPI.md#DnsUpdateZone) | **Put** /v1/dns/zones/{zone} | Update zone



## DnsCreateZone

> DnsZone DnsCreateZone(ctx).DnsZoneCreate(dnsZoneCreate).Execute()

Create zone

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
	dnsZoneCreate := *openapiclient.NewDnsZoneCreate("Zone_example") // DnsZoneCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.DnsCreateZone(context.Background()).DnsZoneCreate(dnsZoneCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsCreateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsCreateZone`: DnsZone
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsCreateZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsZoneCreate** | [**DnsZoneCreate**](DnsZoneCreate.md) |  | 

### Return type

[**DnsZone**](DnsZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsDeleteZone

> map[string]interface{} DnsDeleteZone(ctx, zone).Execute()

Delete zone

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
	resp, r, err := apiClient.DnsZonesAPI.DnsDeleteZone(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsDeleteZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsDeleteZone`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsDeleteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsDeleteZoneRequest struct via the builder pattern


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


## DnsExportZone

> string DnsExportZone(ctx, zone).Execute()

Export zone file



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
	resp, r, err := apiClient.DnsZonesAPI.DnsExportZone(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsExportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsExportZone`: string
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsExportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsExportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsGetZone

> DnsZone DnsGetZone(ctx, zone).Execute()

Get zone

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
	resp, r, err := apiClient.DnsZonesAPI.DnsGetZone(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsGetZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsGetZone`: DnsZone
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsGetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnsZone**](DnsZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsImportZone

> DnsImportZone200Response DnsImportZone(ctx, zone).Body(body).Execute()

Import zone file



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.DnsImportZone(context.Background(), zone).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsImportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsImportZone`: DnsImportZone200Response
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsImportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsImportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**DnsImportZone200Response**](DnsImportZone200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsListZones

> DnsListZones200Response DnsListZones(ctx).Status(status).Name(name).Page(page).PageSize(pageSize).Execute()

List zones

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
	status := "status_example" // string |  (optional)
	name := "name_example" // string | Filter by zone name (substring match) (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.DnsListZones(context.Background()).Status(status).Name(name).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsListZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsListZones`: DnsListZones200Response
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsListZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsListZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **name** | **string** | Filter by zone name (substring match) | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**DnsListZones200Response**](DnsListZones200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsUpdateZone

> DnsZone DnsUpdateZone(ctx, zone).DnsUpdateZoneRequest(dnsUpdateZoneRequest).Execute()

Update zone

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
	dnsUpdateZoneRequest := *openapiclient.NewDnsUpdateZoneRequest() // DnsUpdateZoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.DnsUpdateZone(context.Background(), zone).DnsUpdateZoneRequest(dnsUpdateZoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.DnsUpdateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsUpdateZone`: DnsZone
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.DnsUpdateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsUpdateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsUpdateZoneRequest** | [**DnsUpdateZoneRequest**](DnsUpdateZoneRequest.md) |  | 

### Return type

[**DnsZone**](DnsZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

