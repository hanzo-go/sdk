# \HelpAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHelpArticles**](HelpAPI.md#GetHelpArticles) | **Get** /v1/help/articles | Returns the public knowledge base: the help center&#39;s Published, publicly-visible articles as cards.
[**GetHelpArticlesBySlug**](HelpAPI.md#GetHelpArticlesBySlug) | **Get** /v1/help/articles/{slug} | Returns one public article by slug, with its body.
[**GetHelpCategories**](HelpAPI.md#GetHelpCategories) | **Get** /v1/help/categories | Returns the knowledge-base sections for the public center&#39;s navigation — but ONLY the sections that front at least one Published, public article, so an internal (agent-only) category name or description never leaks.
[**PostHelpTickets**](HelpAPI.md#PostHelpTickets) | **Post** /v1/help/tickets | Files a customer support ticket into the public help center.



## GetHelpArticles

> HelpArticleList GetHelpArticles(ctx).Category(category).Limit(limit).Execute()

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
	resp, r, err := apiClient.HelpAPI.GetHelpArticles(context.Background()).Category(category).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.GetHelpArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelpArticles`: HelpArticleList
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.GetHelpArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHelpArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Category narrows the list to one knowledge-base section, matched against the article&#39;s category by exact name. Empty lists every section. | 
 **limit** | **int32** | Limit caps how many articles are returned. Anything that is not a positive integer uses 50, and values above 200 are clamped to 200. | 

### Return type

[**HelpArticleList**](HelpArticleList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelpArticlesBySlug

> HelpArticle GetHelpArticlesBySlug(ctx, slug).Execute()

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
	resp, r, err := apiClient.HelpAPI.GetHelpArticlesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.GetHelpArticlesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelpArticlesBySlug`: HelpArticle
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.GetHelpArticlesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the article&#39;s public identifier, from the path. It IS the document name in the help center&#39;s store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelpArticlesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelpArticle**](HelpArticle.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelpCategories

> HelpCategoryList GetHelpCategories(ctx).Execute()

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
	resp, r, err := apiClient.HelpAPI.GetHelpCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.GetHelpCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelpCategories`: HelpCategoryList
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.GetHelpCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelpCategoriesRequest struct via the builder pattern


### Return type

[**HelpCategoryList**](HelpCategoryList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHelpTickets

> HelpTicketFiled PostHelpTickets(ctx).HelpTicketIntake(helpTicketIntake).Execute()

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
	helpTicketIntake := *openapiclient.NewHelpTicketIntake() // HelpTicketIntake | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpAPI.PostHelpTickets(context.Background()).HelpTicketIntake(helpTicketIntake).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpAPI.PostHelpTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHelpTickets`: HelpTicketFiled
	fmt.Fprintf(os.Stdout, "Response from `HelpAPI.PostHelpTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHelpTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helpTicketIntake** | [**HelpTicketIntake**](HelpTicketIntake.md) |  | 

### Return type

[**HelpTicketFiled**](HelpTicketFiled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

