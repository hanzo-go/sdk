# \DocumentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDocuments**](DocumentsAPI.md#DeleteDocuments) | **Delete** /v1/documents | Handles DELETE /v1/documents — a JSON array of file_ids.
[**GetDocumentsByFileIdContext**](DocumentsAPI.md#GetDocumentsByFileIdContext) | **Get** /v1/documents/{file_id}/context | Handles GET /v1/documents/:file_id/context — every chunk of a file, as LangChain Documents (used when RAG_USE_FULL_CONTEXT is on).



## DeleteDocuments

> DeleteDocuments(ctx).Execute()

Handles DELETE /v1/documents — a JSON array of file_ids.



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
	r, err := apiClient.DocumentsAPI.DeleteDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DeleteDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentsByFileIdContext

> GetDocumentsByFileIdContext(ctx).Execute()

Handles GET /v1/documents/:file_id/context — every chunk of a file, as LangChain Documents (used when RAG_USE_FULL_CONTEXT is on).



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
	r, err := apiClient.DocumentsAPI.GetDocumentsByFileIdContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentsByFileIdContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsByFileIdContextRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

