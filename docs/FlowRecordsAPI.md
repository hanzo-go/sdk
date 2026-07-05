# \FlowRecordsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCreateRecord**](FlowRecordsAPI.md#FlowCreateRecord) | **Post** /v1/flow/records | Create a record
[**FlowDeleteRecord**](FlowRecordsAPI.md#FlowDeleteRecord) | **Delete** /v1/flow/records/{id} | Delete a record
[**FlowGetRecord**](FlowRecordsAPI.md#FlowGetRecord) | **Get** /v1/flow/records/{id} | Get a record
[**FlowListRecords**](FlowRecordsAPI.md#FlowListRecords) | **Get** /v1/flow/records | List records in a table
[**FlowUpdateRecord**](FlowRecordsAPI.md#FlowUpdateRecord) | **Patch** /v1/flow/records/{id} | Update a record



## FlowCreateRecord

> map[string]interface{} FlowCreateRecord(ctx).AutoCreateRecordRequest(autoCreateRecordRequest).Execute()

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
	autoCreateRecordRequest := *openapiclient.NewAutoCreateRecordRequest("TableId_example", map[string]interface{}(123)) // AutoCreateRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowRecordsAPI.FlowCreateRecord(context.Background()).AutoCreateRecordRequest(autoCreateRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowRecordsAPI.FlowCreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateRecord`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowRecordsAPI.FlowCreateRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateRecordRequest** | [**AutoCreateRecordRequest**](AutoCreateRecordRequest.md) |  | 

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


## FlowDeleteRecord

> FlowDeleteRecord(ctx, id).Execute()

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowRecordsAPI.FlowDeleteRecord(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowRecordsAPI.FlowDeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetRecord

> map[string]interface{} FlowGetRecord(ctx, id).Execute()

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowRecordsAPI.FlowGetRecord(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowRecordsAPI.FlowGetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetRecord`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowRecordsAPI.FlowGetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetRecordRequest struct via the builder pattern


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


## FlowListRecords

> map[string]interface{} FlowListRecords(ctx).TableId(tableId).Cursor(cursor).Limit(limit).Execute()

List records in a table

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
	tableId := "tableId_example" // string | 
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowRecordsAPI.FlowListRecords(context.Background()).TableId(tableId).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowRecordsAPI.FlowListRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListRecords`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowRecordsAPI.FlowListRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableId** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

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


## FlowUpdateRecord

> map[string]interface{} FlowUpdateRecord(ctx, id).FlowUpdateRecordRequest(flowUpdateRecordRequest).Execute()

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
	id := "id_example" // string | 
	flowUpdateRecordRequest := *openapiclient.NewFlowUpdateRecordRequest() // FlowUpdateRecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowRecordsAPI.FlowUpdateRecord(context.Background(), id).FlowUpdateRecordRequest(flowUpdateRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowRecordsAPI.FlowUpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpdateRecord`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowRecordsAPI.FlowUpdateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowUpdateRecordRequest** | [**FlowUpdateRecordRequest**](FlowUpdateRecordRequest.md) |  | 

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

