# \SearchDocsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1SearchDocsIndexes**](SearchDocsAPI.md#CloudGetV1SearchDocsIndexes) | **Get** /v1/search-docs/indexes | Lists the search indexes with their document counts and timestamps.
[**CloudGetV1SearchDocsStats**](SearchDocsAPI.md#CloudGetV1SearchDocsStats) | **Get** /v1/search-docs/stats | Totals the documents across every search index.



## CloudGetV1SearchDocsIndexes

> CloudSearchIndexList CloudGetV1SearchDocsIndexes(ctx).Execute()

Lists the search indexes with their document counts and timestamps.



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
	resp, r, err := apiClient.SearchDocsAPI.CloudGetV1SearchDocsIndexes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocsAPI.CloudGetV1SearchDocsIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SearchDocsIndexes`: CloudSearchIndexList
	fmt.Fprintf(os.Stdout, "Response from `SearchDocsAPI.CloudGetV1SearchDocsIndexes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SearchDocsIndexesRequest struct via the builder pattern


### Return type

[**CloudSearchIndexList**](CloudSearchIndexList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1SearchDocsStats

> CloudSearchStats CloudGetV1SearchDocsStats(ctx).Execute()

Totals the documents across every search index.



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
	resp, r, err := apiClient.SearchDocsAPI.CloudGetV1SearchDocsStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocsAPI.CloudGetV1SearchDocsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SearchDocsStats`: CloudSearchStats
	fmt.Fprintf(os.Stdout, "Response from `SearchDocsAPI.CloudGetV1SearchDocsStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SearchDocsStatsRequest struct via the builder pattern


### Return type

[**CloudSearchStats**](CloudSearchStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

