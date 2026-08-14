# \TemplatesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplatesBySlug**](TemplatesAPI.md#DeleteTemplatesBySlug) | **Delete** /v1/templates/{slug} | Deletes the caller org&#39;s OWN starter kit.
[**GetTemplates**](TemplatesAPI.md#GetTemplates) | **Get** /v1/templates | Lists the public starter-kit catalog plus, for a validated caller, that org&#39;s own private kits.
[**GetTemplatesBySlug**](TemplatesAPI.md#GetTemplatesBySlug) | **Get** /v1/templates/{slug} | Returns one starter kit: the caller org&#39;s own by that slug, else the public catalog&#39;s.
[**PostTemplates**](TemplatesAPI.md#PostTemplates) | **Post** /v1/templates | Creates a starter kit PRIVATE to the caller&#39;s org and answers 201 with the stored kit.
[**PutTemplatesBySlug**](TemplatesAPI.md#PutTemplatesBySlug) | **Put** /v1/templates/{slug} | Overwrites the caller org&#39;s OWN starter kit at the path slug, answering the stored kit.



## DeleteTemplatesBySlug

> DeleteTemplatesBySlug(ctx, slug).Execute()

Deletes the caller org's OWN starter kit.



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
	slug := "acme-portal" // string | Slug is the starter kit to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplatesAPI.DeleteTemplatesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteTemplatesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the starter kit to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> KitList GetTemplates(ctx).Execute()

Lists the public starter-kit catalog plus, for a validated caller, that org's own private kits.



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
	resp, r, err := apiClient.TemplatesAPI.GetTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplates`: KitList
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


### Return type

[**KitList**](KitList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesBySlug

> StarterKit GetTemplatesBySlug(ctx, slug).Execute()

Returns one starter kit: the caller org's own by that slug, else the public catalog's.



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
	slug := "folio" // string | Slug is the starter kit to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplatesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplatesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplatesBySlug`: StarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplatesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the starter kit to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StarterKit**](StarterKit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTemplates

> StarterKit PostTemplates(ctx).PublishKitIn(publishKitIn).Execute()

Creates a starter kit PRIVATE to the caller's org and answers 201 with the stored kit.



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
	publishKitIn := *openapiclient.NewPublishKitIn() // PublishKitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PostTemplates(context.Background()).PublishKitIn(publishKitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PostTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTemplates`: StarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PostTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishKitIn** | [**PublishKitIn**](PublishKitIn.md) |  | 

### Return type

[**StarterKit**](StarterKit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTemplatesBySlug

> StarterKit PutTemplatesBySlug(ctx, slug).ReplaceKitIn(replaceKitIn).Execute()

Overwrites the caller org's OWN starter kit at the path slug, answering the stored kit.



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
	slug := "acme-portal" // string | Slug is the kit to replace, from the path.
	replaceKitIn := *openapiclient.NewReplaceKitIn() // ReplaceKitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PutTemplatesBySlug(context.Background(), slug).ReplaceKitIn(replaceKitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PutTemplatesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTemplatesBySlug`: StarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PutTemplatesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the kit to replace, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTemplatesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceKitIn** | [**ReplaceKitIn**](ReplaceKitIn.md) |  | 

### Return type

[**StarterKit**](StarterKit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

