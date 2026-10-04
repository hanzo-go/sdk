# \AgentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgentByRef**](AgentAPI.md#DeleteAgentByRef) | **Delete** /v1/agent/{ref} | Removes an agent and every run recorded against it.
[**DeleteAgentChatConversationsByIdSharesByShare**](AgentAPI.md#DeleteAgentChatConversationsByIdSharesByShare) | **Delete** /v1/agent/chat/conversations/{id}/shares/{share} | Revoke a link to one of your conversations
[**DeleteAgentChatConversationsByIdSharesByShareViewersByViewer**](AgentAPI.md#DeleteAgentChatConversationsByIdSharesByShareViewersByViewer) | **Delete** /v1/agent/chat/conversations/{id}/shares/{share}/viewers/{viewer} | Remove one viewer from a link
[**DeleteAgentChatSharesByShare**](AgentAPI.md#DeleteAgentChatSharesByShare) | **Delete** /v1/agent/chat/shares/{share} | Revoke any link in your organization
[**DeleteAgentTargetsById**](AgentAPI.md#DeleteAgentTargetsById) | **Delete** /v1/agent/targets/{id} | Deregisters one machine.
[**GetAgent**](AgentAPI.md#GetAgent) | **Get** /v1/agent | Returns every agent defined in the caller&#39;s org, each with the number of runs recorded against it.
[**GetAgentActivity**](AgentAPI.md#GetAgentActivity) | **Get** /v1/agent/activity | Serves the org-wide recent-activity feed.
[**GetAgentBuilds**](AgentAPI.md#GetAgentBuilds) | **Get** /v1/agent/builds | Returns the public index of every published build, most recently updated first, so a gallery can link straight to the story behind each product.
[**GetAgentBuildsByOrgByProject**](AgentAPI.md#GetAgentBuildsByOrgByProject) | **Get** /v1/agent/builds/{org}/{project} | Returns the readable build of one product: the agent session that produced it, turn by turn — the prompts, the reasoning, the commits each turn produced — plus the exact &#x60;git log&#x60; that re-derives every commit binding from git itself, so nothing here has to be taken on trust.
[**GetAgentByRef**](AgentAPI.md#GetAgentByRef) | **Get** /v1/agent/{ref} | Returns one agent with its system prompt and its 20 most recent runs.
[**GetAgentByRefRuns**](AgentAPI.md#GetAgentByRefRuns) | **Get** /v1/agent/{ref}/runs | Returns one agent&#39;s execution history, newest first — each run&#39;s input, its output or its error, and how long it took.
[**GetAgentByRefSpend**](AgentAPI.md#GetAgentByRefSpend) | **Get** /v1/agent/{ref}/spend | Answers what one of your org&#39;s agents has spent, in integer micro-USD.
[**GetAgentChatConversations**](AgentAPI.md#GetAgentChatConversations) | **Get** /v1/agent/chat/conversations | List the agent threads in your org
[**GetAgentChatConversationsById**](AgentAPI.md#GetAgentChatConversationsById) | **Get** /v1/agent/chat/conversations/{id} | Read one agent thread in full
[**GetAgentChatConversationsByIdShares**](AgentAPI.md#GetAgentChatConversationsByIdShares) | **Get** /v1/agent/chat/conversations/{id}/shares | List the live links to one of your conversations
[**GetAgentChatPresets**](AgentAPI.md#GetAgentChatPresets) | **Get** /v1/agent/chat/presets | List the agent presets available to a caller
[**GetAgentChatShared**](AgentAPI.md#GetAgentChatShared) | **Get** /v1/agent/chat/shared | List the chats shared with you
[**GetAgentChatSharedByShare**](AgentAPI.md#GetAgentChatSharedByShare) | **Get** /v1/agent/chat/shared/{share} | Read a chat shared with you
[**GetAgentChatShares**](AgentAPI.md#GetAgentChatShares) | **Get** /v1/agent/chat/shares | List every live link in your organization
[**GetAgentCodingBySessionArtifacts**](AgentAPI.md#GetAgentCodingBySessionArtifacts) | **Get** /v1/agent/coding/{session}/artifacts | Lists what a coding run left, kept after its sandbox is gone: every file it added or changed and its whole change as &#x60;changes.patch&#x60;, stored beside the run; the ports it served, each a preview while its sandbox is kept; its pull request and where its work was published.
[**GetAgentCodingBySessionBlob**](AgentAPI.md#GetAgentCodingBySessionBlob) | **Get** /v1/agent/coding/{session}/blob | Returns one file of a coding run&#39;s repository, read where the tree is read, in the shape GET /v1/git/repos/{name}/blob answers: text verbatim, anything else base64, and a file past the 1 MiB view cap marked truncated with no content.
[**GetAgentCodingBySessionChanges**](AgentAPI.md#GetAgentCodingBySessionChanges) | **Get** /v1/agent/coding/{session}/changes | Returns what a coding run changed, read from the forge it pushed its branch to: the commits on its branch that the base does not have, newest first; the net change of the branch against its base, one entry per file with that file&#39;s patch; and its pull request with the reviews it has had, or null while it has none.
[**GetAgentCodingBySessionTree**](AgentAPI.md#GetAgentCodingBySessionTree) | **Get** /v1/agent/coding/{session}/tree | Lists one directory of a coding run&#39;s repository, one level down with directories first: at the run&#39;s own branch once the forge holds it, and at the branch it started from until then — &#x60;ref&#x60; says which.
[**GetAgentMetrics**](AgentAPI.md#GetAgentMetrics) | **Get** /v1/agent/metrics | Serves the invocations-over-time histogram for the org&#39;s Agents dashboard.
[**GetAgentRuns**](AgentAPI.md#GetAgentRuns) | **Get** /v1/agent/runs | Returns the org&#39;s agent runs across EVERY agent, newest first — what ran here, for whom, on which model, how long it took, and why it failed.
[**GetAgentSessions**](AgentAPI.md#GetAgentSessions) | **Get** /v1/agent/sessions | Returns the caller org&#39;s live sessions, newest first — each with its event count, its direct-child count and a one-line preview of its latest event.
[**GetAgentSessionsById**](AgentAPI.md#GetAgentSessionsById) | **Get** /v1/agent/sessions/{id} | Returns one session with its direct child sessions and its 50 most recent events, oldest of those first.
[**GetAgentSessionsByIdControl**](AgentAPI.md#GetAgentSessionsByIdControl) | **Get** /v1/agent/sessions/{id}/control | Returns the steering commands (pause/resume/stop/message) recorded against the caller&#39;s own session that are newer than the cursor, oldest first, with the cursor to poll from next.
[**GetAgentSessionsByIdProgress**](AgentAPI.md#GetAgentSessionsByIdProgress) | **Get** /v1/agent/sessions/{id}/progress | Returns how far along one run is: the share of its goal that is done, whether it is running, blocked or finished, and a line saying what it is doing right now.
[**GetAgentSessionsByIdTree**](AgentAPI.md#GetAgentSessionsByIdTree) | **Get** /v1/agent/sessions/{id}/tree | Returns the subagent-flow graph rooted at this session: the session, its children, their children, each node carrying its own event count.
[**GetAgentSessionsStream**](AgentAPI.md#GetAgentSessionsStream) | **Get** /v1/agent/sessions/stream | Live session and event updates for the caller&#39;s org, as Server-Sent Events.
[**GetAgentTargets**](AgentAPI.md#GetAgentTargets) | **Get** /v1/agent/targets | Returns every machine registered to the caller&#39;s org, newest first, each with its live session load.
[**GetAgentTargetsById**](AgentAPI.md#GetAgentTargetsById) | **Get** /v1/agent/targets/{id} | Returns one registered machine, with its live session load.
[**PatchAgentByRef**](AgentAPI.md#PatchAgentByRef) | **Patch** /v1/agent/{ref} | Changes an agent in place.
[**PatchAgentSessionsById**](AgentAPI.md#PatchAgentSessionsById) | **Patch** /v1/agent/sessions/{id} | Updates a session&#39;s surface-owned truth: its status, its title, the run-target it is dispatched to, and the product it built plus whether that build&#39;s story is public.
[**PatchAgentTargetsById**](AgentAPI.md#PatchAgentTargetsById) | **Patch** /v1/agent/targets/{id} | Updates one machine in place.
[**PostAgent**](AgentAPI.md#PostAgent) | **Post** /v1/agent | Defines an agent in the caller&#39;s org: a model, a system prompt (instructions) and a set of tool names.
[**PostAgentAsk**](AgentAPI.md#PostAgentAsk) | **Post** /v1/agent/ask | The MCP server a coding run&#39;s harness asks its person through.
[**PostAgentByRefRun**](AgentAPI.md#PostAgentByRefRun) | **Post** /v1/agent/{ref}/run | Run one of your org&#39;s agents and get the recorded run back.
[**PostAgentChat**](AgentAPI.md#PostAgentChat) | **Post** /v1/agent/chat | Run one tool-calling round against your org&#39;s own tools
[**PostAgentChatConversations**](AgentAPI.md#PostAgentChatConversations) | **Post** /v1/agent/chat/conversations | Record turns in a conversation
[**PostAgentChatConversationsByIdShares**](AgentAPI.md#PostAgentChatConversationsByIdShares) | **Post** /v1/agent/chat/conversations/{id}/shares | Share one of your conversations by link
[**PostAgentChatSharesRead**](AgentAPI.md#PostAgentChatSharesRead) | **Post** /v1/agent/chat/shares/read | Open a conversation shared by link
[**PostAgentCoding**](AgentAPI.md#PostAgentCoding) | **Post** /v1/agent/coding | Start one autonomous coding run against a repo in the caller&#39;s org
[**PostAgentCodingBySessionMerge**](AgentAPI.md#PostAgentCodingBySessionMerge) | **Post** /v1/agent/coding/{session}/merge | Merges a coding run&#39;s pull request into the branch it proposes into, on the forge the run pushed to, and answers the pull request after.
[**PostAgentMcpByServer**](AgentAPI.md#PostAgentMcpByServer) | **Post** /v1/agent/mcp/{server} | The MCP address a coding run&#39;s harness reaches one of its org&#39;s MCP servers through.
[**PostAgentSessions**](AgentAPI.md#PostAgentSessions) | **Post** /v1/agent/sessions | Opens a live agent session in the caller&#39;s org — the row every surface (the CLI&#39;s outer agent, hanzo.bot, the console, chat) hangs its activity off.
[**PostAgentSessionsByIdBudget**](AgentAPI.md#PostAgentSessionsByIdBudget) | **Post** /v1/agent/sessions/{id}/budget | Sets, raises, or removes a session&#39;s cap.
[**PostAgentSessionsByIdEvents**](AgentAPI.md#PostAgentSessionsByIdEvents) | **Post** /v1/agent/sessions/{id}/events | Records one turn of a session&#39;s transcript and answers 201 with it.
[**PostAgentSessionsByIdMessage**](AgentAPI.md#PostAgentSessionsByIdMessage) | **Post** /v1/agent/sessions/{id}/message | Sends a steering message to a running session — the endpoint a human or another agent interrupts through.
[**PostAgentSessionsByIdPause**](AgentAPI.md#PostAgentSessionsByIdPause) | **Post** /v1/agent/sessions/{id}/pause | Asks a running session to pause.
[**PostAgentSessionsByIdResume**](AgentAPI.md#PostAgentSessionsByIdResume) | **Post** /v1/agent/sessions/{id}/resume | Asks a paused session to continue, on the same terms as a pause.
[**PostAgentSessionsByIdStop**](AgentAPI.md#PostAgentSessionsByIdStop) | **Post** /v1/agent/sessions/{id}/stop | Ends a running session.
[**PostAgentTargets**](AgentAPI.md#PostAgentTargets) | **Post** /v1/agent/targets | Registers a machine as an agent target, or re-links one that is already registered.
[**PostAgentTargetsByIdClaim**](AgentAPI.md#PostAgentTargetsByIdClaim) | **Post** /v1/agent/targets/{id}/claim | ClaimRoutedRun is the machine&#39;s long poll for work: it authenticates the daemon, stamps the liveness the dispatch gate reads (the poll IS the proof a runner is listening), and waits up to 25 seconds for the next run addressed to THIS machine.
[**PostAgentTargetsByIdKey**](AgentAPI.md#PostAgentTargetsByIdKey) | **Post** /v1/agent/targets/{id}/key | Mints (or rotates) the claim key a &#x60;hanzo code --serve&#x60; daemon presents to claim work for this machine, and returns it ONCE: only its SHA-256 hash is stored.
[**PostAgentTargetsByIdRunsByRunidReport**](AgentAPI.md#PostAgentTargetsByIdRunsByRunidReport) | **Post** /v1/agent/targets/{id}/runs/{runId}/report | Completes a claimed run: it delivers the terminal result to the run&#39;s durable owner, which is what lets that workflow finish.



## DeleteAgentByRef

> DeleteAgentByRef(ctx, ref).Execute()

Removes an agent and every run recorded against it.



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
	ref := "helper" // string | Ref is the agent's public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.DeleteAgentByRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgentByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentByRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteAgentChatConversationsByIdSharesByShare

> DeleteAgentChatConversationsByIdSharesByShare(ctx, id, share).Execute()

Revoke a link to one of your conversations



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
	share := "share_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.DeleteAgentChatConversationsByIdSharesByShare(context.Background(), id, share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgentChatConversationsByIdSharesByShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**share** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentChatConversationsByIdSharesByShareRequest struct via the builder pattern


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


## DeleteAgentChatConversationsByIdSharesByShareViewersByViewer

> DeleteAgentChatConversationsByIdSharesByShareViewersByViewer(ctx, id, share, viewer).Execute()

Remove one viewer from a link



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
	share := "share_example" // string | 
	viewer := "viewer_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.DeleteAgentChatConversationsByIdSharesByShareViewersByViewer(context.Background(), id, share, viewer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgentChatConversationsByIdSharesByShareViewersByViewer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**share** | **string** |  | 
**viewer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentChatConversationsByIdSharesByShareViewersByViewerRequest struct via the builder pattern


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


## DeleteAgentChatSharesByShare

> DeleteAgentChatSharesByShare(ctx, share).Execute()

Revoke any link in your organization



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
	share := "share_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.DeleteAgentChatSharesByShare(context.Background(), share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgentChatSharesByShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**share** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentChatSharesByShareRequest struct via the builder pattern


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


## DeleteAgentTargetsById

> AgentTargetDeleted DeleteAgentTargetsById(ctx, id).Execute()

Deregisters one machine.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.DeleteAgentTargetsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgentTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgentTargetsById`: AgentTargetDeleted
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.DeleteAgentTargetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentTargetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentTargetDeleted**](AgentTargetDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> AgentAgentList GetAgent(ctx).Execute()

Returns every agent defined in the caller's org, each with the number of runs recorded against it.



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
	resp, r, err := apiClient.AgentAPI.GetAgent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgent`: AgentAgentList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRequest struct via the builder pattern


### Return type

[**AgentAgentList**](AgentAgentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentActivity

> AgentActivityFeed GetAgentActivity(ctx).Execute()

Serves the org-wide recent-activity feed.



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
	resp, r, err := apiClient.AgentAPI.GetAgentActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentActivity`: AgentActivityFeed
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentActivityRequest struct via the builder pattern


### Return type

[**AgentActivityFeed**](AgentActivityFeed.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBuilds

> AgentBuildList GetAgentBuilds(ctx).Limit(limit).Execute()

Returns the public index of every published build, most recently updated first, so a gallery can link straight to the story behind each product.



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
	limit := int64(789) // int64 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentBuilds(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentBuilds`: AgentBuildList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the page. Absent, zero or over 500 reads as 100. | 

### Return type

[**AgentBuildList**](AgentBuildList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBuildsByOrgByProject

> AgentBuildView GetAgentBuildsByOrgByProject(ctx, org, project).Execute()

Returns the readable build of one product: the agent session that produced it, turn by turn — the prompts, the reasoning, the commits each turn produced — plus the exact `git log` that re-derives every commit binding from git itself, so nothing here has to be taken on trust.



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
	org := "hanzo" // string | Org is the org that published the build, from the path.
	project := "landing" // string | Project is the product's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentBuildsByOrgByProject(context.Background(), org, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentBuildsByOrgByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentBuildsByOrgByProject`: AgentBuildView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentBuildsByOrgByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the org that published the build, from the path. | 
**project** | **string** | Project is the product&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBuildsByOrgByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentBuildView**](AgentBuildView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentByRef

> AgentAgentDetail GetAgentByRef(ctx, ref).Execute()

Returns one agent with its system prompt and its 20 most recent runs.



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
	ref := "helper" // string | Ref is the agent's public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentByRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentByRef`: AgentAgentDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentByRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentByRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentAgentDetail**](AgentAgentDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentByRefRuns

> AgentRunList GetAgentByRefRuns(ctx, ref).Limit(limit).Execute()

Returns one agent's execution history, newest first — each run's input, its output or its error, and how long it took.



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
	ref := "helper" // string | Ref is the agent's public id or its org-unique name, from the path.
	limit := int64(20) // int64 | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentByRefRuns(context.Background(), ref).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentByRefRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentByRefRuns`: AgentRunList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentByRefRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id or its org-unique name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentByRefRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. | 

### Return type

[**AgentRunList**](AgentRunList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentByRefSpend

> AgentSpendView GetAgentByRefSpend(ctx, ref).By(by).Execute()

Answers what one of your org's agents has spent, in integer micro-USD.



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
	ref := "ref_example" // string | Ref is the agent's public id or its org-unique name.
	by := "by_example" // string | By groups the answer: \"component\" is the only grouping today. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentByRefSpend(context.Background(), ref).By(by).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentByRefSpend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentByRefSpend`: AgentSpendView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentByRefSpend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id or its org-unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentByRefSpendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **by** | **string** | By groups the answer: \&quot;component\&quot; is the only grouping today. | 

### Return type

[**AgentSpendView**](AgentSpendView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentChatConversations

> GetAgentChatConversations(ctx).Execute()

List the agent threads in your org



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
	r, err := apiClient.AgentAPI.GetAgentChatConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatConversationsRequest struct via the builder pattern


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


## GetAgentChatConversationsById

> GetAgentChatConversationsById(ctx, id).Execute()

Read one agent thread in full



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
	r, err := apiClient.AgentAPI.GetAgentChatConversationsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatConversationsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAgentChatConversationsByIdRequest struct via the builder pattern


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


## GetAgentChatConversationsByIdShares

> GetAgentChatConversationsByIdShares(ctx, id).Execute()

List the live links to one of your conversations



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
	r, err := apiClient.AgentAPI.GetAgentChatConversationsByIdShares(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatConversationsByIdShares``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAgentChatConversationsByIdSharesRequest struct via the builder pattern


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


## GetAgentChatPresets

> GetAgentChatPresets(ctx).Execute()

List the agent presets available to a caller



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
	r, err := apiClient.AgentAPI.GetAgentChatPresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatPresetsRequest struct via the builder pattern


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


## GetAgentChatShared

> GetAgentChatShared(ctx).Execute()

List the chats shared with you



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
	r, err := apiClient.AgentAPI.GetAgentChatShared(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatSharedRequest struct via the builder pattern


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


## GetAgentChatSharedByShare

> GetAgentChatSharedByShare(ctx, share).Execute()

Read a chat shared with you



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
	share := "share_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.GetAgentChatSharedByShare(context.Background(), share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatSharedByShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**share** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatSharedByShareRequest struct via the builder pattern


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


## GetAgentChatShares

> GetAgentChatShares(ctx).Execute()

List every live link in your organization



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
	r, err := apiClient.AgentAPI.GetAgentChatShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentChatShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatSharesRequest struct via the builder pattern


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


## GetAgentCodingBySessionArtifacts

> AgentCodingArtifacts GetAgentCodingBySessionArtifacts(ctx, session).Execute()

Lists what a coding run left, kept after its sandbox is gone: every file it added or changed and its whole change as `changes.patch`, stored beside the run; the ports it served, each a preview while its sandbox is kept; its pull request and where its work was published.



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
	session := "sess_0123456789abcdef0123456789abcdef" // string | Session is the run's handle — the sessionId POST /v1/agent/coding answered with — from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentCodingBySessionArtifacts(context.Background(), session).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentCodingBySessionArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentCodingBySessionArtifacts`: AgentCodingArtifacts
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentCodingBySessionArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**session** | **string** | Session is the run&#39;s handle — the sessionId POST /v1/agent/coding answered with — from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentCodingBySessionArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentCodingArtifacts**](AgentCodingArtifacts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentCodingBySessionBlob

> AgentCodingBlob GetAgentCodingBySessionBlob(ctx, session).Path(path).Execute()

Returns one file of a coding run's repository, read where the tree is read, in the shape GET /v1/git/repos/{name}/blob answers: text verbatim, anything else base64, and a file past the 1 MiB view cap marked truncated with no content.



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
	session := "sess_0123456789abcdef0123456789abcdef" // string | Session is the run's handle, from the path.
	path := "api/server.go" // string | Path is repo-relative, from the query. Empty is the repository's root. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentCodingBySessionBlob(context.Background(), session).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentCodingBySessionBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentCodingBySessionBlob`: AgentCodingBlob
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentCodingBySessionBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**session** | **string** | Session is the run&#39;s handle, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentCodingBySessionBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path is repo-relative, from the query. Empty is the repository&#39;s root. | 

### Return type

[**AgentCodingBlob**](AgentCodingBlob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentCodingBySessionChanges

> AgentCodingChanges GetAgentCodingBySessionChanges(ctx, session).Execute()

Returns what a coding run changed, read from the forge it pushed its branch to: the commits on its branch that the base does not have, newest first; the net change of the branch against its base, one entry per file with that file's patch; and its pull request with the reviews it has had, or null while it has none.



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
	session := "sess_0123456789abcdef0123456789abcdef" // string | Session is the run's handle — the sessionId POST /v1/agent/coding answered with — from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentCodingBySessionChanges(context.Background(), session).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentCodingBySessionChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentCodingBySessionChanges`: AgentCodingChanges
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentCodingBySessionChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**session** | **string** | Session is the run&#39;s handle — the sessionId POST /v1/agent/coding answered with — from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentCodingBySessionChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentCodingChanges**](AgentCodingChanges.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentCodingBySessionTree

> AgentCodingTree GetAgentCodingBySessionTree(ctx, session).Path(path).Execute()

Lists one directory of a coding run's repository, one level down with directories first: at the run's own branch once the forge holds it, and at the branch it started from until then — `ref` says which.



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
	session := "sess_0123456789abcdef0123456789abcdef" // string | Session is the run's handle, from the path.
	path := "api" // string | Path is repo-relative, from the query. Empty is the repository's root. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentCodingBySessionTree(context.Background(), session).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentCodingBySessionTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentCodingBySessionTree`: AgentCodingTree
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentCodingBySessionTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**session** | **string** | Session is the run&#39;s handle, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentCodingBySessionTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path is repo-relative, from the query. Empty is the repository&#39;s root. | 

### Return type

[**AgentCodingTree**](AgentCodingTree.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentMetrics

> AgentMetricsView GetAgentMetrics(ctx).Range_(range_).Execute()

Serves the invocations-over-time histogram for the org's Agents dashboard.



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
	range_ := "7D" // string | Range is the window to bucket: 24H, 7D or 30D. Anything else reads as 30D. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentMetrics`: AgentMetricsView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the window to bucket: 24H, 7D or 30D. Anything else reads as 30D. | 

### Return type

[**AgentMetricsView**](AgentMetricsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentRuns

> AgentRunList GetAgentRuns(ctx).Limit(limit).Status(status).Execute()

Returns the org's agent runs across EVERY agent, newest first — what ran here, for whom, on which model, how long it took, and why it failed.



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
	limit := int64(20) // int64 | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. (optional)
	status := "error" // string | Status keeps only runs with this outcome (\"ok\" or \"error\"). Empty keeps both. It is the filter an operator reaches for first — \"show me what broke\" — and answering it here rather than by paging the whole history client-side is the difference between a usable feed and a download. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentRuns(context.Background()).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentRuns`: AgentRunList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. | 
 **status** | **string** | Status keeps only runs with this outcome (\&quot;ok\&quot; or \&quot;error\&quot;). Empty keeps both. It is the filter an operator reaches for first — \&quot;show me what broke\&quot; — and answering it here rather than by paging the whole history client-side is the difference between a usable feed and a download. | 

### Return type

[**AgentRunList**](AgentRunList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessions

> AgentSessionList GetAgentSessions(ctx).Root(root).Parent(parent).Status(status).Project(project).Room(room).Kind(kind).Limit(limit).After(after).Execute()

Returns the caller org's live sessions, newest first — each with its event count, its direct-child count and a one-line preview of its latest event.



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
	root := "root_example" // string | Root scopes the page to one subagent tree (its root session id). (optional)
	parent := "parent_example" // string | Parent scopes the page to the direct children of one session. Ignored when root is set; with neither, only ROOT sessions come back. (optional)
	status := "running" // string | Status filters to running, paused, done or error. (optional)
	project := "project_example" // string | Project filters to the sessions tagged with one product slug. (optional)
	room := "room_example" // string | Room filters to the sessions started in one collaborative room — the query a space view runs to show what has been run in it. (optional)
	kind := "kind_example" // string | Kind filters to the sessions of one kind of run: \"coding\" lists coding runs, each carrying its repo, base, branch, environment and pull request. (optional)
	limit := int64(20) // int64 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)
	after := "after_example" // string | After is the `next` of the previous page. Absent starts at the newest. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentSessions(context.Background()).Root(root).Parent(parent).Status(status).Project(project).Room(room).Kind(kind).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentSessions`: AgentSessionList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **root** | **string** | Root scopes the page to one subagent tree (its root session id). | 
 **parent** | **string** | Parent scopes the page to the direct children of one session. Ignored when root is set; with neither, only ROOT sessions come back. | 
 **status** | **string** | Status filters to running, paused, done or error. | 
 **project** | **string** | Project filters to the sessions tagged with one product slug. | 
 **room** | **string** | Room filters to the sessions started in one collaborative room — the query a space view runs to show what has been run in it. | 
 **kind** | **string** | Kind filters to the sessions of one kind of run: \&quot;coding\&quot; lists coding runs, each carrying its repo, base, branch, environment and pull request. | 
 **limit** | **int64** | Limit caps the page. Absent, zero or over 500 reads as 100. | 
 **after** | **string** | After is the &#x60;next&#x60; of the previous page. Absent starts at the newest. | 

### Return type

[**AgentSessionList**](AgentSessionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsById

> AgentSessionDetail GetAgentSessionsById(ctx, id).Execute()

Returns one session with its direct child sessions and its 50 most recent events, oldest of those first.



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
	id := "sess_1" // string | ID is the session to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentSessionsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentSessionsById`: AgentSessionDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentSessionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSessionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentSessionDetail**](AgentSessionDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsByIdControl

> AgentControlDrain GetAgentSessionsByIdControl(ctx, id).After(after).Execute()

Returns the steering commands (pause/resume/stop/message) recorded against the caller's own session that are newer than the cursor, oldest first, with the cursor to poll from next.



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
	id := "sess_1" // string | ID is the session whose commands are being drained, from the path.
	after := int64(12) // int64 | After is the last seq this poller applied; only commands newer than it come back. Absent or negative reads as 0, which drains from the beginning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentSessionsByIdControl(context.Background(), id).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessionsByIdControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentSessionsByIdControl`: AgentControlDrain
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentSessionsByIdControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session whose commands are being drained, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSessionsByIdControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int64** | After is the last seq this poller applied; only commands newer than it come back. Absent or negative reads as 0, which drains from the beginning. | 

### Return type

[**AgentControlDrain**](AgentControlDrain.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsByIdProgress

> AgentSessionProgress GetAgentSessionsByIdProgress(ctx, id).Execute()

Returns how far along one run is: the share of its goal that is done, whether it is running, blocked or finished, and a line saying what it is doing right now.



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
	id := "sess_1" // string | ID is the session to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentSessionsByIdProgress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessionsByIdProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentSessionsByIdProgress`: AgentSessionProgress
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentSessionsByIdProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSessionsByIdProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentSessionProgress**](AgentSessionProgress.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsByIdTree

> AgentTreeNode GetAgentSessionsByIdTree(ctx, id).Execute()

Returns the subagent-flow graph rooted at this session: the session, its children, their children, each node carrying its own event count.



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
	id := "sess_1" // string | ID is the session to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentSessionsByIdTree(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessionsByIdTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentSessionsByIdTree`: AgentTreeNode
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentSessionsByIdTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSessionsByIdTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentTreeNode**](AgentTreeNode.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsStream

> GetAgentSessionsStream(ctx).Execute()

Live session and event updates for the caller's org, as Server-Sent Events.



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
	r, err := apiClient.AgentAPI.GetAgentSessionsStream(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessionsStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSessionsStreamRequest struct via the builder pattern


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


## GetAgentTargets

> AgentTargetList GetAgentTargets(ctx).Execute()

Returns every machine registered to the caller's org, newest first, each with its live session load.



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
	resp, r, err := apiClient.AgentAPI.GetAgentTargets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentTargets`: AgentTargetList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentTargets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentTargetsRequest struct via the builder pattern


### Return type

[**AgentTargetList**](AgentTargetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentTargetsById

> AgentTargetView GetAgentTargetsById(ctx, id).Execute()

Returns one registered machine, with its live session load.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentTargetsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentTargetsById`: AgentTargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentTargetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentTargetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentTargetView**](AgentTargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentByRef

> AgentAgentView PatchAgentByRef(ctx, ref).AgentUpdateAgentIn(agentUpdateAgentIn).Execute()

Changes an agent in place.



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
	ref := "helper" // string | Ref is the agent to update — its public id or org-unique name, from the path.
	agentUpdateAgentIn := *openapiclient.NewAgentUpdateAgentIn() // AgentUpdateAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PatchAgentByRef(context.Background(), ref).AgentUpdateAgentIn(agentUpdateAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PatchAgentByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentByRef`: AgentAgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PatchAgentByRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent to update — its public id or org-unique name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentByRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentUpdateAgentIn** | [**AgentUpdateAgentIn**](AgentUpdateAgentIn.md) |  | 

### Return type

[**AgentAgentView**](AgentAgentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentSessionsById

> AgentSessionView PatchAgentSessionsById(ctx, id).AgentPatchSessionIn(agentPatchSessionIn).Execute()

Updates a session's surface-owned truth: its status, its title, the run-target it is dispatched to, and the product it built plus whether that build's story is public.



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
	id := "sess_1" // string | ID is the session to update, from the path.
	agentPatchSessionIn := *openapiclient.NewAgentPatchSessionIn() // AgentPatchSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PatchAgentSessionsById(context.Background(), id).AgentPatchSessionIn(agentPatchSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PatchAgentSessionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentSessionsById`: AgentSessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PatchAgentSessionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentSessionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPatchSessionIn** | [**AgentPatchSessionIn**](AgentPatchSessionIn.md) |  | 

### Return type

[**AgentSessionView**](AgentSessionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentTargetsById

> AgentTargetView PatchAgentTargetsById(ctx, id).AgentPatchTargetIn(agentPatchTargetIn).Execute()

Updates one machine in place.



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
	id := "tgt_1" // string | ID is the target to update, from the path.
	agentPatchTargetIn := *openapiclient.NewAgentPatchTargetIn() // AgentPatchTargetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PatchAgentTargetsById(context.Background(), id).AgentPatchTargetIn(agentPatchTargetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PatchAgentTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentTargetsById`: AgentTargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PatchAgentTargetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentTargetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPatchTargetIn** | [**AgentPatchTargetIn**](AgentPatchTargetIn.md) |  | 

### Return type

[**AgentTargetView**](AgentTargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgent

> AgentAgentView PostAgent(ctx).AgentCreateAgentIn(agentCreateAgentIn).Execute()

Defines an agent in the caller's org: a model, a system prompt (instructions) and a set of tool names.



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
	agentCreateAgentIn := *openapiclient.NewAgentCreateAgentIn() // AgentCreateAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgent(context.Background()).AgentCreateAgentIn(agentCreateAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgent`: AgentAgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentCreateAgentIn** | [**AgentCreateAgentIn**](AgentCreateAgentIn.md) |  | 

### Return type

[**AgentAgentView**](AgentAgentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentAsk

> PostAgentAsk(ctx).Execute()

The MCP server a coding run's harness asks its person through.



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
	r, err := apiClient.AgentAPI.PostAgentAsk(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentAskRequest struct via the builder pattern


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


## PostAgentByRefRun

> PostAgentByRefRun(ctx, ref).Execute()

Run one of your org's agents and get the recorded run back.



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
	ref := "ref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.PostAgentByRefRun(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentByRefRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentByRefRunRequest struct via the builder pattern


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


## PostAgentChat

> PostAgentChat(ctx).Execute()

Run one tool-calling round against your org's own tools



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
	r, err := apiClient.AgentAPI.PostAgentChat(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentChatRequest struct via the builder pattern


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


## PostAgentChatConversations

> PostAgentChatConversations(ctx).Execute()

Record turns in a conversation



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
	r, err := apiClient.AgentAPI.PostAgentChatConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentChatConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentChatConversationsRequest struct via the builder pattern


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


## PostAgentChatConversationsByIdShares

> PostAgentChatConversationsByIdShares(ctx, id).Execute()

Share one of your conversations by link



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
	r, err := apiClient.AgentAPI.PostAgentChatConversationsByIdShares(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentChatConversationsByIdShares``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAgentChatConversationsByIdSharesRequest struct via the builder pattern


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


## PostAgentChatSharesRead

> PostAgentChatSharesRead(ctx).Execute()

Open a conversation shared by link



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
	r, err := apiClient.AgentAPI.PostAgentChatSharesRead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentChatSharesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentChatSharesReadRequest struct via the builder pattern


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


## PostAgentCoding

> AgentCodingStarted PostAgentCoding(ctx).AgentCodingStartIn(agentCodingStartIn).Execute()

Start one autonomous coding run against a repo in the caller's org



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
	agentCodingStartIn := *openapiclient.NewAgentCodingStartIn() // AgentCodingStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentCoding(context.Background()).AgentCodingStartIn(agentCodingStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentCoding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentCoding`: AgentCodingStarted
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentCoding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentCodingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentCodingStartIn** | [**AgentCodingStartIn**](AgentCodingStartIn.md) |  | 

### Return type

[**AgentCodingStarted**](AgentCodingStarted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentCodingBySessionMerge

> AgentCodingMerged PostAgentCodingBySessionMerge(ctx, session).Execute()

Merges a coding run's pull request into the branch it proposes into, on the forge the run pushed to, and answers the pull request after.



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
	session := "sess_0123456789abcdef0123456789abcdef" // string | Session is the run's handle — the sessionId POST /v1/agent/coding answered with — from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentCodingBySessionMerge(context.Background(), session).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentCodingBySessionMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentCodingBySessionMerge`: AgentCodingMerged
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentCodingBySessionMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**session** | **string** | Session is the run&#39;s handle — the sessionId POST /v1/agent/coding answered with — from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentCodingBySessionMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentCodingMerged**](AgentCodingMerged.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentMcpByServer

> PostAgentMcpByServer(ctx, server).Execute()

The MCP address a coding run's harness reaches one of its org's MCP servers through.



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
	server := "server_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.PostAgentMcpByServer(context.Background(), server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentMcpByServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**server** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentMcpByServerRequest struct via the builder pattern


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


## PostAgentSessions

> AgentSessionView PostAgentSessions(ctx).AgentRegisterReq(agentRegisterReq).Execute()

Opens a live agent session in the caller's org — the row every surface (the CLI's outer agent, hanzo.bot, the console, chat) hangs its activity off.



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
	agentRegisterReq := *openapiclient.NewAgentRegisterReq() // AgentRegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessions(context.Background()).AgentRegisterReq(agentRegisterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessions`: AgentSessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentRegisterReq** | [**AgentRegisterReq**](AgentRegisterReq.md) |  | 

### Return type

[**AgentSessionView**](AgentSessionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdBudget

> AgentSessionBudgetView PostAgentSessionsByIdBudget(ctx, id).AgentSessionBudgetIn(agentSessionBudgetIn).Execute()

Sets, raises, or removes a session's cap.



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
	id := "id_example" // string | ID is the session, from the path.
	agentSessionBudgetIn := *openapiclient.NewAgentSessionBudgetIn() // AgentSessionBudgetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdBudget(context.Background(), id).AgentSessionBudgetIn(agentSessionBudgetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdBudget`: AgentSessionBudgetView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessionsByIdBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsByIdBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentSessionBudgetIn** | [**AgentSessionBudgetIn**](AgentSessionBudgetIn.md) |  | 

### Return type

[**AgentSessionBudgetView**](AgentSessionBudgetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdEvents

> AgentEventView PostAgentSessionsByIdEvents(ctx, id).AgentEventIn(agentEventIn).Execute()

Records one turn of a session's transcript and answers 201 with it.



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
	id := "id_example" // string | ID is the session to append to, from the path.
	agentEventIn := *openapiclient.NewAgentEventIn() // AgentEventIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdEvents(context.Background(), id).AgentEventIn(agentEventIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdEvents`: AgentEventView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessionsByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to append to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentEventIn** | [**AgentEventIn**](AgentEventIn.md) |  | 

### Return type

[**AgentEventView**](AgentEventView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdMessage

> AgentControlResult PostAgentSessionsByIdMessage(ctx, id).AgentControlIn(agentControlIn).Execute()

Sends a steering message to a running session — the endpoint a human or another agent interrupts through.



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
	id := "id_example" // string | ID is the session to steer, from the path.
	agentControlIn := *openapiclient.NewAgentControlIn() // AgentControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdMessage(context.Background(), id).AgentControlIn(agentControlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdMessage`: AgentControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessionsByIdMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsByIdMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentControlIn** | [**AgentControlIn**](AgentControlIn.md) |  | 

### Return type

[**AgentControlResult**](AgentControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdPause

> AgentControlResult PostAgentSessionsByIdPause(ctx, id).AgentControlIn(agentControlIn).Execute()

Asks a running session to pause.



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
	id := "id_example" // string | ID is the session to steer, from the path.
	agentControlIn := *openapiclient.NewAgentControlIn() // AgentControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdPause(context.Background(), id).AgentControlIn(agentControlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdPause`: AgentControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessionsByIdPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsByIdPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentControlIn** | [**AgentControlIn**](AgentControlIn.md) |  | 

### Return type

[**AgentControlResult**](AgentControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdResume

> AgentControlResult PostAgentSessionsByIdResume(ctx, id).AgentControlIn(agentControlIn).Execute()

Asks a paused session to continue, on the same terms as a pause.



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
	id := "id_example" // string | ID is the session to steer, from the path.
	agentControlIn := *openapiclient.NewAgentControlIn() // AgentControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdResume(context.Background(), id).AgentControlIn(agentControlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdResume`: AgentControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessionsByIdResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsByIdResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentControlIn** | [**AgentControlIn**](AgentControlIn.md) |  | 

### Return type

[**AgentControlResult**](AgentControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdStop

> AgentControlResult PostAgentSessionsByIdStop(ctx, id).AgentControlIn(agentControlIn).Execute()

Ends a running session.



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
	id := "id_example" // string | ID is the session to steer, from the path.
	agentControlIn := *openapiclient.NewAgentControlIn() // AgentControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdStop(context.Background(), id).AgentControlIn(agentControlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdStop`: AgentControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessionsByIdStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsByIdStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentControlIn** | [**AgentControlIn**](AgentControlIn.md) |  | 

### Return type

[**AgentControlResult**](AgentControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargets

> AgentTargetView PostAgentTargets(ctx).AgentTargetReq(agentTargetReq).Execute()

Registers a machine as an agent target, or re-links one that is already registered.



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
	agentTargetReq := *openapiclient.NewAgentTargetReq() // AgentTargetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentTargets(context.Background()).AgentTargetReq(agentTargetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTargets`: AgentTargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentTargetReq** | [**AgentTargetReq**](AgentTargetReq.md) |  | 

### Return type

[**AgentTargetView**](AgentTargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargetsByIdClaim

> AgentRoutedRunOut PostAgentTargetsByIdClaim(ctx, id).Execute()

ClaimRoutedRun is the machine's long poll for work: it authenticates the daemon, stamps the liveness the dispatch gate reads (the poll IS the proof a runner is listening), and waits up to 25 seconds for the next run addressed to THIS machine.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentTargetsByIdClaim(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentTargetsByIdClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTargetsByIdClaim`: AgentRoutedRunOut
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentTargetsByIdClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentTargetsByIdClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentRoutedRunOut**](AgentRoutedRunOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargetsByIdKey

> AgentClaimKeyOut PostAgentTargetsByIdKey(ctx, id).Execute()

Mints (or rotates) the claim key a `hanzo code --serve` daemon presents to claim work for this machine, and returns it ONCE: only its SHA-256 hash is stored.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentTargetsByIdKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentTargetsByIdKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTargetsByIdKey`: AgentClaimKeyOut
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentTargetsByIdKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentTargetsByIdKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentClaimKeyOut**](AgentClaimKeyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargetsByIdRunsByRunidReport

> AgentReportOut PostAgentTargetsByIdRunsByRunidReport(ctx, id, runId).AgentReportRunIn(agentReportRunIn).Execute()

Completes a claimed run: it delivers the terminal result to the run's durable owner, which is what lets that workflow finish.



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
	id := "tgt_1" // string | ID is the machine reporting, from the path.
	runId := "run_1" // string | RunID is the routed run being completed, from the path.
	agentReportRunIn := *openapiclient.NewAgentReportRunIn() // AgentReportRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentTargetsByIdRunsByRunidReport(context.Background(), id, runId).AgentReportRunIn(agentReportRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentTargetsByIdRunsByRunidReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTargetsByIdRunsByRunidReport`: AgentReportOut
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentTargetsByIdRunsByRunidReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine reporting, from the path. | 
**runId** | **string** | RunID is the routed run being completed, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentTargetsByIdRunsByRunidReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentReportRunIn** | [**AgentReportRunIn**](AgentReportRunIn.md) |  | 

### Return type

[**AgentReportOut**](AgentReportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

