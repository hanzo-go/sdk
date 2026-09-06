# \AgentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgentByRef**](AgentAPI.md#DeleteAgentByRef) | **Delete** /v1/agent/{ref} | Removes an agent and every run recorded against it.
[**DeleteAgentTargetsById**](AgentAPI.md#DeleteAgentTargetsById) | **Delete** /v1/agent/targets/{id} | Deregisters one machine.
[**GetAgent**](AgentAPI.md#GetAgent) | **Get** /v1/agent | Returns every agent defined in the caller&#39;s org, each with the number of runs recorded against it.
[**GetAgentActivity**](AgentAPI.md#GetAgentActivity) | **Get** /v1/agent/activity | Serves the org-wide recent-activity feed.
[**GetAgentBuilds**](AgentAPI.md#GetAgentBuilds) | **Get** /v1/agent/builds | Returns the public index of every published build, most recently updated first, so a gallery can link straight to the story behind each product.
[**GetAgentBuildsByOrgByProject**](AgentAPI.md#GetAgentBuildsByOrgByProject) | **Get** /v1/agent/builds/{org}/{project} | Returns the readable build of one product: the agent session that produced it, turn by turn — the prompts, the reasoning, the commits each turn produced — plus the exact &#x60;git log&#x60; that re-derives every commit binding from git itself, so nothing here has to be taken on trust.
[**GetAgentByRef**](AgentAPI.md#GetAgentByRef) | **Get** /v1/agent/{ref} | Returns one agent with its system prompt and its 20 most recent runs.
[**GetAgentByRefRuns**](AgentAPI.md#GetAgentByRefRuns) | **Get** /v1/agent/{ref}/runs | Returns one agent&#39;s execution history, newest first — each run&#39;s input, its output or its error, and how long it took.
[**GetAgentChatConversations**](AgentAPI.md#GetAgentChatConversations) | **Get** /v1/agent/chat/conversations | List the agent threads in your org
[**GetAgentChatConversationsById**](AgentAPI.md#GetAgentChatConversationsById) | **Get** /v1/agent/chat/conversations/{id} | Read one agent thread in full
[**GetAgentChatPresets**](AgentAPI.md#GetAgentChatPresets) | **Get** /v1/agent/chat/presets | List the agent presets available to a caller
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
[**PostAgentByRefRun**](AgentAPI.md#PostAgentByRefRun) | **Post** /v1/agent/{ref}/run | Run one of your org&#39;s agents and get the recorded run back.
[**PostAgentChat**](AgentAPI.md#PostAgentChat) | **Post** /v1/agent/chat | Run one tool-calling round against your org&#39;s own tools
[**PostAgentChatConversations**](AgentAPI.md#PostAgentChatConversations) | **Post** /v1/agent/chat/conversations | Record turns in a conversation
[**PostAgentCoding**](AgentAPI.md#PostAgentCoding) | **Post** /v1/agent/coding | Start one autonomous coding run against a repo in the caller&#39;s org
[**PostAgentSessions**](AgentAPI.md#PostAgentSessions) | **Post** /v1/agent/sessions | Opens a live agent session in the caller&#39;s org — the row every surface (the CLI&#39;s outer agent, hanzo.bot, the console, chat) hangs its activity off.
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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentTargetsById

> TargetDeleted DeleteAgentTargetsById(ctx, id).Execute()

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
	// response from `DeleteAgentTargetsById`: TargetDeleted
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

[**TargetDeleted**](TargetDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> AgentList GetAgent(ctx).Execute()

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
	// response from `GetAgent`: AgentList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRequest struct via the builder pattern


### Return type

[**AgentList**](AgentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentActivity

> ActivityFeed GetAgentActivity(ctx).Execute()

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
	// response from `GetAgentActivity`: ActivityFeed
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentActivityRequest struct via the builder pattern


### Return type

[**ActivityFeed**](ActivityFeed.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBuilds

> BuildList GetAgentBuilds(ctx).Limit(limit).Execute()

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
	// response from `GetAgentBuilds`: BuildList
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

[**BuildList**](BuildList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBuildsByOrgByProject

> BuildView GetAgentBuildsByOrgByProject(ctx, org, project).Execute()

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
	// response from `GetAgentBuildsByOrgByProject`: BuildView
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

[**BuildView**](BuildView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentByRef

> AgentDetail GetAgentByRef(ctx, ref).Execute()

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
	// response from `GetAgentByRef`: AgentDetail
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

[**AgentDetail**](AgentDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentByRefRuns

> RunList GetAgentByRefRuns(ctx, ref).Limit(limit).Execute()

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
	// response from `GetAgentByRefRuns`: RunList
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

[**RunList**](RunList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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


## GetAgentMetrics

> MetricsView GetAgentMetrics(ctx).Range_(range_).Execute()

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
	// response from `GetAgentMetrics`: MetricsView
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

[**MetricsView**](MetricsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentRuns

> RunList GetAgentRuns(ctx).Limit(limit).Status(status).Execute()

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
	// response from `GetAgentRuns`: RunList
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

[**RunList**](RunList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessions

> SessionList GetAgentSessions(ctx).Root(root).Parent(parent).Status(status).Project(project).Room(room).Limit(limit).Execute()

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
	limit := int64(20) // int64 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentSessions(context.Background()).Root(root).Parent(parent).Status(status).Project(project).Room(room).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentSessions`: SessionList
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
 **limit** | **int64** | Limit caps the page. Absent, zero or over 500 reads as 100. | 

### Return type

[**SessionList**](SessionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsById

> SessionDetail GetAgentSessionsById(ctx, id).Execute()

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
	// response from `GetAgentSessionsById`: SessionDetail
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

[**SessionDetail**](SessionDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsByIdControl

> ControlDrain GetAgentSessionsByIdControl(ctx, id).After(after).Execute()

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
	// response from `GetAgentSessionsByIdControl`: ControlDrain
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

[**ControlDrain**](ControlDrain.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsByIdProgress

> SessionProgress GetAgentSessionsByIdProgress(ctx, id).Execute()

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
	// response from `GetAgentSessionsByIdProgress`: SessionProgress
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

[**SessionProgress**](SessionProgress.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSessionsByIdTree

> TreeNode GetAgentSessionsByIdTree(ctx, id).Execute()

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
	// response from `GetAgentSessionsByIdTree`: TreeNode
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

[**TreeNode**](TreeNode.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

> TargetList GetAgentTargets(ctx).Execute()

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
	// response from `GetAgentTargets`: TargetList
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentTargets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentTargetsRequest struct via the builder pattern


### Return type

[**TargetList**](TargetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentTargetsById

> TargetView GetAgentTargetsById(ctx, id).Execute()

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
	// response from `GetAgentTargetsById`: TargetView
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

[**TargetView**](TargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentByRef

> AgentView PatchAgentByRef(ctx, ref).UpdateAgentIn(updateAgentIn).Execute()

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
	updateAgentIn := *openapiclient.NewUpdateAgentIn() // UpdateAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PatchAgentByRef(context.Background(), ref).UpdateAgentIn(updateAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PatchAgentByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentByRef`: AgentView
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

 **updateAgentIn** | [**UpdateAgentIn**](UpdateAgentIn.md) |  | 

### Return type

[**AgentView**](AgentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentSessionsById

> SessionView PatchAgentSessionsById(ctx, id).PatchSessionIn(patchSessionIn).Execute()

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
	patchSessionIn := *openapiclient.NewPatchSessionIn() // PatchSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PatchAgentSessionsById(context.Background(), id).PatchSessionIn(patchSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PatchAgentSessionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentSessionsById`: SessionView
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

 **patchSessionIn** | [**PatchSessionIn**](PatchSessionIn.md) |  | 

### Return type

[**SessionView**](SessionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentTargetsById

> TargetView PatchAgentTargetsById(ctx, id).PatchTargetIn(patchTargetIn).Execute()

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
	patchTargetIn := *openapiclient.NewPatchTargetIn() // PatchTargetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PatchAgentTargetsById(context.Background(), id).PatchTargetIn(patchTargetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PatchAgentTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentTargetsById`: TargetView
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

 **patchTargetIn** | [**PatchTargetIn**](PatchTargetIn.md) |  | 

### Return type

[**TargetView**](TargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgent

> AgentView PostAgent(ctx).CreateAgentIn(createAgentIn).Execute()

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
	createAgentIn := *openapiclient.NewCreateAgentIn() // CreateAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgent(context.Background()).CreateAgentIn(createAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgent`: AgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAgentIn** | [**CreateAgentIn**](CreateAgentIn.md) |  | 

### Return type

[**AgentView**](AgentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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


## PostAgentCoding

> CodingStarted PostAgentCoding(ctx).CodingStartIn(codingStartIn).Execute()

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
	codingStartIn := *openapiclient.NewCodingStartIn() // CodingStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentCoding(context.Background()).CodingStartIn(codingStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentCoding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentCoding`: CodingStarted
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentCoding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentCodingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codingStartIn** | [**CodingStartIn**](CodingStartIn.md) |  | 

### Return type

[**CodingStarted**](CodingStarted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessions

> SessionView PostAgentSessions(ctx).RegisterReq(registerReq).Execute()

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
	registerReq := *openapiclient.NewRegisterReq() // RegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessions(context.Background()).RegisterReq(registerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessions`: SessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerReq** | [**RegisterReq**](RegisterReq.md) |  | 

### Return type

[**SessionView**](SessionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdEvents

> EventView PostAgentSessionsByIdEvents(ctx, id).EventIn(eventIn).Execute()

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
	eventIn := *openapiclient.NewEventIn() // EventIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdEvents(context.Background(), id).EventIn(eventIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdEvents`: EventView
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

 **eventIn** | [**EventIn**](EventIn.md) |  | 

### Return type

[**EventView**](EventView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdMessage

> ControlResult PostAgentSessionsByIdMessage(ctx, id).ControlIn(controlIn).Execute()

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
	controlIn := *openapiclient.NewControlIn() // ControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdMessage(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdMessage`: ControlResult
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

 **controlIn** | [**ControlIn**](ControlIn.md) |  | 

### Return type

[**ControlResult**](ControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdPause

> ControlResult PostAgentSessionsByIdPause(ctx, id).ControlIn(controlIn).Execute()

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
	controlIn := *openapiclient.NewControlIn() // ControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdPause(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdPause`: ControlResult
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

 **controlIn** | [**ControlIn**](ControlIn.md) |  | 

### Return type

[**ControlResult**](ControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdResume

> ControlResult PostAgentSessionsByIdResume(ctx, id).ControlIn(controlIn).Execute()

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
	controlIn := *openapiclient.NewControlIn() // ControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdResume(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdResume`: ControlResult
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

 **controlIn** | [**ControlIn**](ControlIn.md) |  | 

### Return type

[**ControlResult**](ControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentSessionsByIdStop

> ControlResult PostAgentSessionsByIdStop(ctx, id).ControlIn(controlIn).Execute()

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
	controlIn := *openapiclient.NewControlIn() // ControlIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentSessionsByIdStop(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentSessionsByIdStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSessionsByIdStop`: ControlResult
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

 **controlIn** | [**ControlIn**](ControlIn.md) |  | 

### Return type

[**ControlResult**](ControlResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargets

> TargetView PostAgentTargets(ctx).TargetReq(targetReq).Execute()

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
	targetReq := *openapiclient.NewTargetReq() // TargetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentTargets(context.Background()).TargetReq(targetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTargets`: TargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.PostAgentTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetReq** | [**TargetReq**](TargetReq.md) |  | 

### Return type

[**TargetView**](TargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargetsByIdClaim

> RoutedRunOut PostAgentTargetsByIdClaim(ctx, id).Execute()

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
	// response from `PostAgentTargetsByIdClaim`: RoutedRunOut
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

[**RoutedRunOut**](RoutedRunOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargetsByIdKey

> ClaimKeyOut PostAgentTargetsByIdKey(ctx, id).Execute()

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
	// response from `PostAgentTargetsByIdKey`: ClaimKeyOut
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

[**ClaimKeyOut**](ClaimKeyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentTargetsByIdRunsByRunidReport

> ReportOut PostAgentTargetsByIdRunsByRunidReport(ctx, id, runId).ReportRunIn(reportRunIn).Execute()

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
	reportRunIn := *openapiclient.NewReportRunIn() // ReportRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.PostAgentTargetsByIdRunsByRunidReport(context.Background(), id, runId).ReportRunIn(reportRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PostAgentTargetsByIdRunsByRunidReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTargetsByIdRunsByRunidReport`: ReportOut
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


 **reportRunIn** | [**ReportRunIn**](ReportRunIn.md) |  | 

### Return type

[**ReportOut**](ReportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

