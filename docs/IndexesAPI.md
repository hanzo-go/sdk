# \IndexesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCreateIndex**](IndexesAPI.md#SearchCreateIndex) | **Post** /v1/search/indexes | Create a new index
[**SearchDeleteIndex**](IndexesAPI.md#SearchDeleteIndex) | **Delete** /v1/search/indexes/{indexUid} | Delete an index
[**SearchGetIndex**](IndexesAPI.md#SearchGetIndex) | **Get** /v1/search/indexes/{indexUid} | Get index information
[**SearchListIndexes**](IndexesAPI.md#SearchListIndexes) | **Get** /v1/search/indexes | List all indexes
[**SearchUpdateIndex**](IndexesAPI.md#SearchUpdateIndex) | **Patch** /v1/search/indexes/{indexUid} | Update index (primary key)



## SearchCreateIndex

> SearchSummarizedTaskView SearchCreateIndex(ctx).SearchIndexCreateRequest(searchIndexCreateRequest).Execute()

Create a new index

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
	searchIndexCreateRequest := *openapiclient.NewSearchIndexCreateRequest("Uid_example") // SearchIndexCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexesAPI.SearchCreateIndex(context.Background()).SearchIndexCreateRequest(searchIndexCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.SearchCreateIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCreateIndex`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.SearchCreateIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCreateIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchIndexCreateRequest** | [**SearchIndexCreateRequest**](SearchIndexCreateRequest.md) |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeleteIndex

> SearchSummarizedTaskView SearchDeleteIndex(ctx, indexUid).Execute()

Delete an index

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
	indexUid := "indexUid_example" // string | Unique index identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexesAPI.SearchDeleteIndex(context.Background(), indexUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.SearchDeleteIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteIndex`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.SearchDeleteIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SearchGetIndex

> SearchIndexView SearchGetIndex(ctx, indexUid).Execute()

Get index information

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
	indexUid := "indexUid_example" // string | Unique index identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexesAPI.SearchGetIndex(context.Background(), indexUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.SearchGetIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetIndex`: SearchIndexView
	fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.SearchGetIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchIndexView**](SearchIndexView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListIndexes

> SearchPaginatedIndexes SearchListIndexes(ctx).Offset(offset).Limit(limit).Execute()

List all indexes

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
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexesAPI.SearchListIndexes(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.SearchListIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListIndexes`: SearchPaginatedIndexes
	fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.SearchListIndexes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchListIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**SearchPaginatedIndexes**](SearchPaginatedIndexes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateIndex

> SearchSummarizedTaskView SearchUpdateIndex(ctx, indexUid).SearchUpdateIndexRequest(searchUpdateIndexRequest).Execute()

Update index (primary key)

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
	indexUid := "indexUid_example" // string | Unique index identifier
	searchUpdateIndexRequest := *openapiclient.NewSearchUpdateIndexRequest() // SearchUpdateIndexRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexesAPI.SearchUpdateIndex(context.Background(), indexUid).SearchUpdateIndexRequest(searchUpdateIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.SearchUpdateIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateIndex`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.SearchUpdateIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchUpdateIndexRequest** | [**SearchUpdateIndexRequest**](SearchUpdateIndexRequest.md) |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

