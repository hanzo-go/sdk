# \TeamAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTeamAccountCookie**](TeamAPI.md#DeleteTeamAccountCookie) | **Delete** /v1/team/account/cookie | Signs this browser out of team by expiring the HttpOnly account-token cookie the OAuth callback set.
[**DeleteTeamFilesByWorkspaceByFilename**](TeamAPI.md#DeleteTeamFilesByWorkspaceByFilename) | **Delete** /v1/team/files/{workspace}/{filename} | Removes one blob from a workspace&#39;s file store.
[**GetTeamAccountAuthByProvider**](TeamAPI.md#GetTeamAccountAuthByProvider) | **Get** /v1/team/account/auth/{provider} | Start a sign-in at hanzo.id
[**GetTeamAccountAuthByProviderCallback**](TeamAPI.md#GetTeamAccountAuthByProviderCallback) | **Get** /v1/team/account/auth/{provider}/callback | Complete a sign-in and hand the browser its session
[**GetTeamAccountProviders**](TeamAPI.md#GetTeamAccountProviders) | **Get** /v1/team/account/providers | Returns the identity providers this deployment starts a login with.
[**GetTeamBillingPlan**](TeamAPI.md#GetTeamBillingPlan) | **Get** /v1/team/billing/plan | Returns the plan and seat counts for the caller&#39;s OWN org, resolved from the VERIFIED team session token — never a client header.
[**GetTeamBillingUi**](TeamAPI.md#GetTeamBillingUi) | **Get** /v1/team/billing/ui | Open the wallet page
[**GetTeamBots**](TeamAPI.md#GetTeamBots) | **Get** /v1/team/bots | Returns the caller org&#39;s bot members — the org&#39;s agents projected as the workspace Employees they become, each with the member account uuid and Person reference the roster addresses it by.
[**GetTeamCollaborator**](TeamAPI.md#GetTeamCollaborator) | **Get** /v1/team/collaborator | Open the live collaborative-editing socket
[**GetTeamFilesByWorkspaceByFilename**](TeamAPI.md#GetTeamFilesByWorkspaceByFilename) | **Get** /v1/team/files/{workspace}/{filename} | Download a workspace file
[**GetTeamTransactorApiV1Statistics**](TeamAPI.md#GetTeamTransactorApiV1Statistics) | **Get** /v1/team/transactor/api/v1/statistics | Statistics returns the transactor&#39;s live sessions for the workspace the caller&#39;s credential names — the endpoint the front&#39;s workspace switcher and server panel poll on the transactor base.
[**GetTeamTransactorByToken**](TeamAPI.md#GetTeamTransactorByToken) | **Get** /v1/team/transactor/{token} | Open the workspace data-plane socket
[**GetTeamTransactorStatistics**](TeamAPI.md#GetTeamTransactorStatistics) | **Get** /v1/team/transactor/statistics | Statistics returns the transactor&#39;s live sessions for the workspace the caller&#39;s credential names — the endpoint the front&#39;s workspace switcher and server panel poll on the transactor base.
[**PostTeamAccount**](TeamAPI.md#PostTeamAccount) | **Post** /v1/team/account | Read the caller&#39;s account and switch workspace
[**PostTeamBotsSync**](TeamAPI.md#PostTeamBotsSync) | **Post** /v1/team/bots/sync | SyncBots re-projects the caller org&#39;s agents as workspace members into EVERY workspace of the org, and removes the ones whose agent is gone.
[**PostTeamCollaboratorRpcByDocumentid**](TeamAPI.md#PostTeamCollaboratorRpcByDocumentid) | **Post** /v1/team/collaborator/rpc/{documentId} | CollabRPC is the collaborative-markup snapshot plane the Team front&#39;s editor speaks: createContent stores a document field&#39;s markup at a fresh, immutable blob ref and returns it, updateContent stores a new snapshot and answers nothing, and getContent reads back the exact snapshot a ref names.
[**PostTeamFilesByWorkspace**](TeamAPI.md#PostTeamFilesByWorkspace) | **Post** /v1/team/files/{workspace} | Upload a file into a workspace
[**PutTeamAccountCookie**](TeamAPI.md#PutTeamAccountCookie) | **Put** /v1/team/account/cookie | Store the session token as this browser&#39;s cookie



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
	openapiclient "github.com/hanzoai/go-sdk"
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


## DeleteTeamFilesByWorkspaceByFilename

> DeleteTeamFilesByWorkspaceByFilename(ctx, workspace, filename).File(file).Execute()

Removes one blob from a workspace's file store.



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
	workspace := "6579…" // string | Workspace is the workspace uuid the blob belongs to, from the path.
	filename := "filename_example" // string | Filename is the last path segment, which the front sets to the blob id when it sends no explicit `file`.
	file := "0d4f…" // string | File is the blob id, and wins over the path segment when both are present. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeamFilesByWorkspaceByFilename(context.Background(), workspace, filename).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamFilesByWorkspaceByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | Workspace is the workspace uuid the blob belongs to, from the path. | 
**filename** | **string** | Filename is the last path segment, which the front sets to the blob id when it sends no explicit &#x60;file&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamFilesByWorkspaceByFilenameRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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
	openapiclient "github.com/hanzoai/go-sdk"
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

Returns the caller org's bot members — the org's agents projected as the workspace Employees they become, each with the member account uuid and Person reference the roster addresses it by.



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
	openapiclient "github.com/hanzoai/go-sdk"
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


## GetTeamFilesByWorkspaceByFilename

> *os.File GetTeamFilesByWorkspaceByFilename(ctx, workspace, filename).Execute()

Download a workspace file



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
	workspace := "workspace_example" // string | 
	filename := "filename_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamFilesByWorkspaceByFilename(context.Background(), workspace, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamFilesByWorkspaceByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamFilesByWorkspaceByFilename`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamFilesByWorkspaceByFilename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamFilesByWorkspaceByFilenameRequest struct via the builder pattern


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


## GetTeamTransactorApiV1Statistics

> StatsOut GetTeamTransactorApiV1Statistics(ctx).Token(token).Execute()

Statistics returns the transactor's live sessions for the workspace the caller's credential names — the endpoint the front's workspace switcher and server panel poll on the transactor base.



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
	token := "eyJhbGciOiJIUzI1NiJ9…" // string | Token is the workspace token minted by selectWorkspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamTransactorApiV1Statistics(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamTransactorApiV1Statistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamTransactorApiV1Statistics`: StatsOut
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamTransactorApiV1Statistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamTransactorApiV1StatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token is the workspace token minted by selectWorkspace. | 

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


## GetTeamTransactorByToken

> GetTeamTransactorByToken(ctx, token).Execute()

Open the workspace data-plane socket



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

Statistics returns the transactor's live sessions for the workspace the caller's credential names — the endpoint the front's workspace switcher and server panel poll on the transactor base.



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
	token := "eyJhbGciOiJIUzI1NiJ9…" // string | Token is the workspace token minted by selectWorkspace. (optional)

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
 **token** | **string** | Token is the workspace token minted by selectWorkspace. | 

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

Read the caller's account and switch workspace



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

SyncBots re-projects the caller org's agents as workspace members into EVERY workspace of the org, and removes the ones whose agent is gone.



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
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	documentId := "6579…|tracker:class:Issue|issue-1|description" // string | DocumentID addresses the document field, as \"<workspaceUuid>|<objectClass>|<objectId>|<objectAttr>\" — the collaborator-client encodeDocumentId shape, from the path.
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
**documentId** | **string** | DocumentID addresses the document field, as \&quot;&lt;workspaceUuid&gt;|&lt;objectClass&gt;|&lt;objectId&gt;|&lt;objectAttr&gt;\&quot; — the collaborator-client encodeDocumentId shape, from the path. | 

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


## PostTeamFilesByWorkspace

> *os.File PostTeamFilesByWorkspace(ctx, workspace).Body(body).Execute()

Upload a file into a workspace



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
	workspace := "workspace_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.PostTeamFilesByWorkspace(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.PostTeamFilesByWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTeamFilesByWorkspace`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.PostTeamFilesByWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamFilesByWorkspaceRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk"
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

