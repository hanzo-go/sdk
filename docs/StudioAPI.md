# \StudioAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStudioBrandsById**](StudioAPI.md#DeleteStudioBrandsById) | **Delete** /v1/studio/brands/{id} | Deletes one of your brands.
[**DeleteStudioDocumentsById**](StudioAPI.md#DeleteStudioDocumentsById) | **Delete** /v1/studio/documents/{id} | Deletes one of your documents.
[**DeleteStudioFlowsById**](StudioAPI.md#DeleteStudioFlowsById) | **Delete** /v1/studio/flows/{id} | Deletes one of your flows.
[**DeleteStudioTemplatesById**](StudioAPI.md#DeleteStudioTemplatesById) | **Delete** /v1/studio/templates/{id} | Deletes one of your templates.
[**DeleteStudioThreadsById**](StudioAPI.md#DeleteStudioThreadsById) | **Delete** /v1/studio/threads/{id} | Deletes one of your threads.
[**GetStudioBrands**](StudioAPI.md#GetStudioBrands) | **Get** /v1/studio/brands | Lists the brands you may read — palettes and type systems as design tokens.
[**GetStudioBrandsById**](StudioAPI.md#GetStudioBrandsById) | **Get** /v1/studio/brands/{id} | Reads one brand you may read.
[**GetStudioDocuments**](StudioAPI.md#GetStudioDocuments) | **Get** /v1/studio/documents | Lists the documents and decks you may read, most recently changed first.
[**GetStudioDocumentsById**](StudioAPI.md#GetStudioDocumentsById) | **Get** /v1/studio/documents/{id} | Reads one document you may read.
[**GetStudioFlows**](StudioAPI.md#GetStudioFlows) | **Get** /v1/studio/flows | Lists the flows you may read, most recently changed first.
[**GetStudioFlowsById**](StudioAPI.md#GetStudioFlowsById) | **Get** /v1/studio/flows/{id} | Reads one flow you may read.
[**GetStudioGpus**](StudioAPI.md#GetStudioGpus) | **Get** /v1/studio/gpus | Lists the GPUs your org has linked with &#x60;hanzo link&#x60; that can take a studio.render step, each with whether it is free now.
[**GetStudioRuns**](StudioAPI.md#GetStudioRuns) | **Get** /v1/studio/runs | Lists the runs you may read — your own and those your org shares — newest first.
[**GetStudioRunsById**](StudioAPI.md#GetStudioRunsById) | **Get** /v1/studio/runs/{id} | Reads one run: its graph, its input, every step&#39;s state, the assets it made and what it cost.
[**GetStudioRunsByIdEvents**](StudioAPI.md#GetStudioRunsByIdEvents) | **Get** /v1/studio/runs/{id}/events | Stream a run&#39;s progress as Server-Sent Events
[**GetStudioRunsByIdVersions**](StudioAPI.md#GetStudioRunsByIdVersions) | **Get** /v1/studio/runs/{id}/versions | Lists every version of a creation you may read — the runs in this run&#39;s lineage — oldest (version 1) first.
[**GetStudioTemplates**](StudioAPI.md#GetStudioTemplates) | **Get** /v1/studio/templates | Lists templates: the public catalog (org empty) first, then your org&#39;s own that you may read.
[**GetStudioTemplatesById**](StudioAPI.md#GetStudioTemplatesById) | **Get** /v1/studio/templates/{id} | Reads one template: a catalog slug or one of your org&#39;s templates you may read.
[**GetStudioThreads**](StudioAPI.md#GetStudioThreads) | **Get** /v1/studio/threads | Lists your chat threads, most recent first.
[**GetStudioThreadsById**](StudioAPI.md#GetStudioThreadsById) | **Get** /v1/studio/threads/{id} | Reads one of your threads and its messages, oldest first.
[**PatchStudioBrandsById**](StudioAPI.md#PatchStudioBrandsById) | **Patch** /v1/studio/brands/{id} | Changes one of your brands; absent fields stay as they are.
[**PatchStudioDocumentsById**](StudioAPI.md#PatchStudioDocumentsById) | **Patch** /v1/studio/documents/{id} | Changes one of your documents; absent fields stay as they are.
[**PatchStudioFlowsById**](StudioAPI.md#PatchStudioFlowsById) | **Patch** /v1/studio/flows/{id} | Changes one of your flows; absent fields stay as they are.
[**PatchStudioRunsById**](StudioAPI.md#PatchStudioRunsById) | **Patch** /v1/studio/runs/{id} | Renames or shares one of your runs.
[**PatchStudioTemplatesById**](StudioAPI.md#PatchStudioTemplatesById) | **Patch** /v1/studio/templates/{id} | Changes one of your templates; absent fields stay as they are.
[**PostStudioBrands**](StudioAPI.md#PostStudioBrands) | **Post** /v1/studio/brands | Saves a brand: its tokens in the W3C design-token format (color, font, type, space, radius groups) and up to 16 logo items.
[**PostStudioChat**](StudioAPI.md#PostStudioChat) | **Post** /v1/studio/chat | One chat turn that plans and launches runs, streamed
[**PostStudioDocuments**](StudioAPI.md#PostStudioDocuments) | **Post** /v1/studio/documents | Writes a document (kind doc, its IR blocks) or a deck (kind deck, its slides).
[**PostStudioEstimate**](StudioAPI.md#PostStudioEstimate) | **Post** /v1/studio/estimate | Prices a run without starting it: the estimate POST /v1/studio/runs would reserve, step by step.
[**PostStudioFlows**](StudioAPI.md#PostStudioFlows) | **Post** /v1/studio/flows | Saves a graph as a flow of yours.
[**PostStudioRuns**](StudioAPI.md#PostStudioRuns) | **Post** /v1/studio/runs | Starts a run: name exactly one of template, flow or graph, with the input its inputs schema asks for.
[**PostStudioRunsByIdCancel**](StudioAPI.md#PostStudioRunsByIdCancel) | **Post** /v1/studio/runs/{id}/cancel | Cancels one of your runs: a step in flight is told to stop, nothing new starts, and the reservation settles to what finished steps cost.
[**PostStudioRunsByIdRetry**](StudioAPI.md#PostStudioRunsByIdRetry) | **Post** /v1/studio/runs/{id}/retry | Resumes one of your runs that failed or was cancelled: the steps that did not succeed run again in the same run, and the steps that did keep their outputs and are neither run nor charged again.
[**PostStudioTemplates**](StudioAPI.md#PostStudioTemplates) | **Post** /v1/studio/templates | Saves a template to your org: a graph and the inputs schema its runs are started with.



## DeleteStudioBrandsById

> DeleteStudioBrandsById(ctx, id).Execute()

Deletes one of your brands.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudioAPI.DeleteStudioBrandsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.DeleteStudioBrandsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudioBrandsByIdRequest struct via the builder pattern


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


## DeleteStudioDocumentsById

> DeleteStudioDocumentsById(ctx, id).Execute()

Deletes one of your documents.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudioAPI.DeleteStudioDocumentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.DeleteStudioDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudioDocumentsByIdRequest struct via the builder pattern


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


## DeleteStudioFlowsById

> DeleteStudioFlowsById(ctx, id).Execute()

Deletes one of your flows.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudioAPI.DeleteStudioFlowsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.DeleteStudioFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudioFlowsByIdRequest struct via the builder pattern


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


## DeleteStudioTemplatesById

> DeleteStudioTemplatesById(ctx, id).Execute()

Deletes one of your templates.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudioAPI.DeleteStudioTemplatesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.DeleteStudioTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudioTemplatesByIdRequest struct via the builder pattern


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


## DeleteStudioThreadsById

> DeleteStudioThreadsById(ctx, id).Execute()

Deletes one of your threads.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudioAPI.DeleteStudioThreadsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.DeleteStudioThreadsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStudioThreadsByIdRequest struct via the builder pattern


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


## GetStudioBrands

> StudioBrandPage GetStudioBrands(ctx).Limit(limit).Cursor(cursor).Execute()

Lists the brands you may read — palettes and type systems as design tokens.



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
	limit := int64(789) // int64 | Limit is how many rows one page holds, 1 to 100 (default 50). (optional)
	cursor := "cursor_example" // string | Cursor is the previous page's next. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioBrands(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioBrands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioBrands`: StudioBrandPage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit is how many rows one page holds, 1 to 100 (default 50). | 
 **cursor** | **string** | Cursor is the previous page&#39;s next. | 

### Return type

[**StudioBrandPage**](StudioBrandPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioBrandsById

> StudioBrand GetStudioBrandsById(ctx, id).Execute()

Reads one brand you may read.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioBrandsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioBrandsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioBrandsById`: StudioBrand
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioBrandsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioBrandsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioBrand**](StudioBrand.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioDocuments

> StudioDocumentPage GetStudioDocuments(ctx).Kind(kind).Limit(limit).Cursor(cursor).Execute()

Lists the documents and decks you may read, most recently changed first.



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
	kind := "kind_example" // string | Kind is doc or deck. (optional)
	limit := int64(789) // int64 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioDocuments(context.Background()).Kind(kind).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioDocuments`: StudioDocumentPage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind is doc or deck. | 
 **limit** | **int64** |  | 
 **cursor** | **string** |  | 

### Return type

[**StudioDocumentPage**](StudioDocumentPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioDocumentsById

> StudioDocument GetStudioDocumentsById(ctx, id).Execute()

Reads one document you may read.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioDocumentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioDocumentsById`: StudioDocument
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioDocument**](StudioDocument.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioFlows

> StudioFlowPage GetStudioFlows(ctx).Limit(limit).Cursor(cursor).Execute()

Lists the flows you may read, most recently changed first.



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
	limit := int64(789) // int64 | Limit is how many rows one page holds, 1 to 100 (default 50). (optional)
	cursor := "cursor_example" // string | Cursor is the previous page's next. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioFlows(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioFlows`: StudioFlowPage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit is how many rows one page holds, 1 to 100 (default 50). | 
 **cursor** | **string** | Cursor is the previous page&#39;s next. | 

### Return type

[**StudioFlowPage**](StudioFlowPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioFlowsById

> StudioFlow GetStudioFlowsById(ctx, id).Execute()

Reads one flow you may read.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioFlowsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioFlowsById`: StudioFlow
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioFlowsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioFlowsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioFlow**](StudioFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioGpus

> StudioGpuList GetStudioGpus(ctx).Limit(limit).Cursor(cursor).Execute()

Lists the GPUs your org has linked with `hanzo link` that can take a studio.render step, each with whether it is free now.



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
	limit := int64(789) // int64 | Limit is how many rows one page holds, 1 to 100 (default 50). (optional)
	cursor := "cursor_example" // string | Cursor is the previous page's next. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioGpus(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioGpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioGpus`: StudioGpuList
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioGpus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioGpusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit is how many rows one page holds, 1 to 100 (default 50). | 
 **cursor** | **string** | Cursor is the previous page&#39;s next. | 

### Return type

[**StudioGpuList**](StudioGpuList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioRuns

> StudioRunPage GetStudioRuns(ctx).Kind(kind).Status(status).Limit(limit).Cursor(cursor).Execute()

Lists the runs you may read — your own and those your org shares — newest first.



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
	kind := "kind_example" // string |  (optional)
	status := "status_example" // string | Status keeps only runs at this status. (optional)
	limit := int64(789) // int64 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioRuns(context.Background()).Kind(kind).Status(status).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioRuns`: StudioRunPage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 
 **status** | **string** | Status keeps only runs at this status. | 
 **limit** | **int64** |  | 
 **cursor** | **string** |  | 

### Return type

[**StudioRunPage**](StudioRunPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioRunsById

> StudioRun GetStudioRunsById(ctx, id).Execute()

Reads one run: its graph, its input, every step's state, the assets it made and what it cost.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioRunsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioRunsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioRunsById`: StudioRun
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioRunsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioRunsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioRun**](StudioRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioRunsByIdEvents

> *os.File GetStudioRunsByIdEvents(ctx, id).Execute()

Stream a run's progress as Server-Sent Events



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioRunsByIdEvents(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioRunsByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioRunsByIdEvents`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioRunsByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioRunsByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioRunsByIdVersions

> StudioRunPage GetStudioRunsByIdVersions(ctx, id).Execute()

Lists every version of a creation you may read — the runs in this run's lineage — oldest (version 1) first.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioRunsByIdVersions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioRunsByIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioRunsByIdVersions`: StudioRunPage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioRunsByIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioRunsByIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioRunPage**](StudioRunPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioTemplates

> StudioTemplatePage GetStudioTemplates(ctx).Kind(kind).Limit(limit).Cursor(cursor).Execute()

Lists templates: the public catalog (org empty) first, then your org's own that you may read.



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
	kind := "kind_example" // string | Kind keeps only templates that make this: image, poster, document, slides, type, model, ad, video, film, voice, music or other. (optional)
	limit := int64(789) // int64 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioTemplates(context.Background()).Kind(kind).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioTemplates`: StudioTemplatePage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind keeps only templates that make this: image, poster, document, slides, type, model, ad, video, film, voice, music or other. | 
 **limit** | **int64** |  | 
 **cursor** | **string** |  | 

### Return type

[**StudioTemplatePage**](StudioTemplatePage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioTemplatesById

> StudioTemplate GetStudioTemplatesById(ctx, id).Execute()

Reads one template: a catalog slug or one of your org's templates you may read.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioTemplatesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioTemplatesById`: StudioTemplate
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioTemplate**](StudioTemplate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioThreads

> StudioThreadPage GetStudioThreads(ctx).Limit(limit).Cursor(cursor).Execute()

Lists your chat threads, most recent first.



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
	limit := int64(789) // int64 | Limit is how many rows one page holds, 1 to 100 (default 50). (optional)
	cursor := "cursor_example" // string | Cursor is the previous page's next. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioThreads(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioThreads`: StudioThreadPage
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit is how many rows one page holds, 1 to 100 (default 50). | 
 **cursor** | **string** | Cursor is the previous page&#39;s next. | 

### Return type

[**StudioThreadPage**](StudioThreadPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioThreadsById

> StudioThread GetStudioThreadsById(ctx, id).Execute()

Reads one of your threads and its messages, oldest first.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.GetStudioThreadsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.GetStudioThreadsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioThreadsById`: StudioThread
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.GetStudioThreadsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioThreadsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioThread**](StudioThread.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStudioBrandsById

> StudioBrand PatchStudioBrandsById(ctx, id).StudioBrandPatch(studioBrandPatch).Execute()

Changes one of your brands; absent fields stay as they are.



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
	id := "id_example" // string | 
	studioBrandPatch := *openapiclient.NewStudioBrandPatch() // StudioBrandPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PatchStudioBrandsById(context.Background(), id).StudioBrandPatch(studioBrandPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PatchStudioBrandsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchStudioBrandsById`: StudioBrand
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PatchStudioBrandsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStudioBrandsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studioBrandPatch** | [**StudioBrandPatch**](StudioBrandPatch.md) |  | 

### Return type

[**StudioBrand**](StudioBrand.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStudioDocumentsById

> StudioDocument PatchStudioDocumentsById(ctx, id).StudioDocumentPatch(studioDocumentPatch).Execute()

Changes one of your documents; absent fields stay as they are.



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
	id := "id_example" // string | 
	studioDocumentPatch := *openapiclient.NewStudioDocumentPatch() // StudioDocumentPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PatchStudioDocumentsById(context.Background(), id).StudioDocumentPatch(studioDocumentPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PatchStudioDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchStudioDocumentsById`: StudioDocument
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PatchStudioDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStudioDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studioDocumentPatch** | [**StudioDocumentPatch**](StudioDocumentPatch.md) |  | 

### Return type

[**StudioDocument**](StudioDocument.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStudioFlowsById

> StudioFlow PatchStudioFlowsById(ctx, id).StudioFlowPatch(studioFlowPatch).Execute()

Changes one of your flows; absent fields stay as they are.



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
	id := "id_example" // string | 
	studioFlowPatch := *openapiclient.NewStudioFlowPatch() // StudioFlowPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PatchStudioFlowsById(context.Background(), id).StudioFlowPatch(studioFlowPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PatchStudioFlowsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchStudioFlowsById`: StudioFlow
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PatchStudioFlowsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStudioFlowsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studioFlowPatch** | [**StudioFlowPatch**](StudioFlowPatch.md) |  | 

### Return type

[**StudioFlow**](StudioFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStudioRunsById

> StudioRun PatchStudioRunsById(ctx, id).StudioRunPatch(studioRunPatch).Execute()

Renames or shares one of your runs.



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
	id := "id_example" // string | 
	studioRunPatch := *openapiclient.NewStudioRunPatch() // StudioRunPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PatchStudioRunsById(context.Background(), id).StudioRunPatch(studioRunPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PatchStudioRunsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchStudioRunsById`: StudioRun
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PatchStudioRunsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStudioRunsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studioRunPatch** | [**StudioRunPatch**](StudioRunPatch.md) |  | 

### Return type

[**StudioRun**](StudioRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStudioTemplatesById

> StudioTemplate PatchStudioTemplatesById(ctx, id).StudioTemplatePatch(studioTemplatePatch).Execute()

Changes one of your templates; absent fields stay as they are.



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
	id := "id_example" // string | 
	studioTemplatePatch := *openapiclient.NewStudioTemplatePatch() // StudioTemplatePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PatchStudioTemplatesById(context.Background(), id).StudioTemplatePatch(studioTemplatePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PatchStudioTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchStudioTemplatesById`: StudioTemplate
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PatchStudioTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStudioTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **studioTemplatePatch** | [**StudioTemplatePatch**](StudioTemplatePatch.md) |  | 

### Return type

[**StudioTemplate**](StudioTemplate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioBrands

> StudioBrand PostStudioBrands(ctx).StudioBrandWrite(studioBrandWrite).Execute()

Saves a brand: its tokens in the W3C design-token format (color, font, type, space, radius groups) and up to 16 logo items.



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
	studioBrandWrite := *openapiclient.NewStudioBrandWrite() // StudioBrandWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioBrands(context.Background()).StudioBrandWrite(studioBrandWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioBrands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioBrands`: StudioBrand
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studioBrandWrite** | [**StudioBrandWrite**](StudioBrandWrite.md) |  | 

### Return type

[**StudioBrand**](StudioBrand.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioChat

> *os.File PostStudioChat(ctx).ChatTurn(chatTurn).Execute()

One chat turn that plans and launches runs, streamed



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
	chatTurn := *openapiclient.NewChatTurn() // ChatTurn |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioChat(context.Background()).ChatTurn(chatTurn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioChat`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatTurn** | [**ChatTurn**](ChatTurn.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioDocuments

> StudioDocument PostStudioDocuments(ctx).StudioDocumentWrite(studioDocumentWrite).Execute()

Writes a document (kind doc, its IR blocks) or a deck (kind deck, its slides).



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
	studioDocumentWrite := *openapiclient.NewStudioDocumentWrite() // StudioDocumentWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioDocuments(context.Background()).StudioDocumentWrite(studioDocumentWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioDocuments`: StudioDocument
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studioDocumentWrite** | [**StudioDocumentWrite**](StudioDocumentWrite.md) |  | 

### Return type

[**StudioDocument**](StudioDocument.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioEstimate

> StudioEstimate PostStudioEstimate(ctx).StudioRunCreate(studioRunCreate).IdempotencyKey(idempotencyKey).Execute()

Prices a run without starting it: the estimate POST /v1/studio/runs would reserve, step by step.



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
	studioRunCreate := *openapiclient.NewStudioRunCreate() // StudioRunCreate | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioEstimate(context.Background()).StudioRunCreate(studioRunCreate).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioEstimate`: StudioEstimate
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioEstimate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studioRunCreate** | [**StudioRunCreate**](StudioRunCreate.md) |  | 
 **idempotencyKey** | **string** |  | 

### Return type

[**StudioEstimate**](StudioEstimate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioFlows

> StudioFlow PostStudioFlows(ctx).StudioFlowWrite(studioFlowWrite).Execute()

Saves a graph as a flow of yours.



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
	studioFlowWrite := *openapiclient.NewStudioFlowWrite() // StudioFlowWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioFlows(context.Background()).StudioFlowWrite(studioFlowWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioFlows`: StudioFlow
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studioFlowWrite** | [**StudioFlowWrite**](StudioFlowWrite.md) |  | 

### Return type

[**StudioFlow**](StudioFlow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioRuns

> StudioRun PostStudioRuns(ctx).StudioRunCreate(studioRunCreate).IdempotencyKey(idempotencyKey).Execute()

Starts a run: name exactly one of template, flow or graph, with the input its inputs schema asks for.



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
	studioRunCreate := *openapiclient.NewStudioRunCreate() // StudioRunCreate | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioRuns(context.Background()).StudioRunCreate(studioRunCreate).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioRuns`: StudioRun
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studioRunCreate** | [**StudioRunCreate**](StudioRunCreate.md) |  | 
 **idempotencyKey** | **string** |  | 

### Return type

[**StudioRun**](StudioRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioRunsByIdCancel

> StudioRun PostStudioRunsByIdCancel(ctx, id).Execute()

Cancels one of your runs: a step in flight is told to stop, nothing new starts, and the reservation settles to what finished steps cost.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioRunsByIdCancel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioRunsByIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioRunsByIdCancel`: StudioRun
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioRunsByIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioRunsByIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioRun**](StudioRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioRunsByIdRetry

> StudioRun PostStudioRunsByIdRetry(ctx, id).Execute()

Resumes one of your runs that failed or was cancelled: the steps that did not succeed run again in the same run, and the steps that did keep their outputs and are neither run nor charged again.



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
	id := "id_example" // string | ID is the resource's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioRunsByIdRetry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioRunsByIdRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioRunsByIdRetry`: StudioRun
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioRunsByIdRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioRunsByIdRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudioRun**](StudioRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStudioTemplates

> StudioTemplate PostStudioTemplates(ctx).StudioTemplateWrite(studioTemplateWrite).Execute()

Saves a template to your org: a graph and the inputs schema its runs are started with.



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
	studioTemplateWrite := *openapiclient.NewStudioTemplateWrite() // StudioTemplateWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudioAPI.PostStudioTemplates(context.Background()).StudioTemplateWrite(studioTemplateWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudioAPI.PostStudioTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStudioTemplates`: StudioTemplate
	fmt.Fprintf(os.Stdout, "Response from `StudioAPI.PostStudioTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStudioTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studioTemplateWrite** | [**StudioTemplateWrite**](StudioTemplateWrite.md) |  | 

### Return type

[**StudioTemplate**](StudioTemplate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

