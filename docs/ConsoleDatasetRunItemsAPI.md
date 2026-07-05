# \ConsoleDatasetRunItemsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateDatasetRunItem**](ConsoleDatasetRunItemsAPI.md#ConsoleCreateDatasetRunItem) | **Post** /v1/console/dataset-run-items | Create a dataset run item



## ConsoleCreateDatasetRunItem

> ConsoleDatasetRunItem ConsoleCreateDatasetRunItem(ctx).ConsoleCreateDatasetRunItemRequest(consoleCreateDatasetRunItemRequest).Execute()

Create a dataset run item

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
	consoleCreateDatasetRunItemRequest := *openapiclient.NewConsoleCreateDatasetRunItemRequest("DatasetItemId_example", "TraceId_example", "RunName_example") // ConsoleCreateDatasetRunItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleDatasetRunItemsAPI.ConsoleCreateDatasetRunItem(context.Background()).ConsoleCreateDatasetRunItemRequest(consoleCreateDatasetRunItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleDatasetRunItemsAPI.ConsoleCreateDatasetRunItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateDatasetRunItem`: ConsoleDatasetRunItem
	fmt.Fprintf(os.Stdout, "Response from `ConsoleDatasetRunItemsAPI.ConsoleCreateDatasetRunItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateDatasetRunItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateDatasetRunItemRequest** | [**ConsoleCreateDatasetRunItemRequest**](ConsoleCreateDatasetRunItemRequest.md) |  | 

### Return type

[**ConsoleDatasetRunItem**](ConsoleDatasetRunItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

