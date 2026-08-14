# \KbAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKbConnectorsByProvider**](KbAPI.md#DeleteKbConnectorsByProvider) | **Delete** /v1/kb/connectors/{provider} | Revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, purges this provider&#39;s points from the org&#39;s vector namespace, and marks the connector disconnected.
[**GetKbConnectors**](KbAPI.md#GetKbConnectors) | **Get** /v1/kb/connectors | Returns every supported knowledge connector with THIS org&#39;s connection state and the REAL number of documents each has ingested into the org&#39;s store.
[**GetKbConnectorsByProviderCallback**](KbAPI.md#GetKbConnectorsByProviderCallback) | **Get** /v1/kb/connectors/{provider}/callback | CompleteConnectorOAuth finishes an OAuth connection: it exchanges the provider&#39;s code for a token, seals that token in KMS, and records the connection.
[**GetKbConnectorsByProviderConnect**](KbAPI.md#GetKbConnectorsByProviderConnect) | **Get** /v1/kb/connectors/{provider}/connect | StartConnectorOAuth returns the provider authorize URL the console opens to connect this org&#39;s account.
[**GetKbConnectorsCatalog**](KbAPI.md#GetKbConnectorsCatalog) | **Get** /v1/kb/connectors/catalog | Returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.
[**GetKbGraph**](KbAPI.md#GetKbGraph) | **Get** /v1/kb/graph | Returns the caller org&#39;s knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source&#39;s connector provenance as edges.
[**PostKbConnectorsByProviderSync**](KbAPI.md#PostKbConnectorsByProviderSync) | **Post** /v1/kb/connectors/{provider}/sync | Pulls the provider&#39;s documents for the caller&#39;s org and files them as knowledge sources, which the store&#39;s own hook then indexes — so a synced document is retrievable exactly like a hand-written page.
[**PostKbImport**](KbAPI.md#PostKbImport) | **Post** /v1/kb/import | Import an Obsidian, Notion, Roam or Evernote export into the org&#39;s knowledge base
[**PostKbSearch**](KbAPI.md#PostKbSearch) | **Post** /v1/kb/search | Runs a semantic search over the caller org&#39;s own knowledge — its wiki pages, its agent memories and everything its connectors have synced — and returns the matching passages.



## DeleteKbConnectorsByProvider

> ConnectionOut DeleteKbConnectorsByProvider(ctx, provider).Execute()

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
	resp, r, err := apiClient.KbAPI.DeleteKbConnectorsByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.DeleteKbConnectorsByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKbConnectorsByProvider`: ConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.DeleteKbConnectorsByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKbConnectorsByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionOut**](ConnectionOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKbConnectors

> KbConnectorsOut GetKbConnectors(ctx).Execute()

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
	resp, r, err := apiClient.KbAPI.GetKbConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.GetKbConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKbConnectors`: KbConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.GetKbConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKbConnectorsRequest struct via the builder pattern


### Return type

[**KbConnectorsOut**](KbConnectorsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKbConnectorsByProviderCallback

> ConnectionOut GetKbConnectorsByProviderCallback(ctx, provider).Code(code).State(state).Error_(error_).Execute()

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
	resp, r, err := apiClient.KbAPI.GetKbConnectorsByProviderCallback(context.Background(), provider).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.GetKbConnectorsByProviderCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKbConnectorsByProviderCallback`: ConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.GetKbConnectorsByProviderCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector completing its flow, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKbConnectorsByProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code is the provider&#39;s authorization code, exchanged for a token. | 
 **state** | **string** | State is the org-bound value this server signed at connect time. | 
 **error_** | **string** | Error is the provider&#39;s denial reason when the user refused consent. | 

### Return type

[**ConnectionOut**](ConnectionOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKbConnectorsByProviderConnect

> KbAuthorizeOut GetKbConnectorsByProviderConnect(ctx, provider).Execute()

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
	resp, r, err := apiClient.KbAPI.GetKbConnectorsByProviderConnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.GetKbConnectorsByProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKbConnectorsByProviderConnect`: KbAuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.GetKbConnectorsByProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKbConnectorsByProviderConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KbAuthorizeOut**](KbAuthorizeOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKbConnectorsCatalog

> CatalogOut GetKbConnectorsCatalog(ctx).Execute()

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
	resp, r, err := apiClient.KbAPI.GetKbConnectorsCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.GetKbConnectorsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKbConnectorsCatalog`: CatalogOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.GetKbConnectorsCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKbConnectorsCatalogRequest struct via the builder pattern


### Return type

[**CatalogOut**](CatalogOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKbGraph

> GraphOut GetKbGraph(ctx).Project(project).Execute()

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
	resp, r, err := apiClient.KbAPI.GetKbGraph(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.GetKbGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKbGraph`: GraphOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.GetKbGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKbGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows the graph to one project scope. Empty reads the whole org. | 

### Return type

[**GraphOut**](GraphOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKbConnectorsByProviderSync

> KbSyncOut PostKbConnectorsByProviderSync(ctx, provider).Execute()

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
	resp, r, err := apiClient.KbAPI.PostKbConnectorsByProviderSync(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.PostKbConnectorsByProviderSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKbConnectorsByProviderSync`: KbSyncOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.PostKbConnectorsByProviderSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostKbConnectorsByProviderSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KbSyncOut**](KbSyncOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKbImport

> PostKbImport(ctx).Execute()

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
	r, err := apiClient.KbAPI.PostKbImport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.PostKbImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostKbImportRequest struct via the builder pattern


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


## PostKbSearch

> SearchOut PostKbSearch(ctx).SearchIn(searchIn).Execute()

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
	resp, r, err := apiClient.KbAPI.PostKbSearch(context.Background()).SearchIn(searchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.PostKbSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKbSearch`: SearchOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.PostKbSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKbSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchIn** | [**SearchIn**](SearchIn.md) |  | 

### Return type

[**SearchOut**](SearchOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

