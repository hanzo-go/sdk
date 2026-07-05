# \CloudArticleAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddArticle**](CloudArticleAPIAPI.md#CloudApiControllerAddArticle) | **Post** /v1/cloud/add-article | Api Controller Add Article
[**CloudApiControllerDeleteArticle**](CloudArticleAPIAPI.md#CloudApiControllerDeleteArticle) | **Post** /v1/cloud/delete-article | Api Controller Delete Article
[**CloudApiControllerGetArticle**](CloudArticleAPIAPI.md#CloudApiControllerGetArticle) | **Get** /v1/cloud/get-article | Api Controller Get Article
[**CloudApiControllerGetArticles**](CloudArticleAPIAPI.md#CloudApiControllerGetArticles) | **Get** /v1/cloud/get-articles | Api Controller Get Articles
[**CloudApiControllerGetGlobalArticles**](CloudArticleAPIAPI.md#CloudApiControllerGetGlobalArticles) | **Get** /v1/cloud/get-global-articles | Api Controller Get Global Articles
[**CloudApiControllerUpdateArticle**](CloudArticleAPIAPI.md#CloudApiControllerUpdateArticle) | **Post** /v1/cloud/update-article | Api Controller Update Article



## CloudApiControllerAddArticle

> CloudControllersResponse CloudApiControllerAddArticle(ctx).CloudObjectArticle(cloudObjectArticle).Execute()

Api Controller Add Article



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
	cloudObjectArticle := *openapiclient.NewCloudObjectArticle() // CloudObjectArticle | The details of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudArticleAPIAPI.CloudApiControllerAddArticle(context.Background()).CloudObjectArticle(cloudObjectArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudArticleAPIAPI.CloudApiControllerAddArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddArticle`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudArticleAPIAPI.CloudApiControllerAddArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectArticle** | [**CloudObjectArticle**](CloudObjectArticle.md) | The details of the article | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteArticle

> CloudControllersResponse CloudApiControllerDeleteArticle(ctx).CloudObjectArticle(cloudObjectArticle).Execute()

Api Controller Delete Article



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
	cloudObjectArticle := *openapiclient.NewCloudObjectArticle() // CloudObjectArticle | The details of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudArticleAPIAPI.CloudApiControllerDeleteArticle(context.Background()).CloudObjectArticle(cloudObjectArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudArticleAPIAPI.CloudApiControllerDeleteArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteArticle`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudArticleAPIAPI.CloudApiControllerDeleteArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectArticle** | [**CloudObjectArticle**](CloudObjectArticle.md) | The details of the article | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetArticle

> CloudObjectArticle CloudApiControllerGetArticle(ctx).Id(id).Execute()

Api Controller Get Article



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
	id := "id_example" // string | The id (owner/name) of article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudArticleAPIAPI.CloudApiControllerGetArticle(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudArticleAPIAPI.CloudApiControllerGetArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetArticle`: CloudObjectArticle
	fmt.Fprintf(os.Stdout, "Response from `CloudArticleAPIAPI.CloudApiControllerGetArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of article | 

### Return type

[**CloudObjectArticle**](CloudObjectArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetArticles

> []CloudObjectArticle CloudApiControllerGetArticles(ctx).Owner(owner).Execute()

Api Controller Get Articles



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
	owner := "owner_example" // string | The owner of article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudArticleAPIAPI.CloudApiControllerGetArticles(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudArticleAPIAPI.CloudApiControllerGetArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetArticles`: []CloudObjectArticle
	fmt.Fprintf(os.Stdout, "Response from `CloudArticleAPIAPI.CloudApiControllerGetArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of article | 

### Return type

[**[]CloudObjectArticle**](CloudObjectArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalArticles

> []CloudObjectArticle CloudApiControllerGetGlobalArticles(ctx).Execute()

Api Controller Get Global Articles



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
	resp, r, err := apiClient.CloudArticleAPIAPI.CloudApiControllerGetGlobalArticles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudArticleAPIAPI.CloudApiControllerGetGlobalArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalArticles`: []CloudObjectArticle
	fmt.Fprintf(os.Stdout, "Response from `CloudArticleAPIAPI.CloudApiControllerGetGlobalArticles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalArticlesRequest struct via the builder pattern


### Return type

[**[]CloudObjectArticle**](CloudObjectArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateArticle

> CloudControllersResponse CloudApiControllerUpdateArticle(ctx).Id(id).CloudObjectArticle(cloudObjectArticle).Execute()

Api Controller Update Article



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
	cloudObjectArticle := *openapiclient.NewCloudObjectArticle() // CloudObjectArticle | The details of the article

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudArticleAPIAPI.CloudApiControllerUpdateArticle(context.Background()).Id(id).CloudObjectArticle(cloudObjectArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudArticleAPIAPI.CloudApiControllerUpdateArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateArticle`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudArticleAPIAPI.CloudApiControllerUpdateArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the article | 
 **cloudObjectArticle** | [**CloudObjectArticle**](CloudObjectArticle.md) | The details of the article | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

