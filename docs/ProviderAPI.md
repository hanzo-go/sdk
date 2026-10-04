# \ProviderAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnectionById**](ProviderAPI.md#DeleteConnectionById) | **Delete** /v1/connection/{id} | Forgets a connector: every custodied secret, then the row.
[**DeleteProviderGithubReposByRepoPages**](ProviderAPI.md#DeleteProviderGithubReposByRepoPages) | **Delete** /v1/provider/github/repos/{repo}/pages | Deletes the repo&#39;s Pages site.
[**DeleteProviderSlackMessages**](ProviderAPI.md#DeleteProviderSlackMessages) | **Delete** /v1/provider/slack/messages | Takes back one of this app&#39;s messages: DELETE /v1/provider/slack/messages.
[**GetConnection**](ProviderAPI.md#GetConnection) | **Get** /v1/connection | Lists the caller&#39;s OWN connectors across every provider — the set &#x60;hanzo connector ls&#x60; prints.
[**GetConnectionByIdToken**](ProviderAPI.md#GetConnectionByIdToken) | **Get** /v1/connection/{id}/token | Hands the custodied access token to its owner — the ONE place custody exits.
[**GetConnectionProviders**](ProviderAPI.md#GetConnectionProviders) | **Get** /v1/connection/providers | Lists the USER-plane provider cards.
[**GetProvider**](ProviderAPI.md#GetProvider) | **Get** /v1/provider | Returns every registered integration provider together with THIS org&#39;s connection status for it — the catalog the console&#39;s Integrations page renders.
[**GetProviderByProvider**](ProviderAPI.md#GetProviderByProvider) | **Get** /v1/provider/{provider} | Returns ONE provider with this org&#39;s connection status — the same view list carries, for a single id.
[**GetProviderByProviderCallback**](ProviderAPI.md#GetProviderByProviderCallback) | **Get** /v1/provider/{provider}/callback | OAuth return for any connector
[**GetProviderByProviderLogo**](ProviderAPI.md#GetProviderByProviderLogo) | **Get** /v1/provider/{provider}/logo | Serves a connector&#39;s logo from the catalog.
[**GetProviderDiscordLink**](ProviderAPI.md#GetProviderDiscordLink) | **Get** /v1/provider/discord/link | Begin linking a Hanzo account from Discord
[**GetProviderDiscordLinkCallback**](ProviderAPI.md#GetProviderDiscordLinkCallback) | **Get** /v1/provider/discord/link/callback | Complete the Discord account link
[**GetProviderDiscordLinkDiscord**](ProviderAPI.md#GetProviderDiscordLinkDiscord) | **Get** /v1/provider/discord/link/discord | Discord sign-in return leg
[**GetProviderGithubInstallations**](ProviderAPI.md#GetProviderGithubInstallations) | **Get** /v1/provider/github/installations | Lists the GitHub accounts the caller may see the App installed on, each confirmed against the App&#39;s own list, plus where to add another.
[**GetProviderGithubRepos**](ProviderAPI.md#GetProviderGithubRepos) | **Get** /v1/provider/github/repos | Lists the GitHub repositories the caller may work on, most recently pushed first, searched and paged on the server.
[**GetProviderGithubReposByOwnerByRepoBranches**](ProviderAPI.md#GetProviderGithubReposByOwnerByRepoBranches) | **Get** /v1/provider/github/repos/{owner}/{repo}/branches | Lists a repository&#39;s branches, the default first, with prefix search.
[**GetProviderGithubReposByRepoPages**](ProviderAPI.md#GetProviderGithubReposByRepoPages) | **Get** /v1/provider/github/repos/{repo}/pages | Returns the repo&#39;s Pages status, live URL, custom domain and build source.
[**GetProviderGithubUser**](ProviderAPI.md#GetProviderGithubUser) | **Get** /v1/provider/github/user | Reports whether the caller has connected their own GitHub account in the org they act in, and as whom.
[**GetProviderGithubUserCallback**](ProviderAPI.md#GetProviderGithubUserCallback) | **Get** /v1/provider/github/user/callback | Is where GitHub returns the person.
[**GetProviderGitlabProjects**](ProviderAPI.md#GetProviderGitlabProjects) | **Get** /v1/provider/gitlab/projects | Lists the projects the org&#39;s GitLab connection can reach — membership projects, most recently active first.
[**GetProviderSlackChannels**](ProviderAPI.md#GetProviderSlackChannels) | **Get** /v1/provider/slack/channels | Lists Slack conversations for the caller&#39;s connected workspace.
[**GetProviderSlackFile**](ProviderAPI.md#GetProviderSlackFile) | **Get** /v1/provider/slack/file | Read one Slack file&#39;s bytes.
[**GetProviderSlackInstall**](ProviderAPI.md#GetProviderSlackInstall) | **Get** /v1/provider/slack/install | Install the Hanzo app into a Slack workspace
[**GetProviderSlackLink**](ProviderAPI.md#GetProviderSlackLink) | **Get** /v1/provider/slack/link | Begin linking a Hanzo account from Slack
[**GetProviderSlackLinkCallback**](ProviderAPI.md#GetProviderSlackLinkCallback) | **Get** /v1/provider/slack/link/callback | Complete the Slack account link
[**GetProviderSlackLinkSlack**](ProviderAPI.md#GetProviderSlackLinkSlack) | **Get** /v1/provider/slack/link/slack | Slack sign-in return leg
[**GetProviderSlackMessages**](ProviderAPI.md#GetProviderSlackMessages) | **Get** /v1/provider/slack/messages | Reads recent messages from a named Slack channel such as #hanzo-gtm.
[**GetProviderTeamsLink**](ProviderAPI.md#GetProviderTeamsLink) | **Get** /v1/provider/teams/link | Begin linking a Hanzo account from Teams
[**GetProviderTeamsLinkAad**](ProviderAPI.md#GetProviderTeamsLinkAad) | **Get** /v1/provider/teams/link/aad | Microsoft sign-in return leg
[**GetProviderTeamsLinkCallback**](ProviderAPI.md#GetProviderTeamsLinkCallback) | **Get** /v1/provider/teams/link/callback | Complete the Teams account link
[**GetProviderTelegramLink**](ProviderAPI.md#GetProviderTelegramLink) | **Get** /v1/provider/telegram/link | Begin linking a Hanzo account from Telegram
[**GetProviderTelegramLinkAuth**](ProviderAPI.md#GetProviderTelegramLinkAuth) | **Get** /v1/provider/telegram/link/auth | Telegram Login Widget return leg
[**GetProviderTelegramLinkCallback**](ProviderAPI.md#GetProviderTelegramLinkCallback) | **Get** /v1/provider/telegram/link/callback | Complete the Telegram account link
[**GetProviderWhatsappWebhook**](ProviderAPI.md#GetProviderWhatsappWebhook) | **Get** /v1/provider/whatsapp/webhook | WhatsApp Cloud API subscription challenge
[**PostConnectionByIdRefresh**](ProviderAPI.md#PostConnectionByIdRefresh) | **Post** /v1/connection/{id}/refresh | Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.
[**PostConnectionByProviderCredential**](ProviderAPI.md#PostConnectionByProviderCredential) | **Post** /v1/connection/{provider}/credential | Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI&#39;s local PKCE (Adopt).
[**PostConnectionByProviderDevice**](ProviderAPI.md#PostConnectionByProviderDevice) | **Post** /v1/connection/{provider}/device | Begins a device sign-in and returns the code to show the user plus how to poll for completion.
[**PostConnectionByProviderDeviceByFlowPoll**](ProviderAPI.md#PostConnectionByProviderDeviceByFlowPoll) | **Post** /v1/connection/{provider}/device/{flow}/poll | Advances a device sign-in.
[**PostProviderByProviderConnect**](ProviderAPI.md#PostProviderByProviderConnect) | **Post** /v1/provider/{provider}/connect | Acquires the org&#39;s credential for one provider.
[**PostProviderByProviderDisconnect**](ProviderAPI.md#PostProviderByProviderDisconnect) | **Post** /v1/provider/{provider}/disconnect | Revokes (best-effort) and forgets an org&#39;s connection: it deletes every custodied KMS secret and the connection row.
[**PostProviderByProviderRun**](ProviderAPI.md#PostProviderByProviderRun) | **Post** /v1/provider/{provider}/run | Runs one action of a connector as the caller&#39;s org, with the credential the org connected, and answers what the action returned.
[**PostProviderByProviderVerify**](ProviderAPI.md#PostProviderByProviderVerify) | **Post** /v1/provider/{provider}/verify | Re-checks a CONNECTED apikey connector&#39;s stored credential against the provider, live (&#x60;hanzo connector verify&#x60;).
[**PostProviderDiscordInteractions**](ProviderAPI.md#PostProviderDiscordInteractions) | **Post** /v1/provider/discord/interactions | Discord interactions endpoint
[**PostProviderForgeWebhook**](ProviderAPI.md#PostProviderForgeWebhook) | **Post** /v1/provider/forge/webhook | Forge workflow_job webhook
[**PostProviderGithubFork**](ProviderAPI.md#PostProviderGithubFork) | **Post** /v1/provider/github/fork | Forks a granted repository.
[**PostProviderGithubIssuesBackfill**](ProviderAPI.md#PostProviderGithubIssuesBackfill) | **Post** /v1/provider/github/issues/backfill | Seeds the native todo with the EXISTING issues across the org&#39;s granted repos (default state&#x3D;open); the webhook keeps them live thereafter.
[**PostProviderGithubReposByRepoPages**](ProviderAPI.md#PostProviderGithubReposByRepoPages) | **Post** /v1/provider/github/repos/{repo}/pages | Creates the repo&#39;s Pages site and answers 201 Created with it.
[**PostProviderGithubReposByRepoPagesBuilds**](ProviderAPI.md#PostProviderGithubReposByRepoPagesBuilds) | **Post** /v1/provider/github/repos/{repo}/pages/builds | Requests a Pages rebuild and returns the queued build&#39;s status.
[**PostProviderGithubReposImport**](ProviderAPI.md#PostProviderGithubReposImport) | **Post** /v1/provider/github/repos/import | Imports the selected (or all) granted repos into git.hanzo.ai.
[**PostProviderGithubSearch**](ProviderAPI.md#PostProviderGithubSearch) | **Post** /v1/provider/github/search | Finds repositories on GitHub.
[**PostProviderGithubUserComplete**](ProviderAPI.md#PostProviderGithubUserComplete) | **Post** /v1/provider/github/user/complete | Finishes connecting the caller&#39;s GitHub account: it takes the authorization the callback parked, trades it for the person&#39;s token, and seals it.
[**PostProviderGithubUserConnect**](ProviderAPI.md#PostProviderGithubUserConnect) | **Post** /v1/provider/github/user/connect | Begins connecting the caller&#39;s own GitHub account: it answers the App&#39;s authorization page, carrying a state signed for this org and this person.
[**PostProviderGithubUserDisconnect**](ProviderAPI.md#PostProviderGithubUserDisconnect) | **Post** /v1/provider/github/user/disconnect | Forgets the caller&#39;s GitHub connection in this org: the token is revoked at GitHub, and the sealed secrets and the row go.
[**PostProviderGithubWebhook**](ProviderAPI.md#PostProviderGithubWebhook) | **Post** /v1/provider/github/webhook | GitHub App webhook
[**PostProviderLinearClaim**](ProviderAPI.md#PostProviderLinearClaim) | **Post** /v1/provider/linear/claim | Binds the caller&#39;s Linear organization to the org and seals the webhook secret.
[**PostProviderLinearComments**](ProviderAPI.md#PostProviderLinearComments) | **Post** /v1/provider/linear/comments | Posts a comment on a Linear issue with the caller&#39;s own key, so it carries their name.
[**PostProviderLinearIssuesBackfill**](ProviderAPI.md#PostProviderLinearIssuesBackfill) | **Post** /v1/provider/linear/issues/backfill | Seeds the native todo with the EXISTING Linear issues the caller&#39;s key can see (default state&#x3D;open); the webhook keeps them live thereafter.
[**PostProviderLinearWebhook**](ProviderAPI.md#PostProviderLinearWebhook) | **Post** /v1/provider/linear/webhook | Linear webhook
[**PostProviderOpenrouterWebhook**](ProviderAPI.md#PostProviderOpenrouterWebhook) | **Post** /v1/provider/openrouter/webhook | Receive OpenRouter Broadcast traces as usage rows
[**PostProviderSlackCommands**](ProviderAPI.md#PostProviderSlackCommands) | **Post** /v1/provider/slack/commands | Slack slash command webhook
[**PostProviderSlackEvents**](ProviderAPI.md#PostProviderSlackEvents) | **Post** /v1/provider/slack/events | Slack Events API webhook
[**PostProviderSlackJoin**](ProviderAPI.md#PostProviderSlackJoin) | **Post** /v1/provider/slack/join | Joins every public channel in the caller org&#39;s workspace.
[**PostProviderSlackMessages**](ProviderAPI.md#PostProviderSlackMessages) | **Post** /v1/provider/slack/messages | Posts as Hanzo to the caller&#39;s connected Slack workspace.
[**PostProviderSlackReactions**](ProviderAPI.md#PostProviderSlackReactions) | **Post** /v1/provider/slack/reactions | Adds an emoji reaction: POST /v1/provider/slack/reactions.
[**PostProviderSlackSearch**](ProviderAPI.md#PostProviderSlackSearch) | **Post** /v1/provider/slack/search | Answers a workspace question: POST /v1/provider/slack/search.
[**PostProviderTeamsEvents**](ProviderAPI.md#PostProviderTeamsEvents) | **Post** /v1/provider/teams/events | Microsoft Teams Bot Framework webhook
[**PostProviderTelegramConnect**](ProviderAPI.md#PostProviderTelegramConnect) | **Post** /v1/provider/telegram/connect | Mints a short, single-use deep-link code bound to the caller&#39;s org and returns the t.me link the console navigates to.
[**PostProviderTelegramWebhook**](ProviderAPI.md#PostProviderTelegramWebhook) | **Post** /v1/provider/telegram/webhook | Telegram Bot API webhook
[**PostProviderWhatsappWebhook**](ProviderAPI.md#PostProviderWhatsappWebhook) | **Post** /v1/provider/whatsapp/webhook | WhatsApp Cloud API webhook
[**PutProviderGithubReposByRepoPages**](ProviderAPI.md#PutProviderGithubReposByRepoPages) | **Put** /v1/provider/github/repos/{repo}/pages | Sets or clears the custom domain (cname) and updates HTTPS enforcement, build type, or source.
[**PutProviderSlackMessages**](ProviderAPI.md#PutProviderSlackMessages) | **Put** /v1/provider/slack/messages | Rewrites one of this app&#39;s messages: PUT /v1/provider/slack/messages.



## DeleteConnectionById

> ProviderDisconnectOut DeleteConnectionById(ctx, id).Execute()

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
	resp, r, err := apiClient.ProviderAPI.DeleteConnectionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.DeleteConnectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnectionById`: ProviderDisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.DeleteConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderDisconnectOut**](ProviderDisconnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProviderGithubReposByRepoPages

> ProviderGithubPagesDisabledOut DeleteProviderGithubReposByRepoPages(ctx, repo).Execute()

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
	resp, r, err := apiClient.ProviderAPI.DeleteProviderGithubReposByRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.DeleteProviderGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProviderGithubReposByRepoPages`: ProviderGithubPagesDisabledOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.DeleteProviderGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderGithubPagesDisabledOut**](ProviderGithubPagesDisabledOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProviderSlackMessages

> ProviderSlackDeleteMessageOut DeleteProviderSlackMessages(ctx).Channel(channel).Ts(ts).Execute()

Takes back one of this app's messages: DELETE /v1/provider/slack/messages.



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
	channel := "channel_example" // string |  (optional)
	ts := "ts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.DeleteProviderSlackMessages(context.Background()).Channel(channel).Ts(ts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.DeleteProviderSlackMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProviderSlackMessages`: ProviderSlackDeleteMessageOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.DeleteProviderSlackMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderSlackMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** |  | 
 **ts** | **string** |  | 

### Return type

[**ProviderSlackDeleteMessageOut**](ProviderSlackDeleteMessageOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnection

> ProviderConnectionsOut GetConnection(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPI.GetConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnection`: ProviderConnectionsOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


### Return type

[**ProviderConnectionsOut**](ProviderConnectionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionByIdToken

> ProviderTokenOut GetConnectionByIdToken(ctx, id).Execute()

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
	resp, r, err := apiClient.ProviderAPI.GetConnectionByIdToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetConnectionByIdToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionByIdToken`: ProviderTokenOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetConnectionByIdToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionByIdTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderTokenOut**](ProviderTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionProviders

> ProviderUserCatalogOut GetConnectionProviders(ctx).Execute()

Lists the USER-plane provider cards.



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
	resp, r, err := apiClient.ProviderAPI.GetConnectionProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetConnectionProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionProviders`: ProviderUserCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetConnectionProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionProvidersRequest struct via the builder pattern


### Return type

[**ProviderUserCatalogOut**](ProviderUserCatalogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvider

> ProviderListOut GetProvider(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPI.GetProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvider`: ProviderListOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


### Return type

[**ProviderListOut**](ProviderListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderByProvider

> ProviderProviderView GetProviderByProvider(ctx, provider).Execute()

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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/connection) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProviderByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderByProvider`: ProviderProviderView
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/connection) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderProviderView**](ProviderProviderView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderByProviderCallback

> GetProviderByProviderCallback(ctx, provider).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderByProviderCallback(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderByProviderCallback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetProviderByProviderCallbackRequest struct via the builder pattern


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


## GetProviderByProviderLogo

> GetProviderByProviderLogo(ctx, provider).Execute()

Serves a connector's logo from the catalog.



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
	r, err := apiClient.ProviderAPI.GetProviderByProviderLogo(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderByProviderLogo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetProviderByProviderLogoRequest struct via the builder pattern


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


## GetProviderDiscordLink

> GetProviderDiscordLink(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderDiscordLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderDiscordLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderDiscordLinkRequest struct via the builder pattern


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


## GetProviderDiscordLinkCallback

> GetProviderDiscordLinkCallback(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderDiscordLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderDiscordLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderDiscordLinkCallbackRequest struct via the builder pattern


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


## GetProviderDiscordLinkDiscord

> GetProviderDiscordLinkDiscord(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderDiscordLinkDiscord(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderDiscordLinkDiscord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderDiscordLinkDiscordRequest struct via the builder pattern


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


## GetProviderGithubInstallations

> ProviderGithubInstallationsOut GetProviderGithubInstallations(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPI.GetProviderGithubInstallations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGithubInstallations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderGithubInstallations`: ProviderGithubInstallationsOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderGithubInstallations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGithubInstallationsRequest struct via the builder pattern


### Return type

[**ProviderGithubInstallationsOut**](ProviderGithubInstallationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderGithubRepos

> ProviderGithubReposOut GetProviderGithubRepos(ctx).Q(q).Owner(owner).Limit(limit).After(after).Execute()

Lists the GitHub repositories the caller may work on, most recently pushed first, searched and paged on the server.



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
	q := "cloud" // string | Q keeps repositories whose owner/name contains it, case-insensitively. (optional)
	owner := "owner_example" // string | Owner keeps one GitHub account's repositories — an org or a user login. (optional)
	limit := int64(50) // int64 | Limit is the page size: 1 to 100, and 50 when absent or unreadable. (optional)
	after := "after_example" // string | After is the `next` of the previous page. Absent starts at the most recently pushed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProviderGithubRepos(context.Background()).Q(q).Owner(owner).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGithubRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderGithubRepos`: ProviderGithubReposOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderGithubRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGithubReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q keeps repositories whose owner/name contains it, case-insensitively. | 
 **owner** | **string** | Owner keeps one GitHub account&#39;s repositories — an org or a user login. | 
 **limit** | **int64** | Limit is the page size: 1 to 100, and 50 when absent or unreadable. | 
 **after** | **string** | After is the &#x60;next&#x60; of the previous page. Absent starts at the most recently pushed. | 

### Return type

[**ProviderGithubReposOut**](ProviderGithubReposOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderGithubReposByOwnerByRepoBranches

> ProviderGithubBranchesOut GetProviderGithubReposByOwnerByRepoBranches(ctx, owner, repo).Q(q).Limit(limit).After(after).Execute()

Lists a repository's branches, the default first, with prefix search.



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
	owner := "hanzoai" // string | Owner is the GitHub account that holds the repository.
	repo := "cloud" // string | Repo is the repository's name within that account.
	q := "feat/" // string | Q keeps branches whose name starts with it, matched by GitHub. (optional)
	limit := int64(789) // int64 | Limit is the page size: 1 to 100, and 50 when absent or unreadable. (optional)
	after := "after_example" // string | After is the `next` of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProviderGithubReposByOwnerByRepoBranches(context.Background(), owner, repo).Q(q).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGithubReposByOwnerByRepoBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderGithubReposByOwnerByRepoBranches`: ProviderGithubBranchesOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderGithubReposByOwnerByRepoBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner is the GitHub account that holds the repository. | 
**repo** | **string** | Repo is the repository&#39;s name within that account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGithubReposByOwnerByRepoBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Q keeps branches whose name starts with it, matched by GitHub. | 
 **limit** | **int64** | Limit is the page size: 1 to 100, and 50 when absent or unreadable. | 
 **after** | **string** | After is the &#x60;next&#x60; of the previous page. | 

### Return type

[**ProviderGithubBranchesOut**](ProviderGithubBranchesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderGithubReposByRepoPages

> ProviderGithubPagesView GetProviderGithubReposByRepoPages(ctx, repo).Execute()

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
	resp, r, err := apiClient.ProviderAPI.GetProviderGithubReposByRepoPages(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderGithubReposByRepoPages`: ProviderGithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderGithubPagesView**](ProviderGithubPagesView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderGithubUser

> ProviderGithubUserOut GetProviderGithubUser(ctx).Execute()

Reports whether the caller has connected their own GitHub account in the org they act in, and as whom.



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
	resp, r, err := apiClient.ProviderAPI.GetProviderGithubUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGithubUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderGithubUser`: ProviderGithubUserOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderGithubUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGithubUserRequest struct via the builder pattern


### Return type

[**ProviderGithubUserOut**](ProviderGithubUserOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderGithubUserCallback

> GetProviderGithubUserCallback(ctx).Execute()

Is where GitHub returns the person.



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
	r, err := apiClient.ProviderAPI.GetProviderGithubUserCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGithubUserCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGithubUserCallbackRequest struct via the builder pattern


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


## GetProviderGitlabProjects

> ProviderGitlabProjectsOut GetProviderGitlabProjects(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPI.GetProviderGitlabProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderGitlabProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderGitlabProjects`: ProviderGitlabProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderGitlabProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderGitlabProjectsRequest struct via the builder pattern


### Return type

[**ProviderGitlabProjectsOut**](ProviderGitlabProjectsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderSlackChannels

> ProviderSlackChannelsOut GetProviderSlackChannels(ctx).Types(types).Cursor(cursor).Execute()

Lists Slack conversations for the caller's connected workspace.



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
	types := "types_example" // string | Types selects public_channel, private_channel, im or mpim, comma-separated. Defaults to public_channel; each selected type requires its own read scope. (optional)
	cursor := "cursor_example" // string | Cursor continues the previous page's next_cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProviderSlackChannels(context.Background()).Types(types).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderSlackChannels`: ProviderSlackChannelsOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderSlackChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **string** | Types selects public_channel, private_channel, im or mpim, comma-separated. Defaults to public_channel; each selected type requires its own read scope. | 
 **cursor** | **string** | Cursor continues the previous page&#39;s next_cursor. | 

### Return type

[**ProviderSlackChannelsOut**](ProviderSlackChannelsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderSlackFile

> GetProviderSlackFile(ctx).Execute()

Read one Slack file's bytes.



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
	r, err := apiClient.ProviderAPI.GetProviderSlackFile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackFileRequest struct via the builder pattern


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


## GetProviderSlackInstall

> GetProviderSlackInstall(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderSlackInstall(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackInstallRequest struct via the builder pattern


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


## GetProviderSlackLink

> GetProviderSlackLink(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderSlackLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackLinkRequest struct via the builder pattern


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


## GetProviderSlackLinkCallback

> GetProviderSlackLinkCallback(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderSlackLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackLinkCallbackRequest struct via the builder pattern


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


## GetProviderSlackLinkSlack

> GetProviderSlackLinkSlack(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderSlackLinkSlack(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackLinkSlack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackLinkSlackRequest struct via the builder pattern


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


## GetProviderSlackMessages

> ProviderSlackMessagesOut GetProviderSlackMessages(ctx).Channel(channel).Limit(limit).Cursor(cursor).ThreadTs(threadTs).Oldest(oldest).Latest(latest).Execute()

Reads recent messages from a named Slack channel such as #hanzo-gtm.



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
	channel := "channel_example" // string | Channel is a Slack channel ID, name (#hanzo-gtm), or Slack channel mention. (optional)
	limit := int64(789) // int64 | Limit bounds the page to 1-100 messages; zero defaults to 15. (optional)
	cursor := "cursor_example" // string | Cursor continues the previous page's next_cursor. (optional)
	threadTs := "threadTs_example" // string | ThreadTS reads a thread via conversations.replies. Empty reads recent channel messages via conversations.history, which excludes thread replies. (optional)
	oldest := "oldest_example" // string | Oldest filters to messages after this Slack timestamp (exclusive). (optional)
	latest := "latest_example" // string | Latest filters to messages before this Slack timestamp (exclusive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.GetProviderSlackMessages(context.Background()).Channel(channel).Limit(limit).Cursor(cursor).ThreadTs(threadTs).Oldest(oldest).Latest(latest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderSlackMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderSlackMessages`: ProviderSlackMessagesOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.GetProviderSlackMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderSlackMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is a Slack channel ID, name (#hanzo-gtm), or Slack channel mention. | 
 **limit** | **int64** | Limit bounds the page to 1-100 messages; zero defaults to 15. | 
 **cursor** | **string** | Cursor continues the previous page&#39;s next_cursor. | 
 **threadTs** | **string** | ThreadTS reads a thread via conversations.replies. Empty reads recent channel messages via conversations.history, which excludes thread replies. | 
 **oldest** | **string** | Oldest filters to messages after this Slack timestamp (exclusive). | 
 **latest** | **string** | Latest filters to messages before this Slack timestamp (exclusive). | 

### Return type

[**ProviderSlackMessagesOut**](ProviderSlackMessagesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderTeamsLink

> GetProviderTeamsLink(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderTeamsLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderTeamsLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTeamsLinkRequest struct via the builder pattern


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


## GetProviderTeamsLinkAad

> GetProviderTeamsLinkAad(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderTeamsLinkAad(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderTeamsLinkAad``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTeamsLinkAadRequest struct via the builder pattern


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


## GetProviderTeamsLinkCallback

> GetProviderTeamsLinkCallback(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderTeamsLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderTeamsLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTeamsLinkCallbackRequest struct via the builder pattern


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


## GetProviderTelegramLink

> GetProviderTelegramLink(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderTelegramLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderTelegramLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTelegramLinkRequest struct via the builder pattern


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


## GetProviderTelegramLinkAuth

> GetProviderTelegramLinkAuth(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderTelegramLinkAuth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderTelegramLinkAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTelegramLinkAuthRequest struct via the builder pattern


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


## GetProviderTelegramLinkCallback

> GetProviderTelegramLinkCallback(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderTelegramLinkCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderTelegramLinkCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTelegramLinkCallbackRequest struct via the builder pattern


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


## GetProviderWhatsappWebhook

> GetProviderWhatsappWebhook(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.GetProviderWhatsappWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.GetProviderWhatsappWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderWhatsappWebhookRequest struct via the builder pattern


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


## PostConnectionByIdRefresh

> ProviderRefreshOut PostConnectionByIdRefresh(ctx, id).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostConnectionByIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostConnectionByIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectionByIdRefresh`: ProviderRefreshOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostConnectionByIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectionByIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderRefreshOut**](ProviderRefreshOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectionByProviderCredential

> ProviderCredentialOut PostConnectionByProviderCredential(ctx, provider).ProviderCredentialIn(providerCredentialIn).Execute()

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
	providerCredentialIn := *openapiclient.NewProviderCredentialIn() // ProviderCredentialIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostConnectionByProviderCredential(context.Background(), provider).ProviderCredentialIn(providerCredentialIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostConnectionByProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectionByProviderCredential`: ProviderCredentialOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostConnectionByProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectionByProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerCredentialIn** | [**ProviderCredentialIn**](ProviderCredentialIn.md) |  | 

### Return type

[**ProviderCredentialOut**](ProviderCredentialOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectionByProviderDevice

> ProviderDeviceStartOut PostConnectionByProviderDevice(ctx, provider).ProviderDeviceStartIn(providerDeviceStartIn).Execute()

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
	providerDeviceStartIn := *openapiclient.NewProviderDeviceStartIn() // ProviderDeviceStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostConnectionByProviderDevice(context.Background(), provider).ProviderDeviceStartIn(providerDeviceStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostConnectionByProviderDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectionByProviderDevice`: ProviderDeviceStartOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostConnectionByProviderDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectionByProviderDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerDeviceStartIn** | [**ProviderDeviceStartIn**](ProviderDeviceStartIn.md) |  | 

### Return type

[**ProviderDeviceStartOut**](ProviderDeviceStartOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectionByProviderDeviceByFlowPoll

> ProviderDevicePollOut PostConnectionByProviderDeviceByFlowPoll(ctx, provider, flow).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostConnectionByProviderDeviceByFlowPoll(context.Background(), provider, flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostConnectionByProviderDeviceByFlowPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectionByProviderDeviceByFlowPoll`: ProviderDevicePollOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostConnectionByProviderDeviceByFlowPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 
**flow** | **string** | Flow is the id deviceStartOut returned. Expired or another user&#39;s flow is indistinguishable from an unknown one: 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectionByProviderDeviceByFlowPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProviderDevicePollOut**](ProviderDevicePollOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderByProviderConnect

> ProviderConnectOut PostProviderByProviderConnect(ctx, provider).ProviderConnectIn(providerConnectIn).Execute()

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
	providerConnectIn := *openapiclient.NewProviderConnectIn() // ProviderConnectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderByProviderConnect(context.Background(), provider).ProviderConnectIn(providerConnectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderByProviderConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderByProviderConnect`: ProviderConnectOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderByProviderConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderByProviderConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerConnectIn** | [**ProviderConnectIn**](ProviderConnectIn.md) |  | 

### Return type

[**ProviderConnectOut**](ProviderConnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderByProviderDisconnect

> ProviderDisconnectOut PostProviderByProviderDisconnect(ctx, provider).Execute()

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
	provider := "slack" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/connection) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderByProviderDisconnect(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderByProviderDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderByProviderDisconnect`: ProviderDisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderByProviderDisconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/connection) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderByProviderDisconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderDisconnectOut**](ProviderDisconnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderByProviderRun

> ProviderRunOut PostProviderByProviderRun(ctx, provider).ProviderRunIn(providerRunIn).Execute()

Runs one action of a connector as the caller's org, with the credential the org connected, and answers what the action returned.



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
	provider := "bitly" // string | Provider is the connector, from the path.
	providerRunIn := *openapiclient.NewProviderRunIn() // ProviderRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderByProviderRun(context.Background(), provider).ProviderRunIn(providerRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderByProviderRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderByProviderRun`: ProviderRunOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderByProviderRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the connector, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderByProviderRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerRunIn** | [**ProviderRunIn**](ProviderRunIn.md) |  | 

### Return type

[**ProviderRunOut**](ProviderRunOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderByProviderVerify

> ProviderVerifyOut PostProviderByProviderVerify(ctx, provider).Execute()

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
	provider := "cloudflare" // string | Provider is the registry id of the connector — \"slack\", \"github\", \"cloudflare\". Unknown ids are 404, as are the user-plane (/v1/connection) providers, which this surface never resolves.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderByProviderVerify(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderByProviderVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderByProviderVerify`: ProviderVerifyOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderByProviderVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the registry id of the connector — \&quot;slack\&quot;, \&quot;github\&quot;, \&quot;cloudflare\&quot;. Unknown ids are 404, as are the user-plane (/v1/connection) providers, which this surface never resolves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderByProviderVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderVerifyOut**](ProviderVerifyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderDiscordInteractions

> PostProviderDiscordInteractions(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderDiscordInteractions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderDiscordInteractions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderDiscordInteractionsRequest struct via the builder pattern


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


## PostProviderForgeWebhook

> ForgeLaunched PostProviderForgeWebhook(ctx).ForgeJob(forgeJob).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostProviderForgeWebhook(context.Background()).ForgeJob(forgeJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderForgeWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderForgeWebhook`: ForgeLaunched
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderForgeWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderForgeWebhookRequest struct via the builder pattern


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


## PostProviderGithubFork

> ProviderGithubForkOut PostProviderGithubFork(ctx).ProviderGithubForkReq(providerGithubForkReq).Execute()

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
	providerGithubForkReq := *openapiclient.NewProviderGithubForkReq() // ProviderGithubForkReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubFork(context.Background()).ProviderGithubForkReq(providerGithubForkReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubFork`: ProviderGithubForkOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerGithubForkReq** | [**ProviderGithubForkReq**](ProviderGithubForkReq.md) |  | 

### Return type

[**ProviderGithubForkOut**](ProviderGithubForkOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubIssuesBackfill

> ProviderGithubBackfillResult PostProviderGithubIssuesBackfill(ctx).ProviderGithubBackfillIn(providerGithubBackfillIn).Execute()

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
	providerGithubBackfillIn := *openapiclient.NewProviderGithubBackfillIn() // ProviderGithubBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubIssuesBackfill(context.Background()).ProviderGithubBackfillIn(providerGithubBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubIssuesBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubIssuesBackfill`: ProviderGithubBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubIssuesBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubIssuesBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerGithubBackfillIn** | [**ProviderGithubBackfillIn**](ProviderGithubBackfillIn.md) |  | 

### Return type

[**ProviderGithubBackfillResult**](ProviderGithubBackfillResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubReposByRepoPages

> ProviderGithubPagesView PostProviderGithubReposByRepoPages(ctx, repo).ProviderGithubPagesEnableReq(providerGithubPagesEnableReq).Execute()

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
	providerGithubPagesEnableReq := *openapiclient.NewProviderGithubPagesEnableReq() // ProviderGithubPagesEnableReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubReposByRepoPages(context.Background(), repo).ProviderGithubPagesEnableReq(providerGithubPagesEnableReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubReposByRepoPages`: ProviderGithubPagesView
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerGithubPagesEnableReq** | [**ProviderGithubPagesEnableReq**](ProviderGithubPagesEnableReq.md) |  | 

### Return type

[**ProviderGithubPagesView**](ProviderGithubPagesView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubReposByRepoPagesBuilds

> ProviderGithubPagesBuildOut PostProviderGithubReposByRepoPagesBuilds(ctx, repo).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubReposByRepoPagesBuilds(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubReposByRepoPagesBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubReposByRepoPagesBuilds`: ProviderGithubPagesBuildOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubReposByRepoPagesBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository&#39;s short name within the org&#39;s installation, with no owner prefix (the owner is server-derived from the grant). A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubReposByRepoPagesBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderGithubPagesBuildOut**](ProviderGithubPagesBuildOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubReposImport

> ProviderGithubImportOut PostProviderGithubReposImport(ctx).ProviderGithubImportIn(providerGithubImportIn).Execute()

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
	providerGithubImportIn := *openapiclient.NewProviderGithubImportIn() // ProviderGithubImportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubReposImport(context.Background()).ProviderGithubImportIn(providerGithubImportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubReposImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubReposImport`: ProviderGithubImportOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubReposImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubReposImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerGithubImportIn** | [**ProviderGithubImportIn**](ProviderGithubImportIn.md) |  | 

### Return type

[**ProviderGithubImportOut**](ProviderGithubImportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubSearch

> ProviderGithubSearchOut PostProviderGithubSearch(ctx).ProviderGithubSearchReq(providerGithubSearchReq).Execute()

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
	providerGithubSearchReq := *openapiclient.NewProviderGithubSearchReq() // ProviderGithubSearchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubSearch(context.Background()).ProviderGithubSearchReq(providerGithubSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubSearch`: ProviderGithubSearchOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerGithubSearchReq** | [**ProviderGithubSearchReq**](ProviderGithubSearchReq.md) |  | 

### Return type

[**ProviderGithubSearchOut**](ProviderGithubSearchOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubUserComplete

> ProviderGithubUserOut PostProviderGithubUserComplete(ctx).ProviderGithubUserCompleteIn(providerGithubUserCompleteIn).Execute()

Finishes connecting the caller's GitHub account: it takes the authorization the callback parked, trades it for the person's token, and seals it.



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
	providerGithubUserCompleteIn := *openapiclient.NewProviderGithubUserCompleteIn() // ProviderGithubUserCompleteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubUserComplete(context.Background()).ProviderGithubUserCompleteIn(providerGithubUserCompleteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubUserComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubUserComplete`: ProviderGithubUserOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubUserComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubUserCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerGithubUserCompleteIn** | [**ProviderGithubUserCompleteIn**](ProviderGithubUserCompleteIn.md) |  | 

### Return type

[**ProviderGithubUserOut**](ProviderGithubUserOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubUserConnect

> ProviderGithubUserConnectOut PostProviderGithubUserConnect(ctx).ProviderGithubUserConnectIn(providerGithubUserConnectIn).Execute()

Begins connecting the caller's own GitHub account: it answers the App's authorization page, carrying a state signed for this org and this person.



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
	providerGithubUserConnectIn := *openapiclient.NewProviderGithubUserConnectIn() // ProviderGithubUserConnectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubUserConnect(context.Background()).ProviderGithubUserConnectIn(providerGithubUserConnectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubUserConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubUserConnect`: ProviderGithubUserConnectOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubUserConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubUserConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerGithubUserConnectIn** | [**ProviderGithubUserConnectIn**](ProviderGithubUserConnectIn.md) |  | 

### Return type

[**ProviderGithubUserConnectOut**](ProviderGithubUserConnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubUserDisconnect

> ProviderGithubUserDisconnectOut PostProviderGithubUserDisconnect(ctx).Execute()

Forgets the caller's GitHub connection in this org: the token is revoked at GitHub, and the sealed secrets and the row go.



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
	resp, r, err := apiClient.ProviderAPI.PostProviderGithubUserDisconnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubUserDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderGithubUserDisconnect`: ProviderGithubUserDisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderGithubUserDisconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubUserDisconnectRequest struct via the builder pattern


### Return type

[**ProviderGithubUserDisconnectOut**](ProviderGithubUserDisconnectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderGithubWebhook

> PostProviderGithubWebhook(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderGithubWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderGithubWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderGithubWebhookRequest struct via the builder pattern


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


## PostProviderLinearClaim

> ProviderLinearClaimOut PostProviderLinearClaim(ctx).ProviderLinearClaimIn(providerLinearClaimIn).Execute()

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
	providerLinearClaimIn := *openapiclient.NewProviderLinearClaimIn() // ProviderLinearClaimIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderLinearClaim(context.Background()).ProviderLinearClaimIn(providerLinearClaimIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderLinearClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderLinearClaim`: ProviderLinearClaimOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderLinearClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderLinearClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerLinearClaimIn** | [**ProviderLinearClaimIn**](ProviderLinearClaimIn.md) |  | 

### Return type

[**ProviderLinearClaimOut**](ProviderLinearClaimOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderLinearComments

> ProviderLinearCommentOut PostProviderLinearComments(ctx).ProviderLinearCommentIn(providerLinearCommentIn).Execute()

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
	providerLinearCommentIn := *openapiclient.NewProviderLinearCommentIn() // ProviderLinearCommentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderLinearComments(context.Background()).ProviderLinearCommentIn(providerLinearCommentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderLinearComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderLinearComments`: ProviderLinearCommentOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderLinearComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderLinearCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerLinearCommentIn** | [**ProviderLinearCommentIn**](ProviderLinearCommentIn.md) |  | 

### Return type

[**ProviderLinearCommentOut**](ProviderLinearCommentOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderLinearIssuesBackfill

> ProviderLinearBackfillResult PostProviderLinearIssuesBackfill(ctx).ProviderLinearBackfillIn(providerLinearBackfillIn).Execute()

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
	providerLinearBackfillIn := *openapiclient.NewProviderLinearBackfillIn() // ProviderLinearBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderLinearIssuesBackfill(context.Background()).ProviderLinearBackfillIn(providerLinearBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderLinearIssuesBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderLinearIssuesBackfill`: ProviderLinearBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderLinearIssuesBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderLinearIssuesBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerLinearBackfillIn** | [**ProviderLinearBackfillIn**](ProviderLinearBackfillIn.md) |  | 

### Return type

[**ProviderLinearBackfillResult**](ProviderLinearBackfillResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderLinearWebhook

> PostProviderLinearWebhook(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderLinearWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderLinearWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderLinearWebhookRequest struct via the builder pattern


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


## PostProviderOpenrouterWebhook

> map[string]interface{} PostProviderOpenrouterWebhook(ctx).RequestBody(requestBody).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostProviderOpenrouterWebhook(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderOpenrouterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderOpenrouterWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderOpenrouterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderOpenrouterWebhookRequest struct via the builder pattern


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


## PostProviderSlackCommands

> PostProviderSlackCommands(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderSlackCommands(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderSlackCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderSlackCommandsRequest struct via the builder pattern


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


## PostProviderSlackEvents

> PostProviderSlackEvents(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderSlackEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderSlackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderSlackEventsRequest struct via the builder pattern


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


## PostProviderSlackJoin

> ProviderSlackJoinOut PostProviderSlackJoin(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostProviderSlackJoin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderSlackJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderSlackJoin`: ProviderSlackJoinOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderSlackJoin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderSlackJoinRequest struct via the builder pattern


### Return type

[**ProviderSlackJoinOut**](ProviderSlackJoinOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderSlackMessages

> ProviderSlackSendMessageOut PostProviderSlackMessages(ctx).ProviderSlackSendMessageIn(providerSlackSendMessageIn).Execute()

Posts as Hanzo to the caller's connected Slack workspace.



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
	providerSlackSendMessageIn := *openapiclient.NewProviderSlackSendMessageIn() // ProviderSlackSendMessageIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderSlackMessages(context.Background()).ProviderSlackSendMessageIn(providerSlackSendMessageIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderSlackMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderSlackMessages`: ProviderSlackSendMessageOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderSlackMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderSlackMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerSlackSendMessageIn** | [**ProviderSlackSendMessageIn**](ProviderSlackSendMessageIn.md) |  | 

### Return type

[**ProviderSlackSendMessageOut**](ProviderSlackSendMessageOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderSlackReactions

> ProviderSlackReactOut PostProviderSlackReactions(ctx).ProviderSlackReactIn(providerSlackReactIn).Execute()

Adds an emoji reaction: POST /v1/provider/slack/reactions.



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
	providerSlackReactIn := *openapiclient.NewProviderSlackReactIn() // ProviderSlackReactIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderSlackReactions(context.Background()).ProviderSlackReactIn(providerSlackReactIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderSlackReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderSlackReactions`: ProviderSlackReactOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderSlackReactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderSlackReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerSlackReactIn** | [**ProviderSlackReactIn**](ProviderSlackReactIn.md) |  | 

### Return type

[**ProviderSlackReactOut**](ProviderSlackReactOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderSlackSearch

> ProviderSlackSearchOut PostProviderSlackSearch(ctx).ProviderSlackSearchIn(providerSlackSearchIn).Execute()

Answers a workspace question: POST /v1/provider/slack/search.



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
	providerSlackSearchIn := *openapiclient.NewProviderSlackSearchIn() // ProviderSlackSearchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PostProviderSlackSearch(context.Background()).ProviderSlackSearchIn(providerSlackSearchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderSlackSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderSlackSearch`: ProviderSlackSearchOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderSlackSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderSlackSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerSlackSearchIn** | [**ProviderSlackSearchIn**](ProviderSlackSearchIn.md) |  | 

### Return type

[**ProviderSlackSearchOut**](ProviderSlackSearchOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderTeamsEvents

> PostProviderTeamsEvents(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderTeamsEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderTeamsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderTeamsEventsRequest struct via the builder pattern


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


## PostProviderTelegramConnect

> ProviderAuthorizeOut PostProviderTelegramConnect(ctx).Execute()

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
	resp, r, err := apiClient.ProviderAPI.PostProviderTelegramConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderTelegramConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProviderTelegramConnect`: ProviderAuthorizeOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PostProviderTelegramConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderTelegramConnectRequest struct via the builder pattern


### Return type

[**ProviderAuthorizeOut**](ProviderAuthorizeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProviderTelegramWebhook

> PostProviderTelegramWebhook(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderTelegramWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderTelegramWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderTelegramWebhookRequest struct via the builder pattern


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


## PostProviderWhatsappWebhook

> PostProviderWhatsappWebhook(ctx).Execute()

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
	r, err := apiClient.ProviderAPI.PostProviderWhatsappWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PostProviderWhatsappWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostProviderWhatsappWebhookRequest struct via the builder pattern


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


## PutProviderGithubReposByRepoPages

> ProviderGithubPagesUpdatedOut PutProviderGithubReposByRepoPages(ctx, repo).ProviderGithubPagesUpdateReq(providerGithubPagesUpdateReq).Execute()

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
	providerGithubPagesUpdateReq := *openapiclient.NewProviderGithubPagesUpdateReq() // ProviderGithubPagesUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PutProviderGithubReposByRepoPages(context.Background(), repo).ProviderGithubPagesUpdateReq(providerGithubPagesUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PutProviderGithubReposByRepoPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProviderGithubReposByRepoPages`: ProviderGithubPagesUpdatedOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PutProviderGithubReposByRepoPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo is the repository, from the :repo path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProviderGithubReposByRepoPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerGithubPagesUpdateReq** | [**ProviderGithubPagesUpdateReq**](ProviderGithubPagesUpdateReq.md) |  | 

### Return type

[**ProviderGithubPagesUpdatedOut**](ProviderGithubPagesUpdatedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProviderSlackMessages

> ProviderSlackUpdateMessageOut PutProviderSlackMessages(ctx).ProviderSlackUpdateMessageIn(providerSlackUpdateMessageIn).Execute()

Rewrites one of this app's messages: PUT /v1/provider/slack/messages.



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
	providerSlackUpdateMessageIn := *openapiclient.NewProviderSlackUpdateMessageIn() // ProviderSlackUpdateMessageIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderAPI.PutProviderSlackMessages(context.Background()).ProviderSlackUpdateMessageIn(providerSlackUpdateMessageIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderAPI.PutProviderSlackMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProviderSlackMessages`: ProviderSlackUpdateMessageOut
	fmt.Fprintf(os.Stdout, "Response from `ProviderAPI.PutProviderSlackMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutProviderSlackMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerSlackUpdateMessageIn** | [**ProviderSlackUpdateMessageIn**](ProviderSlackUpdateMessageIn.md) |  | 

### Return type

[**ProviderSlackUpdateMessageOut**](ProviderSlackUpdateMessageOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

