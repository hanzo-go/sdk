# \AgentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgentsByRef**](AgentsAPI.md#DeleteAgentsByRef) | **Delete** /v1/agents/{ref} | Removes an agent and every run recorded against it.
[**DeleteAgentsTargetsById**](AgentsAPI.md#DeleteAgentsTargetsById) | **Delete** /v1/agents/targets/{id} | Deregisters one machine.
[**GetAgents**](AgentsAPI.md#GetAgents) | **Get** /v1/agents | Returns every agent defined in the caller&#39;s org, each with the number of runs recorded against it.
[**GetAgentsActivity**](AgentsAPI.md#GetAgentsActivity) | **Get** /v1/agents/activity | Serves the org-wide recent-activity feed.
[**GetAgentsBuilds**](AgentsAPI.md#GetAgentsBuilds) | **Get** /v1/agents/builds | Returns the public index of every published build, most recently updated first, so a gallery can link straight to the story behind each product.
[**GetAgentsBuildsByOrgByProject**](AgentsAPI.md#GetAgentsBuildsByOrgByProject) | **Get** /v1/agents/builds/{org}/{project} | Returns the readable build of one product: the agent session that produced it, turn by turn — the prompts, the reasoning, the commits each turn produced — plus the exact &#x60;git log&#x60; that re-derives every commit binding from git itself, so nothing here has to be taken on trust.
[**GetAgentsByRef**](AgentsAPI.md#GetAgentsByRef) | **Get** /v1/agents/{ref} | Returns one agent with its system prompt and its 20 most recent runs.
[**GetAgentsByRefRuns**](AgentsAPI.md#GetAgentsByRefRuns) | **Get** /v1/agents/{ref}/runs | Returns one agent&#39;s execution history, newest first — each run&#39;s input, its output or its error, and how long it took.
[**GetAgentsChatConversations**](AgentsAPI.md#GetAgentsChatConversations) | **Get** /v1/agents/chat/conversations | List the agent threads in your org
[**GetAgentsChatConversationsById**](AgentsAPI.md#GetAgentsChatConversationsById) | **Get** /v1/agents/chat/conversations/{id} | Read one agent thread in full
[**GetAgentsChatPresets**](AgentsAPI.md#GetAgentsChatPresets) | **Get** /v1/agents/chat/presets | List the agent presets available to a caller
[**GetAgentsMetrics**](AgentsAPI.md#GetAgentsMetrics) | **Get** /v1/agents/metrics | Serves the invocations-over-time histogram for the org&#39;s Agents dashboard.
[**GetAgentsRuns**](AgentsAPI.md#GetAgentsRuns) | **Get** /v1/agents/runs | Returns the org&#39;s agent runs across EVERY agent, newest first — what ran here, for whom, on which model, how long it took, and why it failed.
[**GetAgentsSessions**](AgentsAPI.md#GetAgentsSessions) | **Get** /v1/agents/sessions | Returns the caller org&#39;s live sessions, newest first — each with its event count, its direct-child count and a one-line preview of its latest event.
[**GetAgentsSessionsById**](AgentsAPI.md#GetAgentsSessionsById) | **Get** /v1/agents/sessions/{id} | Returns one session with its direct child sessions and its 50 most recent events, oldest of those first.
[**GetAgentsSessionsByIdControl**](AgentsAPI.md#GetAgentsSessionsByIdControl) | **Get** /v1/agents/sessions/{id}/control | Returns the steering commands (pause/resume/stop/message) recorded against the caller&#39;s own session that are newer than the cursor, oldest first, with the cursor to poll from next.
[**GetAgentsSessionsByIdProgress**](AgentsAPI.md#GetAgentsSessionsByIdProgress) | **Get** /v1/agents/sessions/{id}/progress | Returns how far along one run is: the share of its goal that is done, whether it is running, blocked or finished, and a line saying what it is doing right now.
[**GetAgentsSessionsByIdTree**](AgentsAPI.md#GetAgentsSessionsByIdTree) | **Get** /v1/agents/sessions/{id}/tree | Returns the subagent-flow graph rooted at this session: the session, its children, their children, each node carrying its own event count.
[**GetAgentsSessionsStream**](AgentsAPI.md#GetAgentsSessionsStream) | **Get** /v1/agents/sessions/stream | Live session and event updates for the caller&#39;s org, as Server-Sent Events.
[**GetAgentsTargets**](AgentsAPI.md#GetAgentsTargets) | **Get** /v1/agents/targets | Returns every machine registered to the caller&#39;s org, newest first, each with its live session load.
[**GetAgentsTargetsById**](AgentsAPI.md#GetAgentsTargetsById) | **Get** /v1/agents/targets/{id} | Returns one registered machine, with its live session load.
[**PatchAgentsByRef**](AgentsAPI.md#PatchAgentsByRef) | **Patch** /v1/agents/{ref} | Changes an agent in place.
[**PatchAgentsSessionsById**](AgentsAPI.md#PatchAgentsSessionsById) | **Patch** /v1/agents/sessions/{id} | Updates a session&#39;s surface-owned truth: its status, its title, the run-target it is dispatched to, and the product it built plus whether that build&#39;s story is public.
[**PatchAgentsTargetsById**](AgentsAPI.md#PatchAgentsTargetsById) | **Patch** /v1/agents/targets/{id} | Updates one machine in place.
[**PostAgents**](AgentsAPI.md#PostAgents) | **Post** /v1/agents | Defines an agent in the caller&#39;s org: a model, a system prompt (instructions) and a set of tool names.
[**PostAgentsByRefRun**](AgentsAPI.md#PostAgentsByRefRun) | **Post** /v1/agents/{ref}/run | Run one of your org&#39;s agents and get the recorded run back.
[**PostAgentsChat**](AgentsAPI.md#PostAgentsChat) | **Post** /v1/agents/chat | Run one tool-calling round against your org&#39;s own tools
[**PostAgentsChatConversations**](AgentsAPI.md#PostAgentsChatConversations) | **Post** /v1/agents/chat/conversations | Record turns in a conversation
[**PostAgentsCoding**](AgentsAPI.md#PostAgentsCoding) | **Post** /v1/agents/coding | Start one autonomous coding run against a repo in the caller&#39;s org
[**PostAgentsSessions**](AgentsAPI.md#PostAgentsSessions) | **Post** /v1/agents/sessions | Opens a live agent session in the caller&#39;s org — the row every surface (the CLI&#39;s outer agent, hanzo.bot, the console, chat) hangs its activity off.
[**PostAgentsSessionsByIdEvents**](AgentsAPI.md#PostAgentsSessionsByIdEvents) | **Post** /v1/agents/sessions/{id}/events | Records one turn of a session&#39;s transcript and answers 201 with it.
[**PostAgentsSessionsByIdMessage**](AgentsAPI.md#PostAgentsSessionsByIdMessage) | **Post** /v1/agents/sessions/{id}/message | Sends a steering message to a running session — the endpoint a human or another agent interrupts through.
[**PostAgentsSessionsByIdPause**](AgentsAPI.md#PostAgentsSessionsByIdPause) | **Post** /v1/agents/sessions/{id}/pause | Asks a running session to pause.
[**PostAgentsSessionsByIdResume**](AgentsAPI.md#PostAgentsSessionsByIdResume) | **Post** /v1/agents/sessions/{id}/resume | Asks a paused session to continue, on the same terms as a pause.
[**PostAgentsSessionsByIdStop**](AgentsAPI.md#PostAgentsSessionsByIdStop) | **Post** /v1/agents/sessions/{id}/stop | Ends a running session.
[**PostAgentsTargets**](AgentsAPI.md#PostAgentsTargets) | **Post** /v1/agents/targets | Registers a machine as an agent target, or re-links one that is already registered.
[**PostAgentsTargetsByIdClaim**](AgentsAPI.md#PostAgentsTargetsByIdClaim) | **Post** /v1/agents/targets/{id}/claim | ClaimRoutedRun is the machine&#39;s long poll for work: it authenticates the daemon, stamps the liveness the dispatch gate reads (the poll IS the proof a runner is listening), and waits up to 25 seconds for the next run addressed to THIS machine.
[**PostAgentsTargetsByIdKey**](AgentsAPI.md#PostAgentsTargetsByIdKey) | **Post** /v1/agents/targets/{id}/key | Mints (or rotates) the claim key a &#x60;hanzo code --serve&#x60; daemon presents to claim work for this machine, and returns it ONCE: only its SHA-256 hash is stored.
[**PostAgentsTargetsByIdRunsByRunidReport**](AgentsAPI.md#PostAgentsTargetsByIdRunsByRunidReport) | **Post** /v1/agents/targets/{id}/runs/{runId}/report | Completes a claimed run: it delivers the terminal result to the run&#39;s durable owner, which is what lets that workflow finish.



## DeleteAgentsByRef

> DeleteAgentsByRef(ctx, ref).Execute()

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
	r, err := apiClient.AgentsAPI.DeleteAgentsByRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.DeleteAgentsByRef``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAgentsByRefRequest struct via the builder pattern


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


## DeleteAgentsTargetsById

> TargetDeleted DeleteAgentsTargetsById(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.DeleteAgentsTargetsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.DeleteAgentsTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgentsTargetsById`: TargetDeleted
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.DeleteAgentsTargetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentsTargetsByIdRequest struct via the builder pattern


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


## GetAgents

> AgentList GetAgents(ctx).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgents`: AgentList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


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


## GetAgentsActivity

> ActivityFeed GetAgentsActivity(ctx).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsActivity`: ActivityFeed
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsActivityRequest struct via the builder pattern


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


## GetAgentsBuilds

> BuildList GetAgentsBuilds(ctx).Limit(limit).Execute()

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
	limit := int32(56) // int32 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsBuilds(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsBuilds`: BuildList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the page. Absent, zero or over 500 reads as 100. | 

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


## GetAgentsBuildsByOrgByProject

> BuildView GetAgentsBuildsByOrgByProject(ctx, org, project).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsBuildsByOrgByProject(context.Background(), org, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsBuildsByOrgByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsBuildsByOrgByProject`: BuildView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsBuildsByOrgByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the org that published the build, from the path. | 
**project** | **string** | Project is the product&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsBuildsByOrgByProjectRequest struct via the builder pattern


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


## GetAgentsByRef

> AgentDetail GetAgentsByRef(ctx, ref).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsByRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsByRef`: AgentDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsByRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsByRefRequest struct via the builder pattern


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


## GetAgentsByRefRuns

> RunList GetAgentsByRefRuns(ctx, ref).Limit(limit).Execute()

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
	limit := int32(20) // int32 | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsByRefRuns(context.Background(), ref).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsByRefRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsByRefRuns`: RunList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsByRefRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id or its org-unique name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsByRefRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. | 

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


## GetAgentsChatConversations

> GetAgentsChatConversations(ctx).Execute()

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
	r, err := apiClient.AgentsAPI.GetAgentsChatConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsChatConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsChatConversationsRequest struct via the builder pattern


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


## GetAgentsChatConversationsById

> GetAgentsChatConversationsById(ctx, id).Execute()

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
	r, err := apiClient.AgentsAPI.GetAgentsChatConversationsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsChatConversationsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAgentsChatConversationsByIdRequest struct via the builder pattern


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


## GetAgentsChatPresets

> GetAgentsChatPresets(ctx).Execute()

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
	r, err := apiClient.AgentsAPI.GetAgentsChatPresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsChatPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsChatPresetsRequest struct via the builder pattern


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


## GetAgentsMetrics

> MetricsView GetAgentsMetrics(ctx).Range_(range_).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsMetrics`: MetricsView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsMetricsRequest struct via the builder pattern


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


## GetAgentsRuns

> RunList GetAgentsRuns(ctx).Limit(limit).Status(status).Execute()

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
	limit := int32(20) // int32 | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. (optional)
	status := "error" // string | Status keeps only runs with this outcome (\"ok\" or \"error\"). Empty keeps both. It is the filter an operator reaches for first — \"show me what broke\" — and answering it here rather than by paging the whole history client-side is the difference between a usable feed and a download. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsRuns(context.Background()).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsRuns`: RunList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. | 
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


## GetAgentsSessions

> SessionList GetAgentsSessions(ctx).Root(root).Parent(parent).Status(status).Project(project).Room(room).Limit(limit).Execute()

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
	limit := int32(20) // int32 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsSessions(context.Background()).Root(root).Parent(parent).Status(status).Project(project).Room(room).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsSessions`: SessionList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **root** | **string** | Root scopes the page to one subagent tree (its root session id). | 
 **parent** | **string** | Parent scopes the page to the direct children of one session. Ignored when root is set; with neither, only ROOT sessions come back. | 
 **status** | **string** | Status filters to running, paused, done or error. | 
 **project** | **string** | Project filters to the sessions tagged with one product slug. | 
 **room** | **string** | Room filters to the sessions started in one collaborative room — the query a space view runs to show what has been run in it. | 
 **limit** | **int32** | Limit caps the page. Absent, zero or over 500 reads as 100. | 

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


## GetAgentsSessionsById

> SessionDetail GetAgentsSessionsById(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsSessionsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsSessionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsSessionsById`: SessionDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsSessionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsSessionsByIdRequest struct via the builder pattern


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


## GetAgentsSessionsByIdControl

> ControlDrain GetAgentsSessionsByIdControl(ctx, id).After(after).Execute()

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
	after := int32(12) // int32 | After is the last seq this poller applied; only commands newer than it come back. Absent or negative reads as 0, which drains from the beginning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsSessionsByIdControl(context.Background(), id).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsSessionsByIdControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsSessionsByIdControl`: ControlDrain
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsSessionsByIdControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session whose commands are being drained, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsSessionsByIdControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int32** | After is the last seq this poller applied; only commands newer than it come back. Absent or negative reads as 0, which drains from the beginning. | 

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


## GetAgentsSessionsByIdProgress

> SessionProgress GetAgentsSessionsByIdProgress(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsSessionsByIdProgress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsSessionsByIdProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsSessionsByIdProgress`: SessionProgress
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsSessionsByIdProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsSessionsByIdProgressRequest struct via the builder pattern


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


## GetAgentsSessionsByIdTree

> TreeNode GetAgentsSessionsByIdTree(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsSessionsByIdTree(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsSessionsByIdTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsSessionsByIdTree`: TreeNode
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsSessionsByIdTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsSessionsByIdTreeRequest struct via the builder pattern


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


## GetAgentsSessionsStream

> GetAgentsSessionsStream(ctx).Execute()

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
	r, err := apiClient.AgentsAPI.GetAgentsSessionsStream(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsSessionsStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsSessionsStreamRequest struct via the builder pattern


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


## GetAgentsTargets

> TargetList GetAgentsTargets(ctx).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsTargets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsTargets`: TargetList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsTargets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsTargetsRequest struct via the builder pattern


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


## GetAgentsTargetsById

> TargetView GetAgentsTargetsById(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.GetAgentsTargetsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsTargetsById`: TargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsTargetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsTargetsByIdRequest struct via the builder pattern


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


## PatchAgentsByRef

> AgentView PatchAgentsByRef(ctx, ref).UpdateAgentIn(updateAgentIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PatchAgentsByRef(context.Background(), ref).UpdateAgentIn(updateAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PatchAgentsByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentsByRef`: AgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PatchAgentsByRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent to update — its public id or org-unique name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentsByRefRequest struct via the builder pattern


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


## PatchAgentsSessionsById

> SessionView PatchAgentsSessionsById(ctx, id).PatchSessionIn(patchSessionIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PatchAgentsSessionsById(context.Background(), id).PatchSessionIn(patchSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PatchAgentsSessionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentsSessionsById`: SessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PatchAgentsSessionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentsSessionsByIdRequest struct via the builder pattern


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


## PatchAgentsTargetsById

> TargetView PatchAgentsTargetsById(ctx, id).PatchTargetIn(patchTargetIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PatchAgentsTargetsById(context.Background(), id).PatchTargetIn(patchTargetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PatchAgentsTargetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentsTargetsById`: TargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PatchAgentsTargetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentsTargetsByIdRequest struct via the builder pattern


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


## PostAgents

> AgentView PostAgents(ctx).CreateAgentIn(createAgentIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgents(context.Background()).CreateAgentIn(createAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgents`: AgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsRequest struct via the builder pattern


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


## PostAgentsByRefRun

> PostAgentsByRefRun(ctx, ref).Execute()

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
	r, err := apiClient.AgentsAPI.PostAgentsByRefRun(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsByRefRun``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostAgentsByRefRunRequest struct via the builder pattern


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


## PostAgentsChat

> PostAgentsChat(ctx).Execute()

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
	r, err := apiClient.AgentsAPI.PostAgentsChat(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsChatRequest struct via the builder pattern


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


## PostAgentsChatConversations

> PostAgentsChatConversations(ctx).Execute()

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
	r, err := apiClient.AgentsAPI.PostAgentsChatConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsChatConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsChatConversationsRequest struct via the builder pattern


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


## PostAgentsCoding

> CodingStarted PostAgentsCoding(ctx).CodingStartIn(codingStartIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsCoding(context.Background()).CodingStartIn(codingStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsCoding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsCoding`: CodingStarted
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsCoding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsCodingRequest struct via the builder pattern


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


## PostAgentsSessions

> SessionView PostAgentsSessions(ctx).RegisterReq(registerReq).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsSessions(context.Background()).RegisterReq(registerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsSessions`: SessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsSessionsRequest struct via the builder pattern


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


## PostAgentsSessionsByIdEvents

> EventView PostAgentsSessionsByIdEvents(ctx, id).EventIn(eventIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsSessionsByIdEvents(context.Background(), id).EventIn(eventIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsSessionsByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsSessionsByIdEvents`: EventView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsSessionsByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to append to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsSessionsByIdEventsRequest struct via the builder pattern


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


## PostAgentsSessionsByIdMessage

> ControlResult PostAgentsSessionsByIdMessage(ctx, id).ControlIn(controlIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsSessionsByIdMessage(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsSessionsByIdMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsSessionsByIdMessage`: ControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsSessionsByIdMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsSessionsByIdMessageRequest struct via the builder pattern


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


## PostAgentsSessionsByIdPause

> ControlResult PostAgentsSessionsByIdPause(ctx, id).ControlIn(controlIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsSessionsByIdPause(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsSessionsByIdPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsSessionsByIdPause`: ControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsSessionsByIdPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsSessionsByIdPauseRequest struct via the builder pattern


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


## PostAgentsSessionsByIdResume

> ControlResult PostAgentsSessionsByIdResume(ctx, id).ControlIn(controlIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsSessionsByIdResume(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsSessionsByIdResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsSessionsByIdResume`: ControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsSessionsByIdResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsSessionsByIdResumeRequest struct via the builder pattern


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


## PostAgentsSessionsByIdStop

> ControlResult PostAgentsSessionsByIdStop(ctx, id).ControlIn(controlIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsSessionsByIdStop(context.Background(), id).ControlIn(controlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsSessionsByIdStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsSessionsByIdStop`: ControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsSessionsByIdStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to steer, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsSessionsByIdStopRequest struct via the builder pattern


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


## PostAgentsTargets

> TargetView PostAgentsTargets(ctx).TargetReq(targetReq).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsTargets(context.Background()).TargetReq(targetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsTargets`: TargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsTargetsRequest struct via the builder pattern


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


## PostAgentsTargetsByIdClaim

> RoutedRunOut PostAgentsTargetsByIdClaim(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsTargetsByIdClaim(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsTargetsByIdClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsTargetsByIdClaim`: RoutedRunOut
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsTargetsByIdClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsTargetsByIdClaimRequest struct via the builder pattern


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


## PostAgentsTargetsByIdKey

> ClaimKeyOut PostAgentsTargetsByIdKey(ctx, id).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsTargetsByIdKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsTargetsByIdKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsTargetsByIdKey`: ClaimKeyOut
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsTargetsByIdKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsTargetsByIdKeyRequest struct via the builder pattern


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


## PostAgentsTargetsByIdRunsByRunidReport

> ReportOut PostAgentsTargetsByIdRunsByRunidReport(ctx, id, runId).ReportRunIn(reportRunIn).Execute()

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
	resp, r, err := apiClient.AgentsAPI.PostAgentsTargetsByIdRunsByRunidReport(context.Background(), id, runId).ReportRunIn(reportRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.PostAgentsTargetsByIdRunsByRunidReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsTargetsByIdRunsByRunidReport`: ReportOut
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.PostAgentsTargetsByIdRunsByRunidReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine reporting, from the path. | 
**runId** | **string** | RunID is the routed run being completed, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsTargetsByIdRunsByRunidReportRequest struct via the builder pattern


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

