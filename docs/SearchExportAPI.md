# \SearchExportAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchExportData**](SearchExportAPI.md#SearchExportData) | **Post** /v1/search/export | Export index data



## SearchExportData

> string SearchExportData(ctx).SearchExportDataRequest(searchExportDataRequest).Execute()

Export index data

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
	searchExportDataRequest := *openapiclient.NewSearchExportDataRequest() // SearchExportDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchExportAPI.SearchExportData(context.Background()).SearchExportDataRequest(searchExportDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchExportAPI.SearchExportData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchExportData`: string
	fmt.Fprintf(os.Stdout, "Response from `SearchExportAPI.SearchExportData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchExportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchExportDataRequest** | [**SearchExportDataRequest**](SearchExportDataRequest.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

