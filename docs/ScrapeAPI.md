# \ScrapeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebsearchWebScrape**](ScrapeAPI.md#WebsearchWebScrape) | **Post** /v1/websearch/v1/scrape | Scrape a URL to markdown (Firecrawl response shape)
[**WebsearchWebScrapeBare**](ScrapeAPI.md#WebsearchWebScrapeBare) | **Post** /v1/websearch/scrape | Scrape a URL to markdown (bare alias of /v1/websearch/v1/scrape)



## WebsearchWebScrape

> WebsearchScrapeResponse WebsearchWebScrape(ctx).WebsearchScrapeRequest(websearchScrapeRequest).Execute()

Scrape a URL to markdown (Firecrawl response shape)

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
	websearchScrapeRequest := *openapiclient.NewWebsearchScrapeRequest("Url_example") // WebsearchScrapeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScrapeAPI.WebsearchWebScrape(context.Background()).WebsearchScrapeRequest(websearchScrapeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapeAPI.WebsearchWebScrape``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebsearchWebScrape`: WebsearchScrapeResponse
	fmt.Fprintf(os.Stdout, "Response from `ScrapeAPI.WebsearchWebScrape`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebsearchWebScrapeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websearchScrapeRequest** | [**WebsearchScrapeRequest**](WebsearchScrapeRequest.md) |  | 

### Return type

[**WebsearchScrapeResponse**](WebsearchScrapeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsearchWebScrapeBare

> WebsearchScrapeResponse WebsearchWebScrapeBare(ctx).WebsearchScrapeRequest(websearchScrapeRequest).Execute()

Scrape a URL to markdown (bare alias of /v1/websearch/v1/scrape)

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
	websearchScrapeRequest := *openapiclient.NewWebsearchScrapeRequest("Url_example") // WebsearchScrapeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScrapeAPI.WebsearchWebScrapeBare(context.Background()).WebsearchScrapeRequest(websearchScrapeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapeAPI.WebsearchWebScrapeBare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebsearchWebScrapeBare`: WebsearchScrapeResponse
	fmt.Fprintf(os.Stdout, "Response from `ScrapeAPI.WebsearchWebScrapeBare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebsearchWebScrapeBareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websearchScrapeRequest** | [**WebsearchScrapeRequest**](WebsearchScrapeRequest.md) |  | 

### Return type

[**WebsearchScrapeResponse**](WebsearchScrapeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

