# \WebsearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebsearchSearch**](WebsearchAPI.md#DeleteWebsearchSearch) | **Delete** /v1/websearch/search | Keyless web meta-search, in the SearXNG JSON envelope.
[**GetWebsearchSearch**](WebsearchAPI.md#GetWebsearchSearch) | **Get** /v1/websearch/search | Keyless web meta-search, in the SearXNG JSON envelope.
[**PatchWebsearchSearch**](WebsearchAPI.md#PatchWebsearchSearch) | **Patch** /v1/websearch/search | Keyless web meta-search, in the SearXNG JSON envelope.
[**PostWebsearchSearch**](WebsearchAPI.md#PostWebsearchSearch) | **Post** /v1/websearch/search | Keyless web meta-search, in the SearXNG JSON envelope.
[**PutWebsearchSearch**](WebsearchAPI.md#PutWebsearchSearch) | **Put** /v1/websearch/search | Keyless web meta-search, in the SearXNG JSON envelope.
[**SearchWeb**](WebsearchAPI.md#SearchWeb) | **Post** /v1/websearch | Search the live web



## DeleteWebsearchSearch

> DeleteWebsearchSearch(ctx).Execute()

Keyless web meta-search, in the SearXNG JSON envelope.



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
	r, err := apiClient.WebsearchAPI.DeleteWebsearchSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.DeleteWebsearchSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebsearchSearchRequest struct via the builder pattern


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


## GetWebsearchSearch

> GetWebsearchSearch(ctx).Execute()

Keyless web meta-search, in the SearXNG JSON envelope.



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
	r, err := apiClient.WebsearchAPI.GetWebsearchSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.GetWebsearchSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsearchSearchRequest struct via the builder pattern


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


## PatchWebsearchSearch

> PatchWebsearchSearch(ctx).Execute()

Keyless web meta-search, in the SearXNG JSON envelope.



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
	r, err := apiClient.WebsearchAPI.PatchWebsearchSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.PatchWebsearchSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWebsearchSearchRequest struct via the builder pattern


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


## PostWebsearchSearch

> PostWebsearchSearch(ctx).Execute()

Keyless web meta-search, in the SearXNG JSON envelope.



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
	r, err := apiClient.WebsearchAPI.PostWebsearchSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.PostWebsearchSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebsearchSearchRequest struct via the builder pattern


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


## PutWebsearchSearch

> PutWebsearchSearch(ctx).Execute()

Keyless web meta-search, in the SearXNG JSON envelope.



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
	r, err := apiClient.WebsearchAPI.PutWebsearchSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.PutWebsearchSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebsearchSearchRequest struct via the builder pattern


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


## SearchWeb

> WebSearchResults SearchWeb(ctx).WebSearchQuery(webSearchQuery).Execute()

Search the live web



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
	webSearchQuery := *openapiclient.NewWebSearchQuery() // WebSearchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsearchAPI.SearchWeb(context.Background()).WebSearchQuery(webSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.SearchWeb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchWeb`: WebSearchResults
	fmt.Fprintf(os.Stdout, "Response from `WebsearchAPI.SearchWeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webSearchQuery** | [**WebSearchQuery**](WebSearchQuery.md) |  | 

### Return type

[**WebSearchResults**](WebSearchResults.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

