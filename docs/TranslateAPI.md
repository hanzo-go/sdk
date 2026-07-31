# \TranslateAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1TranslateMemory**](TranslateAPI.md#CloudGetV1TranslateMemory) | **Get** /v1/translate/memory | List returns the org&#39;s own translation-memory entries, newest first, optionally narrowed to one target language and/or one position on the review ladder.
[**CloudPostV1Translate**](TranslateAPI.md#CloudPostV1Translate) | **Post** /v1/translate | 
[**CloudPutV1TranslateMemory**](TranslateAPI.md#CloudPutV1TranslateMemory) | **Put** /v1/translate/memory | Review records a human decision on one translation-memory entry, and returns the entry as stored.



## CloudGetV1TranslateMemory

> CloudMemoryPage CloudGetV1TranslateMemory(ctx).Target(target).State(state).Limit(limit).Execute()

List returns the org's own translation-memory entries, newest first, optionally narrowed to one target language and/or one position on the review ladder.



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
	target := "target_example" // string | Target narrows to one target language tag (BCP-47, e.g. \"es\" or \"pt-BR\"). (optional)
	state := "state_example" // string | State narrows to one position on the review ladder: machine, suggested, approved or published. (optional)
	limit := int32(56) // int32 | Limit caps the rows returned. Non-positive or unparseable means the server default (200); the ceiling is 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.CloudGetV1TranslateMemory(context.Background()).Target(target).State(state).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.CloudGetV1TranslateMemory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TranslateMemory`: CloudMemoryPage
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.CloudGetV1TranslateMemory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TranslateMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | **string** | Target narrows to one target language tag (BCP-47, e.g. \&quot;es\&quot; or \&quot;pt-BR\&quot;). | 
 **state** | **string** | State narrows to one position on the review ladder: machine, suggested, approved or published. | 
 **limit** | **int32** | Limit caps the rows returned. Non-positive or unparseable means the server default (200); the ceiling is 1000. | 

### Return type

[**CloudMemoryPage**](CloudMemoryPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Translate

> CloudPostV1Translate(ctx).Execute()



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
	r, err := apiClient.TranslateAPI.CloudPostV1Translate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.CloudPostV1Translate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TranslateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1TranslateMemory

> CloudMemoryEntry CloudPutV1TranslateMemory(ctx).CloudReviewRequest(cloudReviewRequest).Execute()

Review records a human decision on one translation-memory entry, and returns the entry as stored.



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
	cloudReviewRequest := *openapiclient.NewCloudReviewRequest() // CloudReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.CloudPutV1TranslateMemory(context.Background()).CloudReviewRequest(cloudReviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.CloudPutV1TranslateMemory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1TranslateMemory`: CloudMemoryEntry
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.CloudPutV1TranslateMemory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1TranslateMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudReviewRequest** | [**CloudReviewRequest**](CloudReviewRequest.md) |  | 

### Return type

[**CloudMemoryEntry**](CloudMemoryEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

