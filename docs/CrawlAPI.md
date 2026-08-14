# \CrawlAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPage**](CrawlAPI.md#ReadPage) | **Post** /v1/crawl | Fetch one URL and read it back as markdown



## ReadPage

> CrawlResult ReadPage(ctx).CrawlRequest(crawlRequest).Execute()

Fetch one URL and read it back as markdown



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
	crawlRequest := *openapiclient.NewCrawlRequest() // CrawlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlAPI.ReadPage(context.Background()).CrawlRequest(crawlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlAPI.ReadPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadPage`: CrawlResult
	fmt.Fprintf(os.Stdout, "Response from `CrawlAPI.ReadPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crawlRequest** | [**CrawlRequest**](CrawlRequest.md) |  | 

### Return type

[**CrawlResult**](CrawlResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

