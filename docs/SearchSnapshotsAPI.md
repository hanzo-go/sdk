# \SearchSnapshotsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCreateSnapshot**](SearchSnapshotsAPI.md#SearchCreateSnapshot) | **Post** /v1/search/snapshots | Create a database snapshot



## SearchCreateSnapshot

> SearchSummarizedTaskView SearchCreateSnapshot(ctx).Execute()

Create a database snapshot

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
	resp, r, err := apiClient.SearchSnapshotsAPI.SearchCreateSnapshot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchSnapshotsAPI.SearchCreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCreateSnapshot`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchSnapshotsAPI.SearchCreateSnapshot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCreateSnapshotRequest struct via the builder pattern


### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

