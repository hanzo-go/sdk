# \AIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldClassifyBatch**](AIAPI.md#WorldWorldClassifyBatch) | **Post** /v1/world/classify-batch | Batch headline classification (per-user IAM token)
[**WorldWorldClassifyEvent**](AIAPI.md#WorldWorldClassifyEvent) | **Get** /v1/world/classify-event | Single event classification (per-user IAM token)
[**WorldWorldCountryIntel**](AIAPI.md#WorldWorldCountryIntel) | **Post** /v1/world/country-intel | AI country intelligence brief (per-user IAM token)
[**WorldWorldGroqSummarize**](AIAPI.md#WorldWorldGroqSummarize) | **Post** /v1/world/groq-summarize | World-brief summary via Hanzo inference (forwards the caller IAM token → org/project/billing; anon → skipped)
[**WorldWorldOpenrouterSummarize**](AIAPI.md#WorldWorldOpenrouterSummarize) | **Post** /v1/world/openrouter-summarize | Alt summary path via Hanzo inference (per-user IAM token)



## WorldWorldClassifyBatch

> map[string]interface{} WorldWorldClassifyBatch(ctx).WorldWorldClassifyBatchRequest(worldWorldClassifyBatchRequest).Execute()

Batch headline classification (per-user IAM token)

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
	worldWorldClassifyBatchRequest := *openapiclient.NewWorldWorldClassifyBatchRequest([]string{"Titles_example"}) // WorldWorldClassifyBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.WorldWorldClassifyBatch(context.Background()).WorldWorldClassifyBatchRequest(worldWorldClassifyBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.WorldWorldClassifyBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldClassifyBatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.WorldWorldClassifyBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldClassifyBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worldWorldClassifyBatchRequest** | [**WorldWorldClassifyBatchRequest**](WorldWorldClassifyBatchRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldClassifyEvent

> map[string]interface{} WorldWorldClassifyEvent(ctx).Title(title).Variant(variant).Execute()

Single event classification (per-user IAM token)

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
	title := "title_example" // string | Headline to classify.
	variant := "variant_example" // string | Only 'tech' is special-cased. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.WorldWorldClassifyEvent(context.Background()).Title(title).Variant(variant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.WorldWorldClassifyEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldClassifyEvent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.WorldWorldClassifyEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldClassifyEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Headline to classify. | 
 **variant** | **string** | Only &#39;tech&#39; is special-cased. | 

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


## WorldWorldCountryIntel

> map[string]interface{} WorldWorldCountryIntel(ctx).WorldWorldCountryIntelRequest(worldWorldCountryIntelRequest).Execute()

AI country intelligence brief (per-user IAM token)

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
	worldWorldCountryIntelRequest := *openapiclient.NewWorldWorldCountryIntelRequest("Country_example", "Code_example") // WorldWorldCountryIntelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.WorldWorldCountryIntel(context.Background()).WorldWorldCountryIntelRequest(worldWorldCountryIntelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.WorldWorldCountryIntel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCountryIntel`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.WorldWorldCountryIntel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCountryIntelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worldWorldCountryIntelRequest** | [**WorldWorldCountryIntelRequest**](WorldWorldCountryIntelRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldGroqSummarize

> map[string]interface{} WorldWorldGroqSummarize(ctx).WorldWorldSummarizeRequest(worldWorldSummarizeRequest).Execute()

World-brief summary via Hanzo inference (forwards the caller IAM token → org/project/billing; anon → skipped)

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
	worldWorldSummarizeRequest := *openapiclient.NewWorldWorldSummarizeRequest([]string{"Headlines_example"}) // WorldWorldSummarizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.WorldWorldGroqSummarize(context.Background()).WorldWorldSummarizeRequest(worldWorldSummarizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.WorldWorldGroqSummarize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldGroqSummarize`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.WorldWorldGroqSummarize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldGroqSummarizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worldWorldSummarizeRequest** | [**WorldWorldSummarizeRequest**](WorldWorldSummarizeRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldOpenrouterSummarize

> map[string]interface{} WorldWorldOpenrouterSummarize(ctx).WorldWorldSummarizeRequest(worldWorldSummarizeRequest).Execute()

Alt summary path via Hanzo inference (per-user IAM token)

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
	worldWorldSummarizeRequest := *openapiclient.NewWorldWorldSummarizeRequest([]string{"Headlines_example"}) // WorldWorldSummarizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.WorldWorldOpenrouterSummarize(context.Background()).WorldWorldSummarizeRequest(worldWorldSummarizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.WorldWorldOpenrouterSummarize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldOpenrouterSummarize`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.WorldWorldOpenrouterSummarize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldOpenrouterSummarizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worldWorldSummarizeRequest** | [**WorldWorldSummarizeRequest**](WorldWorldSummarizeRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

