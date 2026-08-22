# \KnowledgeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKnowledgeConnectorsByProvider**](KnowledgeAPI.md#DeleteKnowledgeConnectorsByProvider) | **Delete** /v1/knowledge/connectors/{provider} | Revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, purges this provider&#39;s points from the org&#39;s vector namespace, and marks the connector disconnected.
[**GetKnowledgeConnectors**](KnowledgeAPI.md#GetKnowledgeConnectors) | **Get** /v1/knowledge/connectors | Returns every supported knowledge connector with THIS org&#39;s connection state and the REAL number of documents each has ingested into the org&#39;s store.
[**GetKnowledgeConnectorsByProviderCallback**](KnowledgeAPI.md#GetKnowledgeConnectorsByProviderCallback) | **Get** /v1/knowledge/connectors/{provider}/callback | CompleteConnectorOAuth finishes an OAuth connection: it exchanges the provider&#39;s code for a token, seals that token in KMS, and records the connection.
[**GetKnowledgeConnectorsByProviderConnect**](KnowledgeAPI.md#GetKnowledgeConnectorsByProviderConnect) | **Get** /v1/knowledge/connectors/{provider}/connect | StartConnectorOAuth returns the provider authorize URL the console opens to connect this org&#39;s account.
[**GetKnowledgeConnectorsCatalog**](KnowledgeAPI.md#GetKnowledgeConnectorsCatalog) | **Get** /v1/knowledge/connectors/catalog | Returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.
[**GetKnowledgeGraph**](KnowledgeAPI.md#GetKnowledgeGraph) | **Get** /v1/knowledge/graph | Returns the caller org&#39;s knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source&#39;s connector provenance as edges.
[**PostKnowledgeConnectorsByProviderSync**](KnowledgeAPI.md#PostKnowledgeConnectorsByProviderSync) | **Post** /v1/knowledge/connectors/{provider}/sync | Pulls the provider&#39;s documents for the caller&#39;s org and files them as knowledge sources, which the store&#39;s own hook then indexes — so a synced document is retrievable exactly like a hand-written page.
[**PostKnowledgeImport**](KnowledgeAPI.md#PostKnowledgeImport) | **Post** /v1/knowledge/import | Import an Obsidian, Notion, Roam or Evernote export into the org&#39;s knowledge base
[**PostKnowledgeSearch**](KnowledgeAPI.md#PostKnowledgeSearch) | **Post** /v1/knowledge/search | Runs a semantic search over the caller org&#39;s own knowledge — its wiki pages, its agent memories and everything its connectors have synced — and returns the matching passages.



## DeleteKnowledgeConnectorsByProvider

> ConnectionOut DeleteKnowledgeConnectorsByProvider(ctx, provider).Execute()

Revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, purges this provider's points from the org's vector namespace, and marks the connector disconnected.



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
	provider := "provider_example" // string | Provider is the connector to act on: github, slack, google or notion.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.DeleteKnowledgeConnectorsByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.DeleteKnowledgeConnectorsByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKnowledgeConnectorsByProvider`: ConnectionOut
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

[**ConnectionOut**](ConnectionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectors

> KbConnectorsOut GetKnowledgeConnectors(ctx).Execute()

Returns every supported knowledge connector with THIS org's connection state and the REAL number of documents each has ingested into the org's store.



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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectors`: KbConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeConnectorsRequest struct via the builder pattern


### Return type

[**KbConnectorsOut**](KbConnectorsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectorsByProviderCallback

> ConnectionOut GetKnowledgeConnectorsByProviderCallback(ctx, provider).Code(code).State(state).Error_(error_).Execute()

CompleteConnectorOAuth finishes an OAuth connection: it exchanges the provider's code for a token, seals that token in KMS, and records the connection.



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
	// response from `GetKnowledgeConnectorsByProviderCallback`: ConnectionOut
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

[**ConnectionOut**](ConnectionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectorsByProviderConnect

> KbAuthorizeOut GetKnowledgeConnectorsByProviderConnect(ctx, provider).Execute()

StartConnectorOAuth returns the provider authorize URL the console opens to connect this org's account.



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
	provider := "provider_example" // string | Provider is the connector to act on: github, slack, google or notion.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectorsByProviderConnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectorsByProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectorsByProviderConnect`: KbAuthorizeOut
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

[**KbAuthorizeOut**](KbAuthorizeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeConnectorsCatalog

> CatalogOut GetKnowledgeConnectorsCatalog(ctx).Execute()

Returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.



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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeConnectorsCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeConnectorsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeConnectorsCatalog`: CatalogOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeConnectorsCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeConnectorsCatalogRequest struct via the builder pattern


### Return type

[**CatalogOut**](CatalogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeGraph

> GraphOut GetKnowledgeGraph(ctx).Project(project).Execute()

Returns the caller org's knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source's connector provenance as edges.



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
	project := "project_example" // string | Project narrows the graph to one project scope. Empty reads the whole org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeGraph(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeGraph`: GraphOut
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

[**GraphOut**](GraphOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKnowledgeConnectorsByProviderSync

> KbSyncOut PostKnowledgeConnectorsByProviderSync(ctx, provider).Execute()

Pulls the provider's documents for the caller's org and files them as knowledge sources, which the store's own hook then indexes — so a synced document is retrievable exactly like a hand-written page.



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
	provider := "provider_example" // string | Provider is the connector to act on: github, slack, google or notion.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeConnectorsByProviderSync(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeConnectorsByProviderSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeConnectorsByProviderSync`: KbSyncOut
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

[**KbSyncOut**](KbSyncOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk"
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


## PostKnowledgeSearch

> SearchOut PostKnowledgeSearch(ctx).SearchIn(searchIn).Execute()

Runs a semantic search over the caller org's own knowledge — its wiki pages, its agent memories and everything its connectors have synced — and returns the matching passages.



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
	searchIn := *openapiclient.NewSearchIn() // SearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.PostKnowledgeSearch(context.Background()).SearchIn(searchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.PostKnowledgeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKnowledgeSearch`: SearchOut
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.PostKnowledgeSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKnowledgeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchIn** | [**SearchIn**](SearchIn.md) |  | 

### Return type

[**SearchOut**](SearchOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

