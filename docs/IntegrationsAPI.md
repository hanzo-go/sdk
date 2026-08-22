# \IntegrationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIntegrationsConnectorsById**](IntegrationsAPI.md#DeleteIntegrationsConnectorsById) | **Delete** /v1/integrations/connectors/{id} | Forgets a connector: every custodied secret, then the row.
[**DeleteIntegrationsGithubReposByRepoPages**](IntegrationsAPI.md#DeleteIntegrationsGithubReposByRepoPages) | **Delete** /v1/integrations/github/repos/{repo}/pages | Deletes the repo&#39;s Pages site.
[**GetIntegrations**](IntegrationsAPI.md#GetIntegrations) | **Get** /v1/integrations | Returns every registered integration provider together with THIS org&#39;s connection status for it — the catalog the console&#39;s Integrations page renders.
[**GetIntegrationsByProvider**](IntegrationsAPI.md#GetIntegrationsByProvider) | **Get** /v1/integrations/{provider} | Returns ONE provider with this org&#39;s connection status — the same view list carries, for a single id.
[**GetIntegrationsByProviderCallback**](IntegrationsAPI.md#GetIntegrationsByProviderCallback) | **Get** /v1/integrations/{provider}/callback | OAuth return for any connector
[**GetIntegrationsConnectors**](IntegrationsAPI.md#GetIntegrationsConnectors) | **Get** /v1/integrations/connectors | Lists the caller&#39;s OWN connectors across every provider — the set &#x60;hanzo connector ls&#x60; prints.
[**GetIntegrationsConnectorsByIdToken**](IntegrationsAPI.md#GetIntegrationsConnectorsByIdToken) | **Get** /v1/integrations/connectors/{id}/token | Hands the custodied access token to its owner — the ONE place custody exits.
[**GetIntegrationsConnectorsProviders**](IntegrationsAPI.md#GetIntegrationsConnectorsProviders) | **Get** /v1/integrations/connectors/providers | Lists the user-scoped provider cards — the catalog of what a user can connect, and how.
[**GetIntegrationsDiscordLink**](IntegrationsAPI.md#GetIntegrationsDiscordLink) | **Get** /v1/integrations/discord/link | Begin linking a Hanzo account from Discord
[**GetIntegrationsDiscordLinkCallback**](IntegrationsAPI.md#GetIntegrationsDiscordLinkCallback) | **Get** /v1/integrations/discord/link/callback | Complete the Discord account link
[**GetIntegrationsDiscordLinkDiscord**](IntegrationsAPI.md#GetIntegrationsDiscordLinkDiscord) | **Get** /v1/integrations/discord/link/discord | Discord sign-in return leg
[**GetIntegrationsGithubInstallations**](IntegrationsAPI.md#GetIntegrationsGithubInstallations) | **Get** /v1/integrations/github/installations | Lists the GitHub accounts the caller may see the App installed on, each confirmed against the App&#39;s own list, plus where to add another.
[**GetIntegrationsGithubRepos**](IntegrationsAPI.md#GetIntegrationsGithubRepos) | **Get** /v1/integrations/github/repos | Lists the org&#39;s granted GitHub repositories, each annotated with its native import + sync status from the git object plane.
[**GetIntegrationsGithubReposByRepoPages**](IntegrationsAPI.md#GetIntegrationsGithubReposByRepoPages) | **Get** /v1/integrations/github/repos/{repo}/pages | Returns the repo&#39;s Pages status, live URL, custom domain and build source.
[**GetIntegrationsGitlabProjects**](IntegrationsAPI.md#GetIntegrationsGitlabProjects) | **Get** /v1/integrations/gitlab/projects | Lists the projects the org&#39;s GitLab connection can reach — membership projects, most recently active first.
[**GetIntegrationsSlackInstall**](IntegrationsAPI.md#GetIntegrationsSlackInstall) | **Get** /v1/integrations/slack/install | Install the Hanzo app into a Slack workspace
[**GetIntegrationsSlackLink**](IntegrationsAPI.md#GetIntegrationsSlackLink) | **Get** /v1/integrations/slack/link | Begin linking a Hanzo account from Slack
[**GetIntegrationsSlackLinkCallback**](IntegrationsAPI.md#GetIntegrationsSlackLinkCallback) | **Get** /v1/integrations/slack/link/callback | Complete the Slack account link
[**GetIntegrationsSlackLinkSlack**](IntegrationsAPI.md#GetIntegrationsSlackLinkSlack) | **Get** /v1/integrations/slack/link/slack | Slack sign-in return leg
[**GetIntegrationsTeamsLink**](IntegrationsAPI.md#GetIntegrationsTeamsLink) | **Get** /v1/integrations/teams/link | Begin linking a Hanzo account from Teams
[**GetIntegrationsTeamsLinkAad**](IntegrationsAPI.md#GetIntegrationsTeamsLinkAad) | **Get** /v1/integrations/teams/link/aad | Microsoft sign-in return leg
[**GetIntegrationsTeamsLinkCallback**](IntegrationsAPI.md#GetIntegrationsTeamsLinkCallback) | **Get** /v1/integrations/teams/link/callback | Complete the Teams account link
[**GetIntegrationsTelegramLink**](IntegrationsAPI.md#GetIntegrationsTelegramLink) | **Get** /v1/integrations/telegram/link | Begin linking a Hanzo account from Telegram
[**GetIntegrationsTelegramLinkAuth**](IntegrationsAPI.md#GetIntegrationsTelegramLinkAuth) | **Get** /v1/integrations/telegram/link/auth | Telegram Login Widget return leg
[**GetIntegrationsTelegramLinkCallback**](IntegrationsAPI.md#GetIntegrationsTelegramLinkCallback) | **Get** /v1/integrations/telegram/link/callback | Complete the Telegram account link
[**PostIntegrationsByProviderConnect**](IntegrationsAPI.md#PostIntegrationsByProviderConnect) | **Post** /v1/integrations/{provider}/connect | Acquires the org&#39;s credential for one provider.
[**PostIntegrationsByProviderDisconnect**](IntegrationsAPI.md#PostIntegrationsByProviderDisconnect) | **Post** /v1/integrations/{provider}/disconnect | Revokes (best-effort) and forgets an org&#39;s connection: it deletes every custodied KMS secret and the connection row.
[**PostIntegrationsByProviderVerify**](IntegrationsAPI.md#PostIntegrationsByProviderVerify) | **Post** /v1/integrations/{provider}/verify | Re-checks a CONNECTED apikey connector&#39;s stored credential against the provider, live (&#x60;hanzo connector verify&#x60;).
[**PostIntegrationsConnectorsByIdRefresh**](IntegrationsAPI.md#PostIntegrationsConnectorsByIdRefresh) | **Post** /v1/integrations/connectors/{id}/refresh | Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.
[**PostIntegrationsConnectorsByProviderCredential**](IntegrationsAPI.md#PostIntegrationsConnectorsByProviderCredential) | **Post** /v1/integrations/connectors/{provider}/credential | Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI&#39;s local PKCE (Adopt).
[**PostIntegrationsConnectorsByProviderDevice**](IntegrationsAPI.md#PostIntegrationsConnectorsByProviderDevice) | **Post** /v1/integrations/connectors/{provider}/device | Begins a device sign-in and returns the code to show the user plus how to poll for completion.
[**PostIntegrationsConnectorsByProviderDeviceByFlowPoll**](IntegrationsAPI.md#PostIntegrationsConnectorsByProviderDeviceByFlowPoll) | **Post** /v1/integrations/connectors/{provider}/device/{flow}/poll | Advances a device sign-in.
[**PostIntegrationsDiscordInteractions**](IntegrationsAPI.md#PostIntegrationsDiscordInteractions) | **Post** /v1/integrations/discord/interactions | Discord interactions endpoint
[**PostIntegrationsGithubClaim**](IntegrationsAPI.md#PostIntegrationsGithubClaim) | **Post** /v1/integrations/github/claim | Binds installations the App ALREADY holds to the org the caller is acting in — the reconciliation for a grant that happened outside our connect flow.
[**PostIntegrationsGithubFork**](IntegrationsAPI.md#PostIntegrationsGithubFork) | **Post** /v1/integrations/github/fork | Forks a granted repository.
[**PostIntegrationsGithubIssuesBackfill**](IntegrationsAPI.md#PostIntegrationsGithubIssuesBackfill) | **Post** /v1/integrations/github/issues/backfill | Seeds the native todo with the EXISTING issues across the org&#39;s granted repos (default state&#x3D;open); the webhook keeps them live thereafter.
[**PostIntegrationsGithubReposByRepoPages**](IntegrationsAPI.md#PostIntegrationsGithubReposByRepoPages) | **Post** /v1/integrations/github/repos/{repo}/pages | Creates the repo&#39;s Pages site and answers 201 Created with it.
[**PostIntegrationsGithubReposByRepoPagesBuilds**](IntegrationsAPI.md#PostIntegrationsGithubReposByRepoPagesBuilds) | **Post** /v1/integrations/github/repos/{repo}/pages/builds | Requests a Pages rebuild and returns the queued build&#39;s status.
[**PostIntegrationsGithubReposImport**](IntegrationsAPI.md#PostIntegrationsGithubReposImport) | **Post** /v1/integrations/github/repos/import | Imports the selected (or all) granted repos into git.hanzo.ai.
[**PostIntegrationsGithubSearch**](IntegrationsAPI.md#PostIntegrationsGithubSearch) | **Post** /v1/integrations/github/search | Finds repositories on GitHub.
[**PostIntegrationsGithubWebhook**](IntegrationsAPI.md#PostIntegrationsGithubWebhook) | **Post** /v1/integrations/github/webhook | GitHub App webhook
[**PostIntegrationsOpenrouterWebhook**](IntegrationsAPI.md#PostIntegrationsOpenrouterWebhook) | **Post** /v1/integrations/openrouter/webhook | Receive OpenRouter Broadcast traces as usage rows
[**PostIntegrationsSlackCommands**](IntegrationsAPI.md#PostIntegrationsSlackCommands) | **Post** /v1/integrations/slack/commands | Slack slash command webhook
[**PostIntegrationsSlackEvents**](IntegrationsAPI.md#PostIntegrationsSlackEvents) | **Post** /v1/integrations/slack/events | Slack Events API webhook
[**PostIntegrationsTeamsEvents**](IntegrationsAPI.md#PostIntegrationsTeamsEvents) | **Post** /v1/integrations/teams/events | Microsoft Teams Bot Framework webhook
[**PostIntegrationsTelegramConnect**](IntegrationsAPI.md#PostIntegrationsTelegramConnect) | **Post** /v1/integrations/telegram/connect | Mints a short, single-use deep-link code bound to the caller&#39;s org and returns the t.me link the console navigates to.
[**PostIntegrationsTelegramWebhook**](IntegrationsAPI.md#PostIntegrationsTelegramWebhook) | **Post** /v1/integrations/telegram/webhook | Telegram Bot API webhook
[**PutIntegrationsGithubReposByRepoPages**](IntegrationsAPI.md#PutIntegrationsGithubReposByRepoPages) | **Put** /v1/integrations/github/repos/{repo}/pages | Sets or clears the custom domain (cname) and updates HTTPS enforcement, build type, or source.



## DeleteIntegrationsConnectorsById

> DisconnectOut DeleteIntegrationsConnectorsById(ctx, id).Execute()

Forgets a connector: every custodied secret, then the row.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.DeleteIntegrationsConnectorsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.DeleteIntegrationsConnectorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntegrationsConnectorsById`: DisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.DeleteIntegrationsConnectorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationsConnectorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisconnectOut**](DisconnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationsGithubReposByRepoPages

> GithubPagesDisabledOut DeleteIntegrationsGithubReposByRepoPages(ctx, repo).Execute()

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
	resp, r, err := apiClient.IntegrationsAPI.DeleteIntegrationsGithubReposByRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.DeleteIntegrationsGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntegrationsGithubReposByRepoPages`: GithubPagesDisabledOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.DeleteIntegrationsGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationsGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GithubPagesDisabledOut**](GithubPagesDisabledOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrations

> ListOut GetIntegrations(ctx).Execute()

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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrations`: ListOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsRequest struct via the builder pattern


### Return type

[**ListOut**](ListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsByProvider

> ProviderView GetIntegrationsByProvider(ctx, provider).Execute()

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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/integrations/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsByProvider`: ProviderView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/integrations/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderView**](ProviderView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsByProviderCallback

> GetIntegrationsByProviderCallback(ctx, provider).Execute()

OAuth return for any connector



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsByProviderCallback(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsByProviderCallback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetIntegrationsByProviderCallbackRequest struct via the builder pattern


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


## GetIntegrationsConnectors

> ConnectorsOut GetIntegrationsConnectors(ctx).Execute()

Lists the caller's OWN connectors across every provider — the set `hanzo connector ls` prints.



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsConnectors`: ConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsConnectorsRequest struct via the builder pattern


### Return type

[**ConnectorsOut**](ConnectorsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsConnectorsByIdToken

> ConnectorTokenOut GetIntegrationsConnectorsByIdToken(ctx, id).Execute()

Hands the custodied access token to its owner — the ONE place custody exits.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsConnectorsByIdToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsConnectorsByIdToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsConnectorsByIdToken`: ConnectorTokenOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsConnectorsByIdToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsConnectorsByIdTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorTokenOut**](ConnectorTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsConnectorsProviders

> ConnectorProvidersOut GetIntegrationsConnectorsProviders(ctx).Execute()

Lists the user-scoped provider cards — the catalog of what a user can connect, and how.



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsConnectorsProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsConnectorsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsConnectorsProviders`: ConnectorProvidersOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsConnectorsProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsConnectorsProvidersRequest struct via the builder pattern


### Return type

[**ConnectorProvidersOut**](ConnectorProvidersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsDiscordLink

> GetIntegrationsDiscordLink(ctx).Execute()

Begin linking a Hanzo account from Discord



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsDiscordLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsDiscordLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsDiscordLinkRequest struct via the builder pattern


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


## GetIntegrationsDiscordLinkCallback

> GetIntegrationsDiscordLinkCallback(ctx).Execute()

Complete the Discord account link



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsDiscordLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsDiscordLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsDiscordLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationsDiscordLinkDiscord

> GetIntegrationsDiscordLinkDiscord(ctx).Execute()

Discord sign-in return leg



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsDiscordLinkDiscord(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsDiscordLinkDiscord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsDiscordLinkDiscordRequest struct via the builder pattern


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


## GetIntegrationsGithubInstallations

> GithubInstallationsOut GetIntegrationsGithubInstallations(ctx).Execute()

Lists the GitHub accounts the caller may see the App installed on, each confirmed against the App's own list, plus where to add another.



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsGithubInstallations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsGithubInstallations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsGithubInstallations`: GithubInstallationsOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsGithubInstallations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsGithubInstallationsRequest struct via the builder pattern


### Return type

[**GithubInstallationsOut**](GithubInstallationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsGithubRepos

> GithubReposOut GetIntegrationsGithubRepos(ctx).Execute()

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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsGithubRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsGithubRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsGithubRepos`: GithubReposOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsGithubRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsGithubReposRequest struct via the builder pattern


### Return type

[**GithubReposOut**](GithubReposOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsGithubReposByRepoPages

> GithubPagesView GetIntegrationsGithubReposByRepoPages(ctx, repo).Execute()

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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsGithubReposByRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsGithubReposByRepoPages`: GithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GithubPagesView**](GithubPagesView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsGitlabProjects

> GitlabProjectsOut GetIntegrationsGitlabProjects(ctx).Execute()

Lists the projects the org's GitLab connection can reach — membership projects, most recently active first.



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
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationsGitlabProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsGitlabProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsGitlabProjects`: GitlabProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationsGitlabProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsGitlabProjectsRequest struct via the builder pattern


### Return type

[**GitlabProjectsOut**](GitlabProjectsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsSlackInstall

> GetIntegrationsSlackInstall(ctx).Execute()

Install the Hanzo app into a Slack workspace



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsSlackInstall(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsSlackInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsSlackInstallRequest struct via the builder pattern


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


## GetIntegrationsSlackLink

> GetIntegrationsSlackLink(ctx).Execute()

Begin linking a Hanzo account from Slack



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsSlackLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsSlackLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsSlackLinkRequest struct via the builder pattern


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


## GetIntegrationsSlackLinkCallback

> GetIntegrationsSlackLinkCallback(ctx).Execute()

Complete the Slack account link



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsSlackLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsSlackLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsSlackLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationsSlackLinkSlack

> GetIntegrationsSlackLinkSlack(ctx).Execute()

Slack sign-in return leg



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsSlackLinkSlack(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsSlackLinkSlack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsSlackLinkSlackRequest struct via the builder pattern


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


## GetIntegrationsTeamsLink

> GetIntegrationsTeamsLink(ctx).Execute()

Begin linking a Hanzo account from Teams



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsTeamsLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsTeamsLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsTeamsLinkRequest struct via the builder pattern


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


## GetIntegrationsTeamsLinkAad

> GetIntegrationsTeamsLinkAad(ctx).Execute()

Microsoft sign-in return leg



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsTeamsLinkAad(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsTeamsLinkAad``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsTeamsLinkAadRequest struct via the builder pattern


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


## GetIntegrationsTeamsLinkCallback

> GetIntegrationsTeamsLinkCallback(ctx).Execute()

Complete the Teams account link



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsTeamsLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsTeamsLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsTeamsLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationsTelegramLink

> GetIntegrationsTelegramLink(ctx).Execute()

Begin linking a Hanzo account from Telegram



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsTelegramLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsTelegramLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsTelegramLinkRequest struct via the builder pattern


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


## GetIntegrationsTelegramLinkAuth

> GetIntegrationsTelegramLinkAuth(ctx).Execute()

Telegram Login Widget return leg



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsTelegramLinkAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsTelegramLinkAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsTelegramLinkAuthRequest struct via the builder pattern


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


## GetIntegrationsTelegramLinkCallback

> GetIntegrationsTelegramLinkCallback(ctx).Execute()

Complete the Telegram account link



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
	r, err := apiClient.IntegrationsAPI.GetIntegrationsTelegramLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationsTelegramLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsTelegramLinkCallbackRequest struct via the builder pattern


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


## PostIntegrationsByProviderConnect

> ConnectOut PostIntegrationsByProviderConnect(ctx, provider).ConnectIn(connectIn).Execute()

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
	connectIn := *openapiclient.NewConnectIn() // ConnectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsByProviderConnect(context.Background(), provider).ConnectIn(connectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsByProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsByProviderConnect`: ConnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsByProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsByProviderConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectIn** | [**ConnectIn**](ConnectIn.md) |  | 

### Return type

[**ConnectOut**](ConnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsByProviderDisconnect

> DisconnectOut PostIntegrationsByProviderDisconnect(ctx, provider).Execute()

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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/integrations/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsByProviderDisconnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsByProviderDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsByProviderDisconnect`: DisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsByProviderDisconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/integrations/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsByProviderDisconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisconnectOut**](DisconnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsByProviderVerify

> VerifyOut PostIntegrationsByProviderVerify(ctx, provider).Execute()

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
	provider := "cloudflare" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/integrations/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsByProviderVerify(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsByProviderVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsByProviderVerify`: VerifyOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsByProviderVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/integrations/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsByProviderVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerifyOut**](VerifyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsConnectorsByIdRefresh

> RefreshOut PostIntegrationsConnectorsByIdRefresh(ctx, id).Execute()

Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsConnectorsByIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsConnectorsByIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsConnectorsByIdRefresh`: RefreshOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsConnectorsByIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsConnectorsByIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshOut**](RefreshOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsConnectorsByProviderCredential

> CredentialOut PostIntegrationsConnectorsByProviderCredential(ctx, provider).CredentialIn(credentialIn).Execute()

Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI's local PKCE (Adopt).



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	credentialIn := *openapiclient.NewCredentialIn() // CredentialIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsConnectorsByProviderCredential(context.Background(), provider).CredentialIn(credentialIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsConnectorsByProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsConnectorsByProviderCredential`: CredentialOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsConnectorsByProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsConnectorsByProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialIn** | [**CredentialIn**](CredentialIn.md) |  | 

### Return type

[**CredentialOut**](CredentialOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsConnectorsByProviderDevice

> DeviceStartOut PostIntegrationsConnectorsByProviderDevice(ctx, provider).DeviceStartIn(deviceStartIn).Execute()

Begins a device sign-in and returns the code to show the user plus how to poll for completion.



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	deviceStartIn := *openapiclient.NewDeviceStartIn() // DeviceStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsConnectorsByProviderDevice(context.Background(), provider).DeviceStartIn(deviceStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsConnectorsByProviderDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsConnectorsByProviderDevice`: DeviceStartOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsConnectorsByProviderDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsConnectorsByProviderDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceStartIn** | [**DeviceStartIn**](DeviceStartIn.md) |  | 

### Return type

[**DeviceStartOut**](DeviceStartOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsConnectorsByProviderDeviceByFlowPoll

> DevicePollOut PostIntegrationsConnectorsByProviderDeviceByFlowPoll(ctx, provider, flow).Execute()

Advances a device sign-in.



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	flow := "g_7f2c" // string | Flow is the id deviceStartOut returned. Expired or another user's flow is indistinguishable from an unknown one: 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsConnectorsByProviderDeviceByFlowPoll(context.Background(), provider, flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsConnectorsByProviderDeviceByFlowPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsConnectorsByProviderDeviceByFlowPoll`: DevicePollOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsConnectorsByProviderDeviceByFlowPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 
**flow** | **string** | Flow is the id deviceStartOut returned. Expired or another user&#39;s flow is indistinguishable from an unknown one: 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsConnectorsByProviderDeviceByFlowPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicePollOut**](DevicePollOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsDiscordInteractions

> PostIntegrationsDiscordInteractions(ctx).Execute()

Discord interactions endpoint



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
	r, err := apiClient.IntegrationsAPI.PostIntegrationsDiscordInteractions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsDiscordInteractions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsDiscordInteractionsRequest struct via the builder pattern


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


## PostIntegrationsGithubClaim

> GithubClaimOut PostIntegrationsGithubClaim(ctx).GithubClaimIn(githubClaimIn).Execute()

Binds installations the App ALREADY holds to the org the caller is acting in — the reconciliation for a grant that happened outside our connect flow.



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
	githubClaimIn := *openapiclient.NewGithubClaimIn() // GithubClaimIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubClaim(context.Background()).GithubClaimIn(githubClaimIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubClaim`: GithubClaimOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubClaimIn** | [**GithubClaimIn**](GithubClaimIn.md) |  | 

### Return type

[**GithubClaimOut**](GithubClaimOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubFork

> GithubForkOut PostIntegrationsGithubFork(ctx).GithubForkReq(githubForkReq).Execute()

Forks a granted repository.



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
	githubForkReq := *openapiclient.NewGithubForkReq() // GithubForkReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubFork(context.Background()).GithubForkReq(githubForkReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubFork`: GithubForkOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubForkReq** | [**GithubForkReq**](GithubForkReq.md) |  | 

### Return type

[**GithubForkOut**](GithubForkOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubIssuesBackfill

> GithubBackfillResult PostIntegrationsGithubIssuesBackfill(ctx).GithubBackfillIn(githubBackfillIn).Execute()

Seeds the native todo with the EXISTING issues across the org's granted repos (default state=open); the webhook keeps them live thereafter.



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
	githubBackfillIn := *openapiclient.NewGithubBackfillIn() // GithubBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubIssuesBackfill(context.Background()).GithubBackfillIn(githubBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubIssuesBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubIssuesBackfill`: GithubBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubIssuesBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubIssuesBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubBackfillIn** | [**GithubBackfillIn**](GithubBackfillIn.md) |  | 

### Return type

[**GithubBackfillResult**](GithubBackfillResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubReposByRepoPages

> GithubPagesView PostIntegrationsGithubReposByRepoPages(ctx, repo).GithubPagesEnableReq(githubPagesEnableReq).Execute()

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
	githubPagesEnableReq := *openapiclient.NewGithubPagesEnableReq() // GithubPagesEnableReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubReposByRepoPages(context.Background(), repo).GithubPagesEnableReq(githubPagesEnableReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubReposByRepoPages`: GithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubPagesEnableReq** | [**GithubPagesEnableReq**](GithubPagesEnableReq.md) |  | 

### Return type

[**GithubPagesView**](GithubPagesView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubReposByRepoPagesBuilds

> GithubPagesBuildOut PostIntegrationsGithubReposByRepoPagesBuilds(ctx, repo).Execute()

Requests a Pages rebuild and returns the queued build's status.



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
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubReposByRepoPagesBuilds(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubReposByRepoPagesBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubReposByRepoPagesBuilds`: GithubPagesBuildOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubReposByRepoPagesBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubReposByRepoPagesBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GithubPagesBuildOut**](GithubPagesBuildOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubReposImport

> GithubImportOut PostIntegrationsGithubReposImport(ctx).GithubImportIn(githubImportIn).Execute()

Imports the selected (or all) granted repos into git.hanzo.ai.



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
	githubImportIn := *openapiclient.NewGithubImportIn() // GithubImportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubReposImport(context.Background()).GithubImportIn(githubImportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubReposImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubReposImport`: GithubImportOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubReposImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubReposImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubImportIn** | [**GithubImportIn**](GithubImportIn.md) |  | 

### Return type

[**GithubImportOut**](GithubImportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubSearch

> GithubSearchOut PostIntegrationsGithubSearch(ctx).GithubSearchReq(githubSearchReq).Execute()

Finds repositories on GitHub.



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
	githubSearchReq := *openapiclient.NewGithubSearchReq() // GithubSearchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubSearch(context.Background()).GithubSearchReq(githubSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsGithubSearch`: GithubSearchOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsGithubSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubSearchReq** | [**GithubSearchReq**](GithubSearchReq.md) |  | 

### Return type

[**GithubSearchOut**](GithubSearchOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsGithubWebhook

> PostIntegrationsGithubWebhook(ctx).Execute()

GitHub App webhook



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
	r, err := apiClient.IntegrationsAPI.PostIntegrationsGithubWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsGithubWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsGithubWebhookRequest struct via the builder pattern


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


## PostIntegrationsOpenrouterWebhook

> map[string]interface{} PostIntegrationsOpenrouterWebhook(ctx).RequestBody(requestBody).Execute()

Receive OpenRouter Broadcast traces as usage rows



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsOpenrouterWebhook(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsOpenrouterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsOpenrouterWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsOpenrouterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsOpenrouterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsSlackCommands

> PostIntegrationsSlackCommands(ctx).Execute()

Slack slash command webhook



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
	r, err := apiClient.IntegrationsAPI.PostIntegrationsSlackCommands(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsSlackCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsSlackCommandsRequest struct via the builder pattern


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


## PostIntegrationsSlackEvents

> PostIntegrationsSlackEvents(ctx).Execute()

Slack Events API webhook



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
	r, err := apiClient.IntegrationsAPI.PostIntegrationsSlackEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsSlackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsSlackEventsRequest struct via the builder pattern


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


## PostIntegrationsTeamsEvents

> PostIntegrationsTeamsEvents(ctx).Execute()

Microsoft Teams Bot Framework webhook



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
	r, err := apiClient.IntegrationsAPI.PostIntegrationsTeamsEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsTeamsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsTeamsEventsRequest struct via the builder pattern


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


## PostIntegrationsTelegramConnect

> AuthorizeOut PostIntegrationsTelegramConnect(ctx).Execute()

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
	resp, r, err := apiClient.IntegrationsAPI.PostIntegrationsTelegramConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsTelegramConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationsTelegramConnect`: AuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PostIntegrationsTelegramConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsTelegramConnectRequest struct via the builder pattern


### Return type

[**AuthorizeOut**](AuthorizeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationsTelegramWebhook

> PostIntegrationsTelegramWebhook(ctx).Execute()

Telegram Bot API webhook



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
	r, err := apiClient.IntegrationsAPI.PostIntegrationsTelegramWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PostIntegrationsTelegramWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationsTelegramWebhookRequest struct via the builder pattern


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


## PutIntegrationsGithubReposByRepoPages

> GithubPagesUpdatedOut PutIntegrationsGithubReposByRepoPages(ctx, repo).GithubPagesUpdateReq(githubPagesUpdateReq).Execute()

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
	githubPagesUpdateReq := *openapiclient.NewGithubPagesUpdateReq() // GithubPagesUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PutIntegrationsGithubReposByRepoPages(context.Background(), repo).GithubPagesUpdateReq(githubPagesUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PutIntegrationsGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIntegrationsGithubReposByRepoPages`: GithubPagesUpdatedOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PutIntegrationsGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIntegrationsGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **githubPagesUpdateReq** | [**GithubPagesUpdateReq**](GithubPagesUpdateReq.md) |  | 

### Return type

[**GithubPagesUpdatedOut**](GithubPagesUpdatedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

