# \PlatformAiAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformAiSuggest**](PlatformAiAPI.md#PlatformAiSuggest) | **Post** /v1/platform/ai/suggest | Get AI deployment suggestions



## PlatformAiSuggest

> PlatformTRPCResult PlatformAiSuggest(ctx).PlatformAiSuggestRequest(platformAiSuggestRequest).Execute()

Get AI deployment suggestions

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
	platformAiSuggestRequest := *openapiclient.NewPlatformAiSuggestRequest() // PlatformAiSuggestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAiAPI.PlatformAiSuggest(context.Background()).PlatformAiSuggestRequest(platformAiSuggestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAiAPI.PlatformAiSuggest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformAiSuggest`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformAiAPI.PlatformAiSuggest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformAiSuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformAiSuggestRequest** | [**PlatformAiSuggestRequest**](PlatformAiSuggestRequest.md) |  | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

