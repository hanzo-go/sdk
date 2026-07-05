# \NexusArticleAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddArticle**](NexusArticleAPIAPI.md#NexusAddArticle) | **Post** /v1/nexus/add-article | add Article
[**NexusDeleteArticle**](NexusArticleAPIAPI.md#NexusDeleteArticle) | **Post** /v1/nexus/delete-article | delete Article
[**NexusGetArticle**](NexusArticleAPIAPI.md#NexusGetArticle) | **Get** /v1/nexus/get-article | get Article
[**NexusGetArticles**](NexusArticleAPIAPI.md#NexusGetArticles) | **Get** /v1/nexus/get-articles | get Articles
[**NexusGetGlobalArticles**](NexusArticleAPIAPI.md#NexusGetGlobalArticles) | **Get** /v1/nexus/get-global-articles | get Global Articles
[**NexusUpdateArticle**](NexusArticleAPIAPI.md#NexusUpdateArticle) | **Post** /v1/nexus/update-article | update Article



## NexusAddArticle

> NexusResponse NexusAddArticle(ctx).NexusArticle(nexusArticle).Execute()

add Article



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
	nexusArticle := *openapiclient.NewNexusArticle() // NexusArticle | The details of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusArticleAPIAPI.NexusAddArticle(context.Background()).NexusArticle(nexusArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusArticleAPIAPI.NexusAddArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddArticle`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusArticleAPIAPI.NexusAddArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusArticle** | [**NexusArticle**](NexusArticle.md) | The details of the article | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteArticle

> NexusResponse NexusDeleteArticle(ctx).NexusArticle(nexusArticle).Execute()

delete Article



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
	nexusArticle := *openapiclient.NewNexusArticle() // NexusArticle | The details of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusArticleAPIAPI.NexusDeleteArticle(context.Background()).NexusArticle(nexusArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusArticleAPIAPI.NexusDeleteArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteArticle`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusArticleAPIAPI.NexusDeleteArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusArticle** | [**NexusArticle**](NexusArticle.md) | The details of the article | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetArticle

> NexusArticle NexusGetArticle(ctx).Id(id).Execute()

get Article



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
	id := "id_example" // string | The id (owner/name) of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusArticleAPIAPI.NexusGetArticle(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusArticleAPIAPI.NexusGetArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetArticle`: NexusArticle
	fmt.Fprintf(os.Stdout, "Response from `NexusArticleAPIAPI.NexusGetArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the article | 

### Return type

[**NexusArticle**](NexusArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetArticles

> []NexusArticle NexusGetArticles(ctx).Owner(owner).Execute()

get Articles



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
	owner := "owner_example" // string | The owner of the articles

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusArticleAPIAPI.NexusGetArticles(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusArticleAPIAPI.NexusGetArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetArticles`: []NexusArticle
	fmt.Fprintf(os.Stdout, "Response from `NexusArticleAPIAPI.NexusGetArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the articles | 

### Return type

[**[]NexusArticle**](NexusArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalArticles

> []NexusArticle NexusGetGlobalArticles(ctx).Execute()

get Global Articles



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
	resp, r, err := apiClient.NexusArticleAPIAPI.NexusGetGlobalArticles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusArticleAPIAPI.NexusGetGlobalArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalArticles`: []NexusArticle
	fmt.Fprintf(os.Stdout, "Response from `NexusArticleAPIAPI.NexusGetGlobalArticles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalArticlesRequest struct via the builder pattern


### Return type

[**[]NexusArticle**](NexusArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateArticle

> NexusResponse NexusUpdateArticle(ctx).Id(id).NexusArticle(nexusArticle).Execute()

update Article



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
	id := "id_example" // string | The id (owner/name) of the article
	nexusArticle := *openapiclient.NewNexusArticle() // NexusArticle | The details of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusArticleAPIAPI.NexusUpdateArticle(context.Background()).Id(id).NexusArticle(nexusArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusArticleAPIAPI.NexusUpdateArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateArticle`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusArticleAPIAPI.NexusUpdateArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the article | 
 **nexusArticle** | [**NexusArticle**](NexusArticle.md) | The details of the article | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

