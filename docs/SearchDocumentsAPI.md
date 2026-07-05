# \SearchDocumentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAddOrReplaceDocuments**](SearchDocumentsAPI.md#SearchAddOrReplaceDocuments) | **Post** /v1/search/indexes/{indexUid}/documents | Add or replace documents
[**SearchAddOrUpdateDocuments**](SearchDocumentsAPI.md#SearchAddOrUpdateDocuments) | **Put** /v1/search/indexes/{indexUid}/documents | Add or update documents (partial)
[**SearchDeleteAllDocuments**](SearchDocumentsAPI.md#SearchDeleteAllDocuments) | **Delete** /v1/search/indexes/{indexUid}/documents | Delete all documents in the index
[**SearchDeleteDocument**](SearchDocumentsAPI.md#SearchDeleteDocument) | **Delete** /v1/search/indexes/{indexUid}/documents/{documentId} | Delete a single document
[**SearchDeleteDocumentsBatch**](SearchDocumentsAPI.md#SearchDeleteDocumentsBatch) | **Post** /v1/search/indexes/{indexUid}/documents/delete-batch | Delete documents by IDs
[**SearchDeleteDocumentsByFilter**](SearchDocumentsAPI.md#SearchDeleteDocumentsByFilter) | **Post** /v1/search/indexes/{indexUid}/documents/delete | Delete documents by filter
[**SearchEditDocumentsByFunction**](SearchDocumentsAPI.md#SearchEditDocumentsByFunction) | **Post** /v1/search/indexes/{indexUid}/documents/edit | Edit documents using a function
[**SearchGetDocument**](SearchDocumentsAPI.md#SearchGetDocument) | **Get** /v1/search/indexes/{indexUid}/documents/{documentId} | Get a single document
[**SearchGetDocuments**](SearchDocumentsAPI.md#SearchGetDocuments) | **Get** /v1/search/indexes/{indexUid}/documents | Browse documents



## SearchAddOrReplaceDocuments

> SearchSummarizedTaskView SearchAddOrReplaceDocuments(ctx, indexUid).RequestBody(requestBody).PrimaryKey(primaryKey).Execute()

Add or replace documents

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
	requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | 
	primaryKey := "primaryKey_example" // string | Primary key field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchAddOrReplaceDocuments(context.Background(), indexUid).RequestBody(requestBody).PrimaryKey(primaryKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchAddOrReplaceDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAddOrReplaceDocuments`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchAddOrReplaceDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAddOrReplaceDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** |  | 
 **primaryKey** | **string** | Primary key field name | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-ndjson
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAddOrUpdateDocuments

> SearchSummarizedTaskView SearchAddOrUpdateDocuments(ctx, indexUid).RequestBody(requestBody).PrimaryKey(primaryKey).Execute()

Add or update documents (partial)

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
	requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | 
	primaryKey := "primaryKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchAddOrUpdateDocuments(context.Background(), indexUid).RequestBody(requestBody).PrimaryKey(primaryKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchAddOrUpdateDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAddOrUpdateDocuments`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchAddOrUpdateDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAddOrUpdateDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** |  | 
 **primaryKey** | **string** |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-ndjson
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeleteAllDocuments

> SearchSummarizedTaskView SearchDeleteAllDocuments(ctx, indexUid).Execute()

Delete all documents in the index

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
	resp, r, err := apiClient.SearchDocumentsAPI.SearchDeleteAllDocuments(context.Background(), indexUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchDeleteAllDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteAllDocuments`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchDeleteAllDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteAllDocumentsRequest struct via the builder pattern


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


## SearchDeleteDocument

> SearchSummarizedTaskView SearchDeleteDocument(ctx, indexUid, documentId).Execute()

Delete a single document

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
	documentId := "documentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchDeleteDocument(context.Background(), indexUid, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchDeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteDocument`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchDeleteDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteDocumentRequest struct via the builder pattern


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


## SearchDeleteDocumentsBatch

> SearchSummarizedTaskView SearchDeleteDocumentsBatch(ctx, indexUid).RequestBody(requestBody).Execute()

Delete documents by IDs

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
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchDeleteDocumentsBatch(context.Background(), indexUid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchDeleteDocumentsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteDocumentsBatch`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchDeleteDocumentsBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteDocumentsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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


## SearchDeleteDocumentsByFilter

> SearchSummarizedTaskView SearchDeleteDocumentsByFilter(ctx, indexUid).SearchDeleteDocumentsByFilterRequest(searchDeleteDocumentsByFilterRequest).Execute()

Delete documents by filter

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
	searchDeleteDocumentsByFilterRequest := *openapiclient.NewSearchDeleteDocumentsByFilterRequest(openapiclient.search_deleteDocumentsByFilter_request_filter{ArrayOfString: new([]string)}) // SearchDeleteDocumentsByFilterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchDeleteDocumentsByFilter(context.Background(), indexUid).SearchDeleteDocumentsByFilterRequest(searchDeleteDocumentsByFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchDeleteDocumentsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteDocumentsByFilter`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchDeleteDocumentsByFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteDocumentsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchDeleteDocumentsByFilterRequest** | [**SearchDeleteDocumentsByFilterRequest**](SearchDeleteDocumentsByFilterRequest.md) |  | 

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


## SearchEditDocumentsByFunction

> SearchSummarizedTaskView SearchEditDocumentsByFunction(ctx, indexUid).SearchEditDocumentsByFunctionRequest(searchEditDocumentsByFunctionRequest).Execute()

Edit documents using a function

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
	searchEditDocumentsByFunctionRequest := *openapiclient.NewSearchEditDocumentsByFunctionRequest("Function_example") // SearchEditDocumentsByFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchEditDocumentsByFunction(context.Background(), indexUid).SearchEditDocumentsByFunctionRequest(searchEditDocumentsByFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchEditDocumentsByFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEditDocumentsByFunction`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchEditDocumentsByFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEditDocumentsByFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchEditDocumentsByFunctionRequest** | [**SearchEditDocumentsByFunctionRequest**](SearchEditDocumentsByFunctionRequest.md) |  | 

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


## SearchGetDocument

> map[string]interface{} SearchGetDocument(ctx, indexUid, documentId).Fields(fields).Execute()

Get a single document

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
	documentId := "documentId_example" // string | 
	fields := "fields_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchGetDocument(context.Background(), indexUid, documentId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchGetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetDocument`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchGetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** |  | 

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


## SearchGetDocuments

> SearchPaginatedDocuments SearchGetDocuments(ctx, indexUid).Offset(offset).Limit(limit).Fields(fields).Filter(filter).Execute()

Browse documents

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
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)
	fields := "fields_example" // string | Comma-separated fields to return (optional)
	filter := "filter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchDocumentsAPI.SearchGetDocuments(context.Background(), indexUid).Offset(offset).Limit(limit).Fields(fields).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchDocumentsAPI.SearchGetDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetDocuments`: SearchPaginatedDocuments
	fmt.Fprintf(os.Stdout, "Response from `SearchDocumentsAPI.SearchGetDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **fields** | **string** | Comma-separated fields to return | 
 **filter** | **string** |  | 

### Return type

[**SearchPaginatedDocuments**](SearchPaginatedDocuments.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

