# \IntegrationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotGetIntegration**](IntegrationsAPI.md#BotGetIntegration) | **Get** /v1/bot/integrations/{slug} | Get integration detail with latest version
[**BotListIntegrations**](IntegrationsAPI.md#BotListIntegrations) | **Get** /v1/bot/integrations | List integrations (paginated)
[**CloudDeleteV1IntegrationsGithubReposRepoPages**](IntegrationsAPI.md#CloudDeleteV1IntegrationsGithubReposRepoPages) | **Delete** /v1/integrations/github/repos/{repo}/pages | Deletes the repo&#39;s Pages site.
[**CloudGetV1Integrations**](IntegrationsAPI.md#CloudGetV1Integrations) | **Get** /v1/integrations | Returns every registered integration provider together with THIS org&#39;s connection status for it — the catalog the console&#39;s Integrations page renders.
[**CloudGetV1IntegrationsByProviderCallback**](IntegrationsAPI.md#CloudGetV1IntegrationsByProviderCallback) | **Get** /v1/integrations/{provider}/callback | 
[**CloudGetV1IntegrationsDiscordLink**](IntegrationsAPI.md#CloudGetV1IntegrationsDiscordLink) | **Get** /v1/integrations/discord/link | 
[**CloudGetV1IntegrationsDiscordLinkCallback**](IntegrationsAPI.md#CloudGetV1IntegrationsDiscordLinkCallback) | **Get** /v1/integrations/discord/link/callback | 
[**CloudGetV1IntegrationsDiscordLinkDiscord**](IntegrationsAPI.md#CloudGetV1IntegrationsDiscordLinkDiscord) | **Get** /v1/integrations/discord/link/discord | 
[**CloudGetV1IntegrationsGithubRepos**](IntegrationsAPI.md#CloudGetV1IntegrationsGithubRepos) | **Get** /v1/integrations/github/repos | Lists the org&#39;s granted GitHub repositories, each annotated with its native import + sync status from the git object plane.
[**CloudGetV1IntegrationsGithubReposRepoPages**](IntegrationsAPI.md#CloudGetV1IntegrationsGithubReposRepoPages) | **Get** /v1/integrations/github/repos/{repo}/pages | Returns the repo&#39;s Pages status, live URL, custom domain and build source.
[**CloudGetV1IntegrationsProvider**](IntegrationsAPI.md#CloudGetV1IntegrationsProvider) | **Get** /v1/integrations/{provider} | Returns ONE provider with this org&#39;s connection status — the same view list carries, for a single id.
[**CloudGetV1IntegrationsSlackLink**](IntegrationsAPI.md#CloudGetV1IntegrationsSlackLink) | **Get** /v1/integrations/slack/link | 
[**CloudGetV1IntegrationsSlackLinkCallback**](IntegrationsAPI.md#CloudGetV1IntegrationsSlackLinkCallback) | **Get** /v1/integrations/slack/link/callback | 
[**CloudGetV1IntegrationsSlackLinkSlack**](IntegrationsAPI.md#CloudGetV1IntegrationsSlackLinkSlack) | **Get** /v1/integrations/slack/link/slack | 
[**CloudGetV1IntegrationsTeamsLink**](IntegrationsAPI.md#CloudGetV1IntegrationsTeamsLink) | **Get** /v1/integrations/teams/link | 
[**CloudGetV1IntegrationsTeamsLinkAad**](IntegrationsAPI.md#CloudGetV1IntegrationsTeamsLinkAad) | **Get** /v1/integrations/teams/link/aad | 
[**CloudGetV1IntegrationsTeamsLinkCallback**](IntegrationsAPI.md#CloudGetV1IntegrationsTeamsLinkCallback) | **Get** /v1/integrations/teams/link/callback | 
[**CloudGetV1IntegrationsTelegramLink**](IntegrationsAPI.md#CloudGetV1IntegrationsTelegramLink) | **Get** /v1/integrations/telegram/link | 
[**CloudGetV1IntegrationsTelegramLinkAuth**](IntegrationsAPI.md#CloudGetV1IntegrationsTelegramLinkAuth) | **Get** /v1/integrations/telegram/link/auth | 
[**CloudGetV1IntegrationsTelegramLinkCallback**](IntegrationsAPI.md#CloudGetV1IntegrationsTelegramLinkCallback) | **Get** /v1/integrations/telegram/link/callback | 
[**CloudPostV1IntegrationsDiscordInteractions**](IntegrationsAPI.md#CloudPostV1IntegrationsDiscordInteractions) | **Post** /v1/integrations/discord/interactions | 
[**CloudPostV1IntegrationsGithubIssuesBackfill**](IntegrationsAPI.md#CloudPostV1IntegrationsGithubIssuesBackfill) | **Post** /v1/integrations/github/issues/backfill | Seeds the native tracker with the EXISTING issues across the org&#39;s granted repos (default state&#x3D;open); the webhook keeps them live thereafter.
[**CloudPostV1IntegrationsGithubReposImport**](IntegrationsAPI.md#CloudPostV1IntegrationsGithubReposImport) | **Post** /v1/integrations/github/repos/import | GithubImport imports the selected (or all) granted repos into git.hanzo.ai.
[**CloudPostV1IntegrationsGithubReposRepoPages**](IntegrationsAPI.md#CloudPostV1IntegrationsGithubReposRepoPages) | **Post** /v1/integrations/github/repos/{repo}/pages | Creates the repo&#39;s Pages site and answers 201 Created with it.
[**CloudPostV1IntegrationsGithubReposRepoPagesBuilds**](IntegrationsAPI.md#CloudPostV1IntegrationsGithubReposRepoPagesBuilds) | **Post** /v1/integrations/github/repos/{repo}/pages/builds | GithubPagesBuild requests a Pages rebuild and returns the queued build&#39;s status.
[**CloudPostV1IntegrationsProviderConnect**](IntegrationsAPI.md#CloudPostV1IntegrationsProviderConnect) | **Post** /v1/integrations/{provider}/connect | Acquires the org&#39;s credential for one provider.
[**CloudPostV1IntegrationsProviderDisconnect**](IntegrationsAPI.md#CloudPostV1IntegrationsProviderDisconnect) | **Post** /v1/integrations/{provider}/disconnect | Revokes (best-effort) and forgets an org&#39;s connection: it deletes every custodied KMS secret and the connection row.
[**CloudPostV1IntegrationsProviderVerify**](IntegrationsAPI.md#CloudPostV1IntegrationsProviderVerify) | **Post** /v1/integrations/{provider}/verify | Re-checks a CONNECTED apikey connector&#39;s stored credential against the provider, live (&#x60;hanzo connector verify&#x60;).
[**CloudPostV1IntegrationsSlackCommands**](IntegrationsAPI.md#CloudPostV1IntegrationsSlackCommands) | **Post** /v1/integrations/slack/commands | 
[**CloudPostV1IntegrationsSlackEvents**](IntegrationsAPI.md#CloudPostV1IntegrationsSlackEvents) | **Post** /v1/integrations/slack/events | 
[**CloudPostV1IntegrationsTeamsEvents**](IntegrationsAPI.md#CloudPostV1IntegrationsTeamsEvents) | **Post** /v1/integrations/teams/events | 
[**CloudPostV1IntegrationsTelegramConnect**](IntegrationsAPI.md#CloudPostV1IntegrationsTelegramConnect) | **Post** /v1/integrations/telegram/connect | Mints a short, single-use deep-link code bound to the caller&#39;s org and returns the t.me link the console navigates to.
[**CloudPostV1IntegrationsTelegramWebhook**](IntegrationsAPI.md#CloudPostV1IntegrationsTelegramWebhook) | **Post** /v1/integrations/telegram/webhook | 
[**CloudPutV1IntegrationsGithubReposRepoPages**](IntegrationsAPI.md#CloudPutV1IntegrationsGithubReposRepoPages) | **Put** /v1/integrations/github/repos/{repo}/pages | Sets or clears the custom domain (cname) and updates HTTPS enforcement, build type, or source.



## BotGetIntegration

> BotGetIntegration200Response BotGetIntegration(ctx, slug).Execute()

Get integration detail with latest version

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.BotGetIntegration(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.BotGetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetIntegration`: BotGetIntegration200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.BotGetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotGetIntegration200Response**](BotGetIntegration200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListIntegrations

> BotListIntegrations200Response BotListIntegrations(ctx).Sort(sort).Limit(limit).Cursor(cursor).Execute()

List integrations (paginated)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	sort := "sort_example" // string |  (optional) (default to "updated")
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.BotListIntegrations(context.Background()).Sort(sort).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.BotListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListIntegrations`: BotListIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.BotListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** |  | [default to &quot;updated&quot;]
 **limit** | **int32** |  | [default to 50]
 **cursor** | **time.Time** |  | 

### Return type

[**BotListIntegrations200Response**](BotListIntegrations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1IntegrationsGithubReposRepoPages

> CloudGithubPagesDisabledOut CloudDeleteV1IntegrationsGithubReposRepoPages(ctx, repo).Execute()

Deletes the repo's Pages site.



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
	repo := "widgets" // string | Repo is the repository's short name within the org's installation, with no owner prefix (the owner is server-derived from the grant). A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudDeleteV1IntegrationsGithubReposRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudDeleteV1IntegrationsGithubReposRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1IntegrationsGithubReposRepoPages`: CloudGithubPagesDisabledOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudDeleteV1IntegrationsGithubReposRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1IntegrationsGithubReposRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudGithubPagesDisabledOut**](CloudGithubPagesDisabledOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Integrations

> CloudListOut CloudGetV1Integrations(ctx).Execute()

Returns every registered integration provider together with THIS org's connection status for it — the catalog the console's Integrations page renders.



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
	resp, r, err := apiClient.IntegrationsAPI.CloudGetV1Integrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1Integrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Integrations`: CloudListOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudGetV1Integrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsRequest struct via the builder pattern


### Return type

[**CloudListOut**](CloudListOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IntegrationsByProviderCallback

> CloudGetV1IntegrationsByProviderCallback(ctx, provider).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsByProviderCallback(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsByProviderCallback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsByProviderCallbackRequest struct via the builder pattern


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


## CloudGetV1IntegrationsDiscordLink

> CloudGetV1IntegrationsDiscordLink(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsDiscordLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsDiscordLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsDiscordLinkRequest struct via the builder pattern


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


## CloudGetV1IntegrationsDiscordLinkCallback

> CloudGetV1IntegrationsDiscordLinkCallback(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsDiscordLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsDiscordLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsDiscordLinkCallbackRequest struct via the builder pattern


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


## CloudGetV1IntegrationsDiscordLinkDiscord

> CloudGetV1IntegrationsDiscordLinkDiscord(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsDiscordLinkDiscord(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsDiscordLinkDiscord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsDiscordLinkDiscordRequest struct via the builder pattern


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


## CloudGetV1IntegrationsGithubRepos

> CloudGithubReposOut CloudGetV1IntegrationsGithubRepos(ctx).Execute()

Lists the org's granted GitHub repositories, each annotated with its native import + sync status from the git object plane.



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
	resp, r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsGithubRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsGithubRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IntegrationsGithubRepos`: CloudGithubReposOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudGetV1IntegrationsGithubRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsGithubReposRequest struct via the builder pattern


### Return type

[**CloudGithubReposOut**](CloudGithubReposOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IntegrationsGithubReposRepoPages

> CloudGithubPagesView CloudGetV1IntegrationsGithubReposRepoPages(ctx, repo).Execute()

Returns the repo's Pages status, live URL, custom domain and build source.



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
	repo := "widgets" // string | Repo is the repository's short name within the org's installation, with no owner prefix (the owner is server-derived from the grant). A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsGithubReposRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsGithubReposRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IntegrationsGithubReposRepoPages`: CloudGithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudGetV1IntegrationsGithubReposRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsGithubReposRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudGithubPagesView**](CloudGithubPagesView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IntegrationsProvider

> CloudProviderView CloudGetV1IntegrationsProvider(ctx, provider).Execute()

Returns ONE provider with this org's connection status — the same view list carries, for a single id.



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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IntegrationsProvider`: CloudProviderView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudGetV1IntegrationsProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProviderView**](CloudProviderView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IntegrationsSlackLink

> CloudGetV1IntegrationsSlackLink(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsSlackLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsSlackLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsSlackLinkRequest struct via the builder pattern


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


## CloudGetV1IntegrationsSlackLinkCallback

> CloudGetV1IntegrationsSlackLinkCallback(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsSlackLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsSlackLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsSlackLinkCallbackRequest struct via the builder pattern


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


## CloudGetV1IntegrationsSlackLinkSlack

> CloudGetV1IntegrationsSlackLinkSlack(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsSlackLinkSlack(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsSlackLinkSlack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsSlackLinkSlackRequest struct via the builder pattern


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


## CloudGetV1IntegrationsTeamsLink

> CloudGetV1IntegrationsTeamsLink(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsTeamsLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsTeamsLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsTeamsLinkRequest struct via the builder pattern


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


## CloudGetV1IntegrationsTeamsLinkAad

> CloudGetV1IntegrationsTeamsLinkAad(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsTeamsLinkAad(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsTeamsLinkAad``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsTeamsLinkAadRequest struct via the builder pattern


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


## CloudGetV1IntegrationsTeamsLinkCallback

> CloudGetV1IntegrationsTeamsLinkCallback(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsTeamsLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsTeamsLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsTeamsLinkCallbackRequest struct via the builder pattern


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


## CloudGetV1IntegrationsTelegramLink

> CloudGetV1IntegrationsTelegramLink(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsTelegramLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsTelegramLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsTelegramLinkRequest struct via the builder pattern


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


## CloudGetV1IntegrationsTelegramLinkAuth

> CloudGetV1IntegrationsTelegramLinkAuth(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsTelegramLinkAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsTelegramLinkAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsTelegramLinkAuthRequest struct via the builder pattern


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


## CloudGetV1IntegrationsTelegramLinkCallback

> CloudGetV1IntegrationsTelegramLinkCallback(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudGetV1IntegrationsTelegramLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudGetV1IntegrationsTelegramLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IntegrationsTelegramLinkCallbackRequest struct via the builder pattern


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


## CloudPostV1IntegrationsDiscordInteractions

> CloudPostV1IntegrationsDiscordInteractions(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsDiscordInteractions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsDiscordInteractions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsDiscordInteractionsRequest struct via the builder pattern


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


## CloudPostV1IntegrationsGithubIssuesBackfill

> CloudGithubBackfillResult CloudPostV1IntegrationsGithubIssuesBackfill(ctx).CloudGithubBackfillIn(cloudGithubBackfillIn).Execute()

Seeds the native tracker with the EXISTING issues across the org's granted repos (default state=open); the webhook keeps them live thereafter.



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
	cloudGithubBackfillIn := *openapiclient.NewCloudGithubBackfillIn() // CloudGithubBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsGithubIssuesBackfill(context.Background()).CloudGithubBackfillIn(cloudGithubBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsGithubIssuesBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsGithubIssuesBackfill`: CloudGithubBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsGithubIssuesBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsGithubIssuesBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudGithubBackfillIn** | [**CloudGithubBackfillIn**](CloudGithubBackfillIn.md) |  | 

### Return type

[**CloudGithubBackfillResult**](CloudGithubBackfillResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsGithubReposImport

> CloudGithubImportOut CloudPostV1IntegrationsGithubReposImport(ctx).CloudGithubImportIn(cloudGithubImportIn).Execute()

GithubImport imports the selected (or all) granted repos into git.hanzo.ai.



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
	cloudGithubImportIn := *openapiclient.NewCloudGithubImportIn() // CloudGithubImportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsGithubReposImport(context.Background()).CloudGithubImportIn(cloudGithubImportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsGithubReposImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsGithubReposImport`: CloudGithubImportOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsGithubReposImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsGithubReposImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudGithubImportIn** | [**CloudGithubImportIn**](CloudGithubImportIn.md) |  | 

### Return type

[**CloudGithubImportOut**](CloudGithubImportOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsGithubReposRepoPages

> CloudGithubPagesView CloudPostV1IntegrationsGithubReposRepoPages(ctx, repo).CloudGithubPagesEnableReq(cloudGithubPagesEnableReq).Execute()

Creates the repo's Pages site and answers 201 Created with it.



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
	repo := "widgets" // string | Repo is the repository, from the :repo path segment.
	cloudGithubPagesEnableReq := *openapiclient.NewCloudGithubPagesEnableReq() // CloudGithubPagesEnableReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsGithubReposRepoPages(context.Background(), repo).CloudGithubPagesEnableReq(cloudGithubPagesEnableReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsGithubReposRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsGithubReposRepoPages`: CloudGithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsGithubReposRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsGithubReposRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGithubPagesEnableReq** | [**CloudGithubPagesEnableReq**](CloudGithubPagesEnableReq.md) |  | 

### Return type

[**CloudGithubPagesView**](CloudGithubPagesView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsGithubReposRepoPagesBuilds

> CloudGithubPagesBuildOut CloudPostV1IntegrationsGithubReposRepoPagesBuilds(ctx, repo).Execute()

GithubPagesBuild requests a Pages rebuild and returns the queued build's status.



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
	repo := "widgets" // string | Repo is the repository's short name within the org's installation, with no owner prefix (the owner is server-derived from the grant). A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsGithubReposRepoPagesBuilds(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsGithubReposRepoPagesBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsGithubReposRepoPagesBuilds`: CloudGithubPagesBuildOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsGithubReposRepoPagesBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsGithubReposRepoPagesBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudGithubPagesBuildOut**](CloudGithubPagesBuildOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsProviderConnect

> CloudConnectOut CloudPostV1IntegrationsProviderConnect(ctx, provider).CloudConnectIn(cloudConnectIn).Execute()

Acquires the org's credential for one provider.



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
	provider := "cloudflare" // string | Provider is the connector's registry id, from the :provider path segment.
	cloudConnectIn := *openapiclient.NewCloudConnectIn() // CloudConnectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsProviderConnect(context.Background(), provider).CloudConnectIn(cloudConnectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsProviderConnect`: CloudConnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsProviderConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudConnectIn** | [**CloudConnectIn**](CloudConnectIn.md) |  | 

### Return type

[**CloudConnectOut**](CloudConnectOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsProviderDisconnect

> CloudDisconnectOut CloudPostV1IntegrationsProviderDisconnect(ctx, provider).Execute()

Revokes (best-effort) and forgets an org's connection: it deletes every custodied KMS secret and the connection row.



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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsProviderDisconnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsProviderDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsProviderDisconnect`: CloudDisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsProviderDisconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsProviderDisconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDisconnectOut**](CloudDisconnectOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsProviderVerify

> CloudVerifyOut CloudPostV1IntegrationsProviderVerify(ctx, provider).Execute()

Re-checks a CONNECTED apikey connector's stored credential against the provider, live (`hanzo connector verify`).



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
	provider := "cloudflare" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsProviderVerify(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsProviderVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsProviderVerify`: CloudVerifyOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsProviderVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsProviderVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudVerifyOut**](CloudVerifyOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsSlackCommands

> CloudPostV1IntegrationsSlackCommands(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsSlackCommands(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsSlackCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsSlackCommandsRequest struct via the builder pattern


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


## CloudPostV1IntegrationsSlackEvents

> CloudPostV1IntegrationsSlackEvents(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsSlackEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsSlackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsSlackEventsRequest struct via the builder pattern


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


## CloudPostV1IntegrationsTeamsEvents

> CloudPostV1IntegrationsTeamsEvents(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsTeamsEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsTeamsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsTeamsEventsRequest struct via the builder pattern


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


## CloudPostV1IntegrationsTelegramConnect

> CloudAuthorizeOut CloudPostV1IntegrationsTelegramConnect(ctx).Execute()

Mints a short, single-use deep-link code bound to the caller's org and returns the t.me link the console navigates to.



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
	resp, r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsTelegramConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsTelegramConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IntegrationsTelegramConnect`: CloudAuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPostV1IntegrationsTelegramConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsTelegramConnectRequest struct via the builder pattern


### Return type

[**CloudAuthorizeOut**](CloudAuthorizeOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IntegrationsTelegramWebhook

> CloudPostV1IntegrationsTelegramWebhook(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.CloudPostV1IntegrationsTelegramWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPostV1IntegrationsTelegramWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IntegrationsTelegramWebhookRequest struct via the builder pattern


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


## CloudPutV1IntegrationsGithubReposRepoPages

> CloudGithubPagesUpdatedOut CloudPutV1IntegrationsGithubReposRepoPages(ctx, repo).CloudGithubPagesUpdateReq(cloudGithubPagesUpdateReq).Execute()

Sets or clears the custom domain (cname) and updates HTTPS enforcement, build type, or source.



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
	repo := "widgets" // string | Repo is the repository, from the :repo path segment.
	cloudGithubPagesUpdateReq := *openapiclient.NewCloudGithubPagesUpdateReq() // CloudGithubPagesUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CloudPutV1IntegrationsGithubReposRepoPages(context.Background(), repo).CloudGithubPagesUpdateReq(cloudGithubPagesUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CloudPutV1IntegrationsGithubReposRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1IntegrationsGithubReposRepoPages`: CloudGithubPagesUpdatedOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CloudPutV1IntegrationsGithubReposRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1IntegrationsGithubReposRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGithubPagesUpdateReq** | [**CloudGithubPagesUpdateReq**](CloudGithubPagesUpdateReq.md) |  | 

### Return type

[**CloudGithubPagesUpdatedOut**](CloudGithubPagesUpdatedOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

