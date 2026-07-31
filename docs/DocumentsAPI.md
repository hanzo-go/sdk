# \DocumentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworkCancelDocument**](DocumentsAPI.md#FrameworkCancelDocument) | **Post** /v1/framework/{doctype}/{name}/cancel | Cancel a submitted document (docstatus 1→2)
[**FrameworkCreateDocument**](DocumentsAPI.md#FrameworkCreateDocument) | **Post** /v1/framework/{doctype} | Create a document
[**FrameworkDeleteDocument**](DocumentsAPI.md#FrameworkDeleteDocument) | **Delete** /v1/framework/{doctype}/{name} | Delete a document
[**FrameworkGetDocument**](DocumentsAPI.md#FrameworkGetDocument) | **Get** /v1/framework/{doctype}/{name} | Get a document
[**FrameworkListDocuments**](DocumentsAPI.md#FrameworkListDocuments) | **Get** /v1/framework/{doctype} | List documents of a DocType
[**FrameworkSubmitDocument**](DocumentsAPI.md#FrameworkSubmitDocument) | **Post** /v1/framework/{doctype}/{name}/submit | Submit a document (docstatus 0→1)
[**FrameworkUpdateDocument**](DocumentsAPI.md#FrameworkUpdateDocument) | **Put** /v1/framework/{doctype}/{name} | Update a draft document
[**SearchAddOrReplaceDocuments**](DocumentsAPI.md#SearchAddOrReplaceDocuments) | **Post** /v1/search/indexes/{indexUid}/documents | Add or replace documents
[**SearchAddOrUpdateDocuments**](DocumentsAPI.md#SearchAddOrUpdateDocuments) | **Put** /v1/search/indexes/{indexUid}/documents | Add or update documents (partial)
[**SearchDeleteAllDocuments**](DocumentsAPI.md#SearchDeleteAllDocuments) | **Delete** /v1/search/indexes/{indexUid}/documents | Delete all documents in the index
[**SearchDeleteDocument**](DocumentsAPI.md#SearchDeleteDocument) | **Delete** /v1/search/indexes/{indexUid}/documents/{documentId} | Delete a single document
[**SearchDeleteDocumentsBatch**](DocumentsAPI.md#SearchDeleteDocumentsBatch) | **Post** /v1/search/indexes/{indexUid}/documents/delete-batch | Delete documents by IDs
[**SearchDeleteDocumentsByFilter**](DocumentsAPI.md#SearchDeleteDocumentsByFilter) | **Post** /v1/search/indexes/{indexUid}/documents/delete | Delete documents by filter
[**SearchEditDocumentsByFunction**](DocumentsAPI.md#SearchEditDocumentsByFunction) | **Post** /v1/search/indexes/{indexUid}/documents/edit | Edit documents using a function
[**SearchGetDocument**](DocumentsAPI.md#SearchGetDocument) | **Get** /v1/search/indexes/{indexUid}/documents/{documentId} | Get a single document
[**SearchGetDocuments**](DocumentsAPI.md#SearchGetDocuments) | **Get** /v1/search/indexes/{indexUid}/documents | Browse documents



## FrameworkCancelDocument

> FrameworkDocument FrameworkCancelDocument(ctx, doctype, name).Execute()

Cancel a submitted document (docstatus 1→2)

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.FrameworkCancelDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkCancelDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkCancelDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.FrameworkCancelDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkCancelDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkCreateDocument

> FrameworkDocument FrameworkCreateDocument(ctx, doctype).FrameworkDocument(frameworkDocument).Execute()

Create a document

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
	doctype := "doctype_example" // string | 
	frameworkDocument := *openapiclient.NewFrameworkDocument() // FrameworkDocument | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.FrameworkCreateDocument(context.Background(), doctype).FrameworkDocument(frameworkDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkCreateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkCreateDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.FrameworkCreateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameworkDocument** | [**FrameworkDocument**](FrameworkDocument.md) |  | 

### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkDeleteDocument

> FrameworkDeleteDocument(ctx, doctype, name).Execute()

Delete a document

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DocumentsAPI.FrameworkDeleteDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkDeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkDeleteDocumentRequest struct via the builder pattern


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


## FrameworkGetDocument

> FrameworkDocument FrameworkGetDocument(ctx, doctype, name).Execute()

Get a document

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.FrameworkGetDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkGetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkGetDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.FrameworkGetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkListDocuments

> FrameworkListDocuments200Response FrameworkListDocuments(ctx, doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()

List documents of a DocType

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
	doctype := "doctype_example" // string | 
	filters := "filters_example" // string | 'JSON object of field→value filters (declared fields, name, or docstatus)' (optional)
	fields := "fields_example" // string | Comma list or JSON array of field names to project (optional)
	orderBy := "orderBy_example" // string | field [asc|desc] (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.FrameworkListDocuments(context.Background(), doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkListDocuments`: FrameworkListDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.FrameworkListDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | &#39;JSON object of field→value filters (declared fields, name, or docstatus)&#39; | 
 **fields** | **string** | Comma list or JSON array of field names to project | 
 **orderBy** | **string** | field [asc|desc] | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**FrameworkListDocuments200Response**](FrameworkListDocuments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkSubmitDocument

> FrameworkDocument FrameworkSubmitDocument(ctx, doctype, name).Execute()

Submit a document (docstatus 0→1)

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.FrameworkSubmitDocument(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkSubmitDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkSubmitDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.FrameworkSubmitDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkSubmitDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkUpdateDocument

> FrameworkDocument FrameworkUpdateDocument(ctx, doctype, name).FrameworkDocument(frameworkDocument).Execute()

Update a draft document

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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 
	frameworkDocument := *openapiclient.NewFrameworkDocument() // FrameworkDocument | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.FrameworkUpdateDocument(context.Background(), doctype, name).FrameworkDocument(frameworkDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.FrameworkUpdateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkUpdateDocument`: FrameworkDocument
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.FrameworkUpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **frameworkDocument** | [**FrameworkDocument**](FrameworkDocument.md) |  | 

### Return type

[**FrameworkDocument**](FrameworkDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.DocumentsAPI.SearchAddOrReplaceDocuments(context.Background(), indexUid).RequestBody(requestBody).PrimaryKey(primaryKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchAddOrReplaceDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAddOrReplaceDocuments`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchAddOrReplaceDocuments`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchAddOrUpdateDocuments(context.Background(), indexUid).RequestBody(requestBody).PrimaryKey(primaryKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchAddOrUpdateDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAddOrUpdateDocuments`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchAddOrUpdateDocuments`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchDeleteAllDocuments(context.Background(), indexUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchDeleteAllDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteAllDocuments`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchDeleteAllDocuments`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchDeleteDocument(context.Background(), indexUid, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchDeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteDocument`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchDeleteDocument`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchDeleteDocumentsBatch(context.Background(), indexUid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchDeleteDocumentsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteDocumentsBatch`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchDeleteDocumentsBatch`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchDeleteDocumentsByFilter(context.Background(), indexUid).SearchDeleteDocumentsByFilterRequest(searchDeleteDocumentsByFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchDeleteDocumentsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteDocumentsByFilter`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchDeleteDocumentsByFilter`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchEditDocumentsByFunction(context.Background(), indexUid).SearchEditDocumentsByFunctionRequest(searchEditDocumentsByFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchEditDocumentsByFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEditDocumentsByFunction`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchEditDocumentsByFunction`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchGetDocument(context.Background(), indexUid, documentId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchGetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetDocument`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchGetDocument`: %v\n", resp)
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
	resp, r, err := apiClient.DocumentsAPI.SearchGetDocuments(context.Background(), indexUid).Offset(offset).Limit(limit).Fields(fields).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchGetDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetDocuments`: SearchPaginatedDocuments
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchGetDocuments`: %v\n", resp)
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

