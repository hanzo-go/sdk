# \TeamAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTeamAccountCookie**](TeamAPI.md#DeleteTeamAccountCookie) | **Delete** /v1/team/account/cookie | Signs this browser out of team by expiring the HttpOnly account-token cookie the OAuth callback set.
[**DeleteTeamFilesBySpaceByFilename**](TeamAPI.md#DeleteTeamFilesBySpaceByFilename) | **Delete** /v1/team/files/{space}/{filename} | Removes one blob from a space&#39;s file store.
[**GetTeamAccountAuthByProvider**](TeamAPI.md#GetTeamAccountAuthByProvider) | **Get** /v1/team/account/auth/{provider} | Start a sign-in at hanzo.id
[**GetTeamAccountAuthByProviderCallback**](TeamAPI.md#GetTeamAccountAuthByProviderCallback) | **Get** /v1/team/account/auth/{provider}/callback | Complete a sign-in and hand the browser its session
[**GetTeamAccountProviders**](TeamAPI.md#GetTeamAccountProviders) | **Get** /v1/team/account/providers | Returns the identity providers this deployment starts a login with.
[**GetTeamBillingPlan**](TeamAPI.md#GetTeamBillingPlan) | **Get** /v1/team/billing/plan | Returns the plan and seat counts for the caller&#39;s OWN org, resolved from the VERIFIED team session token — never a client header.
[**GetTeamBillingUi**](TeamAPI.md#GetTeamBillingUi) | **Get** /v1/team/billing/ui | Open the wallet page
[**GetTeamBots**](TeamAPI.md#GetTeamBots) | **Get** /v1/team/bots | Returns the caller org&#39;s bot members — the org&#39;s agents projected as the space Employees they become, each with the member account uuid and Person reference the roster addresses it by.
[**GetTeamCollaborator**](TeamAPI.md#GetTeamCollaborator) | **Get** /v1/team/collaborator | Open the live collaborative-editing socket
[**GetTeamFilesBySpaceByFilename**](TeamAPI.md#GetTeamFilesBySpaceByFilename) | **Get** /v1/team/files/{space}/{filename} | Download a space file
[**GetTeamPublic**](TeamAPI.md#GetTeamPublic) | **Get** /v1/team/public | Lists the rooms orgs have published, across every org.
[**GetTeamRooms**](TeamAPI.md#GetTeamRooms) | **Get** /v1/team/rooms | Returns every room of the caller&#39;s org, across the spaces it owns, with the work facet each carries.
[**GetTeamRoomsByIdMessages**](TeamAPI.md#GetTeamRoomsByIdMessages) | **Get** /v1/team/rooms/{id}/messages | Returns the tail of one room&#39;s conversation, oldest first.
[**GetTeamTransactorByToken**](TeamAPI.md#GetTeamTransactorByToken) | **Get** /v1/team/transactor/{token} | Open the space data-plane socket
[**GetTeamTransactorStatistics**](TeamAPI.md#GetTeamTransactorStatistics) | **Get** /v1/team/transactor/statistics | Statistics returns the transactor&#39;s live sessions for the space the caller&#39;s credential names — the endpoint the front&#39;s space switcher and server panel poll on the transactor base.
[**PostTeamAccount**](TeamAPI.md#PostTeamAccount) | **Post** /v1/team/account | Read the caller&#39;s account and switch space
[**PostTeamBotsSync**](TeamAPI.md#PostTeamBotsSync) | **Post** /v1/team/bots/sync | SyncBots re-projects the caller org&#39;s agents as space members into EVERY space of the org, and removes the ones whose agent is gone.
[**PostTeamCollaboratorRpcByDocumentid**](TeamAPI.md#PostTeamCollaboratorRpcByDocumentid) | **Post** /v1/team/collaborator/rpc/{documentId} | CollabRPC is the collaborative-markup snapshot plane the Team front&#39;s editor speaks: createContent stores a document field&#39;s markup at a fresh, immutable blob ref and returns it, updateContent stores a new snapshot and answers nothing, and getContent reads back the exact snapshot a ref names.
[**PostTeamFilesBySpace**](TeamAPI.md#PostTeamFilesBySpace) | **Post** /v1/team/files/{space} | Upload a file into a space
[**PostTeamRooms**](TeamAPI.md#PostTeamRooms) | **Post** /v1/team/rooms | Opens a named room and answers it as the store now holds it.
[**PostTeamRoomsByIdMessages**](TeamAPI.md#PostTeamRoomsByIdMessages) | **Post** /v1/team/rooms/{id}/messages | Says one thing in a room, as the caller.
[**PutTeamAccountCookie**](TeamAPI.md#PutTeamAccountCookie) | **Put** /v1/team/account/cookie | Store the session token as this browser&#39;s cookie
[**PutTeamRoomsById**](TeamAPI.md#PutTeamRoomsById) | **Put** /v1/team/rooms/{id} | States what a room is for: its lifecycle intent, and what it is about.



## DeleteTeamAccountCookie

> CookieAck DeleteTeamAccountCookie(ctx).Execute()

Signs this browser out of team by expiring the HttpOnly account-token cookie the OAuth callback set.



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
	resp, r, err := apiClient.TeamAPI.DeleteTeamAccountCookie(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamAccountCookie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeamAccountCookie`: CookieAck
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.DeleteTeamAccountCookie`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamAccountCookieRequest struct via the builder pattern


### Return type

[**CookieAck**](CookieAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamFilesBySpaceByFilename

> DeleteTeamFilesBySpaceByFilename(ctx, space, filename).File(file).Execute()

Removes one blob from a space's file store.



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
	space := "6579…" // string | Space is the space uuid the blob belongs to, from the path.
	filename := "filename_example" // string | Filename is the last path segment, which the front sets to the blob id when it sends no explicit `file`.
	file := "0d4f…" // string | File is the blob id, and wins over the path segment when both are present. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeamFilesBySpaceByFilename(context.Background(), space, filename).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamFilesBySpaceByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Space is the space uuid the blob belongs to, from the path. | 
**filename** | **string** | Filename is the last path segment, which the front sets to the blob id when it sends no explicit &#x60;file&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamFilesBySpaceByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **string** | File is the blob id, and wins over the path segment when both are present. | 

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


## GetTeamAccountAuthByProvider

> GetTeamAccountAuthByProvider(ctx, provider).Execute()

Start a sign-in at hanzo.id



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.GetTeamAccountAuthByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamAccountAuthByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamAccountAuthByProviderRequest struct via the builder pattern


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


## GetTeamAccountAuthByProviderCallback

> GetTeamAccountAuthByProviderCallback(ctx, provider).Execute()

Complete a sign-in and hand the browser its session



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.GetTeamAccountAuthByProviderCallback(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamAccountAuthByProviderCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamAccountAuthByProviderCallbackRequest struct via the builder pattern


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


## GetTeamAccountProviders

> []ProviderInfo GetTeamAccountProviders(ctx).Execute()

Returns the identity providers this deployment starts a login with.



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
	resp, r, err := apiClient.TeamAPI.GetTeamAccountProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamAccountProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamAccountProviders`: []ProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamAccountProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamAccountProvidersRequest struct via the builder pattern


### Return type

[**[]ProviderInfo**](ProviderInfo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamBillingPlan

> PlanInfo GetTeamBillingPlan(ctx).Execute()

Returns the plan and seat counts for the caller's OWN org, resolved from the VERIFIED team session token — never a client header.



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
	resp, r, err := apiClient.TeamAPI.GetTeamBillingPlan(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamBillingPlan`: PlanInfo
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamBillingPlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamBillingPlanRequest struct via the builder pattern


### Return type

[**PlanInfo**](PlanInfo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamBillingUi

> *os.File GetTeamBillingUi(ctx).Execute()

Open the wallet page



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
	resp, r, err := apiClient.TeamAPI.GetTeamBillingUi(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamBillingUi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamBillingUi`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamBillingUi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamBillingUiRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamBots

> BotRoster GetTeamBots(ctx).Execute()

Returns the caller org's bot members — the org's agents projected as the space Employees they become, each with the member account uuid and Person reference the roster addresses it by.



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
	resp, r, err := apiClient.TeamAPI.GetTeamBots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamBots`: BotRoster
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamBotsRequest struct via the builder pattern


### Return type

[**BotRoster**](BotRoster.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamCollaborator

> GetTeamCollaborator(ctx).Execute()

Open the live collaborative-editing socket



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
	r, err := apiClient.TeamAPI.GetTeamCollaborator(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamCollaborator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamCollaboratorRequest struct via the builder pattern


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


## GetTeamFilesBySpaceByFilename

> *os.File GetTeamFilesBySpaceByFilename(ctx, space, filename).Execute()

Download a space file



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
	space := "space_example" // string | 
	filename := "filename_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamFilesBySpaceByFilename(context.Background(), space, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamFilesBySpaceByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamFilesBySpaceByFilename`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamFilesBySpaceByFilename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamFilesBySpaceByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamPublic

> PublicRooms GetTeamPublic(ctx).Q(q).Org(org).Limit(limit).Execute()

Lists the rooms orgs have published, across every org.



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
	q := "q_example" // string | Q matches a room's name or its topic. (optional)
	org := "org_example" // string | Org narrows to one org's published rooms. (optional)
	limit := int64(789) // int64 | Limit caps the page, 50 when unstated and 200 at most. An unparseable value reads as unstated rather than as zero — zero pages is not an answer anybody asked for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamPublic(context.Background()).Q(q).Org(org).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamPublic`: PublicRooms
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q matches a room&#39;s name or its topic. | 
 **org** | **string** | Org narrows to one org&#39;s published rooms. | 
 **limit** | **int64** | Limit caps the page, 50 when unstated and 200 at most. An unparseable value reads as unstated rather than as zero — zero pages is not an answer anybody asked for. | 

### Return type

[**PublicRooms**](PublicRooms.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamRooms

> TeamRooms GetTeamRooms(ctx).Execute()

Returns every room of the caller's org, across the spaces it owns, with the work facet each carries.



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
	resp, r, err := apiClient.TeamAPI.GetTeamRooms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamRooms`: TeamRooms
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamRooms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRoomsRequest struct via the builder pattern


### Return type

[**TeamRooms**](TeamRooms.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamRoomsByIdMessages

> TeamMessages GetTeamRoomsByIdMessages(ctx, id).Space(space).Execute()

Returns the tail of one room's conversation, oldest first.



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
	id := "id_example" // string | ID is the room, from the path. The URL is the authority.
	space := "space_example" // string | Space names the space holding the room, and is required for the reason the bind op requires it: a room id is unique within a space and not across the org, so searching every space for a match would make the answer depend on iteration order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamRoomsByIdMessages(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamRoomsByIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamRoomsByIdMessages`: TeamMessages
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamRoomsByIdMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room, from the path. The URL is the authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRoomsByIdMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space names the space holding the room, and is required for the reason the bind op requires it: a room id is unique within a space and not across the org, so searching every space for a match would make the answer depend on iteration order. | 

### Return type

[**TeamMessages**](TeamMessages.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamTransactorByToken

> GetTeamTransactorByToken(ctx, token).Execute()

Open the space data-plane socket



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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.GetTeamTransactorByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamTransactorByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamTransactorByTokenRequest struct via the builder pattern


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


## GetTeamTransactorStatistics

> StatsOut GetTeamTransactorStatistics(ctx).Token(token).Execute()

Statistics returns the transactor's live sessions for the space the caller's credential names — the endpoint the front's space switcher and server panel poll on the transactor base.



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
	token := "eyJhbGciOiJIUzI1NiJ9…" // string | Token is the space token minted by selectWorkspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamTransactorStatistics(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamTransactorStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamTransactorStatistics`: StatsOut
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamTransactorStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamTransactorStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token is the space token minted by selectWorkspace. | 

### Return type

[**StatsOut**](StatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamAccount

> PostTeamAccount(ctx).Execute()

Read the caller's account and switch space



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
	r, err := apiClient.TeamAPI.PostTeamAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamAccountRequest struct via the builder pattern


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


## PostTeamBotsSync

> BotSync PostTeamBotsSync(ctx).Execute()

SyncBots re-projects the caller org's agents as space members into EVERY space of the org, and removes the ones whose agent is gone.



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
	resp, r, err := apiClient.TeamAPI.PostTeamBotsSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamBotsSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamBotsSync`: BotSync
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamBotsSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamBotsSyncRequest struct via the builder pattern


### Return type

[**BotSync**](BotSync.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamCollaboratorRpcByDocumentid

> CollabResult PostTeamCollaboratorRpcByDocumentid(ctx, documentId).CollabRequest(collabRequest).Execute()

CollabRPC is the collaborative-markup snapshot plane the Team front's editor speaks: createContent stores a document field's markup at a fresh, immutable blob ref and returns it, updateContent stores a new snapshot and answers nothing, and getContent reads back the exact snapshot a ref names.



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
	documentId := "6579…|tracker:class:Issue|issue-1|description" // string | DocumentID addresses the document field, as \"<spaceUuid>|<objectClass>|<objectId>|<objectAttr>\" — the collaborator-client encodeDocumentId shape, from the path.
	collabRequest := *openapiclient.NewCollabRequest() // CollabRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamCollaboratorRpcByDocumentid(context.Background(), documentId).CollabRequest(collabRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamCollaboratorRpcByDocumentid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamCollaboratorRpcByDocumentid`: CollabResult
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamCollaboratorRpcByDocumentid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | DocumentID addresses the document field, as \&quot;&lt;spaceUuid&gt;|&lt;objectClass&gt;|&lt;objectId&gt;|&lt;objectAttr&gt;\&quot; — the collaborator-client encodeDocumentId shape, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamCollaboratorRpcByDocumentidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collabRequest** | [**CollabRequest**](CollabRequest.md) |  | 

### Return type

[**CollabResult**](CollabResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamFilesBySpace

> *os.File PostTeamFilesBySpace(ctx, space).Body(body).Execute()

Upload a file into a space



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
	space := "space_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamFilesBySpace(context.Background(), space).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamFilesBySpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamFilesBySpace`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamFilesBySpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamFilesBySpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamRooms

> TeamRoom PostTeamRooms(ctx).TeamRoomNew(teamRoomNew).Execute()

Opens a named room and answers it as the store now holds it.



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
	teamRoomNew := *openapiclient.NewTeamRoomNew() // TeamRoomNew | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamRooms(context.Background()).TeamRoomNew(teamRoomNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamRooms`: TeamRoom
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamRooms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamRoomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamRoomNew** | [**TeamRoomNew**](TeamRoomNew.md) |  | 

### Return type

[**TeamRoom**](TeamRoom.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamRoomsByIdMessages

> TeamMessage PostTeamRoomsByIdMessages(ctx, id).TeamMessageWrite(teamMessageWrite).Execute()

Says one thing in a room, as the caller.



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
	id := "id_example" // string | ID is the room to say it in, from the path.
	teamMessageWrite := *openapiclient.NewTeamMessageWrite() // TeamMessageWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamRoomsByIdMessages(context.Background(), id).TeamMessageWrite(teamMessageWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamRoomsByIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamRoomsByIdMessages`: TeamMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamRoomsByIdMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room to say it in, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamRoomsByIdMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamMessageWrite** | [**TeamMessageWrite**](TeamMessageWrite.md) |  | 

### Return type

[**TeamMessage**](TeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTeamAccountCookie

> CookieAck PutTeamAccountCookie(ctx).Execute()

Store the session token as this browser's cookie



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
	resp, r, err := apiClient.TeamAPI.PutTeamAccountCookie(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PutTeamAccountCookie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTeamAccountCookie`: CookieAck
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PutTeamAccountCookie`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutTeamAccountCookieRequest struct via the builder pattern


### Return type

[**CookieAck**](CookieAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTeamRoomsById

> TeamRoom PutTeamRoomsById(ctx, id).TeamRoomBind(teamRoomBind).Execute()

States what a room is for: its lifecycle intent, and what it is about.



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
	id := "id_example" // string | ID is the room to bind, from the path. The URL is the authority; a body carrying another id cannot redirect the write.
	teamRoomBind := *openapiclient.NewTeamRoomBind() // TeamRoomBind | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PutTeamRoomsById(context.Background(), id).TeamRoomBind(teamRoomBind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PutTeamRoomsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTeamRoomsById`: TeamRoom
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PutTeamRoomsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room to bind, from the path. The URL is the authority; a body carrying another id cannot redirect the write. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTeamRoomsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamRoomBind** | [**TeamRoomBind**](TeamRoomBind.md) |  | 

### Return type

[**TeamRoom**](TeamRoom.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

