# \SearchAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProductControllerSearchDocs**](SearchAPIAPI.md#CloudProductControllerSearchDocs) | **Post** /v1/search-docs | 



## CloudProductControllerSearchDocs

> []map[string]interface{} CloudProductControllerSearchDocs(ctx).CloudProductControllerSearchDocsRequest(cloudProductControllerSearchDocsRequest).Store(store).Execute()





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
	cloudProductControllerSearchDocsRequest := *openapiclient.NewCloudProductControllerSearchDocsRequest("Query_example") // CloudProductControllerSearchDocsRequest | 
	store := "store_example" // string | Optional store/index selector (e.g. bot-docs). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPIAPI.CloudProductControllerSearchDocs(context.Background()).CloudProductControllerSearchDocsRequest(cloudProductControllerSearchDocsRequest).Store(store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPIAPI.CloudProductControllerSearchDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProductControllerSearchDocs`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SearchAPIAPI.CloudProductControllerSearchDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudProductControllerSearchDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProductControllerSearchDocsRequest** | [**CloudProductControllerSearchDocsRequest**](CloudProductControllerSearchDocsRequest.md) |  | 
 **store** | **string** | Optional store/index selector (e.g. bot-docs). | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

