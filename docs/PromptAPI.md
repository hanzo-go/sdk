# \PromptAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePromptByName**](PromptAPI.md#DeletePromptByName) | **Delete** /v1/prompt/{name} | Delete removes one of the caller org&#39;s prompts and every version of it, answering 204.
[**GetPrompt**](PromptAPI.md#GetPrompt) | **Get** /v1/prompt | List returns the caller org&#39;s prompt library as one row per prompt: its name, type, every version number it has, its taxonomy and when it last changed.
[**GetPromptByName**](PromptAPI.md#GetPromptByName) | **Get** /v1/prompt/{name} | Get returns one of the caller org&#39;s prompts: its CURRENT template text plus the metadata of every version it has had.
[**GetPromptCatalog**](PromptAPI.md#GetPromptCatalog) | **Get** /v1/prompt/catalog | Catalog returns the read-only starter prompt library shipped with the binary — reference content every tenant sees the same, NOT the caller&#39;s own prompts and never mixed into them.
[**GetPromptMetrics**](PromptAPI.md#GetPromptMetrics) | **Get** /v1/prompt/metrics | Metrics returns real per-prompt statistics for the caller&#39;s org: how many versions each prompt has, which one is current, and when it was created and last changed.
[**PostPrompt**](PromptAPI.md#PostPrompt) | **Post** /v1/prompt | Create records a prompt for the caller&#39;s org and answers 201 with it.



## DeletePromptByName

> DeletePromptByName(ctx, name).Execute()

Delete removes one of the caller org's prompts and every version of it, answering 204.



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
	name := "greeting" // string | Name is the prompt to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PromptAPI.DeletePromptByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptAPI.DeletePromptByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the prompt to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePromptByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPrompt

> PromptList GetPrompt(ctx).Execute()

List returns the caller org's prompt library as one row per prompt: its name, type, every version number it has, its taxonomy and when it last changed.



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
	resp, r, err := apiClient.PromptAPI.GetPrompt(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptAPI.GetPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrompt`: PromptList
	fmt.Fprintf(os.Stdout, "Response from `PromptAPI.GetPrompt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptRequest struct via the builder pattern


### Return type

[**PromptList**](PromptList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptByName

> PromptDetail GetPromptByName(ctx, name).Execute()

Get returns one of the caller org's prompts: its CURRENT template text plus the metadata of every version it has had.



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
	name := "greeting" // string | Name is the prompt to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptAPI.GetPromptByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptAPI.GetPromptByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptByName`: PromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptAPI.GetPromptByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the prompt to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromptDetail**](PromptDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptCatalog

> CatalogList GetPromptCatalog(ctx).Execute()

Catalog returns the read-only starter prompt library shipped with the binary — reference content every tenant sees the same, NOT the caller's own prompts and never mixed into them.



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
	resp, r, err := apiClient.PromptAPI.GetPromptCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptAPI.GetPromptCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptCatalog`: CatalogList
	fmt.Fprintf(os.Stdout, "Response from `PromptAPI.GetPromptCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptCatalogRequest struct via the builder pattern


### Return type

[**CatalogList**](CatalogList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptMetrics

> MetricList GetPromptMetrics(ctx).Execute()

Metrics returns real per-prompt statistics for the caller's org: how many versions each prompt has, which one is current, and when it was created and last changed.



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
	resp, r, err := apiClient.PromptAPI.GetPromptMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptAPI.GetPromptMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptMetrics`: MetricList
	fmt.Fprintf(os.Stdout, "Response from `PromptAPI.GetPromptMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptMetricsRequest struct via the builder pattern


### Return type

[**MetricList**](MetricList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPrompt

> PromptDetail PostPrompt(ctx).PromptReq(promptReq).Execute()

Create records a prompt for the caller's org and answers 201 with it.



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
	promptReq := *openapiclient.NewPromptReq() // PromptReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptAPI.PostPrompt(context.Background()).PromptReq(promptReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptAPI.PostPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPrompt`: PromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptAPI.PostPrompt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptReq** | [**PromptReq**](PromptReq.md) |  | 

### Return type

[**PromptDetail**](PromptDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

