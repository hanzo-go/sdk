# \PromptsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1PromptsName**](PromptsAPI.md#CloudDeleteV1PromptsName) | **Delete** /v1/prompts/{name} | Delete removes one of the caller org&#39;s prompts and every version of it, answering 204.
[**CloudGetV1Prompts**](PromptsAPI.md#CloudGetV1Prompts) | **Get** /v1/prompts | List returns the caller org&#39;s prompt library as one row per prompt: its name, type, every version number it has, its taxonomy and when it last changed.
[**CloudGetV1PromptsCatalog**](PromptsAPI.md#CloudGetV1PromptsCatalog) | **Get** /v1/prompts/catalog | Catalog returns the read-only starter prompt library shipped with the binary — reference content every tenant sees the same, NOT the caller&#39;s own prompts and never mixed into them.
[**CloudGetV1PromptsMetrics**](PromptsAPI.md#CloudGetV1PromptsMetrics) | **Get** /v1/prompts/metrics | Metrics returns real per-prompt statistics for the caller&#39;s org: how many versions each prompt has, which one is current, and when it was created and last changed.
[**CloudGetV1PromptsName**](PromptsAPI.md#CloudGetV1PromptsName) | **Get** /v1/prompts/{name} | Get returns one of the caller org&#39;s prompts: its CURRENT template text plus the metadata of every version it has had.
[**CloudPostV1Prompts**](PromptsAPI.md#CloudPostV1Prompts) | **Post** /v1/prompts | Create records a prompt for the caller&#39;s org and answers 201 with it.



## CloudDeleteV1PromptsName

> CloudDeleteV1PromptsName(ctx, name).Execute()

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
	r, err := apiClient.PromptsAPI.CloudDeleteV1PromptsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CloudDeleteV1PromptsName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PromptsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1Prompts

> CloudPromptList CloudGetV1Prompts(ctx).Execute()

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
	resp, r, err := apiClient.PromptsAPI.CloudGetV1Prompts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CloudGetV1Prompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Prompts`: CloudPromptList
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CloudGetV1Prompts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PromptsRequest struct via the builder pattern


### Return type

[**CloudPromptList**](CloudPromptList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PromptsCatalog

> CloudCatalogList CloudGetV1PromptsCatalog(ctx).Execute()

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
	resp, r, err := apiClient.PromptsAPI.CloudGetV1PromptsCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CloudGetV1PromptsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PromptsCatalog`: CloudCatalogList
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CloudGetV1PromptsCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PromptsCatalogRequest struct via the builder pattern


### Return type

[**CloudCatalogList**](CloudCatalogList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PromptsMetrics

> CloudMetricList CloudGetV1PromptsMetrics(ctx).Execute()

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
	resp, r, err := apiClient.PromptsAPI.CloudGetV1PromptsMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CloudGetV1PromptsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PromptsMetrics`: CloudMetricList
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CloudGetV1PromptsMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PromptsMetricsRequest struct via the builder pattern


### Return type

[**CloudMetricList**](CloudMetricList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PromptsName

> CloudPromptDetail CloudGetV1PromptsName(ctx, name).Execute()

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
	resp, r, err := apiClient.PromptsAPI.CloudGetV1PromptsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CloudGetV1PromptsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PromptsName`: CloudPromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CloudGetV1PromptsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the prompt to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PromptsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPromptDetail**](CloudPromptDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Prompts

> CloudPromptDetail CloudPostV1Prompts(ctx).CloudPromptReq(cloudPromptReq).Execute()

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
	cloudPromptReq := *openapiclient.NewCloudPromptReq() // CloudPromptReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.CloudPostV1Prompts(context.Background()).CloudPromptReq(cloudPromptReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CloudPostV1Prompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Prompts`: CloudPromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CloudPostV1Prompts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PromptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPromptReq** | [**CloudPromptReq**](CloudPromptReq.md) |  | 

### Return type

[**CloudPromptDetail**](CloudPromptDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

