# \SentinelAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSentinelByWildcard1**](SentinelAPI.md#DeleteSentinelByWildcard1) | **Delete** /v1/sentinel/{wildcard1} | Delete a Sentry project
[**DeleteSentinelProjectsById**](SentinelAPI.md#DeleteSentinelProjectsById) | **Delete** /v1/sentinel/projects/{id} | Deletes one Sentry project of the caller&#39;s org.
[**GetSentinelByWildcard1**](SentinelAPI.md#GetSentinelByWildcard1) | **Get** /v1/sentinel/{wildcard1} | Read the caller org&#39;s errors on the Sentry surface
[**GetSentinelEventsById**](SentinelAPI.md#GetSentinelEventsById) | **Get** /v1/sentinel/events/{id} | Returns one captured error event of a project, by its id.
[**GetSentinelIssues**](SentinelAPI.md#GetSentinelIssues) | **Get** /v1/sentinel/issues | Lists the caller&#39;s org&#39;s grouped error issues, optionally narrowed to one project and one time window, and filtered by status, level, environment, service, a free-text query and a sort.
[**GetSentinelIssuesById**](SentinelAPI.md#GetSentinelIssuesById) | **Get** /v1/sentinel/issues/{id} | Returns one grouped issue of the caller&#39;s org with its latest occurrence sample.
[**GetSentinelIssuesByIdEvents**](SentinelAPI.md#GetSentinelIssuesByIdEvents) | **Get** /v1/sentinel/issues/{id}/events | Lists one issue&#39;s captured occurrences, scoped to a project — a project is an isolation unit, so the caller declares which project&#39;s occurrences to read.
[**GetSentinelLogs**](SentinelAPI.md#GetSentinelLogs) | **Get** /v1/sentinel/logs | Lists a project&#39;s captured error events, newest first, optionally narrowed to those whose message or exception text contains a search string.
[**GetSentinelProjects**](SentinelAPI.md#GetSentinelProjects) | **Get** /v1/sentinel/projects | Lists the caller&#39;s org&#39;s Sentry projects, each with its freshly-derived DSN.
[**GetSentinelProjectsById**](SentinelAPI.md#GetSentinelProjectsById) | **Get** /v1/sentinel/projects/{id} | Returns one Sentry project of the caller&#39;s org, DSN included.
[**GetSentinelStats**](SentinelAPI.md#GetSentinelStats) | **Get** /v1/sentinel/stats | Returns a project&#39;s event-rate timeseries: one bucket per interval over the requested period, counting the events in it.
[**GetSentinelTraces**](SentinelAPI.md#GetSentinelTraces) | **Get** /v1/sentinel/traces | Lists the traces a project&#39;s captured errors reference, each with how many errors landed on it, when they started and stopped, and the latest message seen — the entry point for \&quot;which requests are failing\&quot;.
[**GetSentinelTracesById**](SentinelAPI.md#GetSentinelTracesById) | **Get** /v1/sentinel/traces/{id} | Returns one trace&#39;s captured errors for a project — every error event that carried the trace id, in the order the events plane holds them.
[**PatchSentinelByWildcard1**](SentinelAPI.md#PatchSentinelByWildcard1) | **Patch** /v1/sentinel/{wildcard1} | Not served — the Sentry surface has no partial update
[**PostSentinelByProjectEnvelope**](SentinelAPI.md#PostSentinelByProjectEnvelope) | **Post** /v1/sentinel/{project}/envelope/ | Receive a Sentry envelope on the runtime&#39;s ingest hatch
[**PostSentinelByProjectStore**](SentinelAPI.md#PostSentinelByProjectStore) | **Post** /v1/sentinel/{project}/store/ | Receive a single event on the runtime&#39;s ingest hatch
[**PostSentinelByWildcard1**](SentinelAPI.md#PostSentinelByWildcard1) | **Post** /v1/sentinel/{wildcard1} | Send events to the Sentry surface, or write on it
[**PostSentinelDiscover**](SentinelAPI.md#PostSentinelDiscover) | **Post** /v1/sentinel/discover | Aggregates a project&#39;s captured errors into a table — the caller names the filters, the groupings and the aggregations, and gets back the columns and rows they asked for.
[**PostSentinelProjects**](SentinelAPI.md#PostSentinelProjects) | **Post** /v1/sentinel/projects | Creates a Sentry project under the caller&#39;s org and returns it, DSN included.
[**PostSentinelProjectsByIdKeysRotate**](SentinelAPI.md#PostSentinelProjectsByIdKeysRotate) | **Post** /v1/sentinel/projects/{id}/keys/rotate | Rotates a project&#39;s DSN key — bumping its rotation watermark so keys below it stop verifying — and returns the project with its new DSN.
[**PutSentinelByWildcard1**](SentinelAPI.md#PutSentinelByWildcard1) | **Put** /v1/sentinel/{wildcard1} | Move an error issue through its lifecycle
[**PutSentinelIssuesById**](SentinelAPI.md#PutSentinelIssuesById) | **Put** /v1/sentinel/issues/{id} | Changes an issue&#39;s lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.



## DeleteSentinelByWildcard1

> DeleteSentinelByWildcard1(ctx, wildcard1).Execute()

Delete a Sentry project



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.DeleteSentinelByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.DeleteSentinelByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSentinelByWildcard1Request struct via the builder pattern


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


## DeleteSentinelProjectsById

> DeleteSentinelProjectsById(ctx, id).Execute()

Deletes one Sentry project of the caller's org.



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
	id := "id_example" // string | ID is the project id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.DeleteSentinelProjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.DeleteSentinelProjectsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSentinelProjectsByIdRequest struct via the builder pattern


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


## GetSentinelByWildcard1

> GetSentinelByWildcard1(ctx, wildcard1).Execute()

Read the caller org's errors on the Sentry surface



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.GetSentinelByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelByWildcard1Request struct via the builder pattern


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


## GetSentinelEventsById

> O11yO11ySentryEventOut GetSentinelEventsById(ctx, id).Project(project).Execute()

Returns one captured error event of a project, by its id.



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
	id := "id_example" // string | ID is the event id.
	project := "project_example" // string | Project is the project the event belongs to, by its id. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelEventsById(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelEventsById`: O11yO11ySentryEventOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the event id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project the event belongs to, by its id. Required. | 

### Return type

[**O11yO11ySentryEventOut**](O11yO11ySentryEventOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelIssues

> O11yO11yErrorIssuesOut GetSentinelIssues(ctx).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Project(project).Period(period).Execute()

Lists the caller's org's grouped error issues, optionally narrowed to one project and one time window, and filtered by status, level, environment, service, a free-text query and a sort.



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
	status := "status_example" // string | Status narrows to one lifecycle state: unresolved, resolved or ignored. (optional)
	level := "level_example" // string | Level narrows to one severity, e.g. error, warning, info. (optional)
	environment := "environment_example" // string | Environment narrows to one deployment environment. (optional)
	serviceName := "serviceName_example" // string | ServiceName narrows to one reporting service. (optional)
	query := "query_example" // string | Query narrows to issues whose text contains it. (optional)
	sort := "sort_example" // string | Sort orders the page, e.g. lastSeen, firstSeen, count. (optional)
	offset := int32(56) // int32 | Offset is how many issues to skip. Zero starts at the first. (optional)
	limit := int32(56) // int32 | Limit caps how many issues come back. Zero means the default. (optional)
	project := "project_example" // string | Project narrows the org's issues to one project, by its id. (optional)
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelIssues(context.Background()).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Project(project).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelIssues`: O11yO11yErrorIssuesOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status narrows to one lifecycle state: unresolved, resolved or ignored. | 
 **level** | **string** | Level narrows to one severity, e.g. error, warning, info. | 
 **environment** | **string** | Environment narrows to one deployment environment. | 
 **serviceName** | **string** | ServiceName narrows to one reporting service. | 
 **query** | **string** | Query narrows to issues whose text contains it. | 
 **sort** | **string** | Sort orders the page, e.g. lastSeen, firstSeen, count. | 
 **offset** | **int32** | Offset is how many issues to skip. Zero starts at the first. | 
 **limit** | **int32** | Limit caps how many issues come back. Zero means the default. | 
 **project** | **string** | Project narrows the org&#39;s issues to one project, by its id. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 

### Return type

[**O11yO11yErrorIssuesOut**](O11yO11yErrorIssuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelIssuesById

> O11yO11yErrorGettableIssueOut GetSentinelIssuesById(ctx, id).Execute()

Returns one grouped issue of the caller's org with its latest occurrence sample.



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
	id := "id_example" // string | ID is the issue id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelIssuesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelIssuesById`: O11yO11yErrorGettableIssueOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yErrorGettableIssueOut**](O11yO11yErrorGettableIssueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelIssuesByIdEvents

> O11yO11ySentryIssueEventsOut GetSentinelIssuesByIdEvents(ctx, id).Project(project).Limit(limit).Execute()

Lists one issue's captured occurrences, scoped to a project — a project is an isolation unit, so the caller declares which project's occurrences to read.



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
	id := "id_example" // string | ID is the issue id.
	project := "project_example" // string | Project is the project whose occurrences to read, by its id. Required.
	limit := int32(56) // int32 | Limit caps how many occurrences come back. Zero means the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelIssuesByIdEvents(context.Background(), id).Project(project).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelIssuesByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelIssuesByIdEvents`: O11yO11ySentryIssueEventsOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelIssuesByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelIssuesByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project whose occurrences to read, by its id. Required. | 
 **limit** | **int32** | Limit caps how many occurrences come back. Zero means the default. | 

### Return type

[**O11yO11ySentryIssueEventsOut**](O11yO11ySentryIssueEventsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelLogs

> O11yO11yLogsOut GetSentinelLogs(ctx).Project(project).Query(query).Period(period).Limit(limit).Execute()

Lists a project's captured error events, newest first, optionally narrowed to those whose message or exception text contains a search string.



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
	project := "project_example" // string | Project is the project to read, as its id. Required.
	query := "query_example" // string | Query narrows the page to events whose text contains it. (optional)
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)
	limit := int32(56) // int32 | Limit caps how many events come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelLogs(context.Background()).Project(project).Query(query).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelLogs`: O11yO11yLogsOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **query** | **string** | Query narrows the page to events whose text contains it. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 
 **limit** | **int32** | Limit caps how many events come back. | 

### Return type

[**O11yO11yLogsOut**](O11yO11yLogsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelProjects

> O11yO11ySentryProjectsOut GetSentinelProjects(ctx).Execute()

Lists the caller's org's Sentry projects, each with its freshly-derived DSN.



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
	resp, r, err := apiClient.SentinelAPI.GetSentinelProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelProjects`: O11yO11ySentryProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelProjectsRequest struct via the builder pattern


### Return type

[**O11yO11ySentryProjectsOut**](O11yO11ySentryProjectsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelProjectsById

> O11yO11ySentryProjectOut GetSentinelProjectsById(ctx, id).Execute()

Returns one Sentry project of the caller's org, DSN included.



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
	id := "id_example" // string | ID is the project id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelProjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelProjectsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelProjectsById`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelProjectsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelProjectsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelStats

> O11yO11yStatsOut GetSentinelStats(ctx).Project(project).Field(field).Period(period).Execute()

Returns a project's event-rate timeseries: one bucket per interval over the requested period, counting the events in it.



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
	project := "project_example" // string | Project is the project to read, as its id. Required.
	field := "field_example" // string | Field is the dimension to count over. Empty counts all events. (optional)
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelStats(context.Background()).Project(project).Field(field).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelStats`: O11yO11yStatsOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **field** | **string** | Field is the dimension to count over. Empty counts all events. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 

### Return type

[**O11yO11yStatsOut**](O11yO11yStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelTraces

> O11yO11yTracesOut GetSentinelTraces(ctx).Project(project).Period(period).Limit(limit).Execute()

Lists the traces a project's captured errors reference, each with how many errors landed on it, when they started and stopped, and the latest message seen — the entry point for \"which requests are failing\".



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
	project := "project_example" // string | Project is the project to read, as its id. Required.
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)
	limit := int32(56) // int32 | Limit caps how many traces come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelTraces(context.Background()).Project(project).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelTraces`: O11yO11yTracesOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 
 **limit** | **int32** | Limit caps how many traces come back. | 

### Return type

[**O11yO11yTracesOut**](O11yO11yTracesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentinelTracesById

> O11yO11yTraceOut GetSentinelTracesById(ctx, id).Project(project).Execute()

Returns one trace's captured errors for a project — every error event that carried the trace id, in the order the events plane holds them.



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
	id := "id_example" // string | ID is the trace id.
	project := "project_example" // string | Project is the project the trace's errors belong to. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.GetSentinelTracesById(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.GetSentinelTracesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentinelTracesById`: O11yO11yTraceOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.GetSentinelTracesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the trace id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentinelTracesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project the trace&#39;s errors belong to. Required. | 

### Return type

[**O11yO11yTraceOut**](O11yO11yTraceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSentinelByWildcard1

> PatchSentinelByWildcard1(ctx, wildcard1).Execute()

Not served — the Sentry surface has no partial update



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.PatchSentinelByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PatchSentinelByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSentinelByWildcard1Request struct via the builder pattern


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


## PostSentinelByProjectEnvelope

> PostSentinelByProjectEnvelope(ctx, project).Execute()

Receive a Sentry envelope on the runtime's ingest hatch



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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.PostSentinelByProjectEnvelope(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PostSentinelByProjectEnvelope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSentinelByProjectEnvelopeRequest struct via the builder pattern


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


## PostSentinelByProjectStore

> PostSentinelByProjectStore(ctx, project).Execute()

Receive a single event on the runtime's ingest hatch



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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.PostSentinelByProjectStore(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PostSentinelByProjectStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSentinelByProjectStoreRequest struct via the builder pattern


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


## PostSentinelByWildcard1

> PostSentinelByWildcard1(ctx, wildcard1).Execute()

Send events to the Sentry surface, or write on it



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.PostSentinelByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PostSentinelByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSentinelByWildcard1Request struct via the builder pattern


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


## PostSentinelDiscover

> O11yO11yDiscoverOut PostSentinelDiscover(ctx).O11yO11yDiscoverIn(o11yO11yDiscoverIn).Execute()

Aggregates a project's captured errors into a table — the caller names the filters, the groupings and the aggregations, and gets back the columns and rows they asked for.



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
	o11yO11yDiscoverIn := *openapiclient.NewO11yO11yDiscoverIn("Project_example") // O11yO11yDiscoverIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.PostSentinelDiscover(context.Background()).O11yO11yDiscoverIn(o11yO11yDiscoverIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PostSentinelDiscover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSentinelDiscover`: O11yO11yDiscoverOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.PostSentinelDiscover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSentinelDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDiscoverIn** | [**O11yO11yDiscoverIn**](O11yO11yDiscoverIn.md) |  | 

### Return type

[**O11yO11yDiscoverOut**](O11yO11yDiscoverOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentinelProjects

> O11yO11ySentryProjectOut PostSentinelProjects(ctx).O11yO11ySentryPostableProject(o11yO11ySentryPostableProject).Execute()

Creates a Sentry project under the caller's org and returns it, DSN included.



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
	o11yO11ySentryPostableProject := *openapiclient.NewO11yO11ySentryPostableProject("Name_example") // O11yO11ySentryPostableProject | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.PostSentinelProjects(context.Background()).O11yO11ySentryPostableProject(o11yO11ySentryPostableProject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PostSentinelProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSentinelProjects`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.PostSentinelProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSentinelProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11ySentryPostableProject** | [**O11yO11ySentryPostableProject**](O11yO11ySentryPostableProject.md) |  | 

### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentinelProjectsByIdKeysRotate

> O11yO11ySentryProjectOut PostSentinelProjectsByIdKeysRotate(ctx, id).Execute()

Rotates a project's DSN key — bumping its rotation watermark so keys below it stop verifying — and returns the project with its new DSN.



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
	id := "id_example" // string | ID is the project id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.PostSentinelProjectsByIdKeysRotate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PostSentinelProjectsByIdKeysRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSentinelProjectsByIdKeysRotate`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.PostSentinelProjectsByIdKeysRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSentinelProjectsByIdKeysRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSentinelByWildcard1

> PutSentinelByWildcard1(ctx, wildcard1).Execute()

Move an error issue through its lifecycle



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SentinelAPI.PutSentinelByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PutSentinelByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSentinelByWildcard1Request struct via the builder pattern


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


## PutSentinelIssuesById

> O11yO11yErrorIssueOut PutSentinelIssuesById(ctx, id).O11yO11ySentryUpdateIssueIn(o11yO11ySentryUpdateIssueIn).Execute()

Changes an issue's lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.



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
	id := "id_example" // string | ID is the issue id.
	o11yO11ySentryUpdateIssueIn := *openapiclient.NewO11yO11ySentryUpdateIssueIn("Id_example") // O11yO11ySentryUpdateIssueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentinelAPI.PutSentinelIssuesById(context.Background(), id).O11yO11ySentryUpdateIssueIn(o11yO11ySentryUpdateIssueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentinelAPI.PutSentinelIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSentinelIssuesById`: O11yO11yErrorIssueOut
	fmt.Fprintf(os.Stdout, "Response from `SentinelAPI.PutSentinelIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSentinelIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySentryUpdateIssueIn** | [**O11yO11ySentryUpdateIssueIn**](O11yO11ySentryUpdateIssueIn.md) |  | 

### Return type

[**O11yO11yErrorIssueOut**](O11yO11yErrorIssueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

