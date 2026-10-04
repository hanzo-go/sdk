# \KnowledgeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKnowledgeConnectorsByProvider**](KnowledgeAPI.md#DeleteKnowledgeConnectorsByProvider) | **Delete** /v1/knowledge/connectors/{provider} | Revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, removes this provider&#39;s passages from the org&#39;s store, and marks the connector disconnected.
[**DeleteKnowledgeFilesById**](KnowledgeAPI.md#DeleteKnowledgeFilesById) | **Delete** /v1/knowledge/files/{id} | Removes one of the caller&#39;s org files from the index — its table of contents, passages, full-text rows, mentions and links — and its record.
[**GetKnowledgeConnectors**](KnowledgeAPI.md#GetKnowledgeConnectors) | **Get** /v1/knowledge/connectors | Returns every supported knowledge connector with THIS org&#39;s connection state and the REAL number of documents each has ingested into the org&#39;s store.
[**GetKnowledgeConnectorsByProviderCallback**](KnowledgeAPI.md#GetKnowledgeConnectorsByProviderCallback) | **Get** /v1/knowledge/connectors/{provider}/callback | CompleteConnectorOAuth finishes an OAuth connection: it exchanges the provider&#39;s code for a token, seals that token in KMS, and records the connection.
[**GetKnowledgeConnectorsByProviderConnect**](KnowledgeAPI.md#GetKnowledgeConnectorsByProviderConnect) | **Get** /v1/knowledge/connectors/{provider}/connect | StartConnectorOAuth returns the provider authorize URL the console opens to connect this org&#39;s account.
[**GetKnowledgeConnectorsCatalog**](KnowledgeAPI.md#GetKnowledgeConnectorsCatalog) | **Get** /v1/knowledge/connectors/catalog | Returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.
[**GetKnowledgeFiles**](KnowledgeAPI.md#GetKnowledgeFiles) | **Get** /v1/knowledge/files | Answers the caller&#39;s org files, most recently changed first, each with its status and stage — optionally only those in one bucket, which is how Drive shows a folder&#39;s files with their index state.
[**GetKnowledgeFilesById**](KnowledgeAPI.md#GetKnowledgeFilesById) | **Get** /v1/knowledge/files/{id} | Answers one of the caller&#39;s org files: what it is, where its bytes are, and how far its ingest has got.
[**GetKnowledgeFilesByIdGraph**](KnowledgeAPI.md#GetKnowledgeFilesByIdGraph) | **Get** /v1/knowledge/files/{id}/graph | Answers one of the caller&#39;s org files&#39; place in the org&#39;s graph: the files it links to and is linked from (a hyperlink, a citation by name), the files that name the same entities, and the entities it names.
[**GetKnowledgeFilesByIdSectionsBySection**](KnowledgeAPI.md#GetKnowledgeFilesByIdSectionsBySection) | **Get** /v1/knowledge/files/{id}/sections/{section} | Answers one section of one of the caller&#39;s org files: its own text, its subsections, the sections it links to or shares entities with — in this file or another — and the entities it names.
[**GetKnowledgeFilesByIdToc**](KnowledgeAPI.md#GetKnowledgeFilesByIdToc) | **Get** /v1/knowledge/files/{id}/toc | Answers one of the caller&#39;s org files&#39; table of contents: every section in document order with its depth, its parent and a one-line summary.
[**GetKnowledgeGraph**](KnowledgeAPI.md#GetKnowledgeGraph) | **Get** /v1/knowledge/graph | Returns the caller org&#39;s knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source&#39;s connector provenance as edges.
[**PostKnowledgeConnectorsByProviderSync**](KnowledgeAPI.md#PostKnowledgeConnectorsByProviderSync) | **Post** /v1/knowledge/connectors/{provider}/sync | Pulls the provider&#39;s documents for the caller&#39;s org and files them as knowledge sources, which the store&#39;s own hook then indexes — so a synced document is retrievable exactly like a hand-written page.
[**PostKnowledgeFiles**](KnowledgeAPI.md#PostKnowledgeFiles) | **Post** /v1/knowledge/files | Makes an object in one of the caller&#39;s org buckets a workspace file.
[**PostKnowledgeFilesRetrieve**](KnowledgeAPI.md#PostKnowledgeFilesRetrieve) | **Post** /v1/knowledge/files/retrieve | Grounds an answer in the caller&#39;s org files, table of contents first: candidate documents are found by searching their passages; a model reads their tables of contents — titles and one-line summaries — and picks the sections the answer is in; hybrid search drills into the passages of those sections; and the graph expands to the sections they link to or name the same entities as, in any file of the workspace.
[**PostKnowledgeFilesSearch**](KnowledgeAPI.md#PostKnowledgeFilesSearch) | **Post** /v1/knowledge/files/search | Answers the passages of the caller&#39;s org files that match a query, each citing its file › section › paragraph: the search behind Drive&#39;s box, and the first step an agent takes across a workspace.
[**PostKnowledgeImport**](KnowledgeAPI.md#PostKnowledgeImport) | **Post** /v1/knowledge/import | Import an Obsidian, Notion, Roam or Evernote export into the org&#39;s knowledge base
[**PostKnowledgeReindex**](KnowledgeAPI.md#PostKnowledgeReindex) | **Post** /v1/knowledge/reindex | Rebuilds the caller org&#39;s retrieval from its documents: every passage is removed, and every page, memory and source is cut into passages and embedded again with the configured model; the lexical index is reconciled to the same documents.
[**PostKnowledgeSearch**](KnowledgeAPI.md#PostKnowledgeSearch) | **Post** /v1/knowledge/search | Runs a hybrid search over the caller org&#39;s own knowledge — its wiki pages, its agent memories and everything its connectors have synced.



## DeleteKnowledgeConnectorsByProvider

> KnowledgeConnectionOut DeleteKnowledgeConnectorsByProvider(ctx, provider).Execute()

Revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, removes this provider's passages from the org's store, and marks the connector disconnected.



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
	provider := "provider_example" // string | Provider is the connector to act on: github, slack, google or notion.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.DeleteKnowledgeConnectorsByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.DeleteKnowledgeConnectorsByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKnowledgeConnectorsByProvider`: KnowledgeConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.DeleteKnowledgeConnectorsByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKnowledgeConnectorsByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeConnectionOut**](KnowledgeConnectionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKnowledgeFilesById

> KnowledgeForgotten DeleteKnowledgeFilesById(ctx, id).Execute()

Removes one of the caller's org files from the index — its table of contents, passages, full-text rows, mentions and links — and its record.



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
	id := "id_example" // string | ID is the file's id, as POST /v1/knowledge/files answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.DeleteKnowledgeFilesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.DeleteKnowledgeFilesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKnowledgeFilesById`: KnowledgeForgotten
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.DeleteKnowledgeFilesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the file&#39;s id, as POST /v1/knowledge/files answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKnowledgeFilesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeForgotten**](KnowledgeForgotten.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectors

> KnowledgeKbConnectorsOut GetKnowledgeConnectors(ctx).Execute()

Returns every supported knowledge connector with THIS org's connection state and the REAL number of documents each has ingested into the org's store.



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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectors`: KnowledgeKbConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeConnectorsRequest struct via the builder pattern


### Return type

[**KnowledgeKbConnectorsOut**](KnowledgeKbConnectorsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectorsByProviderCallback

> KnowledgeConnectionOut GetKnowledgeConnectorsByProviderCallback(ctx, provider).Code(code).State(state).Error_(error_).Execute()

CompleteConnectorOAuth finishes an OAuth connection: it exchanges the provider's code for a token, seals that token in KMS, and records the connection.



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
	provider := "provider_example" // string | Provider is the connector completing its flow, from the path.
	code := "code_example" // string | Code is the provider's authorization code, exchanged for a token. (optional)
	state := "state_example" // string | State is the org-bound value this server signed at connect time. (optional)
	error_ := "error__example" // string | Error is the provider's denial reason when the user refused consent. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectorsByProviderCallback(context.Background(), provider).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectorsByProviderCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectorsByProviderCallback`: KnowledgeConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeConnectorsByProviderCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector completing its flow, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeConnectorsByProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code is the provider&#39;s authorization code, exchanged for a token. | 
 **state** | **string** | State is the org-bound value this server signed at connect time. | 
 **error_** | **string** | Error is the provider&#39;s denial reason when the user refused consent. | 

### Return type

[**KnowledgeConnectionOut**](KnowledgeConnectionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectorsByProviderConnect

> KnowledgeKbAuthorizeOut GetKnowledgeConnectorsByProviderConnect(ctx, provider).Execute()

StartConnectorOAuth returns the provider authorize URL the console opens to connect this org's account.



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
	provider := "provider_example" // string | Provider is the connector to act on: github, slack, google or notion.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectorsByProviderConnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectorsByProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectorsByProviderConnect`: KnowledgeKbAuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeConnectorsByProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeConnectorsByProviderConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeKbAuthorizeOut**](KnowledgeKbAuthorizeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectorsCatalog

> KnowledgeCatalogOut GetKnowledgeConnectorsCatalog(ctx).Execute()

Returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.



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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectorsCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectorsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectorsCatalog`: KnowledgeCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeConnectorsCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeConnectorsCatalogRequest struct via the builder pattern


### Return type

[**KnowledgeCatalogOut**](KnowledgeCatalogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeFiles

> KnowledgeFilesOut GetKnowledgeFiles(ctx).Bucket(bucket).Limit(limit).Execute()

Answers the caller's org files, most recently changed first, each with its status and stage — optionally only those in one bucket, which is how Drive shows a folder's files with their index state.



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
	bucket := "bucket_example" // string |  (optional)
	limit := "limit_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeFiles(context.Background()).Bucket(bucket).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeFiles`: KnowledgeFilesOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string** |  | 
 **limit** | **string** |  | 

### Return type

[**KnowledgeFilesOut**](KnowledgeFilesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeFilesById

> KnowledgeFile GetKnowledgeFilesById(ctx, id).Execute()

Answers one of the caller's org files: what it is, where its bytes are, and how far its ingest has got.



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
	id := "id_example" // string | ID is the file's id, as POST /v1/knowledge/files answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeFilesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeFilesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeFilesById`: KnowledgeFile
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeFilesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the file&#39;s id, as POST /v1/knowledge/files answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeFilesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeFile**](KnowledgeFile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeFilesByIdGraph

> KnowledgeFileGraph GetKnowledgeFilesByIdGraph(ctx, id).Execute()

Answers one of the caller's org files' place in the org's graph: the files it links to and is linked from (a hyperlink, a citation by name), the files that name the same entities, and the entities it names.



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
	id := "id_example" // string | ID is the file's id, as POST /v1/knowledge/files answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeFilesByIdGraph(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeFilesByIdGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeFilesByIdGraph`: KnowledgeFileGraph
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeFilesByIdGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the file&#39;s id, as POST /v1/knowledge/files answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeFilesByIdGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeFileGraph**](KnowledgeFileGraph.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeFilesByIdSectionsBySection

> KnowledgeSectionOut GetKnowledgeFilesByIdSectionsBySection(ctx, id, section).Execute()

Answers one section of one of the caller's org files: its own text, its subsections, the sections it links to or shares entities with — in this file or another — and the entities it names.



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
	id := "id_example" // string | ID is the file's id.
	section := "section_example" // string | Section is the section's number in the file's table of contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeFilesByIdSectionsBySection(context.Background(), id, section).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeFilesByIdSectionsBySection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeFilesByIdSectionsBySection`: KnowledgeSectionOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeFilesByIdSectionsBySection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the file&#39;s id. | 
**section** | **string** | Section is the section&#39;s number in the file&#39;s table of contents. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeFilesByIdSectionsBySectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KnowledgeSectionOut**](KnowledgeSectionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeFilesByIdToc

> KnowledgeTocOut GetKnowledgeFilesByIdToc(ctx, id).Execute()

Answers one of the caller's org files' table of contents: every section in document order with its depth, its parent and a one-line summary.



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
	id := "id_example" // string | ID is the file's id, as POST /v1/knowledge/files answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeFilesByIdToc(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeFilesByIdToc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeFilesByIdToc`: KnowledgeTocOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeFilesByIdToc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the file&#39;s id, as POST /v1/knowledge/files answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeFilesByIdTocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeTocOut**](KnowledgeTocOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeGraph

> KnowledgeGraphOut GetKnowledgeGraph(ctx).Project(project).Execute()

Returns the caller org's knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source's connector provenance as edges.



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
	project := "project_example" // string | Project narrows the graph to one project scope. Empty reads the whole org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeGraph(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeGraph`: KnowledgeGraphOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows the graph to one project scope. Empty reads the whole org. | 

### Return type

[**KnowledgeGraphOut**](KnowledgeGraphOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeConnectorsByProviderSync

> KnowledgeKbSyncOut PostKnowledgeConnectorsByProviderSync(ctx, provider).Execute()

Pulls the provider's documents for the caller's org and files them as knowledge sources, which the store's own hook then indexes — so a synced document is retrievable exactly like a hand-written page.



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
	provider := "provider_example" // string | Provider is the connector to act on: github, slack, google or notion.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeConnectorsByProviderSync(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeConnectorsByProviderSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeConnectorsByProviderSync`: KnowledgeKbSyncOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeConnectorsByProviderSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeConnectorsByProviderSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeKbSyncOut**](KnowledgeKbSyncOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeFiles

> KnowledgeFile PostKnowledgeFiles(ctx).KnowledgeFileIn(knowledgeFileIn).Execute()

Makes an object in one of the caller's org buckets a workspace file.



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
	knowledgeFileIn := *openapiclient.NewKnowledgeFileIn() // KnowledgeFileIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeFiles(context.Background()).KnowledgeFileIn(knowledgeFileIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeFiles`: KnowledgeFile
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeFileIn** | [**KnowledgeFileIn**](KnowledgeFileIn.md) |  | 

### Return type

[**KnowledgeFile**](KnowledgeFile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeFilesRetrieve

> KnowledgeRetrieveOut PostKnowledgeFilesRetrieve(ctx).KnowledgeRetrieveIn(knowledgeRetrieveIn).Execute()

Grounds an answer in the caller's org files, table of contents first: candidate documents are found by searching their passages; a model reads their tables of contents — titles and one-line summaries — and picks the sections the answer is in; hybrid search drills into the passages of those sections; and the graph expands to the sections they link to or name the same entities as, in any file of the workspace.



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
	knowledgeRetrieveIn := *openapiclient.NewKnowledgeRetrieveIn() // KnowledgeRetrieveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeFilesRetrieve(context.Background()).KnowledgeRetrieveIn(knowledgeRetrieveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeFilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeFilesRetrieve`: KnowledgeRetrieveOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeFilesRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeFilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeRetrieveIn** | [**KnowledgeRetrieveIn**](KnowledgeRetrieveIn.md) |  | 

### Return type

[**KnowledgeRetrieveOut**](KnowledgeRetrieveOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeFilesSearch

> KnowledgeFileSearchOut PostKnowledgeFilesSearch(ctx).KnowledgeFileSearchIn(knowledgeFileSearchIn).Execute()

Answers the passages of the caller's org files that match a query, each citing its file › section › paragraph: the search behind Drive's box, and the first step an agent takes across a workspace.



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
	knowledgeFileSearchIn := *openapiclient.NewKnowledgeFileSearchIn() // KnowledgeFileSearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeFilesSearch(context.Background()).KnowledgeFileSearchIn(knowledgeFileSearchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeFilesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeFilesSearch`: KnowledgeFileSearchOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeFilesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeFilesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeFileSearchIn** | [**KnowledgeFileSearchIn**](KnowledgeFileSearchIn.md) |  | 

### Return type

[**KnowledgeFileSearchOut**](KnowledgeFileSearchOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeImport

> PostKnowledgeImport(ctx).Execute()

Import an Obsidian, Notion, Roam or Evernote export into the org's knowledge base



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
	r, err := apiClient.KnowledgeAPI.PostKnowledgeImport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeImportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeReindex

> KnowledgeReindexOut PostKnowledgeReindex(ctx).Execute()

Rebuilds the caller org's retrieval from its documents: every passage is removed, and every page, memory and source is cut into passages and embedded again with the configured model; the lexical index is reconciled to the same documents.



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
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeReindex(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeReindex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeReindex`: KnowledgeReindexOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeReindex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeReindexRequest struct via the builder pattern


### Return type

[**KnowledgeReindexOut**](KnowledgeReindexOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeSearch

> KnowledgeSearchOut PostKnowledgeSearch(ctx).KnowledgeSearchIn(knowledgeSearchIn).Execute()

Runs a hybrid search over the caller org's own knowledge — its wiki pages, its agent memories and everything its connectors have synced.



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
	knowledgeSearchIn := *openapiclient.NewKnowledgeSearchIn() // KnowledgeSearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeSearch(context.Background()).KnowledgeSearchIn(knowledgeSearchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeSearch`: KnowledgeSearchOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeSearchIn** | [**KnowledgeSearchIn**](KnowledgeSearchIn.md) |  | 

### Return type

[**KnowledgeSearchOut**](KnowledgeSearchOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

