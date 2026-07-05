# \PromptsCatalogAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PromptsPromptCatalog**](PromptsCatalogAPI.md#PromptsPromptCatalog) | **Get** /v1/prompts/catalog | Read-only starter prompt library



## PromptsPromptCatalog

> PromptsPromptCatalog200Response PromptsPromptCatalog(ctx).Execute()

Read-only starter prompt library

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
	resp, r, err := apiClient.PromptsCatalogAPI.PromptsPromptCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsCatalogAPI.PromptsPromptCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsPromptCatalog`: PromptsPromptCatalog200Response
	fmt.Fprintf(os.Stdout, "Response from `PromptsCatalogAPI.PromptsPromptCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsPromptCatalogRequest struct via the builder pattern


### Return type

[**PromptsPromptCatalog200Response**](PromptsPromptCatalog200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

