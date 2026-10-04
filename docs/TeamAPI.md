# \TeamAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTeamAccountCookie**](TeamAPI.md#DeleteTeamAccountCookie) | **Delete** /v1/team/account/cookie | Signs this browser out of team by expiring the HttpOnly account-token cookie the OAuth callback set.
[**DeleteTeamDocsById**](TeamAPI.md#DeleteTeamDocsById) | **Delete** /v1/team/docs/{id} | Removes a document with everything nested under it and every comment on any of them — the Team client&#39;s own delete, which takes the subtree with it.
[**DeleteTeamFilesBySpaceByFilename**](TeamAPI.md#DeleteTeamFilesBySpaceByFilename) | **Delete** /v1/team/files/{space}/{filename} | Removes one blob from a space&#39;s file store.
[**DeleteTeamMessagesById**](TeamAPI.md#DeleteTeamMessagesById) | **Delete** /v1/team/messages/{id} | Removes a message, with its replies, reactions, files and the inbox notifications that point at it.
[**DeleteTeamMessagesByIdReactionsByEmoji**](TeamAPI.md#DeleteTeamMessagesByIdReactionsByEmoji) | **Delete** /v1/team/messages/{id}/reactions/{emoji} | Takes back the caller&#39;s reaction to a message and answers the message with its reactions as they now stand.
[**DeleteTeamRoomsByIdMembersByAccount**](TeamAPI.md#DeleteTeamRoomsByIdMembersByAccount) | **Delete** /v1/team/rooms/{id}/members/{account} | Takes one person out of a room — the caller leaving, when the account is their own.
[**GetTeamAccountAuthByProvider**](TeamAPI.md#GetTeamAccountAuthByProvider) | **Get** /v1/team/account/auth/{provider} | Start a sign-in at hanzo.id
[**GetTeamAccountAuthByProviderCallback**](TeamAPI.md#GetTeamAccountAuthByProviderCallback) | **Get** /v1/team/account/auth/{provider}/callback | Complete a sign-in and hand the browser its session
[**GetTeamAccountProviders**](TeamAPI.md#GetTeamAccountProviders) | **Get** /v1/team/account/providers | Returns the identity providers this deployment starts a login with.
[**GetTeamBillingPlan**](TeamAPI.md#GetTeamBillingPlan) | **Get** /v1/team/billing/plan | Returns the plan and seat counts for the caller&#39;s OWN org, resolved from the VERIFIED team session token — never a client header.
[**GetTeamBillingUi**](TeamAPI.md#GetTeamBillingUi) | **Get** /v1/team/billing/ui | Open the wallet page
[**GetTeamCollaborator**](TeamAPI.md#GetTeamCollaborator) | **Get** /v1/team/collaborator | Open the live collaborative-editing socket
[**GetTeamDocs**](TeamAPI.md#GetTeamDocs) | **Get** /v1/team/docs | Returns the documents of a space the caller may see, with the teamspaces they are grouped in.
[**GetTeamDocsById**](TeamAPI.md#GetTeamDocsById) | **Get** /v1/team/docs/{id} | Returns one document the caller may see.
[**GetTeamDocsByIdComments**](TeamAPI.md#GetTeamDocsByIdComments) | **Get** /v1/team/docs/{id}/comments | Returns the comments on a document, oldest first — the same message shape a room&#39;s conversation answers.
[**GetTeamEvents**](TeamAPI.md#GetTeamEvents) | **Get** /v1/team/events | Stream live changes to what the caller may see, as Server-Sent Events
[**GetTeamFilesBySpaceByFilename**](TeamAPI.md#GetTeamFilesBySpaceByFilename) | **Get** /v1/team/files/{space}/{filename} | Download a space file
[**GetTeamInbox**](TeamAPI.md#GetTeamInbox) | **Get** /v1/team/inbox | Returns the caller&#39;s notifications, newest first, each with the room or document it is about and the message that caused it.
[**GetTeamMembers**](TeamAPI.md#GetTeamMembers) | **Get** /v1/team/members | Returns the people and agents in a space, with the name, avatar, role and presence a client draws them with.
[**GetTeamMessagesByIdReplies**](TeamAPI.md#GetTeamMessagesByIdReplies) | **Get** /v1/team/messages/{id}/replies | Returns a message&#39;s thread, oldest first, each reply with its reactions and files.
[**GetTeamPublic**](TeamAPI.md#GetTeamPublic) | **Get** /v1/team/public | Lists the rooms orgs have published, across every org.
[**GetTeamRooms**](TeamAPI.md#GetTeamRooms) | **Get** /v1/team/rooms | Returns the rooms the caller may see, with the kind and work facet each carries.
[**GetTeamRoomsByIdMembers**](TeamAPI.md#GetTeamRoomsByIdMembers) | **Get** /v1/team/rooms/{id}/members | Returns the people and agents in one room, as the roster describes them.
[**GetTeamRoomsByIdMessages**](TeamAPI.md#GetTeamRoomsByIdMessages) | **Get** /v1/team/rooms/{id}/messages | Returns the tail of one room&#39;s conversation, oldest first, each message with its reactions, files and thread count.
[**GetTeamTransactorByToken**](TeamAPI.md#GetTeamTransactorByToken) | **Get** /v1/team/transactor/{token} | Open the space data-plane socket
[**GetTeamTransactorStatistics**](TeamAPI.md#GetTeamTransactorStatistics) | **Get** /v1/team/transactor/statistics | Returns the transactor&#39;s live sessions for the space the caller&#39;s credential names — the endpoint the front&#39;s space switcher and server panel poll on the transactor base.
[**PatchTeamDocsById**](TeamAPI.md#PatchTeamDocsById) | **Patch** /v1/team/docs/{id} | Renames a document or moves it — under another document, to the top of its teamspace, or to another teamspace with everything nested under it.
[**PatchTeamMessagesById**](TeamAPI.md#PatchTeamMessagesById) | **Patch** /v1/team/messages/{id} | Rewrites what a message says.
[**PatchTeamRoomsById**](TeamAPI.md#PatchTeamRoomsById) | **Patch** /v1/team/rooms/{id} | Renames a channel, sets its topic, or archives or reopens it, and answers the room as it now stands.
[**PostTeamAccount**](TeamAPI.md#PostTeamAccount) | **Post** /v1/team/account | Read the caller&#39;s account and switch space
[**PostTeamCollaboratorRpcByDocumentid**](TeamAPI.md#PostTeamCollaboratorRpcByDocumentid) | **Post** /v1/team/collaborator/rpc/{documentId} | CollabRPC is the collaborative-markup snapshot plane the Team front&#39;s editor speaks: createContent stores a document field&#39;s markup at a fresh, immutable blob ref and returns it, updateContent stores a new snapshot and answers nothing, and getContent reads back the exact snapshot a ref names.
[**PostTeamDms**](TeamAPI.md#PostTeamDms) | **Post** /v1/team/dms | Opens the direct message between the caller and the people named, or answers the one that already exists — there is one conversation per set of people, however many times it is asked for.
[**PostTeamDocs**](TeamAPI.md#PostTeamDocs) | **Post** /v1/team/docs | Creates a document, as the caller, after its siblings.
[**PostTeamDocsByIdComments**](TeamAPI.md#PostTeamDocsByIdComments) | **Post** /v1/team/docs/{id}/comments | Comments on a document, as the caller.
[**PostTeamFilesBySpace**](TeamAPI.md#PostTeamFilesBySpace) | **Post** /v1/team/files/{space} | Upload a file into a space
[**PostTeamInboxByIdArchive**](TeamAPI.md#PostTeamInboxByIdArchive) | **Post** /v1/team/inbox/{id}/archive | Archives one of the caller&#39;s notifications — read, and out of the live inbox — and answers it.
[**PostTeamInboxByIdRead**](TeamAPI.md#PostTeamInboxByIdRead) | **Post** /v1/team/inbox/{id}/read | Marks one of the caller&#39;s notifications read and answers it.
[**PostTeamInboxRead**](TeamAPI.md#PostTeamInboxRead) | **Post** /v1/team/inbox/read | Marks every live notification of the caller&#39;s read, and says how many it changed.
[**PostTeamMessagesByIdReplies**](TeamAPI.md#PostTeamMessagesByIdReplies) | **Post** /v1/team/messages/{id}/replies | Answers a message in its thread, as the caller.
[**PostTeamRooms**](TeamAPI.md#PostTeamRooms) | **Post** /v1/team/rooms | Opens a named room and answers it as the store now holds it.
[**PostTeamRoomsByIdMembers**](TeamAPI.md#PostTeamRoomsByIdMembers) | **Post** /v1/team/rooms/{id}/members | Adds people to a room and answers the room as it now stands.
[**PostTeamRoomsByIdMessages**](TeamAPI.md#PostTeamRoomsByIdMessages) | **Post** /v1/team/rooms/{id}/messages | Says one thing in a room, as the caller.
[**PutTeamAccountCookie**](TeamAPI.md#PutTeamAccountCookie) | **Put** /v1/team/account/cookie | Store the session token as this browser&#39;s cookie
[**PutTeamMessagesByIdReactionsByEmoji**](TeamAPI.md#PutTeamMessagesByIdReactionsByEmoji) | **Put** /v1/team/messages/{id}/reactions/{emoji} | Adds the caller&#39;s reaction to a message and answers the message with its reactions as they now stand.
[**PutTeamRoomsById**](TeamAPI.md#PutTeamRoomsById) | **Put** /v1/team/rooms/{id} | States what a room is for: its lifecycle intent, and what it is about.



## DeleteTeamAccountCookie

> TeamCookieAck DeleteTeamAccountCookie(ctx).Execute()

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
	// response from `DeleteTeamAccountCookie`: TeamCookieAck
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.DeleteTeamAccountCookie`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamAccountCookieRequest struct via the builder pattern


### Return type

[**TeamCookieAck**](TeamCookieAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamDocsById

> DeleteTeamDocsById(ctx, id).Space(space).Execute()

Removes a document with everything nested under it and every comment on any of them — the Team client's own delete, which takes the subtree with it.



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
	id := "id_example" // string | ID is the document, from the path.
	space := "0e3c…" // string | Space is the space uuid holding it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeamDocsById(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamDocsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamDocsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space is the space uuid holding it. | 

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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamMessagesById

> DeleteTeamMessagesById(ctx, id).Space(space).Execute()

Removes a message, with its replies, reactions, files and the inbox notifications that point at it.



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
	id := "id_example" // string | ID is the message, from the path.
	space := "0e3c…" // string | Space names the space holding it. A message id is unique within a space, not across the org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeamMessagesById(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamMessagesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the message, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamMessagesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space names the space holding it. A message id is unique within a space, not across the org. | 

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


## DeleteTeamMessagesByIdReactionsByEmoji

> TeamTeamMessage DeleteTeamMessagesByIdReactionsByEmoji(ctx, id, emoji).Space(space).Execute()

Takes back the caller's reaction to a message and answers the message with its reactions as they now stand.



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
	id := "id_example" // string | ID is the message, from the path.
	emoji := "emoji_example" // string | Emoji is the reaction, from the path (percent-encoded on the wire).
	space := "0e3c…" // string | Space names the space holding the message. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.DeleteTeamMessagesByIdReactionsByEmoji(context.Background(), id, emoji).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamMessagesByIdReactionsByEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeamMessagesByIdReactionsByEmoji`: TeamTeamMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.DeleteTeamMessagesByIdReactionsByEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the message, from the path. | 
**emoji** | **string** | Emoji is the reaction, from the path (percent-encoded on the wire). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamMessagesByIdReactionsByEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **space** | **string** | Space names the space holding the message. | 

### Return type

[**TeamTeamMessage**](TeamTeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamRoomsByIdMembersByAccount

> DeleteTeamRoomsByIdMembersByAccount(ctx, id, account).Space(space).Execute()

Takes one person out of a room — the caller leaving, when the account is their own.



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
	id := "id_example" // string | ID is the room, from the path.
	account := "account_example" // string | Account is the account uuid to remove, from the path. Your own is leaving.
	space := "0e3c…" // string | Space is the space uuid holding the room. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeamRoomsByIdMembersByAccount(context.Background(), id, account).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamRoomsByIdMembersByAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room, from the path. | 
**account** | **string** | Account is the account uuid to remove, from the path. Your own is leaving. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRoomsByIdMembersByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **space** | **string** | Space is the space uuid holding the room. | 

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

> []TeamProviderInfo GetTeamAccountProviders(ctx).Execute()

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
	// response from `GetTeamAccountProviders`: []TeamProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamAccountProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamAccountProvidersRequest struct via the builder pattern


### Return type

[**[]TeamProviderInfo**](TeamProviderInfo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamBillingPlan

> TeamPlanInfo GetTeamBillingPlan(ctx).Execute()

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
	// response from `GetTeamBillingPlan`: TeamPlanInfo
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamBillingPlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamBillingPlanRequest struct via the builder pattern


### Return type

[**TeamPlanInfo**](TeamPlanInfo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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


## GetTeamDocs

> TeamTeamDocs GetTeamDocs(ctx).Space(space).Teamspace(teamspace).Parent(parent).Execute()

Returns the documents of a space the caller may see, with the teamspaces they are grouped in.



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
	space := "space_example" // string | Space is the space uuid. Optional for a caller in exactly one space. (optional)
	teamspace := "teamspace_example" // string | Teamspace narrows the answer to one teamspace. (optional)
	parent := "parent_example" // string | Parent narrows the answer to one document's children. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamDocs(context.Background()).Space(space).Teamspace(teamspace).Parent(parent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamDocs`: TeamTeamDocs
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **space** | **string** | Space is the space uuid. Optional for a caller in exactly one space. | 
 **teamspace** | **string** | Teamspace narrows the answer to one teamspace. | 
 **parent** | **string** | Parent narrows the answer to one document&#39;s children. | 

### Return type

[**TeamTeamDocs**](TeamTeamDocs.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamDocsById

> TeamTeamDoc GetTeamDocsById(ctx, id).Space(space).Execute()

Returns one document the caller may see.



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
	id := "d41c" // string | ID is the document, from the path.
	space := "0e3c…" // string | Space is the space uuid holding it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamDocsById(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamDocsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamDocsById`: TeamTeamDoc
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamDocsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamDocsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space is the space uuid holding it. | 

### Return type

[**TeamTeamDoc**](TeamTeamDoc.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamDocsByIdComments

> TeamTeamMessages GetTeamDocsByIdComments(ctx, id).Space(space).Execute()

Returns the comments on a document, oldest first — the same message shape a room's conversation answers.



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
	id := "id_example" // string | ID is the document, from the path.
	space := "space_example" // string | Space is the space uuid holding it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamDocsByIdComments(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamDocsByIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamDocsByIdComments`: TeamTeamMessages
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamDocsByIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamDocsByIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space is the space uuid holding it. | 

### Return type

[**TeamTeamMessages**](TeamTeamMessages.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamEvents

> *os.File GetTeamEvents(ctx).Execute()

Stream live changes to what the caller may see, as Server-Sent Events



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
	resp, r, err := apiClient.TeamAPI.GetTeamEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamEvents`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamEventsRequest struct via the builder pattern


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


## GetTeamInbox

> TeamTeamInbox GetTeamInbox(ctx).Space(space).Archived(archived).Execute()

Returns the caller's notifications, newest first, each with the room or document it is about and the message that caused it.



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
	space := "space_example" // string | Space is the space uuid. Optional for a caller in exactly one space. (optional)
	archived := true // bool | Archived lists the archived notifications instead of the live ones. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamInbox(context.Background()).Space(space).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamInbox`: TeamTeamInbox
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **space** | **string** | Space is the space uuid. Optional for a caller in exactly one space. | 
 **archived** | **bool** | Archived lists the archived notifications instead of the live ones. | 

### Return type

[**TeamTeamInbox**](TeamTeamInbox.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMembers

> TeamTeamMembers GetTeamMembers(ctx).Space(space).Execute()

Returns the people and agents in a space, with the name, avatar, role and presence a client draws them with.



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
	space := "space_example" // string | Space is the space uuid. Optional for a caller in exactly one space. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamMembers(context.Background()).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMembers`: TeamTeamMembers
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **space** | **string** | Space is the space uuid. Optional for a caller in exactly one space. | 

### Return type

[**TeamTeamMembers**](TeamTeamMembers.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMessagesByIdReplies

> TeamTeamMessages GetTeamMessagesByIdReplies(ctx, id).Space(space).Execute()

Returns a message's thread, oldest first, each reply with its reactions and files.



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
	id := "id_example" // string | ID is the message, from the path.
	space := "space_example" // string | Space names the space holding it. A message id is unique within a space, not across the org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamMessagesByIdReplies(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamMessagesByIdReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMessagesByIdReplies`: TeamTeamMessages
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamMessagesByIdReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the message, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMessagesByIdRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space names the space holding it. A message id is unique within a space, not across the org. | 

### Return type

[**TeamTeamMessages**](TeamTeamMessages.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamPublic

> TeamPublicRooms GetTeamPublic(ctx).Q(q).Org(org).Limit(limit).Execute()

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
	// response from `GetTeamPublic`: TeamPublicRooms
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

[**TeamPublicRooms**](TeamPublicRooms.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamRooms

> TeamTeamRooms GetTeamRooms(ctx).Execute()

Returns the rooms the caller may see, with the kind and work facet each carries.



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
	// response from `GetTeamRooms`: TeamTeamRooms
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamRooms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRoomsRequest struct via the builder pattern


### Return type

[**TeamTeamRooms**](TeamTeamRooms.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamRoomsByIdMembers

> TeamTeamRoomMembers GetTeamRoomsByIdMembers(ctx, id).Space(space).Execute()

Returns the people and agents in one room, as the roster describes them.



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
	resp, r, err := apiClient.TeamAPI.GetTeamRoomsByIdMembers(context.Background(), id).Space(space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamRoomsByIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamRoomsByIdMembers`: TeamTeamRoomMembers
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamRoomsByIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room, from the path. The URL is the authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRoomsByIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **string** | Space names the space holding the room, and is required for the reason the bind op requires it: a room id is unique within a space and not across the org, so searching every space for a match would make the answer depend on iteration order. | 

### Return type

[**TeamTeamRoomMembers**](TeamTeamRoomMembers.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamRoomsByIdMessages

> TeamTeamMessages GetTeamRoomsByIdMessages(ctx, id).Space(space).Execute()

Returns the tail of one room's conversation, oldest first, each message with its reactions, files and thread count.



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
	// response from `GetTeamRoomsByIdMessages`: TeamTeamMessages
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

[**TeamTeamMessages**](TeamTeamMessages.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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

> TeamStatsOut GetTeamTransactorStatistics(ctx).Token(token).Execute()

Returns the transactor's live sessions for the space the caller's credential names — the endpoint the front's space switcher and server panel poll on the transactor base.



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
	// response from `GetTeamTransactorStatistics`: TeamStatsOut
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

[**TeamStatsOut**](TeamStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTeamDocsById

> TeamTeamDoc PatchTeamDocsById(ctx, id).TeamTeamDocEdit(teamTeamDocEdit).Execute()

Renames a document or moves it — under another document, to the top of its teamspace, or to another teamspace with everything nested under it.



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
	id := "id_example" // string | ID is the document, from the path.
	teamTeamDocEdit := *openapiclient.NewTeamTeamDocEdit() // TeamTeamDocEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PatchTeamDocsById(context.Background(), id).TeamTeamDocEdit(teamTeamDocEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PatchTeamDocsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTeamDocsById`: TeamTeamDoc
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PatchTeamDocsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamDocsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamDocEdit** | [**TeamTeamDocEdit**](TeamTeamDocEdit.md) |  | 

### Return type

[**TeamTeamDoc**](TeamTeamDoc.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTeamMessagesById

> TeamTeamMessage PatchTeamMessagesById(ctx, id).TeamTeamMessageEdit(teamTeamMessageEdit).Execute()

Rewrites what a message says.



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
	id := "id_example" // string | ID is the message, from the path.
	teamTeamMessageEdit := *openapiclient.NewTeamTeamMessageEdit() // TeamTeamMessageEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PatchTeamMessagesById(context.Background(), id).TeamTeamMessageEdit(teamTeamMessageEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PatchTeamMessagesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTeamMessagesById`: TeamTeamMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PatchTeamMessagesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the message, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamMessagesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamMessageEdit** | [**TeamTeamMessageEdit**](TeamTeamMessageEdit.md) |  | 

### Return type

[**TeamTeamMessage**](TeamTeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTeamRoomsById

> TeamTeamRoom PatchTeamRoomsById(ctx, id).TeamTeamRoomEdit(teamTeamRoomEdit).Execute()

Renames a channel, sets its topic, or archives or reopens it, and answers the room as it now stands.



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
	id := "id_example" // string | ID is the room, from the path.
	teamTeamRoomEdit := *openapiclient.NewTeamTeamRoomEdit() // TeamTeamRoomEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PatchTeamRoomsById(context.Background(), id).TeamTeamRoomEdit(teamTeamRoomEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PatchTeamRoomsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTeamRoomsById`: TeamTeamRoom
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PatchTeamRoomsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamRoomsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamRoomEdit** | [**TeamTeamRoomEdit**](TeamTeamRoomEdit.md) |  | 

### Return type

[**TeamTeamRoom**](TeamTeamRoom.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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


## PostTeamCollaboratorRpcByDocumentid

> TeamCollabResult PostTeamCollaboratorRpcByDocumentid(ctx, documentId).TeamCollabRequest(teamCollabRequest).Execute()

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
	teamCollabRequest := *openapiclient.NewTeamCollabRequest() // TeamCollabRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamCollaboratorRpcByDocumentid(context.Background(), documentId).TeamCollabRequest(teamCollabRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamCollaboratorRpcByDocumentid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamCollaboratorRpcByDocumentid`: TeamCollabResult
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

 **teamCollabRequest** | [**TeamCollabRequest**](TeamCollabRequest.md) |  | 

### Return type

[**TeamCollabResult**](TeamCollabResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamDms

> TeamTeamDirect PostTeamDms(ctx).TeamTeamDirectOpen(teamTeamDirectOpen).Execute()

Opens the direct message between the caller and the people named, or answers the one that already exists — there is one conversation per set of people, however many times it is asked for.



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
	teamTeamDirectOpen := *openapiclient.NewTeamTeamDirectOpen() // TeamTeamDirectOpen | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamDms(context.Background()).TeamTeamDirectOpen(teamTeamDirectOpen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamDms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamDms`: TeamTeamDirect
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamDms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamDmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamDirectOpen** | [**TeamTeamDirectOpen**](TeamTeamDirectOpen.md) |  | 

### Return type

[**TeamTeamDirect**](TeamTeamDirect.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamDocs

> TeamTeamDoc PostTeamDocs(ctx).TeamTeamDocNew(teamTeamDocNew).Execute()

Creates a document, as the caller, after its siblings.



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
	teamTeamDocNew := *openapiclient.NewTeamTeamDocNew() // TeamTeamDocNew | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamDocs(context.Background()).TeamTeamDocNew(teamTeamDocNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamDocs`: TeamTeamDoc
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamDocNew** | [**TeamTeamDocNew**](TeamTeamDocNew.md) |  | 

### Return type

[**TeamTeamDoc**](TeamTeamDoc.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamDocsByIdComments

> TeamTeamMessage PostTeamDocsByIdComments(ctx, id).TeamTeamCommentWrite(teamTeamCommentWrite).Execute()

Comments on a document, as the caller.



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
	id := "id_example" // string | ID is the document, from the path.
	teamTeamCommentWrite := *openapiclient.NewTeamTeamCommentWrite() // TeamTeamCommentWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamDocsByIdComments(context.Background(), id).TeamTeamCommentWrite(teamTeamCommentWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamDocsByIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamDocsByIdComments`: TeamTeamMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamDocsByIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the document, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamDocsByIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamCommentWrite** | [**TeamTeamCommentWrite**](TeamTeamCommentWrite.md) |  | 

### Return type

[**TeamTeamMessage**](TeamTeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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


## PostTeamInboxByIdArchive

> TeamTeamInboxItem PostTeamInboxByIdArchive(ctx, id).TeamTeamInboxAt(teamTeamInboxAt).Execute()

Archives one of the caller's notifications — read, and out of the live inbox — and answers it.



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
	id := "id_example" // string | ID is the notification, from the path.
	teamTeamInboxAt := *openapiclient.NewTeamTeamInboxAt() // TeamTeamInboxAt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamInboxByIdArchive(context.Background(), id).TeamTeamInboxAt(teamTeamInboxAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamInboxByIdArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamInboxByIdArchive`: TeamTeamInboxItem
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamInboxByIdArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the notification, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamInboxByIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamInboxAt** | [**TeamTeamInboxAt**](TeamTeamInboxAt.md) |  | 

### Return type

[**TeamTeamInboxItem**](TeamTeamInboxItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamInboxByIdRead

> TeamTeamInboxItem PostTeamInboxByIdRead(ctx, id).TeamTeamInboxAt(teamTeamInboxAt).Execute()

Marks one of the caller's notifications read and answers it.



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
	id := "id_example" // string | ID is the notification, from the path.
	teamTeamInboxAt := *openapiclient.NewTeamTeamInboxAt() // TeamTeamInboxAt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamInboxByIdRead(context.Background(), id).TeamTeamInboxAt(teamTeamInboxAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamInboxByIdRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamInboxByIdRead`: TeamTeamInboxItem
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamInboxByIdRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the notification, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamInboxByIdReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamInboxAt** | [**TeamTeamInboxAt**](TeamTeamInboxAt.md) |  | 

### Return type

[**TeamTeamInboxItem**](TeamTeamInboxItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamInboxRead

> TeamTeamInboxCleared PostTeamInboxRead(ctx).TeamTeamInboxAll(teamTeamInboxAll).Execute()

Marks every live notification of the caller's read, and says how many it changed.



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
	teamTeamInboxAll := *openapiclient.NewTeamTeamInboxAll() // TeamTeamInboxAll | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamInboxRead(context.Background()).TeamTeamInboxAll(teamTeamInboxAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamInboxRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamInboxRead`: TeamTeamInboxCleared
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamInboxRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamInboxReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamInboxAll** | [**TeamTeamInboxAll**](TeamTeamInboxAll.md) |  | 

### Return type

[**TeamTeamInboxCleared**](TeamTeamInboxCleared.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamMessagesByIdReplies

> TeamTeamMessage PostTeamMessagesByIdReplies(ctx, id).TeamTeamReplyWrite(teamTeamReplyWrite).Execute()

Answers a message in its thread, as the caller.



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
	id := "id_example" // string | ID is the message being answered, from the path.
	teamTeamReplyWrite := *openapiclient.NewTeamTeamReplyWrite() // TeamTeamReplyWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamMessagesByIdReplies(context.Background(), id).TeamTeamReplyWrite(teamTeamReplyWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamMessagesByIdReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamMessagesByIdReplies`: TeamTeamMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamMessagesByIdReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the message being answered, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamMessagesByIdRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamReplyWrite** | [**TeamTeamReplyWrite**](TeamTeamReplyWrite.md) |  | 

### Return type

[**TeamTeamMessage**](TeamTeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamRooms

> TeamTeamRoom PostTeamRooms(ctx).TeamTeamRoomNew(teamTeamRoomNew).Execute()

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
	teamTeamRoomNew := *openapiclient.NewTeamTeamRoomNew() // TeamTeamRoomNew | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamRooms(context.Background()).TeamTeamRoomNew(teamTeamRoomNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamRooms`: TeamTeamRoom
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamRooms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamRoomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamTeamRoomNew** | [**TeamTeamRoomNew**](TeamTeamRoomNew.md) |  | 

### Return type

[**TeamTeamRoom**](TeamTeamRoom.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamRoomsByIdMembers

> TeamTeamRoom PostTeamRoomsByIdMembers(ctx, id).TeamTeamRoomJoin(teamTeamRoomJoin).Execute()

Adds people to a room and answers the room as it now stands.



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
	id := "id_example" // string | ID is the room, from the path.
	teamTeamRoomJoin := *openapiclient.NewTeamTeamRoomJoin() // TeamTeamRoomJoin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamRoomsByIdMembers(context.Background(), id).TeamTeamRoomJoin(teamTeamRoomJoin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamRoomsByIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamRoomsByIdMembers`: TeamTeamRoom
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamRoomsByIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the room, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamRoomsByIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamTeamRoomJoin** | [**TeamTeamRoomJoin**](TeamTeamRoomJoin.md) |  | 

### Return type

[**TeamTeamRoom**](TeamTeamRoom.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamRoomsByIdMessages

> TeamTeamMessage PostTeamRoomsByIdMessages(ctx, id).TeamTeamMessageWrite(teamTeamMessageWrite).Execute()

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
	teamTeamMessageWrite := *openapiclient.NewTeamTeamMessageWrite() // TeamTeamMessageWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamRoomsByIdMessages(context.Background(), id).TeamTeamMessageWrite(teamTeamMessageWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamRoomsByIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamRoomsByIdMessages`: TeamTeamMessage
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

 **teamTeamMessageWrite** | [**TeamTeamMessageWrite**](TeamTeamMessageWrite.md) |  | 

### Return type

[**TeamTeamMessage**](TeamTeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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


## PutTeamMessagesByIdReactionsByEmoji

> TeamTeamMessage PutTeamMessagesByIdReactionsByEmoji(ctx, id, emoji).TeamTeamReactionWrite(teamTeamReactionWrite).Execute()

Adds the caller's reaction to a message and answers the message with its reactions as they now stand.



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
	id := "id_example" // string | ID is the message, from the path.
	emoji := "emoji_example" // string | Emoji is the reaction, from the path (percent-encoded on the wire).
	teamTeamReactionWrite := *openapiclient.NewTeamTeamReactionWrite() // TeamTeamReactionWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PutTeamMessagesByIdReactionsByEmoji(context.Background(), id, emoji).TeamTeamReactionWrite(teamTeamReactionWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PutTeamMessagesByIdReactionsByEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTeamMessagesByIdReactionsByEmoji`: TeamTeamMessage
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PutTeamMessagesByIdReactionsByEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the message, from the path. | 
**emoji** | **string** | Emoji is the reaction, from the path (percent-encoded on the wire). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTeamMessagesByIdReactionsByEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamTeamReactionWrite** | [**TeamTeamReactionWrite**](TeamTeamReactionWrite.md) |  | 

### Return type

[**TeamTeamMessage**](TeamTeamMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTeamRoomsById

> TeamTeamRoom PutTeamRoomsById(ctx, id).TeamTeamRoomBind(teamTeamRoomBind).Execute()

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
	teamTeamRoomBind := *openapiclient.NewTeamTeamRoomBind() // TeamTeamRoomBind | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PutTeamRoomsById(context.Background(), id).TeamTeamRoomBind(teamTeamRoomBind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PutTeamRoomsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTeamRoomsById`: TeamTeamRoom
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

 **teamTeamRoomBind** | [**TeamTeamRoomBind**](TeamTeamRoomBind.md) |  | 

### Return type

[**TeamTeamRoom**](TeamTeamRoom.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

