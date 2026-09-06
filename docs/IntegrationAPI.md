# \IntegrationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIntegrationConnectorsById**](IntegrationAPI.md#DeleteIntegrationConnectorsById) | **Delete** /v1/integration/connectors/{id} | Forgets a connector: every custodied secret, then the row.
[**DeleteIntegrationGithubReposByRepoPages**](IntegrationAPI.md#DeleteIntegrationGithubReposByRepoPages) | **Delete** /v1/integration/github/repos/{repo}/pages | Deletes the repo&#39;s Pages site.
[**GetIntegration**](IntegrationAPI.md#GetIntegration) | **Get** /v1/integration | Returns every registered integration provider together with THIS org&#39;s connection status for it — the catalog the console&#39;s Integrations page renders.
[**GetIntegrationByProvider**](IntegrationAPI.md#GetIntegrationByProvider) | **Get** /v1/integration/{provider} | Returns ONE provider with this org&#39;s connection status — the same view list carries, for a single id.
[**GetIntegrationByProviderCallback**](IntegrationAPI.md#GetIntegrationByProviderCallback) | **Get** /v1/integration/{provider}/callback | OAuth return for any connector
[**GetIntegrationConnectors**](IntegrationAPI.md#GetIntegrationConnectors) | **Get** /v1/integration/connectors | Lists the caller&#39;s OWN connectors across every provider — the set &#x60;hanzo connector ls&#x60; prints.
[**GetIntegrationConnectorsByIdToken**](IntegrationAPI.md#GetIntegrationConnectorsByIdToken) | **Get** /v1/integration/connectors/{id}/token | Hands the custodied access token to its owner — the ONE place custody exits.
[**GetIntegrationConnectorsProviders**](IntegrationAPI.md#GetIntegrationConnectorsProviders) | **Get** /v1/integration/connectors/providers | Lists the user-scoped provider cards — the catalog of what a user can connect, and how.
[**GetIntegrationDiscordLink**](IntegrationAPI.md#GetIntegrationDiscordLink) | **Get** /v1/integration/discord/link | Begin linking a Hanzo account from Discord
[**GetIntegrationDiscordLinkCallback**](IntegrationAPI.md#GetIntegrationDiscordLinkCallback) | **Get** /v1/integration/discord/link/callback | Complete the Discord account link
[**GetIntegrationDiscordLinkDiscord**](IntegrationAPI.md#GetIntegrationDiscordLinkDiscord) | **Get** /v1/integration/discord/link/discord | Discord sign-in return leg
[**GetIntegrationGithubInstallations**](IntegrationAPI.md#GetIntegrationGithubInstallations) | **Get** /v1/integration/github/installations | Lists the GitHub accounts the caller may see the App installed on, each confirmed against the App&#39;s own list, plus where to add another.
[**GetIntegrationGithubRepos**](IntegrationAPI.md#GetIntegrationGithubRepos) | **Get** /v1/integration/github/repos | Lists the org&#39;s granted GitHub repositories, each annotated with its native import + sync status from the git object plane.
[**GetIntegrationGithubReposByRepoPages**](IntegrationAPI.md#GetIntegrationGithubReposByRepoPages) | **Get** /v1/integration/github/repos/{repo}/pages | Returns the repo&#39;s Pages status, live URL, custom domain and build source.
[**GetIntegrationGitlabProjects**](IntegrationAPI.md#GetIntegrationGitlabProjects) | **Get** /v1/integration/gitlab/projects | Lists the projects the org&#39;s GitLab connection can reach — membership projects, most recently active first.
[**GetIntegrationSlackInstall**](IntegrationAPI.md#GetIntegrationSlackInstall) | **Get** /v1/integration/slack/install | Install the Hanzo app into a Slack workspace
[**GetIntegrationSlackLink**](IntegrationAPI.md#GetIntegrationSlackLink) | **Get** /v1/integration/slack/link | Begin linking a Hanzo account from Slack
[**GetIntegrationSlackLinkCallback**](IntegrationAPI.md#GetIntegrationSlackLinkCallback) | **Get** /v1/integration/slack/link/callback | Complete the Slack account link
[**GetIntegrationSlackLinkSlack**](IntegrationAPI.md#GetIntegrationSlackLinkSlack) | **Get** /v1/integration/slack/link/slack | Slack sign-in return leg
[**GetIntegrationTeamsLink**](IntegrationAPI.md#GetIntegrationTeamsLink) | **Get** /v1/integration/teams/link | Begin linking a Hanzo account from Teams
[**GetIntegrationTeamsLinkAad**](IntegrationAPI.md#GetIntegrationTeamsLinkAad) | **Get** /v1/integration/teams/link/aad | Microsoft sign-in return leg
[**GetIntegrationTeamsLinkCallback**](IntegrationAPI.md#GetIntegrationTeamsLinkCallback) | **Get** /v1/integration/teams/link/callback | Complete the Teams account link
[**GetIntegrationTelegramLink**](IntegrationAPI.md#GetIntegrationTelegramLink) | **Get** /v1/integration/telegram/link | Begin linking a Hanzo account from Telegram
[**GetIntegrationTelegramLinkAuth**](IntegrationAPI.md#GetIntegrationTelegramLinkAuth) | **Get** /v1/integration/telegram/link/auth | Telegram Login Widget return leg
[**GetIntegrationTelegramLinkCallback**](IntegrationAPI.md#GetIntegrationTelegramLinkCallback) | **Get** /v1/integration/telegram/link/callback | Complete the Telegram account link
[**GetIntegrationWhatsappWebhook**](IntegrationAPI.md#GetIntegrationWhatsappWebhook) | **Get** /v1/integration/whatsapp/webhook | WhatsApp Cloud API subscription challenge
[**PostIntegrationByProviderConnect**](IntegrationAPI.md#PostIntegrationByProviderConnect) | **Post** /v1/integration/{provider}/connect | Acquires the org&#39;s credential for one provider.
[**PostIntegrationByProviderDisconnect**](IntegrationAPI.md#PostIntegrationByProviderDisconnect) | **Post** /v1/integration/{provider}/disconnect | Revokes (best-effort) and forgets an org&#39;s connection: it deletes every custodied KMS secret and the connection row.
[**PostIntegrationByProviderVerify**](IntegrationAPI.md#PostIntegrationByProviderVerify) | **Post** /v1/integration/{provider}/verify | Re-checks a CONNECTED apikey connector&#39;s stored credential against the provider, live (&#x60;hanzo connector verify&#x60;).
[**PostIntegrationConnectorsByIdRefresh**](IntegrationAPI.md#PostIntegrationConnectorsByIdRefresh) | **Post** /v1/integration/connectors/{id}/refresh | Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.
[**PostIntegrationConnectorsByProviderCredential**](IntegrationAPI.md#PostIntegrationConnectorsByProviderCredential) | **Post** /v1/integration/connectors/{provider}/credential | Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI&#39;s local PKCE (Adopt).
[**PostIntegrationConnectorsByProviderDevice**](IntegrationAPI.md#PostIntegrationConnectorsByProviderDevice) | **Post** /v1/integration/connectors/{provider}/device | Begins a device sign-in and returns the code to show the user plus how to poll for completion.
[**PostIntegrationConnectorsByProviderDeviceByFlowPoll**](IntegrationAPI.md#PostIntegrationConnectorsByProviderDeviceByFlowPoll) | **Post** /v1/integration/connectors/{provider}/device/{flow}/poll | Advances a device sign-in.
[**PostIntegrationDiscordInteractions**](IntegrationAPI.md#PostIntegrationDiscordInteractions) | **Post** /v1/integration/discord/interactions | Discord interactions endpoint
[**PostIntegrationForgeWebhook**](IntegrationAPI.md#PostIntegrationForgeWebhook) | **Post** /v1/integration/forge/webhook | Forge workflow_job webhook
[**PostIntegrationGithubClaim**](IntegrationAPI.md#PostIntegrationGithubClaim) | **Post** /v1/integration/github/claim | Binds installations the App ALREADY holds to the org the caller is acting in — the reconciliation for a grant that happened outside our connect flow.
[**PostIntegrationGithubFork**](IntegrationAPI.md#PostIntegrationGithubFork) | **Post** /v1/integration/github/fork | Forks a granted repository.
[**PostIntegrationGithubIssuesBackfill**](IntegrationAPI.md#PostIntegrationGithubIssuesBackfill) | **Post** /v1/integration/github/issues/backfill | Seeds the native todo with the EXISTING issues across the org&#39;s granted repos (default state&#x3D;open); the webhook keeps them live thereafter.
[**PostIntegrationGithubReposByRepoPages**](IntegrationAPI.md#PostIntegrationGithubReposByRepoPages) | **Post** /v1/integration/github/repos/{repo}/pages | Creates the repo&#39;s Pages site and answers 201 Created with it.
[**PostIntegrationGithubReposByRepoPagesBuilds**](IntegrationAPI.md#PostIntegrationGithubReposByRepoPagesBuilds) | **Post** /v1/integration/github/repos/{repo}/pages/builds | Requests a Pages rebuild and returns the queued build&#39;s status.
[**PostIntegrationGithubReposImport**](IntegrationAPI.md#PostIntegrationGithubReposImport) | **Post** /v1/integration/github/repos/import | Imports the selected (or all) granted repos into git.hanzo.ai.
[**PostIntegrationGithubSearch**](IntegrationAPI.md#PostIntegrationGithubSearch) | **Post** /v1/integration/github/search | Finds repositories on GitHub.
[**PostIntegrationGithubWebhook**](IntegrationAPI.md#PostIntegrationGithubWebhook) | **Post** /v1/integration/github/webhook | GitHub App webhook
[**PostIntegrationLinearClaim**](IntegrationAPI.md#PostIntegrationLinearClaim) | **Post** /v1/integration/linear/claim | Binds the caller&#39;s Linear organization to the org and seals the webhook secret.
[**PostIntegrationLinearComments**](IntegrationAPI.md#PostIntegrationLinearComments) | **Post** /v1/integration/linear/comments | Posts a comment on a Linear issue with the caller&#39;s own key, so it carries their name.
[**PostIntegrationLinearIssuesBackfill**](IntegrationAPI.md#PostIntegrationLinearIssuesBackfill) | **Post** /v1/integration/linear/issues/backfill | Seeds the native todo with the EXISTING Linear issues the caller&#39;s key can see (default state&#x3D;open); the webhook keeps them live thereafter.
[**PostIntegrationLinearWebhook**](IntegrationAPI.md#PostIntegrationLinearWebhook) | **Post** /v1/integration/linear/webhook | Linear webhook
[**PostIntegrationOpenrouterWebhook**](IntegrationAPI.md#PostIntegrationOpenrouterWebhook) | **Post** /v1/integration/openrouter/webhook | Receive OpenRouter Broadcast traces as usage rows
[**PostIntegrationSlackCommands**](IntegrationAPI.md#PostIntegrationSlackCommands) | **Post** /v1/integration/slack/commands | Slack slash command webhook
[**PostIntegrationSlackEvents**](IntegrationAPI.md#PostIntegrationSlackEvents) | **Post** /v1/integration/slack/events | Slack Events API webhook
[**PostIntegrationSlackJoin**](IntegrationAPI.md#PostIntegrationSlackJoin) | **Post** /v1/integration/slack/join | Joins every public channel in the caller org&#39;s workspace.
[**PostIntegrationTeamsEvents**](IntegrationAPI.md#PostIntegrationTeamsEvents) | **Post** /v1/integration/teams/events | Microsoft Teams Bot Framework webhook
[**PostIntegrationTelegramConnect**](IntegrationAPI.md#PostIntegrationTelegramConnect) | **Post** /v1/integration/telegram/connect | Mints a short, single-use deep-link code bound to the caller&#39;s org and returns the t.me link the console navigates to.
[**PostIntegrationTelegramWebhook**](IntegrationAPI.md#PostIntegrationTelegramWebhook) | **Post** /v1/integration/telegram/webhook | Telegram Bot API webhook
[**PostIntegrationWhatsappWebhook**](IntegrationAPI.md#PostIntegrationWhatsappWebhook) | **Post** /v1/integration/whatsapp/webhook | WhatsApp Cloud API webhook
[**PutIntegrationGithubReposByRepoPages**](IntegrationAPI.md#PutIntegrationGithubReposByRepoPages) | **Put** /v1/integration/github/repos/{repo}/pages | Sets or clears the custom domain (cname) and updates HTTPS enforcement, build type, or source.



## DeleteIntegrationConnectorsById

> DisconnectOut DeleteIntegrationConnectorsById(ctx, id).Execute()

Forgets a connector: every custodied secret, then the row.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.DeleteIntegrationConnectorsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteIntegrationConnectorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntegrationConnectorsById`: DisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.DeleteIntegrationConnectorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationConnectorsByIdRequest struct via the builder pattern


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


## DeleteIntegrationGithubReposByRepoPages

> GithubPagesDisabledOut DeleteIntegrationGithubReposByRepoPages(ctx, repo).Execute()

Deletes the repo's Pages site.



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
	repo := "widgets" // string | Repo is the repository's short name within the org's installation, with no owner prefix (the owner is server-derived from the grant). A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.DeleteIntegrationGithubReposByRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteIntegrationGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntegrationGithubReposByRepoPages`: GithubPagesDisabledOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.DeleteIntegrationGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationGithubReposByRepoPagesRequest struct via the builder pattern


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


## GetIntegration

> ListOut GetIntegration(ctx).Execute()

Returns every registered integration provider together with THIS org's connection status for it — the catalog the console's Integrations page renders.



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
	resp, r, err := apiClient.IntegrationAPI.GetIntegration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegration`: ListOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationRequest struct via the builder pattern


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


## GetIntegrationByProvider

> ProviderView GetIntegrationByProvider(ctx, provider).Execute()

Returns ONE provider with this org's connection status — the same view list carries, for a single id.



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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/integration/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationByProvider`: ProviderView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/integration/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationByProviderRequest struct via the builder pattern


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


## GetIntegrationByProviderCallback

> GetIntegrationByProviderCallback(ctx, provider).Execute()

OAuth return for any connector



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
	r, err := apiClient.IntegrationAPI.GetIntegrationByProviderCallback(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationByProviderCallback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetIntegrationByProviderCallbackRequest struct via the builder pattern


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


## GetIntegrationConnectors

> ConnectorsOut GetIntegrationConnectors(ctx).Execute()

Lists the caller's OWN connectors across every provider — the set `hanzo connector ls` prints.



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
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationConnectors`: ConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationConnectorsRequest struct via the builder pattern


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


## GetIntegrationConnectorsByIdToken

> ConnectorTokenOut GetIntegrationConnectorsByIdToken(ctx, id).Execute()

Hands the custodied access token to its owner — the ONE place custody exits.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationConnectorsByIdToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationConnectorsByIdToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationConnectorsByIdToken`: ConnectorTokenOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationConnectorsByIdToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationConnectorsByIdTokenRequest struct via the builder pattern


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


## GetIntegrationConnectorsProviders

> ConnectorProvidersOut GetIntegrationConnectorsProviders(ctx).Execute()

Lists the user-scoped provider cards — the catalog of what a user can connect, and how.



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
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationConnectorsProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationConnectorsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationConnectorsProviders`: ConnectorProvidersOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationConnectorsProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationConnectorsProvidersRequest struct via the builder pattern


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


## GetIntegrationDiscordLink

> GetIntegrationDiscordLink(ctx).Execute()

Begin linking a Hanzo account from Discord



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
	r, err := apiClient.IntegrationAPI.GetIntegrationDiscordLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationDiscordLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDiscordLinkRequest struct via the builder pattern


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


## GetIntegrationDiscordLinkCallback

> GetIntegrationDiscordLinkCallback(ctx).Execute()

Complete the Discord account link



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
	r, err := apiClient.IntegrationAPI.GetIntegrationDiscordLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationDiscordLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDiscordLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationDiscordLinkDiscord

> GetIntegrationDiscordLinkDiscord(ctx).Execute()

Discord sign-in return leg



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
	r, err := apiClient.IntegrationAPI.GetIntegrationDiscordLinkDiscord(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationDiscordLinkDiscord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDiscordLinkDiscordRequest struct via the builder pattern


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


## GetIntegrationGithubInstallations

> GithubInstallationsOut GetIntegrationGithubInstallations(ctx).Execute()

Lists the GitHub accounts the caller may see the App installed on, each confirmed against the App's own list, plus where to add another.



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
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationGithubInstallations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationGithubInstallations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationGithubInstallations`: GithubInstallationsOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationGithubInstallations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationGithubInstallationsRequest struct via the builder pattern


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


## GetIntegrationGithubRepos

> GithubReposOut GetIntegrationGithubRepos(ctx).Execute()

Lists the org's granted GitHub repositories, each annotated with its native import + sync status from the git object plane.



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
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationGithubRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationGithubRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationGithubRepos`: GithubReposOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationGithubRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationGithubReposRequest struct via the builder pattern


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


## GetIntegrationGithubReposByRepoPages

> GithubPagesView GetIntegrationGithubReposByRepoPages(ctx, repo).Execute()

Returns the repo's Pages status, live URL, custom domain and build source.



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
	repo := "widgets" // string | Repo is the repository's short name within the org's installation, with no owner prefix (the owner is server-derived from the grant). A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationGithubReposByRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationGithubReposByRepoPages`: GithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationGithubReposByRepoPagesRequest struct via the builder pattern


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


## GetIntegrationGitlabProjects

> GitlabProjectsOut GetIntegrationGitlabProjects(ctx).Execute()

Lists the projects the org's GitLab connection can reach — membership projects, most recently active first.



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
	resp, r, err := apiClient.IntegrationAPI.GetIntegrationGitlabProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationGitlabProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationGitlabProjects`: GitlabProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetIntegrationGitlabProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationGitlabProjectsRequest struct via the builder pattern


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


## GetIntegrationSlackInstall

> GetIntegrationSlackInstall(ctx).Execute()

Install the Hanzo app into a Slack workspace



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
	r, err := apiClient.IntegrationAPI.GetIntegrationSlackInstall(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationSlackInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationSlackInstallRequest struct via the builder pattern


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


## GetIntegrationSlackLink

> GetIntegrationSlackLink(ctx).Execute()

Begin linking a Hanzo account from Slack



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
	r, err := apiClient.IntegrationAPI.GetIntegrationSlackLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationSlackLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationSlackLinkRequest struct via the builder pattern


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


## GetIntegrationSlackLinkCallback

> GetIntegrationSlackLinkCallback(ctx).Execute()

Complete the Slack account link



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
	r, err := apiClient.IntegrationAPI.GetIntegrationSlackLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationSlackLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationSlackLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationSlackLinkSlack

> GetIntegrationSlackLinkSlack(ctx).Execute()

Slack sign-in return leg



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
	r, err := apiClient.IntegrationAPI.GetIntegrationSlackLinkSlack(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationSlackLinkSlack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationSlackLinkSlackRequest struct via the builder pattern


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


## GetIntegrationTeamsLink

> GetIntegrationTeamsLink(ctx).Execute()

Begin linking a Hanzo account from Teams



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
	r, err := apiClient.IntegrationAPI.GetIntegrationTeamsLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationTeamsLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTeamsLinkRequest struct via the builder pattern


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


## GetIntegrationTeamsLinkAad

> GetIntegrationTeamsLinkAad(ctx).Execute()

Microsoft sign-in return leg



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
	r, err := apiClient.IntegrationAPI.GetIntegrationTeamsLinkAad(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationTeamsLinkAad``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTeamsLinkAadRequest struct via the builder pattern


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


## GetIntegrationTeamsLinkCallback

> GetIntegrationTeamsLinkCallback(ctx).Execute()

Complete the Teams account link



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
	r, err := apiClient.IntegrationAPI.GetIntegrationTeamsLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationTeamsLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTeamsLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationTelegramLink

> GetIntegrationTelegramLink(ctx).Execute()

Begin linking a Hanzo account from Telegram



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
	r, err := apiClient.IntegrationAPI.GetIntegrationTelegramLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationTelegramLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTelegramLinkRequest struct via the builder pattern


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


## GetIntegrationTelegramLinkAuth

> GetIntegrationTelegramLinkAuth(ctx).Execute()

Telegram Login Widget return leg



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
	r, err := apiClient.IntegrationAPI.GetIntegrationTelegramLinkAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationTelegramLinkAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTelegramLinkAuthRequest struct via the builder pattern


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


## GetIntegrationTelegramLinkCallback

> GetIntegrationTelegramLinkCallback(ctx).Execute()

Complete the Telegram account link



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
	r, err := apiClient.IntegrationAPI.GetIntegrationTelegramLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationTelegramLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTelegramLinkCallbackRequest struct via the builder pattern


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


## GetIntegrationWhatsappWebhook

> GetIntegrationWhatsappWebhook(ctx).Execute()

WhatsApp Cloud API subscription challenge



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
	r, err := apiClient.IntegrationAPI.GetIntegrationWhatsappWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetIntegrationWhatsappWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationWhatsappWebhookRequest struct via the builder pattern


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


## PostIntegrationByProviderConnect

> ConnectOut PostIntegrationByProviderConnect(ctx, provider).ConnectIn(connectIn).Execute()

Acquires the org's credential for one provider.



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
	provider := "cloudflare" // string | Provider is the connector's registry id, from the :provider path segment.
	connectIn := *openapiclient.NewConnectIn() // ConnectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationByProviderConnect(context.Background(), provider).ConnectIn(connectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationByProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationByProviderConnect`: ConnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationByProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationByProviderConnectRequest struct via the builder pattern


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


## PostIntegrationByProviderDisconnect

> DisconnectOut PostIntegrationByProviderDisconnect(ctx, provider).Execute()

Revokes (best-effort) and forgets an org's connection: it deletes every custodied KMS secret and the connection row.



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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/integration/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationByProviderDisconnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationByProviderDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationByProviderDisconnect`: DisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationByProviderDisconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/integration/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationByProviderDisconnectRequest struct via the builder pattern


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


## PostIntegrationByProviderVerify

> VerifyOut PostIntegrationByProviderVerify(ctx, provider).Execute()

Re-checks a CONNECTED apikey connector's stored credential against the provider, live (`hanzo connector verify`).



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
	provider := "cloudflare" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/integration/connectors) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationByProviderVerify(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationByProviderVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationByProviderVerify`: VerifyOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationByProviderVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/integration/connectors) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationByProviderVerifyRequest struct via the builder pattern


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


## PostIntegrationConnectorsByIdRefresh

> RefreshOut PostIntegrationConnectorsByIdRefresh(ctx, id).Execute()

Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationConnectorsByIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationConnectorsByIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationConnectorsByIdRefresh`: RefreshOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationConnectorsByIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationConnectorsByIdRefreshRequest struct via the builder pattern


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


## PostIntegrationConnectorsByProviderCredential

> CredentialOut PostIntegrationConnectorsByProviderCredential(ctx, provider).CredentialIn(credentialIn).Execute()

Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI's local PKCE (Adopt).



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	credentialIn := *openapiclient.NewCredentialIn() // CredentialIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationConnectorsByProviderCredential(context.Background(), provider).CredentialIn(credentialIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationConnectorsByProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationConnectorsByProviderCredential`: CredentialOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationConnectorsByProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationConnectorsByProviderCredentialRequest struct via the builder pattern


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


## PostIntegrationConnectorsByProviderDevice

> DeviceStartOut PostIntegrationConnectorsByProviderDevice(ctx, provider).DeviceStartIn(deviceStartIn).Execute()

Begins a device sign-in and returns the code to show the user plus how to poll for completion.



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	deviceStartIn := *openapiclient.NewDeviceStartIn() // DeviceStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationConnectorsByProviderDevice(context.Background(), provider).DeviceStartIn(deviceStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationConnectorsByProviderDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationConnectorsByProviderDevice`: DeviceStartOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationConnectorsByProviderDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationConnectorsByProviderDeviceRequest struct via the builder pattern


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


## PostIntegrationConnectorsByProviderDeviceByFlowPoll

> DevicePollOut PostIntegrationConnectorsByProviderDeviceByFlowPoll(ctx, provider, flow).Execute()

Advances a device sign-in.



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	flow := "g_7f2c" // string | Flow is the id deviceStartOut returned. Expired or another user's flow is indistinguishable from an unknown one: 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationConnectorsByProviderDeviceByFlowPoll(context.Background(), provider, flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationConnectorsByProviderDeviceByFlowPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationConnectorsByProviderDeviceByFlowPoll`: DevicePollOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationConnectorsByProviderDeviceByFlowPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 
**flow** | **string** | Flow is the id deviceStartOut returned. Expired or another user&#39;s flow is indistinguishable from an unknown one: 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationConnectorsByProviderDeviceByFlowPollRequest struct via the builder pattern


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


## PostIntegrationDiscordInteractions

> PostIntegrationDiscordInteractions(ctx).Execute()

Discord interactions endpoint



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
	r, err := apiClient.IntegrationAPI.PostIntegrationDiscordInteractions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationDiscordInteractions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationDiscordInteractionsRequest struct via the builder pattern


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


## PostIntegrationForgeWebhook

> ForgeLaunched PostIntegrationForgeWebhook(ctx).ForgeJob(forgeJob).Execute()

Forge workflow_job webhook



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
	forgeJob := *openapiclient.NewForgeJob() // ForgeJob |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationForgeWebhook(context.Background()).ForgeJob(forgeJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationForgeWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationForgeWebhook`: ForgeLaunched
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationForgeWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationForgeWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgeJob** | [**ForgeJob**](ForgeJob.md) |  | 

### Return type

[**ForgeLaunched**](ForgeLaunched.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationGithubClaim

> GithubClaimOut PostIntegrationGithubClaim(ctx).GithubClaimIn(githubClaimIn).Execute()

Binds installations the App ALREADY holds to the org the caller is acting in — the reconciliation for a grant that happened outside our connect flow.



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
	githubClaimIn := *openapiclient.NewGithubClaimIn() // GithubClaimIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubClaim(context.Background()).GithubClaimIn(githubClaimIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubClaim`: GithubClaimOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubClaimRequest struct via the builder pattern


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


## PostIntegrationGithubFork

> GithubForkOut PostIntegrationGithubFork(ctx).GithubForkReq(githubForkReq).Execute()

Forks a granted repository.



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
	githubForkReq := *openapiclient.NewGithubForkReq() // GithubForkReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubFork(context.Background()).GithubForkReq(githubForkReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubFork`: GithubForkOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubForkRequest struct via the builder pattern


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


## PostIntegrationGithubIssuesBackfill

> GithubBackfillResult PostIntegrationGithubIssuesBackfill(ctx).GithubBackfillIn(githubBackfillIn).Execute()

Seeds the native todo with the EXISTING issues across the org's granted repos (default state=open); the webhook keeps them live thereafter.



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
	githubBackfillIn := *openapiclient.NewGithubBackfillIn() // GithubBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubIssuesBackfill(context.Background()).GithubBackfillIn(githubBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubIssuesBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubIssuesBackfill`: GithubBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubIssuesBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubIssuesBackfillRequest struct via the builder pattern


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


## PostIntegrationGithubReposByRepoPages

> GithubPagesView PostIntegrationGithubReposByRepoPages(ctx, repo).GithubPagesEnableReq(githubPagesEnableReq).Execute()

Creates the repo's Pages site and answers 201 Created with it.



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
	repo := "widgets" // string | Repo is the repository, from the :repo path segment.
	githubPagesEnableReq := *openapiclient.NewGithubPagesEnableReq() // GithubPagesEnableReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubReposByRepoPages(context.Background(), repo).GithubPagesEnableReq(githubPagesEnableReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubReposByRepoPages`: GithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubReposByRepoPagesRequest struct via the builder pattern


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


## PostIntegrationGithubReposByRepoPagesBuilds

> GithubPagesBuildOut PostIntegrationGithubReposByRepoPagesBuilds(ctx, repo).Execute()

Requests a Pages rebuild and returns the queued build's status.



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
	repo := "widgets" // string | Repo is the repository's short name within the org's installation, with no owner prefix (the owner is server-derived from the grant). A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubReposByRepoPagesBuilds(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubReposByRepoPagesBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubReposByRepoPagesBuilds`: GithubPagesBuildOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubReposByRepoPagesBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubReposByRepoPagesBuildsRequest struct via the builder pattern


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


## PostIntegrationGithubReposImport

> GithubImportOut PostIntegrationGithubReposImport(ctx).GithubImportIn(githubImportIn).Execute()

Imports the selected (or all) granted repos into git.hanzo.ai.



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
	githubImportIn := *openapiclient.NewGithubImportIn() // GithubImportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubReposImport(context.Background()).GithubImportIn(githubImportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubReposImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubReposImport`: GithubImportOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubReposImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubReposImportRequest struct via the builder pattern


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


## PostIntegrationGithubSearch

> GithubSearchOut PostIntegrationGithubSearch(ctx).GithubSearchReq(githubSearchReq).Execute()

Finds repositories on GitHub.



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
	githubSearchReq := *openapiclient.NewGithubSearchReq() // GithubSearchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationGithubSearch(context.Background()).GithubSearchReq(githubSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationGithubSearch`: GithubSearchOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationGithubSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubSearchRequest struct via the builder pattern


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


## PostIntegrationGithubWebhook

> PostIntegrationGithubWebhook(ctx).Execute()

GitHub App webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationGithubWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationGithubWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationGithubWebhookRequest struct via the builder pattern


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


## PostIntegrationLinearClaim

> LinearClaimOut PostIntegrationLinearClaim(ctx).LinearClaimIn(linearClaimIn).Execute()

Binds the caller's Linear organization to the org and seals the webhook secret.



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
	linearClaimIn := *openapiclient.NewLinearClaimIn() // LinearClaimIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationLinearClaim(context.Background()).LinearClaimIn(linearClaimIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationLinearClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationLinearClaim`: LinearClaimOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationLinearClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationLinearClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linearClaimIn** | [**LinearClaimIn**](LinearClaimIn.md) |  | 

### Return type

[**LinearClaimOut**](LinearClaimOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationLinearComments

> LinearCommentOut PostIntegrationLinearComments(ctx).LinearCommentIn(linearCommentIn).Execute()

Posts a comment on a Linear issue with the caller's own key, so it carries their name.



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
	linearCommentIn := *openapiclient.NewLinearCommentIn() // LinearCommentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationLinearComments(context.Background()).LinearCommentIn(linearCommentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationLinearComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationLinearComments`: LinearCommentOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationLinearComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationLinearCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linearCommentIn** | [**LinearCommentIn**](LinearCommentIn.md) |  | 

### Return type

[**LinearCommentOut**](LinearCommentOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationLinearIssuesBackfill

> LinearBackfillResult PostIntegrationLinearIssuesBackfill(ctx).LinearBackfillIn(linearBackfillIn).Execute()

Seeds the native todo with the EXISTING Linear issues the caller's key can see (default state=open); the webhook keeps them live thereafter.



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
	linearBackfillIn := *openapiclient.NewLinearBackfillIn() // LinearBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationLinearIssuesBackfill(context.Background()).LinearBackfillIn(linearBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationLinearIssuesBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationLinearIssuesBackfill`: LinearBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationLinearIssuesBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationLinearIssuesBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linearBackfillIn** | [**LinearBackfillIn**](LinearBackfillIn.md) |  | 

### Return type

[**LinearBackfillResult**](LinearBackfillResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationLinearWebhook

> PostIntegrationLinearWebhook(ctx).Execute()

Linear webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationLinearWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationLinearWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationLinearWebhookRequest struct via the builder pattern


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


## PostIntegrationOpenrouterWebhook

> map[string]interface{} PostIntegrationOpenrouterWebhook(ctx).RequestBody(requestBody).Execute()

Receive OpenRouter Broadcast traces as usage rows



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationOpenrouterWebhook(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationOpenrouterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationOpenrouterWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationOpenrouterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationOpenrouterWebhookRequest struct via the builder pattern


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


## PostIntegrationSlackCommands

> PostIntegrationSlackCommands(ctx).Execute()

Slack slash command webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationSlackCommands(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationSlackCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationSlackCommandsRequest struct via the builder pattern


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


## PostIntegrationSlackEvents

> PostIntegrationSlackEvents(ctx).Execute()

Slack Events API webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationSlackEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationSlackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationSlackEventsRequest struct via the builder pattern


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


## PostIntegrationSlackJoin

> SlackJoinOut PostIntegrationSlackJoin(ctx).Execute()

Joins every public channel in the caller org's workspace.



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
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationSlackJoin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationSlackJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationSlackJoin`: SlackJoinOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationSlackJoin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationSlackJoinRequest struct via the builder pattern


### Return type

[**SlackJoinOut**](SlackJoinOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIntegrationTeamsEvents

> PostIntegrationTeamsEvents(ctx).Execute()

Microsoft Teams Bot Framework webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationTeamsEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationTeamsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationTeamsEventsRequest struct via the builder pattern


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


## PostIntegrationTelegramConnect

> AuthorizeOut PostIntegrationTelegramConnect(ctx).Execute()

Mints a short, single-use deep-link code bound to the caller's org and returns the t.me link the console navigates to.



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
	resp, r, err := apiClient.IntegrationAPI.PostIntegrationTelegramConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationTelegramConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIntegrationTelegramConnect`: AuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PostIntegrationTelegramConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationTelegramConnectRequest struct via the builder pattern


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


## PostIntegrationTelegramWebhook

> PostIntegrationTelegramWebhook(ctx).Execute()

Telegram Bot API webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationTelegramWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationTelegramWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationTelegramWebhookRequest struct via the builder pattern


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


## PostIntegrationWhatsappWebhook

> PostIntegrationWhatsappWebhook(ctx).Execute()

WhatsApp Cloud API webhook



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
	r, err := apiClient.IntegrationAPI.PostIntegrationWhatsappWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PostIntegrationWhatsappWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostIntegrationWhatsappWebhookRequest struct via the builder pattern


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


## PutIntegrationGithubReposByRepoPages

> GithubPagesUpdatedOut PutIntegrationGithubReposByRepoPages(ctx, repo).GithubPagesUpdateReq(githubPagesUpdateReq).Execute()

Sets or clears the custom domain (cname) and updates HTTPS enforcement, build type, or source.



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
	repo := "widgets" // string | Repo is the repository, from the :repo path segment.
	githubPagesUpdateReq := *openapiclient.NewGithubPagesUpdateReq() // GithubPagesUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.PutIntegrationGithubReposByRepoPages(context.Background(), repo).GithubPagesUpdateReq(githubPagesUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.PutIntegrationGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIntegrationGithubReposByRepoPages`: GithubPagesUpdatedOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.PutIntegrationGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIntegrationGithubReposByRepoPagesRequest struct via the builder pattern


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

