# \DocsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDocsIngest**](DocsAPI.md#PostDocsIngest) | **Post** /v1/docs/ingest | Unified RAG ingest: parse + chunk + embed documents and pipe them to BOTH Hanzo Vector (semantic) AND Hanzo Search (keyword) under the tenant index {owner}-{store}-docs — the same index /v1/chat retrieval reads.



## PostDocsIngest

> PostDocsIngest(ctx).Execute()

Unified RAG ingest: parse + chunk + embed documents and pipe them to BOTH Hanzo Vector (semantic) AND Hanzo Search (keyword) under the tenant index {owner}-{store}-docs — the same index /v1/chat retrieval reads.



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
	r, err := apiClient.DocsAPI.PostDocsIngest(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.PostDocsIngest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDocsIngestRequest struct via the builder pattern


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

