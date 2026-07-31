# \TeamAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1TeamAccountCookie**](TeamAPI.md#CloudDeleteV1TeamAccountCookie) | **Delete** /v1/team/account/cookie | ClearCookie signs this browser out of team by expiring the HttpOnly account-token cookie the OAuth callback set.
[**CloudDeleteV1TeamFilesWorkspaceFilename**](TeamAPI.md#CloudDeleteV1TeamFilesWorkspaceFilename) | **Delete** /v1/team/files/{workspace}/{filename} | DeleteBlob removes one blob from a workspace&#39;s file store.
[**CloudGetV1TeamAccountAuthByProvider**](TeamAPI.md#CloudGetV1TeamAccountAuthByProvider) | **Get** /v1/team/account/auth/{provider} | 
[**CloudGetV1TeamAccountAuthByProviderCallback**](TeamAPI.md#CloudGetV1TeamAccountAuthByProviderCallback) | **Get** /v1/team/account/auth/{provider}/callback | 
[**CloudGetV1TeamAccountProviders**](TeamAPI.md#CloudGetV1TeamAccountProviders) | **Get** /v1/team/account/providers | ListProviders returns the identity providers this deployment starts a login with.
[**CloudGetV1TeamBillingPlan**](TeamAPI.md#CloudGetV1TeamBillingPlan) | **Get** /v1/team/billing/plan | ReadPlan returns the plan and seat counts for the caller&#39;s OWN org, resolved from the VERIFIED team session token — never a client header.
[**CloudGetV1TeamBillingUi**](TeamAPI.md#CloudGetV1TeamBillingUi) | **Get** /v1/team/billing/ui | 
[**CloudGetV1TeamBillingUiByWildcard1**](TeamAPI.md#CloudGetV1TeamBillingUiByWildcard1) | **Get** /v1/team/billing/ui/{wildcard1} | 
[**CloudGetV1TeamBots**](TeamAPI.md#CloudGetV1TeamBots) | **Get** /v1/team/bots | ListBots returns the caller org&#39;s bot members — the org&#39;s agents projected as the workspace Employees they become, each with the member account uuid and Person reference the roster addresses it by.
[**CloudGetV1TeamFilesByWorkspaceByFilename**](TeamAPI.md#CloudGetV1TeamFilesByWorkspaceByFilename) | **Get** /v1/team/files/{workspace}/{filename} | 
[**CloudGetV1TeamTransactorApiV1Statistics**](TeamAPI.md#CloudGetV1TeamTransactorApiV1Statistics) | **Get** /v1/team/transactor/api/v1/statistics | Statistics returns the transactor&#39;s live sessions for the workspace the caller&#39;s token names — the endpoint the front&#39;s workspace switcher and server panel poll on the transactor base.
[**CloudGetV1TeamTransactorByToken**](TeamAPI.md#CloudGetV1TeamTransactorByToken) | **Get** /v1/team/transactor/{token} | 
[**CloudGetV1TeamTransactorStatistics**](TeamAPI.md#CloudGetV1TeamTransactorStatistics) | **Get** /v1/team/transactor/statistics | Statistics returns the transactor&#39;s live sessions for the workspace the caller&#39;s token names — the endpoint the front&#39;s workspace switcher and server panel poll on the transactor base.
[**CloudPostV1TeamAccount**](TeamAPI.md#CloudPostV1TeamAccount) | **Post** /v1/team/account | 
[**CloudPostV1TeamBotsSync**](TeamAPI.md#CloudPostV1TeamBotsSync) | **Post** /v1/team/bots/sync | SyncBots re-projects the caller org&#39;s agents as workspace members into EVERY workspace of the org, and removes the ones whose agent is gone.
[**CloudPostV1TeamFilesByWorkspace**](TeamAPI.md#CloudPostV1TeamFilesByWorkspace) | **Post** /v1/team/files/{workspace} | 
[**CloudPutV1TeamAccountCookie**](TeamAPI.md#CloudPutV1TeamAccountCookie) | **Put** /v1/team/account/cookie | 



## CloudDeleteV1TeamAccountCookie

> CloudCookieAck CloudDeleteV1TeamAccountCookie(ctx).Execute()

ClearCookie signs this browser out of team by expiring the HttpOnly account-token cookie the OAuth callback set.



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
	resp, r, err := apiClient.TeamAPI.CloudDeleteV1TeamAccountCookie(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudDeleteV1TeamAccountCookie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1TeamAccountCookie`: CloudCookieAck
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudDeleteV1TeamAccountCookie`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1TeamAccountCookieRequest struct via the builder pattern


### Return type

[**CloudCookieAck**](CloudCookieAck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1TeamFilesWorkspaceFilename

> CloudDeleteV1TeamFilesWorkspaceFilename(ctx, workspace, filename).File(file).Execute()

DeleteBlob removes one blob from a workspace's file store.



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
	r, err := apiClient.TeamAPI.CloudDeleteV1TeamFilesWorkspaceFilename(context.Background(), workspace, filename).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudDeleteV1TeamFilesWorkspaceFilename``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1TeamFilesWorkspaceFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **string** | File is the blob id, and wins over the path segment when both are present. | 

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


## CloudGetV1TeamAccountAuthByProvider

> CloudGetV1TeamAccountAuthByProvider(ctx, provider).Execute()



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
	r, err := apiClient.TeamAPI.CloudGetV1TeamAccountAuthByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamAccountAuthByProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1TeamAccountAuthByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1TeamAccountAuthByProviderCallback

> CloudGetV1TeamAccountAuthByProviderCallback(ctx, provider).Execute()



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
	r, err := apiClient.TeamAPI.CloudGetV1TeamAccountAuthByProviderCallback(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamAccountAuthByProviderCallback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1TeamAccountAuthByProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1TeamAccountProviders

> []CloudProviderInfo CloudGetV1TeamAccountProviders(ctx).Execute()

ListProviders returns the identity providers this deployment starts a login with.



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
	resp, r, err := apiClient.TeamAPI.CloudGetV1TeamAccountProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamAccountProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TeamAccountProviders`: []CloudProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudGetV1TeamAccountProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamAccountProvidersRequest struct via the builder pattern


### Return type

[**[]CloudProviderInfo**](CloudProviderInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TeamBillingPlan

> CloudPlanInfo CloudGetV1TeamBillingPlan(ctx).Execute()

ReadPlan returns the plan and seat counts for the caller's OWN org, resolved from the VERIFIED team session token — never a client header.



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
	resp, r, err := apiClient.TeamAPI.CloudGetV1TeamBillingPlan(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TeamBillingPlan`: CloudPlanInfo
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudGetV1TeamBillingPlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamBillingPlanRequest struct via the builder pattern


### Return type

[**CloudPlanInfo**](CloudPlanInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TeamBillingUi

> CloudGetV1TeamBillingUi(ctx).Execute()



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
	r, err := apiClient.TeamAPI.CloudGetV1TeamBillingUi(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamBillingUi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamBillingUiRequest struct via the builder pattern


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


## CloudGetV1TeamBillingUiByWildcard1

> CloudGetV1TeamBillingUiByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.CloudGetV1TeamBillingUiByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamBillingUiByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamBillingUiByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1TeamBots

> CloudBotRoster CloudGetV1TeamBots(ctx).Execute()

ListBots returns the caller org's bot members — the org's agents projected as the workspace Employees they become, each with the member account uuid and Person reference the roster addresses it by.



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
	resp, r, err := apiClient.TeamAPI.CloudGetV1TeamBots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TeamBots`: CloudBotRoster
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudGetV1TeamBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamBotsRequest struct via the builder pattern


### Return type

[**CloudBotRoster**](CloudBotRoster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TeamFilesByWorkspaceByFilename

> CloudGetV1TeamFilesByWorkspaceByFilename(ctx, workspace, filename).Execute()



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
	r, err := apiClient.TeamAPI.CloudGetV1TeamFilesByWorkspaceByFilename(context.Background(), workspace, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamFilesByWorkspaceByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamFilesByWorkspaceByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CloudGetV1TeamTransactorApiV1Statistics

> CloudStatsOut CloudGetV1TeamTransactorApiV1Statistics(ctx).Token(token).Execute()

Statistics returns the transactor's live sessions for the workspace the caller's token names — the endpoint the front's workspace switcher and server panel poll on the transactor base.



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
	resp, r, err := apiClient.TeamAPI.CloudGetV1TeamTransactorApiV1Statistics(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamTransactorApiV1Statistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TeamTransactorApiV1Statistics`: CloudStatsOut
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudGetV1TeamTransactorApiV1Statistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamTransactorApiV1StatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token is the workspace token minted by selectWorkspace. | 

### Return type

[**CloudStatsOut**](CloudStatsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TeamTransactorByToken

> CloudGetV1TeamTransactorByToken(ctx, token).Execute()



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
	r, err := apiClient.TeamAPI.CloudGetV1TeamTransactorByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamTransactorByToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1TeamTransactorByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetV1TeamTransactorStatistics

> CloudStatsOut CloudGetV1TeamTransactorStatistics(ctx).Token(token).Execute()

Statistics returns the transactor's live sessions for the workspace the caller's token names — the endpoint the front's workspace switcher and server panel poll on the transactor base.



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
	resp, r, err := apiClient.TeamAPI.CloudGetV1TeamTransactorStatistics(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudGetV1TeamTransactorStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TeamTransactorStatistics`: CloudStatsOut
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudGetV1TeamTransactorStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TeamTransactorStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token is the workspace token minted by selectWorkspace. | 

### Return type

[**CloudStatsOut**](CloudStatsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1TeamAccount

> CloudPostV1TeamAccount(ctx).Execute()



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
	r, err := apiClient.TeamAPI.CloudPostV1TeamAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudPostV1TeamAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TeamAccountRequest struct via the builder pattern


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


## CloudPostV1TeamBotsSync

> CloudBotSync CloudPostV1TeamBotsSync(ctx).Execute()

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
	resp, r, err := apiClient.TeamAPI.CloudPostV1TeamBotsSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudPostV1TeamBotsSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1TeamBotsSync`: CloudBotSync
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CloudPostV1TeamBotsSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TeamBotsSyncRequest struct via the builder pattern


### Return type

[**CloudBotSync**](CloudBotSync.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1TeamFilesByWorkspace

> CloudPostV1TeamFilesByWorkspace(ctx, workspace).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.CloudPostV1TeamFilesByWorkspace(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudPostV1TeamFilesByWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TeamFilesByWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudPutV1TeamAccountCookie

> CloudPutV1TeamAccountCookie(ctx).Execute()



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
	r, err := apiClient.TeamAPI.CloudPutV1TeamAccountCookie(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CloudPutV1TeamAccountCookie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1TeamAccountCookieRequest struct via the builder pattern


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

