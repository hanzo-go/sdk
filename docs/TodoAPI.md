# \TodoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTodoProjectsByKey**](TodoAPI.md#DeleteTodoProjectsByKey) | **Delete** /v1/todo/projects/{key} | Refused — a board is a repository on the forge
[**GetTodoBoard**](TodoAPI.md#GetTodoBoard) | **Get** /v1/todo/board | Returns a board&#39;s issues — work items with their column, priority, assignee, labels and schedule.
[**GetTodoIssues**](TodoAPI.md#GetTodoIssues) | **Get** /v1/todo/issues | Answers across every project in the org.
[**GetTodoProjects**](TodoAPI.md#GetTodoProjects) | **Get** /v1/todo/projects | Returns the boards of your org — the places your work actually is.
[**GetTodoProjectsByKey**](TodoAPI.md#GetTodoProjectsByKey) | **Get** /v1/todo/projects/{key} | Returns one board of your org by its key — the repository name.
[**GetTodoProjectsByKeyIssues**](TodoAPI.md#GetTodoProjectsByKeyIssues) | **Get** /v1/todo/projects/{key}/issues | Returns a board&#39;s issues — work items with their column, priority, assignee, labels and schedule.
[**GetTodoProjectsByKeyIssuesByNum**](TodoAPI.md#GetTodoProjectsByKeyIssuesByNum) | **Get** /v1/todo/projects/{key}/issues/{num} | Returns ONE work item in full — its description included.
[**GetTodoRoomsByRoom**](TodoAPI.md#GetTodoRoomsByRoom) | **Get** /v1/todo/rooms/{room} | Summarises one room&#39;s work.
[**PatchTodoProjectsByKey**](TodoAPI.md#PatchTodoProjectsByKey) | **Patch** /v1/todo/projects/{key} | Refused — a board is a repository on the forge
[**PatchTodoProjectsByKeyIssuesByNum**](TodoAPI.md#PatchTodoProjectsByKeyIssuesByNum) | **Patch** /v1/todo/projects/{key}/issues/{num} | Edits a work item — rename it, rewrite it, move it to another column, or re-prioritise it.
[**PostTodoProjects**](TodoAPI.md#PostTodoProjects) | **Post** /v1/todo/projects | Refused — a board is a repository on the forge
[**PostTodoProjectsByKeyIssues**](TodoAPI.md#PostTodoProjectsByKeyIssues) | **Post** /v1/todo/projects/{key}/issues | Opens a work item on the board — an issue on that repository on the deployment&#39;s forge, filed as YOU.
[**PostTodoProjectsByKeyIssuesByNumClaim**](TodoAPI.md#PostTodoProjectsByKeyIssuesByNumClaim) | **Post** /v1/todo/projects/{key}/issues/{num}/claim | Takes an issue: it becomes yours and it moves to in_progress.



## DeleteTodoProjectsByKey

> DeleteTodoProjectsByKey(ctx, key).Execute()

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
	r, err := apiClient.TodoAPI.DeleteTodoProjectsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.DeleteTodoProjectsByKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTodoProjectsByKeyRequest struct via the builder pattern


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


## GetTodoBoard

> []IssueView GetTodoBoard(ctx).Key(key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()

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
	resp, r, err := apiClient.TodoAPI.GetTodoBoard(context.Background()).Key(key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoBoard`: []IssueView
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoBoardRequest struct via the builder pattern


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

[**[]IssueView**](IssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTodoIssues

> IssueHits GetTodoIssues(ctx).Q(q).Project(project).Status(status).Kind(kind).Repo(repo).Room(room).Source(source).Assignee(assignee).Limit(limit).Execute()

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
	room := "room_example" // string | Room keeps issues bound to one collaboration room, spelled \"<workspace>_<room>\" — the exact value GET /v1/meet/call answers with, so a channel's call and its todo list name the room the same way. This is the read a channel view runs to draw its own list; it spans every board of the org, because the work a channel is about is not confined to one board. (optional)
	source := "source_example" // string | Source keeps one origin: team, git, crm, helpdesk, cms, agent. \"git\" is how you ask for the mirrored GitHub issues specifically. (optional)
	assignee := "assignee_example" // string | Assignee keeps issues held by one person. Pass \"me\" for yourself. (optional)
	limit := int32(56) // int32 | Limit caps the answer; 0 means the default, and anything above the ceiling is clamped rather than refused — a search that errors on being too broad teaches people to guess. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodoAPI.GetTodoIssues(context.Background()).Q(q).Project(project).Status(status).Kind(kind).Repo(repo).Room(room).Source(source).Assignee(assignee).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoIssues`: IssueHits
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q matches an issue&#39;s title or description. A word from the issue, which is what someone remembers — not its number, which is what they are looking up. | 
 **project** | **string** | Project narrows to one team key; \&quot;\&quot; searches every project in the org, which is the point of this op. | 
 **status** | **string** | Status keeps one board column: backlog, todo, in_progress, done, canceled. | 
 **kind** | **string** | Kind keeps one shape: issue, pr, epic. | 
 **repo** | **string** | Repo keeps issues bound to one git repository. | 
 **room** | **string** | Room keeps issues bound to one collaboration room, spelled \&quot;&lt;workspace&gt;_&lt;room&gt;\&quot; — the exact value GET /v1/meet/call answers with, so a channel&#39;s call and its todo list name the room the same way. This is the read a channel view runs to draw its own list; it spans every board of the org, because the work a channel is about is not confined to one board. | 
 **source** | **string** | Source keeps one origin: team, git, crm, helpdesk, cms, agent. \&quot;git\&quot; is how you ask for the mirrored GitHub issues specifically. | 
 **assignee** | **string** | Assignee keeps issues held by one person. Pass \&quot;me\&quot; for yourself. | 
 **limit** | **int32** | Limit caps the answer; 0 means the default, and anything above the ceiling is clamped rather than refused — a search that errors on being too broad teaches people to guess. | 

### Return type

[**IssueHits**](IssueHits.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTodoProjects

> []TodoProject GetTodoProjects(ctx).Execute()

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
	resp, r, err := apiClient.TodoAPI.GetTodoProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoProjects`: []TodoProject
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoProjectsRequest struct via the builder pattern


### Return type

[**[]TodoProject**](TodoProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTodoProjectsByKey

> TodoProject GetTodoProjectsByKey(ctx, key).Execute()

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
	resp, r, err := apiClient.TodoAPI.GetTodoProjectsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoProjectsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoProjectsByKey`: TodoProject
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoProjectsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project&#39;s org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\&quot;ENG\&quot;, \&quot;OPS2\&quot;). Matched case-insensitively. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoProjectsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TodoProject**](TodoProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTodoProjectsByKeyIssues

> []IssueView GetTodoProjectsByKeyIssues(ctx, key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()

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
	resp, r, err := apiClient.TodoAPI.GetTodoProjectsByKeyIssues(context.Background(), key).Status(status).Kind(kind).Repo(repo).Label(label).Source(source).Scheduled(scheduled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoProjectsByKeyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoProjectsByKeyIssues`: []IssueView
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoProjectsByKeyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project whose issues to list, from the path. EMPTY means every project in the org — the global board. It is a filter like the rest of this struct rather than an address, which is what lets one op answer both \&quot;this board\&quot; and \&quot;all the work\&quot; without a second surface disagreeing with the first about what a column is. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoProjectsByKeyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. | 
 **kind** | **string** | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. | 
 **repo** | **string** | Repo keeps only issues bound to that git repository. | 
 **label** | **string** | Label keeps only issues carrying that label, compared case-insensitively.  This is how a board narrows to something SMALLER than a repository — the one mechanism for it. An estate whose apps are directories inside one repository (hanzoai/cloud carries ~140 of them) has no repository per app to address, so the app is a label: &#x60;label&#x3D;app/meet&#x60; is the meet board. Nothing is provisioned to make one exist; a board is the query. | 
 **source** | **string** | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. | 
 **scheduled** | **bool** | Scheduled keeps only issues that carry a date — a start, a due date or both. This is the timeline&#39;s slice of the board: pass scheduled&#x3D;true to get exactly the rows a gantt has somewhere to draw, instead of fetching every issue and discarding the undated ones client-side. | 

### Return type

[**[]IssueView**](IssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTodoProjectsByKeyIssuesByNum

> IssueView GetTodoProjectsByKeyIssuesByNum(ctx, key, num).Execute()

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
	num := int32(56) // int32 | Num is the issue's number on that board.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodoAPI.GetTodoProjectsByKeyIssuesByNum(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoProjectsByKeyIssuesByNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoProjectsByKeyIssuesByNum`: IssueView
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoProjectsByKeyIssuesByNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the board — the repository name, or an index board&#39;s key. | 
**num** | **int32** | Num is the issue&#39;s number on that board. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoProjectsByKeyIssuesByNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IssueView**](IssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTodoRoomsByRoom

> RoomWork GetTodoRoomsByRoom(ctx, room).Execute()

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
	room := "room_example" // string | Room is the room, spelled \"<workspace>_<room>\" — the same value GET /v1/meet/call answers with, so a channel's call and its work name the room identically. From the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodoAPI.GetTodoRoomsByRoom(context.Background(), room).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.GetTodoRoomsByRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTodoRoomsByRoom`: RoomWork
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.GetTodoRoomsByRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**room** | **string** | Room is the room, spelled \&quot;&lt;workspace&gt;_&lt;room&gt;\&quot; — the same value GET /v1/meet/call answers with, so a channel&#39;s call and its work name the room identically. From the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTodoRoomsByRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoomWork**](RoomWork.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTodoProjectsByKey

> PatchTodoProjectsByKey(ctx, key).Execute()

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
	r, err := apiClient.TodoAPI.PatchTodoProjectsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.PatchTodoProjectsByKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchTodoProjectsByKeyRequest struct via the builder pattern


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


## PatchTodoProjectsByKeyIssuesByNum

> IssueView PatchTodoProjectsByKeyIssuesByNum(ctx, key, num).IssueEdit(issueEdit).Execute()

Edits a work item — rename it, rewrite it, move it to another column, or re-prioritise it.



### Example

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
	num := int32(56) // int32 | Num is the issue number on that repository, from the path.
	issueEdit := *openapiclient.NewIssueEdit() // IssueEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodoAPI.PatchTodoProjectsByKeyIssuesByNum(context.Background(), key, num).IssueEdit(issueEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.PatchTodoProjectsByKeyIssuesByNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTodoProjectsByKeyIssuesByNum`: IssueView
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.PatchTodoProjectsByKeyIssuesByNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the board — the repository name, from the path. | 
**num** | **int32** | Num is the issue number on that repository, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTodoProjectsByKeyIssuesByNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issueEdit** | [**IssueEdit**](IssueEdit.md) |  | 

### Return type

[**IssueView**](IssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTodoProjects

> PostTodoProjects(ctx).Execute()

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
	r, err := apiClient.TodoAPI.PostTodoProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.PostTodoProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostTodoProjectsRequest struct via the builder pattern


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


## PostTodoProjectsByKeyIssues

> IssueView PostTodoProjectsByKeyIssues(ctx, key).NewIssue(newIssue).Execute()

Opens a work item on the board — an issue on that repository on the deployment's forge, filed as YOU.



### Example

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
	newIssue := *openapiclient.NewNewIssue() // NewIssue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodoAPI.PostTodoProjectsByKeyIssues(context.Background(), key).NewIssue(newIssue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.PostTodoProjectsByKeyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTodoProjectsByKeyIssues`: IssueView
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.PostTodoProjectsByKeyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the board — the repository name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTodoProjectsByKeyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newIssue** | [**NewIssue**](NewIssue.md) |  | 

### Return type

[**IssueView**](IssueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTodoProjectsByKeyIssuesByNumClaim

> IssueHit PostTodoProjectsByKeyIssuesByNumClaim(ctx, key, num).Execute()

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
	num := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodoAPI.PostTodoProjectsByKeyIssuesByNumClaim(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodoAPI.PostTodoProjectsByKeyIssuesByNumClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTodoProjectsByKeyIssuesByNumClaim`: IssueHit
	fmt.Fprintf(os.Stdout, "Response from `TodoAPI.PostTodoProjectsByKeyIssuesByNumClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**num** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTodoProjectsByKeyIssuesByNumClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IssueHit**](IssueHit.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

