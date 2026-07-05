# \CloudSearchAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProductControllerChatDocs**](CloudSearchAPIAPI.md#CloudProductControllerChatDocs) | **Post** /v1/chat-docs | 
[**CloudProductControllerIndexDocs**](CloudSearchAPIAPI.md#CloudProductControllerIndexDocs) | **Post** /v1/index-docs | 
[**CloudProductControllerSearchDocs**](CloudSearchAPIAPI.md#CloudProductControllerSearchDocs) | **Post** /v1/search-docs | 
[**CloudProductControllerSearchIndexes**](CloudSearchAPIAPI.md#CloudProductControllerSearchIndexes) | **Get** /v1/search-docs/indexes | 
[**CloudProductControllerSearchStats**](CloudSearchAPIAPI.md#CloudProductControllerSearchStats) | **Get** /v1/search-docs/stats | 



## CloudProductControllerChatDocs

> map[string]interface{} CloudProductControllerChatDocs(ctx).CloudProductControllerChatDocsRequest(cloudProductControllerChatDocsRequest).Store(store).Execute()





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
	cloudProductControllerChatDocsRequest := *openapiclient.NewCloudProductControllerChatDocsRequest([]openapiclient.CloudProductControllerChatDocsRequestMessagesInner{*openapiclient.NewCloudProductControllerChatDocsRequestMessagesInner()}) // CloudProductControllerChatDocsRequest | 
	store := "store_example" // string | Optional store/index selector (e.g. bot-docs). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudSearchAPIAPI.CloudProductControllerChatDocs(context.Background()).CloudProductControllerChatDocsRequest(cloudProductControllerChatDocsRequest).Store(store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSearchAPIAPI.CloudProductControllerChatDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProductControllerChatDocs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudSearchAPIAPI.CloudProductControllerChatDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudProductControllerChatDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProductControllerChatDocsRequest** | [**CloudProductControllerChatDocsRequest**](CloudProductControllerChatDocsRequest.md) |  | 
 **store** | **string** | Optional store/index selector (e.g. bot-docs). | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProductControllerIndexDocs

> map[string]interface{} CloudProductControllerIndexDocs(ctx).CloudProductControllerIndexDocsRequest(cloudProductControllerIndexDocsRequest).Store(store).Execute()





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
	cloudProductControllerIndexDocsRequest := *openapiclient.NewCloudProductControllerIndexDocsRequest([]map[string]interface{}{map[string]interface{}(123)}) // CloudProductControllerIndexDocsRequest | 
	store := "store_example" // string | Optional store/index selector (e.g. bot-docs). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudSearchAPIAPI.CloudProductControllerIndexDocs(context.Background()).CloudProductControllerIndexDocsRequest(cloudProductControllerIndexDocsRequest).Store(store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSearchAPIAPI.CloudProductControllerIndexDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProductControllerIndexDocs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudSearchAPIAPI.CloudProductControllerIndexDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudProductControllerIndexDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProductControllerIndexDocsRequest** | [**CloudProductControllerIndexDocsRequest**](CloudProductControllerIndexDocsRequest.md) |  | 
 **store** | **string** | Optional store/index selector (e.g. bot-docs). | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.CloudSearchAPIAPI.CloudProductControllerSearchDocs(context.Background()).CloudProductControllerSearchDocsRequest(cloudProductControllerSearchDocsRequest).Store(store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSearchAPIAPI.CloudProductControllerSearchDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProductControllerSearchDocs`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudSearchAPIAPI.CloudProductControllerSearchDocs`: %v\n", resp)
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


## CloudProductControllerSearchIndexes

> CloudProductControllerSearchIndexes200Response CloudProductControllerSearchIndexes(ctx).Execute()





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
	resp, r, err := apiClient.CloudSearchAPIAPI.CloudProductControllerSearchIndexes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSearchAPIAPI.CloudProductControllerSearchIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProductControllerSearchIndexes`: CloudProductControllerSearchIndexes200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudSearchAPIAPI.CloudProductControllerSearchIndexes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProductControllerSearchIndexesRequest struct via the builder pattern


### Return type

[**CloudProductControllerSearchIndexes200Response**](CloudProductControllerSearchIndexes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProductControllerSearchStats

> CloudProductControllerSearchStats200Response CloudProductControllerSearchStats(ctx).Execute()





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
	resp, r, err := apiClient.CloudSearchAPIAPI.CloudProductControllerSearchStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSearchAPIAPI.CloudProductControllerSearchStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProductControllerSearchStats`: CloudProductControllerSearchStats200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudSearchAPIAPI.CloudProductControllerSearchStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProductControllerSearchStatsRequest struct via the builder pattern


### Return type

[**CloudProductControllerSearchStats200Response**](CloudProductControllerSearchStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

