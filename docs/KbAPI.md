# \KbAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1KbConnectorsProvider**](KbAPI.md#CloudDeleteV1KbConnectorsProvider) | **Delete** /v1/kb/connectors/{provider} | DisconnectConnector revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, purges this provider&#39;s points from the org&#39;s vector namespace, and marks the connector disconnected.
[**CloudGetV1KbConnectors**](KbAPI.md#CloudGetV1KbConnectors) | **Get** /v1/kb/connectors | ListConnectors returns every supported knowledge connector with THIS org&#39;s connection state and the REAL number of documents each has ingested into the org&#39;s store.
[**CloudGetV1KbConnectorsCatalog**](KbAPI.md#CloudGetV1KbConnectorsCatalog) | **Get** /v1/kb/connectors/catalog | ListConnectorCatalog returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.
[**CloudGetV1KbConnectorsProviderCallback**](KbAPI.md#CloudGetV1KbConnectorsProviderCallback) | **Get** /v1/kb/connectors/{provider}/callback | CompleteConnectorOAuth finishes an OAuth connection: it exchanges the provider&#39;s code for a token, seals that token in KMS, and records the connection.
[**CloudGetV1KbConnectorsProviderConnect**](KbAPI.md#CloudGetV1KbConnectorsProviderConnect) | **Get** /v1/kb/connectors/{provider}/connect | StartConnectorOAuth returns the provider authorize URL the console opens to connect this org&#39;s account.
[**CloudGetV1KbGraph**](KbAPI.md#CloudGetV1KbGraph) | **Get** /v1/kb/graph | GetKnowledgeGraph returns the caller org&#39;s knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source&#39;s connector provenance as edges.
[**CloudPostV1KbConnectorsProviderSync**](KbAPI.md#CloudPostV1KbConnectorsProviderSync) | **Post** /v1/kb/connectors/{provider}/sync | SyncConnector pulls the provider&#39;s documents for the caller&#39;s org and files them as knowledge sources, which the store&#39;s own hook then indexes — so a synced document is retrievable exactly like a hand-written page.
[**CloudPostV1KbImport**](KbAPI.md#CloudPostV1KbImport) | **Post** /v1/kb/import | 
[**CloudPostV1KbSearch**](KbAPI.md#CloudPostV1KbSearch) | **Post** /v1/kb/search | SearchKnowledge runs a semantic search over the caller org&#39;s own knowledge — its wiki pages, its agent memories and everything its connectors have synced — and returns the matching passages.



## CloudDeleteV1KbConnectorsProvider

> CloudConnectionOut CloudDeleteV1KbConnectorsProvider(ctx, provider).Execute()

DisconnectConnector revokes a connection: it tombstones the stored credential so a later sync cannot reuse it, purges this provider's points from the org's vector namespace, and marks the connector disconnected.



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
	resp, r, err := apiClient.KbAPI.CloudDeleteV1KbConnectorsProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudDeleteV1KbConnectorsProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1KbConnectorsProvider`: CloudConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudDeleteV1KbConnectorsProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1KbConnectorsProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudConnectionOut**](CloudConnectionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1KbConnectors

> CloudKbConnectorsOut CloudGetV1KbConnectors(ctx).Execute()

ListConnectors returns every supported knowledge connector with THIS org's connection state and the REAL number of documents each has ingested into the org's store.



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
	resp, r, err := apiClient.KbAPI.CloudGetV1KbConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudGetV1KbConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1KbConnectors`: CloudKbConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudGetV1KbConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KbConnectorsRequest struct via the builder pattern


### Return type

[**CloudKbConnectorsOut**](CloudKbConnectorsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1KbConnectorsCatalog

> CloudCatalogOut CloudGetV1KbConnectorsCatalog(ctx).Execute()

ListConnectorCatalog returns the ONE catalog of everything a caller can connect: every first-party connector and every long-tail one, in a single list sorted by provider.



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
	resp, r, err := apiClient.KbAPI.CloudGetV1KbConnectorsCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudGetV1KbConnectorsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1KbConnectorsCatalog`: CloudCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudGetV1KbConnectorsCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KbConnectorsCatalogRequest struct via the builder pattern


### Return type

[**CloudCatalogOut**](CloudCatalogOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1KbConnectorsProviderCallback

> CloudConnectionOut CloudGetV1KbConnectorsProviderCallback(ctx, provider).Code(code).State(state).Error_(error_).Execute()

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
	resp, r, err := apiClient.KbAPI.CloudGetV1KbConnectorsProviderCallback(context.Background(), provider).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudGetV1KbConnectorsProviderCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1KbConnectorsProviderCallback`: CloudConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudGetV1KbConnectorsProviderCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector completing its flow, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KbConnectorsProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code is the provider&#39;s authorization code, exchanged for a token. | 
 **state** | **string** | State is the org-bound value this server signed at connect time. | 
 **error_** | **string** | Error is the provider&#39;s denial reason when the user refused consent. | 

### Return type

[**CloudConnectionOut**](CloudConnectionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1KbConnectorsProviderConnect

> CloudKbAuthorizeOut CloudGetV1KbConnectorsProviderConnect(ctx, provider).Execute()

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
	resp, r, err := apiClient.KbAPI.CloudGetV1KbConnectorsProviderConnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudGetV1KbConnectorsProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1KbConnectorsProviderConnect`: CloudKbAuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudGetV1KbConnectorsProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KbConnectorsProviderConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudKbAuthorizeOut**](CloudKbAuthorizeOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1KbGraph

> CloudGraphOut CloudGetV1KbGraph(ctx).Project(project).Execute()

GetKnowledgeGraph returns the caller org's knowledge as a node/edge graph shaped for a force-directed renderer: pages, memories and synced sources as nodes; the page parent tree, the wikilinks between pages, and each source's connector provenance as edges.



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
	resp, r, err := apiClient.KbAPI.CloudGetV1KbGraph(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudGetV1KbGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1KbGraph`: CloudGraphOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudGetV1KbGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KbGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows the graph to one project scope. Empty reads the whole org. | 

### Return type

[**CloudGraphOut**](CloudGraphOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1KbConnectorsProviderSync

> CloudKbSyncOut CloudPostV1KbConnectorsProviderSync(ctx, provider).Execute()

SyncConnector pulls the provider's documents for the caller's org and files them as knowledge sources, which the store's own hook then indexes — so a synced document is retrievable exactly like a hand-written page.



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
	resp, r, err := apiClient.KbAPI.CloudPostV1KbConnectorsProviderSync(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudPostV1KbConnectorsProviderSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1KbConnectorsProviderSync`: CloudKbSyncOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudPostV1KbConnectorsProviderSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector to act on: github, slack, google or notion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1KbConnectorsProviderSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudKbSyncOut**](CloudKbSyncOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1KbImport

> CloudPostV1KbImport(ctx).Execute()



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
	r, err := apiClient.KbAPI.CloudPostV1KbImport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudPostV1KbImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1KbImportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1KbSearch

> CloudSearchOut CloudPostV1KbSearch(ctx).CloudSearchIn(cloudSearchIn).Execute()

SearchKnowledge runs a semantic search over the caller org's own knowledge — its wiki pages, its agent memories and everything its connectors have synced — and returns the matching passages.



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
	cloudSearchIn := *openapiclient.NewCloudSearchIn() // CloudSearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbAPI.CloudPostV1KbSearch(context.Background()).CloudSearchIn(cloudSearchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbAPI.CloudPostV1KbSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1KbSearch`: CloudSearchOut
	fmt.Fprintf(os.Stdout, "Response from `KbAPI.CloudPostV1KbSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1KbSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSearchIn** | [**CloudSearchIn**](CloudSearchIn.md) |  | 

### Return type

[**CloudSearchOut**](CloudSearchOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

