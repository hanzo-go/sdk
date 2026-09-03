# \TranslateAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTranslateMemory**](TranslateAPI.md#GetTranslateMemory) | **Get** /v1/translate/memory | List returns the org&#39;s own translation-memory entries, newest first, optionally narrowed to one target language and/or one position on the review ladder.
[**PostTranslate**](TranslateAPI.md#PostTranslate) | **Post** /v1/translate | Translate a string or a batch into one target language
[**PutTranslateMemory**](TranslateAPI.md#PutTranslateMemory) | **Put** /v1/translate/memory | Review records a human decision on one translation-memory entry, and returns the entry as stored.



## GetTranslateMemory

> MemoryPage GetTranslateMemory(ctx).Target(target).State(state).Limit(limit).Execute()

List returns the org's own translation-memory entries, newest first, optionally narrowed to one target language and/or one position on the review ladder.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	target := "target_example" // string | Target narrows to one target language tag (BCP-47, e.g. \"es\" or \"pt-BR\"). (optional)
	state := "state_example" // string | State narrows to one position on the review ladder: machine, suggested, approved or published. (optional)
	limit := int64(789) // int64 | Limit caps the rows returned. Non-positive or unparseable means the server default (200); the ceiling is 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.GetTranslateMemory(context.Background()).Target(target).State(state).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.GetTranslateMemory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTranslateMemory`: MemoryPage
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.GetTranslateMemory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslateMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | **string** | Target narrows to one target language tag (BCP-47, e.g. \&quot;es\&quot; or \&quot;pt-BR\&quot;). | 
 **state** | **string** | State narrows to one position on the review ladder: machine, suggested, approved or published. | 
 **limit** | **int64** | Limit caps the rows returned. Non-positive or unparseable means the server default (200); the ceiling is 1000. | 

### Return type

[**MemoryPage**](MemoryPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTranslate

> PostTranslate(ctx).Execute()

Translate a string or a batch into one target language



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TranslateAPI.PostTranslate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.PostTranslate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostTranslateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTranslateMemory

> MemoryEntry PutTranslateMemory(ctx).ReviewRequest(reviewRequest).Execute()

Review records a human decision on one translation-memory entry, and returns the entry as stored.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	reviewRequest := *openapiclient.NewReviewRequest() // ReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.PutTranslateMemory(context.Background()).ReviewRequest(reviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.PutTranslateMemory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTranslateMemory`: MemoryEntry
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.PutTranslateMemory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTranslateMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewRequest** | [**ReviewRequest**](ReviewRequest.md) |  | 

### Return type

[**MemoryEntry**](MemoryEntry.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

