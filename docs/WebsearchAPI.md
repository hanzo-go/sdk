# \WebsearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSearchEngineKey**](WebsearchAPI.md#DeleteSearchEngineKey) | **Delete** /v1/websearch/engines/{name}/key | Removes an engine&#39;s API key from KMS.
[**DeleteWebsearchSearch**](WebsearchAPI.md#DeleteWebsearchSearch) | **Delete** /v1/websearch/search | Web meta-search, in the SearXNG JSON envelope.
[**GetWebsearchSearch**](WebsearchAPI.md#GetWebsearchSearch) | **Get** /v1/websearch/search | Web meta-search, in the SearXNG JSON envelope.
[**ListSearchEngines**](WebsearchAPI.md#ListSearchEngines) | **Get** /v1/websearch/engines | Lists every engine web search can run and how the platform runs it.
[**PatchWebsearchSearch**](WebsearchAPI.md#PatchWebsearchSearch) | **Patch** /v1/websearch/search | Web meta-search, in the SearXNG JSON envelope.
[**PostWebsearchScrape**](WebsearchAPI.md#PostWebsearchScrape) | **Post** /v1/websearch/scrape | Fetch one page and get its extracted markdown, in the firecrawl envelope.
[**PostWebsearchSearch**](WebsearchAPI.md#PostWebsearchSearch) | **Post** /v1/websearch/search | Web meta-search, in the SearXNG JSON envelope.
[**PutWebsearchSearch**](WebsearchAPI.md#PutWebsearchSearch) | **Put** /v1/websearch/search | Web meta-search, in the SearXNG JSON envelope.
[**SearchWeb**](WebsearchAPI.md#SearchWeb) | **Post** /v1/websearch | Search the live web
[**SetSearchEngineKey**](WebsearchAPI.md#SetSearchEngineKey) | **Put** /v1/websearch/engines/{name}/key | Seals an engine&#39;s API key in KMS, replacing any key held.
[**SetSearchEngines**](WebsearchAPI.md#SetSearchEngines) | **Put** /v1/websearch/engines | Replaces the platform&#39;s web search engine setting and answers the result.



## DeleteSearchEngineKey

> WebsearchSearchEngineKeyed DeleteSearchEngineKey(ctx, name).Execute()

Removes an engine's API key from KMS.



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
	name := "name_example" // string | Name is the engine, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsearchAPI.DeleteSearchEngineKey(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.DeleteSearchEngineKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSearchEngineKey`: WebsearchSearchEngineKeyed
	fmt.Fprintf(os.Stdout, "Response from `WebsearchAPI.DeleteSearchEngineKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the engine, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSearchEngineKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebsearchSearchEngineKeyed**](WebsearchSearchEngineKeyed.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebsearchSearch

> DeleteWebsearchSearch(ctx).Execute()

Web meta-search, in the SearXNG JSON envelope.



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

Web meta-search, in the SearXNG JSON envelope.



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


## ListSearchEngines

> WebsearchSearchEngines ListSearchEngines(ctx).Execute()

Lists every engine web search can run and how the platform runs it.



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
	resp, r, err := apiClient.WebsearchAPI.ListSearchEngines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.ListSearchEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSearchEngines`: WebsearchSearchEngines
	fmt.Fprintf(os.Stdout, "Response from `WebsearchAPI.ListSearchEngines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSearchEnginesRequest struct via the builder pattern


### Return type

[**WebsearchSearchEngines**](WebsearchSearchEngines.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWebsearchSearch

> PatchWebsearchSearch(ctx).Execute()

Web meta-search, in the SearXNG JSON envelope.



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


## PostWebsearchScrape

> PostWebsearchScrape(ctx).Execute()

Fetch one page and get its extracted markdown, in the firecrawl envelope.



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
	r, err := apiClient.WebsearchAPI.PostWebsearchScrape(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.PostWebsearchScrape``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebsearchScrapeRequest struct via the builder pattern


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

Web meta-search, in the SearXNG JSON envelope.



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

Web meta-search, in the SearXNG JSON envelope.



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

> WebsearchWebSearchResults SearchWeb(ctx).WebsearchWebSearchQuery(websearchWebSearchQuery).Execute()

Search the live web



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
	websearchWebSearchQuery := *openapiclient.NewWebsearchWebSearchQuery() // WebsearchWebSearchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsearchAPI.SearchWeb(context.Background()).WebsearchWebSearchQuery(websearchWebSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.SearchWeb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchWeb`: WebsearchWebSearchResults
	fmt.Fprintf(os.Stdout, "Response from `WebsearchAPI.SearchWeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websearchWebSearchQuery** | [**WebsearchWebSearchQuery**](WebsearchWebSearchQuery.md) |  | 

### Return type

[**WebsearchWebSearchResults**](WebsearchWebSearchResults.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSearchEngineKey

> WebsearchSearchEngineKeyed SetSearchEngineKey(ctx, name).WebsearchSearchEngineKey(websearchSearchEngineKey).Execute()

Seals an engine's API key in KMS, replacing any key held.



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
	name := "name_example" // string | Name is the engine, from the path.
	websearchSearchEngineKey := *openapiclient.NewWebsearchSearchEngineKey() // WebsearchSearchEngineKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsearchAPI.SetSearchEngineKey(context.Background(), name).WebsearchSearchEngineKey(websearchSearchEngineKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.SetSearchEngineKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSearchEngineKey`: WebsearchSearchEngineKeyed
	fmt.Fprintf(os.Stdout, "Response from `WebsearchAPI.SetSearchEngineKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the engine, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSearchEngineKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **websearchSearchEngineKey** | [**WebsearchSearchEngineKey**](WebsearchSearchEngineKey.md) |  | 

### Return type

[**WebsearchSearchEngineKeyed**](WebsearchSearchEngineKeyed.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSearchEngines

> WebsearchSearchEngines SetSearchEngines(ctx).WebsearchSearchEnginesSet(websearchSearchEnginesSet).Execute()

Replaces the platform's web search engine setting and answers the result.



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
	websearchSearchEnginesSet := *openapiclient.NewWebsearchSearchEnginesSet() // WebsearchSearchEnginesSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsearchAPI.SetSearchEngines(context.Background()).WebsearchSearchEnginesSet(websearchSearchEnginesSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchAPI.SetSearchEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSearchEngines`: WebsearchSearchEngines
	fmt.Fprintf(os.Stdout, "Response from `WebsearchAPI.SetSearchEngines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSearchEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websearchSearchEnginesSet** | [**WebsearchSearchEnginesSet**](WebsearchSearchEnginesSet.md) |  | 

### Return type

[**WebsearchSearchEngines**](WebsearchSearchEngines.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

