# \TaskAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaskProjectsByKey**](TaskAPI.md#DeleteTaskProjectsByKey) | **Delete** /v1/task/projects/{key} | Refused — a board is a repository on the forge
[**GetTaskBoard**](TaskAPI.md#GetTaskBoard) | **Get** /v1/task/board | Returns a board&#39;s issues — work items with their column, priority, assignee, labels and schedule.
[**GetTaskIssues**](TaskAPI.md#GetTaskIssues) | **Get** /v1/task/issues | Answers across every project in the org.
[**GetTaskProjects**](TaskAPI.md#GetTaskProjects) | **Get** /v1/task/projects | Returns the boards of your org — the places your work actually is.
[**GetTaskProjectsByKey**](TaskAPI.md#GetTaskProjectsByKey) | **Get** /v1/task/projects/{key} | Returns one board of your org by its key — the repository name.
[**GetTaskProjectsByKeyIssues**](TaskAPI.md#GetTaskProjectsByKeyIssues) | **Get** /v1/task/projects/{key}/issues | Returns a board&#39;s issues — work items with their column, priority, assignee, labels and schedule.
[**GetTaskProjectsByKeyIssuesByNum**](TaskAPI.md#GetTaskProjectsByKeyIssuesByNum) | **Get** /v1/task/projects/{key}/issues/{num} | Returns ONE work item in full — its description included.
[**GetTaskRoomsByRoom**](TaskAPI.md#GetTaskRoomsByRoom) | **Get** /v1/task/rooms/{room} | Summarises one room&#39;s work.
[**PatchTaskProjectsByKey**](TaskAPI.md#PatchTaskProjectsByKey) | **Patch** /v1/task/projects/{key} | Refused — a board is a repository on the forge
[**PatchTaskProjectsByKeyIssuesByNum**](TaskAPI.md#PatchTaskProjectsByKeyIssuesByNum) | **Patch** /v1/task/projects/{key}/issues/{num} | Edits a work item — rename it, rewrite it, move it to another column, re-prioritise it, hand it to somebody, or comment on it.
[**PostTaskProjects**](TaskAPI.md#PostTaskProjects) | **Post** /v1/task/projects | Refused — a board is a repository on the forge
[**PostTaskProjectsByKeyIssues**](TaskAPI.md#PostTaskProjectsByKeyIssues) | **Post** /v1/task/projects/{key}/issues | Opens a work item on the board, filed as YOU, where the repository lives: an issue on the deployment&#39;s forge, or on GitHub for a repository the forge mirrors from there or does not hold at all.
[**PostTaskProjectsByKeyIssuesByNumClaim**](TaskAPI.md#PostTaskProjectsByKeyIssuesByNumClaim) | **Post** /v1/task/projects/{key}/issues/{num}/claim | Takes an issue: it becomes yours and it moves to in_progress.



## DeleteTaskProjectsByKey

> DeleteTaskProjectsByKey(ctx, key).Execute()

Refused — a board is a repository on the forge



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.DeleteTaskProjectsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.DeleteTaskProjectsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskProjectsByKeyRequest struct via the builder pattern


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


## GetTaskBoard

> []TaskIssueView GetTaskBoard(ctx).Key(key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()

Returns a board's issues — work items with their column, priority, assignee, labels and schedule.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | Key is the project whose issues to list, from the path. EMPTY means every project in the org — the global board. It is a filter like the rest of this struct rather than an address, which is what lets one op answer both \"this board\" and \"all the work\" without a second surface disagreeing with the first about what a column is. (optional)
	status := "status_example" // string | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. (optional)
	kind := "kind_example" // string | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. (optional)
	repo := "repo_example" // string | Repo keeps only issues bound to that git repository. (optional)
	label := "label_example" // string | Label keeps only issues carrying that label, compared case-insensitively.  This is how a board narrows to something SMALLER than a repository — the one mechanism for it. An estate whose apps are directories inside one repository (hanzoai/cloud carries ~140 of them) has no repository per app to address, so the app is a label: `label=app/meet` is the meet board. Nothing is provisioned to make one exist; a board is the query. (optional)
	source := "source_example" // string | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. (optional)
	scheduled := true // bool | Scheduled keeps only issues that carry a date — a start, a due date or both. This is the timeline's slice of the board: pass scheduled=true to get exactly the rows a gantt has somewhere to draw, instead of fetching every issue and discarding the undated ones client-side. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskBoard(context.Background()).Key(key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskBoard`: []TaskIssueView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Key is the project whose issues to list, from the path. EMPTY means every project in the org — the global board. It is a filter like the rest of this struct rather than an address, which is what lets one op answer both \&quot;this board\&quot; and \&quot;all the work\&quot; without a second surface disagreeing with the first about what a column is. | 
 **status** | **string** | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. | 
 **kind** | **string** | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. | 
 **repo** | **string** | Repo keeps only issues bound to that git repository. | 
 **label** | **string** | Label keeps only issues carrying that label, compared case-insensitively.  This is how a board narrows to something SMALLER than a repository — the one mechanism for it. An estate whose apps are directories inside one repository (hanzoai/cloud carries ~140 of them) has no repository per app to address, so the app is a label: &#x60;label&#x3D;app/meet&#x60; is the meet board. Nothing is provisioned to make one exist; a board is the query. | 
 **source** | **string** | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. | 
 **scheduled** | **bool** | Scheduled keeps only issues that carry a date — a start, a due date or both. This is the timeline&#39;s slice of the board: pass scheduled&#x3D;true to get exactly the rows a gantt has somewhere to draw, instead of fetching every issue and discarding the undated ones client-side. | 

### Return type

[**[]TaskIssueView**](TaskIssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskIssues

> TaskIssueHits GetTaskIssues(ctx).Q(q).Project(project).Status(status).Kind(kind).Repo(repo).Room(room).Source(source).Assignee(assignee).Limit(limit).Execute()

Answers across every project in the org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	q := "q_example" // string | Q matches an issue's title or description. A word from the issue, which is what someone remembers — not its number, which is what they are looking up. (optional)
	project := "project_example" // string | Project narrows to one team key; \"\" searches every project in the org, which is the point of this op. (optional)
	status := "status_example" // string | Status keeps one board column: backlog, todo, in_progress, done, canceled. (optional)
	kind := "kind_example" // string | Kind keeps one shape: issue, pr, epic. (optional)
	repo := "repo_example" // string | Repo keeps issues bound to one git repository. (optional)
	room := "room_example" // string | Room keeps issues bound to one collaboration room, spelled \"<space>_<room>\" — the exact value GET /v1/meet/call answers with, so a channel's call and its task list name the room the same way. This is the read a channel view runs to draw its own list; it spans every board of the org, because the work a channel is about is not confined to one board. (optional)
	source := "source_example" // string | Source keeps one origin: team, git, crm, helpdesk, cms, agent. \"git\" is how you ask for the mirrored GitHub issues specifically. (optional)
	assignee := "assignee_example" // string | Assignee keeps issues held by one person. Pass \"me\" for yourself. (optional)
	limit := int64(789) // int64 | Limit caps the answer; 0 means the default, and anything above the ceiling is clamped rather than refused — a search that errors on being too broad teaches people to guess. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskIssues(context.Background()).Q(q).Project(project).Status(status).Kind(kind).Repo(repo).Room(room).Source(source).Assignee(assignee).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskIssues`: TaskIssueHits
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q matches an issue&#39;s title or description. A word from the issue, which is what someone remembers — not its number, which is what they are looking up. | 
 **project** | **string** | Project narrows to one team key; \&quot;\&quot; searches every project in the org, which is the point of this op. | 
 **status** | **string** | Status keeps one board column: backlog, todo, in_progress, done, canceled. | 
 **kind** | **string** | Kind keeps one shape: issue, pr, epic. | 
 **repo** | **string** | Repo keeps issues bound to one git repository. | 
 **room** | **string** | Room keeps issues bound to one collaboration room, spelled \&quot;&lt;space&gt;_&lt;room&gt;\&quot; — the exact value GET /v1/meet/call answers with, so a channel&#39;s call and its task list name the room the same way. This is the read a channel view runs to draw its own list; it spans every board of the org, because the work a channel is about is not confined to one board. | 
 **source** | **string** | Source keeps one origin: team, git, crm, helpdesk, cms, agent. \&quot;git\&quot; is how you ask for the mirrored GitHub issues specifically. | 
 **assignee** | **string** | Assignee keeps issues held by one person. Pass \&quot;me\&quot; for yourself. | 
 **limit** | **int64** | Limit caps the answer; 0 means the default, and anything above the ceiling is clamped rather than refused — a search that errors on being too broad teaches people to guess. | 

### Return type

[**TaskIssueHits**](TaskIssueHits.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskProjects

> []TaskBoardView GetTaskProjects(ctx).Execute()

Returns the boards of your org — the places your work actually is.



### Example

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
	resp, r, err := apiClient.TaskAPI.GetTaskProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskProjects`: []TaskBoardView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskProjectsRequest struct via the builder pattern


### Return type

[**[]TaskBoardView**](TaskBoardView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskProjectsByKey

> TaskBoardView GetTaskProjectsByKey(ctx, key).Execute()

Returns one board of your org by its key — the repository name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | Key is the project's org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\"ENG\", \"OPS2\"). Matched case-insensitively.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskProjectsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskProjectsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskProjectsByKey`: TaskBoardView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskProjectsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project&#39;s org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\&quot;ENG\&quot;, \&quot;OPS2\&quot;). Matched case-insensitively. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskProjectsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskBoardView**](TaskBoardView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskProjectsByKeyIssues

> []TaskIssueView GetTaskProjectsByKeyIssues(ctx, key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()

Returns a board's issues — work items with their column, priority, assignee, labels and schedule.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | Key is the project whose issues to list, from the path. EMPTY means every project in the org — the global board. It is a filter like the rest of this struct rather than an address, which is what lets one op answer both \"this board\" and \"all the work\" without a second surface disagreeing with the first about what a column is.
	status := "status_example" // string | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. (optional)
	kind := "kind_example" // string | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. (optional)
	repo := "repo_example" // string | Repo keeps only issues bound to that git repository. (optional)
	label := "label_example" // string | Label keeps only issues carrying that label, compared case-insensitively.  This is how a board narrows to something SMALLER than a repository — the one mechanism for it. An estate whose apps are directories inside one repository (hanzoai/cloud carries ~140 of them) has no repository per app to address, so the app is a label: `label=app/meet` is the meet board. Nothing is provisioned to make one exist; a board is the query. (optional)
	source := "source_example" // string | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. (optional)
	scheduled := true // bool | Scheduled keeps only issues that carry a date — a start, a due date or both. This is the timeline's slice of the board: pass scheduled=true to get exactly the rows a gantt has somewhere to draw, instead of fetching every issue and discarding the undated ones client-side. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskProjectsByKeyIssues(context.Background(), key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskProjectsByKeyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskProjectsByKeyIssues`: []TaskIssueView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskProjectsByKeyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project whose issues to list, from the path. EMPTY means every project in the org — the global board. It is a filter like the rest of this struct rather than an address, which is what lets one op answer both \&quot;this board\&quot; and \&quot;all the work\&quot; without a second surface disagreeing with the first about what a column is. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskProjectsByKeyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. | 
 **kind** | **string** | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. | 
 **repo** | **string** | Repo keeps only issues bound to that git repository. | 
 **label** | **string** | Label keeps only issues carrying that label, compared case-insensitively.  This is how a board narrows to something SMALLER than a repository — the one mechanism for it. An estate whose apps are directories inside one repository (hanzoai/cloud carries ~140 of them) has no repository per app to address, so the app is a label: &#x60;label&#x3D;app/meet&#x60; is the meet board. Nothing is provisioned to make one exist; a board is the query. | 
 **source** | **string** | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. | 
 **scheduled** | **bool** | Scheduled keeps only issues that carry a date — a start, a due date or both. This is the timeline&#39;s slice of the board: pass scheduled&#x3D;true to get exactly the rows a gantt has somewhere to draw, instead of fetching every issue and discarding the undated ones client-side. | 

### Return type

[**[]TaskIssueView**](TaskIssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskProjectsByKeyIssuesByNum

> TaskIssueView GetTaskProjectsByKeyIssuesByNum(ctx, key, num).Execute()

Returns ONE work item in full — its description included.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | Key is the board — the repository name, or an index board's key.
	num := int64(789) // int64 | Num is the issue's number on that board.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskProjectsByKeyIssuesByNum(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskProjectsByKeyIssuesByNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskProjectsByKeyIssuesByNum`: TaskIssueView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskProjectsByKeyIssuesByNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the board — the repository name, or an index board&#39;s key. | 
**num** | **int64** | Num is the issue&#39;s number on that board. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskProjectsByKeyIssuesByNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskIssueView**](TaskIssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskRoomsByRoom

> TaskRoomWork GetTaskRoomsByRoom(ctx, room).Execute()

Summarises one room's work.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	room := "room_example" // string | Room is the room, spelled \"<space>_<room>\" — the same value GET /v1/meet/call answers with, so a channel's call and its work name the room identically. From the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskRoomsByRoom(context.Background(), room).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskRoomsByRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskRoomsByRoom`: TaskRoomWork
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskRoomsByRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**room** | **string** | Room is the room, spelled \&quot;&lt;space&gt;_&lt;room&gt;\&quot; — the same value GET /v1/meet/call answers with, so a channel&#39;s call and its work name the room identically. From the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRoomsByRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskRoomWork**](TaskRoomWork.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTaskProjectsByKey

> PatchTaskProjectsByKey(ctx, key).Execute()

Refused — a board is a repository on the forge



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.PatchTaskProjectsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PatchTaskProjectsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTaskProjectsByKeyRequest struct via the builder pattern


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


## PatchTaskProjectsByKeyIssuesByNum

> TaskIssueView PatchTaskProjectsByKeyIssuesByNum(ctx, key, num).TaskIssueEdit(taskIssueEdit).Execute()

Edits a work item — rename it, rewrite it, move it to another column, re-prioritise it, hand it to somebody, or comment on it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | Key is the board — the repository name, from the path.
	num := int64(789) // int64 | Num is the issue number on that repository, from the path.
	taskIssueEdit := *openapiclient.NewTaskIssueEdit() // TaskIssueEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.PatchTaskProjectsByKeyIssuesByNum(context.Background(), key, num).TaskIssueEdit(taskIssueEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PatchTaskProjectsByKeyIssuesByNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTaskProjectsByKeyIssuesByNum`: TaskIssueView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.PatchTaskProjectsByKeyIssuesByNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the board — the repository name, from the path. | 
**num** | **int64** | Num is the issue number on that repository, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTaskProjectsByKeyIssuesByNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taskIssueEdit** | [**TaskIssueEdit**](TaskIssueEdit.md) |  | 

### Return type

[**TaskIssueView**](TaskIssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaskProjects

> PostTaskProjects(ctx).Execute()

Refused — a board is a repository on the forge



### Example

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
	r, err := apiClient.TaskAPI.PostTaskProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PostTaskProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaskProjectsRequest struct via the builder pattern


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


## PostTaskProjectsByKeyIssues

> TaskIssueView PostTaskProjectsByKeyIssues(ctx, key).TaskNewIssue(taskNewIssue).Execute()

Opens a work item on the board, filed as YOU, where the repository lives: an issue on the deployment's forge, or on GitHub for a repository the forge mirrors from there or does not hold at all.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | Key is the board — the repository name, from the path. `owner/name`, or `github.com/owner/name`, names a repository on GitHub that no board on the forge stands for.
	taskNewIssue := *openapiclient.NewTaskNewIssue() // TaskNewIssue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.PostTaskProjectsByKeyIssues(context.Background(), key).TaskNewIssue(taskNewIssue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PostTaskProjectsByKeyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaskProjectsByKeyIssues`: TaskIssueView
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.PostTaskProjectsByKeyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the board — the repository name, from the path. &#x60;owner/name&#x60;, or &#x60;github.com/owner/name&#x60;, names a repository on GitHub that no board on the forge stands for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaskProjectsByKeyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskNewIssue** | [**TaskNewIssue**](TaskNewIssue.md) |  | 

### Return type

[**TaskIssueView**](TaskIssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTaskProjectsByKeyIssuesByNumClaim

> TaskIssueHit PostTaskProjectsByKeyIssuesByNumClaim(ctx, key, num).Execute()

Takes an issue: it becomes yours and it moves to in_progress.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	key := "key_example" // string | 
	num := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.PostTaskProjectsByKeyIssuesByNumClaim(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PostTaskProjectsByKeyIssuesByNumClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTaskProjectsByKeyIssuesByNumClaim`: TaskIssueHit
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.PostTaskProjectsByKeyIssuesByNumClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**num** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTaskProjectsByKeyIssuesByNumClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskIssueHit**](TaskIssueHit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

