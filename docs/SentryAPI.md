# \SentryAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSentryByWildcard1**](SentryAPI.md#DeleteSentryByWildcard1) | **Delete** /v1/sentry/{wildcard1} | Delete a Sentry project
[**DeleteSentryProjectsById**](SentryAPI.md#DeleteSentryProjectsById) | **Delete** /v1/sentry/projects/{id} | Deletes one Sentry project of the caller&#39;s org.
[**GetSentryByWildcard1**](SentryAPI.md#GetSentryByWildcard1) | **Get** /v1/sentry/{wildcard1} | Read the caller org&#39;s errors on the Sentry surface
[**GetSentryEventsById**](SentryAPI.md#GetSentryEventsById) | **Get** /v1/sentry/events/{id} | Returns one captured error event of a project, by its id.
[**GetSentryIssues**](SentryAPI.md#GetSentryIssues) | **Get** /v1/sentry/issues | Lists the caller&#39;s org&#39;s grouped error issues, optionally narrowed to one project and one time window, and filtered by status, level, environment, service, a free-text query and a sort.
[**GetSentryIssuesById**](SentryAPI.md#GetSentryIssuesById) | **Get** /v1/sentry/issues/{id} | Returns one grouped issue of the caller&#39;s org with its latest occurrence sample.
[**GetSentryIssuesByIdEvents**](SentryAPI.md#GetSentryIssuesByIdEvents) | **Get** /v1/sentry/issues/{id}/events | Lists one issue&#39;s captured occurrences, scoped to a project — a project is an isolation unit, so the caller declares which project&#39;s occurrences to read.
[**GetSentryLogs**](SentryAPI.md#GetSentryLogs) | **Get** /v1/sentry/logs | Lists a project&#39;s captured error events, newest first, optionally narrowed to those whose message or exception text contains a search string.
[**GetSentryProjects**](SentryAPI.md#GetSentryProjects) | **Get** /v1/sentry/projects | Lists the caller&#39;s org&#39;s Sentry projects, each with its freshly-derived DSN.
[**GetSentryProjectsById**](SentryAPI.md#GetSentryProjectsById) | **Get** /v1/sentry/projects/{id} | Returns one Sentry project of the caller&#39;s org, DSN included.
[**GetSentryStats**](SentryAPI.md#GetSentryStats) | **Get** /v1/sentry/stats | Returns a project&#39;s event-rate timeseries: one bucket per interval over the requested period, counting the events in it.
[**GetSentryTraces**](SentryAPI.md#GetSentryTraces) | **Get** /v1/sentry/traces | Lists the traces a project&#39;s captured errors reference, each with how many errors landed on it, when they started and stopped, and the latest message seen — the entry point for \&quot;which requests are failing\&quot;.
[**GetSentryTracesById**](SentryAPI.md#GetSentryTracesById) | **Get** /v1/sentry/traces/{id} | Returns one trace&#39;s captured errors for a project — every error event that carried the trace id, in the order the events plane holds them.
[**PatchSentryByWildcard1**](SentryAPI.md#PatchSentryByWildcard1) | **Patch** /v1/sentry/{wildcard1} | Not served — the Sentry surface has no partial update
[**PostSentryByProjectEnvelope**](SentryAPI.md#PostSentryByProjectEnvelope) | **Post** /v1/sentry/{project}/envelope/ | Receive a Sentry envelope on the clean root
[**PostSentryByProjectStore**](SentryAPI.md#PostSentryByProjectStore) | **Post** /v1/sentry/{project}/store/ | Receive a single Sentry event on the clean root
[**PostSentryByWildcard1**](SentryAPI.md#PostSentryByWildcard1) | **Post** /v1/sentry/{wildcard1} | Send events to the Sentry surface, or write on it
[**PostSentryDiscover**](SentryAPI.md#PostSentryDiscover) | **Post** /v1/sentry/discover | Aggregates a project&#39;s captured errors into a table — the caller names the filters, the groupings and the aggregations, and gets back the columns and rows they asked for.
[**PostSentryProjects**](SentryAPI.md#PostSentryProjects) | **Post** /v1/sentry/projects | Creates a Sentry project under the caller&#39;s org and returns it, DSN included.
[**PostSentryProjectsByIdKeysRotate**](SentryAPI.md#PostSentryProjectsByIdKeysRotate) | **Post** /v1/sentry/projects/{id}/keys/rotate | Rotates a project&#39;s DSN key — bumping its rotation watermark so keys below it stop verifying — and returns the project with its new DSN.
[**PutSentryByWildcard1**](SentryAPI.md#PutSentryByWildcard1) | **Put** /v1/sentry/{wildcard1} | Move an error issue through its lifecycle
[**PutSentryIssuesById**](SentryAPI.md#PutSentryIssuesById) | **Put** /v1/sentry/issues/{id} | Changes an issue&#39;s lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.



## DeleteSentryByWildcard1

> DeleteSentryByWildcard1(ctx, wildcard1).Execute()

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
	r, err := apiClient.SentryAPI.DeleteSentryByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.DeleteSentryByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSentryByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSentryProjectsById

> DeleteSentryProjectsById(ctx, id).Execute()

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
	r, err := apiClient.SentryAPI.DeleteSentryProjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.DeleteSentryProjectsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSentryProjectsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryByWildcard1

> GetSentryByWildcard1(ctx, wildcard1).Execute()

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
	r, err := apiClient.SentryAPI.GetSentryByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSentryByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryEventsById

> O11yO11ySentryEventOut GetSentryEventsById(ctx, id).Project(project).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryEventsById(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryEventsById`: O11yO11ySentryEventOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the event id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project the event belongs to, by its id. Required. | 

### Return type

[**O11yO11ySentryEventOut**](O11yO11ySentryEventOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryIssues

> O11yO11yErrorIssuesOut GetSentryIssues(ctx).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Project(project).Period(period).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryIssues(context.Background()).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Project(project).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryIssues`: O11yO11yErrorIssuesOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryIssuesRequest struct via the builder pattern


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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryIssuesById

> O11yO11yErrorGettableIssueOut GetSentryIssuesById(ctx, id).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryIssuesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryIssuesById`: O11yO11yErrorGettableIssueOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yErrorGettableIssueOut**](O11yO11yErrorGettableIssueOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryIssuesByIdEvents

> O11yO11ySentryIssueEventsOut GetSentryIssuesByIdEvents(ctx, id).Project(project).Limit(limit).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryIssuesByIdEvents(context.Background(), id).Project(project).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryIssuesByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryIssuesByIdEvents`: O11yO11ySentryIssueEventsOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryIssuesByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryIssuesByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project whose occurrences to read, by its id. Required. | 
 **limit** | **int32** | Limit caps how many occurrences come back. Zero means the default. | 

### Return type

[**O11yO11ySentryIssueEventsOut**](O11yO11ySentryIssueEventsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryLogs

> O11yO11yLogsOut GetSentryLogs(ctx).Project(project).Query(query).Period(period).Limit(limit).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryLogs(context.Background()).Project(project).Query(query).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryLogs`: O11yO11yLogsOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **query** | **string** | Query narrows the page to events whose text contains it. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 
 **limit** | **int32** | Limit caps how many events come back. | 

### Return type

[**O11yO11yLogsOut**](O11yO11yLogsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryProjects

> O11yO11ySentryProjectsOut GetSentryProjects(ctx).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryProjects`: O11yO11ySentryProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryProjectsRequest struct via the builder pattern


### Return type

[**O11yO11ySentryProjectsOut**](O11yO11ySentryProjectsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryProjectsById

> O11yO11ySentryProjectOut GetSentryProjectsById(ctx, id).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryProjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryProjectsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryProjectsById`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryProjectsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryProjectsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryStats

> O11yO11yStatsOut GetSentryStats(ctx).Project(project).Field(field).Period(period).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryStats(context.Background()).Project(project).Field(field).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryStats`: O11yO11yStatsOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **field** | **string** | Field is the dimension to count over. Empty counts all events. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 

### Return type

[**O11yO11yStatsOut**](O11yO11yStatsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryTraces

> O11yO11yTracesOut GetSentryTraces(ctx).Project(project).Period(period).Limit(limit).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryTraces(context.Background()).Project(project).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryTraces`: O11yO11yTracesOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 
 **limit** | **int32** | Limit caps how many traces come back. | 

### Return type

[**O11yO11yTracesOut**](O11yO11yTracesOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentryTracesById

> O11yO11yTraceOut GetSentryTracesById(ctx, id).Project(project).Execute()

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
	resp, r, err := apiClient.SentryAPI.GetSentryTracesById(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.GetSentryTracesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentryTracesById`: O11yO11yTraceOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.GetSentryTracesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the trace id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSentryTracesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project the trace&#39;s errors belong to. Required. | 

### Return type

[**O11yO11yTraceOut**](O11yO11yTraceOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSentryByWildcard1

> PatchSentryByWildcard1(ctx, wildcard1).Execute()

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
	r, err := apiClient.SentryAPI.PatchSentryByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PatchSentryByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchSentryByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentryByProjectEnvelope

> PostSentryByProjectEnvelope(ctx, project).Execute()

Receive a Sentry envelope on the clean root



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
	r, err := apiClient.SentryAPI.PostSentryByProjectEnvelope(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PostSentryByProjectEnvelope``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostSentryByProjectEnvelopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentryByProjectStore

> PostSentryByProjectStore(ctx, project).Execute()

Receive a single Sentry event on the clean root



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
	r, err := apiClient.SentryAPI.PostSentryByProjectStore(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PostSentryByProjectStore``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostSentryByProjectStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentryByWildcard1

> PostSentryByWildcard1(ctx, wildcard1).Execute()

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
	r, err := apiClient.SentryAPI.PostSentryByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PostSentryByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostSentryByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentryDiscover

> O11yO11yDiscoverOut PostSentryDiscover(ctx).O11yO11yDiscoverIn(o11yO11yDiscoverIn).Execute()

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
	resp, r, err := apiClient.SentryAPI.PostSentryDiscover(context.Background()).O11yO11yDiscoverIn(o11yO11yDiscoverIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PostSentryDiscover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSentryDiscover`: O11yO11yDiscoverOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.PostSentryDiscover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSentryDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDiscoverIn** | [**O11yO11yDiscoverIn**](O11yO11yDiscoverIn.md) |  | 

### Return type

[**O11yO11yDiscoverOut**](O11yO11yDiscoverOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentryProjects

> O11yO11ySentryProjectOut PostSentryProjects(ctx).O11yO11ySentryPostableProject(o11yO11ySentryPostableProject).Execute()

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
	resp, r, err := apiClient.SentryAPI.PostSentryProjects(context.Background()).O11yO11ySentryPostableProject(o11yO11ySentryPostableProject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PostSentryProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSentryProjects`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.PostSentryProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSentryProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11ySentryPostableProject** | [**O11yO11ySentryPostableProject**](O11yO11ySentryPostableProject.md) |  | 

### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSentryProjectsByIdKeysRotate

> O11yO11ySentryProjectOut PostSentryProjectsByIdKeysRotate(ctx, id).Execute()

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
	resp, r, err := apiClient.SentryAPI.PostSentryProjectsByIdKeysRotate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PostSentryProjectsByIdKeysRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSentryProjectsByIdKeysRotate`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.PostSentryProjectsByIdKeysRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSentryProjectsByIdKeysRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSentryByWildcard1

> PutSentryByWildcard1(ctx, wildcard1).Execute()

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
	r, err := apiClient.SentryAPI.PutSentryByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PutSentryByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutSentryByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSentryIssuesById

> O11yO11yErrorIssueOut PutSentryIssuesById(ctx, id).O11yO11ySentryUpdateIssueIn(o11yO11ySentryUpdateIssueIn).Execute()

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
	resp, r, err := apiClient.SentryAPI.PutSentryIssuesById(context.Background(), id).O11yO11ySentryUpdateIssueIn(o11yO11ySentryUpdateIssueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentryAPI.PutSentryIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSentryIssuesById`: O11yO11yErrorIssueOut
	fmt.Fprintf(os.Stdout, "Response from `SentryAPI.PutSentryIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSentryIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySentryUpdateIssueIn** | [**O11yO11ySentryUpdateIssueIn**](O11yO11ySentryUpdateIssueIn.md) |  | 

### Return type

[**O11yO11yErrorIssueOut**](O11yO11yErrorIssueOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

