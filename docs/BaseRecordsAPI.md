# \BaseRecordsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BaseCreateRecord**](BaseRecordsAPI.md#BaseCreateRecord) | **Post** /v1/collections/{collection}/records | Create a record
[**BaseDeleteRecord**](BaseRecordsAPI.md#BaseDeleteRecord) | **Delete** /v1/collections/{collection}/records/{id} | Delete a record
[**BaseGetRecord**](BaseRecordsAPI.md#BaseGetRecord) | **Get** /v1/collections/{collection}/records/{id} | Get a record
[**BaseListRecords**](BaseRecordsAPI.md#BaseListRecords) | **Get** /v1/collections/{collection}/records | List records
[**BaseUpdateRecord**](BaseRecordsAPI.md#BaseUpdateRecord) | **Patch** /v1/collections/{collection}/records/{id} | Update a record



## BaseCreateRecord

> BaseRecord BaseCreateRecord(ctx, collection).BaseRecord(baseRecord).Execute()

Create a record

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
	collection := "collection_example" // string | Collection name or id (e.g. `site_drafts`).
	baseRecord := *openapiclient.NewBaseRecord() // BaseRecord | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaseRecordsAPI.BaseCreateRecord(context.Background(), collection).BaseRecord(baseRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseRecordsAPI.BaseCreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BaseCreateRecord`: BaseRecord
	fmt.Fprintf(os.Stdout, "Response from `BaseRecordsAPI.BaseCreateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | Collection name or id (e.g. &#x60;site_drafts&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBaseCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseRecord** | [**BaseRecord**](BaseRecord.md) |  | 

### Return type

[**BaseRecord**](BaseRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BaseDeleteRecord

> BaseDeleteRecord(ctx, collection, id).Execute()

Delete a record

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
	collection := "collection_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BaseRecordsAPI.BaseDeleteRecord(context.Background(), collection, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseRecordsAPI.BaseDeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBaseDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BaseGetRecord

> BaseRecord BaseGetRecord(ctx, collection, id).Execute()

Get a record

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
	collection := "collection_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaseRecordsAPI.BaseGetRecord(context.Background(), collection, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseRecordsAPI.BaseGetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BaseGetRecord`: BaseRecord
	fmt.Fprintf(os.Stdout, "Response from `BaseRecordsAPI.BaseGetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBaseGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseRecord**](BaseRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BaseListRecords

> BaseRecordList BaseListRecords(ctx, collection).Filter(filter).Sort(sort).Page(page).PerPage(perPage).Execute()

List records



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
	collection := "collection_example" // string | Collection name or id (e.g. `site_drafts`).
	filter := "filter_example" // string | Filter expression, e.g. org='hanzo' && slug='home'. (optional)
	sort := "sort_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaseRecordsAPI.BaseListRecords(context.Background(), collection).Filter(filter).Sort(sort).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseRecordsAPI.BaseListRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BaseListRecords`: BaseRecordList
	fmt.Fprintf(os.Stdout, "Response from `BaseRecordsAPI.BaseListRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | Collection name or id (e.g. &#x60;site_drafts&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBaseListRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter expression, e.g. org&#x3D;&#39;hanzo&#39; &amp;&amp; slug&#x3D;&#39;home&#39;. | 
 **sort** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 30]

### Return type

[**BaseRecordList**](BaseRecordList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BaseUpdateRecord

> BaseRecord BaseUpdateRecord(ctx, collection, id).BaseRecord(baseRecord).Execute()

Update a record

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
	collection := "collection_example" // string | 
	id := "id_example" // string | 
	baseRecord := *openapiclient.NewBaseRecord() // BaseRecord | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaseRecordsAPI.BaseUpdateRecord(context.Background(), collection, id).BaseRecord(baseRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaseRecordsAPI.BaseUpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BaseUpdateRecord`: BaseRecord
	fmt.Fprintf(os.Stdout, "Response from `BaseRecordsAPI.BaseUpdateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBaseUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **baseRecord** | [**BaseRecord**](BaseRecord.md) |  | 

### Return type

[**BaseRecord**](BaseRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

