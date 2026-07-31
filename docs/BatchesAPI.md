# \BatchesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGetBatch**](BatchesAPI.md#SearchGetBatch) | **Get** /v1/search/batches/{batchUid} | Get batch details
[**SearchListBatches**](BatchesAPI.md#SearchListBatches) | **Get** /v1/search/batches | List task batches



## SearchGetBatch

> SearchBatchView SearchGetBatch(ctx, batchUid).Execute()

Get batch details

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
	batchUid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.SearchGetBatch(context.Background(), batchUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.SearchGetBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetBatch`: SearchBatchView
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.SearchGetBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchUid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchBatchView**](SearchBatchView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListBatches

> SearchListBatches200Response SearchListBatches(ctx).Limit(limit).From(from).Execute()

List task batches

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
	limit := int32(56) // int32 |  (optional) (default to 20)
	from := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.SearchListBatches(context.Background()).Limit(limit).From(from).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.SearchListBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListBatches`: SearchListBatches200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.SearchListBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchListBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **from** | **int32** |  | 

### Return type

[**SearchListBatches200Response**](SearchListBatches200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

