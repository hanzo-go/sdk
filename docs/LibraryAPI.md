# \LibraryAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLibraryById**](LibraryAPI.md#DeleteLibraryById) | **Delete** /v1/library/{id} | Deletes one of your items and its bytes.
[**GetLibrary**](LibraryAPI.md#GetLibrary) | **Get** /v1/library | Lists what you made and what members of your organization shared with you, newest first.
[**GetLibraryById**](LibraryAPI.md#GetLibraryById) | **Get** /v1/library/{id} | Reads one item with fresh signed URLs.
[**GetLibraryUnowned**](LibraryAPI.md#GetLibraryUnowned) | **Get** /v1/library/unowned | ListUnowned reads, for a SuperAdmin alone, the organization&#39;s items that record no owner.
[**PatchLibraryById**](LibraryAPI.md#PatchLibraryById) | **Patch** /v1/library/{id} | Renames one of your items, shares it with your organization or stops sharing it, moves it to a project, or commits a reserved upload.
[**PostLibrary**](LibraryAPI.md#PostLibrary) | **Post** /v1/library | Keeps one thing you made in your Library: its bytes go to your own prefix of the Library&#39;s store, and it is listed to you alone until you share it.



## DeleteLibraryById

> DeleteLibraryById(ctx, id).Execute()

Deletes one of your items and its bytes.



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
	id := "lib_0f3c…" // string | ID is the item, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.DeleteLibraryById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.DeleteLibraryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the item, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLibraryByIdRequest struct via the builder pattern


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


## GetLibrary

> LibraryItemPage GetLibrary(ctx).Kind(kind).Source(source).Ref(ref).Project(project).Q(q).Since(since).Until(until).Scope(scope).Limit(limit).Cursor(cursor).Execute()

Lists what you made and what members of your organization shared with you, newest first.



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
	kind := "image" // string | Kind keeps one kind: image, video, audio, document, slides, type, model, code or file. (optional)
	source := "studio" // string | Source keeps what one product made: chat, dev, studio or workflow. (optional)
	ref := "ref_example" // string | Ref keeps what one conversation, run or Studio run made. (optional)
	project := "project_example" // string | Project keeps one project's items. (optional)
	q := "q_example" // string | Q keeps the items whose name contains it, ignoring case. (optional)
	since := "since_example" // string | Since and Until bound when it was made: RFC 3339, or a date (YYYY-MM-DD). (optional)
	until := "until_example" // string |  (optional)
	scope := "scope_example" // string | Scope is all (yours and what your org shared, the default), mine, or shared. (optional)
	limit := int64(24) // int64 | Limit bounds the page: default 60, at most 200. (optional)
	cursor := "cursor_example" // string | Cursor continues a list from the page before's `next`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetLibrary(context.Background()).Kind(kind).Source(source).Ref(ref).Project(project).Q(q).Since(since).Until(until).Scope(scope).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLibrary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibrary`: LibraryItemPage
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind keeps one kind: image, video, audio, document, slides, type, model, code or file. | 
 **source** | **string** | Source keeps what one product made: chat, dev, studio or workflow. | 
 **ref** | **string** | Ref keeps what one conversation, run or Studio run made. | 
 **project** | **string** | Project keeps one project&#39;s items. | 
 **q** | **string** | Q keeps the items whose name contains it, ignoring case. | 
 **since** | **string** | Since and Until bound when it was made: RFC 3339, or a date (YYYY-MM-DD). | 
 **until** | **string** |  | 
 **scope** | **string** | Scope is all (yours and what your org shared, the default), mine, or shared. | 
 **limit** | **int64** | Limit bounds the page: default 60, at most 200. | 
 **cursor** | **string** | Cursor continues a list from the page before&#39;s &#x60;next&#x60;. | 

### Return type

[**LibraryItemPage**](LibraryItemPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryById

> LibraryLibraryItem GetLibraryById(ctx, id).Execute()

Reads one item with fresh signed URLs.



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
	id := "lib_0f3c…" // string | ID is the item, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetLibraryById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLibraryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryById`: LibraryLibraryItem
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLibraryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the item, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryLibraryItem**](LibraryLibraryItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryUnowned

> LibraryItemPage GetLibraryUnowned(ctx).Limit(limit).Execute()

ListUnowned reads, for a SuperAdmin alone, the organization's items that record no owner.



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
	limit := int64(789) // int64 | Limit bounds the page: default 60, at most 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetLibraryUnowned(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLibraryUnowned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryUnowned`: LibraryItemPage
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLibraryUnowned`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryUnownedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit bounds the page: default 60, at most 200. | 

### Return type

[**LibraryItemPage**](LibraryItemPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLibraryById

> LibraryLibraryItem PatchLibraryById(ctx, id).LibraryItemPatch(libraryItemPatch).Execute()

Renames one of your items, shares it with your organization or stops sharing it, moves it to a project, or commits a reserved upload.



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
	id := "lib_0f3c…" // string | ID is the item, from the path.
	libraryItemPatch := *openapiclient.NewLibraryItemPatch() // LibraryItemPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.PatchLibraryById(context.Background(), id).LibraryItemPatch(libraryItemPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PatchLibraryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLibraryById`: LibraryLibraryItem
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.PatchLibraryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the item, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLibraryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libraryItemPatch** | [**LibraryItemPatch**](LibraryItemPatch.md) |  | 

### Return type

[**LibraryLibraryItem**](LibraryLibraryItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibrary

> LibraryLibraryItem PostLibrary(ctx).LibraryKeepIn(libraryKeepIn).Execute()

Keeps one thing you made in your Library: its bytes go to your own prefix of the Library's store, and it is listed to you alone until you share it.



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
	libraryKeepIn := *openapiclient.NewLibraryKeepIn() // LibraryKeepIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.PostLibrary(context.Background()).LibraryKeepIn(libraryKeepIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PostLibrary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLibrary`: LibraryLibraryItem
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.PostLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryKeepIn** | [**LibraryKeepIn**](LibraryKeepIn.md) |  | 

### Return type

[**LibraryLibraryItem**](LibraryLibraryItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

