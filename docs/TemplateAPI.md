# \TemplateAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplateBySlug**](TemplateAPI.md#DeleteTemplateBySlug) | **Delete** /v1/template/{slug} | Deletes the caller org&#39;s OWN starter kit.
[**GetTemplate**](TemplateAPI.md#GetTemplate) | **Get** /v1/template | Lists the public starter-kit catalog plus, for a validated caller, that org&#39;s own private kits.
[**GetTemplateBySlug**](TemplateAPI.md#GetTemplateBySlug) | **Get** /v1/template/{slug} | Returns one starter kit: the caller org&#39;s own by that slug, else the public catalog&#39;s.
[**PostTemplate**](TemplateAPI.md#PostTemplate) | **Post** /v1/template | Creates a starter kit PRIVATE to the caller&#39;s org and answers 201 with the stored kit.
[**PutTemplateBySlug**](TemplateAPI.md#PutTemplateBySlug) | **Put** /v1/template/{slug} | Overwrites the caller org&#39;s OWN starter kit at the path slug, answering the stored kit.



## DeleteTemplateBySlug

> DeleteTemplateBySlug(ctx, slug).Execute()

Deletes the caller org's OWN starter kit.



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
	slug := "acme-portal" // string | Slug is the starter kit to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplateAPI.DeleteTemplateBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.DeleteTemplateBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTemplateBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> TemplateKitList GetTemplate(ctx).Execute()

Lists the public starter-kit catalog plus, for a validated caller, that org's own private kits.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.GetTemplate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: TemplateKitList
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


### Return type

[**TemplateKitList**](TemplateKitList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateBySlug

> TemplateStarterKit GetTemplateBySlug(ctx, slug).Execute()

Returns one starter kit: the caller org's own by that slug, else the public catalog's.



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
	slug := "folio" // string | Slug is the starter kit to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.GetTemplateBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.GetTemplateBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateBySlug`: TemplateStarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.GetTemplateBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the starter kit to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateStarterKit**](TemplateStarterKit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTemplate

> TemplateStarterKit PostTemplate(ctx).TemplatePublishKitIn(templatePublishKitIn).Execute()

Creates a starter kit PRIVATE to the caller's org and answers 201 with the stored kit.



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
	templatePublishKitIn := *openapiclient.NewTemplatePublishKitIn() // TemplatePublishKitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.PostTemplate(context.Background()).TemplatePublishKitIn(templatePublishKitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.PostTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTemplate`: TemplateStarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.PostTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templatePublishKitIn** | [**TemplatePublishKitIn**](TemplatePublishKitIn.md) |  | 

### Return type

[**TemplateStarterKit**](TemplateStarterKit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTemplateBySlug

> TemplateStarterKit PutTemplateBySlug(ctx, slug).TemplateReplaceKitIn(templateReplaceKitIn).Execute()

Overwrites the caller org's OWN starter kit at the path slug, answering the stored kit.



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
	slug := "acme-portal" // string | Slug is the kit to replace, from the path.
	templateReplaceKitIn := *openapiclient.NewTemplateReplaceKitIn() // TemplateReplaceKitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAPI.PutTemplateBySlug(context.Background(), slug).TemplateReplaceKitIn(templateReplaceKitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.PutTemplateBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTemplateBySlug`: TemplateStarterKit
	fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.PutTemplateBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the kit to replace, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTemplateBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateReplaceKitIn** | [**TemplateReplaceKitIn**](TemplateReplaceKitIn.md) |  | 

### Return type

[**TemplateStarterKit**](TemplateStarterKit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

