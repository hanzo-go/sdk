# \PromptsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePromptsByName**](PromptsAPI.md#DeletePromptsByName) | **Delete** /v1/prompts/{name} | Delete removes one of the caller org&#39;s prompts and every version of it, answering 204.
[**GetPrompts**](PromptsAPI.md#GetPrompts) | **Get** /v1/prompts | List returns the caller org&#39;s prompt library as one row per prompt: its name, type, every version number it has, its taxonomy and when it last changed.
[**GetPromptsByName**](PromptsAPI.md#GetPromptsByName) | **Get** /v1/prompts/{name} | Get returns one of the caller org&#39;s prompts: its CURRENT template text plus the metadata of every version it has had.
[**GetPromptsCatalog**](PromptsAPI.md#GetPromptsCatalog) | **Get** /v1/prompts/catalog | Catalog returns the read-only starter prompt library shipped with the binary — reference content every tenant sees the same, NOT the caller&#39;s own prompts and never mixed into them.
[**GetPromptsMetrics**](PromptsAPI.md#GetPromptsMetrics) | **Get** /v1/prompts/metrics | Metrics returns real per-prompt statistics for the caller&#39;s org: how many versions each prompt has, which one is current, and when it was created and last changed.
[**PostPrompts**](PromptsAPI.md#PostPrompts) | **Post** /v1/prompts | Create records a prompt for the caller&#39;s org and answers 201 with it.



## DeletePromptsByName

> DeletePromptsByName(ctx, name).Execute()

Delete removes one of the caller org's prompts and every version of it, answering 204.



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
	name := "greeting" // string | Name is the prompt to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PromptsAPI.DeletePromptsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.DeletePromptsByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePromptsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPrompts

> PromptList GetPrompts(ctx).Execute()

List returns the caller org's prompt library as one row per prompt: its name, type, every version number it has, its taxonomy and when it last changed.



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
	resp, r, err := apiClient.PromptsAPI.GetPrompts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrompts`: PromptList
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPrompts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsRequest struct via the builder pattern


### Return type

[**PromptList**](PromptList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptsByName

> PromptDetail GetPromptsByName(ctx, name).Execute()

Get returns one of the caller org's prompts: its CURRENT template text plus the metadata of every version it has had.



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
	name := "greeting" // string | Name is the prompt to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.GetPromptsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptsByName`: PromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the prompt to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromptDetail**](PromptDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptsCatalog

> CatalogList GetPromptsCatalog(ctx).Execute()

Catalog returns the read-only starter prompt library shipped with the binary — reference content every tenant sees the same, NOT the caller's own prompts and never mixed into them.



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
	resp, r, err := apiClient.PromptsAPI.GetPromptsCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptsCatalog`: CatalogList
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptsCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsCatalogRequest struct via the builder pattern


### Return type

[**CatalogList**](CatalogList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptsMetrics

> MetricList GetPromptsMetrics(ctx).Execute()

Metrics returns real per-prompt statistics for the caller's org: how many versions each prompt has, which one is current, and when it was created and last changed.



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
	resp, r, err := apiClient.PromptsAPI.GetPromptsMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptsMetrics`: MetricList
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptsMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsMetricsRequest struct via the builder pattern


### Return type

[**MetricList**](MetricList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPrompts

> PromptDetail PostPrompts(ctx).PromptReq(promptReq).Execute()

Create records a prompt for the caller's org and answers 201 with it.



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
	promptReq := *openapiclient.NewPromptReq() // PromptReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.PostPrompts(context.Background()).PromptReq(promptReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.PostPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPrompts`: PromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.PostPrompts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptReq** | [**PromptReq**](PromptReq.md) |  | 

### Return type

[**PromptDetail**](PromptDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

