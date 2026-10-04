# \CrawlAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPage**](CrawlAPI.md#ReadPage) | **Post** /v1/crawl | Fetch one URL and read it back as markdown



## ReadPage

> CrawlCrawlResult ReadPage(ctx).CrawlCrawlRequest(crawlCrawlRequest).Execute()

Fetch one URL and read it back as markdown



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
	crawlCrawlRequest := *openapiclient.NewCrawlCrawlRequest() // CrawlCrawlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlAPI.ReadPage(context.Background()).CrawlCrawlRequest(crawlCrawlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlAPI.ReadPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadPage`: CrawlCrawlResult
	fmt.Fprintf(os.Stdout, "Response from `CrawlAPI.ReadPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crawlCrawlRequest** | [**CrawlCrawlRequest**](CrawlCrawlRequest.md) |  | 

### Return type

[**CrawlCrawlResult**](CrawlCrawlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

