# \DataroomAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataroomAnalyticsDataroomByDataroomid**](DataroomAPI.md#GetDataroomAnalyticsDataroomByDataroomid) | **Get** /v1/dataroom/analytics/dataroom/{dataroomId} | Rolls up every share link pointing at one data room: session and page-view totals for the room, plus the per-page breakdown for each link beneath it.
[**GetDataroomAnalyticsLinkByLinkid**](DataroomAPI.md#GetDataroomAnalyticsLinkByLinkid) | **Get** /v1/dataroom/analytics/link/{linkId} | Reports how one share link was actually read: total viewing sessions, total page views, and per page the view count, the summed dwell measure and its average.
[**GetDataroomDatarooms**](DataroomAPI.md#GetDataroomDatarooms) | **Get** /v1/dataroom/datarooms | Returns every data room in the caller org&#39;s own store, newest first, with its short public id, name, description and timestamps.
[**GetDataroomDataroomsById**](DataroomAPI.md#GetDataroomDataroomsById) | **Get** /v1/dataroom/datarooms/{id} | Reads one of the caller org&#39;s data rooms together with every document in it, each carrying its membership id and order index.
[**GetDataroomDocuments**](DataroomAPI.md#GetDataroomDocuments) | **Get** /v1/dataroom/documents | Returns every document in the caller org&#39;s own store, newest first — name, opaque storage key, content type, page count, size and timestamps.
[**GetDataroomDocumentsById**](DataroomAPI.md#GetDataroomDocumentsById) | **Get** /v1/dataroom/documents/{id} | Reads one of the caller org&#39;s documents — its name, opaque storage key, content type, page count, size and timestamps.
[**GetDataroomDocumentsByIdFile**](DataroomAPI.md#GetDataroomDocumentsByIdFile) | **Get** /v1/dataroom/documents/{id}/file | Download a document&#39;s bytes as its owner
[**GetDataroomHealth**](DataroomAPI.md#GetDataroomHealth) | **Get** /v1/dataroom/health | Reports that the data room subsystem is up.
[**GetDataroomLinks**](DataroomAPI.md#GetDataroomLinks) | **Get** /v1/dataroom/links | Returns every live share link in the caller org&#39;s own store, newest first, with the controls a visitor will meet: whether an address is required, whether a password is set, the allow and deny lists, whether download is permitted, and when the link expires.
[**GetDataroomTrust**](DataroomAPI.md#GetDataroomTrust) | **Get** /v1/dataroom/trust | Answers the caller org&#39;s OWN trust centre: its settings, every item it holds in both tiers, the requests waiting on it, and the grants it has made.
[**GetDataroomTrustCenterBySlug**](DataroomAPI.md#GetDataroomTrustCenterBySlug) | **Get** /v1/dataroom/trust/center/{slug} | Answers an org&#39;s public trust centre: its name, the text a party must accept to ask for a document, and every item it publishes.
[**GetDataroomTrustCenterBySlugFileByItem**](DataroomAPI.md#GetDataroomTrustCenterBySlugFileByItem) | **Get** /v1/dataroom/trust/center/{slug}/file/{item} | Read a public trust-centre item&#39;s bytes
[**GetDataroomViewByLinkid**](DataroomAPI.md#GetDataroomViewByLinkid) | **Get** /v1/dataroom/view/{linkId} | What a share link&#39;s visitor sees before authenticating
[**GetDataroomViewByLinkidDocumentByDocumentidFile**](DataroomAPI.md#GetDataroomViewByLinkidDocumentByDocumentidFile) | **Get** /v1/dataroom/view/{linkId}/document/{documentId}/file | Read a document&#39;s bytes as an authorised link visitor
[**PatchDataroomTrustArtifactsById**](DataroomAPI.md#PatchDataroomTrustArtifactsById) | **Patch** /v1/dataroom/trust/artifacts/{id} | Changes an item on the caller org&#39;s trust centre — replace its file with a newer edition, move it between public and gated, rewrite what it says, or retire it — and answers with the item as it now stands.
[**PostDataroomDatarooms**](DataroomAPI.md#PostDataroomDatarooms) | **Post** /v1/dataroom/datarooms | Opens a new data room for the caller org and answers with it, including the short public id it is addressed by.
[**PostDataroomDataroomsByIdDocuments**](DataroomAPI.md#PostDataroomDataroomsByIdDocuments) | **Post** /v1/dataroom/datarooms/{id}/documents | Puts an already-uploaded document into one of the caller org&#39;s data rooms and answers with the new membership id.
[**PostDataroomDocuments**](DataroomAPI.md#PostDataroomDocuments) | **Post** /v1/dataroom/documents | Upload a document&#39;s bytes and record it
[**PostDataroomLinks**](DataroomAPI.md#PostDataroomLinks) | **Post** /v1/dataroom/links | Grants access: it mints a public share link over one data room (&#x60;dataroomId&#x60;) or one document (&#x60;documentId&#x60;) — one of the two is required — and answers with the link, whose &#x60;id&#x60; is the token a visitor opens it with.
[**PostDataroomTrustArtifacts**](DataroomAPI.md#PostDataroomTrustArtifacts) | **Post** /v1/dataroom/trust/artifacts | Puts an item on the caller org&#39;s trust centre and answers with it.
[**PostDataroomTrustCenterBySlugRequests**](DataroomAPI.md#PostDataroomTrustCenterBySlugRequests) | **Post** /v1/dataroom/trust/center/{slug}/requests | Records a request to read what an independent auditor signed, and answers with its id.
[**PostDataroomTrustRequestsByIdGrant**](DataroomAPI.md#PostDataroomTrustRequestsByIdGrant) | **Post** /v1/dataroom/trust/requests/{id}/grant | Answers a request by opening access: it mints a share link over what was asked for, addressed to the address that asked and closing at expiry, records the decision, and mails the asker.
[**PostDataroomTrustRequestsByIdRefuse**](DataroomAPI.md#PostDataroomTrustRequestsByIdRefuse) | **Post** /v1/dataroom/trust/requests/{id}/refuse | Answers a request by declining it, recording who declined and why.
[**PostDataroomViewByLinkidAuthenticate**](DataroomAPI.md#PostDataroomViewByLinkidAuthenticate) | **Post** /v1/dataroom/view/{linkId}/authenticate | Pass a share link&#39;s gates and open a viewing session
[**PostDataroomViewByLinkidPageview**](DataroomAPI.md#PostDataroomViewByLinkidPageview) | **Post** /v1/dataroom/view/{linkId}/pageview | Record one page-view against an open viewing session
[**PutDataroomTrust**](DataroomAPI.md#PutDataroomTrust) | **Put** /v1/dataroom/trust | SetCenter opens, publishes or withdraws the caller org&#39;s trust centre and answers with the centre as it now stands.



## GetDataroomAnalyticsDataroomByDataroomid

> DataroomDataroomStats GetDataroomAnalyticsDataroomByDataroomid(ctx, dataroomId).Execute()

Rolls up every share link pointing at one data room: session and page-view totals for the room, plus the per-page breakdown for each link beneath it.



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
	dataroomId := "dataroomId_example" // string | DataroomID is the room to report on. It is the path segment, resolved in the caller's own tenant store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.GetDataroomAnalyticsDataroomByDataroomid(context.Background(), dataroomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomAnalyticsDataroomByDataroomid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomAnalyticsDataroomByDataroomid`: DataroomDataroomStats
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomAnalyticsDataroomByDataroomid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataroomId** | **string** | DataroomID is the room to report on. It is the path segment, resolved in the caller&#39;s own tenant store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomAnalyticsDataroomByDataroomidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataroomDataroomStats**](DataroomDataroomStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomAnalyticsLinkByLinkid

> DataroomDataroomLinkStats GetDataroomAnalyticsLinkByLinkid(ctx, linkId).Execute()

Reports how one share link was actually read: total viewing sessions, total page views, and per page the view count, the summed dwell measure and its average.



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
	linkId := "linkId_example" // string | LinkID is the link to report on. It is the path segment, resolved in the caller's own tenant store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.GetDataroomAnalyticsLinkByLinkid(context.Background(), linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomAnalyticsLinkByLinkid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomAnalyticsLinkByLinkid`: DataroomDataroomLinkStats
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomAnalyticsLinkByLinkid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** | LinkID is the link to report on. It is the path segment, resolved in the caller&#39;s own tenant store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomAnalyticsLinkByLinkidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataroomDataroomLinkStats**](DataroomDataroomLinkStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomDatarooms

> DataroomDataroomRooms GetDataroomDatarooms(ctx).Execute()

Returns every data room in the caller org's own store, newest first, with its short public id, name, description and timestamps.



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
	resp, r, err := apiClient.DataroomAPI.GetDataroomDatarooms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomDatarooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomDatarooms`: DataroomDataroomRooms
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomDatarooms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomDataroomsRequest struct via the builder pattern


### Return type

[**DataroomDataroomRooms**](DataroomDataroomRooms.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomDataroomsById

> DataroomDataroomRoomDetailOne GetDataroomDataroomsById(ctx, id).Execute()

Reads one of the caller org's data rooms together with every document in it, each carrying its membership id and order index.



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
	id := "id_example" // string | ID is the room to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.GetDataroomDataroomsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomDataroomsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomDataroomsById`: DataroomDataroomRoomDetailOne
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomDataroomsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomDataroomsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataroomDataroomRoomDetailOne**](DataroomDataroomRoomDetailOne.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomDocuments

> DataroomDataroomDocuments GetDataroomDocuments(ctx).Execute()

Returns every document in the caller org's own store, newest first — name, opaque storage key, content type, page count, size and timestamps.



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
	resp, r, err := apiClient.DataroomAPI.GetDataroomDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomDocuments`: DataroomDataroomDocuments
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomDocumentsRequest struct via the builder pattern


### Return type

[**DataroomDataroomDocuments**](DataroomDataroomDocuments.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomDocumentsById

> DataroomDataroomDocumentOne GetDataroomDocumentsById(ctx, id).Execute()

Reads one of the caller org's documents — its name, opaque storage key, content type, page count, size and timestamps.



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
	id := "id_example" // string | ID is the document to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.GetDataroomDocumentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomDocumentsById`: DataroomDataroomDocumentOne
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataroomDataroomDocumentOne**](DataroomDataroomDocumentOne.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomDocumentsByIdFile

> GetDataroomDocumentsByIdFile(ctx, id).Execute()

Download a document's bytes as its owner



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
	r, err := apiClient.DataroomAPI.GetDataroomDocumentsByIdFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomDocumentsByIdFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomDocumentsByIdFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDataroomHealth

> DataroomDataroomLiveness GetDataroomHealth(ctx).Execute()

Reports that the data room subsystem is up.



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
	resp, r, err := apiClient.DataroomAPI.GetDataroomHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomHealth`: DataroomDataroomLiveness
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomHealthRequest struct via the builder pattern


### Return type

[**DataroomDataroomLiveness**](DataroomDataroomLiveness.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomLinks

> DataroomDataroomLinks GetDataroomLinks(ctx).Execute()

Returns every live share link in the caller org's own store, newest first, with the controls a visitor will meet: whether an address is required, whether a password is set, the allow and deny lists, whether download is permitted, and when the link expires.



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
	resp, r, err := apiClient.DataroomAPI.GetDataroomLinks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomLinks`: DataroomDataroomLinks
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomLinksRequest struct via the builder pattern


### Return type

[**DataroomDataroomLinks**](DataroomDataroomLinks.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomTrust

> DataroomTrustDesk GetDataroomTrust(ctx).Execute()

Answers the caller org's OWN trust centre: its settings, every item it holds in both tiers, the requests waiting on it, and the grants it has made.



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
	resp, r, err := apiClient.DataroomAPI.GetDataroomTrust(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomTrust``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomTrust`: DataroomTrustDesk
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomTrust`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomTrustRequest struct via the builder pattern


### Return type

[**DataroomTrustDesk**](DataroomTrustDesk.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomTrustCenterBySlug

> DataroomTrustPage GetDataroomTrustCenterBySlug(ctx, slug).Execute()

Answers an org's public trust centre: its name, the text a party must accept to ask for a document, and every item it publishes.



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
	slug := "slug_example" // string | Slug is the centre's public address. It resolves only for an org that has published; anything else is not found, so this cannot be used to learn which orgs exist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.GetDataroomTrustCenterBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomTrustCenterBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataroomTrustCenterBySlug`: DataroomTrustPage
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.GetDataroomTrustCenterBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the centre&#39;s public address. It resolves only for an org that has published; anything else is not found, so this cannot be used to learn which orgs exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomTrustCenterBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataroomTrustPage**](DataroomTrustPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataroomTrustCenterBySlugFileByItem

> GetDataroomTrustCenterBySlugFileByItem(ctx, slug, item).Execute()

Read a public trust-centre item's bytes



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
	slug := "slug_example" // string | 
	item := "item_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataroomAPI.GetDataroomTrustCenterBySlugFileByItem(context.Background(), slug, item).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomTrustCenterBySlugFileByItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**item** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomTrustCenterBySlugFileByItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetDataroomViewByLinkid

> GetDataroomViewByLinkid(ctx, linkId).Execute()

What a share link's visitor sees before authenticating



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
	linkId := "linkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataroomAPI.GetDataroomViewByLinkid(context.Background(), linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomViewByLinkid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomViewByLinkidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDataroomViewByLinkidDocumentByDocumentidFile

> GetDataroomViewByLinkidDocumentByDocumentidFile(ctx, linkId, documentId).Execute()

Read a document's bytes as an authorised link visitor



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
	linkId := "linkId_example" // string | 
	documentId := "documentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataroomAPI.GetDataroomViewByLinkidDocumentByDocumentidFile(context.Background(), linkId, documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.GetDataroomViewByLinkidDocumentByDocumentidFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataroomViewByLinkidDocumentByDocumentidFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PatchDataroomTrustArtifactsById

> DataroomTrustItemView PatchDataroomTrustArtifactsById(ctx, id).DataroomTrustEdit(dataroomTrustEdit).Execute()

Changes an item on the caller org's trust centre — replace its file with a newer edition, move it between public and gated, rewrite what it says, or retire it — and answers with the item as it now stands.



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
	id := "id_example" // string | ID is the item to change, taken from the path.
	dataroomTrustEdit := *openapiclient.NewDataroomTrustEdit() // DataroomTrustEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PatchDataroomTrustArtifactsById(context.Background(), id).DataroomTrustEdit(dataroomTrustEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PatchDataroomTrustArtifactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataroomTrustArtifactsById`: DataroomTrustItemView
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PatchDataroomTrustArtifactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the item to change, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataroomTrustArtifactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataroomTrustEdit** | [**DataroomTrustEdit**](DataroomTrustEdit.md) |  | 

### Return type

[**DataroomTrustItemView**](DataroomTrustItemView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomDatarooms

> DataroomDataroomRoomOne PostDataroomDatarooms(ctx).DataroomDataroomCreate(dataroomDataroomCreate).Execute()

Opens a new data room for the caller org and answers with it, including the short public id it is addressed by.



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
	dataroomDataroomCreate := *openapiclient.NewDataroomDataroomCreate() // DataroomDataroomCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomDatarooms(context.Background()).DataroomDataroomCreate(dataroomDataroomCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomDatarooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomDatarooms`: DataroomDataroomRoomOne
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomDatarooms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomDataroomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataroomDataroomCreate** | [**DataroomDataroomCreate**](DataroomDataroomCreate.md) |  | 

### Return type

[**DataroomDataroomRoomOne**](DataroomDataroomRoomOne.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomDataroomsByIdDocuments

> DataroomDataroomMembership PostDataroomDataroomsByIdDocuments(ctx, id).DataroomDataroomAddDocument(dataroomDataroomAddDocument).Execute()

Puts an already-uploaded document into one of the caller org's data rooms and answers with the new membership id.



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
	id := "id_example" // string | ID is the room to add to. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.
	dataroomDataroomAddDocument := *openapiclient.NewDataroomDataroomAddDocument() // DataroomDataroomAddDocument | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomDataroomsByIdDocuments(context.Background(), id).DataroomDataroomAddDocument(dataroomDataroomAddDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomDataroomsByIdDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomDataroomsByIdDocuments`: DataroomDataroomMembership
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomDataroomsByIdDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room to add to. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomDataroomsByIdDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataroomDataroomAddDocument** | [**DataroomDataroomAddDocument**](DataroomDataroomAddDocument.md) |  | 

### Return type

[**DataroomDataroomMembership**](DataroomDataroomMembership.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomDocuments

> PostDataroomDocuments(ctx).Execute()

Upload a document's bytes and record it



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
	r, err := apiClient.DataroomAPI.PostDataroomDocuments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomDocumentsRequest struct via the builder pattern


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


## PostDataroomLinks

> DataroomDataroomLinkOne PostDataroomLinks(ctx).DataroomDataroomLinkCreate(dataroomDataroomLinkCreate).Execute()

Grants access: it mints a public share link over one data room (`dataroomId`) or one document (`documentId`) — one of the two is required — and answers with the link, whose `id` is the token a visitor opens it with.



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
	dataroomDataroomLinkCreate := *openapiclient.NewDataroomDataroomLinkCreate() // DataroomDataroomLinkCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomLinks(context.Background()).DataroomDataroomLinkCreate(dataroomDataroomLinkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomLinks`: DataroomDataroomLinkOne
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataroomDataroomLinkCreate** | [**DataroomDataroomLinkCreate**](DataroomDataroomLinkCreate.md) |  | 

### Return type

[**DataroomDataroomLinkOne**](DataroomDataroomLinkOne.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomTrustArtifacts

> DataroomTrustItemView PostDataroomTrustArtifacts(ctx).DataroomTrustPublish(dataroomTrustPublish).Execute()

Puts an item on the caller org's trust centre and answers with it.



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
	dataroomTrustPublish := *openapiclient.NewDataroomTrustPublish() // DataroomTrustPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomTrustArtifacts(context.Background()).DataroomTrustPublish(dataroomTrustPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomTrustArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomTrustArtifacts`: DataroomTrustItemView
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomTrustArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomTrustArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataroomTrustPublish** | [**DataroomTrustPublish**](DataroomTrustPublish.md) |  | 

### Return type

[**DataroomTrustItemView**](DataroomTrustItemView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomTrustCenterBySlugRequests

> DataroomTrustAsked PostDataroomTrustCenterBySlugRequests(ctx, slug).DataroomTrustAsk(dataroomTrustAsk).Execute()

Records a request to read what an independent auditor signed, and answers with its id.



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
	slug := "slug_example" // string | Slug is the centre's public address, taken from the path.
	dataroomTrustAsk := *openapiclient.NewDataroomTrustAsk() // DataroomTrustAsk | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomTrustCenterBySlugRequests(context.Background(), slug).DataroomTrustAsk(dataroomTrustAsk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomTrustCenterBySlugRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomTrustCenterBySlugRequests`: DataroomTrustAsked
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomTrustCenterBySlugRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the centre&#39;s public address, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomTrustCenterBySlugRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataroomTrustAsk** | [**DataroomTrustAsk**](DataroomTrustAsk.md) |  | 

### Return type

[**DataroomTrustAsked**](DataroomTrustAsked.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomTrustRequestsByIdGrant

> DataroomTrustGranted PostDataroomTrustRequestsByIdGrant(ctx, id).DataroomTrustDecision(dataroomTrustDecision).Execute()

Answers a request by opening access: it mints a share link over what was asked for, addressed to the address that asked and closing at expiry, records the decision, and mails the asker.



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
	id := "id_example" // string | ID is the request to answer, taken from the path.
	dataroomTrustDecision := *openapiclient.NewDataroomTrustDecision() // DataroomTrustDecision | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomTrustRequestsByIdGrant(context.Background(), id).DataroomTrustDecision(dataroomTrustDecision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomTrustRequestsByIdGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomTrustRequestsByIdGrant`: DataroomTrustGranted
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomTrustRequestsByIdGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the request to answer, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomTrustRequestsByIdGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataroomTrustDecision** | [**DataroomTrustDecision**](DataroomTrustDecision.md) |  | 

### Return type

[**DataroomTrustGranted**](DataroomTrustGranted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomTrustRequestsByIdRefuse

> DataroomTrustRefused PostDataroomTrustRequestsByIdRefuse(ctx, id).DataroomTrustDecision(dataroomTrustDecision).Execute()

Answers a request by declining it, recording who declined and why.



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
	id := "id_example" // string | ID is the request to answer, taken from the path.
	dataroomTrustDecision := *openapiclient.NewDataroomTrustDecision() // DataroomTrustDecision | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PostDataroomTrustRequestsByIdRefuse(context.Background(), id).DataroomTrustDecision(dataroomTrustDecision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomTrustRequestsByIdRefuse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataroomTrustRequestsByIdRefuse`: DataroomTrustRefused
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PostDataroomTrustRequestsByIdRefuse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the request to answer, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomTrustRequestsByIdRefuseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataroomTrustDecision** | [**DataroomTrustDecision**](DataroomTrustDecision.md) |  | 

### Return type

[**DataroomTrustRefused**](DataroomTrustRefused.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDataroomViewByLinkidAuthenticate

> PostDataroomViewByLinkidAuthenticate(ctx, linkId).Execute()

Pass a share link's gates and open a viewing session



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
	linkId := "linkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataroomAPI.PostDataroomViewByLinkidAuthenticate(context.Background(), linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomViewByLinkidAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomViewByLinkidAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostDataroomViewByLinkidPageview

> PostDataroomViewByLinkidPageview(ctx, linkId).Execute()

Record one page-view against an open viewing session



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
	linkId := "linkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataroomAPI.PostDataroomViewByLinkidPageview(context.Background(), linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PostDataroomViewByLinkidPageview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDataroomViewByLinkidPageviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PutDataroomTrust

> DataroomTrustDesk PutDataroomTrust(ctx).DataroomTrustSettings(dataroomTrustSettings).Execute()

SetCenter opens, publishes or withdraws the caller org's trust centre and answers with the centre as it now stands.



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
	dataroomTrustSettings := *openapiclient.NewDataroomTrustSettings() // DataroomTrustSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataroomAPI.PutDataroomTrust(context.Background()).DataroomTrustSettings(dataroomTrustSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataroomAPI.PutDataroomTrust``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataroomTrust`: DataroomTrustDesk
	fmt.Fprintf(os.Stdout, "Response from `DataroomAPI.PutDataroomTrust`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDataroomTrustRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataroomTrustSettings** | [**DataroomTrustSettings**](DataroomTrustSettings.md) |  | 

### Return type

[**DataroomTrustDesk**](DataroomTrustDesk.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

