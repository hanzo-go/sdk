# \AutoRecordsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreateRecord**](AutoRecordsAPI.md#AutoCreateRecord) | **Post** /v1/auto/records | Create a record
[**AutoListRecords**](AutoRecordsAPI.md#AutoListRecords) | **Get** /v1/auto/records | List records in a table



## AutoCreateRecord

> map[string]interface{} AutoCreateRecord(ctx).AutoCreateRecordRequest(autoCreateRecordRequest).Execute()

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
	resp, r, err := apiClient.AutoRecordsAPI.AutoCreateRecord(context.Background()).AutoCreateRecordRequest(autoCreateRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoRecordsAPI.AutoCreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateRecord`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoRecordsAPI.AutoCreateRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateRecordRequest struct via the builder pattern


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


## AutoListRecords

> map[string]interface{} AutoListRecords(ctx).TableId(tableId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoRecordsAPI.AutoListRecords(context.Background()).TableId(tableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoRecordsAPI.AutoListRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListRecords`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoRecordsAPI.AutoListRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableId** | **string** |  | 

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

