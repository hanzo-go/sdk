# \DnsRecordsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsCreateRecord**](DnsRecordsAPI.md#DnsCreateRecord) | **Post** /v1/dns/zones/{zone}/records | Create DNS record
[**DnsDeleteRecord**](DnsRecordsAPI.md#DnsDeleteRecord) | **Delete** /v1/dns/zones/{zone}/records/{id} | Delete DNS record
[**DnsGetRecord**](DnsRecordsAPI.md#DnsGetRecord) | **Get** /v1/dns/zones/{zone}/records/{id} | Get DNS record
[**DnsListRecords**](DnsRecordsAPI.md#DnsListRecords) | **Get** /v1/dns/zones/{zone}/records | List DNS records
[**DnsUpdateRecord**](DnsRecordsAPI.md#DnsUpdateRecord) | **Put** /v1/dns/zones/{zone}/records/{id} | Update DNS record



## DnsCreateRecord

> DnsRecord DnsCreateRecord(ctx, zone).DnsRecordCreate(dnsRecordCreate).Execute()

Create DNS record

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
	dnsRecordCreate := *openapiclient.NewDnsRecordCreate("Name_example", "Type_example", "Content_example") // DnsRecordCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordsAPI.DnsCreateRecord(context.Background(), zone).DnsRecordCreate(dnsRecordCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordsAPI.DnsCreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsCreateRecord`: DnsRecord
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordsAPI.DnsCreateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsRecordCreate** | [**DnsRecordCreate**](DnsRecordCreate.md) |  | 

### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsDeleteRecord

> map[string]interface{} DnsDeleteRecord(ctx, zone, id).Execute()

Delete DNS record

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordsAPI.DnsDeleteRecord(context.Background(), zone, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordsAPI.DnsDeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsDeleteRecord`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordsAPI.DnsDeleteRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsDeleteRecordRequest struct via the builder pattern


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


## DnsGetRecord

> DnsRecord DnsGetRecord(ctx, zone, id).Execute()

Get DNS record

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordsAPI.DnsGetRecord(context.Background(), zone, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordsAPI.DnsGetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsGetRecord`: DnsRecord
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordsAPI.DnsGetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsListRecords

> DnsListRecords200Response DnsListRecords(ctx, zone).Type_(type_).Name(name).Page(page).PageSize(pageSize).Execute()

List DNS records

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
	type_ := "type__example" // string |  (optional)
	name := "name_example" // string | Filter by record name (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordsAPI.DnsListRecords(context.Background(), zone).Type_(type_).Name(name).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordsAPI.DnsListRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsListRecords`: DnsListRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordsAPI.DnsListRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsListRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 
 **name** | **string** | Filter by record name | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 100]

### Return type

[**DnsListRecords200Response**](DnsListRecords200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsUpdateRecord

> DnsRecord DnsUpdateRecord(ctx, zone, id).DnsUpdateRecordRequest(dnsUpdateRecordRequest).Execute()

Update DNS record

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dnsUpdateRecordRequest := *openapiclient.NewDnsUpdateRecordRequest() // DnsUpdateRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsRecordsAPI.DnsUpdateRecord(context.Background(), zone, id).DnsUpdateRecordRequest(dnsUpdateRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsRecordsAPI.DnsUpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsUpdateRecord`: DnsRecord
	fmt.Fprintf(os.Stdout, "Response from `DnsRecordsAPI.DnsUpdateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsUpdateRecordRequest** | [**DnsUpdateRecordRequest**](DnsUpdateRecordRequest.md) |  | 

### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

