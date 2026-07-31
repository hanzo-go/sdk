# \CrawlAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudPostV1Crawl**](CrawlAPI.md#CloudPostV1Crawl) | **Post** /v1/crawl | 



## CloudPostV1Crawl

> CloudCrawlResult CloudPostV1Crawl(ctx).CloudCrawlRequest(cloudCrawlRequest).Execute()



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
	cloudCrawlRequest := *openapiclient.NewCloudCrawlRequest() // CloudCrawlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlAPI.CloudPostV1Crawl(context.Background()).CloudCrawlRequest(cloudCrawlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlAPI.CloudPostV1Crawl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Crawl`: CloudCrawlResult
	fmt.Fprintf(os.Stdout, "Response from `CrawlAPI.CloudPostV1Crawl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CrawlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCrawlRequest** | [**CloudCrawlRequest**](CloudCrawlRequest.md) |  | 

### Return type

[**CloudCrawlResult**](CloudCrawlResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

