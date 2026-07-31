# \HelpAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1HelpArticles**](HelpAPI.md#CloudGetV1HelpArticles) | **Get** /v1/help/articles | Returns the public knowledge base: the help center&#39;s Published, publicly-visible articles as cards.
[**CloudGetV1HelpArticlesSlug**](HelpAPI.md#CloudGetV1HelpArticlesSlug) | **Get** /v1/help/articles/{slug} | Returns one public article by slug, with its body.
[**CloudGetV1HelpCategories**](HelpAPI.md#CloudGetV1HelpCategories) | **Get** /v1/help/categories | Returns the knowledge-base sections for the public center&#39;s navigation — but ONLY the sections that front at least one Published, public article, so an internal (agent-only) category name or description never leaks.
[**CloudPostV1HelpTickets**](HelpAPI.md#CloudPostV1HelpTickets) | **Post** /v1/help/tickets | Files a customer support ticket into the public help center.



## CloudGetV1HelpArticles

> CloudHelpArticleList CloudGetV1HelpArticles(ctx).Category(category).Limit(limit).Execute()

Returns the public knowledge base: the help center's Published, publicly-visible articles as cards.



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
	category := "category_example" // string | Category narrows the list to one knowledge-base section, matched against the article's category by exact name. Empty lists every section. (optional)
	limit := int32(56) // int32 | Limit caps how many articles are returned. Anything that is not a positive integer uses 50, and values above 200 are clamped to 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpAPI.CloudGetV1HelpArticles(context.Background()).Category(category).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.CloudGetV1HelpArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1HelpArticles`: CloudHelpArticleList
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.CloudGetV1HelpArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1HelpArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Category narrows the list to one knowledge-base section, matched against the article&#39;s category by exact name. Empty lists every section. | 
 **limit** | **int32** | Limit caps how many articles are returned. Anything that is not a positive integer uses 50, and values above 200 are clamped to 200. | 

### Return type

[**CloudHelpArticleList**](CloudHelpArticleList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1HelpArticlesSlug

> CloudHelpArticle CloudGetV1HelpArticlesSlug(ctx, slug).Execute()

Returns one public article by slug, with its body.



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
	slug := "slug_example" // string | Slug is the article's public identifier, from the path. It IS the document name in the help center's store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpAPI.CloudGetV1HelpArticlesSlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.CloudGetV1HelpArticlesSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1HelpArticlesSlug`: CloudHelpArticle
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.CloudGetV1HelpArticlesSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the article&#39;s public identifier, from the path. It IS the document name in the help center&#39;s store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1HelpArticlesSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudHelpArticle**](CloudHelpArticle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1HelpCategories

> CloudHelpCategoryList CloudGetV1HelpCategories(ctx).Execute()

Returns the knowledge-base sections for the public center's navigation — but ONLY the sections that front at least one Published, public article, so an internal (agent-only) category name or description never leaks.



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
	resp, r, err := apiClient.HelpAPI.CloudGetV1HelpCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.CloudGetV1HelpCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1HelpCategories`: CloudHelpCategoryList
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.CloudGetV1HelpCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1HelpCategoriesRequest struct via the builder pattern


### Return type

[**CloudHelpCategoryList**](CloudHelpCategoryList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1HelpTickets

> CloudHelpTicketFiled CloudPostV1HelpTickets(ctx).CloudHelpTicketIntake(cloudHelpTicketIntake).Execute()

Files a customer support ticket into the public help center.



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
	cloudHelpTicketIntake := *openapiclient.NewCloudHelpTicketIntake() // CloudHelpTicketIntake | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpAPI.CloudPostV1HelpTickets(context.Background()).CloudHelpTicketIntake(cloudHelpTicketIntake).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.CloudPostV1HelpTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1HelpTickets`: CloudHelpTicketFiled
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.CloudPostV1HelpTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1HelpTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudHelpTicketIntake** | [**CloudHelpTicketIntake**](CloudHelpTicketIntake.md) |  | 

### Return type

[**CloudHelpTicketFiled**](CloudHelpTicketFiled.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

