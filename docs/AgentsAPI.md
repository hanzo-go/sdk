# \AgentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1AgentsRef**](AgentsAPI.md#CloudDeleteV1AgentsRef) | **Delete** /v1/agents/{ref} | DeleteAgent removes an agent and every run recorded against it.
[**CloudDeleteV1AgentsTargetsId**](AgentsAPI.md#CloudDeleteV1AgentsTargetsId) | **Delete** /v1/agents/targets/{id} | DeleteTarget deregisters one machine.
[**CloudGetV1Agents**](AgentsAPI.md#CloudGetV1Agents) | **Get** /v1/agents | ListAgents returns every agent defined in the caller&#39;s org, each with the number of runs recorded against it.
[**CloudGetV1AgentsActivity**](AgentsAPI.md#CloudGetV1AgentsActivity) | **Get** /v1/agents/activity | AgentActivity serves the org-wide recent-activity feed.
[**CloudGetV1AgentsBuilds**](AgentsAPI.md#CloudGetV1AgentsBuilds) | **Get** /v1/agents/builds | ListBuilds returns the public index of every published build, most recently updated first, so a gallery can link straight to the story behind each product.
[**CloudGetV1AgentsBuildsOrgProject**](AgentsAPI.md#CloudGetV1AgentsBuildsOrgProject) | **Get** /v1/agents/builds/{org}/{project} | ReadBuild returns the readable build of one product: the agent session that produced it, turn by turn — the prompts, the reasoning, the commits each turn produced — plus the exact &#x60;git log&#x60; that re-derives every commit binding from git itself, so nothing here has to be taken on trust.
[**CloudGetV1AgentsMetrics**](AgentsAPI.md#CloudGetV1AgentsMetrics) | **Get** /v1/agents/metrics | AgentMetrics serves the invocations-over-time histogram for the org&#39;s Agents dashboard.
[**CloudGetV1AgentsRef**](AgentsAPI.md#CloudGetV1AgentsRef) | **Get** /v1/agents/{ref} | GetAgent returns one agent with its system prompt and its 20 most recent runs.
[**CloudGetV1AgentsRefRuns**](AgentsAPI.md#CloudGetV1AgentsRefRuns) | **Get** /v1/agents/{ref}/runs | ListAgentRuns returns one agent&#39;s execution history, newest first — each run&#39;s input, its output or its error, and how long it took.
[**CloudGetV1AgentsSessions**](AgentsAPI.md#CloudGetV1AgentsSessions) | **Get** /v1/agents/sessions | ListSessions returns the caller org&#39;s live sessions, newest first — each with its event count, its direct-child count and a one-line preview of its latest event.
[**CloudGetV1AgentsSessionsId**](AgentsAPI.md#CloudGetV1AgentsSessionsId) | **Get** /v1/agents/sessions/{id} | GetSession returns one session with its direct child sessions and its 50 most recent events, oldest of those first.
[**CloudGetV1AgentsSessionsIdControl**](AgentsAPI.md#CloudGetV1AgentsSessionsIdControl) | **Get** /v1/agents/sessions/{id}/control | DrainSessionControl returns the steering commands (pause/resume/stop/message) recorded against the caller&#39;s own session that are newer than the cursor, oldest first, with the cursor to poll from next.
[**CloudGetV1AgentsSessionsIdTree**](AgentsAPI.md#CloudGetV1AgentsSessionsIdTree) | **Get** /v1/agents/sessions/{id}/tree | SessionTree returns the subagent-flow graph rooted at this session: the session, its children, their children, each node carrying its own event count.
[**CloudGetV1AgentsSessionsStream**](AgentsAPI.md#CloudGetV1AgentsSessionsStream) | **Get** /v1/agents/sessions/stream | 
[**CloudGetV1AgentsTargets**](AgentsAPI.md#CloudGetV1AgentsTargets) | **Get** /v1/agents/targets | ListTargets returns every machine registered to the caller&#39;s org, newest first, each with its live session load.
[**CloudGetV1AgentsTargetsId**](AgentsAPI.md#CloudGetV1AgentsTargetsId) | **Get** /v1/agents/targets/{id} | GetTarget returns one registered machine, with its live session load.
[**CloudPatchV1AgentsRef**](AgentsAPI.md#CloudPatchV1AgentsRef) | **Patch** /v1/agents/{ref} | UpdateAgent changes an agent in place.
[**CloudPatchV1AgentsSessionsId**](AgentsAPI.md#CloudPatchV1AgentsSessionsId) | **Patch** /v1/agents/sessions/{id} | PatchSession updates a session&#39;s surface-owned truth: its status, its title, the run-target it is dispatched to, and the product it built plus whether that build&#39;s story is public.
[**CloudPatchV1AgentsTargetsId**](AgentsAPI.md#CloudPatchV1AgentsTargetsId) | **Patch** /v1/agents/targets/{id} | PatchTarget updates one machine in place.
[**CloudPostV1Agents**](AgentsAPI.md#CloudPostV1Agents) | **Post** /v1/agents | CreateAgent defines an agent in the caller&#39;s org: a model, a system prompt (instructions) and a set of tool names.
[**CloudPostV1AgentsByRefRun**](AgentsAPI.md#CloudPostV1AgentsByRefRun) | **Post** /v1/agents/{ref}/run | 
[**CloudPostV1AgentsSessions**](AgentsAPI.md#CloudPostV1AgentsSessions) | **Post** /v1/agents/sessions | RegisterSession opens a live agent session in the caller&#39;s org — the row every surface (the CLI&#39;s outer agent, hanzo.bot, the console, chat) hangs its activity off.
[**CloudPostV1AgentsSessionsByIdEvents**](AgentsAPI.md#CloudPostV1AgentsSessionsByIdEvents) | **Post** /v1/agents/sessions/{id}/events | 
[**CloudPostV1AgentsSessionsByIdMessage**](AgentsAPI.md#CloudPostV1AgentsSessionsByIdMessage) | **Post** /v1/agents/sessions/{id}/message | 
[**CloudPostV1AgentsSessionsByIdPause**](AgentsAPI.md#CloudPostV1AgentsSessionsByIdPause) | **Post** /v1/agents/sessions/{id}/pause | 
[**CloudPostV1AgentsSessionsByIdResume**](AgentsAPI.md#CloudPostV1AgentsSessionsByIdResume) | **Post** /v1/agents/sessions/{id}/resume | 
[**CloudPostV1AgentsSessionsByIdStop**](AgentsAPI.md#CloudPostV1AgentsSessionsByIdStop) | **Post** /v1/agents/sessions/{id}/stop | 
[**CloudPostV1AgentsTargets**](AgentsAPI.md#CloudPostV1AgentsTargets) | **Post** /v1/agents/targets | RegisterTarget registers a machine as an agent target, or re-links one that is already registered.
[**CloudPostV1AgentsTargetsIdClaim**](AgentsAPI.md#CloudPostV1AgentsTargetsIdClaim) | **Post** /v1/agents/targets/{id}/claim | ClaimRoutedRun is the machine&#39;s long poll for work: it authenticates the daemon, stamps the liveness the dispatch gate reads (the poll IS the proof a runner is listening), and waits up to 25 seconds for the next run addressed to THIS machine.
[**CloudPostV1AgentsTargetsIdClaimKey**](AgentsAPI.md#CloudPostV1AgentsTargetsIdClaimKey) | **Post** /v1/agents/targets/{id}/claim-key | MintTargetClaimKey mints (or rotates) the claim key a &#x60;hanzo code --serve&#x60; daemon presents to claim work for this machine, and returns it ONCE: only its SHA-256 hash is stored.
[**CloudPostV1AgentsTargetsIdRunsRunIdReport**](AgentsAPI.md#CloudPostV1AgentsTargetsIdRunsRunIdReport) | **Post** /v1/agents/targets/{id}/runs/{runId}/report | ReportRoutedRun completes a claimed run: it delivers the terminal result to the run&#39;s durable owner, which is what lets that workflow finish.



## CloudDeleteV1AgentsRef

> CloudDeleteV1AgentsRef(ctx, ref).Execute()

DeleteAgent removes an agent and every run recorded against it.



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
	ref := "helper" // string | Ref is the agent's public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudDeleteV1AgentsRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudDeleteV1AgentsRef``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1AgentsRefRequest struct via the builder pattern


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


## CloudDeleteV1AgentsTargetsId

> CloudTargetDeleted CloudDeleteV1AgentsTargetsId(ctx, id).Execute()

DeleteTarget deregisters one machine.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudDeleteV1AgentsTargetsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudDeleteV1AgentsTargetsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1AgentsTargetsId`: CloudTargetDeleted
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudDeleteV1AgentsTargetsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1AgentsTargetsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTargetDeleted**](CloudTargetDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Agents

> CloudAgentList CloudGetV1Agents(ctx).Execute()

ListAgents returns every agent defined in the caller's org, each with the number of runs recorded against it.



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
	resp, r, err := apiClient.AgentsAPI.CloudGetV1Agents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1Agents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Agents`: CloudAgentList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1Agents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsRequest struct via the builder pattern


### Return type

[**CloudAgentList**](CloudAgentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsActivity

> CloudActivityFeed CloudGetV1AgentsActivity(ctx).Execute()

AgentActivity serves the org-wide recent-activity feed.



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
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsActivity`: CloudActivityFeed
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsActivityRequest struct via the builder pattern


### Return type

[**CloudActivityFeed**](CloudActivityFeed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsBuilds

> CloudBuildList CloudGetV1AgentsBuilds(ctx).Limit(limit).Execute()

ListBuilds returns the public index of every published build, most recently updated first, so a gallery can link straight to the story behind each product.



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
	limit := int32(56) // int32 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsBuilds(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsBuilds`: CloudBuildList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the page. Absent, zero or over 500 reads as 100. | 

### Return type

[**CloudBuildList**](CloudBuildList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsBuildsOrgProject

> CloudBuildView CloudGetV1AgentsBuildsOrgProject(ctx, org, project).Execute()

ReadBuild returns the readable build of one product: the agent session that produced it, turn by turn — the prompts, the reasoning, the commits each turn produced — plus the exact `git log` that re-derives every commit binding from git itself, so nothing here has to be taken on trust.



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
	org := "hanzo" // string | Org is the org that published the build, from the path.
	project := "landing" // string | Project is the product's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsBuildsOrgProject(context.Background(), org, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsBuildsOrgProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsBuildsOrgProject`: CloudBuildView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsBuildsOrgProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the org that published the build, from the path. | 
**project** | **string** | Project is the product&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsBuildsOrgProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudBuildView**](CloudBuildView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsMetrics

> CloudMetricsView CloudGetV1AgentsMetrics(ctx).Range_(range_).Execute()

AgentMetrics serves the invocations-over-time histogram for the org's Agents dashboard.



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
	range_ := "7D" // string | Range is the window to bucket: 24H, 7D or 30D. Anything else reads as 30D. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsMetrics`: CloudMetricsView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the window to bucket: 24H, 7D or 30D. Anything else reads as 30D. | 

### Return type

[**CloudMetricsView**](CloudMetricsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsRef

> CloudAgentDetail CloudGetV1AgentsRef(ctx, ref).Execute()

GetAgent returns one agent with its system prompt and its 20 most recent runs.



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
	ref := "helper" // string | Ref is the agent's public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsRef`: CloudAgentDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id (the agent_… handle create and list return) or its org-unique name, from the path. Either resolves the same agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAgentDetail**](CloudAgentDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsRefRuns

> CloudRunList CloudGetV1AgentsRefRuns(ctx, ref).Limit(limit).Execute()

ListAgentRuns returns one agent's execution history, newest first — each run's input, its output or its error, and how long it took.



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
	ref := "helper" // string | Ref is the agent's public id or its org-unique name, from the path.
	limit := int32(20) // int32 | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsRefRuns(context.Background(), ref).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsRefRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsRefRuns`: CloudRunList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsRefRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s public id or its org-unique name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsRefRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps how many runs come back, newest first. Absent, zero or out of range (1..200) reads as 50. | 

### Return type

[**CloudRunList**](CloudRunList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsSessions

> CloudSessionList CloudGetV1AgentsSessions(ctx).Root(root).Parent(parent).Status(status).Project(project).Limit(limit).Execute()

ListSessions returns the caller org's live sessions, newest first — each with its event count, its direct-child count and a one-line preview of its latest event.



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
	root := "root_example" // string | Root scopes the page to one subagent tree (its root session id). (optional)
	parent := "parent_example" // string | Parent scopes the page to the direct children of one session. Ignored when root is set; with neither, only ROOT sessions come back. (optional)
	status := "running" // string | Status filters to running, paused, done or error. (optional)
	project := "project_example" // string | Project filters to the sessions tagged with one product slug. (optional)
	limit := int32(20) // int32 | Limit caps the page. Absent, zero or over 500 reads as 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsSessions(context.Background()).Root(root).Parent(parent).Status(status).Project(project).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsSessions`: CloudSessionList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **root** | **string** | Root scopes the page to one subagent tree (its root session id). | 
 **parent** | **string** | Parent scopes the page to the direct children of one session. Ignored when root is set; with neither, only ROOT sessions come back. | 
 **status** | **string** | Status filters to running, paused, done or error. | 
 **project** | **string** | Project filters to the sessions tagged with one product slug. | 
 **limit** | **int32** | Limit caps the page. Absent, zero or over 500 reads as 100. | 

### Return type

[**CloudSessionList**](CloudSessionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsSessionsId

> CloudSessionDetail CloudGetV1AgentsSessionsId(ctx, id).Execute()

GetSession returns one session with its direct child sessions and its 50 most recent events, oldest of those first.



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
	id := "sess_1" // string | ID is the session to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsSessionsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsSessionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsSessionsId`: CloudSessionDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsSessionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsSessionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSessionDetail**](CloudSessionDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsSessionsIdControl

> CloudControlDrain CloudGetV1AgentsSessionsIdControl(ctx, id).After(after).Execute()

DrainSessionControl returns the steering commands (pause/resume/stop/message) recorded against the caller's own session that are newer than the cursor, oldest first, with the cursor to poll from next.



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
	id := "sess_1" // string | ID is the session whose commands are being drained, from the path.
	after := int32(12) // int32 | After is the last seq this poller applied; only commands newer than it come back. Absent or negative reads as 0, which drains from the beginning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsSessionsIdControl(context.Background(), id).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsSessionsIdControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsSessionsIdControl`: CloudControlDrain
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsSessionsIdControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session whose commands are being drained, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsSessionsIdControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int32** | After is the last seq this poller applied; only commands newer than it come back. Absent or negative reads as 0, which drains from the beginning. | 

### Return type

[**CloudControlDrain**](CloudControlDrain.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsSessionsIdTree

> CloudTreeNode CloudGetV1AgentsSessionsIdTree(ctx, id).Execute()

SessionTree returns the subagent-flow graph rooted at this session: the session, its children, their children, each node carrying its own event count.



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
	id := "sess_1" // string | ID is the session to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsSessionsIdTree(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsSessionsIdTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsSessionsIdTree`: CloudTreeNode
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsSessionsIdTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsSessionsIdTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTreeNode**](CloudTreeNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsSessionsStream

> CloudGetV1AgentsSessionsStream(ctx).Execute()



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
	r, err := apiClient.AgentsAPI.CloudGetV1AgentsSessionsStream(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsSessionsStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsSessionsStreamRequest struct via the builder pattern


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


## CloudGetV1AgentsTargets

> CloudTargetList CloudGetV1AgentsTargets(ctx).Execute()

ListTargets returns every machine registered to the caller's org, newest first, each with its live session load.



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
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsTargets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsTargets`: CloudTargetList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsTargets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsTargetsRequest struct via the builder pattern


### Return type

[**CloudTargetList**](CloudTargetList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AgentsTargetsId

> CloudTargetView CloudGetV1AgentsTargetsId(ctx, id).Execute()

GetTarget returns one registered machine, with its live session load.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudGetV1AgentsTargetsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudGetV1AgentsTargetsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AgentsTargetsId`: CloudTargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudGetV1AgentsTargetsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AgentsTargetsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTargetView**](CloudTargetView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AgentsRef

> CloudAgentView CloudPatchV1AgentsRef(ctx, ref).CloudUpdateAgentIn(cloudUpdateAgentIn).Execute()

UpdateAgent changes an agent in place.



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
	ref := "helper" // string | Ref is the agent to update — its public id or org-unique name, from the path.
	cloudUpdateAgentIn := *openapiclient.NewCloudUpdateAgentIn() // CloudUpdateAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPatchV1AgentsRef(context.Background(), ref).CloudUpdateAgentIn(cloudUpdateAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPatchV1AgentsRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1AgentsRef`: CloudAgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPatchV1AgentsRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent to update — its public id or org-unique name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AgentsRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudUpdateAgentIn** | [**CloudUpdateAgentIn**](CloudUpdateAgentIn.md) |  | 

### Return type

[**CloudAgentView**](CloudAgentView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AgentsSessionsId

> CloudSessionView CloudPatchV1AgentsSessionsId(ctx, id).CloudPatchSessionIn(cloudPatchSessionIn).Execute()

PatchSession updates a session's surface-owned truth: its status, its title, the run-target it is dispatched to, and the product it built plus whether that build's story is public.



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
	id := "sess_1" // string | ID is the session to update, from the path.
	cloudPatchSessionIn := *openapiclient.NewCloudPatchSessionIn() // CloudPatchSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPatchV1AgentsSessionsId(context.Background(), id).CloudPatchSessionIn(cloudPatchSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPatchV1AgentsSessionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1AgentsSessionsId`: CloudSessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPatchV1AgentsSessionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the session to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AgentsSessionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPatchSessionIn** | [**CloudPatchSessionIn**](CloudPatchSessionIn.md) |  | 

### Return type

[**CloudSessionView**](CloudSessionView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AgentsTargetsId

> CloudTargetView CloudPatchV1AgentsTargetsId(ctx, id).CloudPatchTargetIn(cloudPatchTargetIn).Execute()

PatchTarget updates one machine in place.



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
	id := "tgt_1" // string | ID is the target to update, from the path.
	cloudPatchTargetIn := *openapiclient.NewCloudPatchTargetIn() // CloudPatchTargetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPatchV1AgentsTargetsId(context.Background(), id).CloudPatchTargetIn(cloudPatchTargetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPatchV1AgentsTargetsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1AgentsTargetsId`: CloudTargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPatchV1AgentsTargetsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AgentsTargetsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPatchTargetIn** | [**CloudPatchTargetIn**](CloudPatchTargetIn.md) |  | 

### Return type

[**CloudTargetView**](CloudTargetView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Agents

> CloudAgentView CloudPostV1Agents(ctx).CloudCreateAgentIn(cloudCreateAgentIn).Execute()

CreateAgent defines an agent in the caller's org: a model, a system prompt (instructions) and a set of tool names.



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
	cloudCreateAgentIn := *openapiclient.NewCloudCreateAgentIn() // CloudCreateAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPostV1Agents(context.Background()).CloudCreateAgentIn(cloudCreateAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1Agents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Agents`: CloudAgentView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPostV1Agents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateAgentIn** | [**CloudCreateAgentIn**](CloudCreateAgentIn.md) |  | 

### Return type

[**CloudAgentView**](CloudAgentView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AgentsByRefRun

> CloudPostV1AgentsByRefRun(ctx, ref).Execute()



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
	ref := "ref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudPostV1AgentsByRefRun(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsByRefRun``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AgentsByRefRunRequest struct via the builder pattern


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


## CloudPostV1AgentsSessions

> CloudSessionView CloudPostV1AgentsSessions(ctx).CloudRegisterReq(cloudRegisterReq).Execute()

RegisterSession opens a live agent session in the caller's org — the row every surface (the CLI's outer agent, hanzo.bot, the console, chat) hangs its activity off.



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
	cloudRegisterReq := *openapiclient.NewCloudRegisterReq() // CloudRegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPostV1AgentsSessions(context.Background()).CloudRegisterReq(cloudRegisterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AgentsSessions`: CloudSessionView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPostV1AgentsSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AgentsSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRegisterReq** | [**CloudRegisterReq**](CloudRegisterReq.md) |  | 

### Return type

[**CloudSessionView**](CloudSessionView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AgentsSessionsByIdEvents

> CloudPostV1AgentsSessionsByIdEvents(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudPostV1AgentsSessionsByIdEvents(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsSessionsByIdEvents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AgentsSessionsByIdEventsRequest struct via the builder pattern


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


## CloudPostV1AgentsSessionsByIdMessage

> CloudPostV1AgentsSessionsByIdMessage(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudPostV1AgentsSessionsByIdMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsSessionsByIdMessage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AgentsSessionsByIdMessageRequest struct via the builder pattern


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


## CloudPostV1AgentsSessionsByIdPause

> CloudPostV1AgentsSessionsByIdPause(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudPostV1AgentsSessionsByIdPause(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsSessionsByIdPause``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AgentsSessionsByIdPauseRequest struct via the builder pattern


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


## CloudPostV1AgentsSessionsByIdResume

> CloudPostV1AgentsSessionsByIdResume(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudPostV1AgentsSessionsByIdResume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsSessionsByIdResume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AgentsSessionsByIdResumeRequest struct via the builder pattern


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


## CloudPostV1AgentsSessionsByIdStop

> CloudPostV1AgentsSessionsByIdStop(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.CloudPostV1AgentsSessionsByIdStop(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsSessionsByIdStop``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1AgentsSessionsByIdStopRequest struct via the builder pattern


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


## CloudPostV1AgentsTargets

> CloudTargetView CloudPostV1AgentsTargets(ctx).CloudTargetReq(cloudTargetReq).Execute()

RegisterTarget registers a machine as an agent target, or re-links one that is already registered.



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
	cloudTargetReq := *openapiclient.NewCloudTargetReq() // CloudTargetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPostV1AgentsTargets(context.Background()).CloudTargetReq(cloudTargetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AgentsTargets`: CloudTargetView
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPostV1AgentsTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AgentsTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudTargetReq** | [**CloudTargetReq**](CloudTargetReq.md) |  | 

### Return type

[**CloudTargetView**](CloudTargetView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AgentsTargetsIdClaim

> CloudRoutedRunOut CloudPostV1AgentsTargetsIdClaim(ctx, id).Execute()

ClaimRoutedRun is the machine's long poll for work: it authenticates the daemon, stamps the liveness the dispatch gate reads (the poll IS the proof a runner is listening), and waits up to 25 seconds for the next run addressed to THIS machine.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPostV1AgentsTargetsIdClaim(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsTargetsIdClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AgentsTargetsIdClaim`: CloudRoutedRunOut
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPostV1AgentsTargetsIdClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AgentsTargetsIdClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRoutedRunOut**](CloudRoutedRunOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AgentsTargetsIdClaimKey

> CloudClaimKeyOut CloudPostV1AgentsTargetsIdClaimKey(ctx, id).Execute()

MintTargetClaimKey mints (or rotates) the claim key a `hanzo code --serve` daemon presents to claim work for this machine, and returns it ONCE: only its SHA-256 hash is stored.



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
	id := "tgt_1" // string | ID is the target to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPostV1AgentsTargetsIdClaimKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsTargetsIdClaimKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AgentsTargetsIdClaimKey`: CloudClaimKeyOut
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPostV1AgentsTargetsIdClaimKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the target to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AgentsTargetsIdClaimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudClaimKeyOut**](CloudClaimKeyOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AgentsTargetsIdRunsRunIdReport

> CloudReportOut CloudPostV1AgentsTargetsIdRunsRunIdReport(ctx, id, runId).CloudReportRunIn(cloudReportRunIn).Execute()

ReportRoutedRun completes a claimed run: it delivers the terminal result to the run's durable owner, which is what lets that workflow finish.



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
	id := "tgt_1" // string | ID is the machine reporting, from the path.
	runId := "run_1" // string | RunID is the routed run being completed, from the path.
	cloudReportRunIn := *openapiclient.NewCloudReportRunIn() // CloudReportRunIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CloudPostV1AgentsTargetsIdRunsRunIdReport(context.Background(), id, runId).CloudReportRunIn(cloudReportRunIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CloudPostV1AgentsTargetsIdRunsRunIdReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AgentsTargetsIdRunsRunIdReport`: CloudReportOut
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CloudPostV1AgentsTargetsIdRunsRunIdReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine reporting, from the path. | 
**runId** | **string** | RunID is the routed run being completed, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AgentsTargetsIdRunsRunIdReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudReportRunIn** | [**CloudReportRunIn**](CloudReportRunIn.md) |  | 

### Return type

[**CloudReportOut**](CloudReportOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

